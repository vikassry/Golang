package fundamentals_2

import "testing"
import "fmt"

func TestShouldMakeSliceOfLength5Capacity10(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	if len(s) != 5 {
		t.Error("Slice length is not 5")
	}

	t.Skip() // skips the statements after this line
	if cap(s) != 10 {
		t.Error("Slice capacity is not 10")
	}
}

func TestAppendSliceAtoB(t *testing.T) {
	a := []int{1}[:]
	b := []int{2, 3}[:]
	c := a[:]
	fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", a, len(a), cap(a)))
	fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", b, len(b), cap(b)))
	fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", c, len(c), cap(c)))
	// t.Skip()
	for _, item := range b {
		c = append(c, item)
	}
	fmt.Println(fmt.Sprintf("\n%v is len(%d) cap(%d)", a, len(a), cap(a)))
	fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", c, len(c), cap(c)))

	if len(c) != 3 {
		t.Error("Array not appended properly")
	}
}
