package repo

import "testing"

func TestSelectOneTask(t *testing.T) {
	count := 12
	data := []struct {
		name           string
		mock           repoMock
		taskID         int
		userID         int
		expErr         error
		expTaskID      int
		expUserID      int
		expPhraseCount int
	}{
		{
			"no error",
			repoMock{
				t: t,
				values: []any{
					1,                    // Id
					1,                    // UserId
					"input.txt",          // Filename
					"/uploads/input.txt", // Filepath
					"pending",            // Status
					"txt",                // Typ
					&count,               // PhraseCountTxt (*int)
					nil,                  // Lines
					nil,                  // MainColorsJpg
					nil,                  // PhotoResolutionJpg
					nil,                  // AudioLengthMp3
					nil,                  // NumOfLinesCsv
					nil,                  // SizeAfterUnzipZip
					nil,                  // ZipFilesZip
				},
			},
			1, 1, nil, 1, 1, count,
		},
	}
	// / add []any to mock so I can input task s values

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewRepoManager(tt.mock)
			task, err := manager.SelectOneTask(tt.taskID, tt.userID)
			if err != tt.expErr {
				t.Fatalf("Expected %v, got %v", tt.expErr, err)
			}
			if task.Id != tt.expTaskID {
				t.Errorf("Expected taskID==%d, got %d", tt.expTaskID, task.Id)
			}
			if task.UserId != tt.expUserID {
				t.Errorf("Expected userID==%d, got %d", tt.expUserID, task.UserId)
			}
			if *task.PhraseCountTxt != tt.expPhraseCount {
				t.Errorf("Expected PhraseCount==%d, got %v", tt.expPhraseCount, *task.PhraseCountTxt)
			}
		})
	}
}
