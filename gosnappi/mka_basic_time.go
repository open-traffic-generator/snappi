package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicTime *****
type mkaBasicTime struct {
	validation
	obj          *otg.MkaBasicTime
	marshaller   marshalMkaBasicTime
	unMarshaller unMarshalMkaBasicTime
	utcHolder    MkaBasicTimeUtc
}

func NewMkaBasicTime() MkaBasicTime {
	obj := mkaBasicTime{obj: &otg.MkaBasicTime{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicTime) msg() *otg.MkaBasicTime {
	return obj.obj
}

func (obj *mkaBasicTime) setMsg(msg *otg.MkaBasicTime) MkaBasicTime {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicTime struct {
	obj *mkaBasicTime
}

type marshalMkaBasicTime interface {
	// ToProto marshals MkaBasicTime to protobuf object *otg.MkaBasicTime
	ToProto() (*otg.MkaBasicTime, error)
	// ToPbText marshals MkaBasicTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicTime to JSON text
	ToJson() (string, error)
}

type unMarshalmkaBasicTime struct {
	obj *mkaBasicTime
}

type unMarshalMkaBasicTime interface {
	// FromProto unmarshals MkaBasicTime from protobuf object *otg.MkaBasicTime
	FromProto(msg *otg.MkaBasicTime) (MkaBasicTime, error)
	// FromPbText unmarshals MkaBasicTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicTime from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicTime) Marshal() marshalMkaBasicTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicTime) Unmarshal() unMarshalMkaBasicTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicTime) ToProto() (*otg.MkaBasicTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicTime) FromProto(msg *otg.MkaBasicTime) (MkaBasicTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicTime) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicTime) FromPbText(value string) error {
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

func (m *marshalmkaBasicTime) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicTime) FromYaml(value string) error {
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

func (m *marshalmkaBasicTime) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicTime) FromJson(value string) error {
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

func (obj *mkaBasicTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicTime) Clone() (MkaBasicTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicTime()
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

func (obj *mkaBasicTime) setNil() {
	obj.utcHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaBasicTime is absolute time in a timezone.
type MkaBasicTime interface {
	Validation
	// msg marshals MkaBasicTime to protobuf object *otg.MkaBasicTime
	// and doesn't set defaults
	msg() *otg.MkaBasicTime
	// setMsg unmarshals MkaBasicTime from protobuf object *otg.MkaBasicTime
	// and doesn't set defaults
	setMsg(*otg.MkaBasicTime) MkaBasicTime
	// provides marshal interface
	Marshal() marshalMkaBasicTime
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicTime
	// validate validates MkaBasicTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MkaBasicTimeChoiceEnum, set in MkaBasicTime
	Choice() MkaBasicTimeChoiceEnum
	// setChoice assigns MkaBasicTimeChoiceEnum provided by user to MkaBasicTime
	setChoice(value MkaBasicTimeChoiceEnum) MkaBasicTime
	// HasChoice checks if Choice has been set in MkaBasicTime
	HasChoice() bool
	// Utc returns MkaBasicTimeUtc, set in MkaBasicTime.
	// MkaBasicTimeUtc is coordinated Universal Time(UTC).
	Utc() MkaBasicTimeUtc
	// SetUtc assigns MkaBasicTimeUtc provided by user to MkaBasicTime.
	// MkaBasicTimeUtc is coordinated Universal Time(UTC).
	SetUtc(value MkaBasicTimeUtc) MkaBasicTime
	// HasUtc checks if Utc has been set in MkaBasicTime
	HasUtc() bool
	setNil()
}

type MkaBasicTimeChoiceEnum string

// Enum of Choice on MkaBasicTime
var MkaBasicTimeChoice = struct {
	UTC MkaBasicTimeChoiceEnum
}{
	UTC: MkaBasicTimeChoiceEnum("utc"),
}

func (obj *mkaBasicTime) Choice() MkaBasicTimeChoiceEnum {
	return MkaBasicTimeChoiceEnum(obj.obj.Choice.Enum().String())
}

// Timezone choice. Currently only Coordinated Universal Time(UTC) is supported.
// Choice returns a string
func (obj *mkaBasicTime) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *mkaBasicTime) setChoice(value MkaBasicTimeChoiceEnum) MkaBasicTime {
	intValue, ok := otg.MkaBasicTime_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicTimeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasicTime_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Utc = nil
	obj.utcHolder = nil

	if value == MkaBasicTimeChoice.UTC {
		obj.obj.Utc = NewMkaBasicTimeUtc().msg()
	}

	return obj
}

// Coordinated Universal Time(UTC) time.
// Utc returns a MkaBasicTimeUtc
func (obj *mkaBasicTime) Utc() MkaBasicTimeUtc {
	if obj.obj.Utc == nil {
		obj.setChoice(MkaBasicTimeChoice.UTC)
	}
	if obj.utcHolder == nil {
		obj.utcHolder = &mkaBasicTimeUtc{obj: obj.obj.Utc}
	}
	return obj.utcHolder
}

// Coordinated Universal Time(UTC) time.
// Utc returns a MkaBasicTimeUtc
func (obj *mkaBasicTime) HasUtc() bool {
	return obj.obj.Utc != nil
}

// Coordinated Universal Time(UTC) time.
// SetUtc sets the MkaBasicTimeUtc value in the MkaBasicTime object
func (obj *mkaBasicTime) SetUtc(value MkaBasicTimeUtc) MkaBasicTime {
	obj.setChoice(MkaBasicTimeChoice.UTC)
	obj.utcHolder = nil
	obj.obj.Utc = value.msg()

	return obj
}

func (obj *mkaBasicTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Utc != nil {

		obj.Utc().validateObj(vObj, set_default)
	}

}

func (obj *mkaBasicTime) setDefault() {
	var choices_set int = 0
	var choice MkaBasicTimeChoiceEnum

	if obj.obj.Utc != nil {
		choices_set += 1
		choice = MkaBasicTimeChoice.UTC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MkaBasicTimeChoice.UTC)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MkaBasicTime")
			}
		} else {
			intVal := otg.MkaBasicTime_Choice_Enum_value[string(choice)]
			enumValue := otg.MkaBasicTime_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
