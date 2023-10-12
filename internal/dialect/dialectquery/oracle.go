package dialectquery

import (
	"fmt"
	"log"
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
	log.Println("insert")
	q := `INSERT INTO %s (id,version_id, is_applied) VALUES (migrate_seq.nextval,?, ?)`
	return fmt.Sprintf(q, tableName)
}

func (m *Oracle) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id=?`
	return fmt.Sprintf(q, tableName)
}

func (m *Oracle) GetMigrationByVersion(tableName string) string {

	q := `SELECT * FROM (SELECT tstamp, is_applied FROM %s WHERE version_id=? ORDER BY tstamp DESC) where rownum <= 1`
	return fmt.Sprintf(q, tableName)
}

func (m *Oracle) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied from %s ORDER BY id DESC`
	return fmt.Sprintf(q, tableName)
}
