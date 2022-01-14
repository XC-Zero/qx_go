## 清墟

[![GitHub](https://img.shields.io/github/license/XC-Zero/golang-learning-bible)](https://github.com/XC-Zero/qx_go/blob/dev/LICENSE)

清墟是一个网页题库系统，基于 Go 语言开发 ，作为练手项目，内部冗余技术栈可能较多

`Just look look ok !`

- 服务:
    - web:
        - gin
        - consul
    - rpc
        - twrip

- 消息队列：
    - kafka

- 数据库：
    - 已引入：
        - mysql
        - redis
        - minio (考虑换成seaweed)
        - mongoDB
        - elasticsearch
        - influxDB (考虑是否替换为questDB)

    - 考虑引入
        - TiDB
        - Drios
        - clickhose

    - 数据库连接
        - gorm
        - 其他都是官方连接库
            