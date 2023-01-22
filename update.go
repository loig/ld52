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
	//"log"
)

func (g *game) updateWheat() {
	g.wheat += g.h.actualSpeed * g.h.bladeSize
	if g.wheat > 99999000 {
		g.wheat = 99999000
	}
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
		g.h.initialSpeed = g.h.speed
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
	boing := g.h.update()
	if boing {
		g.playSound(soundReboundID)
	}
	g.reached -= g.h.ySpeed
	g.ps.genHarvesterParticles(g.h.xBladeLeft, g.h.xBladeRight, g.h.yBladeLeft, g.h.yBladeRight, g.h.bladeSize, g.h.nitro > 0)
	if g.h.nitro <= 0 {
		if g.h.gas > 0 {
			g.ps.genConsumeParticles(screenWidth-32, screenHeight-spriteSize-10-g.h.gas/gasDivider*spriteSize, 0)
		} else {
			speedDivider := (maxSpeed[len(maxSpeed)-1] * gasDivider) / gasTank[len(gasTank)-1]
			g.ps.genConsumeParticles(20, screenHeight-spriteSize-10-g.h.speed/speedDivider*spriteSize, 1)
		}
	}
	g.updateWheat()
	g.t.update(g.h.xPosition, g.h.yPosition, g.h.xSpeed, g.h.ySpeed, g.h.xBladeLeft, g.h.xBladeRight, g.h.yBladeLeft, g.h.yBladeRight)
	gas, nitro, stone := g.s.update(g.h.collideBox, g.h.ySpeed, g.gasRate, g.nitroRate, g.stoneRate, &(g.ps), g.reached)
	soundID := g.h.consume(gas, nitro, stone)
	if soundID >= 0 {
		g.playSound(soundID)
	}
	g.updateField()
	done = g.h.actualSpeed <= 0 || (!g.infiniteMode && g.reached >= goalDistance)
	return done
}

func (g *game) updateShop() (done bool) {
	spent, done, newButton := g.shop.update(g.getWheatForDisplay())
	if spent > 0 {
		g.wheat -= float64(spent) * wheatConversionRate
		g.playSound(soundBuyID)
	} else if spent < 0 {
		g.playSound(soundMissBuyID)
	} else if newButton {
		g.playSound(soundMenuButtonID)
	}
	return done
}

func (g *game) Update() error {

	yHarvesterMove := 0.0

	g.tutoUpdate()

	switch g.state {
	case stateLaunch1:
		g.updateMusic(music1ID, 0.8)
		if g.updateLaunch1() {
			g.state++
		}
	case stateLaunch2:
		g.updateMusic(music1ID, 0.8)
		if g.updateLaunch2() {
			g.state++
			g.t.setup(g.h.xPosition, g.h.yPosition, g.h.xBladeLeft, g.h.xBladeRight, g.h.yBladeLeft, g.h.yBladeRight)
			g.numRun++
			if g.numRun > 99 {
				g.numRun = 99
			}
		}
	case stateRun:
		if g.h.nitro > 0 {
			g.updateMusic(music3ID, 1)
		} else {
			volume := 0.0
			if g.h.initialSpeed != 0 {
				volume = g.h.speed / g.h.initialSpeed
			}
			g.updateMusic(music2ID, volume)
		}
		if g.reached > g.bestReached {
			g.bestReached = g.reached
		}
		if g.updateRun() {
			if g.reached >= goalDistance && !g.infiniteMode {
				g.state = stateEnd
				g.stopMusic()
				return nil
			}
			g.state++
			g.trans.setToShop()
		}
		yHarvesterMove = g.h.ySpeed
	case stateShop:
		g.updateMusic(music1ID, 0.8)
		if g.updateShop() {
			g.state++
			g.trans.setFromShop()
			g.playSound(soundMenuExitID)
		}
	case stateTransToShop, stateTransFromShop, stateTransToLaunch, stateTransToTitle:
		if g.state != stateTransToShop && g.state != stateTransToTitle {
			g.updateMusic(music1ID, 0.8)
		}
		if g.trans.update() {
			if g.state == stateTransToShop {
				g.state = stateShop
			} else if g.state == stateTransFromShop {
				g.state = stateTransToLaunch
				g.trans.setToLaunch()
				g.reset()
			} else if g.state == stateTransToLaunch {
				g.state = stateLaunch1
			} else if g.state == stateTransToTitle {
				g.shop.reset()
				g.wheat = 0
				g.reset()
				g.numRun = 0
				g.state = stateTitle
			}
		}
	case stateTitle:
		g.updateMusic(music1ID, 1)
		x, y := ebiten.CursorPosition()
		onButton := g.onInfiniteButton(float64(x), float64(y)) || g.onNormalButton(float64(x), float64(y))
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if onButton {
				g.state++
				g.playSound(soundStartID)
			} else {
				g.playSound(soundMissBuyID)
			}
		}
	case stateEnd:
		if g.ps.genVictoryParticles() {
			g.playSound(soundFireworkID)
		}
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.state = stateTransToTitle
			g.trans.setToTitle()
			if g.minNumRun == 0 || g.minNumRun > g.numRun {
				g.minNumRun = g.numRun
			}
		}
	}

	g.ps.update(yHarvesterMove)
	return nil
}
