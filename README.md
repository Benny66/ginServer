# gingen

自动生成框架内容可前往开源项目[gingen](https://github.com/Benny66/gingen.git)按如下命令执行

```
mkdir testServer
cd testServer
./gingen init --mod testServer
```

# 基于go-gin框架的web服务框架
项目地址：[https://github.com/Benny66/ginServer](https://github.com/Benny66/ginServer)

## app 
项目工程主要代码文件夹目录，包括api层、model模型数据层、service逻辑层；

- api层请求入口处理，参数校验，数据返回
- model模型数据层是数据交互层，常见的数据库操作方法，数据聚合方法
- 逻辑层是项目核心业务逻辑的处理层；
### api层
api层连同schema模块接受处理请求参数，做数据校验、清洗返回等
```
type UserInterface interface {
	Login(context *gin.Context)
	Refresh(context *gin.Context)
	Logout(context *gin.Context)
	UpdatePassword(context *gin.Context)
}

var UserApi UserInterface = &userApi{}

type userApi struct{}
```

## model【模型数据层】

model就是对数据库表名和表内字段进行模型定义的模块。
- ModelTime定义自动转换存储和查询时间格式
- 可定义模型对应的表名称和表字段

dao 是基于gorm对数据进行增删查改的模块，通过inteface接口暴露调用接口
- Create和update 对数据的创建和修改操作均需要开启事务，在逻辑层进行控制开启、回滚和提交。
- 常见封装的方法包括增删查改、分页查询（Paginate）、查询全部（FindAll）、按条件查询（WhereQuery）、关联查询（Joins）、预加载（Preloads）等等方法

```
func (dao *userDao) Create(tx *gorm.DB, data *model.User) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}
func (dao *userDao) WhereQuery(query interface{}, args ...interface{}) *userDao {
	return &userDao{
		dao.gm.Where(query, args...),
	}
}
func (dao *userDao) Joins(query string, args ...interface{}) *userDao {
	return &userDao{
		dao.gm.Joins(query, args),
	}
}
```
## service【逻辑层】
service【逻辑层】是主要的代码层，开发人员基本上在这个模块上进行开发和修复bug，实现各自项目的逻辑，是最核心的内容
- api模块是对接路由的方法入口，基本上一个业务模块对应一个文件，例如登录模块的各接口可以命名为user.go
- define模块用于定义数据结构类型的目录，不同的业务类型和逻辑，需要定义不同的请求参数和返回参数结构，例如定义type UserLoginApiReq struct来接收登录接口的参数类型
- service模块，顾名思义就是业务逻辑服务的处理模块，其中包括对请求数据参数的校验、业务逻辑处理数据，调用数据层进行保存数据库
- router.go对接路由,统一定义

## config【系统配置】
系统的配置模块，config.go，包括服务信息、数据库信息、日志配置信息、ws配置信息等等

## db【数据库】
数据库文件夹模块，目前使用的mysql、redis数据存储

## migrations【数据迁移】
数据迁移文件模块，项目初始化或升级的时候可进行数据库的数据库迁移脚本命令；

需要在根目录添加install.lock才可进行数据库迁移执行

## public【公共模块】
项目的公共模块，包括：image、html、css、js等文件

## routers【路由】
启动web服务时，初始化gin的路由模块，启动服务

## runtime
runtime模块，保存日志logs、缓存cache等文件

## utils
项目框架需要的工具包，包括：自我封装的库以及调用第三方封装的库

