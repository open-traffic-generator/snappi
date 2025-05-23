package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicPskChainStartTime *****
type mkaBasicPskChainStartTime struct {
	validation
	obj          *otg.MkaBasicPskChainStartTime
	marshaller   marshalMkaBasicPskChainStartTime
	unMarshaller unMarshalMkaBasicPskChainStartTime
	utcHolder    MkaBasicTimeUtc
}

func NewMkaBasicPskChainStartTime() MkaBasicPskChainStartTime {
	obj := mkaBasicPskChainStartTime{obj: &otg.MkaBasicPskChainStartTime{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicPskChainStartTime) msg() *otg.MkaBasicPskChainStartTime {
	return obj.obj
}

func (obj *mkaBasicPskChainStartTime) setMsg(msg *otg.MkaBasicPskChainStartTime) MkaBasicPskChainStartTime {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicPskChainStartTime struct {
	obj *mkaBasicPskChainStartTime
}

type marshalMkaBasicPskChainStartTime interface {
	// ToProto marshals MkaBasicPskChainStartTime to protobuf object *otg.MkaBasicPskChainStartTime
	ToProto() (*otg.MkaBasicPskChainStartTime, error)
	// ToPbText marshals MkaBasicPskChainStartTime to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicPskChainStartTime to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicPskChainStartTime to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MkaBasicPskChainStartTime to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmkaBasicPskChainStartTime struct {
	obj *mkaBasicPskChainStartTime
}

type unMarshalMkaBasicPskChainStartTime interface {
	// FromProto unmarshals MkaBasicPskChainStartTime from protobuf object *otg.MkaBasicPskChainStartTime
	FromProto(msg *otg.MkaBasicPskChainStartTime) (MkaBasicPskChainStartTime, error)
	// FromPbText unmarshals MkaBasicPskChainStartTime from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicPskChainStartTime from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicPskChainStartTime from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicPskChainStartTime) Marshal() marshalMkaBasicPskChainStartTime {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicPskChainStartTime{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicPskChainStartTime) Unmarshal() unMarshalMkaBasicPskChainStartTime {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicPskChainStartTime{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicPskChainStartTime) ToProto() (*otg.MkaBasicPskChainStartTime, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicPskChainStartTime) FromProto(msg *otg.MkaBasicPskChainStartTime) (MkaBasicPskChainStartTime, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicPskChainStartTime) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicPskChainStartTime) FromPbText(value string) error {
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

func (m *marshalmkaBasicPskChainStartTime) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicPskChainStartTime) FromYaml(value string) error {
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

func (m *marshalmkaBasicPskChainStartTime) ToJsonRaw() (string, error) {
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

func (m *marshalmkaBasicPskChainStartTime) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicPskChainStartTime) FromJson(value string) error {
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

func (obj *mkaBasicPskChainStartTime) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicPskChainStartTime) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicPskChainStartTime) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicPskChainStartTime) Clone() (MkaBasicPskChainStartTime, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicPskChainStartTime()
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

func (obj *mkaBasicPskChainStartTime) setNil() {
	obj.utcHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MkaBasicPskChainStartTime is pre-shared key(PSK) chain start time in UTC time format DD-MM-YYYY HH:MM:SS. If this time is set, the key start time will be relative to this value. Otherwise if this value is not set, key start time will be relative to test start time.
type MkaBasicPskChainStartTime interface {
	Validation
	// msg marshals MkaBasicPskChainStartTime to protobuf object *otg.MkaBasicPskChainStartTime
	// and doesn't set defaults
	msg() *otg.MkaBasicPskChainStartTime
	// setMsg unmarshals MkaBasicPskChainStartTime from protobuf object *otg.MkaBasicPskChainStartTime
	// and doesn't set defaults
	setMsg(*otg.MkaBasicPskChainStartTime) MkaBasicPskChainStartTime
	// provides marshal interface
	Marshal() marshalMkaBasicPskChainStartTime
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicPskChainStartTime
	// validate validates MkaBasicPskChainStartTime
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicPskChainStartTime, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MkaBasicPskChainStartTimeChoiceEnum, set in MkaBasicPskChainStartTime
	Choice() MkaBasicPskChainStartTimeChoiceEnum
	// setChoice assigns MkaBasicPskChainStartTimeChoiceEnum provided by user to MkaBasicPskChainStartTime
	setChoice(value MkaBasicPskChainStartTimeChoiceEnum) MkaBasicPskChainStartTime
	// HasChoice checks if Choice has been set in MkaBasicPskChainStartTime
	HasChoice() bool
	// Utc returns MkaBasicTimeUtc, set in MkaBasicPskChainStartTime.
	// MkaBasicTimeUtc is coordinated Universal Time(UTC).
	Utc() MkaBasicTimeUtc
	// SetUtc assigns MkaBasicTimeUtc provided by user to MkaBasicPskChainStartTime.
	// MkaBasicTimeUtc is coordinated Universal Time(UTC).
	SetUtc(value MkaBasicTimeUtc) MkaBasicPskChainStartTime
	// HasUtc checks if Utc has been set in MkaBasicPskChainStartTime
	HasUtc() bool
	setNil()
}

type MkaBasicPskChainStartTimeChoiceEnum string

// Enum of Choice on MkaBasicPskChainStartTime
var MkaBasicPskChainStartTimeChoice = struct {
	UTC MkaBasicPskChainStartTimeChoiceEnum
}{
	UTC: MkaBasicPskChainStartTimeChoiceEnum("utc"),
}

func (obj *mkaBasicPskChainStartTime) Choice() MkaBasicPskChainStartTimeChoiceEnum {
	return MkaBasicPskChainStartTimeChoiceEnum(obj.obj.Choice.Enum().String())
}

// Timezone choice. Currently only Coordinated Universal Time(UTC) is supported.
// Choice returns a string
func (obj *mkaBasicPskChainStartTime) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *mkaBasicPskChainStartTime) setChoice(value MkaBasicPskChainStartTimeChoiceEnum) MkaBasicPskChainStartTime {
	intValue, ok := otg.MkaBasicPskChainStartTime_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MkaBasicPskChainStartTimeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.MkaBasicPskChainStartTime_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Utc = nil
	obj.utcHolder = nil

	if value == MkaBasicPskChainStartTimeChoice.UTC {
		obj.obj.Utc = NewMkaBasicTimeUtc().msg()
	}

	return obj
}

// Coordinated Universal Time(UTC) time.
// Utc returns a MkaBasicTimeUtc
func (obj *mkaBasicPskChainStartTime) Utc() MkaBasicTimeUtc {
	if obj.obj.Utc == nil {
		obj.setChoice(MkaBasicPskChainStartTimeChoice.UTC)
	}
	if obj.utcHolder == nil {
		obj.utcHolder = &mkaBasicTimeUtc{obj: obj.obj.Utc}
	}
	return obj.utcHolder
}

// Coordinated Universal Time(UTC) time.
// Utc returns a MkaBasicTimeUtc
func (obj *mkaBasicPskChainStartTime) HasUtc() bool {
	return obj.obj.Utc != nil
}

// Coordinated Universal Time(UTC) time.
// SetUtc sets the MkaBasicTimeUtc value in the MkaBasicPskChainStartTime object
func (obj *mkaBasicPskChainStartTime) SetUtc(value MkaBasicTimeUtc) MkaBasicPskChainStartTime {
	obj.setChoice(MkaBasicPskChainStartTimeChoice.UTC)
	obj.utcHolder = nil
	obj.obj.Utc = value.msg()

	return obj
}

func (obj *mkaBasicPskChainStartTime) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Utc != nil {

		obj.Utc().validateObj(vObj, set_default)
	}

}

func (obj *mkaBasicPskChainStartTime) setDefault() {
	var choices_set int = 0
	var choice MkaBasicPskChainStartTimeChoiceEnum

	if obj.obj.Utc != nil {
		choices_set += 1
		choice = MkaBasicPskChainStartTimeChoice.UTC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MkaBasicPskChainStartTimeChoice.UTC)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MkaBasicPskChainStartTime")
			}
		} else {
			intVal := otg.MkaBasicPskChainStartTime_Choice_Enum_value[string(choice)]
			enumValue := otg.MkaBasicPskChainStartTime_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
