package main

import (
	"fmt"
	"math"
	"math/bits"
)

const (
	maxNodes = 1000000
)

var (
	encBits [maxNodes]int             // numbers of bits necessary to encode the tree of size n 
	minHeight [maxNodes]int32            // minimal height of a balanced tree of size n
	maxHeight [maxNodes]int32            // maximum height of a balanced tree of size n
	maxSplit  [maxNodes]float64
)

// Computes comp[n], minHeight[n], maxHeight[n]
func compute(n int) {
	if n == 1 {
		minHeight[n] = 1
		maxHeight[n] = 1
		encBits[n] = 0
		return
	}
	if n == 2 {
		minHeight[n] = 2
		maxHeight[n] = 2
		encBits[n] = 1
		return		
	}
	var maxH int32 = 0
	var minH int32 = math.MaxInt32
	var maxS float64 = 0.0
	var rootPositions uint = 0 // Number of positions of the root for the tree of this size
	var maxChildBits int = 0
	for splitIndex := 1; splitIndex <= (n-1)/2; splitIndex++ {
		nRight := n - splitIndex - 1
		// See if there is an overlap
		if minHeight[nRight] <= maxHeight[splitIndex] + 1 {
			if minHeight[nRight] + 1 < minH {
				minH = minHeight[nRight] + 1
			}
			if maxHeight[nRight] > maxHeight[splitIndex] {
				if maxHeight[splitIndex] + 1 + 1 > maxH {
					maxH = maxHeight[splitIndex] + 1 + 1
				}
			} else {
				if maxHeight[splitIndex] + 1 > maxH {
					maxH = maxHeight[splitIndex] + 1
				}
			}
			if splitIndex*2+1 == n {
				rootPositions++
			} else {
				rootPositions+=2
			}
			var childBits int = encBits[splitIndex] + encBits[nRight]
			if childBits > maxChildBits {
				maxChildBits = childBits
			}
			split := float64(nRight)/float64(splitIndex)
			if split > maxS {
				maxS = split
			}
		}
	}
	minHeight[n] = minH
	maxHeight[n] = maxH
	maxSplit[n] = maxS
	encBits[n] = bits.Len(rootPositions-1) + maxChildBits
}

func main() {
	for n := 1; n < maxNodes; n++ {
		compute(n)
		if (n == maxNodes - 1) || (n % 1000 == 0) {
			fmt.Printf("n = %d, c = %d, minH = %d, maxH = %d, maxS = %.4f\n", n, encBits[n], minHeight[n], maxHeight[n], maxSplit[n])
		}
	}
}