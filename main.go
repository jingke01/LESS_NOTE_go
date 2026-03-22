package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Job struct {
	Id      int
	RandNum int
}
type Result struct {
	job *Job
	sum int
}

func main() {
	var wg = sync.WaitGroup{}
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	creatPool(10, jobChan, resultChan, &wg)
	var PrintWg = sync.WaitGroup{}
	PrintWg.Add(1)
	go func() {
		defer PrintWg.Done()
		for result := range resultChan {
			fmt.Printf("job id %v randnum %d result %d \n", result.job.Id, result.job.RandNum, result.sum)
		}
	}()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		job := &Job{
			Id:      i,
			RandNum: rand.Intn(1000),
		}
		jobChan <- job
	}
	wg.Wait()
	close(jobChan)
	close(resultChan)
	PrintWg.Wait()
	fmt.Println("over")
	//time.Sleep(time.Second)
}
func creatPool(num int, jobChan chan *Job, resultChan chan *Result, wg *sync.WaitGroup) {
	for i := 0; i < num; i++ {
		go func() {
			for job := range jobChan {
				r_num := job.RandNum
				var sum = 0
				for r_num != 0 {
					sum += r_num % 10
					r_num /= 10
				}
				resultChan <- &Result{
					job: job,
					sum: sum,
				}
				wg.Done()
			}
		}()
	}
}
