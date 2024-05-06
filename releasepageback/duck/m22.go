package duck

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	_ "github.com/marcboeker/go-duckdb"
	"os"
	"releasepageback/model"
	"strings"
)

func main1() {
	connector, err := duckdb.NewConnector("duck/foo.db?access_mode=READ_WRITE&threads=4", func(execer driver.ExecerContext) error {
		bootQueries := []string{
			"INSTALL 'json'",
			"LOAD 'json'",
			"create table if not  exists releaseinfo (" +
				"id int ," +
				"name varchar ," +
				"path varchar ," +
				"type  varchar ," +
				"version varchar ," +
				"md5path varchar ," +
				"md5name varchar ," +
				"comment varchar , )",
			"create table if not  exists buildimageconninfo (" +
				"id int," +
				"envname varchar ," +
				"type varchar ," +
				"account  varchar ," +
				"passwd varchar )",
			"create table if not  exists commentfiles (" +
				"id  int ," +
				"name varchar ," +
				"path varchar[80] )",
		}

		for _, query := range bootQueries {
			_, err := execer.ExecContext(context.Background(), query, nil)
			if err != nil {
				fmt.Println("22222222222222222", err)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("3333333333333", err)
	}
	// 读取 JSON 文件
	jsonFile, err := os.Open("config/releasinfo.json")
	if err != nil {
		fmt.Println("无法打开 JSON 文件:", err)
		return
	}
	defer jsonFile.Close()

	// 解码 JSON 数据
	var releaseinfo []model.Release
	err = json.NewDecoder(jsonFile).Decode(&releaseinfo)
	if err != nil {
		fmt.Println("无法解码 JSON 数据:", err)
		return
	}
	db := sql.OpenDB(connector)
	// 插入数据
	for _, release := range releaseinfo {
		//commentJSON, err := json.Marshal(release.Comment)
		//if err != nil {
		//	fmt.Println("无法转换 comment 到 JSON:", err)
		//	return
		//}
		commentStr := strings.Join(release.Comment, ",") // 将字符串切片转换为逗号分隔的字符串
		_, err = db.Exec("INSERT INTO releaseinfo (id,name,path ,type ,version ,md5path,md5name,comment) VALUES (?,?,?,?,?,?,?,?)", release.Id, release.Name, release.Path, release.Type, release.Version, release.Md5Path, release.Md5Path, commentStr)
		if err != nil {
			fmt.Println("无法插入数据:", err)
			return
		}
	}
	rows, err := db.Query("select * from releaseinfo ")

	for rows.Next() {
		var release model.Release
		var commentStr string // 用于接收逗号分隔的字符串

		err := rows.Scan(&release.Id, &release.Name, &release.Path, &release.Type, &release.Version, &release.Md5Path, &release.Md5name, &commentStr)
		if err != nil {
			fmt.Println("无法扫描数据:", err)
			return
		}

		// 将逗号分隔的字符串解析为字符串切片对象
		release.Comment = strings.Split(commentStr, ",")

		// 打印 release 信息
		fmt.Printf("release: %p\n", release)
	}
	rows.Close()

}
