package utils

import (
	"math/rand"
	"myapp/graph/model"
	guuid "github.com/google/uuid"
)

func FilterSubscribers(ss []*model.Subscriber, currentId int) (ret []*model.Subscriber) {
	for _, s := range ss {
		if s.ID != currentId {
			ret = append(ret, s)
		}
	}
	return
}

func FindNewJoinedSubscribers(ss []*model.Subscriber, id int) int {
	for index, s := range ss {
		if s.ID == id {
			return index
		}
	}

	return -1
}

func RandString(n int) string {

    var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890#@!$%^&*()")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


func UniqueIntSlice(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GenerateUniqueID() string {
	return guuid.New().String()
}