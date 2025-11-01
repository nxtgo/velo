package transformer

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"strconv"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Transformer func(img image.Image, arg string) (image.Image, error)

var registry = map[string]Transformer{
	"s":         resizeTransform,
	"size":      resizeTransform,
	"grayscale": grayscaleTransform,
	"gray":      grayscaleTransform,
}

func Process(img image.Image, params map[string]string) (data []byte, format string, err error) {
	quality := 100

	for key, val := range params {
		switch key {
		case "q":
			if v, e := strconv.Atoi(val); e == nil {
				quality = v
			}
			continue
		}

		if fn, ok := registry[key]; ok {
			if img, err = fn(img, val); err != nil {
				return nil, "unknown", err
			}
		}
	}

	var buf bytes.Buffer
	switch quality {
	case 100:
		format = "png"
		err = png.Encode(&buf, img)
	default:
		format = "jpeg"
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	}

	if err != nil {
		return nil, "unknown", err
	}

	return buf.Bytes(), format, nil
}
