package ds

import (
	"testing"
)

func TestRingSliceWrite(t *testing.T) {
	t.Run("Write to empty RingSlice", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		if r.Get(0) != 1 {
			t.Errorf("expected 1, got %d", r.Get(0))
		}
	})

	t.Run("Write to partially filled RingSlice", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		r.Write(2)
		if r.Get(0) != 1 || r.Get(1) != 2 {
			t.Errorf("expected [1, 2], got [%d, %d]", r.Get(0), r.Get(1))
		}
	})

	t.Run("Write to full RingSlice (overwrite)", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		r.Write(2)
		r.Write(3)
		r.Write(4) // Overwrites the first element
		if r.Get(0) != 2 || r.Get(1) != 3 || r.Get(2) != 4 {
			t.Errorf("expected [2, 3, 4], got [%d, %d, %d]", r.Get(0), r.Get(1), r.Get(2))
		}
	})

	t.Run("Write updates readPointer when full", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		r.Write(2)
		r.Write(3)
		r.Write(4)
		if r.readPointer != r.writePointer {
			t.Errorf("expected readPointer to equal writePointer, got readPointer=%d, writePointer=%d", r.readPointer, r.writePointer)
		}
	})
}
func TestRingSliceBinarySearch(t *testing.T) {
	t.Run("BinarySearch in empty RingSlice", func(t *testing.T) {
		r := NewRingSlice[int](3)
		index := r.BinarySearch(1)
		if index != 0 {
			t.Errorf("expected 0, got %d", index)
		}
	})

	t.Run("BinarySearch in partially filled RingSlice", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		r.Write(3)
		r.Write(5)
		index := r.BinarySearch(3)
		if index != 1 {
			t.Errorf("expected 1, got %d", index)
		}
	})

	t.Run("BinarySearch in full RingSlice with wrap-around", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(3)
		r.Write(5)
		r.Write(7)
		r.Write(9)
		index := r.BinarySearch(5)
		if r.Get(index) == 0 {
			t.Errorf("expected 0, got %d", index)
		}
	})
}

func TestRingSliceDo(t *testing.T) {
	t.Run("Do on empty RingSlice", func(t *testing.T) {
		r := NewRingSlice[int](3)
		sum := 0
		r.Do(func(x int) {
			sum += x
		})
		if sum != 0 {
			t.Errorf("expected sum 0, got %d", sum)
		}
	})

	t.Run("Do on partially filled RingSlice", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		r.Write(2)
		sum := 0
		r.Do(func(x int) {
			sum += x
		})
		if sum != 3 {
			t.Errorf("expected sum 3, got %d", sum)
		}
	})

	t.Run("Do on full RingSlice with wrap-around", func(t *testing.T) {
		r := NewRingSlice[int](3)
		r.Write(1)
		r.Write(2)
		r.Write(3)
		r.Write(4) // Overwrites the first element
		sum := 0
		r.Do(func(x int) {
			sum += x
		})
		if sum != 9 {
			t.Errorf("expected sum 9, got %d", sum)
		}
	})
}
