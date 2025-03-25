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
	Status     bool
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

func (UE *UploadEngine) addWorker(numOfJobs int, porpuse string) int {

	workerId := len(UE.workers) + 1
	Worker := newWorker(workerId, numOfJobs, porpuse)
	UE.workers[workerId] = Worker

	return workerId
}

func (UE *UploadEngine) deleteWorker(id int) {

	Worker := UE.workers[id]

	Worker.CancelChan <- struct{}{}

	delete(UE.workers, id)
}

func (UE *UploadEngine) Upload(ImageToUpload []byte, fileName, fileExtension string) error {
	if ImageToUpload == nil || len(ImageToUpload) == 0 || fileName == "" || fileExtension == "" {
		return errors.New("arguments needed to upload image")
	}

	Worker := UE.addWorker(len(UE.workers), "upload")

	return errors.New("not implemented yet")

}

func (UE *UploadEngine) Get(fileName string) ([]bytes, error) {
	if fileName == "" {
		return nil, errors.New("arguments needed to get an image")
	}

	Worker := UE.addWorker(len(UE.workers), "get")

	return nil, errors.New("not implemented yet")
}
