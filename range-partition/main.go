package main

import (
	"bufio"
	"fmt"
	"os"
	// range_partition "xenowits.com/m/range-partition/range-partition"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	_, err := fmt.Fprintf(writer, f, a...)
	if err != nil {
		panic("cannot print")
	}
}

func scanf(f string, a ...interface{}) {
	_, err := fmt.Fscanf(reader, f, a...)
	if err != nil {
		panic("cannot scan")
	}
}

// input represents individual input test case.
type input struct {
	t int // test case no
	n int // {1,2,3,...n}
	x int // ratio x:y
	y int // ratio x:y
}

func main() {
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {
			panic("cannot flush to output")
		}
	}(writer)

	var iter, t int
	scanf("%d\n", &t)

	var inputs []input
	for iter < t {
		var n, x, y int
		scanf("%d %d %d\n", &n, &x, &y)

		inputs = append(inputs, input{
			t: iter + 1,
			n: n,
			x: x,
			y: y,
		})

		iter += 1
	}

	for _, input := range inputs {
		printf("Case #%d: ", input.t)
		possible, alanSize, alanInts := RangePartition(input.n, input.x, input.y)
		if !possible {
			printf("IMPOSSIBLE\n")
			continue
		}

		printf("POSSIBLE\n")
		printf("%d\n", alanSize)
		for _, v := range alanInts {
			printf("%d ", v)
		}
		printf("\n")
	}
}

func RangePartition(n, x, y int) (bool, int, []int) {
	totalSum := (n * (n + 1)) / 2 // s = Sum(1,2,....n)

	// Alan's sum (s1) = s*x/(x+y)
	// Barbara's sum (s2) = s*y/(x+y)

	if totalSum%(x+y) != 0 {
		return false, 0, nil // Impossible
	}

	var p, q, r int
	r = totalSum * x / (x + y)
	var tmp []int

	// Find p < r < q
	for p1 := 1; p1 <= n; p1++ {
		p += p1          // sum upto p1
		q = p + (p1 + 1) // sum upto p1+1
		tmp = append(tmp, p1)

		if p == r {
			return true, p1, tmp // perfect sum, ie, p == r == Sum{1,2,..p1}
		}
		if r == q {
			tmp = append(tmp, p1+1)
			return true, p1 + 1, tmp // perfect sum, ie, q == r == Sum{1,2,..p1+1}
		}

		// get to the point where r lies between p and q
		if !(p < r && r < q) {
			continue // we need to move forward, we haven't reached there yet
		}

		// we have got to the point where p < r < q
		// The answer here is {1,2,3...q} - {X}, where X = (q-r)
		var ans []int
		for i := 1; i <= p1+1; i++ {
			if i != (q - r) { // exclude X
				ans = append(ans, i)
			}
		}

		return true, p1, ans
	}

	return false, 0, nil
}
