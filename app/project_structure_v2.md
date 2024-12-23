# 项目结构说明 
blackapp/
├── cmd/ # 应用程序入口
│ └── server/ # 服务器入口
│ └── main.go # 主程序
├── configs/ # 配置文件目录
│ └── dev/ # 开发环境配置
│ └── config.yaml # 配置文件
├── docs/ # 文档目录
│ └── swagger.go # Swagger文档
├── internal/ # 内部包
│ ├── api/ # API层
│ │ ├── handler/ # 处理器
│ │ ├── middleware/ # 中间件
│ │ ├── response/ # 响应封装
│ │ └── router/ # 路由配置
│ ├── domain/ # 领域层
│ │ ├── entity/ # 实体
│ │ └── repository/ # 仓储接口
│ ├── infrastructure/ # 基础设施层
│ │ └── persistence/ # 持久化实现
│ └── service/ # 服务层
│ ├── dto/ # 数据传输对象
│ └── impl/ # 服务实现
├── pkg/ # 公共包
│ ├── config/ # 配置管理
│ ├── database/ # 数据库连接
│ └── logger/ # 日志管理
├── test/ # 测试目录
│ └── integration/ # 集成测试
├── .gitignore # Git忽略文件
├── Dockerfile # Docker构建文件
├── Makefile # 构建脚本
├── README.md # 项目说明
├── docker-compose.yml # Docker编排配置
├── go.mod # Go模块文件
└── go.sum # Go依赖版本锁定文件
