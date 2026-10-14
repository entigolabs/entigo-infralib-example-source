## Dummy module for testing ##
The Oracle twin of aws/hello-world: no providers and no resources, it only outputs "${var.greeting}, ${var.prefix}!" so an Oracle environment has a terraform module of this repository to test. Its test reads the output from the agent's state bucket in OCI Object Storage through the framework's oracle package.

### Example code ###

```
    modules:
      - name: hello
        source: oracle/hello-world

```

The `greeting` variable (default `Hello`) sets the word the output greets with.
