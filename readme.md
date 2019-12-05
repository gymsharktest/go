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