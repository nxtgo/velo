package transformer

import (
	"image"
	"math"
	"strconv"
	"strings"
)

func resizeTransform(img image.Image, arg string) (image.Image, error) {
	parts := strings.Split(arg, "x")
	if len(parts) != 2 {
		return img, nil
	}
	w, _ := strconv.Atoi(parts[0])
	h, _ := strconv.Atoi(parts[1])
	if w <= 0 || h <= 0 {
		return img, nil
	}

	srcBounds := img.Bounds()
	sw, sh := srcBounds.Dx(), srcBounds.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, w, h))

	xRatio := float64(sw) / float64(w)
	yRatio := float64(sh) / float64(h)

	for y := range h {
		sy := int(math.Min(float64(sh-1), math.Floor(float64(y)*yRatio)))
		for x := range w {
			sx := int(math.Min(float64(sw-1), math.Floor(float64(x)*xRatio)))
			dst.Set(x, y, img.At(srcBounds.Min.X+sx, srcBounds.Min.Y+sy))
		}
	}

	return dst, nil
}
