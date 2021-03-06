package noop

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/oakmound/oak/v3/shiny/driver/internal/event"
	"github.com/oakmound/oak/v3/shiny/screen"
	"golang.org/x/image/math/f64"
)

func Main(f func(screen.Screen)) {
	f(screenImpl{})
}

type screenImpl struct{}

func (screenImpl) NewImage(size image.Point) (screen.Image, error) {
	return imageImpl{
		size: size,
		rgba: image.NewRGBA(image.Rect(0, 0, size.X, size.Y)),
	}, nil
}

func (screenImpl) NewTexture(size image.Point) (screen.Texture, error) {
	return textureImpl{
		size: size,
	}, nil
}

func (screenImpl) NewWindow(opts screen.WindowGenerator) (screen.Window, error) {
	return &windowImpl{}, nil
}

type imageImpl struct {
	size image.Point
	rgba *image.RGBA
}

func (ii imageImpl) Size() image.Point {
	return ii.size
}

func (ii imageImpl) Bounds() image.Rectangle {
	return image.Rect(0, 0, ii.size.X, ii.size.Y)
}

func (imageImpl) Release() {}

func (ii imageImpl) RGBA() *image.RGBA {
	return ii.rgba
}

type textureImpl struct {
	size image.Point
}

func (ti textureImpl) Size() image.Point {
	return ti.size
}

func (ti textureImpl) Bounds() image.Rectangle {
	return image.Rect(0, 0, ti.size.X, ti.size.Y)
}

func (textureImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle) {}
func (textureImpl) Fill(dr image.Rectangle, src color.Color, op draw.Op)        {}
func (textureImpl) Release()                                                    {}

type windowImpl struct {
	event.Deque
}

func (*windowImpl) Release()                                                                      {}
func (*windowImpl) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op)     {}
func (*windowImpl) DrawUniform(src2dst f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op) {}
func (*windowImpl) Copy(dp image.Point, src screen.Texture, sr image.Rectangle, op draw.Op)       {}
func (*windowImpl) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op)  {}
func (*windowImpl) Upload(dp image.Point, src screen.Image, sr image.Rectangle)                   {}
func (*windowImpl) Fill(dr image.Rectangle, src color.Color, op draw.Op)                          {}

func (*windowImpl) Publish() screen.PublishResult {
	return screen.PublishResult{}
}
