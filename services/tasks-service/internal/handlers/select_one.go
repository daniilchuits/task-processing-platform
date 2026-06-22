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

	"github.com/go-chi/chi/v5"
)

type selOneTask struct {
	uc usecases.SelectOneTask
}

func NewSelectOneTaskHandler(sel interfaces.Selecter) *selOneTask {
	return &selOneTask{
		uc: usecases.SelectOneTask{
			Selecter: sel,
		},
	}
}

func (sel *selOneTask) SelectTaskById(w http.ResponseWriter, r *http.Request) {

	userIdStr := r.Header.Get("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(domain.ErrConvUserId, ":", userIdStr, " err:", err)
		http.Error(w, domain.ErrConvUserId.Error(), 400)
		return
	}

	taskIdStr := chi.URLParam(r, "id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		log.Println(domain.ErrConvUserId, ":", err)
		http.Error(w, domain.ErrConvUserId.Error(), 400)
		return
	}

	taskDomain, err := sel.uc.Exec(taskId, userId)
	if err != nil {

		if errors.Is(err, domain.ErrSelectingOne) {
			http.Error(w, domain.ErrSelectingOne.Error(), 500)
			return
		} else if errors.Is(err, domain.ErrStrconvId) {
			http.Error(w, domain.ErrStrconvId.Error(), 400)
			return
		}

		log.Println(err)
		http.Error(w, domain.ErrInternalServer.Error(), 500)
		return
	}

	taskHTTP := transport.ToHTTPTask(*taskDomain)
	if err = json.NewEncoder(w).Encode(taskHTTP); err != nil {
		http.Error(w, domain.ErrEncoding.Error(), 500)
		return
	}
}
