package main

import (
	"fmt"

	oak "github.com/oakmound/oak/v3"
	"github.com/oakmound/oak/v3/event"
	"github.com/oakmound/oak/v3/key"
	"github.com/oakmound/oak/v3/render"
	"github.com/oakmound/oak/v3/scene"
)

const (
	borderlessAtStart = false
	fullscreenAtStart = false
	topMostAtStart    = false
)

func main() {
	oak.AddScene("demo", scene.Scene{Start: func(*scene.Context) {
		txt := render.NewText("Press F to toggle fullscreen. Press B to toggle borderless. Press T to toggle topmost", 50, 50)
		render.Draw(txt)

		borderless := borderlessAtStart
		fullscreen := fullscreenAtStart
		topMost := topMostAtStart

		event.GlobalBind(key.Down+key.F, func(event.CID, interface{}) int {
			fullscreen = !fullscreen
			fmt.Println("Setting fullscreen:", fullscreen)
			err := oak.SetFullScreen(fullscreen)
			if err != nil {
				fullscreen = !fullscreen
				fmt.Println(err)
			}
			return 0
		})
		event.GlobalBind(key.Down+key.B, func(event.CID, interface{}) int {
			borderless = !borderless
			fmt.Println("Setting borderless:", borderless)
			err := oak.SetBorderless(borderless)
			if err != nil {
				borderless = !borderless
				fmt.Println(err)
			}
			return 0
		})
		event.GlobalBind(key.Down+key.T, func(event.CID, interface{}) int {
			topMost = !topMost
			fmt.Println("Setting top most:", topMost)
			err := oak.SetTopMost(topMost)
			if err != nil {
				topMost = !topMost
				fmt.Println(err)
			}
			return 0
		})

	}})

	oak.Init("demo", func(c oak.Config) (oak.Config, error) {
		c.TopMost = topMostAtStart
		// Both cannot be true at once!
		c.Borderless = borderlessAtStart
		c.Fullscreen = fullscreenAtStart
		return c, nil
	})
}
