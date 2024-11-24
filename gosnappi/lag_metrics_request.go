package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagMetricsRequest *****
type lagMetricsRequest struct {
	validation
	obj          *otg.LagMetricsRequest
	marshaller   marshalLagMetricsRequest
	unMarshaller unMarshalLagMetricsRequest
}

func NewLagMetricsRequest() LagMetricsRequest {
	obj := lagMetricsRequest{obj: &otg.LagMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *lagMetricsRequest) msg() *otg.LagMetricsRequest {
	return obj.obj
}

func (obj *lagMetricsRequest) setMsg(msg *otg.LagMetricsRequest) LagMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagMetricsRequest struct {
	obj *lagMetricsRequest
}

type marshalLagMetricsRequest interface {
	// ToProto marshals LagMetricsRequest to protobuf object *otg.LagMetricsRequest
	ToProto() (*otg.LagMetricsRequest, error)
	// ToPbText marshals LagMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagMetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LagMetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallagMetricsRequest struct {
	obj *lagMetricsRequest
}

type unMarshalLagMetricsRequest interface {
	// FromProto unmarshals LagMetricsRequest from protobuf object *otg.LagMetricsRequest
	FromProto(msg *otg.LagMetricsRequest) (LagMetricsRequest, error)
	// FromPbText unmarshals LagMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *lagMetricsRequest) Marshal() marshalLagMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagMetricsRequest) Unmarshal() unMarshalLagMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagMetricsRequest) ToProto() (*otg.LagMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagMetricsRequest) FromProto(msg *otg.LagMetricsRequest) (LagMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshallagMetricsRequest) FromPbText(value string) error {
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

func (m *marshallagMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshallagMetricsRequest) FromYaml(value string) error {
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

func (m *marshallagMetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshallagMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshallagMetricsRequest) FromJson(value string) error {
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

func (obj *lagMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagMetricsRequest) Clone() (LagMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagMetricsRequest()
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

// LagMetricsRequest is the request to retrieve per LAG metrics/statistics.
type LagMetricsRequest interface {
	Validation
	// msg marshals LagMetricsRequest to protobuf object *otg.LagMetricsRequest
	// and doesn't set defaults
	msg() *otg.LagMetricsRequest
	// setMsg unmarshals LagMetricsRequest from protobuf object *otg.LagMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.LagMetricsRequest) LagMetricsRequest
	// provides marshal interface
	Marshal() marshalLagMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalLagMetricsRequest
	// validate validates LagMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LagNames returns []string, set in LagMetricsRequest.
	LagNames() []string
	// SetLagNames assigns []string provided by user to LagMetricsRequest
	SetLagNames(value []string) LagMetricsRequest
	// ColumnNames returns []LagMetricsRequestColumnNamesEnum, set in LagMetricsRequest
	ColumnNames() []LagMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []LagMetricsRequestColumnNamesEnum provided by user to LagMetricsRequest
	SetColumnNames(value []LagMetricsRequestColumnNamesEnum) LagMetricsRequest
}

// The names of LAGs to return results for. An empty list will return results for all LAGs.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// LagNames returns a []string
func (obj *lagMetricsRequest) LagNames() []string {
	if obj.obj.LagNames == nil {
		obj.obj.LagNames = make([]string, 0)
	}
	return obj.obj.LagNames
}

// The names of LAGs to return results for. An empty list will return results for all LAGs.
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// x-constraint:
// - /components/schemas/Lag/properties/name
//
// SetLagNames sets the []string value in the LagMetricsRequest object
func (obj *lagMetricsRequest) SetLagNames(value []string) LagMetricsRequest {

	if obj.obj.LagNames == nil {
		obj.obj.LagNames = make([]string, 0)
	}
	obj.obj.LagNames = value

	return obj
}

type LagMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on LagMetricsRequest
var LagMetricsRequestColumnNames = struct {
	OPER_STATUS     LagMetricsRequestColumnNamesEnum
	MEMBER_PORTS_UP LagMetricsRequestColumnNamesEnum
	FRAMES_TX       LagMetricsRequestColumnNamesEnum
	FRAMES_RX       LagMetricsRequestColumnNamesEnum
	BYTES_TX        LagMetricsRequestColumnNamesEnum
	BYTES_RX        LagMetricsRequestColumnNamesEnum
	FRAMES_TX_RATE  LagMetricsRequestColumnNamesEnum
	FRAMES_RX_RATE  LagMetricsRequestColumnNamesEnum
	BYTES_TX_RATE   LagMetricsRequestColumnNamesEnum
	BYTES_RX_RATE   LagMetricsRequestColumnNamesEnum
}{
	OPER_STATUS:     LagMetricsRequestColumnNamesEnum("oper_status"),
	MEMBER_PORTS_UP: LagMetricsRequestColumnNamesEnum("member_ports_up"),
	FRAMES_TX:       LagMetricsRequestColumnNamesEnum("frames_tx"),
	FRAMES_RX:       LagMetricsRequestColumnNamesEnum("frames_rx"),
	BYTES_TX:        LagMetricsRequestColumnNamesEnum("bytes_tx"),
	BYTES_RX:        LagMetricsRequestColumnNamesEnum("bytes_rx"),
	FRAMES_TX_RATE:  LagMetricsRequestColumnNamesEnum("frames_tx_rate"),
	FRAMES_RX_RATE:  LagMetricsRequestColumnNamesEnum("frames_rx_rate"),
	BYTES_TX_RATE:   LagMetricsRequestColumnNamesEnum("bytes_tx_rate"),
	BYTES_RX_RATE:   LagMetricsRequestColumnNamesEnum("bytes_rx_rate"),
}

func (obj *lagMetricsRequest) ColumnNames() []LagMetricsRequestColumnNamesEnum {
	items := []LagMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, LagMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The list of column names that the returned result set will contain. If the list is empty then all columns will be returned. The name of the LAG cannot be excluded.
// SetColumnNames sets the []string value in the LagMetricsRequest object
func (obj *lagMetricsRequest) SetColumnNames(value []LagMetricsRequestColumnNamesEnum) LagMetricsRequest {

	items := []otg.LagMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.LagMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.LagMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *lagMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagMetricsRequest) setDefault() {

}
