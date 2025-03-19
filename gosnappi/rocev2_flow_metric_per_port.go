package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2FlowMetricPerPort *****
type rocev2FlowMetricPerPort struct {
	validation
	obj          *otg.Rocev2FlowMetricPerPort
	marshaller   marshalRocev2FlowMetricPerPort
	unMarshaller unMarshalRocev2FlowMetricPerPort
}

func NewRocev2FlowMetricPerPort() Rocev2FlowMetricPerPort {
	obj := rocev2FlowMetricPerPort{obj: &otg.Rocev2FlowMetricPerPort{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2FlowMetricPerPort) msg() *otg.Rocev2FlowMetricPerPort {
	return obj.obj
}

func (obj *rocev2FlowMetricPerPort) setMsg(msg *otg.Rocev2FlowMetricPerPort) Rocev2FlowMetricPerPort {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2FlowMetricPerPort struct {
	obj *rocev2FlowMetricPerPort
}

type marshalRocev2FlowMetricPerPort interface {
	// ToProto marshals Rocev2FlowMetricPerPort to protobuf object *otg.Rocev2FlowMetricPerPort
	ToProto() (*otg.Rocev2FlowMetricPerPort, error)
	// ToPbText marshals Rocev2FlowMetricPerPort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2FlowMetricPerPort to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2FlowMetricPerPort to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2FlowMetricPerPort struct {
	obj *rocev2FlowMetricPerPort
}

type unMarshalRocev2FlowMetricPerPort interface {
	// FromProto unmarshals Rocev2FlowMetricPerPort from protobuf object *otg.Rocev2FlowMetricPerPort
	FromProto(msg *otg.Rocev2FlowMetricPerPort) (Rocev2FlowMetricPerPort, error)
	// FromPbText unmarshals Rocev2FlowMetricPerPort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2FlowMetricPerPort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2FlowMetricPerPort from JSON text
	FromJson(value string) error
}

func (obj *rocev2FlowMetricPerPort) Marshal() marshalRocev2FlowMetricPerPort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2FlowMetricPerPort{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2FlowMetricPerPort) Unmarshal() unMarshalRocev2FlowMetricPerPort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2FlowMetricPerPort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2FlowMetricPerPort) ToProto() (*otg.Rocev2FlowMetricPerPort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2FlowMetricPerPort) FromProto(msg *otg.Rocev2FlowMetricPerPort) (Rocev2FlowMetricPerPort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2FlowMetricPerPort) ToPbText() (string, error) {
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

func (m *unMarshalrocev2FlowMetricPerPort) FromPbText(value string) error {
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

func (m *marshalrocev2FlowMetricPerPort) ToYaml() (string, error) {
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

func (m *unMarshalrocev2FlowMetricPerPort) FromYaml(value string) error {
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

func (m *marshalrocev2FlowMetricPerPort) ToJson() (string, error) {
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

func (m *unMarshalrocev2FlowMetricPerPort) FromJson(value string) error {
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

func (obj *rocev2FlowMetricPerPort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2FlowMetricPerPort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2FlowMetricPerPort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2FlowMetricPerPort) Clone() (Rocev2FlowMetricPerPort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2FlowMetricPerPort()
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

// Rocev2FlowMetricPerPort is roCEv2 Flow statistics information.
type Rocev2FlowMetricPerPort interface {
	Validation
	// msg marshals Rocev2FlowMetricPerPort to protobuf object *otg.Rocev2FlowMetricPerPort
	// and doesn't set defaults
	msg() *otg.Rocev2FlowMetricPerPort
	// setMsg unmarshals Rocev2FlowMetricPerPort from protobuf object *otg.Rocev2FlowMetricPerPort
	// and doesn't set defaults
	setMsg(*otg.Rocev2FlowMetricPerPort) Rocev2FlowMetricPerPort
	// provides marshal interface
	Marshal() marshalRocev2FlowMetricPerPort
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2FlowMetricPerPort
	// validate validates Rocev2FlowMetricPerPort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2FlowMetricPerPort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowName returns string, set in Rocev2FlowMetricPerPort.
	FlowName() string
	// SetFlowName assigns string provided by user to Rocev2FlowMetricPerPort
	SetFlowName(value string) Rocev2FlowMetricPerPort
	// HasFlowName checks if FlowName has been set in Rocev2FlowMetricPerPort
	HasFlowName() bool
	// TrafficItem returns string, set in Rocev2FlowMetricPerPort.
	TrafficItem() string
	// SetTrafficItem assigns string provided by user to Rocev2FlowMetricPerPort
	SetTrafficItem(value string) Rocev2FlowMetricPerPort
	// HasTrafficItem checks if TrafficItem has been set in Rocev2FlowMetricPerPort
	HasTrafficItem() bool
	// PortTx returns string, set in Rocev2FlowMetricPerPort.
	PortTx() string
	// SetPortTx assigns string provided by user to Rocev2FlowMetricPerPort
	SetPortTx(value string) Rocev2FlowMetricPerPort
	// HasPortTx checks if PortTx has been set in Rocev2FlowMetricPerPort
	HasPortTx() bool
	// PortRx returns string, set in Rocev2FlowMetricPerPort.
	PortRx() string
	// SetPortRx assigns string provided by user to Rocev2FlowMetricPerPort
	SetPortRx(value string) Rocev2FlowMetricPerPort
	// HasPortRx checks if PortRx has been set in Rocev2FlowMetricPerPort
	HasPortRx() bool
	// SrcQp returns uint64, set in Rocev2FlowMetricPerPort.
	SrcQp() uint64
	// SetSrcQp assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetSrcQp(value uint64) Rocev2FlowMetricPerPort
	// HasSrcQp checks if SrcQp has been set in Rocev2FlowMetricPerPort
	HasSrcQp() bool
	// DestQp returns uint64, set in Rocev2FlowMetricPerPort.
	DestQp() uint64
	// SetDestQp assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetDestQp(value uint64) Rocev2FlowMetricPerPort
	// HasDestQp checks if DestQp has been set in Rocev2FlowMetricPerPort
	HasDestQp() bool
	// SrcIpv4 returns string, set in Rocev2FlowMetricPerPort.
	SrcIpv4() string
	// SetSrcIpv4 assigns string provided by user to Rocev2FlowMetricPerPort
	SetSrcIpv4(value string) Rocev2FlowMetricPerPort
	// HasSrcIpv4 checks if SrcIpv4 has been set in Rocev2FlowMetricPerPort
	HasSrcIpv4() bool
	// DestIpv4 returns string, set in Rocev2FlowMetricPerPort.
	DestIpv4() string
	// SetDestIpv4 assigns string provided by user to Rocev2FlowMetricPerPort
	SetDestIpv4(value string) Rocev2FlowMetricPerPort
	// HasDestIpv4 checks if DestIpv4 has been set in Rocev2FlowMetricPerPort
	HasDestIpv4() bool
	// DataFrameTx returns uint64, set in Rocev2FlowMetricPerPort.
	DataFrameTx() uint64
	// SetDataFrameTx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetDataFrameTx(value uint64) Rocev2FlowMetricPerPort
	// HasDataFrameTx checks if DataFrameTx has been set in Rocev2FlowMetricPerPort
	HasDataFrameTx() bool
	// DataFramesRx returns uint64, set in Rocev2FlowMetricPerPort.
	DataFramesRx() uint64
	// SetDataFramesRx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetDataFramesRx(value uint64) Rocev2FlowMetricPerPort
	// HasDataFramesRx checks if DataFramesRx has been set in Rocev2FlowMetricPerPort
	HasDataFramesRx() bool
	// FrameDelta returns uint64, set in Rocev2FlowMetricPerPort.
	FrameDelta() uint64
	// SetFrameDelta assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetFrameDelta(value uint64) Rocev2FlowMetricPerPort
	// HasFrameDelta checks if FrameDelta has been set in Rocev2FlowMetricPerPort
	HasFrameDelta() bool
	// DataFramesRetransmitted returns uint64, set in Rocev2FlowMetricPerPort.
	DataFramesRetransmitted() uint64
	// SetDataFramesRetransmitted assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetDataFramesRetransmitted(value uint64) Rocev2FlowMetricPerPort
	// HasDataFramesRetransmitted checks if DataFramesRetransmitted has been set in Rocev2FlowMetricPerPort
	HasDataFramesRetransmitted() bool
	// FrameSequenceError returns uint64, set in Rocev2FlowMetricPerPort.
	FrameSequenceError() uint64
	// SetFrameSequenceError assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetFrameSequenceError(value uint64) Rocev2FlowMetricPerPort
	// HasFrameSequenceError checks if FrameSequenceError has been set in Rocev2FlowMetricPerPort
	HasFrameSequenceError() bool
	// TxBytes returns uint64, set in Rocev2FlowMetricPerPort.
	TxBytes() uint64
	// SetTxBytes assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetTxBytes(value uint64) Rocev2FlowMetricPerPort
	// HasTxBytes checks if TxBytes has been set in Rocev2FlowMetricPerPort
	HasTxBytes() bool
	// RxBytes returns uint64, set in Rocev2FlowMetricPerPort.
	RxBytes() uint64
	// SetRxBytes assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetRxBytes(value uint64) Rocev2FlowMetricPerPort
	// HasRxBytes checks if RxBytes has been set in Rocev2FlowMetricPerPort
	HasRxBytes() bool
	// DataTxRate returns uint64, set in Rocev2FlowMetricPerPort.
	DataTxRate() uint64
	// SetDataTxRate assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetDataTxRate(value uint64) Rocev2FlowMetricPerPort
	// HasDataTxRate checks if DataTxRate has been set in Rocev2FlowMetricPerPort
	HasDataTxRate() bool
	// DataRxRate returns uint64, set in Rocev2FlowMetricPerPort.
	DataRxRate() uint64
	// SetDataRxRate assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetDataRxRate(value uint64) Rocev2FlowMetricPerPort
	// HasDataRxRate checks if DataRxRate has been set in Rocev2FlowMetricPerPort
	HasDataRxRate() bool
	// MessageTx returns uint64, set in Rocev2FlowMetricPerPort.
	MessageTx() uint64
	// SetMessageTx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetMessageTx(value uint64) Rocev2FlowMetricPerPort
	// HasMessageTx checks if MessageTx has been set in Rocev2FlowMetricPerPort
	HasMessageTx() bool
	// MessageCompleteRx returns uint64, set in Rocev2FlowMetricPerPort.
	MessageCompleteRx() uint64
	// SetMessageCompleteRx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetMessageCompleteRx(value uint64) Rocev2FlowMetricPerPort
	// HasMessageCompleteRx checks if MessageCompleteRx has been set in Rocev2FlowMetricPerPort
	HasMessageCompleteRx() bool
	// MessageFail returns uint64, set in Rocev2FlowMetricPerPort.
	MessageFail() uint64
	// SetMessageFail assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetMessageFail(value uint64) Rocev2FlowMetricPerPort
	// HasMessageFail checks if MessageFail has been set in Rocev2FlowMetricPerPort
	HasMessageFail() bool
	// FlowCompletionTime returns uint64, set in Rocev2FlowMetricPerPort.
	FlowCompletionTime() uint64
	// SetFlowCompletionTime assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetFlowCompletionTime(value uint64) Rocev2FlowMetricPerPort
	// HasFlowCompletionTime checks if FlowCompletionTime has been set in Rocev2FlowMetricPerPort
	HasFlowCompletionTime() bool
	// AvgLatency returns uint64, set in Rocev2FlowMetricPerPort.
	AvgLatency() uint64
	// SetAvgLatency assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetAvgLatency(value uint64) Rocev2FlowMetricPerPort
	// HasAvgLatency checks if AvgLatency has been set in Rocev2FlowMetricPerPort
	HasAvgLatency() bool
	// MinLatency returns uint64, set in Rocev2FlowMetricPerPort.
	MinLatency() uint64
	// SetMinLatency assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetMinLatency(value uint64) Rocev2FlowMetricPerPort
	// HasMinLatency checks if MinLatency has been set in Rocev2FlowMetricPerPort
	HasMinLatency() bool
	// MaxLatency returns uint64, set in Rocev2FlowMetricPerPort.
	MaxLatency() uint64
	// SetMaxLatency assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetMaxLatency(value uint64) Rocev2FlowMetricPerPort
	// HasMaxLatency checks if MaxLatency has been set in Rocev2FlowMetricPerPort
	HasMaxLatency() bool
	// EcnCeRx returns uint64, set in Rocev2FlowMetricPerPort.
	EcnCeRx() uint64
	// SetEcnCeRx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetEcnCeRx(value uint64) Rocev2FlowMetricPerPort
	// HasEcnCeRx checks if EcnCeRx has been set in Rocev2FlowMetricPerPort
	HasEcnCeRx() bool
	// CnpTx returns uint64, set in Rocev2FlowMetricPerPort.
	CnpTx() uint64
	// SetCnpTx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetCnpTx(value uint64) Rocev2FlowMetricPerPort
	// HasCnpTx checks if CnpTx has been set in Rocev2FlowMetricPerPort
	HasCnpTx() bool
	// CnpRx returns uint64, set in Rocev2FlowMetricPerPort.
	CnpRx() uint64
	// SetCnpRx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetCnpRx(value uint64) Rocev2FlowMetricPerPort
	// HasCnpRx checks if CnpRx has been set in Rocev2FlowMetricPerPort
	HasCnpRx() bool
	// AckTx returns uint64, set in Rocev2FlowMetricPerPort.
	AckTx() uint64
	// SetAckTx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetAckTx(value uint64) Rocev2FlowMetricPerPort
	// HasAckTx checks if AckTx has been set in Rocev2FlowMetricPerPort
	HasAckTx() bool
	// AckRx returns uint64, set in Rocev2FlowMetricPerPort.
	AckRx() uint64
	// SetAckRx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetAckRx(value uint64) Rocev2FlowMetricPerPort
	// HasAckRx checks if AckRx has been set in Rocev2FlowMetricPerPort
	HasAckRx() bool
	// NakTx returns uint64, set in Rocev2FlowMetricPerPort.
	NakTx() uint64
	// SetNakTx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetNakTx(value uint64) Rocev2FlowMetricPerPort
	// HasNakTx checks if NakTx has been set in Rocev2FlowMetricPerPort
	HasNakTx() bool
	// NakRx returns uint64, set in Rocev2FlowMetricPerPort.
	NakRx() uint64
	// SetNakRx assigns uint64 provided by user to Rocev2FlowMetricPerPort
	SetNakRx(value uint64) Rocev2FlowMetricPerPort
	// HasNakRx checks if NakRx has been set in Rocev2FlowMetricPerPort
	HasNakRx() bool
	// FirstTimestamp returns string, set in Rocev2FlowMetricPerPort.
	FirstTimestamp() string
	// SetFirstTimestamp assigns string provided by user to Rocev2FlowMetricPerPort
	SetFirstTimestamp(value string) Rocev2FlowMetricPerPort
	// HasFirstTimestamp checks if FirstTimestamp has been set in Rocev2FlowMetricPerPort
	HasFirstTimestamp() bool
	// LastTimestamp returns string, set in Rocev2FlowMetricPerPort.
	LastTimestamp() string
	// SetLastTimestamp assigns string provided by user to Rocev2FlowMetricPerPort
	SetLastTimestamp(value string) Rocev2FlowMetricPerPort
	// HasLastTimestamp checks if LastTimestamp has been set in Rocev2FlowMetricPerPort
	HasLastTimestamp() bool
}

// Flow Name.
// FlowName returns a string
func (obj *rocev2FlowMetricPerPort) FlowName() string {

	return *obj.obj.FlowName

}

// Flow Name.
// FlowName returns a string
func (obj *rocev2FlowMetricPerPort) HasFlowName() bool {
	return obj.obj.FlowName != nil
}

// Flow Name.
// SetFlowName sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetFlowName(value string) Rocev2FlowMetricPerPort {

	obj.obj.FlowName = &value
	return obj
}

// Traffic Item Name.
// TrafficItem returns a string
func (obj *rocev2FlowMetricPerPort) TrafficItem() string {

	return *obj.obj.TrafficItem

}

// Traffic Item Name.
// TrafficItem returns a string
func (obj *rocev2FlowMetricPerPort) HasTrafficItem() bool {
	return obj.obj.TrafficItem != nil
}

// Traffic Item Name.
// SetTrafficItem sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetTrafficItem(value string) Rocev2FlowMetricPerPort {

	obj.obj.TrafficItem = &value
	return obj
}

// The name of the transmit port
// PortTx returns a string
func (obj *rocev2FlowMetricPerPort) PortTx() string {

	return *obj.obj.PortTx

}

// The name of the transmit port
// PortTx returns a string
func (obj *rocev2FlowMetricPerPort) HasPortTx() bool {
	return obj.obj.PortTx != nil
}

// The name of the transmit port
// SetPortTx sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetPortTx(value string) Rocev2FlowMetricPerPort {

	obj.obj.PortTx = &value
	return obj
}

// The name of the receive port
// PortRx returns a string
func (obj *rocev2FlowMetricPerPort) PortRx() string {

	return *obj.obj.PortRx

}

// The name of the receive port
// PortRx returns a string
func (obj *rocev2FlowMetricPerPort) HasPortRx() bool {
	return obj.obj.PortRx != nil
}

// The name of the receive port
// SetPortRx sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetPortRx(value string) Rocev2FlowMetricPerPort {

	obj.obj.PortRx = &value
	return obj
}

// Current Source QP number.
// SrcQp returns a uint64
func (obj *rocev2FlowMetricPerPort) SrcQp() uint64 {

	return *obj.obj.SrcQp

}

// Current Source QP number.
// SrcQp returns a uint64
func (obj *rocev2FlowMetricPerPort) HasSrcQp() bool {
	return obj.obj.SrcQp != nil
}

// Current Source QP number.
// SetSrcQp sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetSrcQp(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.SrcQp = &value
	return obj
}

// Current destination QP number.
// DestQp returns a uint64
func (obj *rocev2FlowMetricPerPort) DestQp() uint64 {

	return *obj.obj.DestQp

}

// Current destination QP number.
// DestQp returns a uint64
func (obj *rocev2FlowMetricPerPort) HasDestQp() bool {
	return obj.obj.DestQp != nil
}

// Current destination QP number.
// SetDestQp sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDestQp(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.DestQp = &value
	return obj
}

// Current source address.
// SrcIpv4 returns a string
func (obj *rocev2FlowMetricPerPort) SrcIpv4() string {

	return *obj.obj.SrcIpv4

}

// Current source address.
// SrcIpv4 returns a string
func (obj *rocev2FlowMetricPerPort) HasSrcIpv4() bool {
	return obj.obj.SrcIpv4 != nil
}

// Current source address.
// SetSrcIpv4 sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetSrcIpv4(value string) Rocev2FlowMetricPerPort {

	obj.obj.SrcIpv4 = &value
	return obj
}

// Current destination address.
// DestIpv4 returns a string
func (obj *rocev2FlowMetricPerPort) DestIpv4() string {

	return *obj.obj.DestIpv4

}

// Current destination address.
// DestIpv4 returns a string
func (obj *rocev2FlowMetricPerPort) HasDestIpv4() bool {
	return obj.obj.DestIpv4 != nil
}

// Current destination address.
// SetDestIpv4 sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDestIpv4(value string) Rocev2FlowMetricPerPort {

	obj.obj.DestIpv4 = &value
	return obj
}

// Current number of data frames transmitted.
// DataFrameTx returns a uint64
func (obj *rocev2FlowMetricPerPort) DataFrameTx() uint64 {

	return *obj.obj.DataFrameTx

}

// Current number of data frames transmitted.
// DataFrameTx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasDataFrameTx() bool {
	return obj.obj.DataFrameTx != nil
}

// Current number of data frames transmitted.
// SetDataFrameTx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDataFrameTx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.DataFrameTx = &value
	return obj
}

// Current number of data frames received.
// DataFramesRx returns a uint64
func (obj *rocev2FlowMetricPerPort) DataFramesRx() uint64 {

	return *obj.obj.DataFramesRx

}

// Current number of data frames received.
// DataFramesRx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasDataFramesRx() bool {
	return obj.obj.DataFramesRx != nil
}

// Current number of data frames received.
// SetDataFramesRx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDataFramesRx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.DataFramesRx = &value
	return obj
}

// Current differnece between tx and rx data frames
// FrameDelta returns a uint64
func (obj *rocev2FlowMetricPerPort) FrameDelta() uint64 {

	return *obj.obj.FrameDelta

}

// Current differnece between tx and rx data frames
// FrameDelta returns a uint64
func (obj *rocev2FlowMetricPerPort) HasFrameDelta() bool {
	return obj.obj.FrameDelta != nil
}

// Current differnece between tx and rx data frames
// SetFrameDelta sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetFrameDelta(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.FrameDelta = &value
	return obj
}

// Current number of data frames re-transmitted.
// DataFramesRetransmitted returns a uint64
func (obj *rocev2FlowMetricPerPort) DataFramesRetransmitted() uint64 {

	return *obj.obj.DataFramesRetransmitted

}

// Current number of data frames re-transmitted.
// DataFramesRetransmitted returns a uint64
func (obj *rocev2FlowMetricPerPort) HasDataFramesRetransmitted() bool {
	return obj.obj.DataFramesRetransmitted != nil
}

// Current number of data frames re-transmitted.
// SetDataFramesRetransmitted sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDataFramesRetransmitted(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.DataFramesRetransmitted = &value
	return obj
}

// Current number of frame sequence errors.
// FrameSequenceError returns a uint64
func (obj *rocev2FlowMetricPerPort) FrameSequenceError() uint64 {

	return *obj.obj.FrameSequenceError

}

// Current number of frame sequence errors.
// FrameSequenceError returns a uint64
func (obj *rocev2FlowMetricPerPort) HasFrameSequenceError() bool {
	return obj.obj.FrameSequenceError != nil
}

// Current number of frame sequence errors.
// SetFrameSequenceError sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetFrameSequenceError(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.FrameSequenceError = &value
	return obj
}

// Current number of bytes transmitted.
// TxBytes returns a uint64
func (obj *rocev2FlowMetricPerPort) TxBytes() uint64 {

	return *obj.obj.TxBytes

}

// Current number of bytes transmitted.
// TxBytes returns a uint64
func (obj *rocev2FlowMetricPerPort) HasTxBytes() bool {
	return obj.obj.TxBytes != nil
}

// Current number of bytes transmitted.
// SetTxBytes sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetTxBytes(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.TxBytes = &value
	return obj
}

// Current number of bytes received.
// RxBytes returns a uint64
func (obj *rocev2FlowMetricPerPort) RxBytes() uint64 {

	return *obj.obj.RxBytes

}

// Current number of bytes received.
// RxBytes returns a uint64
func (obj *rocev2FlowMetricPerPort) HasRxBytes() bool {
	return obj.obj.RxBytes != nil
}

// Current number of bytes received.
// SetRxBytes sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetRxBytes(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.RxBytes = &value
	return obj
}

// Current rate at which data is transmitted in Gbps.
// DataTxRate returns a uint64
func (obj *rocev2FlowMetricPerPort) DataTxRate() uint64 {

	return *obj.obj.DataTxRate

}

// Current rate at which data is transmitted in Gbps.
// DataTxRate returns a uint64
func (obj *rocev2FlowMetricPerPort) HasDataTxRate() bool {
	return obj.obj.DataTxRate != nil
}

// Current rate at which data is transmitted in Gbps.
// SetDataTxRate sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDataTxRate(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.DataTxRate = &value
	return obj
}

// Current rate at which data is received in Gbps.
// DataRxRate returns a uint64
func (obj *rocev2FlowMetricPerPort) DataRxRate() uint64 {

	return *obj.obj.DataRxRate

}

// Current rate at which data is received in Gbps.
// DataRxRate returns a uint64
func (obj *rocev2FlowMetricPerPort) HasDataRxRate() bool {
	return obj.obj.DataRxRate != nil
}

// Current rate at which data is received in Gbps.
// SetDataRxRate sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetDataRxRate(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.DataRxRate = &value
	return obj
}

// Current number of Message transmitted.
// MessageTx returns a uint64
func (obj *rocev2FlowMetricPerPort) MessageTx() uint64 {

	return *obj.obj.MessageTx

}

// Current number of Message transmitted.
// MessageTx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasMessageTx() bool {
	return obj.obj.MessageTx != nil
}

// Current number of Message transmitted.
// SetMessageTx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetMessageTx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.MessageTx = &value
	return obj
}

// Current number of Message Complete received.
// MessageCompleteRx returns a uint64
func (obj *rocev2FlowMetricPerPort) MessageCompleteRx() uint64 {

	return *obj.obj.MessageCompleteRx

}

// Current number of Message Complete received.
// MessageCompleteRx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasMessageCompleteRx() bool {
	return obj.obj.MessageCompleteRx != nil
}

// Current number of Message Complete received.
// SetMessageCompleteRx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetMessageCompleteRx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.MessageCompleteRx = &value
	return obj
}

// Current number of Message Fail count.
// MessageFail returns a uint64
func (obj *rocev2FlowMetricPerPort) MessageFail() uint64 {

	return *obj.obj.MessageFail

}

// Current number of Message Fail count.
// MessageFail returns a uint64
func (obj *rocev2FlowMetricPerPort) HasMessageFail() bool {
	return obj.obj.MessageFail != nil
}

// Current number of Message Fail count.
// SetMessageFail sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetMessageFail(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.MessageFail = &value
	return obj
}

// Current flow comletion time in ms.
// FlowCompletionTime returns a uint64
func (obj *rocev2FlowMetricPerPort) FlowCompletionTime() uint64 {

	return *obj.obj.FlowCompletionTime

}

// Current flow comletion time in ms.
// FlowCompletionTime returns a uint64
func (obj *rocev2FlowMetricPerPort) HasFlowCompletionTime() bool {
	return obj.obj.FlowCompletionTime != nil
}

// Current flow comletion time in ms.
// SetFlowCompletionTime sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetFlowCompletionTime(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.FlowCompletionTime = &value
	return obj
}

// Current average latency measured in ns.
// AvgLatency returns a uint64
func (obj *rocev2FlowMetricPerPort) AvgLatency() uint64 {

	return *obj.obj.AvgLatency

}

// Current average latency measured in ns.
// AvgLatency returns a uint64
func (obj *rocev2FlowMetricPerPort) HasAvgLatency() bool {
	return obj.obj.AvgLatency != nil
}

// Current average latency measured in ns.
// SetAvgLatency sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetAvgLatency(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.AvgLatency = &value
	return obj
}

// Current minimum latency measured in ns.
// MinLatency returns a uint64
func (obj *rocev2FlowMetricPerPort) MinLatency() uint64 {

	return *obj.obj.MinLatency

}

// Current minimum latency measured in ns.
// MinLatency returns a uint64
func (obj *rocev2FlowMetricPerPort) HasMinLatency() bool {
	return obj.obj.MinLatency != nil
}

// Current minimum latency measured in ns.
// SetMinLatency sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetMinLatency(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.MinLatency = &value
	return obj
}

// Current maximum latency measured in ns.
// MaxLatency returns a uint64
func (obj *rocev2FlowMetricPerPort) MaxLatency() uint64 {

	return *obj.obj.MaxLatency

}

// Current maximum latency measured in ns.
// MaxLatency returns a uint64
func (obj *rocev2FlowMetricPerPort) HasMaxLatency() bool {
	return obj.obj.MaxLatency != nil
}

// Current maximum latency measured in ns.
// SetMaxLatency sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetMaxLatency(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.MaxLatency = &value
	return obj
}

// Current number of ECN-CE Recevied.
// EcnCeRx returns a uint64
func (obj *rocev2FlowMetricPerPort) EcnCeRx() uint64 {

	return *obj.obj.EcnCeRx

}

// Current number of ECN-CE Recevied.
// EcnCeRx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasEcnCeRx() bool {
	return obj.obj.EcnCeRx != nil
}

// Current number of ECN-CE Recevied.
// SetEcnCeRx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetEcnCeRx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.EcnCeRx = &value
	return obj
}

// Current number of CNP transmitted.
// CnpTx returns a uint64
func (obj *rocev2FlowMetricPerPort) CnpTx() uint64 {

	return *obj.obj.CnpTx

}

// Current number of CNP transmitted.
// CnpTx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasCnpTx() bool {
	return obj.obj.CnpTx != nil
}

// Current number of CNP transmitted.
// SetCnpTx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetCnpTx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.CnpTx = &value
	return obj
}

// Current number of CNP received.
// CnpRx returns a uint64
func (obj *rocev2FlowMetricPerPort) CnpRx() uint64 {

	return *obj.obj.CnpRx

}

// Current number of CNP received.
// CnpRx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasCnpRx() bool {
	return obj.obj.CnpRx != nil
}

// Current number of CNP received.
// SetCnpRx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetCnpRx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.CnpRx = &value
	return obj
}

// Current number of ACK transmitted.
// AckTx returns a uint64
func (obj *rocev2FlowMetricPerPort) AckTx() uint64 {

	return *obj.obj.AckTx

}

// Current number of ACK transmitted.
// AckTx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasAckTx() bool {
	return obj.obj.AckTx != nil
}

// Current number of ACK transmitted.
// SetAckTx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetAckTx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.AckTx = &value
	return obj
}

// Current number of ACK received.
// AckRx returns a uint64
func (obj *rocev2FlowMetricPerPort) AckRx() uint64 {

	return *obj.obj.AckRx

}

// Current number of ACK received.
// AckRx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasAckRx() bool {
	return obj.obj.AckRx != nil
}

// Current number of ACK received.
// SetAckRx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetAckRx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.AckRx = &value
	return obj
}

// Current number of NAK transmitted.
// NakTx returns a uint64
func (obj *rocev2FlowMetricPerPort) NakTx() uint64 {

	return *obj.obj.NakTx

}

// Current number of NAK transmitted.
// NakTx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasNakTx() bool {
	return obj.obj.NakTx != nil
}

// Current number of NAK transmitted.
// SetNakTx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetNakTx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.NakTx = &value
	return obj
}

// Current number of NAK received.
// NakRx returns a uint64
func (obj *rocev2FlowMetricPerPort) NakRx() uint64 {

	return *obj.obj.NakRx

}

// Current number of NAK received.
// NakRx returns a uint64
func (obj *rocev2FlowMetricPerPort) HasNakRx() bool {
	return obj.obj.NakRx != nil
}

// Current number of NAK received.
// SetNakRx sets the uint64 value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetNakRx(value uint64) Rocev2FlowMetricPerPort {

	obj.obj.NakRx = &value
	return obj
}

// First Timestamp.
// FirstTimestamp returns a string
func (obj *rocev2FlowMetricPerPort) FirstTimestamp() string {

	return *obj.obj.FirstTimestamp

}

// First Timestamp.
// FirstTimestamp returns a string
func (obj *rocev2FlowMetricPerPort) HasFirstTimestamp() bool {
	return obj.obj.FirstTimestamp != nil
}

// First Timestamp.
// SetFirstTimestamp sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetFirstTimestamp(value string) Rocev2FlowMetricPerPort {

	obj.obj.FirstTimestamp = &value
	return obj
}

// Last Timestamp.
// LastTimestamp returns a string
func (obj *rocev2FlowMetricPerPort) LastTimestamp() string {

	return *obj.obj.LastTimestamp

}

// Last Timestamp.
// LastTimestamp returns a string
func (obj *rocev2FlowMetricPerPort) HasLastTimestamp() bool {
	return obj.obj.LastTimestamp != nil
}

// Last Timestamp.
// SetLastTimestamp sets the string value in the Rocev2FlowMetricPerPort object
func (obj *rocev2FlowMetricPerPort) SetLastTimestamp(value string) Rocev2FlowMetricPerPort {

	obj.obj.LastTimestamp = &value
	return obj
}

func (obj *rocev2FlowMetricPerPort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrcIpv4 != nil {

		err := obj.validateIpv4(obj.SrcIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2FlowMetricPerPort.SrcIpv4"))
		}

	}

	if obj.obj.DestIpv4 != nil {

		err := obj.validateIpv4(obj.DestIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2FlowMetricPerPort.DestIpv4"))
		}

	}

}

func (obj *rocev2FlowMetricPerPort) setDefault() {

}
