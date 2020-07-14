package gobuf

import "math"

type Grow func(size, want int) int

func FixedGrow(fixedSize int) Grow {
	return func(size, want int) int {
		newSize := size
		for newSize < want {
			newSize += fixedSize
		}
		return newSize
	}
}

func MultiplyGrow(factor float64) Grow {
	return func(size, want int) int {
		newSize := size
		for newSize < want {
			newSize = int(math.Ceil(float64(newSize) * factor))
		}
		return newSize
	}
}
