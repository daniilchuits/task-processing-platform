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

type insertHandler struct {
	uc usecases.InsertUsecase
}

func NewInsertHandler(
	exists interfaces.CheckingExistence,
	insert interfaces.NoteInserter,
) *insertHandler {
	return &insertHandler{
		uc: usecases.InsertUsecase{
			Exists: exists,
			Insert: insert,
		},
	}
}

func (ins *insertHandler) InsertTask(w http.ResponseWriter, r *http.Request) {

	var filenameHTTP transport.Filename
	if err := json.NewDecoder(r.Body).Decode(&filenameHTTP); err != nil {
		log.Println(domain.ErrDecodingFilename, ":", err)
		http.Error(w, domain.ErrDecodingFilename.Error(), 400)
		return
	}

	filename := transport.FilenameToDomain(filenameHTTP)

	userIdStr := r.Header.Get("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Println(domain.ErrConvUserId, ":", userIdStr, " err:", err)
		http.Error(w, domain.ErrConvUserId.Error(), 400)
		return
	}

	task, err := ins.uc.Exec(userId, filename)
	if err != nil {

		switch {
		case errors.Is(err, domain.ErrChecking):
			http.Error(w, domain.ErrChecking.Error(), 500)
			return
		case errors.Is(err, domain.ErrExists):
			http.Error(w, domain.ErrExists.Error(), 400)
			return
		case errors.Is(err, domain.ErrInvalidExtension):
			http.Error(w, domain.ErrInvalidExtension.Error(), 400)
			return
		case errors.Is(err, domain.ErrInserting):
			http.Error(w, domain.ErrInserting.Error(), 500)
			return
		case errors.Is(err, domain.ErrRegexp):
			http.Error(w, domain.ErrRegexp.Error(), 400)
			return
		default:
			log.Println("Inserting error:", err)
			http.Error(w, domain.ErrInternalServer.Error(), 500)
			return
		}
	}

	taskHTTP := transport.ToHTTPTask(*task)
	if err = json.NewEncoder(w).Encode(taskHTTP); err != nil {
		log.Println(domain.ErrEncoding, ":", err)
		http.Error(w, domain.ErrEncoding.Error(), 500)
		return
	}
}
