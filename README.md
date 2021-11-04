# go-demo

1、下载  

2、配置文件  

```sh
cp .env.example .env
```

然后修改 `.env` 文件中的配置项。  

3、数据准备  

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
COMMIT;
```

4、启动  

```sh
go run main.go
```
