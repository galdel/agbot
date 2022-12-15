package main

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

func main() {
	offset := 0
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
				go start(update.Message.Chat.ChatId, update.Message.Text)
				offset = update.UpdateId + 1
			}
		}

	}
}

func start(chat_id int, text string) {
	_, status, step, err := getStatus(chat_id)
	if err == sql.ErrNoRows {
		status = 0
		step = 0
		err = newStatus(chat_id, text, status, step)
		err = getPairsAndSaveToBase(chat_id)
		if strings.Contains(strings.ToLower(text), "go") {
			pair, _ := getPairFromBase(chat_id, step)
			respond(pair[0], chat_id)
			status = 1
			step = 1
			updateStatus(chat_id, text, status, step)
		} else {
			respond("Нужно написать: go", chat_id)
			status = 0
			step = 0
		}
	} else {
		if status == 0 {
			pair, _ := getPairFromBase(chat_id, step)
			if strings.Contains(strings.ToLower(text), "go") {
				respond(pair[0], chat_id)
				status = 1
				step = 1
				updateStatus(chat_id, text, status, step)
			} else {
				respond("Нужно написать: go", chat_id)
			}
		} else if status == 1 {
			oldPair, _ := getPairFromBase(chat_id, step-1)
			pair, _ := getPairFromBase(chat_id, step)
			if strings.ToLower(text) == strings.ToLower(oldPair[1]) {
				respond("Верно!", chat_id)
			} else {
				respond("Не верно! "+strings.ToUpper(oldPair[1]), chat_id)
			}
			respond(pair[0], chat_id)
			status = 1
			pairs, _ := getPairsFromBase(chat_id)
			if step < len(pairs)-1 {
				step++
			} else {
				status = 2
			}
			updateStatus(chat_id, text, status, step)

		} else if status == 2 {
			//oldPair, _ := getPairFromBase(chat_id, step-1)
			pair, _ := getPairFromBase(chat_id, step)
			if strings.ToLower(text) == strings.ToLower(pair[1]) {
				respond("Верно!", chat_id)
			} else {
				respond("Не верно! "+strings.ToUpper(pair[1]), chat_id)
			}
			respond("Заново? Напиши go", chat_id)
			status = 0
			step = 0
			updateStatus(chat_id, text, status, step)
		}

	}
}
