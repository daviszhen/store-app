# store-app

移动端商店模板，支持多业务类型切换（便利店 / 咖啡店）。

## 技术栈

| 层 | 技术 |
|---|---|
| 前端 | React + TypeScript + Vite |
| 后端 | Go + Gin |
| 数据库 | MatrixOne |

## 项目结构

```
store-app/
├── frontend/          # React 前端
│   ├── src/
│   │   ├── api/       # API 请求封装
│   │   ├── context/   # React Context（购物车状态）
│   │   ├── pages/     # 页面组件（Home / Category / Product / Cart）
│   │   └── types.ts   # 类型定义
│   └── .env           # 前端环境配置
├── server/            # Go 后端
│   ├── cmd/           # 入口 main.go
│   └── internal/
│       ├── config/    # 配置加载（含业务类型种子数据）
│       ├── database/  # 数据库连接与初始化
│       ├── handler/   # API 处理器
│       └── model/     # 数据模型
└── sql/
    └── init.sql       # 数据库初始化脚本
```

## 快速启动

### 1. 初始化数据库

```bash
# 在 MatrixOne 中执行
mysql -h 127.0.0.1 -P 6001 -u root -p111 < sql/init.sql
```

### 2. 启动后端

```bash
cd server
# 可选：设置环境变量
export SERVER_PORT=8080          # 默认 8080
export BUSINESS_TYPE=grocery     # grocery | coffee
./server
```

### 3. 启动前端

```bash
cd frontend
npm install
npm run dev
```

## 配置说明

### 前端 `.env`

```env
# 业务类型：grocery（便利店）/ coffee（咖啡店）
VITE_BUSINESS_TYPE=grocery

# 后端 API 地址
VITE_API_BASE_URL=http://localhost:8080/api
```

### 后端环境变量

| 变量 | 默认值 | 说明 |
|---|---|---|
| `SERVER_PORT` | `8080` | 服务端口 |
| `BUSINESS_TYPE` | `grocery` | 业务类型 |
| `DB_HOST` | `127.0.0.1` | 数据库地址 |
| `DB_PORT` | `6001` | 数据库端口 |
| `DB_USER` | `root` | 数据库用户 |
| `DB_PASSWORD` | `111` | 数据库密码 |
| `DB_NAME` | `store_app` | 数据库名 |

## 业务类型

通过 `VITE_BUSINESS_TYPE` / `BUSINESS_TYPE` 切换：

| 类型 | 说明 |
|---|---|
| `grocery` | 便利店：水果、蔬菜、肉禽、酒水等 7 分类 42 商品 |
| `coffee` | 咖啡店："慢时光咖啡"，4 分类 20 商品 |

新增业务类型：在 `server/internal/config/business/` 创建 Loader 并 `init()` 注册即可。

## API 端点

| 方法 | 路径 | 说明 |
|---|---|---|
| GET/PUT | `/api/store` | 店铺信息 |
| CRUD | `/api/categories` | 商品分类 |
| CRUD | `/api/products` | 商品（支持 `?category_id=` 筛选） |
| CRUD | `/api/cart` | 购物车（按 `user_id`） |
| POST/GET | `/api/orders` | 订单 |
