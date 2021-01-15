ackage main

import "fmt"
import "time"

type Job struct {
  duration time.Duration
  name string
  lock bool
}

func (j *Job)Run()  {
  fmt.Printf("Job Starting - %s with duration %v\n", j.name, j.duration)
  time.Sleep(j.duration)
  fmt.Printf("Job Finished - %s with duration %v\n", j.name, j.duration)
}



func NewWorkerPool(max int) *WorkerPool {
  workers := make([]Worker, max)
  for i := 0; i < len(max); i++ {
    workers[i] = Worker{}
  }
  return &WorkerPool{
    workers: workers,
  }
}

type Worker struct {
  Running bool
  jobs *[]Jobs
}

func (w *Worker) Run(job Job, chan bool conn) {
  w.Running = true
  job.Run()
  w.Running = false
  conn<- = bool
}

  

type WorkerPool struct {
  workers []Worker
  jobs []Job
  conn chan bool
}

func(w *WorkerPool) submitAndFeed (job Job) error {
  for _, wo := range w.Workers {
    wo.Running == false {
      go func() {
        for {
          wo.Run(job)
          <-c.conn
        }
      }
      return
    }
  }
  return errors.New("no available workers")
}

func(w *WorkerPool) submit (job Job) {
  for len(availableJobs) > 0{
    
  }
}


func main() {
  wp = NewWorkerPool(3)
  wp.submit()

}