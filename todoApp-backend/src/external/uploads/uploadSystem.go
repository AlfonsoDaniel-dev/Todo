package uploads

import (
	"errors"
	"os"
)

type UploadAttempt struct {
	Id       int
	data     []byte
	Status   bool
	DoneChan chan struct{}
}

type UploadEngine struct {
	workers        map[int]*worker
	MaxThreads     int
	RepositoryPath string
}

func NewUploadService(maxThreads int, repositoryPath string) (*UploadEngine, error) {
	if maxThreads < 1 || repositoryPath == "" {
		return nil, errors.New("Not enough arguments to create upload engine instance")
	}

	return &UploadEngine{
		workers:        make(map[int]*worker),
		MaxThreads:     maxThreads,
		RepositoryPath: repositoryPath,
	}, nil
}

func (UE *UploadEngine) addWorker(numOfJobs int) {

	workerId := len(UE.workers) + 1
	Worker := newWorker(workerId, numOfJobs)
	UE.workers[workerId] = Worker
}

func (UE *UploadEngine) deleteWorker(id int) {

	Worker := UE.workers[id]

	Worker.CancelChan <- struct{}{}

	delete(UE.workers, id)
}

func (UE *UploadEngine) Upload(Image []byte, fileName string) error {
	if Image == nil || len(Image) == 0 || fileName == "" {
		return errors.New("Image is nil")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	file.Write(Image)

	return nil
}
