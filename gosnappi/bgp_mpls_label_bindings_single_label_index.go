package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpMplsLabelBindingsSingleLabelIndex *****
type bgpMplsLabelBindingsSingleLabelIndex struct {
	validation
	obj          *otg.BgpMplsLabelBindingsSingleLabelIndex
	marshaller   marshalBgpMplsLabelBindingsSingleLabelIndex
	unMarshaller unMarshalBgpMplsLabelBindingsSingleLabelIndex
	valueHolder  RouteMplsLabelIndex
}

func NewBgpMplsLabelBindingsSingleLabelIndex() BgpMplsLabelBindingsSingleLabelIndex {
	obj := bgpMplsLabelBindingsSingleLabelIndex{obj: &otg.BgpMplsLabelBindingsSingleLabelIndex{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) msg() *otg.BgpMplsLabelBindingsSingleLabelIndex {
	return obj.obj
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) setMsg(msg *otg.BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsSingleLabelIndex {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpMplsLabelBindingsSingleLabelIndex struct {
	obj *bgpMplsLabelBindingsSingleLabelIndex
}

type marshalBgpMplsLabelBindingsSingleLabelIndex interface {
	// ToProto marshals BgpMplsLabelBindingsSingleLabelIndex to protobuf object *otg.BgpMplsLabelBindingsSingleLabelIndex
	ToProto() (*otg.BgpMplsLabelBindingsSingleLabelIndex, error)
	// ToPbText marshals BgpMplsLabelBindingsSingleLabelIndex to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpMplsLabelBindingsSingleLabelIndex to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpMplsLabelBindingsSingleLabelIndex to JSON text
	ToJson() (string, error)
}

type unMarshalbgpMplsLabelBindingsSingleLabelIndex struct {
	obj *bgpMplsLabelBindingsSingleLabelIndex
}

type unMarshalBgpMplsLabelBindingsSingleLabelIndex interface {
	// FromProto unmarshals BgpMplsLabelBindingsSingleLabelIndex from protobuf object *otg.BgpMplsLabelBindingsSingleLabelIndex
	FromProto(msg *otg.BgpMplsLabelBindingsSingleLabelIndex) (BgpMplsLabelBindingsSingleLabelIndex, error)
	// FromPbText unmarshals BgpMplsLabelBindingsSingleLabelIndex from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpMplsLabelBindingsSingleLabelIndex from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpMplsLabelBindingsSingleLabelIndex from JSON text
	FromJson(value string) error
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) Marshal() marshalBgpMplsLabelBindingsSingleLabelIndex {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpMplsLabelBindingsSingleLabelIndex{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) Unmarshal() unMarshalBgpMplsLabelBindingsSingleLabelIndex {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpMplsLabelBindingsSingleLabelIndex{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpMplsLabelBindingsSingleLabelIndex) ToProto() (*otg.BgpMplsLabelBindingsSingleLabelIndex, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpMplsLabelBindingsSingleLabelIndex) FromProto(msg *otg.BgpMplsLabelBindingsSingleLabelIndex) (BgpMplsLabelBindingsSingleLabelIndex, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpMplsLabelBindingsSingleLabelIndex) ToPbText() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindingsSingleLabelIndex) FromPbText(value string) error {
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

func (m *marshalbgpMplsLabelBindingsSingleLabelIndex) ToYaml() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindingsSingleLabelIndex) FromYaml(value string) error {
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

func (m *marshalbgpMplsLabelBindingsSingleLabelIndex) ToJson() (string, error) {
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

func (m *unMarshalbgpMplsLabelBindingsSingleLabelIndex) FromJson(value string) error {
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

func (obj *bgpMplsLabelBindingsSingleLabelIndex) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) Clone() (BgpMplsLabelBindingsSingleLabelIndex, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpMplsLabelBindingsSingleLabelIndex()
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

func (obj *bgpMplsLabelBindingsSingleLabelIndex) setNil() {
	obj.valueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpMplsLabelBindingsSingleLabelIndex is container for Single MPLS SR Index.
type BgpMplsLabelBindingsSingleLabelIndex interface {
	Validation
	// msg marshals BgpMplsLabelBindingsSingleLabelIndex to protobuf object *otg.BgpMplsLabelBindingsSingleLabelIndex
	// and doesn't set defaults
	msg() *otg.BgpMplsLabelBindingsSingleLabelIndex
	// setMsg unmarshals BgpMplsLabelBindingsSingleLabelIndex from protobuf object *otg.BgpMplsLabelBindingsSingleLabelIndex
	// and doesn't set defaults
	setMsg(*otg.BgpMplsLabelBindingsSingleLabelIndex) BgpMplsLabelBindingsSingleLabelIndex
	// provides marshal interface
	Marshal() marshalBgpMplsLabelBindingsSingleLabelIndex
	// provides unmarshal interface
	Unmarshal() unMarshalBgpMplsLabelBindingsSingleLabelIndex
	// validate validates BgpMplsLabelBindingsSingleLabelIndex
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpMplsLabelBindingsSingleLabelIndex, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns RouteMplsLabelIndex, set in BgpMplsLabelBindingsSingleLabelIndex.
	// RouteMplsLabelIndex is a container of MPLS Prefix Label Index in the address range.
	Value() RouteMplsLabelIndex
	// SetValue assigns RouteMplsLabelIndex provided by user to BgpMplsLabelBindingsSingleLabelIndex.
	// RouteMplsLabelIndex is a container of MPLS Prefix Label Index in the address range.
	SetValue(value RouteMplsLabelIndex) BgpMplsLabelBindingsSingleLabelIndex
	// HasValue checks if Value has been set in BgpMplsLabelBindingsSingleLabelIndex
	HasValue() bool
	setNil()
}

// description is TBD
// Value returns a RouteMplsLabelIndex
func (obj *bgpMplsLabelBindingsSingleLabelIndex) Value() RouteMplsLabelIndex {
	if obj.obj.Value == nil {
		obj.obj.Value = NewRouteMplsLabelIndex().msg()
	}
	if obj.valueHolder == nil {
		obj.valueHolder = &routeMplsLabelIndex{obj: obj.obj.Value}
	}
	return obj.valueHolder
}

// description is TBD
// Value returns a RouteMplsLabelIndex
func (obj *bgpMplsLabelBindingsSingleLabelIndex) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the RouteMplsLabelIndex value in the BgpMplsLabelBindingsSingleLabelIndex object
func (obj *bgpMplsLabelBindingsSingleLabelIndex) SetValue(value RouteMplsLabelIndex) BgpMplsLabelBindingsSingleLabelIndex {

	obj.valueHolder = nil
	obj.obj.Value = value.msg()

	return obj
}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		obj.Value().validateObj(vObj, set_default)
	}

}

func (obj *bgpMplsLabelBindingsSingleLabelIndex) setDefault() {

}
