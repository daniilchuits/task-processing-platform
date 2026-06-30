package usecases

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
	"time"
)

type PostUsecase struct {
	Checker   interfaces.Checker
	Poster    interfaces.Poster
	Publisher interfaces.Publisher
	QueueName string
}

func (post *PostUsecase) Exec(userID int, file multipart.File, header *multipart.FileHeader) (*domain.Task, error) {

	exists, err := post.Checker.Check(userID, header.Filename)
	if err != nil {
		return nil, domain.ErrDuringCheckingExistence
	}
	if exists {
		return nil, domain.ErrExists
	}

	fileType, err := domain.DetermineType(header.Filename) // type занято Goлангом :D
	if err != nil {
		return nil, err
	}

	if err = domain.CreateFile(header.Filename, file); err != nil {
		return nil, err
	}

	path := filepath.Join(domain.UploadDir, header.Filename)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task, err := post.Poster.Insert(
		userID,
		header.Filename,
		path,
		fileType,
	)
	if err != nil {
		return nil, domain.ErrInserting
	}

	if err = post.Publisher.PublishMsg(ctx, task.Id, path, fileType, post.QueueName); err != nil {
		return nil, err
	}

	return task, nil
}
