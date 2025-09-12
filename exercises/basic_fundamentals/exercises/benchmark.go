package exercises

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var numWorkers = runtime.NumCPU()
var textForSorting = " Sorting is the act of imposing order on chaos, and here is a collection of chaos for you:" +
	"banana, Zebra, apple, Mango, apricot, 42, 101, 7, 9999, 2025-09-12, 1999-12-31, asterisk, #hashtag, elephant, Dog, cat, aardvark, dolphin, dragonfly, Mango (again), 🦊, ⚡,"

var jobs = []struct {
	name string
	fn   func()
}{
	{"Mmc", func() { _ = Mmc(254, 1220) }},
	{"Quicksort", runQuicksort},
	{"Factorial", func() { _ = Factorial(10000) }},
	{"CharactersSort", func() { CharactersSort(textForSorting) }},
	{"AccumulativeSum", accumulativeSum},
}

func fuzzyWorker(id int, wg *sync.WaitGroup, fn func(), jobName string) {
	defer wg.Done()
	fmt.Printf("Worker %d starting job: %s\n", id, jobName)
	fn()
	fmt.Printf("Worker %d finished job: %s\n", id, jobName)
}

func runQuicksort() {
	quicksort := Quicksort
	arr := make([]int, 100000)
	for i := range arr {
		arr[i] = i
	}
	quicksort(arr, 0, len(arr)-1)
}

func accumulativeSum() {
	sum := 0
	for i := range 100000000 {
		sum += i
	}
}

func ParallelExecution() {
	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(numWorkers * len(jobs))
	for i := range numWorkers {
		for _, job := range jobs {
			go fuzzyWorker(i, &wg, job.fn, job.name)
		}
	}
	wg.Wait()
	fmt.Println("All workers done!")

	fmt.Println("Parallel Benchmarking completed in", time.Since(start).Milliseconds(), "ms.")
}

func SequentialExecution() {
	start := time.Now()
	fmt.Println("Starting Sequential Benchmarking at", time.Since(start).Milliseconds(), "ms")
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}

	for i := range numWorkers {
		fmt.Println("Starting batch", i+1)
		for _, job := range jobs {
			job.fn()
		}
		fmt.Println("Finished batch", i+1)
	}

	fmt.Println("Sequential Benchmarking completed at", time.Since(start).Milliseconds(), "ms")
}

func RunBenchmark() {
	fmt.Println("Number of CPU cores:", numWorkers)
	SequentialExecution()
	ParallelExecution()
}
