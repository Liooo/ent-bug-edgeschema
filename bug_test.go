package bug

import (
	"context"
	"testing"

	"ariga.io/atlas/sql/migrate"
	entm "entgo.io/bug/ent/migrate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"entgo.io/bug/ent/enttest"
)

func TestBugSQLite(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	migrationFilePath := "./migrations"
	dir, err := migrate.NewLocalDir(migrationFilePath)

	// delete everything and zero start
	if err != nil {
		panic(err)
	}
	_, err = client.ExecContext(ctx, "DROP TABLE cells;")
	if err != nil {
		panic(err)
	}
	_, err = client.ExecContext(ctx, "DROP TABLE records;")
	if err != nil {
		panic(err)
	}
	_, err = client.ExecContext(ctx, "DROP TABLE data_fields;")
	if err != nil {
		panic(err)
	}

	var migrationOptions = []schema.MigrateOption{
		entm.WithGlobalUniqueID(false),
		entm.WithForeignKeys(true),
		entm.WithDropIndex(true),
		entm.WithDropColumn(true),
		schema.WithAtlas(true),
		schema.WithDir(dir),
	}

	err = client.Schema.NamedDiff(ctx, "test", migrationOptions...)
	if err != nil {
		panic(err)
	}

	println("done!")
}
