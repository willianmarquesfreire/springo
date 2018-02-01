package logger

import (
	"springo/config"
	"log"
)

const (
	ExceptionApiNotStarted       = "API not started"
	ExceptionDatabaseError       = "Database error"
	ExceptionDomainAlreadyExists = "Domain already exists"
	ExceptionInvalidId           = "Invalid id"
	ExceptionFailedFind          = "Failed find"
	ExceptionIncorrectBody       = "Incorrect body"
	ExceptionNotFound            = "Not found"
	ExceptionFailedDelete        = "Failed Delete"
	ExceptionFailedGetAll        = "Failed get all"
	ExceptionNoToken             = "No Token"
	ExceptionInvalidToken        = "Invalid Token"
	ExceptionInstanceToken       = "Token belongs to Instance"
	ExceptionUserHasExists       = "User has exists"
)
func ExceptionApiNotStartedLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionApiNotStarted, " ", message)
}
func ExceptionDatabaseErrorLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionDatabaseError, " ", message)
}
func ExceptionDomainAlreadyExistsLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionDomainAlreadyExists, " ", message)
}
func ExceptionInvalidIdLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionInvalidId, " ", message)
}
func ExceptionFailedFindLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionFailedFind, " ", message)
}
func ExceptionIncorrectBodyLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionIncorrectBody, " ", message)
}
func ExceptionNotFoundLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionNotFound, " ", message)
}
func ExceptionFailedDeleteLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionFailedDelete, " ", message)
}
func ExceptionFailedGetAllLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionFailedGetAll, " ", message)
}
func ExceptionNoTokenLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionNoToken, " ", message)
}
func ExceptionInvalidTokenLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionInvalidToken, " ", message)
}
func ExceptionInstanceTokenLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionInstanceToken, " ", message)
}
func ExceptionUserHasExistsLog(message string) {
	log.Fatal(config.MainConfiguration.Identifier, " - ", ExceptionUserHasExists, " ", message)
}
