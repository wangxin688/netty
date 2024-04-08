package zbx

type ZbxObjectName string

const (
	Host          ZbxObjectName = "host"
	Hostgroup     ZbxObjectName = "hostgroup"
	Template      ZbxObjectName = "template"
	TemplateGroup ZbxObjectName = "templategroup"
	Proxy         ZbxObjectName = "proxy"
	Item          ZbxObjectName = "item"
	History       ZbxObjectName = "history"
	MediaType     ZbxObjectName = "mediatype"
	Action        ZbxObjectName = "action"
	Configuration ZbxObjectName = "configuration"
	HostInterface ZbxObjectName = "hostinterface"
	UserMacro     ZbxObjectName = "usermacro"
)

type ZbxCreateSchema [T]map[string]any

type BaseClient[ZbxCreateSchema, ZbxUpdateSchema, ZbxGetSchema any] struct {
	ZbxUrl    string
	ZbxObject ZbxObjectName
	Headers   map[string]string
}

func NewClient[ZbxCreateSchema, ZbxUpdateSchema, ZbxGetSchema any](
	zbxUrl string, token string) *BaseClient[ZbxCreateSchema, ZbxUpdateSchema, ZbxGetSchema] {

	return &BaseClient[ZbxCreateSchema, ZbxUpdateSchema, ZbxGetSchema]{
		ZbxUrl: zbxUrl + "/api_jsonrpc.php",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorzation": "Bearer " + token,
		},
	}
}

// func create[](createSchema T)
