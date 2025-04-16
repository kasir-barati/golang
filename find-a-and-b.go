package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var averages = [6]float64 {10, 62, -18, 89, 49, 37};
	var standardDeviations = [6]float64 {44, 12, 7, 10, 74, 19};
	var minimumResults = [6]float64 {10, 70, -20, 95, 30, 25};
	var maximumResults = [6]float64 {11, 71, -19, 96, 31, 26};
	calculateCoefficient(averages, standardDeviations, minimumResults, maximumResults);
}

func test1() {
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


func calculateCoefficient(
	average [6]float64, 
	standardDeviation [6]float64, 
	minimumResult [6]float64, 
	maximumResult [6]float64, 
) {
	rand.Seed(time.Now().UTC().UnixNano())
	var averageCoefficient = rand.Float64();;
	var standardDeviationCoefficient  = rand.Float64();
	var condition int = 10000;
	index := 1

	for {
		index++
		if (
			((average[0] * averageCoefficient) + (standardDeviation[0] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult[0] &&
			((average[0] * averageCoefficient) + (standardDeviation[0] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult[0] &&
			((average[1] * averageCoefficient) + (standardDeviation[1] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult[1] &&
			((average[1] * averageCoefficient) + (standardDeviation[1] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult[1] && 
			((average[2] * averageCoefficient) + (standardDeviation[2] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult[2] &&
			((average[2] * averageCoefficient) + (standardDeviation[2] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult[2] && 
			((average[3] * averageCoefficient) + (standardDeviation[3] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult[3] &&
			((average[3] * averageCoefficient) + (standardDeviation[3] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult[3] && 
			((average[4] * averageCoefficient) + (standardDeviation[4] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult[4] &&
			((average[4] * averageCoefficient) + (standardDeviation[4] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult[4] &&  
			((average[5] * averageCoefficient) + (standardDeviation[5] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) >= minimumResult[5] &&
			((average[5] * averageCoefficient) + (standardDeviation[5] * standardDeviationCoefficient)) / (averageCoefficient + standardDeviationCoefficient) <= maximumResult[5] ) {
			
			fmt.Println();
			fmt.Printf("=======================\n");
			fmt.Println("Coefficient found in round: ", index);
			fmt.Println();
			fmt.Println("minimum result is: ", minimumResult);
			fmt.Println();
			fmt.Println("maximum result is: ", maximumResult);
			fmt.Println();
			fmt.Println("\tAverage coefficient is: ", averageCoefficient);
			fmt.Println();
			fmt.Println("\tStandard deviation coefficient is: ", standardDeviationCoefficient);
			fmt.Printf("=======================");
			fmt.Println();
			break;
		}
		if (index > condition) {
			condition += 10000;
			for index := 1; index < 6; index++ {
				minimumResult[index] -= 0.25;
				maximumResult[index] += 0.25;
			}
		}
		averageCoefficient = rand.Float64();
		standardDeviationCoefficient = rand.Float64();
	}
}
