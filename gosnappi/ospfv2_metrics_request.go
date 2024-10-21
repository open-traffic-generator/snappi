package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2MetricsRequest *****
type ospfv2MetricsRequest struct {
	validation
	obj          *otg.Ospfv2MetricsRequest
	marshaller   marshalOspfv2MetricsRequest
	unMarshaller unMarshalOspfv2MetricsRequest
}

func NewOspfv2MetricsRequest() Ospfv2MetricsRequest {
	obj := ospfv2MetricsRequest{obj: &otg.Ospfv2MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2MetricsRequest) msg() *otg.Ospfv2MetricsRequest {
	return obj.obj
}

func (obj *ospfv2MetricsRequest) setMsg(msg *otg.Ospfv2MetricsRequest) Ospfv2MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2MetricsRequest struct {
	obj *ospfv2MetricsRequest
}

type marshalOspfv2MetricsRequest interface {
	// ToProto marshals Ospfv2MetricsRequest to protobuf object *otg.Ospfv2MetricsRequest
	ToProto() (*otg.Ospfv2MetricsRequest, error)
	// ToPbText marshals Ospfv2MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2MetricsRequest struct {
	obj *ospfv2MetricsRequest
}

type unMarshalOspfv2MetricsRequest interface {
	// FromProto unmarshals Ospfv2MetricsRequest from protobuf object *otg.Ospfv2MetricsRequest
	FromProto(msg *otg.Ospfv2MetricsRequest) (Ospfv2MetricsRequest, error)
	// FromPbText unmarshals Ospfv2MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *ospfv2MetricsRequest) Marshal() marshalOspfv2MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2MetricsRequest) Unmarshal() unMarshalOspfv2MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2MetricsRequest) ToProto() (*otg.Ospfv2MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2MetricsRequest) FromProto(msg *otg.Ospfv2MetricsRequest) (Ospfv2MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalospfv2MetricsRequest) FromPbText(value string) error {
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

func (m *marshalospfv2MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalospfv2MetricsRequest) FromYaml(value string) error {
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

func (m *marshalospfv2MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalospfv2MetricsRequest) FromJson(value string) error {
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

func (obj *ospfv2MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2MetricsRequest) Clone() (Ospfv2MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2MetricsRequest()
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

// Ospfv2MetricsRequest is the request to retrieve OSPFv2 per Router metrics/statistics.
type Ospfv2MetricsRequest interface {
	Validation
	// msg marshals Ospfv2MetricsRequest to protobuf object *otg.Ospfv2MetricsRequest
	// and doesn't set defaults
	msg() *otg.Ospfv2MetricsRequest
	// setMsg unmarshals Ospfv2MetricsRequest from protobuf object *otg.Ospfv2MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Ospfv2MetricsRequest) Ospfv2MetricsRequest
	// provides marshal interface
	Marshal() marshalOspfv2MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2MetricsRequest
	// validate validates Ospfv2MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in Ospfv2MetricsRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to Ospfv2MetricsRequest
	SetRouterNames(value []string) Ospfv2MetricsRequest
	// ColumnNames returns []Ospfv2MetricsRequestColumnNamesEnum, set in Ospfv2MetricsRequest
	ColumnNames() []Ospfv2MetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Ospfv2MetricsRequestColumnNamesEnum provided by user to Ospfv2MetricsRequest
	SetColumnNames(value []Ospfv2MetricsRequestColumnNamesEnum) Ospfv2MetricsRequest
}

// The names of OSPFv2 routers to return results for. An empty list will return results for all OSPFv2 router.
//
// x-constraint:
// - /components/schemas/Device.Ospfv2/properties/name
//
// RouterNames returns a []string
func (obj *ospfv2MetricsRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of OSPFv2 routers to return results for. An empty list will return results for all OSPFv2 router.
//
// x-constraint:
// - /components/schemas/Device.Ospfv2/properties/name
//
// SetRouterNames sets the []string value in the Ospfv2MetricsRequest object
func (obj *ospfv2MetricsRequest) SetRouterNames(value []string) Ospfv2MetricsRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type Ospfv2MetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Ospfv2MetricsRequest
var Ospfv2MetricsRequestColumnNames = struct {
	FULL_STATE_COUNT       Ospfv2MetricsRequestColumnNamesEnum
	DOWN_STATE_COUNT       Ospfv2MetricsRequestColumnNamesEnum
	SESSIONS_FLAP          Ospfv2MetricsRequestColumnNamesEnum
	HELLOS_SENT            Ospfv2MetricsRequestColumnNamesEnum
	HELLOS_RECEIVED        Ospfv2MetricsRequestColumnNamesEnum
	DBD_SENT               Ospfv2MetricsRequestColumnNamesEnum
	DBD_RECEIVED           Ospfv2MetricsRequestColumnNamesEnum
	LS_REQUEST_SENT        Ospfv2MetricsRequestColumnNamesEnum
	LS_REQUEST_RECEIVED    Ospfv2MetricsRequestColumnNamesEnum
	LS_UPDATE_SENT         Ospfv2MetricsRequestColumnNamesEnum
	LS_UPDATE_RECEIVED     Ospfv2MetricsRequestColumnNamesEnum
	LS_ACK_SENT            Ospfv2MetricsRequestColumnNamesEnum
	LS_ACK_RECEIVED        Ospfv2MetricsRequestColumnNamesEnum
	LSA_SENT               Ospfv2MetricsRequestColumnNamesEnum
	LSA_RECEIVED           Ospfv2MetricsRequestColumnNamesEnum
	LSA_ACK_SENT           Ospfv2MetricsRequestColumnNamesEnum
	LSA_ACK_RECEIVED       Ospfv2MetricsRequestColumnNamesEnum
	ROUTER_LSA_SENT        Ospfv2MetricsRequestColumnNamesEnum
	ROUTER_LSA_RECEIVED    Ospfv2MetricsRequestColumnNamesEnum
	NETWORK_LSA_SENT       Ospfv2MetricsRequestColumnNamesEnum
	NETWORK_LSA_RECEIVED   Ospfv2MetricsRequestColumnNamesEnum
	SUMMARY_LSA_SENT       Ospfv2MetricsRequestColumnNamesEnum
	SUMMARY_LSA_RECEIVED   Ospfv2MetricsRequestColumnNamesEnum
	EXTERNAL_LSA_SENT      Ospfv2MetricsRequestColumnNamesEnum
	EXTERNAL_LSA_RECEIVED  Ospfv2MetricsRequestColumnNamesEnum
	NSSA_LSA_SENT          Ospfv2MetricsRequestColumnNamesEnum
	NSSA_LSA_RECEIVED      Ospfv2MetricsRequestColumnNamesEnum
	OPAQUE_LOCAL_SENT      Ospfv2MetricsRequestColumnNamesEnum
	OPAQUE_LOCAL_RECEIVED  Ospfv2MetricsRequestColumnNamesEnum
	OPAQUE_AREA_SENT       Ospfv2MetricsRequestColumnNamesEnum
	OPAQUE_AREA_RECEIVED   Ospfv2MetricsRequestColumnNamesEnum
	OPAQUE_DOMAIN_SENT     Ospfv2MetricsRequestColumnNamesEnum
	OPAQUE_DOMAIN_RECEIVED Ospfv2MetricsRequestColumnNamesEnum
}{
	FULL_STATE_COUNT:       Ospfv2MetricsRequestColumnNamesEnum("full_state_count"),
	DOWN_STATE_COUNT:       Ospfv2MetricsRequestColumnNamesEnum("down_state_count"),
	SESSIONS_FLAP:          Ospfv2MetricsRequestColumnNamesEnum("sessions_flap"),
	HELLOS_SENT:            Ospfv2MetricsRequestColumnNamesEnum("hellos_sent"),
	HELLOS_RECEIVED:        Ospfv2MetricsRequestColumnNamesEnum("hellos_received"),
	DBD_SENT:               Ospfv2MetricsRequestColumnNamesEnum("dbd_sent"),
	DBD_RECEIVED:           Ospfv2MetricsRequestColumnNamesEnum("dbd_received"),
	LS_REQUEST_SENT:        Ospfv2MetricsRequestColumnNamesEnum("ls_request_sent"),
	LS_REQUEST_RECEIVED:    Ospfv2MetricsRequestColumnNamesEnum("ls_request_received"),
	LS_UPDATE_SENT:         Ospfv2MetricsRequestColumnNamesEnum("ls_update_sent"),
	LS_UPDATE_RECEIVED:     Ospfv2MetricsRequestColumnNamesEnum("ls_update_received"),
	LS_ACK_SENT:            Ospfv2MetricsRequestColumnNamesEnum("ls_ack_sent"),
	LS_ACK_RECEIVED:        Ospfv2MetricsRequestColumnNamesEnum("ls_ack_received"),
	LSA_SENT:               Ospfv2MetricsRequestColumnNamesEnum("lsa_sent"),
	LSA_RECEIVED:           Ospfv2MetricsRequestColumnNamesEnum("lsa_received"),
	LSA_ACK_SENT:           Ospfv2MetricsRequestColumnNamesEnum("lsa_ack_sent"),
	LSA_ACK_RECEIVED:       Ospfv2MetricsRequestColumnNamesEnum("lsa_ack_received"),
	ROUTER_LSA_SENT:        Ospfv2MetricsRequestColumnNamesEnum("router_lsa_sent"),
	ROUTER_LSA_RECEIVED:    Ospfv2MetricsRequestColumnNamesEnum("router_lsa_received"),
	NETWORK_LSA_SENT:       Ospfv2MetricsRequestColumnNamesEnum("network_lsa_sent"),
	NETWORK_LSA_RECEIVED:   Ospfv2MetricsRequestColumnNamesEnum("network_lsa_received"),
	SUMMARY_LSA_SENT:       Ospfv2MetricsRequestColumnNamesEnum("summary_lsa_sent"),
	SUMMARY_LSA_RECEIVED:   Ospfv2MetricsRequestColumnNamesEnum("summary_lsa_received"),
	EXTERNAL_LSA_SENT:      Ospfv2MetricsRequestColumnNamesEnum("external_lsa_sent"),
	EXTERNAL_LSA_RECEIVED:  Ospfv2MetricsRequestColumnNamesEnum("external_lsa_received"),
	NSSA_LSA_SENT:          Ospfv2MetricsRequestColumnNamesEnum("nssa_lsa_sent"),
	NSSA_LSA_RECEIVED:      Ospfv2MetricsRequestColumnNamesEnum("nssa_lsa_received"),
	OPAQUE_LOCAL_SENT:      Ospfv2MetricsRequestColumnNamesEnum("opaque_local_sent"),
	OPAQUE_LOCAL_RECEIVED:  Ospfv2MetricsRequestColumnNamesEnum("opaque_local_received"),
	OPAQUE_AREA_SENT:       Ospfv2MetricsRequestColumnNamesEnum("opaque_area_sent"),
	OPAQUE_AREA_RECEIVED:   Ospfv2MetricsRequestColumnNamesEnum("opaque_area_received"),
	OPAQUE_DOMAIN_SENT:     Ospfv2MetricsRequestColumnNamesEnum("opaque_domain_sent"),
	OPAQUE_DOMAIN_RECEIVED: Ospfv2MetricsRequestColumnNamesEnum("opaque_domain_received"),
}

func (obj *ospfv2MetricsRequest) ColumnNames() []Ospfv2MetricsRequestColumnNamesEnum {
	items := []Ospfv2MetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Ospfv2MetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain.
// If the list is empty then all columns will be returned except for
// any result_groups.
// The name of the OSPFv2 Router cannot be excluded.
// SetColumnNames sets the []string value in the Ospfv2MetricsRequest object
func (obj *ospfv2MetricsRequest) SetColumnNames(value []Ospfv2MetricsRequestColumnNamesEnum) Ospfv2MetricsRequest {

	items := []otg.Ospfv2MetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Ospfv2MetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Ospfv2MetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *ospfv2MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv2MetricsRequest) setDefault() {

}
