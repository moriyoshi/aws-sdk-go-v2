// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BackupPolicy string

// Enum values for BackupPolicy
const (
	BackupPolicyDefault BackupPolicy = "DEFAULT"
)

// Values returns all known values for BackupPolicy. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BackupPolicy) Values() []BackupPolicy {
	return []BackupPolicy{
		"DEFAULT",
	}
}

type BackupState string

// Enum values for BackupState
const (
	BackupStateCreate_in_progress BackupState = "CREATE_IN_PROGRESS"
	BackupStateReady              BackupState = "READY"
	BackupStateDeleted            BackupState = "DELETED"
	BackupStatePending_deletion   BackupState = "PENDING_DELETION"
)

// Values returns all known values for BackupState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BackupState) Values() []BackupState {
	return []BackupState{
		"CREATE_IN_PROGRESS",
		"READY",
		"DELETED",
		"PENDING_DELETION",
	}
}

type ClusterState string

// Enum values for ClusterState
const (
	ClusterStateCreate_in_progress     ClusterState = "CREATE_IN_PROGRESS"
	ClusterStateUninitialized          ClusterState = "UNINITIALIZED"
	ClusterStateInitialize_in_progress ClusterState = "INITIALIZE_IN_PROGRESS"
	ClusterStateInitialized            ClusterState = "INITIALIZED"
	ClusterStateActive                 ClusterState = "ACTIVE"
	ClusterStateUpdate_in_progress     ClusterState = "UPDATE_IN_PROGRESS"
	ClusterStateDelete_in_progress     ClusterState = "DELETE_IN_PROGRESS"
	ClusterStateDeleted                ClusterState = "DELETED"
	ClusterStateDegraded               ClusterState = "DEGRADED"
)

// Values returns all known values for ClusterState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ClusterState) Values() []ClusterState {
	return []ClusterState{
		"CREATE_IN_PROGRESS",
		"UNINITIALIZED",
		"INITIALIZE_IN_PROGRESS",
		"INITIALIZED",
		"ACTIVE",
		"UPDATE_IN_PROGRESS",
		"DELETE_IN_PROGRESS",
		"DELETED",
		"DEGRADED",
	}
}

type HsmState string

// Enum values for HsmState
const (
	HsmStateCreate_in_progress HsmState = "CREATE_IN_PROGRESS"
	HsmStateActive             HsmState = "ACTIVE"
	HsmStateDegraded           HsmState = "DEGRADED"
	HsmStateDelete_in_progress HsmState = "DELETE_IN_PROGRESS"
	HsmStateDeleted            HsmState = "DELETED"
)

// Values returns all known values for HsmState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (HsmState) Values() []HsmState {
	return []HsmState{
		"CREATE_IN_PROGRESS",
		"ACTIVE",
		"DEGRADED",
		"DELETE_IN_PROGRESS",
		"DELETED",
	}
}
