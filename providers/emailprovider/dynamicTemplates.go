package emailprovider

import (
	"GO-08/providers"
)

func contactUserTemplate() (*providers.DynamicTemplate, error) {
	return &providers.DynamicTemplate{
		TemplateID:  "d-162a7456fc274f08ba5b46ad1336b1e8", // todo change template ID
		Categories:  []string{"Contact User"},
		DynamicData: make(map[string]interface{}),
	}, nil
}
