package server

import (
	"encoding/json"
	"fmt"

	"velo/internal/cache"
	"velo/internal/fetcher"
	"velo/internal/transformer"

	"github.com/valyala/fasthttp"
)

func (s *Server) HandleImage(ctx *fasthttp.RequestCtx) {
	s.stats.IncRequests()

	var paramsStr string
	if v := ctx.UserValue("params"); v != nil {
		paramsStr = v.(string)
	}

	src := string(ctx.QueryArgs().Peek("src"))
	if src == "" {
		ctx.Error("missing src", fasthttp.StatusBadRequest)
		return
	}

	params := parseParams(paramsStr)
	key := cache.Key(src, params)

	if data := cache.MemoryGet(key); data != nil {
		s.stats.IncCacheHit()
		ctx.SetContentType("image/jpeg")
		ctx.Write(data)
		return
	}

	s.stats.IncCacheMiss()

	img, err := fetcher.Fetch(src, s.cfg)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadGateway)
		return
	}

	out, format, err := transformer.Process(img, params)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	cache.MemorySet(key, out)
	cache.DiskSave(key, out, s.cfg.CacheDir)

	ctx.SetContentType(fmt.Sprintf("image/%s", format))
	ctx.Write(out)
}

func (s *Server) HandleStats(ctx *fasthttp.RequestCtx) {
	snapshot := s.stats.Snapshot()

	data, _ := json.MarshalIndent(snapshot, "", "  ")
	ctx.SetContentType("application/json")
	ctx.Write(data)
}
