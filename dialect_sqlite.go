package schema

import (
	"database/sql"
)

const sqliteAllColumns = `SELECT * FROM %s LIMIT 0`

const sqliteTableNames = `SELECT name FROM sqlite_master WHERE type = 'table' ORDER BY name`

const sqliteTableNamesWithSchema = `SELECT "" AS schema, name FROM sqlite_master WHERE type = 'table' ORDER BY name`

const sqliteViewNames = `SELECT name FROM sqlite_master WHERE type = 'view' ORDER BY name`

const sqliteViewNamesWithSchema = `SELECT "" AS schema, name FROM sqlite_master WHERE type = 'view' ORDER BY name`

const sqlitePrimaryKey = `SELECT name FROM pragma_table_info(?) WHERE pk > 0 ORDER BY pk`

type sqliteDialect struct{}

func (sqliteDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (sqliteDialect) PrimaryKey(db *sql.DB, schema, name string) ([]string, error) {
	if schema == "" {
		return fetchNames(db, sqlitePrimaryKey, "", name)
	}
	return fetchNames(db, sqlitePrimaryKey, "", name)
}

func (d sqliteDialect) Table(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, sqliteAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, sqliteAllColumns, schema, name, d.escapeIdent)
}

func (sqliteDialect) TableNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, sqliteTableNamesWithSchema, "", "")
}

func (d sqliteDialect) View(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, sqliteAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, sqliteAllColumns, schema, name, d.escapeIdent)
}

func (sqliteDialect) ViewNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, sqliteViewNamesWithSchema, "", "")
}
