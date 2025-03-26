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
	return sort.Search(r.Len(), func(i int) bool {
		return r.Get(i) >= target
	})
}

func (r *RingSlice[T]) Get(index int) T {
	return r.data[(r.readPointer+index)%r.size]
}

func (r *RingSlice[T]) Do(f func(x T)) {
	if r.full {
		for i := r.readPointer; i < r.size; i++ {
			f(r.data[i])
		}
		for i := 0; i < r.readPointer; i++ {
			f(r.data[i])
		}
		return
	}

	for i := r.readPointer; i < r.writePointer; i++ {
		f(r.data[i])
	}
}

func (r *RingSlice[T]) Len() int {
	n := r.writePointer - r.readPointer
	if r.full {
		n = r.size
	}
	return n
}
