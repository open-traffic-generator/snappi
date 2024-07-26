package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EventLink *****
type eventLink struct {
	validation
	obj          *otg.EventLink
	marshaller   marshalEventLink
	unMarshaller unMarshalEventLink
}

func NewEventLink() EventLink {
	obj := eventLink{obj: &otg.EventLink{}}
	obj.setDefault()
	return &obj
}

func (obj *eventLink) msg() *otg.EventLink {
	return obj.obj
}

func (obj *eventLink) setMsg(msg *otg.EventLink) EventLink {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaleventLink struct {
	obj *eventLink
}

type marshalEventLink interface {
	// ToProto marshals EventLink to protobuf object *otg.EventLink
	ToProto() (*otg.EventLink, error)
	// ToPbText marshals EventLink to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EventLink to YAML text
	ToYaml() (string, error)
	// ToJson marshals EventLink to JSON text
	ToJson() (string, error)
}

type unMarshaleventLink struct {
	obj *eventLink
}

type unMarshalEventLink interface {
	// FromProto unmarshals EventLink from protobuf object *otg.EventLink
	FromProto(msg *otg.EventLink) (EventLink, error)
	// FromPbText unmarshals EventLink from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EventLink from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EventLink from JSON text
	FromJson(value string) error
}

func (obj *eventLink) Marshal() marshalEventLink {
	if obj.marshaller == nil {
		obj.marshaller = &marshaleventLink{obj: obj}
	}
	return obj.marshaller
}

func (obj *eventLink) Unmarshal() unMarshalEventLink {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaleventLink{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaleventLink) ToProto() (*otg.EventLink, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaleventLink) FromProto(msg *otg.EventLink) (EventLink, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaleventLink) ToPbText() (string, error) {
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

func (m *unMarshaleventLink) FromPbText(value string) error {
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

func (m *marshaleventLink) ToYaml() (string, error) {
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

func (m *unMarshaleventLink) FromYaml(value string) error {
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

func (m *marshaleventLink) ToJson() (string, error) {
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

func (m *unMarshaleventLink) FromJson(value string) error {
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

func (obj *eventLink) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eventLink) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eventLink) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eventLink) Clone() (EventLink, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEventLink()
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

// EventLink is the optional container for link up/down event configuration.
type EventLink interface {
	Validation
	// msg marshals EventLink to protobuf object *otg.EventLink
	// and doesn't set defaults
	msg() *otg.EventLink
	// setMsg unmarshals EventLink from protobuf object *otg.EventLink
	// and doesn't set defaults
	setMsg(*otg.EventLink) EventLink
	// provides marshal interface
	Marshal() marshalEventLink
	// provides unmarshal interface
	Unmarshal() unMarshalEventLink
	// validate validates EventLink
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EventLink, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Enable returns bool, set in EventLink.
	Enable() bool
	// SetEnable assigns bool provided by user to EventLink
	SetEnable(value bool) EventLink
	// HasEnable checks if Enable has been set in EventLink
	HasEnable() bool
}

// True to enable notifications when a link up/down event occurs.
// Enable returns a bool
func (obj *eventLink) Enable() bool {

	return *obj.obj.Enable

}

// True to enable notifications when a link up/down event occurs.
// Enable returns a bool
func (obj *eventLink) HasEnable() bool {
	return obj.obj.Enable != nil
}

// True to enable notifications when a link up/down event occurs.
// SetEnable sets the bool value in the EventLink object
func (obj *eventLink) SetEnable(value bool) EventLink {

	obj.obj.Enable = &value
	return obj
}

func (obj *eventLink) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *eventLink) setDefault() {
	if obj.obj.Enable == nil {
		obj.SetEnable(false)
	}

}
