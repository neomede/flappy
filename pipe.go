package main

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	minPipeSize = 100
	gap         = 150
)

type pipes struct {
	pipes       []*pipe
	speed       int32
	currentPipe int
	texture     *sdl.Texture
}

func (pp *pipes) update() {
	for _, p := range pp.pipes {
		p.x -= pp.speed
		if p.x < -p.width {
			p.update()
		}
	}
}

func (pp *pipes) draw(r *sdl.Renderer) {
	for _, p := range pp.pipes {
		rectBottom := &sdl.Rect{
			X: p.x,
			Y: p.upHeight + p.gap,
			W: p.width,
			H: windowHeight - p.upHeight + p.gap,
		}
		r.CopyEx(pp.texture, nil, rectBottom, 0, nil, sdl.FLIP_NONE)

		rectUp := &sdl.Rect{
			X: p.x,
			Y: 0,
			W: p.width,
			H: p.upHeight,
		}
		r.CopyEx(pp.texture, nil, rectUp, 0, nil, sdl.FLIP_VERTICAL)
	}
}

func (pp *pipes) hits(b *bird) bool {
	for _, p := range pp.pipes {
		if p.hits(b) {
			return true
		}
	}
	return false
}

type pipe struct {
	x        int32
	width    int32
	upHeight int32
	gap      int32
	scored   bool
}

func newPipe(x int32) *pipe {
	return &pipe{
		x:        x,
		width:    52,
		upHeight: upHeight(),
		gap:      gap,
	}
}

func upHeight() int32 {
	return int32(rand.Intn(windowHeight-minPipeSize-gap) + minPipeSize)
}

func (p *pipe) update() {
	p.x = windowWidth + minPipeDist
	p.upHeight = upHeight()
	p.scored = false
}

func (p *pipe) hits(b *bird) bool {
	if b.x+b.w <= p.x || b.x >= p.x+p.width {
		return false
	}

	if !p.scored {
		b.score++
		p.scored = true
	}

	return b.y <= p.upHeight || b.y+b.h >= p.upHeight+p.gap
}
