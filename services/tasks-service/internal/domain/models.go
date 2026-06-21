package domain

type Filename struct {
	Name string
}

type Task struct {
	id                   int
	userId               string
	filename             string
	status               string
	typ                  string
	phrase_count_txt     int
	main_colors_jpg      string
	photo_resolution_jpg string
	audio_length_mp3     int
	num_of_lines_csv     int
	pages_pdf            int
	photo_in_doc_pdf     bool
	size_after_unzip_zip int
	zip_files_zip        string
	// TODO
}
