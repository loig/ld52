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
	"image/color"
)

func (g *game) drawLaunch(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	g.drawHUD(screen)
	g.h.draw(screen)
}

func (g *game) drawRun(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 255, G: 255, B: 0, A: 255})
	g.drawHUD(screen)
	g.t.draw(screen)
	g.s.draw(screen)
	g.h.draw(screen)
}

func (g *game) drawShop(screen *ebiten.Image) {

}

func (g *game) Draw(screen *ebiten.Image) {
	switch g.state {
	case stateLaunch1, stateLaunch2:
		g.drawLaunch(screen)
	case stateRun:
		g.drawRun(screen)
	case stateShop:
		g.drawShop(screen)
	}
}
