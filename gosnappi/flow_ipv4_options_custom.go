package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsCustom *****
type flowIpv4OptionsCustom struct {
	validation
	obj          *otg.FlowIpv4OptionsCustom
	marshaller   marshalFlowIpv4OptionsCustom
	unMarshaller unMarshalFlowIpv4OptionsCustom
	typeHolder   FlowIpv4OptionsCustomType
	lengthHolder FlowIpv4OptionsCustomLength
}

func NewFlowIpv4OptionsCustom() FlowIpv4OptionsCustom {
	obj := flowIpv4OptionsCustom{obj: &otg.FlowIpv4OptionsCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsCustom) msg() *otg.FlowIpv4OptionsCustom {
	return obj.obj
}

func (obj *flowIpv4OptionsCustom) setMsg(msg *otg.FlowIpv4OptionsCustom) FlowIpv4OptionsCustom {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsCustom struct {
	obj *flowIpv4OptionsCustom
}

type marshalFlowIpv4OptionsCustom interface {
	// ToProto marshals FlowIpv4OptionsCustom to protobuf object *otg.FlowIpv4OptionsCustom
	ToProto() (*otg.FlowIpv4OptionsCustom, error)
	// ToPbText marshals FlowIpv4OptionsCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsCustom to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsCustom struct {
	obj *flowIpv4OptionsCustom
}

type unMarshalFlowIpv4OptionsCustom interface {
	// FromProto unmarshals FlowIpv4OptionsCustom from protobuf object *otg.FlowIpv4OptionsCustom
	FromProto(msg *otg.FlowIpv4OptionsCustom) (FlowIpv4OptionsCustom, error)
	// FromPbText unmarshals FlowIpv4OptionsCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsCustom from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsCustom) Marshal() marshalFlowIpv4OptionsCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsCustom) Unmarshal() unMarshalFlowIpv4OptionsCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsCustom) ToProto() (*otg.FlowIpv4OptionsCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsCustom) FromProto(msg *otg.FlowIpv4OptionsCustom) (FlowIpv4OptionsCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsCustom) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustom) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsCustom) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustom) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsCustom) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustom) FromJson(value string) error {
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

func (obj *flowIpv4OptionsCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsCustom) Clone() (FlowIpv4OptionsCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsCustom()
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

func (obj *flowIpv4OptionsCustom) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsCustom is user defined IP options to be appended to the IPv4 header.
type FlowIpv4OptionsCustom interface {
	Validation
	// msg marshals FlowIpv4OptionsCustom to protobuf object *otg.FlowIpv4OptionsCustom
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsCustom
	// setMsg unmarshals FlowIpv4OptionsCustom from protobuf object *otg.FlowIpv4OptionsCustom
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsCustom) FlowIpv4OptionsCustom
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsCustom
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsCustom
	// validate validates FlowIpv4OptionsCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns FlowIpv4OptionsCustomType, set in FlowIpv4OptionsCustom.
	// FlowIpv4OptionsCustomType is type options for custom options.
	Type() FlowIpv4OptionsCustomType
	// SetType assigns FlowIpv4OptionsCustomType provided by user to FlowIpv4OptionsCustom.
	// FlowIpv4OptionsCustomType is type options for custom options.
	SetType(value FlowIpv4OptionsCustomType) FlowIpv4OptionsCustom
	// HasType checks if Type has been set in FlowIpv4OptionsCustom
	HasType() bool
	// Length returns FlowIpv4OptionsCustomLength, set in FlowIpv4OptionsCustom.
	// FlowIpv4OptionsCustomLength is length for custom options.
	Length() FlowIpv4OptionsCustomLength
	// SetLength assigns FlowIpv4OptionsCustomLength provided by user to FlowIpv4OptionsCustom.
	// FlowIpv4OptionsCustomLength is length for custom options.
	SetLength(value FlowIpv4OptionsCustomLength) FlowIpv4OptionsCustom
	// HasLength checks if Length has been set in FlowIpv4OptionsCustom
	HasLength() bool
	// Value returns string, set in FlowIpv4OptionsCustom.
	Value() string
	// SetValue assigns string provided by user to FlowIpv4OptionsCustom
	SetValue(value string) FlowIpv4OptionsCustom
	// HasValue checks if Value has been set in FlowIpv4OptionsCustom
	HasValue() bool
	setNil()
}

// description is TBD
// Type returns a FlowIpv4OptionsCustomType
func (obj *flowIpv4OptionsCustom) Type() FlowIpv4OptionsCustomType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewFlowIpv4OptionsCustomType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &flowIpv4OptionsCustomType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a FlowIpv4OptionsCustomType
func (obj *flowIpv4OptionsCustom) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the FlowIpv4OptionsCustomType value in the FlowIpv4OptionsCustom object
func (obj *flowIpv4OptionsCustom) SetType(value FlowIpv4OptionsCustomType) FlowIpv4OptionsCustom {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a FlowIpv4OptionsCustomLength
func (obj *flowIpv4OptionsCustom) Length() FlowIpv4OptionsCustomLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowIpv4OptionsCustomLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowIpv4OptionsCustomLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a FlowIpv4OptionsCustomLength
func (obj *flowIpv4OptionsCustom) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the FlowIpv4OptionsCustomLength value in the FlowIpv4OptionsCustom object
func (obj *flowIpv4OptionsCustom) SetLength(value FlowIpv4OptionsCustomLength) FlowIpv4OptionsCustom {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Value of the option field should not excced 38 bytes since maximum 40 bytes can be added as options in IPv4 header. For type and length requires 2 bytes, hence maximum of 38 bytes are expected. Maximum length of this attribute is 76 (38 * 2 hex character per byte).
// Value returns a string
func (obj *flowIpv4OptionsCustom) Value() string {

	return *obj.obj.Value

}

// Value of the option field should not excced 38 bytes since maximum 40 bytes can be added as options in IPv4 header. For type and length requires 2 bytes, hence maximum of 38 bytes are expected. Maximum length of this attribute is 76 (38 * 2 hex character per byte).
// Value returns a string
func (obj *flowIpv4OptionsCustom) HasValue() bool {
	return obj.obj.Value != nil
}

// Value of the option field should not excced 38 bytes since maximum 40 bytes can be added as options in IPv4 header. For type and length requires 2 bytes, hence maximum of 38 bytes are expected. Maximum length of this attribute is 76 (38 * 2 hex character per byte).
// SetValue sets the string value in the FlowIpv4OptionsCustom object
func (obj *flowIpv4OptionsCustom) SetValue(value string) FlowIpv4OptionsCustom {

	obj.obj.Value = &value
	return obj
}

func (obj *flowIpv4OptionsCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Value != nil {

		if len(*obj.obj.Value) < 0 || len(*obj.obj.Value) > 76 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"0 <= length of FlowIpv4OptionsCustom.Value <= 76 but Got %d",
					len(*obj.obj.Value)))
		}

	}

}

func (obj *flowIpv4OptionsCustom) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue("0000")
	}

}
