package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpFecLabel *****
type ldpFecLabel struct {
	validation
	obj             *otg.LdpFecLabel
	marshaller      marshalLdpFecLabel
	unMarshaller    unMarshalLdpFecLabel
	incrementHolder LdpFecLabelIncrement
}

func NewLdpFecLabel() LdpFecLabel {
	obj := ldpFecLabel{obj: &otg.LdpFecLabel{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpFecLabel) msg() *otg.LdpFecLabel {
	return obj.obj
}

func (obj *ldpFecLabel) setMsg(msg *otg.LdpFecLabel) LdpFecLabel {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpFecLabel struct {
	obj *ldpFecLabel
}

type marshalLdpFecLabel interface {
	// ToProto marshals LdpFecLabel to protobuf object *otg.LdpFecLabel
	ToProto() (*otg.LdpFecLabel, error)
	// ToPbText marshals LdpFecLabel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpFecLabel to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpFecLabel to JSON text
	ToJson() (string, error)
}

type unMarshalldpFecLabel struct {
	obj *ldpFecLabel
}

type unMarshalLdpFecLabel interface {
	// FromProto unmarshals LdpFecLabel from protobuf object *otg.LdpFecLabel
	FromProto(msg *otg.LdpFecLabel) (LdpFecLabel, error)
	// FromPbText unmarshals LdpFecLabel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpFecLabel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpFecLabel from JSON text
	FromJson(value string) error
}

func (obj *ldpFecLabel) Marshal() marshalLdpFecLabel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpFecLabel{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpFecLabel) Unmarshal() unMarshalLdpFecLabel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpFecLabel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpFecLabel) ToProto() (*otg.LdpFecLabel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpFecLabel) FromProto(msg *otg.LdpFecLabel) (LdpFecLabel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpFecLabel) ToPbText() (string, error) {
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

func (m *unMarshalldpFecLabel) FromPbText(value string) error {
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

func (m *marshalldpFecLabel) ToYaml() (string, error) {
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

func (m *unMarshalldpFecLabel) FromYaml(value string) error {
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

func (m *marshalldpFecLabel) ToJson() (string, error) {
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

func (m *unMarshalldpFecLabel) FromJson(value string) error {
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

func (obj *ldpFecLabel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpFecLabel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpFecLabel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpFecLabel) Clone() (LdpFecLabel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpFecLabel()
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

func (obj *ldpFecLabel) setNil() {
	obj.incrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LdpFecLabel is the Generic Label (RFC 5036 Section 3.4.2.1) bound to each Prefix FEC of a Ldp.Ipv4FecRange. increment - the n-th prefix generated by addresses (counting across all entries, in order, from 0) gets label start + n. fixed - every prefix of the range gets the same label. Use 3 (Implicit NULL) or 0 (IPv4 Explicit NULL) to emulate an egress LSR (RFC 3032 Section 2.1).
type LdpFecLabel interface {
	Validation
	// msg marshals LdpFecLabel to protobuf object *otg.LdpFecLabel
	// and doesn't set defaults
	msg() *otg.LdpFecLabel
	// setMsg unmarshals LdpFecLabel from protobuf object *otg.LdpFecLabel
	// and doesn't set defaults
	setMsg(*otg.LdpFecLabel) LdpFecLabel
	// provides marshal interface
	Marshal() marshalLdpFecLabel
	// provides unmarshal interface
	Unmarshal() unMarshalLdpFecLabel
	// validate validates LdpFecLabel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpFecLabel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LdpFecLabelChoiceEnum, set in LdpFecLabel
	Choice() LdpFecLabelChoiceEnum
	// setChoice assigns LdpFecLabelChoiceEnum provided by user to LdpFecLabel
	setChoice(value LdpFecLabelChoiceEnum) LdpFecLabel
	// HasChoice checks if Choice has been set in LdpFecLabel
	HasChoice() bool
	// Increment returns LdpFecLabelIncrement, set in LdpFecLabel.
	// LdpFecLabelIncrement is labels that increase by 1 from one prefix to the next.
	Increment() LdpFecLabelIncrement
	// SetIncrement assigns LdpFecLabelIncrement provided by user to LdpFecLabel.
	// LdpFecLabelIncrement is labels that increase by 1 from one prefix to the next.
	SetIncrement(value LdpFecLabelIncrement) LdpFecLabel
	// HasIncrement checks if Increment has been set in LdpFecLabel
	HasIncrement() bool
	// Fixed returns uint32, set in LdpFecLabel.
	Fixed() uint32
	// SetFixed assigns uint32 provided by user to LdpFecLabel
	SetFixed(value uint32) LdpFecLabel
	// HasFixed checks if Fixed has been set in LdpFecLabel
	HasFixed() bool
	setNil()
}

type LdpFecLabelChoiceEnum string

// Enum of Choice on LdpFecLabel
var LdpFecLabelChoice = struct {
	INCREMENT LdpFecLabelChoiceEnum
	FIXED     LdpFecLabelChoiceEnum
}{
	INCREMENT: LdpFecLabelChoiceEnum("increment"),
	FIXED:     LdpFecLabelChoiceEnum("fixed"),
}

func (obj *ldpFecLabel) Choice() LdpFecLabelChoiceEnum {
	return LdpFecLabelChoiceEnum(obj.obj.Choice.Enum().String())
}

// How labels are assigned to the prefixes of the range.
// Choice returns a string
func (obj *ldpFecLabel) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ldpFecLabel) setChoice(value LdpFecLabelChoiceEnum) LdpFecLabel {
	intValue, ok := otg.LdpFecLabel_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LdpFecLabelChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LdpFecLabel_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Fixed = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil

	if value == LdpFecLabelChoice.INCREMENT {
		obj.obj.Increment = NewLdpFecLabelIncrement().msg()
	}

	if value == LdpFecLabelChoice.FIXED {
		defaultValue := uint32(3)
		obj.obj.Fixed = &defaultValue
	}

	return obj
}

// Labels that increase by 1 from one prefix to the next.
// Increment returns a LdpFecLabelIncrement
func (obj *ldpFecLabel) Increment() LdpFecLabelIncrement {
	if obj.obj.Increment == nil {
		obj.setChoice(LdpFecLabelChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &ldpFecLabelIncrement{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// Labels that increase by 1 from one prefix to the next.
// Increment returns a LdpFecLabelIncrement
func (obj *ldpFecLabel) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// Labels that increase by 1 from one prefix to the next.
// SetIncrement sets the LdpFecLabelIncrement value in the LdpFecLabel object
func (obj *ldpFecLabel) SetIncrement(value LdpFecLabelIncrement) LdpFecLabel {
	obj.setChoice(LdpFecLabelChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// The label bound to every prefix of the range. Meaningful values are 3 (Implicit NULL), 0 (IPv4 Explicit NULL) and 16 to 1048575. Values 1, 2 and 4 to 15 are reserved (RFC 3032 Section 2.1).
// Fixed returns a uint32
func (obj *ldpFecLabel) Fixed() uint32 {

	if obj.obj.Fixed == nil {
		obj.setChoice(LdpFecLabelChoice.FIXED)
	}

	return *obj.obj.Fixed

}

// The label bound to every prefix of the range. Meaningful values are 3 (Implicit NULL), 0 (IPv4 Explicit NULL) and 16 to 1048575. Values 1, 2 and 4 to 15 are reserved (RFC 3032 Section 2.1).
// Fixed returns a uint32
func (obj *ldpFecLabel) HasFixed() bool {
	return obj.obj.Fixed != nil
}

// The label bound to every prefix of the range. Meaningful values are 3 (Implicit NULL), 0 (IPv4 Explicit NULL) and 16 to 1048575. Values 1, 2 and 4 to 15 are reserved (RFC 3032 Section 2.1).
// SetFixed sets the uint32 value in the LdpFecLabel object
func (obj *ldpFecLabel) SetFixed(value uint32) LdpFecLabel {
	obj.setChoice(LdpFecLabelChoice.FIXED)
	obj.obj.Fixed = &value
	return obj
}

func (obj *ldpFecLabel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Fixed != nil {

		if *obj.obj.Fixed > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LdpFecLabel.Fixed <= 1048575 but Got %d", *obj.obj.Fixed))
		}

	}

}

func (obj *ldpFecLabel) setDefault() {
	var choices_set int = 0
	var choice LdpFecLabelChoiceEnum

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = LdpFecLabelChoice.INCREMENT
	}

	if obj.obj.Fixed != nil {
		choices_set += 1
		choice = LdpFecLabelChoice.FIXED
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LdpFecLabelChoice.INCREMENT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LdpFecLabel")
			}
		} else {
			intVal := otg.LdpFecLabel_Choice_Enum_value[string(choice)]
			enumValue := otg.LdpFecLabel_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
