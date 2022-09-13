package set

import (
	"testing"
	"fmt"
)

func TestSet(t *testing.T) {
	set := New()

	fmt.Println(set.IsEmpty())
	set.Add("李")
	set.Add("里")
	set.Add("理")
	set.Add("李")
	set.Add(1)

	fmt.Println(set.GetLength())
	fmt.Println(set.Contains("理"))
	set.Remove("理")
	fmt.Println(set.Contains("理"))
	fmt.Println(set.IsEmpty())
	fmt.Println(set.GetLength())
}
