package game

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/meagar/sorts"
)

type bar struct {
	n      int
	height float64
	img    *ebiten.Image
}

type Alg struct {
	// The name of the algorithm
	Name string

	// The actual sorting function from the sorts package
	Fn func(sorts.SortInput)

	// How many bars to sort through
	DefaultN int

	// How many frames to "sleep" each swap
	Sleep int
}

type Game struct {
	Algs []*Alg

	screenWidth, screenHeight int
	bars                      []*bar
	barWidth                  float64

	// The sorting algorithm waits for us to signal to proceed on this channel
	tick chan struct{}

	// The sorting algorithm tells us it's safe to draw the screen on this channel
	tock chan struct{}

	// Used to delay transitions
	pause int

	// Which algorithm case we're displaying
	active int
	alg    *Alg

	// Sorting state
	lastSwap [2]int    // the last two indexes that were swapped (to draw colored bars)
	lastAt   []int     // the elements that were accessed during the last update
	done     chan bool // whether the algorithm has reported completion
	finished bool

	swaps   int // the total number of times the algorithm swapped two elements
	lookups int // the total number of times the array was indexed into
	cmps    int // the total number of comparisons performed (>, <, >=, <=, >, etc):w
}

var _ ebiten.Game = &Game{}

func (g *Game) init() {
	g.tick = make(chan struct{})
	g.tock = make(chan struct{})
	g.nextAlg()
}

func (g *Game) nextAlg() {
	alg := g.Algs[g.active]
	g.active++
	if g.active >= len(g.Algs) {
		g.active = 0
	}

	ints := shuffle(alg.DefaultN)
	bars := make([]*bar, len(ints))

	// Create our bars, scaled to the (screenWidth/numBars) width and (n/max * screenHeight) height
	g.barWidth = float64(g.screenWidth)/float64(len(ints)) - 1
	max := float64(max(ints))
	for i, n := range ints {
		h := (float64(n)/max)*float64(g.screenHeight) + 1
		bars[i] = &bar{
			n:      n,
			height: h,
			img:    ebiten.NewImage(int(g.barWidth), int(h)),
		}

		bars[i].img.Fill(color.White)
	}

	// Reset state
	g.swaps = 0
	g.lookups = 0
	g.bars = bars
	g.alg = alg
	g.finished = false
	go alg.Fn(g)
}

func max(ints []int) int {
	max := 0
	for _, n := range ints {
		if n > max {
			max = n
		}
	}
	return max
}

func (g *Game) Layout(outerWidth, outerHeight int) (int, int) {
	g.screenWidth = outerWidth
	g.screenHeight = outerHeight

	return outerWidth, outerHeight
}

var n = 0

func (g *Game) Update() error {
	if g.alg == nil {
		g.init()
		return nil
	}

	// n++
	// if n > 1 {
	// 	n = 0
	// } else {
	// 	return nil
	// }

	if g.finished {
		g.pause -= 1
		if g.pause <= 0 {
			g.nextAlg()
		}
	} else {
		g.tick <- struct{}{}
		<-g.tock
	}

	return nil
}

func (g *Game) lastAtIncludes(idx int) bool {
	for _, at := range g.lastAt {
		if at == idx {
			return true
		}
	}
	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	for i, b := range g.bars {
		op.GeoM.Reset()
		op.ColorM.Reset()
		if i == g.lastSwap[0] {
			op.ColorM.Scale(1, 0, 0, 1)
		} else if i == g.lastSwap[1] {
			op.ColorM.Scale(0, 0, 1, 1)
		} else if g.lastAtIncludes(i) {
			op.ColorM.Scale(0, 1, 0, 1)
		}
		op.GeoM.Translate((g.barWidth+1)*float64(i), math.Ceil(float64(g.screenHeight)-b.height))
		screen.DrawImage(b.img, &op)
	}
	debugMsg := fmt.Sprintf("%s %d elements\nSwaps(red/blue): %d\nLookups (green): %d\nFPS: %0.2f", g.alg.Name, len(g.bars), g.swaps, g.lookups, ebiten.CurrentFPS())
	if g.finished {
		debugMsg += "\nDone"
	}
	// g.lastAt = []int{}
	ebitenutil.DebugPrint(screen, debugMsg)
}

//
// sorts.SortInput implementation
//

var _ sorts.SortInput = &Game{}

func (g *Game) At(idx int) int {
	<-g.tick
	g.lookups++
	g.lastAt = append(g.lastAt, idx)
	g.tock <- struct{}{}
	return g.bars[idx].n
}

func (g *Game) ResetIteration() {
	g.lastAt = []int{}
}

func (g *Game) Swap(a, b int) {
	// Block until the update loop tells us to proceed
	<-g.tick

	g.lastSwap[0] = a
	g.lastSwap[1] = b

	g.bars[a], g.bars[b] = g.bars[b], g.bars[a]
	g.swaps++

	// Unblock the update loop
	g.tock <- struct{}{}
}

func (g *Game) Length() int {
	return len(g.bars)
}

func (g *Game) Done() {
	g.ResetIteration()
	g.finished = true
	// 60 ticks per second -> 3 second pause
	g.pause = 60 * 3
	g.lastSwap[0] = -1
	g.lastSwap[1] = -1
}

func shuffle(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	// For each index, swap its value with a random index
	for i := 0; i < n; i++ {
		src := rand.Intn(n)
		arr[i], arr[src] = arr[src], arr[i]
	}

	return arr
}
