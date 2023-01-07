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
	"math"
)

type harvester struct {
	actualSpeed                                      float64
	speed                                            float64
	speedLoss                                        float64
	stoneSpeedLoss                                   float64
	maxSpeed                                         float64
	xSpeed, ySpeed                                   float64
	gas                                              float64
	gasConsumption                                   float64
	gasProduction                                    float64
	maxGas                                           float64
	nitroSpeed                                       float64
	nitro                                            float64
	maxNitro                                         float64
	nitroLoss                                        float64
	orientation                                      float64
	xPosition, yPosition                             float64
	bladeSize                                        float64
	xBladeLeft, yBladeLeft, xBladeRight, yBladeRight float64
	collideBox                                       box
}

func (h *harvester) update() {
	h.updateGas()
	h.updateSpeed()
	h.updateNitro()
	h.updatePosition()
}

func (h *harvester) updateGas() {
	if h.gas > 0 {
		h.gas -= h.gasConsumption
	} else {
		h.gas = 0
	}
}

func (h *harvester) updateSpeed() {
	if h.gas <= 0 {
		if h.speed > 0 {
			h.speed -= h.speedLoss
		} else {
			h.speed = 0
		}
	} else {
		h.speed += h.speedLoss
		if h.speed > h.maxSpeed {
			h.speed = h.maxSpeed
		}
	}
	h.actualSpeed = h.speed
}

func (h *harvester) updateNitro() {
	if h.nitro > 0 {
		h.nitro -= h.nitroLoss
		if h.nitro < 0 {
			h.nitro = 0
		}
		h.actualSpeed = h.nitroSpeed
	}
}

func (h *harvester) updatePosition() {

	h.xSpeed = math.Cos(h.orientation) * h.actualSpeed
	h.ySpeed = math.Sin(h.orientation) * h.actualSpeed

	h.xPosition += h.xSpeed
	//h.yPosition += h.ySpeed (background should move instead)
	if (h.xPosition < 0 && h.xSpeed < 0) ||
		(h.xPosition > screenWidth && h.xSpeed > 0) {
		h.orientation = -(math.Pi + h.orientation)
	}

	h.collideBox.q.x = h.xBladeLeft
	h.collideBox.q.y = h.yBladeLeft
	h.collideBox.s.x = h.xBladeRight
	h.collideBox.s.y = h.yBladeRight

	h.xBladeLeft = h.bladeSize/2*math.Cos(h.orientation-math.Pi/2) + h.xPosition
	h.yBladeLeft = h.bladeSize/2*math.Sin(h.orientation-math.Pi/2) + h.yPosition
	h.xBladeRight = -h.bladeSize/2*math.Cos(h.orientation-math.Pi/2) + h.xPosition
	h.yBladeRight = -h.bladeSize/2*math.Sin(h.orientation-math.Pi/2) + h.yPosition

	h.collideBox.p.x = h.xBladeLeft
	h.collideBox.p.y = h.yBladeLeft
	h.collideBox.r.x = h.xBladeRight
	h.collideBox.r.y = h.yBladeRight
}

func (h *harvester) draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, h.xPosition, h.yPosition, 5, color.RGBA{R: 255, A: 255})
	ebitenutil.DrawLine(screen, h.xBladeLeft, h.yBladeLeft, h.xBladeRight, h.yBladeRight, color.RGBA{R: 255, A: 255})
	h.collideBox.draw(screen)
}

func (h *harvester) drawHUD(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 10, h.gas/h.maxGas*100, 10, color.RGBA{B: 255, A: 255})
	ebitenutil.DrawRect(screen, 0, 30, h.actualSpeed/h.maxSpeed*100, 10, color.RGBA{G: 255, A: 255})
}

func (h *harvester) consume(gas, nitro, stone int) {
	h.gas += float64(gas) * h.gasProduction
	if h.gas > h.maxGas {
		h.gas = h.maxGas
	}

	if nitro > 0 {
		h.nitro = h.maxNitro
	}

	if h.nitro <= 0 {
		h.speed -= float64(stone) * h.stoneSpeedLoss
		if h.speed < 0 {
			h.speed = 0
		}
	}
}
