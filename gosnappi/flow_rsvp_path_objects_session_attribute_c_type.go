package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowRSVPPathObjectsSessionAttributeCType *****
type flowRSVPPathObjectsSessionAttributeCType struct {
	validation
	obj               *otg.FlowRSVPPathObjectsSessionAttributeCType
	marshaller        marshalFlowRSVPPathObjectsSessionAttributeCType
	unMarshaller      unMarshalFlowRSVPPathObjectsSessionAttributeCType
	lspTunnelHolder   FlowRSVPPathSessionAttributeLspTunnel
	lspTunnelRaHolder FlowRSVPPathSessionAttributeLspTunnelRa
}

func NewFlowRSVPPathObjectsSessionAttributeCType() FlowRSVPPathObjectsSessionAttributeCType {
	obj := flowRSVPPathObjectsSessionAttributeCType{obj: &otg.FlowRSVPPathObjectsSessionAttributeCType{}}
	obj.setDefault()
	return &obj
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) msg() *otg.FlowRSVPPathObjectsSessionAttributeCType {
	return obj.obj
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) setMsg(msg *otg.FlowRSVPPathObjectsSessionAttributeCType) FlowRSVPPathObjectsSessionAttributeCType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowRSVPPathObjectsSessionAttributeCType struct {
	obj *flowRSVPPathObjectsSessionAttributeCType
}

type marshalFlowRSVPPathObjectsSessionAttributeCType interface {
	// ToProto marshals FlowRSVPPathObjectsSessionAttributeCType to protobuf object *otg.FlowRSVPPathObjectsSessionAttributeCType
	ToProto() (*otg.FlowRSVPPathObjectsSessionAttributeCType, error)
	// ToPbText marshals FlowRSVPPathObjectsSessionAttributeCType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowRSVPPathObjectsSessionAttributeCType to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowRSVPPathObjectsSessionAttributeCType to JSON text
	ToJson() (string, error)
}

type unMarshalflowRSVPPathObjectsSessionAttributeCType struct {
	obj *flowRSVPPathObjectsSessionAttributeCType
}

type unMarshalFlowRSVPPathObjectsSessionAttributeCType interface {
	// FromProto unmarshals FlowRSVPPathObjectsSessionAttributeCType from protobuf object *otg.FlowRSVPPathObjectsSessionAttributeCType
	FromProto(msg *otg.FlowRSVPPathObjectsSessionAttributeCType) (FlowRSVPPathObjectsSessionAttributeCType, error)
	// FromPbText unmarshals FlowRSVPPathObjectsSessionAttributeCType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowRSVPPathObjectsSessionAttributeCType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowRSVPPathObjectsSessionAttributeCType from JSON text
	FromJson(value string) error
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) Marshal() marshalFlowRSVPPathObjectsSessionAttributeCType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowRSVPPathObjectsSessionAttributeCType{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) Unmarshal() unMarshalFlowRSVPPathObjectsSessionAttributeCType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowRSVPPathObjectsSessionAttributeCType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowRSVPPathObjectsSessionAttributeCType) ToProto() (*otg.FlowRSVPPathObjectsSessionAttributeCType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowRSVPPathObjectsSessionAttributeCType) FromProto(msg *otg.FlowRSVPPathObjectsSessionAttributeCType) (FlowRSVPPathObjectsSessionAttributeCType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowRSVPPathObjectsSessionAttributeCType) ToPbText() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSessionAttributeCType) FromPbText(value string) error {
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

func (m *marshalflowRSVPPathObjectsSessionAttributeCType) ToYaml() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSessionAttributeCType) FromYaml(value string) error {
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

func (m *marshalflowRSVPPathObjectsSessionAttributeCType) ToJson() (string, error) {
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

func (m *unMarshalflowRSVPPathObjectsSessionAttributeCType) FromJson(value string) error {
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

func (obj *flowRSVPPathObjectsSessionAttributeCType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) Clone() (FlowRSVPPathObjectsSessionAttributeCType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowRSVPPathObjectsSessionAttributeCType()
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

func (obj *flowRSVPPathObjectsSessionAttributeCType) setNil() {
	obj.lspTunnelHolder = nil
	obj.lspTunnelRaHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowRSVPPathObjectsSessionAttributeCType is object for SESSION_ATTRIBUTE class. Currently supported c-type is LSP_Tunnel_RA (1) and LSP_Tunnel (7).
type FlowRSVPPathObjectsSessionAttributeCType interface {
	Validation
	// msg marshals FlowRSVPPathObjectsSessionAttributeCType to protobuf object *otg.FlowRSVPPathObjectsSessionAttributeCType
	// and doesn't set defaults
	msg() *otg.FlowRSVPPathObjectsSessionAttributeCType
	// setMsg unmarshals FlowRSVPPathObjectsSessionAttributeCType from protobuf object *otg.FlowRSVPPathObjectsSessionAttributeCType
	// and doesn't set defaults
	setMsg(*otg.FlowRSVPPathObjectsSessionAttributeCType) FlowRSVPPathObjectsSessionAttributeCType
	// provides marshal interface
	Marshal() marshalFlowRSVPPathObjectsSessionAttributeCType
	// provides unmarshal interface
	Unmarshal() unMarshalFlowRSVPPathObjectsSessionAttributeCType
	// validate validates FlowRSVPPathObjectsSessionAttributeCType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowRSVPPathObjectsSessionAttributeCType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum, set in FlowRSVPPathObjectsSessionAttributeCType
	Choice() FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum
	// setChoice assigns FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum provided by user to FlowRSVPPathObjectsSessionAttributeCType
	setChoice(value FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum) FlowRSVPPathObjectsSessionAttributeCType
	// HasChoice checks if Choice has been set in FlowRSVPPathObjectsSessionAttributeCType
	HasChoice() bool
	// LspTunnel returns FlowRSVPPathSessionAttributeLspTunnel, set in FlowRSVPPathObjectsSessionAttributeCType.
	// FlowRSVPPathSessionAttributeLspTunnel is sESSION_ATTRIBUTE class = 207, LSP_TUNNEL_RA C-Type = 7, resource affinity information.
	LspTunnel() FlowRSVPPathSessionAttributeLspTunnel
	// SetLspTunnel assigns FlowRSVPPathSessionAttributeLspTunnel provided by user to FlowRSVPPathObjectsSessionAttributeCType.
	// FlowRSVPPathSessionAttributeLspTunnel is sESSION_ATTRIBUTE class = 207, LSP_TUNNEL_RA C-Type = 7, resource affinity information.
	SetLspTunnel(value FlowRSVPPathSessionAttributeLspTunnel) FlowRSVPPathObjectsSessionAttributeCType
	// HasLspTunnel checks if LspTunnel has been set in FlowRSVPPathObjectsSessionAttributeCType
	HasLspTunnel() bool
	// LspTunnelRa returns FlowRSVPPathSessionAttributeLspTunnelRa, set in FlowRSVPPathObjectsSessionAttributeCType.
	// FlowRSVPPathSessionAttributeLspTunnelRa is sESSION_ATTRIBUTE class = 207, LSP_TUNNEL_RA C-Type = 1, it carries resource affinity information.
	LspTunnelRa() FlowRSVPPathSessionAttributeLspTunnelRa
	// SetLspTunnelRa assigns FlowRSVPPathSessionAttributeLspTunnelRa provided by user to FlowRSVPPathObjectsSessionAttributeCType.
	// FlowRSVPPathSessionAttributeLspTunnelRa is sESSION_ATTRIBUTE class = 207, LSP_TUNNEL_RA C-Type = 1, it carries resource affinity information.
	SetLspTunnelRa(value FlowRSVPPathSessionAttributeLspTunnelRa) FlowRSVPPathObjectsSessionAttributeCType
	// HasLspTunnelRa checks if LspTunnelRa has been set in FlowRSVPPathObjectsSessionAttributeCType
	HasLspTunnelRa() bool
	setNil()
}

type FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum string

// Enum of Choice on FlowRSVPPathObjectsSessionAttributeCType
var FlowRSVPPathObjectsSessionAttributeCTypeChoice = struct {
	LSP_TUNNEL    FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum
	LSP_TUNNEL_RA FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum
}{
	LSP_TUNNEL:    FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum("lsp_tunnel"),
	LSP_TUNNEL_RA: FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum("lsp_tunnel_ra"),
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) Choice() FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum {
	return FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowRSVPPathObjectsSessionAttributeCType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) setChoice(value FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum) FlowRSVPPathObjectsSessionAttributeCType {
	intValue, ok := otg.FlowRSVPPathObjectsSessionAttributeCType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowRSVPPathObjectsSessionAttributeCTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowRSVPPathObjectsSessionAttributeCType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.LspTunnelRa = nil
	obj.lspTunnelRaHolder = nil
	obj.obj.LspTunnel = nil
	obj.lspTunnelHolder = nil

	if value == FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL {
		obj.obj.LspTunnel = NewFlowRSVPPathSessionAttributeLspTunnel().msg()
	}

	if value == FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL_RA {
		obj.obj.LspTunnelRa = NewFlowRSVPPathSessionAttributeLspTunnelRa().msg()
	}

	return obj
}

// description is TBD
// LspTunnel returns a FlowRSVPPathSessionAttributeLspTunnel
func (obj *flowRSVPPathObjectsSessionAttributeCType) LspTunnel() FlowRSVPPathSessionAttributeLspTunnel {
	if obj.obj.LspTunnel == nil {
		obj.setChoice(FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL)
	}
	if obj.lspTunnelHolder == nil {
		obj.lspTunnelHolder = &flowRSVPPathSessionAttributeLspTunnel{obj: obj.obj.LspTunnel}
	}
	return obj.lspTunnelHolder
}

// description is TBD
// LspTunnel returns a FlowRSVPPathSessionAttributeLspTunnel
func (obj *flowRSVPPathObjectsSessionAttributeCType) HasLspTunnel() bool {
	return obj.obj.LspTunnel != nil
}

// description is TBD
// SetLspTunnel sets the FlowRSVPPathSessionAttributeLspTunnel value in the FlowRSVPPathObjectsSessionAttributeCType object
func (obj *flowRSVPPathObjectsSessionAttributeCType) SetLspTunnel(value FlowRSVPPathSessionAttributeLspTunnel) FlowRSVPPathObjectsSessionAttributeCType {
	obj.setChoice(FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL)
	obj.lspTunnelHolder = nil
	obj.obj.LspTunnel = value.msg()

	return obj
}

// description is TBD
// LspTunnelRa returns a FlowRSVPPathSessionAttributeLspTunnelRa
func (obj *flowRSVPPathObjectsSessionAttributeCType) LspTunnelRa() FlowRSVPPathSessionAttributeLspTunnelRa {
	if obj.obj.LspTunnelRa == nil {
		obj.setChoice(FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL_RA)
	}
	if obj.lspTunnelRaHolder == nil {
		obj.lspTunnelRaHolder = &flowRSVPPathSessionAttributeLspTunnelRa{obj: obj.obj.LspTunnelRa}
	}
	return obj.lspTunnelRaHolder
}

// description is TBD
// LspTunnelRa returns a FlowRSVPPathSessionAttributeLspTunnelRa
func (obj *flowRSVPPathObjectsSessionAttributeCType) HasLspTunnelRa() bool {
	return obj.obj.LspTunnelRa != nil
}

// description is TBD
// SetLspTunnelRa sets the FlowRSVPPathSessionAttributeLspTunnelRa value in the FlowRSVPPathObjectsSessionAttributeCType object
func (obj *flowRSVPPathObjectsSessionAttributeCType) SetLspTunnelRa(value FlowRSVPPathSessionAttributeLspTunnelRa) FlowRSVPPathObjectsSessionAttributeCType {
	obj.setChoice(FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL_RA)
	obj.lspTunnelRaHolder = nil
	obj.obj.LspTunnelRa = value.msg()

	return obj
}

func (obj *flowRSVPPathObjectsSessionAttributeCType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LspTunnel != nil {

		obj.LspTunnel().validateObj(vObj, set_default)
	}

	if obj.obj.LspTunnelRa != nil {

		obj.LspTunnelRa().validateObj(vObj, set_default)
	}

}

func (obj *flowRSVPPathObjectsSessionAttributeCType) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(FlowRSVPPathObjectsSessionAttributeCTypeChoice.LSP_TUNNEL)

	}

}
