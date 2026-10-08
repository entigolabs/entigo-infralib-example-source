# hello-world-pinned

A copy of [hello-world](../hello-world) whose `test/module.yaml` sets `pin_step: true`.

Nothing about this chart prevents two installations on one cluster. It exists to show the framework's handling of modules that cannot have a branch copy, such as external-dns or argocd: a pull request on this module replaces the regular `hello-world-pinned-<prefix>` application instead of installing a branch-prefixed copy next to it, and the pull request pipeline waits for every other run before doing so. A merge of an unrelated module leaves it as the pull request deployed it.

Pinned-wait scenario one: two pull requests on this module opened together; the second must queue behind the first.
