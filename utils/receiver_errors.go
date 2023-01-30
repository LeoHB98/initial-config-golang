package utils

import (
	"log"
	"reflect"

	"gitalpha.com/go/nome_projeto/config"
	"gitalpha.com/go/nome_projeto/models"
)

func ReceiverErrors(err interface{}, mail *config.Email) {

	if err != nil {

		log.Println(err)
		var errors []models.Log
		er := reflect.ValueOf(err)

		if er.Kind() == reflect.Slice {
			for i := 0; i < er.Len(); i++ {
				ReceiverErrors(er.Index(i).Interface(), mail)
			}
		}

		if er.Kind() == reflect.Struct {
			for i := 0; i < er.NumField(); i++ {
				ReceiverErrors(er.Field(i).Interface(), mail)
			}
		}

		anyErr, ok := err.(error)
		if !ok {
			return
		}
		errors = append(errors, models.Log{Errr: anyErr})
		SendEmail(errors, mail)

	}

}
