package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingTxOffset *****
type egressOnlyTrackingTxOffset struct {
	validation
	obj          *otg.EgressOnlyTrackingTxOffset
	marshaller   marshalEgressOnlyTrackingTxOffset
	unMarshaller unMarshalEgressOnlyTrackingTxOffset
	customHolder EgressOnlyTrackingTxOffsetCustom
}

func NewEgressOnlyTrackingTxOffset() EgressOnlyTrackingTxOffset {
	obj := egressOnlyTrackingTxOffset{obj: &otg.EgressOnlyTrackingTxOffset{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingTxOffset) msg() *otg.EgressOnlyTrackingTxOffset {
	return obj.obj
}

func (obj *egressOnlyTrackingTxOffset) setMsg(msg *otg.EgressOnlyTrackingTxOffset) EgressOnlyTrackingTxOffset {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingTxOffset struct {
	obj *egressOnlyTrackingTxOffset
}

type marshalEgressOnlyTrackingTxOffset interface {
	// ToProto marshals EgressOnlyTrackingTxOffset to protobuf object *otg.EgressOnlyTrackingTxOffset
	ToProto() (*otg.EgressOnlyTrackingTxOffset, error)
	// ToPbText marshals EgressOnlyTrackingTxOffset to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingTxOffset to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingTxOffset to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingTxOffset struct {
	obj *egressOnlyTrackingTxOffset
}

type unMarshalEgressOnlyTrackingTxOffset interface {
	// FromProto unmarshals EgressOnlyTrackingTxOffset from protobuf object *otg.EgressOnlyTrackingTxOffset
	FromProto(msg *otg.EgressOnlyTrackingTxOffset) (EgressOnlyTrackingTxOffset, error)
	// FromPbText unmarshals EgressOnlyTrackingTxOffset from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingTxOffset from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingTxOffset from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingTxOffset) Marshal() marshalEgressOnlyTrackingTxOffset {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingTxOffset{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingTxOffset) Unmarshal() unMarshalEgressOnlyTrackingTxOffset {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingTxOffset{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingTxOffset) ToProto() (*otg.EgressOnlyTrackingTxOffset, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingTxOffset) FromProto(msg *otg.EgressOnlyTrackingTxOffset) (EgressOnlyTrackingTxOffset, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingTxOffset) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxOffset) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingTxOffset) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxOffset) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingTxOffset) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxOffset) FromJson(value string) error {
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

func (obj *egressOnlyTrackingTxOffset) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTxOffset) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTxOffset) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingTxOffset) Clone() (EgressOnlyTrackingTxOffset, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingTxOffset()
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

func (obj *egressOnlyTrackingTxOffset) setNil() {
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// EgressOnlyTrackingTxOffset is a container of Tx offset properties.
type EgressOnlyTrackingTxOffset interface {
	Validation
	// msg marshals EgressOnlyTrackingTxOffset to protobuf object *otg.EgressOnlyTrackingTxOffset
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingTxOffset
	// setMsg unmarshals EgressOnlyTrackingTxOffset from protobuf object *otg.EgressOnlyTrackingTxOffset
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingTxOffset) EgressOnlyTrackingTxOffset
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingTxOffset
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingTxOffset
	// validate validates EgressOnlyTrackingTxOffset
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingTxOffset, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns EgressOnlyTrackingTxOffsetChoiceEnum, set in EgressOnlyTrackingTxOffset
	Choice() EgressOnlyTrackingTxOffsetChoiceEnum
	// setChoice assigns EgressOnlyTrackingTxOffsetChoiceEnum provided by user to EgressOnlyTrackingTxOffset
	setChoice(value EgressOnlyTrackingTxOffsetChoiceEnum) EgressOnlyTrackingTxOffset
	// HasChoice checks if Choice has been set in EgressOnlyTrackingTxOffset
	HasChoice() bool
	// getter for Auto to set choice.
	Auto()
	// Custom returns EgressOnlyTrackingTxOffsetCustom, set in EgressOnlyTrackingTxOffset.
	// EgressOnlyTrackingTxOffsetCustom is a container of custom Tx offset properties.
	Custom() EgressOnlyTrackingTxOffsetCustom
	// SetCustom assigns EgressOnlyTrackingTxOffsetCustom provided by user to EgressOnlyTrackingTxOffset.
	// EgressOnlyTrackingTxOffsetCustom is a container of custom Tx offset properties.
	SetCustom(value EgressOnlyTrackingTxOffsetCustom) EgressOnlyTrackingTxOffset
	// HasCustom checks if Custom has been set in EgressOnlyTrackingTxOffset
	HasCustom() bool
	setNil()
}

type EgressOnlyTrackingTxOffsetChoiceEnum string

// Enum of Choice on EgressOnlyTrackingTxOffset
var EgressOnlyTrackingTxOffsetChoice = struct {
	AUTO   EgressOnlyTrackingTxOffsetChoiceEnum
	CUSTOM EgressOnlyTrackingTxOffsetChoiceEnum
}{
	AUTO:   EgressOnlyTrackingTxOffsetChoiceEnum("auto"),
	CUSTOM: EgressOnlyTrackingTxOffsetChoiceEnum("custom"),
}

func (obj *egressOnlyTrackingTxOffset) Choice() EgressOnlyTrackingTxOffsetChoiceEnum {
	return EgressOnlyTrackingTxOffsetChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Auto to set choice
func (obj *egressOnlyTrackingTxOffset) Auto() {
	obj.setChoice(EgressOnlyTrackingTxOffsetChoice.AUTO)
}

// Choose "auto" when optional Tx statistics in egress only tracking are required while fetching egress only stats by opting "tx_metric" in get_metrics - egress_only_tracking - tagged_metric - metric_names.  or both Tx and Rx offset of tracked field in packet is the same. Otherwise choose "custom" when the Tx and Rx offsets are different due to DUT adding/ modifying or deleting encapsulation protocol header e.g. when the egress-only tracked packets are MACsec encapulated and the Tx and Rx side have different configuration such that the MACSec header is added/modified or removed.
// Choice returns a string
func (obj *egressOnlyTrackingTxOffset) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *egressOnlyTrackingTxOffset) setChoice(value EgressOnlyTrackingTxOffsetChoiceEnum) EgressOnlyTrackingTxOffset {
	intValue, ok := otg.EgressOnlyTrackingTxOffset_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on EgressOnlyTrackingTxOffsetChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.EgressOnlyTrackingTxOffset_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil

	if value == EgressOnlyTrackingTxOffsetChoice.CUSTOM {
		obj.obj.Custom = NewEgressOnlyTrackingTxOffsetCustom().msg()
	}

	return obj
}

// A container of custom Tx offset.
// Custom returns a EgressOnlyTrackingTxOffsetCustom
func (obj *egressOnlyTrackingTxOffset) Custom() EgressOnlyTrackingTxOffsetCustom {
	if obj.obj.Custom == nil {
		obj.setChoice(EgressOnlyTrackingTxOffsetChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &egressOnlyTrackingTxOffsetCustom{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// A container of custom Tx offset.
// Custom returns a EgressOnlyTrackingTxOffsetCustom
func (obj *egressOnlyTrackingTxOffset) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A container of custom Tx offset.
// SetCustom sets the EgressOnlyTrackingTxOffsetCustom value in the EgressOnlyTrackingTxOffset object
func (obj *egressOnlyTrackingTxOffset) SetCustom(value EgressOnlyTrackingTxOffsetCustom) EgressOnlyTrackingTxOffset {
	obj.setChoice(EgressOnlyTrackingTxOffsetChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *egressOnlyTrackingTxOffset) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

}

func (obj *egressOnlyTrackingTxOffset) setDefault() {
	var choices_set int = 0
	var choice EgressOnlyTrackingTxOffsetChoiceEnum

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = EgressOnlyTrackingTxOffsetChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(EgressOnlyTrackingTxOffsetChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in EgressOnlyTrackingTxOffset")
			}
		} else {
			intVal := otg.EgressOnlyTrackingTxOffset_Choice_Enum_value[string(choice)]
			enumValue := otg.EgressOnlyTrackingTxOffset_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
