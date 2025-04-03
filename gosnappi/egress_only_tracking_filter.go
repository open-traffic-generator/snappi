package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingFilter *****
type egressOnlyTrackingFilter struct {
	validation
	obj          *otg.EgressOnlyTrackingFilter
	marshaller   marshalEgressOnlyTrackingFilter
	unMarshaller unMarshalEgressOnlyTrackingFilter
}

func NewEgressOnlyTrackingFilter() EgressOnlyTrackingFilter {
	obj := egressOnlyTrackingFilter{obj: &otg.EgressOnlyTrackingFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingFilter) msg() *otg.EgressOnlyTrackingFilter {
	return obj.obj
}

func (obj *egressOnlyTrackingFilter) setMsg(msg *otg.EgressOnlyTrackingFilter) EgressOnlyTrackingFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingFilter struct {
	obj *egressOnlyTrackingFilter
}

type marshalEgressOnlyTrackingFilter interface {
	// ToProto marshals EgressOnlyTrackingFilter to protobuf object *otg.EgressOnlyTrackingFilter
	ToProto() (*otg.EgressOnlyTrackingFilter, error)
	// ToPbText marshals EgressOnlyTrackingFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingFilter to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingFilter struct {
	obj *egressOnlyTrackingFilter
}

type unMarshalEgressOnlyTrackingFilter interface {
	// FromProto unmarshals EgressOnlyTrackingFilter from protobuf object *otg.EgressOnlyTrackingFilter
	FromProto(msg *otg.EgressOnlyTrackingFilter) (EgressOnlyTrackingFilter, error)
	// FromPbText unmarshals EgressOnlyTrackingFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingFilter from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingFilter) Marshal() marshalEgressOnlyTrackingFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingFilter) Unmarshal() unMarshalEgressOnlyTrackingFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingFilter) ToProto() (*otg.EgressOnlyTrackingFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingFilter) FromProto(msg *otg.EgressOnlyTrackingFilter) (EgressOnlyTrackingFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingFilter) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingFilter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalegressOnlyTrackingFilter) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingFilter) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalegressOnlyTrackingFilter) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingFilter) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *egressOnlyTrackingFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingFilter) Clone() (EgressOnlyTrackingFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingFilter()
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

// EgressOnlyTrackingFilter is description is TBD
type EgressOnlyTrackingFilter interface {
	Validation
	// msg marshals EgressOnlyTrackingFilter to protobuf object *otg.EgressOnlyTrackingFilter
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingFilter
	// setMsg unmarshals EgressOnlyTrackingFilter from protobuf object *otg.EgressOnlyTrackingFilter
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingFilter) EgressOnlyTrackingFilter
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingFilter
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingFilter
	// validate validates EgressOnlyTrackingFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns EgressOnlyTrackingFilterChoiceEnum, set in EgressOnlyTrackingFilter
	Choice() EgressOnlyTrackingFilterChoiceEnum
	// setChoice assigns EgressOnlyTrackingFilterChoiceEnum provided by user to EgressOnlyTrackingFilter
	setChoice(value EgressOnlyTrackingFilterChoiceEnum) EgressOnlyTrackingFilter
	// HasChoice checks if Choice has been set in EgressOnlyTrackingFilter
	HasChoice() bool
	// getter for None to set choice.
	None()
	// getter for AutoMacsec to set choice.
	AutoMacsec()
}

type EgressOnlyTrackingFilterChoiceEnum string

// Enum of Choice on EgressOnlyTrackingFilter
var EgressOnlyTrackingFilterChoice = struct {
	NONE        EgressOnlyTrackingFilterChoiceEnum
	AUTO_MACSEC EgressOnlyTrackingFilterChoiceEnum
}{
	NONE:        EgressOnlyTrackingFilterChoiceEnum("none"),
	AUTO_MACSEC: EgressOnlyTrackingFilterChoiceEnum("auto_macsec"),
}

func (obj *egressOnlyTrackingFilter) Choice() EgressOnlyTrackingFilterChoiceEnum {
	return EgressOnlyTrackingFilterChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for None to set choice
func (obj *egressOnlyTrackingFilter) None() {
	obj.setChoice(EgressOnlyTrackingFilterChoice.NONE)
}

// getter for AutoMacsec to set choice
func (obj *egressOnlyTrackingFilter) AutoMacsec() {
	obj.setChoice(EgressOnlyTrackingFilterChoice.AUTO_MACSEC)
}

// If a packet does not match the filter it will not be considered for egress tracking. Currently two options are provided: none: All packets will be considered for egress only tracking. auto_macsec: This requires that MACsec enabled Ethernet interface should be configured on this port.  This filter will ensure that only packets with Ethernet Type set to MACsec (0x88E5) and  destined to traffic Rx device(s) will be considered for egress only tracking
// Choice returns a string
func (obj *egressOnlyTrackingFilter) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *egressOnlyTrackingFilter) setChoice(value EgressOnlyTrackingFilterChoiceEnum) EgressOnlyTrackingFilter {
	intValue, ok := otg.EgressOnlyTrackingFilter_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on EgressOnlyTrackingFilterChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.EgressOnlyTrackingFilter_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *egressOnlyTrackingFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *egressOnlyTrackingFilter) setDefault() {
	var choices_set int = 0
	var choice EgressOnlyTrackingFilterChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(EgressOnlyTrackingFilterChoice.NONE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in EgressOnlyTrackingFilter")
			}
		} else {
			intVal := otg.EgressOnlyTrackingFilter_Choice_Enum_value[string(choice)]
			enumValue := otg.EgressOnlyTrackingFilter_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
