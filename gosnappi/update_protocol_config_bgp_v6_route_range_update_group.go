package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigBgpV6RouteRangeUpdateGroup *****
type updateProtocolConfigBgpV6RouteRangeUpdateGroup struct {
	validation
	obj          *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	marshaller   marshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	unMarshaller unMarshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup
}

func NewUpdateProtocolConfigBgpV6RouteRangeUpdateGroup() UpdateProtocolConfigBgpV6RouteRangeUpdateGroup {
	obj := updateProtocolConfigBgpV6RouteRangeUpdateGroup{obj: &otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) msg() *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpV6RouteRangeUpdateGroup {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup struct {
	obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup
}

type marshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup to protobuf object *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup struct {
	obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup
}

type unMarshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup from protobuf object *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) (UpdateProtocolConfigBgpV6RouteRangeUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) Marshal() marshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) ToProto() (*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) (UpdateProtocolConfigBgpV6RouteRangeUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpV6RouteRangeUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) Clone() (UpdateProtocolConfigBgpV6RouteRangeUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigBgpV6RouteRangeUpdateGroup()
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

// UpdateProtocolConfigBgpV6RouteRangeUpdateGroup is a BGP router entry for router-level attribute updates. Router-level and interface-level updates are controlled independently.
type UpdateProtocolConfigBgpV6RouteRangeUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup to protobuf object *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigBgpV6RouteRangeUpdateGroup from protobuf object *otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	// validate validates UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigBgpV6RouteRangeUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigBgpV6RouteRangeUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	SetNames(value []string) UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
}

// The names of the BGP IPv6 route ranges to be updated.
//
// x-constraint:
// - /components/schemas/Device.BgpRouter/properties/name
//
// Names returns a []string
func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of the BGP IPv6 route ranges to be updated.
//
// x-constraint:
// - /components/schemas/Device.BgpRouter/properties/name
//
// SetNames sets the []string value in the UpdateProtocolConfigBgpV6RouteRangeUpdateGroup object
func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) SetNames(value []string) UpdateProtocolConfigBgpV6RouteRangeUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *updateProtocolConfigBgpV6RouteRangeUpdateGroup) setDefault() {

}
