# Espresso

    .
    ├── api
    │   ├── main.go
    │   └── user.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
    │   ├── init.go
    │   └── users.go
    ├── README.md
    ├── serialization
    │   └── base.go
    ├── server
    │   └── router.go
    ├── service
    │   ├── user_login_service.go
    │   ├── user_register_service.go
    │   └── user_register_service_test.go
    └── test
        └── test.go

需要討論的問題
1. Get 一個行程的參數要如何制定
2. 兩個行程同時段問題
3. 統一response那部份有些問題, 大概知道要如何解決


