package main

import (
	"bufio"
	"fmt"
	"os"
	"xenowits.com/m/mentors"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

// input represents individual input test case.
type input struct {
	t       int // test case no
	n       int
	ratings []int
}

func main() {
	defer writer.Flush()
	var iter, t int
	scanf("%d\n", &t)

	var inputs []input

	for iter < t {
		var n int
		scanf("%d\n", &n)

		var ratings []int
		for i := 0; i < n; i++ {
			var rating int
			scanf("%d", &rating)

			ratings = append(ratings, rating)
		}
		scanf("\n")

		inputs = append(inputs, input{
			t:       iter + 1,
			n:       n,
			ratings: ratings,
		})

		iter += 1
	}

	for _, input := range inputs {
		assignedMentors := mentors.AssignMentors(input.n, input.ratings)

		printf("Case #%d:", input.t)
		for i := 0; i < input.n; i++ {
			printf(" %d", assignedMentors[input.ratings[i]])
		}
		printf("\n")

		iter++
	}
}
