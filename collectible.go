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
	"log"
	"math/rand"
)

type collectible struct {
	kind         int
	x, y         float64
	sizeX, sizeY float64
	collideBox   box
}

type collectibleSet struct {
	content []collectible
}

func (s *collectibleSet) update(harvesterBox box, ySpeedHarvester float64) {
	s.move(ySpeedHarvester)
	s.collect(harvesterBox)
	s.generate()
}

func (s *collectibleSet) move(ySpeedHarvester float64) {
	for i := 0; i < len(s.content); i++ {
		s.content[i].collideBox.p.x = s.content[i].x - s.content[i].sizeX/2
		s.content[i].collideBox.p.y = s.content[i].y - s.content[i].sizeY/2
		s.content[i].collideBox.r.x = s.content[i].x + s.content[i].sizeX/2
		s.content[i].collideBox.r.y = s.content[i].y - s.content[i].sizeY/2

		s.content[i].y -= ySpeedHarvester

		s.content[i].collideBox.q.x = s.content[i].x - s.content[i].sizeX/2
		s.content[i].collideBox.q.y = s.content[i].y + s.content[i].sizeY/2
		s.content[i].collideBox.s.x = s.content[i].x + s.content[i].sizeX/2
		s.content[i].collideBox.s.y = s.content[i].y + s.content[i].sizeY/2

		if s.content[i].y-s.content[i].sizeY/2 > screenHeight {
			copy(s.content[i:], s.content[i+1:])
			s.content = s.content[:len(s.content)-1]
			i--
			log.Print("out")
		}
	}
}

func (s *collectibleSet) collect(harvesterBox box) {
	for i := 0; i < len(s.content); i++ {
		if intersectBox(harvesterBox, s.content[i].collideBox) {
			copy(s.content[i:], s.content[i+1:])
			s.content = s.content[:len(s.content)-1]
			i--
		}
	}
}

func (s *collectibleSet) generate() {
	if rand.Intn(100) == 0 {
		s.content = append(s.content, collectible{
			kind: 0,
			x:    screenWidth / 2, y: 0,
			sizeX: 20,
			sizeY: 50,
		})
	}
}

func (s *collectibleSet) draw(screen *ebiten.Image) {
	for _, c := range s.content {
		ebitenutil.DrawRect(screen, c.x-c.sizeX/2, c.y-c.sizeY/2, c.sizeX, c.sizeY, color.RGBA{R: 255, A: 255})
		c.collideBox.draw(screen)
	}
}