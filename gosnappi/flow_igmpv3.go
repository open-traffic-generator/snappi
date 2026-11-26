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

// FlowIgmpv3 is packet Header for IGMPv3 (RFC 9776, which obsoletes RFC 3376), carried directly in IPv4. The choice selects the Membership Query or the IGMPv3 Membership Report. A conforming message is sent with IP protocol number two, TTL one, IP precedence Internetwork Control (ToS 0xc0) and the IP Router Alert option (RFC 2113). The model does not enforce these requirements, so any value the fields can encode may be configured for negative testing.
// A Query of eight octets is IGMPv1 or IGMPv2 and a Query of twelve or more octets is IGMPv3 (RFC 9776 Section 7.1), so membership_query always builds an IGMPv3 Query. The eight octet messages use the igmpv1 and igmpv2 headers.
// The IPv4 and Ethernet destinations are set in their own headers. Use these values (RFC 9776 Sections 4.1.12 and 4.2.15):
// General Query:
// IPv4 dst = 224.0.0.1
// Group-Specific and Group-and-Source Specific Query:
// IPv4 dst = group_address
// IGMPv3 Membership Report:
// IPv4 dst = 224.0.0.22, IPv4 src = a unicast address or 0.0.0.0
//
// Ethernet dst = 01:00:5e + low 23 bits of the IPv4 dst (RFC 1112 Section 6.4), for example 224.0.0.22 gives 01:00:5e:00:00:16.
// The flow size overrides the packet length, so set it large enough for the whole message. A Query is twelve octets plus four per source. A Report is eight octets plus, per group record, eight octets plus four per source plus four per aux_data_length word. Ethernet, IPv4 with Router Alert and the FCS add 42 octets. The default flow size of 64 fits a Query with up to two sources or a Report with one record of up to one source.
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
	// FlowIgmpv3MembershipReport is iGMPv3 Membership Report message (RFC 9776 Section 4.2).
	MembershipReport() FlowIgmpv3MembershipReport
	// SetMembershipReport assigns FlowIgmpv3MembershipReport provided by user to FlowIgmpv3.
	// FlowIgmpv3MembershipReport is iGMPv3 Membership Report message (RFC 9776 Section 4.2).
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
