package main

import (
	"fmt"
	"testing"

	tests "github.com/LanternNassi/IMSController/internal/database/tests"
	"github.com/stretchr/testify/suite"
)

func TestDatabaseSuite(t *testing.T) {

	databaseSuite, err := tests.NewTestDatabaseSuite()

	if err != nil {
		fmt.Println(err)
	}

	suite.Run(t, databaseSuite)
}
