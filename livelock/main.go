package main

import (
	"fmt"
	"time"
)

type hallway struct {
	gopherCrossing *gopher
}

func (s hallway) crossing() {
	fmt.Printf("%s has crossing!\n", s.gopherCrossing.name)
}

type gopher struct {
	name           string
	isNeedCrossing bool
}

func (g *gopher) walk(hallway *hallway, otherGopher *gopher) {
	for g.isNeedCrossing {
		if hallway.gopherCrossing != g {
			time.Sleep(1 * time.Millisecond * 100)
			continue
		}

		// if otherGopher.isNeedCrossing {
		// 	fmt.Printf("%s: please, You crossing first %s!\n", g.name, otherGopher.name)
		// 	hallway.gopherCrossing = otherGopher
		// 	continue
		// }

		hallway.crossing()
		g.isNeedCrossing = false
		fmt.Printf("%s: I have crossed the hallway, thanks you %s!\n", g.name, otherGopher.name)
		hallway.gopherCrossing = otherGopher
		return
	}
}

func main() {
	gopherPurple := &gopher{
		name:           "Gopher Purple",
		isNeedCrossing: true,
	}

	gopherGreen := &gopher{
		name:           "Gopher Green",
		isNeedCrossing: true,
	}

	hallway := &hallway{gopherCrossing: gopherPurple}

	go func() {
		gopherPurple.walk(hallway, gopherGreen)
	}()

	go func() {
		gopherGreen.walk(hallway, gopherPurple)
	}()

	time.Sleep(time.Second) // hack
}
