package migrations

import (
	"errors"
	"io"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/log"
)

func TestK20230518180017_Up(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20230518180017{}

	mock.ExpectExec(DeleteNotNullColumn).WillReturnResult(sqlmock.NewResult(1, 0))
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

func TestK20230518180017_Down(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20230518180017{}
	
	mock.ExpectExec(AddNotNullColumn).WillReturnResult(sqlmock.NewResult(1, 0))
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