package logger

import (
	"springo/config"
	"log"
)

const (
	MessageApiStarted = "API Started"
	MessageResourceStarted = "Started Resource:"
	MessageRequestPost = "Post:"
	MessageRequestGet = "Get:"
	MessageRequestPut = "Put:"
	MessageRequestPatch = "Patch:"
	MessageRequestGetAll = "Get All:"
	MessageRequestDelete = "Delete:"
	MessageRequestError = "Error:"
	MessageRequestWarning = "Warning:"
)

func MessageApiStartedLog(baseUrl string, addr string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageApiStarted, " on address ", addr, " in URL ", baseUrl)
}

func MessageResourceStartedLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageResourceStarted, " ", message)
}
func MessageRequestPostLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestPost, " ", message)
}
func MessageRequestGetLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestGet, " ", message)
}
func MessageRequestPutLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestPut, " ", message)
}
func MessageRequestPatchLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestPatch, " ", message)
}
func MessageRequestGetAllLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestGetAll, " ", message)
}
func MessageRequestDeleteLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestDelete, " ", message)
}
func MessageRequestErrorLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestError, " ", message)
}
func MessageRequestWarningLog(message string) {
	log.Println(config.MainConfiguration.Identifier, " - ", MessageRequestWarning, " ", message)
}