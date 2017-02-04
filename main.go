package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	err = ttf.Init()
	if err != nil {
		return fmt.Errorf("could not initialize TTF: %v", err)
	}
	defer ttf.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer w.Destroy()

	err = drawTitle(r)
	if err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}
	sdl.Delay(5000)

	s, err := newScene(r, 10, 1)
	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}
	go s.run(context.Background(), 50)

loop:
	for {
		switch event := sdl.WaitEvent().(type) {
		case *sdl.QuitEvent:
			break loop
		case *sdl.KeyUpEvent, *sdl.MouseButtonEvent:
			if s.bird.dead {
				s.restart()
				continue
			}
			s.bird.jump()
		default:
			log.Printf("igoring event of type %T", event)
		}
	}

	return nil
}
