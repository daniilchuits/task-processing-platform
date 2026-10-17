package repo

import "testing"

func TestSelectOneTask(t *testing.T) {
	data := []struct {
		name      string
		mock      repoMock
		taskID    int
		userID    int
		expErr    error
		expTaskID int
		expUserID int
	}{
		{
			"no error", repoMock{values: []any{1, 1}}, 1, 1, nil, 1, 1,
		},
	}
	// / add []any to mock so I can input task s valuesы

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
		})
	}
}
