package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type c2addDropCountry struct {
}

func (k c2addDropCountry) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running migration Up:customers_employee_update_columns")

	_, err := d.DB().Exec(AddCountry)
	if err != nil {
		return err
	}

	return nil
}

func (k c2addDropCountry) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running migration Down:customers_employee_update_columns")

	_, err := d.DB().Exec(DropCountry)
	if err != nil {
		return err
	}

	return nil
}