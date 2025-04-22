package statistic

type Operation int32

const (
	ConditionUndefined Operation = iota
	ConditionEq
	ConditionNotEq
	ConditionGt
	ConditionGtEq
	ConditionLt
	ConditionLtEq
	ConditionIn
	ConditionNotIn
	ConditionBetween
	ConditionNotBetween
	ConditionLike
	ConditionNotLike
	ConditionIsNull
	ConditionIsNotNull
)
