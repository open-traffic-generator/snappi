package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagMetric *****
type lagMetric struct {
	validation
	obj          *otg.LagMetric
	marshaller   marshalLagMetric
	unMarshaller unMarshalLagMetric
}

func NewLagMetric() LagMetric {
	obj := lagMetric{obj: &otg.LagMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *lagMetric) msg() *otg.LagMetric {
	return obj.obj
}

func (obj *lagMetric) setMsg(msg *otg.LagMetric) LagMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagMetric struct {
	obj *lagMetric
}

type marshalLagMetric interface {
	// ToProto marshals LagMetric to protobuf object *otg.LagMetric
	ToProto() (*otg.LagMetric, error)
	// ToPbText marshals LagMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LagMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallagMetric struct {
	obj *lagMetric
}

type unMarshalLagMetric interface {
	// FromProto unmarshals LagMetric from protobuf object *otg.LagMetric
	FromProto(msg *otg.LagMetric) (LagMetric, error)
	// FromPbText unmarshals LagMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagMetric from JSON text
	FromJson(value string) error
}

func (obj *lagMetric) Marshal() marshalLagMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagMetric) Unmarshal() unMarshalLagMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagMetric) ToProto() (*otg.LagMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagMetric) FromProto(msg *otg.LagMetric) (LagMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagMetric) ToPbText() (string, error) {
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

func (m *unMarshallagMetric) FromPbText(value string) error {
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

func (m *marshallagMetric) ToYaml() (string, error) {
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

func (m *unMarshallagMetric) FromYaml(value string) error {
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

func (m *marshallagMetric) ToJsonRaw() (string, error) {
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

func (m *marshallagMetric) ToJson() (string, error) {
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

func (m *unMarshallagMetric) FromJson(value string) error {
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

func (obj *lagMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagMetric) Clone() (LagMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagMetric()
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

// LagMetric is description is TBD
type LagMetric interface {
	Validation
	// msg marshals LagMetric to protobuf object *otg.LagMetric
	// and doesn't set defaults
	msg() *otg.LagMetric
	// setMsg unmarshals LagMetric from protobuf object *otg.LagMetric
	// and doesn't set defaults
	setMsg(*otg.LagMetric) LagMetric
	// provides marshal interface
	Marshal() marshalLagMetric
	// provides unmarshal interface
	Unmarshal() unMarshalLagMetric
	// validate validates LagMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LagMetric.
	Name() string
	// SetName assigns string provided by user to LagMetric
	SetName(value string) LagMetric
	// HasName checks if Name has been set in LagMetric
	HasName() bool
	// OperStatus returns LagMetricOperStatusEnum, set in LagMetric
	OperStatus() LagMetricOperStatusEnum
	// SetOperStatus assigns LagMetricOperStatusEnum provided by user to LagMetric
	SetOperStatus(value LagMetricOperStatusEnum) LagMetric
	// HasOperStatus checks if OperStatus has been set in LagMetric
	HasOperStatus() bool
	// MemberPortsUp returns uint32, set in LagMetric.
	MemberPortsUp() uint32
	// SetMemberPortsUp assigns uint32 provided by user to LagMetric
	SetMemberPortsUp(value uint32) LagMetric
	// HasMemberPortsUp checks if MemberPortsUp has been set in LagMetric
	HasMemberPortsUp() bool
	// FramesTx returns uint64, set in LagMetric.
	FramesTx() uint64
	// SetFramesTx assigns uint64 provided by user to LagMetric
	SetFramesTx(value uint64) LagMetric
	// HasFramesTx checks if FramesTx has been set in LagMetric
	HasFramesTx() bool
	// FramesRx returns uint64, set in LagMetric.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to LagMetric
	SetFramesRx(value uint64) LagMetric
	// HasFramesRx checks if FramesRx has been set in LagMetric
	HasFramesRx() bool
	// BytesTx returns uint64, set in LagMetric.
	BytesTx() uint64
	// SetBytesTx assigns uint64 provided by user to LagMetric
	SetBytesTx(value uint64) LagMetric
	// HasBytesTx checks if BytesTx has been set in LagMetric
	HasBytesTx() bool
	// BytesRx returns uint64, set in LagMetric.
	BytesRx() uint64
	// SetBytesRx assigns uint64 provided by user to LagMetric
	SetBytesRx(value uint64) LagMetric
	// HasBytesRx checks if BytesRx has been set in LagMetric
	HasBytesRx() bool
	// FramesTxRate returns float32, set in LagMetric.
	FramesTxRate() float32
	// SetFramesTxRate assigns float32 provided by user to LagMetric
	SetFramesTxRate(value float32) LagMetric
	// HasFramesTxRate checks if FramesTxRate has been set in LagMetric
	HasFramesTxRate() bool
	// FramesRxRate returns float32, set in LagMetric.
	FramesRxRate() float32
	// SetFramesRxRate assigns float32 provided by user to LagMetric
	SetFramesRxRate(value float32) LagMetric
	// HasFramesRxRate checks if FramesRxRate has been set in LagMetric
	HasFramesRxRate() bool
	// BytesTxRate returns float32, set in LagMetric.
	BytesTxRate() float32
	// SetBytesTxRate assigns float32 provided by user to LagMetric
	SetBytesTxRate(value float32) LagMetric
	// HasBytesTxRate checks if BytesTxRate has been set in LagMetric
	HasBytesTxRate() bool
	// BytesRxRate returns float32, set in LagMetric.
	BytesRxRate() float32
	// SetBytesRxRate assigns float32 provided by user to LagMetric
	SetBytesRxRate(value float32) LagMetric
	// HasBytesRxRate checks if BytesRxRate has been set in LagMetric
	HasBytesRxRate() bool
}

// The name of a configured LAG
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// Name returns a string
func (obj *lagMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured LAG
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// Name returns a string
func (obj *lagMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured LAG
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// SetName sets the string value in the LagMetric object
func (obj *lagMetric) SetName(value string) LagMetric {

	obj.obj.Name = &value
	return obj
}

type LagMetricOperStatusEnum string

// Enum of OperStatus on LagMetric
var LagMetricOperStatus = struct {
	UP   LagMetricOperStatusEnum
	DOWN LagMetricOperStatusEnum
}{
	UP:   LagMetricOperStatusEnum("up"),
	DOWN: LagMetricOperStatusEnum("down"),
}

func (obj *lagMetric) OperStatus() LagMetricOperStatusEnum {
	return LagMetricOperStatusEnum(obj.obj.OperStatus.Enum().String())
}

// The current operational state of the LAG. The state can be up or down. State 'up' indicates member_ports_up >= min_links.
// OperStatus returns a string
func (obj *lagMetric) HasOperStatus() bool {
	return obj.obj.OperStatus != nil
}

func (obj *lagMetric) SetOperStatus(value LagMetricOperStatusEnum) LagMetric {
	intValue, ok := otg.LagMetric_OperStatus_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LagMetricOperStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.LagMetric_OperStatus_Enum(intValue)
	obj.obj.OperStatus = &enumValue

	return obj
}

// The number of LAG member ports up.
// MemberPortsUp returns a uint32
func (obj *lagMetric) MemberPortsUp() uint32 {

	return *obj.obj.MemberPortsUp

}

// The number of LAG member ports up.
// MemberPortsUp returns a uint32
func (obj *lagMetric) HasMemberPortsUp() bool {
	return obj.obj.MemberPortsUp != nil
}

// The number of LAG member ports up.
// SetMemberPortsUp sets the uint32 value in the LagMetric object
func (obj *lagMetric) SetMemberPortsUp(value uint32) LagMetric {

	obj.obj.MemberPortsUp = &value
	return obj
}

// The current total number of frames transmitted.
// FramesTx returns a uint64
func (obj *lagMetric) FramesTx() uint64 {

	return *obj.obj.FramesTx

}

// The current total number of frames transmitted.
// FramesTx returns a uint64
func (obj *lagMetric) HasFramesTx() bool {
	return obj.obj.FramesTx != nil
}

// The current total number of frames transmitted.
// SetFramesTx sets the uint64 value in the LagMetric object
func (obj *lagMetric) SetFramesTx(value uint64) LagMetric {

	obj.obj.FramesTx = &value
	return obj
}

// The current total number of valid frames received.
// FramesRx returns a uint64
func (obj *lagMetric) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// The current total number of valid frames received.
// FramesRx returns a uint64
func (obj *lagMetric) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// The current total number of valid frames received.
// SetFramesRx sets the uint64 value in the LagMetric object
func (obj *lagMetric) SetFramesRx(value uint64) LagMetric {

	obj.obj.FramesRx = &value
	return obj
}

// The current total number of bytes transmitted.
// BytesTx returns a uint64
func (obj *lagMetric) BytesTx() uint64 {

	return *obj.obj.BytesTx

}

// The current total number of bytes transmitted.
// BytesTx returns a uint64
func (obj *lagMetric) HasBytesTx() bool {
	return obj.obj.BytesTx != nil
}

// The current total number of bytes transmitted.
// SetBytesTx sets the uint64 value in the LagMetric object
func (obj *lagMetric) SetBytesTx(value uint64) LagMetric {

	obj.obj.BytesTx = &value
	return obj
}

// The current total number of valid bytes received.
// BytesRx returns a uint64
func (obj *lagMetric) BytesRx() uint64 {

	return *obj.obj.BytesRx

}

// The current total number of valid bytes received.
// BytesRx returns a uint64
func (obj *lagMetric) HasBytesRx() bool {
	return obj.obj.BytesRx != nil
}

// The current total number of valid bytes received.
// SetBytesRx sets the uint64 value in the LagMetric object
func (obj *lagMetric) SetBytesRx(value uint64) LagMetric {

	obj.obj.BytesRx = &value
	return obj
}

// The current rate of frames transmitted.
// FramesTxRate returns a float32
func (obj *lagMetric) FramesTxRate() float32 {

	return *obj.obj.FramesTxRate

}

// The current rate of frames transmitted.
// FramesTxRate returns a float32
func (obj *lagMetric) HasFramesTxRate() bool {
	return obj.obj.FramesTxRate != nil
}

// The current rate of frames transmitted.
// SetFramesTxRate sets the float32 value in the LagMetric object
func (obj *lagMetric) SetFramesTxRate(value float32) LagMetric {

	obj.obj.FramesTxRate = &value
	return obj
}

// The current rate of valid frames received.
// FramesRxRate returns a float32
func (obj *lagMetric) FramesRxRate() float32 {

	return *obj.obj.FramesRxRate

}

// The current rate of valid frames received.
// FramesRxRate returns a float32
func (obj *lagMetric) HasFramesRxRate() bool {
	return obj.obj.FramesRxRate != nil
}

// The current rate of valid frames received.
// SetFramesRxRate sets the float32 value in the LagMetric object
func (obj *lagMetric) SetFramesRxRate(value float32) LagMetric {

	obj.obj.FramesRxRate = &value
	return obj
}

// The current rate of bytes transmitted.
// BytesTxRate returns a float32
func (obj *lagMetric) BytesTxRate() float32 {

	return *obj.obj.BytesTxRate

}

// The current rate of bytes transmitted.
// BytesTxRate returns a float32
func (obj *lagMetric) HasBytesTxRate() bool {
	return obj.obj.BytesTxRate != nil
}

// The current rate of bytes transmitted.
// SetBytesTxRate sets the float32 value in the LagMetric object
func (obj *lagMetric) SetBytesTxRate(value float32) LagMetric {

	obj.obj.BytesTxRate = &value
	return obj
}

// The current rate of bytes received.
// BytesRxRate returns a float32
func (obj *lagMetric) BytesRxRate() float32 {

	return *obj.obj.BytesRxRate

}

// The current rate of bytes received.
// BytesRxRate returns a float32
func (obj *lagMetric) HasBytesRxRate() bool {
	return obj.obj.BytesRxRate != nil
}

// The current rate of bytes received.
// SetBytesRxRate sets the float32 value in the LagMetric object
func (obj *lagMetric) SetBytesRxRate(value float32) LagMetric {

	obj.obj.BytesRxRate = &value
	return obj
}

func (obj *lagMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.BytesRx != nil {

		if *obj.obj.BytesRx > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagMetric.BytesRx <= 18446744073709551615 but Got %d", *obj.obj.BytesRx))
		}

	}

}

func (obj *lagMetric) setDefault() {

}
