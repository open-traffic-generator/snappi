package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortOptionsFrameOrderingMode *****
type portOptionsFrameOrderingMode struct {
	validation
	obj          *otg.PortOptionsFrameOrderingMode
	marshaller   marshalPortOptionsFrameOrderingMode
	unMarshaller unMarshalPortOptionsFrameOrderingMode
}

func NewPortOptionsFrameOrderingMode() PortOptionsFrameOrderingMode {
	obj := portOptionsFrameOrderingMode{obj: &otg.PortOptionsFrameOrderingMode{}}
	obj.setDefault()
	return &obj
}

func (obj *portOptionsFrameOrderingMode) msg() *otg.PortOptionsFrameOrderingMode {
	return obj.obj
}

func (obj *portOptionsFrameOrderingMode) setMsg(msg *otg.PortOptionsFrameOrderingMode) PortOptionsFrameOrderingMode {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportOptionsFrameOrderingMode struct {
	obj *portOptionsFrameOrderingMode
}

type marshalPortOptionsFrameOrderingMode interface {
	// ToProto marshals PortOptionsFrameOrderingMode to protobuf object *otg.PortOptionsFrameOrderingMode
	ToProto() (*otg.PortOptionsFrameOrderingMode, error)
	// ToPbText marshals PortOptionsFrameOrderingMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortOptionsFrameOrderingMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortOptionsFrameOrderingMode to JSON text
	ToJson() (string, error)
}

type unMarshalportOptionsFrameOrderingMode struct {
	obj *portOptionsFrameOrderingMode
}

type unMarshalPortOptionsFrameOrderingMode interface {
	// FromProto unmarshals PortOptionsFrameOrderingMode from protobuf object *otg.PortOptionsFrameOrderingMode
	FromProto(msg *otg.PortOptionsFrameOrderingMode) (PortOptionsFrameOrderingMode, error)
	// FromPbText unmarshals PortOptionsFrameOrderingMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortOptionsFrameOrderingMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortOptionsFrameOrderingMode from JSON text
	FromJson(value string) error
}

func (obj *portOptionsFrameOrderingMode) Marshal() marshalPortOptionsFrameOrderingMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportOptionsFrameOrderingMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *portOptionsFrameOrderingMode) Unmarshal() unMarshalPortOptionsFrameOrderingMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportOptionsFrameOrderingMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportOptionsFrameOrderingMode) ToProto() (*otg.PortOptionsFrameOrderingMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportOptionsFrameOrderingMode) FromProto(msg *otg.PortOptionsFrameOrderingMode) (PortOptionsFrameOrderingMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportOptionsFrameOrderingMode) ToPbText() (string, error) {
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

func (m *unMarshalportOptionsFrameOrderingMode) FromPbText(value string) error {
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

func (m *marshalportOptionsFrameOrderingMode) ToYaml() (string, error) {
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

func (m *unMarshalportOptionsFrameOrderingMode) FromYaml(value string) error {
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

func (m *marshalportOptionsFrameOrderingMode) ToJson() (string, error) {
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

func (m *unMarshalportOptionsFrameOrderingMode) FromJson(value string) error {
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

func (obj *portOptionsFrameOrderingMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portOptionsFrameOrderingMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portOptionsFrameOrderingMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portOptionsFrameOrderingMode) Clone() (PortOptionsFrameOrderingMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortOptionsFrameOrderingMode()
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

// PortOptionsFrameOrderingMode is controls how an implementation arranges (interleaves) the frames of multiple flows that share source ports and destination ports.
// A flow transmits frames from its source port(s) to its destination port(s). When multiple flows transmit simultaneously towards a common destination port, the transmit ordering determines how the frames of the different flows are interleaved on the wire from the source ports. This is independent of the per-flow payload sequence used for loss and latency measurement.
type PortOptionsFrameOrderingMode interface {
	Validation
	// msg marshals PortOptionsFrameOrderingMode to protobuf object *otg.PortOptionsFrameOrderingMode
	// and doesn't set defaults
	msg() *otg.PortOptionsFrameOrderingMode
	// setMsg unmarshals PortOptionsFrameOrderingMode from protobuf object *otg.PortOptionsFrameOrderingMode
	// and doesn't set defaults
	setMsg(*otg.PortOptionsFrameOrderingMode) PortOptionsFrameOrderingMode
	// provides marshal interface
	Marshal() marshalPortOptionsFrameOrderingMode
	// provides unmarshal interface
	Unmarshal() unMarshalPortOptionsFrameOrderingMode
	// validate validates PortOptionsFrameOrderingMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortOptionsFrameOrderingMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PortOptionsFrameOrderingModeChoiceEnum, set in PortOptionsFrameOrderingMode
	Choice() PortOptionsFrameOrderingModeChoiceEnum
	// setChoice assigns PortOptionsFrameOrderingModeChoiceEnum provided by user to PortOptionsFrameOrderingMode
	setChoice(value PortOptionsFrameOrderingModeChoiceEnum) PortOptionsFrameOrderingMode
	// HasChoice checks if Choice has been set in PortOptionsFrameOrderingMode
	HasChoice() bool
	// getter for NoOrdering to set choice.
	NoOrdering()
	// getter for Rfc2889 to set choice.
	Rfc2889()
}

type PortOptionsFrameOrderingModeChoiceEnum string

// Enum of Choice on PortOptionsFrameOrderingMode
var PortOptionsFrameOrderingModeChoice = struct {
	NO_ORDERING PortOptionsFrameOrderingModeChoiceEnum
	RFC2889     PortOptionsFrameOrderingModeChoiceEnum
}{
	NO_ORDERING: PortOptionsFrameOrderingModeChoiceEnum("no_ordering"),
	RFC2889:     PortOptionsFrameOrderingModeChoiceEnum("rfc2889"),
}

func (obj *portOptionsFrameOrderingMode) Choice() PortOptionsFrameOrderingModeChoiceEnum {
	return PortOptionsFrameOrderingModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoOrdering to set choice
func (obj *portOptionsFrameOrderingMode) NoOrdering() {
	obj.setChoice(PortOptionsFrameOrderingModeChoice.NO_ORDERING)
}

// getter for Rfc2889 to set choice
func (obj *portOptionsFrameOrderingMode) Rfc2889() {
	obj.setChoice(PortOptionsFrameOrderingModeChoice.RFC2889)
}

// The type of frame ordering mode.
// Choice returns a string
func (obj *portOptionsFrameOrderingMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *portOptionsFrameOrderingMode) setChoice(value PortOptionsFrameOrderingModeChoiceEnum) PortOptionsFrameOrderingMode {
	intValue, ok := otg.PortOptionsFrameOrderingMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PortOptionsFrameOrderingModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.PortOptionsFrameOrderingMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *portOptionsFrameOrderingMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *portOptionsFrameOrderingMode) setDefault() {
	var choices_set int = 0
	var choice PortOptionsFrameOrderingModeChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PortOptionsFrameOrderingModeChoice.NO_ORDERING)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PortOptionsFrameOrderingMode")
			}
		} else {
			intVal := otg.PortOptionsFrameOrderingMode_Choice_Enum_value[string(choice)]
			enumValue := otg.PortOptionsFrameOrderingMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
