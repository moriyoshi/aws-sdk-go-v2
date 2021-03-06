// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type InventoryFilterCondition string

// Enum values for InventoryFilterCondition
const (
	InventoryFilterConditionEquals      InventoryFilterCondition = "EQUALS"
	InventoryFilterConditionNot_equals  InventoryFilterCondition = "NOT_EQUALS"
	InventoryFilterConditionBegins_with InventoryFilterCondition = "BEGINS_WITH"
	InventoryFilterConditionContains    InventoryFilterCondition = "CONTAINS"
)

// Values returns all known values for InventoryFilterCondition. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InventoryFilterCondition) Values() []InventoryFilterCondition {
	return []InventoryFilterCondition{
		"EQUALS",
		"NOT_EQUALS",
		"BEGINS_WITH",
		"CONTAINS",
	}
}

type LicenseConfigurationStatus string

// Enum values for LicenseConfigurationStatus
const (
	LicenseConfigurationStatusAvailable LicenseConfigurationStatus = "AVAILABLE"
	LicenseConfigurationStatusDisabled  LicenseConfigurationStatus = "DISABLED"
)

// Values returns all known values for LicenseConfigurationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LicenseConfigurationStatus) Values() []LicenseConfigurationStatus {
	return []LicenseConfigurationStatus{
		"AVAILABLE",
		"DISABLED",
	}
}

type LicenseCountingType string

// Enum values for LicenseCountingType
const (
	LicenseCountingTypeVcpu     LicenseCountingType = "vCPU"
	LicenseCountingTypeInstance LicenseCountingType = "Instance"
	LicenseCountingTypeCore     LicenseCountingType = "Core"
	LicenseCountingTypeSocket   LicenseCountingType = "Socket"
)

// Values returns all known values for LicenseCountingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LicenseCountingType) Values() []LicenseCountingType {
	return []LicenseCountingType{
		"vCPU",
		"Instance",
		"Core",
		"Socket",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeEc2_instance                     ResourceType = "EC2_INSTANCE"
	ResourceTypeEc2_host                         ResourceType = "EC2_HOST"
	ResourceTypeEc2_ami                          ResourceType = "EC2_AMI"
	ResourceTypeRds                              ResourceType = "RDS"
	ResourceTypeSystems_manager_managed_instance ResourceType = "SYSTEMS_MANAGER_MANAGED_INSTANCE"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"EC2_INSTANCE",
		"EC2_HOST",
		"EC2_AMI",
		"RDS",
		"SYSTEMS_MANAGER_MANAGED_INSTANCE",
	}
}
