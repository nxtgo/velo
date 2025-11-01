package transformer

import (
	"image"
	"image/color"
)

func grayscaleTransform(img image.Image, _ string) (image.Image, error) {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			gray := uint8((r*299 + g*587 + b*114 + 500) / 1000 >> 8)
			dst.Set(x, y, color.NRGBA{gray, gray, gray, uint8(a >> 8)})
		}
	}

	return dst, nil
}
