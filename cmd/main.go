package main

import (
	"bytes"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	watch2 "k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
)

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	kClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	watch, err := kClient.CoreV1().Pods("default").Watch(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	results := watch.ResultChan()

	msg := make(map[string]string)
	msg["text"] = "deleted the pod -- kael"
	data, err := json.Marshal(msg)
	buf := bytes.NewBuffer(data)
	if err != nil {
		panic(err)
	}

	for result := range results {
		if result.Type == watch2.Deleted {
			_, err := http.Post("https://hooks.slack.com/services/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", "application/json", buf)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
