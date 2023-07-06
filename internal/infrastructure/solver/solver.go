package solver

import (
	"errors"
	"math"
	"strings"
)

type TaskSolver struct{}

func New() TaskSolver {
	return TaskSolver{}
}

func (ts TaskSolver) SolveTask1(nums []int) (int, error) {
	var min, max int
	isDecreasing := true

	if len(nums) < 2 {
		return 0, errors.New("nums cannot have a length less than 2")
	}

	min = nums[0]
	max = min

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			isDecreasing = false
		}

		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}

	}

	if isDecreasing {
		return 0, nil
	}

	return max - min, nil
}

func (ts TaskSolver) SolveTask2(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	var result byte

	for i := 0; i < len(str1); i++ {
		result ^= str1[i] ^ str2[i]
	}

	return result == 0
}

func (ts TaskSolver) SolveTask3(s string, numRows int) (string, error) {

	if len(s) == 0 {
		return "", errors.New("string cannot have zero length")
	}

	if numRows <= 0 {
		return "", errors.New("num rows cannot be less than or equal to zero")
	}

	var builder strings.Builder
	builder.Grow(len(s))
	var segmentCount int
	k := len(s)

	if numRows != 1 {
		segmentCount = int(math.Ceil(float64(k) / float64(numRows-1)))
	} else {
		return s, nil
	}

	var ind int
	var maxSegmentLen = numRows - 1

	for i := 0; i < numRows; i++ {

		if i == 0 || i == numRows-1 {
			for j := 0; j < segmentCount; {
				ind = i + maxSegmentLen*j
				if ind < len(s) {
					builder.WriteByte(s[ind])
				}
				j += 2
			}
			continue
		}

		for j := 0; j < segmentCount; j++ {

			if j%2 == 0 {
				ind = i + maxSegmentLen*j
			} else {
				ind = (maxSegmentLen * (j + 1)) - i
			}

			if ind < len(s) {
				builder.WriteByte(s[ind])
			}
		}

	}

	return builder.String(), nil
}
