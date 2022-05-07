package captcha

import (
	"io"

	gocaptcha "github.com/steambap/captcha"
)

// Captcha is a captcha generator.
type Captcha struct {
	core *gocaptcha.Data
}

// Config is the configuration for captcha.
type Config struct {
	// Length is the length of captcha text.
	Length int

	// BaseChars is the base chars to use.
	// Default: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// You can custom number only: "0123456789"
	BaseChars string

	// Noise is the number of noise dots.
	// Default: 1.0
	Noise float64

	// Witdh is the width of captcha image.
	// Default: 150px
	Width int

	// Height is the height of captcha image.
	// Default: 50px
	Height int
}

// New creates a new captcha.
func New(config ...*Config) *Captcha {
	width := 150
	height := 50
	textLength := 4
	baseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	noise := 1.0

	if len(config) > 0 {
		if cfg := config[0]; cfg != nil {
			if cfg.Length > 0 {
				textLength = cfg.Length
			}
			if cfg.BaseChars != "" {
				baseChars = cfg.BaseChars
			}
			if cfg.Noise > 0 {
				noise = cfg.Noise
			}
			if cfg.Width > 0 {
				width = cfg.Width
			}
			if cfg.Height > 0 {
				height = cfg.Height
			}
		}
	}

	core, err := gocaptcha.New(width, height, func(options *gocaptcha.Options) {
		options.TextLength = textLength
		options.CharPreset = baseChars
		options.Noise = noise
	})
	if err != nil {
		panic(err)
	}

	return &Captcha{
		core: core,
	}
}

// Text returns the captcha text.
func (c *Captcha) Text() string {
	return c.core.Text
}

// Write writes the captcha image to writer.
func (c *Captcha) Write(w io.Writer) error {
	return c.core.WriteImage(w)
}
