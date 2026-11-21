package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/entigolabs/entigo-infralib-test/azure"
	"github.com/entigolabs/entigo-infralib-test/env"
	"github.com/entigolabs/entigo-infralib-test/tf"
)

// The output read back from the agent's state in Azure Blob Storage.
func TestHelloWorld(t *testing.T) {
	env.Run(t, func(t *testing.T, e *env.Environment) {
		p := env.ModulePlacement(t, e)
		want := fmt.Sprintf("Hello, %s-%s-%s!", e.Prefix, p.Step.Name, p.Module.Name)
		assert.Equal(t, want, tf.Get(t, e).String(t, "hello-world__hello_world"))
	})
}
