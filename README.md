# go-phonebooks
Save your phonebooks with go as u go :) (In progress)

## For run locally

### Installing depedencies
```
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
go get github.com/dgrijalva/jwt-go
go get github.com/joho/godotenv
```

### First copy .env.example to .env file
```
cp .env.example .env
```
### Fill the variables with your environment
```
DB_NAME = go_phonebooks
DB_PASSWORD =
DB_USERNAME = root
DB_DIALECT = mysql
DB_HOST = localhost
DB_PORT = 3306
DB_CHARSET = utf8
JWT_TOKEN =
PORT = 5000
```

### Run it...

```
go run main.go
// your app will be served at http://localhost:${PORT}
```

## For build

```
go build -o build/main main.go && build/main
```