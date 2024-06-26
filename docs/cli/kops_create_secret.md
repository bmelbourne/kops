
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops create secret

Create a secret.

### Options

```
  -h, --help   help for secret
```

### Options inherited from parent commands

```
      --config string   yaml config file (default is $HOME/.kops.yaml)
      --name string     Name of cluster. Overrides KOPS_CLUSTER_NAME environment variable
      --state string    Location of state storage (kops 'config' file). Overrides KOPS_STATE_STORE environment variable
  -v, --v Level         number for the log level verbosity
```

### SEE ALSO

* [kops create](kops_create.md)	 - Create a resource by command line, filename or stdin.
* [kops create secret ciliumpassword](kops_create_secret_ciliumpassword.md)	 - Create a Cilium IPsec configuration.
* [kops create secret dockerconfig](kops_create_secret_dockerconfig.md)	 - Create a Docker config.
* [kops create secret encryptionconfig](kops_create_secret_encryptionconfig.md)	 - Create an encryption config.

