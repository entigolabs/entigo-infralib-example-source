## Dummy module for testing ##
The Google twin of aws/hello-world: no providers and no resources, it only outputs "${var.greeting}, ${var.prefix}!" so a Google environment has a terraform module of this repository to test. Its test reads the output from the agent's state bucket in Google Cloud Storage through the framework's google package.

### Example code ###

```
    modules:
      - name: hello
        source: google/hello-world

```

The `greeting` variable (default `Hello`) sets the word the output greets with.
