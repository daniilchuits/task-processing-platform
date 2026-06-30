package interfaces

import "worker/domain"

type ZipUpdater interface {
	ZipUpdate(domain.ZipData) error
}
