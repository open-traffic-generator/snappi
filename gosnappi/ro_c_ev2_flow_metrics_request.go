package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowMetricsRequest *****
type roCEv2FlowMetricsRequest struct {
	validation
	obj          *otg.RoCEv2FlowMetricsRequest
	marshaller   marshalRoCEv2FlowMetricsRequest
	unMarshaller unMarshalRoCEv2FlowMetricsRequest
}

func NewRoCEv2FlowMetricsRequest() RoCEv2FlowMetricsRequest {
	obj := roCEv2FlowMetricsRequest{obj: &otg.RoCEv2FlowMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowMetricsRequest) msg() *otg.RoCEv2FlowMetricsRequest {
	return obj.obj
}

func (obj *roCEv2FlowMetricsRequest) setMsg(msg *otg.RoCEv2FlowMetricsRequest) RoCEv2FlowMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowMetricsRequest struct {
	obj *roCEv2FlowMetricsRequest
}

type marshalRoCEv2FlowMetricsRequest interface {
	// ToProto marshals RoCEv2FlowMetricsRequest to protobuf object *otg.RoCEv2FlowMetricsRequest
	ToProto() (*otg.RoCEv2FlowMetricsRequest, error)
	// ToPbText marshals RoCEv2FlowMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowMetricsRequest struct {
	obj *roCEv2FlowMetricsRequest
}

type unMarshalRoCEv2FlowMetricsRequest interface {
	// FromProto unmarshals RoCEv2FlowMetricsRequest from protobuf object *otg.RoCEv2FlowMetricsRequest
	FromProto(msg *otg.RoCEv2FlowMetricsRequest) (RoCEv2FlowMetricsRequest, error)
	// FromPbText unmarshals RoCEv2FlowMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowMetricsRequest) Marshal() marshalRoCEv2FlowMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowMetricsRequest) Unmarshal() unMarshalRoCEv2FlowMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowMetricsRequest) ToProto() (*otg.RoCEv2FlowMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowMetricsRequest) FromProto(msg *otg.RoCEv2FlowMetricsRequest) (RoCEv2FlowMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowMetricsRequest) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowMetricsRequest) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowMetricsRequest) FromJson(value string) error {
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

func (obj *roCEv2FlowMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowMetricsRequest) Clone() (RoCEv2FlowMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowMetricsRequest()
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

// RoCEv2FlowMetricsRequest is the request to retrieve RoCEv2 FLow statistics.
type RoCEv2FlowMetricsRequest interface {
	Validation
	// msg marshals RoCEv2FlowMetricsRequest to protobuf object *otg.RoCEv2FlowMetricsRequest
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowMetricsRequest
	// setMsg unmarshals RoCEv2FlowMetricsRequest from protobuf object *otg.RoCEv2FlowMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowMetricsRequest) RoCEv2FlowMetricsRequest
	// provides marshal interface
	Marshal() marshalRoCEv2FlowMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowMetricsRequest
	// validate validates RoCEv2FlowMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowNames returns []string, set in RoCEv2FlowMetricsRequest.
	FlowNames() []string
	// SetFlowNames assigns []string provided by user to RoCEv2FlowMetricsRequest
	SetFlowNames(value []string) RoCEv2FlowMetricsRequest
	// ColumnNames returns []RoCEv2FlowMetricsRequestColumnNamesEnum, set in RoCEv2FlowMetricsRequest
	ColumnNames() []RoCEv2FlowMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []RoCEv2FlowMetricsRequestColumnNamesEnum provided by user to RoCEv2FlowMetricsRequest
	SetColumnNames(value []RoCEv2FlowMetricsRequestColumnNamesEnum) RoCEv2FlowMetricsRequest
}

// The names of RoCEv2 FLow to return results for. An empty list will return results for all RoCEv2 FLow.
// FlowNames returns a []string
func (obj *roCEv2FlowMetricsRequest) FlowNames() []string {
	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	return obj.obj.FlowNames
}

// The names of RoCEv2 FLow to return results for. An empty list will return results for all RoCEv2 FLow.
// SetFlowNames sets the []string value in the RoCEv2FlowMetricsRequest object
func (obj *roCEv2FlowMetricsRequest) SetFlowNames(value []string) RoCEv2FlowMetricsRequest {

	if obj.obj.FlowNames == nil {
		obj.obj.FlowNames = make([]string, 0)
	}
	obj.obj.FlowNames = value

	return obj
}

type RoCEv2FlowMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on RoCEv2FlowMetricsRequest
var RoCEv2FlowMetricsRequestColumnNames = struct {
	FLOW_NAME                 RoCEv2FlowMetricsRequestColumnNamesEnum
	TRAFFIC_ITEM              RoCEv2FlowMetricsRequestColumnNamesEnum
	PORT_TX                   RoCEv2FlowMetricsRequestColumnNamesEnum
	PORT_RX                   RoCEv2FlowMetricsRequestColumnNamesEnum
	SRC_QP                    RoCEv2FlowMetricsRequestColumnNamesEnum
	DEST_QP                   RoCEv2FlowMetricsRequestColumnNamesEnum
	SRC_IPV4                  RoCEv2FlowMetricsRequestColumnNamesEnum
	DEST_IPV4                 RoCEv2FlowMetricsRequestColumnNamesEnum
	DATA_FRAMES_TX            RoCEv2FlowMetricsRequestColumnNamesEnum
	DATA_FRAME_RX             RoCEv2FlowMetricsRequestColumnNamesEnum
	FRAME_DELTA               RoCEv2FlowMetricsRequestColumnNamesEnum
	DATA_FRAMES_RETRANSMITTED RoCEv2FlowMetricsRequestColumnNamesEnum
	FRAME_SEQUENCE_ERROR      RoCEv2FlowMetricsRequestColumnNamesEnum
	TX_BYTES                  RoCEv2FlowMetricsRequestColumnNamesEnum
	RX_BYTES                  RoCEv2FlowMetricsRequestColumnNamesEnum
	DATA_TX_RATE              RoCEv2FlowMetricsRequestColumnNamesEnum
	DATA_RX_RATE              RoCEv2FlowMetricsRequestColumnNamesEnum
	MESSAGE_TX                RoCEv2FlowMetricsRequestColumnNamesEnum
	MESSAGE_COMPLETE_RX       RoCEv2FlowMetricsRequestColumnNamesEnum
	MESSAGE_FAIL              RoCEv2FlowMetricsRequestColumnNamesEnum
	FLOW_COMPLETION_TIME      RoCEv2FlowMetricsRequestColumnNamesEnum
	AVG_LATESNCY              RoCEv2FlowMetricsRequestColumnNamesEnum
	MIN_LATENCY               RoCEv2FlowMetricsRequestColumnNamesEnum
	MAX_LATENCY               RoCEv2FlowMetricsRequestColumnNamesEnum
	ECN_CE_RX                 RoCEv2FlowMetricsRequestColumnNamesEnum
	CNP_TX                    RoCEv2FlowMetricsRequestColumnNamesEnum
	CNP_RX                    RoCEv2FlowMetricsRequestColumnNamesEnum
	ACK_TX                    RoCEv2FlowMetricsRequestColumnNamesEnum
	ACK_RX                    RoCEv2FlowMetricsRequestColumnNamesEnum
	NAK_TX                    RoCEv2FlowMetricsRequestColumnNamesEnum
	NAK_RX                    RoCEv2FlowMetricsRequestColumnNamesEnum
	FIRST_TIMESTAMP           RoCEv2FlowMetricsRequestColumnNamesEnum
	LAST_TIMESTAMP            RoCEv2FlowMetricsRequestColumnNamesEnum
}{
	FLOW_NAME:                 RoCEv2FlowMetricsRequestColumnNamesEnum("flow_name"),
	TRAFFIC_ITEM:              RoCEv2FlowMetricsRequestColumnNamesEnum("traffic_item"),
	PORT_TX:                   RoCEv2FlowMetricsRequestColumnNamesEnum("port_tx"),
	PORT_RX:                   RoCEv2FlowMetricsRequestColumnNamesEnum("port_rx"),
	SRC_QP:                    RoCEv2FlowMetricsRequestColumnNamesEnum("src_qp"),
	DEST_QP:                   RoCEv2FlowMetricsRequestColumnNamesEnum("dest_qp"),
	SRC_IPV4:                  RoCEv2FlowMetricsRequestColumnNamesEnum("src_ipv4"),
	DEST_IPV4:                 RoCEv2FlowMetricsRequestColumnNamesEnum("dest_ipv4"),
	DATA_FRAMES_TX:            RoCEv2FlowMetricsRequestColumnNamesEnum("data_frames_tx"),
	DATA_FRAME_RX:             RoCEv2FlowMetricsRequestColumnNamesEnum("data_frame_rx"),
	FRAME_DELTA:               RoCEv2FlowMetricsRequestColumnNamesEnum("frame_delta"),
	DATA_FRAMES_RETRANSMITTED: RoCEv2FlowMetricsRequestColumnNamesEnum("data_frames_retransmitted"),
	FRAME_SEQUENCE_ERROR:      RoCEv2FlowMetricsRequestColumnNamesEnum("frame_sequence_error"),
	TX_BYTES:                  RoCEv2FlowMetricsRequestColumnNamesEnum("tx_bytes"),
	RX_BYTES:                  RoCEv2FlowMetricsRequestColumnNamesEnum("rx_bytes"),
	DATA_TX_RATE:              RoCEv2FlowMetricsRequestColumnNamesEnum("data_tx_rate"),
	DATA_RX_RATE:              RoCEv2FlowMetricsRequestColumnNamesEnum("data_rx_rate"),
	MESSAGE_TX:                RoCEv2FlowMetricsRequestColumnNamesEnum("message_tx"),
	MESSAGE_COMPLETE_RX:       RoCEv2FlowMetricsRequestColumnNamesEnum("message_complete_rx"),
	MESSAGE_FAIL:              RoCEv2FlowMetricsRequestColumnNamesEnum("message_fail"),
	FLOW_COMPLETION_TIME:      RoCEv2FlowMetricsRequestColumnNamesEnum("flow_completion_time"),
	AVG_LATESNCY:              RoCEv2FlowMetricsRequestColumnNamesEnum("avg_latesncy"),
	MIN_LATENCY:               RoCEv2FlowMetricsRequestColumnNamesEnum("min_latency"),
	MAX_LATENCY:               RoCEv2FlowMetricsRequestColumnNamesEnum("max_latency"),
	ECN_CE_RX:                 RoCEv2FlowMetricsRequestColumnNamesEnum("ecn_ce_rx"),
	CNP_TX:                    RoCEv2FlowMetricsRequestColumnNamesEnum("cnp_tx"),
	CNP_RX:                    RoCEv2FlowMetricsRequestColumnNamesEnum("cnp_rx"),
	ACK_TX:                    RoCEv2FlowMetricsRequestColumnNamesEnum("ack_tx"),
	ACK_RX:                    RoCEv2FlowMetricsRequestColumnNamesEnum("ack_rx"),
	NAK_TX:                    RoCEv2FlowMetricsRequestColumnNamesEnum("nak_tx"),
	NAK_RX:                    RoCEv2FlowMetricsRequestColumnNamesEnum("nak_rx"),
	FIRST_TIMESTAMP:           RoCEv2FlowMetricsRequestColumnNamesEnum("first_timestamp"),
	LAST_TIMESTAMP:            RoCEv2FlowMetricsRequestColumnNamesEnum("last_timestamp"),
}

func (obj *roCEv2FlowMetricsRequest) ColumnNames() []RoCEv2FlowMetricsRequestColumnNamesEnum {
	items := []RoCEv2FlowMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, RoCEv2FlowMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the Flow cannot be excluded.
// SetColumnNames sets the []string value in the RoCEv2FlowMetricsRequest object
func (obj *roCEv2FlowMetricsRequest) SetColumnNames(value []RoCEv2FlowMetricsRequestColumnNamesEnum) RoCEv2FlowMetricsRequest {

	items := []otg.RoCEv2FlowMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.RoCEv2FlowMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.RoCEv2FlowMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *roCEv2FlowMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *roCEv2FlowMetricsRequest) setDefault() {

}
