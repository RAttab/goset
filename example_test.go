// Copyright (c) 2014 Datacratic. All rights reserved.

package set_test

import (
	"github.com/RAttab/goset"

	"fmt"
)

func Example_Basics() {

	x := set.NewString("a", "b", "c")

	fmt.Println("len(x):", len(x))
	fmt.Println("x.Test(a):", x.Test("a"))
	fmt.Println("x.Test(d):", x.Test("d"))

	y := x.Copy()
	y.Del("b")
	fmt.Println("y:", y)

	z := set.NewString()
	for v := range y {
		z.Put(v)
	}

	fmt.Println("z:", z.Array())

	// Output:
	// len(x): 3
	// x.Test(a): true
	// x.Test(d): false
	// y: { a c }
	// z: [a c]
}

func Example_Operands() {

	x := set.NewString("a", "b", "c")
	y := set.NewString("b", "c", "d")

	fmt.Println("x.Union(y):", x.Union(y))
	fmt.Println("x.Intersect(y):", x.Intersect(y))
	fmt.Println("x.Difference(y):", x.Difference(y))

	// Output:
	// x.Union(y): { a b c d }
	// x.Intersect(y): { b c }
	// x.Difference(y): { a }
}

func Example_Types() {

	unsigned := set.NewUint(uint64(1), uint64(2), uint64(3))
	fmt.Println("unsigned:", unsigned)

	signed := set.NewInt(int64(-1), int64(-2), int64(-3))
	fmt.Println("signed:", signed)

	// Output:
	// unsigned: { 1 2 3 }
	// signed: { -3 -2 -1 }
}
