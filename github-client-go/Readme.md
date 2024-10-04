# Project Documentation: Custom Kubernetes Event Recorder in Go

This project is a Go application that demonstrates how to create and record custom events in a Kubernetes cluster using the client-go library. It connects to a Kubernetes cluster, defines a custom event recorder, and logs a custom event for a Kubernetes object.

---

## Features

1. **Kubernetes Client Initialization**:
   - Connects to a Kubernetes cluster using the kubeconfig file.
   - Utilizes `client-go` to interact with the Kubernetes API.

2. **Custom Event Recorder**:
   - Defines a `CustomEventRecorder` struct that wraps the standard event recorder.
   - Implements an `Event` method to create and send custom events to the cluster.
   - Uses the Kubernetes `Event` API to log events associated with Kubernetes objects.

3. **Event Creation for Kubernetes Objects**:
   - Creates a sample `Node` object.
   - Records a custom event (`NodeUpdated`) for the node, indicating that it was successfully updated.

4. **Logging and Error Handling**:
   - Uses `klog` for structured logging and error reporting.
   - Handles errors gracefully during event creation.

---

## How to Run

### Prerequisites

- **Go Environment**: Go installed on your system (version 1.16 or higher recommended).
- **Kubernetes Cluster**: Access to a running Kubernetes cluster.
- **Kubeconfig File**: A valid `kubeconfig` file to authenticate with the Kubernetes cluster.

### Steps

1. **Clone the Repository**:

   ```sh
   git clone repo_url
   cd project_directory
   ```

2. **Install Dependencies**:

   Ensure that you have the necessary Go modules. Initialize and download dependencies:

   ```sh
   go mod init yourmodule
   go mod tidy
   ```

3. **Modify Namespace (Optional)**:

   In the `main` function, you can set the desired namespace where the event will be recorded:

   ```go
   customRecorder := &CustomEventRecorder{
       recorder:      recorder,
       namespace:     "default", // Change to your desired namespace
       kubeClientset: kubeClientset,
   }
   ```

4. **Build the Application**:

   Compile the Go application:

   ```sh
   go build -o custom-event-recorder main.go
   ```

5. **Run the Application**:

   Execute the binary, specifying the path to your kubeconfig file if it's not in the default location:

   ```sh
   ./custom-event-recorder -kubeconfig /path/to/your/kubeconfig
   ```

   If your kubeconfig is in the default location (`$HOME/.kube/config`), you can omit the `-kubeconfig` flag.

6. **Verify the Event in Kubernetes**:

   Check the events in the specified namespace to see if the custom event was recorded:

   ```sh
   kubectl get events -n default
   ```

   Look for an event with the reason `NodeUpdated` and message `Node was successfully updated`.

---

## Technologies Used

- **Go (Golang)**: Programming language used for the application.
- **Kubernetes Client-Go Library**:
  - **client-go**: Official Kubernetes client library for Go.
  - **Modules Used**:
    - `k8s.io/client-go/kubernetes`
    - `k8s.io/client-go/tools/clientcmd`
    - `k8s.io/client-go/tools/record`
    - `k8s.io/client-go/tools/reference`
    - `k8s.io/apimachinery/pkg/apis/meta/v1`
    - `k8s.io/apimachinery/pkg/runtime`
    - `k8s.io/api/core/v1`
- **Klog**:
  - **k8s.io/klog/v2**: Kubernetes logging library used for logging and error messages.
- **XDG Base Directory Support**:
  - **path/filepath**, **flag**, and **homedir**: Used for handling file paths and command-line flags.

---

## Output [here](output.md)