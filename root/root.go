package root

import "fmt"

var values []int

func init() {
	fmt.Println("inside root init")
	for i := 0; i < 10; i++ {
		values = append(values, i)
	}
}

func AddValue(i int) {
	values = append(values, i)
}

func Noop() {
	fmt.Println("inside noop")
}
