package arrslices

func Sum(arr []int) (sum int) {

	for _, val := range arr {
		sum += val
	}
	return
}

func SumAll(numsToSum ...[]int) (sum []int) {

	for _, numbers := range numsToSum {
		sum = append(sum, Sum(numbers))
	}
	return
}

func SumAllTails(numsToSum ...[]int) (sum []int) {
	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers[1:]))
		}
	}
	return
}
