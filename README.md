# entigo-infralib-example-source

The smallest complete repository of [Entigo Infralib](https://github.com/entigolabs/entigo-infralib) modules, tested and released with [entigo-infralib-test](https://github.com/entigolabs/entigo-infralib-test). Copy it to start a module repository of your own, or read it to see how the pieces fit.

It holds two modules:

- `modules/aws/hello-world`, an OpenTofu module with one output,
- `modules/k8s/hello-world`, a Helm chart served through the platform's gateway.

Two modules do not make a platform. `environments.yaml` therefore lists the network, cluster, ArgoCD, gateway and DNS modules of the released [entigo-infralib](https://github.com/entigolabs/entigo-infralib-release) as *external* modules of the same steps, with their inputs under `environments/aws_demo/`. The agent provisions all of it with one configuration; this repository's modules are tested on top.

Releases of this repository go to [entigo-infralib-example-release](https://github.com/entigolabs/entigo-infralib-example-release), the way entigo-infralib releases go to entigo-infralib-release.

## Layout

```
environments.yaml              environments, steps, the external source
environments/aws_demo/         agent inputs of the external modules, per step
modules/aws/hello-world/       module + test/aws_demo.yaml (input) + test/*_test.go
modules/k8s/hello-world/       chart  + test/aws_demo.yaml (input) + test/*_test.go
go.mod                         requires github.com/entigolabs/entigo-infralib-test
test.sh                        bootstrap of the orchestrator, pins the framework version
.github/workflows/             calls the framework's reusable workflows
```

## Running

With AWS credentials in your shell:

```
./test.sh                          provision aws_demo and test both modules
./test.sh modules/k8s/hello-world  test one module in a step of its own
./test.sh --help
```

The kubeconfig is yours to provide, for example `aws eks update-kubeconfig --region eu-north-1 --name demo-infra-eks`; the test only picks the `kube_context` named in `environments.yaml`.

## Status

Scaffolding, not yet run against AWS. The first real run will settle the external module inputs under `environments/aws_demo/`.
