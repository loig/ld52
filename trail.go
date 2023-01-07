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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type trailPart struct {
	x, y float64
}

type trail struct {
	parts      []trailPart
	goingRight bool
}

func (t *trail) update(xHarvester, yHarvester, xSpeedHarvester, ySpeedHarvester float64) {
	directionChanged :=
		(xSpeedHarvester >= 0 && !t.goingRight) ||
			(xSpeedHarvester < 0 && t.goingRight)
	for i := 0; i < len(t.parts); i++ {
		t.parts[i].y -= ySpeedHarvester
	}
	if len(t.parts) > 2 && t.parts[len(t.parts)-2].y > screenHeight {
		t.parts = t.parts[:len(t.parts)-1]
	}
	if directionChanged || len(t.parts) <= 0 {
		t.parts = append(t.parts, trailPart{})
		copy(t.parts[1:], t.parts[:len(t.parts)-1])
	}
	t.parts[0].x = xHarvester
	t.parts[0].y = yHarvester
}

func (t *trail) draw(screen *ebiten.Image) {
	for i := 0; i < len(t.parts)-1; i++ {
		ebitenutil.DrawLine(screen, t.parts[i].x, t.parts[i].y, t.parts[i+1].x, t.parts[i+1].y, color.RGBA{R: 139, G: 69, B: 19, A: 255})
	}
}
