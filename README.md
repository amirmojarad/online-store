# Online Store
web api for online shopping store(REST API)

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
