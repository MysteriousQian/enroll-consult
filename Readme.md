# 版本 V1.0.0

# enroll-consult

## 目录
- [项目说明](#项目说明)
- [安装依赖](#安装依赖)

## 项目说明
招生咨询系统
#### 目录结构
```
.
├── cmd/                        
│   ├── api/   API服务入口,主服务向外提供接口
│   │   ├── config.yaml         配置文件
│   │   ├── locales             国际化语言文件
│   │   ├── log                 日志文件
│   │   ├── static/             静态文件图片资源等(即将废弃⚠️)
│   │   ├── main_test.go        测试入口文件
│   │   └── main.go             服务入口文件
├── docker/                  docker镜像打包配置文件(DockerFile)
├── internal/                服务内部逻辑处理
│   ├── api/                    外部接口
│   ├── config/                 应用配置
│   ├── db/                     数据库
│   ├── handler/                中间件 
│   ├── scheduler/              调度任务
│   └── services/               业务逻辑
├── pkg/                     公共库   
├── test/                    单元测试 
├── .env                     环境变量
├── tob_group.yml            项目DockerCompose启动配置文件             
├── go.mod
├── go.sum           
├── vendor/                  第三方依赖                
└── README.md                项目说明
```
## 安装依赖
```
go mod tidy

protobuf 格式生成
https://github.com/protocolbuffers/protobuf/releases 安装对应版本 添加到环境变量使用
vi ~/.zshrc or ~/.bash_profile
source ~/.bash_profile
source ~/.zshrc
下载项目 google.golang.org/protobuf/cmd/protoc-gen-go
进入 cmd/protoc-gen-go 目录下运行 go build
生成可执行二进制文件 丢到GOPATH/bin目录下
将 GOPATH/bin 添加到你的 PATH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
protoc-gen-go --version 测试是否安装完成

protoc --version 测试是否正确安装
protoc --go_out=. device.proto #生成 Protocol协议对应的go消息结构文件
```
## Docker部署服务
```
docker compose -f tob_group.yml up -d 启动所有容器服务
```
 