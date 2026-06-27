package interfaces

import "worker/domain"

type TxtUpdater interface {
	TxtUpdate(data domain.DataTxt) error
}
