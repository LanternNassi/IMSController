package tests

import (
	"context"
	"strconv"

	"github.com/LanternNassi/IMSController/internal/models"
)

func (s *DatabaseSuite) Test_013_AddBackUp() {
	backup := &models.Backup{
		ClientID: *s.test_client_id,
		Name:     "TestBackup",
		Backup:   "https://testbackup.com",
		Size:     9,
		BillID:   s.test_bill_id,
	}

	created_backup, err := s.databaseOperations.AddBackup(context.Background(), backup)

	s.test_backup_id = created_backup.ID
	s.NoError(err)
	s.Equal(created_backup.Name, backup.Name)

}

func (s *DatabaseSuite) Test_014_GetBackUps() {
	backups, err := s.databaseOperations.Getbackups(context.Background(), &models.Backup{})
	s.NoError(err)
	s.Equal(1, len(backups))

}

func (s *DatabaseSuite) Test_015_GetBackUpById() {
	backup, err := s.databaseOperations.GetBackUpById(context.Background(), strconv.FormatUint(uint64(s.test_backup_id), 10))
	s.NoError(err)
	s.Equal("TestBackup", backup.Name)
}
