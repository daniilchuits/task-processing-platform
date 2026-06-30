package interfaces

import "notifycation-service/internal/domain"

type Updater interface {
	UpdateTasks(data domain.Msg) error
}
