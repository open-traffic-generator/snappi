package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetMetricsRequest *****
type ultraEthernetMetricsRequest struct {
	validation
	obj          *otg.UltraEthernetMetricsRequest
	marshaller   marshalUltraEthernetMetricsRequest
	unMarshaller unMarshalUltraEthernetMetricsRequest
}

func NewUltraEthernetMetricsRequest() UltraEthernetMetricsRequest {
	obj := ultraEthernetMetricsRequest{obj: &otg.UltraEthernetMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetMetricsRequest) msg() *otg.UltraEthernetMetricsRequest {
	return obj.obj
}

func (obj *ultraEthernetMetricsRequest) setMsg(msg *otg.UltraEthernetMetricsRequest) UltraEthernetMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetMetricsRequest struct {
	obj *ultraEthernetMetricsRequest
}

type marshalUltraEthernetMetricsRequest interface {
	// ToProto marshals UltraEthernetMetricsRequest to protobuf object *otg.UltraEthernetMetricsRequest
	ToProto() (*otg.UltraEthernetMetricsRequest, error)
	// ToPbText marshals UltraEthernetMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetMetricsRequest struct {
	obj *ultraEthernetMetricsRequest
}

type unMarshalUltraEthernetMetricsRequest interface {
	// FromProto unmarshals UltraEthernetMetricsRequest from protobuf object *otg.UltraEthernetMetricsRequest
	FromProto(msg *otg.UltraEthernetMetricsRequest) (UltraEthernetMetricsRequest, error)
	// FromPbText unmarshals UltraEthernetMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetMetricsRequest) Marshal() marshalUltraEthernetMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetMetricsRequest) Unmarshal() unMarshalUltraEthernetMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetMetricsRequest) ToProto() (*otg.UltraEthernetMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetMetricsRequest) FromProto(msg *otg.UltraEthernetMetricsRequest) (UltraEthernetMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetMetricsRequest) FromPbText(value string) error {
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

func (m *marshalultraEthernetMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetMetricsRequest) FromYaml(value string) error {
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

func (m *marshalultraEthernetMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetMetricsRequest) FromJson(value string) error {
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

func (obj *ultraEthernetMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetMetricsRequest) Clone() (UltraEthernetMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetMetricsRequest()
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

// UltraEthernetMetricsRequest is the request to retrieve Ultra Ethernet (UE) per port metrics/statistics covering the Link Layer Retry (LLR), Credit-based Flow Control (CBFC) and PHY Control Ordered Set (CtlOS) features.
type UltraEthernetMetricsRequest interface {
	Validation
	// msg marshals UltraEthernetMetricsRequest to protobuf object *otg.UltraEthernetMetricsRequest
	// and doesn't set defaults
	msg() *otg.UltraEthernetMetricsRequest
	// setMsg unmarshals UltraEthernetMetricsRequest from protobuf object *otg.UltraEthernetMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetMetricsRequest) UltraEthernetMetricsRequest
	// provides marshal interface
	Marshal() marshalUltraEthernetMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetMetricsRequest
	// validate validates UltraEthernetMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// UltraEthernetNames returns []string, set in UltraEthernetMetricsRequest.
	UltraEthernetNames() []string
	// SetUltraEthernetNames assigns []string provided by user to UltraEthernetMetricsRequest
	SetUltraEthernetNames(value []string) UltraEthernetMetricsRequest
	// ColumnNames returns []UltraEthernetMetricsRequestColumnNamesEnum, set in UltraEthernetMetricsRequest
	ColumnNames() []UltraEthernetMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []UltraEthernetMetricsRequestColumnNamesEnum provided by user to UltraEthernetMetricsRequest
	SetColumnNames(value []UltraEthernetMetricsRequestColumnNamesEnum) UltraEthernetMetricsRequest
}

// The names of Ultra Ethernet configurations to return results for. An empty list will return results for all Ultra Ethernet configurations.
//
// x-constraint:
// - /components/schemas/UltraEthernet/properties/name
//
// UltraEthernetNames returns a []string
func (obj *ultraEthernetMetricsRequest) UltraEthernetNames() []string {
	if obj.obj.UltraEthernetNames == nil {
		obj.obj.UltraEthernetNames = make([]string, 0)
	}
	return obj.obj.UltraEthernetNames
}

// The names of Ultra Ethernet configurations to return results for. An empty list will return results for all Ultra Ethernet configurations.
//
// x-constraint:
// - /components/schemas/UltraEthernet/properties/name
//
// SetUltraEthernetNames sets the []string value in the UltraEthernetMetricsRequest object
func (obj *ultraEthernetMetricsRequest) SetUltraEthernetNames(value []string) UltraEthernetMetricsRequest {

	if obj.obj.UltraEthernetNames == nil {
		obj.obj.UltraEthernetNames = make([]string, 0)
	}
	obj.obj.UltraEthernetNames = value

	return obj
}

type UltraEthernetMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on UltraEthernetMetricsRequest
var UltraEthernetMetricsRequestColumnNames = struct {
	TX_INIT_CTL_OS                         UltraEthernetMetricsRequestColumnNamesEnum
	TX_INIT_ECHO_CTL_OS                    UltraEthernetMetricsRequestColumnNamesEnum
	TX_ACK_CTL_OS                          UltraEthernetMetricsRequestColumnNamesEnum
	TX_NACK_CTL_OS                         UltraEthernetMetricsRequestColumnNamesEnum
	TX_DISCARD                             UltraEthernetMetricsRequestColumnNamesEnum
	TX_OK                                  UltraEthernetMetricsRequestColumnNamesEnum
	TX_POISONED                            UltraEthernetMetricsRequestColumnNamesEnum
	TX_REPLAY_EVENT                        UltraEthernetMetricsRequestColumnNamesEnum
	REPLAYED_PACKET                        UltraEthernetMetricsRequestColumnNamesEnum
	REPLAYED_BYTE                          UltraEthernetMetricsRequestColumnNamesEnum
	TX_SEQ                                 UltraEthernetMetricsRequestColumnNamesEnum
	TX_OUTSTANDING_SEQ                     UltraEthernetMetricsRequestColumnNamesEnum
	RX_INIT_CTL_OS                         UltraEthernetMetricsRequestColumnNamesEnum
	RX_INIT_ECHO_CTL_OS                    UltraEthernetMetricsRequestColumnNamesEnum
	RX_ACK_CTL_OS                          UltraEthernetMetricsRequestColumnNamesEnum
	RX_NACK_CTL_OS                         UltraEthernetMetricsRequestColumnNamesEnum
	RX_ACK_NACK_SEQ_ERROR                  UltraEthernetMetricsRequestColumnNamesEnum
	RX_OK                                  UltraEthernetMetricsRequestColumnNamesEnum
	RX_POISONED                            UltraEthernetMetricsRequestColumnNamesEnum
	RX_BAD                                 UltraEthernetMetricsRequestColumnNamesEnum
	RX_EXPECTED_SEQ_GOOD                   UltraEthernetMetricsRequestColumnNamesEnum
	RX_EXPECTED_SEQ_POISONED               UltraEthernetMetricsRequestColumnNamesEnum
	RX_EXPECTED_SEQ_BAD                    UltraEthernetMetricsRequestColumnNamesEnum
	RX_MISSING_SEQ                         UltraEthernetMetricsRequestColumnNamesEnum
	RX_DUPLICATE_SEQ                       UltraEthernetMetricsRequestColumnNamesEnum
	RX_REPLAY                              UltraEthernetMetricsRequestColumnNamesEnum
	RX_NEXT_SEQ                            UltraEthernetMetricsRequestColumnNamesEnum
	LLR_INIT_CTL_OS_SPACING_MIN            UltraEthernetMetricsRequestColumnNamesEnum
	LLR_INIT_CTL_OS_SPACING_ERROR          UltraEthernetMetricsRequestColumnNamesEnum
	LLR_ACK_NACK_CTL_OS_SPACING_MIN        UltraEthernetMetricsRequestColumnNamesEnum
	LLR_ACK_NACK_CTL_OS_SPACING_ERROR      UltraEthernetMetricsRequestColumnNamesEnum
	LLR_INIT_ECHO_INIT_SEQ_MISMATCH        UltraEthernetMetricsRequestColumnNamesEnum
	RX_ACK_CTL_OS_DROPPED                  UltraEthernetMetricsRequestColumnNamesEnum
	RX_NACK_CTL_OS_DROPPED                 UltraEthernetMetricsRequestColumnNamesEnum
	RX_INIT_CTL_OS_DROPPED                 UltraEthernetMetricsRequestColumnNamesEnum
	RX_INIT_ECHO_CTL_OS_DROPPED            UltraEthernetMetricsRequestColumnNamesEnum
	UE_RX_CTL_OS_FRAME_HEADER_ERROR        UltraEthernetMetricsRequestColumnNamesEnum
	UE_RX_CTL_OS_INTRA_FRAME_SPACING_ERROR UltraEthernetMetricsRequestColumnNamesEnum
	UE_CTL_OS_SPACING_MIN                  UltraEthernetMetricsRequestColumnNamesEnum
	UE_CTL_OS_SPACING_ERROR                UltraEthernetMetricsRequestColumnNamesEnum
	LLR_MODE_LOCAL                         UltraEthernetMetricsRequestColumnNamesEnum
	LLR_MODE_REMOTE                        UltraEthernetMetricsRequestColumnNamesEnum
	LLR_TRANSMIT_STATE                     UltraEthernetMetricsRequestColumnNamesEnum
	LLR_ACK_NACK_TRANSMIT_STATE            UltraEthernetMetricsRequestColumnNamesEnum
	ROUND_TRIP_TIME_NS                     UltraEthernetMetricsRequestColumnNamesEnum
	MEAN_TIME_BETWEEN_PHY_ERRORS           UltraEthernetMetricsRequestColumnNamesEnum
	ROUND_TRIP_TIME_VALID                  UltraEthernetMetricsRequestColumnNamesEnum
}{
	TX_INIT_CTL_OS:                         UltraEthernetMetricsRequestColumnNamesEnum("tx_init_ctl_os"),
	TX_INIT_ECHO_CTL_OS:                    UltraEthernetMetricsRequestColumnNamesEnum("tx_init_echo_ctl_os"),
	TX_ACK_CTL_OS:                          UltraEthernetMetricsRequestColumnNamesEnum("tx_ack_ctl_os"),
	TX_NACK_CTL_OS:                         UltraEthernetMetricsRequestColumnNamesEnum("tx_nack_ctl_os"),
	TX_DISCARD:                             UltraEthernetMetricsRequestColumnNamesEnum("tx_discard"),
	TX_OK:                                  UltraEthernetMetricsRequestColumnNamesEnum("tx_ok"),
	TX_POISONED:                            UltraEthernetMetricsRequestColumnNamesEnum("tx_poisoned"),
	TX_REPLAY_EVENT:                        UltraEthernetMetricsRequestColumnNamesEnum("tx_replay_event"),
	REPLAYED_PACKET:                        UltraEthernetMetricsRequestColumnNamesEnum("replayed_packet"),
	REPLAYED_BYTE:                          UltraEthernetMetricsRequestColumnNamesEnum("replayed_byte"),
	TX_SEQ:                                 UltraEthernetMetricsRequestColumnNamesEnum("tx_seq"),
	TX_OUTSTANDING_SEQ:                     UltraEthernetMetricsRequestColumnNamesEnum("tx_outstanding_seq"),
	RX_INIT_CTL_OS:                         UltraEthernetMetricsRequestColumnNamesEnum("rx_init_ctl_os"),
	RX_INIT_ECHO_CTL_OS:                    UltraEthernetMetricsRequestColumnNamesEnum("rx_init_echo_ctl_os"),
	RX_ACK_CTL_OS:                          UltraEthernetMetricsRequestColumnNamesEnum("rx_ack_ctl_os"),
	RX_NACK_CTL_OS:                         UltraEthernetMetricsRequestColumnNamesEnum("rx_nack_ctl_os"),
	RX_ACK_NACK_SEQ_ERROR:                  UltraEthernetMetricsRequestColumnNamesEnum("rx_ack_nack_seq_error"),
	RX_OK:                                  UltraEthernetMetricsRequestColumnNamesEnum("rx_ok"),
	RX_POISONED:                            UltraEthernetMetricsRequestColumnNamesEnum("rx_poisoned"),
	RX_BAD:                                 UltraEthernetMetricsRequestColumnNamesEnum("rx_bad"),
	RX_EXPECTED_SEQ_GOOD:                   UltraEthernetMetricsRequestColumnNamesEnum("rx_expected_seq_good"),
	RX_EXPECTED_SEQ_POISONED:               UltraEthernetMetricsRequestColumnNamesEnum("rx_expected_seq_poisoned"),
	RX_EXPECTED_SEQ_BAD:                    UltraEthernetMetricsRequestColumnNamesEnum("rx_expected_seq_bad"),
	RX_MISSING_SEQ:                         UltraEthernetMetricsRequestColumnNamesEnum("rx_missing_seq"),
	RX_DUPLICATE_SEQ:                       UltraEthernetMetricsRequestColumnNamesEnum("rx_duplicate_seq"),
	RX_REPLAY:                              UltraEthernetMetricsRequestColumnNamesEnum("rx_replay"),
	RX_NEXT_SEQ:                            UltraEthernetMetricsRequestColumnNamesEnum("rx_next_seq"),
	LLR_INIT_CTL_OS_SPACING_MIN:            UltraEthernetMetricsRequestColumnNamesEnum("llr_init_ctl_os_spacing_min"),
	LLR_INIT_CTL_OS_SPACING_ERROR:          UltraEthernetMetricsRequestColumnNamesEnum("llr_init_ctl_os_spacing_error"),
	LLR_ACK_NACK_CTL_OS_SPACING_MIN:        UltraEthernetMetricsRequestColumnNamesEnum("llr_ack_nack_ctl_os_spacing_min"),
	LLR_ACK_NACK_CTL_OS_SPACING_ERROR:      UltraEthernetMetricsRequestColumnNamesEnum("llr_ack_nack_ctl_os_spacing_error"),
	LLR_INIT_ECHO_INIT_SEQ_MISMATCH:        UltraEthernetMetricsRequestColumnNamesEnum("llr_init_echo_init_seq_mismatch"),
	RX_ACK_CTL_OS_DROPPED:                  UltraEthernetMetricsRequestColumnNamesEnum("rx_ack_ctl_os_dropped"),
	RX_NACK_CTL_OS_DROPPED:                 UltraEthernetMetricsRequestColumnNamesEnum("rx_nack_ctl_os_dropped"),
	RX_INIT_CTL_OS_DROPPED:                 UltraEthernetMetricsRequestColumnNamesEnum("rx_init_ctl_os_dropped"),
	RX_INIT_ECHO_CTL_OS_DROPPED:            UltraEthernetMetricsRequestColumnNamesEnum("rx_init_echo_ctl_os_dropped"),
	UE_RX_CTL_OS_FRAME_HEADER_ERROR:        UltraEthernetMetricsRequestColumnNamesEnum("ue_rx_ctl_os_frame_header_error"),
	UE_RX_CTL_OS_INTRA_FRAME_SPACING_ERROR: UltraEthernetMetricsRequestColumnNamesEnum("ue_rx_ctl_os_intra_frame_spacing_error"),
	UE_CTL_OS_SPACING_MIN:                  UltraEthernetMetricsRequestColumnNamesEnum("ue_ctl_os_spacing_min"),
	UE_CTL_OS_SPACING_ERROR:                UltraEthernetMetricsRequestColumnNamesEnum("ue_ctl_os_spacing_error"),
	LLR_MODE_LOCAL:                         UltraEthernetMetricsRequestColumnNamesEnum("llr_mode_local"),
	LLR_MODE_REMOTE:                        UltraEthernetMetricsRequestColumnNamesEnum("llr_mode_remote"),
	LLR_TRANSMIT_STATE:                     UltraEthernetMetricsRequestColumnNamesEnum("llr_transmit_state"),
	LLR_ACK_NACK_TRANSMIT_STATE:            UltraEthernetMetricsRequestColumnNamesEnum("llr_ack_nack_transmit_state"),
	ROUND_TRIP_TIME_NS:                     UltraEthernetMetricsRequestColumnNamesEnum("round_trip_time_ns"),
	MEAN_TIME_BETWEEN_PHY_ERRORS:           UltraEthernetMetricsRequestColumnNamesEnum("mean_time_between_phy_errors"),
	ROUND_TRIP_TIME_VALID:                  UltraEthernetMetricsRequestColumnNamesEnum("round_trip_time_valid"),
}

func (obj *ultraEthernetMetricsRequest) ColumnNames() []UltraEthernetMetricsRequestColumnNamesEnum {
	items := []UltraEthernetMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, UltraEthernetMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The requested list of column names for the result set. If the list is empty then metrics for all columns will be returned. The name of the Ultra Ethernet configuration can not be excluded. The cbfc_sender_virtual_channels and cbfc_receiver_virtual_channels per-VC array metrics are always returned and cannot be individually selected or excluded via column_names.
// SetColumnNames sets the []string value in the UltraEthernetMetricsRequest object
func (obj *ultraEthernetMetricsRequest) SetColumnNames(value []UltraEthernetMetricsRequestColumnNamesEnum) UltraEthernetMetricsRequest {

	items := []otg.UltraEthernetMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.UltraEthernetMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.UltraEthernetMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *ultraEthernetMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ultraEthernetMetricsRequest) setDefault() {

}
