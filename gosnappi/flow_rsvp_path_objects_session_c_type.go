package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsSessionCType *****
type flowRSVPPathObjectsSessionCType struct {
	validation
	obj                 *otg.FlowRSVPPathObjectsSessionCType
	marshaller          marshalFlowRSVPPathObjectsSessionCType
	unMarshaller        unMarshalFlowRSVPPathObjectsSessionCType
	lspTunnelIpv4Holder FlowRSVPPathSessionLspTunnelIpv4
}

func NewFlowRSVPPathObjectsSessionCType() FlowRSVPPathObjectsSessionCType {
	obj := flowRSVPPathObjectsSessionCType{obj: &otg.FlowRSVPPathObjectsSessionCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsSessionCType) msg() *otg.FlowRSVPPathObjectsSessionCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsSessionCType) setMsg(msg *otg.FlowRSVPPathObjectsSessionCType) FlowRSVPPathObjectsSessionCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsSessionCType struct {
	obj *flowRSVPPathObjectsSessionCType
}

type marshalFlowRSVPPathObjectsSessionCType interface {
	// ToProto marshals FlowRSVPPathObjectsSessionCType to protobuf object *otg.FlowRSVPPathObjectsSessionCType
	ToProto() (*otg.FlowRSVPPathObjectsSessionCType, error)
	// ToPbText marshals FlowRSVPPathObjectsSessionCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsSessionCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsSessionCType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsSessionCType struct {
	obj *flowRSVPPathObjectsSessionCType
}

type unMarshalFlowRSVPPathObjectsSessionCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsSessionCType from protobuf object *otg.FlowRSVPPathObjectsSessionCType
	FromProto(msg *otg.FlowRSVPPathObjectsSessionCType) (FlowRSVPPathObjectsSessionCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsSessionCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsSessionCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsSessionCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsSessionCType) Marshal() marshalFlowRSVPPathObjectsSessionCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsSessionCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsSessionCType) Unmarshal() unMarshalFlowRSVPPathObjectsSessionCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsSessionCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsSessionCType) ToProto() (*otg.FlowRSVPPathObjectsSessionCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsSessionCType) FromProto(msg *otg.FlowRSVPPathObjectsSessionCType) (FlowRSVPPathObjectsSessionCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsSessionCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSessionCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsSessionCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSessionCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsSessionCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSessionCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsSessionCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSessionCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSessionCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsSessionCType) Clone() (FlowRSVPPathObjectsSessionCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsSessionCType()
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

func (obj *flowRSVPPathObjectsSessionCType) setNil() {
	obj.lspTunnelIpv4Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsSessionCType is the body of an object corresponding to the class number and c-type. Currently supported c-type for SESSION object is LSP Tunnel IPv4 (7).
type FlowRSVPPathObjectsSessionCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsSessionCType to protobuf object *otg.FlowRSVPPathObjectsSessionCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsSessionCType
	// setMsg unmarshals FlowRSVPPathObjectsSessionCType from protobuf object *otg.FlowRSVPPathObjectsSessionCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsSessionCType) FlowRSVPPathObjectsSessionCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsSessionCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsSessionCType
	// validate validates FlowRSVPPathObjectsSessionCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsSessionCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsSessionCTypeChoiceEnum, set in FlowRSVPPathObjectsSessionCType
	Choice() FlowRSVPPathObjectsSessionCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsSessionCTypeChoiceEnum provided by user to FlowRSVPPathObjectsSessionCType
	setChoice(value FlowRSVPPathObjectsSessionCTypeChoiceEnum) FlowRSVPPathObjectsSessionCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsSessionCType
	HasChoice() bool
	// LspTunnelIpv4 returns FlowRSVPPathSessionLspTunnelIpv4, set in FlowRSVPPathObjectsSessionCType.
	// FlowRSVPPathSessionLspTunnelIpv4 is class = SESSION, LSP_TUNNEL_IPv4 C-Type = 7.
	LspTunnelIpv4() FlowRSVPPathSessionLspTunnelIpv4
	// SetLspTunnelIpv4 assigns FlowRSVPPathSessionLspTunnelIpv4 provided by user to FlowRSVPPathObjectsSessionCType.
	// FlowRSVPPathSessionLspTunnelIpv4 is class = SESSION, LSP_TUNNEL_IPv4 C-Type = 7.
	SetLspTunnelIpv4(value FlowRSVPPathSessionLspTunnelIpv4) FlowRSVPPathObjectsSessionCType
	// HasLspTunnelIpv4 checks if LspTunnelIpv4 has been set in FlowRSVPPathObjectsSessionCType
	HasLspTunnelIpv4() bool
	setNil()
}

type FlowRSVPPathObjectsSessionCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsSessionCType
var FlowRSVPPathObjectsSessionCTypeChoice = struct {
	LSP_TUNNEL_IPV4 FlowRSVPPathObjectsSessionCTypeChoiceEnum
}{
	LSP_TUNNEL_IPV4: FlowRSVPPathObjectsSessionCTypeChoiceEnum("lsp_tunnel_ipv4"),
}

func (obj *flowRSVPPathObjectsSessionCType) Choice() FlowRSVPPathObjectsSessionCTypeChoiceEnum {
	return FlowRSVPPathObjectsSessionCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsSessionCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsSessionCType) setChoice(value FlowRSVPPathObjectsSessionCTypeChoiceEnum) FlowRSVPPathObjectsSessionCType {
	intValue, ok := otg.FlowRSVPPathObjectsSessionCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsSessionCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsSessionCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LspTunnelIpv4 = nil
	obj.lspTunnelIpv4Holder = nil

	if value == FlowRSVPPathObjectsSessionCTypeChoice.LSP_TUNNEL_IPV4 {
		obj.obj.LspTunnelIpv4 = NewFlowRSVPPathSessionLspTunnelIpv4().msg()
	}

	return obj
}

// description is TBD
// LspTunnelIpv4 returns a FlowRSVPPathSessionLspTunnelIpv4
func (obj *flowRSVPPathObjectsSessionCType) LspTunnelIpv4() FlowRSVPPathSessionLspTunnelIpv4 {
	if obj.obj.LspTunnelIpv4 == nil {
		obj.setChoice(FlowRSVPPathObjectsSessionCTypeChoice.LSP_TUNNEL_IPV4)
	}
	if obj.lspTunnelIpv4Holder == nil {
		obj.lspTunnelIpv4Holder = &flowRSVPPathSessionLspTunnelIpv4{obj: obj.obj.LspTunnelIpv4}
	}
	return obj.lspTunnelIpv4Holder
}

// description is TBD
// LspTunnelIpv4 returns a FlowRSVPPathSessionLspTunnelIpv4
func (obj *flowRSVPPathObjectsSessionCType) HasLspTunnelIpv4() bool {
	return obj.obj.LspTunnelIpv4 != nil
}

// description is TBD
// SetLspTunnelIpv4 sets the FlowRSVPPathSessionLspTunnelIpv4 value in the FlowRSVPPathObjectsSessionCType object
func (obj *flowRSVPPathObjectsSessionCType) SetLspTunnelIpv4(value FlowRSVPPathSessionLspTunnelIpv4) FlowRSVPPathObjectsSessionCType {
	obj.setChoice(FlowRSVPPathObjectsSessionCTypeChoice.LSP_TUNNEL_IPV4)
	obj.lspTunnelIpv4Holder = nil
	obj.obj.LspTunnelIpv4 = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsSessionCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LspTunnelIpv4 != nil {

		obj.LspTunnelIpv4().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsSessionCType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPPathObjectsSessionCTypeChoice.LSP_TUNNEL_IPV4)

	}

}
