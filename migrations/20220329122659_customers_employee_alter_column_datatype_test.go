package migrations

import (
	"errors"
	"io"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/log"
)

func TestK20220329122659_Up(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20220329122659{}

	mock.ExpectExec(AlterType).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectExec(invalidQuery).WillReturnError(errors.New("invalid migrations"))

	testCases := []struct {
		desc string 
		err error
	}{
		{"success", nil},
		{"failure", errors.New("invalid migrations")},
	}

	for i, tc := range testCases {
		err := k.Up(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "Test[%d], failed - %s", i, tc.desc)
	}
}

func TestK20220329122659_Down(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20220329122659{}
	
	mock.ExpectExec(ResetType).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectExec(invalidQuery).WillReturnError(errors.New("invalid migrations"))

	Cases := []struct {
		desc string
		err error
	}{
		{"success", nil},
		{"failure", errors.New("invalid migration")},
	}

	for i, tc := range Cases {
		err := k.Down(&db, log.NewMockLogger(io.Discard))
		assert.IsType(t, tc.err, err, "Test %d - failed - %s", i, tc.desc)
	}
}