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

	"entgo.io/bug/ent"
)

//func TestBugSQLite(t *testing.T) {
//	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
//	defer client.Close()
//	test(t, client)
//}

//func TestBugMySQL(t *testing.T) {
//	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
//		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
//		t.Run(version, func(t *testing.T) {
//			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
//			defer client.Close()
//			test(t, client)
//		})
//	}
//}
//
func TestBugPostgres(t *testing.T) {
	//for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
    client,err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=liooo dbname=hiway_test sslmode=disable")
    if err != nil {
    	panic(err)
    }

    defer client.Close()
    test(t, client)
	//}
}

//func TestBugMaria(t *testing.T) {
//	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
//		t.Run(version, func(t *testing.T) {
//			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
//			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
//			defer client.Close()
//			test(t, client)
//		})
//	}
//}

func test(t *testing.T, client *ent.Client) {
	migrationFilePath := "./migrations"
	dir, err := migrate.NewLocalDir(migrationFilePath)
    if err != nil {
        panic(err)
    }
	ctx := context.Background()

	//delete everything and zero start
    _, err = client.ExecContext(ctx, "DROP TABLE IF EXISTS likes;")
    if err != nil {
        panic(err)
    }
	_, err = client.ExecContext(ctx, "DROP TABLE IF EXISTS users;")
	if err != nil {
		panic(err)
	}
	_, err = client.ExecContext(ctx, "DROP TABLE IF EXISTS tweets;")
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
    println("yes")
}
