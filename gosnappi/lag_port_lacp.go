package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortLacp *****
type lagPortLacp struct {
	validation
	obj          *otg.LagPortLacp
	marshaller   marshalLagPortLacp
	unMarshaller unMarshalLagPortLacp
}

func NewLagPortLacp() LagPortLacp {
	obj := lagPortLacp{obj: &otg.LagPortLacp{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortLacp) msg() *otg.LagPortLacp {
	return obj.obj
}

func (obj *lagPortLacp) setMsg(msg *otg.LagPortLacp) LagPortLacp {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortLacp struct {
	obj *lagPortLacp
}

type marshalLagPortLacp interface {
	// ToProto marshals LagPortLacp to protobuf object *otg.LagPortLacp
	ToProto() (*otg.LagPortLacp, error)
	// ToPbText marshals LagPortLacp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortLacp to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortLacp to JSON text
	ToJson() (string, error)
}

type unMarshallagPortLacp struct {
	obj *lagPortLacp
}

type unMarshalLagPortLacp interface {
	// FromProto unmarshals LagPortLacp from protobuf object *otg.LagPortLacp
	FromProto(msg *otg.LagPortLacp) (LagPortLacp, error)
	// FromPbText unmarshals LagPortLacp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortLacp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortLacp from JSON text
	FromJson(value string) error
}

func (obj *lagPortLacp) Marshal() marshalLagPortLacp {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortLacp{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortLacp) Unmarshal() unMarshalLagPortLacp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortLacp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortLacp) ToProto() (*otg.LagPortLacp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortLacp) FromProto(msg *otg.LagPortLacp) (LagPortLacp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortLacp) ToPbText() (string, error) {
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

func (m *unMarshallagPortLacp) FromPbText(value string) error {
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

func (m *marshallagPortLacp) ToYaml() (string, error) {
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

func (m *unMarshallagPortLacp) FromYaml(value string) error {
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

func (m *marshallagPortLacp) ToJson() (string, error) {
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

func (m *unMarshallagPortLacp) FromJson(value string) error {
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

func (obj *lagPortLacp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortLacp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortLacp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortLacp) Clone() (LagPortLacp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortLacp()
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

// LagPortLacp is the container for link aggregation control protocol settings of a LAG member (port).
type LagPortLacp interface {
	Validation
	// msg marshals LagPortLacp to protobuf object *otg.LagPortLacp
	// and doesn't set defaults
	msg() *otg.LagPortLacp
	// setMsg unmarshals LagPortLacp from protobuf object *otg.LagPortLacp
	// and doesn't set defaults
	setMsg(*otg.LagPortLacp) LagPortLacp
	// provides marshal interface
	Marshal() marshalLagPortLacp
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortLacp
	// validate validates LagPortLacp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortLacp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ActorPortNumber returns uint32, set in LagPortLacp.
	ActorPortNumber() uint32
	// SetActorPortNumber assigns uint32 provided by user to LagPortLacp
	SetActorPortNumber(value uint32) LagPortLacp
	// HasActorPortNumber checks if ActorPortNumber has been set in LagPortLacp
	HasActorPortNumber() bool
	// ActorPortPriority returns uint32, set in LagPortLacp.
	ActorPortPriority() uint32
	// SetActorPortPriority assigns uint32 provided by user to LagPortLacp
	SetActorPortPriority(value uint32) LagPortLacp
	// HasActorPortPriority checks if ActorPortPriority has been set in LagPortLacp
	HasActorPortPriority() bool
	// ActorActivity returns LagPortLacpActorActivityEnum, set in LagPortLacp
	ActorActivity() LagPortLacpActorActivityEnum
	// SetActorActivity assigns LagPortLacpActorActivityEnum provided by user to LagPortLacp
	SetActorActivity(value LagPortLacpActorActivityEnum) LagPortLacp
	// HasActorActivity checks if ActorActivity has been set in LagPortLacp
	HasActorActivity() bool
	// LacpduPeriodicTimeInterval returns uint32, set in LagPortLacp.
	LacpduPeriodicTimeInterval() uint32
	// SetLacpduPeriodicTimeInterval assigns uint32 provided by user to LagPortLacp
	SetLacpduPeriodicTimeInterval(value uint32) LagPortLacp
	// HasLacpduPeriodicTimeInterval checks if LacpduPeriodicTimeInterval has been set in LagPortLacp
	HasLacpduPeriodicTimeInterval() bool
	// LacpduTimeout returns uint32, set in LagPortLacp.
	LacpduTimeout() uint32
	// SetLacpduTimeout assigns uint32 provided by user to LagPortLacp
	SetLacpduTimeout(value uint32) LagPortLacp
	// HasLacpduTimeout checks if LacpduTimeout has been set in LagPortLacp
	HasLacpduTimeout() bool
}

// The actor port number
// ActorPortNumber returns a uint32
func (obj *lagPortLacp) ActorPortNumber() uint32 {

	return *obj.obj.ActorPortNumber

}

// The actor port number
// ActorPortNumber returns a uint32
func (obj *lagPortLacp) HasActorPortNumber() bool {
	return obj.obj.ActorPortNumber != nil
}

// The actor port number
// SetActorPortNumber sets the uint32 value in the LagPortLacp object
func (obj *lagPortLacp) SetActorPortNumber(value uint32) LagPortLacp {

	obj.obj.ActorPortNumber = &value
	return obj
}

// The actor port priority
// ActorPortPriority returns a uint32
func (obj *lagPortLacp) ActorPortPriority() uint32 {

	return *obj.obj.ActorPortPriority

}

// The actor port priority
// ActorPortPriority returns a uint32
func (obj *lagPortLacp) HasActorPortPriority() bool {
	return obj.obj.ActorPortPriority != nil
}

// The actor port priority
// SetActorPortPriority sets the uint32 value in the LagPortLacp object
func (obj *lagPortLacp) SetActorPortPriority(value uint32) LagPortLacp {

	obj.obj.ActorPortPriority = &value
	return obj
}

type LagPortLacpActorActivityEnum string

// Enum of ActorActivity on LagPortLacp
var LagPortLacpActorActivity = struct {
	PASSIVE LagPortLacpActorActivityEnum
	ACTIVE  LagPortLacpActorActivityEnum
}{
	PASSIVE: LagPortLacpActorActivityEnum("passive"),
	ACTIVE:  LagPortLacpActorActivityEnum("active"),
}

func (obj *lagPortLacp) ActorActivity() LagPortLacpActorActivityEnum {
	return LagPortLacpActorActivityEnum(obj.obj.ActorActivity.Enum().String())
}

// Sets the value of LACP actor activity as either passive or active.
// Passive indicates the port's preference for not transmitting  LACPDUs unless its partner's control is Active.
// Active indicates the port's preference to participate in the  protocol regardless of the partner's control value.
// ActorActivity returns a string
func (obj *lagPortLacp) HasActorActivity() bool {
	return obj.obj.ActorActivity != nil
}

func (obj *lagPortLacp) SetActorActivity(value LagPortLacpActorActivityEnum) LagPortLacp {
	intValue, ok := otg.LagPortLacp_ActorActivity_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagPortLacpActorActivityEnum", string(value)))
		return obj
	}
	enumValue := otg.LagPortLacp_ActorActivity_Enum(intValue)
	obj.obj.ActorActivity = &enumValue

	return obj
}

// This field defines how frequently LACPDUs are sent to the link partner
// LacpduPeriodicTimeInterval returns a uint32
func (obj *lagPortLacp) LacpduPeriodicTimeInterval() uint32 {

	return *obj.obj.LacpduPeriodicTimeInterval

}

// This field defines how frequently LACPDUs are sent to the link partner
// LacpduPeriodicTimeInterval returns a uint32
func (obj *lagPortLacp) HasLacpduPeriodicTimeInterval() bool {
	return obj.obj.LacpduPeriodicTimeInterval != nil
}

// This field defines how frequently LACPDUs are sent to the link partner
// SetLacpduPeriodicTimeInterval sets the uint32 value in the LagPortLacp object
func (obj *lagPortLacp) SetLacpduPeriodicTimeInterval(value uint32) LagPortLacp {

	obj.obj.LacpduPeriodicTimeInterval = &value
	return obj
}

// This timer is used to detect whether received protocol information has expired
// LacpduTimeout returns a uint32
func (obj *lagPortLacp) LacpduTimeout() uint32 {

	return *obj.obj.LacpduTimeout

}

// This timer is used to detect whether received protocol information has expired
// LacpduTimeout returns a uint32
func (obj *lagPortLacp) HasLacpduTimeout() bool {
	return obj.obj.LacpduTimeout != nil
}

// This timer is used to detect whether received protocol information has expired
// SetLacpduTimeout sets the uint32 value in the LagPortLacp object
func (obj *lagPortLacp) SetLacpduTimeout(value uint32) LagPortLacp {

	obj.obj.LacpduTimeout = &value
	return obj
}

func (obj *lagPortLacp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ActorPortNumber != nil {

		if *obj.obj.ActorPortNumber > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagPortLacp.ActorPortNumber <= 65535 but Got %d", *obj.obj.ActorPortNumber))
		}

	}

	if obj.obj.ActorPortPriority != nil {

		if *obj.obj.ActorPortPriority > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagPortLacp.ActorPortPriority <= 65535 but Got %d", *obj.obj.ActorPortPriority))
		}

	}

	if obj.obj.LacpduPeriodicTimeInterval != nil {

		if *obj.obj.LacpduPeriodicTimeInterval > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagPortLacp.LacpduPeriodicTimeInterval <= 65535 but Got %d", *obj.obj.LacpduPeriodicTimeInterval))
		}

	}

	if obj.obj.LacpduTimeout != nil {

		if *obj.obj.LacpduTimeout > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagPortLacp.LacpduTimeout <= 65535 but Got %d", *obj.obj.LacpduTimeout))
		}

	}

}

func (obj *lagPortLacp) setDefault() {
	if obj.obj.ActorPortNumber == nil {
		obj.SetActorPortNumber(0)
	}
	if obj.obj.ActorPortPriority == nil {
		obj.SetActorPortPriority(1)
	}
	if obj.obj.ActorActivity == nil {
		obj.SetActorActivity(LagPortLacpActorActivity.ACTIVE)

	}
	if obj.obj.LacpduPeriodicTimeInterval == nil {
		obj.SetLacpduPeriodicTimeInterval(0)
	}
	if obj.obj.LacpduTimeout == nil {
		obj.SetLacpduTimeout(0)
	}

}
