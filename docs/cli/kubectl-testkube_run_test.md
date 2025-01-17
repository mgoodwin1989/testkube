## kubectl-testkube run test

Starts new test

### Synopsis

Starts new test based on Test Custom Resource name, returns results to console

```
kubectl-testkube run test <testName> [flags]
```

### Options

```
      --args stringArray        executor binary additional arguments
  -a, --download-artifacts      downlaod artifacts automatically
      --download-dir string     download dir (default "artifacts")
  -h, --help                    help for test
  -n, --name string             execution name, if empty will be autogenerated
  -p, --param stringToString    execution envs passed to executor (default [])
      --params-file string      params file path, e.g. postman env file - will be passed to executor if supported
      --secret stringToString   secret envs in a form of secret_name1=secret_key1 passed to executor (default [])
  -f, --watch                   watch for changes after start
```

### Options inherited from parent commands

```
      --analytics-enabled   enable analytics (default true)
  -c, --client string       client used for connecting to Testkube API one of proxy|direct (default "proxy")
  -s, --namespace string    Kubernetes namespace, default value read from config if set (default "testkube")
  -v, --verbose             show additional debug messages
```

### SEE ALSO

* [kubectl-testkube run](kubectl-testkube_run.md)	 - Runs tests or test suites

