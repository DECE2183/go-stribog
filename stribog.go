package stribog

import (
	"sync"

	bindings "github.com/dece2183/go-stribog/stribog_bindings"
)

const (
	BlockSize = 64
)

type Stribog struct {
	mux  sync.RWMutex
	data []byte
	size int
}

func (s *Stribog) BlockSize() int {
	return BlockSize
}

func (s *Stribog) Size() int {
	return s.size
}

func (s *Stribog) Reset() {
	s.mux.Lock()
	s.data = s.data[:0]
	s.mux.Unlock()
}

func (s *Stribog) Write(p []byte) (n int, err error) {
	s.mux.Lock()
	s.data = append(s.data, p...)
	s.mux.Unlock()
	return len(p), nil
}

func (s *Stribog) Sum(sum []byte) []byte {
	var out []byte

	s.mux.RLock()
	in := make([]byte, len(s.data))
	copy(in, s.data)
	s.mux.RUnlock()

	if s.size == 256/8 {
		out = bindings.Hash256(in)
	} else {
		out = bindings.Hash512(in)
	}

	return append(sum, out...)
}

func (s *Stribog) CheckSum(p []byte) []byte {
	if s.size == 256/8 {
		return bindings.Hash256(p)
	} else {
		return bindings.Hash512(p)
	}
}

func New256() *Stribog {
	return &Stribog{size: 256 / 8}
}

func New512() *Stribog {
	return &Stribog{size: 512 / 8}
}

func Sum256(p []byte) []byte {
	return bindings.Hash256(p)
}

func Sum512(p []byte) []byte {
	return bindings.Hash512(p)
}
