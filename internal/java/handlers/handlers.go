package handlers

import (
	"net/http"

	"github.com/grbit/swagger-example/pkg/java/models"
)

var DB map[string]models.PropertyInfoData

func HandleCreateUpdatePropertyInfoById(r *http.Request, body models.PropertyInfoData, propertyID string,
) (
	models.PropertyInfoResponse, error,
) {
	if DB == nil {
		DB = make(map[string]models.PropertyInfoData)
	}

	DB[propertyID] = body

	return models.PropertyInfoResponse{}, nil
}

func HandleGetPropertyInfoById(r *http.Request, propertyID string) (models.PropertyInfoResponse, error) {
	p, ok := DB[propertyID]
	if !ok {
		return models.PropertyInfoResponse{
			Data:      &models.PropertyInfoData{},
			RequestId: "",
		}, nil
	}

	return models.PropertyInfoResponse{
		Data:      &p,
		RequestId: "",
	}, nil
}

func HandleGetPropertiesInfo(r *http.Request) (models.PropertiesInfoResponse, error) {
	pp := []models.PropertyInfoData{}

	for _, p := range DB {
		pp = append(pp, p)
	}

	return models.PropertiesInfoResponse{
		Data:      pp,
		RequestId: "",
	}, nil
}
