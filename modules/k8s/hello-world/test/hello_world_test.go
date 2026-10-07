package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	kubernetesErrors "k8s.io/apimachinery/pkg/api/errors"

	_ "github.com/entigolabs/entigo-infralib-test/aws" // k8s.Gateway reads the route53 outputs through the aws package
	"github.com/entigolabs/entigo-infralib-test/env"
	"github.com/entigolabs/entigo-infralib-test/k8s"
)

// One test per environment; they run in parallel.
func TestHelloWorld(t *testing.T) {
	env.RunEach(t, map[string]env.TestFunc{
		"aws_exbiz": testHelloWorldPublic,
		"aws_expri": testHelloWorldInternal,
	})
}

// The deployment wait allows for a Spot node being replaced underneath the
// pods, which takes several minutes including the image pull.

// aws_exbiz: deployed and reachable through the external gateway (derived
// from the aws-alb and route53 modules of the environment), with the gateway
// address pinned so the check does not wait for public DNS.
func testHelloWorldPublic(t *testing.T, e *env.Environment) {
	c := k8s.Connect(t, e)
	k8s.WaitUntilDeploymentAvailable(t, c, c.Namespace, 60, 6*time.Second)

	_, err := k8s.WaitUntilK8SHTTPRouteAvailable(t, c, c.Namespace, 20, 6*time.Second)
	require.NoError(t, err, "HTTPRoute %s not accepted", c.Namespace)

	gateway := k8s.Gateway(t, e, "external")
	err = k8s.WaitUntilHostnameAvailable(t, c, gateway, "https://"+gateway.Hostname(c.Namespace), "200", gateway.Retries, 6*time.Second)
	require.NoError(t, err)
}

// aws_expri: deployed, and deliberately not exposed.
func testHelloWorldInternal(t *testing.T, e *env.Environment) {
	c := k8s.Connect(t, e)
	k8s.WaitUntilDeploymentAvailable(t, c, c.Namespace, 60, 6*time.Second)

	_, err := c.GetObjectE(k8s.HTTPRoutes, c.Namespace, c.Namespace)
	require.True(t, kubernetesErrors.IsNotFound(err), "pri must not expose hello-world, got HTTPRoute lookup result: %v", err)
}
