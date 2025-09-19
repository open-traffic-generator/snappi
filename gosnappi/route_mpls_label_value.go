package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RouteMplsLabelValue *****
type routeMplsLabelValue struct {
	validation
	obj          *otg.RouteMplsLabelValue
	marshaller   marshalRouteMplsLabelValue
	unMarshaller unMarshalRouteMplsLabelValue
}

func NewRouteMplsLabelValue() RouteMplsLabelValue {
	obj := routeMplsLabelValue{obj: &otg.RouteMplsLabelValue{}}
	obj.setDefault()
	return &obj
}

func (obj *routeMplsLabelValue) msg() *otg.RouteMplsLabelValue {
	return obj.obj
}

func (obj *routeMplsLabelValue) setMsg(msg *otg.RouteMplsLabelValue) RouteMplsLabelValue {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrouteMplsLabelValue struct {
	obj *routeMplsLabelValue
}

type marshalRouteMplsLabelValue interface {
	// ToProto marshals RouteMplsLabelValue to protobuf object *otg.RouteMplsLabelValue
	ToProto() (*otg.RouteMplsLabelValue, error)
	// ToPbText marshals RouteMplsLabelValue to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RouteMplsLabelValue to YAML text
	ToYaml() (string, error)
	// ToJson marshals RouteMplsLabelValue to JSON text
	ToJson() (string, error)
}

type unMarshalrouteMplsLabelValue struct {
	obj *routeMplsLabelValue
}

type unMarshalRouteMplsLabelValue interface {
	// FromProto unmarshals RouteMplsLabelValue from protobuf object *otg.RouteMplsLabelValue
	FromProto(msg *otg.RouteMplsLabelValue) (RouteMplsLabelValue, error)
	// FromPbText unmarshals RouteMplsLabelValue from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RouteMplsLabelValue from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RouteMplsLabelValue from JSON text
	FromJson(value string) error
}

func (obj *routeMplsLabelValue) Marshal() marshalRouteMplsLabelValue {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrouteMplsLabelValue{obj: obj}
	}
	return obj.marshaller
}

func (obj *routeMplsLabelValue) Unmarshal() unMarshalRouteMplsLabelValue {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrouteMplsLabelValue{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrouteMplsLabelValue) ToProto() (*otg.RouteMplsLabelValue, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrouteMplsLabelValue) FromProto(msg *otg.RouteMplsLabelValue) (RouteMplsLabelValue, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrouteMplsLabelValue) ToPbText() (string, error) {
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

func (m *unMarshalrouteMplsLabelValue) FromPbText(value string) error {
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

func (m *marshalrouteMplsLabelValue) ToYaml() (string, error) {
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

func (m *unMarshalrouteMplsLabelValue) FromYaml(value string) error {
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

func (m *marshalrouteMplsLabelValue) ToJson() (string, error) {
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

func (m *unMarshalrouteMplsLabelValue) FromJson(value string) error {
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

func (obj *routeMplsLabelValue) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *routeMplsLabelValue) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *routeMplsLabelValue) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *routeMplsLabelValue) Clone() (RouteMplsLabelValue, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRouteMplsLabelValue()
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

// RouteMplsLabelValue is a container of MPLS Prefix Label Value/Index in the address range.
type RouteMplsLabelValue interface {
	Validation
	// msg marshals RouteMplsLabelValue to protobuf object *otg.RouteMplsLabelValue
	// and doesn't set defaults
	msg() *otg.RouteMplsLabelValue
	// setMsg unmarshals RouteMplsLabelValue from protobuf object *otg.RouteMplsLabelValue
	// and doesn't set defaults
	setMsg(*otg.RouteMplsLabelValue) RouteMplsLabelValue
	// provides marshal interface
	Marshal() marshalRouteMplsLabelValue
	// provides unmarshal interface
	Unmarshal() unMarshalRouteMplsLabelValue
	// validate validates RouteMplsLabelValue
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RouteMplsLabelValue, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in RouteMplsLabelValue.
	Start() uint32
	// SetStart assigns uint32 provided by user to RouteMplsLabelValue
	SetStart(value uint32) RouteMplsLabelValue
	// HasStart checks if Start has been set in RouteMplsLabelValue
	HasStart() bool
	// Max returns uint32, set in RouteMplsLabelValue.
	Max() uint32
	// SetMax assigns uint32 provided by user to RouteMplsLabelValue
	SetMax(value uint32) RouteMplsLabelValue
	// HasMax checks if Max has been set in RouteMplsLabelValue
	HasMax() bool
	// Step returns uint32, set in RouteMplsLabelValue.
	Step() uint32
	// SetStep assigns uint32 provided by user to RouteMplsLabelValue
	SetStep(value uint32) RouteMplsLabelValue
	// HasStep checks if Step has been set in RouteMplsLabelValue
	HasStep() bool
}

// The starting Label Value/Index.
// Start returns a uint32
func (obj *routeMplsLabelValue) Start() uint32 {

	return *obj.obj.Start

}

// The starting Label Value/Index.
// Start returns a uint32
func (obj *routeMplsLabelValue) HasStart() bool {
	return obj.obj.Start != nil
}

// The starting Label Value/Index.
// SetStart sets the uint32 value in the RouteMplsLabelValue object
func (obj *routeMplsLabelValue) SetStart(value uint32) RouteMplsLabelValue {

	obj.obj.Start = &value
	return obj
}

// The maximum value of Label Value/Index in the range after which next Label will start from the 'start' Label.
// Max returns a uint32
func (obj *routeMplsLabelValue) Max() uint32 {

	return *obj.obj.Max

}

// The maximum value of Label Value/Index in the range after which next Label will start from the 'start' Label.
// Max returns a uint32
func (obj *routeMplsLabelValue) HasMax() bool {
	return obj.obj.Max != nil
}

// The maximum value of Label Value/Index in the range after which next Label will start from the 'start' Label.
// SetMax sets the uint32 value in the RouteMplsLabelValue object
func (obj *routeMplsLabelValue) SetMax(value uint32) RouteMplsLabelValue {

	obj.obj.Max = &value
	return obj
}

// Increments the Label Value/Index within a route range where multiple routes are present.
// Step returns a uint32
func (obj *routeMplsLabelValue) Step() uint32 {

	return *obj.obj.Step

}

// Increments the Label Value/Index within a route range where multiple routes are present.
// Step returns a uint32
func (obj *routeMplsLabelValue) HasStep() bool {
	return obj.obj.Step != nil
}

// Increments the Label Value/Index within a route range where multiple routes are present.
// SetStep sets the uint32 value in the RouteMplsLabelValue object
func (obj *routeMplsLabelValue) SetStep(value uint32) RouteMplsLabelValue {

	obj.obj.Step = &value
	return obj
}

func (obj *routeMplsLabelValue) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RouteMplsLabelValue.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Max != nil {

		if *obj.obj.Max > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RouteMplsLabelValue.Max <= 1048575 but Got %d", *obj.obj.Max))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RouteMplsLabelValue.Step <= 1048575 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *routeMplsLabelValue) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(16)
	}
	if obj.obj.Max == nil {
		obj.SetMax(1048575)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
