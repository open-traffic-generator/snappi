package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowLacp *****
type flowLacp struct {
	validation
	obj           *otg.FlowLacp
	marshaller    marshalFlowLacp
	unMarshaller  unMarshalFlowLacp
	versionHolder PatternFlowLacpVersion
	lacpduHolder  FlowLacpLacpdu
}

func NewFlowLacp() FlowLacp {
	obj := flowLacp{obj: &otg.FlowLacp{}}
	obj.setDefault()
	return &obj
}

func (obj *flowLacp) msg() *otg.FlowLacp {
	return obj.obj
}

func (obj *flowLacp) setMsg(msg *otg.FlowLacp) FlowLacp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowLacp struct {
	obj *flowLacp
}

type marshalFlowLacp interface {
	// ToProto marshals FlowLacp to protobuf object *otg.FlowLacp
	ToProto() (*otg.FlowLacp, error)
	// ToPbText marshals FlowLacp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowLacp to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowLacp to JSON text
	ToJson() (string, error)
}

type unMarshalflowLacp struct {
	obj *flowLacp
}

type unMarshalFlowLacp interface {
	// FromProto unmarshals FlowLacp from protobuf object *otg.FlowLacp
	FromProto(msg *otg.FlowLacp) (FlowLacp, error)
	// FromPbText unmarshals FlowLacp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowLacp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowLacp from JSON text
	FromJson(value string) error
}

func (obj *flowLacp) Marshal() marshalFlowLacp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowLacp{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowLacp) Unmarshal() unMarshalFlowLacp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowLacp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowLacp) ToProto() (*otg.FlowLacp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowLacp) FromProto(msg *otg.FlowLacp) (FlowLacp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowLacp) ToPbText() (string, error) {
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

func (m *unMarshalflowLacp) FromPbText(value string) error {
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

func (m *marshalflowLacp) ToYaml() (string, error) {
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

func (m *unMarshalflowLacp) FromYaml(value string) error {
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

func (m *marshalflowLacp) ToJson() (string, error) {
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

func (m *unMarshalflowLacp) FromJson(value string) error {
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

func (obj *flowLacp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowLacp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowLacp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowLacp) Clone() (FlowLacp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowLacp()
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

func (obj *flowLacp) setNil() {
	obj.versionHolder = nil
	obj.lacpduHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowLacp is defines the fields of a Link Aggregation Control Protocol (LACP) Data Unit (PDU) as specified by IEEE 802.3ad. The ethernet ether type, if set to auto, will be updated to '0x8809' and the destination address, if set to auto, will be updated to "01:80:C2:00:00:02" for LACP. Flows[i].metrics.enable should be set to false to avoid insertion of additional instrumentation bytes in the generated packets.
type FlowLacp interface {
	Validation
	// msg marshals FlowLacp to protobuf object *otg.FlowLacp
	// and doesn't set defaults
	msg() *otg.FlowLacp
	// setMsg unmarshals FlowLacp from protobuf object *otg.FlowLacp
	// and doesn't set defaults
	setMsg(*otg.FlowLacp) FlowLacp
	// provides marshal interface
	Marshal() marshalFlowLacp
	// provides unmarshal interface
	Unmarshal() unMarshalFlowLacp
	// validate validates FlowLacp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowLacp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowLacpChoiceEnum, set in FlowLacp
	Choice() FlowLacpChoiceEnum
	// setChoice assigns FlowLacpChoiceEnum provided by user to FlowLacp
	setChoice(value FlowLacpChoiceEnum) FlowLacp
	// HasChoice checks if Choice has been set in FlowLacp
	HasChoice() bool
	// Version returns PatternFlowLacpVersion, set in FlowLacp.
	// PatternFlowLacpVersion is the LACP version number.
	Version() PatternFlowLacpVersion
	// SetVersion assigns PatternFlowLacpVersion provided by user to FlowLacp.
	// PatternFlowLacpVersion is the LACP version number.
	SetVersion(value PatternFlowLacpVersion) FlowLacp
	// HasVersion checks if Version has been set in FlowLacp
	HasVersion() bool
	// Lacpdu returns FlowLacpLacpdu, set in FlowLacp.
	// FlowLacpLacpdu is defines the TLV fields of a Link Aggregation Control Protocol (LACP)
	// Data Unit (PDU) as specified by IEEE 802.3ad.
	Lacpdu() FlowLacpLacpdu
	// SetLacpdu assigns FlowLacpLacpdu provided by user to FlowLacp.
	// FlowLacpLacpdu is defines the TLV fields of a Link Aggregation Control Protocol (LACP)
	// Data Unit (PDU) as specified by IEEE 802.3ad.
	SetLacpdu(value FlowLacpLacpdu) FlowLacp
	// HasLacpdu checks if Lacpdu has been set in FlowLacp
	HasLacpdu() bool
	setNil()
}

type FlowLacpChoiceEnum string

// Enum of Choice on FlowLacp
var FlowLacpChoice = struct {
	LACPDU FlowLacpChoiceEnum
}{
	LACPDU: FlowLacpChoiceEnum("lacpdu"),
}

func (obj *flowLacp) Choice() FlowLacpChoiceEnum {
	return FlowLacpChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowLacp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowLacp) setChoice(value FlowLacpChoiceEnum) FlowLacp {
	intValue, ok := otg.FlowLacp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowLacpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowLacp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Lacpdu = nil
	obj.lacpduHolder = nil

	if value == FlowLacpChoice.LACPDU {
		obj.obj.Lacpdu = NewFlowLacpLacpdu().msg()
	}

	return obj
}

// description is TBD
// Version returns a PatternFlowLacpVersion
func (obj *flowLacp) Version() PatternFlowLacpVersion {
	if obj.obj.Version == nil {
		obj.obj.Version = NewPatternFlowLacpVersion().msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &patternFlowLacpVersion{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a PatternFlowLacpVersion
func (obj *flowLacp) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the PatternFlowLacpVersion value in the FlowLacp object
func (obj *flowLacp) SetVersion(value PatternFlowLacpVersion) FlowLacp {

	obj.versionHolder = nil
	obj.obj.Version = value.msg()

	return obj
}

// description is TBD
// Lacpdu returns a FlowLacpLacpdu
func (obj *flowLacp) Lacpdu() FlowLacpLacpdu {
	if obj.obj.Lacpdu == nil {
		obj.setChoice(FlowLacpChoice.LACPDU)
	}
	if obj.lacpduHolder == nil {
		obj.lacpduHolder = &flowLacpLacpdu{obj: obj.obj.Lacpdu}
	}
	return obj.lacpduHolder
}

// description is TBD
// Lacpdu returns a FlowLacpLacpdu
func (obj *flowLacp) HasLacpdu() bool {
	return obj.obj.Lacpdu != nil
}

// description is TBD
// SetLacpdu sets the FlowLacpLacpdu value in the FlowLacp object
func (obj *flowLacp) SetLacpdu(value FlowLacpLacpdu) FlowLacp {
	obj.setChoice(FlowLacpChoice.LACPDU)
	obj.lacpduHolder = nil
	obj.obj.Lacpdu = value.msg()

	return obj
}

func (obj *flowLacp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

	if obj.obj.Lacpdu != nil {

		obj.Lacpdu().validateObj(vObj, set_default)
	}

}

func (obj *flowLacp) setDefault() {
	var choices_set int = 0
	var choice FlowLacpChoiceEnum

	if obj.obj.Lacpdu != nil {
		choices_set += 1
		choice = FlowLacpChoice.LACPDU
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowLacpChoice.LACPDU)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowLacp")
			}
		} else {
			intVal := otg.FlowLacp_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowLacp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
