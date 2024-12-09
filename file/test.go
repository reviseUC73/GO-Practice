package file

import (
	"fmt"
)

func TestPrint() {
	fmt.Println("test print")
}

func privteTestPrint() { // letter func is private
	fmt.Println("test print private")
}