package interfaces

import "worker/domain"

type CSVUpdater interface {
	CSVUpdate(domain.CsvData) error
}
