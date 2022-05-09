package homework

func reverse(input []int64) (result []int64) {
	var maxIndex = len(input) - 1
	for i := range input {
		result = append(result, input[maxIndex-i])
	}
	return
}
