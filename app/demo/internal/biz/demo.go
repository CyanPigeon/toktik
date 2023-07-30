package biz

// TODO 此处定义错误返回值。如果无错误返回值可忽略。

// Entity 实体结构体，与物理表一一对应。
// TODO 'Entity'需要替换为对应的实体表名
// TODO 如果有多个实体，请拆分为多个文件，一个文件一个实体
type Entity struct {
	// TODO 此处填写实体表字段名
}

// EntityRepo 实体的读/写/查询操作接口。
// TODO 'EntityRepo'需要替换为对应的实体表名，后缀统一为Repo
type EntityRepo interface {
	// TODO 此处填写对实体的操作，即增删改查
}
