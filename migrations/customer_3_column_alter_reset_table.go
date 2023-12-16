package migrations

import (
	"gofr.dev/pkg/log"

	"gofr.dev/pkg/datastore"
)

type c3alterResetType struct {
}

func (k c3alterResetType) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Up:customer_column_alter_reset_table")
	_, err := d.DB().Exec(AlterType)
	if err != nil {
		return err
	}
	return nil
}

func (k c3alterResetType) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Down:customer_column_alter_reset_table")
	_, err := d.DB().Exec(ResetType)
	if err != nil {
		return err
	}
	return nil
}
