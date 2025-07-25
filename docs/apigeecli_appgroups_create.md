## apigeecli appgroups create

Create an AppGroup

### Synopsis

Create an AppGroup

```
apigeecli appgroups create [flags]
```

### Options

```
      --attrs stringToString   Custom attributes (default [])
  -i, --channelid string       channel identifier identifies the owner maintaining this grouping
  -u, --channelurl string      A reference to the associated storefront/marketplace
      --devs stringToString    Developer details; set developerid=role. NOTE: only 1 role is supported (default [])
  -d, --display-name string    app group name displayed in the UI
  -h, --help                   help for create
  -n, --name string            Name of the AppGroup
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

* [apigeecli appgroups](apigeecli_appgroups.md)	 - Manage Apigee Application Groups

###### Auto generated by spf13/cobra on 1-Jul-2025
