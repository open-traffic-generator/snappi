package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowPpp *****
type flowPpp struct {
	validation
	obj                *otg.FlowPpp
	marshaller         marshalFlowPpp
	unMarshaller       unMarshalFlowPpp
	addressHolder      PatternFlowPppAddress
	controlHolder      PatternFlowPppControl
	protocolTypeHolder PatternFlowPppProtocolType
}

func NewFlowPpp() FlowPpp {
	obj := flowPpp{obj: &otg.FlowPpp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowPpp) msg() *otg.FlowPpp {
	return obj.obj
}

func (obj *flowPpp) setMsg(msg *otg.FlowPpp) FlowPpp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowPpp struct {
	obj *flowPpp
}

type marshalFlowPpp interface {
	// ToProto marshals FlowPpp to protobuf object *otg.FlowPpp
	ToProto() (*otg.FlowPpp, error)
	// ToPbText marshals FlowPpp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowPpp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowPpp to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowPpp to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowPpp struct {
	obj *flowPpp
}

type unMarshalFlowPpp interface {
	// FromProto unmarshals FlowPpp from protobuf object *otg.FlowPpp
	FromProto(msg *otg.FlowPpp) (FlowPpp, error)
	// FromPbText unmarshals FlowPpp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowPpp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowPpp from JSON text
	FromJson(value string) error
}

func (obj *flowPpp) Marshal() marshalFlowPpp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowPpp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowPpp) Unmarshal() unMarshalFlowPpp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowPpp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowPpp) ToProto() (*otg.FlowPpp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowPpp) FromProto(msg *otg.FlowPpp) (FlowPpp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowPpp) ToPbText() (string, error) {
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

func (m *unMarshalflowPpp) FromPbText(value string) error {
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

func (m *marshalflowPpp) ToYaml() (string, error) {
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

func (m *unMarshalflowPpp) FromYaml(value string) error {
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

func (m *marshalflowPpp) ToJsonRaw() (string, error) {
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

func (m *marshalflowPpp) ToJson() (string, error) {
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

func (m *unMarshalflowPpp) FromJson(value string) error {
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

func (obj *flowPpp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowPpp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowPpp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowPpp) Clone() (FlowPpp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowPpp()
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

func (obj *flowPpp) setNil() {
	obj.addressHolder = nil
	obj.controlHolder = nil
	obj.protocolTypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowPpp is pPP packet header
type FlowPpp interface {
	Validation
	// msg marshals FlowPpp to protobuf object *otg.FlowPpp
	// and doesn't set defaults
	msg() *otg.FlowPpp
	// setMsg unmarshals FlowPpp from protobuf object *otg.FlowPpp
	// and doesn't set defaults
	setMsg(*otg.FlowPpp) FlowPpp
	// provides marshal interface
	Marshal() marshalFlowPpp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowPpp
	// validate validates FlowPpp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowPpp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Address returns PatternFlowPppAddress, set in FlowPpp.
	// PatternFlowPppAddress is pPP address
	Address() PatternFlowPppAddress
	// SetAddress assigns PatternFlowPppAddress provided by user to FlowPpp.
	// PatternFlowPppAddress is pPP address
	SetAddress(value PatternFlowPppAddress) FlowPpp
	// HasAddress checks if Address has been set in FlowPpp
	HasAddress() bool
	// Control returns PatternFlowPppControl, set in FlowPpp.
	// PatternFlowPppControl is pPP control
	Control() PatternFlowPppControl
	// SetControl assigns PatternFlowPppControl provided by user to FlowPpp.
	// PatternFlowPppControl is pPP control
	SetControl(value PatternFlowPppControl) FlowPpp
	// HasControl checks if Control has been set in FlowPpp
	HasControl() bool
	// ProtocolType returns PatternFlowPppProtocolType, set in FlowPpp.
	// PatternFlowPppProtocolType is pPP protocol type
	ProtocolType() PatternFlowPppProtocolType
	// SetProtocolType assigns PatternFlowPppProtocolType provided by user to FlowPpp.
	// PatternFlowPppProtocolType is pPP protocol type
	SetProtocolType(value PatternFlowPppProtocolType) FlowPpp
	// HasProtocolType checks if ProtocolType has been set in FlowPpp
	HasProtocolType() bool
	setNil()
}

// description is TBD
// Address returns a PatternFlowPppAddress
func (obj *flowPpp) Address() PatternFlowPppAddress {
	if obj.obj.Address == nil {
		obj.obj.Address = NewPatternFlowPppAddress().msg()
	}
	if obj.addressHolder == nil {
		obj.addressHolder = &patternFlowPppAddress{obj: obj.obj.Address}
	}
	return obj.addressHolder
}

// description is TBD
// Address returns a PatternFlowPppAddress
func (obj *flowPpp) HasAddress() bool {
	return obj.obj.Address != nil
}

// description is TBD
// SetAddress sets the PatternFlowPppAddress value in the FlowPpp object
func (obj *flowPpp) SetAddress(value PatternFlowPppAddress) FlowPpp {

	obj.addressHolder = nil
	obj.obj.Address = value.msg()

	return obj
}

// description is TBD
// Control returns a PatternFlowPppControl
func (obj *flowPpp) Control() PatternFlowPppControl {
	if obj.obj.Control == nil {
		obj.obj.Control = NewPatternFlowPppControl().msg()
	}
	if obj.controlHolder == nil {
		obj.controlHolder = &patternFlowPppControl{obj: obj.obj.Control}
	}
	return obj.controlHolder
}

// description is TBD
// Control returns a PatternFlowPppControl
func (obj *flowPpp) HasControl() bool {
	return obj.obj.Control != nil
}

// description is TBD
// SetControl sets the PatternFlowPppControl value in the FlowPpp object
func (obj *flowPpp) SetControl(value PatternFlowPppControl) FlowPpp {

	obj.controlHolder = nil
	obj.obj.Control = value.msg()

	return obj
}

// description is TBD
// ProtocolType returns a PatternFlowPppProtocolType
func (obj *flowPpp) ProtocolType() PatternFlowPppProtocolType {
	if obj.obj.ProtocolType == nil {
		obj.obj.ProtocolType = NewPatternFlowPppProtocolType().msg()
	}
	if obj.protocolTypeHolder == nil {
		obj.protocolTypeHolder = &patternFlowPppProtocolType{obj: obj.obj.ProtocolType}
	}
	return obj.protocolTypeHolder
}

// description is TBD
// ProtocolType returns a PatternFlowPppProtocolType
func (obj *flowPpp) HasProtocolType() bool {
	return obj.obj.ProtocolType != nil
}

// description is TBD
// SetProtocolType sets the PatternFlowPppProtocolType value in the FlowPpp object
func (obj *flowPpp) SetProtocolType(value PatternFlowPppProtocolType) FlowPpp {

	obj.protocolTypeHolder = nil
	obj.obj.ProtocolType = value.msg()

	return obj
}

func (obj *flowPpp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Address != nil {

		obj.Address().validateObj(vObj, set_default)
	}

	if obj.obj.Control != nil {

		obj.Control().validateObj(vObj, set_default)
	}

	if obj.obj.ProtocolType != nil {

		obj.ProtocolType().validateObj(vObj, set_default)
	}

}

func (obj *flowPpp) setDefault() {

}
