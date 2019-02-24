package function

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Handle a serverless request
func Handle(req []byte) string {
	config, err := rest.InClusterConfig()

	if err != nil {
		fmt.Println("couldn't get cluster config - err:", err)
		os.Exit(-1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	data, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	namespaces := "Namespaces"

	for _, namespace := range data.Items {
		namespaces = fmt.Sprintf("%s\n%s", namespaces, namespace.GetName())
	}

	return fmt.Sprintf(namespaces)
}
