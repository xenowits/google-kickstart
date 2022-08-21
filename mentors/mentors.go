package mentors

import "sort"

func AssignMentors(n int, ratings []int) map[int]int {
	sortedRatings := make([]int, n)
	copy(sortedRatings[:], ratings)

	sort.Ints(sortedRatings)

	mentors := make(map[int]int)

	for i := 0; i < n; i++ {
		mentee := sortedRatings[i]

		// binary search from arr[i+1:n-1] for 2*mentee
		var (
			start  = i + 1
			end    = n - 1
			mid    int
			mentor = -1
		)

		for start <= end {
			mid = (start + end) / 2
			if sortedRatings[mid] <= 2*mentee {
				// okay for mentorship
				mentor = sortedRatings[mid]
				start = mid + 1
			} else {
				// not okay for mentorship, need a lower rating
				end = mid - 1
			}
		}

		if mentor == -1 && i > 0 {
			// higher rating mentor not found
			mentor = sortedRatings[i-1]
		}

		mentors[mentee] = mentor
	}

	return mentors
}
