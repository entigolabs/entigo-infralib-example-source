package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/entigolabs/entigo-infralib-test/env"
	"github.com/entigolabs/entigo-infralib-test/k8s"
)

func TestHelloWorld(t *testing.T) {
	for _, e := range env.Selected(t) {
		t.Run(e.Name, func(t *testing.T) {
			t.Parallel()
			// Connect picks the environment's kube_context and the namespace
			// of this module's ArgoCD application.
			c := k8s.Connect(t, e)
			k8s.WaitUntilDeploymentAvailable(t, c, c.Namespace, 20, 6*time.Second)

			// Reachable through the external gateway, with the address pinned so
			// the check does not wait for public DNS.
			gateway := e.Gateway(t, "external")
			err := k8s.WaitUntilHostnameAvailable(t, c, gateway, "https://"+gateway.Hostname(c.Namespace), "200", gateway.Retries, 6*time.Second)
			require.NoError(t, err)
		})
	}
}
