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
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"image/color"
)

func (g *game) drawLaunch(screen *ebiten.Image) {
	g.drawField(screen, false)
	g.h.draw(screen)
	g.drawHUD(screen)
}

func (g *game) drawField(screen *ebiten.Image, drawTrail bool) {

	shift := g.fieldShift
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, -fieldTileSize)
	op.GeoM.Translate(0, shift)

	for filled := shift; filled <= screenHeight+fieldTileSize; filled += fieldTileSize {
		fgImage.DrawImage(fieldImage, op)
		op.GeoM.Translate(0, fieldTileSize)
	}

	shift = g.fieldShift
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, -fieldTileSize)
	op.GeoM.Translate(0, shift)

	for filled := shift; filled <= screenHeight+fieldTileSize; filled += fieldTileSize {
		bgImage.DrawImage(groundImage, op)
		op.GeoM.Translate(0, fieldTileSize)
	}

	if drawTrail {
		g.t.applyOnImage(bgImage, fgImage)
	}
	op = &ebiten.DrawImageOptions{}

	screen.DrawImage(fgImage, op)

}

func (g *game) drawRun(screen *ebiten.Image) {
	//ebitenutil.DrawRect(screen, fieldStart, 0, fieldWidth, screenHeight, color.RGBA{R: 255, G: 255, B: 0, A: 255})
	g.drawField(screen, true)
	g.t.draw(screen)
	g.s.draw(screen)
	g.h.draw(screen)
	g.drawHUD(screen)
}

func (g *game) drawShop(screen *ebiten.Image) {
	g.drawField(screen, true)
	g.s.draw(screen)
	g.h.draw(screen)
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, 0.5)
	screen.DrawImage(blackbgImage, op)

	g.drawWheatHUD(screen)
	g.shop.draw(screen, g.getWheatForDisplay())
}

func (g *game) Draw(screen *ebiten.Image) {
	switch g.state {
	case stateLaunch1, stateLaunch2:
		g.drawLaunch(screen)
		g.ps.draw(screen)
	case stateRun:
		g.drawRun(screen)
		g.ps.draw(screen)
	case stateShop:
		g.drawShop(screen)
	case stateTransToShop:
		g.drawRun(screen)
		g.ps.draw(screen)
		g.trans.draw(screen)
	case stateTransFromShop:
		g.drawShop(screen)
		g.ps.draw(screen)
		g.trans.draw(screen)
	case stateTransToLaunch:
		g.drawLaunch(screen)
		g.ps.draw(screen)
		g.trans.draw(screen)
	}

}
