package two

import (
	"fmt"

	"github.com/surajssd/study/pkgsall/root"
)

func init() {
	fmt.Println("inside two init")
	root.AddValue(-2)
}
