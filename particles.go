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
	//"log"
	"math/rand"
)

type particle struct {
	x, y           float64
	vx, vy         float64
	a              float64
	va             float64
	r, g, b        float64
	life           int
	lifeTime       int
	moveWithScreen bool
}

type particleSys struct {
	content   []particle
	lastAlive int
}

func (p *particle) update(yHarvesterMove float64) (dead bool) {
	p.life++
	if p.lifeTime <= p.life {
		return true
	}
	p.x += p.vx
	p.y += p.vy
	if p.moveWithScreen {
		p.y -= yHarvesterMove
	}
	p.a -= p.va
	if p.a <= 0 {
		return true
	}
	return
}

func (ps *particleSys) draw(screen *ebiten.Image) {
	for i := 0; i <= ps.lastAlive; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ps.content[i].x, ps.content[i].y)
		op.ColorM.Scale(ps.content[i].r, ps.content[i].g, ps.content[i].b, ps.content[i].a)
		screen.DrawImage(particuleImg, op)
	}
}

func (ps *particleSys) update(yHarvesterMove float64) {
	for i := 0; i <= ps.lastAlive; i++ {
		if ps.content[i].update(yHarvesterMove) {
			if ps.lastAlive >= 0 {
				ps.content[i], ps.content[ps.lastAlive] = ps.content[ps.lastAlive], ps.content[i]
				ps.lastAlive--
				i--
			}
		}
	}
}

func (ps *particleSys) addParticle(p particle) {
	ps.lastAlive++
	if ps.lastAlive < len(ps.content) {
		ps.content[ps.lastAlive] = p
	} else {
		ps.content = append(ps.content, p)
	}
}

func (ps *particleSys) genHarvesterParticles(xLeft, xRight, yLeft, yRight float64, size float64, nitro bool) {
	a := (yRight - yLeft) / (xRight - xLeft)
	b := yLeft - a*xLeft
	num := rand.Intn(int(size/6)) + 1
	if nitro {
		num += rand.Intn(50) + 50
	}
	for i := 0; i < num; i++ {
		xx := xLeft + rand.Float64()*(xRight-xLeft)
		yy := a*xx + b
		rr := 0.9
		gg := 0.9
		bb := 0.3
		va := 0.01
		if nitro {
			rr = 0
			gg = 1
			bb = 0
			va = 0.1
			yy += rand.Float64() * 100
		}
		ps.addParticle(
			particle{
				x: xx, y: yy,
				vx: 0, vy: 0,
				a:  0.6,
				va: va,
				r:  rr, g: gg, b: bb,
				life:           0,
				lifeTime:       60,
				moveWithScreen: true},
		)
	}
	//log.Print(len(ps.content), ps.lastAlive, ps)
}

func (ps *particleSys) genCollectParticles(x, y float64, kind int) {
	num := 50
	for i := 0; i < num; i++ {
		xx := x + float64(rand.Intn(16)-8)
		yy := y + float64(rand.Intn(16)-8)
		rr := 0.9
		gg := 0.9
		bb := 0.0
		if kind == 1 {
			rr = 0
			gg = 1
		}
		if kind == 2 {
			rr = 0.3
			gg = 0.3
			bb = 0.3
		}
		ps.addParticle(
			particle{
				x: xx, y: yy,
				vx: 0, vy: 0,
				a:  0.6,
				va: 0.01,
				r:  rr, g: gg, b: bb,
				life:           0,
				lifeTime:       60,
				moveWithScreen: true},
		)
	}
	//log.Print(len(ps.content), ps.lastAlive, ps)
}

func (ps *particleSys) genConsumeParticles(x, y float64, kind int) {
	if rand.Intn(3) == 0 {
		xx := x + float64(rand.Intn(10))
		yy := y + float64(rand.Intn(3))
		rr := 1.0
		gg := 1.0
		bb := 0.2
		if kind == 1 {
			rr = 0
			gg = 0.65
			bb = 0.75
		}
		ps.addParticle(
			particle{
				x: xx, y: yy,
				vx: 0, vy: 0,
				a:  0.6,
				va: 0.01,
				r:  rr, g: gg, b: bb,
				life:           0,
				lifeTime:       60,
				moveWithScreen: false},
		)
	}
	//log.Print(len(ps.content), ps.lastAlive, ps)
}

func (ps *particleSys) genVictoryParticles() (playSound bool) {

	if rand.Intn(10) == 0 {
		playSound = true
		num := rand.Intn(50) + 40
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		x := rand.Float64() * screenWidth
		y := rand.Float64() * screenHeight
		for i := 0; i < num; i++ {
			ps.addParticle(
				particle{
					x: x, y: y,
					vx: 2 * (0.5 - rand.Float64()), vy: 2 * (0.5 - rand.Float64()),
					a:  1,
					va: 0.01,
					r:  r, g: g, b: b,
					life:           0,
					lifeTime:       120,
					moveWithScreen: false},
			)
		}
	}

	return

}
