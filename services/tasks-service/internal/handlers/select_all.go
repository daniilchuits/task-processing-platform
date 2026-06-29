package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"task-service/internal/domain"
	"task-service/internal/interfaces"
	"task-service/internal/transport"
	"task-service/internal/usecases"
)

type selectHandler struct {
	uc usecases.SelectAllTasksUsecase
}

func NewSelectHandler(sel interfaces.SelecterAll) *selectHandler {
	return &selectHandler{
		uc: usecases.SelectAllTasksUsecase{
			Selecter: sel,
		},
	}
}

func (sel *selectHandler) SelectAllTasks(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.Header.Get("user_id")
	if userIdStr == "" {
		http.Error(w, domain.ErrEmptyUserId.Error(), 400)
		return
	}
	log.Println(1)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(domain.ErrConvUserId, ":", userIdStr, " err:", err)
		http.Error(w, domain.ErrConvUserId.Error(), 400)
		return
	}
	log.Println(2)
	tasksDomain, err := sel.uc.Exec(userId)
	if err != nil {

		if errors.Is(err, domain.ErrSelectingAll) {
			http.Error(w, domain.ErrSelectingAll.Error(), 500)
			return
		}

		log.Println("Selecting error:", err)
		http.Error(w, domain.ErrInternalServer.Error(), 500)
		return
	}

	var tasks []transport.Task

	for _, taskDomain := range *tasksDomain {
		httpTask := transport.ToHTTPTask(taskDomain)
		log.Printf("%+v\n", httpTask)

		tasks = append(tasks, httpTask)
	}
	log.Println(3)
	if err = json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println(domain.ErrEncoding, ":", err)
		http.Error(w, domain.ErrEncoding.Error(), 500)
		return
	}
}
