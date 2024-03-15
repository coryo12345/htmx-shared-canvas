package canvas

import (
	"errors"
	"sync"
)

type Canvas struct {
	height int
	width  int
	pixels []Pixel
	mutex  sync.Mutex
}

type Pixel struct {
	R uint8
	G uint8
	B uint8
}

func New(height int, width int) *Canvas {
	size := height * width
	pixels := make([]Pixel, size)

	for i := range pixels {
		pixels[i].R = uint8((i*2)%50) + 200
		pixels[i].G = uint8((i*2)%50) + 200
		pixels[i].B = uint8((i*2)%50) + 200
	}

	canvas := Canvas{
		height: height,
		width:  width,
		pixels: pixels,
		mutex:  sync.Mutex{},
	}
	return &canvas
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) Pixels() []Pixel {
	c.mutex.Lock()
	pixelsCopy := make([]Pixel, len(c.pixels))
	copy(pixelsCopy, c.pixels)
	c.mutex.Unlock()
	return pixelsCopy
}

func (c *Canvas) GetPixel(pos int) Pixel {
	c.mutex.Lock()
	pixel := c.pixels[pos]
	c.mutex.Unlock()
	return pixel
}

func (c *Canvas) SetPixel(pos int, pixel Pixel) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if pos > len(c.pixels) || pos < 0 {
		return errors.New("provided position is outside canvas")
	}

	c.pixels[pos] = pixel
	return nil
}

func (c *Canvas) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	pixels := make([]Pixel, c.height*c.width)
	for i := range pixels {
		pixels[i] = Pixel{
			R: 255,
			G: 255,
			B: 255,
		}
	}
	c.pixels = pixels
}
