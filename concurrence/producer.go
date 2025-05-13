package main

import (
    "fmt"
    "sync"
)

/*
10个生产者，5个消费者
共1000个物料，
5个消费者并行消费


*/
func main() {
    const (
        numProducers = 10
        numConsumers = 5
        totalJobs    = 1000
    )
    jobs := make(chan int, totalJobs)
    materials := make(chan int, numProducers) // ?

    go func() {
        for i := 1; i <= totalJobs; i++ {
            jobs <- i
        }
        close(jobs)
    }()

    var wgProducers sync.WaitGroup
    wgProducers.Add(numProducers)
    for i := 0; i < numProducers; i++ {
        go func() {
            defer wgProducers.Done()
            for job := range jobs {
                materials <- job
            }
        }()
    }

    // wait produce done
    go func() {
        wgProducers.Wait()
        close(materials)
    }()

    var wgConsumers sync.WaitGroup
    wgConsumers.Add(numConsumers)

    for i := 0; i < numConsumers; i++ {
        go func() {
            defer wgConsumers.Done()
            for material := range materials {
                fmt.Printf("consumer eat:%d \n", material)
            }
        }()
    }
    wgConsumers.Wait()
    fmt.Println("work done")

    sync.WaitGroup{}
    sync
}
