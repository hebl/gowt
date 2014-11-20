## Go Web 模板

### 思路

使用Go语言本身的

#### 工具

- http Server
  * net/http
- 路由
  * [gorilla/mux](https://github.com/gorilla/mux)
- 数据库操作，不使用`ORM`， 使用原生`SQL`
  * [gopkg.in/pg.v3](https://github.com/go-pg/pg)
- Session， Session存储在Cookie或者本地文件中
  * [gorilla/sessions](https://github.com/gorilla/sessions)

#### 文件结构

- `main.go` 主程序
- `core` 程序目录
  * `global.go` 全局变量、常量
  * `db.go` 数据库模型与操作
  * `controller` 控制器、过滤器等
  * `view.go` 输出渲染
  * `utils.go` 工具类
- `templates` HTML模板目录
- `assets` 静态目录，javascript、CSS、图片等 

#### 数据库结构

一个用户表，一个内容表

```sql
CREATE TABLE users(
	Id integer,
	name varchar(30),
	email varchar(30),
	dtime timestamp
);

CREATE TABLE post(
	Id integer,
	title varchar(100),
	content text,
	created timestamp
);
```