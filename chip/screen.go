// Package chip provide functionality for the chip-8 interpreter
package chip

import "github.com/hajimehoshi/ebiten"

// Game loop used for ebiten game engine
type game struct {
	// Logical Screen Width
	width int

	// Logical Screen Height
	height int

	// pixels matrix that is used to render screen
	pixels []byte

	// Update function that logic every tick
	update func()
}

func (g *game) Update(screen *ebiten.Image) error {
	g.update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.pixels)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

// Screen is a monochrome display
type Screen struct {

	// ebiten Game Engine
	g *game
}

// NewScreen creates a new Screen if no previous screen was created. It ensures
// that only one screen exists (singelton).
// width: logical Screen width
// height: logcial screen height
// windowWidth: The displayed window width
// windowHeight: The displayed window height
// update: Update function that handles logic and is called every tick
func NewScreen(width, height, windowWidth, windowHeight int,
	update func()) (s *Screen) {

	g := &game{
		pixels: make([]byte, width*height*4),
		width:  width,
		height: height,
		update: update,
	}
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("A Chip 8")

	s = &Screen{g: g}

	return s
}

// Show the window and start main loop
func (s *Screen) Show() {
	ebiten.RunGame(s.g)
}

// Reset clears the screen
func (s *Screen) Reset() {
	s.g.pixels = make([]byte, s.g.width*s.g.height*4)
}

// SetPixels update the pixels matrix
func (s *Screen) SetPixels(pix []byte) {
	s.g.pixels = pix
}

// PixelAt retrieve if pixel is writen at given coordinates
func (s *Screen) PixelAt(x, y int) bool {
	return s.g.pixels[(x*4)+(y*s.g.width*4)] == 1
}

// Draw a white pixel at given position or erase it
func (s *Screen) Draw(x, y int, write bool) {
	pos := (x * 4) + (y * s.g.width * 4)
	var value byte = 0xFF
	if !write {
		value = 0
	}
	s.g.pixels[pos] = value
	s.g.pixels[pos+1] = value
	s.g.pixels[pos+2] = value
	s.g.pixels[pos+3] = value
}
