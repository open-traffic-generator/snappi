package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LacpMetricsRequest *****
type lacpMetricsRequest struct {
	validation
	obj          *otg.LacpMetricsRequest
	marshaller   marshalLacpMetricsRequest
	unMarshaller unMarshalLacpMetricsRequest
}

func NewLacpMetricsRequest() LacpMetricsRequest {
	obj := lacpMetricsRequest{obj: &otg.LacpMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *lacpMetricsRequest) msg() *otg.LacpMetricsRequest {
	return obj.obj
}

func (obj *lacpMetricsRequest) setMsg(msg *otg.LacpMetricsRequest) LacpMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallacpMetricsRequest struct {
	obj *lacpMetricsRequest
}

type marshalLacpMetricsRequest interface {
	// ToProto marshals LacpMetricsRequest to protobuf object *otg.LacpMetricsRequest
	ToProto() (*otg.LacpMetricsRequest, error)
	// ToPbText marshals LacpMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LacpMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals LacpMetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LacpMetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallacpMetricsRequest struct {
	obj *lacpMetricsRequest
}

type unMarshalLacpMetricsRequest interface {
	// FromProto unmarshals LacpMetricsRequest from protobuf object *otg.LacpMetricsRequest
	FromProto(msg *otg.LacpMetricsRequest) (LacpMetricsRequest, error)
	// FromPbText unmarshals LacpMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LacpMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LacpMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *lacpMetricsRequest) Marshal() marshalLacpMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshallacpMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *lacpMetricsRequest) Unmarshal() unMarshalLacpMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallacpMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallacpMetricsRequest) ToProto() (*otg.LacpMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallacpMetricsRequest) FromProto(msg *otg.LacpMetricsRequest) (LacpMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallacpMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshallacpMetricsRequest) FromPbText(value string) error {
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

func (m *marshallacpMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshallacpMetricsRequest) FromYaml(value string) error {
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

func (m *marshallacpMetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshallacpMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshallacpMetricsRequest) FromJson(value string) error {
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

func (obj *lacpMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lacpMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lacpMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lacpMetricsRequest) Clone() (LacpMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLacpMetricsRequest()
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

// LacpMetricsRequest is the request to retrieve LACP per LAG member metrics/statistics.
type LacpMetricsRequest interface {
	Validation
	// msg marshals LacpMetricsRequest to protobuf object *otg.LacpMetricsRequest
	// and doesn't set defaults
	msg() *otg.LacpMetricsRequest
	// setMsg unmarshals LacpMetricsRequest from protobuf object *otg.LacpMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.LacpMetricsRequest) LacpMetricsRequest
	// provides marshal interface
	Marshal() marshalLacpMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalLacpMetricsRequest
	// validate validates LacpMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LacpMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LagNames returns []string, set in LacpMetricsRequest.
	LagNames() []string
	// SetLagNames assigns []string provided by user to LacpMetricsRequest
	SetLagNames(value []string) LacpMetricsRequest
	// LagMemberPortNames returns []string, set in LacpMetricsRequest.
	LagMemberPortNames() []string
	// SetLagMemberPortNames assigns []string provided by user to LacpMetricsRequest
	SetLagMemberPortNames(value []string) LacpMetricsRequest
	// ColumnNames returns []LacpMetricsRequestColumnNamesEnum, set in LacpMetricsRequest
	ColumnNames() []LacpMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []LacpMetricsRequestColumnNamesEnum provided by user to LacpMetricsRequest
	SetColumnNames(value []LacpMetricsRequestColumnNamesEnum) LacpMetricsRequest
}

// The names of LAG (ports group) for which LACP metrics to be returned. An empty list will return metrics for all LAGs.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// LagNames returns a []string
func (obj *lacpMetricsRequest) LagNames() []string {
	if obj.obj.LagNames == nil {
		obj.obj.LagNames = make([]string, 0)
	}
	return obj.obj.LagNames
}

// The names of LAG (ports group) for which LACP metrics to be returned. An empty list will return metrics for all LAGs.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// SetLagNames sets the []string value in the LacpMetricsRequest object
func (obj *lacpMetricsRequest) SetLagNames(value []string) LacpMetricsRequest {

	if obj.obj.LagNames == nil {
		obj.obj.LagNames = make([]string, 0)
	}
	obj.obj.LagNames = value

	return obj
}

// The names of LAG members (ports) for which LACP metrics to be returned. An empty list will return metrics for all LAG members.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// LagMemberPortNames returns a []string
func (obj *lacpMetricsRequest) LagMemberPortNames() []string {
	if obj.obj.LagMemberPortNames == nil {
		obj.obj.LagMemberPortNames = make([]string, 0)
	}
	return obj.obj.LagMemberPortNames
}

// The names of LAG members (ports) for which LACP metrics to be returned. An empty list will return metrics for all LAG members.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetLagMemberPortNames sets the []string value in the LacpMetricsRequest object
func (obj *lacpMetricsRequest) SetLagMemberPortNames(value []string) LacpMetricsRequest {

	if obj.obj.LagMemberPortNames == nil {
		obj.obj.LagMemberPortNames = make([]string, 0)
	}
	obj.obj.LagMemberPortNames = value

	return obj
}

type LacpMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on LacpMetricsRequest
var LacpMetricsRequestColumnNames = struct {
	LACP_PACKETS_RX  LacpMetricsRequestColumnNamesEnum
	LACP_PACKETS_TX  LacpMetricsRequestColumnNamesEnum
	LACP_RX_ERRORS   LacpMetricsRequestColumnNamesEnum
	ACTIVITY         LacpMetricsRequestColumnNamesEnum
	TIMEOUT          LacpMetricsRequestColumnNamesEnum
	SYNCHRONIZATION  LacpMetricsRequestColumnNamesEnum
	AGGREGATABLE     LacpMetricsRequestColumnNamesEnum
	COLLECTING       LacpMetricsRequestColumnNamesEnum
	DISTRIBUTING     LacpMetricsRequestColumnNamesEnum
	SYSTEM_ID        LacpMetricsRequestColumnNamesEnum
	OPER_KEY         LacpMetricsRequestColumnNamesEnum
	PARTNER_ID       LacpMetricsRequestColumnNamesEnum
	PARTNER_KEY      LacpMetricsRequestColumnNamesEnum
	PORT_NUM         LacpMetricsRequestColumnNamesEnum
	PARTNER_PORT_NUM LacpMetricsRequestColumnNamesEnum
}{
	LACP_PACKETS_RX:  LacpMetricsRequestColumnNamesEnum("lacp_packets_rx"),
	LACP_PACKETS_TX:  LacpMetricsRequestColumnNamesEnum("lacp_packets_tx"),
	LACP_RX_ERRORS:   LacpMetricsRequestColumnNamesEnum("lacp_rx_errors"),
	ACTIVITY:         LacpMetricsRequestColumnNamesEnum("activity"),
	TIMEOUT:          LacpMetricsRequestColumnNamesEnum("timeout"),
	SYNCHRONIZATION:  LacpMetricsRequestColumnNamesEnum("synchronization"),
	AGGREGATABLE:     LacpMetricsRequestColumnNamesEnum("aggregatable"),
	COLLECTING:       LacpMetricsRequestColumnNamesEnum("collecting"),
	DISTRIBUTING:     LacpMetricsRequestColumnNamesEnum("distributing"),
	SYSTEM_ID:        LacpMetricsRequestColumnNamesEnum("system_id"),
	OPER_KEY:         LacpMetricsRequestColumnNamesEnum("oper_key"),
	PARTNER_ID:       LacpMetricsRequestColumnNamesEnum("partner_id"),
	PARTNER_KEY:      LacpMetricsRequestColumnNamesEnum("partner_key"),
	PORT_NUM:         LacpMetricsRequestColumnNamesEnum("port_num"),
	PARTNER_PORT_NUM: LacpMetricsRequestColumnNamesEnum("partner_port_num"),
}

func (obj *lacpMetricsRequest) ColumnNames() []LacpMetricsRequestColumnNamesEnum {
	items := []LacpMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, LacpMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned. The name of LAG and LAG member can not be excluded.
// SetColumnNames sets the []string value in the LacpMetricsRequest object
func (obj *lacpMetricsRequest) SetColumnNames(value []LacpMetricsRequestColumnNamesEnum) LacpMetricsRequest {

	items := []otg.LacpMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.LacpMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.LacpMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *lacpMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lacpMetricsRequest) setDefault() {

}
