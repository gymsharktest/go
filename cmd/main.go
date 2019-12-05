package main

import (
	api "radu/api/internal"
)

func main() {
	api.AddPack(api.Pack{Count: 250})
	api.AddPack(api.Pack{Count: 500})
	api.AddPack(api.Pack{Count: 1000})
	api.AddPack(api.Pack{Count: 2000})
	api.AddPack(api.Pack{Count: 5000})
	api.StartServer()
}
