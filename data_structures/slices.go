package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(operation mapOperation, values []int32) []int32 {
	var mappedValues []int32
	for _, value := range values {
		mappedValues = append(mappedValues, operation(value))
	}
	return mappedValues

}

func filterInts(operation filterOperation, values []int32) []int32 {
	var filterValues []int32
	for _, value := range values {
		if operation(value) {
			filterValues = append(filterValues, value)
		}
	}
	return filterValues
}

func concatenate(destination []string, newValues ...string) []string {
	for _, newVal := range newValues {
		destination = append(destination, newVal)
	}
	return destination
}

func equals(list1 []string, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}
	isEqual := true
	for index, value := range list2 {
		if list1[index] != value {
			isEqual = false
		}
	}
	return isEqual
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
