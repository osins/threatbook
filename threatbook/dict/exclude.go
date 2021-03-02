package dict

type ExcludeType string

const (
	//ASN: asn信息。
	EXCLUDE_ASN ExcludeType = "asn"
	//PORTS: 端口信息。
	EXCLUDE_PORTS ExcludeType = "ports"
	//CAS: SSL证书等信息。
	EXCLUDE_CAS ExcludeType = "cas"
	//RDNS_LIST: rdns记录信息。
	EXCLUDE_RDNS_LIST ExcludeType = "rdns_list"
	//INTELLIGENCES: 威胁情报。
	EXCLUDE_INTELLIGENCES ExcludeType = "intelligences"
	//JUDGMENTS: 从威胁情报中分析，综合判定的威胁类型。
	EXCLUDE_JUDGMENTS ExcludeType = "judgments"
	//TAGS_CLASSES: 相关攻击团伙或安全事件信息标签等。
	EXCLUDE_TAGS_CLASSES ExcludeType = "tags_classes"
	//SAMPLES: 相关样本。
	EXCLUDE_SAMPLES ExcludeType = "samples"
	//UPDATE_TIME: 情报最近更新时间。
	EXCLUDE_UPDATE_TIME ExcludeType = "update_time"
	//SUM_CUR_DOMAINS: 当前指向域名的数量。
	EXCLUDE_SUM_CUR_DOMAINS ExcludeType = "sum_cur_domains"
	//SCENE: 应用场景。
	EXCLUDE_SCENE ExcludeType = "scene"
)
