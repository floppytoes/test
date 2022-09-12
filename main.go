package main

import (
	"fmt"
	//"digomodule/packageA"
)

func main() {
	fmt.Println("Hello, this is main!")

	//fmt.Println(packageA.x) - not exported
	// fmt.Println(packageA.Y)
	// fmt.Printf("%v %T\n",packageA.UntypedStringConstant, packageA.UntypedStringConstant)
	// fmt.Printf("%v %T\n",packageA.TypedStringConstant, packageA.TypedStringConstant)

	// packageA.PrintHello()

	// x := packageA.NewBloomString("foo")
	// y := packageA.NewBloomStringRecord1("foobar", 0)
	// z1 := packageA.NewBloomStringRecord1("foobar", 0)
	// z2 := packageA.NewBloomStringRecord2("foobar", 0)
	// z3 := packageA.NewBloomStringRecord3("foobar", 0)
	// y.EditBloomStringRecord("bacon")
	// y.EditBloomStringRecord("fritter")
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z1, z2, z3)

}

