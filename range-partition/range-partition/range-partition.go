// Problem: https://codingcompetitions.withgoogle.com/kickstart/round/00000000008cb4d1/0000000000b20deb
//
// Alan and Barbara suddenly felt like playing with numbers. Alan chooses a non-empty subset from the set of first N
// positive integers (1,2,…,N). Barbara takes the rest of the numbers (if any) from the set. And then they both calculate the sum of the elements in their respective sets.
//
// Alan believes in a magic ratio, which is X:Y. Hence, Alan wants to choose the subset in such a way that the ratio between the sum of Alan's subset and the sum of Barbara's subset is exactly X:Y.
//
// Can you help Alan to choose a subset that can achieve the desired ratio?
// Input
//
// The first line of the input gives the number of test cases, T. T test cases follow.
// Each test case has a single line containing three integers, N, X and Y, as described above.
// Output
//
// For each test case, output the first line containing Case #x: y, where x is the test case number (starting from 1) and y is POSSIBLE, if Alan can choose such a non-empty subset, and IMPOSSIBLE otherwise.
// If you print POSSIBLE, then output two more lines for that test case.
// In the second line, print a single integer, which denotes the size of Alan's subset.
// In the third line, print the integers present in Alan's subset.
// If there are multiple solutions, you can print any of them.
// Limits
//
// Time limit: 5 seconds.
// Memory limit: 1 GB.
// 1≤T≤100
// .
// 1≤X≤108.
// 1≤Y≤108.
// gcd(X,Y)=1, where gcd is Greatest common divisor.
// Test Set 1
//
// 1≤N≤15.
// Test Set 2
//
// 1≤N≤5000.

package range_partition

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
