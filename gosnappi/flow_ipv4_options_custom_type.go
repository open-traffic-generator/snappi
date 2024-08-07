package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv4OptionsCustomType *****
type flowIpv4OptionsCustomType struct {
	validation
	obj                *otg.FlowIpv4OptionsCustomType
	marshaller         marshalFlowIpv4OptionsCustomType
	unMarshaller       unMarshalFlowIpv4OptionsCustomType
	copiedFlagHolder   PatternFlowIpv4OptionsCustomTypeCopiedFlag
	optionClassHolder  PatternFlowIpv4OptionsCustomTypeOptionClass
	optionNumberHolder PatternFlowIpv4OptionsCustomTypeOptionNumber
}

func NewFlowIpv4OptionsCustomType() FlowIpv4OptionsCustomType {
	obj := flowIpv4OptionsCustomType{obj: &otg.FlowIpv4OptionsCustomType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv4OptionsCustomType) msg() *otg.FlowIpv4OptionsCustomType {
	return obj.obj
}

func (obj *flowIpv4OptionsCustomType) setMsg(msg *otg.FlowIpv4OptionsCustomType) FlowIpv4OptionsCustomType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv4OptionsCustomType struct {
	obj *flowIpv4OptionsCustomType
}

type marshalFlowIpv4OptionsCustomType interface {
	// ToProto marshals FlowIpv4OptionsCustomType to protobuf object *otg.FlowIpv4OptionsCustomType
	ToProto() (*otg.FlowIpv4OptionsCustomType, error)
	// ToPbText marshals FlowIpv4OptionsCustomType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv4OptionsCustomType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv4OptionsCustomType to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv4OptionsCustomType struct {
	obj *flowIpv4OptionsCustomType
}

type unMarshalFlowIpv4OptionsCustomType interface {
	// FromProto unmarshals FlowIpv4OptionsCustomType from protobuf object *otg.FlowIpv4OptionsCustomType
	FromProto(msg *otg.FlowIpv4OptionsCustomType) (FlowIpv4OptionsCustomType, error)
	// FromPbText unmarshals FlowIpv4OptionsCustomType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv4OptionsCustomType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv4OptionsCustomType from JSON text
	FromJson(value string) error
}

func (obj *flowIpv4OptionsCustomType) Marshal() marshalFlowIpv4OptionsCustomType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv4OptionsCustomType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv4OptionsCustomType) Unmarshal() unMarshalFlowIpv4OptionsCustomType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv4OptionsCustomType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv4OptionsCustomType) ToProto() (*otg.FlowIpv4OptionsCustomType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv4OptionsCustomType) FromProto(msg *otg.FlowIpv4OptionsCustomType) (FlowIpv4OptionsCustomType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv4OptionsCustomType) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustomType) FromPbText(value string) error {
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

func (m *marshalflowIpv4OptionsCustomType) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustomType) FromYaml(value string) error {
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

func (m *marshalflowIpv4OptionsCustomType) ToJson() (string, error) {
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

func (m *unMarshalflowIpv4OptionsCustomType) FromJson(value string) error {
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

func (obj *flowIpv4OptionsCustomType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsCustomType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv4OptionsCustomType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv4OptionsCustomType) Clone() (FlowIpv4OptionsCustomType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv4OptionsCustomType()
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

func (obj *flowIpv4OptionsCustomType) setNil() {
	obj.copiedFlagHolder = nil
	obj.optionClassHolder = nil
	obj.optionNumberHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv4OptionsCustomType is type options for custom options.
type FlowIpv4OptionsCustomType interface {
	Validation
	// msg marshals FlowIpv4OptionsCustomType to protobuf object *otg.FlowIpv4OptionsCustomType
	// and doesn't set defaults
	msg() *otg.FlowIpv4OptionsCustomType
	// setMsg unmarshals FlowIpv4OptionsCustomType from protobuf object *otg.FlowIpv4OptionsCustomType
	// and doesn't set defaults
	setMsg(*otg.FlowIpv4OptionsCustomType) FlowIpv4OptionsCustomType
	// provides marshal interface
	Marshal() marshalFlowIpv4OptionsCustomType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv4OptionsCustomType
	// validate validates FlowIpv4OptionsCustomType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv4OptionsCustomType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CopiedFlag returns PatternFlowIpv4OptionsCustomTypeCopiedFlag, set in FlowIpv4OptionsCustomType.
	// PatternFlowIpv4OptionsCustomTypeCopiedFlag is this flag indicates this option is copied to all fragments on fragmentations.
	CopiedFlag() PatternFlowIpv4OptionsCustomTypeCopiedFlag
	// SetCopiedFlag assigns PatternFlowIpv4OptionsCustomTypeCopiedFlag provided by user to FlowIpv4OptionsCustomType.
	// PatternFlowIpv4OptionsCustomTypeCopiedFlag is this flag indicates this option is copied to all fragments on fragmentations.
	SetCopiedFlag(value PatternFlowIpv4OptionsCustomTypeCopiedFlag) FlowIpv4OptionsCustomType
	// HasCopiedFlag checks if CopiedFlag has been set in FlowIpv4OptionsCustomType
	HasCopiedFlag() bool
	// OptionClass returns PatternFlowIpv4OptionsCustomTypeOptionClass, set in FlowIpv4OptionsCustomType.
	// PatternFlowIpv4OptionsCustomTypeOptionClass is option class [Ref:https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml#ip-parameters-1].
	OptionClass() PatternFlowIpv4OptionsCustomTypeOptionClass
	// SetOptionClass assigns PatternFlowIpv4OptionsCustomTypeOptionClass provided by user to FlowIpv4OptionsCustomType.
	// PatternFlowIpv4OptionsCustomTypeOptionClass is option class [Ref:https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml#ip-parameters-1].
	SetOptionClass(value PatternFlowIpv4OptionsCustomTypeOptionClass) FlowIpv4OptionsCustomType
	// HasOptionClass checks if OptionClass has been set in FlowIpv4OptionsCustomType
	HasOptionClass() bool
	// OptionNumber returns PatternFlowIpv4OptionsCustomTypeOptionNumber, set in FlowIpv4OptionsCustomType.
	// PatternFlowIpv4OptionsCustomTypeOptionNumber is option Number [Ref:https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml#ip-parameters-1].
	OptionNumber() PatternFlowIpv4OptionsCustomTypeOptionNumber
	// SetOptionNumber assigns PatternFlowIpv4OptionsCustomTypeOptionNumber provided by user to FlowIpv4OptionsCustomType.
	// PatternFlowIpv4OptionsCustomTypeOptionNumber is option Number [Ref:https://www.iana.org/assignments/ip-parameters/ip-parameters.xhtml#ip-parameters-1].
	SetOptionNumber(value PatternFlowIpv4OptionsCustomTypeOptionNumber) FlowIpv4OptionsCustomType
	// HasOptionNumber checks if OptionNumber has been set in FlowIpv4OptionsCustomType
	HasOptionNumber() bool
	setNil()
}

// description is TBD
// CopiedFlag returns a PatternFlowIpv4OptionsCustomTypeCopiedFlag
func (obj *flowIpv4OptionsCustomType) CopiedFlag() PatternFlowIpv4OptionsCustomTypeCopiedFlag {
	if obj.obj.CopiedFlag == nil {
		obj.obj.CopiedFlag = NewPatternFlowIpv4OptionsCustomTypeCopiedFlag().msg()
	}
	if obj.copiedFlagHolder == nil {
		obj.copiedFlagHolder = &patternFlowIpv4OptionsCustomTypeCopiedFlag{obj: obj.obj.CopiedFlag}
	}
	return obj.copiedFlagHolder
}

// description is TBD
// CopiedFlag returns a PatternFlowIpv4OptionsCustomTypeCopiedFlag
func (obj *flowIpv4OptionsCustomType) HasCopiedFlag() bool {
	return obj.obj.CopiedFlag != nil
}

// description is TBD
// SetCopiedFlag sets the PatternFlowIpv4OptionsCustomTypeCopiedFlag value in the FlowIpv4OptionsCustomType object
func (obj *flowIpv4OptionsCustomType) SetCopiedFlag(value PatternFlowIpv4OptionsCustomTypeCopiedFlag) FlowIpv4OptionsCustomType {

	obj.copiedFlagHolder = nil
	obj.obj.CopiedFlag = value.msg()

	return obj
}

// description is TBD
// OptionClass returns a PatternFlowIpv4OptionsCustomTypeOptionClass
func (obj *flowIpv4OptionsCustomType) OptionClass() PatternFlowIpv4OptionsCustomTypeOptionClass {
	if obj.obj.OptionClass == nil {
		obj.obj.OptionClass = NewPatternFlowIpv4OptionsCustomTypeOptionClass().msg()
	}
	if obj.optionClassHolder == nil {
		obj.optionClassHolder = &patternFlowIpv4OptionsCustomTypeOptionClass{obj: obj.obj.OptionClass}
	}
	return obj.optionClassHolder
}

// description is TBD
// OptionClass returns a PatternFlowIpv4OptionsCustomTypeOptionClass
func (obj *flowIpv4OptionsCustomType) HasOptionClass() bool {
	return obj.obj.OptionClass != nil
}

// description is TBD
// SetOptionClass sets the PatternFlowIpv4OptionsCustomTypeOptionClass value in the FlowIpv4OptionsCustomType object
func (obj *flowIpv4OptionsCustomType) SetOptionClass(value PatternFlowIpv4OptionsCustomTypeOptionClass) FlowIpv4OptionsCustomType {

	obj.optionClassHolder = nil
	obj.obj.OptionClass = value.msg()

	return obj
}

// description is TBD
// OptionNumber returns a PatternFlowIpv4OptionsCustomTypeOptionNumber
func (obj *flowIpv4OptionsCustomType) OptionNumber() PatternFlowIpv4OptionsCustomTypeOptionNumber {
	if obj.obj.OptionNumber == nil {
		obj.obj.OptionNumber = NewPatternFlowIpv4OptionsCustomTypeOptionNumber().msg()
	}
	if obj.optionNumberHolder == nil {
		obj.optionNumberHolder = &patternFlowIpv4OptionsCustomTypeOptionNumber{obj: obj.obj.OptionNumber}
	}
	return obj.optionNumberHolder
}

// description is TBD
// OptionNumber returns a PatternFlowIpv4OptionsCustomTypeOptionNumber
func (obj *flowIpv4OptionsCustomType) HasOptionNumber() bool {
	return obj.obj.OptionNumber != nil
}

// description is TBD
// SetOptionNumber sets the PatternFlowIpv4OptionsCustomTypeOptionNumber value in the FlowIpv4OptionsCustomType object
func (obj *flowIpv4OptionsCustomType) SetOptionNumber(value PatternFlowIpv4OptionsCustomTypeOptionNumber) FlowIpv4OptionsCustomType {

	obj.optionNumberHolder = nil
	obj.obj.OptionNumber = value.msg()

	return obj
}

func (obj *flowIpv4OptionsCustomType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CopiedFlag != nil {

		obj.CopiedFlag().validateObj(vObj, set_default)
	}

	if obj.obj.OptionClass != nil {

		obj.OptionClass().validateObj(vObj, set_default)
	}

	if obj.obj.OptionNumber != nil {

		obj.OptionNumber().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv4OptionsCustomType) setDefault() {

}
