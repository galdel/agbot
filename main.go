package main

import (
	"log"
	"strings"
	"time"
)

func main() {
	offset := 0
	status := make(map[int]int)
	words := make(map[int][][2]string)
	step := make(map[int]int)

	for {
		updates, err := getUpdates(offset)
		if err != nil {
			log.Println(err.Error())
		}
		//fmt.Println(updates)
		if len(updates) == 0 {
			time.Sleep(1000)
		} else {
			for _, update := range updates {
				go start(update.Message.Chat.ChatId, update.Message.Text, status, words, step)
				offset = update.UpdateId + 1
			}
		}

	}
}

func start(chat_id int, text string, status map[int]int, words map[int][][2]string, step map[int]int) {
	if status[chat_id] == 0 {
		if strings.Contains(strings.ToLower(text), "go") {
			words[chat_id], _ = getPairs()
			respond(words[chat_id][step[chat_id]][0], chat_id)
			status[chat_id] = 1
			step[chat_id] = 1
		} else {
			respond("Нужно написать: go", chat_id)
		}
	} else if status[chat_id] == 1 {
		if strings.ToLower(text) == strings.ToLower(words[chat_id][step[chat_id]-1][1]) {
			respond("Верно!", chat_id)
		} else {
			respond("Не верно! "+strings.ToUpper(words[chat_id][step[chat_id]-1][1]), chat_id)
		}
		respond(words[chat_id][step[chat_id]][0], chat_id)
		status[chat_id] = 1
		if step[chat_id] < getLen()-1 {
			step[chat_id]++
		} else {
			status[chat_id] = 2
		}

	} else if status[chat_id] == 2 {
		if strings.ToLower(text) == strings.ToLower(words[chat_id][step[chat_id]][1]) {
			respond("Верно!", chat_id)
		} else {
			respond("Не верно! "+strings.ToUpper(words[chat_id][step[chat_id]][1]), chat_id)
		}
		respond("Заново? Напиши go", chat_id)
		status[chat_id] = 0
		step[chat_id] = 0
	}
}
