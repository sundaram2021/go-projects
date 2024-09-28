package processor

import (
    "context"
    "fmt"
    "math"
    "sync"
    "time"
)

// Task represents a task to be processed.
type Task struct {
    ID      int
    Name    string
    Retries int
}

// WorkerPool manages a pool of workers to process tasks concurrently.
type WorkerPool struct {
    taskQueue    chan Task    // Channel for task queue
    numWorkers   int          // Initial number of workers
    maxWorkers   int          // Maximum number of workers
    dynamicScale bool         // Whether to dynamically scale workers
    wg           sync.WaitGroup
    shutdownChan chan struct{}  // Channel for graceful shutdown
    rateLimit    int           // Number of tasks per second
    retryLimit   int           // Max number of retries
    ctx          context.Context
    cancel       context.CancelFunc
}

// NewWorkerPool creates a new worker pool with dynamic scaling, rate limiting, and retry options.
func NewWorkerPool(numWorkers, maxWorkers, rateLimit, retryLimit int, dynamicScale bool) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    return &WorkerPool{
        taskQueue:    make(chan Task),
        numWorkers:   numWorkers,
        maxWorkers:   maxWorkers,
        dynamicScale: dynamicScale,
        shutdownChan: make(chan struct{}),
        rateLimit:    rateLimit,
        retryLimit:   retryLimit,
        ctx:          ctx,
        cancel:       cancel,
    }
}

// Start initializes the worker pool and starts the workers.
func (wp *WorkerPool) Start() {
    for i := 1; i <= wp.numWorkers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }

    if wp.dynamicScale {
        go wp.autoScaleWorkers()
    }
}

// Stop gracefully shuts down the worker pool.
func (wp *WorkerPool) Stop() {
    close(wp.shutdownChan) // Notify workers to stop
    close(wp.taskQueue)    // Stop accepting new tasks
    wp.wg.Wait()           // Wait for all workers to complete
}

// SubmitTask adds a new task to the task queue for processing.
func (wp *WorkerPool) SubmitTask(task Task) error {
    select {
    case wp.taskQueue <- task:
        return nil
    case <-wp.ctx.Done():
        return fmt.Errorf("worker pool shutting down, unable to submit task")
    }
}

// worker processes tasks from the taskQueue with retries and rate limiting.
func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    
    rateLimitTicker := time.NewTicker(time.Second / time.Duration(wp.rateLimit)) // Rate limiting
    defer rateLimitTicker.Stop()

    for {
        select {
        case task, ok := <-wp.taskQueue:
            if !ok {
                return // Exit if the task queue is closed
            }
            <-rateLimitTicker.C // Apply rate limiting by waiting for the ticker

            fmt.Printf("Worker %d processing task: %d - %s\n", id, task.ID, task.Name)

            if success := wp.processTask(task); !success {
                if task.Retries < wp.retryLimit {
                    task.Retries++
                    fmt.Printf("Worker %d retrying task: %d - %s (Retry #%d)\n", id, task.ID, task.Name, task.Retries)
                    time.Sleep(wp.backoffDuration(task.Retries)) // Exponential backoff
                    wp.SubmitTask(task) // Retry the task
                } else {
                    fmt.Printf("Worker %d failed task after retries: %d - %s\n", id, task.ID, task.Name)
                }
            }
        case <-wp.shutdownChan:
            fmt.Printf("Worker %d shutting down gracefully.\n", id)
            return
        }
    }
}

// processTask simulates task processing and randomly fails some tasks for retry.
func (wp *WorkerPool) processTask(task Task) bool {
    // Simulate random task failure based on ID (fail if the ID is odd for example)
    if task.ID%2 == 1 {
        fmt.Printf("Task %d failed!\n", task.ID)
        return false
    }

    fmt.Printf("Task %d processed successfully.\n", task.ID)
    return true
}

// backoffDuration calculates an exponential backoff duration for retries.
func (wp *WorkerPool) backoffDuration(retries int) time.Duration {
    // Calculate exponential backoff, e.g., 2^retries seconds
    return time.Duration(math.Pow(2, float64(retries))) * time.Second
}

// autoScaleWorkers scales the worker pool based on the length of the task queue.
func (wp *WorkerPool) autoScaleWorkers() {
    ticker := time.NewTicker(5 * time.Second) // Check every 5 seconds
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            queueLength := len(wp.taskQueue)
            fmt.Printf("Task queue length: %d\n", queueLength)
            if queueLength > wp.numWorkers && wp.numWorkers < wp.maxWorkers {
                additionalWorkers := wp.maxWorkers - wp.numWorkers
                for i := 0; i < additionalWorkers && wp.numWorkers < wp.maxWorkers; i++ {
                    wp.numWorkers++
                    wp.wg.Add(1)
                    go wp.worker(wp.numWorkers)
                    fmt.Printf("Added worker %d, total workers: %d\n", wp.numWorkers, wp.numWorkers)
                }
            }
        case <-wp.shutdownChan:
            fmt.Println("Stopping auto-scaler.")
            return
        }
    }
}
