package interfaces

import "worker/domain"

type PDFUpdater interface {
	PdfUpdate(data domain.PDFData) error
}
