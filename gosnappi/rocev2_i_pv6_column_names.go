package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2IPv6ColumnNames *****
type rocev2IPv6ColumnNames struct {
	validation
	obj          *otg.Rocev2IPv6ColumnNames
	marshaller   marshalRocev2IPv6ColumnNames
	unMarshaller unMarshalRocev2IPv6ColumnNames
}

func NewRocev2IPv6ColumnNames() Rocev2IPv6ColumnNames {
	obj := rocev2IPv6ColumnNames{obj: &otg.Rocev2IPv6ColumnNames{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2IPv6ColumnNames) msg() *otg.Rocev2IPv6ColumnNames {
	return obj.obj
}

func (obj *rocev2IPv6ColumnNames) setMsg(msg *otg.Rocev2IPv6ColumnNames) Rocev2IPv6ColumnNames {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2IPv6ColumnNames struct {
	obj *rocev2IPv6ColumnNames
}

type marshalRocev2IPv6ColumnNames interface {
	// ToProto marshals Rocev2IPv6ColumnNames to protobuf object *otg.Rocev2IPv6ColumnNames
	ToProto() (*otg.Rocev2IPv6ColumnNames, error)
	// ToPbText marshals Rocev2IPv6ColumnNames to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2IPv6ColumnNames to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2IPv6ColumnNames to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2IPv6ColumnNames struct {
	obj *rocev2IPv6ColumnNames
}

type unMarshalRocev2IPv6ColumnNames interface {
	// FromProto unmarshals Rocev2IPv6ColumnNames from protobuf object *otg.Rocev2IPv6ColumnNames
	FromProto(msg *otg.Rocev2IPv6ColumnNames) (Rocev2IPv6ColumnNames, error)
	// FromPbText unmarshals Rocev2IPv6ColumnNames from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2IPv6ColumnNames from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2IPv6ColumnNames from JSON text
	FromJson(value string) error
}

func (obj *rocev2IPv6ColumnNames) Marshal() marshalRocev2IPv6ColumnNames {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2IPv6ColumnNames{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2IPv6ColumnNames) Unmarshal() unMarshalRocev2IPv6ColumnNames {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2IPv6ColumnNames{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2IPv6ColumnNames) ToProto() (*otg.Rocev2IPv6ColumnNames, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2IPv6ColumnNames) FromProto(msg *otg.Rocev2IPv6ColumnNames) (Rocev2IPv6ColumnNames, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2IPv6ColumnNames) ToPbText() (string, error) {
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

func (m *unMarshalrocev2IPv6ColumnNames) FromPbText(value string) error {
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

func (m *marshalrocev2IPv6ColumnNames) ToYaml() (string, error) {
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

func (m *unMarshalrocev2IPv6ColumnNames) FromYaml(value string) error {
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

func (m *marshalrocev2IPv6ColumnNames) ToJson() (string, error) {
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

func (m *unMarshalrocev2IPv6ColumnNames) FromJson(value string) error {
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

func (obj *rocev2IPv6ColumnNames) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2IPv6ColumnNames) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2IPv6ColumnNames) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2IPv6ColumnNames) Clone() (Rocev2IPv6ColumnNames, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2IPv6ColumnNames()
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

// Rocev2IPv6ColumnNames is the names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
type Rocev2IPv6ColumnNames interface {
	Validation
	// msg marshals Rocev2IPv6ColumnNames to protobuf object *otg.Rocev2IPv6ColumnNames
	// and doesn't set defaults
	msg() *otg.Rocev2IPv6ColumnNames
	// setMsg unmarshals Rocev2IPv6ColumnNames from protobuf object *otg.Rocev2IPv6ColumnNames
	// and doesn't set defaults
	setMsg(*otg.Rocev2IPv6ColumnNames) Rocev2IPv6ColumnNames
	// provides marshal interface
	Marshal() marshalRocev2IPv6ColumnNames
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2IPv6ColumnNames
	// validate validates Rocev2IPv6ColumnNames
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2IPv6ColumnNames, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in Rocev2IPv6ColumnNames.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to Rocev2IPv6ColumnNames
	SetPeerNames(value []string) Rocev2IPv6ColumnNames
	// ColumnNames returns []Rocev2IPv6ColumnNamesColumnNamesEnum, set in Rocev2IPv6ColumnNames
	ColumnNames() []Rocev2IPv6ColumnNamesColumnNamesEnum
	// SetColumnNames assigns []Rocev2IPv6ColumnNamesColumnNamesEnum provided by user to Rocev2IPv6ColumnNames
	SetColumnNames(value []Rocev2IPv6ColumnNamesColumnNamesEnum) Rocev2IPv6ColumnNames
}

// The names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Rocev2.V6peer/properties/name
//
// PeerNames returns a []string
func (obj *rocev2IPv6ColumnNames) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
//
// x-constraint:
// - /components/schemas/Rocev2.V6peer/properties/name
//
// SetPeerNames sets the []string value in the Rocev2IPv6ColumnNames object
func (obj *rocev2IPv6ColumnNames) SetPeerNames(value []string) Rocev2IPv6ColumnNames {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type Rocev2IPv6ColumnNamesColumnNamesEnum string

// Enum of ColumnNames on Rocev2IPv6ColumnNames
var Rocev2IPv6ColumnNamesColumnNames = struct {
	QP_CONFIGURED         Rocev2IPv6ColumnNamesColumnNamesEnum
	QP_UP                 Rocev2IPv6ColumnNamesColumnNamesEnum
	QP_DOWN               Rocev2IPv6ColumnNamesColumnNamesEnum
	CONNECT_REQUEST_TX    Rocev2IPv6ColumnNamesColumnNamesEnum
	CONNECT_REQUEST_RX    Rocev2IPv6ColumnNamesColumnNamesEnum
	CONNECT_REPLY_TX      Rocev2IPv6ColumnNamesColumnNamesEnum
	CONNECT_REPLY_RX      Rocev2IPv6ColumnNamesColumnNamesEnum
	READY_TX              Rocev2IPv6ColumnNamesColumnNamesEnum
	READY_RX              Rocev2IPv6ColumnNamesColumnNamesEnum
	DISCONNECT_REQUEST_TX Rocev2IPv6ColumnNamesColumnNamesEnum
	DISCONNECT_REQUEST_RX Rocev2IPv6ColumnNamesColumnNamesEnum
	DISCONNECT_REPLY_TX   Rocev2IPv6ColumnNamesColumnNamesEnum
	DISCONNECT_REPLY_RX   Rocev2IPv6ColumnNamesColumnNamesEnum
	REJECT_TX             Rocev2IPv6ColumnNamesColumnNamesEnum
	REJECT_RX             Rocev2IPv6ColumnNamesColumnNamesEnum
	UNKNOWN_MSG_RX        Rocev2IPv6ColumnNamesColumnNamesEnum
}{
	QP_CONFIGURED:         Rocev2IPv6ColumnNamesColumnNamesEnum("qp_configured"),
	QP_UP:                 Rocev2IPv6ColumnNamesColumnNamesEnum("qp_up"),
	QP_DOWN:               Rocev2IPv6ColumnNamesColumnNamesEnum("qp_down"),
	CONNECT_REQUEST_TX:    Rocev2IPv6ColumnNamesColumnNamesEnum("connect_request_tx"),
	CONNECT_REQUEST_RX:    Rocev2IPv6ColumnNamesColumnNamesEnum("connect_request_rx"),
	CONNECT_REPLY_TX:      Rocev2IPv6ColumnNamesColumnNamesEnum("connect_reply_tx"),
	CONNECT_REPLY_RX:      Rocev2IPv6ColumnNamesColumnNamesEnum("connect_reply_rx"),
	READY_TX:              Rocev2IPv6ColumnNamesColumnNamesEnum("ready_tx"),
	READY_RX:              Rocev2IPv6ColumnNamesColumnNamesEnum("ready_rx"),
	DISCONNECT_REQUEST_TX: Rocev2IPv6ColumnNamesColumnNamesEnum("disconnect_request_tx"),
	DISCONNECT_REQUEST_RX: Rocev2IPv6ColumnNamesColumnNamesEnum("disconnect_request_rx"),
	DISCONNECT_REPLY_TX:   Rocev2IPv6ColumnNamesColumnNamesEnum("disconnect_reply_tx"),
	DISCONNECT_REPLY_RX:   Rocev2IPv6ColumnNamesColumnNamesEnum("disconnect_reply_rx"),
	REJECT_TX:             Rocev2IPv6ColumnNamesColumnNamesEnum("reject_tx"),
	REJECT_RX:             Rocev2IPv6ColumnNamesColumnNamesEnum("reject_rx"),
	UNKNOWN_MSG_RX:        Rocev2IPv6ColumnNamesColumnNamesEnum("unknown_msg_rx"),
}

func (obj *rocev2IPv6ColumnNames) ColumnNames() []Rocev2IPv6ColumnNamesColumnNamesEnum {
	items := []Rocev2IPv6ColumnNamesColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Rocev2IPv6ColumnNamesColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the RoCEv2 peer cannot be excluded.
// SetColumnNames sets the []string value in the Rocev2IPv6ColumnNames object
func (obj *rocev2IPv6ColumnNames) SetColumnNames(value []Rocev2IPv6ColumnNamesColumnNamesEnum) Rocev2IPv6ColumnNames {

	items := []otg.Rocev2IPv6ColumnNames_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Rocev2IPv6ColumnNames_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Rocev2IPv6ColumnNames_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *rocev2IPv6ColumnNames) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rocev2IPv6ColumnNames) setDefault() {

}
