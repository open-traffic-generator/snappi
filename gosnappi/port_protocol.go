package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortProtocol *****
type portProtocol struct {
	validation
	obj          *otg.PortProtocol
	marshaller   marshalPortProtocol
	unMarshaller unMarshalPortProtocol
	rocev2Holder Rocev2PerPortSettings
	macsecHolder MacsecPerPortSettings
}

func NewPortProtocol() PortProtocol {
	obj := portProtocol{obj: &otg.PortProtocol{}}
	obj.setDefault()
	return &obj
}

func (obj *portProtocol) msg() *otg.PortProtocol {
	return obj.obj
}

func (obj *portProtocol) setMsg(msg *otg.PortProtocol) PortProtocol {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportProtocol struct {
	obj *portProtocol
}

type marshalPortProtocol interface {
	// ToProto marshals PortProtocol to protobuf object *otg.PortProtocol
	ToProto() (*otg.PortProtocol, error)
	// ToPbText marshals PortProtocol to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortProtocol to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortProtocol to JSON text
	ToJson() (string, error)
}

type unMarshalportProtocol struct {
	obj *portProtocol
}

type unMarshalPortProtocol interface {
	// FromProto unmarshals PortProtocol from protobuf object *otg.PortProtocol
	FromProto(msg *otg.PortProtocol) (PortProtocol, error)
	// FromPbText unmarshals PortProtocol from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortProtocol from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortProtocol from JSON text
	FromJson(value string) error
}

func (obj *portProtocol) Marshal() marshalPortProtocol {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportProtocol{obj: obj}
	}
	return obj.marshaller
}

func (obj *portProtocol) Unmarshal() unMarshalPortProtocol {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportProtocol{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportProtocol) ToProto() (*otg.PortProtocol, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportProtocol) FromProto(msg *otg.PortProtocol) (PortProtocol, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportProtocol) ToPbText() (string, error) {
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

func (m *unMarshalportProtocol) FromPbText(value string) error {
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

func (m *marshalportProtocol) ToYaml() (string, error) {
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

func (m *unMarshalportProtocol) FromYaml(value string) error {
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

func (m *marshalportProtocol) ToJson() (string, error) {
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

func (m *unMarshalportProtocol) FromJson(value string) error {
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

func (obj *portProtocol) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portProtocol) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portProtocol) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portProtocol) Clone() (PortProtocol, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortProtocol()
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

func (obj *portProtocol) setNil() {
	obj.rocev2Holder = nil
	obj.macsecHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PortProtocol is supprted protocols.
type PortProtocol interface {
	Validation
	// msg marshals PortProtocol to protobuf object *otg.PortProtocol
	// and doesn't set defaults
	msg() *otg.PortProtocol
	// setMsg unmarshals PortProtocol from protobuf object *otg.PortProtocol
	// and doesn't set defaults
	setMsg(*otg.PortProtocol) PortProtocol
	// provides marshal interface
	Marshal() marshalPortProtocol
	// provides unmarshal interface
	Unmarshal() unMarshalPortProtocol
	// validate validates PortProtocol
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortProtocol, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PortProtocolChoiceEnum, set in PortProtocol
	Choice() PortProtocolChoiceEnum
	// setChoice assigns PortProtocolChoiceEnum provided by user to PortProtocol
	setChoice(value PortProtocolChoiceEnum) PortProtocol
	// HasChoice checks if Choice has been set in PortProtocol
	HasChoice() bool
	// Rocev2 returns Rocev2PerPortSettings, set in PortProtocol.
	// Rocev2PerPortSettings is data plane traffic flow configuration for a test port.
	Rocev2() Rocev2PerPortSettings
	// SetRocev2 assigns Rocev2PerPortSettings provided by user to PortProtocol.
	// Rocev2PerPortSettings is data plane traffic flow configuration for a test port.
	SetRocev2(value Rocev2PerPortSettings) PortProtocol
	// HasRocev2 checks if Rocev2 has been set in PortProtocol
	HasRocev2() bool
	// Macsec returns MacsecPerPortSettings, set in PortProtocol.
	// MacsecPerPortSettings is mACsec hardware acceleration settings of a test port.
	Macsec() MacsecPerPortSettings
	// SetMacsec assigns MacsecPerPortSettings provided by user to PortProtocol.
	// MacsecPerPortSettings is mACsec hardware acceleration settings of a test port.
	SetMacsec(value MacsecPerPortSettings) PortProtocol
	// HasMacsec checks if Macsec has been set in PortProtocol
	HasMacsec() bool
	setNil()
}

type PortProtocolChoiceEnum string

// Enum of Choice on PortProtocol
var PortProtocolChoice = struct {
	ROCEV2 PortProtocolChoiceEnum
	MACSEC PortProtocolChoiceEnum
}{
	ROCEV2: PortProtocolChoiceEnum("rocev2"),
	MACSEC: PortProtocolChoiceEnum("macsec"),
}

func (obj *portProtocol) Choice() PortProtocolChoiceEnum {
	return PortProtocolChoiceEnum(obj.obj.Choice.Enum().String())
}

// list of protocols that have per port settings
// Choice returns a string
func (obj *portProtocol) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *portProtocol) setChoice(value PortProtocolChoiceEnum) PortProtocol {
	intValue, ok := otg.PortProtocol_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PortProtocolChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PortProtocol_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Macsec = nil
	obj.macsecHolder = nil
	obj.obj.Rocev2 = nil
	obj.rocev2Holder = nil

	if value == PortProtocolChoice.ROCEV2 {
		obj.obj.Rocev2 = NewRocev2PerPortSettings().msg()
	}

	if value == PortProtocolChoice.MACSEC {
		obj.obj.Macsec = NewMacsecPerPortSettings().msg()
	}

	return obj
}

// description is TBD
// Rocev2 returns a Rocev2PerPortSettings
func (obj *portProtocol) Rocev2() Rocev2PerPortSettings {
	if obj.obj.Rocev2 == nil {
		obj.setChoice(PortProtocolChoice.ROCEV2)
	}
	if obj.rocev2Holder == nil {
		obj.rocev2Holder = &rocev2PerPortSettings{obj: obj.obj.Rocev2}
	}
	return obj.rocev2Holder
}

// description is TBD
// Rocev2 returns a Rocev2PerPortSettings
func (obj *portProtocol) HasRocev2() bool {
	return obj.obj.Rocev2 != nil
}

// description is TBD
// SetRocev2 sets the Rocev2PerPortSettings value in the PortProtocol object
func (obj *portProtocol) SetRocev2(value Rocev2PerPortSettings) PortProtocol {
	obj.setChoice(PortProtocolChoice.ROCEV2)
	obj.rocev2Holder = nil
	obj.obj.Rocev2 = value.msg()

	return obj
}

// description is TBD
// Macsec returns a MacsecPerPortSettings
func (obj *portProtocol) Macsec() MacsecPerPortSettings {
	if obj.obj.Macsec == nil {
		obj.setChoice(PortProtocolChoice.MACSEC)
	}
	if obj.macsecHolder == nil {
		obj.macsecHolder = &macsecPerPortSettings{obj: obj.obj.Macsec}
	}
	return obj.macsecHolder
}

// description is TBD
// Macsec returns a MacsecPerPortSettings
func (obj *portProtocol) HasMacsec() bool {
	return obj.obj.Macsec != nil
}

// description is TBD
// SetMacsec sets the MacsecPerPortSettings value in the PortProtocol object
func (obj *portProtocol) SetMacsec(value MacsecPerPortSettings) PortProtocol {
	obj.setChoice(PortProtocolChoice.MACSEC)
	obj.macsecHolder = nil
	obj.obj.Macsec = value.msg()

	return obj
}

func (obj *portProtocol) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Rocev2 != nil {

		obj.Rocev2().validateObj(vObj, set_default)
	}

	if obj.obj.Macsec != nil {

		obj.Macsec().validateObj(vObj, set_default)
	}

}

func (obj *portProtocol) setDefault() {
	var choices_set int = 0
	var choice PortProtocolChoiceEnum

	if obj.obj.Rocev2 != nil {
		choices_set += 1
		choice = PortProtocolChoice.ROCEV2
	}

	if obj.obj.Macsec != nil {
		choices_set += 1
		choice = PortProtocolChoice.MACSEC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PortProtocolChoice.ROCEV2)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PortProtocol")
			}
		} else {
			intVal := otg.PortProtocol_Choice_Enum_value[string(choice)]
			enumValue := otg.PortProtocol_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
