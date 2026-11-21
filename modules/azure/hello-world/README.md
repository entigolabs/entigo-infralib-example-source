## Dummy module for testing ##
The Azure twin of aws/hello-world: no providers and no resources, it only outputs "${var.greeting}, ${var.prefix}!" so an Azure environment has a terraform module of this repository to test. Its test reads the output from the agent's state container in Azure Blob Storage through the framework's azure package.

### Example code ###

```
    modules:
      - name: hello
        source: azure/hello-world

```

The `greeting` variable (default `Hello`) sets the word the output greets with.
