package openapi

type Swagger struct {
	Version    string           `json:"openapi,omitempty" yaml:"openapi,omitempty"`
	Info       *Info            `json:"info,omitempty" yaml:"info,omitempty"`
	Servers    []*Server        `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths      map[string]*Path `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components *Components      `json:"components,omitempty" yaml:"components,omitempty"`
}
type Info struct {
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Version     string `json:"version,omitempty" yaml:"version,omitempty"`
}
type Server struct {
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
type Path struct {
	Get    *Method `json:"get,omitempty" yaml:"get,omitempty"`
	Put    *Method `json:"put,omitempty" yaml:"put,omitempty"`
	Post   *Method `json:"post,omitempty" yaml:"post,omitempty"`
	Delete *Method `json:"delete,omitempty" yaml:"delete,omitempty"`
}
type Method struct {
	OperationId string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Summary     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Parameters  []*Parameter          `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Request     *Request              `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   map[int]*Response     `json:"responses,omitempty" yaml:"responses,omitempty"`
	Security    []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
}
type Components struct {
	SecuritySchemes map[string]*Security  `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Parameters      map[string]*Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses       map[int]*Response     `json:"responses,omitempty" yaml:"responses,omitempty"`
	Schemas         map[string]*Schema    `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}
type Security struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
}
type Response struct {
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*Content `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
type Parameter struct {
	Name        string  `json:"name,omitempty" yaml:"name,omitempty"`
	In          string  `json:"in,omitempty" yaml:"in,omitempty"`
	Schema      *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example     string  `json:"example,omitempty" yaml:"example,omitempty"`
	Required    bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
type Request struct {
	Description string              `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*Content `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string              `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
type Content struct {
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}
type Schema struct {
	Type                 string                `json:"type,omitempty" yaml:"type,omitempty"`
	Format               string                `json:"format,omitempty" yaml:"format,omitempty"`
	ReadOnly             bool                  `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	Required             bool                  `json:"required,omitempty" yaml:"required,omitempty"`
	AdditionalProperties *AdditionalProperties `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Items                *Schema               `json:"items,omitempty" yaml:"items,omitempty"`
	Properties           map[string]*Schema    `json:"properties,omitempty" yaml:"properties,omitempty"`
	Description          string                `json:"description,omitempty" yaml:"description,omitempty"`
	Ref                  string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
type AdditionalProperties struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
