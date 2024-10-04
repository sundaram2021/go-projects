# Project Documentation: Dynamic Worker Pool Implementation in Go

This repository contains a Go application designed to manage and process tasks concurrently using a dynamic worker pool. It showcases features such as rate limiting, task retries with exponential backoff, and the ability to dynamically scale the number of workers based on the workload.

---

## Features

1. **Task Processing**:
   - Handles tasks defined by a unique ID, name, and retry count.
   - Processes tasks concurrently using a pool of worker goroutines.

2. **Dynamic Worker Scaling**:
   - Dynamically adjusts the number of workers based on the task queue length.
   - Scales up to a predefined maximum number of workers.

3. **Rate Limiting**:
   - Enforces a limit on the number of tasks processed per second to prevent overload.

4. **Retry Mechanism**:
   - Automatically retries tasks that fail during processing.
   - Uses exponential backoff to delay retries and reduce load.

5. **Graceful Shutdown**:
   - Supports graceful shutdown of workers, ensuring all ongoing tasks are completed before stopping.

6. **Context Handling**:
   - Utilizes context to handle cancellation and shutdown of the worker pool operations effectively.

---

## How to Use

### Prerequisites

- **Go Environment**: Go installed on your machine (version 1.15+ recommended).
- **Understanding of Concurrency**: Basic knowledge of Go's concurrency model, including goroutines and channels.

### Running the Application

1. **Clone the Repository**:
   ```bash
   git clone repo_url
   cd task_processor
   ```

2. **Build and Run**:
   - Compile the application:
     ```bash
     go build main.go
     ```
   - Run the compiled binary:
     ```bash
     ./main
     ```

### Example Usage

- **Start the Worker Pool**:
  - Initialize the worker pool with a specific configuration (number of workers, max workers, rate limit, etc.).
  - Submit tasks to the worker pool using the `SubmitTask` function.

- **Monitoring and Scaling**:
  - Observe the automatic scaling behavior as the task queue length changes over time.
  - View console outputs to monitor task processing, retries, and scaling activities.

---

## API Reference

- **`NewWorkerPool(numWorkers, maxWorkers, rateLimit, retryLimit int, dynamicScale bool) *WorkerPool`**:
  - Initializes and returns a new instance of `WorkerPool`.

- **`Start()`**:
  - Starts the initial set of worker goroutines and begins processing tasks.

- **`Stop()`**:
  - Gracefully stops all workers and closes the task queue.

- **`SubmitTask(task Task) error`**:
  - Submits a new task to the worker pool for processing.

- **`worker(id int)`**:
  - Represents a worker goroutine that processes tasks from the task queue.

- **`processTask(task Task) bool`**:
  - Processes a given task and returns a success status.

- **`backoffDuration(retries int) time.Duration`**:
  - Calculates the delay for retries based on exponential backoff.

- **`autoScaleWorkers()`**:
  - Monitors the task queue and scales the number of workers dynamically.

---

## Technologies Used

- **Go (Golang)**: Core programming language used to implement the worker pool and its functionalities.

---

## Conclusion

This implementation of a dynamic worker pool in Go provides a robust foundation for building applications that require concurrent task processing with features like rate limiting, retries, and auto-scaling. The application is ideal for developers looking to understand or implement advanced concurrency patterns in Go.