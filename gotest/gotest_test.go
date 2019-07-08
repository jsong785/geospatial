package gotest

import "testing"

var previousErrorMessage string
var previousWarningMessage string

type MockTestInterface struct{}

func (i MockTestInterface) OutputErrorMessage(errorMessage string) {
	previousErrorMessage = errorMessage
}

func (i MockTestInterface) OutputWarningMessage(warningMessage string) {
	previousErrorMessage = warningMessage
}

type mockStruct struct {
	x int
	y string
}

func TestAssertEqual(t *testing.T) {
	Initialize(MockTestInterface{})
	defer Cleanup()

	previousErrorMessage = ""
	ASSERT_EQ(1, 1)
	if len(previousErrorMessage) != 0 {
		t.Error("unepxected error")
	}

	previousErrorMessage = ""
	ASSERT_EQ(1, 2)
	if len(previousErrorMessage) == 0 {
		t.Error("unepxected error")
	}

	previousErrorMessage = ""
	ASSERT_EQ(mockStruct{1, "same"}, mockStruct{1, "same"})
	if len(previousErrorMessage) != 0 {
		t.Error("unepxected error")
	}

	previousErrorMessage = ""
	ASSERT_EQ(mockStruct{1, "diff"}, mockStruct{1, "same"})
	if len(previousErrorMessage) == 0 {
		t.Error("unepxected error")
	}
}

func TestExpectEqual(t *testing.T) {
	Initialize(MockTestInterface{})
	defer Cleanup()

	previousWarningMessage = ""
	EXPECT_EQ(1, 1)
	if len(previousWarningMessage) != 0 {
		t.Log("unepxected error")
	}

	previousWarningMessage = ""
	EXPECT_EQ(1, 2)
	if len(previousWarningMessage) == 0 {
		t.Log("unepxected error")
	}

	previousWarningMessage = ""
	EXPECT_EQ(mockStruct{1, "same"}, mockStruct{1, "same"})
	if len(previousWarningMessage) != 0 {
		t.Log("unepxected error")
	}

	previousWarningMessage = ""
	EXPECT_EQ(mockStruct{1, "diff"}, mockStruct{1, "same"})
	if len(previousWarningMessage) == 0 {
		t.Log("unepxected error")
	}
}
func TestAssertNotEqual(t *testing.T) {
	Initialize(MockTestInterface{})
	defer Cleanup()

	previousErrorMessage = ""
	ASSERT_NE(1, 2)
	if len(previousErrorMessage) != 0 {
		t.Error("unepxected error")
	}

	previousErrorMessage = ""
	ASSERT_NE(1, 1)
	if len(previousErrorMessage) == 0 {
		t.Error("unepxected error")
	}

	previousErrorMessage = ""
	ASSERT_NE(mockStruct{1, "diff"}, mockStruct{1, "same"})
	if len(previousErrorMessage) != 0 {
		t.Error("unepxected error")
	}

	previousErrorMessage = ""
	ASSERT_NE(mockStruct{1, "same"}, mockStruct{1, "same"})
	if len(previousErrorMessage) == 0 {
		t.Error("unepxected error")
	}
}

func TestExpectNotEqual(t *testing.T) {
	Initialize(MockTestInterface{})
	defer Cleanup()

	previousWarningMessage = ""
	EXPECT_NE(1, 2)
	if len(previousWarningMessage) != 0 {
		t.Log("unepxected error")
	}

	previousWarningMessage = ""
	EXPECT_NE(1, 1)
	if len(previousWarningMessage) == 0 {
		t.Log("unepxected error")
	}

	previousWarningMessage = ""
	EXPECT_NE(mockStruct{1, "diff"}, mockStruct{1, "same"})
	if len(previousWarningMessage) != 0 {
		t.Log("unepxected error")
	}

	previousWarningMessage = ""
	EXPECT_NE(mockStruct{1, "same"}, mockStruct{1, "same"})
	if len(previousWarningMessage) == 0 {
		t.Log("unepxected error")
	}
}
