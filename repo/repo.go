package repo

import (
	"context"
	"database/sql"
	"time"

	"todo/config"
	"todo/logger"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type Repo struct {
	DB *sql.DB
}

func NewRepo(ctx context.Context, conf *config.Config) (*Repo, error) {
	dbconn, err := newDBClient(conf.DBConfig.User, conf.DBConfig.Pass, conf.DBConfig.Host, conf.DBConfig.Name)
	if err != nil {
		return nil, err
	}

	if _, err = dbconn.DB.Exec(`CREATE TABLE IF NOT EXISTS todo(id INT PRIMARY KEY AUTO_INCREMENT, name TEXT, content TEXT, time DATETIME)`); err != nil {
		logger.Error(err)
		return nil, err
	}
	return &Repo{dbconn.DB}, nil
}

func newDBClient(user, pass, host, dbname string) (*Repo, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	mc := mysql.Config{
		User:                 user,
		Passwd:               pass,
		Net:                  "tcp",
		Addr:                 host + ":" + "3307",
		DBName:               dbname,
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  jst,
	}

	db, err := sql.Open("mysql", mc.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &Repo{db}, nil
}
