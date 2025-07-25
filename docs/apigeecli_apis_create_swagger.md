## apigeecli apis create swagger

Creates an API proxy from a Swagger Spec

### Synopsis

Creates an API proxy from a Swagger Spec for Cloud Endpoints/API Gateway

```
apigeecli apis create swagger [flags]
```

### Options

```
      --add-cors             Add a CORS policy
  -d, --desc string          API Proxy description; Merges with the existing description in the spec.
  -h, --help                 help for swagger
      --import               Import API Proxy after generation from spec (default true)
  -n, --name string          API Proxy name. If not specified, will look for the x-google-api-name extension
      --space string         Apigee Space to associate to
  -f, --swaggerfile string   Path to a Swagger Specification file with API Gateway or Cloud Endpoints extensions
  -u, --swaggeruri string    URI to a Swagger Specification file with API Gateway or Cloud Endpoints extensions
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --api api          Sets the control plane API. Must be one of prod, autopush or staging; default is prod
      --default-token    Use Google default application credentials access token
      --disable-check    Disable check for newer versions
      --metadata-token   Metadata OAuth2 access token
      --no-output        Disable printing all statements to stdout
      --no-warnings      Disable printing warnings to stderr
  -o, --org string       Apigee organization name
      --print-output     Control printing of info log statements (default true)
  -r, --region string    Apigee control plane region name; default is https://apigee.googleapis.com
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli apis create](apigeecli_apis_create.md)	 - Creates an API proxy in an Apigee Org

###### Auto generated by spf13/cobra on 1-Jul-2025
