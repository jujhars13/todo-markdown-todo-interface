package scan

import (
	"reflect"
	"testing"
)

func TestStringForReadmeItems(t *testing.T) {

	// see cases_test.go for test cases struct
	for _, tc := range stringForTodoItemsTestCases {
		observed := StringForTodoItems(tc.testMarkdownString)
		if !reflect.DeepEqual(tc.expected, observed) {
			t.Fatalf("ScanFileForReadmeItems= %v, wanted %v in testString \n %v", observed, tc.expected, tc.testMarkdownString)
		}
	}
}
