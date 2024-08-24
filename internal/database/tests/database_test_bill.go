package tests

import (
	"context"
	"strconv"

	"github.com/LanternNassi/IMSController/internal/models"
	"github.com/shopspring/decimal"
)

func (s *DatabaseSuite) Test_012_AddBill() {
	_bill := &models.Bill{
		ClientID:    *s.test_client_id,
		BackupCount: 1,
		BackupSize:  20,
		TotalCost:   decimal.NewFromInt(20000),
		Billed:      true,
	}

	created_bill, err := s.databaseOperations.AddBill(context.Background(), _bill)
	s.test_bill_id = created_bill.ID
	s.NoError(err)
	s.Equal(_bill, created_bill)
}

func (s *DatabaseSuite) Test_013_GetBills() {
	bills, err := s.databaseOperations.GetBills(context.Background(), &models.Bill{})

	s.NoError(err)
	s.Greater(len(bills), 0)

}

func (s *DatabaseSuite) Test_014_GetBillById() {
	bill, err := s.databaseOperations.GetBillById(context.Background(), strconv.FormatUint(uint64(s.test_bill_id), 10))

	s.NoError(err)
	s.Equal(int64(20), bill.BackupSize)
}

func (s *DatabaseSuite) Test_015_UpdateBill() {
	update_bill, err := s.databaseOperations.UpdateBill(context.Background(), &models.Bill{BackupSize: 200}, strconv.FormatUint(uint64(s.test_bill_id), 10))

	s.NoError(err)
	s.Equal(int64(200), update_bill.BackupSize)
}
