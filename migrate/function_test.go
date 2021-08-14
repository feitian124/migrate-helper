package migrate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFunction(t *testing.T) {
	table := []struct {
		line   string
		result *FunctionResult
	}{
		{"  func (s *testStmtSummarySuite) TestSimpleStmtSummaryByDigestEvicted(c *C) {  ",
			&FunctionResult{match: true, function: "func", caller: "(s *testStmtSummarySuite)", name: "TestSimpleStmtSummaryByDigestEvicted", rest: "(c *C) {"},
		},
		{
			"func (s *testStmtSummarySuite) TestNewStmtSummaryByDigestEvictedElement(c *C) {",
			&FunctionResult{match: true, function: "func", caller: "(s *testStmtSummarySuite)", name: "TestNewStmtSummaryByDigestEvictedElement", rest: "(c *C) {"},
		},
		{
			"func NewStmtSummaryByDigestEvictedElement(c *C) {",
			&FunctionResult{match: false, function: "", caller: "", name: "", rest: ""},
		},
	}

	for _, v := range table {
		x, err := Function(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}
