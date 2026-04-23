package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmMDNameSpecified *****
type flowCfmMDNameSpecified struct {
	validation
	obj                *otg.FlowCfmMDNameSpecified
	marshaller         marshalFlowCfmMDNameSpecified
	unMarshaller       unMarshalFlowCfmMDNameSpecified
	mdNameLengthHolder PatternFlowCfmMDNameSpecifiedMdNameLength
}

func NewFlowCfmMDNameSpecified() FlowCfmMDNameSpecified {
	obj := flowCfmMDNameSpecified{obj: &otg.FlowCfmMDNameSpecified{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmMDNameSpecified) msg() *otg.FlowCfmMDNameSpecified {
	return obj.obj
}

func (obj *flowCfmMDNameSpecified) setMsg(msg *otg.FlowCfmMDNameSpecified) FlowCfmMDNameSpecified {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmMDNameSpecified struct {
	obj *flowCfmMDNameSpecified
}

type marshalFlowCfmMDNameSpecified interface {
	// ToProto marshals FlowCfmMDNameSpecified to protobuf object *otg.FlowCfmMDNameSpecified
	ToProto() (*otg.FlowCfmMDNameSpecified, error)
	// ToPbText marshals FlowCfmMDNameSpecified to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmMDNameSpecified to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmMDNameSpecified to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmMDNameSpecified struct {
	obj *flowCfmMDNameSpecified
}

type unMarshalFlowCfmMDNameSpecified interface {
	// FromProto unmarshals FlowCfmMDNameSpecified from protobuf object *otg.FlowCfmMDNameSpecified
	FromProto(msg *otg.FlowCfmMDNameSpecified) (FlowCfmMDNameSpecified, error)
	// FromPbText unmarshals FlowCfmMDNameSpecified from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmMDNameSpecified from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmMDNameSpecified from JSON text
	FromJson(value string) error
}

func (obj *flowCfmMDNameSpecified) Marshal() marshalFlowCfmMDNameSpecified {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmMDNameSpecified{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmMDNameSpecified) Unmarshal() unMarshalFlowCfmMDNameSpecified {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmMDNameSpecified{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmMDNameSpecified) ToProto() (*otg.FlowCfmMDNameSpecified, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmMDNameSpecified) FromProto(msg *otg.FlowCfmMDNameSpecified) (FlowCfmMDNameSpecified, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmMDNameSpecified) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmMDNameSpecified) FromPbText(value string) error {
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

func (m *marshalflowCfmMDNameSpecified) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmMDNameSpecified) FromYaml(value string) error {
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

func (m *marshalflowCfmMDNameSpecified) ToJson() (string, error) {
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

func (m *unMarshalflowCfmMDNameSpecified) FromJson(value string) error {
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

func (obj *flowCfmMDNameSpecified) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmMDNameSpecified) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmMDNameSpecified) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmMDNameSpecified) Clone() (FlowCfmMDNameSpecified, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmMDNameSpecified()
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

func (obj *flowCfmMDNameSpecified) setNil() {
	obj.mdNameLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmMDNameSpecified is maintenance domain name
type FlowCfmMDNameSpecified interface {
	Validation
	// msg marshals FlowCfmMDNameSpecified to protobuf object *otg.FlowCfmMDNameSpecified
	// and doesn't set defaults
	msg() *otg.FlowCfmMDNameSpecified
	// setMsg unmarshals FlowCfmMDNameSpecified from protobuf object *otg.FlowCfmMDNameSpecified
	// and doesn't set defaults
	setMsg(*otg.FlowCfmMDNameSpecified) FlowCfmMDNameSpecified
	// provides marshal interface
	Marshal() marshalFlowCfmMDNameSpecified
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmMDNameSpecified
	// validate validates FlowCfmMDNameSpecified
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmMDNameSpecified, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MdNameFormat returns FlowCfmMDNameSpecifiedMdNameFormatEnum, set in FlowCfmMDNameSpecified
	MdNameFormat() FlowCfmMDNameSpecifiedMdNameFormatEnum
	// SetMdNameFormat assigns FlowCfmMDNameSpecifiedMdNameFormatEnum provided by user to FlowCfmMDNameSpecified
	SetMdNameFormat(value FlowCfmMDNameSpecifiedMdNameFormatEnum) FlowCfmMDNameSpecified
	// HasMdNameFormat checks if MdNameFormat has been set in FlowCfmMDNameSpecified
	HasMdNameFormat() bool
	// MdNameLength returns PatternFlowCfmMDNameSpecifiedMdNameLength, set in FlowCfmMDNameSpecified.
	// PatternFlowCfmMDNameSpecifiedMdNameLength is maintenance domain name length
	MdNameLength() PatternFlowCfmMDNameSpecifiedMdNameLength
	// SetMdNameLength assigns PatternFlowCfmMDNameSpecifiedMdNameLength provided by user to FlowCfmMDNameSpecified.
	// PatternFlowCfmMDNameSpecifiedMdNameLength is maintenance domain name length
	SetMdNameLength(value PatternFlowCfmMDNameSpecifiedMdNameLength) FlowCfmMDNameSpecified
	// HasMdNameLength checks if MdNameLength has been set in FlowCfmMDNameSpecified
	HasMdNameLength() bool
	// MdName returns string, set in FlowCfmMDNameSpecified.
	MdName() string
	// SetMdName assigns string provided by user to FlowCfmMDNameSpecified
	SetMdName(value string) FlowCfmMDNameSpecified
	// HasMdName checks if MdName has been set in FlowCfmMDNameSpecified
	HasMdName() bool
	setNil()
}

type FlowCfmMDNameSpecifiedMdNameFormatEnum string

// Enum of MdNameFormat on FlowCfmMDNameSpecified
var FlowCfmMDNameSpecifiedMdNameFormat = struct {
	DOMAIN_NAME_STR FlowCfmMDNameSpecifiedMdNameFormatEnum
	MAC_ADDR        FlowCfmMDNameSpecifiedMdNameFormatEnum
	CHAR_STR        FlowCfmMDNameSpecifiedMdNameFormatEnum
}{
	DOMAIN_NAME_STR: FlowCfmMDNameSpecifiedMdNameFormatEnum("domain_name_str"),
	MAC_ADDR:        FlowCfmMDNameSpecifiedMdNameFormatEnum("mac_addr"),
	CHAR_STR:        FlowCfmMDNameSpecifiedMdNameFormatEnum("char_str"),
}

func (obj *flowCfmMDNameSpecified) MdNameFormat() FlowCfmMDNameSpecifiedMdNameFormatEnum {
	return FlowCfmMDNameSpecifiedMdNameFormatEnum(obj.obj.MdNameFormat.Enum().String())
}

// Maintenance domain name format
// MdNameFormat returns a string
func (obj *flowCfmMDNameSpecified) HasMdNameFormat() bool {
	return obj.obj.MdNameFormat != nil
}

func (obj *flowCfmMDNameSpecified) SetMdNameFormat(value FlowCfmMDNameSpecifiedMdNameFormatEnum) FlowCfmMDNameSpecified {
	intValue, ok := otg.FlowCfmMDNameSpecified_MdNameFormat_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmMDNameSpecifiedMdNameFormatEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmMDNameSpecified_MdNameFormat_Enum(intValue)
	obj.obj.MdNameFormat = &enumValue

	return obj
}

// description is TBD
// MdNameLength returns a PatternFlowCfmMDNameSpecifiedMdNameLength
func (obj *flowCfmMDNameSpecified) MdNameLength() PatternFlowCfmMDNameSpecifiedMdNameLength {
	if obj.obj.MdNameLength == nil {
		obj.obj.MdNameLength = NewPatternFlowCfmMDNameSpecifiedMdNameLength().msg()
	}
	if obj.mdNameLengthHolder == nil {
		obj.mdNameLengthHolder = &patternFlowCfmMDNameSpecifiedMdNameLength{obj: obj.obj.MdNameLength}
	}
	return obj.mdNameLengthHolder
}

// description is TBD
// MdNameLength returns a PatternFlowCfmMDNameSpecifiedMdNameLength
func (obj *flowCfmMDNameSpecified) HasMdNameLength() bool {
	return obj.obj.MdNameLength != nil
}

// description is TBD
// SetMdNameLength sets the PatternFlowCfmMDNameSpecifiedMdNameLength value in the FlowCfmMDNameSpecified object
func (obj *flowCfmMDNameSpecified) SetMdNameLength(value PatternFlowCfmMDNameSpecifiedMdNameLength) FlowCfmMDNameSpecified {

	obj.mdNameLengthHolder = nil
	obj.obj.MdNameLength = value.msg()

	return obj
}

// Maintenance domain name
// MdName returns a string
func (obj *flowCfmMDNameSpecified) MdName() string {

	return *obj.obj.MdName

}

// Maintenance domain name
// MdName returns a string
func (obj *flowCfmMDNameSpecified) HasMdName() bool {
	return obj.obj.MdName != nil
}

// Maintenance domain name
// SetMdName sets the string value in the FlowCfmMDNameSpecified object
func (obj *flowCfmMDNameSpecified) SetMdName(value string) FlowCfmMDNameSpecified {

	obj.obj.MdName = &value
	return obj
}

func (obj *flowCfmMDNameSpecified) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MdNameLength != nil {

		obj.MdNameLength().validateObj(vObj, set_default)
	}

	if obj.obj.MdName != nil {

		err := obj.validateHex(obj.MdName())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmMDNameSpecified.MdName"))
		}

	}

}

func (obj *flowCfmMDNameSpecified) setDefault() {
	if obj.obj.MdNameFormat == nil {
		obj.SetMdNameFormat(FlowCfmMDNameSpecifiedMdNameFormat.CHAR_STR)

	}
	if obj.obj.MdName == nil {
		obj.SetMdName("00")
	}

}
