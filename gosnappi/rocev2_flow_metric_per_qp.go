package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2FlowMetricPerQP *****
type rocev2FlowMetricPerQP struct {
	validation
	obj          *otg.Rocev2FlowMetricPerQP
	marshaller   marshalRocev2FlowMetricPerQP
	unMarshaller unMarshalRocev2FlowMetricPerQP
}

func NewRocev2FlowMetricPerQP() Rocev2FlowMetricPerQP {
	obj := rocev2FlowMetricPerQP{obj: &otg.Rocev2FlowMetricPerQP{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2FlowMetricPerQP) msg() *otg.Rocev2FlowMetricPerQP {
	return obj.obj
}

func (obj *rocev2FlowMetricPerQP) setMsg(msg *otg.Rocev2FlowMetricPerQP) Rocev2FlowMetricPerQP {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2FlowMetricPerQP struct {
	obj *rocev2FlowMetricPerQP
}

type marshalRocev2FlowMetricPerQP interface {
	// ToProto marshals Rocev2FlowMetricPerQP to protobuf object *otg.Rocev2FlowMetricPerQP
	ToProto() (*otg.Rocev2FlowMetricPerQP, error)
	// ToPbText marshals Rocev2FlowMetricPerQP to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2FlowMetricPerQP to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2FlowMetricPerQP to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2FlowMetricPerQP struct {
	obj *rocev2FlowMetricPerQP
}

type unMarshalRocev2FlowMetricPerQP interface {
	// FromProto unmarshals Rocev2FlowMetricPerQP from protobuf object *otg.Rocev2FlowMetricPerQP
	FromProto(msg *otg.Rocev2FlowMetricPerQP) (Rocev2FlowMetricPerQP, error)
	// FromPbText unmarshals Rocev2FlowMetricPerQP from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2FlowMetricPerQP from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2FlowMetricPerQP from JSON text
	FromJson(value string) error
}

func (obj *rocev2FlowMetricPerQP) Marshal() marshalRocev2FlowMetricPerQP {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2FlowMetricPerQP{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2FlowMetricPerQP) Unmarshal() unMarshalRocev2FlowMetricPerQP {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2FlowMetricPerQP{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2FlowMetricPerQP) ToProto() (*otg.Rocev2FlowMetricPerQP, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2FlowMetricPerQP) FromProto(msg *otg.Rocev2FlowMetricPerQP) (Rocev2FlowMetricPerQP, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2FlowMetricPerQP) ToPbText() (string, error) {
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

func (m *unMarshalrocev2FlowMetricPerQP) FromPbText(value string) error {
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

func (m *marshalrocev2FlowMetricPerQP) ToYaml() (string, error) {
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

func (m *unMarshalrocev2FlowMetricPerQP) FromYaml(value string) error {
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

func (m *marshalrocev2FlowMetricPerQP) ToJson() (string, error) {
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

func (m *unMarshalrocev2FlowMetricPerQP) FromJson(value string) error {
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

func (obj *rocev2FlowMetricPerQP) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2FlowMetricPerQP) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2FlowMetricPerQP) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2FlowMetricPerQP) Clone() (Rocev2FlowMetricPerQP, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2FlowMetricPerQP()
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

// Rocev2FlowMetricPerQP is roCEv2 Flow statistics information.
type Rocev2FlowMetricPerQP interface {
	Validation
	// msg marshals Rocev2FlowMetricPerQP to protobuf object *otg.Rocev2FlowMetricPerQP
	// and doesn't set defaults
	msg() *otg.Rocev2FlowMetricPerQP
	// setMsg unmarshals Rocev2FlowMetricPerQP from protobuf object *otg.Rocev2FlowMetricPerQP
	// and doesn't set defaults
	setMsg(*otg.Rocev2FlowMetricPerQP) Rocev2FlowMetricPerQP
	// provides marshal interface
	Marshal() marshalRocev2FlowMetricPerQP
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2FlowMetricPerQP
	// validate validates Rocev2FlowMetricPerQP
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2FlowMetricPerQP, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowName returns string, set in Rocev2FlowMetricPerQP.
	FlowName() string
	// SetFlowName assigns string provided by user to Rocev2FlowMetricPerQP
	SetFlowName(value string) Rocev2FlowMetricPerQP
	// HasFlowName checks if FlowName has been set in Rocev2FlowMetricPerQP
	HasFlowName() bool
	// PortTx returns string, set in Rocev2FlowMetricPerQP.
	PortTx() string
	// SetPortTx assigns string provided by user to Rocev2FlowMetricPerQP
	SetPortTx(value string) Rocev2FlowMetricPerQP
	// HasPortTx checks if PortTx has been set in Rocev2FlowMetricPerQP
	HasPortTx() bool
	// PortRx returns string, set in Rocev2FlowMetricPerQP.
	PortRx() string
	// SetPortRx assigns string provided by user to Rocev2FlowMetricPerQP
	SetPortRx(value string) Rocev2FlowMetricPerQP
	// HasPortRx checks if PortRx has been set in Rocev2FlowMetricPerQP
	HasPortRx() bool
	// SrcQp returns uint64, set in Rocev2FlowMetricPerQP.
	SrcQp() uint64
	// SetSrcQp assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetSrcQp(value uint64) Rocev2FlowMetricPerQP
	// HasSrcQp checks if SrcQp has been set in Rocev2FlowMetricPerQP
	HasSrcQp() bool
	// DestQp returns uint64, set in Rocev2FlowMetricPerQP.
	DestQp() uint64
	// SetDestQp assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetDestQp(value uint64) Rocev2FlowMetricPerQP
	// HasDestQp checks if DestQp has been set in Rocev2FlowMetricPerQP
	HasDestQp() bool
	// SrcIpv4 returns string, set in Rocev2FlowMetricPerQP.
	SrcIpv4() string
	// SetSrcIpv4 assigns string provided by user to Rocev2FlowMetricPerQP
	SetSrcIpv4(value string) Rocev2FlowMetricPerQP
	// HasSrcIpv4 checks if SrcIpv4 has been set in Rocev2FlowMetricPerQP
	HasSrcIpv4() bool
	// DestIpv4 returns string, set in Rocev2FlowMetricPerQP.
	DestIpv4() string
	// SetDestIpv4 assigns string provided by user to Rocev2FlowMetricPerQP
	SetDestIpv4(value string) Rocev2FlowMetricPerQP
	// HasDestIpv4 checks if DestIpv4 has been set in Rocev2FlowMetricPerQP
	HasDestIpv4() bool
	// DataFramesTx returns uint64, set in Rocev2FlowMetricPerQP.
	DataFramesTx() uint64
	// SetDataFramesTx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetDataFramesTx(value uint64) Rocev2FlowMetricPerQP
	// HasDataFramesTx checks if DataFramesTx has been set in Rocev2FlowMetricPerQP
	HasDataFramesTx() bool
	// DataFramesRx returns uint64, set in Rocev2FlowMetricPerQP.
	DataFramesRx() uint64
	// SetDataFramesRx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetDataFramesRx(value uint64) Rocev2FlowMetricPerQP
	// HasDataFramesRx checks if DataFramesRx has been set in Rocev2FlowMetricPerQP
	HasDataFramesRx() bool
	// FrameDelta returns uint64, set in Rocev2FlowMetricPerQP.
	FrameDelta() uint64
	// SetFrameDelta assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetFrameDelta(value uint64) Rocev2FlowMetricPerQP
	// HasFrameDelta checks if FrameDelta has been set in Rocev2FlowMetricPerQP
	HasFrameDelta() bool
	// DataFramesRetransmitted returns uint64, set in Rocev2FlowMetricPerQP.
	DataFramesRetransmitted() uint64
	// SetDataFramesRetransmitted assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetDataFramesRetransmitted(value uint64) Rocev2FlowMetricPerQP
	// HasDataFramesRetransmitted checks if DataFramesRetransmitted has been set in Rocev2FlowMetricPerQP
	HasDataFramesRetransmitted() bool
	// FrameSequenceError returns uint64, set in Rocev2FlowMetricPerQP.
	FrameSequenceError() uint64
	// SetFrameSequenceError assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetFrameSequenceError(value uint64) Rocev2FlowMetricPerQP
	// HasFrameSequenceError checks if FrameSequenceError has been set in Rocev2FlowMetricPerQP
	HasFrameSequenceError() bool
	// TxBytes returns uint64, set in Rocev2FlowMetricPerQP.
	TxBytes() uint64
	// SetTxBytes assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetTxBytes(value uint64) Rocev2FlowMetricPerQP
	// HasTxBytes checks if TxBytes has been set in Rocev2FlowMetricPerQP
	HasTxBytes() bool
	// RxBytes returns uint64, set in Rocev2FlowMetricPerQP.
	RxBytes() uint64
	// SetRxBytes assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetRxBytes(value uint64) Rocev2FlowMetricPerQP
	// HasRxBytes checks if RxBytes has been set in Rocev2FlowMetricPerQP
	HasRxBytes() bool
	// DataTxRate returns uint64, set in Rocev2FlowMetricPerQP.
	DataTxRate() uint64
	// SetDataTxRate assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetDataTxRate(value uint64) Rocev2FlowMetricPerQP
	// HasDataTxRate checks if DataTxRate has been set in Rocev2FlowMetricPerQP
	HasDataTxRate() bool
	// DataRxRate returns uint64, set in Rocev2FlowMetricPerQP.
	DataRxRate() uint64
	// SetDataRxRate assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetDataRxRate(value uint64) Rocev2FlowMetricPerQP
	// HasDataRxRate checks if DataRxRate has been set in Rocev2FlowMetricPerQP
	HasDataRxRate() bool
	// MessageTx returns uint64, set in Rocev2FlowMetricPerQP.
	MessageTx() uint64
	// SetMessageTx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetMessageTx(value uint64) Rocev2FlowMetricPerQP
	// HasMessageTx checks if MessageTx has been set in Rocev2FlowMetricPerQP
	HasMessageTx() bool
	// MessageCompleteRx returns uint64, set in Rocev2FlowMetricPerQP.
	MessageCompleteRx() uint64
	// SetMessageCompleteRx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetMessageCompleteRx(value uint64) Rocev2FlowMetricPerQP
	// HasMessageCompleteRx checks if MessageCompleteRx has been set in Rocev2FlowMetricPerQP
	HasMessageCompleteRx() bool
	// MessageFail returns uint64, set in Rocev2FlowMetricPerQP.
	MessageFail() uint64
	// SetMessageFail assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetMessageFail(value uint64) Rocev2FlowMetricPerQP
	// HasMessageFail checks if MessageFail has been set in Rocev2FlowMetricPerQP
	HasMessageFail() bool
	// FlowCompletionTime returns uint64, set in Rocev2FlowMetricPerQP.
	FlowCompletionTime() uint64
	// SetFlowCompletionTime assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetFlowCompletionTime(value uint64) Rocev2FlowMetricPerQP
	// HasFlowCompletionTime checks if FlowCompletionTime has been set in Rocev2FlowMetricPerQP
	HasFlowCompletionTime() bool
	// AvgLatency returns uint64, set in Rocev2FlowMetricPerQP.
	AvgLatency() uint64
	// SetAvgLatency assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetAvgLatency(value uint64) Rocev2FlowMetricPerQP
	// HasAvgLatency checks if AvgLatency has been set in Rocev2FlowMetricPerQP
	HasAvgLatency() bool
	// MinLatency returns uint64, set in Rocev2FlowMetricPerQP.
	MinLatency() uint64
	// SetMinLatency assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetMinLatency(value uint64) Rocev2FlowMetricPerQP
	// HasMinLatency checks if MinLatency has been set in Rocev2FlowMetricPerQP
	HasMinLatency() bool
	// MaxLatency returns uint64, set in Rocev2FlowMetricPerQP.
	MaxLatency() uint64
	// SetMaxLatency assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetMaxLatency(value uint64) Rocev2FlowMetricPerQP
	// HasMaxLatency checks if MaxLatency has been set in Rocev2FlowMetricPerQP
	HasMaxLatency() bool
	// EcnCeRx returns uint64, set in Rocev2FlowMetricPerQP.
	EcnCeRx() uint64
	// SetEcnCeRx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetEcnCeRx(value uint64) Rocev2FlowMetricPerQP
	// HasEcnCeRx checks if EcnCeRx has been set in Rocev2FlowMetricPerQP
	HasEcnCeRx() bool
	// CnpTx returns uint64, set in Rocev2FlowMetricPerQP.
	CnpTx() uint64
	// SetCnpTx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetCnpTx(value uint64) Rocev2FlowMetricPerQP
	// HasCnpTx checks if CnpTx has been set in Rocev2FlowMetricPerQP
	HasCnpTx() bool
	// CnpRx returns uint64, set in Rocev2FlowMetricPerQP.
	CnpRx() uint64
	// SetCnpRx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetCnpRx(value uint64) Rocev2FlowMetricPerQP
	// HasCnpRx checks if CnpRx has been set in Rocev2FlowMetricPerQP
	HasCnpRx() bool
	// AckTx returns uint64, set in Rocev2FlowMetricPerQP.
	AckTx() uint64
	// SetAckTx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetAckTx(value uint64) Rocev2FlowMetricPerQP
	// HasAckTx checks if AckTx has been set in Rocev2FlowMetricPerQP
	HasAckTx() bool
	// AckRx returns uint64, set in Rocev2FlowMetricPerQP.
	AckRx() uint64
	// SetAckRx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetAckRx(value uint64) Rocev2FlowMetricPerQP
	// HasAckRx checks if AckRx has been set in Rocev2FlowMetricPerQP
	HasAckRx() bool
	// NakTx returns uint64, set in Rocev2FlowMetricPerQP.
	NakTx() uint64
	// SetNakTx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetNakTx(value uint64) Rocev2FlowMetricPerQP
	// HasNakTx checks if NakTx has been set in Rocev2FlowMetricPerQP
	HasNakTx() bool
	// NakRx returns uint64, set in Rocev2FlowMetricPerQP.
	NakRx() uint64
	// SetNakRx assigns uint64 provided by user to Rocev2FlowMetricPerQP
	SetNakRx(value uint64) Rocev2FlowMetricPerQP
	// HasNakRx checks if NakRx has been set in Rocev2FlowMetricPerQP
	HasNakRx() bool
	// FirstTimestamp returns string, set in Rocev2FlowMetricPerQP.
	FirstTimestamp() string
	// SetFirstTimestamp assigns string provided by user to Rocev2FlowMetricPerQP
	SetFirstTimestamp(value string) Rocev2FlowMetricPerQP
	// HasFirstTimestamp checks if FirstTimestamp has been set in Rocev2FlowMetricPerQP
	HasFirstTimestamp() bool
	// LastTimestamp returns string, set in Rocev2FlowMetricPerQP.
	LastTimestamp() string
	// SetLastTimestamp assigns string provided by user to Rocev2FlowMetricPerQP
	SetLastTimestamp(value string) Rocev2FlowMetricPerQP
	// HasLastTimestamp checks if LastTimestamp has been set in Rocev2FlowMetricPerQP
	HasLastTimestamp() bool
}

// Flow Name.
// FlowName returns a string
func (obj *rocev2FlowMetricPerQP) FlowName() string {

	return *obj.obj.FlowName

}

// Flow Name.
// FlowName returns a string
func (obj *rocev2FlowMetricPerQP) HasFlowName() bool {
	return obj.obj.FlowName != nil
}

// Flow Name.
// SetFlowName sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetFlowName(value string) Rocev2FlowMetricPerQP {

	obj.obj.FlowName = &value
	return obj
}

// The name of the transmit port
// PortTx returns a string
func (obj *rocev2FlowMetricPerQP) PortTx() string {

	return *obj.obj.PortTx

}

// The name of the transmit port
// PortTx returns a string
func (obj *rocev2FlowMetricPerQP) HasPortTx() bool {
	return obj.obj.PortTx != nil
}

// The name of the transmit port
// SetPortTx sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetPortTx(value string) Rocev2FlowMetricPerQP {

	obj.obj.PortTx = &value
	return obj
}

// The name of the receive port
// PortRx returns a string
func (obj *rocev2FlowMetricPerQP) PortRx() string {

	return *obj.obj.PortRx

}

// The name of the receive port
// PortRx returns a string
func (obj *rocev2FlowMetricPerQP) HasPortRx() bool {
	return obj.obj.PortRx != nil
}

// The name of the receive port
// SetPortRx sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetPortRx(value string) Rocev2FlowMetricPerQP {

	obj.obj.PortRx = &value
	return obj
}

// Current source QP number.
// SrcQp returns a uint64
func (obj *rocev2FlowMetricPerQP) SrcQp() uint64 {

	return *obj.obj.SrcQp

}

// Current source QP number.
// SrcQp returns a uint64
func (obj *rocev2FlowMetricPerQP) HasSrcQp() bool {
	return obj.obj.SrcQp != nil
}

// Current source QP number.
// SetSrcQp sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetSrcQp(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.SrcQp = &value
	return obj
}

// Current destination QP number.
// DestQp returns a uint64
func (obj *rocev2FlowMetricPerQP) DestQp() uint64 {

	return *obj.obj.DestQp

}

// Current destination QP number.
// DestQp returns a uint64
func (obj *rocev2FlowMetricPerQP) HasDestQp() bool {
	return obj.obj.DestQp != nil
}

// Current destination QP number.
// SetDestQp sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDestQp(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.DestQp = &value
	return obj
}

// Current source address.
// SrcIpv4 returns a string
func (obj *rocev2FlowMetricPerQP) SrcIpv4() string {

	return *obj.obj.SrcIpv4

}

// Current source address.
// SrcIpv4 returns a string
func (obj *rocev2FlowMetricPerQP) HasSrcIpv4() bool {
	return obj.obj.SrcIpv4 != nil
}

// Current source address.
// SetSrcIpv4 sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetSrcIpv4(value string) Rocev2FlowMetricPerQP {

	obj.obj.SrcIpv4 = &value
	return obj
}

// Current destination address.
// DestIpv4 returns a string
func (obj *rocev2FlowMetricPerQP) DestIpv4() string {

	return *obj.obj.DestIpv4

}

// Current destination address.
// DestIpv4 returns a string
func (obj *rocev2FlowMetricPerQP) HasDestIpv4() bool {
	return obj.obj.DestIpv4 != nil
}

// Current destination address.
// SetDestIpv4 sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDestIpv4(value string) Rocev2FlowMetricPerQP {

	obj.obj.DestIpv4 = &value
	return obj
}

// Current number of data frames transmitted.
// DataFramesTx returns a uint64
func (obj *rocev2FlowMetricPerQP) DataFramesTx() uint64 {

	return *obj.obj.DataFramesTx

}

// Current number of data frames transmitted.
// DataFramesTx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasDataFramesTx() bool {
	return obj.obj.DataFramesTx != nil
}

// Current number of data frames transmitted.
// SetDataFramesTx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDataFramesTx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.DataFramesTx = &value
	return obj
}

// Current number of data frames received.
// DataFramesRx returns a uint64
func (obj *rocev2FlowMetricPerQP) DataFramesRx() uint64 {

	return *obj.obj.DataFramesRx

}

// Current number of data frames received.
// DataFramesRx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasDataFramesRx() bool {
	return obj.obj.DataFramesRx != nil
}

// Current number of data frames received.
// SetDataFramesRx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDataFramesRx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.DataFramesRx = &value
	return obj
}

// Current differnece between tx and rx data frames
// FrameDelta returns a uint64
func (obj *rocev2FlowMetricPerQP) FrameDelta() uint64 {

	return *obj.obj.FrameDelta

}

// Current differnece between tx and rx data frames
// FrameDelta returns a uint64
func (obj *rocev2FlowMetricPerQP) HasFrameDelta() bool {
	return obj.obj.FrameDelta != nil
}

// Current differnece between tx and rx data frames
// SetFrameDelta sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetFrameDelta(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.FrameDelta = &value
	return obj
}

// Current number of data frames re-transmitted.
// DataFramesRetransmitted returns a uint64
func (obj *rocev2FlowMetricPerQP) DataFramesRetransmitted() uint64 {

	return *obj.obj.DataFramesRetransmitted

}

// Current number of data frames re-transmitted.
// DataFramesRetransmitted returns a uint64
func (obj *rocev2FlowMetricPerQP) HasDataFramesRetransmitted() bool {
	return obj.obj.DataFramesRetransmitted != nil
}

// Current number of data frames re-transmitted.
// SetDataFramesRetransmitted sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDataFramesRetransmitted(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.DataFramesRetransmitted = &value
	return obj
}

// Current number of frame sequence errors.
// FrameSequenceError returns a uint64
func (obj *rocev2FlowMetricPerQP) FrameSequenceError() uint64 {

	return *obj.obj.FrameSequenceError

}

// Current number of frame sequence errors.
// FrameSequenceError returns a uint64
func (obj *rocev2FlowMetricPerQP) HasFrameSequenceError() bool {
	return obj.obj.FrameSequenceError != nil
}

// Current number of frame sequence errors.
// SetFrameSequenceError sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetFrameSequenceError(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.FrameSequenceError = &value
	return obj
}

// Current number of bytes transmitted.
// TxBytes returns a uint64
func (obj *rocev2FlowMetricPerQP) TxBytes() uint64 {

	return *obj.obj.TxBytes

}

// Current number of bytes transmitted.
// TxBytes returns a uint64
func (obj *rocev2FlowMetricPerQP) HasTxBytes() bool {
	return obj.obj.TxBytes != nil
}

// Current number of bytes transmitted.
// SetTxBytes sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetTxBytes(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.TxBytes = &value
	return obj
}

// Current number of bytes received.
// RxBytes returns a uint64
func (obj *rocev2FlowMetricPerQP) RxBytes() uint64 {

	return *obj.obj.RxBytes

}

// Current number of bytes received.
// RxBytes returns a uint64
func (obj *rocev2FlowMetricPerQP) HasRxBytes() bool {
	return obj.obj.RxBytes != nil
}

// Current number of bytes received.
// SetRxBytes sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetRxBytes(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.RxBytes = &value
	return obj
}

// Current rate at which data is transmitted in Gbps.
// DataTxRate returns a uint64
func (obj *rocev2FlowMetricPerQP) DataTxRate() uint64 {

	return *obj.obj.DataTxRate

}

// Current rate at which data is transmitted in Gbps.
// DataTxRate returns a uint64
func (obj *rocev2FlowMetricPerQP) HasDataTxRate() bool {
	return obj.obj.DataTxRate != nil
}

// Current rate at which data is transmitted in Gbps.
// SetDataTxRate sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDataTxRate(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.DataTxRate = &value
	return obj
}

// Current rate at which data is received in Gbps.
// DataRxRate returns a uint64
func (obj *rocev2FlowMetricPerQP) DataRxRate() uint64 {

	return *obj.obj.DataRxRate

}

// Current rate at which data is received in Gbps.
// DataRxRate returns a uint64
func (obj *rocev2FlowMetricPerQP) HasDataRxRate() bool {
	return obj.obj.DataRxRate != nil
}

// Current rate at which data is received in Gbps.
// SetDataRxRate sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetDataRxRate(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.DataRxRate = &value
	return obj
}

// Current number of Message transmitted.
// MessageTx returns a uint64
func (obj *rocev2FlowMetricPerQP) MessageTx() uint64 {

	return *obj.obj.MessageTx

}

// Current number of Message transmitted.
// MessageTx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasMessageTx() bool {
	return obj.obj.MessageTx != nil
}

// Current number of Message transmitted.
// SetMessageTx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetMessageTx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.MessageTx = &value
	return obj
}

// Current number of Message Complete received.
// MessageCompleteRx returns a uint64
func (obj *rocev2FlowMetricPerQP) MessageCompleteRx() uint64 {

	return *obj.obj.MessageCompleteRx

}

// Current number of Message Complete received.
// MessageCompleteRx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasMessageCompleteRx() bool {
	return obj.obj.MessageCompleteRx != nil
}

// Current number of Message Complete received.
// SetMessageCompleteRx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetMessageCompleteRx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.MessageCompleteRx = &value
	return obj
}

// Current number of Message Fail count.
// MessageFail returns a uint64
func (obj *rocev2FlowMetricPerQP) MessageFail() uint64 {

	return *obj.obj.MessageFail

}

// Current number of Message Fail count.
// MessageFail returns a uint64
func (obj *rocev2FlowMetricPerQP) HasMessageFail() bool {
	return obj.obj.MessageFail != nil
}

// Current number of Message Fail count.
// SetMessageFail sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetMessageFail(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.MessageFail = &value
	return obj
}

// Current flow comletion time in ms.
// FlowCompletionTime returns a uint64
func (obj *rocev2FlowMetricPerQP) FlowCompletionTime() uint64 {

	return *obj.obj.FlowCompletionTime

}

// Current flow comletion time in ms.
// FlowCompletionTime returns a uint64
func (obj *rocev2FlowMetricPerQP) HasFlowCompletionTime() bool {
	return obj.obj.FlowCompletionTime != nil
}

// Current flow comletion time in ms.
// SetFlowCompletionTime sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetFlowCompletionTime(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.FlowCompletionTime = &value
	return obj
}

// Current average latency measured in ns.
// AvgLatency returns a uint64
func (obj *rocev2FlowMetricPerQP) AvgLatency() uint64 {

	return *obj.obj.AvgLatency

}

// Current average latency measured in ns.
// AvgLatency returns a uint64
func (obj *rocev2FlowMetricPerQP) HasAvgLatency() bool {
	return obj.obj.AvgLatency != nil
}

// Current average latency measured in ns.
// SetAvgLatency sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetAvgLatency(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.AvgLatency = &value
	return obj
}

// Current minimum latency measured in ns.
// MinLatency returns a uint64
func (obj *rocev2FlowMetricPerQP) MinLatency() uint64 {

	return *obj.obj.MinLatency

}

// Current minimum latency measured in ns.
// MinLatency returns a uint64
func (obj *rocev2FlowMetricPerQP) HasMinLatency() bool {
	return obj.obj.MinLatency != nil
}

// Current minimum latency measured in ns.
// SetMinLatency sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetMinLatency(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.MinLatency = &value
	return obj
}

// Current maximum latency measured in ns.
// MaxLatency returns a uint64
func (obj *rocev2FlowMetricPerQP) MaxLatency() uint64 {

	return *obj.obj.MaxLatency

}

// Current maximum latency measured in ns.
// MaxLatency returns a uint64
func (obj *rocev2FlowMetricPerQP) HasMaxLatency() bool {
	return obj.obj.MaxLatency != nil
}

// Current maximum latency measured in ns.
// SetMaxLatency sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetMaxLatency(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.MaxLatency = &value
	return obj
}

// Current number of ECN-CE Recevied.
// EcnCeRx returns a uint64
func (obj *rocev2FlowMetricPerQP) EcnCeRx() uint64 {

	return *obj.obj.EcnCeRx

}

// Current number of ECN-CE Recevied.
// EcnCeRx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasEcnCeRx() bool {
	return obj.obj.EcnCeRx != nil
}

// Current number of ECN-CE Recevied.
// SetEcnCeRx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetEcnCeRx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.EcnCeRx = &value
	return obj
}

// Current number of CNP transmitted.
// CnpTx returns a uint64
func (obj *rocev2FlowMetricPerQP) CnpTx() uint64 {

	return *obj.obj.CnpTx

}

// Current number of CNP transmitted.
// CnpTx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasCnpTx() bool {
	return obj.obj.CnpTx != nil
}

// Current number of CNP transmitted.
// SetCnpTx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetCnpTx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.CnpTx = &value
	return obj
}

// Current number of CNP received.
// CnpRx returns a uint64
func (obj *rocev2FlowMetricPerQP) CnpRx() uint64 {

	return *obj.obj.CnpRx

}

// Current number of CNP received.
// CnpRx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasCnpRx() bool {
	return obj.obj.CnpRx != nil
}

// Current number of CNP received.
// SetCnpRx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetCnpRx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.CnpRx = &value
	return obj
}

// Current number of ACK transmitted.
// AckTx returns a uint64
func (obj *rocev2FlowMetricPerQP) AckTx() uint64 {

	return *obj.obj.AckTx

}

// Current number of ACK transmitted.
// AckTx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasAckTx() bool {
	return obj.obj.AckTx != nil
}

// Current number of ACK transmitted.
// SetAckTx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetAckTx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.AckTx = &value
	return obj
}

// Current number of ACK received.
// AckRx returns a uint64
func (obj *rocev2FlowMetricPerQP) AckRx() uint64 {

	return *obj.obj.AckRx

}

// Current number of ACK received.
// AckRx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasAckRx() bool {
	return obj.obj.AckRx != nil
}

// Current number of ACK received.
// SetAckRx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetAckRx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.AckRx = &value
	return obj
}

// Current number of NAK transmitted.
// NakTx returns a uint64
func (obj *rocev2FlowMetricPerQP) NakTx() uint64 {

	return *obj.obj.NakTx

}

// Current number of NAK transmitted.
// NakTx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasNakTx() bool {
	return obj.obj.NakTx != nil
}

// Current number of NAK transmitted.
// SetNakTx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetNakTx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.NakTx = &value
	return obj
}

// Current number of NAK received.
// NakRx returns a uint64
func (obj *rocev2FlowMetricPerQP) NakRx() uint64 {

	return *obj.obj.NakRx

}

// Current number of NAK received.
// NakRx returns a uint64
func (obj *rocev2FlowMetricPerQP) HasNakRx() bool {
	return obj.obj.NakRx != nil
}

// Current number of NAK received.
// SetNakRx sets the uint64 value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetNakRx(value uint64) Rocev2FlowMetricPerQP {

	obj.obj.NakRx = &value
	return obj
}

// First Timestamp.
// FirstTimestamp returns a string
func (obj *rocev2FlowMetricPerQP) FirstTimestamp() string {

	return *obj.obj.FirstTimestamp

}

// First Timestamp.
// FirstTimestamp returns a string
func (obj *rocev2FlowMetricPerQP) HasFirstTimestamp() bool {
	return obj.obj.FirstTimestamp != nil
}

// First Timestamp.
// SetFirstTimestamp sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetFirstTimestamp(value string) Rocev2FlowMetricPerQP {

	obj.obj.FirstTimestamp = &value
	return obj
}

// Last Timestamp.
// LastTimestamp returns a string
func (obj *rocev2FlowMetricPerQP) LastTimestamp() string {

	return *obj.obj.LastTimestamp

}

// Last Timestamp.
// LastTimestamp returns a string
func (obj *rocev2FlowMetricPerQP) HasLastTimestamp() bool {
	return obj.obj.LastTimestamp != nil
}

// Last Timestamp.
// SetLastTimestamp sets the string value in the Rocev2FlowMetricPerQP object
func (obj *rocev2FlowMetricPerQP) SetLastTimestamp(value string) Rocev2FlowMetricPerQP {

	obj.obj.LastTimestamp = &value
	return obj
}

func (obj *rocev2FlowMetricPerQP) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrcIpv4 != nil {

		err := obj.validateIpv4(obj.SrcIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2FlowMetricPerQP.SrcIpv4"))
		}

	}

	if obj.obj.DestIpv4 != nil {

		err := obj.validateIpv4(obj.DestIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Rocev2FlowMetricPerQP.DestIpv4"))
		}

	}

}

func (obj *rocev2FlowMetricPerQP) setDefault() {

}
