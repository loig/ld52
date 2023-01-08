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
	shop                          *shop
	gasRate, nitroRate, stoneRate int
	wheat                         float64
	fieldShift                    float64
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

	g.shop = initShop()

	g.h.speedLoss = 0.01
	g.h.speedStep = 0.1
	g.h.gasConsumption = 2.5
	g.h.nitroLoss = 1
	g.h.maxNitro = 250

	g.reset()

	return
}

func (g *game) reset() {
	g.fieldShift = 0

	g.h.xPosition = startPositionX
	g.h.yPosition = startPositionY
	g.h.orientation = -math.Pi / 2
	g.h.rotationStep = 0.03
	g.h.animationStep = 0
	g.h.animationFrame = 0

	// Speed
	g.h.maxSpeed = maxSpeed[g.shop.speedLevel]
	g.h.speed = 0
	g.h.xSpeed = 0
	g.h.ySpeed = 0
	g.h.actualSpeed = 0

	// Gas
	g.h.maxGas = gasTank[g.shop.gasTankLevel]
	g.h.gas = g.h.maxGas
	g.h.gasProduction = gasEfficiency[g.shop.gasEfficiencyLevel]
	g.gasRate = gasRate[g.shop.gasOnFieldLevel]

	// Nitro
	g.h.nitro = 0
	g.h.nitroSpeed = nitroEfficiency[g.shop.nitroEfficiencyLevel]
	g.nitroRate = nitroOnField[g.shop.nitroOnFieldLevel]

	// Stone
	g.h.stoneSpeedLoss = stoneLoss[g.shop.stoneProtectionLevel]
	g.stoneRate = stoneOnField[g.shop.stoneOnFieldLevel]

	// Blade
	g.h.bladeSize = bladeSize[g.shop.bladeLevel]
	g.h.bladeLevel = g.shop.bladeLevel

	g.t.parts = g.t.parts[0:0]

	g.s.content = g.s.content[0:0]
}

func (g game) getWheatForDisplay() int {
	return int(g.wheat / wheatConversionRate)
}
