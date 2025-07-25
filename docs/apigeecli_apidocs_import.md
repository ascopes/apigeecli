## apigeecli apidocs import

Import from a folder containing apidocs

### Synopsis

Import from a folder containing apidocs

```
apigeecli apidocs import [flags]
```

### Examples

```
Assume data is exported from org1 which contains siteid site1. The fles are exported as
	site_org1-site1.json and apidocs_org1-site1_00000.json

	When importing this data into a new org, the siteid changes (since it is a combination of
	org name and siteid).

	Now import the data to org2 as
apigeecli apidocs import -f samples/apidocs  -s $siteId
```

### Options

```
  -f, --folder string           Folder containing site_<siteid>.json and apidocs_<siteid>_<id>.json files
  -h, --help                    help for import
      --use-src-siteid string   Use source siteid when importing; useful whem importing data between two orgs
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
  -s, --siteid string    Name or siteid of the portal
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli apidocs](apigeecli_apidocs.md)	 - Manage Apigee API catalog item through ApiDoc

###### Auto generated by spf13/cobra on 1-Jul-2025
