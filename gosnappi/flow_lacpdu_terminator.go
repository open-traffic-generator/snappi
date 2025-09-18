package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacpduTerminator *****
type flowLacpduTerminator struct {
	validation
	obj          *otg.FlowLacpduTerminator
	marshaller   marshalFlowLacpduTerminator
	unMarshaller unMarshalFlowLacpduTerminator
	typeHolder   PatternFlowLacpduTerminatorType
	lengthHolder PatternFlowLacpduTerminatorLength
}

func NewFlowLacpduTerminator() FlowLacpduTerminator {
	obj := flowLacpduTerminator{obj: &otg.FlowLacpduTerminator{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacpduTerminator) msg() *otg.FlowLacpduTerminator {
	return obj.obj
}

func (obj *flowLacpduTerminator) setMsg(msg *otg.FlowLacpduTerminator) FlowLacpduTerminator {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacpduTerminator struct {
	obj *flowLacpduTerminator
}

type marshalFlowLacpduTerminator interface {
	// ToProto marshals FlowLacpduTerminator to protobuf object *otg.FlowLacpduTerminator
	ToProto() (*otg.FlowLacpduTerminator, error)
	// ToPbText marshals FlowLacpduTerminator to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacpduTerminator to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacpduTerminator to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacpduTerminator struct {
	obj *flowLacpduTerminator
}

type unMarshalFlowLacpduTerminator interface {
	// FromProto unmarshals FlowLacpduTerminator from protobuf object *otg.FlowLacpduTerminator
	FromProto(msg *otg.FlowLacpduTerminator) (FlowLacpduTerminator, error)
	// FromPbText unmarshals FlowLacpduTerminator from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacpduTerminator from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacpduTerminator from JSON text
	FromJson(value string) error
}

func (obj *flowLacpduTerminator) Marshal() marshalFlowLacpduTerminator {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacpduTerminator{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacpduTerminator) Unmarshal() unMarshalFlowLacpduTerminator {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacpduTerminator{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacpduTerminator) ToProto() (*otg.FlowLacpduTerminator, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacpduTerminator) FromProto(msg *otg.FlowLacpduTerminator) (FlowLacpduTerminator, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacpduTerminator) ToPbText() (string, error) {
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

func (m *unMarshalflowLacpduTerminator) FromPbText(value string) error {
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

func (m *marshalflowLacpduTerminator) ToYaml() (string, error) {
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

func (m *unMarshalflowLacpduTerminator) FromYaml(value string) error {
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

func (m *marshalflowLacpduTerminator) ToJson() (string, error) {
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

func (m *unMarshalflowLacpduTerminator) FromJson(value string) error {
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

func (obj *flowLacpduTerminator) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacpduTerminator) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacpduTerminator) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacpduTerminator) Clone() (FlowLacpduTerminator, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacpduTerminator()
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

func (obj *flowLacpduTerminator) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacpduTerminator is marks the end of the LACPDU message.
type FlowLacpduTerminator interface {
	Validation
	// msg marshals FlowLacpduTerminator to protobuf object *otg.FlowLacpduTerminator
	// and doesn't set defaults
	msg() *otg.FlowLacpduTerminator
	// setMsg unmarshals FlowLacpduTerminator from protobuf object *otg.FlowLacpduTerminator
	// and doesn't set defaults
	setMsg(*otg.FlowLacpduTerminator) FlowLacpduTerminator
	// provides marshal interface
	Marshal() marshalFlowLacpduTerminator
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacpduTerminator
	// validate validates FlowLacpduTerminator
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacpduTerminator, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowLacpduTerminatorType, set in FlowLacpduTerminator.
	// PatternFlowLacpduTerminatorType is tLV Type for Terminator Information. The value 0x00 indicates the end of the TLV list.
	Type() PatternFlowLacpduTerminatorType
	// SetType assigns PatternFlowLacpduTerminatorType provided by user to FlowLacpduTerminator.
	// PatternFlowLacpduTerminatorType is tLV Type for Terminator Information. The value 0x00 indicates the end of the TLV list.
	SetType(value PatternFlowLacpduTerminatorType) FlowLacpduTerminator
	// HasType checks if Type has been set in FlowLacpduTerminator
	HasType() bool
	// Length returns PatternFlowLacpduTerminatorLength, set in FlowLacpduTerminator.
	// PatternFlowLacpduTerminatorLength is the length of the Terminator TLV payload. The value must be 0.
	Length() PatternFlowLacpduTerminatorLength
	// SetLength assigns PatternFlowLacpduTerminatorLength provided by user to FlowLacpduTerminator.
	// PatternFlowLacpduTerminatorLength is the length of the Terminator TLV payload. The value must be 0.
	SetLength(value PatternFlowLacpduTerminatorLength) FlowLacpduTerminator
	// HasLength checks if Length has been set in FlowLacpduTerminator
	HasLength() bool
	// Reserved returns string, set in FlowLacpduTerminator.
	Reserved() string
	// SetReserved assigns string provided by user to FlowLacpduTerminator
	SetReserved(value string) FlowLacpduTerminator
	// HasReserved checks if Reserved has been set in FlowLacpduTerminator
	HasReserved() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowLacpduTerminatorType
func (obj *flowLacpduTerminator) Type() PatternFlowLacpduTerminatorType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowLacpduTerminatorType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowLacpduTerminatorType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowLacpduTerminatorType
func (obj *flowLacpduTerminator) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowLacpduTerminatorType value in the FlowLacpduTerminator object
func (obj *flowLacpduTerminator) SetType(value PatternFlowLacpduTerminatorType) FlowLacpduTerminator {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a PatternFlowLacpduTerminatorLength
func (obj *flowLacpduTerminator) Length() PatternFlowLacpduTerminatorLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowLacpduTerminatorLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowLacpduTerminatorLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowLacpduTerminatorLength
func (obj *flowLacpduTerminator) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowLacpduTerminatorLength value in the FlowLacpduTerminator object
func (obj *flowLacpduTerminator) SetLength(value PatternFlowLacpduTerminatorLength) FlowLacpduTerminator {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Final block of reserved bytes to pad the PDU. It is 50 bytes long.
// Reserved returns a string
func (obj *flowLacpduTerminator) Reserved() string {

	return *obj.obj.Reserved

}

// Final block of reserved bytes to pad the PDU. It is 50 bytes long.
// Reserved returns a string
func (obj *flowLacpduTerminator) HasReserved() bool {
	return obj.obj.Reserved != nil
}

// Final block of reserved bytes to pad the PDU. It is 50 bytes long.
// SetReserved sets the string value in the FlowLacpduTerminator object
func (obj *flowLacpduTerminator) SetReserved(value string) FlowLacpduTerminator {

	obj.obj.Reserved = &value
	return obj
}

func (obj *flowLacpduTerminator) validateObj(vObj *validation, set_default bool) {
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
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowLacpduTerminator.Reserved"))
		}

	}

}

func (obj *flowLacpduTerminator) setDefault() {
	if obj.obj.Reserved == nil {
		obj.SetReserved("0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	}

}
