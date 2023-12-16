package migrations

import (
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
)

func All() map[string]dbmigration.Migrator {
	return map[string]dbmigration.Migrator{

		"createdrop": c1createdrop{},
		"addDropCountry": c2addDropCountry{},
		"alterResetType": c3alterResetType{},
		"nullAddDelete": c4nullAddDelete{},
		"alterResetPrimary": c5alterResetPrimary{},
		"addDropPhone": c6addDropPhone{},
	}
}