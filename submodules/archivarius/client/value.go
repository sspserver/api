package client

import (
	"errors"
	"net"
	"time"

	"github.com/geniusrabbit/archivarius/internal/server/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ErrUnknownValueType is an error for unknown value type
var ErrUnknownValueType = errors.New("unknown value type")

// AnyExtValue converts any value to grpc.ExtValue
func AnyExtValue(val any) (grpc.ExtValue, error) {
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
	return nil, ErrUnknownValueType
}

// MustExtValue converts any value to grpc.ExtValue and panics on error
func MustExtValue(val any) grpc.ExtValue {
	v, err := AnyExtValue(val)
	if err != nil {
		panic(err)
	}
	return v
}

// MakeExtValue converts any value to grpc.ExtValue and ignores error
func MakeExtValue(val any) grpc.ExtValue {
	v, _ := AnyExtValue(val)
	return v
}

// ValueTypeEnum is a value type enum
type ValueTypeEnum int

const (
	ValueTypeUnknown ValueTypeEnum = iota
	ValueTypeString
	ValueTypeInt
	ValueTypeUint
	ValueTypeFloat
	ValueTypeTime
	ValueTypeIp
)

// ValueExtType returns value type
func ValueExtType(v any) ValueTypeEnum {
	switch v.(type) {
	case *grpc.Value_StringValue:
		return ValueTypeString
	case *grpc.Value_IntValue:
		return ValueTypeInt
	case *grpc.Value_UintValue:
		return ValueTypeUint
	case *grpc.Value_FloatValue:
		return ValueTypeFloat
	case *grpc.Value_TimeValue:
		return ValueTypeTime
	case *grpc.Value_IpValue:
		return ValueTypeIp
	}
	return ValueTypeUnknown
}

// ValueExtFrom converts grpc.Value to any
func ValueExtFrom(v any) any {
	switch v := v.(type) {
	case nil:
	case *grpc.Value_StringValue:
		return v.StringValue
	case *grpc.Value_IntValue:
		return v.IntValue
	case *grpc.Value_UintValue:
		return v.UintValue
	case *grpc.Value_FloatValue:
		return v.FloatValue
	case *grpc.Value_TimeValue:
		return v.TimeValue.AsTime()
	case *grpc.Value_IpValue:
		return net.ParseIP(v.IpValue)
	}
	return nil
}
