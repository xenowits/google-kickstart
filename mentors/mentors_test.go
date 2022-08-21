package mentors

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

// input represents individual input test case.
type input struct {
	t       int   // test case no
	n       int   // no of ratings
	ratings []int // ratings of students
}

func TestAssignMentors(tt *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "test_data/test_set_1/ts1_input.txt",
			output: "test_data/test_set_1/ts1_output.txt",
		},
		{
			input:  "test_data/test_set_2/ts2_input.txt",
			output: "test_data/test_set_2/ts2_output.txt",
		},
	}

	for _, test := range tests {
		f, err := os.Open(test.input)
		if err != nil {
			fmt.Println(err)
			panic("cannot open file")
		}

		defer f.Close()

		var (
			iter    int
			scanner = bufio.NewScanner(f)
			inputs  []input
		)

		const maxCapacity int = 1000005 // your required line length
		buf := make([]byte, maxCapacity)
		scanner.Buffer(buf, maxCapacity)

		// Get no of test cases
		if !scanner.Scan() {
			panic("scan error")
		}

		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("cannot convert no of test cases to integer")
		}

		for iter < t {
			if !scanner.Scan() {
				panic("scan error")
			}

			n, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("cannot convert no of ratings to integer")
			}

			if !scanner.Scan() {
				panic("scan error")
			}

			var ratings []int
			rates := strings.Fields(scanner.Text())
			for _, rate := range rates {
				r, err := strconv.Atoi(rate)
				if err != nil {
					fmt.Println("cannot convert rating to integer")
				}

				ratings = append(ratings, r)
			}

			inputs = append(inputs, input{
				t:       iter + 1,
				n:       n,
				ratings: ratings,
			})

			iter += 1
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		var sb strings.Builder
		for _, input := range inputs {
			mentors := AssignMentors(input.n, input.ratings)

			sb.WriteString(fmt.Sprintf("Case #%d:", input.t))
			for i := 0; i < input.n; i++ {
				sb.WriteString(fmt.Sprintf(" %d", mentors[input.ratings[i]]))
			}
			sb.WriteString("\n")

			iter++
		}

		actual, err := os.ReadFile(test.output)
		if err != nil {
			panic("couldn't read actual output file")
		}

		require.Equal(tt, sb.String(), string(actual))
	}
}
