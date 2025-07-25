## apigeecli environments traceconfig overrides create

Create a new Distributed Trace config override

### Synopsis

Create a new Distributed Trace config override

```
apigeecli environments traceconfig overrides create [flags]
```

### Options

```
  -p, --endpoint-uri string   Trace endpoint, used only with JAEGER
  -x, --exporter string       Trace exporter can be JAEGER or CLOUD_TRACE
  -h, --help                  help for create
  -n, --name string           API Proxy name
      --rate string           Sampler Rate
  -s, --sampler string        Sampler can be set to PROBABILITY or OFF (default "PROBABILITY")
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --api api          Sets the control plane API. Must be one of prod, autopush or staging; default is prod
      --default-token    Use Google default application credentials access token
      --disable-check    Disable check for newer versions
  -e, --env string       Apigee environment name
      --metadata-token   Metadata OAuth2 access token
      --no-output        Disable printing all statements to stdout
      --no-warnings      Disable printing warnings to stderr
  -o, --org string       Apigee organization name
      --print-output     Control printing of info log statements (default true)
  -r, --region string    Apigee control plane region name; default is https://apigee.googleapis.com
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli environments traceconfig overrides](apigeecli_environments_traceconfig_overrides.md)	 - Manage Distributed Trace config overrides for the environment

###### Auto generated by spf13/cobra on 1-Jul-2025
