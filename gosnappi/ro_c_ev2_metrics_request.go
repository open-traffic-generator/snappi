package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2MetricsRequest *****
type roCEv2MetricsRequest struct {
	validation
	obj          *otg.RoCEv2MetricsRequest
	marshaller   marshalRoCEv2MetricsRequest
	unMarshaller unMarshalRoCEv2MetricsRequest
}

func NewRoCEv2MetricsRequest() RoCEv2MetricsRequest {
	obj := roCEv2MetricsRequest{obj: &otg.RoCEv2MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2MetricsRequest) msg() *otg.RoCEv2MetricsRequest {
	return obj.obj
}

func (obj *roCEv2MetricsRequest) setMsg(msg *otg.RoCEv2MetricsRequest) RoCEv2MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2MetricsRequest struct {
	obj *roCEv2MetricsRequest
}

type marshalRoCEv2MetricsRequest interface {
	// ToProto marshals RoCEv2MetricsRequest to protobuf object *otg.RoCEv2MetricsRequest
	ToProto() (*otg.RoCEv2MetricsRequest, error)
	// ToPbText marshals RoCEv2MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2MetricsRequest struct {
	obj *roCEv2MetricsRequest
}

type unMarshalRoCEv2MetricsRequest interface {
	// FromProto unmarshals RoCEv2MetricsRequest from protobuf object *otg.RoCEv2MetricsRequest
	FromProto(msg *otg.RoCEv2MetricsRequest) (RoCEv2MetricsRequest, error)
	// FromPbText unmarshals RoCEv2MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *roCEv2MetricsRequest) Marshal() marshalRoCEv2MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2MetricsRequest) Unmarshal() unMarshalRoCEv2MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2MetricsRequest) ToProto() (*otg.RoCEv2MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2MetricsRequest) FromProto(msg *otg.RoCEv2MetricsRequest) (RoCEv2MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2MetricsRequest) FromPbText(value string) error {
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

func (m *marshalroCEv2MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2MetricsRequest) FromYaml(value string) error {
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

func (m *marshalroCEv2MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalroCEv2MetricsRequest) FromJson(value string) error {
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

func (obj *roCEv2MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2MetricsRequest) Clone() (RoCEv2MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2MetricsRequest()
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

// RoCEv2MetricsRequest is the request to retrieve RoCEv2 over IPv4 per peer metrics/statistics.
type RoCEv2MetricsRequest interface {
	Validation
	// msg marshals RoCEv2MetricsRequest to protobuf object *otg.RoCEv2MetricsRequest
	// and doesn't set defaults
	msg() *otg.RoCEv2MetricsRequest
	// setMsg unmarshals RoCEv2MetricsRequest from protobuf object *otg.RoCEv2MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.RoCEv2MetricsRequest) RoCEv2MetricsRequest
	// provides marshal interface
	Marshal() marshalRoCEv2MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2MetricsRequest
	// validate validates RoCEv2MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in RoCEv2MetricsRequest.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to RoCEv2MetricsRequest
	SetPeerNames(value []string) RoCEv2MetricsRequest
	// ColumnNames returns []RoCEv2MetricsRequestColumnNamesEnum, set in RoCEv2MetricsRequest
	ColumnNames() []RoCEv2MetricsRequestColumnNamesEnum
	// SetColumnNames assigns []RoCEv2MetricsRequestColumnNamesEnum provided by user to RoCEv2MetricsRequest
	SetColumnNames(value []RoCEv2MetricsRequestColumnNamesEnum) RoCEv2MetricsRequest
}

// The names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/RoCEv2.V4peer/properties/name
//
// PeerNames returns a []string
func (obj *roCEv2MetricsRequest) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/RoCEv2.V4peer/properties/name
//
// SetPeerNames sets the []string value in the RoCEv2MetricsRequest object
func (obj *roCEv2MetricsRequest) SetPeerNames(value []string) RoCEv2MetricsRequest {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type RoCEv2MetricsRequestColumnNamesEnum string

// Enum of ColumnNames on RoCEv2MetricsRequest
var RoCEv2MetricsRequestColumnNames = struct {
	SESSION_STATE         RoCEv2MetricsRequestColumnNamesEnum
	QP_CONFUGRED          RoCEv2MetricsRequestColumnNamesEnum
	QP_UP                 RoCEv2MetricsRequestColumnNamesEnum
	QP_DOWN               RoCEv2MetricsRequestColumnNamesEnum
	CONNECT_REQUEST_TX    RoCEv2MetricsRequestColumnNamesEnum
	CONNECT_REQUEST_RX    RoCEv2MetricsRequestColumnNamesEnum
	CONNECT_REPLY_TX      RoCEv2MetricsRequestColumnNamesEnum
	CONNECT_REPLY_RX      RoCEv2MetricsRequestColumnNamesEnum
	READY_TX              RoCEv2MetricsRequestColumnNamesEnum
	READY_RX              RoCEv2MetricsRequestColumnNamesEnum
	DISCONNECT_REQUEST_TX RoCEv2MetricsRequestColumnNamesEnum
	DISCONNECT_REQUEST_RX RoCEv2MetricsRequestColumnNamesEnum
	DISCONNECT_REPLY_TX   RoCEv2MetricsRequestColumnNamesEnum
	DISCONNECT_REPLY_RX   RoCEv2MetricsRequestColumnNamesEnum
	REJECT_TX             RoCEv2MetricsRequestColumnNamesEnum
	REJECT_RX             RoCEv2MetricsRequestColumnNamesEnum
	UNKNOWN_MSG_RX        RoCEv2MetricsRequestColumnNamesEnum
}{
	SESSION_STATE:         RoCEv2MetricsRequestColumnNamesEnum("session_state"),
	QP_CONFUGRED:          RoCEv2MetricsRequestColumnNamesEnum("qp_confugred"),
	QP_UP:                 RoCEv2MetricsRequestColumnNamesEnum("qp_up"),
	QP_DOWN:               RoCEv2MetricsRequestColumnNamesEnum("qp_down"),
	CONNECT_REQUEST_TX:    RoCEv2MetricsRequestColumnNamesEnum("connect_request_tx"),
	CONNECT_REQUEST_RX:    RoCEv2MetricsRequestColumnNamesEnum("connect_request_rx"),
	CONNECT_REPLY_TX:      RoCEv2MetricsRequestColumnNamesEnum("connect_reply_tx"),
	CONNECT_REPLY_RX:      RoCEv2MetricsRequestColumnNamesEnum("connect_reply_rx"),
	READY_TX:              RoCEv2MetricsRequestColumnNamesEnum("ready_tx"),
	READY_RX:              RoCEv2MetricsRequestColumnNamesEnum("ready_rx"),
	DISCONNECT_REQUEST_TX: RoCEv2MetricsRequestColumnNamesEnum("disconnect_request_tx"),
	DISCONNECT_REQUEST_RX: RoCEv2MetricsRequestColumnNamesEnum("disconnect_request_rx"),
	DISCONNECT_REPLY_TX:   RoCEv2MetricsRequestColumnNamesEnum("disconnect_reply_tx"),
	DISCONNECT_REPLY_RX:   RoCEv2MetricsRequestColumnNamesEnum("disconnect_reply_rx"),
	REJECT_TX:             RoCEv2MetricsRequestColumnNamesEnum("reject_tx"),
	REJECT_RX:             RoCEv2MetricsRequestColumnNamesEnum("reject_rx"),
	UNKNOWN_MSG_RX:        RoCEv2MetricsRequestColumnNamesEnum("unknown_msg_rx"),
}

func (obj *roCEv2MetricsRequest) ColumnNames() []RoCEv2MetricsRequestColumnNamesEnum {
	items := []RoCEv2MetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, RoCEv2MetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the RoCEv2 peer cannot be excluded.
// SetColumnNames sets the []string value in the RoCEv2MetricsRequest object
func (obj *roCEv2MetricsRequest) SetColumnNames(value []RoCEv2MetricsRequestColumnNamesEnum) RoCEv2MetricsRequest {

	items := []otg.RoCEv2MetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.RoCEv2MetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.RoCEv2MetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *roCEv2MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *roCEv2MetricsRequest) setDefault() {

}
