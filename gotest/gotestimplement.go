package gotest

import "fmt"
import "testing"

// TestImplement ...
type TestImplement struct {
	Test *testing.T
}

// OutputErrorMessage ...
func (i TestImplement) OutputErrorMessage(errorMessage string) {
	//i.Test.Log(errorMessage)
	fmt.Print(errorMessage)
	fmt.Print("\n")
	i.Test.Fail()
}

// OutputWarningMessage ...
func (i TestImplement) OutputWarningMessage(warningMessage string) {
	i.Test.Log(warningMessage)
	fmt.Print(warningMessage)
	fmt.Print("\n")
}
