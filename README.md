# go-phonebooks
Save your phonebooks with go as u go :) (In progress)

## For run locally

### Installing depedencies
```
go get github/gorilla/mux
go get github/jinzhu/gorm
go get github.com/go-sql-driver/mysql
go get github/dgrijalva/jwt-go
go get github/joho/godotenv
```

### First copy .env.example to .env file
```
cp .env.example .env
```
### Fill the variables with your environment
```
db_name = go_phonebooks
db_pass =
db_user = root
db_type = mysql
db_host = localhost
db_port = 3306
jwt_token =
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