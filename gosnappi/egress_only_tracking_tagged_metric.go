package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EgressOnlyTrackingTaggedMetric *****
type egressOnlyTrackingTaggedMetric struct {
	validation
	obj              *otg.EgressOnlyTrackingTaggedMetric
	marshaller       marshalEgressOnlyTrackingTaggedMetric
	unMarshaller     unMarshalEgressOnlyTrackingTaggedMetric
	tagsHolder       EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	timestampsHolder EgressOnlyTrackingTimestamp
	txMetricsHolder  EgressOnlyTrackingTxMetrics
}

func NewEgressOnlyTrackingTaggedMetric() EgressOnlyTrackingTaggedMetric {
	obj := egressOnlyTrackingTaggedMetric{obj: &otg.EgressOnlyTrackingTaggedMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *egressOnlyTrackingTaggedMetric) msg() *otg.EgressOnlyTrackingTaggedMetric {
	return obj.obj
}

func (obj *egressOnlyTrackingTaggedMetric) setMsg(msg *otg.EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingTaggedMetric {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalegressOnlyTrackingTaggedMetric struct {
	obj *egressOnlyTrackingTaggedMetric
}

type marshalEgressOnlyTrackingTaggedMetric interface {
	// ToProto marshals EgressOnlyTrackingTaggedMetric to protobuf object *otg.EgressOnlyTrackingTaggedMetric
	ToProto() (*otg.EgressOnlyTrackingTaggedMetric, error)
	// ToPbText marshals EgressOnlyTrackingTaggedMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EgressOnlyTrackingTaggedMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals EgressOnlyTrackingTaggedMetric to JSON text
	ToJson() (string, error)
}

type unMarshalegressOnlyTrackingTaggedMetric struct {
	obj *egressOnlyTrackingTaggedMetric
}

type unMarshalEgressOnlyTrackingTaggedMetric interface {
	// FromProto unmarshals EgressOnlyTrackingTaggedMetric from protobuf object *otg.EgressOnlyTrackingTaggedMetric
	FromProto(msg *otg.EgressOnlyTrackingTaggedMetric) (EgressOnlyTrackingTaggedMetric, error)
	// FromPbText unmarshals EgressOnlyTrackingTaggedMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EgressOnlyTrackingTaggedMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EgressOnlyTrackingTaggedMetric from JSON text
	FromJson(value string) error
}

func (obj *egressOnlyTrackingTaggedMetric) Marshal() marshalEgressOnlyTrackingTaggedMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalegressOnlyTrackingTaggedMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *egressOnlyTrackingTaggedMetric) Unmarshal() unMarshalEgressOnlyTrackingTaggedMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalegressOnlyTrackingTaggedMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalegressOnlyTrackingTaggedMetric) ToProto() (*otg.EgressOnlyTrackingTaggedMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalegressOnlyTrackingTaggedMetric) FromProto(msg *otg.EgressOnlyTrackingTaggedMetric) (EgressOnlyTrackingTaggedMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalegressOnlyTrackingTaggedMetric) ToPbText() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTaggedMetric) FromPbText(value string) error {
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

func (m *marshalegressOnlyTrackingTaggedMetric) ToYaml() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTaggedMetric) FromYaml(value string) error {
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

func (m *marshalegressOnlyTrackingTaggedMetric) ToJson() (string, error) {
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

func (m *unMarshalegressOnlyTrackingTaggedMetric) FromJson(value string) error {
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

func (obj *egressOnlyTrackingTaggedMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTaggedMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *egressOnlyTrackingTaggedMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *egressOnlyTrackingTaggedMetric) Clone() (EgressOnlyTrackingTaggedMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEgressOnlyTrackingTaggedMetric()
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

func (obj *egressOnlyTrackingTaggedMetric) setNil() {
	obj.tagsHolder = nil
	obj.timestampsHolder = nil
	obj.txMetricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// EgressOnlyTrackingTaggedMetric is metrics for each set of values applicable for configured
// metric tags in egress packet header fields.
// The container is keyed by list of tag-value pairs.
type EgressOnlyTrackingTaggedMetric interface {
	Validation
	// msg marshals EgressOnlyTrackingTaggedMetric to protobuf object *otg.EgressOnlyTrackingTaggedMetric
	// and doesn't set defaults
	msg() *otg.EgressOnlyTrackingTaggedMetric
	// setMsg unmarshals EgressOnlyTrackingTaggedMetric from protobuf object *otg.EgressOnlyTrackingTaggedMetric
	// and doesn't set defaults
	setMsg(*otg.EgressOnlyTrackingTaggedMetric) EgressOnlyTrackingTaggedMetric
	// provides marshal interface
	Marshal() marshalEgressOnlyTrackingTaggedMetric
	// provides unmarshal interface
	Unmarshal() unMarshalEgressOnlyTrackingTaggedMetric
	// validate validates EgressOnlyTrackingTaggedMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EgressOnlyTrackingTaggedMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Tags returns EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIterIter, set in EgressOnlyTrackingTaggedMetric
	Tags() EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	// FramesRx returns uint64, set in EgressOnlyTrackingTaggedMetric.
	FramesRx() uint64
	// SetFramesRx assigns uint64 provided by user to EgressOnlyTrackingTaggedMetric
	SetFramesRx(value uint64) EgressOnlyTrackingTaggedMetric
	// HasFramesRx checks if FramesRx has been set in EgressOnlyTrackingTaggedMetric
	HasFramesRx() bool
	// BytesRx returns uint64, set in EgressOnlyTrackingTaggedMetric.
	BytesRx() uint64
	// SetBytesRx assigns uint64 provided by user to EgressOnlyTrackingTaggedMetric
	SetBytesRx(value uint64) EgressOnlyTrackingTaggedMetric
	// HasBytesRx checks if BytesRx has been set in EgressOnlyTrackingTaggedMetric
	HasBytesRx() bool
	// FramesRxRate returns float32, set in EgressOnlyTrackingTaggedMetric.
	FramesRxRate() float32
	// SetFramesRxRate assigns float32 provided by user to EgressOnlyTrackingTaggedMetric
	SetFramesRxRate(value float32) EgressOnlyTrackingTaggedMetric
	// HasFramesRxRate checks if FramesRxRate has been set in EgressOnlyTrackingTaggedMetric
	HasFramesRxRate() bool
	// Timestamps returns EgressOnlyTrackingTimestamp, set in EgressOnlyTrackingTaggedMetric.
	// EgressOnlyTrackingTimestamp is the container for timestamp metrics.
	// The container will be empty if the timestamp has not been configured for the flow.
	Timestamps() EgressOnlyTrackingTimestamp
	// SetTimestamps assigns EgressOnlyTrackingTimestamp provided by user to EgressOnlyTrackingTaggedMetric.
	// EgressOnlyTrackingTimestamp is the container for timestamp metrics.
	// The container will be empty if the timestamp has not been configured for the flow.
	SetTimestamps(value EgressOnlyTrackingTimestamp) EgressOnlyTrackingTaggedMetric
	// HasTimestamps checks if Timestamps has been set in EgressOnlyTrackingTaggedMetric
	HasTimestamps() bool
	// RxL1RateBps returns float32, set in EgressOnlyTrackingTaggedMetric.
	RxL1RateBps() float32
	// SetRxL1RateBps assigns float32 provided by user to EgressOnlyTrackingTaggedMetric
	SetRxL1RateBps(value float32) EgressOnlyTrackingTaggedMetric
	// HasRxL1RateBps checks if RxL1RateBps has been set in EgressOnlyTrackingTaggedMetric
	HasRxL1RateBps() bool
	// RxRateBytes returns float32, set in EgressOnlyTrackingTaggedMetric.
	RxRateBytes() float32
	// SetRxRateBytes assigns float32 provided by user to EgressOnlyTrackingTaggedMetric
	SetRxRateBytes(value float32) EgressOnlyTrackingTaggedMetric
	// HasRxRateBytes checks if RxRateBytes has been set in EgressOnlyTrackingTaggedMetric
	HasRxRateBytes() bool
	// RxRateBps returns float32, set in EgressOnlyTrackingTaggedMetric.
	RxRateBps() float32
	// SetRxRateBps assigns float32 provided by user to EgressOnlyTrackingTaggedMetric
	SetRxRateBps(value float32) EgressOnlyTrackingTaggedMetric
	// HasRxRateBps checks if RxRateBps has been set in EgressOnlyTrackingTaggedMetric
	HasRxRateBps() bool
	// RxRateKbps returns float32, set in EgressOnlyTrackingTaggedMetric.
	RxRateKbps() float32
	// SetRxRateKbps assigns float32 provided by user to EgressOnlyTrackingTaggedMetric
	SetRxRateKbps(value float32) EgressOnlyTrackingTaggedMetric
	// HasRxRateKbps checks if RxRateKbps has been set in EgressOnlyTrackingTaggedMetric
	HasRxRateKbps() bool
	// RxRateMbps returns float32, set in EgressOnlyTrackingTaggedMetric.
	RxRateMbps() float32
	// SetRxRateMbps assigns float32 provided by user to EgressOnlyTrackingTaggedMetric
	SetRxRateMbps(value float32) EgressOnlyTrackingTaggedMetric
	// HasRxRateMbps checks if RxRateMbps has been set in EgressOnlyTrackingTaggedMetric
	HasRxRateMbps() bool
	// TxMetrics returns EgressOnlyTrackingTxMetrics, set in EgressOnlyTrackingTaggedMetric.
	// EgressOnlyTrackingTxMetrics is the container for tx metrics.
	// The container will be empty if the tx metrics has not been configured.
	TxMetrics() EgressOnlyTrackingTxMetrics
	// SetTxMetrics assigns EgressOnlyTrackingTxMetrics provided by user to EgressOnlyTrackingTaggedMetric.
	// EgressOnlyTrackingTxMetrics is the container for tx metrics.
	// The container will be empty if the tx metrics has not been configured.
	SetTxMetrics(value EgressOnlyTrackingTxMetrics) EgressOnlyTrackingTaggedMetric
	// HasTxMetrics checks if TxMetrics has been set in EgressOnlyTrackingTaggedMetric
	HasTxMetrics() bool
	setNil()
}

// List of tag and value pairs
// Tags returns a []EgressOnlyTrackingMetricTag
func (obj *egressOnlyTrackingTaggedMetric) Tags() EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	if len(obj.obj.Tags) == 0 {
		obj.obj.Tags = []*otg.EgressOnlyTrackingMetricTag{}
	}
	if obj.tagsHolder == nil {
		obj.tagsHolder = newEgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter(&obj.obj.Tags).setMsg(obj)
	}
	return obj.tagsHolder
}

type egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter struct {
	obj                              *egressOnlyTrackingTaggedMetric
	egressOnlyTrackingMetricTagSlice []EgressOnlyTrackingMetricTag
	fieldPtr                         *[]*otg.EgressOnlyTrackingMetricTag
}

func newEgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter(ptr *[]*otg.EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	return &egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter{fieldPtr: ptr}
}

type EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter interface {
	setMsg(*egressOnlyTrackingTaggedMetric) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	Items() []EgressOnlyTrackingMetricTag
	Add() EgressOnlyTrackingMetricTag
	Append(items ...EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	Set(index int, newObj EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	Clear() EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	clearHolderSlice() EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
	appendHolderSlice(item EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter
}

func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) setMsg(msg *egressOnlyTrackingTaggedMetric) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&egressOnlyTrackingMetricTag{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) Items() []EgressOnlyTrackingMetricTag {
	return obj.egressOnlyTrackingMetricTagSlice
}

func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) Add() EgressOnlyTrackingMetricTag {
	newObj := &otg.EgressOnlyTrackingMetricTag{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &egressOnlyTrackingMetricTag{obj: newObj}
	newLibObj.setDefault()
	obj.egressOnlyTrackingMetricTagSlice = append(obj.egressOnlyTrackingMetricTagSlice, newLibObj)
	return newLibObj
}

func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) Append(items ...EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.egressOnlyTrackingMetricTagSlice = append(obj.egressOnlyTrackingMetricTagSlice, item)
	}
	return obj
}

func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) Set(index int, newObj EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.egressOnlyTrackingMetricTagSlice[index] = newObj
	return obj
}
func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) Clear() EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.EgressOnlyTrackingMetricTag{}
		obj.egressOnlyTrackingMetricTagSlice = []EgressOnlyTrackingMetricTag{}
	}
	return obj
}
func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) clearHolderSlice() EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	if len(obj.egressOnlyTrackingMetricTagSlice) > 0 {
		obj.egressOnlyTrackingMetricTagSlice = []EgressOnlyTrackingMetricTag{}
	}
	return obj
}
func (obj *egressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter) appendHolderSlice(item EgressOnlyTrackingMetricTag) EgressOnlyTrackingTaggedMetricEgressOnlyTrackingMetricTagIter {
	obj.egressOnlyTrackingMetricTagSlice = append(obj.egressOnlyTrackingMetricTagSlice, item)
	return obj
}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *egressOnlyTrackingTaggedMetric) FramesRx() uint64 {

	return *obj.obj.FramesRx

}

// The current total number of valid frames received
// FramesRx returns a uint64
func (obj *egressOnlyTrackingTaggedMetric) HasFramesRx() bool {
	return obj.obj.FramesRx != nil
}

// The current total number of valid frames received
// SetFramesRx sets the uint64 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetFramesRx(value uint64) EgressOnlyTrackingTaggedMetric {

	obj.obj.FramesRx = &value
	return obj
}

// The current total number of bytes received
// BytesRx returns a uint64
func (obj *egressOnlyTrackingTaggedMetric) BytesRx() uint64 {

	return *obj.obj.BytesRx

}

// The current total number of bytes received
// BytesRx returns a uint64
func (obj *egressOnlyTrackingTaggedMetric) HasBytesRx() bool {
	return obj.obj.BytesRx != nil
}

// The current total number of bytes received
// SetBytesRx sets the uint64 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetBytesRx(value uint64) EgressOnlyTrackingTaggedMetric {

	obj.obj.BytesRx = &value
	return obj
}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *egressOnlyTrackingTaggedMetric) FramesRxRate() float32 {

	return *obj.obj.FramesRxRate

}

// The current rate of valid frames received
// FramesRxRate returns a float32
func (obj *egressOnlyTrackingTaggedMetric) HasFramesRxRate() bool {
	return obj.obj.FramesRxRate != nil
}

// The current rate of valid frames received
// SetFramesRxRate sets the float32 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetFramesRxRate(value float32) EgressOnlyTrackingTaggedMetric {

	obj.obj.FramesRxRate = &value
	return obj
}

// description is TBD
// Timestamps returns a EgressOnlyTrackingTimestamp
func (obj *egressOnlyTrackingTaggedMetric) Timestamps() EgressOnlyTrackingTimestamp {
	if obj.obj.Timestamps == nil {
		obj.obj.Timestamps = NewEgressOnlyTrackingTimestamp().msg()
	}
	if obj.timestampsHolder == nil {
		obj.timestampsHolder = &egressOnlyTrackingTimestamp{obj: obj.obj.Timestamps}
	}
	return obj.timestampsHolder
}

// description is TBD
// Timestamps returns a EgressOnlyTrackingTimestamp
func (obj *egressOnlyTrackingTaggedMetric) HasTimestamps() bool {
	return obj.obj.Timestamps != nil
}

// description is TBD
// SetTimestamps sets the EgressOnlyTrackingTimestamp value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetTimestamps(value EgressOnlyTrackingTimestamp) EgressOnlyTrackingTaggedMetric {

	obj.timestampsHolder = nil
	obj.obj.Timestamps = value.msg()

	return obj
}

// The Layer 1 receive rate in bits per second.
// RxL1RateBps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) RxL1RateBps() float32 {

	return *obj.obj.RxL1RateBps

}

// The Layer 1 receive rate in bits per second.
// RxL1RateBps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) HasRxL1RateBps() bool {
	return obj.obj.RxL1RateBps != nil
}

// The Layer 1 receive rate in bits per second.
// SetRxL1RateBps sets the float32 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetRxL1RateBps(value float32) EgressOnlyTrackingTaggedMetric {

	obj.obj.RxL1RateBps = &value
	return obj
}

// The receive rate in bytes per second.
// RxRateBytes returns a float32
func (obj *egressOnlyTrackingTaggedMetric) RxRateBytes() float32 {

	return *obj.obj.RxRateBytes

}

// The receive rate in bytes per second.
// RxRateBytes returns a float32
func (obj *egressOnlyTrackingTaggedMetric) HasRxRateBytes() bool {
	return obj.obj.RxRateBytes != nil
}

// The receive rate in bytes per second.
// SetRxRateBytes sets the float32 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetRxRateBytes(value float32) EgressOnlyTrackingTaggedMetric {

	obj.obj.RxRateBytes = &value
	return obj
}

// The receive rate in bits per second.
// RxRateBps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) RxRateBps() float32 {

	return *obj.obj.RxRateBps

}

// The receive rate in bits per second.
// RxRateBps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) HasRxRateBps() bool {
	return obj.obj.RxRateBps != nil
}

// The receive rate in bits per second.
// SetRxRateBps sets the float32 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetRxRateBps(value float32) EgressOnlyTrackingTaggedMetric {

	obj.obj.RxRateBps = &value
	return obj
}

// The receive rate in Kilobits per second.
// RxRateKbps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) RxRateKbps() float32 {

	return *obj.obj.RxRateKbps

}

// The receive rate in Kilobits per second.
// RxRateKbps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) HasRxRateKbps() bool {
	return obj.obj.RxRateKbps != nil
}

// The receive rate in Kilobits per second.
// SetRxRateKbps sets the float32 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetRxRateKbps(value float32) EgressOnlyTrackingTaggedMetric {

	obj.obj.RxRateKbps = &value
	return obj
}

// The receive rate in Megabits per second.
// RxRateMbps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) RxRateMbps() float32 {

	return *obj.obj.RxRateMbps

}

// The receive rate in Megabits per second.
// RxRateMbps returns a float32
func (obj *egressOnlyTrackingTaggedMetric) HasRxRateMbps() bool {
	return obj.obj.RxRateMbps != nil
}

// The receive rate in Megabits per second.
// SetRxRateMbps sets the float32 value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetRxRateMbps(value float32) EgressOnlyTrackingTaggedMetric {

	obj.obj.RxRateMbps = &value
	return obj
}

// description is TBD
// TxMetrics returns a EgressOnlyTrackingTxMetrics
func (obj *egressOnlyTrackingTaggedMetric) TxMetrics() EgressOnlyTrackingTxMetrics {
	if obj.obj.TxMetrics == nil {
		obj.obj.TxMetrics = NewEgressOnlyTrackingTxMetrics().msg()
	}
	if obj.txMetricsHolder == nil {
		obj.txMetricsHolder = &egressOnlyTrackingTxMetrics{obj: obj.obj.TxMetrics}
	}
	return obj.txMetricsHolder
}

// description is TBD
// TxMetrics returns a EgressOnlyTrackingTxMetrics
func (obj *egressOnlyTrackingTaggedMetric) HasTxMetrics() bool {
	return obj.obj.TxMetrics != nil
}

// description is TBD
// SetTxMetrics sets the EgressOnlyTrackingTxMetrics value in the EgressOnlyTrackingTaggedMetric object
func (obj *egressOnlyTrackingTaggedMetric) SetTxMetrics(value EgressOnlyTrackingTxMetrics) EgressOnlyTrackingTaggedMetric {

	obj.txMetricsHolder = nil
	obj.obj.TxMetrics = value.msg()

	return obj
}

func (obj *egressOnlyTrackingTaggedMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Tags) != 0 {

		if set_default {
			obj.Tags().clearHolderSlice()
			for _, item := range obj.obj.Tags {
				obj.Tags().appendHolderSlice(&egressOnlyTrackingMetricTag{obj: item})
			}
		}
		for _, item := range obj.Tags().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Timestamps != nil {

		obj.Timestamps().validateObj(vObj, set_default)
	}

	if obj.obj.TxMetrics != nil {

		obj.TxMetrics().validateObj(vObj, set_default)
	}

}

func (obj *egressOnlyTrackingTaggedMetric) setDefault() {

}
