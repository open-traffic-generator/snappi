package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortOptions *****
type portOptions struct {
	validation
	obj          *otg.PortOptions
	marshaller   marshalPortOptions
	unMarshaller unMarshalPortOptions
}

func NewPortOptions() PortOptions {
	obj := portOptions{obj: &otg.PortOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *portOptions) msg() *otg.PortOptions {
	return obj.obj
}

func (obj *portOptions) setMsg(msg *otg.PortOptions) PortOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportOptions struct {
	obj *portOptions
}

type marshalPortOptions interface {
	// ToProto marshals PortOptions to protobuf object *otg.PortOptions
	ToProto() (*otg.PortOptions, error)
	// ToPbText marshals PortOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortOptions to JSON text
	ToJson() (string, error)
}

type unMarshalportOptions struct {
	obj *portOptions
}

type unMarshalPortOptions interface {
	// FromProto unmarshals PortOptions from protobuf object *otg.PortOptions
	FromProto(msg *otg.PortOptions) (PortOptions, error)
	// FromPbText unmarshals PortOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortOptions from JSON text
	FromJson(value string) error
}

func (obj *portOptions) Marshal() marshalPortOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *portOptions) Unmarshal() unMarshalPortOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportOptions) ToProto() (*otg.PortOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportOptions) FromProto(msg *otg.PortOptions) (PortOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportOptions) ToPbText() (string, error) {
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

func (m *unMarshalportOptions) FromPbText(value string) error {
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

func (m *marshalportOptions) ToYaml() (string, error) {
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

func (m *unMarshalportOptions) FromYaml(value string) error {
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

func (m *marshalportOptions) ToJson() (string, error) {
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

func (m *unMarshalportOptions) FromJson(value string) error {
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

func (obj *portOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portOptions) Clone() (PortOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortOptions()
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

// PortOptions is common port options that apply to all configured Port objects.
type PortOptions interface {
	Validation
	// msg marshals PortOptions to protobuf object *otg.PortOptions
	// and doesn't set defaults
	msg() *otg.PortOptions
	// setMsg unmarshals PortOptions from protobuf object *otg.PortOptions
	// and doesn't set defaults
	setMsg(*otg.PortOptions) PortOptions
	// provides marshal interface
	Marshal() marshalPortOptions
	// provides unmarshal interface
	Unmarshal() unMarshalPortOptions
	// validate validates PortOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocationPreemption returns bool, set in PortOptions.
	LocationPreemption() bool
	// SetLocationPreemption assigns bool provided by user to PortOptions
	SetLocationPreemption(value bool) PortOptions
	// HasLocationPreemption checks if LocationPreemption has been set in PortOptions
	HasLocationPreemption() bool
	// DataIntegrity returns bool, set in PortOptions.
	DataIntegrity() bool
	// SetDataIntegrity assigns bool provided by user to PortOptions
	SetDataIntegrity(value bool) PortOptions
	// HasDataIntegrity checks if DataIntegrity has been set in PortOptions
	HasDataIntegrity() bool
}

// Preempt all the test port locations as defined by the  Port.Port.properties.location. If the test ports defined by their location values are in use and  this value is true, the test ports will be preempted.
// LocationPreemption returns a bool
func (obj *portOptions) LocationPreemption() bool {

	return *obj.obj.LocationPreemption

}

// Preempt all the test port locations as defined by the  Port.Port.properties.location. If the test ports defined by their location values are in use and  this value is true, the test ports will be preempted.
// LocationPreemption returns a bool
func (obj *portOptions) HasLocationPreemption() bool {
	return obj.obj.LocationPreemption != nil
}

// Preempt all the test port locations as defined by the  Port.Port.properties.location. If the test ports defined by their location values are in use and  this value is true, the test ports will be preempted.
// SetLocationPreemption sets the bool value in the PortOptions object
func (obj *portOptions) SetLocationPreemption(value bool) PortOptions {

	obj.obj.LocationPreemption = &value
	return obj
}

// Check and validate that flow payload has not been modified when received on Rx port. It applies to all flows and if set to true, it is reported in port statistics with  frame count for both tarnished and untarnished payload.
// DataIntegrity returns a bool
func (obj *portOptions) DataIntegrity() bool {

	return *obj.obj.DataIntegrity

}

// Check and validate that flow payload has not been modified when received on Rx port. It applies to all flows and if set to true, it is reported in port statistics with  frame count for both tarnished and untarnished payload.
// DataIntegrity returns a bool
func (obj *portOptions) HasDataIntegrity() bool {
	return obj.obj.DataIntegrity != nil
}

// Check and validate that flow payload has not been modified when received on Rx port. It applies to all flows and if set to true, it is reported in port statistics with  frame count for both tarnished and untarnished payload.
// SetDataIntegrity sets the bool value in the PortOptions object
func (obj *portOptions) SetDataIntegrity(value bool) PortOptions {

	obj.obj.DataIntegrity = &value
	return obj
}

func (obj *portOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *portOptions) setDefault() {
	if obj.obj.LocationPreemption == nil {
		obj.SetLocationPreemption(false)
	}
	if obj.obj.DataIntegrity == nil {
		obj.SetDataIntegrity(false)
	}

}
