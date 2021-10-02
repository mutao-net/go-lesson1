package mylib

func Avereage(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
