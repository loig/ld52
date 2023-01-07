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
	"math"
)

type game struct {
	state                         int
	h                             harvester
	t                             trail
	s                             collectibleSet
	gasRate, nitroRate, stoneRate int
	wheat                         float64
}

const (
	stateLaunch1 int = iota
	stateLaunch2
	stateRun
	stateShop
)

func initGame() (g *game) {
	g = &game{}

	g.state = stateLaunch1
	g.reset()

	g.h.speedLoss = 0.01
	g.h.stoneSpeedLoss = 2
	g.h.maxSpeed = 5
	g.h.gas = 1000
	g.h.gasConsumption = 2.5
	g.h.gasProduction = 250
	g.h.nitroLoss = 1
	g.h.maxNitro = 250
	g.h.nitroSpeed = 10
	g.h.bladeSize = 100

	g.gasRate = 2500
	g.nitroRate = 10000
	g.stoneRate = 250

	return
}

func (g *game) reset() {
	g.h.xPosition = screenWidth / 2
	g.h.yPosition = screenHeight - screenHeight/3
	g.h.speed = 0
	g.h.xSpeed = 0
	g.h.ySpeed = 0
	g.h.actualSpeed = 0
	g.h.speedStep = 0.1
	g.h.gas = 1000
	g.h.gasConsumption = 0.5
	g.h.gasProduction = 250
	g.h.maxGas = 1000
	g.h.nitro = 0
	g.h.orientation = -math.Pi / 2
	g.h.rotationStep = 0.03

	g.t.parts = g.t.parts[0:0]

	g.s.content = g.s.content[0:0]
}

func (g game) getWheatForDisplay() int {
	return int(g.wheat / 1000)
}
