// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CardLog struct {
	CreateTime time.Time  `json:"createTime"`
	ID         string     `json:"id"`
	Word       *Word      `json:"word"`
	LastResult CardResult `json:"lastResult"`
}

type CardSchedule struct {
	CreateTime   time.Time `json:"createTime"`
	ID           string    `json:"id"`
	Word         *Word     `json:"word"`
	ScheduledFor time.Time `json:"scheduledFor"`
}

type CardStatus struct {
	CardID string     `json:"cardId"`
	Result CardResult `json:"result"`
}

type NewWord struct {
	Lang1 string `json:"lang1"`
	Lang2 string `json:"lang2"`
	Word1 string `json:"word1"`
	Word2 string `json:"word2"`
}

type ScheduledWordsResponse struct {
	Cards []*Word `json:"cards"`
}

type SetSettings struct {
	NewCardsPerDay int `json:"newCardsPerDay"`
}

type UpdateWord struct {
	ID    string `json:"id"`
	Lang1 string `json:"lang1"`
	Lang2 string `json:"lang2"`
	Word1 string `json:"word1"`
	Word2 string `json:"word2"`
}

type UserSettings struct {
	NewCardsPerDay int `json:"newCardsPerDay"`
}

type Word struct {
	ID         string    `json:"id"`
	Lang1      string    `json:"lang1"`
	Lang2      string    `json:"lang2"`
	Word1      string    `json:"word1"`
	Word2      string    `json:"word2"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type CardResult string

const (
	CardResultEasy  CardResult = "Easy"
	CardResultGood  CardResult = "Good"
	CardResultBad   CardResult = "Bad"
	CardResultRetry CardResult = "Retry"
)

var AllCardResult = []CardResult{
	CardResultEasy,
	CardResultGood,
	CardResultBad,
	CardResultRetry,
}

func (e CardResult) IsValid() bool {
	switch e {
	case CardResultEasy, CardResultGood, CardResultBad, CardResultRetry:
		return true
	}
	return false
}

func (e CardResult) String() string {
	return string(e)
}

func (e *CardResult) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CardResult(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CardResult", str)
	}
	return nil
}

func (e CardResult) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
