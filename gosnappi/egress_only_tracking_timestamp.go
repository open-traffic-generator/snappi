package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingTimestamp *****
type egressOnlyTrackingTimestamp struct {
	validation
	obj          *otg.EgressOnlyTrackingTimestamp
	marshaller   marshalEgressOnlyTrackingTimestamp
	unMarshaller unMarshalEgressOnlyTrackingTimestamp
}

func NewEgressOnlyTrackingTimestamp() EgressOnlyTrackingTimestamp {
	obj := egressOnlyTrackingTimestamp{obj: &otg.EgressOnlyTrackingTimestamp{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingTimestamp) msg() *otg.EgressOnlyTrackingTimestamp {
	return obj.obj
}

func (obj *egressOnlyTrackingTimestamp) setMsg(msg *otg.EgressOnlyTrackingTimestamp) EgressOnlyTrackingTimestamp {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingTimestamp struct {
	obj *egressOnlyTrackingTimestamp
}

type marshalEgressOnlyTrackingTimestamp interface {
	// ToProto marshals EgressOnlyTrackingTimestamp to protobuf object *otg.EgressOnlyTrackingTimestamp
	ToProto() (*otg.EgressOnlyTrackingTimestamp, error)
	// ToPbText marshals EgressOnlyTrackingTimestamp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingTimestamp to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingTimestamp to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals EgressOnlyTrackingTimestamp to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalegressOnlyTrackingTimestamp struct {
	obj *egressOnlyTrackingTimestamp
}

type unMarshalEgressOnlyTrackingTimestamp interface {
	// FromProto unmarshals EgressOnlyTrackingTimestamp from protobuf object *otg.EgressOnlyTrackingTimestamp
	FromProto(msg *otg.EgressOnlyTrackingTimestamp) (EgressOnlyTrackingTimestamp, error)
	// FromPbText unmarshals EgressOnlyTrackingTimestamp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingTimestamp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingTimestamp from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingTimestamp) Marshal() marshalEgressOnlyTrackingTimestamp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingTimestamp{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingTimestamp) Unmarshal() unMarshalEgressOnlyTrackingTimestamp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingTimestamp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingTimestamp) ToProto() (*otg.EgressOnlyTrackingTimestamp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingTimestamp) FromProto(msg *otg.EgressOnlyTrackingTimestamp) (EgressOnlyTrackingTimestamp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingTimestamp) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTimestamp) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingTimestamp) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTimestamp) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingTimestamp) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalegressOnlyTrackingTimestamp) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTimestamp) FromJson(value string) error {
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

func (obj *egressOnlyTrackingTimestamp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTimestamp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTimestamp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingTimestamp) Clone() (EgressOnlyTrackingTimestamp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingTimestamp()
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

// EgressOnlyTrackingTimestamp is the container for timestamp metrics.
// The container will be empty if the timestamp has not been configured for the flow.
type EgressOnlyTrackingTimestamp interface {
	Validation
	// msg marshals EgressOnlyTrackingTimestamp to protobuf object *otg.EgressOnlyTrackingTimestamp
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingTimestamp
	// setMsg unmarshals EgressOnlyTrackingTimestamp from protobuf object *otg.EgressOnlyTrackingTimestamp
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingTimestamp) EgressOnlyTrackingTimestamp
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingTimestamp
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingTimestamp
	// validate validates EgressOnlyTrackingTimestamp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingTimestamp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FirstTimestampNs returns float64, set in EgressOnlyTrackingTimestamp.
	FirstTimestampNs() float64
	// SetFirstTimestampNs assigns float64 provided by user to EgressOnlyTrackingTimestamp
	SetFirstTimestampNs(value float64) EgressOnlyTrackingTimestamp
	// HasFirstTimestampNs checks if FirstTimestampNs has been set in EgressOnlyTrackingTimestamp
	HasFirstTimestampNs() bool
	// LastTimestampNs returns float64, set in EgressOnlyTrackingTimestamp.
	LastTimestampNs() float64
	// SetLastTimestampNs assigns float64 provided by user to EgressOnlyTrackingTimestamp
	SetLastTimestampNs(value float64) EgressOnlyTrackingTimestamp
	// HasLastTimestampNs checks if LastTimestampNs has been set in EgressOnlyTrackingTimestamp
	HasLastTimestampNs() bool
}

// First timestamp in nanoseconds
// FirstTimestampNs returns a float64
func (obj *egressOnlyTrackingTimestamp) FirstTimestampNs() float64 {

	return *obj.obj.FirstTimestampNs

}

// First timestamp in nanoseconds
// FirstTimestampNs returns a float64
func (obj *egressOnlyTrackingTimestamp) HasFirstTimestampNs() bool {
	return obj.obj.FirstTimestampNs != nil
}

// First timestamp in nanoseconds
// SetFirstTimestampNs sets the float64 value in the EgressOnlyTrackingTimestamp object
func (obj *egressOnlyTrackingTimestamp) SetFirstTimestampNs(value float64) EgressOnlyTrackingTimestamp {

	obj.obj.FirstTimestampNs = &value
	return obj
}

// Last timestamp in nanoseconds
// LastTimestampNs returns a float64
func (obj *egressOnlyTrackingTimestamp) LastTimestampNs() float64 {

	return *obj.obj.LastTimestampNs

}

// Last timestamp in nanoseconds
// LastTimestampNs returns a float64
func (obj *egressOnlyTrackingTimestamp) HasLastTimestampNs() bool {
	return obj.obj.LastTimestampNs != nil
}

// Last timestamp in nanoseconds
// SetLastTimestampNs sets the float64 value in the EgressOnlyTrackingTimestamp object
func (obj *egressOnlyTrackingTimestamp) SetLastTimestampNs(value float64) EgressOnlyTrackingTimestamp {

	obj.obj.LastTimestampNs = &value
	return obj
}

func (obj *egressOnlyTrackingTimestamp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *egressOnlyTrackingTimestamp) setDefault() {

}
