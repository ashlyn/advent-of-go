package ranges

// Range is a range of integers, inclusive
type Range struct {
	Start int
	End   int
}

// New returns a new Range
func New(start int, end int) Range {
	return Range{start, end}
}

// NewWithLength returns a new Range with a given length
func NewWithLength(start int, length int) Range {
	return Range{start, start + length - 1}
}

// Length returns the length of the range
func (r Range) Length() int {
	return r.End - r.Start + 1
}

// Contains returns true if the range contains the given integer
func (r Range) Contains(i int) bool {
	return i >= r.Start && i <= r.End
}

// ContainsRange returns true if the range fully contains the given range
func (r Range) ContainsRange(r2 Range) bool {
	return r.Contains(r2.Start) && r.Contains(r2.End)
}

// Overlaps returns true if the range overlaps with the given range
func (r Range) Overlaps(r2 Range) bool {
	return r.Contains(r2.Start) || r.Contains(r2.End) || r2.Contains(r.Start) || r2.Contains(r.End)
}

// Iterator returns a slice of integers in the range
func (r Range) Iterator() []int {
	iterator := []int{}
	for i := r.Start; i <= r.End; i++ {
		iterator = append(iterator, i)
	}
	return iterator
}
