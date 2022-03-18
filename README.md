# Online Store
web api for online shopping store(REST API)

## Installation
- `First` make a db.env file in path `controllers/db/` and fill env variables like pattern below:
```env
USERNAME=[YOUR_USERNAME]
PASSWORD=[YOUR_PASSWORD]
HOST=[YOUR_HOST]
DB_NAME=[YOUR_DB_NAME]
```
- `Second` Create a new Docker Container with execute the command below.
```bash
$ docker run -p 5432:5432 -d \
    -e POSTGRES_PASSWORD=[YOUR_PASSWORD] \
	-e POSTGRES_USER=[YOUR_USERNAME] \
	-e POSTGRES_DB=[YOUR_DB_NAME] \
	-v pgdata:/var/lib/postgres/data \
	postgres
```
- `Third` in root of project, run this command:
```bash
$ go run main.go
```
## Tools
- Used ORM: Entgo
- Auth: JWT
- Database: Postgresql


## Schema 
```
➜  entgo-example git:(main) go run entgo.io/ent/cmd/ent describe ./controllers/ent/schema
Category:
        +-------------+--------+--------+----------+----------+---------+---------------+-----------+------------------------------+------------+
        |    Field    |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |          StructTag           | Validators |
        +-------------+--------+--------+----------+----------+---------+---------------+-----------+------------------------------+------------+
        | id          | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"          |          0 |
        | name        | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty"        |          0 |
        | description | string | false  | false    | false    | false   | false         | false     | json:"description,omitempty" |          0 |
        | thumbnail   | string | false  | false    | false    | false   | false         | false     | json:"thumbnail,omitempty"   |          0 |
        +-------------+--------+--------+----------+----------+---------+---------------+-----------+------------------------------+------------+

Customer:
        +-----------------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------------+------------+
        |      Field      |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |            StructTag             | Validators |
        +-----------------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------------+------------+
        | id              | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"              |          0 |
        | email           | string | false  | false    | false    | false   | false         | false     | json:"email,omitempty"           |          1 |
        | password        | string | false  | false    | false    | false   | false         | false     | json:"password,omitempty"        |          0 |
        | full_name       | string | false  | true     | false    | false   | false         | false     | json:"full_name,omitempty"       |          0 |
        | billing_address | string | false  | true     | false    | false   | false         | false     | json:"billing_address,omitempty" |          0 |
        | country         | string | false  | true     | false    | false   | false         | false     | json:"country,omitempty"         |          0 |
        | phone           | string | false  | false    | false    | false   | false         | false     | json:"phone,omitempty"           |          0 |
        +-----------------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------------+------------+

Order:
        +------------------+--------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+
        |      Field       |     Type     | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |             StructTag             | Validators |
        +------------------+--------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+
        | id               | int          | false  | false    | false    | false   | false         | false     | json:"id,omitempty"               |          0 |
        | ammount          | float64      | false  | false    | false    | false   | false         | false     | json:"ammount,omitempty"          |          1 |
        | shipping_address | string       | false  | false    | false    | false   | false         | false     | json:"shipping_address,omitempty" |          0 |
        | email            | string       | false  | false    | false    | false   | false         | false     | json:"email,omitempty"            |          0 |
        | date             | time.Time    | false  | false    | false    | true    | false         | false     | json:"date,omitempty"             |          0 |
        | status           | order.Status | false  | false    | false    | false   | false         | false     | json:"status,omitempty"           |          0 |
        +------------------+--------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+

Product:
        +--------------+-----------+--------+----------+----------+---------+---------------+-----------+-------------------------------+------------+
        |    Field     |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag           | Validators |
        +--------------+-----------+--------+----------+----------+---------+---------------+-----------+-------------------------------+------------+
        | id           | int       | false  | false    | false    | false   | false         | false     | json:"id,omitempty"           |          0 |
        | sku          | string    | false  | false    | false    | false   | false         | false     | json:"sku,omitempty"          |          0 |
        | name         | string    | false  | false    | false    | false   | false         | false     | json:"name,omitempty"         |          0 |
        | price        | float64   | false  | false    | false    | false   | false         | false     | json:"price,omitempty"        |          0 |
        | weight       | float64   | false  | false    | false    | false   | false         | false     | json:"weight,omitempty"       |          0 |
        | descriptions | string    | false  | false    | false    | false   | false         | false     | json:"descriptions,omitempty" |          0 |
        | thumbnail    | string    | false  | false    | false    | false   | false         | false     | json:"thumbnail,omitempty"    |          0 |
        | create_date  | time.Time | false  | false    | false    | true    | false         | false     | json:"create_date,omitempty"  |          0 |
        | stock        | int       | false  | false    | false    | true    | false         | false     | json:"stock,omitempty"        |          0 |
        +--------------+-----------+--------+----------+----------+---------+---------------+-----------+-------------------------------+------------+

```

## API Routers
- `Auth`
```
POST
        - /users/login
        - /users/signup
```
- `Customer`
```
POST
        - /users
        - /users/:id/orders
        - /users/:id/purchase
        - /users/:id/cart

GET
        - /users/:id/orders/
        - /users/:id/purchase/:id
        - /users/:id/purchase/all
        - /users/:id/cart/:id
        - /users/:id/cart/all

DELETE
        - /users/:id/orders/:id
        - /users/:id/cart/:id
        - /users/:id/purchase/:id

PATCH
        - /users/:id
```
