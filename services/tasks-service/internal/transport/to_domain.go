package transport

import "task-service/internal/domain"

func FilenameToDomain(name Filename) domain.Filename {
	return domain.Filename{Name: name.Name}
}
