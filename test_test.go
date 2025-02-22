package partitionlabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var initString = "ababcbacadefegdehijhklij" // Input array
	var expectedResult = []int{9, 7, 8}         // The expected answer with correct length.
	res := partitionLabels(initString)          // Calls your implementation

	assert.Equal(t, expectedResult, res)
}
func partitionLabels(s string) []int {
	lastOccurrence := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		lastOccurrence[s[i]] = i
	}
	var result []int
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if lastOccurrence[s[i]] > end {
			end = lastOccurrence[s[i]]
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}
