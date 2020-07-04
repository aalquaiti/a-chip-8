// Package chip provide functionality for the chip-8 interpreter
package chip

import "github.com/hajimehoshi/ebiten"

// Logical Screen Width
var screenWidth int

// Logical Screen Height
var screenHeight int

// Reference to the screen
// Only one screen will exist within the process
var screen *Screen

// ebiten Game Engine
var g *game

// Update function that handles interpreter logic every tick
var update func()

// Screen is a monochrome display
type Screen struct{}

// pixels matrix that is used to render screen
var pixels []byte

// NewScreen creates a new Screen if no previous screen was created. It ensures
// that only one screen exists (singelton).
// width: logical Screen width
// height: logcial screen height
// windowWidth: The displayed window width
// windowHeight: The displayed window height
// updatef: Update function that handles logic and is called every tick
func NewScreen(width, height, windowWidth, windowHeight int,
	updateF func()) *Screen {
	if screen != nil {
		return screen
	}

	screen = new(Screen)
	update = updateF
	g = &game{}

	screenWidth = width
	screenHeight = height
	pixels = make([]byte, screenWidth*screenHeight*4)

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("A Chip 8")

	return screen
}

// Show the window and start main loop
func (s *Screen) Show() {
	ebiten.RunGame(g)
}

// Reset clears the screen
func (s *Screen) Reset() {
	pixels = make([]byte, screenWidth*screenHeight*4)
}

// SetPixels update the pixels matrix
func (s *Screen) SetPixels(pix []byte) {
	pixels = pix
}

// PixelAt retrieve if pixel is writen at given coordinates
func (s *Screen) PixelAt(x, y int) bool {
	return pixels[(x*4)+(y*screenWidth*4)] == 1
}

// Draw a white pixel at given position or erase it
func (s *Screen) Draw(x, y int, write bool) {
	pos := (x * 4) + (y * screenWidth * 4)
	var value byte = 0xFF
	if !write {
		value = 0
	}
	pixels[pos] = value
	pixels[pos+1] = value
	pixels[pos+2] = value
	pixels[pos+3] = value
}

// Game loop used for ebiten game engine
type game struct {
}

func (g *game) Update(screen *ebiten.Image) error {
	update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(pixels)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
