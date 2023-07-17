package service

import (
	"database/sql"
	"log"

	"github.com/programzheng/black-key-proxy/internal/model"
)

type RelayEventDataObject struct {
	Identifier string
	Key        string
}

func CreateRelayEventByRelayEventDataObject(relayEventDataObject *RelayEventDataObject) *model.RelayEvent {
	identifier := &sql.NullString{}
	if relayEventDataObject.Identifier != "" {
		identifier.String = relayEventDataObject.Identifier
	}

	re := &model.RelayEvent{
		Identifier: *identifier,
		Key:        relayEventDataObject.Key,
	}
	model.DB.Create(re)

	return re
}

func getRelayEventByRelayEventDataObject(relayEventDataObject *RelayEventDataObject) (*model.RelayEvent, error) {
	re := &model.RelayEvent{Key: relayEventDataObject.Key}

	rew := &model.RelayEvent{Key: relayEventDataObject.Key}
	identifier := &sql.NullString{String: "NULL"}
	if relayEventDataObject.Identifier != "" {
		identifier.String = relayEventDataObject.Identifier
		identifier.Valid = true
	}
	rew.Identifier = *identifier

	err := model.DB.Where(rew).Find(re).Error
	if err != nil {
		log.Printf("GetEventByKey error: %v", err)
		return nil, err
	}
	return re, nil
}

type SendGetImageUrlByRelayEventDataObjectResults struct {
	StatusCode string
	Url        *string
}

func getImageUrlByReplyEvent(re *model.RelayEvent) (string, error) {
	r := &model.Relay{}
	err := model.DB.Model(re).Association("Relay").Find(r)
	if err != nil {
		return "", err
	}

	return r.Uri, nil
}

func SendGetImageUrlByRelayEventDataObject(do *RelayEventDataObject) *SendGetImageUrlByRelayEventDataObjectResults {
	re, err := getRelayEventByRelayEventDataObject(&RelayEventDataObject{
		Key:        do.Key,
		Identifier: do.Identifier,
	})
	if err != nil {
		return &SendGetImageUrlByRelayEventDataObjectResults{
			StatusCode: "error",
			Url:        nil,
		}
	}

	url, err := getImageUrlByReplyEvent(re)
	if err != nil {
		return &SendGetImageUrlByRelayEventDataObjectResults{
			StatusCode: "error",
			Url:        nil,
		}
	}

	return &SendGetImageUrlByRelayEventDataObjectResults{
		StatusCode: "success",
		Url:        &url,
	}
}
