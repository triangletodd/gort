# What
![gort](https://github.com/triangletodd/gort/blob/master/static/gort.jpg?raw=true)

Gort is a URL shortener intended to run on a Kubernetes cluster which uses the native Kubernetes API as its datastore via a Custom Resource Definition (gorturl).

# Why
I wanted to write some Golang and play with the Kubernetes API and Kubernetes CRD's.

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
  - [example/client/gorturls.yaml](example/client/gorturls.yaml)
- A quick and dirty client to test that everything was happy
  - [example/client/gort_client.go](example/client/gort_client.go)
- Used [gin](https://github.com/gin-gonic/gin) to create an http server with the following verbs#endpoints
  - GET#service/status -- Intended to be used for healthchecking.
  - GET#url -- Lists of all GortURLs (JSON)
  - POST#url -- Creates a GortURL (JSON)
  - GET#url/:short -- Returns a GortURL (JSON)
  - GET#/:short -- Redirects to a GortURL's long address
    - Hacky, but I used gin router's NoRoute functionality to do this
    - httprouter v1 only supports explicit matches. see: [#73](https://github.com/julienschmidt/httprouter/issues/73)
- Wrote a "db" adapter using my new generated client [internal/k8s/k8s.go](internal/k8s/k8s.go)
- Wrapped my v1.GortURL custom resource in a URL model [internal/models/url.go](internal/models/url.go)

# Gotchas
I've temporarily hardcoded the path to my DigitalOcean cluster's kubeconfig on line 13 in [gort_client.go](example/client/gort_client.go#13). For now, if you want this to run locally you'll need to do one of two things:

- Make sure you have a valid kubeconfig at $HOME/.kubeenv/kubedo-1
- Modify [gort_client.go](example/client/gort_client.go#13) to point to the kubeconfig for your cluster

I will, eventually, rework the code to support in-cluster authentication using a k8s service account as well as support local dev via the KUBECONFIG env variable (my local preference over contexts).

# TODO
  - Support in-cluster auth as-well as KUBECONFIG auth for local development
  - Add delete and update functionality to the URL model
  - Add delete and update functionality to the API
  - Create a Dockerfile, automate the build on Dockerhub
  - Create a Kubernetes Deployment file
  - GortURL should bootstrap itself (create its own CRD)
