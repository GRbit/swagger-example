package handlers

import (
	"net/http"

	"github.com/grbit/swagger-example/pkg/java/models"
)

func HandleCreateUpdatePropertyInfoById(r *http.Request, body models.PropertyInfoData, propertyID string,
) (
	models.PropertyInfoResponse, error,
) {
	return models.PropertyInfoResponse{}, nil
}

func HandleGetPropertyInfoById(r *http.Request, propertyID string) (models.PropertyInfoResponse, error) {
	return models.PropertyInfoResponse{}, nil
}

func HandleGetPropertiesInfo(r *http.Request) (models.PropertiesInfoResponse, error) {
	return models.PropertiesInfoResponse{}, nil
}
