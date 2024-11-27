# Go E-commerce Backend

使用 Golang 開發的電商網站後端 API

## 技術棧

- Golang
- Gin (Web框架)
- PostgreSQL (主資料庫)
- GORM (ORM)
- Redis (快取、購物車)
- JWT (認證)

## 功能

- [x] 使用者認證
  - [x] 註冊
  - [x] 登入 (JWT)
- [ ] 商品管理
- [ ] 購物車
- [ ] 訂單系統

## 開始使用

### 前置需求

- Go 1.20+
- PostgreSQL
- Redis
- Docker (選用)

### 環境設定

1. 複製環境設定檔
```bash
cp config.example.go config.go
```

2. 啟動資料庫
```bash
docker run --name my-postgres -e POSTGRES_PASSWORD=your_password -p 5432:5432 -d postgres
docker run --name my-redis -p 6379:6379 -d redis
```

3. 執行
```bash
go run cmd/main.go
```

## API文檔

### 使用者相關

- POST `/register` - 註冊
- POST `/login` - 登入

## 專案結構

```
.
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── models/
│   ├── repository/
│   ├── service/
│   └── handler/
└── pkg/
    ├── database/
    └── middleware/
```

## 開發中功能

- Redis 整合
- 商品管理
- Functional Programming 實作

## License

MIT
