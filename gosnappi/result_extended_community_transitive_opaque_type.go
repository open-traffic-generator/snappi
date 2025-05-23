package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultExtendedCommunityTransitiveOpaqueType *****
type resultExtendedCommunityTransitiveOpaqueType struct {
	validation
	obj                        *otg.ResultExtendedCommunityTransitiveOpaqueType
	marshaller                 marshalResultExtendedCommunityTransitiveOpaqueType
	unMarshaller               unMarshalResultExtendedCommunityTransitiveOpaqueType
	colorSubtypeHolder         ResultExtendedCommunityTransitiveOpaqueTypeColor
	encapsulationSubtypeHolder ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
}

func NewResultExtendedCommunityTransitiveOpaqueType() ResultExtendedCommunityTransitiveOpaqueType {
	obj := resultExtendedCommunityTransitiveOpaqueType{obj: &otg.ResultExtendedCommunityTransitiveOpaqueType{}}
	obj.setDefault()
	return &obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) msg() *otg.ResultExtendedCommunityTransitiveOpaqueType {
	return obj.obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) setMsg(msg *otg.ResultExtendedCommunityTransitiveOpaqueType) ResultExtendedCommunityTransitiveOpaqueType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultExtendedCommunityTransitiveOpaqueType struct {
	obj *resultExtendedCommunityTransitiveOpaqueType
}

type marshalResultExtendedCommunityTransitiveOpaqueType interface {
	// ToProto marshals ResultExtendedCommunityTransitiveOpaqueType to protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueType
	ToProto() (*otg.ResultExtendedCommunityTransitiveOpaqueType, error)
	// ToPbText marshals ResultExtendedCommunityTransitiveOpaqueType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultExtendedCommunityTransitiveOpaqueType to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultExtendedCommunityTransitiveOpaqueType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals ResultExtendedCommunityTransitiveOpaqueType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalresultExtendedCommunityTransitiveOpaqueType struct {
	obj *resultExtendedCommunityTransitiveOpaqueType
}

type unMarshalResultExtendedCommunityTransitiveOpaqueType interface {
	// FromProto unmarshals ResultExtendedCommunityTransitiveOpaqueType from protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueType
	FromProto(msg *otg.ResultExtendedCommunityTransitiveOpaqueType) (ResultExtendedCommunityTransitiveOpaqueType, error)
	// FromPbText unmarshals ResultExtendedCommunityTransitiveOpaqueType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultExtendedCommunityTransitiveOpaqueType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultExtendedCommunityTransitiveOpaqueType from JSON text
	FromJson(value string) error
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) Marshal() marshalResultExtendedCommunityTransitiveOpaqueType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultExtendedCommunityTransitiveOpaqueType{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) Unmarshal() unMarshalResultExtendedCommunityTransitiveOpaqueType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultExtendedCommunityTransitiveOpaqueType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueType) ToProto() (*otg.ResultExtendedCommunityTransitiveOpaqueType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueType) FromProto(msg *otg.ResultExtendedCommunityTransitiveOpaqueType) (ResultExtendedCommunityTransitiveOpaqueType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueType) ToPbText() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueType) FromPbText(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueType) ToYaml() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueType) FromYaml(value string) error {
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

func (m *marshalresultExtendedCommunityTransitiveOpaqueType) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalresultExtendedCommunityTransitiveOpaqueType) ToJson() (string, error) {
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

func (m *unMarshalresultExtendedCommunityTransitiveOpaqueType) FromJson(value string) error {
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

func (obj *resultExtendedCommunityTransitiveOpaqueType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) Clone() (ResultExtendedCommunityTransitiveOpaqueType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultExtendedCommunityTransitiveOpaqueType()
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

func (obj *resultExtendedCommunityTransitiveOpaqueType) setNil() {
	obj.colorSubtypeHolder = nil
	obj.encapsulationSubtypeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultExtendedCommunityTransitiveOpaqueType is the Transitive Opaque Extended Community is sent as type 0x03.
type ResultExtendedCommunityTransitiveOpaqueType interface {
	Validation
	// msg marshals ResultExtendedCommunityTransitiveOpaqueType to protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueType
	// and doesn't set defaults
	msg() *otg.ResultExtendedCommunityTransitiveOpaqueType
	// setMsg unmarshals ResultExtendedCommunityTransitiveOpaqueType from protobuf object *otg.ResultExtendedCommunityTransitiveOpaqueType
	// and doesn't set defaults
	setMsg(*otg.ResultExtendedCommunityTransitiveOpaqueType) ResultExtendedCommunityTransitiveOpaqueType
	// provides marshal interface
	Marshal() marshalResultExtendedCommunityTransitiveOpaqueType
	// provides unmarshal interface
	Unmarshal() unMarshalResultExtendedCommunityTransitiveOpaqueType
	// validate validates ResultExtendedCommunityTransitiveOpaqueType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultExtendedCommunityTransitiveOpaqueType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum, set in ResultExtendedCommunityTransitiveOpaqueType
	Choice() ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum
	// setChoice assigns ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum provided by user to ResultExtendedCommunityTransitiveOpaqueType
	setChoice(value ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum) ResultExtendedCommunityTransitiveOpaqueType
	// HasChoice checks if Choice has been set in ResultExtendedCommunityTransitiveOpaqueType
	HasChoice() bool
	// ColorSubtype returns ResultExtendedCommunityTransitiveOpaqueTypeColor, set in ResultExtendedCommunityTransitiveOpaqueType.
	// ResultExtendedCommunityTransitiveOpaqueTypeColor is the Color Community contains locally administrator defined 'color' value which is used in conjunction with Encapsulation  attribute to decide whether a data packet can be transmitted on a certain tunnel or not. It is defined in RFC9012 and sent with sub-type as 0x0b.
	ColorSubtype() ResultExtendedCommunityTransitiveOpaqueTypeColor
	// SetColorSubtype assigns ResultExtendedCommunityTransitiveOpaqueTypeColor provided by user to ResultExtendedCommunityTransitiveOpaqueType.
	// ResultExtendedCommunityTransitiveOpaqueTypeColor is the Color Community contains locally administrator defined 'color' value which is used in conjunction with Encapsulation  attribute to decide whether a data packet can be transmitted on a certain tunnel or not. It is defined in RFC9012 and sent with sub-type as 0x0b.
	SetColorSubtype(value ResultExtendedCommunityTransitiveOpaqueTypeColor) ResultExtendedCommunityTransitiveOpaqueType
	// HasColorSubtype checks if ColorSubtype has been set in ResultExtendedCommunityTransitiveOpaqueType
	HasColorSubtype() bool
	// EncapsulationSubtype returns ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation, set in ResultExtendedCommunityTransitiveOpaqueType.
	// ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation is this identifies the type of tunneling technology being signalled. It is defined in RFC9012 and sent with sub-type as 0x0c.
	EncapsulationSubtype() ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
	// SetEncapsulationSubtype assigns ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation provided by user to ResultExtendedCommunityTransitiveOpaqueType.
	// ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation is this identifies the type of tunneling technology being signalled. It is defined in RFC9012 and sent with sub-type as 0x0c.
	SetEncapsulationSubtype(value ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ResultExtendedCommunityTransitiveOpaqueType
	// HasEncapsulationSubtype checks if EncapsulationSubtype has been set in ResultExtendedCommunityTransitiveOpaqueType
	HasEncapsulationSubtype() bool
	setNil()
}

type ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum string

// Enum of Choice on ResultExtendedCommunityTransitiveOpaqueType
var ResultExtendedCommunityTransitiveOpaqueTypeChoice = struct {
	COLOR_SUBTYPE         ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum
	ENCAPSULATION_SUBTYPE ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum
}{
	COLOR_SUBTYPE:         ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum("color_subtype"),
	ENCAPSULATION_SUBTYPE: ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum("encapsulation_subtype"),
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) Choice() ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum {
	return ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *resultExtendedCommunityTransitiveOpaqueType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) setChoice(value ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum) ResultExtendedCommunityTransitiveOpaqueType {
	intValue, ok := otg.ResultExtendedCommunityTransitiveOpaqueType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ResultExtendedCommunityTransitiveOpaqueType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.EncapsulationSubtype = nil
	obj.encapsulationSubtypeHolder = nil
	obj.obj.ColorSubtype = nil
	obj.colorSubtypeHolder = nil

	if value == ResultExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE {
		obj.obj.ColorSubtype = NewResultExtendedCommunityTransitiveOpaqueTypeColor().msg()
	}

	if value == ResultExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE {
		obj.obj.EncapsulationSubtype = NewResultExtendedCommunityTransitiveOpaqueTypeEncapsulation().msg()
	}

	return obj
}

// description is TBD
// ColorSubtype returns a ResultExtendedCommunityTransitiveOpaqueTypeColor
func (obj *resultExtendedCommunityTransitiveOpaqueType) ColorSubtype() ResultExtendedCommunityTransitiveOpaqueTypeColor {
	if obj.obj.ColorSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE)
	}
	if obj.colorSubtypeHolder == nil {
		obj.colorSubtypeHolder = &resultExtendedCommunityTransitiveOpaqueTypeColor{obj: obj.obj.ColorSubtype}
	}
	return obj.colorSubtypeHolder
}

// description is TBD
// ColorSubtype returns a ResultExtendedCommunityTransitiveOpaqueTypeColor
func (obj *resultExtendedCommunityTransitiveOpaqueType) HasColorSubtype() bool {
	return obj.obj.ColorSubtype != nil
}

// description is TBD
// SetColorSubtype sets the ResultExtendedCommunityTransitiveOpaqueTypeColor value in the ResultExtendedCommunityTransitiveOpaqueType object
func (obj *resultExtendedCommunityTransitiveOpaqueType) SetColorSubtype(value ResultExtendedCommunityTransitiveOpaqueTypeColor) ResultExtendedCommunityTransitiveOpaqueType {
	obj.setChoice(ResultExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE)
	obj.colorSubtypeHolder = nil
	obj.obj.ColorSubtype = value.msg()

	return obj
}

// description is TBD
// EncapsulationSubtype returns a ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
func (obj *resultExtendedCommunityTransitiveOpaqueType) EncapsulationSubtype() ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation {
	if obj.obj.EncapsulationSubtype == nil {
		obj.setChoice(ResultExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE)
	}
	if obj.encapsulationSubtypeHolder == nil {
		obj.encapsulationSubtypeHolder = &resultExtendedCommunityTransitiveOpaqueTypeEncapsulation{obj: obj.obj.EncapsulationSubtype}
	}
	return obj.encapsulationSubtypeHolder
}

// description is TBD
// EncapsulationSubtype returns a ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation
func (obj *resultExtendedCommunityTransitiveOpaqueType) HasEncapsulationSubtype() bool {
	return obj.obj.EncapsulationSubtype != nil
}

// description is TBD
// SetEncapsulationSubtype sets the ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation value in the ResultExtendedCommunityTransitiveOpaqueType object
func (obj *resultExtendedCommunityTransitiveOpaqueType) SetEncapsulationSubtype(value ResultExtendedCommunityTransitiveOpaqueTypeEncapsulation) ResultExtendedCommunityTransitiveOpaqueType {
	obj.setChoice(ResultExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE)
	obj.encapsulationSubtypeHolder = nil
	obj.obj.EncapsulationSubtype = value.msg()

	return obj
}

func (obj *resultExtendedCommunityTransitiveOpaqueType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ColorSubtype != nil {

		obj.ColorSubtype().validateObj(vObj, set_default)
	}

	if obj.obj.EncapsulationSubtype != nil {

		obj.EncapsulationSubtype().validateObj(vObj, set_default)
	}

}

func (obj *resultExtendedCommunityTransitiveOpaqueType) setDefault() {
	var choices_set int = 0
	var choice ResultExtendedCommunityTransitiveOpaqueTypeChoiceEnum

	if obj.obj.ColorSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitiveOpaqueTypeChoice.COLOR_SUBTYPE
	}

	if obj.obj.EncapsulationSubtype != nil {
		choices_set += 1
		choice = ResultExtendedCommunityTransitiveOpaqueTypeChoice.ENCAPSULATION_SUBTYPE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ResultExtendedCommunityTransitiveOpaqueType")
			}
		} else {
			intVal := otg.ResultExtendedCommunityTransitiveOpaqueType_Choice_Enum_value[string(choice)]
			enumValue := otg.ResultExtendedCommunityTransitiveOpaqueType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
