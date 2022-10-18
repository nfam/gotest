package gotest

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Equal writes the difference to console if the actual result does not equal
// the expected one.
func Equal(t *testing.T, actual interface{}, expect interface{}, samples ...interface{}) {
	t.Helper()
	if diff := cmp.Diff(expect, actual); diff != "" {
		for _, s := range samples {
			d, _ := json.MarshalIndent(s, "", "    ")
			diff += "sample\t" + string(d)
		}
		t.Error(diff)
	}
}

// Error writes log to console if the actual result does not equal
// the expected one.
func Error(t *testing.T, actual error, expect string, samples ...interface{}) {
	t.Helper()
	var b strings.Builder

	if actual == nil {
		if expect != "" {
			fmt.Fprintf(&b, "expect\t%q\nreturn\tnil", expect)
		}
	} else if expect == "" {
		fmt.Fprintf(&b, "expect\tnil\nreturn\t%q", actual.Error())
	} else if actual.Error() != expect {
		fmt.Fprintf(&b, "expect\t%q\nreturn\t%q", expect, actual.Error())
	}

	if b.Len() > 0 {
		for _, s := range samples {
			d, _ := json.MarshalIndent(s, "", "    ")
			b.WriteString("\nsample\t")
			b.Write(d)
		}
		t.Error("\n" + b.String())
	}
}

// True writes log to console if the value is not true.
func True(t *testing.T, value bool, samples ...interface{}) {
	t.Helper()
	Equal(t, value, true, samples...)
}

// False writes log to console if the value is not false.
func False(t *testing.T, value bool, samples ...interface{}) {
	t.Helper()
	Equal(t, value, false, samples...)
}
