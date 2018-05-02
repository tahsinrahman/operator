---
title: Uninstall
description: Vault operator Uninstall
menu:
  product_vault-operator_0.1.0:
    identifier: uninstall-stash
    name: Uninstall
    parent: setup
    weight: 20
product_name: vault-operator
menu_name: product_vault-operator_0.1.0
section_menu_id: setup
---
# Uninstall Vault operator

To uninstall Vault operator, run the following command:

```console
$ curl -fsSL https://raw.githubusercontent.com/kubevault/operator/0.1.0/hack/deploy/stash.sh \
    | bash -s -- --uninstall [--namespace=NAMESPACE]

+ kubectl delete deployment -l app=stash -n kube-system
deployment "stash-operator" deleted
+ kubectl delete service -l app=stash -n kube-system
service "stash-operator" deleted
+ kubectl delete secret -l app=stash -n kube-system
No resources found
+ kubectl delete serviceaccount -l app=stash -n kube-system
No resources found
+ kubectl delete clusterrolebindings -l app=stash -n kube-system
No resources found
+ kubectl delete clusterrole -l app=stash -n kube-system
No resources found
+ kubectl delete initializerconfiguration -l app=stash
initializerconfiguration "stash-initializer" deleted
```

The above command will leave the Vault operator crd objects as-is. If you wish to **nuke** all Vault operator crd objects, also pass the `--purge` flag. This will keep a copy of Vault operator crd objects in your current directory.
