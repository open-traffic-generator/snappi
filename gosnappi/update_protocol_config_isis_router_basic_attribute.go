package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisRouterBasicAttribute *****
type updateProtocolConfigIsisRouterBasicAttribute struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisRouterBasicAttribute
	marshaller   marshalUpdateProtocolConfigIsisRouterBasicAttribute
	unMarshaller unMarshalUpdateProtocolConfigIsisRouterBasicAttribute
}

func NewUpdateProtocolConfigIsisRouterBasicAttribute() UpdateProtocolConfigIsisRouterBasicAttribute {
	obj := updateProtocolConfigIsisRouterBasicAttribute{obj: &otg.UpdateProtocolConfigIsisRouterBasicAttribute{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) msg() *otg.UpdateProtocolConfigIsisRouterBasicAttribute {
	return obj.obj
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) setMsg(msg *otg.UpdateProtocolConfigIsisRouterBasicAttribute) UpdateProtocolConfigIsisRouterBasicAttribute {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisRouterBasicAttribute struct {
	obj *updateProtocolConfigIsisRouterBasicAttribute
}

type marshalUpdateProtocolConfigIsisRouterBasicAttribute interface {
	// ToProto marshals UpdateProtocolConfigIsisRouterBasicAttribute to protobuf object *otg.UpdateProtocolConfigIsisRouterBasicAttribute
	ToProto() (*otg.UpdateProtocolConfigIsisRouterBasicAttribute, error)
	// ToPbText marshals UpdateProtocolConfigIsisRouterBasicAttribute to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisRouterBasicAttribute to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisRouterBasicAttribute to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisRouterBasicAttribute struct {
	obj *updateProtocolConfigIsisRouterBasicAttribute
}

type unMarshalUpdateProtocolConfigIsisRouterBasicAttribute interface {
	// FromProto unmarshals UpdateProtocolConfigIsisRouterBasicAttribute from protobuf object *otg.UpdateProtocolConfigIsisRouterBasicAttribute
	FromProto(msg *otg.UpdateProtocolConfigIsisRouterBasicAttribute) (UpdateProtocolConfigIsisRouterBasicAttribute, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisRouterBasicAttribute from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisRouterBasicAttribute from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisRouterBasicAttribute from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) Marshal() marshalUpdateProtocolConfigIsisRouterBasicAttribute {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisRouterBasicAttribute{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) Unmarshal() unMarshalUpdateProtocolConfigIsisRouterBasicAttribute {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisRouterBasicAttribute{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisRouterBasicAttribute) ToProto() (*otg.UpdateProtocolConfigIsisRouterBasicAttribute, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisRouterBasicAttribute) FromProto(msg *otg.UpdateProtocolConfigIsisRouterBasicAttribute) (UpdateProtocolConfigIsisRouterBasicAttribute, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisRouterBasicAttribute) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterBasicAttribute) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisRouterBasicAttribute) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterBasicAttribute) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisRouterBasicAttribute) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterBasicAttribute) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisRouterBasicAttribute) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) Clone() (UpdateProtocolConfigIsisRouterBasicAttribute, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisRouterBasicAttribute()
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

// UpdateProtocolConfigIsisRouterBasicAttribute is the basic properties of an ISIS Router to be updated.
type UpdateProtocolConfigIsisRouterBasicAttribute interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisRouterBasicAttribute to protobuf object *otg.UpdateProtocolConfigIsisRouterBasicAttribute
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisRouterBasicAttribute
	// setMsg unmarshals UpdateProtocolConfigIsisRouterBasicAttribute from protobuf object *otg.UpdateProtocolConfigIsisRouterBasicAttribute
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisRouterBasicAttribute) UpdateProtocolConfigIsisRouterBasicAttribute
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisRouterBasicAttribute
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisRouterBasicAttribute
	// validate validates UpdateProtocolConfigIsisRouterBasicAttribute
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisRouterBasicAttribute, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LearnedLspFilter returns bool, set in UpdateProtocolConfigIsisRouterBasicAttribute.
	LearnedLspFilter() bool
	// SetLearnedLspFilter assigns bool provided by user to UpdateProtocolConfigIsisRouterBasicAttribute
	SetLearnedLspFilter(value bool) UpdateProtocolConfigIsisRouterBasicAttribute
	// HasLearnedLspFilter checks if LearnedLspFilter has been set in UpdateProtocolConfigIsisRouterBasicAttribute
	HasLearnedLspFilter() bool
	// EnableWideMetric returns bool, set in UpdateProtocolConfigIsisRouterBasicAttribute.
	EnableWideMetric() bool
	// SetEnableWideMetric assigns bool provided by user to UpdateProtocolConfigIsisRouterBasicAttribute
	SetEnableWideMetric(value bool) UpdateProtocolConfigIsisRouterBasicAttribute
	// HasEnableWideMetric checks if EnableWideMetric has been set in UpdateProtocolConfigIsisRouterBasicAttribute
	HasEnableWideMetric() bool
}

// Configuration for controlling storage of ISIS learned LSPs are received from the neighbors.
// LearnedLspFilter returns a bool
func (obj *updateProtocolConfigIsisRouterBasicAttribute) LearnedLspFilter() bool {

	return *obj.obj.LearnedLspFilter

}

// Configuration for controlling storage of ISIS learned LSPs are received from the neighbors.
// LearnedLspFilter returns a bool
func (obj *updateProtocolConfigIsisRouterBasicAttribute) HasLearnedLspFilter() bool {
	return obj.obj.LearnedLspFilter != nil
}

// Configuration for controlling storage of ISIS learned LSPs are received from the neighbors.
// SetLearnedLspFilter sets the bool value in the UpdateProtocolConfigIsisRouterBasicAttribute object
func (obj *updateProtocolConfigIsisRouterBasicAttribute) SetLearnedLspFilter(value bool) UpdateProtocolConfigIsisRouterBasicAttribute {

	obj.obj.LearnedLspFilter = &value
	return obj
}

// When set to true, it allows sending of more detailed metric information  for the routes using 32-bit wide values using TLV 135 IP reachability and more detailed reachability information for IS reachability by using TLV 22.  The detailed usage is described in RFC3784.
// EnableWideMetric returns a bool
func (obj *updateProtocolConfigIsisRouterBasicAttribute) EnableWideMetric() bool {

	return *obj.obj.EnableWideMetric

}

// When set to true, it allows sending of more detailed metric information  for the routes using 32-bit wide values using TLV 135 IP reachability and more detailed reachability information for IS reachability by using TLV 22.  The detailed usage is described in RFC3784.
// EnableWideMetric returns a bool
func (obj *updateProtocolConfigIsisRouterBasicAttribute) HasEnableWideMetric() bool {
	return obj.obj.EnableWideMetric != nil
}

// When set to true, it allows sending of more detailed metric information  for the routes using 32-bit wide values using TLV 135 IP reachability and more detailed reachability information for IS reachability by using TLV 22.  The detailed usage is described in RFC3784.
// SetEnableWideMetric sets the bool value in the UpdateProtocolConfigIsisRouterBasicAttribute object
func (obj *updateProtocolConfigIsisRouterBasicAttribute) SetEnableWideMetric(value bool) UpdateProtocolConfigIsisRouterBasicAttribute {

	obj.obj.EnableWideMetric = &value
	return obj
}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *updateProtocolConfigIsisRouterBasicAttribute) setDefault() {
	if obj.obj.LearnedLspFilter == nil {
		obj.SetLearnedLspFilter(false)
	}
	if obj.obj.EnableWideMetric == nil {
		obj.SetEnableWideMetric(true)
	}

}
