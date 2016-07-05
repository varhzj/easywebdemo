package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Conf struct {
	Db DbConf `json:"db"`
}

type DbConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Name string `json:"name"`
}

const DbFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true"

var dbStore *DBStore

type DBStore struct {
	DB *sql.DB
}

func GetDBStore() *DBStore {
	return dbStore
}

func (ds *DBStore) GetDB() *sql.DB {
	return ds.DB
}

func (ds *DBStore) Close() error {
	return ds.DB.Close()
}

func init() {
	data, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		panic(err)
	}

	var conf Conf
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}
	log.Printf("DbConf:{"+DbFormat+"}", conf.Db.User, conf.Db.Pwd, conf.Db.Host, conf.Db.Port, conf.Db.Name)

	dbConnStr := fmt.Sprintf(DbFormat, conf.Db.User, conf.Db.Pwd, conf.Db.Host, conf.Db.Port, conf.Db.Name)
	db, err := sql.Open("mysql", dbConnStr)
	err = db.Ping()
	if err != nil {
		log.Fatal("Can not connect to db")
		panic(err)
	}
	dbStore = &DBStore{DB: db}

	//FuncMap["str2html"] = Str2html
}
