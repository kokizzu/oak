package main

import (
	"image/color"

	"github.com/oakmound/oak/v3/collision"

	"github.com/oakmound/oak/v3/physics"

	"github.com/oakmound/oak/v3/event"
	"github.com/oakmound/oak/v3/key"

	oak "github.com/oakmound/oak/v3"
	"github.com/oakmound/oak/v3/entities"
	"github.com/oakmound/oak/v3/render"
	"github.com/oakmound/oak/v3/scene"
)

// Collision labels
const (
	// The only collision label we need for this demo is 'ground',
	// indicating something we shouldn't be able to fall or walk through
	Ground collision.Label = 1
)

func main() {
	oak.AddScene("platformer", scene.Scene{Start: func(*scene.Context) {

		char := entities.NewMoving(100, 100, 16, 32,
			render.NewColorBox(16, 32, color.RGBA{255, 0, 0, 255}),
			nil, 0, 0)

		render.Draw(char.R)

		char.Speed = physics.NewVector(3, 3)

		fallSpeed := .1

		char.Bind(event.Enter, func(id event.CID, nothing interface{}) int {
			char := event.GetEntity(id).(*entities.Moving)
			// Move left and right with A and D
			if oak.IsDown(key.A) {
				char.ShiftX(-char.Speed.X())
			}
			if oak.IsDown(key.D) {
				char.ShiftX(char.Speed.X())
			}
			oldY := char.Y()
			char.ShiftY(char.Delta.Y())
			hit := collision.HitLabel(char.Space, Ground)

			// If we've moved in y value this frame and in the last frame,
			// we were below what we're trying to hit, we are still falling
			if hit != nil && !(oldY != char.Y() && oldY+char.H > hit.Y()) {
				// Correct our y if we started falling into the ground
				char.SetY(hit.Y() - char.H)
				char.Delta.SetY(0)
				// Jump with Space
				if oak.IsDown(key.Spacebar) {
					char.Delta.ShiftY(-char.Speed.Y())
				}
			} else {
				// Fall if there's no ground
				char.Delta.ShiftY(fallSpeed)
			}
			return 0
		})

		ground := entities.NewSolid(0, 400, 500, 20,
			render.NewColorBox(500, 20, color.RGBA{0, 0, 255, 255}),
			nil, 0)
		ground.UpdateLabel(Ground)

		render.Draw(ground.R)

	}})
	oak.Init("platformer")
}
