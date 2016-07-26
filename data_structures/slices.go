package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(operation mapOperation, vals []int32) []int32 {
	var mapped_values []int32
	for _, value := range vals {
		mapped_values = append(mapped_values, operation(value))
	}
	return mapped_values

}

func filterInts(operation filterOperation, vals []int32) []int32 {
	return []int32{3}
}

func concatenate(destination []string, newValues ...string) []string {
	// for _, newVal := range newValues {
	// 	destination = append(destination, newVal)
	// }
	return destination
}

func equals(list1 []string, list2 []string) bool {
	return false
	// if len(list1) != len(list2){
	// 	return false
	// }
	// isEqual := true
	// for index, value := range list2{
	// 	if list1[index] != value{
	// 		isEqual = false
	// 	}
	// }
	// return isEqual
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
