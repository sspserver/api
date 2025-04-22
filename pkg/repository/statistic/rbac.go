package statistic

import "github.com/geniusrabbit/blaze-api/pkg/acl"

const (
	RBACStatisticObjectName = `statistic`
)

var (
	RBACStatisticObject = acl.RBACType{
		ResourceName: RBACStatisticObjectName,
	}
)
