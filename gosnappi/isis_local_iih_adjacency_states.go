package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLocalIIHAdjacencyStates *****
type isisLocalIIHAdjacencyStates struct {
	validation
	obj                 *otg.IsisLocalIIHAdjacencyStates
	marshaller          marshalIsisLocalIIHAdjacencyStates
	unMarshaller        unMarshalIsisLocalIIHAdjacencyStates
	localStateHolder    IsisLocalIIHState
	neighborStateHolder IsisNeighborIIHState
}

func NewIsisLocalIIHAdjacencyStates() IsisLocalIIHAdjacencyStates {
	obj := isisLocalIIHAdjacencyStates{obj: &otg.IsisLocalIIHAdjacencyStates{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLocalIIHAdjacencyStates) msg() *otg.IsisLocalIIHAdjacencyStates {
	return obj.obj
}

func (obj *isisLocalIIHAdjacencyStates) setMsg(msg *otg.IsisLocalIIHAdjacencyStates) IsisLocalIIHAdjacencyStates {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLocalIIHAdjacencyStates struct {
	obj *isisLocalIIHAdjacencyStates
}

type marshalIsisLocalIIHAdjacencyStates interface {
	// ToProto marshals IsisLocalIIHAdjacencyStates to protobuf object *otg.IsisLocalIIHAdjacencyStates
	ToProto() (*otg.IsisLocalIIHAdjacencyStates, error)
	// ToPbText marshals IsisLocalIIHAdjacencyStates to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLocalIIHAdjacencyStates to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLocalIIHAdjacencyStates to JSON text
	ToJson() (string, error)
}

type unMarshalisisLocalIIHAdjacencyStates struct {
	obj *isisLocalIIHAdjacencyStates
}

type unMarshalIsisLocalIIHAdjacencyStates interface {
	// FromProto unmarshals IsisLocalIIHAdjacencyStates from protobuf object *otg.IsisLocalIIHAdjacencyStates
	FromProto(msg *otg.IsisLocalIIHAdjacencyStates) (IsisLocalIIHAdjacencyStates, error)
	// FromPbText unmarshals IsisLocalIIHAdjacencyStates from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLocalIIHAdjacencyStates from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLocalIIHAdjacencyStates from JSON text
	FromJson(value string) error
}

func (obj *isisLocalIIHAdjacencyStates) Marshal() marshalIsisLocalIIHAdjacencyStates {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLocalIIHAdjacencyStates{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLocalIIHAdjacencyStates) Unmarshal() unMarshalIsisLocalIIHAdjacencyStates {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLocalIIHAdjacencyStates{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLocalIIHAdjacencyStates) ToProto() (*otg.IsisLocalIIHAdjacencyStates, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLocalIIHAdjacencyStates) FromProto(msg *otg.IsisLocalIIHAdjacencyStates) (IsisLocalIIHAdjacencyStates, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLocalIIHAdjacencyStates) ToPbText() (string, error) {
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

func (m *unMarshalisisLocalIIHAdjacencyStates) FromPbText(value string) error {
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

func (m *marshalisisLocalIIHAdjacencyStates) ToYaml() (string, error) {
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

func (m *unMarshalisisLocalIIHAdjacencyStates) FromYaml(value string) error {
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

func (m *marshalisisLocalIIHAdjacencyStates) ToJson() (string, error) {
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

func (m *unMarshalisisLocalIIHAdjacencyStates) FromJson(value string) error {
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

func (obj *isisLocalIIHAdjacencyStates) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLocalIIHAdjacencyStates) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLocalIIHAdjacencyStates) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLocalIIHAdjacencyStates) Clone() (IsisLocalIIHAdjacencyStates, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLocalIIHAdjacencyStates()
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

func (obj *isisLocalIIHAdjacencyStates) setNil() {
	obj.localStateHolder = nil
	obj.neighborStateHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLocalIIHAdjacencyStates is information for a local adjacency.
type IsisLocalIIHAdjacencyStates interface {
	Validation
	// msg marshals IsisLocalIIHAdjacencyStates to protobuf object *otg.IsisLocalIIHAdjacencyStates
	// and doesn't set defaults
	msg() *otg.IsisLocalIIHAdjacencyStates
	// setMsg unmarshals IsisLocalIIHAdjacencyStates from protobuf object *otg.IsisLocalIIHAdjacencyStates
	// and doesn't set defaults
	setMsg(*otg.IsisLocalIIHAdjacencyStates) IsisLocalIIHAdjacencyStates
	// provides marshal interface
	Marshal() marshalIsisLocalIIHAdjacencyStates
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLocalIIHAdjacencyStates
	// validate validates IsisLocalIIHAdjacencyStates
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLocalIIHAdjacencyStates, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NeighborSystemId returns string, set in IsisLocalIIHAdjacencyStates.
	NeighborSystemId() string
	// SetNeighborSystemId assigns string provided by user to IsisLocalIIHAdjacencyStates
	SetNeighborSystemId(value string) IsisLocalIIHAdjacencyStates
	// HasNeighborSystemId checks if NeighborSystemId has been set in IsisLocalIIHAdjacencyStates
	HasNeighborSystemId() bool
	// LocalState returns IsisLocalIIHState, set in IsisLocalIIHAdjacencyStates.
	// IsisLocalIIHState is information for a local adjacency.
	LocalState() IsisLocalIIHState
	// SetLocalState assigns IsisLocalIIHState provided by user to IsisLocalIIHAdjacencyStates.
	// IsisLocalIIHState is information for a local adjacency.
	SetLocalState(value IsisLocalIIHState) IsisLocalIIHAdjacencyStates
	// HasLocalState checks if LocalState has been set in IsisLocalIIHAdjacencyStates
	HasLocalState() bool
	// NeighborState returns IsisNeighborIIHState, set in IsisLocalIIHAdjacencyStates.
	// IsisNeighborIIHState is information for neighbor adjacency State.
	NeighborState() IsisNeighborIIHState
	// SetNeighborState assigns IsisNeighborIIHState provided by user to IsisLocalIIHAdjacencyStates.
	// IsisNeighborIIHState is information for neighbor adjacency State.
	SetNeighborState(value IsisNeighborIIHState) IsisLocalIIHAdjacencyStates
	// HasNeighborState checks if NeighborState has been set in IsisLocalIIHAdjacencyStates
	HasNeighborState() bool
	setNil()
}

// System ID of a neighbor in the hex format, e.g. '650000000001'.
// NeighborSystemId returns a string
func (obj *isisLocalIIHAdjacencyStates) NeighborSystemId() string {

	return *obj.obj.NeighborSystemId

}

// System ID of a neighbor in the hex format, e.g. '650000000001'.
// NeighborSystemId returns a string
func (obj *isisLocalIIHAdjacencyStates) HasNeighborSystemId() bool {
	return obj.obj.NeighborSystemId != nil
}

// System ID of a neighbor in the hex format, e.g. '650000000001'.
// SetNeighborSystemId sets the string value in the IsisLocalIIHAdjacencyStates object
func (obj *isisLocalIIHAdjacencyStates) SetNeighborSystemId(value string) IsisLocalIIHAdjacencyStates {

	obj.obj.NeighborSystemId = &value
	return obj
}

// Local adjacency state of this ISIS router.
// LocalState returns a IsisLocalIIHState
func (obj *isisLocalIIHAdjacencyStates) LocalState() IsisLocalIIHState {
	if obj.obj.LocalState == nil {
		obj.obj.LocalState = NewIsisLocalIIHState().msg()
	}
	if obj.localStateHolder == nil {
		obj.localStateHolder = &isisLocalIIHState{obj: obj.obj.LocalState}
	}
	return obj.localStateHolder
}

// Local adjacency state of this ISIS router.
// LocalState returns a IsisLocalIIHState
func (obj *isisLocalIIHAdjacencyStates) HasLocalState() bool {
	return obj.obj.LocalState != nil
}

// Local adjacency state of this ISIS router.
// SetLocalState sets the IsisLocalIIHState value in the IsisLocalIIHAdjacencyStates object
func (obj *isisLocalIIHAdjacencyStates) SetLocalState(value IsisLocalIIHState) IsisLocalIIHAdjacencyStates {

	obj.localStateHolder = nil
	obj.obj.LocalState = value.msg()

	return obj
}

// A IS-IS neighbor that are learned by this ISIS router.
// NeighborState returns a IsisNeighborIIHState
func (obj *isisLocalIIHAdjacencyStates) NeighborState() IsisNeighborIIHState {
	if obj.obj.NeighborState == nil {
		obj.obj.NeighborState = NewIsisNeighborIIHState().msg()
	}
	if obj.neighborStateHolder == nil {
		obj.neighborStateHolder = &isisNeighborIIHState{obj: obj.obj.NeighborState}
	}
	return obj.neighborStateHolder
}

// A IS-IS neighbor that are learned by this ISIS router.
// NeighborState returns a IsisNeighborIIHState
func (obj *isisLocalIIHAdjacencyStates) HasNeighborState() bool {
	return obj.obj.NeighborState != nil
}

// A IS-IS neighbor that are learned by this ISIS router.
// SetNeighborState sets the IsisNeighborIIHState value in the IsisLocalIIHAdjacencyStates object
func (obj *isisLocalIIHAdjacencyStates) SetNeighborState(value IsisNeighborIIHState) IsisLocalIIHAdjacencyStates {

	obj.neighborStateHolder = nil
	obj.obj.NeighborState = value.msg()

	return obj
}

func (obj *isisLocalIIHAdjacencyStates) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocalState != nil {

		obj.LocalState().validateObj(vObj, set_default)
	}

	if obj.obj.NeighborState != nil {

		obj.NeighborState().validateObj(vObj, set_default)
	}

}

func (obj *isisLocalIIHAdjacencyStates) setDefault() {

}
