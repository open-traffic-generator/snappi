package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Bgpv4MetricsRequest *****
type bgpv4MetricsRequest struct {
	validation
	obj          *otg.Bgpv4MetricsRequest
	marshaller   marshalBgpv4MetricsRequest
	unMarshaller unMarshalBgpv4MetricsRequest
}

func NewBgpv4MetricsRequest() Bgpv4MetricsRequest {
	obj := bgpv4MetricsRequest{obj: &otg.Bgpv4MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpv4MetricsRequest) msg() *otg.Bgpv4MetricsRequest {
	return obj.obj
}

func (obj *bgpv4MetricsRequest) setMsg(msg *otg.Bgpv4MetricsRequest) Bgpv4MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpv4MetricsRequest struct {
	obj *bgpv4MetricsRequest
}

type marshalBgpv4MetricsRequest interface {
	// ToProto marshals Bgpv4MetricsRequest to protobuf object *otg.Bgpv4MetricsRequest
	ToProto() (*otg.Bgpv4MetricsRequest, error)
	// ToPbText marshals Bgpv4MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Bgpv4MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Bgpv4MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalbgpv4MetricsRequest struct {
	obj *bgpv4MetricsRequest
}

type unMarshalBgpv4MetricsRequest interface {
	// FromProto unmarshals Bgpv4MetricsRequest from protobuf object *otg.Bgpv4MetricsRequest
	FromProto(msg *otg.Bgpv4MetricsRequest) (Bgpv4MetricsRequest, error)
	// FromPbText unmarshals Bgpv4MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Bgpv4MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Bgpv4MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *bgpv4MetricsRequest) Marshal() marshalBgpv4MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpv4MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpv4MetricsRequest) Unmarshal() unMarshalBgpv4MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpv4MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpv4MetricsRequest) ToProto() (*otg.Bgpv4MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpv4MetricsRequest) FromProto(msg *otg.Bgpv4MetricsRequest) (Bgpv4MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpv4MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalbgpv4MetricsRequest) FromPbText(value string) error {
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

func (m *marshalbgpv4MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalbgpv4MetricsRequest) FromYaml(value string) error {
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

func (m *marshalbgpv4MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalbgpv4MetricsRequest) FromJson(value string) error {
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

func (obj *bgpv4MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpv4MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpv4MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpv4MetricsRequest) Clone() (Bgpv4MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpv4MetricsRequest()
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

// Bgpv4MetricsRequest is the request to retrieve BGPv4 per peer metrics/statistics.
type Bgpv4MetricsRequest interface {
	Validation
	// msg marshals Bgpv4MetricsRequest to protobuf object *otg.Bgpv4MetricsRequest
	// and doesn't set defaults
	msg() *otg.Bgpv4MetricsRequest
	// setMsg unmarshals Bgpv4MetricsRequest from protobuf object *otg.Bgpv4MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Bgpv4MetricsRequest) Bgpv4MetricsRequest
	// provides marshal interface
	Marshal() marshalBgpv4MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalBgpv4MetricsRequest
	// validate validates Bgpv4MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Bgpv4MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in Bgpv4MetricsRequest.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to Bgpv4MetricsRequest
	SetPeerNames(value []string) Bgpv4MetricsRequest
	// ColumnNames returns []Bgpv4MetricsRequestColumnNamesEnum, set in Bgpv4MetricsRequest
	ColumnNames() []Bgpv4MetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Bgpv4MetricsRequestColumnNamesEnum provided by user to Bgpv4MetricsRequest
	SetColumnNames(value []Bgpv4MetricsRequestColumnNamesEnum) Bgpv4MetricsRequest
}

// The names of BGPv4 peers to return results for. An empty list will return results for all BGPv4 peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4peer/properties/name
//
// PeerNames returns a []string
func (obj *bgpv4MetricsRequest) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of BGPv4 peers to return results for. An empty list will return results for all BGPv4 peers.
//
// x-constraint:
// - /components/schemas/Bgp.V4peer/properties/name
//
// SetPeerNames sets the []string value in the Bgpv4MetricsRequest object
func (obj *bgpv4MetricsRequest) SetPeerNames(value []string) Bgpv4MetricsRequest {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

type Bgpv4MetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Bgpv4MetricsRequest
var Bgpv4MetricsRequestColumnNames = struct {
	SESSION_STATE            Bgpv4MetricsRequestColumnNamesEnum
	SESSION_FLAP_COUNT       Bgpv4MetricsRequestColumnNamesEnum
	ROUTES_ADVERTISED        Bgpv4MetricsRequestColumnNamesEnum
	ROUTES_RECEIVED          Bgpv4MetricsRequestColumnNamesEnum
	ROUTE_WITHDRAWS_SENT     Bgpv4MetricsRequestColumnNamesEnum
	ROUTE_WITHDRAWS_RECEIVED Bgpv4MetricsRequestColumnNamesEnum
	UPDATES_SENT             Bgpv4MetricsRequestColumnNamesEnum
	UPDATES_RECEIVED         Bgpv4MetricsRequestColumnNamesEnum
	OPENS_SENT               Bgpv4MetricsRequestColumnNamesEnum
	OPENS_RECEIVED           Bgpv4MetricsRequestColumnNamesEnum
	KEEPALIVES_SENT          Bgpv4MetricsRequestColumnNamesEnum
	KEEPALIVES_RECEIVED      Bgpv4MetricsRequestColumnNamesEnum
	NOTIFICATIONS_SENT       Bgpv4MetricsRequestColumnNamesEnum
	NOTIFICATIONS_RECEIVED   Bgpv4MetricsRequestColumnNamesEnum
	FSM_STATE                Bgpv4MetricsRequestColumnNamesEnum
	END_OF_RIB_RECEIVED      Bgpv4MetricsRequestColumnNamesEnum
}{
	SESSION_STATE:            Bgpv4MetricsRequestColumnNamesEnum("session_state"),
	SESSION_FLAP_COUNT:       Bgpv4MetricsRequestColumnNamesEnum("session_flap_count"),
	ROUTES_ADVERTISED:        Bgpv4MetricsRequestColumnNamesEnum("routes_advertised"),
	ROUTES_RECEIVED:          Bgpv4MetricsRequestColumnNamesEnum("routes_received"),
	ROUTE_WITHDRAWS_SENT:     Bgpv4MetricsRequestColumnNamesEnum("route_withdraws_sent"),
	ROUTE_WITHDRAWS_RECEIVED: Bgpv4MetricsRequestColumnNamesEnum("route_withdraws_received"),
	UPDATES_SENT:             Bgpv4MetricsRequestColumnNamesEnum("updates_sent"),
	UPDATES_RECEIVED:         Bgpv4MetricsRequestColumnNamesEnum("updates_received"),
	OPENS_SENT:               Bgpv4MetricsRequestColumnNamesEnum("opens_sent"),
	OPENS_RECEIVED:           Bgpv4MetricsRequestColumnNamesEnum("opens_received"),
	KEEPALIVES_SENT:          Bgpv4MetricsRequestColumnNamesEnum("keepalives_sent"),
	KEEPALIVES_RECEIVED:      Bgpv4MetricsRequestColumnNamesEnum("keepalives_received"),
	NOTIFICATIONS_SENT:       Bgpv4MetricsRequestColumnNamesEnum("notifications_sent"),
	NOTIFICATIONS_RECEIVED:   Bgpv4MetricsRequestColumnNamesEnum("notifications_received"),
	FSM_STATE:                Bgpv4MetricsRequestColumnNamesEnum("fsm_state"),
	END_OF_RIB_RECEIVED:      Bgpv4MetricsRequestColumnNamesEnum("end_of_rib_received"),
}

func (obj *bgpv4MetricsRequest) ColumnNames() []Bgpv4MetricsRequestColumnNamesEnum {
	items := []Bgpv4MetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Bgpv4MetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the BGPv4 peer cannot be excluded.
// SetColumnNames sets the []string value in the Bgpv4MetricsRequest object
func (obj *bgpv4MetricsRequest) SetColumnNames(value []Bgpv4MetricsRequestColumnNamesEnum) Bgpv4MetricsRequest {

	items := []otg.Bgpv4MetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Bgpv4MetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Bgpv4MetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *bgpv4MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpv4MetricsRequest) setDefault() {

}
