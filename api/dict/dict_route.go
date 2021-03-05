package dict

// RouteURL struct type(inherit string)
type RouteURL string

const (
	// RouteRoot type of RouteURL
	RouteRoot RouteURL = ""
	// RouteV1IpQuery type of RouteURL
	RouteV1IpQuery RouteURL = RouteRoot + "/ip/query"
	// RouteV1IpAdvQuery type of RouteURL
	RouteV1IpAdvQuery RouteURL = RouteRoot + "/ip/adv_query"

	// RouteV1DomainQuery type of RouteURL
	RouteV1DomainQuery RouteURL = RouteRoot + "/domain/query"
	// RouteV1DomainAdvQuery type of RouteURL
	RouteV1DomainAdvQuery RouteURL = RouteRoot + "/domain/adv_query"
	// RouteV1DomainSubDomains type of RouteURL
	RouteV1DomainSubDomains RouteURL = RouteRoot + "/domain/sub_domains"

	// RouteV1SceneIPReputation type of RouteURL
	RouteV1SceneIPReputation RouteURL = RouteRoot + "/scene/ip_reputation"
	// RouteV1SceneDNS type of RouteURL
	RouteV1SceneDNS RouteURL = RouteRoot + "/scene/DNS"
	// RouteV1SceneDomainContext type of RouteURL
	RouteV1SceneDomainContext RouteURL = RouteRoot + "/scene/domain_context"

	// RouteV1FileUpload type of RouteURL
	RouteV1FileUpload RouteURL = RouteRoot + "/file/upload"
	// RouteV1FileReport type of RouteURL
	RouteV1FileReport RouteURL = RouteRoot + "/file/report"
	// RouteV1FileReportMultiengines type of RouteURL
	RouteV1FileReportMultiengines RouteURL = RouteRoot + "/file/report/multiengines"

	// RouteV1UrlScan type of RouteURL
	RouteV1UrlScan RouteURL = RouteRoot + "/url/scan"
	// RouteV1UrlReport type of RouteURL
	RouteV1UrlReport RouteURL = RouteRoot + "/url/report"
)
