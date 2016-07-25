package control_structures

func collatzChainLength(n int) int {
	i := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		i++
	}
	return i

}
