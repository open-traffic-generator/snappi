package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3MetricsRequest *****
type ospfv3MetricsRequest struct {
	validation
	obj          *otg.Ospfv3MetricsRequest
	marshaller   marshalOspfv3MetricsRequest
	unMarshaller unMarshalOspfv3MetricsRequest
}

func NewOspfv3MetricsRequest() Ospfv3MetricsRequest {
	obj := ospfv3MetricsRequest{obj: &otg.Ospfv3MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3MetricsRequest) msg() *otg.Ospfv3MetricsRequest {
	return obj.obj
}

func (obj *ospfv3MetricsRequest) setMsg(msg *otg.Ospfv3MetricsRequest) Ospfv3MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3MetricsRequest struct {
	obj *ospfv3MetricsRequest
}

type marshalOspfv3MetricsRequest interface {
	// ToProto marshals Ospfv3MetricsRequest to protobuf object *otg.Ospfv3MetricsRequest
	ToProto() (*otg.Ospfv3MetricsRequest, error)
	// ToPbText marshals Ospfv3MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3MetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv3MetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv3MetricsRequest struct {
	obj *ospfv3MetricsRequest
}

type unMarshalOspfv3MetricsRequest interface {
	// FromProto unmarshals Ospfv3MetricsRequest from protobuf object *otg.Ospfv3MetricsRequest
	FromProto(msg *otg.Ospfv3MetricsRequest) (Ospfv3MetricsRequest, error)
	// FromPbText unmarshals Ospfv3MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *ospfv3MetricsRequest) Marshal() marshalOspfv3MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3MetricsRequest) Unmarshal() unMarshalOspfv3MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3MetricsRequest) ToProto() (*otg.Ospfv3MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3MetricsRequest) FromProto(msg *otg.Ospfv3MetricsRequest) (Ospfv3MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalospfv3MetricsRequest) FromPbText(value string) error {
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

func (m *marshalospfv3MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalospfv3MetricsRequest) FromYaml(value string) error {
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

func (m *marshalospfv3MetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshalospfv3MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalospfv3MetricsRequest) FromJson(value string) error {
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

func (obj *ospfv3MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3MetricsRequest) Clone() (Ospfv3MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3MetricsRequest()
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

// Ospfv3MetricsRequest is the request to retrieve OSPFv3 per router metrics/statistics.
type Ospfv3MetricsRequest interface {
	Validation
	// msg marshals Ospfv3MetricsRequest to protobuf object *otg.Ospfv3MetricsRequest
	// and doesn't set defaults
	msg() *otg.Ospfv3MetricsRequest
	// setMsg unmarshals Ospfv3MetricsRequest from protobuf object *otg.Ospfv3MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Ospfv3MetricsRequest) Ospfv3MetricsRequest
	// provides marshal interface
	Marshal() marshalOspfv3MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3MetricsRequest
	// validate validates Ospfv3MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in Ospfv3MetricsRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to Ospfv3MetricsRequest
	SetRouterNames(value []string) Ospfv3MetricsRequest
	// ColumnNames returns []Ospfv3MetricsRequestColumnNamesEnum, set in Ospfv3MetricsRequest
	ColumnNames() []Ospfv3MetricsRequestColumnNamesEnum
	// SetColumnNames assigns []Ospfv3MetricsRequestColumnNamesEnum provided by user to Ospfv3MetricsRequest
	SetColumnNames(value []Ospfv3MetricsRequestColumnNamesEnum) Ospfv3MetricsRequest
}

// The names of OSPFv3 routers to return results for. An empty list will return results for all OSPFv3 routers.
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// RouterNames returns a []string
func (obj *ospfv3MetricsRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of OSPFv3 routers to return results for. An empty list will return results for all OSPFv3 routers.
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// x-constraint:
// - /components/schemas/Ospfv3.RouterInstance/properties/name
//
// SetRouterNames sets the []string value in the Ospfv3MetricsRequest object
func (obj *ospfv3MetricsRequest) SetRouterNames(value []string) Ospfv3MetricsRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type Ospfv3MetricsRequestColumnNamesEnum string

// Enum of ColumnNames on Ospfv3MetricsRequest
var Ospfv3MetricsRequestColumnNames = struct {
	FULL_STATE_COUNT               Ospfv3MetricsRequestColumnNamesEnum
	DOWN_STATE_COUNT               Ospfv3MetricsRequestColumnNamesEnum
	SESSIONS_FLAP                  Ospfv3MetricsRequestColumnNamesEnum
	HELLOS_SENT                    Ospfv3MetricsRequestColumnNamesEnum
	HELLOS_RECEIVED                Ospfv3MetricsRequestColumnNamesEnum
	DBD_SENT                       Ospfv3MetricsRequestColumnNamesEnum
	DBD_RECEIVED                   Ospfv3MetricsRequestColumnNamesEnum
	LS_REQUEST_SENT                Ospfv3MetricsRequestColumnNamesEnum
	LS_REQUEST_RECEIVED            Ospfv3MetricsRequestColumnNamesEnum
	LS_UPDATE_SENT                 Ospfv3MetricsRequestColumnNamesEnum
	LS_UPDATE_RECEIVED             Ospfv3MetricsRequestColumnNamesEnum
	LS_ACK_SENT                    Ospfv3MetricsRequestColumnNamesEnum
	LS_ACK_RECEIVED                Ospfv3MetricsRequestColumnNamesEnum
	LSA_SENT                       Ospfv3MetricsRequestColumnNamesEnum
	LSA_RECEIVED                   Ospfv3MetricsRequestColumnNamesEnum
	ROUTER_LSA_SENT                Ospfv3MetricsRequestColumnNamesEnum
	ROUTER_LSA_RECEIVED            Ospfv3MetricsRequestColumnNamesEnum
	NETWORK_LSA_SENT               Ospfv3MetricsRequestColumnNamesEnum
	NETWORK_LSA_RECEIVED           Ospfv3MetricsRequestColumnNamesEnum
	INTER_AREA_PREFIX_LSA_SENT     Ospfv3MetricsRequestColumnNamesEnum
	INTER_AREA_PREFIX_LSA_RECEIVED Ospfv3MetricsRequestColumnNamesEnum
	INTER_AREA_ROUTER_LSA_SENT     Ospfv3MetricsRequestColumnNamesEnum
	INTER_AREA_ROUTER_LSA_RECEIVED Ospfv3MetricsRequestColumnNamesEnum
	EXTERNAL_LSA_SENT              Ospfv3MetricsRequestColumnNamesEnum
	EXTERNAL_LSA_RECEIVED          Ospfv3MetricsRequestColumnNamesEnum
	NSSA_LSA_SENT                  Ospfv3MetricsRequestColumnNamesEnum
	NSSA_LSA_RECEIVED              Ospfv3MetricsRequestColumnNamesEnum
	LINK_LSA_SENT                  Ospfv3MetricsRequestColumnNamesEnum
	LINK_LSA_RECEIVED              Ospfv3MetricsRequestColumnNamesEnum
	INTRA_AREA_PREFIX_LSA_SENT     Ospfv3MetricsRequestColumnNamesEnum
	INTRA_AREA_PREFIX_LSA_RECEIVED Ospfv3MetricsRequestColumnNamesEnum
}{
	FULL_STATE_COUNT:               Ospfv3MetricsRequestColumnNamesEnum("full_state_count"),
	DOWN_STATE_COUNT:               Ospfv3MetricsRequestColumnNamesEnum("down_state_count"),
	SESSIONS_FLAP:                  Ospfv3MetricsRequestColumnNamesEnum("sessions_flap"),
	HELLOS_SENT:                    Ospfv3MetricsRequestColumnNamesEnum("hellos_sent"),
	HELLOS_RECEIVED:                Ospfv3MetricsRequestColumnNamesEnum("hellos_received"),
	DBD_SENT:                       Ospfv3MetricsRequestColumnNamesEnum("dbd_sent"),
	DBD_RECEIVED:                   Ospfv3MetricsRequestColumnNamesEnum("dbd_received"),
	LS_REQUEST_SENT:                Ospfv3MetricsRequestColumnNamesEnum("ls_request_sent"),
	LS_REQUEST_RECEIVED:            Ospfv3MetricsRequestColumnNamesEnum("ls_request_received"),
	LS_UPDATE_SENT:                 Ospfv3MetricsRequestColumnNamesEnum("ls_update_sent"),
	LS_UPDATE_RECEIVED:             Ospfv3MetricsRequestColumnNamesEnum("ls_update_received"),
	LS_ACK_SENT:                    Ospfv3MetricsRequestColumnNamesEnum("ls_ack_sent"),
	LS_ACK_RECEIVED:                Ospfv3MetricsRequestColumnNamesEnum("ls_ack_received"),
	LSA_SENT:                       Ospfv3MetricsRequestColumnNamesEnum("lsa_sent"),
	LSA_RECEIVED:                   Ospfv3MetricsRequestColumnNamesEnum("lsa_received"),
	ROUTER_LSA_SENT:                Ospfv3MetricsRequestColumnNamesEnum("router_lsa_sent"),
	ROUTER_LSA_RECEIVED:            Ospfv3MetricsRequestColumnNamesEnum("router_lsa_received"),
	NETWORK_LSA_SENT:               Ospfv3MetricsRequestColumnNamesEnum("network_lsa_sent"),
	NETWORK_LSA_RECEIVED:           Ospfv3MetricsRequestColumnNamesEnum("network_lsa_received"),
	INTER_AREA_PREFIX_LSA_SENT:     Ospfv3MetricsRequestColumnNamesEnum("inter_area_prefix_lsa_sent"),
	INTER_AREA_PREFIX_LSA_RECEIVED: Ospfv3MetricsRequestColumnNamesEnum("inter_area_prefix_lsa_received"),
	INTER_AREA_ROUTER_LSA_SENT:     Ospfv3MetricsRequestColumnNamesEnum("inter_area_router_lsa_sent"),
	INTER_AREA_ROUTER_LSA_RECEIVED: Ospfv3MetricsRequestColumnNamesEnum("inter_area_router_lsa_received"),
	EXTERNAL_LSA_SENT:              Ospfv3MetricsRequestColumnNamesEnum("external_lsa_sent"),
	EXTERNAL_LSA_RECEIVED:          Ospfv3MetricsRequestColumnNamesEnum("external_lsa_received"),
	NSSA_LSA_SENT:                  Ospfv3MetricsRequestColumnNamesEnum("nssa_lsa_sent"),
	NSSA_LSA_RECEIVED:              Ospfv3MetricsRequestColumnNamesEnum("nssa_lsa_received"),
	LINK_LSA_SENT:                  Ospfv3MetricsRequestColumnNamesEnum("link_lsa_sent"),
	LINK_LSA_RECEIVED:              Ospfv3MetricsRequestColumnNamesEnum("link_lsa_received"),
	INTRA_AREA_PREFIX_LSA_SENT:     Ospfv3MetricsRequestColumnNamesEnum("intra_area_prefix_lsa_sent"),
	INTRA_AREA_PREFIX_LSA_RECEIVED: Ospfv3MetricsRequestColumnNamesEnum("intra_area_prefix_lsa_received"),
}

func (obj *ospfv3MetricsRequest) ColumnNames() []Ospfv3MetricsRequestColumnNamesEnum {
	items := []Ospfv3MetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, Ospfv3MetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain.
// If the list is empty then all columns will be returned except for
// any result_groups.
// The name of the OSPFv3 Router cannot be excluded.
// SetColumnNames sets the []string value in the Ospfv3MetricsRequest object
func (obj *ospfv3MetricsRequest) SetColumnNames(value []Ospfv3MetricsRequestColumnNamesEnum) Ospfv3MetricsRequest {

	items := []otg.Ospfv3MetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.Ospfv3MetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.Ospfv3MetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *ospfv3MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ospfv3MetricsRequest) setDefault() {

}
