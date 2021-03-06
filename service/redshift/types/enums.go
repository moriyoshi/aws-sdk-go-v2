// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionType string

// Enum values for ActionType
const (
	ActionTypeRestore_cluster       ActionType = "restore-cluster"
	ActionTypeRecommend_node_config ActionType = "recommend-node-config"
	ActionTypeResize_cluster        ActionType = "resize-cluster"
)

// Values returns all known values for ActionType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionType) Values() []ActionType {
	return []ActionType{
		"restore-cluster",
		"recommend-node-config",
		"resize-cluster",
	}
}

type Mode string

// Enum values for Mode
const (
	ModeStandard         Mode = "standard"
	ModeHigh_performance Mode = "high-performance"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"standard",
		"high-performance",
	}
}

type NodeConfigurationOptionsFilterName string

// Enum values for NodeConfigurationOptionsFilterName
const (
	NodeConfigurationOptionsFilterNameNode_type                          NodeConfigurationOptionsFilterName = "NodeType"
	NodeConfigurationOptionsFilterNameNum_nodes                          NodeConfigurationOptionsFilterName = "NumberOfNodes"
	NodeConfigurationOptionsFilterNameEstimated_disk_utilization_percent NodeConfigurationOptionsFilterName = "EstimatedDiskUtilizationPercent"
	NodeConfigurationOptionsFilterNameMode                               NodeConfigurationOptionsFilterName = "Mode"
)

// Values returns all known values for NodeConfigurationOptionsFilterName. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (NodeConfigurationOptionsFilterName) Values() []NodeConfigurationOptionsFilterName {
	return []NodeConfigurationOptionsFilterName{
		"NodeType",
		"NumberOfNodes",
		"EstimatedDiskUtilizationPercent",
		"Mode",
	}
}

type OperatorType string

// Enum values for OperatorType
const (
	OperatorTypeEq      OperatorType = "eq"
	OperatorTypeLt      OperatorType = "lt"
	OperatorTypeGt      OperatorType = "gt"
	OperatorTypeLe      OperatorType = "le"
	OperatorTypeGe      OperatorType = "ge"
	OperatorTypeIn      OperatorType = "in"
	OperatorTypeBetween OperatorType = "between"
)

// Values returns all known values for OperatorType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OperatorType) Values() []OperatorType {
	return []OperatorType{
		"eq",
		"lt",
		"gt",
		"le",
		"ge",
		"in",
		"between",
	}
}

type ParameterApplyType string

// Enum values for ParameterApplyType
const (
	ParameterApplyTypeStatic  ParameterApplyType = "static"
	ParameterApplyTypeDynamic ParameterApplyType = "dynamic"
)

// Values returns all known values for ParameterApplyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ParameterApplyType) Values() []ParameterApplyType {
	return []ParameterApplyType{
		"static",
		"dynamic",
	}
}

type ReservedNodeOfferingType string

// Enum values for ReservedNodeOfferingType
const (
	ReservedNodeOfferingTypeRegular    ReservedNodeOfferingType = "Regular"
	ReservedNodeOfferingTypeUpgradable ReservedNodeOfferingType = "Upgradable"
)

// Values returns all known values for ReservedNodeOfferingType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReservedNodeOfferingType) Values() []ReservedNodeOfferingType {
	return []ReservedNodeOfferingType{
		"Regular",
		"Upgradable",
	}
}

type ScheduledActionFilterName string

// Enum values for ScheduledActionFilterName
const (
	ScheduledActionFilterNameCluster_identifier ScheduledActionFilterName = "cluster-identifier"
	ScheduledActionFilterNameIam_role           ScheduledActionFilterName = "iam-role"
)

// Values returns all known values for ScheduledActionFilterName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledActionFilterName) Values() []ScheduledActionFilterName {
	return []ScheduledActionFilterName{
		"cluster-identifier",
		"iam-role",
	}
}

type ScheduledActionState string

// Enum values for ScheduledActionState
const (
	ScheduledActionStateActive   ScheduledActionState = "ACTIVE"
	ScheduledActionStateDisabled ScheduledActionState = "DISABLED"
)

// Values returns all known values for ScheduledActionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledActionState) Values() []ScheduledActionState {
	return []ScheduledActionState{
		"ACTIVE",
		"DISABLED",
	}
}

type ScheduledActionTypeValues string

// Enum values for ScheduledActionTypeValues
const (
	ScheduledActionTypeValuesResize_cluster ScheduledActionTypeValues = "ResizeCluster"
	ScheduledActionTypeValuesPause_cluster  ScheduledActionTypeValues = "PauseCluster"
	ScheduledActionTypeValuesResume_cluster ScheduledActionTypeValues = "ResumeCluster"
)

// Values returns all known values for ScheduledActionTypeValues. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScheduledActionTypeValues) Values() []ScheduledActionTypeValues {
	return []ScheduledActionTypeValues{
		"ResizeCluster",
		"PauseCluster",
		"ResumeCluster",
	}
}

type ScheduleState string

// Enum values for ScheduleState
const (
	ScheduleStateModifying ScheduleState = "MODIFYING"
	ScheduleStateActive    ScheduleState = "ACTIVE"
	ScheduleStateFailed    ScheduleState = "FAILED"
)

// Values returns all known values for ScheduleState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScheduleState) Values() []ScheduleState {
	return []ScheduleState{
		"MODIFYING",
		"ACTIVE",
		"FAILED",
	}
}

type SnapshotAttributeToSortBy string

// Enum values for SnapshotAttributeToSortBy
const (
	SnapshotAttributeToSortBySource_type SnapshotAttributeToSortBy = "SOURCE_TYPE"
	SnapshotAttributeToSortByTotal_size  SnapshotAttributeToSortBy = "TOTAL_SIZE"
	SnapshotAttributeToSortByCreate_time SnapshotAttributeToSortBy = "CREATE_TIME"
)

// Values returns all known values for SnapshotAttributeToSortBy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SnapshotAttributeToSortBy) Values() []SnapshotAttributeToSortBy {
	return []SnapshotAttributeToSortBy{
		"SOURCE_TYPE",
		"TOTAL_SIZE",
		"CREATE_TIME",
	}
}

type SortByOrder string

// Enum values for SortByOrder
const (
	SortByOrderAscending  SortByOrder = "ASC"
	SortByOrderDescending SortByOrder = "DESC"
)

// Values returns all known values for SortByOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortByOrder) Values() []SortByOrder {
	return []SortByOrder{
		"ASC",
		"DESC",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeCluster               SourceType = "cluster"
	SourceTypeClusterParameterGroup SourceType = "cluster-parameter-group"
	SourceTypeClusterSecurityGroup  SourceType = "cluster-security-group"
	SourceTypeClusterSnapshot       SourceType = "cluster-snapshot"
	SourceTypeScheduledAction       SourceType = "scheduled-action"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"cluster",
		"cluster-parameter-group",
		"cluster-security-group",
		"cluster-snapshot",
		"scheduled-action",
	}
}

type TableRestoreStatusType string

// Enum values for TableRestoreStatusType
const (
	TableRestoreStatusTypePending     TableRestoreStatusType = "PENDING"
	TableRestoreStatusTypeIn_progress TableRestoreStatusType = "IN_PROGRESS"
	TableRestoreStatusTypeSucceeded   TableRestoreStatusType = "SUCCEEDED"
	TableRestoreStatusTypeFailed      TableRestoreStatusType = "FAILED"
	TableRestoreStatusTypeCanceled    TableRestoreStatusType = "CANCELED"
)

// Values returns all known values for TableRestoreStatusType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TableRestoreStatusType) Values() []TableRestoreStatusType {
	return []TableRestoreStatusType{
		"PENDING",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
	}
}

type UsageLimitBreachAction string

// Enum values for UsageLimitBreachAction
const (
	UsageLimitBreachActionLog         UsageLimitBreachAction = "log"
	UsageLimitBreachActionEmit_metric UsageLimitBreachAction = "emit-metric"
	UsageLimitBreachActionDisable     UsageLimitBreachAction = "disable"
)

// Values returns all known values for UsageLimitBreachAction. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitBreachAction) Values() []UsageLimitBreachAction {
	return []UsageLimitBreachAction{
		"log",
		"emit-metric",
		"disable",
	}
}

type UsageLimitFeatureType string

// Enum values for UsageLimitFeatureType
const (
	UsageLimitFeatureTypeSpectrum            UsageLimitFeatureType = "spectrum"
	UsageLimitFeatureTypeConcurrency_scaling UsageLimitFeatureType = "concurrency-scaling"
)

// Values returns all known values for UsageLimitFeatureType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitFeatureType) Values() []UsageLimitFeatureType {
	return []UsageLimitFeatureType{
		"spectrum",
		"concurrency-scaling",
	}
}

type UsageLimitLimitType string

// Enum values for UsageLimitLimitType
const (
	UsageLimitLimitTypeTime         UsageLimitLimitType = "time"
	UsageLimitLimitTypeData_scanned UsageLimitLimitType = "data-scanned"
)

// Values returns all known values for UsageLimitLimitType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitLimitType) Values() []UsageLimitLimitType {
	return []UsageLimitLimitType{
		"time",
		"data-scanned",
	}
}

type UsageLimitPeriod string

// Enum values for UsageLimitPeriod
const (
	UsageLimitPeriodDaily   UsageLimitPeriod = "daily"
	UsageLimitPeriodWeekly  UsageLimitPeriod = "weekly"
	UsageLimitPeriodMonthly UsageLimitPeriod = "monthly"
)

// Values returns all known values for UsageLimitPeriod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageLimitPeriod) Values() []UsageLimitPeriod {
	return []UsageLimitPeriod{
		"daily",
		"weekly",
		"monthly",
	}
}
