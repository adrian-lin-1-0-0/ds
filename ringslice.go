package ds

import (
	"cmp"
	"sort"
)

type RingSlice[T cmp.Ordered] struct {
	size         int
	data         []T
	readPointer  int
	writePointer int
	full         bool
}

func NewRingSlice[T cmp.Ordered](size int) *RingSlice[T] {
	return &RingSlice[T]{
		size: size,
		data: make([]T, size),
	}
}

func (r *RingSlice[T]) Write(v T) {
	r.data[r.writePointer] = v
	r.writePointer = (r.writePointer + 1) % r.size
	if r.writePointer == 0 {
		r.full = true
	}
	if r.full {
		r.readPointer = r.writePointer
	}
}

func (r *RingSlice[T]) BinarySearch(target T) int {
	n := r.writePointer - r.readPointer
	if r.full {
		n = r.size
	}
	return sort.Search(n, func(i int) bool {
		return r.Get(i) >= target
	})
}

func (r *RingSlice[T]) Get(index int) T {
	return r.data[(r.readPointer+index)%r.size]
}

func (r *RingSlice[T]) Do(f func(x T)) {
	for i := 0; i < r.size; i++ {
		f(r.Get(i))
	}
}
