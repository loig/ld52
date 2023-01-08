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

type point struct {
	x, y float64
}

type box struct {
	p, q, r, s point
}

// assumes that b1 and b2 are rectangles
// assumes that b2 sides are parallel to the axis
// assumes that, for b2, p is top left, q is bottom left, r is top right, s is bottom right
func intersectBox(b1, b2 box) bool {
	if b2.p.x > b1.p.x && b2.p.x > b1.q.x && b2.p.x > b1.r.x && b2.p.x > b1.s.x {
		return false
	}

	if b2.r.x < b1.p.x && b2.r.x < b1.q.x && b2.r.x < b1.r.x && b2.r.x < b1.s.x {
		return false
	}

	if b2.p.y > b1.p.y && b2.p.y > b1.q.y && b2.p.y > b1.r.y && b2.p.y > b1.s.y {
		return false
	}

	if b2.q.y < b1.p.y && b2.q.y < b1.q.y && b2.p.y < b1.r.y && b2.p.y < b1.s.y {
		return false
	}

	return true
}

func (b box) draw(screen *ebiten.Image) {
	//ebitenutil.DrawLine(screen, b.p.x, b.p.y, b.r.x, b.r.y, color.RGBA{B: 255, A: 255})
	//ebitenutil.DrawLine(screen, b.r.x, b.r.y, b.s.x, b.s.y, color.RGBA{B: 255, A: 255})
	//ebitenutil.DrawLine(screen, b.s.x, b.s.y, b.q.x, b.q.y, color.RGBA{B: 255, A: 255})
	//ebitenutil.DrawLine(screen, b.q.x, b.q.y, b.p.x, b.p.y, color.RGBA{B: 255, A: 255})
}

func getDigits(num int) (digits []int) {
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	return
}
