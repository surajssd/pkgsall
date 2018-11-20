package one

import (
	"fmt"

	"github.com/surajssd/study/pkgsall/root"
)

func init() {
	fmt.Println("inside one init")
	root.AddValue(1)
}
