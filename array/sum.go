package array

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers{
		sum += v
	}
 return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sum := Sum(numbers)
		sums = append(sums, sum) 
	}
	return sums
}

func SumAllTails(numbers ...[]int ) []int {
	var result [] int
	for _, v := range numbers{
		if len(v) == 0 {
			result = append(result, 0)
		} else {
		result = append(result, Sum(v[1:]))
		}
	}
	return result
}