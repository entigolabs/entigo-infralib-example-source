# entigo-infralib-example-source

The smallest complete repository of [Entigo Infralib](https://github.com/entigolabs/entigo-infralib) modules, tested and released with [entigo-infralib-test](https://github.com/entigolabs/entigo-infralib-test). Copy it to start a module repository of your own, or read it to see how the pieces fit.

It holds two modules, tested on two environments:

- `modules/aws/hello-world`, an OpenTofu module with one output. It greets with "Hello" on `aws_exbiz` and, through its input file, with "Tere" on `aws_expri`.
- `modules/k8s/hello-world`, a Helm chart. Public through the platform's gateway on `aws_exbiz`, internal only on `aws_expri`.

Each module has one input file and one test function per environment (`env.RunEach`), and the environments' tests run in parallel. That is the shape entigo-infralib's own biz and pri scenarios take.

Two modules do not make a platform. Each file under `environments/` is therefore a complete agent configuration that lists the network, cluster, ArgoCD, gateway and DNS modules of the released [entigo-infralib](https://github.com/entigolabs/entigo-infralib-release) beside this repository's own, with the external modules' inputs inline. The agent provisions all of it; this repository's modules are tested on top.

An environment file is named `<cloud>_<prefix>.yaml`. The environments live in the entigo-infralib AWS test account next to its own biz and pri, so they are `aws_exbiz` and `aws_expri`. The region comes from `AWS_REGION`, as for the agent.

Releases of this repository go to [entigo-infralib-example-release](https://github.com/entigolabs/entigo-infralib-example-release), the way entigo-infralib releases go to entigo-infralib-release.

## Layout

```
environments/aws_exbiz.yaml    agent configuration of each environment
environments/aws_expri.yaml
modules/aws/hello-world/       module + test.sh + test/aws_{exbiz,expri}.yaml (inputs) + test/*_test.go
modules/k8s/hello-world/       chart  + test.sh + test/aws_{exbiz,expri}.yaml (inputs) + test/*_test.go
go.mod                         requires github.com/entigolabs/entigo-infralib-test
test.sh                        bootstrap of the orchestrator, pins the framework version
.github/workflows/             calls the framework's reusable workflows
```

## Running

With AWS credentials and `AWS_REGION=eu-north-1` in your shell:

```
./test.sh                          provision aws_exbiz and aws_expri and test both modules on both
./test.sh --env aws_expri          one environment only
./test.sh modules/k8s/hello-world  test one module in a step of its own
modules/k8s/hello-world/test.sh    the same, from the module's directory
./test.sh --help
```

The kubeconfig is yours to provide, for example `aws eks update-kubeconfig --region eu-north-1 --name exbiz-infra-eks`; the test picks the context that command creates for the environment's `aws/eks` module.

## Status

Scaffolding, not yet run against AWS. The first real run will settle the external module inputs under `environments/`.
