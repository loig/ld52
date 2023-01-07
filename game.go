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

type game struct {
	h                             harvester
	t                             trail
	s                             collectibleSet
	gasRate, nitroRate, stoneRate int
	wheat                         float64
}

func initGame() (g *game) {
	g = &game{}

	g.h.xPosition = screenWidth / 2
	g.h.yPosition = screenHeight - screenHeight/3
	g.h.speed = 1
	g.h.speedLoss = 0.01
	g.h.stoneSpeedLoss = 2
	g.h.maxSpeed = 5
	g.h.gas = 1000
	g.h.gasConsumption = 0.5
	g.h.gasProduction = 250
	g.h.maxGas = 1000
	g.h.nitroLoss = 1
	g.h.maxNitro = 250
	g.h.nitroSpeed = 10
	g.h.orientation = -1.5
	g.h.bladeSize = 100

	g.gasRate = 2500
	g.nitroRate = 10000
	g.stoneRate = 250

	return
}

func (g game) getWheatForDisplay() int {
	return int(g.wheat / 1000)
}
