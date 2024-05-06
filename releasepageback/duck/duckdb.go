package duck

//
//import (
//	"context"
//	"database/sql"
//	"database/sql/driver"
//	"fmt"
//	//"github.com/marcboeker/go-duckdb"
//	"log"
//)
//
///*
//dsn   format：
//
//execute init sql to duckdb
//*/
//func Bootqueery(execer driver.ExecerContext) error {
//	bootQueries := []string{
//		"INSTALL 'json'",
//		"LOAD 'json'",
//		"create table if not  exists releaseinfo (" +
//			"id int ," +
//			"name varchar ," +
//			"path varchar ," +
//			"type  varchar ," +
//			"version varchar ," +
//			"md5path varchar ," +
//			"md5name varchar ," +
//			"comment varchar , )",
//		"create table if not  exists buildimageconninfo (" +
//			"id int," +
//			"envname varchar ," +
//			"type varchar ," +
//			"account  varchar ," +
//			"passwd varchar )",
//		"create table if not  exists commentfiles (" +
//			"id  int ," +
//			"name varchar ," +
//			"path varchar[80] )",
//	}
//	for _, query := range bootQueries {
//		_, err := execer.ExecContext(context.Background(), query, nil)
//		if err != nil {
//			fmt.Println("execute sql faild", err)
//		}
//	}
//	return nil
//}
//
//// create connect
//func ConnDB(dsn string, bootqueery func(driver.ExecerContext) error) *sql.DB {
//
//	connector, err := duckdb.NewConnector(dsn, bootqueery)
//	if err != nil {
//		fmt.Println("create connector faild", err)
//	}
//	db := sql.OpenDB(connector)
//	return db
//}
//
//func Query(db *sql.DB, sql string, args ...any) (rows *sql.Rows, err error) {
//	defer db.Close()
//
//	rows, err = db.Query(sql)
//	if err != nil {
//		log.Fatal(err)
//		return nil, err
//	}
//	return rows, nil
//}
