package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureRequestCaptureSlice *****
type captureRequestCaptureSlice struct {
	validation
	obj           *otg.CaptureRequestCaptureSlice
	marshaller    marshalCaptureRequestCaptureSlice
	unMarshaller  unMarshalCaptureRequestCaptureSlice
	initialHolder CaptureRequestCaptureSliceInitial
}

func NewCaptureRequestCaptureSlice() CaptureRequestCaptureSlice {
	obj := captureRequestCaptureSlice{obj: &otg.CaptureRequestCaptureSlice{}}
	obj.setDefault()
	return &obj
}

func (obj *captureRequestCaptureSlice) msg() *otg.CaptureRequestCaptureSlice {
	return obj.obj
}

func (obj *captureRequestCaptureSlice) setMsg(msg *otg.CaptureRequestCaptureSlice) CaptureRequestCaptureSlice {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureRequestCaptureSlice struct {
	obj *captureRequestCaptureSlice
}

type marshalCaptureRequestCaptureSlice interface {
	// ToProto marshals CaptureRequestCaptureSlice to protobuf object *otg.CaptureRequestCaptureSlice
	ToProto() (*otg.CaptureRequestCaptureSlice, error)
	// ToPbText marshals CaptureRequestCaptureSlice to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureRequestCaptureSlice to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureRequestCaptureSlice to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureRequestCaptureSlice struct {
	obj *captureRequestCaptureSlice
}

type unMarshalCaptureRequestCaptureSlice interface {
	// FromProto unmarshals CaptureRequestCaptureSlice from protobuf object *otg.CaptureRequestCaptureSlice
	FromProto(msg *otg.CaptureRequestCaptureSlice) (CaptureRequestCaptureSlice, error)
	// FromPbText unmarshals CaptureRequestCaptureSlice from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureRequestCaptureSlice from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureRequestCaptureSlice from JSON text
	FromJson(value string) error
}

func (obj *captureRequestCaptureSlice) Marshal() marshalCaptureRequestCaptureSlice {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureRequestCaptureSlice{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureRequestCaptureSlice) Unmarshal() unMarshalCaptureRequestCaptureSlice {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureRequestCaptureSlice{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureRequestCaptureSlice) ToProto() (*otg.CaptureRequestCaptureSlice, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureRequestCaptureSlice) FromProto(msg *otg.CaptureRequestCaptureSlice) (CaptureRequestCaptureSlice, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureRequestCaptureSlice) ToPbText() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSlice) FromPbText(value string) error {
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

func (m *marshalcaptureRequestCaptureSlice) ToYaml() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSlice) FromYaml(value string) error {
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

func (m *marshalcaptureRequestCaptureSlice) ToJson() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSlice) FromJson(value string) error {
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

func (obj *captureRequestCaptureSlice) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureRequestCaptureSlice) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureRequestCaptureSlice) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureRequestCaptureSlice) Clone() (CaptureRequestCaptureSlice, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureRequestCaptureSlice()
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

func (obj *captureRequestCaptureSlice) setNil() {
	obj.initialHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CaptureRequestCaptureSlice is packets to be captured based on specification of capture slice i.e.
// position of first packet and count of packets to capture.
// To be noted,
// - definition of capture slice works in conjunction with capture filter parameters in set_config.
// - to get definitive outcome with large number of captured packets, 'overwrite' attribute in 'captures'
// settings of set_config should be disabled.
type CaptureRequestCaptureSlice interface {
	Validation
	// msg marshals CaptureRequestCaptureSlice to protobuf object *otg.CaptureRequestCaptureSlice
	// and doesn't set defaults
	msg() *otg.CaptureRequestCaptureSlice
	// setMsg unmarshals CaptureRequestCaptureSlice from protobuf object *otg.CaptureRequestCaptureSlice
	// and doesn't set defaults
	setMsg(*otg.CaptureRequestCaptureSlice) CaptureRequestCaptureSlice
	// provides marshal interface
	Marshal() marshalCaptureRequestCaptureSlice
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureRequestCaptureSlice
	// validate validates CaptureRequestCaptureSlice
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureRequestCaptureSlice, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns CaptureRequestCaptureSliceChoiceEnum, set in CaptureRequestCaptureSlice
	Choice() CaptureRequestCaptureSliceChoiceEnum
	// setChoice assigns CaptureRequestCaptureSliceChoiceEnum provided by user to CaptureRequestCaptureSlice
	setChoice(value CaptureRequestCaptureSliceChoiceEnum) CaptureRequestCaptureSlice
	// HasChoice checks if Choice has been set in CaptureRequestCaptureSlice
	HasChoice() bool
	// Initial returns CaptureRequestCaptureSliceInitial, set in CaptureRequestCaptureSlice.
	// CaptureRequestCaptureSliceInitial is specification of capture slice to capture packets from begining of captured packet sequence.
	Initial() CaptureRequestCaptureSliceInitial
	// SetInitial assigns CaptureRequestCaptureSliceInitial provided by user to CaptureRequestCaptureSlice.
	// CaptureRequestCaptureSliceInitial is specification of capture slice to capture packets from begining of captured packet sequence.
	SetInitial(value CaptureRequestCaptureSliceInitial) CaptureRequestCaptureSlice
	// HasInitial checks if Initial has been set in CaptureRequestCaptureSlice
	HasInitial() bool
	setNil()
}

type CaptureRequestCaptureSliceChoiceEnum string

// Enum of Choice on CaptureRequestCaptureSlice
var CaptureRequestCaptureSliceChoice = struct {
	INITIAL CaptureRequestCaptureSliceChoiceEnum
}{
	INITIAL: CaptureRequestCaptureSliceChoiceEnum("initial"),
}

func (obj *captureRequestCaptureSlice) Choice() CaptureRequestCaptureSliceChoiceEnum {
	return CaptureRequestCaptureSliceChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *captureRequestCaptureSlice) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *captureRequestCaptureSlice) setChoice(value CaptureRequestCaptureSliceChoiceEnum) CaptureRequestCaptureSlice {
	intValue, ok := otg.CaptureRequestCaptureSlice_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on CaptureRequestCaptureSliceChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.CaptureRequestCaptureSlice_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Initial = nil
	obj.initialHolder = nil

	if value == CaptureRequestCaptureSliceChoice.INITIAL {
		obj.obj.Initial = NewCaptureRequestCaptureSliceInitial().msg()
	}

	return obj
}

// description is TBD
// Initial returns a CaptureRequestCaptureSliceInitial
func (obj *captureRequestCaptureSlice) Initial() CaptureRequestCaptureSliceInitial {
	if obj.obj.Initial == nil {
		obj.setChoice(CaptureRequestCaptureSliceChoice.INITIAL)
	}
	if obj.initialHolder == nil {
		obj.initialHolder = &captureRequestCaptureSliceInitial{obj: obj.obj.Initial}
	}
	return obj.initialHolder
}

// description is TBD
// Initial returns a CaptureRequestCaptureSliceInitial
func (obj *captureRequestCaptureSlice) HasInitial() bool {
	return obj.obj.Initial != nil
}

// description is TBD
// SetInitial sets the CaptureRequestCaptureSliceInitial value in the CaptureRequestCaptureSlice object
func (obj *captureRequestCaptureSlice) SetInitial(value CaptureRequestCaptureSliceInitial) CaptureRequestCaptureSlice {
	obj.setChoice(CaptureRequestCaptureSliceChoice.INITIAL)
	obj.initialHolder = nil
	obj.obj.Initial = value.msg()

	return obj
}

func (obj *captureRequestCaptureSlice) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Initial != nil {

		obj.Initial().validateObj(vObj, set_default)
	}

}

func (obj *captureRequestCaptureSlice) setDefault() {
	var choices_set int = 0
	var choice CaptureRequestCaptureSliceChoiceEnum

	if obj.obj.Initial != nil {
		choices_set += 1
		choice = CaptureRequestCaptureSliceChoice.INITIAL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(CaptureRequestCaptureSliceChoice.INITIAL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in CaptureRequestCaptureSlice")
			}
		} else {
			intVal := otg.CaptureRequestCaptureSlice_Choice_Enum_value[string(choice)]
			enumValue := otg.CaptureRequestCaptureSlice_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
