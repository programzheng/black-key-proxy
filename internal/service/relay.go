package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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
	if r.Decode != "" && r.Uri != "" {
		switch r.Decode {
		case model.ProxyImageUrl:
			piudo, err := sendProxyImageUrlHttpRequest(string(r.Method), r.Uri)
			if err != nil {
				return "", err
			}
			return piudo.Url, nil
		}
	}

	return r.Uri, nil
}

type ProxyImageUrlDataObject struct {
	Url string `json:"url"`
}

func sendProxyImageUrlHttpRequest(method string, uri string) (*ProxyImageUrlDataObject, error) {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response status code:%d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	piudo := &ProxyImageUrlDataObject{}
	err = json.Unmarshal(body, &piudo)
	if err != nil {
		return nil, err
	}

	return piudo, nil
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
