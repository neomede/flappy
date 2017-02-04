package main

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type bird struct {
	x, y    int32
	w, h    int32
	speed   float64
	gravity float64
	dead    bool
	frames  []*sdl.Texture
	frame   int
	mu      sync.Mutex
}

func (b *bird) update() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.y += int32(b.speed)

	b.speed += b.gravity
	if b.y > windowHeight {
		b.dead = true
	}
	b.frame = (b.frame + 1) % len(b.frames)
}

func (b *bird) draw(r *sdl.Renderer) {
	rect := &sdl.Rect{X: b.x, Y: b.y, W: b.w, H: b.h}
	r.Copy(b.frames[b.frame], nil, rect)
}

func (b *bird) jump() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.speed = -10
}
