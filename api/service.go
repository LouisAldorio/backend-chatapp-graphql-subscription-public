package api

import (
	"encoding/json"
	"math/rand"
	"myapp/api/model"
	"time"
)

func RandomOneImageFromPicsum() string {

	var value []*model.PicsumResponse

	picsum := PicsumApi{
		Route: "/v2/list",
	}

	response := picsum.New().Prepare().Hit()
	if err := json.Unmarshal(response, &value); err != nil {
		panic(err)
	}

	min := 1
	max := len(value) - 1

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(max-min) + min
	return value[index].DownloadUrl
}
