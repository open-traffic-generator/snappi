package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpMetricsRequest *****
type rsvpMetricsRequest struct {
	validation
	obj          *otg.RsvpMetricsRequest
	marshaller   marshalRsvpMetricsRequest
	unMarshaller unMarshalRsvpMetricsRequest
}

func NewRsvpMetricsRequest() RsvpMetricsRequest {
	obj := rsvpMetricsRequest{obj: &otg.RsvpMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpMetricsRequest) msg() *otg.RsvpMetricsRequest {
	return obj.obj
}

func (obj *rsvpMetricsRequest) setMsg(msg *otg.RsvpMetricsRequest) RsvpMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpMetricsRequest struct {
	obj *rsvpMetricsRequest
}

type marshalRsvpMetricsRequest interface {
	// ToProto marshals RsvpMetricsRequest to protobuf object *otg.RsvpMetricsRequest
	ToProto() (*otg.RsvpMetricsRequest, error)
	// ToPbText marshals RsvpMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpMetricsRequest struct {
	obj *rsvpMetricsRequest
}

type unMarshalRsvpMetricsRequest interface {
	// FromProto unmarshals RsvpMetricsRequest from protobuf object *otg.RsvpMetricsRequest
	FromProto(msg *otg.RsvpMetricsRequest) (RsvpMetricsRequest, error)
	// FromPbText unmarshals RsvpMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *rsvpMetricsRequest) Marshal() marshalRsvpMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpMetricsRequest) Unmarshal() unMarshalRsvpMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpMetricsRequest) ToProto() (*otg.RsvpMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpMetricsRequest) FromProto(msg *otg.RsvpMetricsRequest) (RsvpMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalrsvpMetricsRequest) FromPbText(value string) error {
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

func (m *marshalrsvpMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalrsvpMetricsRequest) FromYaml(value string) error {
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

func (m *marshalrsvpMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalrsvpMetricsRequest) FromJson(value string) error {
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

func (obj *rsvpMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpMetricsRequest) Clone() (RsvpMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpMetricsRequest()
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

// RsvpMetricsRequest is the request to retrieve RSVP-TE per Router metrics/statistics.
type RsvpMetricsRequest interface {
	Validation
	// msg marshals RsvpMetricsRequest to protobuf object *otg.RsvpMetricsRequest
	// and doesn't set defaults
	msg() *otg.RsvpMetricsRequest
	// setMsg unmarshals RsvpMetricsRequest from protobuf object *otg.RsvpMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.RsvpMetricsRequest) RsvpMetricsRequest
	// provides marshal interface
	Marshal() marshalRsvpMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpMetricsRequest
	// validate validates RsvpMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in RsvpMetricsRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to RsvpMetricsRequest
	SetRouterNames(value []string) RsvpMetricsRequest
	// ColumnNames returns []RsvpMetricsRequestColumnNamesEnum, set in RsvpMetricsRequest
	ColumnNames() []RsvpMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []RsvpMetricsRequestColumnNamesEnum provided by user to RsvpMetricsRequest
	SetColumnNames(value []RsvpMetricsRequestColumnNamesEnum) RsvpMetricsRequest
}

// The names of RSVP-TE Routers to return results for. An empty list as input will return results for all RSVP-TE routers.
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// RouterNames returns a []string
func (obj *rsvpMetricsRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of RSVP-TE Routers to return results for. An empty list as input will return results for all RSVP-TE routers.
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// x-constraint:
// - /components/schemas/Device.Rsvp/properties/name
//
// SetRouterNames sets the []string value in the RsvpMetricsRequest object
func (obj *rsvpMetricsRequest) SetRouterNames(value []string) RsvpMetricsRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type RsvpMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on RsvpMetricsRequest
var RsvpMetricsRequestColumnNames = struct {
	INGRESS_P2P_LSPS_CONFIGURED  RsvpMetricsRequestColumnNamesEnum
	INGRESS_P2P_LSPS_UP          RsvpMetricsRequestColumnNamesEnum
	EGRESS_P2P_LSPS_UP           RsvpMetricsRequestColumnNamesEnum
	LSP_FLAP_COUNT               RsvpMetricsRequestColumnNamesEnum
	PATHS_TX                     RsvpMetricsRequestColumnNamesEnum
	PATHS_RX                     RsvpMetricsRequestColumnNamesEnum
	RESVS_TX                     RsvpMetricsRequestColumnNamesEnum
	RESVS_RX                     RsvpMetricsRequestColumnNamesEnum
	PATH_TEARS_TX                RsvpMetricsRequestColumnNamesEnum
	PATH_TEARS_RX                RsvpMetricsRequestColumnNamesEnum
	RESV_TEARS_TX                RsvpMetricsRequestColumnNamesEnum
	RESV_TEARS_RX                RsvpMetricsRequestColumnNamesEnum
	PATH_ERRORS_TX               RsvpMetricsRequestColumnNamesEnum
	PATH_ERRORS_RX               RsvpMetricsRequestColumnNamesEnum
	RESV_ERRORS_TX               RsvpMetricsRequestColumnNamesEnum
	RESV_ERRORS_RX               RsvpMetricsRequestColumnNamesEnum
	RESV_CONF_TX                 RsvpMetricsRequestColumnNamesEnum
	RESV_CONF_RX                 RsvpMetricsRequestColumnNamesEnum
	HELLOS_TX                    RsvpMetricsRequestColumnNamesEnum
	HELLOS_RX                    RsvpMetricsRequestColumnNamesEnum
	ACKS_TX                      RsvpMetricsRequestColumnNamesEnum
	ACKS_RX                      RsvpMetricsRequestColumnNamesEnum
	NACKS_TX                     RsvpMetricsRequestColumnNamesEnum
	NACKS_RX                     RsvpMetricsRequestColumnNamesEnum
	SREFRESH_TX                  RsvpMetricsRequestColumnNamesEnum
	SREFRESH_RX                  RsvpMetricsRequestColumnNamesEnum
	BUNDLE_TX                    RsvpMetricsRequestColumnNamesEnum
	BUNDLE_RX                    RsvpMetricsRequestColumnNamesEnum
	PATH_REEVALUATION_REQUEST_TX RsvpMetricsRequestColumnNamesEnum
	PATH_REOPTIMIZATIONS         RsvpMetricsRequestColumnNamesEnum
}{
	INGRESS_P2P_LSPS_CONFIGURED:  RsvpMetricsRequestColumnNamesEnum("ingress_p2p_lsps_configured"),
	INGRESS_P2P_LSPS_UP:          RsvpMetricsRequestColumnNamesEnum("ingress_p2p_lsps_up"),
	EGRESS_P2P_LSPS_UP:           RsvpMetricsRequestColumnNamesEnum("egress_p2p_lsps_up"),
	LSP_FLAP_COUNT:               RsvpMetricsRequestColumnNamesEnum("lsp_flap_count"),
	PATHS_TX:                     RsvpMetricsRequestColumnNamesEnum("paths_tx"),
	PATHS_RX:                     RsvpMetricsRequestColumnNamesEnum("paths_rx"),
	RESVS_TX:                     RsvpMetricsRequestColumnNamesEnum("resvs_tx"),
	RESVS_RX:                     RsvpMetricsRequestColumnNamesEnum("resvs_rx"),
	PATH_TEARS_TX:                RsvpMetricsRequestColumnNamesEnum("path_tears_tx"),
	PATH_TEARS_RX:                RsvpMetricsRequestColumnNamesEnum("path_tears_rx"),
	RESV_TEARS_TX:                RsvpMetricsRequestColumnNamesEnum("resv_tears_tx"),
	RESV_TEARS_RX:                RsvpMetricsRequestColumnNamesEnum("resv_tears_rx"),
	PATH_ERRORS_TX:               RsvpMetricsRequestColumnNamesEnum("path_errors_tx"),
	PATH_ERRORS_RX:               RsvpMetricsRequestColumnNamesEnum("path_errors_rx"),
	RESV_ERRORS_TX:               RsvpMetricsRequestColumnNamesEnum("resv_errors_tx"),
	RESV_ERRORS_RX:               RsvpMetricsRequestColumnNamesEnum("resv_errors_rx"),
	RESV_CONF_TX:                 RsvpMetricsRequestColumnNamesEnum("resv_conf_tx"),
	RESV_CONF_RX:                 RsvpMetricsRequestColumnNamesEnum("resv_conf_rx"),
	HELLOS_TX:                    RsvpMetricsRequestColumnNamesEnum("hellos_tx"),
	HELLOS_RX:                    RsvpMetricsRequestColumnNamesEnum("hellos_rx"),
	ACKS_TX:                      RsvpMetricsRequestColumnNamesEnum("acks_tx"),
	ACKS_RX:                      RsvpMetricsRequestColumnNamesEnum("acks_rx"),
	NACKS_TX:                     RsvpMetricsRequestColumnNamesEnum("nacks_tx"),
	NACKS_RX:                     RsvpMetricsRequestColumnNamesEnum("nacks_rx"),
	SREFRESH_TX:                  RsvpMetricsRequestColumnNamesEnum("srefresh_tx"),
	SREFRESH_RX:                  RsvpMetricsRequestColumnNamesEnum("srefresh_rx"),
	BUNDLE_TX:                    RsvpMetricsRequestColumnNamesEnum("bundle_tx"),
	BUNDLE_RX:                    RsvpMetricsRequestColumnNamesEnum("bundle_rx"),
	PATH_REEVALUATION_REQUEST_TX: RsvpMetricsRequestColumnNamesEnum("path_reevaluation_request_tx"),
	PATH_REOPTIMIZATIONS:         RsvpMetricsRequestColumnNamesEnum("path_reoptimizations"),
}

func (obj *rsvpMetricsRequest) ColumnNames() []RsvpMetricsRequestColumnNamesEnum {
	items := []RsvpMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, RsvpMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the input list is empty then all columns will be returned except for any result_groups.
// SetColumnNames sets the []string value in the RsvpMetricsRequest object
func (obj *rsvpMetricsRequest) SetColumnNames(value []RsvpMetricsRequestColumnNamesEnum) RsvpMetricsRequest {

	items := []otg.RsvpMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.RsvpMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.RsvpMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *rsvpMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rsvpMetricsRequest) setDefault() {

}
