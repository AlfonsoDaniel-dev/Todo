package uploads

type worker struct {
	Id         int
	Free       bool
	Propose    string
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

				switch W.Propose {
				case "upload":
					err := uploadImage(job.data, job.FileName, job.Repository)
					if err != nil {
						job.Status = false
					}
					job.Status = true

					W.Free = true

					job.DoneChan <- struct{}{}
				case "get":
					img, err := getImage(job.FileName, job.Repository)
					if err != nil {
						job.Status = false
					}
					job.Status = true

					job.data = img

					W.Free = true

					job.DoneChan <- struct{}{}
				}

			case _ = <-W.CancelChan:
				return
			}
		}
	}()
}

func newWorker(id, numJobs int, porpuse string) *worker {

	Worker := &worker{
		Id:         id,
		Propose:    porpuse,
		CancelChan: make(chan struct{}),
		Queue:      make(chan UploadAttempt, numJobs),
		Result:     make(chan UploadAttempt, numJobs),
	}

	Worker.work()

	return Worker
}
