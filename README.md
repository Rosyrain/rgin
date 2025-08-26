# rgin命令行工具--一键生成gin框架的开发脚手架

<hr>

**前言**

​		在使用 Gin 框架进行开发学习的过程中，脚手架的搭建是一个常见且必要的步骤。尽管可以提前准备一个简单的脚手架模板并通过复制粘贴来快速启动项目，但这种方式往往需要对模块名称（如 `mod name`）等关键命名进行手动修改。然而，由于项目名称可能与代码中的某些变量名重复，导致无法简单地通过全局搜索替换（如 `Shift` + `Ctrl` + `F`）来一次性完成所有修改，反而需要逐一甄别和调整，增加了不必要的繁琐操作。

​		鉴于 Gin 框架本身以简洁著称，并未提供一键生成项目模板的命令行工具，因此，基于上述痛点，我开发了一个可以根据个人开发习惯生成通用模板的命令行工具。该工具不仅能够快速生成符合 Gin 框架开发规范的项目结构，还能灵活适配不同的命名需求，在一定程度上提高了开发效率。

​		值得一提的是，本文所介绍的方法并不仅限于 Gin 框架的脚手架生成。你可以基于本文的思路和方法，开发适用于其他常用框架或脚手架的个性化命令行工具，从而为日常开发工作提供更多便利。通过这种方式，开发者可以根据自身需求定制专属的模板生成工具，进一步优化开发流程，提升工作效率。

<hr>

### 0、快速使用

项目地址：[Rosyrain/rgin](https://github.com/Rosyrain/rgin/tree/main)

如果有帮助到你，希望可以得到你的**star**ヾ(≧▽≦*)o。

```shell
go install github.com/rosyrain/rgin@0.1.0
rgin -l zh # 全局切换为中文
rgin init myapp --with-example # --with-example 生成blubell示例代码
cd myapp
sqlite3 ./data/app.db < models/create_table.sql #创建sqlite数据库并建表
go run main.go #启动基础项目
```

生成的脚手架结构参考了七米老师的`bluebell`项目，支持 MySQL/Redis/SQLite 等多种数据库，内置 JWT 认证、限流等功能。


```
├── conf/           # 配置文件
├── controller/     # 控制器/处理函数
├── dao/            # 数据访问层(根据需求可选)
│   ├── mysql/
│   ├── redis/
|   └── sqlite/
├── example/        # 示例项目（可选,通过--with-example生成）
├── logger/         # 日志组件
├── logic/          # 业务逻辑
├── middlewares/    # 中间件（JWT、限流等）
├── models/         # 数据模型
├── pkg/            # 第三方库封装（jwt、snowflask等）
│   ├── jwt/
│   └── snowflask/
├── router/         # 路由注册
├── settings/       # 配置加载
├── main.go         # 启动入口
├── Dockerfile      # Docker 部署
└── wait-for.sh     # Docker 启动检测脚本
```


### 1、rgin 工具自身结构

```
rgin/
├── cmd/              # 命令行入口及子命令
│   ├── help.go
│   ├── init.go
│   └── root.go
├── conf/             # 工具自身配置
├── internal/         # 内部实现
│   ├── generator/    # 代码生成逻辑
│   ├── i18n/         # 国际化支持
│   ├── project/      # 项目结构定义
│   └── template/     # 模板加载与渲染
│       ├── templates/  # 各类模板文件(放在此处是为了与embed进行适配)
            ├── main.go.tmpl
            ...
│       └── template.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```


---
**更多说明与实现参考**

> 下面文章实现内容属于初次搭建时的思路，目前项目结构与初次搭建有较大差距，具体可借助AI理解分析

- [CSDN 博客：rgin命令行工具–一键生成gin框架的开发脚手架](https://blog.csdn.net/meng7000/article/details/145829359)
- [稀土掘金：rgin命令行工具–一键生成gin框架的开发脚手架](https://juejin.cn/spost/7474781404163522611)

**相关资料**

- [cobra：Go 命令行/应用程序生成工具](https://www.cnblogs.com/jiujuan/p/15487918.html)
- [Go 标准库 template 文档](https://www.topgoer.com/常用标准库/template.html)
- [go:embed 用法详解](https://blog.axiaoxin.com/post/go-embed/)

