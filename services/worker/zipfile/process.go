package zipfile

import "archive/zip"

func process(reader *zip.Reader) ([]string, int64) {

	var (
		sizeAfterUnzip int64
		files          []string
	)

	for _, file := range reader.File {

		sizeAfterUnzip += file.FileInfo().Size()
		files = append(files, file.Name)
	}
	return files, sizeAfterUnzip
}
