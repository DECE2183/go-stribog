package main

import (
	"fmt"
	"os"
	"sync"

	bindings "github.com/dece2183/go-stribog/stribog"
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

func New256() *Stribog {
	return &Stribog{size: 256 / 8}
}

func New512() *Stribog {
	return &Stribog{size: 512 / 8}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing message")
		os.Exit(1)
	}

	h := New256()
	h.Write([]byte(os.Args[1]))
	hash := h.Sum([]byte{})

	for i := 0; i < h.Size(); i++ {
		fmt.Printf("%X", hash[i])
	}
	fmt.Println()
}
