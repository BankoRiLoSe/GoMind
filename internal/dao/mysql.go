package dao

import (
	"database/sql"
	"fmt"
	"strings"

	"gomind/internal/config"

	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(cfg config.MySQLConfig, dsn string) (*gorm.DB, error) {
	if err := ensureMySQLDatabase(cfg); err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect mysql: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get mysql sql db: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql: %w", err)
	}

	return db, nil
}

func ensureMySQLDatabase(cfg config.MySQLConfig) error {
	serverDSN := (&driver.Config{
		User:                 cfg.Username,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		AllowNativePasswords: true,
		ParseTime:            cfg.ParseTime,
		Params: map[string]string{
			"charset": cfg.Charset,
			"loc":     cfg.Loc,
		},
	}).FormatDSN()

	db, err := sql.Open("mysql", serverDSN)
	if err != nil {
		return fmt.Errorf("open mysql server connection: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping mysql server: %w", err)
	}

	query := fmt.Sprintf(
		"CREATE DATABASE IF NOT EXISTS %s CHARACTER SET %s",
		quoteMySQLIdentifier(cfg.Database),
		cfg.Charset,
	)
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("create mysql database %q: %w", cfg.Database, err)
	}

	return nil
}

func quoteMySQLIdentifier(name string) string {
	return "`" + strings.ReplaceAll(name, "`", "``") + "`"
}
