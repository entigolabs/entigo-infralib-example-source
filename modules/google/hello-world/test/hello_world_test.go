package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/entigolabs/entigo-infralib-test/env"
	_ "github.com/entigolabs/entigo-infralib-test/google"
	"github.com/entigolabs/entigo-infralib-test/tf"
)

// The output read back from the agent's state in Google Cloud Storage.
func TestHelloWorld(t *testing.T) {
	env.Run(t, func(t *testing.T, e *env.Environment) {
		p := env.ModulePlacement(t, e)
		want := fmt.Sprintf("Hello, %s-%s-%s!", e.Prefix, p.Step.Name, p.Module.Name)
		assert.Equal(t, want, tf.Get(t, e).String(t, "hello-world__hello_world"))
	})
}
