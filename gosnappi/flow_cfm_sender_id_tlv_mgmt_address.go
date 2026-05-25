package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmSenderIdTlvMgmtAddress *****
type flowCfmSenderIdTlvMgmtAddress struct {
	validation
	obj          *otg.FlowCfmSenderIdTlvMgmtAddress
	marshaller   marshalFlowCfmSenderIdTlvMgmtAddress
	unMarshaller unMarshalFlowCfmSenderIdTlvMgmtAddress
	lengthHolder FlowCfmLength
}

func NewFlowCfmSenderIdTlvMgmtAddress() FlowCfmSenderIdTlvMgmtAddress {
	obj := flowCfmSenderIdTlvMgmtAddress{obj: &otg.FlowCfmSenderIdTlvMgmtAddress{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmSenderIdTlvMgmtAddress) msg() *otg.FlowCfmSenderIdTlvMgmtAddress {
	return obj.obj
}

func (obj *flowCfmSenderIdTlvMgmtAddress) setMsg(msg *otg.FlowCfmSenderIdTlvMgmtAddress) FlowCfmSenderIdTlvMgmtAddress {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmSenderIdTlvMgmtAddress struct {
	obj *flowCfmSenderIdTlvMgmtAddress
}

type marshalFlowCfmSenderIdTlvMgmtAddress interface {
	// ToProto marshals FlowCfmSenderIdTlvMgmtAddress to protobuf object *otg.FlowCfmSenderIdTlvMgmtAddress
	ToProto() (*otg.FlowCfmSenderIdTlvMgmtAddress, error)
	// ToPbText marshals FlowCfmSenderIdTlvMgmtAddress to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmSenderIdTlvMgmtAddress to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmSenderIdTlvMgmtAddress to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmSenderIdTlvMgmtAddress struct {
	obj *flowCfmSenderIdTlvMgmtAddress
}

type unMarshalFlowCfmSenderIdTlvMgmtAddress interface {
	// FromProto unmarshals FlowCfmSenderIdTlvMgmtAddress from protobuf object *otg.FlowCfmSenderIdTlvMgmtAddress
	FromProto(msg *otg.FlowCfmSenderIdTlvMgmtAddress) (FlowCfmSenderIdTlvMgmtAddress, error)
	// FromPbText unmarshals FlowCfmSenderIdTlvMgmtAddress from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmSenderIdTlvMgmtAddress from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmSenderIdTlvMgmtAddress from JSON text
	FromJson(value string) error
}

func (obj *flowCfmSenderIdTlvMgmtAddress) Marshal() marshalFlowCfmSenderIdTlvMgmtAddress {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmSenderIdTlvMgmtAddress{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmSenderIdTlvMgmtAddress) Unmarshal() unMarshalFlowCfmSenderIdTlvMgmtAddress {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmSenderIdTlvMgmtAddress{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmSenderIdTlvMgmtAddress) ToProto() (*otg.FlowCfmSenderIdTlvMgmtAddress, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmSenderIdTlvMgmtAddress) FromProto(msg *otg.FlowCfmSenderIdTlvMgmtAddress) (FlowCfmSenderIdTlvMgmtAddress, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmSenderIdTlvMgmtAddress) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvMgmtAddress) FromPbText(value string) error {
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

func (m *marshalflowCfmSenderIdTlvMgmtAddress) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvMgmtAddress) FromYaml(value string) error {
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

func (m *marshalflowCfmSenderIdTlvMgmtAddress) ToJson() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlvMgmtAddress) FromJson(value string) error {
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

func (obj *flowCfmSenderIdTlvMgmtAddress) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlvMgmtAddress) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlvMgmtAddress) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmSenderIdTlvMgmtAddress) Clone() (FlowCfmSenderIdTlvMgmtAddress, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmSenderIdTlvMgmtAddress()
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

func (obj *flowCfmSenderIdTlvMgmtAddress) setNil() {
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmSenderIdTlvMgmtAddress is sender ID TLV with Management Address
type FlowCfmSenderIdTlvMgmtAddress interface {
	Validation
	// msg marshals FlowCfmSenderIdTlvMgmtAddress to protobuf object *otg.FlowCfmSenderIdTlvMgmtAddress
	// and doesn't set defaults
	msg() *otg.FlowCfmSenderIdTlvMgmtAddress
	// setMsg unmarshals FlowCfmSenderIdTlvMgmtAddress from protobuf object *otg.FlowCfmSenderIdTlvMgmtAddress
	// and doesn't set defaults
	setMsg(*otg.FlowCfmSenderIdTlvMgmtAddress) FlowCfmSenderIdTlvMgmtAddress
	// provides marshal interface
	Marshal() marshalFlowCfmSenderIdTlvMgmtAddress
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmSenderIdTlvMgmtAddress
	// validate validates FlowCfmSenderIdTlvMgmtAddress
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmSenderIdTlvMgmtAddress, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns FlowCfmLength, set in FlowCfmSenderIdTlvMgmtAddress.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	Length() FlowCfmLength
	// SetLength assigns FlowCfmLength provided by user to FlowCfmSenderIdTlvMgmtAddress.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	SetLength(value FlowCfmLength) FlowCfmSenderIdTlvMgmtAddress
	// HasLength checks if Length has been set in FlowCfmSenderIdTlvMgmtAddress
	HasLength() bool
	// Address returns string, set in FlowCfmSenderIdTlvMgmtAddress.
	Address() string
	// SetAddress assigns string provided by user to FlowCfmSenderIdTlvMgmtAddress
	SetAddress(value string) FlowCfmSenderIdTlvMgmtAddress
	// HasAddress checks if Address has been set in FlowCfmSenderIdTlvMgmtAddress
	HasAddress() bool
	setNil()
}

// description is TBD
// Length returns a FlowCfmLength
func (obj *flowCfmSenderIdTlvMgmtAddress) Length() FlowCfmLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowCfmLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowCfmLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a FlowCfmLength
func (obj *flowCfmSenderIdTlvMgmtAddress) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the FlowCfmLength value in the FlowCfmSenderIdTlvMgmtAddress object
func (obj *flowCfmSenderIdTlvMgmtAddress) SetLength(value FlowCfmLength) FlowCfmSenderIdTlvMgmtAddress {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// The address is in the format defined by the domain. This lets the receiver know exactly how to dial back into the sender's management stack.
// Address returns a string
func (obj *flowCfmSenderIdTlvMgmtAddress) Address() string {

	return *obj.obj.Address

}

// The address is in the format defined by the domain. This lets the receiver know exactly how to dial back into the sender's management stack.
// Address returns a string
func (obj *flowCfmSenderIdTlvMgmtAddress) HasAddress() bool {
	return obj.obj.Address != nil
}

// The address is in the format defined by the domain. This lets the receiver know exactly how to dial back into the sender's management stack.
// SetAddress sets the string value in the FlowCfmSenderIdTlvMgmtAddress object
func (obj *flowCfmSenderIdTlvMgmtAddress) SetAddress(value string) FlowCfmSenderIdTlvMgmtAddress {

	obj.obj.Address = &value
	return obj
}

func (obj *flowCfmSenderIdTlvMgmtAddress) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Address != nil {

		err := obj.validateHex(obj.Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmSenderIdTlvMgmtAddress.Address"))
		}

	}

}

func (obj *flowCfmSenderIdTlvMgmtAddress) setDefault() {
	if obj.obj.Address == nil {
		obj.SetAddress("00")
	}

}
