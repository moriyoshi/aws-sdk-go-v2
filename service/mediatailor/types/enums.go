// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Mode string

// Enum values for Mode
const (
	ModeOff              Mode = "OFF"
	ModeBehind_live_edge Mode = "BEHIND_LIVE_EDGE"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"OFF",
		"BEHIND_LIVE_EDGE",
	}
}

type OriginManifestType string

// Enum values for OriginManifestType
const (
	OriginManifestTypeSingle_period OriginManifestType = "SINGLE_PERIOD"
	OriginManifestTypeMulti_period  OriginManifestType = "MULTI_PERIOD"
)

// Values returns all known values for OriginManifestType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginManifestType) Values() []OriginManifestType {
	return []OriginManifestType{
		"SINGLE_PERIOD",
		"MULTI_PERIOD",
	}
}
