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
	"image"
	//"image/color"
	"math"
)

type harvester struct {
	actualSpeed                                      float64
	speed                                            float64
	speedLoss                                        float64
	stoneSpeedLoss                                   float64
	maxSpeed                                         float64
	speedStep                                        float64
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
	rotationStep                                     float64
	xPosition, yPosition                             float64
	bladeSize                                        float64
	bladeLevel                                       int
	xBladeLeft, yBladeLeft, xBladeRight, yBladeRight float64
	xSprite, ySprite                                 float64
	collideBox                                       box
	animationStep                                    int
	animationFrame                                   int
}

func (h *harvester) update() {
	h.updateGas()
	h.updateSpeed()
	h.updateNitro()
	h.updatePosition()
	h.updateCollideBox()
	h.updateAnimation()
}

func (h *harvester) updateAnimation() {
	h.animationFrame++
	if h.animationFrame >= harvesterAnimationFrames {
		h.animationStep = (h.animationStep + 1) % harvesterAnimationSteps
		h.animationFrame = 0
	}
}

func (h *harvester) updateGas() {
	if h.nitro <= 0 {
		if h.gas > 0 {
			h.gas -= h.gasConsumption
		} else {
			h.gas = 0
		}
	}
}

func (h *harvester) updateSpeed() {
	if h.gas <= 0 {
		if h.speed > 0 {
			h.speed -= h.speedLoss
		} else {
			h.speed = 0
		}
	} /* else {
		h.speed += h.speedLoss
		if h.speed > h.maxSpeed {
			h.speed = h.maxSpeed
		}
	} */
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

	if h.yPosition > goalPositionY {
		h.yPosition += h.ySpeed
		h.ySpeed = 0
	}
	if (h.xPosition < fieldStart && h.xSpeed < 0) ||
		(h.xPosition > fieldEnd && h.xSpeed > 0) {
		h.orientation = -(math.Pi + h.orientation)
		h.xSpeed = -h.xSpeed
	}
}

func (h *harvester) updateCollideBox() {

	xShift := math.Cos(h.orientation) * -bladeHeight
	yShift := math.Sin(h.orientation) * -bladeHeight

	h.collideBox.q.x = h.xBladeLeft + xShift
	h.collideBox.q.y = h.yBladeLeft + yShift
	h.collideBox.s.x = h.xBladeRight + xShift
	h.collideBox.s.y = h.yBladeRight + yShift

	h.xBladeLeft = h.bladeSize/2*math.Cos(h.orientation-math.Pi/2) + h.xPosition
	h.yBladeLeft = h.bladeSize/2*math.Sin(h.orientation-math.Pi/2) + h.yPosition
	h.xBladeRight = -h.bladeSize/2*math.Cos(h.orientation-math.Pi/2) + h.xPosition
	h.yBladeRight = -h.bladeSize/2*math.Sin(h.orientation-math.Pi/2) + h.yPosition

	h.collideBox.p.x = h.xBladeLeft
	h.collideBox.p.y = h.yBladeLeft
	h.collideBox.r.x = h.xBladeRight
	h.collideBox.r.y = h.yBladeRight

	h.xSprite = spriteSize/2*math.Cos(h.orientation-math.Pi/2) + h.xPosition
	h.ySprite = spriteSize/2*math.Sin(h.orientation-math.Pi/2) + h.yPosition
}

func (h *harvester) draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Rotate(h.orientation + math.Pi/2)
	options.GeoM.Translate(
		h.xSprite,
		h.ySprite,
	)
	screen.DrawImage(moissImages[h.animationStep], &options)

	options = ebiten.DrawImageOptions{}
	options.GeoM.Rotate(h.orientation + math.Pi/2)
	options.GeoM.Translate(
		h.xBladeLeft,
		h.yBladeLeft,
	)
	screen.DrawImage(moissLameImages[h.animationStep+h.bladeLevel*harvesterAnimationSteps], &options)

	//ebitenutil.DrawCircle(screen, h.xPosition, h.yPosition, 5, color.RGBA{R: 255, A: 255})
	//ebitenutil.DrawLine(screen, h.xBladeLeft, h.yBladeLeft, h.xBladeRight, h.yBladeRight, color.RGBA{R: 255, A: 255})
	//h.collideBox.draw(screen)
}

func (h *harvester) drawHUD(screen *ebiten.Image) {
	//ebitenutil.DrawRect(screen, 0, 10, h.gas/h.maxGas*100, 10, color.RGBA{B: 255, A: 255})
	//ebitenutil.DrawRect(screen, 0, 30, h.actualSpeed/h.maxSpeed*100, 10, color.RGBA{G: 255, A: 255})

	margin := 10.0

	// gas tank
	// Logo
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth-spriteSize-margin, screenHeight-spriteSize-margin)
	screen.DrawImage(gasLogoImage, op)

	tankDivider := 500.0

	// bg
	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(tankbgImage, op)

	for i := 0.0; i < (h.maxGas/tankDivider)-2; i++ {
		op.GeoM.Translate(0, -spriteSize)
		screen.DrawImage(tankbgImage, op)
	}

	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(tankbgImage, op)

	// Content
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth-spriteSize-margin, screenHeight-spriteSize-margin)
	gasHeight := h.gas / tankDivider * spriteSize
	gasDrawn := 0.0
	for gasDrawn+spriteSize <= gasHeight {
		op.GeoM.Translate(0, -spriteSize)
		screen.DrawImage(gasTankImages[3], op)
		gasDrawn += spriteSize
	}
	remaining := int(gasHeight - gasDrawn)
	op.GeoM.Translate(0, -float64(remaining))
	screen.DrawImage(gasTankImages[3].SubImage(image.Rect(0, 0, spriteSize, remaining)).(*ebiten.Image), op)

	// Container
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth-spriteSize-margin, screenHeight-spriteSize-margin)
	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(gasTankImages[0], op)

	for i := 0.0; i < (h.maxGas/tankDivider)-2; i++ {
		op.GeoM.Translate(0, -spriteSize)
		screen.DrawImage(gasTankImages[1], op)
	}

	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(gasTankImages[2], op)

	// speed display
	// logo
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(margin, screenHeight-spriteSize-margin)
	screen.DrawImage(speedLogoImage, op)

	speedDivider := (maxSpeed[len(maxSpeed)-1] * tankDivider) / gasTank[len(gasTank)-1]

	// bg
	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(tankbgImage, op)

	for i := 0.0; i < (h.maxGas/tankDivider)-2; i++ {
		op.GeoM.Translate(0, -spriteSize)
		screen.DrawImage(tankbgImage, op)
	}

	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(tankbgImage, op)

	// Content
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(margin, screenHeight-spriteSize-margin)
	speedHeight := h.speed / speedDivider * spriteSize
	speedDrawn := 0.0
	for speedDrawn+spriteSize <= speedHeight {
		op.GeoM.Translate(0, -spriteSize)
		screen.DrawImage(speedValueImage, op)
		speedDrawn += spriteSize
	}
	remaining = int(speedHeight - speedDrawn)
	op.GeoM.Translate(0, -float64(remaining))
	screen.DrawImage(speedValueImage.SubImage(image.Rect(0, 0, spriteSize, remaining)).(*ebiten.Image), op)

	// Container
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(margin, screenHeight-spriteSize-margin)
	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(gasTankImages[0], op)

	for i := 0.0; i < (h.maxSpeed/speedDivider)-2; i++ {
		op.GeoM.Translate(0, -spriteSize)
		screen.DrawImage(gasTankImages[1], op)
	}

	op.GeoM.Translate(0, -spriteSize)
	screen.DrawImage(gasTankImages[2], op)

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
