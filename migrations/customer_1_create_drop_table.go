package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type c1createdrop struct {
}

func (k c1createdrop) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Up:customers_employee_create_drop_table")

	_, err := d.DB().Exec(CreateTable)
	if err != nil {
		return err
	}
	return nil
}

func (k c1createdrop) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Down:customers_employee_create_drop_table")

	_, err := d.DB().Exec(DroopTable)
	if err != nil {
		return err
	}
	return nil
}
