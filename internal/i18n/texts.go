package i18n

// 英文文本
var enTexts = map[string]string{
	"cmd.root.short": "A modern scaffolding tool for Gin framework",
	"cmd.root.long": `Rgin is a modern scaffolding tool for the Gin framework that helps you:
  - Create new Gin projects with best practices
  - Add components (controllers, models, services) to your project
  - Include optional features like database support
  - Generate example code to help you get started

To get started, run:
  rgin init myapp

For more examples and documentation, visit:
  https://github.com/rosyrain/rgin

[Note: Use --lang zh for Chinese language support]`,

	"cmd.init.short": "Initialize a new Gin project",
	"cmd.init.long": `Initialize a new Gin project with the recommended project structure and basic setup.

Usage Examples:
  1. Create basic project:
     rgin init myapp

  2. With SQLite database:
     rgin init myapp --with-db sqlite

  3. Include examples:
     rgin init myapp --with-examples

  4. Full example:
     rgin init myapp --with-db sqlite --with-examples

Project Structure:
  myapp/
  ├── cmd/          - Command line entry
  ├── internal/     - Internal code
  │   ├── controller/ - Controllers
  │   ├── model/    - Data models
  │   └── service/  - Business services
  ├── pkg/         - Public packages
  ├── api/         - API documentation
  └── config/      - Configuration files

Notes:
  - Project name should be a valid Go package name
  - SQLite is used by default (lightweight, no configuration needed)
  - Use --with-examples to generate annotated example code`,

	"cmd.add.short": "Add a component to the project",
	"cmd.add.long": `Add a new component to your Gin project.

Available Components:
  1. controller - Controller
     For handling HTTP requests and responses
     Example: rgin add controller user

  2. model - Data Model
     For defining data structures and database interactions
     Example: rgin add model product

  3. service - Business Service
     For implementing business logic
     Example: rgin add service auth

Usage Examples:
  1. Add user controller:
     rgin add controller user
     # Generates internal/controller/user.go
     # Includes: Basic CRUD operations

  2. Add product model:
     rgin add model product
     # Generates internal/model/product.go
     # Includes: Struct definition and basic database operations

  3. Add auth service:
     rgin add service auth
     # Generates internal/service/auth.go
     # Includes: Basic authentication logic framework

Notes:
  - Make sure to run this command in the project root
  - Component name should be a valid Go identifier
  - Existing components won't be overwritten (manual deletion required)`,
}

// 中文文本
var zhTexts = map[string]string{
	"cmd.root.short": "现代化的 Gin 框架脚手架工具",
	"cmd.root.long": `Rgin 是一个现代化的 Gin 框架脚手架工具，可以帮助您：
  - 使用最佳实践创建新的 Gin 项目
  - 添加组件（控制器、模型、服务）到项目中
  - 包含可选功能，如数据库支持
  - 生成示例代码帮助您快速上手

开始使用：
  rgin init myapp

更多示例和文档：
  https://github.com/rosyrain/rgin

[注意：使用 --lang en 切换为英文]`,

	"cmd.init.short": "初始化一个新的 Gin 项目",
	"cmd.init.long": `初始化一个新的 Gin 项目，包含推荐的项目结构和基本设置。

使用示例：
  1. 创建基础项目：
     rgin init myapp

  2. 使用 SQLite 数据库：
     rgin init myapp --with-db sqlite

  3. 包含示例代码：
     rgin init myapp --with-examples

  4. 完整示例：
     rgin init myapp --with-db sqlite --with-examples

项目结构：
  myapp/
  ├── cmd/          - 命令行入口
  ├── internal/     - 内部代码
  │   ├── controller/ - 控制器
  │   ├── model/    - 数据模型
  │   └── service/  - 业务服务
  ├── pkg/         - 公共包
  ├── api/         - API 文档
  └── config/      - 配置文件

注意事项：
  - 项目名称应该是有效的 Go 包名
  - 默认使用 SQLite 数据库（轻量级，无需配置）
  - 使用 --with-examples 可以生成带有注释的示例代码`,

	"cmd.add.short": "向项目添加新组件",
	"cmd.add.long": `向现有的 Gin 项目添加新组件。

可用组件类型：
  1. controller - 控制器
     用于处理 HTTP 请求和响应
     示例：rgin add controller user

  2. model - 数据模型
     用于定义数据结构和数据库交互
     示例：rgin add model product

  3. service - 业务服务
     用于实现业务逻辑
     示例：rgin add service auth

使用示例：
  1. 添加用户控制器：
     rgin add controller user
     # 生成 internal/controller/user.go
     # 包含：基本的 CRUD 操作

  2. 添加产品模型：
     rgin add model product
     # 生成 internal/model/product.go
     # 包含：结构体定义和基本的数据库操作

  3. 添加认证服务：
     rgin add service auth
     # 生成 internal/service/auth.go
     # 包含：基本的认证逻辑框架

注意事项：
  - 确保在项目根目录下执行此命令
  - 组件名称应该是有效的 Go 标识符
  - 已存在的组件不会被覆盖（需要手动删除）`,
} 