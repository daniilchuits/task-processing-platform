package transport

import "task-service/internal/domain"

func ToHTTPTask(task domain.Task) Task {
	taskHTTP := Task{
		Id:       task.Id,
		UserId:   task.UserId,
		Filename: task.Filename,
		Filepath: task.Filepath,
		Status:   task.Status,
		Typ:      task.Typ,
	}

	if task.PhraseCountTxt != nil {
		taskHTTP.PhraseCountTxt = *task.PhraseCountTxt
	}
	if task.Lines != nil {
		taskHTTP.LinesTxt = *task.Lines
	}
	if task.MainColorsJpg != nil {
		taskHTTP.MainColorsJpg = *task.MainColorsJpg
	}
	if task.PhotoResolutionJpg != nil {
		taskHTTP.PhotoResolutionJpg = *task.PhotoResolutionJpg
	}
	if task.AudioLengthMp3 != nil {
		taskHTTP.AudioLengthMp3 = *task.AudioLengthMp3
	}
	if task.NumOfLinesCsv != nil {
		taskHTTP.NumOfLinesCsv = *task.NumOfLinesCsv
	}
	if task.PagesPdf != nil {
		taskHTTP.PagesPdf = *task.PagesPdf
	}
	if task.PhotoInDocPdf != nil {
		taskHTTP.PhotoInDocPdf = *task.PhotoInDocPdf
	}
	if task.SizeAfterUnzipZip != nil {
		taskHTTP.SizeAfterUnzipZip = *task.SizeAfterUnzipZip
	}
	if task.ZipFilesZip != nil {
		taskHTTP.ZipFilesZip = *task.ZipFilesZip
	}
	return taskHTTP
}
