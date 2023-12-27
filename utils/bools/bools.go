package bools

// NullableBool is an integer-based nullable (ternary) boolean representation
// See the Nullable, False, and True constants for the convenience values
type NullableBool int
const (
	// Nullable represents a null boolean value (value is -1)
	Nullable = iota - 1
	// False represnts a false boolean value (value is 0)
	False = iota - 1
	// True represents a true boolean value (value is 1)
	True = iota - 1
)
