package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingTxOffsetCustom *****
type egressOnlyTrackingTxOffsetCustom struct {
	validation
	obj          *otg.EgressOnlyTrackingTxOffsetCustom
	marshaller   marshalEgressOnlyTrackingTxOffsetCustom
	unMarshaller unMarshalEgressOnlyTrackingTxOffsetCustom
}

func NewEgressOnlyTrackingTxOffsetCustom() EgressOnlyTrackingTxOffsetCustom {
	obj := egressOnlyTrackingTxOffsetCustom{obj: &otg.EgressOnlyTrackingTxOffsetCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingTxOffsetCustom) msg() *otg.EgressOnlyTrackingTxOffsetCustom {
	return obj.obj
}

func (obj *egressOnlyTrackingTxOffsetCustom) setMsg(msg *otg.EgressOnlyTrackingTxOffsetCustom) EgressOnlyTrackingTxOffsetCustom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingTxOffsetCustom struct {
	obj *egressOnlyTrackingTxOffsetCustom
}

type marshalEgressOnlyTrackingTxOffsetCustom interface {
	// ToProto marshals EgressOnlyTrackingTxOffsetCustom to protobuf object *otg.EgressOnlyTrackingTxOffsetCustom
	ToProto() (*otg.EgressOnlyTrackingTxOffsetCustom, error)
	// ToPbText marshals EgressOnlyTrackingTxOffsetCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingTxOffsetCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingTxOffsetCustom to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingTxOffsetCustom struct {
	obj *egressOnlyTrackingTxOffsetCustom
}

type unMarshalEgressOnlyTrackingTxOffsetCustom interface {
	// FromProto unmarshals EgressOnlyTrackingTxOffsetCustom from protobuf object *otg.EgressOnlyTrackingTxOffsetCustom
	FromProto(msg *otg.EgressOnlyTrackingTxOffsetCustom) (EgressOnlyTrackingTxOffsetCustom, error)
	// FromPbText unmarshals EgressOnlyTrackingTxOffsetCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingTxOffsetCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingTxOffsetCustom from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingTxOffsetCustom) Marshal() marshalEgressOnlyTrackingTxOffsetCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingTxOffsetCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingTxOffsetCustom) Unmarshal() unMarshalEgressOnlyTrackingTxOffsetCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingTxOffsetCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingTxOffsetCustom) ToProto() (*otg.EgressOnlyTrackingTxOffsetCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingTxOffsetCustom) FromProto(msg *otg.EgressOnlyTrackingTxOffsetCustom) (EgressOnlyTrackingTxOffsetCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingTxOffsetCustom) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxOffsetCustom) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingTxOffsetCustom) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxOffsetCustom) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingTxOffsetCustom) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxOffsetCustom) FromJson(value string) error {
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

func (obj *egressOnlyTrackingTxOffsetCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTxOffsetCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTxOffsetCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingTxOffsetCustom) Clone() (EgressOnlyTrackingTxOffsetCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingTxOffsetCustom()
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

// EgressOnlyTrackingTxOffsetCustom is a container of custom Tx offset properties.
type EgressOnlyTrackingTxOffsetCustom interface {
	Validation
	// msg marshals EgressOnlyTrackingTxOffsetCustom to protobuf object *otg.EgressOnlyTrackingTxOffsetCustom
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingTxOffsetCustom
	// setMsg unmarshals EgressOnlyTrackingTxOffsetCustom from protobuf object *otg.EgressOnlyTrackingTxOffsetCustom
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingTxOffsetCustom) EgressOnlyTrackingTxOffsetCustom
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingTxOffsetCustom
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingTxOffsetCustom
	// validate validates EgressOnlyTrackingTxOffsetCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingTxOffsetCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in EgressOnlyTrackingTxOffsetCustom.
	Value() uint32
	// SetValue assigns uint32 provided by user to EgressOnlyTrackingTxOffsetCustom
	SetValue(value uint32) EgressOnlyTrackingTxOffsetCustom
	// HasValue checks if Value has been set in EgressOnlyTrackingTxOffsetCustom
	HasValue() bool
}

// Tx Offset in bits relative to start of the packet.
// Value returns a uint32
func (obj *egressOnlyTrackingTxOffsetCustom) Value() uint32 {

	return *obj.obj.Value

}

// Tx Offset in bits relative to start of the packet.
// Value returns a uint32
func (obj *egressOnlyTrackingTxOffsetCustom) HasValue() bool {
	return obj.obj.Value != nil
}

// Tx Offset in bits relative to start of the packet.
// SetValue sets the uint32 value in the EgressOnlyTrackingTxOffsetCustom object
func (obj *egressOnlyTrackingTxOffsetCustom) SetValue(value uint32) EgressOnlyTrackingTxOffsetCustom {

	obj.obj.Value = &value
	return obj
}

func (obj *egressOnlyTrackingTxOffsetCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *egressOnlyTrackingTxOffsetCustom) setDefault() {

}
