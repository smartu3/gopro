# gopro

## 简单的go项目实践

### 结构

#### conf 配置文件统一存放目录

#### config 专门用来处理配置和配置文件的go package

#### db.sql 执行source db.sql 创建数据库和表，新环境初始化使用

#### handler 类似于MVC架构里的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果

#### main.go 程序唯一入口

#### model 数据库相关的统一操作，包括数据库初始化和对表的增删改查

#### pkg 引用的包

#### router 路由相关处理

#### service 实际业务处理函数的位置

#### util 工具类函数存放目录

#### vendor vendor目录管理依赖包