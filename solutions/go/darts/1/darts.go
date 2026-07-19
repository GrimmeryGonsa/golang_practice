package darts

import "math"

func Score(x, y float64) int {
	square := x*x + y*y
    squareRoot := math.Ceil( math.Sqrt(square))
    if squareRoot <= 1{
        return 10
    } else if squareRoot <= 5{
        return 5
    } else if squareRoot <= 10 {
        return 1
    }
    return 0
}
