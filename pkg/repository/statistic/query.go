package statistic

import (
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/model"
	"github.com/geniusrabbit/blaze-api/repository"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/adformat"
	adfrepo "github.com/sspserver/api/pkg/repository/adformat/repository"
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
	// Convert KeyFormatCode to KeyFormatID
	if cond.Key == KeyFormatCode {
		vals := gocast.AnySlice[string](cond.Value)
		fmts := adfrepo.New()
		formats, err := fmts.FetchList(query.Statement.Context, &adformat.Filter{Codename: vals})
		if err != nil {
			panic("failed to fetch ad formats for condition: " + err.Error())
		}
		cond.Value = xtypes.SliceApply(formats, func(format *models.Format) any {
			return format.ID
		})
		cond.Key = KeyFormatID
	}
	// Process conditions
	switch cond.Op {
	case ConditionIn:
		if len(cond.Value) > 0 {
			query = query.Where(cond.Key.String()+" IN (?)", cond.Value)
		}
	case ConditionNotIn:
		if len(cond.Value) > 0 {
			query = query.Where(cond.Key.String()+" NOT IN (?)", cond.Value)
		}
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
		query = query.Where("timemark >= ?", time.Date(
			fl.StartDate.Year(),
			fl.StartDate.Month(),
			fl.StartDate.Day(),
			fl.StartDate.Hour(),
			fl.StartDate.Minute(),
			fl.StartDate.Second(),
			0, time.UTC,
		))
	}
	if !fl.EndDate.IsZero() {
		query = query.Where("timemark < ?", time.Date(
			fl.EndDate.Year(),
			fl.EndDate.Month(),
			fl.EndDate.Day(),
			fl.EndDate.Hour(),
			fl.EndDate.Minute(),
			fl.EndDate.Second(),
			0, time.UTC,
		))
	}
	for _, cond := range fl.Conditions {
		query = cond.PrepareQuery(query)
	}
	return query
}

// ListOrder of the objects list
type ListOrder struct {
	Orders map[OrderingKey]models.Order
}

func (ol *ListOrder) SetOrder(key OrderingKey, order models.Order) {
	if ol == nil {
		return
	}
	if ol.Orders == nil {
		ol.Orders = make(map[OrderingKey]models.Order)
	}
	if key == OrderingKeyFormatCode {
		key = OrderingKeyFormatID
	}
	ol.Orders[key] = order
}

func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	if ol == nil || len(ol.Orders) == 0 {
		return query
	}
	for key, order := range ol.Orders {
		if key == OrderingKeyFormatCode {
			key = OrderingKeyFormatID
		}
		// query = order.PrepareQuery(query, key.String())
		query = query.Order(clause.OrderByColumn{
			Column: clause.Column{Name: key.String(), Raw: strings.ContainsAny(key.String(), " \t\n\r()<>=!@#$%^&*|`~{}[]'\"+-*/\\")},
			Desc:   order == model.OrderDesc,
		})
	}
	return query
}

func WithGroup(fields ...Key) *repository.GroupOption {
	return &repository.GroupOption{
		Groups: xtypes.SliceApply(fields, func(key Key) string {
			if key == KeyFormatCode {
				return KeyFormatID.String()
			}
			return key.String()
		}),
		SummingFields: counterFields,
		Ext:           fields,
	}
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
