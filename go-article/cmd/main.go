package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Opanpan/go-article-service/config"
	"github.com/Opanpan/go-article-service/internal/controller"
	"github.com/Opanpan/go-article-service/internal/repository"
	database "github.com/Opanpan/go-article-service/internal/repository/mysql"
	"github.com/Opanpan/go-article-service/internal/router"
	"github.com/Opanpan/go-article-service/internal/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

// @title Article Service API
// @version 1.0
// @description This is an API for managing articles
// @host localhost:8080
// @BasePath /
func main() {
	// Load configurations
	configurations := config.LoadConfiguration()

	// Connect to MySQL database
	db, err := database.ConnectDBMysql(configurations)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := runMigrate(db, configurations); err != nil {
		log.Errorf("Error running migrations: %v", err)
	} else {
		log.Info("Migrations run successfully")
	}

	// repository
	articleRepo := repository.NewArticleRepository(db)

	// service
	articleService := service.NewArticleService(articleRepo)

	//contorller
	articleController := controller.NewArticleController(articleService)

	app := router.NewRouter(articleController)

	app.SetupRouter(configurations.Get("PORT"))
}

func runMigrate(db *sql.DB, config config.Config) error {
	if tableExists, err := checkTableExists(db, "article"); err != nil {
		return err
	} else if !tableExists {
		if err := runMigrations(db, config); err != nil {
			return err
		}
	}
	return nil
}

func checkTableExists(db *sql.DB, tableName string) (bool, error) {
	rows, err := db.Query(fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName))
	if err != nil {
		return false, fmt.Errorf("could not check if table '%s' exists: %w", tableName, err)
	}
	return rows.Next(), nil
}

func runMigrations(db *sql.DB, config config.Config) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("could not start SQL driver: %w", err)
	}

	fmt.Println("INIII", config.Get("MIGRATION_PATH"))
	m, err := migrate.NewWithDatabaseInstance(
		config.Get("MIGRATION_PATH"),
		config.Get("DB_TYPE"), driver)
	if err != nil {
		return fmt.Errorf("could not initialize migration instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if strings.Contains(err.Error(), "Dirty database") {
			return handleDirtyDatabase(m, config)
		} else if err != migrate.ErrNoChange {
			return fmt.Errorf("could not run up migrations: %w", err)
		}
	}
	return nil
}

func handleDirtyDatabase(m *migrate.Migrate, config config.Config) error {
	versionMigrationFail, err := strconv.Atoi(config.Get("MIGRATION_RETRY_VERSION"))
	if err != nil {
		return fmt.Errorf("could not convert migration retry version: %w", err)
	}
	logrus.Warn("Database is in a dirty state. Attempting to fix...")
	if err := m.Force(versionMigrationFail); err != nil {
		return fmt.Errorf("could not force version: %w", err)
	}
	logrus.Info("Forced version successfully. Retrying migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run up migrations after forcing: %w", err)
	}
	return nil
}
