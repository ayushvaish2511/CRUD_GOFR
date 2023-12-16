package migrations

import (
	"gofr.dev/pkg/log"

	"gofr.dev/pkg/datastore"
)

type c4nullAddDelete struct {
}

func (k c4nullAddDelete) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Up:customer_null_add_delete_table")
	_, err := d.DB().Exec(AddNotNullColumn)
	if err != nil {
		return err
	}
	return nil
}

func (k c4nullAddDelete) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Down:customer_null_add_delete_table")
	_, err := d.DB().Exec(DeleteNotNullColumn)
	if err != nil {
		return err
	}
	return nil
}
