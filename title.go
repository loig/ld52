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
)

var xNormalModeTop float64 = float64(screenWidth-128) / 2
var xNormalModeBottom float64 = xNormalModeTop + 128
var yNormalModeTop float64 = 2 * float64(screenHeight) / 3
var yNormalModeBottom float64 = yNormalModeTop + 32

func (g *game) onNormalButton(x, y float64) bool {
	onButton := x >= xNormalModeTop && x <= xNormalModeBottom && y >= yNormalModeTop && y <= yNormalModeBottom
	if !g.onNButton {
		if onButton {
			g.onNButton = true
			g.playSound(soundMenuButtonID)
		}
	} else {
		g.onNButton = onButton
	}
	return onButton
}

var xInfiniteModeTop float64 = float64(screenWidth-128) / 2
var xInfiniteModeBottom float64 = xInfiniteModeTop + 128
var yInfiniteModeTop float64 = yNormalModeBottom + 4
var yInfiniteModeBottom float64 = yInfiniteModeTop + 32

func (g *game) onInfiniteButton(x, y float64) bool {
	onButton := x >= xInfiniteModeTop && x <= xInfiniteModeBottom && y >= yInfiniteModeTop && y <= yInfiniteModeBottom
	if !g.onIButton {
		if onButton {
			g.onIButton = true
			g.playSound(soundMenuButtonID)
		}
	} else {
		g.onIButton = onButton
	}
	g.infiniteMode = onButton && g.bestReached > 0
	return g.infiniteMode
}

func (g *game) drawTitle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if g.onNButton {
		op.ColorM.Scale(0.8, 1, 1, 1)
	}
	op.GeoM.Translate(xNormalModeTop, yNormalModeTop)
	screen.DrawImage(bgbuttonImage, op)
	screen.DrawImage(normalMImg, op)
	op = &ebiten.DrawImageOptions{}
	if g.onIButton {
		op.ColorM.Scale(0.8, 1, 1, 1)
	}
	if g.bestReached <= 0 {
		op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	}
	op.GeoM.Translate(xInfiniteModeTop, yInfiniteModeTop)
	screen.DrawImage(bgbuttonImage, op)
	screen.DrawImage(infiniteMImg, op)
}
