package doc

import _ "embed"

var (

	//go:embed pkg/swagger/template_service/template_service.swagger.json
	SwaggerDocService []byte
)
