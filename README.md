# GoBnB

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

> To run tests, `cd` from the project root to the directory of each package and run the commands:


|  Package |  Directory |
|---|---|
|  main |  /cmd/web |
|  handlers | /internal/handlers  |
|  render |  /internal/render |
|  forms |  /internal/forms |

```
go test -v
go test -cover
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

### Run Application

```
cd gobnb
go mod download
go run cmd/web/main.go cmd/web/middlewares.go cmd/web/routes.go
```

### Tech Used

- Bootstrap
- JavaScript
- Golang (using the in-built net/http package for handling requests, chi for routing, and html/template for serving the pages)
- SQL

### To be Deployed On

- VPS: Linode
- OS: Ubuntu 20.04
- Server: Caddy
- Database: PostgreSQL
