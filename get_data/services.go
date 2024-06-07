package get_data

import (
	"errors"
	"strings"
)

type Service struct {
	ID   int64
	Name string
}

var servicesData []Service = []Service{
	{ID: 1, Name: "Покупка привилегии"},
}

func GetServices() []string {
	var services []string

	for _, service := range servicesData {
		services = append(services, service.Name)
	}

	return services
}

func GetServiceFromName(serviceName string) (Service, error) {
	var out Service

	var check bool = false
	for _, service := range servicesData {
		if strings.EqualFold(serviceName, service.Name) {
			out = service
			check = true
			break
		}
	}

	if !check {
		return out, errors.New("ServiceNotFound")
	}

	return out, nil
}
