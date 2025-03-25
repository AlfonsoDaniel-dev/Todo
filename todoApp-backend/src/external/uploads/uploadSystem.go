package uploads

import (
	"errors"
)

type UploadAttempt struct {
	Id         int
	data       []byte
	FileName   string
	FileExt    string
	Repository string
	Status     error
	DoneChan   chan struct{}
}

type UploadEngine struct {
	workers        map[int]*worker
	MaxThreads     int
	RepositoryPath string
}

func NewUploadEngine(maxThreads int, repositoryPath string) (*UploadEngine, error) {
	if maxThreads < 1 || repositoryPath == "" {
		return nil, errors.New("Not enough arguments to create upload engine instance")
	}

	return &UploadEngine{
		workers:        make(map[int]*worker),
		MaxThreads:     maxThreads,
		RepositoryPath: repositoryPath,
	}, nil
}

func (UE *UploadEngine) startWorker(numOfJobs int, porpuse string) *worker {

	workerId := len(UE.workers) + 1
	Worker := newWorker(workerId, numOfJobs, porpuse)
	UE.workers[workerId] = Worker

	return Worker
}

func (UE *UploadEngine) deleteWorker(id int) {

	Worker := UE.workers[id]

	close(Worker.Queue)

	Worker.CancelChan <- struct{}{}

	close(Worker.CancelChan)

	delete(UE.workers, id)
}

func (UE *UploadEngine) Upload(ImageToUpload []byte, fileName, fileExtension string) error {
	if ImageToUpload == nil || len(ImageToUpload) == 0 || fileName == "" || fileExtension == "" {
		return errors.New("arguments needed to upload image")
	}

	Worker := UE.startWorker(len(UE.workers), "upload")

	attempt := UploadAttempt{
		Id:         1,
		data:       ImageToUpload,
		FileName:   fileName,
		FileExt:    fileExtension,
		Repository: UE.RepositoryPath,
		DoneChan:   make(chan struct{}),
	}

	Worker.Queue <- attempt

	<-attempt.DoneChan

	if attempt.Status != nil {
		UE.deleteWorker(Worker.Id)
		return attempt.Status
	}

	UE.deleteWorker(Worker.Id)

	return errors.New("not implemented yet")

}

func (UE *UploadEngine) Get(fileName, fileRepository string) ([]byte, error) {
	if fileName == "" || fileRepository == "" {
		return nil, errors.New("arguments needed to get an image")
	}

	Worker := UE.startWorker(len(UE.workers), "get")
	attempt := UploadAttempt{
		Id:         1,
		data:       nil,
		FileName:   fileName,
		FileExt:    "",
		Repository: fileRepository,
		DoneChan:   make(chan struct{}),
	}

	Worker.Queue <- attempt

	<-attempt.DoneChan

	if attempt.Status != nil {
		UE.deleteWorker(Worker.Id)
		return nil, attempt.Status
	}

	UE.deleteWorker(Worker.Id)

	return attempt.data, errors.New("not implemented yet")
}
