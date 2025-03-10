package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MkaBasicTimeUtc *****
type mkaBasicTimeUtc struct {
	validation
	obj          *otg.MkaBasicTimeUtc
	marshaller   marshalMkaBasicTimeUtc
	unMarshaller unMarshalMkaBasicTimeUtc
}

func NewMkaBasicTimeUtc() MkaBasicTimeUtc {
	obj := mkaBasicTimeUtc{obj: &otg.MkaBasicTimeUtc{}}
	obj.setDefault()
	return &obj
}

func (obj *mkaBasicTimeUtc) msg() *otg.MkaBasicTimeUtc {
	return obj.obj
}

func (obj *mkaBasicTimeUtc) setMsg(msg *otg.MkaBasicTimeUtc) MkaBasicTimeUtc {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmkaBasicTimeUtc struct {
	obj *mkaBasicTimeUtc
}

type marshalMkaBasicTimeUtc interface {
	// ToProto marshals MkaBasicTimeUtc to protobuf object *otg.MkaBasicTimeUtc
	ToProto() (*otg.MkaBasicTimeUtc, error)
	// ToPbText marshals MkaBasicTimeUtc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MkaBasicTimeUtc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MkaBasicTimeUtc to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals MkaBasicTimeUtc to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalmkaBasicTimeUtc struct {
	obj *mkaBasicTimeUtc
}

type unMarshalMkaBasicTimeUtc interface {
	// FromProto unmarshals MkaBasicTimeUtc from protobuf object *otg.MkaBasicTimeUtc
	FromProto(msg *otg.MkaBasicTimeUtc) (MkaBasicTimeUtc, error)
	// FromPbText unmarshals MkaBasicTimeUtc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MkaBasicTimeUtc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MkaBasicTimeUtc from JSON text
	FromJson(value string) error
}

func (obj *mkaBasicTimeUtc) Marshal() marshalMkaBasicTimeUtc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmkaBasicTimeUtc{obj: obj}
	}
	return obj.marshaller
}

func (obj *mkaBasicTimeUtc) Unmarshal() unMarshalMkaBasicTimeUtc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmkaBasicTimeUtc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmkaBasicTimeUtc) ToProto() (*otg.MkaBasicTimeUtc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmkaBasicTimeUtc) FromProto(msg *otg.MkaBasicTimeUtc) (MkaBasicTimeUtc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmkaBasicTimeUtc) ToPbText() (string, error) {
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

func (m *unMarshalmkaBasicTimeUtc) FromPbText(value string) error {
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

func (m *marshalmkaBasicTimeUtc) ToYaml() (string, error) {
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

func (m *unMarshalmkaBasicTimeUtc) FromYaml(value string) error {
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

func (m *marshalmkaBasicTimeUtc) ToJsonRaw() (string, error) {
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

func (m *marshalmkaBasicTimeUtc) ToJson() (string, error) {
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

func (m *unMarshalmkaBasicTimeUtc) FromJson(value string) error {
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

func (obj *mkaBasicTimeUtc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mkaBasicTimeUtc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mkaBasicTimeUtc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mkaBasicTimeUtc) Clone() (MkaBasicTimeUtc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMkaBasicTimeUtc()
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

// MkaBasicTimeUtc is coordinated Universal Time(UTC).
type MkaBasicTimeUtc interface {
	Validation
	// msg marshals MkaBasicTimeUtc to protobuf object *otg.MkaBasicTimeUtc
	// and doesn't set defaults
	msg() *otg.MkaBasicTimeUtc
	// setMsg unmarshals MkaBasicTimeUtc from protobuf object *otg.MkaBasicTimeUtc
	// and doesn't set defaults
	setMsg(*otg.MkaBasicTimeUtc) MkaBasicTimeUtc
	// provides marshal interface
	Marshal() marshalMkaBasicTimeUtc
	// provides unmarshal interface
	Unmarshal() unMarshalMkaBasicTimeUtc
	// validate validates MkaBasicTimeUtc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MkaBasicTimeUtc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Day returns uint32, set in MkaBasicTimeUtc.
	Day() uint32
	// SetDay assigns uint32 provided by user to MkaBasicTimeUtc
	SetDay(value uint32) MkaBasicTimeUtc
	// HasDay checks if Day has been set in MkaBasicTimeUtc
	HasDay() bool
	// Month returns uint32, set in MkaBasicTimeUtc.
	Month() uint32
	// SetMonth assigns uint32 provided by user to MkaBasicTimeUtc
	SetMonth(value uint32) MkaBasicTimeUtc
	// HasMonth checks if Month has been set in MkaBasicTimeUtc
	HasMonth() bool
	// Year returns uint32, set in MkaBasicTimeUtc.
	Year() uint32
	// SetYear assigns uint32 provided by user to MkaBasicTimeUtc
	SetYear(value uint32) MkaBasicTimeUtc
	// HasYear checks if Year has been set in MkaBasicTimeUtc
	HasYear() bool
	// Hour returns uint32, set in MkaBasicTimeUtc.
	Hour() uint32
	// SetHour assigns uint32 provided by user to MkaBasicTimeUtc
	SetHour(value uint32) MkaBasicTimeUtc
	// HasHour checks if Hour has been set in MkaBasicTimeUtc
	HasHour() bool
	// Minute returns uint32, set in MkaBasicTimeUtc.
	Minute() uint32
	// SetMinute assigns uint32 provided by user to MkaBasicTimeUtc
	SetMinute(value uint32) MkaBasicTimeUtc
	// HasMinute checks if Minute has been set in MkaBasicTimeUtc
	HasMinute() bool
	// Second returns uint32, set in MkaBasicTimeUtc.
	Second() uint32
	// SetSecond assigns uint32 provided by user to MkaBasicTimeUtc
	SetSecond(value uint32) MkaBasicTimeUtc
	// HasSecond checks if Second has been set in MkaBasicTimeUtc
	HasSecond() bool
}

// Day of the month in DD format.
// Day returns a uint32
func (obj *mkaBasicTimeUtc) Day() uint32 {

	return *obj.obj.Day

}

// Day of the month in DD format.
// Day returns a uint32
func (obj *mkaBasicTimeUtc) HasDay() bool {
	return obj.obj.Day != nil
}

// Day of the month in DD format.
// SetDay sets the uint32 value in the MkaBasicTimeUtc object
func (obj *mkaBasicTimeUtc) SetDay(value uint32) MkaBasicTimeUtc {

	obj.obj.Day = &value
	return obj
}

// Month of the year in MM format.
// Month returns a uint32
func (obj *mkaBasicTimeUtc) Month() uint32 {

	return *obj.obj.Month

}

// Month of the year in MM format.
// Month returns a uint32
func (obj *mkaBasicTimeUtc) HasMonth() bool {
	return obj.obj.Month != nil
}

// Month of the year in MM format.
// SetMonth sets the uint32 value in the MkaBasicTimeUtc object
func (obj *mkaBasicTimeUtc) SetMonth(value uint32) MkaBasicTimeUtc {

	obj.obj.Month = &value
	return obj
}

// Year from the start of common era(CE) in YYYY format.
// Year returns a uint32
func (obj *mkaBasicTimeUtc) Year() uint32 {

	return *obj.obj.Year

}

// Year from the start of common era(CE) in YYYY format.
// Year returns a uint32
func (obj *mkaBasicTimeUtc) HasYear() bool {
	return obj.obj.Year != nil
}

// Year from the start of common era(CE) in YYYY format.
// SetYear sets the uint32 value in the MkaBasicTimeUtc object
func (obj *mkaBasicTimeUtc) SetYear(value uint32) MkaBasicTimeUtc {

	obj.obj.Year = &value
	return obj
}

// Hour of the day in HH format.
// Hour returns a uint32
func (obj *mkaBasicTimeUtc) Hour() uint32 {

	return *obj.obj.Hour

}

// Hour of the day in HH format.
// Hour returns a uint32
func (obj *mkaBasicTimeUtc) HasHour() bool {
	return obj.obj.Hour != nil
}

// Hour of the day in HH format.
// SetHour sets the uint32 value in the MkaBasicTimeUtc object
func (obj *mkaBasicTimeUtc) SetHour(value uint32) MkaBasicTimeUtc {

	obj.obj.Hour = &value
	return obj
}

// Minute of the hour in MM format.
// Minute returns a uint32
func (obj *mkaBasicTimeUtc) Minute() uint32 {

	return *obj.obj.Minute

}

// Minute of the hour in MM format.
// Minute returns a uint32
func (obj *mkaBasicTimeUtc) HasMinute() bool {
	return obj.obj.Minute != nil
}

// Minute of the hour in MM format.
// SetMinute sets the uint32 value in the MkaBasicTimeUtc object
func (obj *mkaBasicTimeUtc) SetMinute(value uint32) MkaBasicTimeUtc {

	obj.obj.Minute = &value
	return obj
}

// Second of the minute in SS format.
// Second returns a uint32
func (obj *mkaBasicTimeUtc) Second() uint32 {

	return *obj.obj.Second

}

// Second of the minute in SS format.
// Second returns a uint32
func (obj *mkaBasicTimeUtc) HasSecond() bool {
	return obj.obj.Second != nil
}

// Second of the minute in SS format.
// SetSecond sets the uint32 value in the MkaBasicTimeUtc object
func (obj *mkaBasicTimeUtc) SetSecond(value uint32) MkaBasicTimeUtc {

	obj.obj.Second = &value
	return obj
}

func (obj *mkaBasicTimeUtc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Day != nil {

		if *obj.obj.Day < 1 || *obj.obj.Day > 31 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MkaBasicTimeUtc.Day <= 31 but Got %d", *obj.obj.Day))
		}

	}

	if obj.obj.Month != nil {

		if *obj.obj.Month < 1 || *obj.obj.Month > 12 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MkaBasicTimeUtc.Month <= 12 but Got %d", *obj.obj.Month))
		}

	}

	if obj.obj.Year != nil {

		if *obj.obj.Year > 9999 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicTimeUtc.Year <= 9999 but Got %d", *obj.obj.Year))
		}

	}

	if obj.obj.Hour != nil {

		if *obj.obj.Hour > 23 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicTimeUtc.Hour <= 23 but Got %d", *obj.obj.Hour))
		}

	}

	if obj.obj.Minute != nil {

		if *obj.obj.Minute > 59 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicTimeUtc.Minute <= 59 but Got %d", *obj.obj.Minute))
		}

	}

	if obj.obj.Second != nil {

		if *obj.obj.Second > 59 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MkaBasicTimeUtc.Second <= 59 but Got %d", *obj.obj.Second))
		}

	}

}

func (obj *mkaBasicTimeUtc) setDefault() {

}
