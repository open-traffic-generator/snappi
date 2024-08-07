package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathTimeValuesType1 *****
type flowRSVPPathTimeValuesType1 struct {
	validation
	obj                  *otg.FlowRSVPPathTimeValuesType1
	marshaller           marshalFlowRSVPPathTimeValuesType1
	unMarshaller         unMarshalFlowRSVPPathTimeValuesType1
	refreshPeriodRHolder PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
}

func NewFlowRSVPPathTimeValuesType1() FlowRSVPPathTimeValuesType1 {
	obj := flowRSVPPathTimeValuesType1{obj: &otg.FlowRSVPPathTimeValuesType1{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathTimeValuesType1) msg() *otg.FlowRSVPPathTimeValuesType1 {
	return obj.obj
}

func (obj *flowRSVPPathTimeValuesType1) setMsg(msg *otg.FlowRSVPPathTimeValuesType1) FlowRSVPPathTimeValuesType1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathTimeValuesType1 struct {
	obj *flowRSVPPathTimeValuesType1
}

type marshalFlowRSVPPathTimeValuesType1 interface {
	// ToProto marshals FlowRSVPPathTimeValuesType1 to protobuf object *otg.FlowRSVPPathTimeValuesType1
	ToProto() (*otg.FlowRSVPPathTimeValuesType1, error)
	// ToPbText marshals FlowRSVPPathTimeValuesType1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathTimeValuesType1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathTimeValuesType1 to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathTimeValuesType1 struct {
	obj *flowRSVPPathTimeValuesType1
}

type unMarshalFlowRSVPPathTimeValuesType1 interface {
	// FromProto unmarshals FlowRSVPPathTimeValuesType1 from protobuf object *otg.FlowRSVPPathTimeValuesType1
	FromProto(msg *otg.FlowRSVPPathTimeValuesType1) (FlowRSVPPathTimeValuesType1, error)
	// FromPbText unmarshals FlowRSVPPathTimeValuesType1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathTimeValuesType1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathTimeValuesType1 from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathTimeValuesType1) Marshal() marshalFlowRSVPPathTimeValuesType1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathTimeValuesType1{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathTimeValuesType1) Unmarshal() unMarshalFlowRSVPPathTimeValuesType1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathTimeValuesType1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathTimeValuesType1) ToProto() (*otg.FlowRSVPPathTimeValuesType1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathTimeValuesType1) FromProto(msg *otg.FlowRSVPPathTimeValuesType1) (FlowRSVPPathTimeValuesType1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathTimeValuesType1) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathTimeValuesType1) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathTimeValuesType1) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathTimeValuesType1) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathTimeValuesType1) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathTimeValuesType1) FromJson(value string) error {
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

func (obj *flowRSVPPathTimeValuesType1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathTimeValuesType1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathTimeValuesType1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathTimeValuesType1) Clone() (FlowRSVPPathTimeValuesType1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathTimeValuesType1()
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

func (obj *flowRSVPPathTimeValuesType1) setNil() {
	obj.refreshPeriodRHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathTimeValuesType1 is tIME_VALUES Object: Class = 5, C-Type = 1
type FlowRSVPPathTimeValuesType1 interface {
	Validation
	// msg marshals FlowRSVPPathTimeValuesType1 to protobuf object *otg.FlowRSVPPathTimeValuesType1
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathTimeValuesType1
	// setMsg unmarshals FlowRSVPPathTimeValuesType1 from protobuf object *otg.FlowRSVPPathTimeValuesType1
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathTimeValuesType1) FlowRSVPPathTimeValuesType1
	// provides marshal interface
	Marshal() marshalFlowRSVPPathTimeValuesType1
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathTimeValuesType1
	// validate validates FlowRSVPPathTimeValuesType1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathTimeValuesType1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RefreshPeriodR returns PatternFlowRSVPPathTimeValuesType1RefreshPeriodR, set in FlowRSVPPathTimeValuesType1.
	// PatternFlowRSVPPathTimeValuesType1RefreshPeriodR is the refresh timeout period R used to generate this message;in milliseconds.
	RefreshPeriodR() PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
	// SetRefreshPeriodR assigns PatternFlowRSVPPathTimeValuesType1RefreshPeriodR provided by user to FlowRSVPPathTimeValuesType1.
	// PatternFlowRSVPPathTimeValuesType1RefreshPeriodR is the refresh timeout period R used to generate this message;in milliseconds.
	SetRefreshPeriodR(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodR) FlowRSVPPathTimeValuesType1
	// HasRefreshPeriodR checks if RefreshPeriodR has been set in FlowRSVPPathTimeValuesType1
	HasRefreshPeriodR() bool
	setNil()
}

// description is TBD
// RefreshPeriodR returns a PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
func (obj *flowRSVPPathTimeValuesType1) RefreshPeriodR() PatternFlowRSVPPathTimeValuesType1RefreshPeriodR {
	if obj.obj.RefreshPeriodR == nil {
		obj.obj.RefreshPeriodR = NewPatternFlowRSVPPathTimeValuesType1RefreshPeriodR().msg()
	}
	if obj.refreshPeriodRHolder == nil {
		obj.refreshPeriodRHolder = &patternFlowRSVPPathTimeValuesType1RefreshPeriodR{obj: obj.obj.RefreshPeriodR}
	}
	return obj.refreshPeriodRHolder
}

// description is TBD
// RefreshPeriodR returns a PatternFlowRSVPPathTimeValuesType1RefreshPeriodR
func (obj *flowRSVPPathTimeValuesType1) HasRefreshPeriodR() bool {
	return obj.obj.RefreshPeriodR != nil
}

// description is TBD
// SetRefreshPeriodR sets the PatternFlowRSVPPathTimeValuesType1RefreshPeriodR value in the FlowRSVPPathTimeValuesType1 object
func (obj *flowRSVPPathTimeValuesType1) SetRefreshPeriodR(value PatternFlowRSVPPathTimeValuesType1RefreshPeriodR) FlowRSVPPathTimeValuesType1 {

	obj.refreshPeriodRHolder = nil
	obj.obj.RefreshPeriodR = value.msg()

	return obj
}

func (obj *flowRSVPPathTimeValuesType1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RefreshPeriodR != nil {

		obj.RefreshPeriodR().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathTimeValuesType1) setDefault() {

}
