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
	"image/color"
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

//go:embed assets/moissonneuse-lame-moyennea-1.png
var mlm1 []byte

//go:embed assets/moissonneuse-lame-moyennea-2.png
var mlm2 []byte

//go:embed assets/moissonneuse-lame-grande1.png
var mlg1 []byte

//go:embed assets/moissonneuse-lame-grande2.png
var mlg2 []byte

var moissLameImages [6]*ebiten.Image

//go:embed assets/field.png
var fieldBytes []byte

var fieldImage *ebiten.Image

//go:embed assets/ground.png
var groundBytes []byte

var groundImage *ebiten.Image

var fgImage *ebiten.Image
var bgImage *ebiten.Image

//go:embed assets/gas.png
var gasLogoBytes []byte
var gasLogoImage *ebiten.Image

//go:embed assets/wheat.png
var wheatLogoBytes []byte
var wheatLogoImage *ebiten.Image

//go:embed assets/jaugebas.png
var gasTankBottomBytes []byte

//go:embed assets/jaugemid.png
var gasTankMidBytes []byte

//go:embed assets/jaugehaut.png
var gasTankTopBytes []byte

//go:embed assets/jaugecontenu.png
var gasTankContentBytes []byte

var gasTankImages [4]*ebiten.Image

//go:embed assets/digits.png
var digitsBytes []byte
var digitsImage *ebiten.Image

//go:embed assets/stone.png
var stoneBytes []byte
var stoneImage *ebiten.Image

//go:embed assets/gascollect.png
var gasCollectBytes []byte
var gasCollectImage *ebiten.Image

//go:embed assets/nitro.png
var nitroBytes []byte
var nitroImage *ebiten.Image

//go:embed assets/speed.png
var speedLogoBytes []byte
var speedLogoImage *ebiten.Image

//go:embed assets/speedjaugecontenu.png
var speedValueBytes []byte
var speedValueImage *ebiten.Image

//go:embed assets/wheatbg.png
var wbgBytes []byte
var wbgImage *ebiten.Image

//go:embed assets/fondbouton.png
var bgbuttonBytes []byte
var bgbuttonImage *ebiten.Image

var blackbgImage *ebiten.Image

//go:embed assets/buttongastank.png
var butgtBytes []byte
var butgtImage *ebiten.Image

//go:embed assets/buttongascollect.png
var butgcBytes []byte
var butgcImage *ebiten.Image

//go:embed assets/buttongaseff.png
var butgeBytes []byte
var butgeImage *ebiten.Image

//go:embed assets/buttonnitrocollect.png
var butncBytes []byte
var butncImage *ebiten.Image

//go:embed assets/buttonnitroeff.png
var butneBytes []byte
var butneImage *ebiten.Image

//go:embed assets/buttonbladeeff.png
var butbeBytes []byte
var butbeImage *ebiten.Image

//go:embed assets/buttonspeed.png
var butseBytes []byte
var butseImage *ebiten.Image

//go:embed assets/buttonstonecollect.png
var butstcBytes []byte
var butstcImage *ebiten.Image

//go:embed assets/buttonstoneeff.png
var butsteBytes []byte
var butsteImage *ebiten.Image

//go:embed assets/shoplevel.png
var slevBytes []byte
var slevImage *ebiten.Image

//go:embed assets/shoplevelok.png
var slevOkBytes []byte
var slevOKImage *ebiten.Image

//go:embed assets/buttonout.png
var boutBytes []byte
var boutImage *ebiten.Image

//go:embed assets/jaugebg.png
var tankbgBytes []byte
var tankbgImage *ebiten.Image

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
	moissLameImages[0] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlp2))
	if err != nil {
		log.Fatal(err)
	}
	moissLameImages[1] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlm1))
	if err != nil {
		log.Fatal(err)
	}
	moissLameImages[2] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlm2))
	if err != nil {
		log.Fatal(err)
	}
	moissLameImages[3] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlg1))
	if err != nil {
		log.Fatal(err)
	}
	moissLameImages[4] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(mlg2))
	if err != nil {
		log.Fatal(err)
	}
	moissLameImages[5] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(fieldBytes))
	if err != nil {
		log.Fatal(err)
	}
	fieldImage = ebiten.NewImageFromImage(decoded)

	fgImage = ebiten.NewImage(screenWidth, screenHeight)

	decoded, _, err = image.Decode(bytes.NewReader(groundBytes))
	if err != nil {
		log.Fatal(err)
	}
	groundImage = ebiten.NewImageFromImage(decoded)

	bgImage = ebiten.NewImage(screenWidth, screenHeight)
	bgImage.Fill(color.RGBA{R: 106, G: 35, B: 0, A: 255})

	decoded, _, err = image.Decode(bytes.NewReader(gasLogoBytes))
	if err != nil {
		log.Fatal(err)
	}
	gasLogoImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(wheatLogoBytes))
	if err != nil {
		log.Fatal(err)
	}
	wheatLogoImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(gasTankBottomBytes))
	if err != nil {
		log.Fatal(err)
	}
	gasTankImages[0] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(gasTankMidBytes))
	if err != nil {
		log.Fatal(err)
	}
	gasTankImages[1] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(gasTankTopBytes))
	if err != nil {
		log.Fatal(err)
	}
	gasTankImages[2] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(gasTankContentBytes))
	if err != nil {
		log.Fatal(err)
	}
	gasTankImages[3] = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(digitsBytes))
	if err != nil {
		log.Fatal(err)
	}
	digitsImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(stoneBytes))
	if err != nil {
		log.Fatal(err)
	}
	stoneImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(gasCollectBytes))
	if err != nil {
		log.Fatal(err)
	}
	gasCollectImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(nitroBytes))
	if err != nil {
		log.Fatal(err)
	}
	nitroImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(speedLogoBytes))
	if err != nil {
		log.Fatal(err)
	}
	speedLogoImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(speedValueBytes))
	if err != nil {
		log.Fatal(err)
	}
	speedValueImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(wbgBytes))
	if err != nil {
		log.Fatal(err)
	}
	wbgImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(bgbuttonBytes))
	if err != nil {
		log.Fatal(err)
	}
	bgbuttonImage = ebiten.NewImageFromImage(decoded)

	blackbgImage = ebiten.NewImage(screenWidth, screenHeight)
	blackbgImage.Fill(color.RGBA{A: 255})

	decoded, _, err = image.Decode(bytes.NewReader(butgtBytes))
	if err != nil {
		log.Fatal(err)
	}
	butgtImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butgcBytes))
	if err != nil {
		log.Fatal(err)
	}
	butgcImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butgeBytes))
	if err != nil {
		log.Fatal(err)
	}
	butgeImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butncBytes))
	if err != nil {
		log.Fatal(err)
	}
	butncImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butneBytes))
	if err != nil {
		log.Fatal(err)
	}
	butneImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butbeBytes))
	if err != nil {
		log.Fatal(err)
	}
	butbeImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butseBytes))
	if err != nil {
		log.Fatal(err)
	}
	butseImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butstcBytes))
	if err != nil {
		log.Fatal(err)
	}
	butstcImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(butsteBytes))
	if err != nil {
		log.Fatal(err)
	}
	butsteImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(slevBytes))
	if err != nil {
		log.Fatal(err)
	}
	slevImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(slevOkBytes))
	if err != nil {
		log.Fatal(err)
	}
	slevOKImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(boutBytes))
	if err != nil {
		log.Fatal(err)
	}
	boutImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(tankbgBytes))
	if err != nil {
		log.Fatal(err)
	}
	tankbgImage = ebiten.NewImageFromImage(decoded)

}
