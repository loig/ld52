/*
ld52, a game for Ludum Dare 52
Copyright (C) 2023 Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

type shop struct {
	gasTankLevel                                               int
	gasOnFieldLevel                                            int
	gasEfficiencyLevel                                         int
	nitroOnFieldLevel                                          int
	nitroEfficiencyLevel                                       int
	bladeLevel                                                 int
	speedLevel                                                 int
	stoneOnFieldLevel                                          int
	stoneProtectionLevel                                       int
	updaters                                                   []shopUpdater
	xOutTopLeft, yOutTopLeft, xOutBottomRight, yOutBottomRight float64
}

type shopUpdater struct {
	xTopLeft, yTopLeft, xBottomRight, yBottomRight float64
	level                                          *int
	price                                          []int
	name                                           string
}

var gasTank []float64 = []float64{1000, 1500, 2500, 4000}
var gasTankPrice []int = []int{0, 300, 1000, 5000}

var gasRate []int = []int{-1, 4000, 3000, 2000}
var gasRatePrice []int = []int{0, 500, 1500, 8000}

var gasEfficiency []float64 = []float64{250, 375, 625, 1000}
var gasEfficiencyPrice []int = []int{0, 500, 2000, 7000}

var nitroOnField []int = []int{-1, 10000, 5000, 2000}
var nitroOnFieldPrice []int = []int{0, 2000, 7000, 15000}

var nitroEfficiency []float64 = []float64{20, 40, 100}
var nitroEfficiencyPrice []int = []int{0, 1000, 3000}

var bladeSize []float64 = []float64{32, 64, 128, 256}
var bladeSizePrice []int = []int{0, 300, 1500, 5000}

var maxSpeed []float64 = []float64{3, 5, 7, 10}
var maxSpeedPrice []int = []int{0, 150, 250, 500}

var stoneOnField []int = []int{100, 200, 500}
var stoneOnFieldPrice []int = []int{0, 150, 500}

var stoneLoss []float64 = []float64{7, 5, 3, 1}
var stoneLossPrice []int = []int{0, 1000, 2000, 5000}

func initShop() (s shop) {

	s.xOutTopLeft = screenWidth - 50
	s.yOutTopLeft = screenHeight - 50
	s.xOutBottomRight = screenWidth - 10
	s.yOutBottomRight = screenHeight - 10

	updaterXMargin := 10.0
	updaterXSize := screenWidth - 2*updaterXMargin
	updaterYMargin := 20.0
	updaterYSize := 40.0
	updaterYSpace := 10.0

	xTop := updaterXMargin
	yTop := updaterYMargin
	xBottom := xTop + updaterXSize
	yBottom := yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.gasTankLevel),
		price: gasTankPrice,
		name:  "Gas Tank Size",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.gasOnFieldLevel),
		price: gasRatePrice,
		name:  "Gas on Field Rate",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.gasEfficiencyLevel),
		price: gasEfficiencyPrice,
		name:  "Gas Efficiency",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.nitroOnFieldLevel),
		price: nitroOnFieldPrice,
		name:  "Nitro on Field Rate",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	if s.nitroOnFieldLevel > 0 {
		s.updaters = append(s.updaters, shopUpdater{
			xTopLeft: xTop, yTopLeft: yTop,
			xBottomRight: xBottom, yBottomRight: yBottom,
			level: &(s.nitroEfficiencyLevel),
			price: nitroEfficiencyPrice,
			name:  "Nitro Efficiency",
		})

		yTop = yBottom + updaterYSpace
		yBottom = yTop + updaterYSize
	}

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.bladeLevel),
		price: bladeSizePrice,
		name:  "Blade Size",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.speedLevel),
		price: maxSpeedPrice,
		name:  "Max Speed",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.stoneOnFieldLevel),
		price: stoneOnFieldPrice,
		name:  "Stone on Field Rate",
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.stoneProtectionLevel),
		price: stoneLossPrice,
		name:  "Stone Protection",
	})

	return
}

func (s *shop) update(wheat int) (spent int, done bool) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if float64(x) > s.xOutTopLeft && float64(x) < s.xOutBottomRight &&
			float64(y) > s.yOutTopLeft && float64(y) < s.yOutBottomRight {
			done = true
			return
		}
		for _, su := range s.updaters {
			if float64(x) > su.xTopLeft && float64(x) < su.xBottomRight &&
				float64(y) > su.yTopLeft && float64(y) < su.yBottomRight {
				if len(su.price) > *(su.level)+1 && wheat >= su.price[*(su.level)+1] {
					spent = su.price[*(su.level)+1]
					*(su.level)++
				}
				break
			}
		}
	}
	return
}

func (s shop) draw(screen *ebiten.Image, wheat int) {

	ebitenutil.DebugPrint(screen, fmt.Sprint(wheat))

	for _, su := range s.updaters {
		col := color.RGBA{R: 125, A: 255}
		price := ""
		if *(su.level)+1 >= len(su.price) {
			col.R = 25
		} else {
			price = fmt.Sprint(su.price[*(su.level)+1])
		}
		ebitenutil.DrawRect(screen, su.xTopLeft, su.yTopLeft, su.xBottomRight-su.xTopLeft, su.yBottomRight-su.yTopLeft, color.RGBA{R: 125, A: 255})
		ebitenutil.DebugPrintAt(screen, su.name, int(su.xTopLeft), int(su.yTopLeft))
		ebitenutil.DebugPrintAt(screen, price, int(su.xTopLeft), int(su.yTopLeft)+10)
	}

	ebitenutil.DrawRect(screen, s.xOutTopLeft, s.yOutTopLeft, s.xOutBottomRight-s.xOutTopLeft, s.yOutBottomRight-s.yOutTopLeft, color.RGBA{R: 125, A: 255})

}
