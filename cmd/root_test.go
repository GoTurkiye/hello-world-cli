package cmd

import (
	"bytes"
	"testing"
)

func TestRootCmda(t *testing.T) {
	buf := bytes.NewBufferString("")
	rootCmd.SetOut(buf)

	rootCmd.Execute()

	out := buf.String()
	if out != "Merhaba Go Türkiye!!!" {
		t.Errorf("it should be equal %s == Merhaba Go Türkiye!!", out)
	}
}
