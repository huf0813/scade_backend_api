package covid_hospital

import (
	"encoding/json"
	"github.com/huf0813/scade_backend_api/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const upstream = "https://dekontaminasi.com/api/id/covid19/hospitals"

func extract() ([]domain.Hospital, error) {
	res, err := http.Get(upstream)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var hospitals []domain.Hospital
	if err := json.Unmarshal(data, &hospitals); err != nil {
		panic(err)
	}
	return hospitals, nil
}

func transform(hospital *domain.Hospital) *domain.Hospital {
	splitName := hospital.Region
	splitResult := strings.Split(splitName, ",")
	findCity := strings.ReplaceAll(splitResult[0], "KOTA ", "")
	hospital.Region = findCity
	return hospital
}

func Load() ([]domain.Hospital, error) {
	result, err := extract()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(result); i++ {
		tmp := result[i]
		newTransform := transform(&tmp)
		result[i].Region = newTransform.Region
	}
	return result, nil
}
