package test

import (
	"bytes"
	"testing"

	"github.com/krrishmahar/ministats/cmd"
	"github.com/stretchr/testify/assert"
)

func TestSystemInfoCmd(t *testing.T)  {
	var buf bytes.Buffer
	root := cmd.NewRootCmd()
	root.SetOut(&buf)
	root.SetArgs([] string{"sysinfo"})

	err := root.Execute()
	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "OS")
	assert.Contains(t, buf.String(), "Uptime")
	assert.Contains(t, buf.String(), "User")
}
