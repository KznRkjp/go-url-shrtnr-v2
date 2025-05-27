package urlgen

import (
	"math/rand"
	"sync"
	"time"
)

const (
	charset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyLength = 8
)

var (
	rng  = rand.New(rand.NewSource(time.Now().UnixNano()))
	lock sync.Mutex
)

// GenerateShortKey - генерирует случайную строку длиной в 8 символов
func GenerateShortKey() string {
	lock.Lock()
	defer lock.Unlock()
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rng.Intn(len(charset))]
	}
	return string(shortKey)
}
