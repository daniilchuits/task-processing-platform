package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
	"task-service/internal/messaging/rabbitmq"
	"task-service/internal/transport"
	"task-service/internal/usecases"
)

type postHandler struct {
	uc usecases.PostUsecase
}

func NewPostHandler(
	check interfaces.Checker,
	post interfaces.Poster,
	pub rabbitmq.MyPublisher,
) *postHandler {
	return &postHandler{
		uc: usecases.PostUsecase{
			Check:   check,
			Post:    post,
			Publish: &pub,
		},
	}
}

func (post *postHandler) PostTask(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.Header.Get("user_id")
	if userIdStr == "" {
		http.Error(w, domain.ErrEmptyUserId.Error(), 400)
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, domain.ErrConvUserId.Error(), 400)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, domain.ErrGetting.Error(), 400)
		return
	}

	domainTask, err := post.uc.Exec(userId, file, header)
	if err != nil {

		switch {
		case errors.Is(err, domain.ErrDuringCheckingExistence):
			http.Error(w, domain.ErrDuringCheckingExistence.Error(), 500)
			return
		case errors.Is(err, domain.ErrExists):
			http.Error(w, domain.ErrExists.Error(), 400)
			return
		case errors.Is(err, domain.ErrInvalidExtension):
			http.Error(w, domain.ErrInvalidExtension.Error(), 400)
			return
		case errors.Is(err, domain.ErrCopy):
			http.Error(w, domain.ErrCopy.Error(), 500)
			return
		case errors.Is(err, domain.ErrCreating):
			http.Error(w, domain.ErrCreating.Error(), 500)
			return
		case errors.Is(err, domain.ErrFileExists):
			http.Error(w, domain.ErrFileExists.Error(), 400)
			return
		case errors.Is(err, domain.ErrInserting):
			http.Error(w, domain.ErrInserting.Error(), 500)
			return
		case errors.Is(err, domain.ErrPublishingMessageToRabbitMQ):
			http.Error(w, domain.ErrPublishingMessageToRabbitMQ.Error(), 500)
			return
		default:
			log.Println("Internal error:", err)
			http.Error(w, domain.ErrInternalServer.Error(), 500)
			return
		}
	}

	task := transport.ToHTTPTask(*domainTask)
	if err = json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, domain.ErrEncoding.Error(), 500)
		return
	}

}
