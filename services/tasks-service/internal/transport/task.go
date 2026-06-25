package transport

type Task struct {
	Id                 int    `json:"id"`
	UserId             string `json:"user_id"`
	Filename           string `json:"filename"`
	Filepath           string `json:"filepath"`
	Status             string `json:"status"`
	Typ                string `json:"type"`
	PhraseCountTxt     int    `json:"phrase_count,omitempty"`
	MainColorsJpg      string `json:"main_colors,omitempty"`
	PhotoResolutionJpg string `json:"photo_resolution,omitempty"`
	AudioLengthMp3     int    `json:"audio_length,omitempty"`
	NumOfLinesCsv      int    `json:"number_of_lines,omitempty"`
	PagesPdf           int    `json:"pages_count,omitempty"`
	PhotoInDocPdf      bool   `json:"photo_in_document,omitempty"`
	SizeAfterUnzipZip  int    `json:"size_after_unzip,omitempty"`
	ZipFilesZip        string `json:"files_in_zip,omitempty"`
}
