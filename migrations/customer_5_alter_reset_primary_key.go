package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type c5alterResetPrimary struct {
}

func (k c5alterResetPrimary) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running migration Up:customer_alter_reset_primary_key")
	_, err := d.DB().Exec(AlterPrimaryKey)
	if err != nil {
		return err
	}
	return nil
}

func (k c5alterResetPrimary) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Down:customer_alter_reset_primary_key")
	_, err := d.DB().Exec(ResetPrimaryKey)
	if err != nil {
		return err
	}
	return nil
}
