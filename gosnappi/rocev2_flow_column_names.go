package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2FlowColumnNames *****
type rocev2FlowColumnNames struct {
	validation
	obj          *otg.Rocev2FlowColumnNames
	marshaller   marshalRocev2FlowColumnNames
	unMarshaller unMarshalRocev2FlowColumnNames
}

func NewRocev2FlowColumnNames() Rocev2FlowColumnNames {
	obj := rocev2FlowColumnNames{obj: &otg.Rocev2FlowColumnNames{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2FlowColumnNames) msg() *otg.Rocev2FlowColumnNames {
	return obj.obj
}

func (obj *rocev2FlowColumnNames) setMsg(msg *otg.Rocev2FlowColumnNames) Rocev2FlowColumnNames {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2FlowColumnNames struct {
	obj *rocev2FlowColumnNames
}

type marshalRocev2FlowColumnNames interface {
	// ToProto marshals Rocev2FlowColumnNames to protobuf object *otg.Rocev2FlowColumnNames
	ToProto() (*otg.Rocev2FlowColumnNames, error)
	// ToPbText marshals Rocev2FlowColumnNames to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2FlowColumnNames to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2FlowColumnNames to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2FlowColumnNames to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2FlowColumnNames struct {
	obj *rocev2FlowColumnNames
}

type unMarshalRocev2FlowColumnNames interface {
	// FromProto unmarshals Rocev2FlowColumnNames from protobuf object *otg.Rocev2FlowColumnNames
	FromProto(msg *otg.Rocev2FlowColumnNames) (Rocev2FlowColumnNames, error)
	// FromPbText unmarshals Rocev2FlowColumnNames from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2FlowColumnNames from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2FlowColumnNames from JSON text
	FromJson(value string) error
}

func (obj *rocev2FlowColumnNames) Marshal() marshalRocev2FlowColumnNames {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2FlowColumnNames{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2FlowColumnNames) Unmarshal() unMarshalRocev2FlowColumnNames {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2FlowColumnNames{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2FlowColumnNames) ToProto() (*otg.Rocev2FlowColumnNames, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2FlowColumnNames) FromProto(msg *otg.Rocev2FlowColumnNames) (Rocev2FlowColumnNames, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2FlowColumnNames) ToPbText() (string, error) {
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

func (m *unMarshalrocev2FlowColumnNames) FromPbText(value string) error {
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

func (m *marshalrocev2FlowColumnNames) ToYaml() (string, error) {
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

func (m *unMarshalrocev2FlowColumnNames) FromYaml(value string) error {
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

func (m *marshalrocev2FlowColumnNames) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2FlowColumnNames) ToJson() (string, error) {
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

func (m *unMarshalrocev2FlowColumnNames) FromJson(value string) error {
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

func (obj *rocev2FlowColumnNames) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2FlowColumnNames) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2FlowColumnNames) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2FlowColumnNames) Clone() (Rocev2FlowColumnNames, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2FlowColumnNames()
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

// Rocev2FlowColumnNames is the names of RoCEv2 flows. An empty list will return results for all RoCEv2 flows.
type Rocev2FlowColumnNames interface {
	Validation
	// msg marshals Rocev2FlowColumnNames to protobuf object *otg.Rocev2FlowColumnNames
	// and doesn't set defaults
	msg() *otg.Rocev2FlowColumnNames
	// setMsg unmarshals Rocev2FlowColumnNames from protobuf object *otg.Rocev2FlowColumnNames
	// and doesn't set defaults
	setMsg(*otg.Rocev2FlowColumnNames) Rocev2FlowColumnNames
	// provides marshal interface
	Marshal() marshalRocev2FlowColumnNames
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2FlowColumnNames
	// validate validates Rocev2FlowColumnNames
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2FlowColumnNames, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ColumnNames returns []Rocev2FlowColumnNamesColumnNamesEnum, set in Rocev2FlowColumnNames
	ColumnNames() []Rocev2FlowColumnNamesColumnNamesEnum
	// SetColumnNames assigns []Rocev2FlowColumnNamesColumnNamesEnum provided by user to Rocev2FlowColumnNames
	SetColumnNames(value []Rocev2FlowColumnNamesColumnNamesEnum) Rocev2FlowColumnNames
}

type Rocev2FlowColumnNamesColumnNamesEnum string

// Enum of ColumnNames on Rocev2FlowColumnNames
var Rocev2FlowColumnNamesColumnNames = struct {
	FLOW_NAME                 Rocev2FlowColumnNamesColumnNamesEnum
	PORT_TX                   Rocev2FlowColumnNamesColumnNamesEnum
	PORT_RX                   Rocev2FlowColumnNamesColumnNamesEnum
	SRC_QP                    Rocev2FlowColumnNamesColumnNamesEnum
	DEST_QP                   Rocev2FlowColumnNamesColumnNamesEnum
	SRC_IPV4                  Rocev2FlowColumnNamesColumnNamesEnum
	DEST_IPV4                 Rocev2FlowColumnNamesColumnNamesEnum
	DATA_FRAMES_TX            Rocev2FlowColumnNamesColumnNamesEnum
	DATA_FRAMES_RX            Rocev2FlowColumnNamesColumnNamesEnum
	FRAME_DELTA               Rocev2FlowColumnNamesColumnNamesEnum
	DATA_FRAMES_RETRANSMITTED Rocev2FlowColumnNamesColumnNamesEnum
	FRAME_SEQUENCE_ERROR      Rocev2FlowColumnNamesColumnNamesEnum
	TX_BYTES                  Rocev2FlowColumnNamesColumnNamesEnum
	RX_BYTES                  Rocev2FlowColumnNamesColumnNamesEnum
	DATA_TX_RATE              Rocev2FlowColumnNamesColumnNamesEnum
	DATA_RX_RATE              Rocev2FlowColumnNamesColumnNamesEnum
	MESSAGE_TX                Rocev2FlowColumnNamesColumnNamesEnum
	MESSAGE_COMPLETE_RX       Rocev2FlowColumnNamesColumnNamesEnum
	MESSAGE_FAIL              Rocev2FlowColumnNamesColumnNamesEnum
	FLOW_COMPLETION_TIME      Rocev2FlowColumnNamesColumnNamesEnum
	AVG_LATENCY               Rocev2FlowColumnNamesColumnNamesEnum
	MIN_LATENCY               Rocev2FlowColumnNamesColumnNamesEnum
	MAX_LATENCY               Rocev2FlowColumnNamesColumnNamesEnum
	ECN_CE_RX                 Rocev2FlowColumnNamesColumnNamesEnum
	CNP_TX                    Rocev2FlowColumnNamesColumnNamesEnum
	CNP_RX                    Rocev2FlowColumnNamesColumnNamesEnum
	ACK_TX                    Rocev2FlowColumnNamesColumnNamesEnum
	ACK_RX                    Rocev2FlowColumnNamesColumnNamesEnum
	NAK_TX                    Rocev2FlowColumnNamesColumnNamesEnum
	NAK_RX                    Rocev2FlowColumnNamesColumnNamesEnum
	FIRST_TIMESTAMP           Rocev2FlowColumnNamesColumnNamesEnum
	LAST_TIMESTAMP            Rocev2FlowColumnNamesColumnNamesEnum
}{
	FLOW_NAME:                 Rocev2FlowColumnNamesColumnNamesEnum("flow_name"),
	PORT_TX:                   Rocev2FlowColumnNamesColumnNamesEnum("port_tx"),
	PORT_RX:                   Rocev2FlowColumnNamesColumnNamesEnum("port_rx"),
	SRC_QP:                    Rocev2FlowColumnNamesColumnNamesEnum("src_qp"),
	DEST_QP:                   Rocev2FlowColumnNamesColumnNamesEnum("dest_qp"),
	SRC_IPV4:                  Rocev2FlowColumnNamesColumnNamesEnum("src_ipv4"),
	DEST_IPV4:                 Rocev2FlowColumnNamesColumnNamesEnum("dest_ipv4"),
	DATA_FRAMES_TX:            Rocev2FlowColumnNamesColumnNamesEnum("data_frames_tx"),
	DATA_FRAMES_RX:            Rocev2FlowColumnNamesColumnNamesEnum("data_frames_rx"),
	FRAME_DELTA:               Rocev2FlowColumnNamesColumnNamesEnum("frame_delta"),
	DATA_FRAMES_RETRANSMITTED: Rocev2FlowColumnNamesColumnNamesEnum("data_frames_retransmitted"),
	FRAME_SEQUENCE_ERROR:      Rocev2FlowColumnNamesColumnNamesEnum("frame_sequence_error"),
	TX_BYTES:                  Rocev2FlowColumnNamesColumnNamesEnum("tx_bytes"),
	RX_BYTES:                  Rocev2FlowColumnNamesColumnNamesEnum("rx_bytes"),
	DATA_TX_RATE:              Rocev2FlowColumnNamesColumnNamesEnum("data_tx_rate"),
	DATA_RX_RATE:              Rocev2FlowColumnNamesColumnNamesEnum("data_rx_rate"),
	MESSAGE_TX:                Rocev2FlowColumnNamesColumnNamesEnum("message_tx"),
	MESSAGE_COMPLETE_RX:       Rocev2FlowColumnNamesColumnNamesEnum("message_complete_rx"),
	MESSAGE_FAIL:              Rocev2FlowColumnNamesColumnNamesEnum("message_fail"),
	FLOW_COMPLETION_TIME:      Rocev2FlowColumnNamesColumnNamesEnum("flow_completion_time"),
	AVG_LATENCY:               Rocev2FlowColumnNamesColumnNamesEnum("avg_latency"),
	MIN_LATENCY:               Rocev2FlowColumnNamesColumnNamesEnum("min_latency"),
	MAX_LATENCY:               Rocev2FlowColumnNamesColumnNamesEnum("max_latency"),
	ECN_CE_RX:                 Rocev2FlowColumnNamesColumnNamesEnum("ecn_ce_rx"),
	CNP_TX:                    Rocev2FlowColumnNamesColumnNamesEnum("cnp_tx"),
	CNP_RX:                    Rocev2FlowColumnNamesColumnNamesEnum("cnp_rx"),
	ACK_TX:                    Rocev2FlowColumnNamesColumnNamesEnum("ack_tx"),
	ACK_RX:                    Rocev2FlowColumnNamesColumnNamesEnum("ack_rx"),
	NAK_TX:                    Rocev2FlowColumnNamesColumnNamesEnum("nak_tx"),
	NAK_RX:                    Rocev2FlowColumnNamesColumnNamesEnum("nak_rx"),
	FIRST_TIMESTAMP:           Rocev2FlowColumnNamesColumnNamesEnum("first_timestamp"),
	LAST_TIMESTAMP:            Rocev2FlowColumnNamesColumnNamesEnum("last_timestamp"),
}

func (obj *rocev2FlowColumnNames) ColumnNames() []Rocev2FlowColumnNamesColumnNamesEnum {
	items := []Rocev2FlowColumnNamesColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Rocev2FlowColumnNamesColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the Flow cannot be excluded.
// SetColumnNames sets the []string value in the Rocev2FlowColumnNames object
func (obj *rocev2FlowColumnNames) SetColumnNames(value []Rocev2FlowColumnNamesColumnNamesEnum) Rocev2FlowColumnNames {

	items := []otg.Rocev2FlowColumnNames_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Rocev2FlowColumnNames_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Rocev2FlowColumnNames_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *rocev2FlowColumnNames) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rocev2FlowColumnNames) setDefault() {

}
