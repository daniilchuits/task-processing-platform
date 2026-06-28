package interfaces

import "worker/domain"

type MP3Updater interface {
	Mp3Udate(domain.MP3Data) error
}
