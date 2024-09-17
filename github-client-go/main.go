package main

import (
    "context"
    "flag"
    "fmt"
    "time"

    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/kubernetes/scheme"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/record"
    "k8s.io/client-go/util/homedir"
	"k8s.io/client-go/tools/reference"
    "k8s.io/klog/v2"
    "path/filepath"
)

type CustomEventRecorder struct {
    recorder     record.EventRecorder
    namespace    string
    kubeClientset kubernetes.Interface
}

func (c *CustomEventRecorder) Event(object runtime.Object, eventtype, reason, message string) {
    ref, err := reference.GetReference(scheme.Scheme, object)
    if err != nil {
        return
    }

    event := &corev1.Event{
        ObjectMeta: metav1.ObjectMeta{
            Name:      fmt.Sprintf("%s.%x", ref.Name, time.Now().UnixNano()),
            Namespace: c.namespace,
        },
        InvolvedObject: *ref,
        Reason:         reason,
        Message:        message,
        Source: corev1.EventSource{
            Component: "controllerAgentName",
        },
        FirstTimestamp: metav1.NewTime(time.Now()),
        LastTimestamp:  metav1.NewTime(time.Now()),
        Count:          1,
        Type:           eventtype,
    }

    _, err = c.kubeClientset.CoreV1().Events(c.namespace).Create(context.TODO(), event, metav1.CreateOptions{})
    if err != nil {
        klog.ErrorS(err, "Failed to create event")
    } else {
        klog.V(4).InfoS("Event created successfully", "event", event.Name, "namespace", c.namespace)
    }
}

func main() {
    klog.InitFlags(nil)
    defer klog.Flush()

    var kubeconfig string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = filepath.Join(home, ".kube", "config")
    }
    flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "(optional) absolute path to the kubeconfig file")
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        klog.Fatalf("Error building kubeconfig: %v", err)
    }

    kubeClientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        klog.Fatalf("Error creating Kubernetes client: %v", err)
    }

    eventBroadcaster := record.NewBroadcaster()
    recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "custom-event-recorder"})

    customRecorder := &CustomEventRecorder{
        recorder:     recorder,
        namespace:    "default", // Set to your desired namespace
        kubeClientset: kubeClientset,
    }

    
    node := &corev1.Node{
        ObjectMeta: metav1.ObjectMeta{
            Name: "test-node",
        },
    }

    customRecorder.Event(node, corev1.EventTypeNormal, "NodeUpdated", "Node was successfully updated")
}
