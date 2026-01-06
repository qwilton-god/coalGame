package company

import "errors"

var ErrInsufficientCoals = errors.New("Недостаточно угля на балансе компании.")
var ErrUnknownMinerType = errors.New("Неизвестный тип шахтера.")
var ErrUnknownProductName = errors.New("Неизвестный продукт.")
var ErrProductAlreadyBuyed = errors.New("Данный предмет уже был куплен.")
var ErrDontCompletedConditions = errors.New("Не выполены все условия.")
var ErrNotFound = errors.New("Не найдено.")
