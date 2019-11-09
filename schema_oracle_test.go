package schema_test

import (
	"fmt"

	// _ "gopkg.in/goracle.v2" // goracle
	// _ "github.com/mattn/go-oci8" // oci8
	// _ "gopkg.in/rana/ora.v4" // ora

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = XDescribe("schema", func() {
	Context("using github.com/go-goracle/goracle (Oracle)", func() {

		const (
			user = "test_user"
			pass = "password"
			host = "localhost"
			port = "32772"
			dbs  = "xe"
		)

		var oracle = &testParams{
			DriverName: "goracle",
			// DriverName: "oci8",
			// DriverName: "ora",
			ConnStr: fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbs),

			CreateDDL: []string{
				`CREATE TABLE web_resource (
					id				NUMBER NOT NULL,
					url				NVARCHAR2(1024) NOT NULL UNIQUE,
					content			BLOB,
					compressed_size	NUMBER NOT NULL,
					content_length	NUMBER NOT NULL,
					content_type	NVARCHAR2(128) NOT NULL,
					etag			NVARCHAR2(128) NOT NULL,
					last_modified	NVARCHAR2(128) NOT NULL,
					created_at		TIMESTAMP WITH TIME ZONE NOT NULL,
					modified_at		TIMESTAMP WITH TIME ZONE,
					PRIMARY KEY (id)
				)`,
				// "CREATE INDEX idx_web_resource_url ON web_resource(url)",
				"CREATE INDEX idx_web_resource_created_at ON web_resource(created_at)",
				"CREATE INDEX idx_web_resource_modified_at ON web_resource(modified_at)",
				"CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource",
				// Tests for correct identifer escaping.
				"CREATE TABLE {blanks in name} (id NUMBER, PRIMARY KEY (id))",
				"CREATE TABLE {[brackets] in name} (id INTEGER, PRIMARY KEY (id))",
				`CREATE TABLE {"d.quotes" in name} (id INTEGER, PRIMARY KEY (id))`,
				"CREATE TABLE {'s.quotes' in name} (id INTEGER, PRIMARY KEY (id))",
				"CREATE TABLE {{braces}} in name} (id INTEGER, PRIMARY KEY (id))",
				"CREATE TABLE {`backticks` in name} (id INTEGER, PRIMARY KEY (id))",
				`CREATE TABLE {backslashes\in\name} (id INTEGER, PRIMARY KEY (id))`,
			},
			DropDDL: []string{
				`DROP TABLE {backslashes\in\name}`,
				"DROP TABLE {`backticks` in name}",
				"DROP TABLE {{braces}} in name}",
				"DROP TABLE {'s.quotes' in name}",
				`DROP TABLE {"d.quotes" in name}`,
				"DROP TABLE {[brackets] in name}",
				"DROP TABLE {blanks in name}",
				"DROP VIEW web_resource_view",
				"DROP INDEX idx_web_resource_modified_at",
				"DROP INDEX idx_web_resource_created_at",
				// "DROP INDEX idx_web_resource_url",
				"DROP TABLE web_resource",
			},

			TableExpRes: []string{
				"ID NUMBER",
				"URL NVARCHAR2",
				"CONTENT BLOB",
				"COMPRESSED_SIZE NUMBER",
				"CONTENT_LENGTH NUMBER",
				"CONTENT_TYPE NVARCHAR2",
				"ETAG NVARCHAR2",
				"LAST_MODIFIED NVARCHAR2",
				"CREATED_AT TIMESTAMP WITH TIMEZONE",
				"MODIFIED_AT TIMESTAMP WITH TIMEZONE",
			},
			ViewExpRes: []string{
				"ID NUMBER",
				"URL NVARCHAR2",
			},

			TableNamesExpRes: []string{
				"WEB_RESOURCE",
				// Tests for correct identifer escaping.
				"BLANKS IN NAME",
				"[BRACKETS] IN NAME",
				`"D.QUOTES" IN NAME`,
				"'S.QUOTES' IN NAME",
				"{BRACES} IN NAME",
				"`BACKTICKS` IN NAME",
				`BACKSLASHES\IN\NAME`,
			},
			ViewNamesExpRes: []string{
				"WEB_RESOURCE_VIEW",
			},
		}

		SchemaTestRunner(oracle)
	})
})

// func oraDump(db *sql.DB) error {

// 	//SELECT table_name FROM user_tables
// 	rows, err := db.Query(`
// 		SELECT *
// 		  FROM user_tables
//    `)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	ci, err := rows.ColumnTypes()
// 	if err != nil {
// 		return err
// 	}
// 	for _, c := range ci {
// 		log.Printf("%v", c)
// 	}

// 	cols, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}
// 	vals := make([]interface{}, len(cols))
// 	for i, _ := range cols {
// 		vals[i] = new(sql.RawBytes)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(vals...)
// 		if err != nil {
// 			// return err
// 			log.Printf("%v", err)
// 		}
// 		s := ""
// 		for _, v := range vals {
// 			s = s + fmt.Sprintf("%s ", v)
// 		}
// 		log.Print(s)
// 	}
// 	return nil
// }
