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
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

//go:embed assets/moissonneuse-corps1.png
var mc1 []byte

//go:embed assets/moissonneuse-corps2.png
var mc2 []byte

var moissImages [2]*ebiten.Image

//go:embed assets/moissonneuse-lame-petite1.png
var mlp1 []byte

//go:embed assets/moissonneuse-lame-petite2.png
var mlp2 []byte

var moissLamePetiteImages [2]*ebiten.Image

func loadAssets() {
	var err error

	decoded, _, err := image.Decode(bytes.NewReader(mc1))
	if err != nil {
		log.Fatal(err)
	}
	moissImages[0] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mc2))
	if err != nil {
		log.Fatal(err)
	}
	moissImages[1] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlp1))
	if err != nil {
		log.Fatal(err)
	}
	moissLamePetiteImages[0] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlp2))
	if err != nil {
		log.Fatal(err)
	}
	moissLamePetiteImages[1] = ebiten.NewImageFromImage(decoded)
}
