package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {

	// 用户
	//genModel("tiktok_user", "user", "User")
	// video
	//genModel("video", "video", "Video")
	// 评论
	//genModel("comment", "comment", "Comment")
	// 关注
	genModel("follow", "follow", "Follow")

	//genModel("message", "message", "Message")

}

func genModel(tableName string, outFileName string, modelName string) {
	// 设置生成模式
	g := gen.NewGenerator(gen.Config{
		// dal的位置
		OutPath:      "./internal/data/model/dao/" + outFileName,    //dao方法 查询类文件路径
		ModelPkgPath: "./internal/data/model/",                      // 模型
		OutFile:      outFileName + "Dao.go",                        // 输出dao的名称
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	//gormdb, _ := gorm.Open(mysql.Open("root:qq1448265203@(127.0.0.1:3306)/testDb?charset=utf8mb4&parseTime=True&loc=Local"))

	dsn := "host=localhost user=postgres password=Ymriri dbname=testDb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	gormdb, _ := gorm.Open(postgres.Open(dsn))

	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(biz.Student{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, biz.Student{})
	//g.ApplyBasic(g.GenerateModelAs("students", "Student"))
	// 为所有表生自接口
	g.ApplyBasic(
		g.GenerateModelAs(tableName, modelName),
		// Generate structs from all tables of current database
		//g.GenerateAllTable()...,
	)
	// 使用如下代码默认在model下面生成，仅仅生成表结构
	//g.GenerateModel("students", gen.WithMethod(biz.Student{}))
	// 表换个名称
	//g.GenerateModelAs("students", "Student")
	// Generate the code
	g.Execute()
}
