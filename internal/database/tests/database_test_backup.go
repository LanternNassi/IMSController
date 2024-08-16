package tests

import (
	"context"
	"strconv"

	"github.com/LanternNassi/IMSController/internal/models"
)

func (s *DatabaseSuite) TestAddBackUp() {
	backup := &models.Backup{
		Name:   "TestBackup",
		Backup: []byte("This is the backup"),
		Size:   9,
		Bill:   20,
	}

	created_backup, err := s.databaseOperations.AddBackup(context.Background(), backup)

	s.test_backup_id = created_backup.ID
	s.NoError(err)
	s.Equal(created_backup.Name, backup.Name)

}

func (s *DatabaseSuite) TestGetBackUps() {
	backups, err := s.databaseOperations.Getbackups(context.Background(), &models.Backup{})
	s.NoError(err)
	s.Equal(len(backups), 1)

}

func (s *DatabaseSuite) TestGetBackUpById() {
	backup, err := s.databaseOperations.GetBackUpById(context.Background(), strconv.FormatUint(uint64(s.test_backup_id), 10))
	s.NoError(err)
	s.Equal(backup.Name, "TestBackup")
}
