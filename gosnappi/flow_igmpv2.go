package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv2 *****
type flowIgmpv2 struct {
	validation
	obj                   *otg.FlowIgmpv2
	marshaller            marshalFlowIgmpv2
	unMarshaller          unMarshalFlowIgmpv2
	typeHolder            PatternFlowIgmpv2Type
	maxResponseTimeHolder PatternFlowIgmpv2MaxResponseTime
	checksumHolder        PatternFlowIgmpv2Checksum
	groupAddressHolder    PatternFlowIgmpv2GroupAddress
}

func NewFlowIgmpv2() FlowIgmpv2 {
	obj := flowIgmpv2{obj: &otg.FlowIgmpv2{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv2) msg() *otg.FlowIgmpv2 {
	return obj.obj
}

func (obj *flowIgmpv2) setMsg(msg *otg.FlowIgmpv2) FlowIgmpv2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv2 struct {
	obj *flowIgmpv2
}

type marshalFlowIgmpv2 interface {
	// ToProto marshals FlowIgmpv2 to protobuf object *otg.FlowIgmpv2
	ToProto() (*otg.FlowIgmpv2, error)
	// ToPbText marshals FlowIgmpv2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv2 to JSON text
	ToJson() (string, error)
}

type unMarshalflowIgmpv2 struct {
	obj *flowIgmpv2
}

type unMarshalFlowIgmpv2 interface {
	// FromProto unmarshals FlowIgmpv2 from protobuf object *otg.FlowIgmpv2
	FromProto(msg *otg.FlowIgmpv2) (FlowIgmpv2, error)
	// FromPbText unmarshals FlowIgmpv2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv2 from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv2) Marshal() marshalFlowIgmpv2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv2{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv2) Unmarshal() unMarshalFlowIgmpv2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv2) ToProto() (*otg.FlowIgmpv2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv2) FromProto(msg *otg.FlowIgmpv2) (FlowIgmpv2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv2) ToPbText() (string, error) {
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

func (m *unMarshalflowIgmpv2) FromPbText(value string) error {
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

func (m *marshalflowIgmpv2) ToYaml() (string, error) {
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

func (m *unMarshalflowIgmpv2) FromYaml(value string) error {
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

func (m *marshalflowIgmpv2) ToJson() (string, error) {
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

func (m *unMarshalflowIgmpv2) FromJson(value string) error {
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

func (obj *flowIgmpv2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv2) Clone() (FlowIgmpv2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv2()
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

func (obj *flowIgmpv2) setNil() {
	obj.typeHolder = nil
	obj.maxResponseTimeHolder = nil
	obj.checksumHolder = nil
	obj.groupAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv2 is iGMPv2 packet header (RFC 2236)
type FlowIgmpv2 interface {
	Validation
	// msg marshals FlowIgmpv2 to protobuf object *otg.FlowIgmpv2
	// and doesn't set defaults
	msg() *otg.FlowIgmpv2
	// setMsg unmarshals FlowIgmpv2 from protobuf object *otg.FlowIgmpv2
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv2) FlowIgmpv2
	// provides marshal interface
	Marshal() marshalFlowIgmpv2
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv2
	// validate validates FlowIgmpv2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIgmpv2Type, set in FlowIgmpv2.
	// PatternFlowIgmpv2Type is iGMPv2 message type
	Type() PatternFlowIgmpv2Type
	// SetType assigns PatternFlowIgmpv2Type provided by user to FlowIgmpv2.
	// PatternFlowIgmpv2Type is iGMPv2 message type
	SetType(value PatternFlowIgmpv2Type) FlowIgmpv2
	// HasType checks if Type has been set in FlowIgmpv2
	HasType() bool
	// MaxResponseTime returns PatternFlowIgmpv2MaxResponseTime, set in FlowIgmpv2.
	// PatternFlowIgmpv2MaxResponseTime is maximum response time in units of 1/10 second. Meaningful only in Membership Query messages; ignored in Reports and Leave Group (RFC 2236 Section 2.2).
	MaxResponseTime() PatternFlowIgmpv2MaxResponseTime
	// SetMaxResponseTime assigns PatternFlowIgmpv2MaxResponseTime provided by user to FlowIgmpv2.
	// PatternFlowIgmpv2MaxResponseTime is maximum response time in units of 1/10 second. Meaningful only in Membership Query messages; ignored in Reports and Leave Group (RFC 2236 Section 2.2).
	SetMaxResponseTime(value PatternFlowIgmpv2MaxResponseTime) FlowIgmpv2
	// HasMaxResponseTime checks if MaxResponseTime has been set in FlowIgmpv2
	HasMaxResponseTime() bool
	// Checksum returns PatternFlowIgmpv2Checksum, set in FlowIgmpv2.
	// PatternFlowIgmpv2Checksum is checksum
	Checksum() PatternFlowIgmpv2Checksum
	// SetChecksum assigns PatternFlowIgmpv2Checksum provided by user to FlowIgmpv2.
	// PatternFlowIgmpv2Checksum is checksum
	SetChecksum(value PatternFlowIgmpv2Checksum) FlowIgmpv2
	// HasChecksum checks if Checksum has been set in FlowIgmpv2
	HasChecksum() bool
	// GroupAddress returns PatternFlowIgmpv2GroupAddress, set in FlowIgmpv2.
	// PatternFlowIgmpv2GroupAddress is group address is zero in a general Membership Query. It is set to the group being reported/left in Reports and Leave Group messages (RFC 2236 Section 2.4).
	GroupAddress() PatternFlowIgmpv2GroupAddress
	// SetGroupAddress assigns PatternFlowIgmpv2GroupAddress provided by user to FlowIgmpv2.
	// PatternFlowIgmpv2GroupAddress is group address is zero in a general Membership Query. It is set to the group being reported/left in Reports and Leave Group messages (RFC 2236 Section 2.4).
	SetGroupAddress(value PatternFlowIgmpv2GroupAddress) FlowIgmpv2
	// HasGroupAddress checks if GroupAddress has been set in FlowIgmpv2
	HasGroupAddress() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIgmpv2Type
func (obj *flowIgmpv2) Type() PatternFlowIgmpv2Type {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIgmpv2Type().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIgmpv2Type{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIgmpv2Type
func (obj *flowIgmpv2) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIgmpv2Type value in the FlowIgmpv2 object
func (obj *flowIgmpv2) SetType(value PatternFlowIgmpv2Type) FlowIgmpv2 {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// MaxResponseTime returns a PatternFlowIgmpv2MaxResponseTime
func (obj *flowIgmpv2) MaxResponseTime() PatternFlowIgmpv2MaxResponseTime {
	if obj.obj.MaxResponseTime == nil {
		obj.obj.MaxResponseTime = NewPatternFlowIgmpv2MaxResponseTime().msg()
	}
	if obj.maxResponseTimeHolder == nil {
		obj.maxResponseTimeHolder = &patternFlowIgmpv2MaxResponseTime{obj: obj.obj.MaxResponseTime}
	}
	return obj.maxResponseTimeHolder
}

// description is TBD
// MaxResponseTime returns a PatternFlowIgmpv2MaxResponseTime
func (obj *flowIgmpv2) HasMaxResponseTime() bool {
	return obj.obj.MaxResponseTime != nil
}

// description is TBD
// SetMaxResponseTime sets the PatternFlowIgmpv2MaxResponseTime value in the FlowIgmpv2 object
func (obj *flowIgmpv2) SetMaxResponseTime(value PatternFlowIgmpv2MaxResponseTime) FlowIgmpv2 {

	obj.maxResponseTimeHolder = nil
	obj.obj.MaxResponseTime = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIgmpv2Checksum
func (obj *flowIgmpv2) Checksum() PatternFlowIgmpv2Checksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIgmpv2Checksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIgmpv2Checksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIgmpv2Checksum
func (obj *flowIgmpv2) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIgmpv2Checksum value in the FlowIgmpv2 object
func (obj *flowIgmpv2) SetChecksum(value PatternFlowIgmpv2Checksum) FlowIgmpv2 {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// GroupAddress returns a PatternFlowIgmpv2GroupAddress
func (obj *flowIgmpv2) GroupAddress() PatternFlowIgmpv2GroupAddress {
	if obj.obj.GroupAddress == nil {
		obj.obj.GroupAddress = NewPatternFlowIgmpv2GroupAddress().msg()
	}
	if obj.groupAddressHolder == nil {
		obj.groupAddressHolder = &patternFlowIgmpv2GroupAddress{obj: obj.obj.GroupAddress}
	}
	return obj.groupAddressHolder
}

// description is TBD
// GroupAddress returns a PatternFlowIgmpv2GroupAddress
func (obj *flowIgmpv2) HasGroupAddress() bool {
	return obj.obj.GroupAddress != nil
}

// description is TBD
// SetGroupAddress sets the PatternFlowIgmpv2GroupAddress value in the FlowIgmpv2 object
func (obj *flowIgmpv2) SetGroupAddress(value PatternFlowIgmpv2GroupAddress) FlowIgmpv2 {

	obj.groupAddressHolder = nil
	obj.obj.GroupAddress = value.msg()

	return obj
}

func (obj *flowIgmpv2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.MaxResponseTime != nil {

		obj.MaxResponseTime().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.GroupAddress != nil {

		obj.GroupAddress().validateObj(vObj, set_default)
	}

}

func (obj *flowIgmpv2) setDefault() {

}
