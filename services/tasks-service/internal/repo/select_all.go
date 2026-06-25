package repo

import (
	"log"
	"task-service/internal/domain"
)

func (manager repoManager) SelectAllTasks(userId int) (*[]domain.Task, error) {

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
		WHERE user_id=$1
	`

	var tasks []domain.Task

	rows, err := manager.db.Query(query, userId)
	if err != nil {
		log.Println("Error selecting all tasks:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var task domain.Task

		if err = rows.Scan(
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
		); err != nil {
			log.Println("Error scaning one task:", err)
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error after scaning rows:", err)
		return nil, err
	}
	return &tasks, nil
}
