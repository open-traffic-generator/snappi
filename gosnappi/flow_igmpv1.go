package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv1 *****
type flowIgmpv1 struct {
	validation
	obj                *otg.FlowIgmpv1
	marshaller         marshalFlowIgmpv1
	unMarshaller       unMarshalFlowIgmpv1
	versionHolder      PatternFlowIgmpv1Version
	typeHolder         PatternFlowIgmpv1Type
	unusedHolder       PatternFlowIgmpv1Unused
	checksumHolder     PatternFlowIgmpv1Checksum
	groupAddressHolder PatternFlowIgmpv1GroupAddress
}

func NewFlowIgmpv1() FlowIgmpv1 {
	obj := flowIgmpv1{obj: &otg.FlowIgmpv1{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv1) msg() *otg.FlowIgmpv1 {
	return obj.obj
}

func (obj *flowIgmpv1) setMsg(msg *otg.FlowIgmpv1) FlowIgmpv1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv1 struct {
	obj *flowIgmpv1
}

type marshalFlowIgmpv1 interface {
	// ToProto marshals FlowIgmpv1 to protobuf object *otg.FlowIgmpv1
	ToProto() (*otg.FlowIgmpv1, error)
	// ToPbText marshals FlowIgmpv1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv1 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowIgmpv1 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowIgmpv1 struct {
	obj *flowIgmpv1
}

type unMarshalFlowIgmpv1 interface {
	// FromProto unmarshals FlowIgmpv1 from protobuf object *otg.FlowIgmpv1
	FromProto(msg *otg.FlowIgmpv1) (FlowIgmpv1, error)
	// FromPbText unmarshals FlowIgmpv1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv1 from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv1) Marshal() marshalFlowIgmpv1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv1{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv1) Unmarshal() unMarshalFlowIgmpv1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv1) ToProto() (*otg.FlowIgmpv1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv1) FromProto(msg *otg.FlowIgmpv1) (FlowIgmpv1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv1) ToPbText() (string, error) {
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

func (m *unMarshalflowIgmpv1) FromPbText(value string) error {
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

func (m *marshalflowIgmpv1) ToYaml() (string, error) {
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

func (m *unMarshalflowIgmpv1) FromYaml(value string) error {
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

func (m *marshalflowIgmpv1) ToJsonRaw() (string, error) {
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

func (m *marshalflowIgmpv1) ToJson() (string, error) {
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

func (m *unMarshalflowIgmpv1) FromJson(value string) error {
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

func (obj *flowIgmpv1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv1) Clone() (FlowIgmpv1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv1()
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

func (obj *flowIgmpv1) setNil() {
	obj.versionHolder = nil
	obj.typeHolder = nil
	obj.unusedHolder = nil
	obj.checksumHolder = nil
	obj.groupAddressHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv1 is iGMPv1 packet header
type FlowIgmpv1 interface {
	Validation
	// msg marshals FlowIgmpv1 to protobuf object *otg.FlowIgmpv1
	// and doesn't set defaults
	msg() *otg.FlowIgmpv1
	// setMsg unmarshals FlowIgmpv1 from protobuf object *otg.FlowIgmpv1
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv1) FlowIgmpv1
	// provides marshal interface
	Marshal() marshalFlowIgmpv1
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv1
	// validate validates FlowIgmpv1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns PatternFlowIgmpv1Version, set in FlowIgmpv1.
	// PatternFlowIgmpv1Version is version number
	Version() PatternFlowIgmpv1Version
	// SetVersion assigns PatternFlowIgmpv1Version provided by user to FlowIgmpv1.
	// PatternFlowIgmpv1Version is version number
	SetVersion(value PatternFlowIgmpv1Version) FlowIgmpv1
	// HasVersion checks if Version has been set in FlowIgmpv1
	HasVersion() bool
	// Type returns PatternFlowIgmpv1Type, set in FlowIgmpv1.
	// PatternFlowIgmpv1Type is type of message
	Type() PatternFlowIgmpv1Type
	// SetType assigns PatternFlowIgmpv1Type provided by user to FlowIgmpv1.
	// PatternFlowIgmpv1Type is type of message
	SetType(value PatternFlowIgmpv1Type) FlowIgmpv1
	// HasType checks if Type has been set in FlowIgmpv1
	HasType() bool
	// Unused returns PatternFlowIgmpv1Unused, set in FlowIgmpv1.
	// PatternFlowIgmpv1Unused is unused
	Unused() PatternFlowIgmpv1Unused
	// SetUnused assigns PatternFlowIgmpv1Unused provided by user to FlowIgmpv1.
	// PatternFlowIgmpv1Unused is unused
	SetUnused(value PatternFlowIgmpv1Unused) FlowIgmpv1
	// HasUnused checks if Unused has been set in FlowIgmpv1
	HasUnused() bool
	// Checksum returns PatternFlowIgmpv1Checksum, set in FlowIgmpv1.
	// PatternFlowIgmpv1Checksum is checksum
	Checksum() PatternFlowIgmpv1Checksum
	// SetChecksum assigns PatternFlowIgmpv1Checksum provided by user to FlowIgmpv1.
	// PatternFlowIgmpv1Checksum is checksum
	SetChecksum(value PatternFlowIgmpv1Checksum) FlowIgmpv1
	// HasChecksum checks if Checksum has been set in FlowIgmpv1
	HasChecksum() bool
	// GroupAddress returns PatternFlowIgmpv1GroupAddress, set in FlowIgmpv1.
	// PatternFlowIgmpv1GroupAddress is group address
	GroupAddress() PatternFlowIgmpv1GroupAddress
	// SetGroupAddress assigns PatternFlowIgmpv1GroupAddress provided by user to FlowIgmpv1.
	// PatternFlowIgmpv1GroupAddress is group address
	SetGroupAddress(value PatternFlowIgmpv1GroupAddress) FlowIgmpv1
	// HasGroupAddress checks if GroupAddress has been set in FlowIgmpv1
	HasGroupAddress() bool
	setNil()
}

// description is TBD
// Version returns a PatternFlowIgmpv1Version
func (obj *flowIgmpv1) Version() PatternFlowIgmpv1Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowIgmpv1Version().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowIgmpv1Version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowIgmpv1Version
func (obj *flowIgmpv1) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowIgmpv1Version value in the FlowIgmpv1 object
func (obj *flowIgmpv1) SetVersion(value PatternFlowIgmpv1Version) FlowIgmpv1 {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// Type returns a PatternFlowIgmpv1Type
func (obj *flowIgmpv1) Type() PatternFlowIgmpv1Type {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIgmpv1Type().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIgmpv1Type{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIgmpv1Type
func (obj *flowIgmpv1) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIgmpv1Type value in the FlowIgmpv1 object
func (obj *flowIgmpv1) SetType(value PatternFlowIgmpv1Type) FlowIgmpv1 {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Unused returns a PatternFlowIgmpv1Unused
func (obj *flowIgmpv1) Unused() PatternFlowIgmpv1Unused {
	if obj.obj.Unused == nil {
		obj.obj.Unused = NewPatternFlowIgmpv1Unused().msg()
	}
	if obj.unusedHolder == nil {
		obj.unusedHolder = &patternFlowIgmpv1Unused{obj: obj.obj.Unused}
	}
	return obj.unusedHolder
}

// description is TBD
// Unused returns a PatternFlowIgmpv1Unused
func (obj *flowIgmpv1) HasUnused() bool {
	return obj.obj.Unused != nil
}

// description is TBD
// SetUnused sets the PatternFlowIgmpv1Unused value in the FlowIgmpv1 object
func (obj *flowIgmpv1) SetUnused(value PatternFlowIgmpv1Unused) FlowIgmpv1 {

	obj.unusedHolder = nil
	obj.obj.Unused = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIgmpv1Checksum
func (obj *flowIgmpv1) Checksum() PatternFlowIgmpv1Checksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIgmpv1Checksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIgmpv1Checksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIgmpv1Checksum
func (obj *flowIgmpv1) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIgmpv1Checksum value in the FlowIgmpv1 object
func (obj *flowIgmpv1) SetChecksum(value PatternFlowIgmpv1Checksum) FlowIgmpv1 {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// GroupAddress returns a PatternFlowIgmpv1GroupAddress
func (obj *flowIgmpv1) GroupAddress() PatternFlowIgmpv1GroupAddress {
	if obj.obj.GroupAddress == nil {
		obj.obj.GroupAddress = NewPatternFlowIgmpv1GroupAddress().msg()
	}
	if obj.groupAddressHolder == nil {
		obj.groupAddressHolder = &patternFlowIgmpv1GroupAddress{obj: obj.obj.GroupAddress}
	}
	return obj.groupAddressHolder
}

// description is TBD
// GroupAddress returns a PatternFlowIgmpv1GroupAddress
func (obj *flowIgmpv1) HasGroupAddress() bool {
	return obj.obj.GroupAddress != nil
}

// description is TBD
// SetGroupAddress sets the PatternFlowIgmpv1GroupAddress value in the FlowIgmpv1 object
func (obj *flowIgmpv1) SetGroupAddress(value PatternFlowIgmpv1GroupAddress) FlowIgmpv1 {

	obj.groupAddressHolder = nil
	obj.obj.GroupAddress = value.msg()

	return obj
}

func (obj *flowIgmpv1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Unused != nil {

		obj.Unused().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.GroupAddress != nil {

		obj.GroupAddress().validateObj(vObj, set_default)
	}

}

func (obj *flowIgmpv1) setDefault() {

}
