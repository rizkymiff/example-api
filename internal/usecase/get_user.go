package usecase

import "kirito/internal/domain/services"

func GetDummyUser() interface{} {
    return services.GetDummyUser()
}