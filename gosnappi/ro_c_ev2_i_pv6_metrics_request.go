package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2IPv6MetricsRequest *****
type roCEv2IPv6MetricsRequest struct {
	validation
	obj          *otg.RoCEv2IPv6MetricsRequest
	marshaller   marshalRoCEv2IPv6MetricsRequest
	unMarshaller unMarshalRoCEv2IPv6MetricsRequest
}

func NewRoCEv2IPv6MetricsRequest() RoCEv2IPv6MetricsRequest {
	obj := roCEv2IPv6MetricsRequest{obj: &otg.RoCEv2IPv6MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2IPv6MetricsRequest) msg() *otg.RoCEv2IPv6MetricsRequest {
	return obj.obj
}

func (obj *roCEv2IPv6MetricsRequest) setMsg(msg *otg.RoCEv2IPv6MetricsRequest) RoCEv2IPv6MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2IPv6MetricsRequest struct {
	obj *roCEv2IPv6MetricsRequest
}

type marshalRoCEv2IPv6MetricsRequest interface {
	// ToProto marshals RoCEv2IPv6MetricsRequest to protobuf object *otg.RoCEv2IPv6MetricsRequest
	ToProto() (*otg.RoCEv2IPv6MetricsRequest, error)
	// ToPbText marshals RoCEv2IPv6MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2IPv6MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2IPv6MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2IPv6MetricsRequest struct {
	obj *roCEv2IPv6MetricsRequest
}

type unMarshalRoCEv2IPv6MetricsRequest interface {
	// FromProto unmarshals RoCEv2IPv6MetricsRequest from protobuf object *otg.RoCEv2IPv6MetricsRequest
	FromProto(msg *otg.RoCEv2IPv6MetricsRequest) (RoCEv2IPv6MetricsRequest, error)
	// FromPbText unmarshals RoCEv2IPv6MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2IPv6MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2IPv6MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *roCEv2IPv6MetricsRequest) Marshal() marshalRoCEv2IPv6MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2IPv6MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2IPv6MetricsRequest) Unmarshal() unMarshalRoCEv2IPv6MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2IPv6MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2IPv6MetricsRequest) ToProto() (*otg.RoCEv2IPv6MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2IPv6MetricsRequest) FromProto(msg *otg.RoCEv2IPv6MetricsRequest) (RoCEv2IPv6MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2IPv6MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2IPv6MetricsRequest) FromPbText(value string) error {
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

func (m *marshalroCEv2IPv6MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2IPv6MetricsRequest) FromYaml(value string) error {
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

func (m *marshalroCEv2IPv6MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalroCEv2IPv6MetricsRequest) FromJson(value string) error {
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

func (obj *roCEv2IPv6MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2IPv6MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2IPv6MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2IPv6MetricsRequest) Clone() (RoCEv2IPv6MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2IPv6MetricsRequest()
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

// RoCEv2IPv6MetricsRequest is the request to retrieve RoCEv2 over IPv6 per peer metrics/statistics.
type RoCEv2IPv6MetricsRequest interface {
	Validation
	// msg marshals RoCEv2IPv6MetricsRequest to protobuf object *otg.RoCEv2IPv6MetricsRequest
	// and doesn't set defaults
	msg() *otg.RoCEv2IPv6MetricsRequest
	// setMsg unmarshals RoCEv2IPv6MetricsRequest from protobuf object *otg.RoCEv2IPv6MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.RoCEv2IPv6MetricsRequest) RoCEv2IPv6MetricsRequest
	// provides marshal interface
	Marshal() marshalRoCEv2IPv6MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2IPv6MetricsRequest
	// validate validates RoCEv2IPv6MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2IPv6MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in RoCEv2IPv6MetricsRequest.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to RoCEv2IPv6MetricsRequest
	SetPeerNames(value []string) RoCEv2IPv6MetricsRequest
	// ColumnNames returns []RoCEv2IPv6MetricsRequestColumnNamesEnum, set in RoCEv2IPv6MetricsRequest
	ColumnNames() []RoCEv2IPv6MetricsRequestColumnNamesEnum
	// SetColumnNames assigns []RoCEv2IPv6MetricsRequestColumnNamesEnum provided by user to RoCEv2IPv6MetricsRequest
	SetColumnNames(value []RoCEv2IPv6MetricsRequestColumnNamesEnum) RoCEv2IPv6MetricsRequest
}

// The names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/RoCEv2.V6peer/properties/name
//
// PeerNames returns a []string
func (obj *roCEv2IPv6MetricsRequest) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/RoCEv2.V6peer/properties/name
//
// SetPeerNames sets the []string value in the RoCEv2IPv6MetricsRequest object
func (obj *roCEv2IPv6MetricsRequest) SetPeerNames(value []string) RoCEv2IPv6MetricsRequest {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type RoCEv2IPv6MetricsRequestColumnNamesEnum string

// Enum of ColumnNames on RoCEv2IPv6MetricsRequest
var RoCEv2IPv6MetricsRequestColumnNames = struct {
	SESSION_STATE         RoCEv2IPv6MetricsRequestColumnNamesEnum
	QP_CONFUGRED          RoCEv2IPv6MetricsRequestColumnNamesEnum
	QP_UP                 RoCEv2IPv6MetricsRequestColumnNamesEnum
	QP_DOWN               RoCEv2IPv6MetricsRequestColumnNamesEnum
	CONNECT_REQUEST_TX    RoCEv2IPv6MetricsRequestColumnNamesEnum
	CONNECT_REQUEST_RX    RoCEv2IPv6MetricsRequestColumnNamesEnum
	CONNECT_REPLY_TX      RoCEv2IPv6MetricsRequestColumnNamesEnum
	CONNECT_REPLY_RX      RoCEv2IPv6MetricsRequestColumnNamesEnum
	READY_TX              RoCEv2IPv6MetricsRequestColumnNamesEnum
	READY_RX              RoCEv2IPv6MetricsRequestColumnNamesEnum
	DISCONNECT_REQUEST_TX RoCEv2IPv6MetricsRequestColumnNamesEnum
	DISCONNECT_REQUEST_RX RoCEv2IPv6MetricsRequestColumnNamesEnum
	DISCONNECT_REPLY_TX   RoCEv2IPv6MetricsRequestColumnNamesEnum
	DISCONNECT_REPLY_RX   RoCEv2IPv6MetricsRequestColumnNamesEnum
	REJECT_TX             RoCEv2IPv6MetricsRequestColumnNamesEnum
	REJECT_RX             RoCEv2IPv6MetricsRequestColumnNamesEnum
	UNKNOWN_MSG_RX        RoCEv2IPv6MetricsRequestColumnNamesEnum
}{
	SESSION_STATE:         RoCEv2IPv6MetricsRequestColumnNamesEnum("session_state"),
	QP_CONFUGRED:          RoCEv2IPv6MetricsRequestColumnNamesEnum("qp_confugred"),
	QP_UP:                 RoCEv2IPv6MetricsRequestColumnNamesEnum("qp_up"),
	QP_DOWN:               RoCEv2IPv6MetricsRequestColumnNamesEnum("qp_down"),
	CONNECT_REQUEST_TX:    RoCEv2IPv6MetricsRequestColumnNamesEnum("connect_request_tx"),
	CONNECT_REQUEST_RX:    RoCEv2IPv6MetricsRequestColumnNamesEnum("connect_request_rx"),
	CONNECT_REPLY_TX:      RoCEv2IPv6MetricsRequestColumnNamesEnum("connect_reply_tx"),
	CONNECT_REPLY_RX:      RoCEv2IPv6MetricsRequestColumnNamesEnum("connect_reply_rx"),
	READY_TX:              RoCEv2IPv6MetricsRequestColumnNamesEnum("ready_tx"),
	READY_RX:              RoCEv2IPv6MetricsRequestColumnNamesEnum("ready_rx"),
	DISCONNECT_REQUEST_TX: RoCEv2IPv6MetricsRequestColumnNamesEnum("disconnect_request_tx"),
	DISCONNECT_REQUEST_RX: RoCEv2IPv6MetricsRequestColumnNamesEnum("disconnect_request_rx"),
	DISCONNECT_REPLY_TX:   RoCEv2IPv6MetricsRequestColumnNamesEnum("disconnect_reply_tx"),
	DISCONNECT_REPLY_RX:   RoCEv2IPv6MetricsRequestColumnNamesEnum("disconnect_reply_rx"),
	REJECT_TX:             RoCEv2IPv6MetricsRequestColumnNamesEnum("reject_tx"),
	REJECT_RX:             RoCEv2IPv6MetricsRequestColumnNamesEnum("reject_rx"),
	UNKNOWN_MSG_RX:        RoCEv2IPv6MetricsRequestColumnNamesEnum("unknown_msg_rx"),
}

func (obj *roCEv2IPv6MetricsRequest) ColumnNames() []RoCEv2IPv6MetricsRequestColumnNamesEnum {
	items := []RoCEv2IPv6MetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, RoCEv2IPv6MetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the RoCEv2 peer cannot be excluded.
// SetColumnNames sets the []string value in the RoCEv2IPv6MetricsRequest object
func (obj *roCEv2IPv6MetricsRequest) SetColumnNames(value []RoCEv2IPv6MetricsRequestColumnNamesEnum) RoCEv2IPv6MetricsRequest {

	items := []otg.RoCEv2IPv6MetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.RoCEv2IPv6MetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.RoCEv2IPv6MetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *roCEv2IPv6MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *roCEv2IPv6MetricsRequest) setDefault() {

}
