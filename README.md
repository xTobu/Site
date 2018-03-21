
# Site（暫停更新

> 基於 Vue.js + Golang 的後台管理系統

**Front-end**
 - vue
 - vuex
 - vue-router
 - pug
 - sass-loader
 - axios
 - qs（qs.stringify for axios） 
 - ElementUI
 - url-loader（引用ElementUI樣式時需要編譯字型檔）
 
**Back-end**
 - Gin
 - contrib/static
 - go-mssqldb
 - go-sql-driver/mysql
 
##  Setup

**Front-end**
```bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

**Back-end**
```bash
# serve at localhost:8000 with static "/"
cd ./backend && go run ./main.go
```