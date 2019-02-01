package main

import (
	"errors"
)

func main() {

}

// contains the x/y points
type point struct {
	x float64
	y float64
}

// holds the equation coefficients
type equation struct {
	x, y, c float64
}

// Finding the intersection using Advanced Linear Algebra, check README.md
func findIntersection(eq1, eq2 equation) (point, error) {
	if eq1.x == eq2.x && eq1.y == eq2.y && eq1.c == eq2.c {
		return point{0,0}, errors.New("lines are the same, infinite solutions")
	}
	if eq1.x/eq1.y == eq2.x/eq2.y {
		return point{0,0}, errors.New("lines are parallel, no solutions")
	}

	inverse := eq2.y * eq1.c - eq1.y * eq2.c
	determinate := eq2.y * eq1.x - eq1.y * eq2.x

	x :=  inverse / determinate

	line1 := (eq1.c - eq1.x * x ) / eq1.y

	y := line1

	return point{x,y}, nil
}

// insufficient algo, bad approximations and/or infinite values, due to initial guess X_0
func newtonsMethod(eq1, eq2 equation) (point, error) {
	if eq1.x == eq2.x && eq1.y == eq2.y && eq1.c == eq2.c {
		return point{0,0}, errors.New("lines are the same, infinite solutions #Newton")
	}
	if eq1.x/eq1.y == eq2.x/eq2.y || -(eq1.x/ eq1.y + eq2.x/eq2.y) == 0 {
		return point{0,0}, errors.New("lines are parallel, no solutions #Newton")
	}

	// X_n+1 = X_n - f(X_n)/f'(X_n)
	// where f(X_n) = line1 - line2,
	// and f'(X_n) = slope1 - slope2
	// starting X_n at 0
	//line1 = eq1.c/ eq1.y - eq1.x/ eq1.y
	//line2 = eq2.c/eq2.y - eq2.x/eq2.y

	x := 0.0

	for i := 0; i < 100; i++ {
		f := (eq1.c/ eq1.y + eq2.c/eq2.y) - (eq1.x/ eq1.y + eq2.x/eq2.y)
		fPrime := -(eq1.x/ eq1.y + eq2.x/eq2.y)
		x = x - f/fPrime
	}

	line1 := (eq1.c - eq1.x * x ) / eq1.y
	y := line1

	return point{x,y}, nil
}