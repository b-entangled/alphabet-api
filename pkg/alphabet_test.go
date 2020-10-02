package alphabet

import (
	"testing"
)

func TestValidate(t *testing.T) {
	var testCases = []struct {
		Name     string
		Input    string
		Expected bool
	}{
		{
			Name:     "test case 1",
			Input:    "this is my first test case",
			Expected: false,
		},
		{
			Name:     "test case 2",
			Input:    "&823%& this ^&# is my !@$& second test *&^ case with ++==** special characters",
			Expected: false,
		},
		{
			Name:     "test case 3",
			Input:    "&823%& this ^&# is my !@$& (qwertyuiopasdfghjklzxcvbnm) second test *&^ case with ++==** special characters",
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Logf("\nrunning ==== %v\n", testCase.Name)
		result := Validate(testCase.Input)
		if result != testCase.Expected {
			t.Errorf("\nValidate : Expected %v, Actual %v\n", testCase.Expected, result)
		}
	}
}
