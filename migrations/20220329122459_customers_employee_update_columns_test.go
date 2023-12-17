package migrations

import (
	"errors"
	"io"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/log"
)

func TestK20220329122459_Up(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20220329122459{}

	mock.ExpectExec(AddCountry).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectExec(invalidQuery).willReturnError(errors.New("invalid migration"))

	testCases := []struct {
		desc string
		err error
	}{
		{"success", nil},
		{"failure", errors.New("invalid migration")},
	}

	for i, tc := range testCases {
		err := k.Up(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "TEST[%d], failed - %s", i, tc.desc)
	}
}

func TestK20220329122459_Up(t *testing.T) {
	mock, db := initializeTest(t)
	k := K20220329122459{}

	mock.ExpectExec(DropCountry).WillReturnResult(sqlmock.NewResult(1, 0))
	mock.ExpectExec(invalidQuery).willReturnError(errors.New("invalid migration"))

	testCases := []struct {
		desc string
		err error
	}{
		{"success", nil},
		{"failure", errors.New("invalid migration")},
	}

	for i, tc := range testCases {
		err := k.Down(&db, log.NewMockLogger(io.Discard))

		assert.IsType(t, tc.err, err, "TEST[%d], failed - %s", i, tc.desc)
	}
}