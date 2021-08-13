package migrate

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestProcessLine(t *testing.T) {
	table := []struct {
		input  string
		output string
	}{
		{"c.Assert(stmtEvictedElement.beginTime, Equals, now)",
			"require.Equal(t, now, stmtEvictedElement.beginTime)",
		},
		{`c.Assert(getAllEvicted(ssbde), Equals, "{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}")`,
			`require.Equal(t, "{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}", getAllEvicted(ssbde))`,
		},
		{"c.Assert(stmtEvictedElement.beginTime, Greater, now)",
			"require.Greater(t, stmtEvictedElement.beginTime, now)",
		},
		{
			`func (s *testRangerSuite) TestTableRange(c *C) {`,
			`func TestTableRange(t *testing.T) {`,
		},
	}

	for _, v := range table {
		output, err := ProcessLine(v.input)
		require.NoError(t, err)
		require.Equal(t, v.output, strings.TrimSpace(output))
	}
}
