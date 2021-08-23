# GoBnB (wip)

> A fully featured bookings and reservation system for a Bed & Breakfast.

Visitors to the site will be able to search for accommodations by date and make an online reservation, and the site owner will be able to manage reservations from a secure back end.

### Key Functionalities

- Showcase Properties
- Allow bookings for one or more nights
- Check availability
- Book property
- Notifications to guests, and property owner

## Getting Started

### Clone Application

```
git clone https://github.com/Fakorede/gobnb.git
```

### Run Tests

> To run tests on a package level, `cd` from the project root to the directory of each package and run the commands:


|  **Package** |  **Directory** |
|---|---|
|  main |  /cmd/web |
|  handlers | /internal/handlers  |
|  render |  /internal/render |
|  forms |  /internal/forms |

```
$ go test -v
$ go test -cover
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

> To run all tests from the project root

```
$ go test -v ./...
$ go test -cover ./...
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out ./...
```

### Database Setup

```
cp database.yml.example database.yml
```

create db in your database client and add `database`, `user`, and `password` to database.yml file
```
database: gobnb
user:
password:
```

### Run Migrations

> To run migrations, you must have Soda nstalled

```
$ go get github.com/gobuffalo/pop/...

// add path to .zprofile
$ cd ~
$ nano .zprofile
export PATH="$HOME/go/bin:$PATH"

// confirm path has been added
$ which soda

// run migrations
$ soda migrate
```

### Run Application

```
$ cd gobnb
$ go mod download

$ go run cmd/web/main.go cmd/web/middlewares.go cmd/web/routes.go

alternatively,

// On Mac
$ chmod +x run.sh
$ ./run.sh

// On Windows
$ run.bat
```

### Tech Used

- Bootstrap
- JavaScript
- Golang (using the in-built net/http package for handling requests, chi for routing, html/template for serving the pages, soda for database migrations)
- SQL

### Deployed On

- VPS: Linode
- OS: Ubuntu 20.04
- Server: Caddy
- Database: PostgreSQL
