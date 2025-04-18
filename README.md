# Development workflows


## Rebuild the binary
```sh
omega install
```

You can use 
```sh
go install ~/sources/omega/cmd/omega
```
as a fallback if its really bricked, or this is your first install

## Logging into the Omega Cluster
Just need access to the `kubeconfig.json` (ask Ford)

## Port forward argocd server
```sh
devbox run argocd-forward
```

```sh
devbox run argocd-password
```

## Adding new cdk8s code (manifests, helm charts, projects, ect)
Sync the rendered directory. Make sure you rebuild the binary
```sh
omega sync -d
```


## Importing CRDs into cdk8s
You can alter the wellKnownCrds list to include more. This avoids accidentally importing a ton of CRDs
```sh
omega import
```


## Creating a new project
First create a new go file in /projects, the add it to the array in registry.go



