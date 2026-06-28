package interfaces

import "worker/domain"

type JPGUpdater interface {
	JPGUpdate(data domain.DataJPG) error
}
