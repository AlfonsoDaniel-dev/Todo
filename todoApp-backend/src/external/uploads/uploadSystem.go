package uploads

import (
	"bytes"
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type UploadAttempt struct {
	Id               int
	data             []byte
	FileName         string
	FileExt          string
	UploadRepository string
	Status           bool
	DoneChan         chan struct{}
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

func (UE *UploadEngine) Upload(ImageToUpload []byte, fileName, fileExtension string) error {
	if ImageToUpload == nil || len(ImageToUpload) == 0 || fileName == "" || fileExtension == "" {
		return errors.New("arguments needed to upload image")
	}

}
