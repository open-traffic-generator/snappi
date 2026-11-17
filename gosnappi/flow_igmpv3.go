package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIgmpv3 *****
type flowIgmpv3 struct {
	validation
	obj                    *otg.FlowIgmpv3
	marshaller             marshalFlowIgmpv3
	unMarshaller           unMarshalFlowIgmpv3
	membershipQueryHolder  FlowIgmpv3MembershipQuery
	membershipReportHolder FlowIgmpv3MembershipReport
}

func NewFlowIgmpv3() FlowIgmpv3 {
	obj := flowIgmpv3{obj: &otg.FlowIgmpv3{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIgmpv3) msg() *otg.FlowIgmpv3 {
	return obj.obj
}

func (obj *flowIgmpv3) setMsg(msg *otg.FlowIgmpv3) FlowIgmpv3 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIgmpv3 struct {
	obj *flowIgmpv3
}

type marshalFlowIgmpv3 interface {
	// ToProto marshals FlowIgmpv3 to protobuf object *otg.FlowIgmpv3
	ToProto() (*otg.FlowIgmpv3, error)
	// ToPbText marshals FlowIgmpv3 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIgmpv3 to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIgmpv3 to JSON text
	ToJson() (string, error)
}

type unMarshalflowIgmpv3 struct {
	obj *flowIgmpv3
}

type unMarshalFlowIgmpv3 interface {
	// FromProto unmarshals FlowIgmpv3 from protobuf object *otg.FlowIgmpv3
	FromProto(msg *otg.FlowIgmpv3) (FlowIgmpv3, error)
	// FromPbText unmarshals FlowIgmpv3 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIgmpv3 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIgmpv3 from JSON text
	FromJson(value string) error
}

func (obj *flowIgmpv3) Marshal() marshalFlowIgmpv3 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIgmpv3{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIgmpv3) Unmarshal() unMarshalFlowIgmpv3 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIgmpv3{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIgmpv3) ToProto() (*otg.FlowIgmpv3, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIgmpv3) FromProto(msg *otg.FlowIgmpv3) (FlowIgmpv3, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIgmpv3) ToPbText() (string, error) {
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

func (m *unMarshalflowIgmpv3) FromPbText(value string) error {
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

func (m *marshalflowIgmpv3) ToYaml() (string, error) {
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

func (m *unMarshalflowIgmpv3) FromYaml(value string) error {
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

func (m *marshalflowIgmpv3) ToJson() (string, error) {
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

func (m *unMarshalflowIgmpv3) FromJson(value string) error {
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

func (obj *flowIgmpv3) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIgmpv3) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIgmpv3) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIgmpv3) Clone() (FlowIgmpv3, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIgmpv3()
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

func (obj *flowIgmpv3) setNil() {
	obj.membershipQueryHolder = nil
	obj.membershipReportHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIgmpv3 is iGMPv3 packet header (RFC 9776, which obsoletes RFC 3376). The choice selects the Membership Query or the Version 3 Membership Report message layout.
type FlowIgmpv3 interface {
	Validation
	// msg marshals FlowIgmpv3 to protobuf object *otg.FlowIgmpv3
	// and doesn't set defaults
	msg() *otg.FlowIgmpv3
	// setMsg unmarshals FlowIgmpv3 from protobuf object *otg.FlowIgmpv3
	// and doesn't set defaults
	setMsg(*otg.FlowIgmpv3) FlowIgmpv3
	// provides marshal interface
	Marshal() marshalFlowIgmpv3
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIgmpv3
	// validate validates FlowIgmpv3
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIgmpv3, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIgmpv3ChoiceEnum, set in FlowIgmpv3
	Choice() FlowIgmpv3ChoiceEnum
	// setChoice assigns FlowIgmpv3ChoiceEnum provided by user to FlowIgmpv3
	setChoice(value FlowIgmpv3ChoiceEnum) FlowIgmpv3
	// HasChoice checks if Choice has been set in FlowIgmpv3
	HasChoice() bool
	// MembershipQuery returns FlowIgmpv3MembershipQuery, set in FlowIgmpv3.
	// FlowIgmpv3MembershipQuery is iGMPv3 Membership Query message (RFC 9776 Section 4.1).
	MembershipQuery() FlowIgmpv3MembershipQuery
	// SetMembershipQuery assigns FlowIgmpv3MembershipQuery provided by user to FlowIgmpv3.
	// FlowIgmpv3MembershipQuery is iGMPv3 Membership Query message (RFC 9776 Section 4.1).
	SetMembershipQuery(value FlowIgmpv3MembershipQuery) FlowIgmpv3
	// HasMembershipQuery checks if MembershipQuery has been set in FlowIgmpv3
	HasMembershipQuery() bool
	// MembershipReport returns FlowIgmpv3MembershipReport, set in FlowIgmpv3.
	// FlowIgmpv3MembershipReport is iGMPv3 Version 3 Membership Report message (RFC 9776 Section 4.2).
	MembershipReport() FlowIgmpv3MembershipReport
	// SetMembershipReport assigns FlowIgmpv3MembershipReport provided by user to FlowIgmpv3.
	// FlowIgmpv3MembershipReport is iGMPv3 Version 3 Membership Report message (RFC 9776 Section 4.2).
	SetMembershipReport(value FlowIgmpv3MembershipReport) FlowIgmpv3
	// HasMembershipReport checks if MembershipReport has been set in FlowIgmpv3
	HasMembershipReport() bool
	setNil()
}

type FlowIgmpv3ChoiceEnum string

// Enum of Choice on FlowIgmpv3
var FlowIgmpv3Choice = struct {
	MEMBERSHIP_QUERY  FlowIgmpv3ChoiceEnum
	MEMBERSHIP_REPORT FlowIgmpv3ChoiceEnum
}{
	MEMBERSHIP_QUERY:  FlowIgmpv3ChoiceEnum("membership_query"),
	MEMBERSHIP_REPORT: FlowIgmpv3ChoiceEnum("membership_report"),
}

func (obj *flowIgmpv3) Choice() FlowIgmpv3ChoiceEnum {
	return FlowIgmpv3ChoiceEnum(obj.obj.Choice.Enum().String())
}

// Selects the IGMPv3 message, default is membership_query.
// Choice returns a string
func (obj *flowIgmpv3) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIgmpv3) setChoice(value FlowIgmpv3ChoiceEnum) FlowIgmpv3 {
	intValue, ok := otg.FlowIgmpv3_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIgmpv3ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIgmpv3_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MembershipReport = nil
	obj.membershipReportHolder = nil
	obj.obj.MembershipQuery = nil
	obj.membershipQueryHolder = nil

	if value == FlowIgmpv3Choice.MEMBERSHIP_QUERY {
		obj.obj.MembershipQuery = NewFlowIgmpv3MembershipQuery().msg()
	}

	if value == FlowIgmpv3Choice.MEMBERSHIP_REPORT {
		obj.obj.MembershipReport = NewFlowIgmpv3MembershipReport().msg()
	}

	return obj
}

// description is TBD
// MembershipQuery returns a FlowIgmpv3MembershipQuery
func (obj *flowIgmpv3) MembershipQuery() FlowIgmpv3MembershipQuery {
	if obj.obj.MembershipQuery == nil {
		obj.setChoice(FlowIgmpv3Choice.MEMBERSHIP_QUERY)
	}
	if obj.membershipQueryHolder == nil {
		obj.membershipQueryHolder = &flowIgmpv3MembershipQuery{obj: obj.obj.MembershipQuery}
	}
	return obj.membershipQueryHolder
}

// description is TBD
// MembershipQuery returns a FlowIgmpv3MembershipQuery
func (obj *flowIgmpv3) HasMembershipQuery() bool {
	return obj.obj.MembershipQuery != nil
}

// description is TBD
// SetMembershipQuery sets the FlowIgmpv3MembershipQuery value in the FlowIgmpv3 object
func (obj *flowIgmpv3) SetMembershipQuery(value FlowIgmpv3MembershipQuery) FlowIgmpv3 {
	obj.setChoice(FlowIgmpv3Choice.MEMBERSHIP_QUERY)
	obj.membershipQueryHolder = nil
	obj.obj.MembershipQuery = value.msg()

	return obj
}

// description is TBD
// MembershipReport returns a FlowIgmpv3MembershipReport
func (obj *flowIgmpv3) MembershipReport() FlowIgmpv3MembershipReport {
	if obj.obj.MembershipReport == nil {
		obj.setChoice(FlowIgmpv3Choice.MEMBERSHIP_REPORT)
	}
	if obj.membershipReportHolder == nil {
		obj.membershipReportHolder = &flowIgmpv3MembershipReport{obj: obj.obj.MembershipReport}
	}
	return obj.membershipReportHolder
}

// description is TBD
// MembershipReport returns a FlowIgmpv3MembershipReport
func (obj *flowIgmpv3) HasMembershipReport() bool {
	return obj.obj.MembershipReport != nil
}

// description is TBD
// SetMembershipReport sets the FlowIgmpv3MembershipReport value in the FlowIgmpv3 object
func (obj *flowIgmpv3) SetMembershipReport(value FlowIgmpv3MembershipReport) FlowIgmpv3 {
	obj.setChoice(FlowIgmpv3Choice.MEMBERSHIP_REPORT)
	obj.membershipReportHolder = nil
	obj.obj.MembershipReport = value.msg()

	return obj
}

func (obj *flowIgmpv3) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MembershipQuery != nil {

		obj.MembershipQuery().validateObj(vObj, set_default)
	}

	if obj.obj.MembershipReport != nil {

		obj.MembershipReport().validateObj(vObj, set_default)
	}

}

func (obj *flowIgmpv3) setDefault() {
	var choices_set int = 0
	var choice FlowIgmpv3ChoiceEnum

	if obj.obj.MembershipQuery != nil {
		choices_set += 1
		choice = FlowIgmpv3Choice.MEMBERSHIP_QUERY
	}

	if obj.obj.MembershipReport != nil {
		choices_set += 1
		choice = FlowIgmpv3Choice.MEMBERSHIP_REPORT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIgmpv3Choice.MEMBERSHIP_QUERY)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIgmpv3")
			}
		} else {
			intVal := otg.FlowIgmpv3_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIgmpv3_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
