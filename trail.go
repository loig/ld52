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

type trailPart struct {
	x, y           float64
	xLeft, yLeft   float64
	xRight, yRight float64
}

type trail struct {
	parts      []trailPart
	triangles  [][]ebiten.Vertex
	goingRight bool
}

func (t *trail) setup(xHarvester, yHarvester, xHarvesterLeft, xHarvesterRight, yHarvesterLeft, yHarvesterRight float64) {
	t.parts = append(t.parts, trailPart{x: xHarvester, y: yHarvester, xLeft: xHarvesterLeft, xRight: xHarvesterRight, yLeft: yHarvesterLeft, yRight: yHarvesterRight})
	t.parts = append(t.parts, trailPart{x: xHarvester, y: yHarvester, xLeft: xHarvesterLeft, xRight: xHarvesterRight, yLeft: yHarvesterLeft, yRight: yHarvesterRight})
}

func (t *trail) update(xHarvester, yHarvester, xSpeedHarvester, ySpeedHarvester, xHarvesterLeft, xHarvesterRight, yHarvesterLeft, yHarvesterRight float64) {
	directionChanged :=
		(xSpeedHarvester >= 0 && !t.goingRight) ||
			(xSpeedHarvester < 0 && t.goingRight)
	for i := 0; i < len(t.parts); i++ {
		t.parts[i].y -= ySpeedHarvester
		t.parts[i].yLeft -= ySpeedHarvester
		t.parts[i].yRight -= ySpeedHarvester
	}
	for len(t.parts) > 2 && t.parts[len(t.parts)-2].yLeft > screenHeight && t.parts[len(t.parts)-2].yRight > screenHeight {
		t.parts = t.parts[:len(t.parts)-1]
	}
	if directionChanged {
		t.parts = append(t.parts, trailPart{})
		t.parts = append(t.parts, trailPart{})
		copy(t.parts[2:], t.parts[:len(t.parts)-2])
		t.parts[1].x = xHarvester
		t.parts[1].xLeft = xHarvesterLeft
		t.parts[1].xRight = xHarvesterRight
		t.parts[1].y = yHarvester
		t.parts[1].yLeft = yHarvesterLeft
		t.parts[1].yRight = yHarvesterRight
	}
	t.parts[0].x = xHarvester
	t.parts[0].xLeft = xHarvesterLeft
	t.parts[0].xRight = xHarvesterRight
	t.parts[0].y = yHarvester
	t.parts[0].yLeft = yHarvesterLeft
	t.parts[0].yRight = yHarvesterRight
	t.goingRight = xSpeedHarvester >= 0
	t.getTriangles()
}

func makeVertex(x, y float64) ebiten.Vertex {
	return ebiten.Vertex{
		DstX: float32(x), DstY: float32(y),
		SrcX: float32(x), SrcY: float32(y),
		ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1,
	}
}

func (t *trail) getTriangles() {
	num := 0
	for i := 0; i < len(t.parts)-1; i += 2 {
		triangle1 := make([]ebiten.Vertex, 3)
		triangle1[0] = makeVertex(t.parts[i].xLeft, t.parts[i].yLeft)
		triangle1[1] = makeVertex(t.parts[i+1].xLeft, t.parts[i+1].yLeft)
		triangle1[2] = makeVertex(t.parts[i+1].xRight, t.parts[i+1].yRight)
		if len(t.triangles) <= num {
			t.triangles = append(t.triangles, triangle1)
		} else {
			t.triangles[num] = triangle1
		}
		num++
		triangle2 := make([]ebiten.Vertex, 3)
		triangle2[0] = makeVertex(t.parts[i].xRight, t.parts[i].yRight)
		triangle2[1] = makeVertex(t.parts[i].xLeft, t.parts[i].yLeft)
		triangle2[2] = makeVertex(t.parts[i+1].xRight, t.parts[i+1].yRight)
		if len(t.triangles) <= num {
			t.triangles = append(t.triangles, triangle2)
		} else {
			t.triangles[num] = triangle2
		}
		num++
	}
	t.triangles = t.triangles[:num]
	/*
		for i := 0; i < len(t.triangles); i++ {
			for j := 0; j < len(t.triangles[i]); j++ {
				if t.triangles[i][j].DstX < 0 {
					t.triangles[i][j].DstX = 0
				}
				if t.triangles[i][j].DstX >= fieldWidth {
					t.triangles[i][j].DstX = fieldWidth - 1
				}
				if t.triangles[i][j].DstY >= screenWidth {
					t.triangles[i][j].DstY = screenWidth - 1
				}
			}
		}
	*/
}

func (t *trail) applyOnImage(bgimg, fgimg *ebiten.Image) {
	opt := &ebiten.DrawTrianglesOptions{}
	opt.Address = ebiten.AddressRepeat
	opt.CompositeMode = ebiten.CompositeModeCopy
	for _, triangle := range t.triangles {
		fgimg.DrawTriangles(triangle, []uint16{0, 1, 2}, bgimg, opt)
	}
}

func (t *trail) draw(screen *ebiten.Image) {
	/*
		for i := 0; i < len(t.parts)-1; i += 2 {
			ebitenutil.DrawLine(screen, t.parts[i].x, t.parts[i].y, t.parts[i+1].x, t.parts[i+1].y, color.RGBA{R: 139, G: 69, B: 19, A: 255})
			ebitenutil.DrawLine(screen, t.parts[i].xLeft, t.parts[i].yLeft, t.parts[i+1].xLeft, t.parts[i+1].yLeft, color.RGBA{R: 139, G: 69, B: 19, A: 255})
			ebitenutil.DrawLine(screen, t.parts[i].xRight, t.parts[i].yRight, t.parts[i+1].xRight, t.parts[i+1].yRight, color.RGBA{R: 139, G: 69, B: 19, A: 255})
		}
	*/
}
