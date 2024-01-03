package main

import (
	"fmt"
	"math"
)

/*
*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ a t^2 + vo*t + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))

Submit your Go program source code.
*/
func main() {
	a, vo, so, t := 0.0, 0.0, 0.0, 0.0
	fmt.Printf("insert the Aceleration \n")
	_, _ = fmt.Scanf("%f%*c", &a)
	fmt.Printf("insert the Inital Velocity \n")
	_, _ = fmt.Scanf("%f%*c", &vo)
	fmt.Printf("insert the Intial displacement \n")
	_, _ = fmt.Scanf("%f%*c", &so)
	fn := GenDisplaceFn(a, vo, so)
	fmt.Printf("insert the time \n")
	_, _ = fmt.Scanf("%f%*c", &t)
	fmt.Printf("%f", fn(t))
}

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vo*t + so
	}
	return fn
}
