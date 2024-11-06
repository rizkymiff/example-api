package services

import "kirito/internal/domain/entities"

func GetDummyUser() entities.User {
    return entities.User{
        ID:    1,
        Name:  "Achmad Miftahul Rizqi",
        Email: "tizqy.miftaqul77@gmail.com",
    }
}