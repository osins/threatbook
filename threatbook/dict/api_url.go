package dict

type ApiURL string

const (
	API_V3_ROOT         ApiURL = "https://api.threatbook.cn/v3"
	API_V3_IP_QUERY     ApiURL = API_V3_ROOT + "/ip/query"
	API_V3_IP_ADV_QUERY ApiURL = API_V3_ROOT + "/ip/adv_query"

	API_V3_DOMAIN_QUERY       ApiURL = API_V3_ROOT + "/domain/query"
	API_V3_DOMAIN_ADV_QUERY   ApiURL = API_V3_ROOT + "/domain/adv_query"
	API_V3_DOMAIN_SUB_DOMAINS ApiURL = API_V3_ROOT + "/domain/sub_domains"

	API_V3_SCEN_IP_REPUTATION   ApiURL = API_V3_ROOT + "/scene/ip_reputation"
	API_V3_SCENE_DNS            ApiURL = API_V3_ROOT + "/scene/dns"
	API_V3_SCENE_DOMAIN_CONTEXT ApiURL = API_V3_ROOT + "/scene/domain_context"

	API_V3_FILE_UPLOAD              ApiURL = API_V3_ROOT + "/file/upload"
	API_V3_FILE_REPORT              ApiURL = API_V3_ROOT + "/file/report"
	API_V3_FILE_REPORT_MULTIENGINES ApiURL = API_V3_ROOT + "/file/report/multiengines"
)
