package dialectquery

import (
	"fmt"
)

type Oracle struct{}

var _ Querier = (*Oracle)(nil)

func (m *Oracle) CreateTable(tableName string) string {
	q := `CREATE TABLE %s (
		id number NOT NULL,
		version_id number NOT NULL,
		is_applied number NOT NULL,
		tstamp DATE default sysdate,
		PRIMARY KEY(id)
	)`
	//		CREATE SEQUENCE migrate_seq START WITH 1000 INCREMENT BY   1 NOCACHE NOCYCLE;

	return fmt.Sprintf(q, tableName)
}

func (m *Oracle) InsertVersion(tableName string) string {

	q := `INSERT INTO %s (id,version_id, is_applied) VALUES (nvl((select max(id) from %s),0) + 1, :1, :2)`
	return fmt.Sprintf(q, tableName, tableName)
}

func (m *Oracle) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id=:1`
	return fmt.Sprintf(q, tableName)
}

func (m *Oracle) GetMigrationByVersion(tableName string) string {

	q := `SELECT * FROM (SELECT tstamp, is_applied FROM %s WHERE version_id=:1 ORDER BY tstamp DESC) where rownum <= 1`
	return fmt.Sprintf(q, tableName)
}

func (m *Oracle) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied from %s ORDER BY id DESC`
	return fmt.Sprintf(q, tableName)
}
