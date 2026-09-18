package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaLinkTrafficEngineering *****
type ospfv2LsaLinkTrafficEngineering struct {
	validation
	obj           *otg.Ospfv2LsaLinkTrafficEngineering
	marshaller    marshalOspfv2LsaLinkTrafficEngineering
	unMarshaller  unMarshalOspfv2LsaLinkTrafficEngineering
	linkMsdHolder Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
}

func NewOspfv2LsaLinkTrafficEngineering() Ospfv2LsaLinkTrafficEngineering {
	obj := ospfv2LsaLinkTrafficEngineering{obj: &otg.Ospfv2LsaLinkTrafficEngineering{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaLinkTrafficEngineering) msg() *otg.Ospfv2LsaLinkTrafficEngineering {
	return obj.obj
}

func (obj *ospfv2LsaLinkTrafficEngineering) setMsg(msg *otg.Ospfv2LsaLinkTrafficEngineering) Ospfv2LsaLinkTrafficEngineering {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaLinkTrafficEngineering struct {
	obj *ospfv2LsaLinkTrafficEngineering
}

type marshalOspfv2LsaLinkTrafficEngineering interface {
	// ToProto marshals Ospfv2LsaLinkTrafficEngineering to protobuf object *otg.Ospfv2LsaLinkTrafficEngineering
	ToProto() (*otg.Ospfv2LsaLinkTrafficEngineering, error)
	// ToPbText marshals Ospfv2LsaLinkTrafficEngineering to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaLinkTrafficEngineering to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaLinkTrafficEngineering to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaLinkTrafficEngineering struct {
	obj *ospfv2LsaLinkTrafficEngineering
}

type unMarshalOspfv2LsaLinkTrafficEngineering interface {
	// FromProto unmarshals Ospfv2LsaLinkTrafficEngineering from protobuf object *otg.Ospfv2LsaLinkTrafficEngineering
	FromProto(msg *otg.Ospfv2LsaLinkTrafficEngineering) (Ospfv2LsaLinkTrafficEngineering, error)
	// FromPbText unmarshals Ospfv2LsaLinkTrafficEngineering from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaLinkTrafficEngineering from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaLinkTrafficEngineering from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaLinkTrafficEngineering) Marshal() marshalOspfv2LsaLinkTrafficEngineering {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaLinkTrafficEngineering{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaLinkTrafficEngineering) Unmarshal() unMarshalOspfv2LsaLinkTrafficEngineering {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaLinkTrafficEngineering{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaLinkTrafficEngineering) ToProto() (*otg.Ospfv2LsaLinkTrafficEngineering, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaLinkTrafficEngineering) FromProto(msg *otg.Ospfv2LsaLinkTrafficEngineering) (Ospfv2LsaLinkTrafficEngineering, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaLinkTrafficEngineering) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaLinkTrafficEngineering) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalospfv2LsaLinkTrafficEngineering) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaLinkTrafficEngineering) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalospfv2LsaLinkTrafficEngineering) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaLinkTrafficEngineering) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ospfv2LsaLinkTrafficEngineering) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaLinkTrafficEngineering) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaLinkTrafficEngineering) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaLinkTrafficEngineering) Clone() (Ospfv2LsaLinkTrafficEngineering, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaLinkTrafficEngineering()
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

func (obj *ospfv2LsaLinkTrafficEngineering) setNil() {
	obj.linkMsdHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2LsaLinkTrafficEngineering is traffic engineering attributes for a link, decoded from the Link TLV sub-TLVs of the
// Traffic Engineering Opaque LSA (RFC 3630 Section 2.5) and the corresponding sub-TLVs of
// the Extended Link TLV of the OSPFv2 Extended Link Opaque LSA (RFC 9492).
type Ospfv2LsaLinkTrafficEngineering interface {
	Validation
	// msg marshals Ospfv2LsaLinkTrafficEngineering to protobuf object *otg.Ospfv2LsaLinkTrafficEngineering
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaLinkTrafficEngineering
	// setMsg unmarshals Ospfv2LsaLinkTrafficEngineering from protobuf object *otg.Ospfv2LsaLinkTrafficEngineering
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaLinkTrafficEngineering) Ospfv2LsaLinkTrafficEngineering
	// provides marshal interface
	Marshal() marshalOspfv2LsaLinkTrafficEngineering
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaLinkTrafficEngineering
	// validate validates Ospfv2LsaLinkTrafficEngineering
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaLinkTrafficEngineering, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LinkType returns Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum, set in Ospfv2LsaLinkTrafficEngineering
	LinkType() Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum
	// SetLinkType assigns Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum provided by user to Ospfv2LsaLinkTrafficEngineering
	SetLinkType(value Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum) Ospfv2LsaLinkTrafficEngineering
	// HasLinkType checks if LinkType has been set in Ospfv2LsaLinkTrafficEngineering
	HasLinkType() bool
	// LocalInterfaceIpAddresses returns []string, set in Ospfv2LsaLinkTrafficEngineering.
	LocalInterfaceIpAddresses() []string
	// SetLocalInterfaceIpAddresses assigns []string provided by user to Ospfv2LsaLinkTrafficEngineering
	SetLocalInterfaceIpAddresses(value []string) Ospfv2LsaLinkTrafficEngineering
	// RemoteInterfaceIpAddresses returns []string, set in Ospfv2LsaLinkTrafficEngineering.
	RemoteInterfaceIpAddresses() []string
	// SetRemoteInterfaceIpAddresses assigns []string provided by user to Ospfv2LsaLinkTrafficEngineering
	SetRemoteInterfaceIpAddresses(value []string) Ospfv2LsaLinkTrafficEngineering
	// TeMetric returns uint32, set in Ospfv2LsaLinkTrafficEngineering.
	TeMetric() uint32
	// SetTeMetric assigns uint32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetTeMetric(value uint32) Ospfv2LsaLinkTrafficEngineering
	// HasTeMetric checks if TeMetric has been set in Ospfv2LsaLinkTrafficEngineering
	HasTeMetric() bool
	// MaximumBandwidth returns float32, set in Ospfv2LsaLinkTrafficEngineering.
	MaximumBandwidth() float32
	// SetMaximumBandwidth assigns float32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetMaximumBandwidth(value float32) Ospfv2LsaLinkTrafficEngineering
	// HasMaximumBandwidth checks if MaximumBandwidth has been set in Ospfv2LsaLinkTrafficEngineering
	HasMaximumBandwidth() bool
	// MaximumReservableBandwidth returns float32, set in Ospfv2LsaLinkTrafficEngineering.
	MaximumReservableBandwidth() float32
	// SetMaximumReservableBandwidth assigns float32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetMaximumReservableBandwidth(value float32) Ospfv2LsaLinkTrafficEngineering
	// HasMaximumReservableBandwidth checks if MaximumReservableBandwidth has been set in Ospfv2LsaLinkTrafficEngineering
	HasMaximumReservableBandwidth() bool
	// UnreservedBandwidths returns []float32, set in Ospfv2LsaLinkTrafficEngineering.
	UnreservedBandwidths() []float32
	// SetUnreservedBandwidths assigns []float32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetUnreservedBandwidths(value []float32) Ospfv2LsaLinkTrafficEngineering
	// AdministrativeGroup returns uint32, set in Ospfv2LsaLinkTrafficEngineering.
	AdministrativeGroup() uint32
	// SetAdministrativeGroup assigns uint32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetAdministrativeGroup(value uint32) Ospfv2LsaLinkTrafficEngineering
	// HasAdministrativeGroup checks if AdministrativeGroup has been set in Ospfv2LsaLinkTrafficEngineering
	HasAdministrativeGroup() bool
	// ExtendedAdministrativeGroup returns []uint32, set in Ospfv2LsaLinkTrafficEngineering.
	ExtendedAdministrativeGroup() []uint32
	// SetExtendedAdministrativeGroup assigns []uint32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetExtendedAdministrativeGroup(value []uint32) Ospfv2LsaLinkTrafficEngineering
	// Srlg returns []uint32, set in Ospfv2LsaLinkTrafficEngineering.
	Srlg() []uint32
	// SetSrlg assigns []uint32 provided by user to Ospfv2LsaLinkTrafficEngineering
	SetSrlg(value []uint32) Ospfv2LsaLinkTrafficEngineering
	// LinkMsd returns Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIterIter, set in Ospfv2LsaLinkTrafficEngineering
	LinkMsd() Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
	setNil()
}

type Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum string

// Enum of LinkType on Ospfv2LsaLinkTrafficEngineering
var Ospfv2LsaLinkTrafficEngineeringLinkType = struct {
	POINT_TO_POINT Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum
	MULTI_ACCESS   Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum
}{
	POINT_TO_POINT: Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum("point_to_point"),
	MULTI_ACCESS:   Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum("multi_access"),
}

func (obj *ospfv2LsaLinkTrafficEngineering) LinkType() Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum {
	return Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum(obj.obj.LinkType.Enum().String())
}

// The Link Type sub-TLV, sub-type 1 (RFC 3630 Section 2.5.1).
// LinkType returns a string
func (obj *ospfv2LsaLinkTrafficEngineering) HasLinkType() bool {
	return obj.obj.LinkType != nil
}

func (obj *ospfv2LsaLinkTrafficEngineering) SetLinkType(value Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum) Ospfv2LsaLinkTrafficEngineering {
	intValue, ok := otg.Ospfv2LsaLinkTrafficEngineering_LinkType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2LsaLinkTrafficEngineeringLinkTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2LsaLinkTrafficEngineering_LinkType_Enum(intValue)
	obj.obj.LinkType = &enumValue

	return obj
}

// The Local Interface IP Address sub-TLV, sub-type 3 (RFC 3630 Section 2.5.3).
// LocalInterfaceIpAddresses returns a []string
func (obj *ospfv2LsaLinkTrafficEngineering) LocalInterfaceIpAddresses() []string {
	if obj.obj.LocalInterfaceIpAddresses == nil {
		obj.obj.LocalInterfaceIpAddresses = make([]string, 0)
	}
	return obj.obj.LocalInterfaceIpAddresses
}

// The Local Interface IP Address sub-TLV, sub-type 3 (RFC 3630 Section 2.5.3).
// SetLocalInterfaceIpAddresses sets the []string value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetLocalInterfaceIpAddresses(value []string) Ospfv2LsaLinkTrafficEngineering {

	if obj.obj.LocalInterfaceIpAddresses == nil {
		obj.obj.LocalInterfaceIpAddresses = make([]string, 0)
	}
	obj.obj.LocalInterfaceIpAddresses = value

	return obj
}

// The Remote Interface IP Address sub-TLV, sub-type 4 (RFC 3630 Section 2.5.4).
// RemoteInterfaceIpAddresses returns a []string
func (obj *ospfv2LsaLinkTrafficEngineering) RemoteInterfaceIpAddresses() []string {
	if obj.obj.RemoteInterfaceIpAddresses == nil {
		obj.obj.RemoteInterfaceIpAddresses = make([]string, 0)
	}
	return obj.obj.RemoteInterfaceIpAddresses
}

// The Remote Interface IP Address sub-TLV, sub-type 4 (RFC 3630 Section 2.5.4).
// SetRemoteInterfaceIpAddresses sets the []string value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetRemoteInterfaceIpAddresses(value []string) Ospfv2LsaLinkTrafficEngineering {

	if obj.obj.RemoteInterfaceIpAddresses == nil {
		obj.obj.RemoteInterfaceIpAddresses = make([]string, 0)
	}
	obj.obj.RemoteInterfaceIpAddresses = value

	return obj
}

// The Traffic Engineering Metric sub-TLV, sub-type 5 (RFC 3630 Section 2.5.5).
// TeMetric returns a uint32
func (obj *ospfv2LsaLinkTrafficEngineering) TeMetric() uint32 {

	return *obj.obj.TeMetric

}

// The Traffic Engineering Metric sub-TLV, sub-type 5 (RFC 3630 Section 2.5.5).
// TeMetric returns a uint32
func (obj *ospfv2LsaLinkTrafficEngineering) HasTeMetric() bool {
	return obj.obj.TeMetric != nil
}

// The Traffic Engineering Metric sub-TLV, sub-type 5 (RFC 3630 Section 2.5.5).
// SetTeMetric sets the uint32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetTeMetric(value uint32) Ospfv2LsaLinkTrafficEngineering {

	obj.obj.TeMetric = &value
	return obj
}

// The Maximum Bandwidth sub-TLV, sub-type 6, in bytes per second
// (RFC 3630 Section 2.5.6).
// MaximumBandwidth returns a float32
func (obj *ospfv2LsaLinkTrafficEngineering) MaximumBandwidth() float32 {

	return *obj.obj.MaximumBandwidth

}

// The Maximum Bandwidth sub-TLV, sub-type 6, in bytes per second
// (RFC 3630 Section 2.5.6).
// MaximumBandwidth returns a float32
func (obj *ospfv2LsaLinkTrafficEngineering) HasMaximumBandwidth() bool {
	return obj.obj.MaximumBandwidth != nil
}

// The Maximum Bandwidth sub-TLV, sub-type 6, in bytes per second
// (RFC 3630 Section 2.5.6).
// SetMaximumBandwidth sets the float32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetMaximumBandwidth(value float32) Ospfv2LsaLinkTrafficEngineering {

	obj.obj.MaximumBandwidth = &value
	return obj
}

// The Maximum Reservable Bandwidth sub-TLV, sub-type 7, in bytes per second
// (RFC 3630 Section 2.5.7).
// MaximumReservableBandwidth returns a float32
func (obj *ospfv2LsaLinkTrafficEngineering) MaximumReservableBandwidth() float32 {

	return *obj.obj.MaximumReservableBandwidth

}

// The Maximum Reservable Bandwidth sub-TLV, sub-type 7, in bytes per second
// (RFC 3630 Section 2.5.7).
// MaximumReservableBandwidth returns a float32
func (obj *ospfv2LsaLinkTrafficEngineering) HasMaximumReservableBandwidth() bool {
	return obj.obj.MaximumReservableBandwidth != nil
}

// The Maximum Reservable Bandwidth sub-TLV, sub-type 7, in bytes per second
// (RFC 3630 Section 2.5.7).
// SetMaximumReservableBandwidth sets the float32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetMaximumReservableBandwidth(value float32) Ospfv2LsaLinkTrafficEngineering {

	obj.obj.MaximumReservableBandwidth = &value
	return obj
}

// The Unreserved Bandwidth sub-TLV, sub-type 8: bandwidth reservable at each of the
// eight priority levels 0-7, in bytes per second (RFC 3630 Section 2.5.8).
// UnreservedBandwidths returns a []float32
func (obj *ospfv2LsaLinkTrafficEngineering) UnreservedBandwidths() []float32 {
	if obj.obj.UnreservedBandwidths == nil {
		obj.obj.UnreservedBandwidths = make([]float32, 0)
	}
	return obj.obj.UnreservedBandwidths
}

// The Unreserved Bandwidth sub-TLV, sub-type 8: bandwidth reservable at each of the
// eight priority levels 0-7, in bytes per second (RFC 3630 Section 2.5.8).
// SetUnreservedBandwidths sets the []float32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetUnreservedBandwidths(value []float32) Ospfv2LsaLinkTrafficEngineering {

	if obj.obj.UnreservedBandwidths == nil {
		obj.obj.UnreservedBandwidths = make([]float32, 0)
	}
	obj.obj.UnreservedBandwidths = value

	return obj
}

// The Administrative Group sub-TLV bitmask, sourced from either the TE Link TLV
// sub-type 9 (RFC 3630 Section 2.5.9) or the OSPFv2 Extended Link TLV sub-type 19
// (RFC 9492 Section 6.2).
// AdministrativeGroup returns a uint32
func (obj *ospfv2LsaLinkTrafficEngineering) AdministrativeGroup() uint32 {

	return *obj.obj.AdministrativeGroup

}

// The Administrative Group sub-TLV bitmask, sourced from either the TE Link TLV
// sub-type 9 (RFC 3630 Section 2.5.9) or the OSPFv2 Extended Link TLV sub-type 19
// (RFC 9492 Section 6.2).
// AdministrativeGroup returns a uint32
func (obj *ospfv2LsaLinkTrafficEngineering) HasAdministrativeGroup() bool {
	return obj.obj.AdministrativeGroup != nil
}

// The Administrative Group sub-TLV bitmask, sourced from either the TE Link TLV
// sub-type 9 (RFC 3630 Section 2.5.9) or the OSPFv2 Extended Link TLV sub-type 19
// (RFC 9492 Section 6.2).
// SetAdministrativeGroup sets the uint32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetAdministrativeGroup(value uint32) Ospfv2LsaLinkTrafficEngineering {

	obj.obj.AdministrativeGroup = &value
	return obj
}

// The Extended Administrative Group, sourced from either the TE Link TLV sub-type 26
// (RFC 7308) or the OSPFv2 Extended Link TLV sub-type 20 (RFC 9492 Section 6.3), as one
// or more additional 32-bit administrative-group words beyond administrative_group.
// ExtendedAdministrativeGroup returns a []uint32
func (obj *ospfv2LsaLinkTrafficEngineering) ExtendedAdministrativeGroup() []uint32 {
	if obj.obj.ExtendedAdministrativeGroup == nil {
		obj.obj.ExtendedAdministrativeGroup = make([]uint32, 0)
	}
	return obj.obj.ExtendedAdministrativeGroup
}

// The Extended Administrative Group, sourced from either the TE Link TLV sub-type 26
// (RFC 7308) or the OSPFv2 Extended Link TLV sub-type 20 (RFC 9492 Section 6.3), as one
// or more additional 32-bit administrative-group words beyond administrative_group.
// SetExtendedAdministrativeGroup sets the []uint32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetExtendedAdministrativeGroup(value []uint32) Ospfv2LsaLinkTrafficEngineering {

	if obj.obj.ExtendedAdministrativeGroup == nil {
		obj.obj.ExtendedAdministrativeGroup = make([]uint32, 0)
	}
	obj.obj.ExtendedAdministrativeGroup = value

	return obj
}

// The Shared Risk Link Group (SRLG) membership of this link, sourced from either the TE
// Link TLV sub-type 16 (RFC 4203 Section 1.3) or the OSPFv2 Extended Link TLV sub-type 11
// (RFC 9492 Section 6.1).
// Srlg returns a []uint32
func (obj *ospfv2LsaLinkTrafficEngineering) Srlg() []uint32 {
	if obj.obj.Srlg == nil {
		obj.obj.Srlg = make([]uint32, 0)
	}
	return obj.obj.Srlg
}

// The Shared Risk Link Group (SRLG) membership of this link, sourced from either the TE
// Link TLV sub-type 16 (RFC 4203 Section 1.3) or the OSPFv2 Extended Link TLV sub-type 11
// (RFC 9492 Section 6.1).
// SetSrlg sets the []uint32 value in the Ospfv2LsaLinkTrafficEngineering object
func (obj *ospfv2LsaLinkTrafficEngineering) SetSrlg(value []uint32) Ospfv2LsaLinkTrafficEngineering {

	if obj.obj.Srlg == nil {
		obj.obj.Srlg = make([]uint32, 0)
	}
	obj.obj.Srlg = value

	return obj
}

// One or more Maximum SID Depth (MSD) values for this link, decoded from the Link MSD sub-TLV of the OSPFv2 Extended Link TLV, sub-type 6 (RFC 8476 Section 3).
// LinkMsd returns a []Ospfv2LsaMsd
func (obj *ospfv2LsaLinkTrafficEngineering) LinkMsd() Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	if len(obj.obj.LinkMsd) == 0 {
		obj.obj.LinkMsd = []*otg.Ospfv2LsaMsd{}
	}
	if obj.linkMsdHolder == nil {
		obj.linkMsdHolder = newOspfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter(&obj.obj.LinkMsd).setMsg(obj)
	}
	return obj.linkMsdHolder
}

type ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter struct {
	obj               *ospfv2LsaLinkTrafficEngineering
	ospfv2LsaMsdSlice []Ospfv2LsaMsd
	fieldPtr          *[]*otg.Ospfv2LsaMsd
}

func newOspfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter(ptr *[]*otg.Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	return &ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter{fieldPtr: ptr}
}

type Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter interface {
	setMsg(*ospfv2LsaLinkTrafficEngineering) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
	Items() []Ospfv2LsaMsd
	Add() Ospfv2LsaMsd
	Append(items ...Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
	Set(index int, newObj Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
	Clear() Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
	clearHolderSlice() Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
	appendHolderSlice(item Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter
}

func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) setMsg(msg *ospfv2LsaLinkTrafficEngineering) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaMsd{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) Items() []Ospfv2LsaMsd {
	return obj.ospfv2LsaMsdSlice
}

func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) Add() Ospfv2LsaMsd {
	newObj := &otg.Ospfv2LsaMsd{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaMsd{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) Append(items ...Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) Set(index int, newObj Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaMsdSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) Clear() Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaMsd{}
		obj.ospfv2LsaMsdSlice = []Ospfv2LsaMsd{}
	}
	return obj
}
func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) clearHolderSlice() Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	if len(obj.ospfv2LsaMsdSlice) > 0 {
		obj.ospfv2LsaMsdSlice = []Ospfv2LsaMsd{}
	}
	return obj
}
func (obj *ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter) appendHolderSlice(item Ospfv2LsaMsd) Ospfv2LsaLinkTrafficEngineeringOspfv2LsaMsdIter {
	obj.ospfv2LsaMsdSlice = append(obj.ospfv2LsaMsdSlice, item)
	return obj
}

func (obj *ospfv2LsaLinkTrafficEngineering) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocalInterfaceIpAddresses != nil {

		err := obj.validateIpv4Slice(obj.LocalInterfaceIpAddresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2LsaLinkTrafficEngineering.LocalInterfaceIpAddresses"))
		}

	}

	if obj.obj.RemoteInterfaceIpAddresses != nil {

		err := obj.validateIpv4Slice(obj.RemoteInterfaceIpAddresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2LsaLinkTrafficEngineering.RemoteInterfaceIpAddresses"))
		}

	}

	if len(obj.obj.LinkMsd) != 0 {

		if set_default {
			obj.LinkMsd().clearHolderSlice()
			for _, item := range obj.obj.LinkMsd {
				obj.LinkMsd().appendHolderSlice(&ospfv2LsaMsd{obj: item})
			}
		}
		for _, item := range obj.LinkMsd().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2LsaLinkTrafficEngineering) setDefault() {

}
