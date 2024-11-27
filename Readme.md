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
