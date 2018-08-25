package kata

import "math"

// There are two tasks.
// The first one is to calculate the total stopping distance in meters given v, mu (and the reaction time t = 1).
// dist(v, mu) -> d = total stopping distance
// The second task is to calculate v in km per hour knowing d in meters and mu with the supposition that the reaction time is still t = 1.
// speed(d, mu) -> v such that dist(v, mu) = d.

const g = 9.81

func Dist(v, mu float64) float64 { // suppose reaction time is 1
	v = v * 10 / 36
	return v + (v * v / (2 * mu * g))
}

func Speed(d, mu float64) float64 { // suppose reaction time is 1
	k := 2 * mu * g
	v := (-k + math.Sqrt(k*k+4*d*k)) / 2
	return v * 36 / 10
}

// Solution:
// d = vk/k + vv/k
// dk = vk + vv
// dk + 0.25kk = vk + vv + 0.25kk
// dk + 0.25kk = (v + 0.5k)^2
// v + 0.5k = sqrt(dk + 0.25kk)
// v = -0.5k + sqrk(dk + 0.25kk)
