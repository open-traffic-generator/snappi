package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmSelectOpCode *****
type flowCfmSelectOpCode struct {
	validation
	obj          *otg.FlowCfmSelectOpCode
	marshaller   marshalFlowCfmSelectOpCode
	unMarshaller unMarshalFlowCfmSelectOpCode
	ccmHolder    FlowCfmCcm
	lbrHolder    FlowCfmLbr
}

func NewFlowCfmSelectOpCode() FlowCfmSelectOpCode {
	obj := flowCfmSelectOpCode{obj: &otg.FlowCfmSelectOpCode{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmSelectOpCode) msg() *otg.FlowCfmSelectOpCode {
	return obj.obj
}

func (obj *flowCfmSelectOpCode) setMsg(msg *otg.FlowCfmSelectOpCode) FlowCfmSelectOpCode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmSelectOpCode struct {
	obj *flowCfmSelectOpCode
}

type marshalFlowCfmSelectOpCode interface {
	// ToProto marshals FlowCfmSelectOpCode to protobuf object *otg.FlowCfmSelectOpCode
	ToProto() (*otg.FlowCfmSelectOpCode, error)
	// ToPbText marshals FlowCfmSelectOpCode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmSelectOpCode to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmSelectOpCode to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmSelectOpCode struct {
	obj *flowCfmSelectOpCode
}

type unMarshalFlowCfmSelectOpCode interface {
	// FromProto unmarshals FlowCfmSelectOpCode from protobuf object *otg.FlowCfmSelectOpCode
	FromProto(msg *otg.FlowCfmSelectOpCode) (FlowCfmSelectOpCode, error)
	// FromPbText unmarshals FlowCfmSelectOpCode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmSelectOpCode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmSelectOpCode from JSON text
	FromJson(value string) error
}

func (obj *flowCfmSelectOpCode) Marshal() marshalFlowCfmSelectOpCode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmSelectOpCode{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmSelectOpCode) Unmarshal() unMarshalFlowCfmSelectOpCode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmSelectOpCode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmSelectOpCode) ToProto() (*otg.FlowCfmSelectOpCode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmSelectOpCode) FromProto(msg *otg.FlowCfmSelectOpCode) (FlowCfmSelectOpCode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmSelectOpCode) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmSelectOpCode) FromPbText(value string) error {
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

func (m *marshalflowCfmSelectOpCode) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmSelectOpCode) FromYaml(value string) error {
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

func (m *marshalflowCfmSelectOpCode) ToJson() (string, error) {
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

func (m *unMarshalflowCfmSelectOpCode) FromJson(value string) error {
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

func (obj *flowCfmSelectOpCode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmSelectOpCode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmSelectOpCode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmSelectOpCode) Clone() (FlowCfmSelectOpCode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmSelectOpCode()
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

func (obj *flowCfmSelectOpCode) setNil() {
	obj.ccmHolder = nil
	obj.lbrHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmSelectOpCode is select message header. Currently only Continuity Check Message (CCM) and Loobback Reply (LBR) are supported.
type FlowCfmSelectOpCode interface {
	Validation
	// msg marshals FlowCfmSelectOpCode to protobuf object *otg.FlowCfmSelectOpCode
	// and doesn't set defaults
	msg() *otg.FlowCfmSelectOpCode
	// setMsg unmarshals FlowCfmSelectOpCode from protobuf object *otg.FlowCfmSelectOpCode
	// and doesn't set defaults
	setMsg(*otg.FlowCfmSelectOpCode) FlowCfmSelectOpCode
	// provides marshal interface
	Marshal() marshalFlowCfmSelectOpCode
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmSelectOpCode
	// validate validates FlowCfmSelectOpCode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmSelectOpCode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmSelectOpCodeChoiceEnum, set in FlowCfmSelectOpCode
	Choice() FlowCfmSelectOpCodeChoiceEnum
	// setChoice assigns FlowCfmSelectOpCodeChoiceEnum provided by user to FlowCfmSelectOpCode
	setChoice(value FlowCfmSelectOpCodeChoiceEnum) FlowCfmSelectOpCode
	// HasChoice checks if Choice has been set in FlowCfmSelectOpCode
	HasChoice() bool
	// Ccm returns FlowCfmCcm, set in FlowCfmSelectOpCode.
	// FlowCfmCcm is cCM is used to monitor connectivity and configuration errors. Each CCM message has a unique sequence number (Session ID) and unique flow-identifier.
	Ccm() FlowCfmCcm
	// SetCcm assigns FlowCfmCcm provided by user to FlowCfmSelectOpCode.
	// FlowCfmCcm is cCM is used to monitor connectivity and configuration errors. Each CCM message has a unique sequence number (Session ID) and unique flow-identifier.
	SetCcm(value FlowCfmCcm) FlowCfmSelectOpCode
	// HasCcm checks if Ccm has been set in FlowCfmSelectOpCode
	HasCcm() bool
	// Lbr returns FlowCfmLbr, set in FlowCfmSelectOpCode.
	// FlowCfmLbr is loopback Reply (LBR) messages are unicast frames sent by a Maintenance Association End Point (MEP) or Intermediate Point (MIP) to a source MEP in response to a Loopback Message (LBM).
	Lbr() FlowCfmLbr
	// SetLbr assigns FlowCfmLbr provided by user to FlowCfmSelectOpCode.
	// FlowCfmLbr is loopback Reply (LBR) messages are unicast frames sent by a Maintenance Association End Point (MEP) or Intermediate Point (MIP) to a source MEP in response to a Loopback Message (LBM).
	SetLbr(value FlowCfmLbr) FlowCfmSelectOpCode
	// HasLbr checks if Lbr has been set in FlowCfmSelectOpCode
	HasLbr() bool
	setNil()
}

type FlowCfmSelectOpCodeChoiceEnum string

// Enum of Choice on FlowCfmSelectOpCode
var FlowCfmSelectOpCodeChoice = struct {
	CCM FlowCfmSelectOpCodeChoiceEnum
	LBR FlowCfmSelectOpCodeChoiceEnum
}{
	CCM: FlowCfmSelectOpCodeChoiceEnum("ccm"),
	LBR: FlowCfmSelectOpCodeChoiceEnum("lbr"),
}

func (obj *flowCfmSelectOpCode) Choice() FlowCfmSelectOpCodeChoiceEnum {
	return FlowCfmSelectOpCodeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowCfmSelectOpCode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmSelectOpCode) setChoice(value FlowCfmSelectOpCodeChoiceEnum) FlowCfmSelectOpCode {
	intValue, ok := otg.FlowCfmSelectOpCode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmSelectOpCodeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmSelectOpCode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Lbr = nil
	obj.lbrHolder = nil
	obj.obj.Ccm = nil
	obj.ccmHolder = nil

	if value == FlowCfmSelectOpCodeChoice.CCM {
		obj.obj.Ccm = NewFlowCfmCcm().msg()
	}

	if value == FlowCfmSelectOpCodeChoice.LBR {
		obj.obj.Lbr = NewFlowCfmLbr().msg()
	}

	return obj
}

// description is TBD
// Ccm returns a FlowCfmCcm
func (obj *flowCfmSelectOpCode) Ccm() FlowCfmCcm {
	if obj.obj.Ccm == nil {
		obj.setChoice(FlowCfmSelectOpCodeChoice.CCM)
	}
	if obj.ccmHolder == nil {
		obj.ccmHolder = &flowCfmCcm{obj: obj.obj.Ccm}
	}
	return obj.ccmHolder
}

// description is TBD
// Ccm returns a FlowCfmCcm
func (obj *flowCfmSelectOpCode) HasCcm() bool {
	return obj.obj.Ccm != nil
}

// description is TBD
// SetCcm sets the FlowCfmCcm value in the FlowCfmSelectOpCode object
func (obj *flowCfmSelectOpCode) SetCcm(value FlowCfmCcm) FlowCfmSelectOpCode {
	obj.setChoice(FlowCfmSelectOpCodeChoice.CCM)
	obj.ccmHolder = nil
	obj.obj.Ccm = value.msg()

	return obj
}

// description is TBD
// Lbr returns a FlowCfmLbr
func (obj *flowCfmSelectOpCode) Lbr() FlowCfmLbr {
	if obj.obj.Lbr == nil {
		obj.setChoice(FlowCfmSelectOpCodeChoice.LBR)
	}
	if obj.lbrHolder == nil {
		obj.lbrHolder = &flowCfmLbr{obj: obj.obj.Lbr}
	}
	return obj.lbrHolder
}

// description is TBD
// Lbr returns a FlowCfmLbr
func (obj *flowCfmSelectOpCode) HasLbr() bool {
	return obj.obj.Lbr != nil
}

// description is TBD
// SetLbr sets the FlowCfmLbr value in the FlowCfmSelectOpCode object
func (obj *flowCfmSelectOpCode) SetLbr(value FlowCfmLbr) FlowCfmSelectOpCode {
	obj.setChoice(FlowCfmSelectOpCodeChoice.LBR)
	obj.lbrHolder = nil
	obj.obj.Lbr = value.msg()

	return obj
}

func (obj *flowCfmSelectOpCode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ccm != nil {

		obj.Ccm().validateObj(vObj, set_default)
	}

	if obj.obj.Lbr != nil {

		obj.Lbr().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmSelectOpCode) setDefault() {
	var choices_set int = 0
	var choice FlowCfmSelectOpCodeChoiceEnum

	if obj.obj.Ccm != nil {
		choices_set += 1
		choice = FlowCfmSelectOpCodeChoice.CCM
	}

	if obj.obj.Lbr != nil {
		choices_set += 1
		choice = FlowCfmSelectOpCodeChoice.LBR
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmSelectOpCodeChoice.CCM)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmSelectOpCode")
			}
		} else {
			intVal := otg.FlowCfmSelectOpCode_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmSelectOpCode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
