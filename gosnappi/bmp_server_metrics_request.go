package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerMetricsRequest *****
type bmpServerMetricsRequest struct {
	validation
	obj          *otg.BmpServerMetricsRequest
	marshaller   marshalBmpServerMetricsRequest
	unMarshaller unMarshalBmpServerMetricsRequest
}

func NewBmpServerMetricsRequest() BmpServerMetricsRequest {
	obj := bmpServerMetricsRequest{obj: &otg.BmpServerMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerMetricsRequest) msg() *otg.BmpServerMetricsRequest {
	return obj.obj
}

func (obj *bmpServerMetricsRequest) setMsg(msg *otg.BmpServerMetricsRequest) BmpServerMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerMetricsRequest struct {
	obj *bmpServerMetricsRequest
}

type marshalBmpServerMetricsRequest interface {
	// ToProto marshals BmpServerMetricsRequest to protobuf object *otg.BmpServerMetricsRequest
	ToProto() (*otg.BmpServerMetricsRequest, error)
	// ToPbText marshals BmpServerMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerMetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerMetricsRequest struct {
	obj *bmpServerMetricsRequest
}

type unMarshalBmpServerMetricsRequest interface {
	// FromProto unmarshals BmpServerMetricsRequest from protobuf object *otg.BmpServerMetricsRequest
	FromProto(msg *otg.BmpServerMetricsRequest) (BmpServerMetricsRequest, error)
	// FromPbText unmarshals BmpServerMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *bmpServerMetricsRequest) Marshal() marshalBmpServerMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerMetricsRequest) Unmarshal() unMarshalBmpServerMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerMetricsRequest) ToProto() (*otg.BmpServerMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerMetricsRequest) FromProto(msg *otg.BmpServerMetricsRequest) (BmpServerMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerMetricsRequest) FromPbText(value string) error {
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

func (m *marshalbmpServerMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerMetricsRequest) FromYaml(value string) error {
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

func (m *marshalbmpServerMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalbmpServerMetricsRequest) FromJson(value string) error {
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

func (obj *bmpServerMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerMetricsRequest) Clone() (BmpServerMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerMetricsRequest()
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

// BmpServerMetricsRequest is the request to retrieve per BMP Server metrics/statistics.
type BmpServerMetricsRequest interface {
	Validation
	// msg marshals BmpServerMetricsRequest to protobuf object *otg.BmpServerMetricsRequest
	// and doesn't set defaults
	msg() *otg.BmpServerMetricsRequest
	// setMsg unmarshals BmpServerMetricsRequest from protobuf object *otg.BmpServerMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.BmpServerMetricsRequest) BmpServerMetricsRequest
	// provides marshal interface
	Marshal() marshalBmpServerMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerMetricsRequest
	// validate validates BmpServerMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServerNames returns []string, set in BmpServerMetricsRequest.
	ServerNames() []string
	// SetServerNames assigns []string provided by user to BmpServerMetricsRequest
	SetServerNames(value []string) BmpServerMetricsRequest
	// ColumnNames returns []BmpServerMetricsRequestColumnNamesEnum, set in BmpServerMetricsRequest
	ColumnNames() []BmpServerMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []BmpServerMetricsRequestColumnNamesEnum provided by user to BmpServerMetricsRequest
	SetColumnNames(value []BmpServerMetricsRequestColumnNamesEnum) BmpServerMetricsRequest
}

// The names of BMP Servers to return results for. An empty list will return results for all BMP Servers.
//
// x-constraint:
// - /components/schemas/Device.Bmp.ServerV4/properties/name
// - /components/schemas/Device.Bmp.ServerV6/properties/name
//
// ServerNames returns a []string
func (obj *bmpServerMetricsRequest) ServerNames() []string {
	if obj.obj.ServerNames == nil {
		obj.obj.ServerNames = make([]string, 0)
	}
	return obj.obj.ServerNames
}

// The names of BMP Servers to return results for. An empty list will return results for all BMP Servers.
//
// x-constraint:
// - /components/schemas/Device.Bmp.ServerV4/properties/name
// - /components/schemas/Device.Bmp.ServerV6/properties/name
//
// SetServerNames sets the []string value in the BmpServerMetricsRequest object
func (obj *bmpServerMetricsRequest) SetServerNames(value []string) BmpServerMetricsRequest {

	if obj.obj.ServerNames == nil {
		obj.obj.ServerNames = make([]string, 0)
	}
	obj.obj.ServerNames = value

	return obj
}

type BmpServerMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on BmpServerMetricsRequest
var BmpServerMetricsRequestColumnNames = struct {
	SESSION_STATE                            BmpServerMetricsRequestColumnNamesEnum
	FLAP_COUNT                               BmpServerMetricsRequestColumnNamesEnum
	ROUTE_MONITORING_MESSAGES_RECEIVED       BmpServerMetricsRequestColumnNamesEnum
	STATISTICS_MESSAGES_RECEIVED             BmpServerMetricsRequestColumnNamesEnum
	PEER_DOWN_MESSAGES_RECEIVED              BmpServerMetricsRequestColumnNamesEnum
	PEER_UP_MESSAGES_RECEIVED                BmpServerMetricsRequestColumnNamesEnum
	INITIATION_MESSAGES_RECEIVED             BmpServerMetricsRequestColumnNamesEnum
	ROUTE_MIRRORING_MESSAGES_RECEIVED        BmpServerMetricsRequestColumnNamesEnum
	TERMINATION_MESSAGES_RECEIVED            BmpServerMetricsRequestColumnNamesEnum
	PRE_POLICY_IPV4_UNICAST_ROUTES_RECEIVED  BmpServerMetricsRequestColumnNamesEnum
	POST_POLICY_IPV4_UNICAST_ROUTES_RECEIVED BmpServerMetricsRequestColumnNamesEnum
	PRE_POLICY_IPV6_UNICAST_ROUTES_RECEIVED  BmpServerMetricsRequestColumnNamesEnum
	POST_POLICY_IPV6_UNICAST_ROUTES_RECEIVED BmpServerMetricsRequestColumnNamesEnum
}{
	SESSION_STATE:                            BmpServerMetricsRequestColumnNamesEnum("session_state"),
	FLAP_COUNT:                               BmpServerMetricsRequestColumnNamesEnum("flap_count"),
	ROUTE_MONITORING_MESSAGES_RECEIVED:       BmpServerMetricsRequestColumnNamesEnum("route_monitoring_messages_received"),
	STATISTICS_MESSAGES_RECEIVED:             BmpServerMetricsRequestColumnNamesEnum("statistics_messages_received"),
	PEER_DOWN_MESSAGES_RECEIVED:              BmpServerMetricsRequestColumnNamesEnum("peer_down_messages_received"),
	PEER_UP_MESSAGES_RECEIVED:                BmpServerMetricsRequestColumnNamesEnum("peer_up_messages_received"),
	INITIATION_MESSAGES_RECEIVED:             BmpServerMetricsRequestColumnNamesEnum("initiation_messages_received"),
	ROUTE_MIRRORING_MESSAGES_RECEIVED:        BmpServerMetricsRequestColumnNamesEnum("route_mirroring_messages_received"),
	TERMINATION_MESSAGES_RECEIVED:            BmpServerMetricsRequestColumnNamesEnum("termination_messages_received"),
	PRE_POLICY_IPV4_UNICAST_ROUTES_RECEIVED:  BmpServerMetricsRequestColumnNamesEnum("pre_policy_ipv4_unicast_routes_received"),
	POST_POLICY_IPV4_UNICAST_ROUTES_RECEIVED: BmpServerMetricsRequestColumnNamesEnum("post_policy_ipv4_unicast_routes_received"),
	PRE_POLICY_IPV6_UNICAST_ROUTES_RECEIVED:  BmpServerMetricsRequestColumnNamesEnum("pre_policy_ipv6_unicast_routes_received"),
	POST_POLICY_IPV6_UNICAST_ROUTES_RECEIVED: BmpServerMetricsRequestColumnNamesEnum("post_policy_ipv6_unicast_routes_received"),
}

func (obj *bmpServerMetricsRequest) ColumnNames() []BmpServerMetricsRequestColumnNamesEnum {
	items := []BmpServerMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, BmpServerMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned except for any result_groups. The name of the BMP Server cannot be excluded.
// SetColumnNames sets the []string value in the BmpServerMetricsRequest object
func (obj *bmpServerMetricsRequest) SetColumnNames(value []BmpServerMetricsRequestColumnNamesEnum) BmpServerMetricsRequest {

	items := []otg.BmpServerMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.BmpServerMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.BmpServerMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *bmpServerMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bmpServerMetricsRequest) setDefault() {

}
