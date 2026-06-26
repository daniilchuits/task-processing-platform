package usecases

import (
	"context"
	"log"
	"mime/multipart"
	"path/filepath"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
	"time"
)

type PostUsecase struct {
	Check     interfaces.Checker
	Post      interfaces.Poster
	Publish   interfaces.Publisher
	QueueName string
}

func (post *PostUsecase) Exec(userID int, file multipart.File, header *multipart.FileHeader) (*domain.Task, error) {

	exists, err := post.Check.Check(userID, header.Filename)
	if err != nil {
		log.Println("Checking existence err:", err)
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

	if err = post.Publish.PublishMsg(ctx, path, post.QueueName); err != nil {
		return nil, err
	}

	task, err := post.Post.Insert(
		userID,
		header.Filename,
		path,
		fileType,
	)
	if err != nil {
		return nil, domain.ErrInserting
	}
	return task, nil
}
