## Caveat

Current implementation does not actually provide the best answer. It will give the minimum amount of packs for a given order, but not the minimum amount of products sent; For example

Given the following pack sizes: 33 and 90 and a request for an order of 93 products, the api will return:

* 1 x 90
* 1 x 33

Which is 123 products sent to fulfill an order of 93.

The optimal solution would be to return

* 3 x 33

Which is 99 products.

@todo: implement algorithm to calculate the smallest `x`, where, given a number `n` and a set of integrers `P`,

`x > n` and `x = SUM(a * i)` where `a` belongs to `P` and `i` is an integer.

---

Server will run on port 1080.

## Endpoint: `GET /packs - returns all packs

### Response

```json
[
    {
        "count": 5000
    },
    {
        "count": 2000
    },
    {
        "count": 1000
    },
    {
        "count": 500
    },
    {
        "count": 200
    },
    {
        "count": 100
    }
]
```

## Endpoint `POST /packs` - creates a new pack and adds it to the current list (@todo actually save to DB rather than in memory)

### Request
```json
{
	"count": 100
}
```

### Response

```json
{
    "count": 100
}
```

## Endpoint `POST /order` - given the number of products ordered, the breakdown of packs will be returned

### Request

```json
{
	"count": 1650
}
```

### Response

```json
{
    "order": {
        "1000": 1,
        "250": 1,
        "500": 1
    }
}
```
