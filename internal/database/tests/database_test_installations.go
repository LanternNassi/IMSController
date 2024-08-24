package tests

import (
	"context"
	"strconv"

	"github.com/LanternNassi/IMSController/internal/models"
)

func (s *DatabaseSuite) Test_005_AddInstallation() {
	installation := &models.Installation{
		ClientID:          *s.test_client_id,
		Installation_type: "server",
		Computer_name:     "Cyber",
		IMS_version:       "1.1.4",
		Operating_system:  "Win-11",
		RAM:               "8GB",
		Processor:         "Intel core i3",
		Active:            "true",
	}

	created_inst, err := s.databaseOperations.AddInstallation(context.Background(), installation)
	s.test_installation_id = created_inst.ID
	s.NoError(err)
	s.Equal(installation, created_inst)

}

func (s *DatabaseSuite) Test_006_GetInstallations() {
	installations, err := s.databaseOperations.GetInstallations(context.Background(), &models.Installation{})

	s.NoError(err)
	s.Greater(len(installations), 0)

}

func (s *DatabaseSuite) Test_007_GetInstallationById() {
	installation, err := s.databaseOperations.GetInstallationById(context.Background(), strconv.FormatUint(uint64(s.test_installation_id), 10))

	s.NoError(err)

	s.Equal("server", installation.Installation_type)
}

func (s *DatabaseSuite) Tes_008_tUpdateInstallation() {

	updated_inst := models.Installation{
		RAM:           "16GB",
		Computer_name: "Nessim",
		IMS_version:   "1.1.6",
	}

	_inst, err := s.databaseOperations.UpdateInstallation(context.Background(), &updated_inst, strconv.FormatUint(uint64(s.test_installation_id), 10))

	s.NoError(err)
	s.Equal(updated_inst.RAM, _inst.RAM)
}
