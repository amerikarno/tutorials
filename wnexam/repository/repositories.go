package repository

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Data struct create
type Data struct{
	Data []JSONData `json:"Data"`
}

//JSONData struct create
type JSONData struct{
	ConfirmDate string `json:"ConfirmDate"`
	No string `json:"No"`
	Age int `json:"Age"`
	Gender string `json:"Gender"`
	GenderEn string `json:"GenderEn"`
	Nation string `json:"Nation"`
	NationEn string `json:"NationEn"`
	Province string `json:"Province"`
	ProvinceID int `json:"ProvinceId"`
	District string `json:"District"`
	ProvinceEn string `json:"ProvinceEn"`
	StatQuarantine int `json:"StatQuarantine"`
}

//DataRepository interface create
type DataRepository interface {
	DecodeJSONUrl() (Data, error)
	GetJSONUrl() ([]byte,error)
}

type dataRepository struct{
	url string
}

//NewDataRepository create
func NewDataRepository(url string) DataRepository{
	return dataRepository{url:url}
}

func (r dataRepository) DecodeJSONUrl() (Data, error) {
	body, err := r.GetJSONUrl()
	if err != nil {
		log.Panicln(err)
	}
	decode := Data{}
	err = json.Unmarshal(body, &decode)
	if err != nil {
		log.Println(err)
	}
	return decode, nil
}

func (r dataRepository) GetJSONUrl() ([]byte,error) {
	spaceClient := http.Client{
		Timeout: time.Second *2,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				},
				},
			}
	req, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		log.Println(err)
	}

	res, err := spaceClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	
	return body,nil
}


