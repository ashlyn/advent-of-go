package maths

import (
	"math"
)

// Pow is a convenience function for ints using math.Pow
func Pow(base int, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

// Abs returns the absolute value of the input
func Abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

// Lcm calculates the least common multiple of two numbers
func Lcm(first int, second int) int {
	return (first * second) / Gcd(first, second)
}

// Gcd calculates the greatest common denominator of two numbers
func Gcd(first int, second int) int {
	var div int

	for i := 1; i <= first && i <= second; i++ {
		if first%i == 0 && second%i == 0 {
			div = i
		}
	}

	return div
}

// MaxInt returns the maximum value for `int`
func MaxInt() int {
	return int(^uint(0) >> 1)
}

// Divisors returns a list of all whole divisors of n
func Divisors(n int) []int {
	divisors := []int{}
	if n < 2 {
		return divisors
	}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	for i := range divisors {
		divisors = append(divisors, n/divisors[i])
	}
	return divisors
}

// IsPrime returns `true` if n is a prime number, otherwise false
func IsPrime(n int) bool {
	return len(PrimeFactors(n)) == 1
}

// PrimeFactors returns a map where the keys are the prime factors
// of n and the values are the powers
func PrimeFactors(n int) map[int]int {
	primes := map[int]int{}

	// factorize even component
	for n%2 == 0 {
		primes[2]++
		n = n / 2
	}

	// factorize any odd component
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			primes[i]++
			n = n / i
		}
	}

	// n is prime
	if n > 2 {
		primes[n]++
	}

	return primes
}

// PrimeFactorsSlice returns a slice of prime factors of n
func PrimeFactorsSlice(n int) []int {
	factors := []int{}
	for factor := range PrimeFactors(n) {
		factors = append(factors, factor)
	}
	return factors
}

// SumOfDivisors calculates the sum of all divisors of a number n
func SumOfDivisors(n int) int {
	factors := PrimeFactors(n)
	sumOfDivisors := 1
	for factor, power := range factors {
		sum := 0
		for i := 0; i <= power; i++ {
			sum += int(math.Pow(float64(factor), float64(i)))
		}
		sumOfDivisors *= sum
	}
	return sumOfDivisors
}

// Max is a convenience function for max with ints
func Max(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}

// Min is a convenience function for max with ints
func Min(i int, j int) int {
	if i <= j {
		return i
	}
	return j
}

// DegreesToRadians converts an angle in degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// RadiansToDegrees converts an angle in raidans to degrees
func RadiansToDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
