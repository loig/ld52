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
	speed                float64
	speedLoss            float64
	maxSpeed             float64
	xSpeed, ySpeed       float64
	gas                  float64
	gasConsumption       float64
	maxGas               float64
	orientation          float64
	xPosition, yPosition float64
}

func (h *harvester) update() {
	h.updateGas()
	h.updateSpeed()
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
	h.xSpeed = math.Cos(h.orientation) * h.speed
	h.ySpeed = math.Sin(h.orientation) * h.speed
}

func (h *harvester) updatePosition() {
	h.xPosition += h.xSpeed
	//h.yPosition += h.ySpeed (background should move instead)
	if (h.xPosition < 0 && h.xSpeed < 0) ||
		(h.xPosition > screenWidth && h.xSpeed > 0) {
		h.orientation = -(math.Pi + h.orientation)
	}
}

func (h *harvester) draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, h.xPosition, h.yPosition, 5, color.RGBA{R: 255, A: 255})
}

func (h *harvester) drawHUD(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 10, h.gas/h.maxGas*100, 10, color.RGBA{B: 255, A: 255})
	ebitenutil.DrawRect(screen, 0, 30, h.speed/h.maxSpeed*100, 10, color.RGBA{G: 255, A: 255})
}
