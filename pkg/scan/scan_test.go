package scan

import (
	"reflect"
	"testing"
)

func TestScanStringForReadmeItems(t *testing.T) {

	// see cases_test.go for test cases struct
	for _, tc := range scanStringForReadmeItemsTestCases {
		observed := ScanStringForReadmeItems(tc.testMarkdownString)
		if !reflect.DeepEqual(tc.expected, observed) {
			t.Fatalf("ScanFileForReadmeItems= %v, wanted %v in testString \n %v", observed, tc.expected, tc.testMarkdownString)
		}
	}
}

// see https://stackoverflow.com/a/67012273/341156
func TestGetStringAtSpecificLineInFile(t *testing.T) {

}
