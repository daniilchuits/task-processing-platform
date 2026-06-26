package domain

type Task struct {
	Id                 int
	UserId             string
	Filename           string
	Filepath           string
	Status             string
	Typ                string
	PhraseCountTxt     *int
	Lines              *int
	MainColorsJpg      *string
	PhotoResolutionJpg *string
	AudioLengthMp3     *int
	NumOfLinesCsv      *int
	PagesPdf           *int
	PhotoInDocPdf      *bool
	SizeAfterUnzipZip  *int
	ZipFilesZip        *string
}
