package cache

import "sync"

var (
	memory = sync.Map{}
)

func MemoryGet(key string) []byte {
	if v, ok := memory.Load(key); ok {
		return v.([]byte)
	}
	return nil
}

func MemorySet(key string, data []byte) {
	memory.Store(key, data)
}
