## apigeecli environments archives create

Create a new revision of archive in the environment

### Synopsis

Create a new revision of archive in the environment

```
apigeecli environments archives create [flags]
```

### Options

```
  -f, --folder string    Archive Folder
  -h, --help             help for create
  -n, --name string      Archive name
  -w, --wait             Wait for deployment
  -z, --zipfile string   Archive Zip file
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

* [apigeecli environments archives](apigeecli_environments_archives.md)	 - Manage archive deployments for the environment

###### Auto generated by spf13/cobra on 1-Jul-2025
