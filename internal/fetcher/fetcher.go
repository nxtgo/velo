package fetcher

import (
	"bytes"
	"errors"
	"image"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"velo/internal/config"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			ForceAttemptHTTP2:  false,
			DisableCompression: true,
		},
	}
}

func Fetch(src string, cfg config.Config) (image.Image, error) {
	u, err := url.Parse(src)
	if err != nil {
		return nil, err
	}

	allowed := false
	for _, re := range cfg.DomainWhitelist {
		if re.MatchString(u.Host) {
			allowed = true
			break
		}
	}
	if !allowed {
		return nil, errors.New("fetch: domain not allowed")
	}

	req, _ := http.NewRequest("GET", src, nil)
	req.Header.Set("user-agent", "mozilla/5.0 (compatible; velo)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "image/") {
		return nil, errors.New("fetch: not an image")
	}

	reader := io.LimitReader(resp.Body, cfg.MaxImageSize)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}
