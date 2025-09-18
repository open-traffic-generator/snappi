package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6ExtHeader *****
type flowIpv6ExtHeader struct {
	validation
	obj              *otg.FlowIpv6ExtHeader
	marshaller       marshalFlowIpv6ExtHeader
	unMarshaller     unMarshalFlowIpv6ExtHeader
	nextHeaderHolder PatternFlowIpv6ExtHeaderNextHeader
	hdrExtLenHolder  PatternFlowIpv6ExtHeaderHdrExtLen
	routingHolder    FlowIpv6Routing
}

func NewFlowIpv6ExtHeader() FlowIpv6ExtHeader {
	obj := flowIpv6ExtHeader{obj: &otg.FlowIpv6ExtHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6ExtHeader) msg() *otg.FlowIpv6ExtHeader {
	return obj.obj
}

func (obj *flowIpv6ExtHeader) setMsg(msg *otg.FlowIpv6ExtHeader) FlowIpv6ExtHeader {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6ExtHeader struct {
	obj *flowIpv6ExtHeader
}

type marshalFlowIpv6ExtHeader interface {
	// ToProto marshals FlowIpv6ExtHeader to protobuf object *otg.FlowIpv6ExtHeader
	ToProto() (*otg.FlowIpv6ExtHeader, error)
	// ToPbText marshals FlowIpv6ExtHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6ExtHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6ExtHeader to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6ExtHeader struct {
	obj *flowIpv6ExtHeader
}

type unMarshalFlowIpv6ExtHeader interface {
	// FromProto unmarshals FlowIpv6ExtHeader from protobuf object *otg.FlowIpv6ExtHeader
	FromProto(msg *otg.FlowIpv6ExtHeader) (FlowIpv6ExtHeader, error)
	// FromPbText unmarshals FlowIpv6ExtHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6ExtHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6ExtHeader from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6ExtHeader) Marshal() marshalFlowIpv6ExtHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6ExtHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6ExtHeader) Unmarshal() unMarshalFlowIpv6ExtHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6ExtHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6ExtHeader) ToProto() (*otg.FlowIpv6ExtHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6ExtHeader) FromProto(msg *otg.FlowIpv6ExtHeader) (FlowIpv6ExtHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6ExtHeader) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6ExtHeader) FromPbText(value string) error {
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

func (m *marshalflowIpv6ExtHeader) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6ExtHeader) FromYaml(value string) error {
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

func (m *marshalflowIpv6ExtHeader) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6ExtHeader) FromJson(value string) error {
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

func (obj *flowIpv6ExtHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6ExtHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6ExtHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6ExtHeader) Clone() (FlowIpv6ExtHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6ExtHeader()
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

func (obj *flowIpv6ExtHeader) setNil() {
	obj.nextHeaderHolder = nil
	obj.hdrExtLenHolder = nil
	obj.routingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6ExtHeader is ipv6 extension packet headers.
type FlowIpv6ExtHeader interface {
	Validation
	// msg marshals FlowIpv6ExtHeader to protobuf object *otg.FlowIpv6ExtHeader
	// and doesn't set defaults
	msg() *otg.FlowIpv6ExtHeader
	// setMsg unmarshals FlowIpv6ExtHeader from protobuf object *otg.FlowIpv6ExtHeader
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6ExtHeader) FlowIpv6ExtHeader
	// provides marshal interface
	Marshal() marshalFlowIpv6ExtHeader
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6ExtHeader
	// validate validates FlowIpv6ExtHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6ExtHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIpv6ExtHeaderChoiceEnum, set in FlowIpv6ExtHeader
	Choice() FlowIpv6ExtHeaderChoiceEnum
	// setChoice assigns FlowIpv6ExtHeaderChoiceEnum provided by user to FlowIpv6ExtHeader
	setChoice(value FlowIpv6ExtHeaderChoiceEnum) FlowIpv6ExtHeader
	// HasChoice checks if Choice has been set in FlowIpv6ExtHeader
	HasChoice() bool
	// NextHeader returns PatternFlowIpv6ExtHeaderNextHeader, set in FlowIpv6ExtHeader.
	// PatternFlowIpv6ExtHeaderNextHeader is 8-bit selector that identifies the type of header immediately following the SRH. It uses the same values as the IPv4 Protocol field. For TCP and UDP the values are 6 and 17 respectively. For other headers refer RFC 8200.
	NextHeader() PatternFlowIpv6ExtHeaderNextHeader
	// SetNextHeader assigns PatternFlowIpv6ExtHeaderNextHeader provided by user to FlowIpv6ExtHeader.
	// PatternFlowIpv6ExtHeaderNextHeader is 8-bit selector that identifies the type of header immediately following the SRH. It uses the same values as the IPv4 Protocol field. For TCP and UDP the values are 6 and 17 respectively. For other headers refer RFC 8200.
	SetNextHeader(value PatternFlowIpv6ExtHeaderNextHeader) FlowIpv6ExtHeader
	// HasNextHeader checks if NextHeader has been set in FlowIpv6ExtHeader
	HasNextHeader() bool
	// HdrExtLen returns PatternFlowIpv6ExtHeaderHdrExtLen, set in FlowIpv6ExtHeader.
	// PatternFlowIpv6ExtHeaderHdrExtLen is 8-bit unsigned integer specifying the length of the SRH in 8-octet units, not including the first 8 octets of the SRH itself. When AUTO is assigned it is set to the correct value.
	HdrExtLen() PatternFlowIpv6ExtHeaderHdrExtLen
	// SetHdrExtLen assigns PatternFlowIpv6ExtHeaderHdrExtLen provided by user to FlowIpv6ExtHeader.
	// PatternFlowIpv6ExtHeaderHdrExtLen is 8-bit unsigned integer specifying the length of the SRH in 8-octet units, not including the first 8 octets of the SRH itself. When AUTO is assigned it is set to the correct value.
	SetHdrExtLen(value PatternFlowIpv6ExtHeaderHdrExtLen) FlowIpv6ExtHeader
	// HasHdrExtLen checks if HdrExtLen has been set in FlowIpv6ExtHeader
	HasHdrExtLen() bool
	// Routing returns FlowIpv6Routing, set in FlowIpv6ExtHeader.
	// FlowIpv6Routing is ipv6 routing packet headers.
	Routing() FlowIpv6Routing
	// SetRouting assigns FlowIpv6Routing provided by user to FlowIpv6ExtHeader.
	// FlowIpv6Routing is ipv6 routing packet headers.
	SetRouting(value FlowIpv6Routing) FlowIpv6ExtHeader
	// HasRouting checks if Routing has been set in FlowIpv6ExtHeader
	HasRouting() bool
	setNil()
}

type FlowIpv6ExtHeaderChoiceEnum string

// Enum of Choice on FlowIpv6ExtHeader
var FlowIpv6ExtHeaderChoice = struct {
	ROUTING FlowIpv6ExtHeaderChoiceEnum
}{
	ROUTING: FlowIpv6ExtHeaderChoiceEnum("routing"),
}

func (obj *flowIpv6ExtHeader) Choice() FlowIpv6ExtHeaderChoiceEnum {
	return FlowIpv6ExtHeaderChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIpv6ExtHeader) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIpv6ExtHeader) setChoice(value FlowIpv6ExtHeaderChoiceEnum) FlowIpv6ExtHeader {
	intValue, ok := otg.FlowIpv6ExtHeader_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIpv6ExtHeaderChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIpv6ExtHeader_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Routing = nil
	obj.routingHolder = nil

	if value == FlowIpv6ExtHeaderChoice.ROUTING {
		obj.obj.Routing = NewFlowIpv6Routing().msg()
	}

	return obj
}

// description is TBD
// NextHeader returns a PatternFlowIpv6ExtHeaderNextHeader
func (obj *flowIpv6ExtHeader) NextHeader() PatternFlowIpv6ExtHeaderNextHeader {
	if obj.obj.NextHeader == nil {
		obj.obj.NextHeader = NewPatternFlowIpv6ExtHeaderNextHeader().msg()
	}
	if obj.nextHeaderHolder == nil {
		obj.nextHeaderHolder = &patternFlowIpv6ExtHeaderNextHeader{obj: obj.obj.NextHeader}
	}
	return obj.nextHeaderHolder
}

// description is TBD
// NextHeader returns a PatternFlowIpv6ExtHeaderNextHeader
func (obj *flowIpv6ExtHeader) HasNextHeader() bool {
	return obj.obj.NextHeader != nil
}

// description is TBD
// SetNextHeader sets the PatternFlowIpv6ExtHeaderNextHeader value in the FlowIpv6ExtHeader object
func (obj *flowIpv6ExtHeader) SetNextHeader(value PatternFlowIpv6ExtHeaderNextHeader) FlowIpv6ExtHeader {

	obj.nextHeaderHolder = nil
	obj.obj.NextHeader = value.msg()

	return obj
}

// description is TBD
// HdrExtLen returns a PatternFlowIpv6ExtHeaderHdrExtLen
func (obj *flowIpv6ExtHeader) HdrExtLen() PatternFlowIpv6ExtHeaderHdrExtLen {
	if obj.obj.HdrExtLen == nil {
		obj.obj.HdrExtLen = NewPatternFlowIpv6ExtHeaderHdrExtLen().msg()
	}
	if obj.hdrExtLenHolder == nil {
		obj.hdrExtLenHolder = &patternFlowIpv6ExtHeaderHdrExtLen{obj: obj.obj.HdrExtLen}
	}
	return obj.hdrExtLenHolder
}

// description is TBD
// HdrExtLen returns a PatternFlowIpv6ExtHeaderHdrExtLen
func (obj *flowIpv6ExtHeader) HasHdrExtLen() bool {
	return obj.obj.HdrExtLen != nil
}

// description is TBD
// SetHdrExtLen sets the PatternFlowIpv6ExtHeaderHdrExtLen value in the FlowIpv6ExtHeader object
func (obj *flowIpv6ExtHeader) SetHdrExtLen(value PatternFlowIpv6ExtHeaderHdrExtLen) FlowIpv6ExtHeader {

	obj.hdrExtLenHolder = nil
	obj.obj.HdrExtLen = value.msg()

	return obj
}

// description is TBD
// Routing returns a FlowIpv6Routing
func (obj *flowIpv6ExtHeader) Routing() FlowIpv6Routing {
	if obj.obj.Routing == nil {
		obj.setChoice(FlowIpv6ExtHeaderChoice.ROUTING)
	}
	if obj.routingHolder == nil {
		obj.routingHolder = &flowIpv6Routing{obj: obj.obj.Routing}
	}
	return obj.routingHolder
}

// description is TBD
// Routing returns a FlowIpv6Routing
func (obj *flowIpv6ExtHeader) HasRouting() bool {
	return obj.obj.Routing != nil
}

// description is TBD
// SetRouting sets the FlowIpv6Routing value in the FlowIpv6ExtHeader object
func (obj *flowIpv6ExtHeader) SetRouting(value FlowIpv6Routing) FlowIpv6ExtHeader {
	obj.setChoice(FlowIpv6ExtHeaderChoice.ROUTING)
	obj.routingHolder = nil
	obj.obj.Routing = value.msg()

	return obj
}

func (obj *flowIpv6ExtHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NextHeader != nil {

		obj.NextHeader().validateObj(vObj, set_default)
	}

	if obj.obj.HdrExtLen != nil {

		obj.HdrExtLen().validateObj(vObj, set_default)
	}

	if obj.obj.Routing != nil {

		obj.Routing().validateObj(vObj, set_default)
	}

}

func (obj *flowIpv6ExtHeader) setDefault() {
	var choices_set int = 0
	var choice FlowIpv6ExtHeaderChoiceEnum

	if obj.obj.Routing != nil {
		choices_set += 1
		choice = FlowIpv6ExtHeaderChoice.ROUTING
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIpv6ExtHeaderChoice.ROUTING)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIpv6ExtHeader")
			}
		} else {
			intVal := otg.FlowIpv6ExtHeader_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIpv6ExtHeader_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
