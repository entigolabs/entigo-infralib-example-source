# entigo-infralib-example-source

The smallest complete repository of [Entigo Infralib](https://github.com/entigolabs/entigo-infralib) modules, tested and released with [entigo-infralib-test](https://github.com/entigolabs/entigo-infralib-test). Copy it to start a module repository of your own, or read it to see how the pieces fit.

It holds five modules, tested on four environments across three clouds:

- `modules/aws/hello-world`, an OpenTofu module with one output. It greets with "Hello" on `aws_exbiz` and, through its input file, with "Tere" on `aws_expri`.
- `modules/google/hello-world`, its Google twin, on `google_exbiz`; its test reads the output from the agent's state bucket in Cloud Storage.
- `modules/oracle/hello-world`, its Oracle twin, on `oracle_exbiz`; its test reads the output from the agent's state bucket in OCI Object Storage.
- `modules/k8s/hello-world`, a Helm chart. Public through the platform's gateway on `aws_exbiz` and `google_exbiz`, internal on `oracle_exbiz` (the EU Sovereign Cloud realm has no public DNS, so the platform has only an internal gateway behind the OCI native ingress controller), not exposed on `aws_expri`.
- `modules/k8s/hello-world-pinned`, a copy of hello-world with `pin_step: true`, showing how modules that cannot have a branch copy are handled.

Each module has one input file and one test function per environment (`env.RunEach`), and the environments' tests run in parallel. That is the shape entigo-infralib's own biz and pri scenarios take.

Two modules do not make a platform. Each file under `environments/` is therefore a complete agent configuration that lists the network, cluster, ArgoCD, gateway and DNS modules of the released [entigo-infralib](https://github.com/entigolabs/entigo-infralib-release) beside this repository's own, with the external modules' inputs inline. The agent provisions all of it; this repository's modules are tested on top.

An environment file is named `<cloud>_<prefix>.yaml`. The environments live in the entigo-infralib test accounts next to its own biz and pri, so they are `aws_exbiz`, `aws_expri`, `google_exbiz` and `oracle_exbiz`. Region, project, zone and compartment come from `AWS_REGION`, `GOOGLE_PROJECT`, `GOOGLE_REGION`, `GOOGLE_ZONE`, `OCI_REGION` and `OCI_COMPARTMENT_ID`, as for the agent.

Releases of this repository are its git tags and GitHub releases; they are also published to [entigo-infralib-example-release](https://github.com/entigolabs/entigo-infralib-example-release), the way entigo-infralib releases go to entigo-infralib-release, to exercise that optional path. OCI publishing follows.

## Layout

```
environments/aws_exbiz.yaml    agent configuration of each environment
environments/aws_expri.yaml
environments/google_exbiz.yaml
environments/oracle_exbiz.yaml
modules/aws/hello-world/       module + test.sh + test/aws_{exbiz,expri}.yaml (inputs) + test/*_test.go
modules/google/hello-world/    module + test.sh + test/google_exbiz.yaml + test/*_test.go
modules/oracle/hello-world/    module + test.sh + test/oracle_exbiz.yaml + test/*_test.go
modules/k8s/hello-world/       chart  + test.sh + test/{aws_exbiz,aws_expri,google_exbiz,oracle_exbiz}.yaml (inputs) + test/*_test.go
modules/k8s/hello-world-pinned/ the same, plus test/module.yaml with pin_step
go.mod                         requires github.com/entigolabs/entigo-infralib-test
test.sh                        bootstrap of the orchestrator, pins the framework version
release_version.txt            major.minor of the next release
.github/workflows/             pull-request, stable and release, each calling a reusable workflow of the framework
```

## Running

With AWS credentials and `AWS_REGION=eu-north-1` in your shell, for Google `GOOGLE_APPLICATION_CREDENTIALS` (or a gcloud login) with `GOOGLE_PROJECT`, `GOOGLE_REGION` and `GOOGLE_ZONE`, and for Oracle an `OCI_CONFIG_FILE` (or `~/.oci/config`) with `OCI_REGION` and `OCI_COMPARTMENT_ID`; the clouds with credentials are the ones that run:

```
./test.sh                          provision every environment with credentials and test all modules on them
./test.sh --env aws_expri          one environment only
./test.sh modules/k8s/hello-world  test one module in a step of its own
modules/k8s/hello-world/test.sh    the same, from the module's directory
./test.sh --help
```

The kubeconfig is yours to provide, for example `aws eks update-kubeconfig --region eu-north-1 --name exbiz-infra-eks` or `gcloud container clusters get-credentials exbiz-infra-gke --region europe-north1`; the test picks the context that command creates for the environment's cluster module. For OKE it is `oci ce cluster create-kubeconfig --cluster-id <id of exbiz-infra-oke> --token-version 2.0.0 --kube-endpoint PUBLIC_ENDPOINT`; the test matches the `context-exbiz-infra-oke-...` context that creates.

## Pipelines

- **Pull request**: every module a pull request changes is applied in a step of its own on every environment of its cloud and tested. The step stays until the environment is nuked. Needs the `AWS_*`, `GOOGLE_*` and `OCI_*` secrets.
- **Stable** (weekday mornings, or by hand): provisions every environment from the latest release in [entigo-infralib-example-release](https://github.com/entigolabs/entigo-infralib-example-release) and runs that release's tests.
- **Release** (after a green Stable, or by hand): applies `main` to every environment, tests, and when main is ahead of the latest release tags it, creates the GitHub release, publishes the charts, modules and a signed index as OCI packages to ghcr.io mirrored to ECR Public, and publishes `modules/` to [entigo-infralib-example-release](https://github.com/entigolabs/entigo-infralib-example-release). Needs `SSH_PRIVATE_KEY`, a deploy key with write access there.

## Status

All environments provision and all tests pass, from a workstation and from the pipelines.

The release notes list the modules changed since the previous release and the other paths that changed.
