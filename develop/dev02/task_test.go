package main

import "testing"

func TestUnpack(t *testing.T) {
	testData := []struct {
		name          string
		testValue     string
		expectedValue string
	}{
		{
			name:          "test1",
			testValue:     "a4bc2d5e",
			expectedValue: "aaaabccddddde",
		},
		{
			name:          "test2",
			testValue:     "abcd",
			expectedValue: "abcd",
		},
		{
			name:          "test3",
			testValue:     "45",
			expectedValue: "некорректная строка",
		},
		{
			name:          "test4",
			testValue:     "",
			expectedValue: "",
		},
		{
			name:          "test5",
			testValue:     `qwe\4\5`,
			expectedValue: "qwe45",
		},
		{
			name:          "test6",
			testValue:     `qwe\45`,
			expectedValue: "qwe44444",
		},
		{
			name:          "test7",
			testValue:     "qwe\\\\5",
			expectedValue: "qwe\\\\\\\\\\",
		},
	}

	for _, x := range testData {
		t.Run(x.name, func(t *testing.T) {
			result := Unpack(x.testValue)
			if result != x.expectedValue {
				t.Errorf("got %s, expected %s", result, x.expectedValue)
			}
		})
	}
}
