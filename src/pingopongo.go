package main

import (
	"fmt"
	"time"
)

type PingoPongo struct {
	ping chan string
	pong chan string
}

func NewPingoPongo() *PingoPongo {
	return &PingoPongo{
		ping: make(chan string, 1),
		pong: make(chan string, 1),
	}
}

func (pp *PingoPongo) Start() {
	go pp.run()
	pp.ping <- "ping"
}

func (pp *PingoPongo) run() {
	for {
		select {
		case msg1 := <-pp.ping:
			fmt.Println(msg1)
			time.Sleep(1 * time.Second)
			pp.pong <- "pong"
		case msg2 := <-pp.pong:
			fmt.Println(msg2)
			time.Sleep(1 * time.Second)
			pp.ping <- "ping"
		}
	}
}

func main() {
	game := NewPingoPongo()
	game.Start()

	select {}
}
