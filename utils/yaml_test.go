package utils_test

import (
	"testing"

	"github.com/Jasrags/ShadowMUD/utils"
	"github.com/stretchr/testify/assert"
)

func TestFormatFilename(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{
			filename: "My File Name.txt",
			expected: "my_file_name.txt",
		},
		{
			filename: "Another File: Name.docx",
			expected: "another_file__name.docx",
		},
		{
			filename: "File/Path/With:Special*Characters?\"<>|.txt",
			expected: "file_path_with_special_characters_____.txt",
		},
	}

	for _, tc := range testCases {
		result := utils.FormatFilename(tc.filename)
		assert.Equal(t, tc.expected, result)
	}
}
