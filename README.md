# Mini project Pro Class Golang Harisenin.com

## Requirements
For running this project please install go. Go:
```
go version
go1.21.6 darwin/arm64
```

## The software you need
Recommended software to install. Example:
```
Visual Studio Code is editor code
pgAdmin 4 is management tool for PostgreSQL databases
```

## Getting Started
Init folder name:
```
go mod init favorite-book
```

## Setup port
```
Open main.go
edit value port to open in browser as you wish, for example:
port = "8080"
```

## Setup database:
```
Open infrastructure.go
setup dsn value according to your database settings, for example:
dsn = "postgres://postgres:password@localhost:5432/db_book"
postgres is your username
password is your password
localhost is your ip address
5432 is port your database
db_book is your schema
```

## Install library:
```
install several libraries needed for the project, for example:
go get github.com/gofiber/fiber/v2
```

## Run the app:
```
open terminal
go run main.go
open browser
go to url http://localhost:8080/
```

## Steps :
- Clone this repo
- Install go
- Open project with Visual Studio Code
- Setup your database PostgreSQL with pgAdmin 4 
- Follow the steps init
- Edit main.go
- Edit infrastructure.go
- Install library
- Running project on development 

## Reference :
```
https://bit.ly/ProgramDetailsGOLANG

https://dandelion-authority-347.notion.site/Advance-8ddb2fdf00be4a5fa4abda27e8b32fc8
```