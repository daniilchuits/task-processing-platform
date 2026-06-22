package domain

type Filename struct {
	Name string
}

type Task struct {
	Id                 int
	UserId             string
	Filename           string
	Status             string
	Typ                string
	PhraseCountTxt     *int
	MainColorsJpg      *string
	PhotoResolutionJpg *string
	AudioLengthMp3     *int
	NumOfLinesCsv      *int
	PagesPdf           *int
	PhotoInDocPdf      *bool
	SizeAfterUnzipZip  *int
	ZipFilesZip        *string
}
