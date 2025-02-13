
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops

kOps is Kubernetes Operations.

### Synopsis

kOps is Kubernetes Operations.

 kOps is the easiest way to get a production grade Kubernetes cluster up and running. We like to think of it as kubectl for clusters.

 kOps helps you create, destroy, upgrade and maintain production-grade, highly available, Kubernetes clusters from the command line. AWS (Amazon Web Services) is currently officially supported, with Digital Ocean and OpenStack in beta support.

### Options

```
      --config string   yaml config file (default is $HOME/.kops.yaml)
  -h, --help            help for kops
      --name string     Name of cluster. Overrides KOPS_CLUSTER_NAME environment variable
      --state string    Location of state storage (kops 'config' file). Overrides KOPS_STATE_STORE environment variable
  -v, --v Level         number for the log level verbosity
```

### SEE ALSO

* [kops completion](kops_completion.md)	 - Generate the autocompletion script for the specified shell
* [kops create](kops_create.md)	 - Create a resource by command line, filename or stdin.
* [kops delete](kops_delete.md)	 - Delete clusters, instancegroups, instances, and secrets.
* [kops distrust](kops_distrust.md)	 - Distrust keypairs.
* [kops edit](kops_edit.md)	 - Edit clusters and other resources.
* [kops export](kops_export.md)	 - Export configuration.
* [kops get](kops_get.md)	 - Get one or many resources.
* [kops promote](kops_promote.md)	 - Promote a resource.
* [kops reconcile](kops_reconcile.md)	 - Reconcile a cluster.
* [kops replace](kops_replace.md)	 - Replace cluster resources.
* [kops rolling-update](kops_rolling-update.md)	 - Rolling update a cluster.
* [kops toolbox](kops_toolbox.md)	 - Miscellaneous, experimental, or infrequently used commands.
* [kops trust](kops_trust.md)	 - Trust keypairs.
* [kops update](kops_update.md)	 - Update a cluster.
* [kops upgrade](kops_upgrade.md)	 - Upgrade a kubernetes cluster.
* [kops validate](kops_validate.md)	 - Validate a kOps cluster.
* [kops version](kops_version.md)	 - Print the kOps version information.

