package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureField *****
type captureField struct {
	validation
	obj          *otg.CaptureField
	marshaller   marshalCaptureField
	unMarshaller unMarshalCaptureField
}

func NewCaptureField() CaptureField {
	obj := captureField{obj: &otg.CaptureField{}}
	obj.setDefault()
	return &obj
}

func (obj *captureField) msg() *otg.CaptureField {
	return obj.obj
}

func (obj *captureField) setMsg(msg *otg.CaptureField) CaptureField {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureField struct {
	obj *captureField
}

type marshalCaptureField interface {
	// ToProto marshals CaptureField to protobuf object *otg.CaptureField
	ToProto() (*otg.CaptureField, error)
	// ToPbText marshals CaptureField to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureField to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureField to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureField struct {
	obj *captureField
}

type unMarshalCaptureField interface {
	// FromProto unmarshals CaptureField from protobuf object *otg.CaptureField
	FromProto(msg *otg.CaptureField) (CaptureField, error)
	// FromPbText unmarshals CaptureField from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureField from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureField from JSON text
	FromJson(value string) error
}

func (obj *captureField) Marshal() marshalCaptureField {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureField{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureField) Unmarshal() unMarshalCaptureField {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureField{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureField) ToProto() (*otg.CaptureField, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureField) FromProto(msg *otg.CaptureField) (CaptureField, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureField) ToPbText() (string, error) {
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

func (m *unMarshalcaptureField) FromPbText(value string) error {
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

func (m *marshalcaptureField) ToYaml() (string, error) {
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

func (m *unMarshalcaptureField) FromYaml(value string) error {
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

func (m *marshalcaptureField) ToJson() (string, error) {
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

func (m *unMarshalcaptureField) FromJson(value string) error {
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

func (obj *captureField) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureField) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureField) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureField) Clone() (CaptureField, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureField()
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

// CaptureField is description is TBD
type CaptureField interface {
	Validation
	// msg marshals CaptureField to protobuf object *otg.CaptureField
	// and doesn't set defaults
	msg() *otg.CaptureField
	// setMsg unmarshals CaptureField from protobuf object *otg.CaptureField
	// and doesn't set defaults
	setMsg(*otg.CaptureField) CaptureField
	// provides marshal interface
	Marshal() marshalCaptureField
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureField
	// validate validates CaptureField
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureField, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns string, set in CaptureField.
	Value() string
	// SetValue assigns string provided by user to CaptureField
	SetValue(value string) CaptureField
	// HasValue checks if Value has been set in CaptureField
	HasValue() bool
	// Mask returns string, set in CaptureField.
	Mask() string
	// SetMask assigns string provided by user to CaptureField
	SetMask(value string) CaptureField
	// HasMask checks if Mask has been set in CaptureField
	HasMask() bool
	// Negate returns bool, set in CaptureField.
	Negate() bool
	// SetNegate assigns bool provided by user to CaptureField
	SetNegate(value bool) CaptureField
	// HasNegate checks if Negate has been set in CaptureField
	HasNegate() bool
}

// description is TBD
// Value returns a string
func (obj *captureField) Value() string {

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *captureField) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the CaptureField object
func (obj *captureField) SetValue(value string) CaptureField {

	obj.obj.Value = &value
	return obj
}

// description is TBD
// Mask returns a string
func (obj *captureField) Mask() string {

	return *obj.obj.Mask

}

// description is TBD
// Mask returns a string
func (obj *captureField) HasMask() bool {
	return obj.obj.Mask != nil
}

// description is TBD
// SetMask sets the string value in the CaptureField object
func (obj *captureField) SetMask(value string) CaptureField {

	obj.obj.Mask = &value
	return obj
}

// description is TBD
// Negate returns a bool
func (obj *captureField) Negate() bool {

	return *obj.obj.Negate

}

// description is TBD
// Negate returns a bool
func (obj *captureField) HasNegate() bool {
	return obj.obj.Negate != nil
}

// description is TBD
// SetNegate sets the bool value in the CaptureField object
func (obj *captureField) SetNegate(value bool) CaptureField {

	obj.obj.Negate = &value
	return obj
}

func (obj *captureField) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on CaptureField.Value"))
		}

	}

	if obj.obj.Mask != nil {

		err := obj.validateHex(obj.Mask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on CaptureField.Mask"))
		}

	}

}

func (obj *captureField) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue("00")
	}
	if obj.obj.Mask == nil {
		obj.SetMask("00")
	}
	if obj.obj.Negate == nil {
		obj.SetNegate(false)
	}

}
