package file

import "testing"

func TestOpenSuccessCase(t *testing.T) {
	cases := []string{
		"./testdata/pdf-test.pdf",
	}
	for _, testCase := range cases {
		pdf := NewPdfFile(testCase)
		err := pdf.Open()
		if err != nil {
			t.Errorf("Opening pdf with path %s failed with error %v", testCase, err)
		}
	}
}
