package main

import (
	"database/sql"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func getStatus(chat_id int) (text string, status int, step int, err error) {
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return
	}
	defer db.Close()
	err = db.QueryRow("SELECT * FROM statuses WHERE chat_id = $1", chat_id).Scan(&chat_id, &text, &status, &step)
	if err != nil {
		return
	}
	return
}

func newStatus(chat_id int, text string, status int, step int) (err error) {
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO statuses (chat_id, text, status, step) VALUES ($1, $2, $3, $4)", chat_id, text, status, step)
	if err != nil {
		return
	}
	return
}

func updateStatus(chat_id int, text string, status int, step int) (err error) {
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return
	}
	defer db.Close()
	_, err = db.Exec("UPDATE statuses SET text = $2, status = $3, step = $4 WHERE chat_id = $1", chat_id, text, status, step)
	if err != nil {
		return
	}
	return
}

func getPairFromBase(chat_id int, step int) (pair [2]string, err error) {
	var first_word string
	var second_word string
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return
	}
	defer db.Close()
	err = db.QueryRow("SELECT * FROM words WHERE chat_id = $1 AND step = $2", chat_id, step).Scan(&chat_id, &step, &first_word, &second_word)
	if err != nil {
		return
	}
	pair[0] = first_word
	pair[1] = second_word
	return

}
func getPairsFromBase(chat_id int) (pairs [][2]string, err error) {
	//databaseUrl := "postgresql://agbot:123456@192.168.10.177:5432/agbot"
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM words WHERE chat_id = $1", chat_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var chat_id int
		var step int
		var first_word string
		var second_word string
		var pair [2]string
		err = rows.Scan(&chat_id, &step, &first_word, &second_word)
		pair[0] = first_word
		pair[1] = second_word
		pairs = append(pairs, pair)
	}
	return pairs, err
}

func getPairsAndSaveToBase(chat_id int) error {
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return err
	}
	defer db.Close()
	words, err := openFile("words.txt")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	for i, word := range words {
		if rand.Intn(2) == 1 {
			_, err = db.Exec("insert into words (chat_id, step, first_word, second_word) values ($1, $2, $3 , $4)", chat_id, i, word.eng, word.rus)
			if err != nil {
				return err
			}

		} else {
			_, err = db.Exec("insert into words (chat_id, step, first_word, second_word) values ($1, $2, $3 , $4)", chat_id, i, word.rus, word.eng)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func updatePairsAndSaveToBase(chat_id int) error {
	db, err := sql.Open("postgres", "postgresql://agbot:123456@192.168.10.177:5432/agbot")
	if err != nil {
		return err
	}
	defer db.Close()
	words, err := openFile("words.txt")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	for i, word := range words {
		if rand.Intn(2) == 1 {
			_, err = db.Exec("insert into words (chat_id, step, first_word, second_word) values ($1, $2, $3 , $4)", chat_id, i, word.eng, word.rus)
			if err != nil {
				return err
			}

		} else {
			_, err = db.Exec("insert into words (chat_id, step, first_word, second_word) values ($1, $2, $3 , $4)", chat_id, i, word.rus, word.eng)
			if err != nil {
				return err
			}
		}
	}
	return err
}
