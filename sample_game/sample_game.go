package main

import (
	"fmt"
	"termon.googlecode.com/hg"
	"time"
	"rand"
)

type Player struct {
	X, Y int
	Health, Level, Killed int
}

type Monster struct {
	X, Y int
}

var p = Player{0, 0, 100, 1, 0}
var monsters = []Monster{
	{10, 10},
	{20, 5},
	{70, 20},
	{69, 19},
	{50, 5},
}

func draw() {
	term.Clear()
	f := "Health: %d\tLevel: %d\tEnemies killed: %d\tPosition: (%d, %d)"
	term.AddAt(0, 0, fmt.Sprintf(f, p.Health, p.Level, p.Killed, p.X, p.Y))
	for x := 0; x < *term.Cols; x++ {
		for y := 1; y < *term.Rows; y++ {
			ch := '.'
			if x == p.X && y-1 == p.Y {
				ch = '*'
			} else {
				for _, m := range monsters {
					if x == m.X && y-1 == m.Y {
						ch = '$'
						break
					}
				}
			}
			term.AddAt(x, y, ch)
		}
	}
}

func listen() {
	for {
		c := term.GetChar()
		switch c {
			case term.KEY_UP:
				p.Y -= 1
			case term.KEY_DOWN:
				p.Y += 1
			case term.KEY_RIGHT:
				p.X += 1
			case term.KEY_LEFT:
				p.X -= 1
		}
		for i, _ := range monsters {
			m := &monsters[i]
			if p.X == m.X && p.Y == m.Y {
				m.X = 0
				m.Y = 0
				p.Killed++
				p.Level = int(p.Killed / 3.0) + 1
				break
			}
		}
		draw()
	}
}

func moveMonsters() {
	for {
		for i, _ := range monsters {
			m := &monsters[i]
			add := rand.Intn(3) - 1
			if rand.Intn(2) == 0 {
				m.X += add
			} else {
				m.Y += add
			}
		}
		time.Sleep(2000 * 1000 * 1000)
	}
}

func Start() {
	term.Init()
	term.Keypad()
	term.Noecho()
	term.HalfDelay(1)
	
	go moveMonsters()
	draw()
	listen()
	
	term.End()
}
