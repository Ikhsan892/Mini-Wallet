package wallet

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"wallet/application/infrastructures"
)

type App interface {
	DB() *gorm.DB
	Context() context.Context
	WithDBConnection(name string) *gorm.DB
	GetAppDir() string
	CreateSQliteDatabaseFile() error
	Settings() Config
	Migration() *migrate.Migrate
	Shutdown() error
	ZapLogger() *zap.Logger
}

func New() *Wallet {
	ctx := context.Background()
	cfg := newConfig()
	logger := NewZapLogger()
	baseDir, _ := os.Getwd()

	g := &Wallet{
		sqliteConns:  make(map[string]*gorm.DB),
		baseDir:      baseDir,
		notification: make(chan struct{}, 1),
		ctx:          ctx,
		cfg:          cfg,
		logger:       logger,
	}

	return g
}

type Wallet struct {
	sqliteConns  map[string]*gorm.DB
	ctx          context.Context
	cfg          Config
	logger       *zap.Logger
	baseDir      string
	notification chan struct{}
}

func (g *Wallet) GetAppDir() string {
	return g.baseDir
}

func (g *Wallet) DB() *gorm.DB {
	return g.WithDBConnection("default")
}

func (g *Wallet) WithDBConnection(name string) *gorm.DB {
	return g.sqliteConns[name]
}

func (g *Wallet) initDefaultSQlite(path string) {
	sqlite, err := infrastructures.NewSQLiteWithGorm(fmt.Sprintf("%s/%s", path, "wallet.db"))
	if err != nil {
		g.logger.Fatal(err.Error())
		os.Exit(1)
	}

	g.appendDefaultDatabase("default", sqlite)
}

func (g *Wallet) CreateSQliteDatabaseFile() error {
	f, err := os.Create(fmt.Sprintf("%s/%s", g.baseDir, g.cfg.DB.DBFileName))
	if err != nil {
		g.logger.Fatal("cannot create database file")
		return err
	}

	defer f.Close()

	return nil
}

func (g *Wallet) Settings() Config {
	return g.cfg
}

func (g Wallet) ZapLogger() *zap.Logger {
	return g.logger
}

func (g *Wallet) appendDefaultDatabase(key string, conn *gorm.DB) {
	g.sqliteConns[key] = conn
}

func (g Wallet) IsDBFileExist() bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s", g.baseDir, g.cfg.DB.DBFileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		}
		return false
	}

	return true
}

// starting point of application
func (g *Wallet) Start() {
	g.initDefaultSQlite(g.baseDir)

	if !g.IsDBFileExist() {
		g.ZapLogger().Info("Creating new database file...")
		if err := g.CreateSQliteDatabaseFile(); err != nil {
			g.ZapLogger().Error("Error creating database...", zap.Error(err))
			os.Exit(1)
		}
		g.ZapLogger().Info("Database file created...")

		g.ZapLogger().Info("Running migration...")
		if err := g.Migration().Up(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		g.ZapLogger().Info("Migration successfully...")
	}
}

func (g Wallet) Context() context.Context {
	return g.ctx
}

// for shutdown
func (g *Wallet) Shutdown() error {
	close(g.notification)

	db, _ := g.sqliteConns["default"].DB()
	db.Close()

	return nil
}
