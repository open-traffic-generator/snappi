package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureCustom *****
type captureCustom struct {
	validation
	obj          *otg.CaptureCustom
	marshaller   marshalCaptureCustom
	unMarshaller unMarshalCaptureCustom
}

func NewCaptureCustom() CaptureCustom {
	obj := captureCustom{obj: &otg.CaptureCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *captureCustom) msg() *otg.CaptureCustom {
	return obj.obj
}

func (obj *captureCustom) setMsg(msg *otg.CaptureCustom) CaptureCustom {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureCustom struct {
	obj *captureCustom
}

type marshalCaptureCustom interface {
	// ToProto marshals CaptureCustom to protobuf object *otg.CaptureCustom
	ToProto() (*otg.CaptureCustom, error)
	// ToPbText marshals CaptureCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureCustom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals CaptureCustom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalcaptureCustom struct {
	obj *captureCustom
}

type unMarshalCaptureCustom interface {
	// FromProto unmarshals CaptureCustom from protobuf object *otg.CaptureCustom
	FromProto(msg *otg.CaptureCustom) (CaptureCustom, error)
	// FromPbText unmarshals CaptureCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureCustom from JSON text
	FromJson(value string) error
}

func (obj *captureCustom) Marshal() marshalCaptureCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureCustom) Unmarshal() unMarshalCaptureCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureCustom) ToProto() (*otg.CaptureCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureCustom) FromProto(msg *otg.CaptureCustom) (CaptureCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureCustom) ToPbText() (string, error) {
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

func (m *unMarshalcaptureCustom) FromPbText(value string) error {
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

func (m *marshalcaptureCustom) ToYaml() (string, error) {
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

func (m *unMarshalcaptureCustom) FromYaml(value string) error {
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

func (m *marshalcaptureCustom) ToJsonRaw() (string, error) {
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

func (m *marshalcaptureCustom) ToJson() (string, error) {
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

func (m *unMarshalcaptureCustom) FromJson(value string) error {
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

func (obj *captureCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureCustom) Clone() (CaptureCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureCustom()
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

// CaptureCustom is description is TBD
type CaptureCustom interface {
	Validation
	// msg marshals CaptureCustom to protobuf object *otg.CaptureCustom
	// and doesn't set defaults
	msg() *otg.CaptureCustom
	// setMsg unmarshals CaptureCustom from protobuf object *otg.CaptureCustom
	// and doesn't set defaults
	setMsg(*otg.CaptureCustom) CaptureCustom
	// provides marshal interface
	Marshal() marshalCaptureCustom
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureCustom
	// validate validates CaptureCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Offset returns uint32, set in CaptureCustom.
	Offset() uint32
	// SetOffset assigns uint32 provided by user to CaptureCustom
	SetOffset(value uint32) CaptureCustom
	// HasOffset checks if Offset has been set in CaptureCustom
	HasOffset() bool
	// BitLength returns uint32, set in CaptureCustom.
	BitLength() uint32
	// SetBitLength assigns uint32 provided by user to CaptureCustom
	SetBitLength(value uint32) CaptureCustom
	// HasBitLength checks if BitLength has been set in CaptureCustom
	HasBitLength() bool
	// Value returns string, set in CaptureCustom.
	Value() string
	// SetValue assigns string provided by user to CaptureCustom
	SetValue(value string) CaptureCustom
	// HasValue checks if Value has been set in CaptureCustom
	HasValue() bool
	// Mask returns string, set in CaptureCustom.
	Mask() string
	// SetMask assigns string provided by user to CaptureCustom
	SetMask(value string) CaptureCustom
	// HasMask checks if Mask has been set in CaptureCustom
	HasMask() bool
	// Negate returns bool, set in CaptureCustom.
	Negate() bool
	// SetNegate assigns bool provided by user to CaptureCustom
	SetNegate(value bool) CaptureCustom
	// HasNegate checks if Negate has been set in CaptureCustom
	HasNegate() bool
}

// The bit offset of field to filter on
// Offset returns a uint32
func (obj *captureCustom) Offset() uint32 {

	return *obj.obj.Offset

}

// The bit offset of field to filter on
// Offset returns a uint32
func (obj *captureCustom) HasOffset() bool {
	return obj.obj.Offset != nil
}

// The bit offset of field to filter on
// SetOffset sets the uint32 value in the CaptureCustom object
func (obj *captureCustom) SetOffset(value uint32) CaptureCustom {

	obj.obj.Offset = &value
	return obj
}

// The bit length of field to filter on
// BitLength returns a uint32
func (obj *captureCustom) BitLength() uint32 {

	return *obj.obj.BitLength

}

// The bit length of field to filter on
// BitLength returns a uint32
func (obj *captureCustom) HasBitLength() bool {
	return obj.obj.BitLength != nil
}

// The bit length of field to filter on
// SetBitLength sets the uint32 value in the CaptureCustom object
func (obj *captureCustom) SetBitLength(value uint32) CaptureCustom {

	obj.obj.BitLength = &value
	return obj
}

// description is TBD
// Value returns a string
func (obj *captureCustom) Value() string {

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *captureCustom) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the CaptureCustom object
func (obj *captureCustom) SetValue(value string) CaptureCustom {

	obj.obj.Value = &value
	return obj
}

// description is TBD
// Mask returns a string
func (obj *captureCustom) Mask() string {

	return *obj.obj.Mask

}

// description is TBD
// Mask returns a string
func (obj *captureCustom) HasMask() bool {
	return obj.obj.Mask != nil
}

// description is TBD
// SetMask sets the string value in the CaptureCustom object
func (obj *captureCustom) SetMask(value string) CaptureCustom {

	obj.obj.Mask = &value
	return obj
}

// description is TBD
// Negate returns a bool
func (obj *captureCustom) Negate() bool {

	return *obj.obj.Negate

}

// description is TBD
// Negate returns a bool
func (obj *captureCustom) HasNegate() bool {
	return obj.obj.Negate != nil
}

// description is TBD
// SetNegate sets the bool value in the CaptureCustom object
func (obj *captureCustom) SetNegate(value bool) CaptureCustom {

	obj.obj.Negate = &value
	return obj
}

func (obj *captureCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateHex(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on CaptureCustom.Value"))
		}

	}

	if obj.obj.Mask != nil {

		err := obj.validateHex(obj.Mask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on CaptureCustom.Mask"))
		}

	}

}

func (obj *captureCustom) setDefault() {
	if obj.obj.BitLength == nil {
		obj.SetBitLength(8)
	}
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
