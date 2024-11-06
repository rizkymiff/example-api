package main

import "kirito/internal/interfaces"

func main() {
    router := interfaces.SetupRouter()
    router.Run(":8080")
}