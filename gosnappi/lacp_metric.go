package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LacpMetric *****
type lacpMetric struct {
	validation
	obj          *otg.LacpMetric
	marshaller   marshalLacpMetric
	unMarshaller unMarshalLacpMetric
}

func NewLacpMetric() LacpMetric {
	obj := lacpMetric{obj: &otg.LacpMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *lacpMetric) msg() *otg.LacpMetric {
	return obj.obj
}

func (obj *lacpMetric) setMsg(msg *otg.LacpMetric) LacpMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallacpMetric struct {
	obj *lacpMetric
}

type marshalLacpMetric interface {
	// ToProto marshals LacpMetric to protobuf object *otg.LacpMetric
	ToProto() (*otg.LacpMetric, error)
	// ToPbText marshals LacpMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LacpMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals LacpMetric to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LacpMetric to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallacpMetric struct {
	obj *lacpMetric
}

type unMarshalLacpMetric interface {
	// FromProto unmarshals LacpMetric from protobuf object *otg.LacpMetric
	FromProto(msg *otg.LacpMetric) (LacpMetric, error)
	// FromPbText unmarshals LacpMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LacpMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LacpMetric from JSON text
	FromJson(value string) error
}

func (obj *lacpMetric) Marshal() marshalLacpMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshallacpMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *lacpMetric) Unmarshal() unMarshalLacpMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallacpMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallacpMetric) ToProto() (*otg.LacpMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallacpMetric) FromProto(msg *otg.LacpMetric) (LacpMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallacpMetric) ToPbText() (string, error) {
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

func (m *unMarshallacpMetric) FromPbText(value string) error {
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

func (m *marshallacpMetric) ToYaml() (string, error) {
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

func (m *unMarshallacpMetric) FromYaml(value string) error {
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

func (m *marshallacpMetric) ToJsonRaw() (string, error) {
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

func (m *marshallacpMetric) ToJson() (string, error) {
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

func (m *unMarshallacpMetric) FromJson(value string) error {
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

func (obj *lacpMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lacpMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lacpMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lacpMetric) Clone() (LacpMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLacpMetric()
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

// LacpMetric is lACP metrics (statistics) per LAG member.
type LacpMetric interface {
	Validation
	// msg marshals LacpMetric to protobuf object *otg.LacpMetric
	// and doesn't set defaults
	msg() *otg.LacpMetric
	// setMsg unmarshals LacpMetric from protobuf object *otg.LacpMetric
	// and doesn't set defaults
	setMsg(*otg.LacpMetric) LacpMetric
	// provides marshal interface
	Marshal() marshalLacpMetric
	// provides unmarshal interface
	Unmarshal() unMarshalLacpMetric
	// validate validates LacpMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LacpMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LagName returns string, set in LacpMetric.
	LagName() string
	// SetLagName assigns string provided by user to LacpMetric
	SetLagName(value string) LacpMetric
	// HasLagName checks if LagName has been set in LacpMetric
	HasLagName() bool
	// LagMemberPortName returns string, set in LacpMetric.
	LagMemberPortName() string
	// SetLagMemberPortName assigns string provided by user to LacpMetric
	SetLagMemberPortName(value string) LacpMetric
	// HasLagMemberPortName checks if LagMemberPortName has been set in LacpMetric
	HasLagMemberPortName() bool
	// LacpPacketsRx returns uint64, set in LacpMetric.
	LacpPacketsRx() uint64
	// SetLacpPacketsRx assigns uint64 provided by user to LacpMetric
	SetLacpPacketsRx(value uint64) LacpMetric
	// HasLacpPacketsRx checks if LacpPacketsRx has been set in LacpMetric
	HasLacpPacketsRx() bool
	// LacpPacketsTx returns uint64, set in LacpMetric.
	LacpPacketsTx() uint64
	// SetLacpPacketsTx assigns uint64 provided by user to LacpMetric
	SetLacpPacketsTx(value uint64) LacpMetric
	// HasLacpPacketsTx checks if LacpPacketsTx has been set in LacpMetric
	HasLacpPacketsTx() bool
	// LacpRxErrors returns uint64, set in LacpMetric.
	LacpRxErrors() uint64
	// SetLacpRxErrors assigns uint64 provided by user to LacpMetric
	SetLacpRxErrors(value uint64) LacpMetric
	// HasLacpRxErrors checks if LacpRxErrors has been set in LacpMetric
	HasLacpRxErrors() bool
	// Activity returns LacpMetricActivityEnum, set in LacpMetric
	Activity() LacpMetricActivityEnum
	// SetActivity assigns LacpMetricActivityEnum provided by user to LacpMetric
	SetActivity(value LacpMetricActivityEnum) LacpMetric
	// HasActivity checks if Activity has been set in LacpMetric
	HasActivity() bool
	// Timeout returns LacpMetricTimeoutEnum, set in LacpMetric
	Timeout() LacpMetricTimeoutEnum
	// SetTimeout assigns LacpMetricTimeoutEnum provided by user to LacpMetric
	SetTimeout(value LacpMetricTimeoutEnum) LacpMetric
	// HasTimeout checks if Timeout has been set in LacpMetric
	HasTimeout() bool
	// Synchronization returns LacpMetricSynchronizationEnum, set in LacpMetric
	Synchronization() LacpMetricSynchronizationEnum
	// SetSynchronization assigns LacpMetricSynchronizationEnum provided by user to LacpMetric
	SetSynchronization(value LacpMetricSynchronizationEnum) LacpMetric
	// HasSynchronization checks if Synchronization has been set in LacpMetric
	HasSynchronization() bool
	// Aggregatable returns bool, set in LacpMetric.
	Aggregatable() bool
	// SetAggregatable assigns bool provided by user to LacpMetric
	SetAggregatable(value bool) LacpMetric
	// HasAggregatable checks if Aggregatable has been set in LacpMetric
	HasAggregatable() bool
	// Collecting returns bool, set in LacpMetric.
	Collecting() bool
	// SetCollecting assigns bool provided by user to LacpMetric
	SetCollecting(value bool) LacpMetric
	// HasCollecting checks if Collecting has been set in LacpMetric
	HasCollecting() bool
	// Distributing returns bool, set in LacpMetric.
	Distributing() bool
	// SetDistributing assigns bool provided by user to LacpMetric
	SetDistributing(value bool) LacpMetric
	// HasDistributing checks if Distributing has been set in LacpMetric
	HasDistributing() bool
	// SystemId returns string, set in LacpMetric.
	SystemId() string
	// SetSystemId assigns string provided by user to LacpMetric
	SetSystemId(value string) LacpMetric
	// HasSystemId checks if SystemId has been set in LacpMetric
	HasSystemId() bool
	// OperKey returns uint32, set in LacpMetric.
	OperKey() uint32
	// SetOperKey assigns uint32 provided by user to LacpMetric
	SetOperKey(value uint32) LacpMetric
	// HasOperKey checks if OperKey has been set in LacpMetric
	HasOperKey() bool
	// PartnerId returns string, set in LacpMetric.
	PartnerId() string
	// SetPartnerId assigns string provided by user to LacpMetric
	SetPartnerId(value string) LacpMetric
	// HasPartnerId checks if PartnerId has been set in LacpMetric
	HasPartnerId() bool
	// PartnerKey returns uint32, set in LacpMetric.
	PartnerKey() uint32
	// SetPartnerKey assigns uint32 provided by user to LacpMetric
	SetPartnerKey(value uint32) LacpMetric
	// HasPartnerKey checks if PartnerKey has been set in LacpMetric
	HasPartnerKey() bool
	// PortNum returns uint32, set in LacpMetric.
	PortNum() uint32
	// SetPortNum assigns uint32 provided by user to LacpMetric
	SetPortNum(value uint32) LacpMetric
	// HasPortNum checks if PortNum has been set in LacpMetric
	HasPortNum() bool
	// PartnerPortNum returns uint32, set in LacpMetric.
	PartnerPortNum() uint32
	// SetPartnerPortNum assigns uint32 provided by user to LacpMetric
	SetPartnerPortNum(value uint32) LacpMetric
	// HasPartnerPortNum checks if PartnerPortNum has been set in LacpMetric
	HasPartnerPortNum() bool
}

// The name of a LAG (ports group) configured with LACP.
// LagName returns a string
func (obj *lacpMetric) LagName() string {

	return *obj.obj.LagName

}

// The name of a LAG (ports group) configured with LACP.
// LagName returns a string
func (obj *lacpMetric) HasLagName() bool {
	return obj.obj.LagName != nil
}

// The name of a LAG (ports group) configured with LACP.
// SetLagName sets the string value in the LacpMetric object
func (obj *lacpMetric) SetLagName(value string) LacpMetric {

	obj.obj.LagName = &value
	return obj
}

// The name of a LAG member (port) configured with LACP.
// LagMemberPortName returns a string
func (obj *lacpMetric) LagMemberPortName() string {

	return *obj.obj.LagMemberPortName

}

// The name of a LAG member (port) configured with LACP.
// LagMemberPortName returns a string
func (obj *lacpMetric) HasLagMemberPortName() bool {
	return obj.obj.LagMemberPortName != nil
}

// The name of a LAG member (port) configured with LACP.
// SetLagMemberPortName sets the string value in the LacpMetric object
func (obj *lacpMetric) SetLagMemberPortName(value string) LacpMetric {

	obj.obj.LagMemberPortName = &value
	return obj
}

// Number of LACPDUs received.
// LacpPacketsRx returns a uint64
func (obj *lacpMetric) LacpPacketsRx() uint64 {

	return *obj.obj.LacpPacketsRx

}

// Number of LACPDUs received.
// LacpPacketsRx returns a uint64
func (obj *lacpMetric) HasLacpPacketsRx() bool {
	return obj.obj.LacpPacketsRx != nil
}

// Number of LACPDUs received.
// SetLacpPacketsRx sets the uint64 value in the LacpMetric object
func (obj *lacpMetric) SetLacpPacketsRx(value uint64) LacpMetric {

	obj.obj.LacpPacketsRx = &value
	return obj
}

// Number of LACPDUs transmitted.
// LacpPacketsTx returns a uint64
func (obj *lacpMetric) LacpPacketsTx() uint64 {

	return *obj.obj.LacpPacketsTx

}

// Number of LACPDUs transmitted.
// LacpPacketsTx returns a uint64
func (obj *lacpMetric) HasLacpPacketsTx() bool {
	return obj.obj.LacpPacketsTx != nil
}

// Number of LACPDUs transmitted.
// SetLacpPacketsTx sets the uint64 value in the LacpMetric object
func (obj *lacpMetric) SetLacpPacketsTx(value uint64) LacpMetric {

	obj.obj.LacpPacketsTx = &value
	return obj
}

// Number of LACPDUs receive packet errors.
// LacpRxErrors returns a uint64
func (obj *lacpMetric) LacpRxErrors() uint64 {

	return *obj.obj.LacpRxErrors

}

// Number of LACPDUs receive packet errors.
// LacpRxErrors returns a uint64
func (obj *lacpMetric) HasLacpRxErrors() bool {
	return obj.obj.LacpRxErrors != nil
}

// Number of LACPDUs receive packet errors.
// SetLacpRxErrors sets the uint64 value in the LacpMetric object
func (obj *lacpMetric) SetLacpRxErrors(value uint64) LacpMetric {

	obj.obj.LacpRxErrors = &value
	return obj
}

type LacpMetricActivityEnum string

// Enum of Activity on LacpMetric
var LacpMetricActivity = struct {
	ACTIVE  LacpMetricActivityEnum
	PASSIVE LacpMetricActivityEnum
}{
	ACTIVE:  LacpMetricActivityEnum("active"),
	PASSIVE: LacpMetricActivityEnum("passive"),
}

func (obj *lacpMetric) Activity() LacpMetricActivityEnum {
	return LacpMetricActivityEnum(obj.obj.Activity.Enum().String())
}

// Indicates participant is active or passive.
// Activity returns a string
func (obj *lacpMetric) HasActivity() bool {
	return obj.obj.Activity != nil
}

func (obj *lacpMetric) SetActivity(value LacpMetricActivityEnum) LacpMetric {
	intValue, ok := otg.LacpMetric_Activity_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LacpMetricActivityEnum", string(value)))
		return obj
	}
	enumValue := otg.LacpMetric_Activity_Enum(intValue)
	obj.obj.Activity = &enumValue

	return obj
}

type LacpMetricTimeoutEnum string

// Enum of Timeout on LacpMetric
var LacpMetricTimeout = struct {
	SHORT LacpMetricTimeoutEnum
	LONG  LacpMetricTimeoutEnum
}{
	SHORT: LacpMetricTimeoutEnum("short"),
	LONG:  LacpMetricTimeoutEnum("long"),
}

func (obj *lacpMetric) Timeout() LacpMetricTimeoutEnum {
	return LacpMetricTimeoutEnum(obj.obj.Timeout.Enum().String())
}

// The timeout type (short or long) used by the participant.
// Timeout returns a string
func (obj *lacpMetric) HasTimeout() bool {
	return obj.obj.Timeout != nil
}

func (obj *lacpMetric) SetTimeout(value LacpMetricTimeoutEnum) LacpMetric {
	intValue, ok := otg.LacpMetric_Timeout_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LacpMetricTimeoutEnum", string(value)))
		return obj
	}
	enumValue := otg.LacpMetric_Timeout_Enum(intValue)
	obj.obj.Timeout = &enumValue

	return obj
}

type LacpMetricSynchronizationEnum string

// Enum of Synchronization on LacpMetric
var LacpMetricSynchronization = struct {
	IN_SYNC  LacpMetricSynchronizationEnum
	OUT_SYNC LacpMetricSynchronizationEnum
}{
	IN_SYNC:  LacpMetricSynchronizationEnum("in_sync"),
	OUT_SYNC: LacpMetricSynchronizationEnum("out_sync"),
}

func (obj *lacpMetric) Synchronization() LacpMetricSynchronizationEnum {
	return LacpMetricSynchronizationEnum(obj.obj.Synchronization.Enum().String())
}

// Indicates whether the participant is in-sync or out-of-sync.
// Synchronization returns a string
func (obj *lacpMetric) HasSynchronization() bool {
	return obj.obj.Synchronization != nil
}

func (obj *lacpMetric) SetSynchronization(value LacpMetricSynchronizationEnum) LacpMetric {
	intValue, ok := otg.LacpMetric_Synchronization_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LacpMetricSynchronizationEnum", string(value)))
		return obj
	}
	enumValue := otg.LacpMetric_Synchronization_Enum(intValue)
	obj.obj.Synchronization = &enumValue

	return obj
}

// A true value indicates that the participant will allow the link to be used as part of the aggregate. A false value indicates the link should be used as an  individual link.
// Aggregatable returns a bool
func (obj *lacpMetric) Aggregatable() bool {

	return *obj.obj.Aggregatable

}

// A true value indicates that the participant will allow the link to be used as part of the aggregate. A false value indicates the link should be used as an  individual link.
// Aggregatable returns a bool
func (obj *lacpMetric) HasAggregatable() bool {
	return obj.obj.Aggregatable != nil
}

// A true value indicates that the participant will allow the link to be used as part of the aggregate. A false value indicates the link should be used as an  individual link.
// SetAggregatable sets the bool value in the LacpMetric object
func (obj *lacpMetric) SetAggregatable(value bool) LacpMetric {

	obj.obj.Aggregatable = &value
	return obj
}

// If true, the participant is collecting incoming frames on the link, otherwise false.
// Collecting returns a bool
func (obj *lacpMetric) Collecting() bool {

	return *obj.obj.Collecting

}

// If true, the participant is collecting incoming frames on the link, otherwise false.
// Collecting returns a bool
func (obj *lacpMetric) HasCollecting() bool {
	return obj.obj.Collecting != nil
}

// If true, the participant is collecting incoming frames on the link, otherwise false.
// SetCollecting sets the bool value in the LacpMetric object
func (obj *lacpMetric) SetCollecting(value bool) LacpMetric {

	obj.obj.Collecting = &value
	return obj
}

// When true, the participant is distributing outgoing frames; when false, distribution is disabled.
// Distributing returns a bool
func (obj *lacpMetric) Distributing() bool {

	return *obj.obj.Distributing

}

// When true, the participant is distributing outgoing frames; when false, distribution is disabled.
// Distributing returns a bool
func (obj *lacpMetric) HasDistributing() bool {
	return obj.obj.Distributing != nil
}

// When true, the participant is distributing outgoing frames; when false, distribution is disabled.
// SetDistributing sets the bool value in the LacpMetric object
func (obj *lacpMetric) SetDistributing(value bool) LacpMetric {

	obj.obj.Distributing = &value
	return obj
}

// MAC address that defines the local system ID for the aggregate interface.
// SystemId returns a string
func (obj *lacpMetric) SystemId() string {

	return *obj.obj.SystemId

}

// MAC address that defines the local system ID for the aggregate interface.
// SystemId returns a string
func (obj *lacpMetric) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// MAC address that defines the local system ID for the aggregate interface.
// SetSystemId sets the string value in the LacpMetric object
func (obj *lacpMetric) SetSystemId(value string) LacpMetric {

	obj.obj.SystemId = &value
	return obj
}

// Current operational value of the key for the aggregate interface.
// OperKey returns a uint32
func (obj *lacpMetric) OperKey() uint32 {

	return *obj.obj.OperKey

}

// Current operational value of the key for the aggregate interface.
// OperKey returns a uint32
func (obj *lacpMetric) HasOperKey() bool {
	return obj.obj.OperKey != nil
}

// Current operational value of the key for the aggregate interface.
// SetOperKey sets the uint32 value in the LacpMetric object
func (obj *lacpMetric) SetOperKey(value uint32) LacpMetric {

	obj.obj.OperKey = &value
	return obj
}

// MAC address representing the protocol partner's interface system ID.
// PartnerId returns a string
func (obj *lacpMetric) PartnerId() string {

	return *obj.obj.PartnerId

}

// MAC address representing the protocol partner's interface system ID.
// PartnerId returns a string
func (obj *lacpMetric) HasPartnerId() bool {
	return obj.obj.PartnerId != nil
}

// MAC address representing the protocol partner's interface system ID.
// SetPartnerId sets the string value in the LacpMetric object
func (obj *lacpMetric) SetPartnerId(value string) LacpMetric {

	obj.obj.PartnerId = &value
	return obj
}

// Operational value of the protocol partner's key.
// PartnerKey returns a uint32
func (obj *lacpMetric) PartnerKey() uint32 {

	return *obj.obj.PartnerKey

}

// Operational value of the protocol partner's key.
// PartnerKey returns a uint32
func (obj *lacpMetric) HasPartnerKey() bool {
	return obj.obj.PartnerKey != nil
}

// Operational value of the protocol partner's key.
// SetPartnerKey sets the uint32 value in the LacpMetric object
func (obj *lacpMetric) SetPartnerKey(value uint32) LacpMetric {

	obj.obj.PartnerKey = &value
	return obj
}

// Port number of the local (actor) aggregation member.
// PortNum returns a uint32
func (obj *lacpMetric) PortNum() uint32 {

	return *obj.obj.PortNum

}

// Port number of the local (actor) aggregation member.
// PortNum returns a uint32
func (obj *lacpMetric) HasPortNum() bool {
	return obj.obj.PortNum != nil
}

// Port number of the local (actor) aggregation member.
// SetPortNum sets the uint32 value in the LacpMetric object
func (obj *lacpMetric) SetPortNum(value uint32) LacpMetric {

	obj.obj.PortNum = &value
	return obj
}

// Port number of the partner (remote) port for this member port.
// PartnerPortNum returns a uint32
func (obj *lacpMetric) PartnerPortNum() uint32 {

	return *obj.obj.PartnerPortNum

}

// Port number of the partner (remote) port for this member port.
// PartnerPortNum returns a uint32
func (obj *lacpMetric) HasPartnerPortNum() bool {
	return obj.obj.PartnerPortNum != nil
}

// Port number of the partner (remote) port for this member port.
// SetPartnerPortNum sets the uint32 value in the LacpMetric object
func (obj *lacpMetric) SetPartnerPortNum(value uint32) LacpMetric {

	obj.obj.PartnerPortNum = &value
	return obj
}

func (obj *lacpMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LacpMetric.SystemId"))
		}

	}

	if obj.obj.PartnerId != nil {

		err := obj.validateMac(obj.PartnerId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LacpMetric.PartnerId"))
		}

	}

}

func (obj *lacpMetric) setDefault() {

}
