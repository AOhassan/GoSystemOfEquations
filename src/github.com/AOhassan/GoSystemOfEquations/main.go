package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

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


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello, World")
}

func httpRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	httpRequests()


	// Test Cases
	//eq1 := equation{1, 1, 1}
	//eq2 := equation{-1, 1, 0}
	//
	//intersectionPoint, error := findIntersection(eq1, eq2)
	//
	//if error != nil {
	//	fmt.Println(error)
	//} else {
	//	fmt.Println(intersectionPoint)
	//}
	//
	//intersectionPoint0, error0 := newtonsMethod(eq1, eq2)
	//
	//if error0 != nil {
	//	fmt.Println(error0)
	//} else {
	//	fmt.Println(intersectionPoint0)
	//}
	//
	//eq3 := equation{1, 1, 1}
	//eq4 := equation{1, 1, 1}
	//
	//intersectionPoint2, error2 := findIntersection(eq3, eq4)
	//
	//if error2 != nil {
	//	fmt.Println(error2)
	//} else {
	//	fmt.Println(intersectionPoint2)
	//}
	//
	//eq5 := equation{1, 1, 1}
	//eq6 := equation{1, 1, 6}
	//
	//intersectionPoint3, error3 := findIntersection(eq5, eq6)
	//
	//if error3 != nil {
	//	fmt.Println(error3)
	//} else {
	//	fmt.Println(intersectionPoint3)
	//}
}