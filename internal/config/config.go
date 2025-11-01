package config

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	Addr            string
	CacheDir        string
	MaxImageSize    int64
	DomainWhitelist []*regexp.Regexp
}

func Load() Config {
	return Config{
		Addr:            env("VELO_ADDR", ":8080"),
		CacheDir:        env("VELO_CACHE_DIR", "store"),
		MaxImageSize:    envInt64("VELO_MAX_IMAGE_SIZE", 10<<20), //10mb default ok
		DomainWhitelist: parseWhitelist(env("VELO_WHITELISTED_DOMAINS", `^(?:i\.imgur\.com|cdn\.discordapp\.com)$`)),
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt64(key string, def int64) int64 {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			return n
		}
	}
	return def
}

func parseWhitelist(raw string) []*regexp.Regexp {
	var out []*regexp.Regexp
	for p := range strings.SplitSeq(raw, ",") {
		if re, err := regexp.Compile(strings.TrimSpace(p)); err == nil {
			out = append(out, re)
		}
	}
	return out
}
