package scanString

import (
	"reflect"
	"testing"
)

func TestScanFileForReadmeItems(t *testing.T) {

	// see cases_test for test cases struct
	for _, tc := range testCases {
		observed := ScanStringForReadmeItems(tc.testMarkdownString)
		if !reflect.DeepEqual(tc.expected, observed) {
			t.Fatalf("ScanFileForReadmeItems= %v, wanted %v in testString \n %v", observed, tc.expected, tc.testMarkdownString)
		}
	}
}
