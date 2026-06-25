package repo

import "task-service/internal/domain"

func (manager repoManager) SelectOneTask(taskId, userId int) (*domain.Task, error) {

	query := `
		SELECT
			id,
			user_id,
			filename,
			filepath,
			status,
			type,
			phrase_count_txt,
			main_colors_jpg,
			photo_resolution_jpg,
			audio_length_mp3,
			num_of_lines_csv,
			pages_pdf,
			photo_in_doc_pdf,
			size_after_unzip_zip,
			zip_files_zip
		FROM tasks
		WHERE id=$1
			AND user_id=$2
	`

	var task domain.Task
	err := manager.db.QueryRow(query, taskId, userId).Scan(
		&task.Id,
		&task.UserId,
		&task.Filename,
		&task.Filepath,
		&task.Status,
		&task.Typ,
		&task.PhraseCountTxt,
		&task.MainColorsJpg,
		&task.PhotoResolutionJpg,
		&task.AudioLengthMp3,
		&task.NumOfLinesCsv,
		&task.PagesPdf,
		&task.PhotoInDocPdf,
		&task.SizeAfterUnzipZip,
		&task.ZipFilesZip,
	)
	return &task, err
}
