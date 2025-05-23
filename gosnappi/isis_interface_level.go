package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceLevel *****
type isisInterfaceLevel struct {
	validation
	obj          *otg.IsisInterfaceLevel
	marshaller   marshalIsisInterfaceLevel
	unMarshaller unMarshalIsisInterfaceLevel
}

func NewIsisInterfaceLevel() IsisInterfaceLevel {
	obj := isisInterfaceLevel{obj: &otg.IsisInterfaceLevel{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceLevel) msg() *otg.IsisInterfaceLevel {
	return obj.obj
}

func (obj *isisInterfaceLevel) setMsg(msg *otg.IsisInterfaceLevel) IsisInterfaceLevel {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceLevel struct {
	obj *isisInterfaceLevel
}

type marshalIsisInterfaceLevel interface {
	// ToProto marshals IsisInterfaceLevel to protobuf object *otg.IsisInterfaceLevel
	ToProto() (*otg.IsisInterfaceLevel, error)
	// ToPbText marshals IsisInterfaceLevel to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceLevel to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceLevel to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisInterfaceLevel to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisInterfaceLevel struct {
	obj *isisInterfaceLevel
}

type unMarshalIsisInterfaceLevel interface {
	// FromProto unmarshals IsisInterfaceLevel from protobuf object *otg.IsisInterfaceLevel
	FromProto(msg *otg.IsisInterfaceLevel) (IsisInterfaceLevel, error)
	// FromPbText unmarshals IsisInterfaceLevel from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceLevel from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceLevel from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceLevel) Marshal() marshalIsisInterfaceLevel {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceLevel{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceLevel) Unmarshal() unMarshalIsisInterfaceLevel {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceLevel{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceLevel) ToProto() (*otg.IsisInterfaceLevel, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceLevel) FromProto(msg *otg.IsisInterfaceLevel) (IsisInterfaceLevel, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceLevel) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceLevel) FromPbText(value string) error {
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

func (m *marshalisisInterfaceLevel) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceLevel) FromYaml(value string) error {
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

func (m *marshalisisInterfaceLevel) ToJsonRaw() (string, error) {
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

func (m *marshalisisInterfaceLevel) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceLevel) FromJson(value string) error {
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

func (obj *isisInterfaceLevel) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceLevel) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceLevel) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceLevel) Clone() (IsisInterfaceLevel, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceLevel()
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

// IsisInterfaceLevel is configuration for the properties of Level 1 Hello.
type IsisInterfaceLevel interface {
	Validation
	// msg marshals IsisInterfaceLevel to protobuf object *otg.IsisInterfaceLevel
	// and doesn't set defaults
	msg() *otg.IsisInterfaceLevel
	// setMsg unmarshals IsisInterfaceLevel from protobuf object *otg.IsisInterfaceLevel
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceLevel) IsisInterfaceLevel
	// provides marshal interface
	Marshal() marshalIsisInterfaceLevel
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceLevel
	// validate validates IsisInterfaceLevel
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceLevel, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Priority returns uint32, set in IsisInterfaceLevel.
	Priority() uint32
	// SetPriority assigns uint32 provided by user to IsisInterfaceLevel
	SetPriority(value uint32) IsisInterfaceLevel
	// HasPriority checks if Priority has been set in IsisInterfaceLevel
	HasPriority() bool
	// HelloInterval returns uint32, set in IsisInterfaceLevel.
	HelloInterval() uint32
	// SetHelloInterval assigns uint32 provided by user to IsisInterfaceLevel
	SetHelloInterval(value uint32) IsisInterfaceLevel
	// HasHelloInterval checks if HelloInterval has been set in IsisInterfaceLevel
	HasHelloInterval() bool
	// DeadInterval returns uint32, set in IsisInterfaceLevel.
	DeadInterval() uint32
	// SetDeadInterval assigns uint32 provided by user to IsisInterfaceLevel
	SetDeadInterval(value uint32) IsisInterfaceLevel
	// HasDeadInterval checks if DeadInterval has been set in IsisInterfaceLevel
	HasDeadInterval() bool
}

// The Priority setting in Level 1 LAN Hellos for Designated Router election.
// Priority returns a uint32
func (obj *isisInterfaceLevel) Priority() uint32 {

	return *obj.obj.Priority

}

// The Priority setting in Level 1 LAN Hellos for Designated Router election.
// Priority returns a uint32
func (obj *isisInterfaceLevel) HasPriority() bool {
	return obj.obj.Priority != nil
}

// The Priority setting in Level 1 LAN Hellos for Designated Router election.
// SetPriority sets the uint32 value in the IsisInterfaceLevel object
func (obj *isisInterfaceLevel) SetPriority(value uint32) IsisInterfaceLevel {

	obj.obj.Priority = &value
	return obj
}

// The Hello interval for Level 1 Hello messages, in seconds.
// HelloInterval returns a uint32
func (obj *isisInterfaceLevel) HelloInterval() uint32 {

	return *obj.obj.HelloInterval

}

// The Hello interval for Level 1 Hello messages, in seconds.
// HelloInterval returns a uint32
func (obj *isisInterfaceLevel) HasHelloInterval() bool {
	return obj.obj.HelloInterval != nil
}

// The Hello interval for Level 1 Hello messages, in seconds.
// SetHelloInterval sets the uint32 value in the IsisInterfaceLevel object
func (obj *isisInterfaceLevel) SetHelloInterval(value uint32) IsisInterfaceLevel {

	obj.obj.HelloInterval = &value
	return obj
}

// The Dead (Holding Time) interval for Level 1 Hello messages, in seconds.
// DeadInterval returns a uint32
func (obj *isisInterfaceLevel) DeadInterval() uint32 {

	return *obj.obj.DeadInterval

}

// The Dead (Holding Time) interval for Level 1 Hello messages, in seconds.
// DeadInterval returns a uint32
func (obj *isisInterfaceLevel) HasDeadInterval() bool {
	return obj.obj.DeadInterval != nil
}

// The Dead (Holding Time) interval for Level 1 Hello messages, in seconds.
// SetDeadInterval sets the uint32 value in the IsisInterfaceLevel object
func (obj *isisInterfaceLevel) SetDeadInterval(value uint32) IsisInterfaceLevel {

	obj.obj.DeadInterval = &value
	return obj
}

func (obj *isisInterfaceLevel) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisInterfaceLevel) setDefault() {
	if obj.obj.Priority == nil {
		obj.SetPriority(0)
	}
	if obj.obj.HelloInterval == nil {
		obj.SetHelloInterval(10)
	}
	if obj.obj.DeadInterval == nil {
		obj.SetDeadInterval(30)
	}

}
