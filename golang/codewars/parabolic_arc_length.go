// We want to approximate the length of a curve representing a function y = f(x) with a<= x <= b. First, we split the interval [a, b] into n sub-intervals with widths h1, h2, ... , hn by defining points x1, x2 , ... , xn-1 between a and b. This defines points P0, P1, P2, ... , Pn on the curve whose x-coordinates are a, x1, x2 , ... , xn-1, b and y-coordinates f(a), f(x1), ..., f(xn-1), f(b) . By connecting these points, we obtain a polygonal path approximating the curve.

// Our task is to approximate the length of a parabolic arc representing the curve y = x * x with x in the interval [0, 1]. We will take a common step h between the points xi: h1, h2, ... , hn = h = 1/n and we will consider the points P0, P1, P2, ... , Pn on the curve. The coordinates of each Pi are (xi, yi = xi * xi).

// The function len_curve (or similar in other languages) takes n as parameter (number of sub-intervals) and returns the length of the curve truncated to 9 decimal places.

package kata

import (
	"math"
)

func LenCurve(n int) float64 {
	var result float64
	coordinates := coordinatesOfX(n)
	for i := 0; i < len(coordinates)-1; i++ {
		result += distance(coordinates[i], coordinates[i+1])
	}
	return truncate(result)
}

func coordinatesOfX(n int) []float64 {
	distance := float64(1) / float64(n)
	var result []float64
	for i := 0; i <= n; i++ {
		result = append(result, distance*float64(i))
	}
	return result
}

func coordinateOfY(x float64) float64 {
	return x * x
}

func distance(start float64, end float64) float64 {
	valX := math.Pow(end-start, 2)
	valY := math.Pow(coordinateOfY(end)-coordinateOfY(start), 2)
	return math.Sqrt(valX + valY)
}

func truncate(target float64) float64 {
	return math.Round(target*math.Pow(10, 9)) / math.Pow(10, 9)
}
