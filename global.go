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

const (
	screenWidth  = 270
	screenHeight = 480
	fieldWidth   = 200
	fieldStart   = (screenWidth - fieldWidth) / 2
	fieldEnd     = fieldStart + fieldWidth

	maxAngle = -math.Pi / 7
	minAngle = -math.Pi - maxAngle

	minSpeed = 0.5

	wheatConversionRate = 1000

	spriteSize    = 32
	bladeHeight   = 10.0
	fieldTileSize = 32
	digitTileSize = 24

	startPositionX = screenWidth / 2
	startPositionY = screenHeight - screenHeight/8
	goalPositionY  = screenHeight / 2

	harvesterAnimationFrames = 10
	harvesterAnimationSteps  = 2
)
