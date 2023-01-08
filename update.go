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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) updateWheat() {
	g.wheat += g.h.actualSpeed * g.h.bladeSize
}

func (g *game) updateLaunch1() (done bool) {
	g.h.updateCollideBox()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		done = true
	} else {
		g.h.orientation += g.h.rotationStep
		if g.h.orientation >= maxAngle {
			g.h.orientation = maxAngle
			g.h.rotationStep = -g.h.rotationStep
		} else if g.h.orientation <= minAngle {
			g.h.orientation = minAngle
			g.h.rotationStep = -g.h.rotationStep
		}
	}
	return done
}

func (g *game) updateLaunch2() (done bool) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		done = true
	} else {
		g.h.speed += g.h.speedStep
		if g.h.speed > g.h.maxSpeed {
			g.h.speed = minSpeed
		}
	}
	g.h.actualSpeed = g.h.speed
	return done
}

func (g *game) updateField() {
	g.fieldShift -= g.h.ySpeed
	for g.fieldShift > fieldTileSize {
		g.fieldShift -= fieldTileSize
	}
}

func (g *game) updateRun() (done bool) {
	g.h.update()
	g.updateWheat()
	g.t.update(g.h.xPosition, g.h.yPosition, g.h.xSpeed, g.h.ySpeed, g.h.xBladeLeft, g.h.xBladeRight, g.h.yBladeLeft, g.h.yBladeRight)
	gas, nitro, stone := g.s.update(g.h.collideBox, g.h.ySpeed, g.gasRate, g.nitroRate, g.stoneRate)
	g.h.consume(gas, nitro, stone)
	g.updateField()
	done = g.h.speed <= 0
	return done
}

func (g *game) updateShop() (done bool) {
	spent, done := g.shop.update(g.getWheatForDisplay())
	g.wheat -= float64(spent) * wheatConversionRate
	return done
}

func (g *game) Update() error {
	switch g.state {
	case stateLaunch1:
		if g.updateLaunch1() {
			g.state++
		}
	case stateLaunch2:
		if g.updateLaunch2() {
			g.state++
			g.t.setup(g.h.xPosition, g.h.yPosition, g.h.xBladeLeft, g.h.xBladeRight, g.h.yBladeLeft, g.h.yBladeRight)
		}
	case stateRun:
		if g.updateRun() {
			g.state++
		}
	case stateShop:
		if g.updateShop() {
			g.state = stateLaunch1
			g.reset()
		}
	}
	return nil
}
