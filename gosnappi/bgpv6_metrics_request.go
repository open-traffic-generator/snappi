package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Bgpv6MetricsRequest *****
type bgpv6MetricsRequest struct {
	validation
	obj          *otg.Bgpv6MetricsRequest
	marshaller   marshalBgpv6MetricsRequest
	unMarshaller unMarshalBgpv6MetricsRequest
}

func NewBgpv6MetricsRequest() Bgpv6MetricsRequest {
	obj := bgpv6MetricsRequest{obj: &otg.Bgpv6MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpv6MetricsRequest) msg() *otg.Bgpv6MetricsRequest {
	return obj.obj
}

func (obj *bgpv6MetricsRequest) setMsg(msg *otg.Bgpv6MetricsRequest) Bgpv6MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpv6MetricsRequest struct {
	obj *bgpv6MetricsRequest
}

type marshalBgpv6MetricsRequest interface {
	// ToProto marshals Bgpv6MetricsRequest to protobuf object *otg.Bgpv6MetricsRequest
	ToProto() (*otg.Bgpv6MetricsRequest, error)
	// ToPbText marshals Bgpv6MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Bgpv6MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Bgpv6MetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Bgpv6MetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpv6MetricsRequest struct {
	obj *bgpv6MetricsRequest
}

type unMarshalBgpv6MetricsRequest interface {
	// FromProto unmarshals Bgpv6MetricsRequest from protobuf object *otg.Bgpv6MetricsRequest
	FromProto(msg *otg.Bgpv6MetricsRequest) (Bgpv6MetricsRequest, error)
	// FromPbText unmarshals Bgpv6MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Bgpv6MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Bgpv6MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *bgpv6MetricsRequest) Marshal() marshalBgpv6MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpv6MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpv6MetricsRequest) Unmarshal() unMarshalBgpv6MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpv6MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpv6MetricsRequest) ToProto() (*otg.Bgpv6MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpv6MetricsRequest) FromProto(msg *otg.Bgpv6MetricsRequest) (Bgpv6MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpv6MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalbgpv6MetricsRequest) FromPbText(value string) error {
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

func (m *marshalbgpv6MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalbgpv6MetricsRequest) FromYaml(value string) error {
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

func (m *marshalbgpv6MetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshalbgpv6MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalbgpv6MetricsRequest) FromJson(value string) error {
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

func (obj *bgpv6MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpv6MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpv6MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpv6MetricsRequest) Clone() (Bgpv6MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpv6MetricsRequest()
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

// Bgpv6MetricsRequest is the request to retrieve BGPv6 per peer metrics/statistics.
type Bgpv6MetricsRequest interface {
	Validation
	// msg marshals Bgpv6MetricsRequest to protobuf object *otg.Bgpv6MetricsRequest
	// and doesn't set defaults
	msg() *otg.Bgpv6MetricsRequest
	// setMsg unmarshals Bgpv6MetricsRequest from protobuf object *otg.Bgpv6MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Bgpv6MetricsRequest) Bgpv6MetricsRequest
	// provides marshal interface
	Marshal() marshalBgpv6MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalBgpv6MetricsRequest
	// validate validates Bgpv6MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Bgpv6MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in Bgpv6MetricsRequest.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to Bgpv6MetricsRequest
	SetPeerNames(value []string) Bgpv6MetricsRequest
	// ColumnNames returns []Bgpv6MetricsRequestColumnNamesEnum, set in Bgpv6MetricsRequest
	ColumnNames() []Bgpv6MetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Bgpv6MetricsRequestColumnNamesEnum provided by user to Bgpv6MetricsRequest
	SetColumnNames(value []Bgpv6MetricsRequestColumnNamesEnum) Bgpv6MetricsRequest
}

// The names of BGPv6 peers to return results for. An empty list will return results for all BGPv6 peers.
//
// x-constraint:
// - /components/schemas/Bgp.V6peer/properties/name
//
// x-constraint:
// - /components/schemas/Bgp.V6peer/properties/name
//
// PeerNames returns a []string
func (obj *bgpv6MetricsRequest) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of BGPv6 peers to return results for. An empty list will return results for all BGPv6 peers.
//
// x-constraint:
// - /components/schemas/Bgp.V6peer/properties/name
//
// x-constraint:
// - /components/schemas/Bgp.V6peer/properties/name
//
// SetPeerNames sets the []string value in the Bgpv6MetricsRequest object
func (obj *bgpv6MetricsRequest) SetPeerNames(value []string) Bgpv6MetricsRequest {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type Bgpv6MetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Bgpv6MetricsRequest
var Bgpv6MetricsRequestColumnNames = struct {
	SESSION_STATE            Bgpv6MetricsRequestColumnNamesEnum
	SESSION_FLAP_COUNT       Bgpv6MetricsRequestColumnNamesEnum
	ROUTES_ADVERTISED        Bgpv6MetricsRequestColumnNamesEnum
	ROUTES_RECEIVED          Bgpv6MetricsRequestColumnNamesEnum
	ROUTE_WITHDRAWS_SENT     Bgpv6MetricsRequestColumnNamesEnum
	ROUTE_WITHDRAWS_RECEIVED Bgpv6MetricsRequestColumnNamesEnum
	UPDATES_SENT             Bgpv6MetricsRequestColumnNamesEnum
	UPDATES_RECEIVED         Bgpv6MetricsRequestColumnNamesEnum
	OPENS_SENT               Bgpv6MetricsRequestColumnNamesEnum
	OPENS_RECEIVED           Bgpv6MetricsRequestColumnNamesEnum
	KEEPALIVES_SENT          Bgpv6MetricsRequestColumnNamesEnum
	KEEPALIVES_RECEIVED      Bgpv6MetricsRequestColumnNamesEnum
	NOTIFICATIONS_SENT       Bgpv6MetricsRequestColumnNamesEnum
	NOTIFICATIONS_RECEIVED   Bgpv6MetricsRequestColumnNamesEnum
	FSM_STATE                Bgpv6MetricsRequestColumnNamesEnum
	END_OF_RIB_RECEIVED      Bgpv6MetricsRequestColumnNamesEnum
}{
	SESSION_STATE:            Bgpv6MetricsRequestColumnNamesEnum("session_state"),
	SESSION_FLAP_COUNT:       Bgpv6MetricsRequestColumnNamesEnum("session_flap_count"),
	ROUTES_ADVERTISED:        Bgpv6MetricsRequestColumnNamesEnum("routes_advertised"),
	ROUTES_RECEIVED:          Bgpv6MetricsRequestColumnNamesEnum("routes_received"),
	ROUTE_WITHDRAWS_SENT:     Bgpv6MetricsRequestColumnNamesEnum("route_withdraws_sent"),
	ROUTE_WITHDRAWS_RECEIVED: Bgpv6MetricsRequestColumnNamesEnum("route_withdraws_received"),
	UPDATES_SENT:             Bgpv6MetricsRequestColumnNamesEnum("updates_sent"),
	UPDATES_RECEIVED:         Bgpv6MetricsRequestColumnNamesEnum("updates_received"),
	OPENS_SENT:               Bgpv6MetricsRequestColumnNamesEnum("opens_sent"),
	OPENS_RECEIVED:           Bgpv6MetricsRequestColumnNamesEnum("opens_received"),
	KEEPALIVES_SENT:          Bgpv6MetricsRequestColumnNamesEnum("keepalives_sent"),
	KEEPALIVES_RECEIVED:      Bgpv6MetricsRequestColumnNamesEnum("keepalives_received"),
	NOTIFICATIONS_SENT:       Bgpv6MetricsRequestColumnNamesEnum("notifications_sent"),
	NOTIFICATIONS_RECEIVED:   Bgpv6MetricsRequestColumnNamesEnum("notifications_received"),
	FSM_STATE:                Bgpv6MetricsRequestColumnNamesEnum("fsm_state"),
	END_OF_RIB_RECEIVED:      Bgpv6MetricsRequestColumnNamesEnum("end_of_rib_received"),
}

func (obj *bgpv6MetricsRequest) ColumnNames() []Bgpv6MetricsRequestColumnNamesEnum {
	items := []Bgpv6MetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Bgpv6MetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the BGPv6 peer cannot be excluded.
// SetColumnNames sets the []string value in the Bgpv6MetricsRequest object
func (obj *bgpv6MetricsRequest) SetColumnNames(value []Bgpv6MetricsRequestColumnNamesEnum) Bgpv6MetricsRequest {

	items := []otg.Bgpv6MetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Bgpv6MetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Bgpv6MetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *bgpv6MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpv6MetricsRequest) setDefault() {

}
