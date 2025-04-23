package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisMetricsRequest *****
type isisMetricsRequest struct {
	validation
	obj          *otg.IsisMetricsRequest
	marshaller   marshalIsisMetricsRequest
	unMarshaller unMarshalIsisMetricsRequest
}

func NewIsisMetricsRequest() IsisMetricsRequest {
	obj := isisMetricsRequest{obj: &otg.IsisMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *isisMetricsRequest) msg() *otg.IsisMetricsRequest {
	return obj.obj
}

func (obj *isisMetricsRequest) setMsg(msg *otg.IsisMetricsRequest) IsisMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisMetricsRequest struct {
	obj *isisMetricsRequest
}

type marshalIsisMetricsRequest interface {
	// ToProto marshals IsisMetricsRequest to protobuf object *otg.IsisMetricsRequest
	ToProto() (*otg.IsisMetricsRequest, error)
	// ToPbText marshals IsisMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalisisMetricsRequest struct {
	obj *isisMetricsRequest
}

type unMarshalIsisMetricsRequest interface {
	// FromProto unmarshals IsisMetricsRequest from protobuf object *otg.IsisMetricsRequest
	FromProto(msg *otg.IsisMetricsRequest) (IsisMetricsRequest, error)
	// FromPbText unmarshals IsisMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *isisMetricsRequest) Marshal() marshalIsisMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisMetricsRequest) Unmarshal() unMarshalIsisMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisMetricsRequest) ToProto() (*otg.IsisMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisMetricsRequest) FromProto(msg *otg.IsisMetricsRequest) (IsisMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalisisMetricsRequest) FromPbText(value string) error {
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

func (m *marshalisisMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalisisMetricsRequest) FromYaml(value string) error {
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

func (m *marshalisisMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalisisMetricsRequest) FromJson(value string) error {
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

func (obj *isisMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisMetricsRequest) Clone() (IsisMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisMetricsRequest()
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

// IsisMetricsRequest is the request to retrieve ISIS per Router metrics/statistics.
type IsisMetricsRequest interface {
	Validation
	// msg marshals IsisMetricsRequest to protobuf object *otg.IsisMetricsRequest
	// and doesn't set defaults
	msg() *otg.IsisMetricsRequest
	// setMsg unmarshals IsisMetricsRequest from protobuf object *otg.IsisMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.IsisMetricsRequest) IsisMetricsRequest
	// provides marshal interface
	Marshal() marshalIsisMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalIsisMetricsRequest
	// validate validates IsisMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in IsisMetricsRequest.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to IsisMetricsRequest
	SetRouterNames(value []string) IsisMetricsRequest
	// ColumnNames returns []IsisMetricsRequestColumnNamesEnum, set in IsisMetricsRequest
	ColumnNames() []IsisMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []IsisMetricsRequestColumnNamesEnum provided by user to IsisMetricsRequest
	SetColumnNames(value []IsisMetricsRequestColumnNamesEnum) IsisMetricsRequest
}

// The names of ISIS Routers to return results for. An empty list will return results for all ISIS router.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *isisMetricsRequest) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of ISIS Routers to return results for. An empty list will return results for all ISIS router.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the IsisMetricsRequest object
func (obj *isisMetricsRequest) SetRouterNames(value []string) IsisMetricsRequest {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type IsisMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on IsisMetricsRequest
var IsisMetricsRequestColumnNames = struct {
	L1_SESSIONS_UP                    IsisMetricsRequestColumnNamesEnum
	L1_SESSION_FLAP                   IsisMetricsRequestColumnNamesEnum
	L1_DATABASE_SIZE                  IsisMetricsRequestColumnNamesEnum
	L1_BROADCAST_HELLOS_SENT          IsisMetricsRequestColumnNamesEnum
	L1_BROADCAST_HELLOS_RECEIVED      IsisMetricsRequestColumnNamesEnum
	L1_POINT_TO_POINT_HELLOS_SENT     IsisMetricsRequestColumnNamesEnum
	L1_POINT_TO_POINT_HELLOS_RECEIVED IsisMetricsRequestColumnNamesEnum
	L1_PSNP_SENT                      IsisMetricsRequestColumnNamesEnum
	L1_PSNP_RECEIVED                  IsisMetricsRequestColumnNamesEnum
	L1_CSNP_SENT                      IsisMetricsRequestColumnNamesEnum
	L1_CSNP_RECEIVED                  IsisMetricsRequestColumnNamesEnum
	L1_LSP_SENT                       IsisMetricsRequestColumnNamesEnum
	L1_LSP_RECEIVED                   IsisMetricsRequestColumnNamesEnum
	L2_SESSIONS_UP                    IsisMetricsRequestColumnNamesEnum
	L2_SESSION_FLAP                   IsisMetricsRequestColumnNamesEnum
	L2_DATABASE_SIZE                  IsisMetricsRequestColumnNamesEnum
	L2_BROADCAST_HELLOS_SENT          IsisMetricsRequestColumnNamesEnum
	L2_BROADCAST_HELLOS_RECEIVED      IsisMetricsRequestColumnNamesEnum
	L2_POINT_TO_POINT_HELLOS_SENT     IsisMetricsRequestColumnNamesEnum
	L2_POINT_TO_POINT_HELLOS_RECEIVED IsisMetricsRequestColumnNamesEnum
	L2_PSNP_SENT                      IsisMetricsRequestColumnNamesEnum
	L2_PSNP_RECEIVED                  IsisMetricsRequestColumnNamesEnum
	L2_CSNP_SENT                      IsisMetricsRequestColumnNamesEnum
	L2_CSNP_RECEIVED                  IsisMetricsRequestColumnNamesEnum
	L2_LSP_SENT                       IsisMetricsRequestColumnNamesEnum
	L2_LSP_RECEIVED                   IsisMetricsRequestColumnNamesEnum
	L1_RR_SENT                        IsisMetricsRequestColumnNamesEnum
	L1_RR_RECEIVED                    IsisMetricsRequestColumnNamesEnum
	L1_RA_SENT                        IsisMetricsRequestColumnNamesEnum
	L1_RA_RECEIVED                    IsisMetricsRequestColumnNamesEnum
	L1_SA_SENT                        IsisMetricsRequestColumnNamesEnum
	L1_SA_RECEIVED                    IsisMetricsRequestColumnNamesEnum
	L2_RR_SENT                        IsisMetricsRequestColumnNamesEnum
	L2_RR_RECEIVED                    IsisMetricsRequestColumnNamesEnum
	L2_RA_SENT                        IsisMetricsRequestColumnNamesEnum
	L2_RA_RECEIVED                    IsisMetricsRequestColumnNamesEnum
	L2_SA_SENT                        IsisMetricsRequestColumnNamesEnum
	L2_SA_RECEIVED                    IsisMetricsRequestColumnNamesEnum
}{
	L1_SESSIONS_UP:                    IsisMetricsRequestColumnNamesEnum("l1_sessions_up"),
	L1_SESSION_FLAP:                   IsisMetricsRequestColumnNamesEnum("l1_session_flap"),
	L1_DATABASE_SIZE:                  IsisMetricsRequestColumnNamesEnum("l1_database_size"),
	L1_BROADCAST_HELLOS_SENT:          IsisMetricsRequestColumnNamesEnum("l1_broadcast_hellos_sent"),
	L1_BROADCAST_HELLOS_RECEIVED:      IsisMetricsRequestColumnNamesEnum("l1_broadcast_hellos_received"),
	L1_POINT_TO_POINT_HELLOS_SENT:     IsisMetricsRequestColumnNamesEnum("l1_point_to_point_hellos_sent"),
	L1_POINT_TO_POINT_HELLOS_RECEIVED: IsisMetricsRequestColumnNamesEnum("l1_point_to_point_hellos_received"),
	L1_PSNP_SENT:                      IsisMetricsRequestColumnNamesEnum("l1_psnp_sent"),
	L1_PSNP_RECEIVED:                  IsisMetricsRequestColumnNamesEnum("l1_psnp_received"),
	L1_CSNP_SENT:                      IsisMetricsRequestColumnNamesEnum("l1_csnp_sent"),
	L1_CSNP_RECEIVED:                  IsisMetricsRequestColumnNamesEnum("l1_csnp_received"),
	L1_LSP_SENT:                       IsisMetricsRequestColumnNamesEnum("l1_lsp_sent"),
	L1_LSP_RECEIVED:                   IsisMetricsRequestColumnNamesEnum("l1_lsp_received"),
	L2_SESSIONS_UP:                    IsisMetricsRequestColumnNamesEnum("l2_sessions_up"),
	L2_SESSION_FLAP:                   IsisMetricsRequestColumnNamesEnum("l2_session_flap"),
	L2_DATABASE_SIZE:                  IsisMetricsRequestColumnNamesEnum("l2_database_size"),
	L2_BROADCAST_HELLOS_SENT:          IsisMetricsRequestColumnNamesEnum("l2_broadcast_hellos_sent"),
	L2_BROADCAST_HELLOS_RECEIVED:      IsisMetricsRequestColumnNamesEnum("l2_broadcast_hellos_received"),
	L2_POINT_TO_POINT_HELLOS_SENT:     IsisMetricsRequestColumnNamesEnum("l2_point_to_point_hellos_sent"),
	L2_POINT_TO_POINT_HELLOS_RECEIVED: IsisMetricsRequestColumnNamesEnum("l2_point_to_point_hellos_received"),
	L2_PSNP_SENT:                      IsisMetricsRequestColumnNamesEnum("l2_psnp_sent"),
	L2_PSNP_RECEIVED:                  IsisMetricsRequestColumnNamesEnum("l2_psnp_received"),
	L2_CSNP_SENT:                      IsisMetricsRequestColumnNamesEnum("l2_csnp_sent"),
	L2_CSNP_RECEIVED:                  IsisMetricsRequestColumnNamesEnum("l2_csnp_received"),
	L2_LSP_SENT:                       IsisMetricsRequestColumnNamesEnum("l2_lsp_sent"),
	L2_LSP_RECEIVED:                   IsisMetricsRequestColumnNamesEnum("l2_lsp_received"),
	L1_RR_SENT:                        IsisMetricsRequestColumnNamesEnum("l1_rr_sent"),
	L1_RR_RECEIVED:                    IsisMetricsRequestColumnNamesEnum("l1_rr_received"),
	L1_RA_SENT:                        IsisMetricsRequestColumnNamesEnum("l1_ra_sent"),
	L1_RA_RECEIVED:                    IsisMetricsRequestColumnNamesEnum("l1_ra_received"),
	L1_SA_SENT:                        IsisMetricsRequestColumnNamesEnum("l1_sa_sent"),
	L1_SA_RECEIVED:                    IsisMetricsRequestColumnNamesEnum("l1_sa_received"),
	L2_RR_SENT:                        IsisMetricsRequestColumnNamesEnum("l2_rr_sent"),
	L2_RR_RECEIVED:                    IsisMetricsRequestColumnNamesEnum("l2_rr_received"),
	L2_RA_SENT:                        IsisMetricsRequestColumnNamesEnum("l2_ra_sent"),
	L2_RA_RECEIVED:                    IsisMetricsRequestColumnNamesEnum("l2_ra_received"),
	L2_SA_SENT:                        IsisMetricsRequestColumnNamesEnum("l2_sa_sent"),
	L2_SA_RECEIVED:                    IsisMetricsRequestColumnNamesEnum("l2_sa_received"),
}

func (obj *isisMetricsRequest) ColumnNames() []IsisMetricsRequestColumnNamesEnum {
	items := []IsisMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, IsisMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the ISIS Router cannot be excluded.
// SetColumnNames sets the []string value in the IsisMetricsRequest object
func (obj *isisMetricsRequest) SetColumnNames(value []IsisMetricsRequestColumnNamesEnum) IsisMetricsRequest {

	items := []otg.IsisMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.IsisMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.IsisMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *isisMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisMetricsRequest) setDefault() {

}
