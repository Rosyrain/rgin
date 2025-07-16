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

生成的开发脚手架参考了七米老师的bluebell项目  架构如下(需要设置相关配置mysql/redis便可启动预留的示例)：

```
├─conf			//存放配置文件
├─controller	//句柄函数
├─dao			//数据库
│  ├─mysql		//mysql
│  └─redis		//redis
|─example       //示例项目(blubell),通过init --with-example生成
├─logger		//日志加载
├─logic			//逻辑处理函数
├─middlewares	//中间件，已配置jwt认证，令牌桶限流
├─models		//模型定义
├─pkg			//第三方库调用，jwt的具体相关功能实现，snowflask算法
│  ├─jwt
│  └─snowflask
├─router		//路由设置
├─settings		//配置加载
├─main.go		//启动入口，相关配置初始化
├─DOckerfile    //docker部署
└─wait-for.sh   //docker部署时调用的验证脚本
```

### 项目结构设计(预设计)

```
rgin/
├── cmd/                  // 命令行工具入口
    ├── create.go		  // 子命令create定义
│   └── root.go			  // 根命令定义
├── internal/             // 内部实现(功能的具体实现)
│   ├── generator/        // 代码生成逻辑
│   │   └── generator.go
│   ├── project/          // 项目结构定义
│   │   └── project.go
│   └── template/         // 模板文件加载
|       ├──templates/     // 模板文件,放在此处是为了与embed进行适配
|			├──main.go.tmpl...
│       └── template.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

<hr>
**具体实现思路可参考文章：**

[rgin命令行工具–一键生成gin框架的开发脚手架-CSDN博客](https://blog.csdn.net/meng7000/article/details/145829359)

[rgin命令行工具–一键生成gin框架的开发脚手架-稀土掘金](https://juejin.cn/spost/7474781404163522611)

**参考文献**

[golang常用库包：cli命令行/应用程序生成工具-cobra使用 - 九卷 - 博客园](https://www.cnblogs.com/jiujuan/p/15487918.html)

[Template · Go语言中文文档](https://www.topgoer.com/常用标准库/template.html)

[go:embed 用法详解：如何将静态资源文件打包进二进制文件中？ - 阿小信的博客](https://blog.axiaoxin.com/post/go-embed/)

