package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsDuidUuid *****
type dhcpv6ClientOptionsDuidUuid struct {
	validation
	obj           *otg.Dhcpv6ClientOptionsDuidUuid
	marshaller    marshalDhcpv6ClientOptionsDuidUuid
	unMarshaller  unMarshalDhcpv6ClientOptionsDuidUuid
	versionHolder Dhcpv6ClientOptionsDuidUuidVersion
	variantHolder Dhcpv6ClientOptionsDuidUuidVariant
}

func NewDhcpv6ClientOptionsDuidUuid() Dhcpv6ClientOptionsDuidUuid {
	obj := dhcpv6ClientOptionsDuidUuid{obj: &otg.Dhcpv6ClientOptionsDuidUuid{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsDuidUuid) msg() *otg.Dhcpv6ClientOptionsDuidUuid {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsDuidUuid) setMsg(msg *otg.Dhcpv6ClientOptionsDuidUuid) Dhcpv6ClientOptionsDuidUuid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsDuidUuid struct {
	obj *dhcpv6ClientOptionsDuidUuid
}

type marshalDhcpv6ClientOptionsDuidUuid interface {
	// ToProto marshals Dhcpv6ClientOptionsDuidUuid to protobuf object *otg.Dhcpv6ClientOptionsDuidUuid
	ToProto() (*otg.Dhcpv6ClientOptionsDuidUuid, error)
	// ToPbText marshals Dhcpv6ClientOptionsDuidUuid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsDuidUuid to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsDuidUuid to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsDuidUuid to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsDuidUuid struct {
	obj *dhcpv6ClientOptionsDuidUuid
}

type unMarshalDhcpv6ClientOptionsDuidUuid interface {
	// FromProto unmarshals Dhcpv6ClientOptionsDuidUuid from protobuf object *otg.Dhcpv6ClientOptionsDuidUuid
	FromProto(msg *otg.Dhcpv6ClientOptionsDuidUuid) (Dhcpv6ClientOptionsDuidUuid, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsDuidUuid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsDuidUuid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsDuidUuid from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsDuidUuid) Marshal() marshalDhcpv6ClientOptionsDuidUuid {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsDuidUuid{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsDuidUuid) Unmarshal() unMarshalDhcpv6ClientOptionsDuidUuid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsDuidUuid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsDuidUuid) ToProto() (*otg.Dhcpv6ClientOptionsDuidUuid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsDuidUuid) FromProto(msg *otg.Dhcpv6ClientOptionsDuidUuid) (Dhcpv6ClientOptionsDuidUuid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsDuidUuid) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuid) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidUuid) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuid) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsDuidUuid) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshaldhcpv6ClientOptionsDuidUuid) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsDuidUuid) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsDuidUuid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidUuid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsDuidUuid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsDuidUuid) Clone() (Dhcpv6ClientOptionsDuidUuid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsDuidUuid()
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

func (obj *dhcpv6ClientOptionsDuidUuid) setNil() {
	obj.versionHolder = nil
	obj.variantHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsDuidUuid is dUID embedded a Universally Unique IDentifier (UUID). A UUID is an identifier that is unique across both  space and time, with respect to the space of all UUIDs.
type Dhcpv6ClientOptionsDuidUuid interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsDuidUuid to protobuf object *otg.Dhcpv6ClientOptionsDuidUuid
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsDuidUuid
	// setMsg unmarshals Dhcpv6ClientOptionsDuidUuid from protobuf object *otg.Dhcpv6ClientOptionsDuidUuid
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsDuidUuid) Dhcpv6ClientOptionsDuidUuid
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsDuidUuid
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsDuidUuid
	// validate validates Dhcpv6ClientOptionsDuidUuid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsDuidUuid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns Dhcpv6ClientOptionsDuidUuidVersion, set in Dhcpv6ClientOptionsDuidUuid.
	// Dhcpv6ClientOptionsDuidUuidVersion is the version number is in the most significant 4 bits of the timestamp (bits 4 through 7 of the time_hi_and_version field).
	Version() Dhcpv6ClientOptionsDuidUuidVersion
	// SetVersion assigns Dhcpv6ClientOptionsDuidUuidVersion provided by user to Dhcpv6ClientOptionsDuidUuid.
	// Dhcpv6ClientOptionsDuidUuidVersion is the version number is in the most significant 4 bits of the timestamp (bits 4 through 7 of the time_hi_and_version field).
	SetVersion(value Dhcpv6ClientOptionsDuidUuidVersion) Dhcpv6ClientOptionsDuidUuid
	// HasVersion checks if Version has been set in Dhcpv6ClientOptionsDuidUuid
	HasVersion() bool
	// Variant returns Dhcpv6ClientOptionsDuidUuidVariant, set in Dhcpv6ClientOptionsDuidUuid.
	// Dhcpv6ClientOptionsDuidUuidVariant is the variant field determines the layout of the UUID.  That is, the interpretation of all other bits in the  UUID depends on the setting of the bits in the variant field).
	Variant() Dhcpv6ClientOptionsDuidUuidVariant
	// SetVariant assigns Dhcpv6ClientOptionsDuidUuidVariant provided by user to Dhcpv6ClientOptionsDuidUuid.
	// Dhcpv6ClientOptionsDuidUuidVariant is the variant field determines the layout of the UUID.  That is, the interpretation of all other bits in the  UUID depends on the setting of the bits in the variant field).
	SetVariant(value Dhcpv6ClientOptionsDuidUuidVariant) Dhcpv6ClientOptionsDuidUuid
	// HasVariant checks if Variant has been set in Dhcpv6ClientOptionsDuidUuid
	HasVariant() bool
	// TimeLow returns uint32, set in Dhcpv6ClientOptionsDuidUuid.
	TimeLow() uint32
	// SetTimeLow assigns uint32 provided by user to Dhcpv6ClientOptionsDuidUuid
	SetTimeLow(value uint32) Dhcpv6ClientOptionsDuidUuid
	// HasTimeLow checks if TimeLow has been set in Dhcpv6ClientOptionsDuidUuid
	HasTimeLow() bool
	// TimeMid returns uint32, set in Dhcpv6ClientOptionsDuidUuid.
	TimeMid() uint32
	// SetTimeMid assigns uint32 provided by user to Dhcpv6ClientOptionsDuidUuid
	SetTimeMid(value uint32) Dhcpv6ClientOptionsDuidUuid
	// HasTimeMid checks if TimeMid has been set in Dhcpv6ClientOptionsDuidUuid
	HasTimeMid() bool
	// TimeHiAndVersion returns uint32, set in Dhcpv6ClientOptionsDuidUuid.
	TimeHiAndVersion() uint32
	// SetTimeHiAndVersion assigns uint32 provided by user to Dhcpv6ClientOptionsDuidUuid
	SetTimeHiAndVersion(value uint32) Dhcpv6ClientOptionsDuidUuid
	// HasTimeHiAndVersion checks if TimeHiAndVersion has been set in Dhcpv6ClientOptionsDuidUuid
	HasTimeHiAndVersion() bool
	// ClockSeqHiAndReserved returns uint32, set in Dhcpv6ClientOptionsDuidUuid.
	ClockSeqHiAndReserved() uint32
	// SetClockSeqHiAndReserved assigns uint32 provided by user to Dhcpv6ClientOptionsDuidUuid
	SetClockSeqHiAndReserved(value uint32) Dhcpv6ClientOptionsDuidUuid
	// HasClockSeqHiAndReserved checks if ClockSeqHiAndReserved has been set in Dhcpv6ClientOptionsDuidUuid
	HasClockSeqHiAndReserved() bool
	// ClockSeqLow returns uint32, set in Dhcpv6ClientOptionsDuidUuid.
	ClockSeqLow() uint32
	// SetClockSeqLow assigns uint32 provided by user to Dhcpv6ClientOptionsDuidUuid
	SetClockSeqLow(value uint32) Dhcpv6ClientOptionsDuidUuid
	// HasClockSeqLow checks if ClockSeqLow has been set in Dhcpv6ClientOptionsDuidUuid
	HasClockSeqLow() bool
	// Node returns string, set in Dhcpv6ClientOptionsDuidUuid.
	Node() string
	// SetNode assigns string provided by user to Dhcpv6ClientOptionsDuidUuid
	SetNode(value string) Dhcpv6ClientOptionsDuidUuid
	// HasNode checks if Node has been set in Dhcpv6ClientOptionsDuidUuid
	HasNode() bool
	setNil()
}

// The version number is in the most significant 4 bits of the timestamp (bits 4 through 7 of the time_hi_and_version field).
// Version returns a Dhcpv6ClientOptionsDuidUuidVersion
func (obj *dhcpv6ClientOptionsDuidUuid) Version() Dhcpv6ClientOptionsDuidUuidVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = NewDhcpv6ClientOptionsDuidUuidVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &dhcpv6ClientOptionsDuidUuidVersion{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// The version number is in the most significant 4 bits of the timestamp (bits 4 through 7 of the time_hi_and_version field).
// Version returns a Dhcpv6ClientOptionsDuidUuidVersion
func (obj *dhcpv6ClientOptionsDuidUuid) HasVersion() bool {
	return obj.obj.Version != nil
}

// The version number is in the most significant 4 bits of the timestamp (bits 4 through 7 of the time_hi_and_version field).
// SetVersion sets the Dhcpv6ClientOptionsDuidUuidVersion value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetVersion(value Dhcpv6ClientOptionsDuidUuidVersion) Dhcpv6ClientOptionsDuidUuid {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// The variant field determines the layout of the UUID. It is multiplexed with clock_seq_hi_and_reserved.
// Variant returns a Dhcpv6ClientOptionsDuidUuidVariant
func (obj *dhcpv6ClientOptionsDuidUuid) Variant() Dhcpv6ClientOptionsDuidUuidVariant {
	if obj.obj.Variant == nil {
		obj.obj.Variant = NewDhcpv6ClientOptionsDuidUuidVariant().msg()
	}
	if obj.variantHolder == nil {
		obj.variantHolder = &dhcpv6ClientOptionsDuidUuidVariant{obj: obj.obj.Variant}
	}
	return obj.variantHolder
}

// The variant field determines the layout of the UUID. It is multiplexed with clock_seq_hi_and_reserved.
// Variant returns a Dhcpv6ClientOptionsDuidUuidVariant
func (obj *dhcpv6ClientOptionsDuidUuid) HasVariant() bool {
	return obj.obj.Variant != nil
}

// The variant field determines the layout of the UUID. It is multiplexed with clock_seq_hi_and_reserved.
// SetVariant sets the Dhcpv6ClientOptionsDuidUuidVariant value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetVariant(value Dhcpv6ClientOptionsDuidUuidVariant) Dhcpv6ClientOptionsDuidUuid {

	obj.variantHolder = nil
	obj.obj.Variant = value.msg()

	return obj
}

// The low field of the timestamp.
// TimeLow returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) TimeLow() uint32 {

	return *obj.obj.TimeLow

}

// The low field of the timestamp.
// TimeLow returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) HasTimeLow() bool {
	return obj.obj.TimeLow != nil
}

// The low field of the timestamp.
// SetTimeLow sets the uint32 value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetTimeLow(value uint32) Dhcpv6ClientOptionsDuidUuid {

	obj.obj.TimeLow = &value
	return obj
}

// The middle field of the timestamp.
// TimeMid returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) TimeMid() uint32 {

	return *obj.obj.TimeMid

}

// The middle field of the timestamp.
// TimeMid returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) HasTimeMid() bool {
	return obj.obj.TimeMid != nil
}

// The middle field of the timestamp.
// SetTimeMid sets the uint32 value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetTimeMid(value uint32) Dhcpv6ClientOptionsDuidUuid {

	obj.obj.TimeMid = &value
	return obj
}

// The high field of the timestamp multiplexed with the version number.
// TimeHiAndVersion returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) TimeHiAndVersion() uint32 {

	return *obj.obj.TimeHiAndVersion

}

// The high field of the timestamp multiplexed with the version number.
// TimeHiAndVersion returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) HasTimeHiAndVersion() bool {
	return obj.obj.TimeHiAndVersion != nil
}

// The high field of the timestamp multiplexed with the version number.
// SetTimeHiAndVersion sets the uint32 value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetTimeHiAndVersion(value uint32) Dhcpv6ClientOptionsDuidUuid {

	obj.obj.TimeHiAndVersion = &value
	return obj
}

// The high field of the clock sequence multiplexed with the variant.
// ClockSeqHiAndReserved returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) ClockSeqHiAndReserved() uint32 {

	return *obj.obj.ClockSeqHiAndReserved

}

// The high field of the clock sequence multiplexed with the variant.
// ClockSeqHiAndReserved returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) HasClockSeqHiAndReserved() bool {
	return obj.obj.ClockSeqHiAndReserved != nil
}

// The high field of the clock sequence multiplexed with the variant.
// SetClockSeqHiAndReserved sets the uint32 value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetClockSeqHiAndReserved(value uint32) Dhcpv6ClientOptionsDuidUuid {

	obj.obj.ClockSeqHiAndReserved = &value
	return obj
}

// The low field of the clock sequence.
// ClockSeqLow returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) ClockSeqLow() uint32 {

	return *obj.obj.ClockSeqLow

}

// The low field of the clock sequence.
// ClockSeqLow returns a uint32
func (obj *dhcpv6ClientOptionsDuidUuid) HasClockSeqLow() bool {
	return obj.obj.ClockSeqLow != nil
}

// The low field of the clock sequence.
// SetClockSeqLow sets the uint32 value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetClockSeqLow(value uint32) Dhcpv6ClientOptionsDuidUuid {

	obj.obj.ClockSeqLow = &value
	return obj
}

// The spatially unique node identifier.
// Node returns a string
func (obj *dhcpv6ClientOptionsDuidUuid) Node() string {

	return *obj.obj.Node

}

// The spatially unique node identifier.
// Node returns a string
func (obj *dhcpv6ClientOptionsDuidUuid) HasNode() bool {
	return obj.obj.Node != nil
}

// The spatially unique node identifier.
// SetNode sets the string value in the Dhcpv6ClientOptionsDuidUuid object
func (obj *dhcpv6ClientOptionsDuidUuid) SetNode(value string) Dhcpv6ClientOptionsDuidUuid {

	obj.obj.Node = &value
	return obj
}

func (obj *dhcpv6ClientOptionsDuidUuid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.Variant != nil {

		obj.Variant().validateObj(vObj, set_default)
	}

	if obj.obj.TimeMid != nil {

		if *obj.obj.TimeMid > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsDuidUuid.TimeMid <= 65535 but Got %d", *obj.obj.TimeMid))
		}

	}

	if obj.obj.TimeHiAndVersion != nil {

		if *obj.obj.TimeHiAndVersion > 4095 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsDuidUuid.TimeHiAndVersion <= 4095 but Got %d", *obj.obj.TimeHiAndVersion))
		}

	}

	if obj.obj.ClockSeqHiAndReserved != nil {

		if *obj.obj.ClockSeqHiAndReserved > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsDuidUuid.ClockSeqHiAndReserved <= 31 but Got %d", *obj.obj.ClockSeqHiAndReserved))
		}

	}

	if obj.obj.ClockSeqLow != nil {

		if *obj.obj.ClockSeqLow > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsDuidUuid.ClockSeqLow <= 127 but Got %d", *obj.obj.ClockSeqLow))
		}

	}

	if obj.obj.Node != nil {

		err := obj.validateMac(obj.Node())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Dhcpv6ClientOptionsDuidUuid.Node"))
		}

	}

}

func (obj *dhcpv6ClientOptionsDuidUuid) setDefault() {
	if obj.obj.TimeLow == nil {
		obj.SetTimeLow(0)
	}
	if obj.obj.TimeMid == nil {
		obj.SetTimeMid(0)
	}
	if obj.obj.TimeHiAndVersion == nil {
		obj.SetTimeHiAndVersion(0)
	}
	if obj.obj.ClockSeqHiAndReserved == nil {
		obj.SetClockSeqHiAndReserved(0)
	}
	if obj.obj.ClockSeqLow == nil {
		obj.SetClockSeqLow(0)
	}
	if obj.obj.Node == nil {
		obj.SetNode("00:00:00:00:00:00")
	}

}
