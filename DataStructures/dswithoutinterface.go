// _Interfaces_ are named collections of method
// signatures.

package main

import (
	"fmt"
 "math"
)

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect_noint struct {
    width, height float64
}
type circle_noint struct {
    radius float64
}

func (r rect_noint) area() float64 {
    return r.width * r.height
}
func (r rect_noint) perim() float64 {
    return 2*r.width + 2*r.height
}

func (r rect_noint) measure() {
    fmt.Println(r)
    fmt.Println(r.area())
    fmt.Println(r.perim())
}

// The implementation for `circle`s.
func (c circle_noint) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle_noint) perim() float64 {
    return 2 * math.Pi * c.radius
}

func (c circle_noint) measure() {
    fmt.Println(c)
    fmt.Println(c.area())
    fmt.Println(c.perim())
}

func noint() {
    r := rect_noint{width: 3, height: 4}
    c := circle_noint{radius: 5}
	r.measure()
	c.measure()
}
