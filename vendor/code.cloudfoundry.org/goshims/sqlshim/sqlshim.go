// This file was generated by counterfeiter
// with command: counterfeiter -p -o /Users/pivotal/workspace/nfs-volume-release/src/code.cloudfoundry.org/goshims/sqlshim database/sql
package sqlshim

import (
	"database/sql"
	"database/sql/driver"
)

type SqlShim struct{}

func (sh *SqlShim) Register(name string, driver driver.Driver) {
	sql.Register(name, driver)
}

func (sh *SqlShim) Drivers() []string {
	return sql.Drivers()
}

func (sh *SqlShim) Open(driverName string, dataSourceName string) (SqlDB, error) {
	return sql.Open(driverName, dataSourceName)
}
