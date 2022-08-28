package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PicsumApi struct {
	picsumClient *http.Client
	Route        string
	request      *http.Request
}

func (r *PicsumApi) New() *PicsumApi {
	r.picsumClient = &http.Client{}

	return r
}

func (r *PicsumApi) Prepare() *PicsumApi {

	req, err := http.NewRequest("GET", os.Getenv("PICSUM_BASEURL")+r.Route, nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("limit", "1000")

	req.URL.RawQuery = q.Encode()
	r.request = req

	return r
}

func (r *PicsumApi) Hit() []byte {

	resp, err := r.picsumClient.Do(r.request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}
