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
	//"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	//"image/color"
	//"log"
	"image"
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
	displayPrice                                               int
	onOut                                                      bool
}

type shopUpdater struct {
	xTopLeft, yTopLeft, xBottomRight, yBottomRight float64
	level                                          *int
	price                                          []int
	name                                           string
	image                                          *ebiten.Image
	startFromZero                                  bool
	active                                         bool
}

var gasTank []float64 = []float64{1000, 2500, 4000} // 1000
var gasTankPrice []int = []int{0, 1000, 5000}

var gasRate []int = []int{-1, 1000, 500, 200} // -1
var gasRatePrice []int = []int{0, 500, 1500, 8000}

var gasEfficiency []float64 = []float64{250, 500, 1000}
var gasEfficiencyPrice []int = []int{0, 2000, 7000}

var nitroOnField []int = []int{10, 2000, 1000, 200} // -1
var nitroOnFieldPrice []int = []int{0, 2000, 7000, 15000}

var nitroEfficiency []float64 = []float64{10, 20, 50} // 10
var nitroEfficiencyPrice []int = []int{0, 1000, 3000}

var bladeSize []float64 = []float64{32, 64, 128}
var bladeSizePrice []int = []int{0, 10, 30}

var maxSpeed []float64 = []float64{2.5, 5, 10} // 3
var maxSpeedPrice []int = []int{0, 150, 500}

var stoneOnField []int = []int{100, 200, 500} // 100
var stoneOnFieldPrice []int = []int{0, 150, 500}

var stoneLoss []float64 = []float64{7, 5, 1}
var stoneLossPrice []int = []int{0, 1000, 5000}

func initShop() (s *shop) {

	s = &shop{}

	s.xOutTopLeft = screenWidth - 42
	s.yOutTopLeft = screenHeight - 42
	s.xOutBottomRight = screenWidth - 10
	s.yOutBottomRight = screenHeight - 10

	updaterXMargin := float64(screenWidth-128) / 2
	updaterXSize := screenWidth - 2*updaterXMargin
	updaterYMargin := 64.0
	updaterYSize := 32.0
	updaterYSpace := 4.0

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
		image: butgtImage,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level:         &(s.gasOnFieldLevel),
		price:         gasRatePrice,
		name:          "Gas on Field Rate",
		image:         butgcImage,
		startFromZero: true,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.gasEfficiencyLevel),
		price: gasEfficiencyPrice,
		name:  "Gas Efficiency",
		image: butgeImage,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level:         &(s.nitroOnFieldLevel),
		price:         nitroOnFieldPrice,
		name:          "Nitro on Field Rate",
		image:         butncImage,
		startFromZero: true,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.nitroEfficiencyLevel),
		price: nitroEfficiencyPrice,
		name:  "Nitro Efficiency",
		image: butneImage,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.bladeLevel),
		price: bladeSizePrice,
		name:  "Blade Size",
		image: butbeImage,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.speedLevel),
		price: maxSpeedPrice,
		name:  "Max Speed",
		image: butseImage,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.stoneOnFieldLevel),
		price: stoneOnFieldPrice,
		name:  "Stone on Field Rate",
		image: butstcImage,
	})

	yTop = yBottom + updaterYSpace
	yBottom = yTop + updaterYSize

	s.updaters = append(s.updaters, shopUpdater{
		xTopLeft: xTop, yTopLeft: yTop,
		xBottomRight: xBottom, yBottomRight: yBottom,
		level: &(s.stoneProtectionLevel),
		price: stoneLossPrice,
		name:  "Stone Protection",
		image: butsteImage,
	})

	return
}

func (s *shop) update(wheat int) (spent int, done bool) {

	x, y := ebiten.CursorPosition()
	s.displayPrice = 0

	for i, su := range s.updaters {
		s.updaters[i].active = float64(x) > su.xTopLeft && float64(x) < su.xBottomRight &&
			float64(y) > su.yTopLeft && float64(y) < su.yBottomRight
		if s.updaters[i].active && len(su.price) > *(su.level)+1 {
			s.displayPrice = su.price[*(su.level)+1]
		}
	}

	s.onOut = float64(x) > s.xOutTopLeft && float64(x) < s.xOutBottomRight &&
		float64(y) > s.yOutTopLeft && float64(y) < s.yOutBottomRight

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if s.onOut {
			done = true
			return
		}
		for _, su := range s.updaters {
			if su.active {
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

	//ebitenutil.DebugPrint(screen, fmt.Sprint(wheat))

	for _, su := range s.updaters {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(su.xTopLeft, su.yTopLeft)
		if su.active {
			op.ColorM.Scale(0.8, 1, 1, 1)
		}
		if *(su.level)+1 >= len(su.price) ||
			su.price[*(su.level)+1] > wheat ||
			(su.name == "Nitro Efficiency" && s.nitroOnFieldLevel == 0) ||
			(su.name == "Gas Efficiency" && s.gasOnFieldLevel == 0) {
			op.ColorM.Scale(0.5, 0.5, 0.5, 1)
		}
		screen.DrawImage(bgbuttonImage, op)
		screen.DrawImage(su.image, op)

		op.GeoM.Translate(64, 1)
		for i := 0; i < 3; i++ {
			levelImg := slevImage
			if (!su.startFromZero && *(su.level) >= 3-i-1) ||
				(su.startFromZero && *(su.level) >= 3-i) {
				levelImg = slevOKImage
			}
			screen.DrawImage(levelImg, op)
			op.GeoM.Translate(0, 10)
		}

		/*
			col := color.RGBA{R: 125, A: 100}
			price := ""
			if *(su.level)+1 >= len(su.price) {
				col.R = 25
			} else {
				price = fmt.Sprint(su.price[*(su.level)+1])
			}
			//ebitenutil.DrawRect(screen, su.xTopLeft, su.yTopLeft, su.xBottomRight-su.xTopLeft, su.yBottomRight-su.yTopLeft, col)
			ebitenutil.DebugPrintAt(screen, su.name, int(su.xTopLeft), int(su.yTopLeft))
			ebitenutil.DebugPrintAt(screen, price, int(su.xTopLeft), int(su.yTopLeft)+10)
		*/
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.xOutTopLeft, s.yOutTopLeft)
	if s.onOut {
		op.ColorM.Scale(0.8, 1, 1, 1)
	}
	screen.DrawImage(boutImage, op)

	if s.displayPrice > 0 {
		shift := -4.0

		digits := getDigits(s.displayPrice)
		if len(digits) == 0 {
			digits = append(digits, 0)
		}
		wheatNumDigits := len(digits)
		wheatX := float64(screenWidth-(spriteSize+digitTileSize*wheatNumDigits))/2 + shift

		op := &ebiten.DrawImageOptions{}
		op.ColorM.Scale(0.7, 1, 1, 1)
		op.GeoM.Translate(wheatX, screenHeight-74)
		screen.DrawImage(wbgImage, op)
		screen.DrawImage(wheatLogoImage, op)

		op.GeoM.Translate(spriteSize, 0)
		for i := len(digits) - 1; i >= 0; i-- {
			screen.DrawImage(wbgImage, op)
			screen.DrawImage(digitsImage.SubImage(image.Rect(digits[i]*digitTileSize, 0, (digits[i]+1)*digitTileSize, spriteSize)).(*ebiten.Image), op)
			op.GeoM.Translate(digitTileSize, 0)
		}
	}

	//ebitenutil.DrawRect(screen, s.xOutTopLeft, s.yOutTopLeft, s.xOutBottomRight-s.xOutTopLeft, s.yOutBottomRight-s.yOutTopLeft, color.RGBA{R: 125, A: 255})

}
