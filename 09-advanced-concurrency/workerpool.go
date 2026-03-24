package main

import (
	"fmt"
	"sync"
	"time"
)

// Job, işlenecek bir görev birimidir.
type Job struct {
	ID    int
	Value int
}

// Result, işlenen görevin sonucudur.
type Result struct {
	JobID  int
	Output int
}

// worker, kanaldan işleri alır ve sonuçları gönderir.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: İşleniyor %d...\n", id, job.ID)
		time.Sleep(time.Millisecond * 500) // Simüle edilen gecikme
		results <- Result{JobID: job.ID, Output: job.Value * 2}
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// Worker Pool başlatılıyor
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// İşler gönderiliyor
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j * 10}
	}
	close(jobs)

	// Worker'ların bitmesi bekleniyor
	go func() {
		wg.Wait()
		close(results)
	}()

	// Sonuçlar toplanıyor
	for res := range results {
		fmt.Printf("Sonuç: Job %d -> Çıktı %d\n", res.JobID, res.Output)
	}
	fmt.Println("Tüm işler tamamlandı.")
}
