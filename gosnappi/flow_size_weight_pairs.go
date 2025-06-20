package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowSizeWeightPairs *****
type flowSizeWeightPairs struct {
	validation
	obj          *otg.FlowSizeWeightPairs
	marshaller   marshalFlowSizeWeightPairs
	unMarshaller unMarshalFlowSizeWeightPairs
	customHolder FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
}

func NewFlowSizeWeightPairs() FlowSizeWeightPairs {
	obj := flowSizeWeightPairs{obj: &otg.FlowSizeWeightPairs{}}
	obj.setDefault()
	return &obj
}

func (obj *flowSizeWeightPairs) msg() *otg.FlowSizeWeightPairs {
	return obj.obj
}

func (obj *flowSizeWeightPairs) setMsg(msg *otg.FlowSizeWeightPairs) FlowSizeWeightPairs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowSizeWeightPairs struct {
	obj *flowSizeWeightPairs
}

type marshalFlowSizeWeightPairs interface {
	// ToProto marshals FlowSizeWeightPairs to protobuf object *otg.FlowSizeWeightPairs
	ToProto() (*otg.FlowSizeWeightPairs, error)
	// ToPbText marshals FlowSizeWeightPairs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowSizeWeightPairs to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowSizeWeightPairs to JSON text
	ToJson() (string, error)
}

type unMarshalflowSizeWeightPairs struct {
	obj *flowSizeWeightPairs
}

type unMarshalFlowSizeWeightPairs interface {
	// FromProto unmarshals FlowSizeWeightPairs from protobuf object *otg.FlowSizeWeightPairs
	FromProto(msg *otg.FlowSizeWeightPairs) (FlowSizeWeightPairs, error)
	// FromPbText unmarshals FlowSizeWeightPairs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowSizeWeightPairs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowSizeWeightPairs from JSON text
	FromJson(value string) error
}

func (obj *flowSizeWeightPairs) Marshal() marshalFlowSizeWeightPairs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowSizeWeightPairs{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowSizeWeightPairs) Unmarshal() unMarshalFlowSizeWeightPairs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowSizeWeightPairs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowSizeWeightPairs) ToProto() (*otg.FlowSizeWeightPairs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowSizeWeightPairs) FromProto(msg *otg.FlowSizeWeightPairs) (FlowSizeWeightPairs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowSizeWeightPairs) ToPbText() (string, error) {
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

func (m *unMarshalflowSizeWeightPairs) FromPbText(value string) error {
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

func (m *marshalflowSizeWeightPairs) ToYaml() (string, error) {
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

func (m *unMarshalflowSizeWeightPairs) FromYaml(value string) error {
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

func (m *marshalflowSizeWeightPairs) ToJson() (string, error) {
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

func (m *unMarshalflowSizeWeightPairs) FromJson(value string) error {
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

func (obj *flowSizeWeightPairs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowSizeWeightPairs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowSizeWeightPairs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowSizeWeightPairs) Clone() (FlowSizeWeightPairs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowSizeWeightPairs()
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

func (obj *flowSizeWeightPairs) setNil() {
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowSizeWeightPairs is frame size distribution, defined as <size, weight> pairs (including IMIX distribution).
// Frames are randomly generated such that the proportion of each frame size out of the total number of frames
// are matching with the weight value of the <size, weight> pair. However, as with any other probability
// distribution, the sample distribution is close to theoretical value only if the size of the sample is reasonably large.
// When the number of frames is very low the transmitted frames may not come close to the ratio described in the weight.
type FlowSizeWeightPairs interface {
	Validation
	// msg marshals FlowSizeWeightPairs to protobuf object *otg.FlowSizeWeightPairs
	// and doesn't set defaults
	msg() *otg.FlowSizeWeightPairs
	// setMsg unmarshals FlowSizeWeightPairs from protobuf object *otg.FlowSizeWeightPairs
	// and doesn't set defaults
	setMsg(*otg.FlowSizeWeightPairs) FlowSizeWeightPairs
	// provides marshal interface
	Marshal() marshalFlowSizeWeightPairs
	// provides unmarshal interface
	Unmarshal() unMarshalFlowSizeWeightPairs
	// validate validates FlowSizeWeightPairs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowSizeWeightPairs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowSizeWeightPairsChoiceEnum, set in FlowSizeWeightPairs
	Choice() FlowSizeWeightPairsChoiceEnum
	// setChoice assigns FlowSizeWeightPairsChoiceEnum provided by user to FlowSizeWeightPairs
	setChoice(value FlowSizeWeightPairsChoiceEnum) FlowSizeWeightPairs
	// HasChoice checks if Choice has been set in FlowSizeWeightPairs
	HasChoice() bool
	// Predefined returns FlowSizeWeightPairsPredefinedEnum, set in FlowSizeWeightPairs
	Predefined() FlowSizeWeightPairsPredefinedEnum
	// SetPredefined assigns FlowSizeWeightPairsPredefinedEnum provided by user to FlowSizeWeightPairs
	SetPredefined(value FlowSizeWeightPairsPredefinedEnum) FlowSizeWeightPairs
	// HasPredefined checks if Predefined has been set in FlowSizeWeightPairs
	HasPredefined() bool
	// Custom returns FlowSizeWeightPairsFlowSizeWeightPairsCustomIterIter, set in FlowSizeWeightPairs
	Custom() FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
	setNil()
}

type FlowSizeWeightPairsChoiceEnum string

// Enum of Choice on FlowSizeWeightPairs
var FlowSizeWeightPairsChoice = struct {
	PREDEFINED FlowSizeWeightPairsChoiceEnum
	CUSTOM     FlowSizeWeightPairsChoiceEnum
}{
	PREDEFINED: FlowSizeWeightPairsChoiceEnum("predefined"),
	CUSTOM:     FlowSizeWeightPairsChoiceEnum("custom"),
}

func (obj *flowSizeWeightPairs) Choice() FlowSizeWeightPairsChoiceEnum {
	return FlowSizeWeightPairsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowSizeWeightPairs) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowSizeWeightPairs) setChoice(value FlowSizeWeightPairsChoiceEnum) FlowSizeWeightPairs {
	intValue, ok := otg.FlowSizeWeightPairs_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSizeWeightPairsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSizeWeightPairs_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil
	obj.obj.Predefined = otg.FlowSizeWeightPairs_Predefined_unspecified.Enum()

	if value == FlowSizeWeightPairsChoice.CUSTOM {
		obj.obj.Custom = []*otg.FlowSizeWeightPairsCustom{}
	}

	return obj
}

type FlowSizeWeightPairsPredefinedEnum string

// Enum of Predefined on FlowSizeWeightPairs
var FlowSizeWeightPairsPredefined = struct {
	IMIX          FlowSizeWeightPairsPredefinedEnum
	IPSEC_IMIX    FlowSizeWeightPairsPredefinedEnum
	IPV6_IMIX     FlowSizeWeightPairsPredefinedEnum
	STANDARD_IMIX FlowSizeWeightPairsPredefinedEnum
	TCP_IMIX      FlowSizeWeightPairsPredefinedEnum
}{
	IMIX:          FlowSizeWeightPairsPredefinedEnum("imix"),
	IPSEC_IMIX:    FlowSizeWeightPairsPredefinedEnum("ipsec_imix"),
	IPV6_IMIX:     FlowSizeWeightPairsPredefinedEnum("ipv6_imix"),
	STANDARD_IMIX: FlowSizeWeightPairsPredefinedEnum("standard_imix"),
	TCP_IMIX:      FlowSizeWeightPairsPredefinedEnum("tcp_imix"),
}

func (obj *flowSizeWeightPairs) Predefined() FlowSizeWeightPairsPredefinedEnum {
	return FlowSizeWeightPairsPredefinedEnum(obj.obj.Predefined.Enum().String())
}

// Specify predefined frame size distribution <size, weight> pairs (including IMIX distribution).
// The available predefined distribution pairs are:
// - IMIX (64:7, 570:4, and 1518:1)
// - IPSec IMIX (90:58.67, 92:2, 594:23.66 and 1418:15.67)
// - IPv6 IMIX (60:58.67, 496:2, 594:23.66 and 1518:15.67)
// - Standard IMIX (58:58.67, 62:2, 594:23.66 and 1518:15.67)
// - TCP IMIX (90:58.67, 92:2, 594:23.66 and 1518:15.67)
// Predefined returns a string
func (obj *flowSizeWeightPairs) HasPredefined() bool {
	return obj.obj.Predefined != nil
}

func (obj *flowSizeWeightPairs) SetPredefined(value FlowSizeWeightPairsPredefinedEnum) FlowSizeWeightPairs {
	intValue, ok := otg.FlowSizeWeightPairs_Predefined_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowSizeWeightPairsPredefinedEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowSizeWeightPairs_Predefined_Enum(intValue)
	obj.obj.Predefined = &enumValue

	return obj
}

// description is TBD
// Custom returns a []FlowSizeWeightPairsCustom
func (obj *flowSizeWeightPairs) Custom() FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	if len(obj.obj.Custom) == 0 {
		obj.setChoice(FlowSizeWeightPairsChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = newFlowSizeWeightPairsFlowSizeWeightPairsCustomIter(&obj.obj.Custom).setMsg(obj)
	}
	return obj.customHolder
}

type flowSizeWeightPairsFlowSizeWeightPairsCustomIter struct {
	obj                            *flowSizeWeightPairs
	flowSizeWeightPairsCustomSlice []FlowSizeWeightPairsCustom
	fieldPtr                       *[]*otg.FlowSizeWeightPairsCustom
}

func newFlowSizeWeightPairsFlowSizeWeightPairsCustomIter(ptr *[]*otg.FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	return &flowSizeWeightPairsFlowSizeWeightPairsCustomIter{fieldPtr: ptr}
}

type FlowSizeWeightPairsFlowSizeWeightPairsCustomIter interface {
	setMsg(*flowSizeWeightPairs) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
	Items() []FlowSizeWeightPairsCustom
	Add() FlowSizeWeightPairsCustom
	Append(items ...FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
	Set(index int, newObj FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
	Clear() FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
	clearHolderSlice() FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
	appendHolderSlice(item FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter
}

func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) setMsg(msg *flowSizeWeightPairs) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowSizeWeightPairsCustom{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) Items() []FlowSizeWeightPairsCustom {
	return obj.flowSizeWeightPairsCustomSlice
}

func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) Add() FlowSizeWeightPairsCustom {
	newObj := &otg.FlowSizeWeightPairsCustom{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowSizeWeightPairsCustom{obj: newObj}
	newLibObj.setDefault()
	obj.flowSizeWeightPairsCustomSlice = append(obj.flowSizeWeightPairsCustomSlice, newLibObj)
	return newLibObj
}

func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) Append(items ...FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowSizeWeightPairsCustomSlice = append(obj.flowSizeWeightPairsCustomSlice, item)
	}
	return obj
}

func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) Set(index int, newObj FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowSizeWeightPairsCustomSlice[index] = newObj
	return obj
}
func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) Clear() FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowSizeWeightPairsCustom{}
		obj.flowSizeWeightPairsCustomSlice = []FlowSizeWeightPairsCustom{}
	}
	return obj
}
func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) clearHolderSlice() FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	if len(obj.flowSizeWeightPairsCustomSlice) > 0 {
		obj.flowSizeWeightPairsCustomSlice = []FlowSizeWeightPairsCustom{}
	}
	return obj
}
func (obj *flowSizeWeightPairsFlowSizeWeightPairsCustomIter) appendHolderSlice(item FlowSizeWeightPairsCustom) FlowSizeWeightPairsFlowSizeWeightPairsCustomIter {
	obj.flowSizeWeightPairsCustomSlice = append(obj.flowSizeWeightPairsCustomSlice, item)
	return obj
}

func (obj *flowSizeWeightPairs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Custom) != 0 {

		if set_default {
			obj.Custom().clearHolderSlice()
			for _, item := range obj.obj.Custom {
				obj.Custom().appendHolderSlice(&flowSizeWeightPairsCustom{obj: item})
			}
		}
		for _, item := range obj.Custom().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowSizeWeightPairs) setDefault() {
	var choices_set int = 0
	var choice FlowSizeWeightPairsChoiceEnum

	if obj.obj.Predefined != nil && obj.obj.Predefined.Number() != 0 {
		choices_set += 1
		choice = FlowSizeWeightPairsChoice.PREDEFINED
	}

	if len(obj.obj.Custom) > 0 {
		choices_set += 1
		choice = FlowSizeWeightPairsChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowSizeWeightPairsChoice.PREDEFINED)
			if obj.obj.Predefined.Number() == 0 {
				obj.SetPredefined(FlowSizeWeightPairsPredefined.IMIX)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowSizeWeightPairs")
			}
		} else {
			intVal := otg.FlowSizeWeightPairs_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowSizeWeightPairs_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
