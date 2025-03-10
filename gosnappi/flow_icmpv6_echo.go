package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6Echo *****
type flowIcmpv6Echo struct {
	validation
	obj                  *otg.FlowIcmpv6Echo
	marshaller           marshalFlowIcmpv6Echo
	unMarshaller         unMarshalFlowIcmpv6Echo
	typeHolder           PatternFlowIcmpv6EchoType
	codeHolder           PatternFlowIcmpv6EchoCode
	identifierHolder     PatternFlowIcmpv6EchoIdentifier
	sequenceNumberHolder PatternFlowIcmpv6EchoSequenceNumber
	checksumHolder       PatternFlowIcmpv6EchoChecksum
}

func NewFlowIcmpv6Echo() FlowIcmpv6Echo {
	obj := flowIcmpv6Echo{obj: &otg.FlowIcmpv6Echo{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6Echo) msg() *otg.FlowIcmpv6Echo {
	return obj.obj
}

func (obj *flowIcmpv6Echo) setMsg(msg *otg.FlowIcmpv6Echo) FlowIcmpv6Echo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6Echo struct {
	obj *flowIcmpv6Echo
}

type marshalFlowIcmpv6Echo interface {
	// ToProto marshals FlowIcmpv6Echo to protobuf object *otg.FlowIcmpv6Echo
	ToProto() (*otg.FlowIcmpv6Echo, error)
	// ToPbText marshals FlowIcmpv6Echo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6Echo to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6Echo to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowIcmpv6Echo to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowIcmpv6Echo struct {
	obj *flowIcmpv6Echo
}

type unMarshalFlowIcmpv6Echo interface {
	// FromProto unmarshals FlowIcmpv6Echo from protobuf object *otg.FlowIcmpv6Echo
	FromProto(msg *otg.FlowIcmpv6Echo) (FlowIcmpv6Echo, error)
	// FromPbText unmarshals FlowIcmpv6Echo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6Echo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6Echo from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6Echo) Marshal() marshalFlowIcmpv6Echo {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6Echo{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6Echo) Unmarshal() unMarshalFlowIcmpv6Echo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6Echo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6Echo) ToProto() (*otg.FlowIcmpv6Echo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6Echo) FromProto(msg *otg.FlowIcmpv6Echo) (FlowIcmpv6Echo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6Echo) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6Echo) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6Echo) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6Echo) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6Echo) ToJsonRaw() (string, error) {
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

func (m *marshalflowIcmpv6Echo) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6Echo) FromJson(value string) error {
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

func (obj *flowIcmpv6Echo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6Echo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6Echo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6Echo) Clone() (FlowIcmpv6Echo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6Echo()
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

func (obj *flowIcmpv6Echo) setNil() {
	obj.typeHolder = nil
	obj.codeHolder = nil
	obj.identifierHolder = nil
	obj.sequenceNumberHolder = nil
	obj.checksumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6Echo is packet Header for ICMPv6 Echo
type FlowIcmpv6Echo interface {
	Validation
	// msg marshals FlowIcmpv6Echo to protobuf object *otg.FlowIcmpv6Echo
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6Echo
	// setMsg unmarshals FlowIcmpv6Echo from protobuf object *otg.FlowIcmpv6Echo
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6Echo) FlowIcmpv6Echo
	// provides marshal interface
	Marshal() marshalFlowIcmpv6Echo
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6Echo
	// validate validates FlowIcmpv6Echo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6Echo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIcmpv6EchoType, set in FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoType is iCMPv6 echo type
	Type() PatternFlowIcmpv6EchoType
	// SetType assigns PatternFlowIcmpv6EchoType provided by user to FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoType is iCMPv6 echo type
	SetType(value PatternFlowIcmpv6EchoType) FlowIcmpv6Echo
	// HasType checks if Type has been set in FlowIcmpv6Echo
	HasType() bool
	// Code returns PatternFlowIcmpv6EchoCode, set in FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoCode is iCMPv6 echo sub type
	Code() PatternFlowIcmpv6EchoCode
	// SetCode assigns PatternFlowIcmpv6EchoCode provided by user to FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoCode is iCMPv6 echo sub type
	SetCode(value PatternFlowIcmpv6EchoCode) FlowIcmpv6Echo
	// HasCode checks if Code has been set in FlowIcmpv6Echo
	HasCode() bool
	// Identifier returns PatternFlowIcmpv6EchoIdentifier, set in FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoIdentifier is iCMPv6 echo identifier
	Identifier() PatternFlowIcmpv6EchoIdentifier
	// SetIdentifier assigns PatternFlowIcmpv6EchoIdentifier provided by user to FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoIdentifier is iCMPv6 echo identifier
	SetIdentifier(value PatternFlowIcmpv6EchoIdentifier) FlowIcmpv6Echo
	// HasIdentifier checks if Identifier has been set in FlowIcmpv6Echo
	HasIdentifier() bool
	// SequenceNumber returns PatternFlowIcmpv6EchoSequenceNumber, set in FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoSequenceNumber is iCMPv6 echo sequence number
	SequenceNumber() PatternFlowIcmpv6EchoSequenceNumber
	// SetSequenceNumber assigns PatternFlowIcmpv6EchoSequenceNumber provided by user to FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoSequenceNumber is iCMPv6 echo sequence number
	SetSequenceNumber(value PatternFlowIcmpv6EchoSequenceNumber) FlowIcmpv6Echo
	// HasSequenceNumber checks if SequenceNumber has been set in FlowIcmpv6Echo
	HasSequenceNumber() bool
	// Checksum returns PatternFlowIcmpv6EchoChecksum, set in FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoChecksum is iCMPv6 checksum
	Checksum() PatternFlowIcmpv6EchoChecksum
	// SetChecksum assigns PatternFlowIcmpv6EchoChecksum provided by user to FlowIcmpv6Echo.
	// PatternFlowIcmpv6EchoChecksum is iCMPv6 checksum
	SetChecksum(value PatternFlowIcmpv6EchoChecksum) FlowIcmpv6Echo
	// HasChecksum checks if Checksum has been set in FlowIcmpv6Echo
	HasChecksum() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIcmpv6EchoType
func (obj *flowIcmpv6Echo) Type() PatternFlowIcmpv6EchoType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIcmpv6EchoType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIcmpv6EchoType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIcmpv6EchoType
func (obj *flowIcmpv6Echo) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIcmpv6EchoType value in the FlowIcmpv6Echo object
func (obj *flowIcmpv6Echo) SetType(value PatternFlowIcmpv6EchoType) FlowIcmpv6Echo {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Code returns a PatternFlowIcmpv6EchoCode
func (obj *flowIcmpv6Echo) Code() PatternFlowIcmpv6EchoCode {
	if obj.obj.Code == nil {
		obj.obj.Code = NewPatternFlowIcmpv6EchoCode().msg()
	}
	if obj.codeHolder == nil {
		obj.codeHolder = &patternFlowIcmpv6EchoCode{obj: obj.obj.Code}
	}
	return obj.codeHolder
}

// description is TBD
// Code returns a PatternFlowIcmpv6EchoCode
func (obj *flowIcmpv6Echo) HasCode() bool {
	return obj.obj.Code != nil
}

// description is TBD
// SetCode sets the PatternFlowIcmpv6EchoCode value in the FlowIcmpv6Echo object
func (obj *flowIcmpv6Echo) SetCode(value PatternFlowIcmpv6EchoCode) FlowIcmpv6Echo {

	obj.codeHolder = nil
	obj.obj.Code = value.msg()

	return obj
}

// description is TBD
// Identifier returns a PatternFlowIcmpv6EchoIdentifier
func (obj *flowIcmpv6Echo) Identifier() PatternFlowIcmpv6EchoIdentifier {
	if obj.obj.Identifier == nil {
		obj.obj.Identifier = NewPatternFlowIcmpv6EchoIdentifier().msg()
	}
	if obj.identifierHolder == nil {
		obj.identifierHolder = &patternFlowIcmpv6EchoIdentifier{obj: obj.obj.Identifier}
	}
	return obj.identifierHolder
}

// description is TBD
// Identifier returns a PatternFlowIcmpv6EchoIdentifier
func (obj *flowIcmpv6Echo) HasIdentifier() bool {
	return obj.obj.Identifier != nil
}

// description is TBD
// SetIdentifier sets the PatternFlowIcmpv6EchoIdentifier value in the FlowIcmpv6Echo object
func (obj *flowIcmpv6Echo) SetIdentifier(value PatternFlowIcmpv6EchoIdentifier) FlowIcmpv6Echo {

	obj.identifierHolder = nil
	obj.obj.Identifier = value.msg()

	return obj
}

// description is TBD
// SequenceNumber returns a PatternFlowIcmpv6EchoSequenceNumber
func (obj *flowIcmpv6Echo) SequenceNumber() PatternFlowIcmpv6EchoSequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = NewPatternFlowIcmpv6EchoSequenceNumber().msg()
	}
	if obj.sequenceNumberHolder == nil {
		obj.sequenceNumberHolder = &patternFlowIcmpv6EchoSequenceNumber{obj: obj.obj.SequenceNumber}
	}
	return obj.sequenceNumberHolder
}

// description is TBD
// SequenceNumber returns a PatternFlowIcmpv6EchoSequenceNumber
func (obj *flowIcmpv6Echo) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// description is TBD
// SetSequenceNumber sets the PatternFlowIcmpv6EchoSequenceNumber value in the FlowIcmpv6Echo object
func (obj *flowIcmpv6Echo) SetSequenceNumber(value PatternFlowIcmpv6EchoSequenceNumber) FlowIcmpv6Echo {

	obj.sequenceNumberHolder = nil
	obj.obj.SequenceNumber = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIcmpv6EchoChecksum
func (obj *flowIcmpv6Echo) Checksum() PatternFlowIcmpv6EchoChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIcmpv6EchoChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIcmpv6EchoChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIcmpv6EchoChecksum
func (obj *flowIcmpv6Echo) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIcmpv6EchoChecksum value in the FlowIcmpv6Echo object
func (obj *flowIcmpv6Echo) SetChecksum(value PatternFlowIcmpv6EchoChecksum) FlowIcmpv6Echo {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

func (obj *flowIcmpv6Echo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Code != nil {

		obj.Code().validateObj(vObj, set_default)
	}

	if obj.obj.Identifier != nil {

		obj.Identifier().validateObj(vObj, set_default)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpv6Echo) setDefault() {

}
