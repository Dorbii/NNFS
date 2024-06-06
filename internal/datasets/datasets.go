package datasets

import (
	"image/color"
	"math"
	"math/rand/v2"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

func PlotData(plotType string) {
	switch plotType {
	case "spiral":
		p := plot.New()
		dataX, dataY := SpiralData(100, 3)
		cmapBRG := map[float64]color.Color{
			0: color.RGBA{R: 0, G: 0, B: 255, A: 255},
			1: color.RGBA{R: 255, G: 0, B: 0, A: 255},
			2: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		}
		for k, v := range cmapBRG {
			xys := plotter.XYs{}
			for i := range dataX {
				if dataY[i] == k {
					xys = append(xys, plotter.XY{X: dataX[i][0], Y: dataX[i][1]})
				}
			}
			scatter, err := plotter.NewScatter(xys)
			if err != nil {
				panic(err)
			}
			scatter.GlyphStyle.Color = v
			scatter.GlyphStyle.Shape = draw.CircleGlyph{}
			scatter.GlyphStyle.Radius = 3
			p.Add(scatter)
		}
		err := p.Save(450, 400, "spiral.png")
		if err != nil {
			//! add logger here
			panic(err)
		}
	}

}

func SpiralData(samples, classes int) ([][]float64, []float64) {
	//x := mat.NewDense(samples*classes, 2, nil)
	x := [][]float64{}
	y := []float64{}
	for i := 0; i < classes; i++ {
		r := linspace(0.0, 1.0, samples)
		t := linspace(float64(i)*4.0, (float64(i)+1.0)*4.0, samples)
		if len(t) == len(r) {
			for index := range t {
				t[index] += rand.NormFloat64() * 0.2
				sin := math.Sin(t[index] * 2.5)
				cos := math.Cos(t[index] * 2.5)
				x = append(x, []float64{r[index] * sin, r[index] * cos})
			}
		} else {
			return nil, nil
		}
		for j := i * samples; j < (i+1)*samples; j++ {
			y = append(y, float64(i))
		}
	}
	return x, y
}

// replacement for numpy linspace
func linspace(lo, hi float64, num int) []float64 {
	res := make([]float64, num)
	if num == 1 {
		res[0] = lo
		return res
	}
	for i := 0; i < num; i++ {
		res[i] = lo + float64(i)*(hi-lo)/float64(num-1)
	}
	return res
}
