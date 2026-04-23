package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisInterfaceMetric *****
type updateProtocolConfigIsisInterfaceMetric struct {
	validation
	obj          *otg.UpdateProtocolConfigIsisInterfaceMetric
	marshaller   marshalUpdateProtocolConfigIsisInterfaceMetric
	unMarshaller unMarshalUpdateProtocolConfigIsisInterfaceMetric
}

func NewUpdateProtocolConfigIsisInterfaceMetric() UpdateProtocolConfigIsisInterfaceMetric {
	obj := updateProtocolConfigIsisInterfaceMetric{obj: &otg.UpdateProtocolConfigIsisInterfaceMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisInterfaceMetric) msg() *otg.UpdateProtocolConfigIsisInterfaceMetric {
	return obj.obj
}

func (obj *updateProtocolConfigIsisInterfaceMetric) setMsg(msg *otg.UpdateProtocolConfigIsisInterfaceMetric) UpdateProtocolConfigIsisInterfaceMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisInterfaceMetric struct {
	obj *updateProtocolConfigIsisInterfaceMetric
}

type marshalUpdateProtocolConfigIsisInterfaceMetric interface {
	// ToProto marshals UpdateProtocolConfigIsisInterfaceMetric to protobuf object *otg.UpdateProtocolConfigIsisInterfaceMetric
	ToProto() (*otg.UpdateProtocolConfigIsisInterfaceMetric, error)
	// ToPbText marshals UpdateProtocolConfigIsisInterfaceMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisInterfaceMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisInterfaceMetric to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisInterfaceMetric struct {
	obj *updateProtocolConfigIsisInterfaceMetric
}

type unMarshalUpdateProtocolConfigIsisInterfaceMetric interface {
	// FromProto unmarshals UpdateProtocolConfigIsisInterfaceMetric from protobuf object *otg.UpdateProtocolConfigIsisInterfaceMetric
	FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceMetric) (UpdateProtocolConfigIsisInterfaceMetric, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisInterfaceMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisInterfaceMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisInterfaceMetric from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisInterfaceMetric) Marshal() marshalUpdateProtocolConfigIsisInterfaceMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisInterfaceMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisInterfaceMetric) Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisInterfaceMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisInterfaceMetric) ToProto() (*otg.UpdateProtocolConfigIsisInterfaceMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisInterfaceMetric) FromProto(msg *otg.UpdateProtocolConfigIsisInterfaceMetric) (UpdateProtocolConfigIsisInterfaceMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisInterfaceMetric) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceMetric) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceMetric) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceMetric) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisInterfaceMetric) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisInterfaceMetric) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisInterfaceMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisInterfaceMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisInterfaceMetric) Clone() (UpdateProtocolConfigIsisInterfaceMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisInterfaceMetric()
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

// UpdateProtocolConfigIsisInterfaceMetric is the metric cost to be applied to the IS-IS interface.
type UpdateProtocolConfigIsisInterfaceMetric interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisInterfaceMetric to protobuf object *otg.UpdateProtocolConfigIsisInterfaceMetric
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisInterfaceMetric
	// setMsg unmarshals UpdateProtocolConfigIsisInterfaceMetric from protobuf object *otg.UpdateProtocolConfigIsisInterfaceMetric
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisInterfaceMetric) UpdateProtocolConfigIsisInterfaceMetric
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisInterfaceMetric
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisInterfaceMetric
	// validate validates UpdateProtocolConfigIsisInterfaceMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisInterfaceMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Value returns uint32, set in UpdateProtocolConfigIsisInterfaceMetric.
	Value() uint32
	// SetValue assigns uint32 provided by user to UpdateProtocolConfigIsisInterfaceMetric
	SetValue(value uint32) UpdateProtocolConfigIsisInterfaceMetric
	// HasValue checks if Value has been set in UpdateProtocolConfigIsisInterfaceMetric
	HasValue() bool
}

// The metric cost for the interface.
// Value returns a uint32
func (obj *updateProtocolConfigIsisInterfaceMetric) Value() uint32 {

	return *obj.obj.Value

}

// The metric cost for the interface.
// Value returns a uint32
func (obj *updateProtocolConfigIsisInterfaceMetric) HasValue() bool {
	return obj.obj.Value != nil
}

// The metric cost for the interface.
// SetValue sets the uint32 value in the UpdateProtocolConfigIsisInterfaceMetric object
func (obj *updateProtocolConfigIsisInterfaceMetric) SetValue(value uint32) UpdateProtocolConfigIsisInterfaceMetric {

	obj.obj.Value = &value
	return obj
}

func (obj *updateProtocolConfigIsisInterfaceMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigIsisInterfaceMetric.Value <= 16777215 but Got %d", *obj.obj.Value))
		}

	}

}

func (obj *updateProtocolConfigIsisInterfaceMetric) setDefault() {
	if obj.obj.Value == nil {
		obj.SetValue(10)
	}

}
