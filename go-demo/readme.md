## 项目介绍
这是一个基于Go语言的用户管理系统，包含用户注册、登录、查询等功能。

### 功能列表
- 用户注册
- 用户登录
- 查询用户信息 
- 更新用户信息
- 删除用户

## 快速开始
1. 克隆项目到本地
```bash
git clone https://github.com/yourusername/user-management.git
```

2. 安装依赖
```bash
go mod tidy
```
3. 配置数据库连接
在`main.go`中，配置数据库连接信息。
```go
const (
	DBHost     = "localhost"
	DBPort     = "3306"
	DBUser     = "root"
	DBPassword = "123456"
	DBName     = "user_management"
)
```
4. 运行项目
```bash
go run main.go
```

