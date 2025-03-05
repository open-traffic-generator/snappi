package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2InterBatch *****
type roCEv2InterBatch struct {
	validation
	obj          *otg.RoCEv2InterBatch
	marshaller   marshalRoCEv2InterBatch
	unMarshaller unMarshalRoCEv2InterBatch
}

func NewRoCEv2InterBatch() RoCEv2InterBatch {
	obj := roCEv2InterBatch{obj: &otg.RoCEv2InterBatch{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2InterBatch) msg() *otg.RoCEv2InterBatch {
	return obj.obj
}

func (obj *roCEv2InterBatch) setMsg(msg *otg.RoCEv2InterBatch) RoCEv2InterBatch {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2InterBatch struct {
	obj *roCEv2InterBatch
}

type marshalRoCEv2InterBatch interface {
	// ToProto marshals RoCEv2InterBatch to protobuf object *otg.RoCEv2InterBatch
	ToProto() (*otg.RoCEv2InterBatch, error)
	// ToPbText marshals RoCEv2InterBatch to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2InterBatch to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2InterBatch to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2InterBatch struct {
	obj *roCEv2InterBatch
}

type unMarshalRoCEv2InterBatch interface {
	// FromProto unmarshals RoCEv2InterBatch from protobuf object *otg.RoCEv2InterBatch
	FromProto(msg *otg.RoCEv2InterBatch) (RoCEv2InterBatch, error)
	// FromPbText unmarshals RoCEv2InterBatch from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2InterBatch from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2InterBatch from JSON text
	FromJson(value string) error
}

func (obj *roCEv2InterBatch) Marshal() marshalRoCEv2InterBatch {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2InterBatch{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2InterBatch) Unmarshal() unMarshalRoCEv2InterBatch {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2InterBatch{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2InterBatch) ToProto() (*otg.RoCEv2InterBatch, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2InterBatch) FromProto(msg *otg.RoCEv2InterBatch) (RoCEv2InterBatch, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2InterBatch) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2InterBatch) FromPbText(value string) error {
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

func (m *marshalroCEv2InterBatch) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2InterBatch) FromYaml(value string) error {
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

func (m *marshalroCEv2InterBatch) ToJson() (string, error) {
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

func (m *unMarshalroCEv2InterBatch) FromJson(value string) error {
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

func (obj *roCEv2InterBatch) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2InterBatch) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2InterBatch) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2InterBatch) Clone() (RoCEv2InterBatch, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2InterBatch()
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

// RoCEv2InterBatch is roCEv2 inter batch settigns applicable per port
type RoCEv2InterBatch interface {
	Validation
	// msg marshals RoCEv2InterBatch to protobuf object *otg.RoCEv2InterBatch
	// and doesn't set defaults
	msg() *otg.RoCEv2InterBatch
	// setMsg unmarshals RoCEv2InterBatch from protobuf object *otg.RoCEv2InterBatch
	// and doesn't set defaults
	setMsg(*otg.RoCEv2InterBatch) RoCEv2InterBatch
	// provides marshal interface
	Marshal() marshalRoCEv2InterBatch
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2InterBatch
	// validate validates RoCEv2InterBatch
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2InterBatch, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// InterBatchPeriodUnit returns RoCEv2InterBatchInterBatchPeriodUnitEnum, set in RoCEv2InterBatch
	InterBatchPeriodUnit() RoCEv2InterBatchInterBatchPeriodUnitEnum
	// SetInterBatchPeriodUnit assigns RoCEv2InterBatchInterBatchPeriodUnitEnum provided by user to RoCEv2InterBatch
	SetInterBatchPeriodUnit(value RoCEv2InterBatchInterBatchPeriodUnitEnum) RoCEv2InterBatch
	// HasInterBatchPeriodUnit checks if InterBatchPeriodUnit has been set in RoCEv2InterBatch
	HasInterBatchPeriodUnit() bool
	// InterBatchPeriod returns uint64, set in RoCEv2InterBatch.
	InterBatchPeriod() uint64
	// SetInterBatchPeriod assigns uint64 provided by user to RoCEv2InterBatch
	SetInterBatchPeriod(value uint64) RoCEv2InterBatch
	// HasInterBatchPeriod checks if InterBatchPeriod has been set in RoCEv2InterBatch
	HasInterBatchPeriod() bool
}

type RoCEv2InterBatchInterBatchPeriodUnitEnum string

// Enum of InterBatchPeriodUnit on RoCEv2InterBatch
var RoCEv2InterBatchInterBatchPeriodUnit = struct {
	NANOSECONDS  RoCEv2InterBatchInterBatchPeriodUnitEnum
	MICROSENDONS RoCEv2InterBatchInterBatchPeriodUnitEnum
	MILLISECONDS RoCEv2InterBatchInterBatchPeriodUnitEnum
	SECONDS      RoCEv2InterBatchInterBatchPeriodUnitEnum
}{
	NANOSECONDS:  RoCEv2InterBatchInterBatchPeriodUnitEnum("nanoseconds"),
	MICROSENDONS: RoCEv2InterBatchInterBatchPeriodUnitEnum("microsendons"),
	MILLISECONDS: RoCEv2InterBatchInterBatchPeriodUnitEnum("milliseconds"),
	SECONDS:      RoCEv2InterBatchInterBatchPeriodUnitEnum("seconds"),
}

func (obj *roCEv2InterBatch) InterBatchPeriodUnit() RoCEv2InterBatchInterBatchPeriodUnitEnum {
	return RoCEv2InterBatchInterBatchPeriodUnitEnum(obj.obj.InterBatchPeriodUnit.Enum().String())
}

// inter batch period units.
// InterBatchPeriodUnit returns a string
func (obj *roCEv2InterBatch) HasInterBatchPeriodUnit() bool {
	return obj.obj.InterBatchPeriodUnit != nil
}

func (obj *roCEv2InterBatch) SetInterBatchPeriodUnit(value RoCEv2InterBatchInterBatchPeriodUnitEnum) RoCEv2InterBatch {
	intValue, ok := otg.RoCEv2InterBatch_InterBatchPeriodUnit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RoCEv2InterBatchInterBatchPeriodUnitEnum", string(value)))
		return obj
	}
	enumValue := otg.RoCEv2InterBatch_InterBatchPeriodUnit_Enum(intValue)
	obj.obj.InterBatchPeriodUnit = &enumValue

	return obj
}

// Inter batch Period.
// InterBatchPeriod returns a uint64
func (obj *roCEv2InterBatch) InterBatchPeriod() uint64 {

	return *obj.obj.InterBatchPeriod

}

// Inter batch Period.
// InterBatchPeriod returns a uint64
func (obj *roCEv2InterBatch) HasInterBatchPeriod() bool {
	return obj.obj.InterBatchPeriod != nil
}

// Inter batch Period.
// SetInterBatchPeriod sets the uint64 value in the RoCEv2InterBatch object
func (obj *roCEv2InterBatch) SetInterBatchPeriod(value uint64) RoCEv2InterBatch {

	obj.obj.InterBatchPeriod = &value
	return obj
}

func (obj *roCEv2InterBatch) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *roCEv2InterBatch) setDefault() {

}
