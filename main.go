package main

import (
	"BestBet/simulator"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(simulator.GetBestBetSquareRootAverage(0.5, 100, 200000, 0.1))
}
