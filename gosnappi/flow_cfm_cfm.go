package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmCfm *****
type flowCfmCfm struct {
	validation
	obj                     *otg.FlowCfmCfm
	marshaller              marshalFlowCfmCfm
	unMarshaller            unMarshalFlowCfmCfm
	mdNameHolder            FlowCfmMDName
	shortMaNameLengthHolder PatternFlowCfmCfmShortMaNameLength
}

func NewFlowCfmCfm() FlowCfmCfm {
	obj := flowCfmCfm{obj: &otg.FlowCfmCfm{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmCfm) msg() *otg.FlowCfmCfm {
	return obj.obj
}

func (obj *flowCfmCfm) setMsg(msg *otg.FlowCfmCfm) FlowCfmCfm {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmCfm struct {
	obj *flowCfmCfm
}

type marshalFlowCfmCfm interface {
	// ToProto marshals FlowCfmCfm to protobuf object *otg.FlowCfmCfm
	ToProto() (*otg.FlowCfmCfm, error)
	// ToPbText marshals FlowCfmCfm to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmCfm to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmCfm to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmCfm struct {
	obj *flowCfmCfm
}

type unMarshalFlowCfmCfm interface {
	// FromProto unmarshals FlowCfmCfm from protobuf object *otg.FlowCfmCfm
	FromProto(msg *otg.FlowCfmCfm) (FlowCfmCfm, error)
	// FromPbText unmarshals FlowCfmCfm from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmCfm from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmCfm from JSON text
	FromJson(value string) error
}

func (obj *flowCfmCfm) Marshal() marshalFlowCfmCfm {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmCfm{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmCfm) Unmarshal() unMarshalFlowCfmCfm {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmCfm{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmCfm) ToProto() (*otg.FlowCfmCfm, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmCfm) FromProto(msg *otg.FlowCfmCfm) (FlowCfmCfm, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmCfm) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmCfm) FromPbText(value string) error {
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

func (m *marshalflowCfmCfm) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmCfm) FromYaml(value string) error {
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

func (m *marshalflowCfmCfm) ToJson() (string, error) {
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

func (m *unMarshalflowCfmCfm) FromJson(value string) error {
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

func (obj *flowCfmCfm) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmCfm) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmCfm) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmCfm) Clone() (FlowCfmCfm, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmCfm()
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

func (obj *flowCfmCfm) setNil() {
	obj.mdNameHolder = nil
	obj.shortMaNameLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmCfm is connectivity Fault Management
type FlowCfmCfm interface {
	Validation
	// msg marshals FlowCfmCfm to protobuf object *otg.FlowCfmCfm
	// and doesn't set defaults
	msg() *otg.FlowCfmCfm
	// setMsg unmarshals FlowCfmCfm from protobuf object *otg.FlowCfmCfm
	// and doesn't set defaults
	setMsg(*otg.FlowCfmCfm) FlowCfmCfm
	// provides marshal interface
	Marshal() marshalFlowCfmCfm
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmCfm
	// validate validates FlowCfmCfm
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmCfm, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MdName returns FlowCfmMDName, set in FlowCfmCfm.
	// FlowCfmMDName is maintenance domain name, can be auto or specified
	MdName() FlowCfmMDName
	// SetMdName assigns FlowCfmMDName provided by user to FlowCfmCfm.
	// FlowCfmMDName is maintenance domain name, can be auto or specified
	SetMdName(value FlowCfmMDName) FlowCfmCfm
	// HasMdName checks if MdName has been set in FlowCfmCfm
	HasMdName() bool
	// ShortMaNameFormat returns FlowCfmCfmShortMaNameFormatEnum, set in FlowCfmCfm
	ShortMaNameFormat() FlowCfmCfmShortMaNameFormatEnum
	// SetShortMaNameFormat assigns FlowCfmCfmShortMaNameFormatEnum provided by user to FlowCfmCfm
	SetShortMaNameFormat(value FlowCfmCfmShortMaNameFormatEnum) FlowCfmCfm
	// HasShortMaNameFormat checks if ShortMaNameFormat has been set in FlowCfmCfm
	HasShortMaNameFormat() bool
	// ShortMaNameLength returns PatternFlowCfmCfmShortMaNameLength, set in FlowCfmCfm.
	// PatternFlowCfmCfmShortMaNameLength is short MA name length
	ShortMaNameLength() PatternFlowCfmCfmShortMaNameLength
	// SetShortMaNameLength assigns PatternFlowCfmCfmShortMaNameLength provided by user to FlowCfmCfm.
	// PatternFlowCfmCfmShortMaNameLength is short MA name length
	SetShortMaNameLength(value PatternFlowCfmCfmShortMaNameLength) FlowCfmCfm
	// HasShortMaNameLength checks if ShortMaNameLength has been set in FlowCfmCfm
	HasShortMaNameLength() bool
	// ShortMaName returns string, set in FlowCfmCfm.
	ShortMaName() string
	// SetShortMaName assigns string provided by user to FlowCfmCfm
	SetShortMaName(value string) FlowCfmCfm
	// HasShortMaName checks if ShortMaName has been set in FlowCfmCfm
	HasShortMaName() bool
	setNil()
}

// description is TBD
// MdName returns a FlowCfmMDName
func (obj *flowCfmCfm) MdName() FlowCfmMDName {
	if obj.obj.MdName == nil {
		obj.obj.MdName = NewFlowCfmMDName().msg()
	}
	if obj.mdNameHolder == nil {
		obj.mdNameHolder = &flowCfmMDName{obj: obj.obj.MdName}
	}
	return obj.mdNameHolder
}

// description is TBD
// MdName returns a FlowCfmMDName
func (obj *flowCfmCfm) HasMdName() bool {
	return obj.obj.MdName != nil
}

// description is TBD
// SetMdName sets the FlowCfmMDName value in the FlowCfmCfm object
func (obj *flowCfmCfm) SetMdName(value FlowCfmMDName) FlowCfmCfm {

	obj.mdNameHolder = nil
	obj.obj.MdName = value.msg()

	return obj
}

type FlowCfmCfmShortMaNameFormatEnum string

// Enum of ShortMaNameFormat on FlowCfmCfm
var FlowCfmCfmShortMaNameFormat = struct {
	PRIMARY_VID     FlowCfmCfmShortMaNameFormatEnum
	CHAR_STR        FlowCfmCfmShortMaNameFormatEnum
	TWO_OCT_INT     FlowCfmCfmShortMaNameFormatEnum
	RFC_2685_VPN_ID FlowCfmCfmShortMaNameFormatEnum
}{
	PRIMARY_VID:     FlowCfmCfmShortMaNameFormatEnum("primary_vid"),
	CHAR_STR:        FlowCfmCfmShortMaNameFormatEnum("char_str"),
	TWO_OCT_INT:     FlowCfmCfmShortMaNameFormatEnum("two_oct_int"),
	RFC_2685_VPN_ID: FlowCfmCfmShortMaNameFormatEnum("rfc_2685_vpn_id"),
}

func (obj *flowCfmCfm) ShortMaNameFormat() FlowCfmCfmShortMaNameFormatEnum {
	return FlowCfmCfmShortMaNameFormatEnum(obj.obj.ShortMaNameFormat.Enum().String())
}

// Short MA name format
// ShortMaNameFormat returns a string
func (obj *flowCfmCfm) HasShortMaNameFormat() bool {
	return obj.obj.ShortMaNameFormat != nil
}

func (obj *flowCfmCfm) SetShortMaNameFormat(value FlowCfmCfmShortMaNameFormatEnum) FlowCfmCfm {
	intValue, ok := otg.FlowCfmCfm_ShortMaNameFormat_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmCfmShortMaNameFormatEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmCfm_ShortMaNameFormat_Enum(intValue)
	obj.obj.ShortMaNameFormat = &enumValue

	return obj
}

// description is TBD
// ShortMaNameLength returns a PatternFlowCfmCfmShortMaNameLength
func (obj *flowCfmCfm) ShortMaNameLength() PatternFlowCfmCfmShortMaNameLength {
	if obj.obj.ShortMaNameLength == nil {
		obj.obj.ShortMaNameLength = NewPatternFlowCfmCfmShortMaNameLength().msg()
	}
	if obj.shortMaNameLengthHolder == nil {
		obj.shortMaNameLengthHolder = &patternFlowCfmCfmShortMaNameLength{obj: obj.obj.ShortMaNameLength}
	}
	return obj.shortMaNameLengthHolder
}

// description is TBD
// ShortMaNameLength returns a PatternFlowCfmCfmShortMaNameLength
func (obj *flowCfmCfm) HasShortMaNameLength() bool {
	return obj.obj.ShortMaNameLength != nil
}

// description is TBD
// SetShortMaNameLength sets the PatternFlowCfmCfmShortMaNameLength value in the FlowCfmCfm object
func (obj *flowCfmCfm) SetShortMaNameLength(value PatternFlowCfmCfmShortMaNameLength) FlowCfmCfm {

	obj.shortMaNameLengthHolder = nil
	obj.obj.ShortMaNameLength = value.msg()

	return obj
}

// Short MA name
// ShortMaName returns a string
func (obj *flowCfmCfm) ShortMaName() string {

	return *obj.obj.ShortMaName

}

// Short MA name
// ShortMaName returns a string
func (obj *flowCfmCfm) HasShortMaName() bool {
	return obj.obj.ShortMaName != nil
}

// Short MA name
// SetShortMaName sets the string value in the FlowCfmCfm object
func (obj *flowCfmCfm) SetShortMaName(value string) FlowCfmCfm {

	obj.obj.ShortMaName = &value
	return obj
}

func (obj *flowCfmCfm) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MdName != nil {

		obj.MdName().validateObj(vObj, set_default)
	}

	if obj.obj.ShortMaNameLength != nil {

		obj.ShortMaNameLength().validateObj(vObj, set_default)
	}

	if obj.obj.ShortMaName != nil {

		err := obj.validateHex(obj.ShortMaName())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmCfm.ShortMaName"))
		}

	}

}

func (obj *flowCfmCfm) setDefault() {
	if obj.obj.ShortMaNameFormat == nil {
		obj.SetShortMaNameFormat(FlowCfmCfmShortMaNameFormat.CHAR_STR)

	}
	if obj.obj.ShortMaName == nil {
		obj.SetShortMaName("00")
	}

}
