package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerPeerStats *****
type bmpServerPeerStats struct {
	validation
	obj          *otg.BmpServerPeerStats
	marshaller   marshalBmpServerPeerStats
	unMarshaller unMarshalBmpServerPeerStats
}

func NewBmpServerPeerStats() BmpServerPeerStats {
	obj := bmpServerPeerStats{obj: &otg.BmpServerPeerStats{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerPeerStats) msg() *otg.BmpServerPeerStats {
	return obj.obj
}

func (obj *bmpServerPeerStats) setMsg(msg *otg.BmpServerPeerStats) BmpServerPeerStats {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerPeerStats struct {
	obj *bmpServerPeerStats
}

type marshalBmpServerPeerStats interface {
	// ToProto marshals BmpServerPeerStats to protobuf object *otg.BmpServerPeerStats
	ToProto() (*otg.BmpServerPeerStats, error)
	// ToPbText marshals BmpServerPeerStats to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerPeerStats to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerPeerStats to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerPeerStats struct {
	obj *bmpServerPeerStats
}

type unMarshalBmpServerPeerStats interface {
	// FromProto unmarshals BmpServerPeerStats from protobuf object *otg.BmpServerPeerStats
	FromProto(msg *otg.BmpServerPeerStats) (BmpServerPeerStats, error)
	// FromPbText unmarshals BmpServerPeerStats from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerPeerStats from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerPeerStats from JSON text
	FromJson(value string) error
}

func (obj *bmpServerPeerStats) Marshal() marshalBmpServerPeerStats {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerPeerStats{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerPeerStats) Unmarshal() unMarshalBmpServerPeerStats {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerPeerStats{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerPeerStats) ToProto() (*otg.BmpServerPeerStats, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerPeerStats) FromProto(msg *otg.BmpServerPeerStats) (BmpServerPeerStats, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerPeerStats) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalbmpServerPeerStats) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbmpServerPeerStats) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalbmpServerPeerStats) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbmpServerPeerStats) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalbmpServerPeerStats) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bmpServerPeerStats) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerPeerStats) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerPeerStats) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerPeerStats) Clone() (BmpServerPeerStats, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerPeerStats()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// BmpServerPeerStats is the last set of BMP stats sent by the BMP client for this peer as per the message format defined in https://www.rfc-editor.org/rfc/rfc7854.html#section-4.8. It might be absent if no BMP Statistics message have been received yet or peer is still in down state. If the BMP client has sent a subset of the various statistic counters, only those should be set when returning the result.
type BmpServerPeerStats interface {
	Validation
	// msg marshals BmpServerPeerStats to protobuf object *otg.BmpServerPeerStats
	// and doesn't set defaults
	msg() *otg.BmpServerPeerStats
	// setMsg unmarshals BmpServerPeerStats from protobuf object *otg.BmpServerPeerStats
	// and doesn't set defaults
	setMsg(*otg.BmpServerPeerStats) BmpServerPeerStats
	// provides marshal interface
	Marshal() marshalBmpServerPeerStats
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerPeerStats
	// validate validates BmpServerPeerStats
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerPeerStats, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RejectedPrefixes returns uint32, set in BmpServerPeerStats.
	RejectedPrefixes() uint32
	// SetRejectedPrefixes assigns uint32 provided by user to BmpServerPeerStats
	SetRejectedPrefixes(value uint32) BmpServerPeerStats
	// HasRejectedPrefixes checks if RejectedPrefixes has been set in BmpServerPeerStats
	HasRejectedPrefixes() bool
	// DuplicatePrefixes returns uint32, set in BmpServerPeerStats.
	DuplicatePrefixes() uint32
	// SetDuplicatePrefixes assigns uint32 provided by user to BmpServerPeerStats
	SetDuplicatePrefixes(value uint32) BmpServerPeerStats
	// HasDuplicatePrefixes checks if DuplicatePrefixes has been set in BmpServerPeerStats
	HasDuplicatePrefixes() bool
	// DuplicateWithdraws returns uint32, set in BmpServerPeerStats.
	DuplicateWithdraws() uint32
	// SetDuplicateWithdraws assigns uint32 provided by user to BmpServerPeerStats
	SetDuplicateWithdraws(value uint32) BmpServerPeerStats
	// HasDuplicateWithdraws checks if DuplicateWithdraws has been set in BmpServerPeerStats
	HasDuplicateWithdraws() bool
	// ClusterListInvalidatedUpdates returns uint32, set in BmpServerPeerStats.
	ClusterListInvalidatedUpdates() uint32
	// SetClusterListInvalidatedUpdates assigns uint32 provided by user to BmpServerPeerStats
	SetClusterListInvalidatedUpdates(value uint32) BmpServerPeerStats
	// HasClusterListInvalidatedUpdates checks if ClusterListInvalidatedUpdates has been set in BmpServerPeerStats
	HasClusterListInvalidatedUpdates() bool
	// AsPathInvalidatedUpdates returns uint32, set in BmpServerPeerStats.
	AsPathInvalidatedUpdates() uint32
	// SetAsPathInvalidatedUpdates assigns uint32 provided by user to BmpServerPeerStats
	SetAsPathInvalidatedUpdates(value uint32) BmpServerPeerStats
	// HasAsPathInvalidatedUpdates checks if AsPathInvalidatedUpdates has been set in BmpServerPeerStats
	HasAsPathInvalidatedUpdates() bool
	// OriginatorIdInvalidatedUpdates returns uint32, set in BmpServerPeerStats.
	OriginatorIdInvalidatedUpdates() uint32
	// SetOriginatorIdInvalidatedUpdates assigns uint32 provided by user to BmpServerPeerStats
	SetOriginatorIdInvalidatedUpdates(value uint32) BmpServerPeerStats
	// HasOriginatorIdInvalidatedUpdates checks if OriginatorIdInvalidatedUpdates has been set in BmpServerPeerStats
	HasOriginatorIdInvalidatedUpdates() bool
	// AsConfedInvalidatedUpdates returns uint32, set in BmpServerPeerStats.
	AsConfedInvalidatedUpdates() uint32
	// SetAsConfedInvalidatedUpdates assigns uint32 provided by user to BmpServerPeerStats
	SetAsConfedInvalidatedUpdates(value uint32) BmpServerPeerStats
	// HasAsConfedInvalidatedUpdates checks if AsConfedInvalidatedUpdates has been set in BmpServerPeerStats
	HasAsConfedInvalidatedUpdates() bool
	// AdjRibRoutes returns uint64, set in BmpServerPeerStats.
	AdjRibRoutes() uint64
	// SetAdjRibRoutes assigns uint64 provided by user to BmpServerPeerStats
	SetAdjRibRoutes(value uint64) BmpServerPeerStats
	// HasAdjRibRoutes checks if AdjRibRoutes has been set in BmpServerPeerStats
	HasAdjRibRoutes() bool
	// LocalRibRoutes returns uint64, set in BmpServerPeerStats.
	LocalRibRoutes() uint64
	// SetLocalRibRoutes assigns uint64 provided by user to BmpServerPeerStats
	SetLocalRibRoutes(value uint64) BmpServerPeerStats
	// HasLocalRibRoutes checks if LocalRibRoutes has been set in BmpServerPeerStats
	HasLocalRibRoutes() bool
	// WithdrawUpdatesReceived returns uint32, set in BmpServerPeerStats.
	WithdrawUpdatesReceived() uint32
	// SetWithdrawUpdatesReceived assigns uint32 provided by user to BmpServerPeerStats
	SetWithdrawUpdatesReceived(value uint32) BmpServerPeerStats
	// HasWithdrawUpdatesReceived checks if WithdrawUpdatesReceived has been set in BmpServerPeerStats
	HasWithdrawUpdatesReceived() bool
	// WithdrawPrefixesReceived returns uint32, set in BmpServerPeerStats.
	WithdrawPrefixesReceived() uint32
	// SetWithdrawPrefixesReceived assigns uint32 provided by user to BmpServerPeerStats
	SetWithdrawPrefixesReceived(value uint32) BmpServerPeerStats
	// HasWithdrawPrefixesReceived checks if WithdrawPrefixesReceived has been set in BmpServerPeerStats
	HasWithdrawPrefixesReceived() bool
	// DuplicateUpdates returns uint32, set in BmpServerPeerStats.
	DuplicateUpdates() uint32
	// SetDuplicateUpdates assigns uint32 provided by user to BmpServerPeerStats
	SetDuplicateUpdates(value uint32) BmpServerPeerStats
	// HasDuplicateUpdates checks if DuplicateUpdates has been set in BmpServerPeerStats
	HasDuplicateUpdates() bool
}

// Number of prefixes rejected by inbound policy.
// RejectedPrefixes returns a uint32
func (obj *bmpServerPeerStats) RejectedPrefixes() uint32 {

	return *obj.obj.RejectedPrefixes

}

// Number of prefixes rejected by inbound policy.
// RejectedPrefixes returns a uint32
func (obj *bmpServerPeerStats) HasRejectedPrefixes() bool {
	return obj.obj.RejectedPrefixes != nil
}

// Number of prefixes rejected by inbound policy.
// SetRejectedPrefixes sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetRejectedPrefixes(value uint32) BmpServerPeerStats {

	obj.obj.RejectedPrefixes = &value
	return obj
}

// Number of (known) duplicate prefix advertisements.
// DuplicatePrefixes returns a uint32
func (obj *bmpServerPeerStats) DuplicatePrefixes() uint32 {

	return *obj.obj.DuplicatePrefixes

}

// Number of (known) duplicate prefix advertisements.
// DuplicatePrefixes returns a uint32
func (obj *bmpServerPeerStats) HasDuplicatePrefixes() bool {
	return obj.obj.DuplicatePrefixes != nil
}

// Number of (known) duplicate prefix advertisements.
// SetDuplicatePrefixes sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetDuplicatePrefixes(value uint32) BmpServerPeerStats {

	obj.obj.DuplicatePrefixes = &value
	return obj
}

// Number of (known) duplicate withdraws.
// DuplicateWithdraws returns a uint32
func (obj *bmpServerPeerStats) DuplicateWithdraws() uint32 {

	return *obj.obj.DuplicateWithdraws

}

// Number of (known) duplicate withdraws.
// DuplicateWithdraws returns a uint32
func (obj *bmpServerPeerStats) HasDuplicateWithdraws() bool {
	return obj.obj.DuplicateWithdraws != nil
}

// Number of (known) duplicate withdraws.
// SetDuplicateWithdraws sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetDuplicateWithdraws(value uint32) BmpServerPeerStats {

	obj.obj.DuplicateWithdraws = &value
	return obj
}

// Number of updates invalidated due to CLUSTER_LIST loop.
// ClusterListInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) ClusterListInvalidatedUpdates() uint32 {

	return *obj.obj.ClusterListInvalidatedUpdates

}

// Number of updates invalidated due to CLUSTER_LIST loop.
// ClusterListInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) HasClusterListInvalidatedUpdates() bool {
	return obj.obj.ClusterListInvalidatedUpdates != nil
}

// Number of updates invalidated due to CLUSTER_LIST loop.
// SetClusterListInvalidatedUpdates sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetClusterListInvalidatedUpdates(value uint32) BmpServerPeerStats {

	obj.obj.ClusterListInvalidatedUpdates = &value
	return obj
}

// Number of updates invalidated due to AS_PATH loop.
// AsPathInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) AsPathInvalidatedUpdates() uint32 {

	return *obj.obj.AsPathInvalidatedUpdates

}

// Number of updates invalidated due to AS_PATH loop.
// AsPathInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) HasAsPathInvalidatedUpdates() bool {
	return obj.obj.AsPathInvalidatedUpdates != nil
}

// Number of updates invalidated due to AS_PATH loop.
// SetAsPathInvalidatedUpdates sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetAsPathInvalidatedUpdates(value uint32) BmpServerPeerStats {

	obj.obj.AsPathInvalidatedUpdates = &value
	return obj
}

// Number of updates invalidated due to ORIGINATOR_ID.
// OriginatorIdInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) OriginatorIdInvalidatedUpdates() uint32 {

	return *obj.obj.OriginatorIdInvalidatedUpdates

}

// Number of updates invalidated due to ORIGINATOR_ID.
// OriginatorIdInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) HasOriginatorIdInvalidatedUpdates() bool {
	return obj.obj.OriginatorIdInvalidatedUpdates != nil
}

// Number of updates invalidated due to ORIGINATOR_ID.
// SetOriginatorIdInvalidatedUpdates sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetOriginatorIdInvalidatedUpdates(value uint32) BmpServerPeerStats {

	obj.obj.OriginatorIdInvalidatedUpdates = &value
	return obj
}

// Number of updates invalidated due to AS_CONFED loop.
// AsConfedInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) AsConfedInvalidatedUpdates() uint32 {

	return *obj.obj.AsConfedInvalidatedUpdates

}

// Number of updates invalidated due to AS_CONFED loop.
// AsConfedInvalidatedUpdates returns a uint32
func (obj *bmpServerPeerStats) HasAsConfedInvalidatedUpdates() bool {
	return obj.obj.AsConfedInvalidatedUpdates != nil
}

// Number of updates invalidated due to AS_CONFED loop.
// SetAsConfedInvalidatedUpdates sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetAsConfedInvalidatedUpdates(value uint32) BmpServerPeerStats {

	obj.obj.AsConfedInvalidatedUpdates = &value
	return obj
}

// Number of routes in Adj-RIBs-In.
// AdjRibRoutes returns a uint64
func (obj *bmpServerPeerStats) AdjRibRoutes() uint64 {

	return *obj.obj.AdjRibRoutes

}

// Number of routes in Adj-RIBs-In.
// AdjRibRoutes returns a uint64
func (obj *bmpServerPeerStats) HasAdjRibRoutes() bool {
	return obj.obj.AdjRibRoutes != nil
}

// Number of routes in Adj-RIBs-In.
// SetAdjRibRoutes sets the uint64 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetAdjRibRoutes(value uint64) BmpServerPeerStats {

	obj.obj.AdjRibRoutes = &value
	return obj
}

// Number of routes in Loc-RIB.
// LocalRibRoutes returns a uint64
func (obj *bmpServerPeerStats) LocalRibRoutes() uint64 {

	return *obj.obj.LocalRibRoutes

}

// Number of routes in Loc-RIB.
// LocalRibRoutes returns a uint64
func (obj *bmpServerPeerStats) HasLocalRibRoutes() bool {
	return obj.obj.LocalRibRoutes != nil
}

// Number of routes in Loc-RIB.
// SetLocalRibRoutes sets the uint64 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetLocalRibRoutes(value uint64) BmpServerPeerStats {

	obj.obj.LocalRibRoutes = &value
	return obj
}

// Number of updates subjected to treat-as-withdraw treatment [RFC7606].
// WithdrawUpdatesReceived returns a uint32
func (obj *bmpServerPeerStats) WithdrawUpdatesReceived() uint32 {

	return *obj.obj.WithdrawUpdatesReceived

}

// Number of updates subjected to treat-as-withdraw treatment [RFC7606].
// WithdrawUpdatesReceived returns a uint32
func (obj *bmpServerPeerStats) HasWithdrawUpdatesReceived() bool {
	return obj.obj.WithdrawUpdatesReceived != nil
}

// Number of updates subjected to treat-as-withdraw treatment [RFC7606].
// SetWithdrawUpdatesReceived sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetWithdrawUpdatesReceived(value uint32) BmpServerPeerStats {

	obj.obj.WithdrawUpdatesReceived = &value
	return obj
}

// Number of prefixes  subjected to treat-as-withdraw treatment [RFC7606].
// WithdrawPrefixesReceived returns a uint32
func (obj *bmpServerPeerStats) WithdrawPrefixesReceived() uint32 {

	return *obj.obj.WithdrawPrefixesReceived

}

// Number of prefixes  subjected to treat-as-withdraw treatment [RFC7606].
// WithdrawPrefixesReceived returns a uint32
func (obj *bmpServerPeerStats) HasWithdrawPrefixesReceived() bool {
	return obj.obj.WithdrawPrefixesReceived != nil
}

// Number of prefixes  subjected to treat-as-withdraw treatment [RFC7606].
// SetWithdrawPrefixesReceived sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetWithdrawPrefixesReceived(value uint32) BmpServerPeerStats {

	obj.obj.WithdrawPrefixesReceived = &value
	return obj
}

// Number of duplicate update messages received.
// DuplicateUpdates returns a uint32
func (obj *bmpServerPeerStats) DuplicateUpdates() uint32 {

	return *obj.obj.DuplicateUpdates

}

// Number of duplicate update messages received.
// DuplicateUpdates returns a uint32
func (obj *bmpServerPeerStats) HasDuplicateUpdates() bool {
	return obj.obj.DuplicateUpdates != nil
}

// Number of duplicate update messages received.
// SetDuplicateUpdates sets the uint32 value in the BmpServerPeerStats object
func (obj *bmpServerPeerStats) SetDuplicateUpdates(value uint32) BmpServerPeerStats {

	obj.obj.DuplicateUpdates = &value
	return obj
}

func (obj *bmpServerPeerStats) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bmpServerPeerStats) setDefault() {

}
