# OhMyGo

这是一个我借助AI一起做的一个AI应用平台的小项目，功能可能比较少，安全问题可能也没怎么实现，但记录一下。
一个面向个人/团队的 AI 应用平台：支持登录注册、会话管理、AI 聊天（SSE 流式）、RAG 文档问答（Redis 向量检索）、图片识别（ONNXRuntime）、TTS 语音合成（百度）。

## 功能概览

- 用户体系：注册/登录、JWT 鉴权
- 会话体系：新建会话、切换历史会话、拉取历史消息
- AI 聊天：普通回复 + SSE 流式回复（`[DONE]` 作为结束标记）
- RAG：上传 `.md/.txt` → 向量化 → Redis（RediSearch）索引 → 向量检索增强回答
- 图片识别：上传图片 → ONNX 推理 → 返回分类结果
- TTS：创建任务、查询任务结果（百度 TTS）

## 技术栈

- 后端：Go + Gin + Gorm
- 数据：MySQL + Redis（redis-stack/RediSearch）+ RabbitMQ
- 前端：Vue3 + Vue Router + Element Plus
- AI：OpenAI 兼容接口（也可对接 DashScope compatible-mode），RAG 使用 embedding + Redis 向量检索

## 快速开始（Docker Compose 推荐）

### 1) 准备环境变量（包含 API Key）

复制示例文件并填写（不要把真实 `.env` 提交到 GitHub）：

```powershell
cd d:\GopherAI-v2
copy .env.example .env
notepad .env
```

关键变量：

- `OPENAI_API_KEY`
- `OPENAI_MODEL_NAME`
- `OPENAI_BASE_URL`

示例（DashScope OpenAI 兼容）：

```text
OPENAI_BASE_URL=https://dashscope.aliyuncs.com/compatible-mode/v1
OPENAI_MODEL_NAME=qwen-plus
OPENAI_API_KEY=YOUR_KEY
```

### 2) 准备后端配置（TOML）

项目运行依赖 `config/config.toml`。Docker 部署使用 `config/config.docker.toml` 覆盖挂载。

```powershell
copy config\config.docker.example.toml config\config.docker.toml
notepad config\config.docker.toml
```

注意：`emailConfig.authcode` 属于敏感信息，不要提交到仓库。

### 3) 启动

```powershell
docker compose up -d --build
```

访问：

- 前端：`http://localhost:8080`
- 后端：`http://localhost:9090/healthz`
- RabbitMQ 管理：`http://localhost:15672`（root / 123456）
- RedisInsight：`http://localhost:8001`

## 本地开发（不使用 Docker）

### 后端

1. 启动 MySQL/Redis/RabbitMQ（可用 docker 单独跑，也可本机安装）
2. 复制并修改配置：

```powershell
cd d:\GopherAI-v2
copy config\config.example.toml config\config.toml
notepad config\config.toml
```

3. 启动后端：

```powershell
go mod download
go run .
```

默认监听：`0.0.0.0:9090`（由 `mainConfig` 控制）。

### 前端

```powershell
cd d:\GopherAI-v2\vue-frontend
npm install
npm run serve
```

开发代理在 [vue.config.js](file:///d:/GopherAI-v2/vue-frontend/vue.config.js) 中配置：`/api -> http://localhost:9090/api/v1`

## 配置说明

### `.env`（模型相关）

由 [common/aihelper/model.go](file:///d:/GopherAI-v2/common/aihelper/model.go) 读取：

- `OPENAI_API_KEY`：模型服务 Key
- `OPENAI_MODEL_NAME`：模型名（chat model）
- `OPENAI_BASE_URL`：OpenAI 兼容 base url（不要带反引号/引号/多余空格）

### `config/config.toml`（服务端基础配置）

由 [config/config.go](file:///d:/GopherAI-v2/config/config.go#L89-L105) 读取。

重点项：

- `mysqlConfig`：DB 连接信息
- `redisConfig`：Redis 连接信息（RAG 向量索引依赖 RediSearch/FT.*）
- `rabbitmqConfig`：RabbitMQ 连接信息
- `jwtConfig.expire_duration`：JWT 有效期（小时）
- `ragModelConfig`：RAG embedding/chat 模型与向量维度
- `voiceServiceConfig`：百度 TTS 的 `apiKey/secretKey`

## 常见问题

### 1) `http://localhost:9090/` 返回 404

这是正常的：业务接口统一挂在 `/api/v1/...` 下（见 [router/router.go](file:///d:/GopherAI-v2/router/router.go)）。探活用：`GET /healthz`。

### 2) `TLS handshake timeout`（调用大模型失败）

常见原因：

- `.env` 里的 `OPENAI_BASE_URL` 带了引号/反引号/空格导致 URL 非法
- Docker/本机网络无法访问目标域名 `:443`（公司代理/防火墙/DNS/IPv6）

### 3) RAG 是否会持久化？

- 原始文件：保存到 `uploads/<username>/`（在容器内建议挂载 volume）
- 向量索引：保存在 Redis（建议启用持久化/挂载 volume），目前“不会自动从文件重建索引”，重新上传可重建

### 4) 图片识别如何启用？

当前图片识别使用本地 ONNX 模型文件路径（Linux 路径）：

- 模型：`/root/models/mobilenetv2/mobilenetv2-7.onnx`
- 标签：`/root/imagenet_classes.txt`

见 [service/image.go](file:///d:/GopherAI-v2/service/image/image.go#L13-L15)。如需在 Docker 使用，建议将模型文件挂载到上述路径。

## 目录结构（简版）

```text
.
├── main.go                 # 后端入口
├── router/                 # 路由与分组（/api/v1）
├── controller/             # HTTP Handler
├── service/                # 业务逻辑编排
├── dao/                    # 数据访问
├── common/                 # MySQL/Redis/RabbitMQ/RAG/TTS/AI 等能力
├── model/                  # Gorm 模型与 DTO
└── vue-frontend/           # Vue3 前端
```

