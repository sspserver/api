package protobuf

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/demdxx/xtypes"
	blzmodel "github.com/geniusrabbit/blaze-api/model"
	"github.com/geniusrabbit/blaze-api/repository"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/geniusrabbit/archivarius/internal/archivarius"
	"github.com/geniusrabbit/archivarius/internal/config"
	"github.com/geniusrabbit/archivarius/internal/server/grpc"
)

// GRPCServer is a gRPC server for archivarius service that implements grpc.ArchivariusServer interface.
type GRPCServer struct {
	usecase archivarius.Usecase
	grpc.UnimplementedArchivariusServiceServer
	cfg *config.Config
}

// NewGRPCServer creates a new instance of GRPCServer.
func NewGRPCServer(svc archivarius.Usecase, cfg *config.Config) *GRPCServer {
	return &GRPCServer{usecase: svc, cfg: cfg}
}

// Statistic returns a list of items by the given request.
func (g *GRPCServer) Statistic(ctx context.Context, request *grpc.StatisticRequest) (*grpc.StatisticResponse, error) {
	var (
		filter archivarius.Filter
		order  archivarius.ListOrder
	)

	// Check for filter
	if request.Filter != nil {
		// Iterate over conditions and map each
		for _, condition := range request.Filter.Conditions {
			op, err := opCode(condition.Op)
			if err != nil {
				return nil, err
			}

			var val []any
			for _, value := range condition.Value {
				switch v := value.Value.(type) {
				case *grpc.Value_StringValue:
					val = append(val, v.StringValue)
				case *grpc.Value_IntValue:
					val = append(val, v.IntValue)
				case *grpc.Value_FloatValue:
					val = append(val, v.FloatValue)
				case *grpc.Value_TimeValue:
					val = append(val, v.TimeValue.AsTime())
				case *grpc.Value_IpValue:
					val = append(val, net.ParseIP(v.IpValue))
				default:
					panic("unknown value type")
				}
			}
			cond := archivarius.Condition{
				Key:   keyFromPb(condition.Key),
				Op:    op,
				Value: val,
			}
			filter.Conditions = append(filter.Conditions, &cond)
		}
		if request.Filter.StartDate != nil {
			filter.StartDate = request.Filter.StartDate.AsTime()
		}
		if request.Filter.EndDate != nil {
			filter.EndDate = request.Filter.EndDate.AsTime()
		}
	}

	// Check for order
	if request.Order != nil {
		// Iterate over order keys and map each
		for _, o := range request.Order {
			// Check if key is grouping key
			if okey := orderingKeyFromPb(o.Key); okey.IsGroup() {
				// If key is not in group, return error
				if !xtypes.Slice[grpc.Key](request.Group).Has(func(val grpc.Key) bool {
					key := keyFromPb(val)
					if string(key) == string(okey) {
						return true
					}
					// If key is account id, check for pub and adv account id (SPECIFIC CASE)
					if key == archivarius.KeyAccountID {
						return false ||
							okey == archivarius.OrderingKeyPubAccountID ||
							okey == archivarius.OrderingKeyAdvAccountID
					}
					return false
				}) {
					return nil, fmt.Errorf("ordering key[%s] is not in group", string(okey))
				}
			}

			odir := blzmodel.OrderDesc
			if o.Asc {
				odir = blzmodel.OrderAsc
			}

			// Append key to order
			order.Keys = append(order.Keys, archivarius.KeyOrder{
				Key:   orderingKeyFromPb(o.Key),
				Order: odir,
			})
		}
	}

	// Prepare pagination
	page := archivarius.Pagination{
		Offset: int(request.PageOffset),
		Size:   int(request.PageLimit),
	}

	// Map key from protobuf to archivarius
	keys := xtypes.SliceApply(request.Group, func(v grpc.Key) archivarius.Key { return keyFromPb(v) })
	group := archivarius.WithGroup(keys...)

	// Call usecase statistics
	resp, err := g.usecase.Statistic(ctx, []repository.QOption{&filter, &order, group, &page}...)
	if err != nil {
		return nil, err
	}

	// Call usecase count
	total, err := g.usecase.Count(ctx, []repository.QOption{&filter, group}...)
	if err != nil {
		return nil, err
	}

	// Map response to protobuf
	items := make([]*grpc.Item, 0, len(resp.Items))

	// Iterate over items and map each
	for _, item := range resp.Items {
		keys := make([]*grpc.ItemKey, 0, len(item.Keys))
		for _, key := range item.Keys {
			val, err := valueI(key.Value)
			if err != nil {
				return nil, err
			}
			keys = append(keys, &grpc.ItemKey{
				Key:   keyToPb(key.Key),
				Value: &grpc.Value{Value: val},
			})
		}

		items = append(items, &grpc.Item{
			Keys:        keys,
			Spent:       item.Spent,
			Profit:      item.Profit,
			BidPrice:    item.BidPrice,
			Requests:    item.Requests,
			Impressions: item.Impressions,
			Views:       item.Views,
			Directs:     item.Directs,
			Clicks:      item.Clicks,
			Leads:       item.Leads,
			Bids:        item.Bids,
			Wins:        item.Wins,
			Skips:       item.Skips,
			Nobids:      item.Nobids,
			Errors:      item.Errors,
			Ctr:         item.CTR,
			Ecpm:        item.ECPM,
			Ecpc:        item.ECPC,
			Ecpa:        item.ECPA,
		})
	}

	return &grpc.StatisticResponse{
		Items:      items,
		TotalCount: uint64(total),
	}, nil
}

func opCode(cond grpc.Condition) (archivarius.Op, error) {
	switch cond {
	case grpc.Condition_EQ:
		return archivarius.Eq, nil
	case grpc.Condition_NE:
		return archivarius.NotEq, nil
	case grpc.Condition_GT:
		return archivarius.Gt, nil
	case grpc.Condition_GE:
		return archivarius.Gte, nil
	case grpc.Condition_LT:
		return archivarius.Lt, nil
	case grpc.Condition_LE:
		return archivarius.Lte, nil
	case grpc.Condition_IN:
		return archivarius.In, nil
	case grpc.Condition_NI:
		return archivarius.NotIn, nil
	case grpc.Condition_BT:
		return archivarius.Between, nil
	case grpc.Condition_NB:
		return archivarius.NotBetween, nil
	case grpc.Condition_LI:
		return archivarius.Like, nil
	case grpc.Condition_NL:
		return archivarius.NotLike, nil
	}
	return 0, errors.New("unknown condition operator")
}

func valueI(val any) (grpc.ExtValue, error) {
	switch v := val.(type) {
	case string:
		return &grpc.Value_StringValue{StringValue: v}, nil
	case int:
		return &grpc.Value_IntValue{IntValue: int64(v)}, nil
	case int8:
		return &grpc.Value_IntValue{IntValue: int64(v)}, nil
	case int16:
		return &grpc.Value_IntValue{IntValue: int64(v)}, nil
	case int32:
		return &grpc.Value_IntValue{IntValue: int64(v)}, nil
	case int64:
		return &grpc.Value_IntValue{IntValue: int64(v)}, nil
	case uint:
		return &grpc.Value_UintValue{UintValue: uint64(v)}, nil
	case uint8:
		return &grpc.Value_UintValue{UintValue: uint64(v)}, nil
	case uint16:
		return &grpc.Value_UintValue{UintValue: uint64(v)}, nil
	case uint32:
		return &grpc.Value_UintValue{UintValue: uint64(v)}, nil
	case uint64:
		return &grpc.Value_UintValue{UintValue: v}, nil
	case float32:
		return &grpc.Value_FloatValue{FloatValue: float64(v)}, nil
	case float64:
		return &grpc.Value_FloatValue{FloatValue: v}, nil
	case time.Time:
		return &grpc.Value_TimeValue{TimeValue: timestamppb.New(v)}, nil
	case net.IP:
		return &grpc.Value_IpValue{IpValue: v.String()}, nil
	}
	return nil, errors.New("unknown value type")
}

// Ensure GRPCServer implements grpc.ArchivariusServer
var _ grpc.ArchivariusServiceServer = &GRPCServer{}
