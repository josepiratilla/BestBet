package simulator

import (
	"math"
	"math/rand"
)

//Simulator holds the logic to simulate multiple bets
type Simulator struct {
}

//RunOnce executes one exercise and see the result
//Starts by and initial amount of 1
func RunOnce(probabilityWin float64, percentBet float64, iterations int) float64 {
	money := 1.0
	for a := 0; a < iterations; a++ {
		bet := money * percentBet
		if rand.Float64() < probabilityWin {
			money += bet
		} else {
			money -= bet
		}
	}
	return money
}

//RunMultiple executes multiple runs on RunOnce and returns the average.
func RunMultiple(probabilityWin float64, percentBet float64, iterations int, runs int) float64 {
	Total := 0.0
	for a := 0; a < runs; a++ {
		Total += RunOnce(probabilityWin, percentBet, iterations)
	}
	return Total / float64(runs)
}

//RunMultipleSquareRootAverage executes multiple runs on RunOnce and returns the square root average.
func RunMultipleSquareRootAverage(probabilityWin float64, percentBet float64, iterations int, runs int) float64 {
	Total := 0.0
	for a := 0; a < runs; a++ {
		one := RunOnce(probabilityWin, percentBet, iterations)
		Total += one
	}
	return math.Sqrt(Total) / float64(runs)
}

//GetBestBetAverage returns the best bet based on the maximum average depending on the probability Win
func GetBestBetAverage(probabilityWin float64, iterations int, runs int, precission float64) float64 {
	probBoundMin := 0.0
	probBoundMax := 1.0
	probMid := 0.5
	resultMid := RunMultiple(probabilityWin, probMid, iterations, runs)
	for probBoundMax-probBoundMin > precission {
		probLow := (probBoundMin + probMid) / 2
		resultLow := RunMultiple(probabilityWin, probLow, iterations, runs)
		probHigh := (probBoundMax + probMid) / 2
		resultHigh := RunMultiple(probabilityWin, probHigh, iterations, runs)
		if resultLow > resultMid {
			probBoundMax = probMid
			probMid = probLow
			resultMid = resultLow
			continue
		}
		if resultHigh > resultMid {
			probBoundMin = probMid
			probMid = probHigh
			resultMid = resultHigh
			continue
		}
		probBoundMin = probLow
		probBoundMax = probHigh
	}
	return probMid
}

//GetBestBetSquareRootAverage returns the best bet based on the maximum average depending on the probability Win
func GetBestBetSquareRootAverage(probabilityWin float64, iterations int, runs int, precission float64) float64 {
	probBoundMin := 0.0
	probBoundMax := 1.0
	probMid := 0.5
	resultMid := RunMultipleSquareRootAverage(probabilityWin, probMid, iterations, runs)
	for probBoundMax-probBoundMin > precission {
		probLow := (probBoundMin + probMid) / 2
		resultLow := RunMultipleSquareRootAverage(probabilityWin, probLow, iterations, runs)
		probHigh := (probBoundMax + probMid) / 2
		resultHigh := RunMultipleSquareRootAverage(probabilityWin, probHigh, iterations, runs)
		if resultLow > resultMid {
			probBoundMax = probMid
			probMid = probLow
			resultMid = resultLow
			continue
		}
		if resultHigh > resultMid {
			probBoundMin = probMid
			probMid = probHigh
			resultMid = resultHigh
			continue
		}
		probBoundMin = probLow
		probBoundMax = probHigh
	}
	return probMid
}
