package main

import (
	"fmt"
	"math/rand"
)

func main() {
	calculateMatchmaking(10, 44, 10, 11, "10a + 44b / a+b ≈ 10");
	calculateMatchmaking(62, 12, 70, 71, "62a + 12b / a+b ≈ 70");
	// calculateMatchmaking(-30, 0, , , "-30a + 0b / a+b ≈ ");
	// calculateMatchmaking(0, 0, , , "0a + 0b / a+b ≈ ");
	// calculateMatchmaking(100, 0, , , "100a + 0b / a+b ≈ ");
	// calculateMatchmaking(80, 0, , , "80a + 0b / a+b ≈ ");
	calculateMatchmaking(-18, 7, -20, -19, "-18a + 7b / a+b ≈ -20");
	calculateMatchmaking(89, 10, 95, 96, "89a + 10b / a+b ≈ 95");
	calculateMatchmaking(49, 74, 30, 31, "49a + 74b / a+b ≈ 30");
	calculateMatchmaking(37, 19, 25, 26, "37a + 19b / a+b ≈ 25");
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

func calculateMatchmaking(
	average float64, 
	standardDeviation float64, 
	minimumResult float64, 
	maximumResult float64, 
	expression string,
) {
	var averageCoefficient = rand.Float64();
	var standardDeviationCoefficient = rand.Float64();
	var condition int = 1000;
	index := 1

	for {
		index++
		if (((average * averageCoefficient) + (standardDeviation * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult && ((average * averageCoefficient) + (standardDeviation * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult) {
			fmt.Println();
			fmt.Printf("=======================\n");
			fmt.Printf("\texpression %s is equal in round: %d", expression, index);
			fmt.Println();
			fmt.Println("\tAverage coefficient is: ", averageCoefficient);
			fmt.Println();
			fmt.Println("\tStandard deviation coefficient is: ", standardDeviationCoefficient);
			fmt.Printf("=======================");
			fmt.Println();
			break;
		}
		if (index > condition) {
			condition += 1000;
			minimumResult -= 0.25;
			maximumResult += 0.25;
		}
		averageCoefficient = rand.Float64();
		standardDeviationCoefficient = rand.Float64();
	}
}