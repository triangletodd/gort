package k8s

import (
	"path/filepath"

	mtnclient "github.com/triangletodd/gort/pkg/client/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetClient() *mtnclient.Clientset {
	var kubeconfig string

	// FIXME: not the right way to do this.. respect KUBECONFIG
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

	return clientset
}
