package dict

// ExcludeType struct type(inherit string)
type ExcludeType string

const (
	// ExcludeAsn asn信息。
	ExcludeAsn ExcludeType = "asn"
	// ExcludePorts 端口信息。
	ExcludePorts ExcludeType = "ports"
	// ExcludeCas SSL证书等信息。
	ExcludeCas ExcludeType = "cas"
	// ExcludeRdnsList rDNS记录信息。
	ExcludeRdnsList ExcludeType = "rDNS_list"
	// ExcludeIntelligences 威胁情报。
	ExcludeIntelligences ExcludeType = "intelligences"
	// ExcludeJudgments 从威胁情报中分析，综合判定的威胁类型。
	ExcludeJudgments ExcludeType = "judgments"
	// ExcludeTagsClasses 相关攻击团伙或安全事件信息标签等。
	ExcludeTagsClasses ExcludeType = "tags_classes"
	// ExcludeSamples 相关样本。
	ExcludeSamples ExcludeType = "samples"
	// ExcludeUpdateTime 情报最近更新时间。
	ExcludeUpdateTime ExcludeType = "update_time"
	// ExcludeSumCurDomains 当前指向域名的数量。
	ExcludeSumCurDomains ExcludeType = "sum_cur_domains"
	// ExcludeScene 应用场景。
	ExcludeScene ExcludeType = "scene"
)
