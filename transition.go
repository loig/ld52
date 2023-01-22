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
)

type transition struct {
	numFrames    int
	currentFrame int
	alphaChange  float64
	alpha        float64
	alphaGoal    float64
}

func (t *transition) setToLaunch() {
	t.numFrames = 80
	t.currentFrame = 0
	t.alpha = 1
	t.alphaChange = -t.alpha / float64(t.numFrames)
	t.alphaGoal = 0
}

func (t *transition) setToShop() {
	t.numFrames = 40
	t.currentFrame = 0
	t.alpha = 0
	t.alphaGoal = 0.5
	t.alphaChange = t.alphaGoal / float64(t.numFrames)
}

func (t *transition) setFromShop() {
	t.numFrames = 40
	t.currentFrame = 0
	t.alphaChange = +0.01
	t.alpha = 0.5
	t.alphaGoal = 1
}

func (t *transition) update() (done bool) {
	t.currentFrame++
	done = t.currentFrame >= t.numFrames
	t.alpha += t.alphaChange
	if t.alphaChange < 0 && t.alpha < t.alphaGoal {
		t.alpha = t.alphaGoal
	} else if t.alphaChange > 0 && t.alpha > t.alphaGoal {
		t.alpha = t.alphaGoal
	}
	return
}

func (t transition) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1, 1, 1, t.alpha)
	screen.DrawImage(blackbgImage, op)
}
