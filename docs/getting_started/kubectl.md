# kubectl cluster admin configuration

When you run `kops update cluster` during cluster creation, you automatically get a kubectl configuration for accessing the cluster. This configuration gives you full admin access to the cluster.
If you want to create this configuration on other machine, you can run the following as long as you have access to the kOps state store.

To create the kubecfg configuration settings for use with kubectl:

```
export KOPS_STATE_STORE=<location of the kops state store>
NAME=<kubernetes.mydomain.com>
kops export kubecfg ${NAME}
```

Warning: Note that the exported configuration gives you full admin privileges using TLS certificates that are not easy to rotate. For regular kubectl usage, you should consider using another method for authenticating to the cluster.

If you are using kops >= 1.19.0, kops export kubecfg will also require passing either the --admin or --user flag if the context does not already exist. For more information, see the [release notes](https://kops.sigs.k8s.io/releases/1.19-notes/#changes-to-kubernetes-config-export).
