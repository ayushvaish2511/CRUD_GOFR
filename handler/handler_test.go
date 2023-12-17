package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ayushvaish2511/CRUD_GOFR/model"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
	"gofr.dev/pkg/gofr/responder"
)

type mockStore struct{}

// GetById implements store.Store.
func (mockStore) GetById(ctx *gofr.Context, id uuid.UUID) (model.Customer, error) {
	uid := uuid.MustParse("313c08cd-9269-4716-aab7-68342b9efd2b")
	if id == uid {
		return model.Customer{ID: uid, Name: "test", Email: "test@xyz.com", Phone: "1234567890"}, nil
	}

	return model.Customer{}, errors.EntityNotFound{Entity: "customer", ID: "37387615-aead-4b28-9adc-78c1eb714ca7"}
}

// Update implements store.Store.
func (mockStore) Update(ctx *gofr.Context, customer model.Customer) (model.Customer, error) {
	if customer.Name == "some name" {
		return model.Customer{}, nil
	}

	return model.Customer{}, errors.Error("error updating customer")
}

func (m mockStore) Get(ctx *gofr.Context) ([]model.Customer, error) {
	p := ctx.Param("mock")
	if p == "success" {
		return nil, nil
	}

	return nil, errors.Error("error fetching customers")
}

func (m mockStore) Create(_ *gofr.Context, customer model.Customer) (model.Customer, error) {
	uid := uuid.MustParse("313c08cd-9269-4716-aab7-68342b9efd2b")

	switch customer.Name {
	case "test":
		return model.Customer{ID: uid, Name: "success", Email: "success@gmail.com", Phone: "1234567890"}, nil
	case "mock body error":
		return model.Customer{}, errors.InvalidParam{Param: []string{"body"}}
	}

	return model.Customer{}, errors.Error("error adding new customer")
}

func (m mockStore) Delete(ctx *gofr.Context, _ uuid.UUID) error {
	uid := "313c08cd-9269-4716-aab7-68342b9efd2b"
	if ctx.PathParam("id") == uid {
		return nil
	}

	return errors.Error("error deleting customer")
}

func TestModel_AddCustomer(t *testing.T) {
	h := New(mockStore{})

	app := gofr.New()

	Cases := []struct {
		desc string
		body []byte
		err  error
	}{
		{
			"create success",
			[]byte(`{"name":"test","email":"test@xyz.com","phone":1234567890}`),
			nil,
		},
		{
			"create invalid body",
			[]byte(`mock body error`),
			errors.InvalidParam{Param: []string{"body"}},
		},
		{
			"create error",
			[]byte(`{"name":"creation error","email":"name@gmail.com","phone":1234567890}`),
			errors.Error("error adding new customer"),
		},
	}

	for i, tc := range Cases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://dummy", bytes.NewReader(tc.body))

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)

		_, err := h.Create(ctx)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestModel_UpdateCustomer(t *testing.T) {
	h := New(mockStore{})

	app := gofr.New()

	tests := []struct {
		desc string
		body []byte
		err  error
		id   string
	}{
		{
			"missing id",
			nil,
			errors.MissingParam{Param: []string{"id"}},
			"",
		},
		{
			"invalid id",
			nil,
			errors.InvalidParam{Param: []string{"id"}},
			"abc123",
		},
		{
			"invalid body",
			[]byte(`{`), errors.InvalidParam{Param: []string{"body"}},
			"313c08cd-9269-4716-aab7-68342b9efd2b",
		},
		{
			"update success",
			[]byte(`{"name":"test"}`),
			nil,
			"313c08cd-9269-4716-aab7-68342b9efd2b",
		},
		{
			"update error",
			[]byte(`{"name":"creation error"}`),
			errors.Error("error updating customer"),
			"313c08cd-9269-4716-aab7-68342b9efd2b",
		},
	}

	for i, tc := range tests {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "http://dummy", bytes.NewReader(tc.body))

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		_, err := h.Update(ctx)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
func TestModel_GetCustomerById(t *testing.T) {
	h := New(mockStore{})

	app := gofr.New()
	uid := uuid.MustParse("313c08cd-9269-4716-aab7-68342b9efd2b")

	Cases := []struct {
		desc string
		id   string
		resp interface{}
		err  error
	}{
		{
			"get by id success",
			"313c08cd-9269-4716-aab7-68342b9efd2b",
			model.Customer{
				ID:    uid,
				Name:  "test",
				Email: "test@xyz.com",
				Phone: "1234567890",
			},
			nil,
		},
		{
			"invalid id",
			"absd123",
			nil,
			errors.InvalidParam{
				Param: []string{"id"},
			},
		},
		{
			"missing id",
			"",
			nil,
			errors.MissingParam{Param: []string{"id"}},
		},
		{
			"id not found",
			"37387615-aead-4b28-9adc-78c1eb714ca7",
			nil,
			errors.EntityNotFound{
				Entity: "customer",
				ID:     "37387615-aead-4b28-9adc-78c1eb714ca7",
			},
		},
	}

	for i, tc := range Cases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		resp, err := h.GetById(ctx)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)

		assert.Equal(t, tc.resp, resp, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestModel_DeleteCustomer(t *testing.T) {
	h := New(mockStore{})

	app := gofr.New()

	Cases := []struct {
		desc string
		id   string
		err  error
	}{
		{
			"delete success",
			"313c08cd-9269-4716-aab7-68342b9efd2b",
			nil,
		},
		{
			"delete fail",
			"37387615-aead-4b28-9adc-78c1eb714ca7",
			errors.Error("error deleting customer"),
		},
		{
			"invalid id",
			"absd123",
			errors.InvalidParam{Param: []string{"id"}},
		},
		{
			"missing id",
			"",
			errors.MissingParam{Param: []string{"id"}},
		},
	}

	for i, tc := range Cases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy", nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)

		ctx := gofr.NewContext(res, req, app)

		ctx.SetPathParams(map[string]string{
			"id": tc.id,
		})

		_, err := h.Delete(ctx)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}

func TestModel_GetCustomers(t *testing.T) {
	h := New(mockStore{})

	app := gofr.New()

	Cases := []struct {
		desc         string
		mockParamStr string
		err          error
	}{
		{
			"get success",
			"mock=success",
			nil,
		},
		{
			"get fail",
			"",
			errors.Error("error fetching customer listing"),
		},
	}

	for i, tc := range Cases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "http://dummy?"+tc.mockParamStr, nil)

		req := request.NewHTTPRequest(r)
		res := responder.NewContextualResponder(w, r)
		ctx := gofr.NewContext(res, req, app)

		_, err := h.Get(ctx)
		assert.Equal(t, tc.err, err, "TEST[%d], failed.\n%s", i, tc.desc)
	}
}
