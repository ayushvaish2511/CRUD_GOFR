package migrations

import (
	"gofr.dev/pkg/log"

	"gofr.dev/pkg/datastore"
)

type c6addDropPhone struct {
}

func (k c6addDropPhone) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Up:customer_phone_add_drop_table")
	_, err := d.DB().Exec(AddPhone)
	if err != nil {
		return err
	}
	return nil
}

func (k c6addDropPhone) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Infof("Running Migration Down:customer_phone_add_drop_table")
	_, err := d.DB().Exec(DropPhone)
	if err != nil {
		return err
	}
	return nil
}
