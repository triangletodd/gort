# What
![gort](https://github.com/triangletodd/gort/blob/master/static/gort.jpg?raw=true)

Gort is a URL shortener intended to run on a Kubernetes cluster which uses a Kubernetes Custom Resource Definition as its data store.



# Why
I wanted to write some Golang and play with the kubernetes API.

# How
- Created the structure that I wanted for my CRD in [pkg/apis](pkg/apis)
  - It took a few examples to get this working, but here are a few:
    - [OpenShift Blog](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/)
    - [rliang @ ITNEXT](https://itnext.io/how-to-generate-client-codes-for-kubernetes-custom-resource-definitions-crd-b4b9907769ba)
    - [tstringer @ medium](https://medium.com/@trstringer/create-kubernetes-controllers-for-core-and-custom-resources-62fc35ad64a3)
- Ran [update-codegen.sh](hack/update-codegen.sh) to generate the client
  - [update-codegen.sh](hack/update-codegen.sh) uses [kubernetes/code-generator](https://github.com/kubernetes/code-generator) under the covers
  - generated code can be found in the [pkg/client](pkg/client) folder as well as [zz_generated.deepcopy.go](pkg/apis/gorturl/v1/zz_generated.deepcopy.go)
- Manually created my CRD and a few resources to test my client code against
  - [example/client/gorturls.yaml](exmaple/client/gorturls.yaml)
- A quick and dirty client to test that everything was happy
  - [example/client/gort_client.go](example/client/gort_client.go)

# Gotchas
I've temporarily hardcoded the path to my DO kubeconfig on line 13 in [gort_client.go](example/client/gort_client.go#13). If you want this to run locally you'll need to do one of two things:

- Make sure you have a valid kubeconfig at $HOME/.kubeenv/kubedo-1
- Modify [gort_client.go](example/client/gort_client.go#13) to point to the kubeconfig for your cluster

I will, eventually, rework the code to support in-cluster authentication as well as support local dev via the KUBECONFIG env variable (my local preference over contexts).