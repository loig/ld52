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

const (
	mouseTutoLen = 60
	mouseTutoB1  = 20
	mouseTutoB2  = 30
	mouseTutoB3  = 35
	mouseTutoB4  = 45
)

func (g *game) tutoUpdate() {
	g.mouseTutoStep++
	if g.mouseTutoStep >= mouseTutoLen {
		g.mouseTutoStep = 0
	}
}

func (g *game) tutoDraw(screen *ebiten.Image) {

	yshift := yInfiniteModeBottom + 4
	xshift := float64(screenWidth-32) / 2

	if g.state == stateLaunch1 {
		yshift += 48
	}

	if g.state == stateLaunch2 {
		xshift -= 80
	}

	img := souris1Img
	if (g.mouseTutoStep >= mouseTutoB1 && g.mouseTutoStep < mouseTutoB2) ||
		(g.mouseTutoStep >= mouseTutoB3 && g.mouseTutoStep < mouseTutoB4) {
		img = souris2Img
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(xshift, yshift)
	screen.DrawImage(img, op)
}
