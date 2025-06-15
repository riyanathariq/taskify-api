package database

import (
	"flag"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"github.com/riyanathariq/taskify-api/internal/config"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	flags   = flag.NewFlagSet("db:migrate", flag.ExitOnError)
	dir     = flags.String("dir", "database/migration", "directory with migration files")
	table   = flags.String("table", "db_migration", "migrations table name")
	verbose = flags.Bool("verbose", false, "enable verbose mode")
	help    = flags.Bool("guide", false, "print help")
	version = flags.Bool("version", false, "print version")
	dsnParm = flags.String("dsn", "", "database dsn")
)

func DatabaseMigration(cfg *config.Config) {

	var dsn = map[string]func(cfg *config.Config) string{
		"mysql": func(cfg *config.Config) string {
			return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
				cfg.Database.User, cfg.Database.Password,
				cfg.Database.Host, cfg.Database.Port,
				cfg.Database.Name)
		},
		"postgres": func(cfg *config.Config) string {
			return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				cfg.Database.User, cfg.Database.Password,
				cfg.Database.Host, cfg.Database.Port,
				cfg.Database.Name)
		},
	}

	flags.Usage = usage
	flags.Parse(os.Args[2:])

	if *version {
		fmt.Println(goose.MaxVersion)
		return
	}
	if *verbose {
		goose.SetVerbose(true)
	}

	goose.SetTableName(*table)

	args := flags.Args()

	if len(args) == 0 || *help {
		flags.Usage()
		return
	}

	switch args[0] {
	case "create":
		if len(args) < 3 {
			log.Fatalf("invalid arguments: need migration name and type (sql|go), got: %v", args)
		}

		migrationName := args[1]
		migrationType := args[2]

		if migrationType != "sql" && migrationType != "go" {
			log.Fatalf("invalid migration type: %s. Use 'sql' or 'go'", migrationType)
		}

		if err := goose.Run("create", nil, *dir, migrationName, migrationType); err != nil {
			log.Fatalf("goose create: %v", err)
		}

		log.Printf("Migration %s.%s created successfully in %s", migrationName, migrationType, *dir)
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	//dbstring := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	dbStringCon := *dsnParm

	if strings.Trim(dbStringCon, "") == "" {
		fn, ok := dsn["postgres"]
		if !ok {
			log.Printf(fmt.Sprintf("invalid driver %s", "postgres"))
		}

		dbStringCon = fn(cfg)
	}

	db, err := goose.OpenDBWithDriver("postgres", dbStringCon)

	if err != nil {
		log.Printf(err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("db migrate: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("db migrate run: %v", err)
	}
}

func usage() {
	fmt.Println(usageCommands)
}

var (
	usageCommands = `
  --dir string     directory with migration files (default "database/migration")
  --guide          print help
  --table string   migrations table name (default "db_migration")
  --verbose        enable verbose mode
  --version        print version
  --dsn 		   database dsn

Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)

// normalizeMySQLDSN parses the dsn used with the mysql driver to always have
// the parameter `parseTime` set to true. This allows internal goose logic
// to assume that DATETIME/DATE/TIMESTAMP can be scanned into the time.Time
// type.
func normalizeDBString(driver string, str string) string {
	if driver == "mysql" {
		var err error
		str, err = normalizeMySQLDSN(str)
		if err != nil {
			log.Fatalf("failed to normalize MySQL connection string: %v", err)
		}
	}
	return str
}

func normalizeMySQLDSN(dsn string) (string, error) {
	config, err := mysql.ParseDSN(dsn)
	if err != nil {
		return "", err
	}
	config.ParseTime = true
	return config.FormatDSN(), nil
}
