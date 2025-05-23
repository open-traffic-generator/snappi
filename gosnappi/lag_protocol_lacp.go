package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagProtocolLacp *****
type lagProtocolLacp struct {
	validation
	obj          *otg.LagProtocolLacp
	marshaller   marshalLagProtocolLacp
	unMarshaller unMarshalLagProtocolLacp
}

func NewLagProtocolLacp() LagProtocolLacp {
	obj := lagProtocolLacp{obj: &otg.LagProtocolLacp{}}
	obj.setDefault()
	return &obj
}

func (obj *lagProtocolLacp) msg() *otg.LagProtocolLacp {
	return obj.obj
}

func (obj *lagProtocolLacp) setMsg(msg *otg.LagProtocolLacp) LagProtocolLacp {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagProtocolLacp struct {
	obj *lagProtocolLacp
}

type marshalLagProtocolLacp interface {
	// ToProto marshals LagProtocolLacp to protobuf object *otg.LagProtocolLacp
	ToProto() (*otg.LagProtocolLacp, error)
	// ToPbText marshals LagProtocolLacp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagProtocolLacp to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagProtocolLacp to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LagProtocolLacp to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallagProtocolLacp struct {
	obj *lagProtocolLacp
}

type unMarshalLagProtocolLacp interface {
	// FromProto unmarshals LagProtocolLacp from protobuf object *otg.LagProtocolLacp
	FromProto(msg *otg.LagProtocolLacp) (LagProtocolLacp, error)
	// FromPbText unmarshals LagProtocolLacp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagProtocolLacp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagProtocolLacp from JSON text
	FromJson(value string) error
}

func (obj *lagProtocolLacp) Marshal() marshalLagProtocolLacp {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagProtocolLacp{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagProtocolLacp) Unmarshal() unMarshalLagProtocolLacp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagProtocolLacp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagProtocolLacp) ToProto() (*otg.LagProtocolLacp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagProtocolLacp) FromProto(msg *otg.LagProtocolLacp) (LagProtocolLacp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagProtocolLacp) ToPbText() (string, error) {
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

func (m *unMarshallagProtocolLacp) FromPbText(value string) error {
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

func (m *marshallagProtocolLacp) ToYaml() (string, error) {
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

func (m *unMarshallagProtocolLacp) FromYaml(value string) error {
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

func (m *marshallagProtocolLacp) ToJsonRaw() (string, error) {
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

func (m *marshallagProtocolLacp) ToJson() (string, error) {
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

func (m *unMarshallagProtocolLacp) FromJson(value string) error {
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

func (obj *lagProtocolLacp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagProtocolLacp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagProtocolLacp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagProtocolLacp) Clone() (LagProtocolLacp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagProtocolLacp()
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

// LagProtocolLacp is the container for link aggregation control protocol settings of a LAG (ports group).
type LagProtocolLacp interface {
	Validation
	// msg marshals LagProtocolLacp to protobuf object *otg.LagProtocolLacp
	// and doesn't set defaults
	msg() *otg.LagProtocolLacp
	// setMsg unmarshals LagProtocolLacp from protobuf object *otg.LagProtocolLacp
	// and doesn't set defaults
	setMsg(*otg.LagProtocolLacp) LagProtocolLacp
	// provides marshal interface
	Marshal() marshalLagProtocolLacp
	// provides unmarshal interface
	Unmarshal() unMarshalLagProtocolLacp
	// validate validates LagProtocolLacp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagProtocolLacp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ActorSystemId returns string, set in LagProtocolLacp.
	ActorSystemId() string
	// SetActorSystemId assigns string provided by user to LagProtocolLacp
	SetActorSystemId(value string) LagProtocolLacp
	// HasActorSystemId checks if ActorSystemId has been set in LagProtocolLacp
	HasActorSystemId() bool
	// ActorSystemPriority returns uint32, set in LagProtocolLacp.
	ActorSystemPriority() uint32
	// SetActorSystemPriority assigns uint32 provided by user to LagProtocolLacp
	SetActorSystemPriority(value uint32) LagProtocolLacp
	// HasActorSystemPriority checks if ActorSystemPriority has been set in LagProtocolLacp
	HasActorSystemPriority() bool
	// ActorKey returns uint32, set in LagProtocolLacp.
	ActorKey() uint32
	// SetActorKey assigns uint32 provided by user to LagProtocolLacp
	SetActorKey(value uint32) LagProtocolLacp
	// HasActorKey checks if ActorKey has been set in LagProtocolLacp
	HasActorKey() bool
}

// The actor system id
// ActorSystemId returns a string
func (obj *lagProtocolLacp) ActorSystemId() string {

	return *obj.obj.ActorSystemId

}

// The actor system id
// ActorSystemId returns a string
func (obj *lagProtocolLacp) HasActorSystemId() bool {
	return obj.obj.ActorSystemId != nil
}

// The actor system id
// SetActorSystemId sets the string value in the LagProtocolLacp object
func (obj *lagProtocolLacp) SetActorSystemId(value string) LagProtocolLacp {

	obj.obj.ActorSystemId = &value
	return obj
}

// The actor system priority
// ActorSystemPriority returns a uint32
func (obj *lagProtocolLacp) ActorSystemPriority() uint32 {

	return *obj.obj.ActorSystemPriority

}

// The actor system priority
// ActorSystemPriority returns a uint32
func (obj *lagProtocolLacp) HasActorSystemPriority() bool {
	return obj.obj.ActorSystemPriority != nil
}

// The actor system priority
// SetActorSystemPriority sets the uint32 value in the LagProtocolLacp object
func (obj *lagProtocolLacp) SetActorSystemPriority(value uint32) LagProtocolLacp {

	obj.obj.ActorSystemPriority = &value
	return obj
}

// The actor key
// ActorKey returns a uint32
func (obj *lagProtocolLacp) ActorKey() uint32 {

	return *obj.obj.ActorKey

}

// The actor key
// ActorKey returns a uint32
func (obj *lagProtocolLacp) HasActorKey() bool {
	return obj.obj.ActorKey != nil
}

// The actor key
// SetActorKey sets the uint32 value in the LagProtocolLacp object
func (obj *lagProtocolLacp) SetActorKey(value uint32) LagProtocolLacp {

	obj.obj.ActorKey = &value
	return obj
}

func (obj *lagProtocolLacp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ActorSystemId != nil {

		err := obj.validateMac(obj.ActorSystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LagProtocolLacp.ActorSystemId"))
		}

	}

	if obj.obj.ActorSystemPriority != nil {

		if *obj.obj.ActorSystemPriority > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagProtocolLacp.ActorSystemPriority <= 65535 but Got %d", *obj.obj.ActorSystemPriority))
		}

	}

	if obj.obj.ActorKey != nil {

		if *obj.obj.ActorKey > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagProtocolLacp.ActorKey <= 65535 but Got %d", *obj.obj.ActorKey))
		}

	}

}

func (obj *lagProtocolLacp) setDefault() {
	if obj.obj.ActorSystemId == nil {
		obj.SetActorSystemId("00:00:00:00:00:00")
	}
	if obj.obj.ActorSystemPriority == nil {
		obj.SetActorSystemPriority(0)
	}
	if obj.obj.ActorKey == nil {
		obj.SetActorKey(0)
	}

}
