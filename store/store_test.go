package store

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ayushvaish2511/CRUD_GOFR/model"
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func TestCore(t *testing.T) {
	app := gofr.New()

	seeder := datastore.NewSeeder(&app.DataStore, "../db")
	seeder.ResetCounter = true

	createTable(app)
}

func createTable(app *gofr.Gofr) {
	_, err := app.DB().Exec("DROP TABLE IF EXISTS customers;")

	if err != nil {
		return
	}

	_, err = app.DB().Exec("CREATE TABLE IF NOT EXISTS customers" +
		" (id varchar(36) PRIMARY KEY , name varchar(50) , email varchar(50) , phone varchar(50));")
	if err != nil {
		return
	}
}

// To Add Customers into the mock data
func TestAddingCustomer(t *testing.T) {
	ctx := gofr.NewContext(nil, nil, gofr.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("mock connection failed")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()
	Cases := []struct {
		description string
		customer    model.Customer
		mockError   error
		err         error
	}{
		{
			"Valid Case",
			model.Customer{
				Name:  "A",
				Email: "A@XYZ.com",
				Phone: "1111111111",
			},
			nil,
			nil,
		},
		{
			"Database Error",
			model.Customer{
				Name:  "B",
				Email: "B@XYZ.com",
				Phone: "2222222222",
			},
			errors.DB{},
			errors.DB{Err: errors.DB{}},
		},
	}

	for i, tc := range Cases {
		row := mock.NewRows([]string{"name", "email", "phone"}).AddRow(tc.customer.Name, tc.customer.Email, tc.customer.Phone)
		mock.ExpectQuery("INSERT INTO").
			WithArgs(sqlmock.AnyArg(), tc.customer.Name, tc.customer.Email, tc.customer.Phone).
			WillReturnRows(row).WillReturnError(tc.mockError)

		store := New()
		resp, err := store.Create(ctx, tc.customer)

		ctx.Logger.Log(resp)
		assert.Equal(t, tc.err, err, "Test %d failed - %s", i, tc.description)
	}
}

func TestDeletingCustomer(t *testing.T) {
	uid := uuid.MustParse("313c08cd-9269-4716-aab7-68342b9efd2b")
	uid1 := uuid.MustParse("37387615-aead-4b28-9adc-78c1eb714ca1")

	Cases := []struct {
		desc         string
		id           uuid.UUID
		mockError    error
		rowsAffected int64
		err          error
	}{
		{
			"delete success test #1",
			uid,
			nil,
			1,
			nil,
		},
		{
			"delete failure test #2",
			uid1,
			errors.DB{},
			0,
			errors.DB{Err: errors.DB{}},
		},
	}
	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("mockdata is not initialized")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	for i, tc := range Cases {
		mock.ExpectExec("DELETE FROM customers").WithArgs(tc.id).
			WillReturnResult(sqlmock.NewResult(1, tc.rowsAffected)).
			WillReturnError(tc.mockError)

		store := New()

		err := store.Delete(ctx, tc.id)
		if err != tc.err {
			t.Errorf("TEST CASE %v FAILED, Expected: %v, Got: %v", i, nil, err)
		}
	}
}

func TestGetCustomers(t *testing.T) {
	uid := uuid.MustParse("637739bf-9716-431b-9087-546e650c4269")

	Cases := []struct {
		desc      string
		customer  model.Customer
		mockError error
		err       error
	}{
		{
			"Get existent data",
			model.Customer{
				ID:    uid,
				Name:  "B",
				Email: "B@xyz.com",
				Phone: "2222222222",
			},
			nil,
			nil,
		},
		{
			"db connection failed",
			model.Customer{},
			errors.DB{},
			errors.DB{Err: errors.DB{}},
		},
	}

	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("mock not initialized")
	}

	ctx.Context = context.Background()
	ctx.DataStore = datastore.DataStore{ORM: db}

	for i, tc := range Cases {
		rows := mock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone)
		mock.ExpectQuery("SELECT id,name").WillReturnRows(rows).WillReturnError(tc.mockError)

		store := New()
		_, err := store.Get(ctx)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestGetCustomersByID(t *testing.T) {
	uid := uuid.MustParse("637739bf-9716-431b-9087-546e650c4269")
	invalidUID := uuid.MustParse("37387615-aead-4b28-9adc-78c1eb714ca5")

	Cases := []struct {
		description string
		customer    model.Customer
		id          uuid.UUID
		mockError   error
		err         error
	}{
		{
			"Get existent id", 
			model.Customer {
				ID: uid, 
				Name: "B", 
				Email: "B@xyz.com", 
				Phone: "2222222222",
			},
			uid, 
			nil, 
			nil,
		},
		{
			"Get non existent id", 
			model.Customer{}, 
			invalidUID, 
			sql.ErrNoRows,
			errors.EntityNotFound{
				Entity: "customer", 
				ID: "37387615-aead-4b28-9adc-78c1eb714ca5",
			},
		},
	}
	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("mock is not initialized")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	for i, tc := range Cases {
		rows := mock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(tc.customer.ID, tc.customer.Name, tc.customer.Email, tc.customer.Phone)
		mock.ExpectQuery("SELECT id,name").WithArgs(tc.id).WillReturnRows(rows).WillReturnError(tc.mockError)

		store := New()

		_, err := store.GetById(ctx, tc.id)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.description)
	}
}

func TestUpdateCustomer(t *testing.T) {
	uid := uuid.MustParse("637739bf-9716-431b-9087-546e650c4269")
	uid1 := uuid.MustParse("37387615-aead-4b28-9adc-78c1eb714ca1")

	Cases := []struct {
		description     string
		customer model.Customer
		err      error
	}{
		{
			"update success", 
			model.Customer{
				ID: uid, 
				Name: "F",
			}, 
			nil,
		},
		{
			"update fail", 
			model.Customer{
				ID: uid1, 
				Name: "very-long-mock-name-lasdjflsdjfljasdlfjsdlfjsdfljlkj"}, 
				errors.DB{},
			},
	}

	app := gofr.New()
	ctx := gofr.NewContext(nil, nil, app)
	db, mock, err := sqlmock.New()

	if err != nil {
		ctx.Logger.Error("mock is not initialized")
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	for i, tc := range Cases {

		mock.ExpectExec("UPDATE customers").
			WithArgs(tc.customer.Name, tc.customer.Email, tc.customer.Phone, tc.customer.ID).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)

		ctx := gofr.NewContext(nil, nil, app)
		ctx.Context = context.Background()

		store := New()

		_, err := store.Update(ctx, tc.customer)

		if _, ok := err.(errors.DB); err != nil && ok == false {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.err, err, tc.description)
		}
	}
}