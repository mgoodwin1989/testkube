## kubectl-testkube get webhook

Get webhook details

### Synopsis

Get webhook, you can change output format, to get single details pass name as first arg

```
kubectl-testkube get webhook <webhookName> [flags]
```

### Options

```
  -h, --help          help for webhook
  -n, --name string   unique webhook name, you can also pass it as argument
```

### Options inherited from parent commands

```
      --analytics-enabled    enable analytics (default true)
  -c, --client string        client used for connecting to Testkube API one of proxy|direct (default "proxy")
      --go-template string   go template to render (default "{{.}}")
  -s, --namespace string     Kubernetes namespace, default value read from config if set (default "testkube")
  -o, --output string        output type can be one of json|yaml|pretty|go-template (default "pretty")
  -v, --verbose              show additional debug messages
```

### SEE ALSO

* [kubectl-testkube get](kubectl-testkube_get.md)	 - Get resources

