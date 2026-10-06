package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "github.com/entigolabs/entigo-infralib-test/aws"
	"github.com/entigolabs/entigo-infralib-test/env"
	"github.com/entigolabs/entigo-infralib-test/tf"
)

// One test per environment; they run in parallel.
func TestHelloWorld(t *testing.T) {
	env.RunEach(t, map[string]env.TestFunc{
		"aws_biz": testHelloWorldBiz,
		"aws_pri": testHelloWorldPri,
	})
}

func testHelloWorldBiz(t *testing.T, e *env.Environment) {
	assert.Equal(t, greeting(t, e, "Hello"), tf.Get(t, e).String(t, "hello-world__hello_world"))
}

// aws_pri.yaml sets greeting: Tere.
func testHelloWorldPri(t *testing.T, e *env.Environment) {
	assert.Equal(t, greeting(t, e, "Tere"), tf.Get(t, e).String(t, "hello-world__hello_world"))
}

// greeting is what the module outputs: the agent names a module instance
// <prefix>-<step>-<module>, and in a pull request the step is a per-branch
// one, which ModulePlacement reports.
func greeting(t *testing.T, e *env.Environment, word string) string {
	p := env.ModulePlacement(t, e)
	return fmt.Sprintf("%s, %s-%s-%s!", word, e.Prefix, p.Step.Name, p.Module.Name)
}
