package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "task_processor/processor"
    "time"
)

func main() {
    // Initialize a worker pool with rate limiting (2 tasks per second) and retry limit of 3.
    workerPool := processor.NewWorkerPool(3, 10, 2, 3, true)
    workerPool.Start()

    // Simulate submitting tasks.
    for i := 1; i <= 20; i++ {
        task := processor.Task{
            ID:   i,
            Name: fmt.Sprintf("Task-%d", i),
        }
        fmt.Printf("Submitting task: %d - %s\n", task.ID, task.Name)
        err := workerPool.SubmitTask(task)
        if err != nil {
            fmt.Println(err)
            break
        }
        time.Sleep(500 * time.Millisecond) // Simulate delay between task submissions
    }

    // Set up graceful shutdown using OS signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
    fmt.Println("\nReceived shutdown signal. Shutting down gracefully...")
    workerPool.Stop()
    fmt.Println("All tasks processed and worker pool shut down!")
}
