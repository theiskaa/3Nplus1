package main

import (
	"fmt"
	"time"
)

// executeLoop is main function of 3Nplus1.
// Uses threeNplusOne function to do actual job.
func executeLoop(baseN, loopTurn int, logger func(operation string, tryIndex int)) {
	n := baseN

	callback := func(tryIndex int) {
		time.Sleep(800 * time.Millisecond)

		res, operation := threeNplusOne(n)
		n = res

		logger(operation, tryIndex)
	}

	// -1 represents infinity in in our case.
	if loopTurn != -1 {
		for i := 1; i <= loopTurn; i++ {
			callback(i)
		}
		return
	}

	try := 0
	for {
		try++
		callback(try)
	}

}

// threeNplusOne takes N, generates result and full operation history-string.
func threeNplusOne(n int) (int, string) {
	if n%2 == 0 {
		res := n / 2
		operation := fmt.Sprintf("%v is EVEN ──▶ %v / 2 = %v", n, n, res)

		return res, operation
	}

	res := (3 * n) + 1
	operation := fmt.Sprintf("%v is ODD ──▶ (3 * %v) + 1 = %v", n, n, res)

	return res, operation
}
