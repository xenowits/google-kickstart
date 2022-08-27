package main

import (
	"bufio"
	"fmt"
	"os"
	rp "xenowits.com/m/range-partition/range-partition"
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
		possible, alanSize, alanInts := rp.RangePartition(input.n, input.x, input.y)
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
