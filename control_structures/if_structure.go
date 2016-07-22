package control_structures

import (
    "strconv"
)


func fizzBuzz(i int32) string {
	if i % 3 == 0 &&  i % 5 == 0 {
		return "FizzBuzz"
	}
	if i % 3 == 0 {
		return "Fizz"
	}
	if i % 5 == 0 {
		return "Buzz"
	}
	t := strconv.FormatInt(int32(i))
	return t
}
