package statistic

import (
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"
)

var counterFields = []string{
	"potential_revenue",
	"failed_revenue",
	"compromised_revenue",
	"revenue",
	// Counters
	"imps",
	"success_imps",
	"failed_imps",
	"compromised_imps",
	"custom_imps",
	"backup_imps",
	"views",
	"failed_views",
	"compromised_views",
	"custom_views",
	"backup_views",
	"directs",
	"success_directs",
	"failed_directs",
	"compromised_directs",
	"custom_directs",
	"backup_directs",
	"clicks",
	"failed_clicks",
	"compromised_clicks",
	"custom_clicks",
	"backup_clicks",
	"bid_requests",
	"bid_wins",
	"bid_skips",
	"bid_nobids",
	"bid_errors",
	"adblocks",
	"privates",
	"robots",
	"backups",
}

type Condition struct {
	Key   Key
	Op    Operation
	Value []any
}

func (cond *Condition) PrepareQuery(query *gorm.DB) *gorm.DB {
	if cond == nil {
		return query
	}
	if strings.ContainsAny(cond.Key.String(), ".-+*$%=<>!\t\n ") {
		panic("invalid condition key: " + cond.Key)
	}
	switch cond.Op {
	case ConditionIn:
		query = query.Where(cond.Key.String()+" IN (?)", cond.Value)
	case ConditionNotIn:
		query = query.Where(cond.Key.String()+" NOT IN (?)", cond.Value)
	case ConditionLike:
		query = query.Where(cond.Key.String()+" LIKE ?", cond.Value[0])
	case ConditionNotLike:
		query = query.Where(cond.Key.String()+" NOT LIKE ?", cond.Value[0])
	case ConditionEq:
		query = query.Where(cond.Key.String()+" = ?", cond.Value[0])
	case ConditionNotEq:
		query = query.Where(cond.Key.String()+" <> ?", cond.Value[0])
	case ConditionGt:
		query = query.Where(cond.Key.String()+" > ?", cond.Value[0])
	case ConditionGtEq:
		query = query.Where(cond.Key.String()+" >= ?", cond.Value[0])
	case ConditionLt:
		query = query.Where(cond.Key.String()+" < ?", cond.Value[0])
	case ConditionLtEq:
		query = query.Where(cond.Key.String()+" <= ?", cond.Value[0])
	case ConditionBetween:
		query = query.Where(cond.Key.String()+" BETWEEN ? AND ?", cond.Value[0], cond.Value[1])
	case ConditionNotBetween:
		query = query.Where(cond.Key.String()+" NOT BETWEEN ? AND ?", cond.Value[0], cond.Value[1])
	case ConditionIsNull:
		query = query.Where(cond.Key.String() + " IS NULL")
	case ConditionIsNotNull:
		query = query.Where(cond.Key.String() + " IS NOT NULL")
	default:
		panic("unknown condition operator: " + gocast.Str(int(cond.Op)))
	}
	return query
}

// Filter of the objects list
type Filter struct {
	Conditions []*Condition
	StartDate  time.Time
	EndDate    time.Time
}

func (fl *Filter) PrepareQuery(query *gorm.DB) *gorm.DB {
	if fl == nil {
		return query
	}
	if !fl.StartDate.IsZero() {
		query = query.Where("timemark >= ?", fl.StartDate)
	}
	if !fl.EndDate.IsZero() {
		query = query.Where("timemark < ?", fl.EndDate)
	}
	for _, cond := range fl.Conditions {
		query = cond.PrepareQuery(query)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil {
		return query
	}
	return query
}

func WithGroup(fields ...Key) *repository.GroupOption {
	return &repository.GroupOption{
		Groups:        xtypes.SliceApply(fields, func(key Key) string { return key.String() }),
		SummingFields: counterFields,
	}
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
