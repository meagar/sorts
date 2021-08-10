package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

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
	tick                      chan (struct{})

	// Which algorithm case we're displaying
	active int
	alg    *Alg

	// Sorting state
	lastSwap [2]int // the last two indexes that were swapped (to draw colored bars)
	done     bool   // whether the algorithm has reported completion
	swaps    int    // the total number of times the algorithm swapped two elements
}

var _ ebiten.Game = &Game{}

func (g *Game) init() {
	g.nextAlg()
}

func (g *Game) nextAlg() {
	g.alg = nil
	alg := g.Algs[g.active]
	g.active++
	if g.active >= len(g.Algs) {
		g.active = 0
	}

	g.swaps = 0
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

	g.alg = alg
	g.bars = bars
	g.done = false
	g.tick = make(chan struct{})
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

var sleep = 0
var ticks = 0

func (g *Game) Update() error {
	if len(g.bars) == 0 {
		g.init()
	}

	if g.alg == nil {
		return nil
	}

	if !g.done {
		// ticks++
		// if ticks > g.alg().Sleep {
		// ticks = 0
		g.tick <- struct{}{}
		// }
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.alg == nil {
		return
	}
	op := ebiten.DrawImageOptions{}
	for i, b := range g.bars {
		op.GeoM.Reset()
		op.GeoM.Translate((g.barWidth+1)*float64(i), float64(g.screenHeight)-b.height)

		// Ideally, we'd just have one blue sprite and one red sprite, rather than filling/refilling constantly
		// if i == g.lastSwap[0] {
		// 	b.img.Fill(color.RGBA{255, 0, 0, 255})
		// } else if i == g.lastSwap[1] {
		// 	b.img.Fill(color.RGBA{0, 0, 255, 255})
		// } else {
		// 	b.img.Fill(color.White)
		// }

		screen.DrawImage(b.img, &op)
	}
	debugMsg := fmt.Sprintf("%s %d elements: %d swaps\nFPS: %0.2f", g.alg.Name, len(g.bars), g.swaps, ebiten.CurrentFPS())
	if g.done {
		debugMsg += "\nDone"
	}
	ebitenutil.DebugPrint(screen, debugMsg)
}

//
// sorts.SortInput implementation
//

var _ sorts.SortInput = &Game{}

func (g *Game) At(idx int) int {
	return g.bars[idx].n
}

func (g *Game) Swap(a, b int) {
	// Block until the game tells us to proceed
	<-g.tick

	g.lastSwap = [2]int{a, b}
	g.bars[a], g.bars[b] = g.bars[b], g.bars[a]
	g.swaps++
}

func (g *Game) CmpGE(a, b int) bool {
	return a >= b
}
func (g *Game) CmpLE(a, b int) bool {
	return a <= b
}
func (g *Game) Length() int {
	return len(g.bars)
}

func (g *Game) Done() {
	g.done = true
	g.lastSwap = [2]int{-1, -1}
	close(g.tick)

	select {
	case <-time.After(3 * time.Second):
		g.nextAlg()
	}
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
