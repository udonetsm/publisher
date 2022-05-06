package db

import (
	"database/sql"

	"github.com/udonetsm/help/helper"
	"github.com/udonetsm/help/models"
)

var Data_json string

func parseyaml() string {
	pgconf := models.Postgres_conf{}
	pgconf = pgconf.StoreConf(helper.Home() + "/.confs/conn_config.yaml")
	return pgconf.Dbname + pgconf.Dbpassword + pgconf.SslMode +
		pgconf.Dbport + pgconf.Dbuser + pgconf.SslMode
}

func sqlDb() *sql.DB {
	pgconnstr := parseyaml()
	sdb, err := sql.Open("pgx", pgconnstr)
	helper.Errors(err, "sqlopen")
	return sdb
}
