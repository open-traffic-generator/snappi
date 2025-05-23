package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2PriorityValue *****
type rocev2PriorityValue struct {
	validation
	obj          *otg.Rocev2PriorityValue
	marshaller   marshalRocev2PriorityValue
	unMarshaller unMarshalRocev2PriorityValue
}

func NewRocev2PriorityValue() Rocev2PriorityValue {
	obj := rocev2PriorityValue{obj: &otg.Rocev2PriorityValue{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2PriorityValue) msg() *otg.Rocev2PriorityValue {
	return obj.obj
}

func (obj *rocev2PriorityValue) setMsg(msg *otg.Rocev2PriorityValue) Rocev2PriorityValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2PriorityValue struct {
	obj *rocev2PriorityValue
}

type marshalRocev2PriorityValue interface {
	// ToProto marshals Rocev2PriorityValue to protobuf object *otg.Rocev2PriorityValue
	ToProto() (*otg.Rocev2PriorityValue, error)
	// ToPbText marshals Rocev2PriorityValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2PriorityValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2PriorityValue to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2PriorityValue to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2PriorityValue struct {
	obj *rocev2PriorityValue
}

type unMarshalRocev2PriorityValue interface {
	// FromProto unmarshals Rocev2PriorityValue from protobuf object *otg.Rocev2PriorityValue
	FromProto(msg *otg.Rocev2PriorityValue) (Rocev2PriorityValue, error)
	// FromPbText unmarshals Rocev2PriorityValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2PriorityValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2PriorityValue from JSON text
	FromJson(value string) error
}

func (obj *rocev2PriorityValue) Marshal() marshalRocev2PriorityValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2PriorityValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2PriorityValue) Unmarshal() unMarshalRocev2PriorityValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2PriorityValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2PriorityValue) ToProto() (*otg.Rocev2PriorityValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2PriorityValue) FromProto(msg *otg.Rocev2PriorityValue) (Rocev2PriorityValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2PriorityValue) ToPbText() (string, error) {
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

func (m *unMarshalrocev2PriorityValue) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalrocev2PriorityValue) ToYaml() (string, error) {
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

func (m *unMarshalrocev2PriorityValue) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalrocev2PriorityValue) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2PriorityValue) ToJson() (string, error) {
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

func (m *unMarshalrocev2PriorityValue) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *rocev2PriorityValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2PriorityValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2PriorityValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2PriorityValue) Clone() (Rocev2PriorityValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2PriorityValue()
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

// Rocev2PriorityValue is priority value for CNP, ACK or NAK packets.
type Rocev2PriorityValue interface {
	Validation
	// msg marshals Rocev2PriorityValue to protobuf object *otg.Rocev2PriorityValue
	// and doesn't set defaults
	msg() *otg.Rocev2PriorityValue
	// setMsg unmarshals Rocev2PriorityValue from protobuf object *otg.Rocev2PriorityValue
	// and doesn't set defaults
	setMsg(*otg.Rocev2PriorityValue) Rocev2PriorityValue
	// provides marshal interface
	Marshal() marshalRocev2PriorityValue
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2PriorityValue
	// validate validates Rocev2PriorityValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2PriorityValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in Rocev2PriorityValue.
	Value() uint32
	// SetValue assigns uint32 provided by user to Rocev2PriorityValue
	SetValue(value uint32) Rocev2PriorityValue
	// HasValue checks if Value has been set in Rocev2PriorityValue
	HasValue() bool
}

// description is TBD
// Value returns a uint32
func (obj *rocev2PriorityValue) Value() uint32 {

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *rocev2PriorityValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the Rocev2PriorityValue object
func (obj *rocev2PriorityValue) SetValue(value uint32) Rocev2PriorityValue {

	obj.obj.Value = &value
	return obj
}

func (obj *rocev2PriorityValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2PriorityValue.Value <= 63 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *rocev2PriorityValue) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(48)
	}

}
