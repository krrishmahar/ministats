package test

import (
	"bytes"
	"testing"

	"github.com/krrishmahar/ministats/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCpuCmd(t *testing.T)  {
	var buf bytes.Buffer
	root := cmd.NewRootCmd()
	root.SetOut(&buf)
	root.SetArgs([] string{"cpu"})

	err := root.Execute()
	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "Total CPU Usage")
}
