package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RouteMplsLabelIndex *****
type routeMplsLabelIndex struct {
	validation
	obj          *otg.RouteMplsLabelIndex
	marshaller   marshalRouteMplsLabelIndex
	unMarshaller unMarshalRouteMplsLabelIndex
}

func NewRouteMplsLabelIndex() RouteMplsLabelIndex {
	obj := routeMplsLabelIndex{obj: &otg.RouteMplsLabelIndex{}}
	obj.setDefault()
	return &obj
}

func (obj *routeMplsLabelIndex) msg() *otg.RouteMplsLabelIndex {
	return obj.obj
}

func (obj *routeMplsLabelIndex) setMsg(msg *otg.RouteMplsLabelIndex) RouteMplsLabelIndex {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrouteMplsLabelIndex struct {
	obj *routeMplsLabelIndex
}

type marshalRouteMplsLabelIndex interface {
	// ToProto marshals RouteMplsLabelIndex to protobuf object *otg.RouteMplsLabelIndex
	ToProto() (*otg.RouteMplsLabelIndex, error)
	// ToPbText marshals RouteMplsLabelIndex to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RouteMplsLabelIndex to YAML text
	ToYaml() (string, error)
	// ToJson marshals RouteMplsLabelIndex to JSON text
	ToJson() (string, error)
}

type unMarshalrouteMplsLabelIndex struct {
	obj *routeMplsLabelIndex
}

type unMarshalRouteMplsLabelIndex interface {
	// FromProto unmarshals RouteMplsLabelIndex from protobuf object *otg.RouteMplsLabelIndex
	FromProto(msg *otg.RouteMplsLabelIndex) (RouteMplsLabelIndex, error)
	// FromPbText unmarshals RouteMplsLabelIndex from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RouteMplsLabelIndex from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RouteMplsLabelIndex from JSON text
	FromJson(value string) error
}

func (obj *routeMplsLabelIndex) Marshal() marshalRouteMplsLabelIndex {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrouteMplsLabelIndex{obj: obj}
	}
	return obj.marshaller
}

func (obj *routeMplsLabelIndex) Unmarshal() unMarshalRouteMplsLabelIndex {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrouteMplsLabelIndex{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrouteMplsLabelIndex) ToProto() (*otg.RouteMplsLabelIndex, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrouteMplsLabelIndex) FromProto(msg *otg.RouteMplsLabelIndex) (RouteMplsLabelIndex, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrouteMplsLabelIndex) ToPbText() (string, error) {
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

func (m *unMarshalrouteMplsLabelIndex) FromPbText(value string) error {
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

func (m *marshalrouteMplsLabelIndex) ToYaml() (string, error) {
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

func (m *unMarshalrouteMplsLabelIndex) FromYaml(value string) error {
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

func (m *marshalrouteMplsLabelIndex) ToJson() (string, error) {
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

func (m *unMarshalrouteMplsLabelIndex) FromJson(value string) error {
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

func (obj *routeMplsLabelIndex) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *routeMplsLabelIndex) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *routeMplsLabelIndex) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *routeMplsLabelIndex) Clone() (RouteMplsLabelIndex, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRouteMplsLabelIndex()
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

// RouteMplsLabelIndex is a container of MPLS Prefix Label Index in the address range.
type RouteMplsLabelIndex interface {
	Validation
	// msg marshals RouteMplsLabelIndex to protobuf object *otg.RouteMplsLabelIndex
	// and doesn't set defaults
	msg() *otg.RouteMplsLabelIndex
	// setMsg unmarshals RouteMplsLabelIndex from protobuf object *otg.RouteMplsLabelIndex
	// and doesn't set defaults
	setMsg(*otg.RouteMplsLabelIndex) RouteMplsLabelIndex
	// provides marshal interface
	Marshal() marshalRouteMplsLabelIndex
	// provides unmarshal interface
	Unmarshal() unMarshalRouteMplsLabelIndex
	// validate validates RouteMplsLabelIndex
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RouteMplsLabelIndex, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in RouteMplsLabelIndex.
	Start() uint32
	// SetStart assigns uint32 provided by user to RouteMplsLabelIndex
	SetStart(value uint32) RouteMplsLabelIndex
	// HasStart checks if Start has been set in RouteMplsLabelIndex
	HasStart() bool
	// End returns uint32, set in RouteMplsLabelIndex.
	End() uint32
	// SetEnd assigns uint32 provided by user to RouteMplsLabelIndex
	SetEnd(value uint32) RouteMplsLabelIndex
	// HasEnd checks if End has been set in RouteMplsLabelIndex
	HasEnd() bool
	// Step returns uint32, set in RouteMplsLabelIndex.
	Step() uint32
	// SetStep assigns uint32 provided by user to RouteMplsLabelIndex
	SetStep(value uint32) RouteMplsLabelIndex
	// HasStep checks if Step has been set in RouteMplsLabelIndex
	HasStep() bool
}

// The starting Label Index.
// Start returns a uint32
func (obj *routeMplsLabelIndex) Start() uint32 {

	return *obj.obj.Start

}

// The starting Label Index.
// Start returns a uint32
func (obj *routeMplsLabelIndex) HasStart() bool {
	return obj.obj.Start != nil
}

// The starting Label Index.
// SetStart sets the uint32 value in the RouteMplsLabelIndex object
func (obj *routeMplsLabelIndex) SetStart(value uint32) RouteMplsLabelIndex {

	obj.obj.Start = &value
	return obj
}

// The total number of Label Index in the range.
// End returns a uint32
func (obj *routeMplsLabelIndex) End() uint32 {

	return *obj.obj.End

}

// The total number of Label Index in the range.
// End returns a uint32
func (obj *routeMplsLabelIndex) HasEnd() bool {
	return obj.obj.End != nil
}

// The total number of Label Index in the range.
// SetEnd sets the uint32 value in the RouteMplsLabelIndex object
func (obj *routeMplsLabelIndex) SetEnd(value uint32) RouteMplsLabelIndex {

	obj.obj.End = &value
	return obj
}

// Increments the Label Index within a route range where multiple routes are present.
// Step returns a uint32
func (obj *routeMplsLabelIndex) Step() uint32 {

	return *obj.obj.Step

}

// Increments the Label Index within a route range where multiple routes are present.
// Step returns a uint32
func (obj *routeMplsLabelIndex) HasStep() bool {
	return obj.obj.Step != nil
}

// Increments the Label Index within a route range where multiple routes are present.
// SetStep sets the uint32 value in the RouteMplsLabelIndex object
func (obj *routeMplsLabelIndex) SetStep(value uint32) RouteMplsLabelIndex {

	obj.obj.Step = &value
	return obj
}

func (obj *routeMplsLabelIndex) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start < 1 || *obj.obj.Start > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RouteMplsLabelIndex.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.End != nil {

		if *obj.obj.End < 1 || *obj.obj.End > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RouteMplsLabelIndex.End <= 1048575 but Got %d", *obj.obj.End))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < 1 || *obj.obj.Step > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= RouteMplsLabelIndex.Step <= 1048575 but Got %d", *obj.obj.Step))
		}

	}

}

func (obj *routeMplsLabelIndex) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.End == nil {
		obj.SetEnd(1048575)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}

}
