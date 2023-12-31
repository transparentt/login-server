package migration

import (
	"github.com/transparentt/login-server/config"
	"github.com/transparentt/login-server/pkg/rethinkdb/logic"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func migrateV1(session *r.Session) error {

	config := config.LoadConfig()

	_, err := r.DB(config.Database).TableCreate(logic.UserTable).RunWrite(session)
	if err != nil {
		return err
	}

	return nil
}
