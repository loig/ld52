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
	"image"
)

func (g *game) drawHUD(screen *ebiten.Image) {
	g.h.drawHUD(screen)
	g.drawWheatHUD(screen)
	g.drawRunHUD(screen)
	if g.reached > 0 {
		g.drawDistanceHUD(screen, false)
		g.drawDistanceHUD(screen, true)
	}
}

func (g *game) drawDistanceHUD(screen *ebiten.Image, record bool) {

	yShift := float64(10 + spriteSize + 10)

	dist := int(g.reached / 15)
	logo := distImg
	if record {
		dist = int(g.bestReached / 15)
		yShift += float64(10 + spriteSize)
		logo = maxDistImg
	}

	digits := getDigits(dist)
	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	width := float64(spriteSize + digitTileSize*len(digits))
	scale := 1.0
	if width > screenWidth-8 {
		scale = float64(screenWidth-8) / width
		width = width * scale
	}
	distX := (float64(screenWidth)-width)/2 - 4

	op := &ebiten.DrawImageOptions{}
	op2 := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(distX, yShift)
	op2.GeoM.Scale(scale, scale)
	op2.GeoM.Translate(distX, yShift)
	screen.DrawImage(wbgImage, op2)
	screen.DrawImage(logo, op)

	op.GeoM.Translate(spriteSize*scale, 0)
	op2.GeoM.Translate(spriteSize*scale, 0)
	op.ColorM.Scale(0.5, 0.7, 1, 1)
	for i := len(digits) - 1; i >= 0; i-- {
		screen.DrawImage(wbgImage, op2)
		screen.DrawImage(digitsImage.SubImage(image.Rect(digits[i]*digitTileSize, 0, (digits[i]+1)*digitTileSize, spriteSize)).(*ebiten.Image), op)
		op.GeoM.Translate(digitTileSize*scale, 0)
		op2.GeoM.Translate(digitTileSize*scale, 0)
	}
}

func (g *game) drawRunHUD(screen *ebiten.Image) {
	digits := getDigits(g.numRun)
	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	clockX := float64(screenWidth-(spriteSize+digitTileSize*len(digits))) - 12

	op := &ebiten.DrawImageOptions{}
	op2 := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(clockX, 10)
	op2.GeoM.Translate(clockX, 10)
	screen.DrawImage(wbgImage, op2)
	screen.DrawImage(clockImg, op)

	op.GeoM.Translate(spriteSize, 0)
	op2.GeoM.Translate(spriteSize, 0)
	op.ColorM.Scale(0.5, 0.7, 1, 1)
	for i := len(digits) - 1; i >= 0; i-- {
		screen.DrawImage(wbgImage, op2)
		screen.DrawImage(digitsImage.SubImage(image.Rect(digits[i]*digitTileSize, 0, (digits[i]+1)*digitTileSize, spriteSize)).(*ebiten.Image), op)
		op.GeoM.Translate(digitTileSize, 0)
		op2.GeoM.Translate(digitTileSize, 0)
	}

}

func (g *game) drawWheatHUD(screen *ebiten.Image) {
	//ebitenutil.DebugPrintAt(screen, fmt.Sprint("Wheat: ", g.getWheatForDisplay()), 0, 70)

	shift := 4.0

	digits := getDigits(g.getWheatForDisplay())
	if len(digits) == 0 {
		digits = append(digits, 0)
	}
	//wheatNumDigits := len(digits)
	//wheatX := float64(screenWidth-(spriteSize+digitTileSize*wheatNumDigits))/2 + shift
	//wheatX := float64(screenWidth-(spriteSize+digitTileSize*wheatNumDigits)) + shift

	wheatX := shift

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(wheatX, 10)
	screen.DrawImage(wbgImage, op)
	screen.DrawImage(wheatLogoImage, op)

	op.GeoM.Translate(spriteSize, 0)
	for i := len(digits) - 1; i >= 0; i-- {
		screen.DrawImage(wbgImage, op)
		screen.DrawImage(digitsImage.SubImage(image.Rect(digits[i]*digitTileSize, 0, (digits[i]+1)*digitTileSize, spriteSize)).(*ebiten.Image), op)
		op.GeoM.Translate(digitTileSize, 0)
	}
}
