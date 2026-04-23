package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigBgpV4RouteRangeUpdateGroup *****
type updateProtocolConfigBgpV4RouteRangeUpdateGroup struct {
	validation
	obj          *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	marshaller   marshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	unMarshaller unMarshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup
}

func NewUpdateProtocolConfigBgpV4RouteRangeUpdateGroup() UpdateProtocolConfigBgpV4RouteRangeUpdateGroup {
	obj := updateProtocolConfigBgpV4RouteRangeUpdateGroup{obj: &otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) msg() *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpV4RouteRangeUpdateGroup {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup struct {
	obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup
}

type marshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup to protobuf object *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup struct {
	obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup
}

type unMarshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup from protobuf object *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) (UpdateProtocolConfigBgpV4RouteRangeUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) Marshal() marshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) ToProto() (*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) (UpdateProtocolConfigBgpV4RouteRangeUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpV4RouteRangeUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) Clone() (UpdateProtocolConfigBgpV4RouteRangeUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigBgpV4RouteRangeUpdateGroup()
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

// UpdateProtocolConfigBgpV4RouteRangeUpdateGroup is a BGP router entry for route range-level attribute updates. Router-level and interface-level updates are controlled independently.
type UpdateProtocolConfigBgpV4RouteRangeUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup to protobuf object *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigBgpV4RouteRangeUpdateGroup from protobuf object *otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	// validate validates UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigBgpV4RouteRangeUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigBgpV4RouteRangeUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	SetNames(value []string) UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
}

// Names of the BGP IPv4 route ranges to be updated.
//
// x-constraint:
// - /components/schemas/Device.BgpRouter/properties/name
//
// Names returns a []string
func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// Names of the BGP IPv4 route ranges to be updated.
//
// x-constraint:
// - /components/schemas/Device.BgpRouter/properties/name
//
// SetNames sets the []string value in the UpdateProtocolConfigBgpV4RouteRangeUpdateGroup object
func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) SetNames(value []string) UpdateProtocolConfigBgpV4RouteRangeUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *updateProtocolConfigBgpV4RouteRangeUpdateGroup) setDefault() {

}
