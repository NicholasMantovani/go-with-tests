package genericarrays

func Sum(numbers []int) int {

	add := func(acc, x int) int {
		return acc + x
	}

	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {

	sumAll := func(acc, x []int) []int {
		return append(acc, Sum(x))
	}

	return Reduce(numbersToSum, sumAll, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {

	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(x[1:]))
		}
	}

	return Reduce(numbersToSum, sumTail, []int{})
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {

	adjustBalanace := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			currentBalance -= t.Sum
		}
		if t.To == name {
			currentBalance += t.Sum
		}
		return currentBalance
	}

	return ReduceModified(transactions, adjustBalanace, 0.0)
}

func ReduceModified[I, O any](input []I, aggregator func(O, I) O, zeroValue O) O {
	var res = zeroValue
	for _, x := range input {
		res = aggregator(res, x)
	}
	return res
}

func Reduce[T any](input []T, aggregator func(T, T) T, zeroValue T) T {
	res := zeroValue
	for _, v := range input {
		res = aggregator(res, v)
	}
	return res
}

func Find[T any](input []T, findFunction func(T) bool) (value T, found bool) {
	for _, v := range input {
		if findFunction(v) {
			return v, true
		}
	}
	return
}
