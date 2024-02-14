package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsSenderTemplateCType *****
type flowRSVPPathObjectsSenderTemplateCType struct {
	validation
	obj                 *otg.FlowRSVPPathObjectsSenderTemplateCType
	marshaller          marshalFlowRSVPPathObjectsSenderTemplateCType
	unMarshaller        unMarshalFlowRSVPPathObjectsSenderTemplateCType
	lspTunnelIpv4Holder FlowRSVPPathSenderTemplateLspTunnelIpv4
}

func NewFlowRSVPPathObjectsSenderTemplateCType() FlowRSVPPathObjectsSenderTemplateCType {
	obj := flowRSVPPathObjectsSenderTemplateCType{obj: &otg.FlowRSVPPathObjectsSenderTemplateCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) msg() *otg.FlowRSVPPathObjectsSenderTemplateCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) setMsg(msg *otg.FlowRSVPPathObjectsSenderTemplateCType) FlowRSVPPathObjectsSenderTemplateCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsSenderTemplateCType struct {
	obj *flowRSVPPathObjectsSenderTemplateCType
}

type marshalFlowRSVPPathObjectsSenderTemplateCType interface {
	// ToProto marshals FlowRSVPPathObjectsSenderTemplateCType to protobuf object *otg.FlowRSVPPathObjectsSenderTemplateCType
	ToProto() (*otg.FlowRSVPPathObjectsSenderTemplateCType, error)
	// ToPbText marshals FlowRSVPPathObjectsSenderTemplateCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsSenderTemplateCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsSenderTemplateCType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsSenderTemplateCType struct {
	obj *flowRSVPPathObjectsSenderTemplateCType
}

type unMarshalFlowRSVPPathObjectsSenderTemplateCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsSenderTemplateCType from protobuf object *otg.FlowRSVPPathObjectsSenderTemplateCType
	FromProto(msg *otg.FlowRSVPPathObjectsSenderTemplateCType) (FlowRSVPPathObjectsSenderTemplateCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsSenderTemplateCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsSenderTemplateCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsSenderTemplateCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) Marshal() marshalFlowRSVPPathObjectsSenderTemplateCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsSenderTemplateCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) Unmarshal() unMarshalFlowRSVPPathObjectsSenderTemplateCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsSenderTemplateCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsSenderTemplateCType) ToProto() (*otg.FlowRSVPPathObjectsSenderTemplateCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsSenderTemplateCType) FromProto(msg *otg.FlowRSVPPathObjectsSenderTemplateCType) (FlowRSVPPathObjectsSenderTemplateCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsSenderTemplateCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSenderTemplateCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsSenderTemplateCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSenderTemplateCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsSenderTemplateCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSenderTemplateCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsSenderTemplateCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) Clone() (FlowRSVPPathObjectsSenderTemplateCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsSenderTemplateCType()
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

func (obj *flowRSVPPathObjectsSenderTemplateCType) setNil() {
	obj.lspTunnelIpv4Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsSenderTemplateCType is object for SENDER_TEMPLATE class. Currently supported c-type is LSP Tunnel IPv4 (7).
type FlowRSVPPathObjectsSenderTemplateCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsSenderTemplateCType to protobuf object *otg.FlowRSVPPathObjectsSenderTemplateCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsSenderTemplateCType
	// setMsg unmarshals FlowRSVPPathObjectsSenderTemplateCType from protobuf object *otg.FlowRSVPPathObjectsSenderTemplateCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsSenderTemplateCType) FlowRSVPPathObjectsSenderTemplateCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsSenderTemplateCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsSenderTemplateCType
	// validate validates FlowRSVPPathObjectsSenderTemplateCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsSenderTemplateCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum, set in FlowRSVPPathObjectsSenderTemplateCType
	Choice() FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum provided by user to FlowRSVPPathObjectsSenderTemplateCType
	setChoice(value FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum) FlowRSVPPathObjectsSenderTemplateCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsSenderTemplateCType
	HasChoice() bool
	// LspTunnelIpv4 returns FlowRSVPPathSenderTemplateLspTunnelIpv4, set in FlowRSVPPathObjectsSenderTemplateCType.
	// FlowRSVPPathSenderTemplateLspTunnelIpv4 is class = SENDER_TEMPLATE, LSP_TUNNEL_IPv4 C-Type = 7
	LspTunnelIpv4() FlowRSVPPathSenderTemplateLspTunnelIpv4
	// SetLspTunnelIpv4 assigns FlowRSVPPathSenderTemplateLspTunnelIpv4 provided by user to FlowRSVPPathObjectsSenderTemplateCType.
	// FlowRSVPPathSenderTemplateLspTunnelIpv4 is class = SENDER_TEMPLATE, LSP_TUNNEL_IPv4 C-Type = 7
	SetLspTunnelIpv4(value FlowRSVPPathSenderTemplateLspTunnelIpv4) FlowRSVPPathObjectsSenderTemplateCType
	// HasLspTunnelIpv4 checks if LspTunnelIpv4 has been set in FlowRSVPPathObjectsSenderTemplateCType
	HasLspTunnelIpv4() bool
	setNil()
}

type FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsSenderTemplateCType
var FlowRSVPPathObjectsSenderTemplateCTypeChoice = struct {
	LSP_TUNNEL_IPV4 FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum
}{
	LSP_TUNNEL_IPV4: FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum("lsp_tunnel_ipv4"),
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) Choice() FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum {
	return FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsSenderTemplateCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) setChoice(value FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum) FlowRSVPPathObjectsSenderTemplateCType {
	intValue, ok := otg.FlowRSVPPathObjectsSenderTemplateCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsSenderTemplateCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsSenderTemplateCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LspTunnelIpv4 = nil
	obj.lspTunnelIpv4Holder = nil

	if value == FlowRSVPPathObjectsSenderTemplateCTypeChoice.LSP_TUNNEL_IPV4 {
		obj.obj.LspTunnelIpv4 = NewFlowRSVPPathSenderTemplateLspTunnelIpv4().msg()
	}

	return obj
}

// description is TBD
// LspTunnelIpv4 returns a FlowRSVPPathSenderTemplateLspTunnelIpv4
func (obj *flowRSVPPathObjectsSenderTemplateCType) LspTunnelIpv4() FlowRSVPPathSenderTemplateLspTunnelIpv4 {
	if obj.obj.LspTunnelIpv4 == nil {
		obj.setChoice(FlowRSVPPathObjectsSenderTemplateCTypeChoice.LSP_TUNNEL_IPV4)
	}
	if obj.lspTunnelIpv4Holder == nil {
		obj.lspTunnelIpv4Holder = &flowRSVPPathSenderTemplateLspTunnelIpv4{obj: obj.obj.LspTunnelIpv4}
	}
	return obj.lspTunnelIpv4Holder
}

// description is TBD
// LspTunnelIpv4 returns a FlowRSVPPathSenderTemplateLspTunnelIpv4
func (obj *flowRSVPPathObjectsSenderTemplateCType) HasLspTunnelIpv4() bool {
	return obj.obj.LspTunnelIpv4 != nil
}

// description is TBD
// SetLspTunnelIpv4 sets the FlowRSVPPathSenderTemplateLspTunnelIpv4 value in the FlowRSVPPathObjectsSenderTemplateCType object
func (obj *flowRSVPPathObjectsSenderTemplateCType) SetLspTunnelIpv4(value FlowRSVPPathSenderTemplateLspTunnelIpv4) FlowRSVPPathObjectsSenderTemplateCType {
	obj.setChoice(FlowRSVPPathObjectsSenderTemplateCTypeChoice.LSP_TUNNEL_IPV4)
	obj.lspTunnelIpv4Holder = nil
	obj.obj.LspTunnelIpv4 = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsSenderTemplateCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LspTunnelIpv4 != nil {

		obj.LspTunnelIpv4().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsSenderTemplateCType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPPathObjectsSenderTemplateCTypeChoice.LSP_TUNNEL_IPV4)

	}

}
