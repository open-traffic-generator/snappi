package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpTerminator *****
type flowLacpTerminator struct {
	validation
	obj          *otg.FlowLacpTerminator
	marshaller   marshalFlowLacpTerminator
	unMarshaller unMarshalFlowLacpTerminator
	typeHolder   PatternFlowLacpTerminatorType
	lengthHolder PatternFlowLacpTerminatorLength
}

func NewFlowLacpTerminator() FlowLacpTerminator {
	obj := flowLacpTerminator{obj: &otg.FlowLacpTerminator{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpTerminator) msg() *otg.FlowLacpTerminator {
	return obj.obj
}

func (obj *flowLacpTerminator) setMsg(msg *otg.FlowLacpTerminator) FlowLacpTerminator {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpTerminator struct {
	obj *flowLacpTerminator
}

type marshalFlowLacpTerminator interface {
	// ToProto marshals FlowLacpTerminator to protobuf object *otg.FlowLacpTerminator
	ToProto() (*otg.FlowLacpTerminator, error)
	// ToPbText marshals FlowLacpTerminator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpTerminator to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpTerminator to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpTerminator struct {
	obj *flowLacpTerminator
}

type unMarshalFlowLacpTerminator interface {
	// FromProto unmarshals FlowLacpTerminator from protobuf object *otg.FlowLacpTerminator
	FromProto(msg *otg.FlowLacpTerminator) (FlowLacpTerminator, error)
	// FromPbText unmarshals FlowLacpTerminator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpTerminator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpTerminator from JSON text
	FromJson(value string) error
}

func (obj *flowLacpTerminator) Marshal() marshalFlowLacpTerminator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpTerminator{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpTerminator) Unmarshal() unMarshalFlowLacpTerminator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpTerminator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpTerminator) ToProto() (*otg.FlowLacpTerminator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpTerminator) FromProto(msg *otg.FlowLacpTerminator) (FlowLacpTerminator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpTerminator) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpTerminator) FromPbText(value string) error {
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

func (m *marshalflowLacpTerminator) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpTerminator) FromYaml(value string) error {
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

func (m *marshalflowLacpTerminator) ToJson() (string, error) {
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

func (m *unMarshalflowLacpTerminator) FromJson(value string) error {
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

func (obj *flowLacpTerminator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpTerminator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpTerminator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpTerminator) Clone() (FlowLacpTerminator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpTerminator()
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

func (obj *flowLacpTerminator) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpTerminator is marks the end of the LACPDU message.
type FlowLacpTerminator interface {
	Validation
	// msg marshals FlowLacpTerminator to protobuf object *otg.FlowLacpTerminator
	// and doesn't set defaults
	msg() *otg.FlowLacpTerminator
	// setMsg unmarshals FlowLacpTerminator from protobuf object *otg.FlowLacpTerminator
	// and doesn't set defaults
	setMsg(*otg.FlowLacpTerminator) FlowLacpTerminator
	// provides marshal interface
	Marshal() marshalFlowLacpTerminator
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpTerminator
	// validate validates FlowLacpTerminator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpTerminator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpTerminatorType, set in FlowLacpTerminator.
	// PatternFlowLacpTerminatorType is tLV Type for Terminator Information. The value 0x00 indicates the end of the TLV list.
	Type() PatternFlowLacpTerminatorType
	// SetType assigns PatternFlowLacpTerminatorType provided by user to FlowLacpTerminator.
	// PatternFlowLacpTerminatorType is tLV Type for Terminator Information. The value 0x00 indicates the end of the TLV list.
	SetType(value PatternFlowLacpTerminatorType) FlowLacpTerminator
	// HasType checks if Type has been set in FlowLacpTerminator
	HasType() bool
	// Length returns PatternFlowLacpTerminatorLength, set in FlowLacpTerminator.
	// PatternFlowLacpTerminatorLength is the length of the Terminator TLV payload. The value must be 0.
	Length() PatternFlowLacpTerminatorLength
	// SetLength assigns PatternFlowLacpTerminatorLength provided by user to FlowLacpTerminator.
	// PatternFlowLacpTerminatorLength is the length of the Terminator TLV payload. The value must be 0.
	SetLength(value PatternFlowLacpTerminatorLength) FlowLacpTerminator
	// HasLength checks if Length has been set in FlowLacpTerminator
	HasLength() bool
	// Reserved returns string, set in FlowLacpTerminator.
	Reserved() string
	// SetReserved assigns string provided by user to FlowLacpTerminator
	SetReserved(value string) FlowLacpTerminator
	// HasReserved checks if Reserved has been set in FlowLacpTerminator
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpTerminatorType
func (obj *flowLacpTerminator) Type() PatternFlowLacpTerminatorType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpTerminatorType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpTerminatorType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpTerminatorType
func (obj *flowLacpTerminator) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpTerminatorType value in the FlowLacpTerminator object
func (obj *flowLacpTerminator) SetType(value PatternFlowLacpTerminatorType) FlowLacpTerminator {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpTerminatorLength
func (obj *flowLacpTerminator) Length() PatternFlowLacpTerminatorLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpTerminatorLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpTerminatorLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpTerminatorLength
func (obj *flowLacpTerminator) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpTerminatorLength value in the FlowLacpTerminator object
func (obj *flowLacpTerminator) SetLength(value PatternFlowLacpTerminatorLength) FlowLacpTerminator {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Final block of reserved bytes to pad the PDU. It is 50 bytes long.
// Reserved returns a string
func (obj *flowLacpTerminator) Reserved() string {

	return *obj.obj.Reserved

}

// Final block of reserved bytes to pad the PDU. It is 50 bytes long.
// Reserved returns a string
func (obj *flowLacpTerminator) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// Final block of reserved bytes to pad the PDU. It is 50 bytes long.
// SetReserved sets the string value in the FlowLacpTerminator object
func (obj *flowLacpTerminator) SetReserved(value string) FlowLacpTerminator {

	obj.obj.Reserved = &value
	return obj
}

func (obj *flowLacpTerminator) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved != nil {

		err := obj.validateHex(obj.Reserved())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowLacpTerminator.Reserved"))
		}

	}

}

func (obj *flowLacpTerminator) setDefault() {
	if obj.obj.Reserved == nil {
		obj.SetReserved("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	}

}
