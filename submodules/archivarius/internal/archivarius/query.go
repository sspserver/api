package archivarius

import (
	"time"

	"github.com/demdxx/xtypes"
	blzmodel "github.com/geniusrabbit/blaze-api/model"
	"github.com/geniusrabbit/blaze-api/repository"
	"gorm.io/gorm"
)

const moneyDivider = "1000000000"

// summingFields define fields that should be summed
var summingFields = []string{
	"adv_spend",
	"adv_potential_spend",
	"adv_failed_spend",
	"adv_compromised_spend",
	"pub_revenue",
	"sales_budget",
	"buying_budget",
	"network_revenue",
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
	"leads",
	"success_leads",
	"failed_leads",
	"compromised_leads",
	"src_bid_requests",
	"src_bid_wins",
	"src_bid_skips",
	"src_bid_nobids",
	"src_bid_errors",
	"ap_bid_requests",
	"ap_bid_wins",
	"ap_bid_skips",
	"ap_bid_nobids",
	"ap_bid_errors",
	"adblocks",
	"privates",
	"robots",
	"backups",
}

// calcFields define fields that should be calculated
var calcFields = []string{
	"IF(imps   > 0, clicks / imps, 0) AS ctr",
	"IF(imps   > 0, clicks / imps, 0) AS ctr",
	"IF(imps   > 0, toFloat64(adv_spend)/" + moneyDivider + " * 1000 / imps, 0) AS ecpm",
	"IF(clicks > 0, toFloat64(adv_spend)/" + moneyDivider + " * 1000 / clicks, 0) AS ecpc",
	"IF(leads  > 0, toFloat64(adv_spend)/" + moneyDivider + " * 1000 / leads, 0) AS ecpa",
	"sales_budget + buying_budget AS bid_price",
	"src_bid_wins + ap_bid_wins AS wins",
	"ap_bid_requests + ap_bid_requests AS bids",
	"src_bid_skips + ap_bid_skips AS skips",
	"src_bid_nobids + ap_bid_nobids AS nobids",
	"src_bid_errors + ap_bid_errors AS errors",
}

// Op is an operation
type Op int

// Operations
const (
	In Op = iota + 1
	NotIn
	Like
	NotLike
	Eq
	NotEq
	Gt
	Gte
	Lt
	Lte
	Between
	NotBetween
	Null
	NotNull
)

// Condition is a key-op-value condition
type Condition struct {
	Key   Key
	Op    Op
	Value []any
}

// PrepareQuery prepare query with condition
func (cond *Condition) PrepareQuery(query *gorm.DB) *gorm.DB {
	if cond == nil {
		return query
	}

	// If account id, then split to adv and pub account ids as it's a composite key for account in the system
	if cond.Key == KeyAccountID {
		expr1, values1 := condExpr(&Condition{Key: KeyAdvAccountID, Op: Eq, Value: cond.Value})
		expr2, values2 := condExpr(&Condition{Key: KeyPubAccountID, Op: Eq, Value: cond.Value})
		query = query.Where("("+expr1+" OR "+expr2+")", append(values1, values2...)...)
	} else {
		expr, values := condExpr(cond)
		query = query.Where(expr, values...)
	}

	return query
}

func condExpr(cond *Condition) (string, []any) {
	if cond == nil {
		return "", nil
	}

	// Prepare query for operation
	switch cond.Op {
	case In:
		return string(cond.Key) + " IN (?)", []any{cond.Value}
	case NotIn:
		return string(cond.Key) + " NOT IN (?)", []any{cond.Value}
	case Like:
		return string(cond.Key) + " LIKE ?", cond.Value[:1]
	case NotLike:
		return string(cond.Key) + " NOT LIKE ?", cond.Value[:1]
	case Eq:
		return string(cond.Key) + " = ?", cond.Value[:1]
	case NotEq:
		return string(cond.Key) + " <> ?", cond.Value[:1]
	case Gt:
		return string(cond.Key) + " > ?", cond.Value[:1]
	case Lt:
		return string(cond.Key) + " < ?", cond.Value[:1]
	case Gte:
		return string(cond.Key) + " >= ?", cond.Value[:1]
	case Lte:
		return string(cond.Key) + " <= ?", cond.Value[:1]
	case Between:
		return string(cond.Key) + " BETWEEN ? AND ?", cond.Value[:2]
	case NotBetween:
		return string(cond.Key) + " NOT BETWEEN ? AND ?", cond.Value[:2]
	case Null:
		return string(cond.Key) + " IS NULL", []any{}
	case NotNull:
		return string(cond.Key) + " IS NOT NULL", []any{}
	}

	return "", nil
}

// Filter of the objects list
type Filter struct {
	Conditions []*Condition
	StartDate  time.Time
	EndDate    time.Time
}

// PrepareQuery prepare query with filter
func (fl *Filter) PrepareQuery(query *gorm.DB) *gorm.DB {
	// Prepare query with filter
	if fl == nil {
		return query
	}

	// Filter by date range
	if !fl.StartDate.IsZero() {
		query = query.Where("timemark >= ?", fl.StartDate)
	}
	if !fl.EndDate.IsZero() {
		query = query.Where("timemark < ?", fl.EndDate)
	}

	// Prepare all the conditions
	for _, cond := range fl.Conditions {
		query = cond.PrepareQuery(query)
	}
	return query
}

// KeyOrder is an ordering key with direction
type KeyOrder struct {
	Key   OrderingKey
	Order blzmodel.Order
}

// ListOrder of the objects list
type ListOrder struct {
	Keys []KeyOrder
}

// PrepareQuery prepare query with order
func (ol *ListOrder) PrepareQuery(query *gorm.DB) *gorm.DB {
	// Prepare query with order
	if ol == nil {
		return query
	}

	// Prepare all the keys
	for _, k := range ol.Keys {
		query = k.Order.PrepareQuery(query, string(k.Key))
	}

	return query
}

// WithGroup returns a group option
func WithGroup(fields ...Key) *GroupOption {
	return &GroupOption{
		Groups:        xtypes.SliceApply(fields, func(v Key) string { return string(v) }),
		SummingFields: summingFields,
		calcFields:    calcFields,
	}
}

// Pagination of the objects list
type Pagination struct {
	Offset int
	Size   int
}

// PrepareQuery prepare query with pagination
func (p *Pagination) PrepareQuery(q *gorm.DB) *gorm.DB {
	if p.Offset > 0 {
		q = q.Offset(p.Offset)
	}
	if p.Size > 0 {
		q = q.Limit(p.Size)
	}
	return q
}

// GroupOption is a group option with grouping cols, summing fields and calculated fields
type GroupOption struct {
	Groups        []string
	SummingFields []string
	calcFields    []string
}

// PrepareQuery prepare query with group
func (opt *GroupOption) PrepareQuery(query *gorm.DB) *gorm.DB {
	// Prepare query with group
	if opt == nil {
		return query
	}

	// Prepare query with group
	for _, group := range opt.Groups {
		query = query.Group(group)
	}

	// Prepare query with fields
	fields := opt.Groups
	fields = append(fields,
		xtypes.SliceApply(opt.SummingFields, func(field string) string {
			return "SUM(" + field + ") AS " + field
		})...,
	)
	fields = append(fields, opt.calcFields...)

	// Select fields
	return query.Select(fields)
}

// List select options
type (
	Option  = repository.QOption
	Options = repository.ListOptions
)
