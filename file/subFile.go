package file

import (
	"fmt"
	"github.com/reviseUC73/go-test1/file/internal/fileInternal"
)

func MyPrint() {
	fmt.Printf("sub file print")
	TestPrint()
	privteTestPrint()

	// from internal folder
	fileInternal.MyPrintFileInternal()
}
