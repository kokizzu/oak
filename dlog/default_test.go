package dlog_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/oakmound/oak/v3/dlog"
)

func TestLogger(t *testing.T) {
	lgr := dlog.NewLogger()

	err := lgr.SetLogLevel(-1)
	if err == nil {
		t.Fatalf("expected -1 log level to error")
	}

	lgr.SetLogLevel(dlog.VERBOSE)

	var buff = new(bytes.Buffer)

	lgr.SetOutput(buff)
	// This function wrapper corrects the logged file generated
	calllogger := func() {
		lgr.Error("error")
		lgr.Info("info")
		lgr.Verb("verb")

		lgr.SetFilter(func(s string) bool { return strings.Contains(s, "foo") })
		lgr.Verb("bar")
		lgr.Verb("foo")
	}
	calllogger()

	expectedOut := `[default_test:34]  ERROR: error
[default_test:34]  INFO: info
[default_test:34]  VERBOSE: verb
[default_test:34]  VERBOSE: foo
`
	out := buff.String()

	if out != expectedOut {
		t.Fatalf("logged output did not match: got %q expected %q", out, expectedOut)
	}
}
