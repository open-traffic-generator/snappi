package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpEcho *****
type flowIcmpEcho struct {
	validation
	obj                  *otg.FlowIcmpEcho
	marshaller           marshalFlowIcmpEcho
	unMarshaller         unMarshalFlowIcmpEcho
	typeHolder           PatternFlowIcmpEchoType
	codeHolder           PatternFlowIcmpEchoCode
	checksumHolder       PatternFlowIcmpEchoChecksum
	identifierHolder     PatternFlowIcmpEchoIdentifier
	sequenceNumberHolder PatternFlowIcmpEchoSequenceNumber
}

func NewFlowIcmpEcho() FlowIcmpEcho {
	obj := flowIcmpEcho{obj: &otg.FlowIcmpEcho{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpEcho) msg() *otg.FlowIcmpEcho {
	return obj.obj
}

func (obj *flowIcmpEcho) setMsg(msg *otg.FlowIcmpEcho) FlowIcmpEcho {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpEcho struct {
	obj *flowIcmpEcho
}

type marshalFlowIcmpEcho interface {
	// ToProto marshals FlowIcmpEcho to protobuf object *otg.FlowIcmpEcho
	ToProto() (*otg.FlowIcmpEcho, error)
	// ToPbText marshals FlowIcmpEcho to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpEcho to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpEcho to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowIcmpEcho to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowIcmpEcho struct {
	obj *flowIcmpEcho
}

type unMarshalFlowIcmpEcho interface {
	// FromProto unmarshals FlowIcmpEcho from protobuf object *otg.FlowIcmpEcho
	FromProto(msg *otg.FlowIcmpEcho) (FlowIcmpEcho, error)
	// FromPbText unmarshals FlowIcmpEcho from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpEcho from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpEcho from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpEcho) Marshal() marshalFlowIcmpEcho {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpEcho{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpEcho) Unmarshal() unMarshalFlowIcmpEcho {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpEcho{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpEcho) ToProto() (*otg.FlowIcmpEcho, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpEcho) FromProto(msg *otg.FlowIcmpEcho) (FlowIcmpEcho, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpEcho) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpEcho) FromPbText(value string) error {
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

func (m *marshalflowIcmpEcho) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpEcho) FromYaml(value string) error {
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

func (m *marshalflowIcmpEcho) ToJsonRaw() (string, error) {
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

func (m *marshalflowIcmpEcho) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpEcho) FromJson(value string) error {
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

func (obj *flowIcmpEcho) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpEcho) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpEcho) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpEcho) Clone() (FlowIcmpEcho, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpEcho()
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

func (obj *flowIcmpEcho) setNil() {
	obj.typeHolder = nil
	obj.codeHolder = nil
	obj.checksumHolder = nil
	obj.identifierHolder = nil
	obj.sequenceNumberHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpEcho is packet Header for ICMP echo request
type FlowIcmpEcho interface {
	Validation
	// msg marshals FlowIcmpEcho to protobuf object *otg.FlowIcmpEcho
	// and doesn't set defaults
	msg() *otg.FlowIcmpEcho
	// setMsg unmarshals FlowIcmpEcho from protobuf object *otg.FlowIcmpEcho
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpEcho) FlowIcmpEcho
	// provides marshal interface
	Marshal() marshalFlowIcmpEcho
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpEcho
	// validate validates FlowIcmpEcho
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpEcho, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowIcmpEchoType, set in FlowIcmpEcho.
	// PatternFlowIcmpEchoType is the type of ICMP echo packet
	Type() PatternFlowIcmpEchoType
	// SetType assigns PatternFlowIcmpEchoType provided by user to FlowIcmpEcho.
	// PatternFlowIcmpEchoType is the type of ICMP echo packet
	SetType(value PatternFlowIcmpEchoType) FlowIcmpEcho
	// HasType checks if Type has been set in FlowIcmpEcho
	HasType() bool
	// Code returns PatternFlowIcmpEchoCode, set in FlowIcmpEcho.
	// PatternFlowIcmpEchoCode is the ICMP subtype.  The default code for ICMP echo request and reply is 0.
	Code() PatternFlowIcmpEchoCode
	// SetCode assigns PatternFlowIcmpEchoCode provided by user to FlowIcmpEcho.
	// PatternFlowIcmpEchoCode is the ICMP subtype.  The default code for ICMP echo request and reply is 0.
	SetCode(value PatternFlowIcmpEchoCode) FlowIcmpEcho
	// HasCode checks if Code has been set in FlowIcmpEcho
	HasCode() bool
	// Checksum returns PatternFlowIcmpEchoChecksum, set in FlowIcmpEcho.
	// PatternFlowIcmpEchoChecksum is iCMP checksum
	Checksum() PatternFlowIcmpEchoChecksum
	// SetChecksum assigns PatternFlowIcmpEchoChecksum provided by user to FlowIcmpEcho.
	// PatternFlowIcmpEchoChecksum is iCMP checksum
	SetChecksum(value PatternFlowIcmpEchoChecksum) FlowIcmpEcho
	// HasChecksum checks if Checksum has been set in FlowIcmpEcho
	HasChecksum() bool
	// Identifier returns PatternFlowIcmpEchoIdentifier, set in FlowIcmpEcho.
	// PatternFlowIcmpEchoIdentifier is iCMP identifier
	Identifier() PatternFlowIcmpEchoIdentifier
	// SetIdentifier assigns PatternFlowIcmpEchoIdentifier provided by user to FlowIcmpEcho.
	// PatternFlowIcmpEchoIdentifier is iCMP identifier
	SetIdentifier(value PatternFlowIcmpEchoIdentifier) FlowIcmpEcho
	// HasIdentifier checks if Identifier has been set in FlowIcmpEcho
	HasIdentifier() bool
	// SequenceNumber returns PatternFlowIcmpEchoSequenceNumber, set in FlowIcmpEcho.
	// PatternFlowIcmpEchoSequenceNumber is iCMP sequence number
	SequenceNumber() PatternFlowIcmpEchoSequenceNumber
	// SetSequenceNumber assigns PatternFlowIcmpEchoSequenceNumber provided by user to FlowIcmpEcho.
	// PatternFlowIcmpEchoSequenceNumber is iCMP sequence number
	SetSequenceNumber(value PatternFlowIcmpEchoSequenceNumber) FlowIcmpEcho
	// HasSequenceNumber checks if SequenceNumber has been set in FlowIcmpEcho
	HasSequenceNumber() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowIcmpEchoType
func (obj *flowIcmpEcho) Type() PatternFlowIcmpEchoType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowIcmpEchoType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowIcmpEchoType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowIcmpEchoType
func (obj *flowIcmpEcho) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowIcmpEchoType value in the FlowIcmpEcho object
func (obj *flowIcmpEcho) SetType(value PatternFlowIcmpEchoType) FlowIcmpEcho {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Code returns a PatternFlowIcmpEchoCode
func (obj *flowIcmpEcho) Code() PatternFlowIcmpEchoCode {
	if obj.obj.Code == nil {
		obj.obj.Code = NewPatternFlowIcmpEchoCode().msg()
	}
	if obj.codeHolder == nil {
		obj.codeHolder = &patternFlowIcmpEchoCode{obj: obj.obj.Code}
	}
	return obj.codeHolder
}

// description is TBD
// Code returns a PatternFlowIcmpEchoCode
func (obj *flowIcmpEcho) HasCode() bool {
	return obj.obj.Code != nil
}

// description is TBD
// SetCode sets the PatternFlowIcmpEchoCode value in the FlowIcmpEcho object
func (obj *flowIcmpEcho) SetCode(value PatternFlowIcmpEchoCode) FlowIcmpEcho {

	obj.codeHolder = nil
	obj.obj.Code = value.msg()

	return obj
}

// description is TBD
// Checksum returns a PatternFlowIcmpEchoChecksum
func (obj *flowIcmpEcho) Checksum() PatternFlowIcmpEchoChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternFlowIcmpEchoChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternFlowIcmpEchoChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternFlowIcmpEchoChecksum
func (obj *flowIcmpEcho) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternFlowIcmpEchoChecksum value in the FlowIcmpEcho object
func (obj *flowIcmpEcho) SetChecksum(value PatternFlowIcmpEchoChecksum) FlowIcmpEcho {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

// description is TBD
// Identifier returns a PatternFlowIcmpEchoIdentifier
func (obj *flowIcmpEcho) Identifier() PatternFlowIcmpEchoIdentifier {
	if obj.obj.Identifier == nil {
		obj.obj.Identifier = NewPatternFlowIcmpEchoIdentifier().msg()
	}
	if obj.identifierHolder == nil {
		obj.identifierHolder = &patternFlowIcmpEchoIdentifier{obj: obj.obj.Identifier}
	}
	return obj.identifierHolder
}

// description is TBD
// Identifier returns a PatternFlowIcmpEchoIdentifier
func (obj *flowIcmpEcho) HasIdentifier() bool {
	return obj.obj.Identifier != nil
}

// description is TBD
// SetIdentifier sets the PatternFlowIcmpEchoIdentifier value in the FlowIcmpEcho object
func (obj *flowIcmpEcho) SetIdentifier(value PatternFlowIcmpEchoIdentifier) FlowIcmpEcho {

	obj.identifierHolder = nil
	obj.obj.Identifier = value.msg()

	return obj
}

// description is TBD
// SequenceNumber returns a PatternFlowIcmpEchoSequenceNumber
func (obj *flowIcmpEcho) SequenceNumber() PatternFlowIcmpEchoSequenceNumber {
	if obj.obj.SequenceNumber == nil {
		obj.obj.SequenceNumber = NewPatternFlowIcmpEchoSequenceNumber().msg()
	}
	if obj.sequenceNumberHolder == nil {
		obj.sequenceNumberHolder = &patternFlowIcmpEchoSequenceNumber{obj: obj.obj.SequenceNumber}
	}
	return obj.sequenceNumberHolder
}

// description is TBD
// SequenceNumber returns a PatternFlowIcmpEchoSequenceNumber
func (obj *flowIcmpEcho) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// description is TBD
// SetSequenceNumber sets the PatternFlowIcmpEchoSequenceNumber value in the FlowIcmpEcho object
func (obj *flowIcmpEcho) SetSequenceNumber(value PatternFlowIcmpEchoSequenceNumber) FlowIcmpEcho {

	obj.sequenceNumberHolder = nil
	obj.obj.SequenceNumber = value.msg()

	return obj
}

func (obj *flowIcmpEcho) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Code != nil {

		obj.Code().validateObj(vObj, set_default)
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

	if obj.obj.Identifier != nil {

		obj.Identifier().validateObj(vObj, set_default)
	}

	if obj.obj.SequenceNumber != nil {

		obj.SequenceNumber().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpEcho) setDefault() {

}
