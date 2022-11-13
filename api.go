package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}
type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
type Chat struct {
	ChatId   int    `json:"id"`
	UserName string `json:"username"`
}
type RestResponse struct {
	Result []Update `json:"result"`
}
type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

// https://api.telegram.org/bot<token>/METHOD_NAME
var token = "5700847634:AAHcNZbMNQt4jA6nwRzO6LAb6GRDVD3xdWI"
var api = "https://api.telegram.org/bot"
var url = api + token

func getUpdates(offset int) ([]Update, error) {
	resp, err := http.Get(url + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(text string, chatId int) error {
	var botMessage BotMessage
	botMessage.ChatId = chatId
	botMessage.Text = text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	fmt.Println(botMessage)
	if err != nil {
		return err
	}
	return nil
}
