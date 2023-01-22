/*
LD50, a game for Ludum Dare 50

	Copyright (C) 2022  Loïg Jezequel

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see https://www.gnu.org/licenses/.
*/
package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"io/ioutil"
	"log"
	"math"
	"time"
)

const (
	soundStartID int = iota
	soundMenuExitID
	soundBuyID
	soundMenuButtonID
	soundGasID
	soundNitroID
	soundStoneID
	soundFireworkID
	soundReboundID
	soundMissBuyID
)

const (
	music1ID int = iota
	music2ID
	music3ID
)

type soundManager struct {
	audioContext *audio.Context
	musicPlayers [3]*audio.Player
	currentMusic int
}

// loop the music
func (g *game) updateMusic(musicID int, volume float64) {

	if g.audio.musicPlayers[musicID] != nil {
		if musicID != g.audio.currentMusic {
			if g.audio.musicPlayers[g.audio.currentMusic] != nil {
				g.audio.musicPlayers[g.audio.currentMusic].Pause()
			}
		}

		if !g.audio.musicPlayers[musicID].IsPlaying() {
			if (musicID == music3ID && g.audio.currentMusic == music2ID) || (musicID == music2ID && g.audio.currentMusic == music3ID) {
				time := g.audio.musicPlayers[g.audio.currentMusic].Current()
				g.audio.musicPlayers[musicID].Rewind()
				error := g.audio.musicPlayers[musicID].Seek(time)
				if error != nil {
					log.Print(error)
				}
			} else {
				g.audio.musicPlayers[musicID].Rewind()
			}
			g.audio.musicPlayers[musicID].Play()
		}
		g.audio.musicPlayers[musicID].SetVolume(volume)
	}

	g.audio.currentMusic = musicID
}

// stop the music
func (g *game) stopMusic() {
	for i := 0; i < 3; i++ {
		if g.audio.musicPlayers[i] != nil && g.audio.musicPlayers[i].IsPlaying() {
			g.audio.musicPlayers[i].Pause()
		}
	}
}

// play a sound
func (g *game) playSound(sound int) {
	soundBytes := soundStart
	switch sound {
	case soundStartID:
		soundBytes = soundStart
	case soundMenuExitID:
		soundBytes = soundMenuExit
	case soundBuyID:
		soundBytes = soundBuy
	case soundMenuButtonID:
		soundBytes = soundMenuButton
	case soundGasID:
		soundBytes = soundGas
	case soundNitroID:
		soundBytes = soundNitro
	case soundStoneID:
		soundBytes = soundStone
	case soundFireworkID:
		soundBytes = soundFirework
	case soundReboundID:
		soundBytes = soundRebound
	case soundMissBuyID:
		soundBytes = soundMissBuy
	}
	soundPlayer := audio.NewPlayerFromBytes(g.audio.audioContext, soundBytes)
	soundPlayer.Play()
}

// decode music and sounds
func (g *game) initAudio() {

	var error error
	g.audio.audioContext = audio.NewContext(44100)

	// music
	sound, error := mp3.Decode(g.audio.audioContext, bytes.NewReader(music1Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("34s931ms")
	duration := tduration.Seconds()
	theBytes := int64(math.Round(duration * 4 * float64(44100)))
	music1 = audio.NewInfiniteLoop(sound, theBytes)
	g.audio.musicPlayers[0], error = audio.NewPlayer(g.audio.audioContext, music1)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(music2Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ = time.ParseDuration("24s006ms")
	duration = tduration.Seconds()
	theBytes = int64(math.Round(duration * 4 * float64(44100)))
	tIntroDuration, _ := time.ParseDuration("1s879ms")
	introDuration := tIntroDuration.Seconds()
	introTheBytes := int64(math.Round(introDuration * 4 * float64(44100)))

	music2 = audio.NewInfiniteLoopWithIntro(sound, introTheBytes, theBytes)
	g.audio.musicPlayers[1], error = audio.NewPlayer(g.audio.audioContext, music2)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(music3Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	music3 = audio.NewInfiniteLoopWithIntro(sound, introTheBytes, theBytes)
	g.audio.musicPlayers[2], error = audio.NewPlayer(g.audio.audioContext, music3)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	// sounds
	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundStartBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundStart, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundMenuExitBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMenuExit, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundBuyBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundBuy, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundMenuButtonBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMenuButton, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundGasBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundGas, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundNitroBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundNitro, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundStoneBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundStone, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundFireworkBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundFirework, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundReboundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundRebound, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(soundMissBuyBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundMissBuy, error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
}
