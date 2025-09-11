package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpMplsLabelBindingsSingleLabelValue *****
type bgpMplsLabelBindingsSingleLabelValue struct {
	validation
	obj          *otg.BgpMplsLabelBindingsSingleLabelValue
	marshaller   marshalBgpMplsLabelBindingsSingleLabelValue
	unMarshaller unMarshalBgpMplsLabelBindingsSingleLabelValue
	valueHolder  RouteMplsLabelValue
}

func NewBgpMplsLabelBindingsSingleLabelValue() BgpMplsLabelBindingsSingleLabelValue {
	obj := bgpMplsLabelBindingsSingleLabelValue{obj: &otg.BgpMplsLabelBindingsSingleLabelValue{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) msg() *otg.BgpMplsLabelBindingsSingleLabelValue {
	return obj.obj
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) setMsg(msg *otg.BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsSingleLabelValue {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpMplsLabelBindingsSingleLabelValue struct {
	obj *bgpMplsLabelBindingsSingleLabelValue
}

type marshalBgpMplsLabelBindingsSingleLabelValue interface {
	// ToProto marshals BgpMplsLabelBindingsSingleLabelValue to protobuf object *otg.BgpMplsLabelBindingsSingleLabelValue
	ToProto() (*otg.BgpMplsLabelBindingsSingleLabelValue, error)
	// ToPbText marshals BgpMplsLabelBindingsSingleLabelValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpMplsLabelBindingsSingleLabelValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpMplsLabelBindingsSingleLabelValue to JSON text
	ToJson() (string, error)
}

type unMarshalbgpMplsLabelBindingsSingleLabelValue struct {
	obj *bgpMplsLabelBindingsSingleLabelValue
}

type unMarshalBgpMplsLabelBindingsSingleLabelValue interface {
	// FromProto unmarshals BgpMplsLabelBindingsSingleLabelValue from protobuf object *otg.BgpMplsLabelBindingsSingleLabelValue
	FromProto(msg *otg.BgpMplsLabelBindingsSingleLabelValue) (BgpMplsLabelBindingsSingleLabelValue, error)
	// FromPbText unmarshals BgpMplsLabelBindingsSingleLabelValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpMplsLabelBindingsSingleLabelValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpMplsLabelBindingsSingleLabelValue from JSON text
	FromJson(value string) error
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) Marshal() marshalBgpMplsLabelBindingsSingleLabelValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpMplsLabelBindingsSingleLabelValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) Unmarshal() unMarshalBgpMplsLabelBindingsSingleLabelValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpMplsLabelBindingsSingleLabelValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpMplsLabelBindingsSingleLabelValue) ToProto() (*otg.BgpMplsLabelBindingsSingleLabelValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpMplsLabelBindingsSingleLabelValue) FromProto(msg *otg.BgpMplsLabelBindingsSingleLabelValue) (BgpMplsLabelBindingsSingleLabelValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpMplsLabelBindingsSingleLabelValue) ToPbText() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindingsSingleLabelValue) FromPbText(value string) error {
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

func (m *marshalbgpMplsLabelBindingsSingleLabelValue) ToYaml() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindingsSingleLabelValue) FromYaml(value string) error {
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

func (m *marshalbgpMplsLabelBindingsSingleLabelValue) ToJson() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindingsSingleLabelValue) FromJson(value string) error {
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

func (obj *bgpMplsLabelBindingsSingleLabelValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) Clone() (BgpMplsLabelBindingsSingleLabelValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpMplsLabelBindingsSingleLabelValue()
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

func (obj *bgpMplsLabelBindingsSingleLabelValue) setNil() {
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpMplsLabelBindingsSingleLabelValue is container for Single MPLS Value.
type BgpMplsLabelBindingsSingleLabelValue interface {
	Validation
	// msg marshals BgpMplsLabelBindingsSingleLabelValue to protobuf object *otg.BgpMplsLabelBindingsSingleLabelValue
	// and doesn't set defaults
	msg() *otg.BgpMplsLabelBindingsSingleLabelValue
	// setMsg unmarshals BgpMplsLabelBindingsSingleLabelValue from protobuf object *otg.BgpMplsLabelBindingsSingleLabelValue
	// and doesn't set defaults
	setMsg(*otg.BgpMplsLabelBindingsSingleLabelValue) BgpMplsLabelBindingsSingleLabelValue
	// provides marshal interface
	Marshal() marshalBgpMplsLabelBindingsSingleLabelValue
	// provides unmarshal interface
	Unmarshal() unMarshalBgpMplsLabelBindingsSingleLabelValue
	// validate validates BgpMplsLabelBindingsSingleLabelValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpMplsLabelBindingsSingleLabelValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns RouteMplsLabelValue, set in BgpMplsLabelBindingsSingleLabelValue.
	// RouteMplsLabelValue is a container of MPLS Prefix Label Value in the address range.
	Value() RouteMplsLabelValue
	// SetValue assigns RouteMplsLabelValue provided by user to BgpMplsLabelBindingsSingleLabelValue.
	// RouteMplsLabelValue is a container of MPLS Prefix Label Value in the address range.
	SetValue(value RouteMplsLabelValue) BgpMplsLabelBindingsSingleLabelValue
	// HasValue checks if Value has been set in BgpMplsLabelBindingsSingleLabelValue
	HasValue() bool
	setNil()
}

// description is TBD
// Value returns a RouteMplsLabelValue
func (obj *bgpMplsLabelBindingsSingleLabelValue) Value() RouteMplsLabelValue {
	if obj.obj.Value == nil {
		obj.obj.Value = NewRouteMplsLabelValue().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &routeMplsLabelValue{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a RouteMplsLabelValue
func (obj *bgpMplsLabelBindingsSingleLabelValue) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the RouteMplsLabelValue value in the BgpMplsLabelBindingsSingleLabelValue object
func (obj *bgpMplsLabelBindingsSingleLabelValue) SetValue(value RouteMplsLabelValue) BgpMplsLabelBindingsSingleLabelValue {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *bgpMplsLabelBindingsSingleLabelValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *bgpMplsLabelBindingsSingleLabelValue) setDefault() {

}
