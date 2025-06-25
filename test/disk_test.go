package test

import (
	"bytes"
	"testing"

	"github.com/krrishmahar/ministats/cmd"
	"github.com/stretchr/testify/assert"
)

func TestDiskCmd(t *testing.T)  {
	var buf bytes.Buffer
	root := cmd.NewRootCmd()
	root.SetOut(&buf)
	root.SetArgs([] string{"disk"})

	err := root.Execute()
	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "Disk Used")
}
