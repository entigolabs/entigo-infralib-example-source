package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/entigolabs/entigo-infralib-test/aws"
	"github.com/entigolabs/entigo-infralib-test/env"
	"github.com/entigolabs/entigo-infralib-test/tf"
)

func TestHelloWorld(t *testing.T) {
	for _, e := range env.Selected(t) {
		t.Run(e.Name, func(t *testing.T) {
			t.Parallel()
			// The agent names the module <prefix>-<step>-<module>; in a pull
			// request the step is a per-branch one, which Placement reports.
			p := env.ModulePlacement(t, e)
			outputs := tf.Get(t, e)
			want := fmt.Sprintf("Hello, %s-%s-%s!", e.Prefix, p.Step.Name, p.Module.Name)
			assert.Equal(t, want, outputs.String(t, "hello-world__hello_world"))
		})
	}
}
