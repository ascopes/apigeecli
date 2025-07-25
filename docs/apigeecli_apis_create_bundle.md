## apigeecli apis create bundle

Creates an API proxy from an Zip or folder

### Synopsis

Creates an API proxy from an Zip or folder; Optionally deploy the API to an env

```
apigeecli apis create bundle [flags]
```

### Options

```
  -d, --desc string           API Proxy description; Appends to the description if the description already exists
  -e, --env string            Name of the environment to deploy the proxy
  -h, --help                  help for bundle
  -n, --name string           API Proxy name
      --ovr                   Forces deployment of the new revision
  -f, --proxy-folder string   Path to the Proxy Bundle; ex: ./test/apiproxy
  -p, --proxy-zip string      Path to the Proxy bundle/zip file
  -s, --sa string             The format must be {ACCOUNT_ID}@{PROJECT}.iam.gserviceaccount.com.
      --safedeploy            When set to true, generateDeployChangeReport will be executed and deployment will proceed if there are no conflicts
      --sequencedrollout      If set to true, the routing rules will be rolled out in a safe order; default is false
      --space string          Apigee Space to associate to
      --wait                  Waits for the deployment to finish, with success or error
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
