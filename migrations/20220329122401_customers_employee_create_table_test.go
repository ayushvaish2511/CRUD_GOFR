package migrations

import (
	"errors"
	"io"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

const invalidQuery = "invalid query"

func initializeTest(t *testing.T) (sqlmock.Sqlmock, datastore.DataStore) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error %s was not expected while opening mock database connection", err)
	}

	dataStore := datastore.DataStore{ORM: db}

	return mock, dataStore
}

func TestK20220329122401_UP(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20220329122401{}

	mock.ExpectExec(CreateTable).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectExec(invalidQuery).WillReturnError(errors.New("invalid migration"))

	testCases := []struct {
		desc string
		err error
	}{
		{"sucess", nil},
		{"failure", errors.New("invalid migration")},
	}
	
	for i, tc := range testCases {
		err := k.Up(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestK20220329122401_Down(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20220329122401{}

	mock.ExpectExec(DroopTable).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectExec(invalidQuery).WillReturnError(errors.New("invalid migration"))

	testCses := []struct {
		desc string
		err error
	}{
		{"sucess", nil},
		{"failure", errors.New("invalid migration")},
	}

	for i, tc := range testCses {
		err := k.Down(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "Test[%d] - failed - %s", i, tc.desc)
	}
}