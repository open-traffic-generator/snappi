package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingTxMetrics *****
type egressOnlyTrackingTxMetrics struct {
	validation
	obj          *otg.EgressOnlyTrackingTxMetrics
	marshaller   marshalEgressOnlyTrackingTxMetrics
	unMarshaller unMarshalEgressOnlyTrackingTxMetrics
}

func NewEgressOnlyTrackingTxMetrics() EgressOnlyTrackingTxMetrics {
	obj := egressOnlyTrackingTxMetrics{obj: &otg.EgressOnlyTrackingTxMetrics{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingTxMetrics) msg() *otg.EgressOnlyTrackingTxMetrics {
	return obj.obj
}

func (obj *egressOnlyTrackingTxMetrics) setMsg(msg *otg.EgressOnlyTrackingTxMetrics) EgressOnlyTrackingTxMetrics {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingTxMetrics struct {
	obj *egressOnlyTrackingTxMetrics
}

type marshalEgressOnlyTrackingTxMetrics interface {
	// ToProto marshals EgressOnlyTrackingTxMetrics to protobuf object *otg.EgressOnlyTrackingTxMetrics
	ToProto() (*otg.EgressOnlyTrackingTxMetrics, error)
	// ToPbText marshals EgressOnlyTrackingTxMetrics to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingTxMetrics to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingTxMetrics to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals EgressOnlyTrackingTxMetrics to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalegressOnlyTrackingTxMetrics struct {
	obj *egressOnlyTrackingTxMetrics
}

type unMarshalEgressOnlyTrackingTxMetrics interface {
	// FromProto unmarshals EgressOnlyTrackingTxMetrics from protobuf object *otg.EgressOnlyTrackingTxMetrics
	FromProto(msg *otg.EgressOnlyTrackingTxMetrics) (EgressOnlyTrackingTxMetrics, error)
	// FromPbText unmarshals EgressOnlyTrackingTxMetrics from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingTxMetrics from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingTxMetrics from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingTxMetrics) Marshal() marshalEgressOnlyTrackingTxMetrics {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingTxMetrics{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingTxMetrics) Unmarshal() unMarshalEgressOnlyTrackingTxMetrics {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingTxMetrics{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingTxMetrics) ToProto() (*otg.EgressOnlyTrackingTxMetrics, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingTxMetrics) FromProto(msg *otg.EgressOnlyTrackingTxMetrics) (EgressOnlyTrackingTxMetrics, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingTxMetrics) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxMetrics) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingTxMetrics) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxMetrics) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingTxMetrics) ToJsonRaw() (string, error) {
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

func (m *marshalegressOnlyTrackingTxMetrics) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTxMetrics) FromJson(value string) error {
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

func (obj *egressOnlyTrackingTxMetrics) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTxMetrics) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTxMetrics) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingTxMetrics) Clone() (EgressOnlyTrackingTxMetrics, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingTxMetrics()
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

// EgressOnlyTrackingTxMetrics is the container for tx metrics.
// The container will be empty if the tx metrics has not been configured.
type EgressOnlyTrackingTxMetrics interface {
	Validation
	// msg marshals EgressOnlyTrackingTxMetrics to protobuf object *otg.EgressOnlyTrackingTxMetrics
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingTxMetrics
	// setMsg unmarshals EgressOnlyTrackingTxMetrics from protobuf object *otg.EgressOnlyTrackingTxMetrics
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingTxMetrics) EgressOnlyTrackingTxMetrics
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingTxMetrics
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingTxMetrics
	// validate validates EgressOnlyTrackingTxMetrics
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingTxMetrics, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortTx returns string, set in EgressOnlyTrackingTxMetrics.
	PortTx() string
	// SetPortTx assigns string provided by user to EgressOnlyTrackingTxMetrics
	SetPortTx(value string) EgressOnlyTrackingTxMetrics
	// HasPortTx checks if PortTx has been set in EgressOnlyTrackingTxMetrics
	HasPortTx() bool
	// FramesTx returns uint64, set in EgressOnlyTrackingTxMetrics.
	FramesTx() uint64
	// SetFramesTx assigns uint64 provided by user to EgressOnlyTrackingTxMetrics
	SetFramesTx(value uint64) EgressOnlyTrackingTxMetrics
	// HasFramesTx checks if FramesTx has been set in EgressOnlyTrackingTxMetrics
	HasFramesTx() bool
	// BytesTx returns uint64, set in EgressOnlyTrackingTxMetrics.
	BytesTx() uint64
	// SetBytesTx assigns uint64 provided by user to EgressOnlyTrackingTxMetrics
	SetBytesTx(value uint64) EgressOnlyTrackingTxMetrics
	// HasBytesTx checks if BytesTx has been set in EgressOnlyTrackingTxMetrics
	HasBytesTx() bool
	// FramesTxRate returns float32, set in EgressOnlyTrackingTxMetrics.
	FramesTxRate() float32
	// SetFramesTxRate assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetFramesTxRate(value float32) EgressOnlyTrackingTxMetrics
	// HasFramesTxRate checks if FramesTxRate has been set in EgressOnlyTrackingTxMetrics
	HasFramesTxRate() bool
	// Loss returns float32, set in EgressOnlyTrackingTxMetrics.
	Loss() float32
	// SetLoss assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetLoss(value float32) EgressOnlyTrackingTxMetrics
	// HasLoss checks if Loss has been set in EgressOnlyTrackingTxMetrics
	HasLoss() bool
	// TxL1RateBps returns float32, set in EgressOnlyTrackingTxMetrics.
	TxL1RateBps() float32
	// SetTxL1RateBps assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetTxL1RateBps(value float32) EgressOnlyTrackingTxMetrics
	// HasTxL1RateBps checks if TxL1RateBps has been set in EgressOnlyTrackingTxMetrics
	HasTxL1RateBps() bool
	// TxRateBytes returns float32, set in EgressOnlyTrackingTxMetrics.
	TxRateBytes() float32
	// SetTxRateBytes assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetTxRateBytes(value float32) EgressOnlyTrackingTxMetrics
	// HasTxRateBytes checks if TxRateBytes has been set in EgressOnlyTrackingTxMetrics
	HasTxRateBytes() bool
	// TxRateBps returns float32, set in EgressOnlyTrackingTxMetrics.
	TxRateBps() float32
	// SetTxRateBps assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetTxRateBps(value float32) EgressOnlyTrackingTxMetrics
	// HasTxRateBps checks if TxRateBps has been set in EgressOnlyTrackingTxMetrics
	HasTxRateBps() bool
	// TxRateKbps returns float32, set in EgressOnlyTrackingTxMetrics.
	TxRateKbps() float32
	// SetTxRateKbps assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetTxRateKbps(value float32) EgressOnlyTrackingTxMetrics
	// HasTxRateKbps checks if TxRateKbps has been set in EgressOnlyTrackingTxMetrics
	HasTxRateKbps() bool
	// TxRateMbps returns float32, set in EgressOnlyTrackingTxMetrics.
	TxRateMbps() float32
	// SetTxRateMbps assigns float32 provided by user to EgressOnlyTrackingTxMetrics
	SetTxRateMbps(value float32) EgressOnlyTrackingTxMetrics
	// HasTxRateMbps checks if TxRateMbps has been set in EgressOnlyTrackingTxMetrics
	HasTxRateMbps() bool
}

// The name of the transmit port
// PortTx returns a string
func (obj *egressOnlyTrackingTxMetrics) PortTx() string {

	return *obj.obj.PortTx

}

// The name of the transmit port
// PortTx returns a string
func (obj *egressOnlyTrackingTxMetrics) HasPortTx() bool {
	return obj.obj.PortTx != nil
}

// The name of the transmit port
// SetPortTx sets the string value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetPortTx(value string) EgressOnlyTrackingTxMetrics {

	obj.obj.PortTx = &value
	return obj
}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *egressOnlyTrackingTxMetrics) FramesTx() uint64 {

	return *obj.obj.FramesTx

}

// The current total number of frames transmitted
// FramesTx returns a uint64
func (obj *egressOnlyTrackingTxMetrics) HasFramesTx() bool {
	return obj.obj.FramesTx != nil
}

// The current total number of frames transmitted
// SetFramesTx sets the uint64 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetFramesTx(value uint64) EgressOnlyTrackingTxMetrics {

	obj.obj.FramesTx = &value
	return obj
}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *egressOnlyTrackingTxMetrics) BytesTx() uint64 {

	return *obj.obj.BytesTx

}

// The current total number of bytes transmitted
// BytesTx returns a uint64
func (obj *egressOnlyTrackingTxMetrics) HasBytesTx() bool {
	return obj.obj.BytesTx != nil
}

// The current total number of bytes transmitted
// SetBytesTx sets the uint64 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetBytesTx(value uint64) EgressOnlyTrackingTxMetrics {

	obj.obj.BytesTx = &value
	return obj
}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *egressOnlyTrackingTxMetrics) FramesTxRate() float32 {

	return *obj.obj.FramesTxRate

}

// The current rate of frames transmitted
// FramesTxRate returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasFramesTxRate() bool {
	return obj.obj.FramesTxRate != nil
}

// The current rate of frames transmitted
// SetFramesTxRate sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetFramesTxRate(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.FramesTxRate = &value
	return obj
}

// The percentage of lost frames
// Loss returns a float32
func (obj *egressOnlyTrackingTxMetrics) Loss() float32 {

	return *obj.obj.Loss

}

// The percentage of lost frames
// Loss returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasLoss() bool {
	return obj.obj.Loss != nil
}

// The percentage of lost frames
// SetLoss sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetLoss(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.Loss = &value
	return obj
}

// The Layer 1 transmission rate in bits per second.
// TxL1RateBps returns a float32
func (obj *egressOnlyTrackingTxMetrics) TxL1RateBps() float32 {

	return *obj.obj.TxL1RateBps

}

// The Layer 1 transmission rate in bits per second.
// TxL1RateBps returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasTxL1RateBps() bool {
	return obj.obj.TxL1RateBps != nil
}

// The Layer 1 transmission rate in bits per second.
// SetTxL1RateBps sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetTxL1RateBps(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.TxL1RateBps = &value
	return obj
}

// The transmission rate in bytes per second.
// TxRateBytes returns a float32
func (obj *egressOnlyTrackingTxMetrics) TxRateBytes() float32 {

	return *obj.obj.TxRateBytes

}

// The transmission rate in bytes per second.
// TxRateBytes returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasTxRateBytes() bool {
	return obj.obj.TxRateBytes != nil
}

// The transmission rate in bytes per second.
// SetTxRateBytes sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetTxRateBytes(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.TxRateBytes = &value
	return obj
}

// The transmission rate in bits per second.
// TxRateBps returns a float32
func (obj *egressOnlyTrackingTxMetrics) TxRateBps() float32 {

	return *obj.obj.TxRateBps

}

// The transmission rate in bits per second.
// TxRateBps returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasTxRateBps() bool {
	return obj.obj.TxRateBps != nil
}

// The transmission rate in bits per second.
// SetTxRateBps sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetTxRateBps(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.TxRateBps = &value
	return obj
}

// The transmission rate in Kilobits per second.
// TxRateKbps returns a float32
func (obj *egressOnlyTrackingTxMetrics) TxRateKbps() float32 {

	return *obj.obj.TxRateKbps

}

// The transmission rate in Kilobits per second.
// TxRateKbps returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasTxRateKbps() bool {
	return obj.obj.TxRateKbps != nil
}

// The transmission rate in Kilobits per second.
// SetTxRateKbps sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetTxRateKbps(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.TxRateKbps = &value
	return obj
}

// The transmission rate in Megabits per second.
// TxRateMbps returns a float32
func (obj *egressOnlyTrackingTxMetrics) TxRateMbps() float32 {

	return *obj.obj.TxRateMbps

}

// The transmission rate in Megabits per second.
// TxRateMbps returns a float32
func (obj *egressOnlyTrackingTxMetrics) HasTxRateMbps() bool {
	return obj.obj.TxRateMbps != nil
}

// The transmission rate in Megabits per second.
// SetTxRateMbps sets the float32 value in the EgressOnlyTrackingTxMetrics object
func (obj *egressOnlyTrackingTxMetrics) SetTxRateMbps(value float32) EgressOnlyTrackingTxMetrics {

	obj.obj.TxRateMbps = &value
	return obj
}

func (obj *egressOnlyTrackingTxMetrics) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *egressOnlyTrackingTxMetrics) setDefault() {

}
