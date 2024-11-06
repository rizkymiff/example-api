package services

import "kirito/internal/domain/entities"

func GetMessage() entities.Message {
    return entities.Message{Text: "Hello World"}
}