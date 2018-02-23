package lb

import (
	"github.com/ericchiang/k8s"
	"time"
	"context"
)

type K8SPeerWatcher struct {
	podLabel string

	podsAdded   chan podMetaData
	podsDeleted chan podMetaData

	isWatching bool

	namespace string
	labels    map[string]string
}

func NewK8SWatcher(namespace string, labels map[string]string) (*K8SPeerWatcher, error) {

	return &K8SPeerWatcher{
		podsAdded:   make(chan podMetaData, 10),
		podsDeleted: make(chan podMetaData, 10),
		namespace:   namespace,
		labels:      labels,
	}, nil
}

func (k *K8SPeerWatcher) PodsAdded() chan podMetaData {
	return k.podsAdded
}

func (k *K8SPeerWatcher) PodsDeleted() chan podMetaData {
	return k.podsDeleted
}

func (k *K8SPeerWatcher) WatchPods() {
	if k.isWatching {
		panic("WatchPods has already been called.")
	}
	k.isWatching = true

	sleepTime := time.Second
	for {
		time.Sleep(sleepTime)
		w, err := k.getWatcher()
		if err != nil {
			if sleepTime < time.Minute {
				sleepTime = sleepTime * 2
			}
			continue
		} else {
			sleepTime = time.Second
		}
		err = k.watchK8S(w)
		if err != nil {
			if sleepTime < time.Minute {
				sleepTime = sleepTime * 2
			}
		} else {
			sleepTime = time.Second
		}
	}
}

func (k *K8SPeerWatcher) getWatcher() (*k8s.CoreV1PodWatcher, error) {
	client, err := k8s.NewInClusterClient()
	if err != nil {
		return nil, err
	}
	l := new(k8s.LabelSelector)
	for k, v := range k.labels {
		l.Eq(k, v)
	}

	return client.CoreV1().WatchPods(context.Background(), k.namespace, l.Selector())
}

func (k *K8SPeerWatcher) watchK8S(podWatcher *k8s.CoreV1PodWatcher) error {
	defer podWatcher.Close()
	for {
		event, pod, err := podWatcher.Next()
		if err != nil {
			return err
		}
		if *event.Type == "DELETED" || *event.Type == "ERROR" {
			k.podsDeleted <- podMetaData{UID: *pod.Metadata.Uid}
		} else if *event.Type == "ADDED" || *event.Type == "MODIFIED" {
			if *pod.Status.Phase != "Running" || *pod.Status.PodIP == "" {
				k.podsDeleted <- podMetaData{UID: *pod.Metadata.Uid}
			} else {
				k.podsAdded <- podMetaData{podIP: *pod.Status.PodIP, UID: *pod.Metadata.Uid}
			}
		}
	}
}

type podMetaData struct {
	podIP string
	UID   string
}
