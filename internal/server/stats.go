package server

import (
	"os"
	"path/filepath"
	"sync/atomic"
	"time"
)

type Stats struct {
	StartTime     time.Time
	TotalRequests uint64
	CacheHits     uint64
	CacheMisses   uint64
	DiskItems     uint64
	DiskSize      uint64
}

func NewStats() *Stats {
	return &Stats{
		StartTime: time.Now(),
	}
}

func (s *Stats) IncRequests() {
	atomic.AddUint64(&s.TotalRequests, 1)
}

func (s *Stats) IncCacheHit() {
	atomic.AddUint64(&s.CacheHits, 1)
}

func (s *Stats) IncCacheMiss() {
	atomic.AddUint64(&s.CacheMisses, 1)
}

func (s *Stats) Snapshot() map[string]any {
	return map[string]any{
		"requests":       atomic.LoadUint64(&s.TotalRequests),
		"cache_hits":     atomic.LoadUint64(&s.CacheHits),
		"cache_misses":   atomic.LoadUint64(&s.CacheMisses),
		"uptime_seconds": uint64(time.Since(s.StartTime).Seconds()),
		"disk_items":     atomic.LoadUint64(&s.DiskItems),
		"disk_size":      float64(atomic.LoadUint64(&s.DiskSize)) / (1024 * 1024),
	}
}

func (s *Stats) startDiskMonitor(cacheDir string) {
	for {
		var count uint64
		var totalSize uint64

		filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if !d.IsDir() {
				count++
				if info, err := d.Info(); err == nil {
					totalSize += uint64(info.Size())
				}
			}
			return nil
		})

		atomic.StoreUint64(&s.DiskItems, count)
		atomic.StoreUint64(&s.DiskSize, totalSize)

		time.Sleep(30 * time.Second)
	}
}
