package circulararray

// CircularIndex calculates the new index in a circular array given the current index, increment, and length.
func CircularIndex(index int, increment int, length int) int {
	if increment < 0 {
		return circularBackwards(index, increment, length)
	}
	return circularForwards(index, increment, length)
}

func circularBackwards(index int, increment int, length int) int {
	return (index + (length + (increment % length))) % length
}

func circularForwards(index int, increment int, length int) int {
	return (index + increment) % length
}