## apigeecli securityprofiles listrevisions

Returns the revisions of a security profile

### Synopsis

Returns the revisions of a security profile

```
apigeecli securityprofiles listrevisions [flags]
```

### Options

```
  -h, --help          help for listrevisions
  -n, --name string   Name of the security profile
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

* [apigeecli securityprofiles](apigeecli_securityprofiles.md)	 - Manage Adv API Security Profiles

###### Auto generated by spf13/cobra on 1-Jul-2025
