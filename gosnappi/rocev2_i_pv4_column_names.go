package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2IPv4ColumnNames *****
type rocev2IPv4ColumnNames struct {
	validation
	obj          *otg.Rocev2IPv4ColumnNames
	marshaller   marshalRocev2IPv4ColumnNames
	unMarshaller unMarshalRocev2IPv4ColumnNames
}

func NewRocev2IPv4ColumnNames() Rocev2IPv4ColumnNames {
	obj := rocev2IPv4ColumnNames{obj: &otg.Rocev2IPv4ColumnNames{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2IPv4ColumnNames) msg() *otg.Rocev2IPv4ColumnNames {
	return obj.obj
}

func (obj *rocev2IPv4ColumnNames) setMsg(msg *otg.Rocev2IPv4ColumnNames) Rocev2IPv4ColumnNames {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2IPv4ColumnNames struct {
	obj *rocev2IPv4ColumnNames
}

type marshalRocev2IPv4ColumnNames interface {
	// ToProto marshals Rocev2IPv4ColumnNames to protobuf object *otg.Rocev2IPv4ColumnNames
	ToProto() (*otg.Rocev2IPv4ColumnNames, error)
	// ToPbText marshals Rocev2IPv4ColumnNames to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2IPv4ColumnNames to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2IPv4ColumnNames to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2IPv4ColumnNames to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2IPv4ColumnNames struct {
	obj *rocev2IPv4ColumnNames
}

type unMarshalRocev2IPv4ColumnNames interface {
	// FromProto unmarshals Rocev2IPv4ColumnNames from protobuf object *otg.Rocev2IPv4ColumnNames
	FromProto(msg *otg.Rocev2IPv4ColumnNames) (Rocev2IPv4ColumnNames, error)
	// FromPbText unmarshals Rocev2IPv4ColumnNames from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2IPv4ColumnNames from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2IPv4ColumnNames from JSON text
	FromJson(value string) error
}

func (obj *rocev2IPv4ColumnNames) Marshal() marshalRocev2IPv4ColumnNames {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2IPv4ColumnNames{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2IPv4ColumnNames) Unmarshal() unMarshalRocev2IPv4ColumnNames {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2IPv4ColumnNames{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2IPv4ColumnNames) ToProto() (*otg.Rocev2IPv4ColumnNames, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2IPv4ColumnNames) FromProto(msg *otg.Rocev2IPv4ColumnNames) (Rocev2IPv4ColumnNames, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2IPv4ColumnNames) ToPbText() (string, error) {
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

func (m *unMarshalrocev2IPv4ColumnNames) FromPbText(value string) error {
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

func (m *marshalrocev2IPv4ColumnNames) ToYaml() (string, error) {
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

func (m *unMarshalrocev2IPv4ColumnNames) FromYaml(value string) error {
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

func (m *marshalrocev2IPv4ColumnNames) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2IPv4ColumnNames) ToJson() (string, error) {
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

func (m *unMarshalrocev2IPv4ColumnNames) FromJson(value string) error {
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

func (obj *rocev2IPv4ColumnNames) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2IPv4ColumnNames) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2IPv4ColumnNames) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2IPv4ColumnNames) Clone() (Rocev2IPv4ColumnNames, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2IPv4ColumnNames()
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

// Rocev2IPv4ColumnNames is the names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
type Rocev2IPv4ColumnNames interface {
	Validation
	// msg marshals Rocev2IPv4ColumnNames to protobuf object *otg.Rocev2IPv4ColumnNames
	// and doesn't set defaults
	msg() *otg.Rocev2IPv4ColumnNames
	// setMsg unmarshals Rocev2IPv4ColumnNames from protobuf object *otg.Rocev2IPv4ColumnNames
	// and doesn't set defaults
	setMsg(*otg.Rocev2IPv4ColumnNames) Rocev2IPv4ColumnNames
	// provides marshal interface
	Marshal() marshalRocev2IPv4ColumnNames
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2IPv4ColumnNames
	// validate validates Rocev2IPv4ColumnNames
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2IPv4ColumnNames, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in Rocev2IPv4ColumnNames.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to Rocev2IPv4ColumnNames
	SetPeerNames(value []string) Rocev2IPv4ColumnNames
	// ColumnNames returns []Rocev2IPv4ColumnNamesColumnNamesEnum, set in Rocev2IPv4ColumnNames
	ColumnNames() []Rocev2IPv4ColumnNamesColumnNamesEnum
	// SetColumnNames assigns []Rocev2IPv4ColumnNamesColumnNamesEnum provided by user to Rocev2IPv4ColumnNames
	SetColumnNames(value []Rocev2IPv4ColumnNamesColumnNamesEnum) Rocev2IPv4ColumnNames
}

// The names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Rocev2.V4peer/properties/name
//
// x-constraint:
// - /components/schemas/Rocev2.V4peer/properties/name
//
// PeerNames returns a []string
func (obj *rocev2IPv4ColumnNames) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Rocev2.V4peer/properties/name
//
// x-constraint:
// - /components/schemas/Rocev2.V4peer/properties/name
//
// SetPeerNames sets the []string value in the Rocev2IPv4ColumnNames object
func (obj *rocev2IPv4ColumnNames) SetPeerNames(value []string) Rocev2IPv4ColumnNames {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type Rocev2IPv4ColumnNamesColumnNamesEnum string

// Enum of ColumnNames on Rocev2IPv4ColumnNames
var Rocev2IPv4ColumnNamesColumnNames = struct {
	QP_CONFIGURED         Rocev2IPv4ColumnNamesColumnNamesEnum
	QP_UP                 Rocev2IPv4ColumnNamesColumnNamesEnum
	QP_DOWN               Rocev2IPv4ColumnNamesColumnNamesEnum
	CONNECT_REQUEST_TX    Rocev2IPv4ColumnNamesColumnNamesEnum
	CONNECT_REQUEST_RX    Rocev2IPv4ColumnNamesColumnNamesEnum
	CONNECT_REPLY_TX      Rocev2IPv4ColumnNamesColumnNamesEnum
	CONNECT_REPLY_RX      Rocev2IPv4ColumnNamesColumnNamesEnum
	READY_TX              Rocev2IPv4ColumnNamesColumnNamesEnum
	READY_RX              Rocev2IPv4ColumnNamesColumnNamesEnum
	DISCONNECT_REQUEST_TX Rocev2IPv4ColumnNamesColumnNamesEnum
	DISCONNECT_REQUEST_RX Rocev2IPv4ColumnNamesColumnNamesEnum
	DISCONNECT_REPLY_TX   Rocev2IPv4ColumnNamesColumnNamesEnum
	DISCONNECT_REPLY_RX   Rocev2IPv4ColumnNamesColumnNamesEnum
	REJECT_TX             Rocev2IPv4ColumnNamesColumnNamesEnum
	REJECT_RX             Rocev2IPv4ColumnNamesColumnNamesEnum
	UNKNOWN_MSG_RX        Rocev2IPv4ColumnNamesColumnNamesEnum
}{
	QP_CONFIGURED:         Rocev2IPv4ColumnNamesColumnNamesEnum("qp_configured"),
	QP_UP:                 Rocev2IPv4ColumnNamesColumnNamesEnum("qp_up"),
	QP_DOWN:               Rocev2IPv4ColumnNamesColumnNamesEnum("qp_down"),
	CONNECT_REQUEST_TX:    Rocev2IPv4ColumnNamesColumnNamesEnum("connect_request_tx"),
	CONNECT_REQUEST_RX:    Rocev2IPv4ColumnNamesColumnNamesEnum("connect_request_rx"),
	CONNECT_REPLY_TX:      Rocev2IPv4ColumnNamesColumnNamesEnum("connect_reply_tx"),
	CONNECT_REPLY_RX:      Rocev2IPv4ColumnNamesColumnNamesEnum("connect_reply_rx"),
	READY_TX:              Rocev2IPv4ColumnNamesColumnNamesEnum("ready_tx"),
	READY_RX:              Rocev2IPv4ColumnNamesColumnNamesEnum("ready_rx"),
	DISCONNECT_REQUEST_TX: Rocev2IPv4ColumnNamesColumnNamesEnum("disconnect_request_tx"),
	DISCONNECT_REQUEST_RX: Rocev2IPv4ColumnNamesColumnNamesEnum("disconnect_request_rx"),
	DISCONNECT_REPLY_TX:   Rocev2IPv4ColumnNamesColumnNamesEnum("disconnect_reply_tx"),
	DISCONNECT_REPLY_RX:   Rocev2IPv4ColumnNamesColumnNamesEnum("disconnect_reply_rx"),
	REJECT_TX:             Rocev2IPv4ColumnNamesColumnNamesEnum("reject_tx"),
	REJECT_RX:             Rocev2IPv4ColumnNamesColumnNamesEnum("reject_rx"),
	UNKNOWN_MSG_RX:        Rocev2IPv4ColumnNamesColumnNamesEnum("unknown_msg_rx"),
}

func (obj *rocev2IPv4ColumnNames) ColumnNames() []Rocev2IPv4ColumnNamesColumnNamesEnum {
	items := []Rocev2IPv4ColumnNamesColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Rocev2IPv4ColumnNamesColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the RoCEv2 peer cannot be excluded.
// SetColumnNames sets the []string value in the Rocev2IPv4ColumnNames object
func (obj *rocev2IPv4ColumnNames) SetColumnNames(value []Rocev2IPv4ColumnNamesColumnNamesEnum) Rocev2IPv4ColumnNames {

	items := []otg.Rocev2IPv4ColumnNames_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Rocev2IPv4ColumnNames_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Rocev2IPv4ColumnNames_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *rocev2IPv4ColumnNames) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rocev2IPv4ColumnNames) setDefault() {

}
