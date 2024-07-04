package getsa

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"os"
)

var (
	// outside cluster client

	config, _    = clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	clientset, _ = kubernetes.NewForConfig(config)

	/*
		// inside cluster client
		// creates the in-cluster config
		config, _ = rest.InClusterConfig()

		// creates the clientset
		clientset, _ = kubernetes.NewForConfig(config)

	*/
)

func GetSa(w http.ResponseWriter, r *http.Request) {
	// create role binding for service account
	//_, err := clientset.RbacV1().RoleBindings(nsName).Create(context.Background(), &roleBinding, metav1.CreateOptions{})
	list, _ := clientset.CoreV1().ServiceAccounts("").List(context.Background(), v1.ListOptions{})

	// declare empty  slice
	sl := []string{}

	// iterate over list to get sa names and put it to slice
	for _, item := range list.Items {
		//log.Println(item.Name)
		sl = append(sl, item.Name)
	}
	// output with ResponseWriter
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	fmt.Fprintf(w, "That is slice for service account")
	data, err := json.Marshal(sl)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}
