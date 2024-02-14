package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsLabelRequestCType *****
type flowRSVPPathObjectsLabelRequestCType struct {
	validation
	obj                     *otg.FlowRSVPPathObjectsLabelRequestCType
	marshaller              marshalFlowRSVPPathObjectsLabelRequestCType
	unMarshaller            unMarshalFlowRSVPPathObjectsLabelRequestCType
	withoutLabelRangeHolder FlowRSVPPathLabelRequestWithoutLabelRange
}

func NewFlowRSVPPathObjectsLabelRequestCType() FlowRSVPPathObjectsLabelRequestCType {
	obj := flowRSVPPathObjectsLabelRequestCType{obj: &otg.FlowRSVPPathObjectsLabelRequestCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsLabelRequestCType) msg() *otg.FlowRSVPPathObjectsLabelRequestCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsLabelRequestCType) setMsg(msg *otg.FlowRSVPPathObjectsLabelRequestCType) FlowRSVPPathObjectsLabelRequestCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsLabelRequestCType struct {
	obj *flowRSVPPathObjectsLabelRequestCType
}

type marshalFlowRSVPPathObjectsLabelRequestCType interface {
	// ToProto marshals FlowRSVPPathObjectsLabelRequestCType to protobuf object *otg.FlowRSVPPathObjectsLabelRequestCType
	ToProto() (*otg.FlowRSVPPathObjectsLabelRequestCType, error)
	// ToPbText marshals FlowRSVPPathObjectsLabelRequestCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsLabelRequestCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsLabelRequestCType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsLabelRequestCType struct {
	obj *flowRSVPPathObjectsLabelRequestCType
}

type unMarshalFlowRSVPPathObjectsLabelRequestCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsLabelRequestCType from protobuf object *otg.FlowRSVPPathObjectsLabelRequestCType
	FromProto(msg *otg.FlowRSVPPathObjectsLabelRequestCType) (FlowRSVPPathObjectsLabelRequestCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsLabelRequestCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsLabelRequestCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsLabelRequestCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsLabelRequestCType) Marshal() marshalFlowRSVPPathObjectsLabelRequestCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsLabelRequestCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsLabelRequestCType) Unmarshal() unMarshalFlowRSVPPathObjectsLabelRequestCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsLabelRequestCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsLabelRequestCType) ToProto() (*otg.FlowRSVPPathObjectsLabelRequestCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsLabelRequestCType) FromProto(msg *otg.FlowRSVPPathObjectsLabelRequestCType) (FlowRSVPPathObjectsLabelRequestCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsLabelRequestCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsLabelRequestCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsLabelRequestCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsLabelRequestCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsLabelRequestCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsLabelRequestCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsLabelRequestCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsLabelRequestCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsLabelRequestCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsLabelRequestCType) Clone() (FlowRSVPPathObjectsLabelRequestCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsLabelRequestCType()
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

func (obj *flowRSVPPathObjectsLabelRequestCType) setNil() {
	obj.withoutLabelRangeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsLabelRequestCType is object for LABEL_REQUEST class. Currently supported c-type is Without Label Range (1).
type FlowRSVPPathObjectsLabelRequestCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsLabelRequestCType to protobuf object *otg.FlowRSVPPathObjectsLabelRequestCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsLabelRequestCType
	// setMsg unmarshals FlowRSVPPathObjectsLabelRequestCType from protobuf object *otg.FlowRSVPPathObjectsLabelRequestCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsLabelRequestCType) FlowRSVPPathObjectsLabelRequestCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsLabelRequestCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsLabelRequestCType
	// validate validates FlowRSVPPathObjectsLabelRequestCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsLabelRequestCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum, set in FlowRSVPPathObjectsLabelRequestCType
	Choice() FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum provided by user to FlowRSVPPathObjectsLabelRequestCType
	setChoice(value FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum) FlowRSVPPathObjectsLabelRequestCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsLabelRequestCType
	HasChoice() bool
	// WithoutLabelRange returns FlowRSVPPathLabelRequestWithoutLabelRange, set in FlowRSVPPathObjectsLabelRequestCType.
	// FlowRSVPPathLabelRequestWithoutLabelRange is class = LABEL_REQUEST, Without Label Range C-Type = 1
	WithoutLabelRange() FlowRSVPPathLabelRequestWithoutLabelRange
	// SetWithoutLabelRange assigns FlowRSVPPathLabelRequestWithoutLabelRange provided by user to FlowRSVPPathObjectsLabelRequestCType.
	// FlowRSVPPathLabelRequestWithoutLabelRange is class = LABEL_REQUEST, Without Label Range C-Type = 1
	SetWithoutLabelRange(value FlowRSVPPathLabelRequestWithoutLabelRange) FlowRSVPPathObjectsLabelRequestCType
	// HasWithoutLabelRange checks if WithoutLabelRange has been set in FlowRSVPPathObjectsLabelRequestCType
	HasWithoutLabelRange() bool
	setNil()
}

type FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsLabelRequestCType
var FlowRSVPPathObjectsLabelRequestCTypeChoice = struct {
	WITHOUT_LABEL_RANGE FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum
}{
	WITHOUT_LABEL_RANGE: FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum("without_label_range"),
}

func (obj *flowRSVPPathObjectsLabelRequestCType) Choice() FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum {
	return FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsLabelRequestCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsLabelRequestCType) setChoice(value FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum) FlowRSVPPathObjectsLabelRequestCType {
	intValue, ok := otg.FlowRSVPPathObjectsLabelRequestCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsLabelRequestCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsLabelRequestCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.WithoutLabelRange = nil
	obj.withoutLabelRangeHolder = nil

	if value == FlowRSVPPathObjectsLabelRequestCTypeChoice.WITHOUT_LABEL_RANGE {
		obj.obj.WithoutLabelRange = NewFlowRSVPPathLabelRequestWithoutLabelRange().msg()
	}

	return obj
}

// description is TBD
// WithoutLabelRange returns a FlowRSVPPathLabelRequestWithoutLabelRange
func (obj *flowRSVPPathObjectsLabelRequestCType) WithoutLabelRange() FlowRSVPPathLabelRequestWithoutLabelRange {
	if obj.obj.WithoutLabelRange == nil {
		obj.setChoice(FlowRSVPPathObjectsLabelRequestCTypeChoice.WITHOUT_LABEL_RANGE)
	}
	if obj.withoutLabelRangeHolder == nil {
		obj.withoutLabelRangeHolder = &flowRSVPPathLabelRequestWithoutLabelRange{obj: obj.obj.WithoutLabelRange}
	}
	return obj.withoutLabelRangeHolder
}

// description is TBD
// WithoutLabelRange returns a FlowRSVPPathLabelRequestWithoutLabelRange
func (obj *flowRSVPPathObjectsLabelRequestCType) HasWithoutLabelRange() bool {
	return obj.obj.WithoutLabelRange != nil
}

// description is TBD
// SetWithoutLabelRange sets the FlowRSVPPathLabelRequestWithoutLabelRange value in the FlowRSVPPathObjectsLabelRequestCType object
func (obj *flowRSVPPathObjectsLabelRequestCType) SetWithoutLabelRange(value FlowRSVPPathLabelRequestWithoutLabelRange) FlowRSVPPathObjectsLabelRequestCType {
	obj.setChoice(FlowRSVPPathObjectsLabelRequestCTypeChoice.WITHOUT_LABEL_RANGE)
	obj.withoutLabelRangeHolder = nil
	obj.obj.WithoutLabelRange = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsLabelRequestCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.WithoutLabelRange != nil {

		obj.WithoutLabelRange().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsLabelRequestCType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPPathObjectsLabelRequestCTypeChoice.WITHOUT_LABEL_RANGE)

	}

}
