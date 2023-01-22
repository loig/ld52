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
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

//go:embed assets/start.mp3
var soundStartBytes []byte
var soundStart []byte

//go:embed assets/menuexit.mp3
var soundMenuExitBytes []byte
var soundMenuExit []byte

//go:embed assets/buy.mp3
var soundBuyBytes []byte
var soundBuy []byte

//go:embed assets/menubutton.mp3
var soundMenuButtonBytes []byte
var soundMenuButton []byte

//go:embed assets/gas.mp3
var soundGasBytes []byte
var soundGas []byte

//go:embed assets/nitro.mp3
var soundNitroBytes []byte
var soundNitro []byte

//go:embed assets/stone.mp3
var soundStoneBytes []byte
var soundStone []byte

//go:embed assets/firework.mp3
var soundFireworkBytes []byte
var soundFirework []byte

//go:embed assets/rebound.mp3
var soundReboundBytes []byte
var soundRebound []byte

//go:embed assets/missbuy.mp3
var soundMissBuyBytes []byte
var soundMissBuy []byte

//go:embed assets/music1.mp3
var music1Bytes []byte
var music1 *audio.InfiniteLoop

//go:embed assets/music2.mp3
var music2Bytes []byte
var music2 *audio.InfiniteLoop

//go:embed assets/music3.mp3
var music3Bytes []byte
var music3 *audio.InfiniteLoop
