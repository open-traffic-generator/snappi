package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpMetricsRequest *****
type ldpMetricsRequest struct {
	validation
	obj          *otg.LdpMetricsRequest
	marshaller   marshalLdpMetricsRequest
	unMarshaller unMarshalLdpMetricsRequest
}

func NewLdpMetricsRequest() LdpMetricsRequest {
	obj := ldpMetricsRequest{obj: &otg.LdpMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpMetricsRequest) msg() *otg.LdpMetricsRequest {
	return obj.obj
}

func (obj *ldpMetricsRequest) setMsg(msg *otg.LdpMetricsRequest) LdpMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpMetricsRequest struct {
	obj *ldpMetricsRequest
}

type marshalLdpMetricsRequest interface {
	// ToProto marshals LdpMetricsRequest to protobuf object *otg.LdpMetricsRequest
	ToProto() (*otg.LdpMetricsRequest, error)
	// ToPbText marshals LdpMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalldpMetricsRequest struct {
	obj *ldpMetricsRequest
}

type unMarshalLdpMetricsRequest interface {
	// FromProto unmarshals LdpMetricsRequest from protobuf object *otg.LdpMetricsRequest
	FromProto(msg *otg.LdpMetricsRequest) (LdpMetricsRequest, error)
	// FromPbText unmarshals LdpMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *ldpMetricsRequest) Marshal() marshalLdpMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpMetricsRequest) Unmarshal() unMarshalLdpMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpMetricsRequest) ToProto() (*otg.LdpMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpMetricsRequest) FromProto(msg *otg.LdpMetricsRequest) (LdpMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalldpMetricsRequest) FromPbText(value string) error {
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

func (m *marshalldpMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalldpMetricsRequest) FromYaml(value string) error {
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

func (m *marshalldpMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalldpMetricsRequest) FromJson(value string) error {
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

func (obj *ldpMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpMetricsRequest) Clone() (LdpMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpMetricsRequest()
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

// LdpMetricsRequest is the request to retrieve LDP per router metrics/statistics.
type LdpMetricsRequest interface {
	Validation
	// msg marshals LdpMetricsRequest to protobuf object *otg.LdpMetricsRequest
	// and doesn't set defaults
	msg() *otg.LdpMetricsRequest
	// setMsg unmarshals LdpMetricsRequest from protobuf object *otg.LdpMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.LdpMetricsRequest) LdpMetricsRequest
	// provides marshal interface
	Marshal() marshalLdpMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalLdpMetricsRequest
	// validate validates LdpMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in LdpMetricsRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to LdpMetricsRequest
	SetRouterNames(value []string) LdpMetricsRequest
	// ColumnNames returns []LdpMetricsRequestColumnNamesEnum, set in LdpMetricsRequest
	ColumnNames() []LdpMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []LdpMetricsRequestColumnNamesEnum provided by user to LdpMetricsRequest
	SetColumnNames(value []LdpMetricsRequestColumnNamesEnum) LdpMetricsRequest
}

// The names of LDP routers to return results for. An empty list will return results for all LDP routers.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// RouterNames returns a []string
func (obj *ldpMetricsRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of LDP routers to return results for. An empty list will return results for all LDP routers.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// SetRouterNames sets the []string value in the LdpMetricsRequest object
func (obj *ldpMetricsRequest) SetRouterNames(value []string) LdpMetricsRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type LdpMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on LdpMetricsRequest
var LdpMetricsRequestColumnNames = struct {
	SESSIONS_UP                   LdpMetricsRequestColumnNamesEnum
	SESSIONS_FLAP                 LdpMetricsRequestColumnNamesEnum
	SESSIONS_NON_EXISTENT         LdpMetricsRequestColumnNamesEnum
	SESSIONS_INITIALIZED          LdpMetricsRequestColumnNamesEnum
	SESSIONS_OPEN_RECEIVED        LdpMetricsRequestColumnNamesEnum
	SESSIONS_OPEN_SENT            LdpMetricsRequestColumnNamesEnum
	SESSIONS_OPERATIONAL          LdpMetricsRequestColumnNamesEnum
	NOTIFICATIONS_SENT            LdpMetricsRequestColumnNamesEnum
	NOTIFICATIONS_RECEIVED        LdpMetricsRequestColumnNamesEnum
	LABEL_MAPPINGS_SENT           LdpMetricsRequestColumnNamesEnum
	LABEL_MAPPINGS_RECEIVED       LdpMetricsRequestColumnNamesEnum
	LABEL_REQUESTS_SENT           LdpMetricsRequestColumnNamesEnum
	LABEL_REQUESTS_RECEIVED       LdpMetricsRequestColumnNamesEnum
	LABEL_WITHDRAWS_SENT          LdpMetricsRequestColumnNamesEnum
	LABEL_WITHDRAWS_RECEIVED      LdpMetricsRequestColumnNamesEnum
	LABEL_RELEASES_SENT           LdpMetricsRequestColumnNamesEnum
	LABEL_RELEASES_RECEIVED       LdpMetricsRequestColumnNamesEnum
	LABEL_ABORT_REQUESTS_SENT     LdpMetricsRequestColumnNamesEnum
	LABEL_ABORT_REQUESTS_RECEIVED LdpMetricsRequestColumnNamesEnum
	INGRESS_LSPS_UP               LdpMetricsRequestColumnNamesEnum
	EGRESS_LSPS_UP                LdpMetricsRequestColumnNamesEnum
}{
	SESSIONS_UP:                   LdpMetricsRequestColumnNamesEnum("sessions_up"),
	SESSIONS_FLAP:                 LdpMetricsRequestColumnNamesEnum("sessions_flap"),
	SESSIONS_NON_EXISTENT:         LdpMetricsRequestColumnNamesEnum("sessions_non_existent"),
	SESSIONS_INITIALIZED:          LdpMetricsRequestColumnNamesEnum("sessions_initialized"),
	SESSIONS_OPEN_RECEIVED:        LdpMetricsRequestColumnNamesEnum("sessions_open_received"),
	SESSIONS_OPEN_SENT:            LdpMetricsRequestColumnNamesEnum("sessions_open_sent"),
	SESSIONS_OPERATIONAL:          LdpMetricsRequestColumnNamesEnum("sessions_operational"),
	NOTIFICATIONS_SENT:            LdpMetricsRequestColumnNamesEnum("notifications_sent"),
	NOTIFICATIONS_RECEIVED:        LdpMetricsRequestColumnNamesEnum("notifications_received"),
	LABEL_MAPPINGS_SENT:           LdpMetricsRequestColumnNamesEnum("label_mappings_sent"),
	LABEL_MAPPINGS_RECEIVED:       LdpMetricsRequestColumnNamesEnum("label_mappings_received"),
	LABEL_REQUESTS_SENT:           LdpMetricsRequestColumnNamesEnum("label_requests_sent"),
	LABEL_REQUESTS_RECEIVED:       LdpMetricsRequestColumnNamesEnum("label_requests_received"),
	LABEL_WITHDRAWS_SENT:          LdpMetricsRequestColumnNamesEnum("label_withdraws_sent"),
	LABEL_WITHDRAWS_RECEIVED:      LdpMetricsRequestColumnNamesEnum("label_withdraws_received"),
	LABEL_RELEASES_SENT:           LdpMetricsRequestColumnNamesEnum("label_releases_sent"),
	LABEL_RELEASES_RECEIVED:       LdpMetricsRequestColumnNamesEnum("label_releases_received"),
	LABEL_ABORT_REQUESTS_SENT:     LdpMetricsRequestColumnNamesEnum("label_abort_requests_sent"),
	LABEL_ABORT_REQUESTS_RECEIVED: LdpMetricsRequestColumnNamesEnum("label_abort_requests_received"),
	INGRESS_LSPS_UP:               LdpMetricsRequestColumnNamesEnum("ingress_lsps_up"),
	EGRESS_LSPS_UP:                LdpMetricsRequestColumnNamesEnum("egress_lsps_up"),
}

func (obj *ldpMetricsRequest) ColumnNames() []LdpMetricsRequestColumnNamesEnum {
	items := []LdpMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, LdpMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned. The name of the LDP router cannot be excluded.
// SetColumnNames sets the []string value in the LdpMetricsRequest object
func (obj *ldpMetricsRequest) SetColumnNames(value []LdpMetricsRequestColumnNamesEnum) LdpMetricsRequest {

	items := []otg.LdpMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.LdpMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.LdpMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *ldpMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ldpMetricsRequest) setDefault() {

}
