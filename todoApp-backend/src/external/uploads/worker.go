package uploads

type worker struct {
	Id         int
	Free       bool
	CancelChan chan struct{}
	Queue      <-chan UploadAttempt
	Result     chan<- UploadAttempt
}

func (W *worker) work() {
	go func() {
		for {
			select {
			case job := <-W.Queue:

				W.Free = false

				err := uploadImage()
				if err != nil {
					job.Status = false
				}

			case _ = <-W.CancelChan:
				return
			}
		}
	}()
}

func newWorker(id, numJobs int) *worker {

	Worker := &worker{
		Id:         id,
		CancelChan: make(chan struct{}),
		Queue:      make(chan UploadAttempt, numJobs),
		Result:     make(chan UploadAttempt, numJobs),
	}

	Worker.work()

	return Worker
}
