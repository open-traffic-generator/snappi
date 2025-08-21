package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureRequestPackets *****
type captureRequestPackets struct {
	validation
	obj          *otg.CaptureRequestPackets
	marshaller   marshalCaptureRequestPackets
	unMarshaller unMarshalCaptureRequestPackets
	sliceHolder  CaptureRequestCaptureSlice
}

func NewCaptureRequestPackets() CaptureRequestPackets {
	obj := captureRequestPackets{obj: &otg.CaptureRequestPackets{}}
	obj.setDefault()
	return &obj
}

func (obj *captureRequestPackets) msg() *otg.CaptureRequestPackets {
	return obj.obj
}

func (obj *captureRequestPackets) setMsg(msg *otg.CaptureRequestPackets) CaptureRequestPackets {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureRequestPackets struct {
	obj *captureRequestPackets
}

type marshalCaptureRequestPackets interface {
	// ToProto marshals CaptureRequestPackets to protobuf object *otg.CaptureRequestPackets
	ToProto() (*otg.CaptureRequestPackets, error)
	// ToPbText marshals CaptureRequestPackets to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureRequestPackets to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureRequestPackets to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureRequestPackets struct {
	obj *captureRequestPackets
}

type unMarshalCaptureRequestPackets interface {
	// FromProto unmarshals CaptureRequestPackets from protobuf object *otg.CaptureRequestPackets
	FromProto(msg *otg.CaptureRequestPackets) (CaptureRequestPackets, error)
	// FromPbText unmarshals CaptureRequestPackets from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureRequestPackets from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureRequestPackets from JSON text
	FromJson(value string) error
}

func (obj *captureRequestPackets) Marshal() marshalCaptureRequestPackets {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureRequestPackets{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureRequestPackets) Unmarshal() unMarshalCaptureRequestPackets {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureRequestPackets{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureRequestPackets) ToProto() (*otg.CaptureRequestPackets, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureRequestPackets) FromProto(msg *otg.CaptureRequestPackets) (CaptureRequestPackets, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureRequestPackets) ToPbText() (string, error) {
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

func (m *unMarshalcaptureRequestPackets) FromPbText(value string) error {
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

func (m *marshalcaptureRequestPackets) ToYaml() (string, error) {
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

func (m *unMarshalcaptureRequestPackets) FromYaml(value string) error {
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

func (m *marshalcaptureRequestPackets) ToJson() (string, error) {
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

func (m *unMarshalcaptureRequestPackets) FromJson(value string) error {
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

func (obj *captureRequestPackets) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureRequestPackets) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureRequestPackets) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureRequestPackets) Clone() (CaptureRequestPackets, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureRequestPackets()
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

func (obj *captureRequestPackets) setNil() {
	obj.sliceHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CaptureRequestPackets is the packets to be captured on the given port.
type CaptureRequestPackets interface {
	Validation
	// msg marshals CaptureRequestPackets to protobuf object *otg.CaptureRequestPackets
	// and doesn't set defaults
	msg() *otg.CaptureRequestPackets
	// setMsg unmarshals CaptureRequestPackets from protobuf object *otg.CaptureRequestPackets
	// and doesn't set defaults
	setMsg(*otg.CaptureRequestPackets) CaptureRequestPackets
	// provides marshal interface
	Marshal() marshalCaptureRequestPackets
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureRequestPackets
	// validate validates CaptureRequestPackets
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureRequestPackets, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns CaptureRequestPacketsChoiceEnum, set in CaptureRequestPackets
	Choice() CaptureRequestPacketsChoiceEnum
	// setChoice assigns CaptureRequestPacketsChoiceEnum provided by user to CaptureRequestPackets
	setChoice(value CaptureRequestPacketsChoiceEnum) CaptureRequestPackets
	// HasChoice checks if Choice has been set in CaptureRequestPackets
	HasChoice() bool
	// getter for All to set choice.
	All()
	// Slice returns CaptureRequestCaptureSlice, set in CaptureRequestPackets.
	// CaptureRequestCaptureSlice is packets to be captured based on specification of capture slice through start index and packet count.
	Slice() CaptureRequestCaptureSlice
	// SetSlice assigns CaptureRequestCaptureSlice provided by user to CaptureRequestPackets.
	// CaptureRequestCaptureSlice is packets to be captured based on specification of capture slice through start index and packet count.
	SetSlice(value CaptureRequestCaptureSlice) CaptureRequestPackets
	// HasSlice checks if Slice has been set in CaptureRequestPackets
	HasSlice() bool
	setNil()
}

type CaptureRequestPacketsChoiceEnum string

// Enum of Choice on CaptureRequestPackets
var CaptureRequestPacketsChoice = struct {
	ALL   CaptureRequestPacketsChoiceEnum
	SLICE CaptureRequestPacketsChoiceEnum
}{
	ALL:   CaptureRequestPacketsChoiceEnum("all"),
	SLICE: CaptureRequestPacketsChoiceEnum("slice"),
}

func (obj *captureRequestPackets) Choice() CaptureRequestPacketsChoiceEnum {
	return CaptureRequestPacketsChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for All to set choice
func (obj *captureRequestPackets) All() {
	obj.setChoice(CaptureRequestPacketsChoice.ALL)
}

// description is TBD
// Choice returns a string
func (obj *captureRequestPackets) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *captureRequestPackets) setChoice(value CaptureRequestPacketsChoiceEnum) CaptureRequestPackets {
	intValue, ok := otg.CaptureRequestPackets_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on CaptureRequestPacketsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.CaptureRequestPackets_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Slice = nil
	obj.sliceHolder = nil

	if value == CaptureRequestPacketsChoice.SLICE {
		obj.obj.Slice = NewCaptureRequestCaptureSlice().msg()
	}

	return obj
}

// description is TBD
// Slice returns a CaptureRequestCaptureSlice
func (obj *captureRequestPackets) Slice() CaptureRequestCaptureSlice {
	if obj.obj.Slice == nil {
		obj.setChoice(CaptureRequestPacketsChoice.SLICE)
	}
	if obj.sliceHolder == nil {
		obj.sliceHolder = &captureRequestCaptureSlice{obj: obj.obj.Slice}
	}
	return obj.sliceHolder
}

// description is TBD
// Slice returns a CaptureRequestCaptureSlice
func (obj *captureRequestPackets) HasSlice() bool {
	return obj.obj.Slice != nil
}

// description is TBD
// SetSlice sets the CaptureRequestCaptureSlice value in the CaptureRequestPackets object
func (obj *captureRequestPackets) SetSlice(value CaptureRequestCaptureSlice) CaptureRequestPackets {
	obj.setChoice(CaptureRequestPacketsChoice.SLICE)
	obj.sliceHolder = nil
	obj.obj.Slice = value.msg()

	return obj
}

func (obj *captureRequestPackets) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Slice != nil {

		obj.Slice().validateObj(vObj, set_default)
	}

}

func (obj *captureRequestPackets) setDefault() {
	var choices_set int = 0
	var choice CaptureRequestPacketsChoiceEnum

	if obj.obj.Slice != nil {
		choices_set += 1
		choice = CaptureRequestPacketsChoice.SLICE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(CaptureRequestPacketsChoice.ALL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in CaptureRequestPackets")
			}
		} else {
			intVal := otg.CaptureRequestPackets_Choice_Enum_value[string(choice)]
			enumValue := otg.CaptureRequestPackets_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
