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
	obj                 *otg.PortMetric
	marshaller          marshalPortMetric
	unMarshaller        unMarshalPortMetric
	dataIntegrityHolder MetricDataIntegrity
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
	obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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

func (obj *portMetric) setNil() {
	obj.dataIntegrityHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
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
	// LastChange returns uint64, set in PortMetric.
	LastChange() uint64
	// SetLastChange assigns uint64 provided by user to PortMetric
	SetLastChange(value uint64) PortMetric
	// HasLastChange checks if LastChange has been set in PortMetric
	HasLastChange() bool
	// DataIntegrity returns MetricDataIntegrity, set in PortMetric.
	// MetricDataIntegrity is the container for data integrity metrics. The container will be empty if
	// options.port_options.data_integrity has not been enabled during set_config.
	DataIntegrity() MetricDataIntegrity
	// SetDataIntegrity assigns MetricDataIntegrity provided by user to PortMetric.
	// MetricDataIntegrity is the container for data integrity metrics. The container will be empty if
	// options.port_options.data_integrity has not been enabled during set_config.
	SetDataIntegrity(value MetricDataIntegrity) PortMetric
	// HasDataIntegrity checks if DataIntegrity has been set in PortMetric
	HasDataIntegrity() bool
	// Speed returns uint64, set in PortMetric.
	Speed() uint64
	// SetSpeed assigns uint64 provided by user to PortMetric
	SetSpeed(value uint64) PortMetric
	// HasSpeed checks if Speed has been set in PortMetric
	HasSpeed() bool
	// PfcClass0FramesRx returns uint64, set in PortMetric.
	PfcClass0FramesRx() uint64
	// SetPfcClass0FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass0FramesRx(value uint64) PortMetric
	// HasPfcClass0FramesRx checks if PfcClass0FramesRx has been set in PortMetric
	HasPfcClass0FramesRx() bool
	// PfcClass1FramesRx returns uint64, set in PortMetric.
	PfcClass1FramesRx() uint64
	// SetPfcClass1FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass1FramesRx(value uint64) PortMetric
	// HasPfcClass1FramesRx checks if PfcClass1FramesRx has been set in PortMetric
	HasPfcClass1FramesRx() bool
	// PfcClass2FramesRx returns uint64, set in PortMetric.
	PfcClass2FramesRx() uint64
	// SetPfcClass2FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass2FramesRx(value uint64) PortMetric
	// HasPfcClass2FramesRx checks if PfcClass2FramesRx has been set in PortMetric
	HasPfcClass2FramesRx() bool
	// PfcClass3FramesRx returns uint64, set in PortMetric.
	PfcClass3FramesRx() uint64
	// SetPfcClass3FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass3FramesRx(value uint64) PortMetric
	// HasPfcClass3FramesRx checks if PfcClass3FramesRx has been set in PortMetric
	HasPfcClass3FramesRx() bool
	// PfcClass4FramesRx returns uint64, set in PortMetric.
	PfcClass4FramesRx() uint64
	// SetPfcClass4FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass4FramesRx(value uint64) PortMetric
	// HasPfcClass4FramesRx checks if PfcClass4FramesRx has been set in PortMetric
	HasPfcClass4FramesRx() bool
	// PfcClass5FramesRx returns uint64, set in PortMetric.
	PfcClass5FramesRx() uint64
	// SetPfcClass5FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass5FramesRx(value uint64) PortMetric
	// HasPfcClass5FramesRx checks if PfcClass5FramesRx has been set in PortMetric
	HasPfcClass5FramesRx() bool
	// PfcClass6FramesRx returns uint64, set in PortMetric.
	PfcClass6FramesRx() uint64
	// SetPfcClass6FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass6FramesRx(value uint64) PortMetric
	// HasPfcClass6FramesRx checks if PfcClass6FramesRx has been set in PortMetric
	HasPfcClass6FramesRx() bool
	// PfcClass7FramesRx returns uint64, set in PortMetric.
	PfcClass7FramesRx() uint64
	// SetPfcClass7FramesRx assigns uint64 provided by user to PortMetric
	SetPfcClass7FramesRx(value uint64) PortMetric
	// HasPfcClass7FramesRx checks if PfcClass7FramesRx has been set in PortMetric
	HasPfcClass7FramesRx() bool
	setNil()
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

// The timestamp indicates the absolute time of the last
// link state change of the test port (e.g., up-to-down transition).
//
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// LastChange returns a uint64
func (obj *portMetric) LastChange() uint64 {

	return *obj.obj.LastChange

}

// The timestamp indicates the absolute time of the last
// link state change of the test port (e.g., up-to-down transition).
//
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// LastChange returns a uint64
func (obj *portMetric) HasLastChange() bool {
	return obj.obj.LastChange != nil
}

// The timestamp indicates the absolute time of the last
// link state change of the test port (e.g., up-to-down transition).
//
// The value is the timestamp in nanoseconds relative to
// the Unix Epoch (Jan 1, 1970 00:00:00 UTC).
// SetLastChange sets the uint64 value in the PortMetric object
func (obj *portMetric) SetLastChange(value uint64) PortMetric {

	obj.obj.LastChange = &value
	return obj
}

// description is TBD
// DataIntegrity returns a MetricDataIntegrity
func (obj *portMetric) DataIntegrity() MetricDataIntegrity {
	if obj.obj.DataIntegrity == nil {
		obj.obj.DataIntegrity = NewMetricDataIntegrity().msg()
	}
	if obj.dataIntegrityHolder == nil {
		obj.dataIntegrityHolder = &metricDataIntegrity{obj: obj.obj.DataIntegrity}
	}
	return obj.dataIntegrityHolder
}

// description is TBD
// DataIntegrity returns a MetricDataIntegrity
func (obj *portMetric) HasDataIntegrity() bool {
	return obj.obj.DataIntegrity != nil
}

// description is TBD
// SetDataIntegrity sets the MetricDataIntegrity value in the PortMetric object
func (obj *portMetric) SetDataIntegrity(value MetricDataIntegrity) PortMetric {

	obj.dataIntegrityHolder = nil
	obj.obj.DataIntegrity = value.msg()

	return obj
}

// The speed in KBps the frames are transmitted. The calculated speed for  a negotiated line speed of (i) 100 Gbps is 100 * 1000 * 1000 / 8 = 12500000 KBps, (ii) 1.6 Tbps (high performance devices) is 1.6 * 1000 * 1000 * 1000 / 8 = 200000000 KBps, (iii) 10 Mbps (legacy devices) is 10 * 1000 / 8 = 1250 KBps
// Speed returns a uint64
func (obj *portMetric) Speed() uint64 {

	return *obj.obj.Speed

}

// The speed in KBps the frames are transmitted. The calculated speed for  a negotiated line speed of (i) 100 Gbps is 100 * 1000 * 1000 / 8 = 12500000 KBps, (ii) 1.6 Tbps (high performance devices) is 1.6 * 1000 * 1000 * 1000 / 8 = 200000000 KBps, (iii) 10 Mbps (legacy devices) is 10 * 1000 / 8 = 1250 KBps
// Speed returns a uint64
func (obj *portMetric) HasSpeed() bool {
	return obj.obj.Speed != nil
}

// The speed in KBps the frames are transmitted. The calculated speed for  a negotiated line speed of (i) 100 Gbps is 100 * 1000 * 1000 / 8 = 12500000 KBps, (ii) 1.6 Tbps (high performance devices) is 1.6 * 1000 * 1000 * 1000 / 8 = 200000000 KBps, (iii) 10 Mbps (legacy devices) is 10 * 1000 / 8 = 1250 KBps
// SetSpeed sets the uint64 value in the PortMetric object
func (obj *portMetric) SetSpeed(value uint64) PortMetric {

	obj.obj.Speed = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 0 (priority 0) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass0FramesRx returns a uint64
func (obj *portMetric) PfcClass0FramesRx() uint64 {

	return *obj.obj.PfcClass_0FramesRx

}

// The total number of Priority Flow Control (PFC) class 0 (priority 0) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass0FramesRx returns a uint64
func (obj *portMetric) HasPfcClass0FramesRx() bool {
	return obj.obj.PfcClass_0FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 0 (priority 0) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass0FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass0FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_0FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 1 (priority 1) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass1FramesRx returns a uint64
func (obj *portMetric) PfcClass1FramesRx() uint64 {

	return *obj.obj.PfcClass_1FramesRx

}

// The total number of Priority Flow Control (PFC) class 1 (priority 1) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass1FramesRx returns a uint64
func (obj *portMetric) HasPfcClass1FramesRx() bool {
	return obj.obj.PfcClass_1FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 1 (priority 1) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass1FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass1FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_1FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 2 (priority 2) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass2FramesRx returns a uint64
func (obj *portMetric) PfcClass2FramesRx() uint64 {

	return *obj.obj.PfcClass_2FramesRx

}

// The total number of Priority Flow Control (PFC) class 2 (priority 2) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass2FramesRx returns a uint64
func (obj *portMetric) HasPfcClass2FramesRx() bool {
	return obj.obj.PfcClass_2FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 2 (priority 2) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass2FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass2FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_2FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 3 (priority 3) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass3FramesRx returns a uint64
func (obj *portMetric) PfcClass3FramesRx() uint64 {

	return *obj.obj.PfcClass_3FramesRx

}

// The total number of Priority Flow Control (PFC) class 3 (priority 3) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass3FramesRx returns a uint64
func (obj *portMetric) HasPfcClass3FramesRx() bool {
	return obj.obj.PfcClass_3FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 3 (priority 3) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass3FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass3FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_3FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 4 (priority 4) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass4FramesRx returns a uint64
func (obj *portMetric) PfcClass4FramesRx() uint64 {

	return *obj.obj.PfcClass_4FramesRx

}

// The total number of Priority Flow Control (PFC) class 4 (priority 4) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass4FramesRx returns a uint64
func (obj *portMetric) HasPfcClass4FramesRx() bool {
	return obj.obj.PfcClass_4FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 4 (priority 4) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass4FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass4FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_4FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 5 (priority 5) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass5FramesRx returns a uint64
func (obj *portMetric) PfcClass5FramesRx() uint64 {

	return *obj.obj.PfcClass_5FramesRx

}

// The total number of Priority Flow Control (PFC) class 5 (priority 5) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass5FramesRx returns a uint64
func (obj *portMetric) HasPfcClass5FramesRx() bool {
	return obj.obj.PfcClass_5FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 5 (priority 5) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass5FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass5FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_5FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 6 (priority 6) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass6FramesRx returns a uint64
func (obj *portMetric) PfcClass6FramesRx() uint64 {

	return *obj.obj.PfcClass_6FramesRx

}

// The total number of Priority Flow Control (PFC) class 6 (priority 6) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass6FramesRx returns a uint64
func (obj *portMetric) HasPfcClass6FramesRx() bool {
	return obj.obj.PfcClass_6FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 6 (priority 6) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass6FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass6FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_6FramesRx = &value
	return obj
}

// The total number of Priority Flow Control (PFC) class 7 (priority 7) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass7FramesRx returns a uint64
func (obj *portMetric) PfcClass7FramesRx() uint64 {

	return *obj.obj.PfcClass_7FramesRx

}

// The total number of Priority Flow Control (PFC) class 7 (priority 7) pause frames received on this port, as defined by IEEE 802.1Qbb.
// PfcClass7FramesRx returns a uint64
func (obj *portMetric) HasPfcClass7FramesRx() bool {
	return obj.obj.PfcClass_7FramesRx != nil
}

// The total number of Priority Flow Control (PFC) class 7 (priority 7) pause frames received on this port, as defined by IEEE 802.1Qbb.
// SetPfcClass7FramesRx sets the uint64 value in the PortMetric object
func (obj *portMetric) SetPfcClass7FramesRx(value uint64) PortMetric {

	obj.obj.PfcClass_7FramesRx = &value
	return obj
}

func (obj *portMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DataIntegrity != nil {

		obj.DataIntegrity().validateObj(vObj, set_default)
	}

}

func (obj *portMetric) setDefault() {

}
