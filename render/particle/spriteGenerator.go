package particle

import (
	"github.com/oakmound/oak/v3/alg/range/floatrange"

	"github.com/oakmound/oak/v3/alg"
	"github.com/oakmound/oak/v3/render"
)

// A SpriteGenerator generate SpriteParticles
type SpriteGenerator struct {
	BaseGenerator
	SpriteRotation floatrange.Range
	Base           *render.Sprite
}

// NewSpriteGenerator creates a SpriteGenerator
func NewSpriteGenerator(options ...func(Generator)) Generator {
	g := new(SpriteGenerator)
	g.setDefaults()

	for _, opt := range options {
		opt(g)
	}

	return g
}

func (sg *SpriteGenerator) setDefaults() {
	sg.BaseGenerator.setDefaults()
	sg.SpriteRotation = floatrange.NewConstant(0)
}

// Generate creates a source using this generator
func (sg *SpriteGenerator) Generate(layer int) *Source {
	// Convert rotation from degrees to radians
	if sg.Rotation != nil {
		sg.Rotation = sg.Rotation.Mult(alg.DegToRad)
	}
	return NewSource(sg, layer)
}

// GenerateParticle creates a particle from a generator
func (sg *SpriteGenerator) GenerateParticle(bp *baseParticle) Particle {
	return &SpriteParticle{
		baseParticle: bp,
		rotation:     float32(sg.SpriteRotation.Poll()),
	}
}

// A Sprited can have a sprite set to it
type Sprited interface {
	SetSprite(*render.Sprite)
	SetSpriteRotation(f floatrange.Range)
}

// Sprite sets a Sprited's sprite
func Sprite(s *render.Sprite) func(Generator) {
	return func(g Generator) {
		g.(Sprited).SetSprite(s)
	}
}

// SetSprite is the function on a sprite generator that satisfies
// Sprited
func (sg *SpriteGenerator) SetSprite(s *render.Sprite) {
	sg.Base = s
}

// SpriteRotation sets a Sprited's rotation
func SpriteRotation(f floatrange.Range) func(Generator) {
	return func(g Generator) {
		g.(Sprited).SetSpriteRotation(f)
	}
}

// SetSpriteRotation satisfied Sprited for SpriteGenerators
func (sg *SpriteGenerator) SetSpriteRotation(f floatrange.Range) {
	sg.SpriteRotation = f
}

// GetParticleSize returns the size of the sprite that the generator generates
func (sg *SpriteGenerator) GetParticleSize() (w float64, h float64, perParticle bool) {
	bounds := sg.Base.GetRGBA().Rect.Max
	return float64(bounds.X), float64(bounds.Y), false
}
