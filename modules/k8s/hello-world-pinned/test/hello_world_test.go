package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	kubernetesErrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/entigolabs/entigo-infralib-test/env"
	"github.com/entigolabs/entigo-infralib-test/k8s"
)

// The deployment must come up everywhere. The wait allows for a Spot node
// being replaced underneath the pods, which takes minutes with the image pull.
func TestDeployment(t *testing.T) {
	env.Run(t, func(t *testing.T, e *env.Environment) {
		c := k8s.Connect(t, e)
		k8s.WaitUntilDeploymentAvailable(t, c, c.Namespace, 60, 6*time.Second)
	})
}

// Exposure differs per environment: public on exbiz, internal only on expri.
func TestRoute(t *testing.T) {
	env.RunEach(t, map[string]env.TestFunc{
		"aws_exbiz": testRoutePublic,
		"aws_expri": testRouteAbsent,
	})
}

// aws_exbiz: the HTTPRoute is accepted and answers 200 through the gateway.
// Hostname, scheme and the gateway address come from the objects; the probe
// runs inside the cluster pinned to the address, so DNS need not have caught up.
func testRoutePublic(t *testing.T, e *env.Environment) {
	c := k8s.Connect(t, e)
	route := k8s.WaitUntilRouteReachable(t, c, c.Namespace, 100, 6*time.Second)
	require.False(t, route.Internal(), "%s is served by an internal load balancer", route.URL())
}

// aws_expri: deliberately not exposed, no HTTPRoute at all.
func testRouteAbsent(t *testing.T, e *env.Environment) {
	c := k8s.Connect(t, e)
	_, err := c.GetObjectE(k8s.HTTPRoutes, c.Namespace, c.Namespace)
	require.True(t, kubernetesErrors.IsNotFound(err), "expri must not expose hello-world-pinned, got: %v", err)
}
