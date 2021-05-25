package scanfile

import (
	"reflect"
	"testing"
)

func TestScanFileForReadmeItems(t *testing.T) {

	// see cases_test for test cases struct
	for _, tc := range testCases {
		observed := ScanFileForReadmeItems(tc.testMarkdownString)
		if !reflect.DeepEqual(tc.expected, observed) {
			t.Fatalf("ScanFileForReadmeItems= %v, wanted %v", observed, tc.expected)
		}
	}
}
