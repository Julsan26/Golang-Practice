package Golang_Practice

func solution(inputArray []int) int {
	max := inputArray[0] * inputArray[1]

	for i := 0; i < len(inputArray)-1; i++ {
		if max < inputArray[i]*inputArray[i+1] {
			max = inputArray[i] * inputArray[i+1]
		}
	}

	return max
}