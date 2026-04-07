package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisLink *****
type updateProtocolConfigIsisLink struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisLink
	marshaller   marshalUpdateProtocolConfigIsisLink
	unMarshaller unMarshalUpdateProtocolConfigIsisLink
}

func NewUpdateProtocolConfigIsisLink() UpdateProtocolConfigIsisLink {
	obj := updateProtocolConfigIsisLink{obj: &otg.UpdateProtocolConfigIsisLink{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisLink) msg() *otg.UpdateProtocolConfigIsisLink {
	return obj.obj
}

func (obj *updateProtocolConfigIsisLink) setMsg(msg *otg.UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisLink {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisLink struct {
	obj *updateProtocolConfigIsisLink
}

type marshalUpdateProtocolConfigIsisLink interface {
	// ToProto marshals UpdateProtocolConfigIsisLink to protobuf object *otg.UpdateProtocolConfigIsisLink
	ToProto() (*otg.UpdateProtocolConfigIsisLink, error)
	// ToPbText marshals UpdateProtocolConfigIsisLink to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisLink to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisLink to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisLink struct {
	obj *updateProtocolConfigIsisLink
}

type unMarshalUpdateProtocolConfigIsisLink interface {
	// FromProto unmarshals UpdateProtocolConfigIsisLink from protobuf object *otg.UpdateProtocolConfigIsisLink
	FromProto(msg *otg.UpdateProtocolConfigIsisLink) (UpdateProtocolConfigIsisLink, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisLink from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisLink from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisLink from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisLink) Marshal() marshalUpdateProtocolConfigIsisLink {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisLink{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisLink) Unmarshal() unMarshalUpdateProtocolConfigIsisLink {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisLink{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisLink) ToProto() (*otg.UpdateProtocolConfigIsisLink, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisLink) FromProto(msg *otg.UpdateProtocolConfigIsisLink) (UpdateProtocolConfigIsisLink, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisLink) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisLink) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisLink) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisLink) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisLink) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisLink) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisLink) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisLink) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisLink) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisLink) Clone() (UpdateProtocolConfigIsisLink, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisLink()
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

// UpdateProtocolConfigIsisLink is a container of IS-IS with associated properties to be updated that may or may not able to maintain the current IS-IS Up session state always.
type UpdateProtocolConfigIsisLink interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisLink to protobuf object *otg.UpdateProtocolConfigIsisLink
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisLink
	// setMsg unmarshals UpdateProtocolConfigIsisLink from protobuf object *otg.UpdateProtocolConfigIsisLink
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisLink) UpdateProtocolConfigIsisLink
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisLink
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisLink
	// validate validates UpdateProtocolConfigIsisLink
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisLink, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// InterfaceName returns string, set in UpdateProtocolConfigIsisLink.
	InterfaceName() string
	// SetInterfaceName assigns string provided by user to UpdateProtocolConfigIsisLink
	SetInterfaceName(value string) UpdateProtocolConfigIsisLink
	// HasInterfaceName checks if InterfaceName has been set in UpdateProtocolConfigIsisLink
	HasInterfaceName() bool
	// Metric returns uint32, set in UpdateProtocolConfigIsisLink.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to UpdateProtocolConfigIsisLink
	SetMetric(value uint32) UpdateProtocolConfigIsisLink
	// HasMetric checks if Metric has been set in UpdateProtocolConfigIsisLink
	HasMetric() bool
}

// The name of an IS-IS emulated or simulated interface.
//
// x-constraint:
// - /components/schemas/Isis.Interface/properties/name
//
// InterfaceName returns a string
func (obj *updateProtocolConfigIsisLink) InterfaceName() string {

	return *obj.obj.InterfaceName

}

// The name of an IS-IS emulated or simulated interface.
//
// x-constraint:
// - /components/schemas/Isis.Interface/properties/name
//
// InterfaceName returns a string
func (obj *updateProtocolConfigIsisLink) HasInterfaceName() bool {
	return obj.obj.InterfaceName != nil
}

// The name of an IS-IS emulated or simulated interface.
//
// x-constraint:
// - /components/schemas/Isis.Interface/properties/name
//
// SetInterfaceName sets the string value in the UpdateProtocolConfigIsisLink object
func (obj *updateProtocolConfigIsisLink) SetInterfaceName(value string) UpdateProtocolConfigIsisLink {

	obj.obj.InterfaceName = &value
	return obj
}

// The default metric cost for the interface.
// Metric returns a uint32
func (obj *updateProtocolConfigIsisLink) Metric() uint32 {

	return *obj.obj.Metric

}

// The default metric cost for the interface.
// Metric returns a uint32
func (obj *updateProtocolConfigIsisLink) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The default metric cost for the interface.
// SetMetric sets the uint32 value in the UpdateProtocolConfigIsisLink object
func (obj *updateProtocolConfigIsisLink) SetMetric(value uint32) UpdateProtocolConfigIsisLink {

	obj.obj.Metric = &value
	return obj
}

func (obj *updateProtocolConfigIsisLink) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Metric != nil {

		if *obj.obj.Metric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigIsisLink.Metric <= 16777215 but Got %d", *obj.obj.Metric))
		}

	}

}

func (obj *updateProtocolConfigIsisLink) setDefault() {
	if obj.obj.Metric == nil {
		obj.SetMetric(10)
	}

}
