package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2TargetLineRate *****
type rocev2TargetLineRate struct {
	validation
	obj          *otg.Rocev2TargetLineRate
	marshaller   marshalRocev2TargetLineRate
	unMarshaller unMarshalRocev2TargetLineRate
	flowsHolder  Rocev2TargetLineRateRocev2FlowIter
}

func NewRocev2TargetLineRate() Rocev2TargetLineRate {
	obj := rocev2TargetLineRate{obj: &otg.Rocev2TargetLineRate{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2TargetLineRate) msg() *otg.Rocev2TargetLineRate {
	return obj.obj
}

func (obj *rocev2TargetLineRate) setMsg(msg *otg.Rocev2TargetLineRate) Rocev2TargetLineRate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2TargetLineRate struct {
	obj *rocev2TargetLineRate
}

type marshalRocev2TargetLineRate interface {
	// ToProto marshals Rocev2TargetLineRate to protobuf object *otg.Rocev2TargetLineRate
	ToProto() (*otg.Rocev2TargetLineRate, error)
	// ToPbText marshals Rocev2TargetLineRate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2TargetLineRate to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2TargetLineRate to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2TargetLineRate to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2TargetLineRate struct {
	obj *rocev2TargetLineRate
}

type unMarshalRocev2TargetLineRate interface {
	// FromProto unmarshals Rocev2TargetLineRate from protobuf object *otg.Rocev2TargetLineRate
	FromProto(msg *otg.Rocev2TargetLineRate) (Rocev2TargetLineRate, error)
	// FromPbText unmarshals Rocev2TargetLineRate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2TargetLineRate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2TargetLineRate from JSON text
	FromJson(value string) error
}

func (obj *rocev2TargetLineRate) Marshal() marshalRocev2TargetLineRate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2TargetLineRate{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2TargetLineRate) Unmarshal() unMarshalRocev2TargetLineRate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2TargetLineRate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2TargetLineRate) ToProto() (*otg.Rocev2TargetLineRate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2TargetLineRate) FromProto(msg *otg.Rocev2TargetLineRate) (Rocev2TargetLineRate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2TargetLineRate) ToPbText() (string, error) {
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

func (m *unMarshalrocev2TargetLineRate) FromPbText(value string) error {
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

func (m *marshalrocev2TargetLineRate) ToYaml() (string, error) {
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

func (m *unMarshalrocev2TargetLineRate) FromYaml(value string) error {
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

func (m *marshalrocev2TargetLineRate) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2TargetLineRate) ToJson() (string, error) {
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

func (m *unMarshalrocev2TargetLineRate) FromJson(value string) error {
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

func (obj *rocev2TargetLineRate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2TargetLineRate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2TargetLineRate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2TargetLineRate) Clone() (Rocev2TargetLineRate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2TargetLineRate()
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

func (obj *rocev2TargetLineRate) setNil() {
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2TargetLineRate is configure target line rate of traffic rate on this port as percentage of link speed.
type Rocev2TargetLineRate interface {
	Validation
	// msg marshals Rocev2TargetLineRate to protobuf object *otg.Rocev2TargetLineRate
	// and doesn't set defaults
	msg() *otg.Rocev2TargetLineRate
	// setMsg unmarshals Rocev2TargetLineRate from protobuf object *otg.Rocev2TargetLineRate
	// and doesn't set defaults
	setMsg(*otg.Rocev2TargetLineRate) Rocev2TargetLineRate
	// provides marshal interface
	Marshal() marshalRocev2TargetLineRate
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2TargetLineRate
	// validate validates Rocev2TargetLineRate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2TargetLineRate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint64, set in Rocev2TargetLineRate.
	Value() uint64
	// SetValue assigns uint64 provided by user to Rocev2TargetLineRate
	SetValue(value uint64) Rocev2TargetLineRate
	// HasValue checks if Value has been set in Rocev2TargetLineRate
	HasValue() bool
	// Flows returns Rocev2TargetLineRateRocev2FlowIterIter, set in Rocev2TargetLineRate
	Flows() Rocev2TargetLineRateRocev2FlowIter
	setNil()
}

// Target Line Rate as percentage of max line rate.
// Value returns a uint64
func (obj *rocev2TargetLineRate) Value() uint64 {

	return *obj.obj.Value

}

// Target Line Rate as percentage of max line rate.
// Value returns a uint64
func (obj *rocev2TargetLineRate) HasValue() bool {
	return obj.obj.Value != nil
}

// Target Line Rate as percentage of max line rate.
// SetValue sets the uint64 value in the Rocev2TargetLineRate object
func (obj *rocev2TargetLineRate) SetValue(value uint64) Rocev2TargetLineRate {

	obj.obj.Value = &value
	return obj
}

// description is TBD
// Flows returns a []Rocev2Flow
func (obj *rocev2TargetLineRate) Flows() Rocev2TargetLineRateRocev2FlowIter {
	if len(obj.obj.Flows) == 0 {
		obj.obj.Flows = []*otg.Rocev2Flow{}
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newRocev2TargetLineRateRocev2FlowIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type rocev2TargetLineRateRocev2FlowIter struct {
	obj             *rocev2TargetLineRate
	rocev2FlowSlice []Rocev2Flow
	fieldPtr        *[]*otg.Rocev2Flow
}

func newRocev2TargetLineRateRocev2FlowIter(ptr *[]*otg.Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter {
	return &rocev2TargetLineRateRocev2FlowIter{fieldPtr: ptr}
}

type Rocev2TargetLineRateRocev2FlowIter interface {
	setMsg(*rocev2TargetLineRate) Rocev2TargetLineRateRocev2FlowIter
	Items() []Rocev2Flow
	Add() Rocev2Flow
	Append(items ...Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter
	Set(index int, newObj Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter
	Clear() Rocev2TargetLineRateRocev2FlowIter
	clearHolderSlice() Rocev2TargetLineRateRocev2FlowIter
	appendHolderSlice(item Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter
}

func (obj *rocev2TargetLineRateRocev2FlowIter) setMsg(msg *rocev2TargetLineRate) Rocev2TargetLineRateRocev2FlowIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rocev2Flow{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rocev2TargetLineRateRocev2FlowIter) Items() []Rocev2Flow {
	return obj.rocev2FlowSlice
}

func (obj *rocev2TargetLineRateRocev2FlowIter) Add() Rocev2Flow {
	newObj := &otg.Rocev2Flow{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rocev2Flow{obj: newObj}
	newLibObj.setDefault()
	obj.rocev2FlowSlice = append(obj.rocev2FlowSlice, newLibObj)
	return newLibObj
}

func (obj *rocev2TargetLineRateRocev2FlowIter) Append(items ...Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rocev2FlowSlice = append(obj.rocev2FlowSlice, item)
	}
	return obj
}

func (obj *rocev2TargetLineRateRocev2FlowIter) Set(index int, newObj Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rocev2FlowSlice[index] = newObj
	return obj
}
func (obj *rocev2TargetLineRateRocev2FlowIter) Clear() Rocev2TargetLineRateRocev2FlowIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Rocev2Flow{}
		obj.rocev2FlowSlice = []Rocev2Flow{}
	}
	return obj
}
func (obj *rocev2TargetLineRateRocev2FlowIter) clearHolderSlice() Rocev2TargetLineRateRocev2FlowIter {
	if len(obj.rocev2FlowSlice) > 0 {
		obj.rocev2FlowSlice = []Rocev2Flow{}
	}
	return obj
}
func (obj *rocev2TargetLineRateRocev2FlowIter) appendHolderSlice(item Rocev2Flow) Rocev2TargetLineRateRocev2FlowIter {
	obj.rocev2FlowSlice = append(obj.rocev2FlowSlice, item)
	return obj
}

func (obj *rocev2TargetLineRate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Rocev2TargetLineRate.Value <= 100 but Got %d", *obj.obj.Value))
		}

	}

	if len(obj.obj.Flows) != 0 {

		if set_default {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				obj.Flows().appendHolderSlice(&rocev2Flow{obj: item})
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rocev2TargetLineRate) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(98)
	}

}
