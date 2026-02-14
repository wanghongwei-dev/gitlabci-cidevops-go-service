# GitLab CI Go 服务示例

这是一个使用Go语言编写的简单Web服务示例，展示了如何在GitLab CI/CD环境中进行持续集成和部署。

## 功能特性

- 提供基础的HTTP服务
- 健康检查端点
- Docker容器化支持
- Kubernetes部署配置

## 技术栈

- **语言**: Go
- **Web框架**: 标准库net/http
- **容器化**: Docker
- **部署**: Kubernetes

## API接口

### 主页
```
GET /
响应: "欢迎来到GitLab CI Go服务演示!"
```

### 健康检查
```
GET /health
响应: "OK"
状态码: 200
```

## 快速开始

### 本地运行

1. 克隆项目
```bash
git clone <repository-url>
cd gitlabci-cidevops-go-service
```

2. 运行服务
```bash
go run main.go
```

3. 访问服务
```bash
curl http://localhost:8082
curl http://localhost:8082/health
```

### Docker运行

1. 构建镜像
```bash
docker build -t go-service .
```

2. 运行容器
```bash
docker run -p 8082:8082 go-service
```

## 项目结构

```
gitlabci-cidevops-go-service/
├── manifest_tpl/           # Kubernetes部署模板
│   ├── deployment.yaml     # 基础部署配置
│   ├── deployment-stag.yaml # 预发布环境配置
│   ├── deployment-uat.yaml  # UAT环境配置
│   └── deployment-prod.yaml # 生产环境配置
├── Dockerfile              # Docker镜像构建文件
├── main.go                 # 主程序文件
├── main_test.go            # 测试文件
└── README.md               # 项目说明文档
```

## CI/CD配置

项目包含完整的GitLab CI/CD配置，支持：
- 代码构建和测试
- Docker镜像构建
- 自动化部署到不同环境

## 部署配置

### Kubernetes部署

使用提供的YAML模板进行Kubernetes部署：

```bash
# 部署到开发环境
kubectl apply -f manifest_tpl/deployment.yaml

# 部署到预发布环境
kubectl apply -f manifest_tpl/deployment-stag.yaml

# 部署到生产环境
kubectl apply -f manifest_tpl/deployment-prod.yaml
```

## 环境要求

- Go 1.16+
- Docker (可选)
- Kubernetes集群 (可选)

## 测试

运行单元测试：
```bash
go test -v
```

## 许可证

MIT License