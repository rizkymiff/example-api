package usecase

import "kirito/internal/domain/services"

func GetHelloWorldMessage() string {
    msg := services.GetMessage()
    return msg.Text
}