package main

import (
	"fmt"
	"math/rand"
)

func main() {
    var a = rand.Float64();
	var b = rand.Float64();
	var minimumResult = 10.0;
	var maximumResult = 11.0;
	var condition int = 1000;
	index := 1

	for {
		index++
		if (((10 * a) + (44 * b)) / (a + b) >= minimumResult && ((10 * a) + (44 * b)) / (a + b) <= maximumResult) {
			fmt.Printf("expression is equal in round %d", index);
			fmt.Println();
			fmt.Println("A is: ", a);
			fmt.Println();
			fmt.Println("B is: ", b);
			fmt.Println();
			break;
		}
		if (index > condition) {
			condition += 1000;
			minimumResult -= 0.25;
			maximumResult += 0.25;
		}
		a = rand.Float64();
		b = rand.Float64();
	}

	/*
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println()
	fmt.Println(rand.Float64())
    fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println();
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()
	s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100))
    fmt.Print(r3.Intn(100), ",")
	fmt.Print(rand.Intn(100), ",")
	*/
}