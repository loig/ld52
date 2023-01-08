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

const (
	collectibleGas int = iota
	collectibleNitro
	collectibleStone
)

func (s *collectibleSet) update(harvesterBox box, ySpeedHarvester float64, gasRate, nitroRate, stoneRate int, ps *particleSys, reached float64) (gas, nitro, stone int) {
	s.move(ySpeedHarvester)
	gas, nitro, stone = s.collect(harvesterBox, ps)
	if ySpeedHarvester < 0 {
		s.generate(gasRate, nitroRate, stoneRate, reached)
	}
	return
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
		}
	}
}

func (s *collectibleSet) collect(harvesterBox box, ps *particleSys) (gas, nitro, stone int) {
	for i := 0; i < len(s.content); i++ {
		if intersectBox(harvesterBox, s.content[i].collideBox) {
			switch s.content[i].kind {
			case collectibleGas:
				gas++
				ps.genCollectParticles(s.content[i].x, s.content[i].y, 0)
			case collectibleNitro:
				nitro++
				ps.genCollectParticles(s.content[i].x, s.content[i].y, 1)
			case collectibleStone:
				stone++
				ps.genCollectParticles(s.content[i].x, s.content[i].y, 2)
			}
			copy(s.content[i:], s.content[i+1:])
			s.content = s.content[:len(s.content)-1]
			i--
		}
	}
	return
}

func (s *collectibleSet) generate(gasRate, nitroRate, stoneRate int, reached float64) {

	if nitroRate > 0 {
		if rand.Intn(nitroRate) == 0 {
			xPos := fieldStart + rand.Float64()*(fieldWidth)
			s.content = append(s.content, collectible{
				kind: collectibleNitro,
				x:    xPos, y: -8,
				sizeX: 16,
				sizeY: 16,
			})
			return
		}
	}

	if gasRate > 0 {
		if rand.Intn(gasRate) == 0 {
			xPos := fieldStart + rand.Float64()*(fieldWidth)
			s.content = append(s.content, collectible{
				kind: collectibleGas,
				x:    xPos, y: -8,
				sizeX: 16,
				sizeY: 16,
			})
			return
		}
	}

	if rand.Intn(stoneRate) == 0 && reached > stonePoint {
		xPos := fieldStart + rand.Float64()*(fieldWidth)
		s.content = append(s.content, collectible{
			kind: collectibleStone,
			x:    xPos, y: -8,
			sizeX: 16,
			sizeY: 16,
		})
	}
}

func (s *collectibleSet) draw(screen *ebiten.Image) {
	for i := len(s.content) - 1; i >= 0; i-- {
		c := s.content[i]
		img := stoneImage
		switch c.kind {
		case collectibleGas:
			img = gasCollectImage
		case collectibleNitro:
			img = nitroImage
		}
		//ebitenutil.DrawRect(screen, c.x-c.sizeX/2, c.y-c.sizeY/2, c.sizeX, c.sizeY, col)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(c.x-c.sizeX/2, c.y-c.sizeY/2)
		screen.DrawImage(img, op)
		c.collideBox.draw(screen)
	}
}
