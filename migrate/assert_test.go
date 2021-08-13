package migrate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEquals(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(stmtEvictedElement.beginTime, Equals, now)",
			&AssertResult{match: true, caller: "c", actual: "stmtEvictedElement.beginTime", checker: "Equals", expect: "now"},
		},
		{"c.Assert(stmtEvictedElement.beginTime, Greater, now)",
			&AssertResult{match: false, caller: "", actual: "", checker: "", expect: ""},
		},
		{"  c.Assert(stmtEvictedElement.beginTime(), Equals, now1) ",
			&AssertResult{match: true, caller: "c", actual: "stmtEvictedElement.beginTime()", checker: "Equals", expect: "now1"},
		},
		{`c.Assert(getAllEvicted(ssbde), Equals, "{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}")`,
			&AssertResult{match: true, caller: "c", actual: "getAllEvicted(ssbde)", checker: "Equals", expect: `"{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}"`},
		},
	}

	for _, v := range table {
		x, err := Equals(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

func TestDeepEquals(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(stmtEvictedElement.beginTime, DeepEquals, now)",
			&AssertResult{match: true, caller: "c", actual: "stmtEvictedElement.beginTime", checker: "DeepEquals", expect: "now"},
		},
		{"c.Assert(stmtEvictedElement.beginTime, Greater, now)",
			&AssertResult{match: false, caller: "", actual: "", checker: "", expect: ""},
		},
		{"  c.Assert(stmtEvictedElement.beginTime(), DeepEquals, now1) ",
			&AssertResult{match: true, caller: "c", actual: "stmtEvictedElement.beginTime()", checker: "DeepEquals", expect: "now1"},
		},
		{`c.Assert(getAllEvicted(ssbde), DeepEquals, "{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}")`,
			&AssertResult{match: true, caller: "c", actual: "getAllEvicted(ssbde)", checker: "DeepEquals", expect: `"{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}"`},
		},
	}

	for _, v := range table {
		x, err := DeepEquals(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

func TestIsNil(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(err, IsNil)",
			&AssertResult{match: true, caller: "c", actual: "err", checker: "IsNil", expect: ""},
		},
	}

	for _, v := range table {
		x, err := IsNil(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

func TestNotNil(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(err, NotNil)",
			&AssertResult{match: true, caller: "c", actual: "err", checker: "NotNil", expect: ""},
		},
	}

	for _, v := range table {
		x, err := NotNil(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

func TestIsTrue(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(os.IsNotExist(err), IsTrue)",
			&AssertResult{match: true, caller: "c", actual: "os.IsNotExist(err)", checker: "IsTrue", expect: ""},
		},
		{`c.Assert(tk.HasPlan(queryList, "Batch_Point_Get"), IsTrue) `,
			&AssertResult{match: true, caller: "c", actual: `tk.HasPlan(queryList, "Batch_Point_Get")`, checker: "IsTrue"},
		},
	}

	for _, v := range table {
		x, err := IsTrue(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

func TestIsFalse(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(os.IsNotExist(err), IsFalse)",
			&AssertResult{match: true, caller: "c", actual: "os.IsNotExist(err)", checker: "IsFalse", expect: ""},
		},
		{`c.Assert(tk.HasPlan(queryList, "Batch_Point_Get"), IsFalse) `,
			&AssertResult{match: true, caller: "c", actual: `tk.HasPlan(queryList, "Batch_Point_Get")`, checker: "IsFalse"},
		},
	}

	for _, v := range table {
		x, err := IsFalse(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

func TestGreater(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{" c.Assert(tk.Se.GetSessionVars().StmtCtx.MemTracker.MaxConsumed(), Greater, int64(0)) ",
			&AssertResult{match: true, caller: "c", actual: "tk.Se.GetSessionVars().StmtCtx.MemTracker.MaxConsumed()", checker: "Greater", expect: "int64(0)"},
		},
	}

	for _, v := range table {
		x, err := Greater(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}