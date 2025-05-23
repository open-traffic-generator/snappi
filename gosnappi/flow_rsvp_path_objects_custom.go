package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsCustom *****
type flowRSVPPathObjectsCustom struct {
	validation
	obj          *otg.FlowRSVPPathObjectsCustom
	marshaller   marshalFlowRSVPPathObjectsCustom
	unMarshaller unMarshalFlowRSVPPathObjectsCustom
	typeHolder   PatternFlowRSVPPathObjectsCustomType
	lengthHolder FlowRSVPObjectLength
}

func NewFlowRSVPPathObjectsCustom() FlowRSVPPathObjectsCustom {
	obj := flowRSVPPathObjectsCustom{obj: &otg.FlowRSVPPathObjectsCustom{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsCustom) msg() *otg.FlowRSVPPathObjectsCustom {
	return obj.obj
}

func (obj *flowRSVPPathObjectsCustom) setMsg(msg *otg.FlowRSVPPathObjectsCustom) FlowRSVPPathObjectsCustom {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsCustom struct {
	obj *flowRSVPPathObjectsCustom
}

type marshalFlowRSVPPathObjectsCustom interface {
	// ToProto marshals FlowRSVPPathObjectsCustom to protobuf object *otg.FlowRSVPPathObjectsCustom
	ToProto() (*otg.FlowRSVPPathObjectsCustom, error)
	// ToPbText marshals FlowRSVPPathObjectsCustom to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsCustom to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsCustom to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals FlowRSVPPathObjectsCustom to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalflowRSVPPathObjectsCustom struct {
	obj *flowRSVPPathObjectsCustom
}

type unMarshalFlowRSVPPathObjectsCustom interface {
	// FromProto unmarshals FlowRSVPPathObjectsCustom from protobuf object *otg.FlowRSVPPathObjectsCustom
	FromProto(msg *otg.FlowRSVPPathObjectsCustom) (FlowRSVPPathObjectsCustom, error)
	// FromPbText unmarshals FlowRSVPPathObjectsCustom from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsCustom from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsCustom from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsCustom) Marshal() marshalFlowRSVPPathObjectsCustom {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsCustom{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsCustom) Unmarshal() unMarshalFlowRSVPPathObjectsCustom {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsCustom{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsCustom) ToProto() (*otg.FlowRSVPPathObjectsCustom, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsCustom) FromProto(msg *otg.FlowRSVPPathObjectsCustom) (FlowRSVPPathObjectsCustom, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsCustom) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsCustom) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsCustom) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsCustom) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsCustom) ToJsonRaw() (string, error) {
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

func (m *marshalflowRSVPPathObjectsCustom) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsCustom) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsCustom) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsCustom) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsCustom) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsCustom) Clone() (FlowRSVPPathObjectsCustom, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsCustom()
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

func (obj *flowRSVPPathObjectsCustom) setNil() {
	obj.typeHolder = nil
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsCustom is custom packet header
type FlowRSVPPathObjectsCustom interface {
	Validation
	// msg marshals FlowRSVPPathObjectsCustom to protobuf object *otg.FlowRSVPPathObjectsCustom
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsCustom
	// setMsg unmarshals FlowRSVPPathObjectsCustom from protobuf object *otg.FlowRSVPPathObjectsCustom
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsCustom) FlowRSVPPathObjectsCustom
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsCustom
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsCustom
	// validate validates FlowRSVPPathObjectsCustom
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsCustom, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns PatternFlowRSVPPathObjectsCustomType, set in FlowRSVPPathObjectsCustom.
	// PatternFlowRSVPPathObjectsCustomType is user defined object type.
	Type() PatternFlowRSVPPathObjectsCustomType
	// SetType assigns PatternFlowRSVPPathObjectsCustomType provided by user to FlowRSVPPathObjectsCustom.
	// PatternFlowRSVPPathObjectsCustomType is user defined object type.
	SetType(value PatternFlowRSVPPathObjectsCustomType) FlowRSVPPathObjectsCustom
	// HasType checks if Type has been set in FlowRSVPPathObjectsCustom
	HasType() bool
	// Length returns FlowRSVPObjectLength, set in FlowRSVPPathObjectsCustom.
	// FlowRSVPObjectLength is description is TBD
	Length() FlowRSVPObjectLength
	// SetLength assigns FlowRSVPObjectLength provided by user to FlowRSVPPathObjectsCustom.
	// FlowRSVPObjectLength is description is TBD
	SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsCustom
	// HasLength checks if Length has been set in FlowRSVPPathObjectsCustom
	HasLength() bool
	// Bytes returns string, set in FlowRSVPPathObjectsCustom.
	Bytes() string
	// SetBytes assigns string provided by user to FlowRSVPPathObjectsCustom
	SetBytes(value string) FlowRSVPPathObjectsCustom
	// HasBytes checks if Bytes has been set in FlowRSVPPathObjectsCustom
	HasBytes() bool
	setNil()
}

// description is TBD
// Type returns a PatternFlowRSVPPathObjectsCustomType
func (obj *flowRSVPPathObjectsCustom) Type() PatternFlowRSVPPathObjectsCustomType {
	if obj.obj.Type == nil {
		obj.obj.Type = NewPatternFlowRSVPPathObjectsCustomType().msg()
	}
	if obj.typeHolder == nil {
		obj.typeHolder = &patternFlowRSVPPathObjectsCustomType{obj: obj.obj.Type}
	}
	return obj.typeHolder
}

// description is TBD
// Type returns a PatternFlowRSVPPathObjectsCustomType
func (obj *flowRSVPPathObjectsCustom) HasType() bool {
	return obj.obj.Type != nil
}

// description is TBD
// SetType sets the PatternFlowRSVPPathObjectsCustomType value in the FlowRSVPPathObjectsCustom object
func (obj *flowRSVPPathObjectsCustom) SetType(value PatternFlowRSVPPathObjectsCustomType) FlowRSVPPathObjectsCustom {

	obj.typeHolder = nil
	obj.obj.Type = value.msg()

	return obj
}

// description is TBD
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsCustom) Length() FlowRSVPObjectLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowRSVPObjectLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowRSVPObjectLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a FlowRSVPObjectLength
func (obj *flowRSVPPathObjectsCustom) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the FlowRSVPObjectLength value in the FlowRSVPPathObjectsCustom object
func (obj *flowRSVPPathObjectsCustom) SetLength(value FlowRSVPObjectLength) FlowRSVPPathObjectsCustom {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// A custom packet header defined as a string of hex bytes. The string MUST contain sequence of valid hex bytes. Spaces or colons can be part of the bytes but will be discarded. Value of the this field should not excced 65525 bytes since maximum 65528 bytes can be added as object-contents in RSVP header. For type and length requires 3 bytes, hence maximum of 65524 bytes are expected. Maximum length of this attribute is 131050 (65525 * 2 hex character per byte).
// Bytes returns a string
func (obj *flowRSVPPathObjectsCustom) Bytes() string {

	return *obj.obj.Bytes

}

// A custom packet header defined as a string of hex bytes. The string MUST contain sequence of valid hex bytes. Spaces or colons can be part of the bytes but will be discarded. Value of the this field should not excced 65525 bytes since maximum 65528 bytes can be added as object-contents in RSVP header. For type and length requires 3 bytes, hence maximum of 65524 bytes are expected. Maximum length of this attribute is 131050 (65525 * 2 hex character per byte).
// Bytes returns a string
func (obj *flowRSVPPathObjectsCustom) HasBytes() bool {
	return obj.obj.Bytes != nil
}

// A custom packet header defined as a string of hex bytes. The string MUST contain sequence of valid hex bytes. Spaces or colons can be part of the bytes but will be discarded. Value of the this field should not excced 65525 bytes since maximum 65528 bytes can be added as object-contents in RSVP header. For type and length requires 3 bytes, hence maximum of 65524 bytes are expected. Maximum length of this attribute is 131050 (65525 * 2 hex character per byte).
// SetBytes sets the string value in the FlowRSVPPathObjectsCustom object
func (obj *flowRSVPPathObjectsCustom) SetBytes(value string) FlowRSVPPathObjectsCustom {

	obj.obj.Bytes = &value
	return obj
}

func (obj *flowRSVPPathObjectsCustom) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Type != nil {

		obj.Type().validateObj(vObj, set_default)
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.Bytes != nil {

		if len(*obj.obj.Bytes) < 1 || len(*obj.obj.Bytes) > 131050 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of FlowRSVPPathObjectsCustom.Bytes <= 131050 but Got %d",
					len(*obj.obj.Bytes)))
		}

	}

}

func (obj *flowRSVPPathObjectsCustom) setDefault() {
	if obj.obj.Bytes == nil {
		obj.SetBytes("0000")
	}

}
