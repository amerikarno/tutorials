package service

import (
	"wnexam/repository"
)

//Province struct
type Province struct{ 
	ProvinceName string `json:"ProvincName"`
	Total int `json:"Total"`
}

//Ages struct
type Ages struct{
	Below30 int `json:"0-30"`
	Above30lt60 int `json:"30-60"`
	Above60 int `json:"60+"`
	NotAvailable int `json:"N/A"`
}

//WnService interface
type WnService interface{
	GetProvince() (map[string]interface{}, error) 
	GetAges() (*Ages, error)
}

type wnService struct{
	repo repository.DataRepository
}

//NewWnService interface
func NewWnService(repo repository.DataRepository) WnService{
	return wnService{repo}
}

func (s wnService) GetProvince() (map[string]interface{}, error){
	infos, err := s.repo.DecodeJSONUrl()
	if err != nil {
		return nil , err
	}
	pvncs := make(map[string]int)
	for _, data := range infos.Data{
		_, exist := pvncs[data.ProvinceEn]
		if exist {
			pvncs[data.ProvinceEn]++
		} else {
			pvncs[data.ProvinceEn] = 1
		}

		// fmt.Printf("Province: %s\nNation: %s\n",data.ProvinceEn,data.NationEn)
	}
	var provinces []map[string]int
	for k,j := range pvncs{
		// fmt.Printf("\"%s\": %d,",procince,count)
		if k == "" {
			k = "N/A"
		}
		province := map[string]int{
			k:j,
		}
		provinces = append(provinces, province)
	}

	info := map[string]interface{}{
		"Province": provinces,
	}
	// var p map[string]int
	// for _, pv := range provinces{
	// 	p = map[string]int{
	// 		pv.ProvinceName: pv.Total,
	// 	}
	// }
	
	return info, nil
}

func (s wnService) GetAges() (*Ages, error) {
	infos, err := s.repo.DecodeJSONUrl()
	if err != nil {
		return nil , err
	}

	// var ages Ages
	var a60, a3060, a30, na int
	for _, info := range infos.Data{

		switch {
		case info.Age >= 61:
			a60++
		case info.Age >= 31 && info.Age <= 60:
			a3060++
		case info.Age <= 30:
			a30++
		default:
			na++
		}
	}
	ages := Ages{
		Below30: a30,
		Above30lt60: a3060,
		Above60: a60,
		NotAvailable: na,
	}

	return &ages, nil
}