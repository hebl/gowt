Go-Web-Template
===============

Go 语言开发Web程序模板，使用数据库为PostgreSQL

在构建0.1后，考虑到一些扩展以及方便，采用`xorm`成为数据库访问的模型。

## 目录结构

```
- app 程序目录结构
	- controller 控制器
	- model 模型
	- view 渲染
	- lib 工具模块（mail，log等）
	- router.go 路由
	- app.go Web服务启动程序
- etc
	- conf.ini 网站配置文件
- assets 静态文件
	- css
	- js
	- fonts
	- img
- log 日志
- main.go 主程序
- Makefile 程序编译文件
- dist 用户生成发布版本（make dist）
```