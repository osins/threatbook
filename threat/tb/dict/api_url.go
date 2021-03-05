package dict

// APIURL struct type(inherit string)
type APIURL string

const (
	// APIV3Root type of APIURL(string)
	APIV3Root APIURL = "https://API.tb.cn/v3"

	// APIV3IpQuery type of APIURL(string)
	APIV3IpQuery APIURL = APIV3Root + "/ip/query"

	// APIV3IpAdvQuery type of APIURL(string)
	APIV3IpAdvQuery APIURL = APIV3Root + "/ip/adv_query"

	// APIV3DomainQuery type of APIURL(string)
	APIV3DomainQuery APIURL = APIV3Root + "/domain/query"

	// APIV3DomainAdvQuery type of APIURL(string)
	APIV3DomainAdvQuery APIURL = APIV3Root + "/domain/adv_query"

	// APIV3DomainSubDomains type of APIURL(string)
	APIV3DomainSubDomains APIURL = APIV3Root + "/domain/sub_domains"

	// APIV3ScenIPReputation type of APIURL(string)
	APIV3ScenIPReputation APIURL = APIV3Root + "/scene/ip_reputation"

	// APIV3SceneDNS type of APIURL(string)
	APIV3SceneDNS APIURL = APIV3Root + "/scene/DNS"

	// APIV3SceneDomainContext type of APIURL(string)
	APIV3SceneDomainContext APIURL = APIV3Root + "/scene/domain_context"

	// APIV3FileUpload type of APIURL(string)
	APIV3FileUpload APIURL = APIV3Root + "/file/upload"

	// APIV3FileReport type of APIURL(string)
	APIV3FileReport APIURL = APIV3Root + "/file/report"

	// APIV3FileReportMultiengines type of APIURL(string)
	APIV3FileReportMultiengines APIURL = APIV3Root + "/file/report/multiengines"

	// APIV3UrlScan type of APIURL(string)
	APIV3UrlScan APIURL = APIV3Root + "/url/scan"

	// APIV3UrlReport type of APIURL(string)
	APIV3UrlReport APIURL = APIV3Root + "/url/report"
)
