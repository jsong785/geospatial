package gotest

import (
	"fmt"
	"reflect"
	"runtime"
)

// TestInterface - interface for testing
type TestInterface interface {
	OutputWarningMessage(warningMessage string)
	OutputErrorMessage(errorMessage string)
}

var testPtr TestInterface

func init() {
	testPtr = nil
}

// Initialize - initialize tests
func Initialize(t TestInterface) {
	testPtr = t
}

// Cleanup - cleanup tests
func Cleanup() {
	testPtr = nil
}

func outputErrorMessage(stackSkip int, errorMessage string) {
	_, file, line, _ := runtime.Caller(stackSkip)
	formattedErrorMessage := fmt.Sprintf("[%v:%v] error: \n%v", file, line, errorMessage)
	testPtr.OutputErrorMessage(formattedErrorMessage)
}

// ASSERT_EQ ...
func ASSERT_EQ(arg1, arg2 interface{}) {
	if !reflect.DeepEqual(arg1, arg2) {
		message := fmt.Sprintf("%v != %v; expected ==", arg1, arg2)
		outputErrorMessage(2, message)
	}
}

// ASSERT_NE ...
func ASSERT_NE(arg1, arg2 interface{}) {
	if reflect.DeepEqual(arg1, arg2) {
		message := fmt.Sprintf("%v == %v; expected !=", arg1, arg2)
		outputErrorMessage(2, message)
	}
}

// ASSERT_TRUE ...
func ASSERT_TRUE(val bool) {
	if !val {
		message := fmt.Sprintf("%v; expected true", val)
		outputErrorMessage(2, message)
	}
}

// ASSERT_FALSE ...
func ASSERT_FALSE(val bool) {
	if val {
		message := fmt.Sprintf("%v; expected false", val)
		outputErrorMessage(2, message)
	}
}

func outputWarningMessage(stackSkip int, errorMessage string) {
	_, file, line, _ := runtime.Caller(stackSkip)
	formattedErrorMessage := fmt.Sprintf("[%v:%v] warning: \n%v", file, line, errorMessage)
	testPtr.OutputWarningMessage(formattedErrorMessage)
}

// EXPECT_EQ ...
func EXPECT_EQ(arg1, arg2 interface{}) {
	if !reflect.DeepEqual(arg1, arg2) {
		message := fmt.Sprintf("%v != %v; expected ==", arg1, arg2)
		outputWarningMessage(2, message)
	}
}

// EXPECT_NE ...
func EXPECT_NE(arg1, arg2 interface{}) {
	if reflect.DeepEqual(arg1, arg2) {
		message := fmt.Sprintf("%v == %v; expected !=", arg1, arg2)
		outputWarningMessage(2, message)
	}
}

// EXPECT_TRUE ...
func EXPECT_TRUE(val bool) {
	if !val {
		message := fmt.Sprintf("%v; expected true", val)
		outputWarningMessage(2, message)
	}
}

// EXPECT_FALSE ...
func EXPECT_FALSE(val bool) {
	if val {
		message := fmt.Sprintf("%v; expected false", val)
		outputWarningMessage(2, message)
	}
}
