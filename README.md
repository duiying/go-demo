# go-demo

### 介绍

本项目为 Go 语言的练手项目，实现了一个 `user` 模块 CRUD 的 API，实践了如下知识点：    

- 路由
- 中间件
- MySQL
- Redis
- 日志
- 代码分层

代码分层：  

- controller：控制器层，负责校验参数
- logic：逻辑层，负责处理实际业务逻辑
- dao：数据库层，负责执行 SQL
- model：模型层，负责封装结构体以及部分字段 Map

### 如何运行

1、下载  

2、配置文件  

```sh
cp .env.example .env
```

然后修改 `.env` 文件中的配置项（包括日志目录、数据库配置等）。  

3、准备测试数据  

```SQL
CREATE DATABASE passport DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

use passport;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                        `name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
                        `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
                        `root` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'ROOT 用户 {0：否；1：是；}',
                        `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                        `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        PRIMARY KEY (`id`),
                        KEY `idx_name` (`name`),
                        KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

INSERT INTO `user` VALUES (1, 'duiying', 'duiying@gmail.com', 1, '2021-11-04 16:53:33', '2021-11-04 16:53:33');
```

4、启动  

```sh
go mod tidy
go run main.go
```

5、访问 API  

[127.0.0.1:9551](http://127.0.0.1:9551)

### TODO

- 防 SQL 注入
