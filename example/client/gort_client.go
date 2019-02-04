package main

import (
	"fmt"
	"log"
	"path/filepath"

	mtnclient "github.com/triangletodd/gort/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kubeenv", "kubedo-1")
	} else {
		panic("can't find your home directory")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := mtnclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	list, err := clientset.MtnV1().GortURLs("default").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing all gorturls: %v", err)
	}

	for _, gorturl := range list.Items {
		fmt.Printf("gorturl %s with url %s \n", gorturl.Name, gorturl.Spec.Long)
	}
}

// $ go run gort_client.go
// gorturl 101 with url aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbQ==
// gorturl 102 with url aHR0cHM6Ly93d3cueWFob28uY29t
// gorturl 103 with url aHR0cHM6Ly9sb2dzLnZpdGFsYm9vay5jb20v
// gorturl 104 with url aHR0cHM6Ly9ncmFmYW5hLmdjLnZpdGFsYm9vay5jb20v
