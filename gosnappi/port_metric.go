package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortMetric *****
type portMetric struct {
	validation
	obj          *otg.PortMetric
	marshaller   marshalPortMetric
	unMarshaller unMarshalPortMetric
}

func NewPortMetric() PortMetric {
	obj := portMetric{obj: &otg.PortMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *portMetric) msg() *otg.PortMetric {
	return obj.obj
}

func (obj *portMetric) setMsg(msg *otg.PortMetric) PortMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportMetric struct {
	obj *portMetric
}

type marshalPortMetric interface {
	// ToProto marshals PortMetric to protobuf object *otg.PortMetric
	ToProto() (*otg.PortMetric, error)
	// ToPbText marshals PortMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortMetric to JSON text
	ToJson() (string, error)
}

type unMarshalportMetric struct {
	obj *portMetric
}

type unMarshalPortMetric interface {
	// FromProto unmarshals PortMetric from protobuf object *otg.PortMetric
	FromProto(msg *otg.PortMetric) (PortMetric, error)
	// FromPbText unmarshals PortMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortMetric from JSON text
	FromJson(value string) error
}

func (obj *portMetric) Marshal() marshalPortMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *portMetric) Unmarshal() unMarshalPortMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportMetric) ToProto() (*otg.PortMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportMetric) FromProto(msg *otg.PortMetric) (PortMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportMetric) ToPbText() (string, error) {
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

func (m *unMarshalportMetric) FromPbText(value string) error {
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

func (m *marshalportMetric) ToYaml() (string, error) {
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

func (m *unMarshalportMetric) FromYaml(value string) error {
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

func (m *marshalportMetric) ToJson() (string, error) {
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

func (m *unMarshalportMetric) FromJson(value string) error {
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

func (obj *portMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portMetric) Clone() (PortMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortMetric()
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

// PortMetric is description is TBD
type PortMetric interface {
	Validation
	// msg marshals PortMetric to protobuf object *otg.PortMetric
	// and doesn't set defaults
	msg() *otg.PortMetric
	// setMsg unmarshals PortMetric from protobuf object *otg.PortMetric
	// and doesn't set defaults
	setMsg(*otg.PortMetric) PortMetric
	// provides marshal interface
	Marshal() marshalPortMetric
	// provides unmarshal interface
	Unmarshal() unMarshalPortMetric
	// validate validates PortMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PortMetric.
	Name() string
	// SetName assigns string provided by user to PortMetric
	SetName(value string) PortMetric
	// HasName checks if Name has been set in PortMetric
	HasName() bool
	// Location returns string, set in PortMetric.
	Location() string
	// SetLocation assigns string provided by user to PortMetric
	SetLocation(value string) PortMetric
	// HasLocation checks if Location has been set in PortMetric
	HasLocation() bool
	// Link returns PortMetricLinkEnum, set in PortMetric
	Link() PortMetricLinkEnum
	// SetLink assigns PortMetricLinkEnum provided by user to PortMetric
	SetLink(value PortMetricLinkEnum) PortMetric
	// HasLink checks if Link has been set in PortMetric
	HasLink() bool
	// Capture returns PortMetricCaptureEnum, set in PortMetric
	Capture() PortMetricCaptureEnum
	// SetCapture assigns PortMetricCaptureEnum provided by user to PortMetric
	SetCapture(value PortMetricCaptureEnum) PortMetric
	// HasCapture checks if Capture has been set in PortMetric
	HasCapture() bool
	// FramesTx returns uint64, set in PortMetric.
	FramesTx() uint64
	// SetFramesTx assigns uint64 provided by user to PortMetric
	SetFramesTx(value uint64) PortMetric
	// HasFramesTx checks if FramesTx has been set in PortMetric
	HasFramesTx() bool
	// FramesRx returns uint64, set in PortMetric.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to PortMetric
	SetFramesRx(value uint64) PortMetric
	// HasFramesRx checks if FramesRx has been set in PortMetric
	HasFramesRx() bool
	// BytesTx returns uint64, set in PortMetric.
	BytesTx() uint64
	// SetBytesTx assigns uint64 provided by user to PortMetric
	SetBytesTx(value uint64) PortMetric
	// HasBytesTx checks if BytesTx has been set in PortMetric
	HasBytesTx() bool
	// BytesRx returns uint64, set in PortMetric.
	BytesRx() uint64
	// SetBytesRx assigns uint64 provided by user to PortMetric
	SetBytesRx(value uint64) PortMetric
	// HasBytesRx checks if BytesRx has been set in PortMetric
	HasBytesRx() bool
	// FramesTxRate returns float32, set in PortMetric.
	FramesTxRate() float32
	// SetFramesTxRate assigns float32 provided by user to PortMetric
	SetFramesTxRate(value float32) PortMetric
	// HasFramesTxRate checks if FramesTxRate has been set in PortMetric
	HasFramesTxRate() bool
	// FramesRxRate returns float32, set in PortMetric.
	FramesRxRate() float32
	// SetFramesRxRate assigns float32 provided by user to PortMetric
	SetFramesRxRate(value float32) PortMetric
	// HasFramesRxRate checks if FramesRxRate has been set in PortMetric
	HasFramesRxRate() bool
	// BytesTxRate returns float32, set in PortMetric.
	BytesTxRate() float32
	// SetBytesTxRate assigns float32 provided by user to PortMetric
	SetBytesTxRate(value float32) PortMetric
	// HasBytesTxRate checks if BytesTxRate has been set in PortMetric
	HasBytesTxRate() bool
	// BytesRxRate returns float32, set in PortMetric.
	BytesRxRate() float32
	// SetBytesRxRate assigns float32 provided by user to PortMetric
	SetBytesRxRate(value float32) PortMetric
	// HasBytesRxRate checks if BytesRxRate has been set in PortMetric
	HasBytesRxRate() bool
	// Transmit returns PortMetricTransmitEnum, set in PortMetric
	Transmit() PortMetricTransmitEnum
	// SetTransmit assigns PortMetricTransmitEnum provided by user to PortMetric
	SetTransmit(value PortMetricTransmitEnum) PortMetric
	// HasTransmit checks if Transmit has been set in PortMetric
	HasTransmit() bool
	// LastLinkStateChangeTime returns string, set in PortMetric.
	LastLinkStateChangeTime() string
	// SetLastLinkStateChangeTime assigns string provided by user to PortMetric
	SetLastLinkStateChangeTime(value string) PortMetric
	// HasLastLinkStateChangeTime checks if LastLinkStateChangeTime has been set in PortMetric
	HasLastLinkStateChangeTime() bool
}

// The name of a configured port
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// Name returns a string
func (obj *portMetric) Name() string {

	return *obj.obj.Name

}

// The name of a configured port
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// Name returns a string
func (obj *portMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of a configured port
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetName sets the string value in the PortMetric object
func (obj *portMetric) SetName(value string) PortMetric {

	obj.obj.Name = &value
	return obj
}

// The state of the connection to the test port location. The format should be the configured port location along with  any custom connection state message.
// Location returns a string
func (obj *portMetric) Location() string {

	return *obj.obj.Location

}

// The state of the connection to the test port location. The format should be the configured port location along with  any custom connection state message.
// Location returns a string
func (obj *portMetric) HasLocation() bool {
	return obj.obj.Location != nil
}

// The state of the connection to the test port location. The format should be the configured port location along with  any custom connection state message.
// SetLocation sets the string value in the PortMetric object
func (obj *portMetric) SetLocation(value string) PortMetric {

	obj.obj.Location = &value
	return obj
}

type PortMetricLinkEnum string

// Enum of Link on PortMetric
var PortMetricLink = struct {
	UP   PortMetricLinkEnum
	DOWN PortMetricLinkEnum
}{
	UP:   PortMetricLinkEnum("up"),
	DOWN: PortMetricLinkEnum("down"),
}

func (obj *portMetric) Link() PortMetricLinkEnum {
	return PortMetricLinkEnum(obj.obj.Link.Enum().String())
}

// The state of the test port link The string can be up, down or a custom error message.
// Link returns a string
func (obj *portMetric) HasLink() bool {
	return obj.obj.Link != nil
}

func (obj *portMetric) SetLink(value PortMetricLinkEnum) PortMetric {
	intValue, ok := otg.PortMetric_Link_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PortMetricLinkEnum", string(value)))
		return obj
	}
	enumValue := otg.PortMetric_Link_Enum(intValue)
	obj.obj.Link = &enumValue

	return obj
}

type PortMetricCaptureEnum string

// Enum of Capture on PortMetric
var PortMetricCapture = struct {
	STARTED PortMetricCaptureEnum
	STOPPED PortMetricCaptureEnum
}{
	STARTED: PortMetricCaptureEnum("started"),
	STOPPED: PortMetricCaptureEnum("stopped"),
}

func (obj *portMetric) Capture() PortMetricCaptureEnum {
	return PortMetricCaptureEnum(obj.obj.Capture.Enum().String())
}

// The state of the test port capture infrastructure. The string can be started, stopped or a custom error message.
// Capture returns a string
func (obj *portMetric) HasCapture() bool {
	return obj.obj.Capture != nil
}

func (obj *portMetric) SetCapture(value PortMetricCaptureEnum) PortMetric {
	intValue, ok := otg.PortMetric_Capture_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PortMetricCaptureEnum", string(value)))
		return obj
	}
	enumValue := otg.PortMetric_Capture_Enum(intValue)
	obj.obj.Capture = &enumValue

	return obj
}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *portMetric) FramesTx() uint64 {

	return *obj.obj.FramesTx

}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *portMetric) HasFramesTx() bool {
	return obj.obj.FramesTx != nil
}

// The current total number of frames transmitted
// SetFramesTx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetFramesTx(value uint64) PortMetric {

	obj.obj.FramesTx = &value
	return obj
}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *portMetric) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *portMetric) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// The current total number of valid frames received
// SetFramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetFramesRx(value uint64) PortMetric {

	obj.obj.FramesRx = &value
	return obj
}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *portMetric) BytesTx() uint64 {

	return *obj.obj.BytesTx

}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *portMetric) HasBytesTx() bool {
	return obj.obj.BytesTx != nil
}

// The current total number of bytes transmitted
// SetBytesTx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetBytesTx(value uint64) PortMetric {

	obj.obj.BytesTx = &value
	return obj
}

// The current total number of valid bytes received
// BytesRx returns a uint64
func (obj *portMetric) BytesRx() uint64 {

	return *obj.obj.BytesRx

}

// The current total number of valid bytes received
// BytesRx returns a uint64
func (obj *portMetric) HasBytesRx() bool {
	return obj.obj.BytesRx != nil
}

// The current total number of valid bytes received
// SetBytesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetBytesRx(value uint64) PortMetric {

	obj.obj.BytesRx = &value
	return obj
}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *portMetric) FramesTxRate() float32 {

	return *obj.obj.FramesTxRate

}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *portMetric) HasFramesTxRate() bool {
	return obj.obj.FramesTxRate != nil
}

// The current rate of frames transmitted
// SetFramesTxRate sets the float32 value in the PortMetric object
func (obj *portMetric) SetFramesTxRate(value float32) PortMetric {

	obj.obj.FramesTxRate = &value
	return obj
}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *portMetric) FramesRxRate() float32 {

	return *obj.obj.FramesRxRate

}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *portMetric) HasFramesRxRate() bool {
	return obj.obj.FramesRxRate != nil
}

// The current rate of valid frames received
// SetFramesRxRate sets the float32 value in the PortMetric object
func (obj *portMetric) SetFramesRxRate(value float32) PortMetric {

	obj.obj.FramesRxRate = &value
	return obj
}

// The current rate of bytes transmitted
// BytesTxRate returns a float32
func (obj *portMetric) BytesTxRate() float32 {

	return *obj.obj.BytesTxRate

}

// The current rate of bytes transmitted
// BytesTxRate returns a float32
func (obj *portMetric) HasBytesTxRate() bool {
	return obj.obj.BytesTxRate != nil
}

// The current rate of bytes transmitted
// SetBytesTxRate sets the float32 value in the PortMetric object
func (obj *portMetric) SetBytesTxRate(value float32) PortMetric {

	obj.obj.BytesTxRate = &value
	return obj
}

// The current rate of bytes received
// BytesRxRate returns a float32
func (obj *portMetric) BytesRxRate() float32 {

	return *obj.obj.BytesRxRate

}

// The current rate of bytes received
// BytesRxRate returns a float32
func (obj *portMetric) HasBytesRxRate() bool {
	return obj.obj.BytesRxRate != nil
}

// The current rate of bytes received
// SetBytesRxRate sets the float32 value in the PortMetric object
func (obj *portMetric) SetBytesRxRate(value float32) PortMetric {

	obj.obj.BytesRxRate = &value
	return obj
}

type PortMetricTransmitEnum string

// Enum of Transmit on PortMetric
var PortMetricTransmit = struct {
	STARTED PortMetricTransmitEnum
	STOPPED PortMetricTransmitEnum
}{
	STARTED: PortMetricTransmitEnum("started"),
	STOPPED: PortMetricTransmitEnum("stopped"),
}

func (obj *portMetric) Transmit() PortMetricTransmitEnum {
	return PortMetricTransmitEnum(obj.obj.Transmit.Enum().String())
}

// The transmit state of the flow.
// Transmit returns a string
func (obj *portMetric) HasTransmit() bool {
	return obj.obj.Transmit != nil
}

func (obj *portMetric) SetTransmit(value PortMetricTransmitEnum) PortMetric {
	intValue, ok := otg.PortMetric_Transmit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PortMetricTransmitEnum", string(value)))
		return obj
	}
	enumValue := otg.PortMetric_Transmit_Enum(intValue)
	obj.obj.Transmit = &enumValue

	return obj
}

// The timestamp of the last link-state change event
// LastLinkStateChangeTime returns a string
func (obj *portMetric) LastLinkStateChangeTime() string {

	return *obj.obj.LastLinkStateChangeTime

}

// The timestamp of the last link-state change event
// LastLinkStateChangeTime returns a string
func (obj *portMetric) HasLastLinkStateChangeTime() bool {
	return obj.obj.LastLinkStateChangeTime != nil
}

// The timestamp of the last link-state change event
// SetLastLinkStateChangeTime sets the string value in the PortMetric object
func (obj *portMetric) SetLastLinkStateChangeTime(value string) PortMetric {

	obj.obj.LastLinkStateChangeTime = &value
	return obj
}

func (obj *portMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *portMetric) setDefault() {

}
