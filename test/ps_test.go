package test

import (
	"bytes"
	"testing"

	"github.com/krrishmahar/ministats/cmd"
	"github.com/stretchr/testify/assert"
)

func TestPsSortCpu(t *testing.T) {
	var buf bytes.Buffer
	root := cmd.NewRootCmd()
	root.SetOut(&buf)
	root.SetArgs([]string{"ps", "--sort=cpu"})

	err := root.Execute()
	assert.NoError(t, err)
	assert.Contains(t, buf.String(), "PID")
}
