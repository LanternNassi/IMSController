package main

import (
	"fmt"
	"testing"

	database_tests "github.com/LanternNassi/IMSController/internal/database/tests"
	server_tests "github.com/LanternNassi/IMSController/internal/server/tests"
	"github.com/stretchr/testify/suite"
)

func TestDatabaseSuite(t *testing.T) {

	databaseSuite, err := database_tests.NewTestDatabaseSuite()

	if err != nil {
		fmt.Println(err)
	}

	suite.Run(t, databaseSuite)
}

func TestServerSuite(t *testing.T) {
	server_suite, err := server_tests.NewTestServerSuite()

	if err != nil {
		fmt.Println(err)
	}

	suite.Run(t, server_suite)
}
