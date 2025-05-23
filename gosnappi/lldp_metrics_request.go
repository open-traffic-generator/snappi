package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpMetricsRequest *****
type lldpMetricsRequest struct {
	validation
	obj          *otg.LldpMetricsRequest
	marshaller   marshalLldpMetricsRequest
	unMarshaller unMarshalLldpMetricsRequest
}

func NewLldpMetricsRequest() LldpMetricsRequest {
	obj := lldpMetricsRequest{obj: &otg.LldpMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpMetricsRequest) msg() *otg.LldpMetricsRequest {
	return obj.obj
}

func (obj *lldpMetricsRequest) setMsg(msg *otg.LldpMetricsRequest) LldpMetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpMetricsRequest struct {
	obj *lldpMetricsRequest
}

type marshalLldpMetricsRequest interface {
	// ToProto marshals LldpMetricsRequest to protobuf object *otg.LldpMetricsRequest
	ToProto() (*otg.LldpMetricsRequest, error)
	// ToPbText marshals LldpMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpMetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpMetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpMetricsRequest struct {
	obj *lldpMetricsRequest
}

type unMarshalLldpMetricsRequest interface {
	// FromProto unmarshals LldpMetricsRequest from protobuf object *otg.LldpMetricsRequest
	FromProto(msg *otg.LldpMetricsRequest) (LldpMetricsRequest, error)
	// FromPbText unmarshals LldpMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *lldpMetricsRequest) Marshal() marshalLldpMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpMetricsRequest) Unmarshal() unMarshalLldpMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpMetricsRequest) ToProto() (*otg.LldpMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpMetricsRequest) FromProto(msg *otg.LldpMetricsRequest) (LldpMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshallldpMetricsRequest) FromPbText(value string) error {
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

func (m *marshallldpMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshallldpMetricsRequest) FromYaml(value string) error {
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

func (m *marshallldpMetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshallldpMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshallldpMetricsRequest) FromJson(value string) error {
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

func (obj *lldpMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpMetricsRequest) Clone() (LldpMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpMetricsRequest()
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

// LldpMetricsRequest is the request to retrieve LLDP per instance metrics/statistics.
type LldpMetricsRequest interface {
	Validation
	// msg marshals LldpMetricsRequest to protobuf object *otg.LldpMetricsRequest
	// and doesn't set defaults
	msg() *otg.LldpMetricsRequest
	// setMsg unmarshals LldpMetricsRequest from protobuf object *otg.LldpMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.LldpMetricsRequest) LldpMetricsRequest
	// provides marshal interface
	Marshal() marshalLldpMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalLldpMetricsRequest
	// validate validates LldpMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LldpNames returns []string, set in LldpMetricsRequest.
	LldpNames() []string
	// SetLldpNames assigns []string provided by user to LldpMetricsRequest
	SetLldpNames(value []string) LldpMetricsRequest
	// ColumnNames returns []LldpMetricsRequestColumnNamesEnum, set in LldpMetricsRequest
	ColumnNames() []LldpMetricsRequestColumnNamesEnum
	// SetColumnNames assigns []LldpMetricsRequestColumnNamesEnum provided by user to LldpMetricsRequest
	SetColumnNames(value []LldpMetricsRequestColumnNamesEnum) LldpMetricsRequest
}

// The names of LLDP instances to return results for. An empty list will return results for all LLDP instances.
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// LldpNames returns a []string
func (obj *lldpMetricsRequest) LldpNames() []string {
	if obj.obj.LldpNames == nil {
		obj.obj.LldpNames = make([]string, 0)
	}
	return obj.obj.LldpNames
}

// The names of LLDP instances to return results for. An empty list will return results for all LLDP instances.
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// SetLldpNames sets the []string value in the LldpMetricsRequest object
func (obj *lldpMetricsRequest) SetLldpNames(value []string) LldpMetricsRequest {

	if obj.obj.LldpNames == nil {
		obj.obj.LldpNames = make([]string, 0)
	}
	obj.obj.LldpNames = value

	return obj
}

type LldpMetricsRequestColumnNamesEnum string

// Enum of ColumnNames on LldpMetricsRequest
var LldpMetricsRequestColumnNames = struct {
	FRAMES_RX       LldpMetricsRequestColumnNamesEnum
	FRAMES_TX       LldpMetricsRequestColumnNamesEnum
	FRAMES_ERROR_RX LldpMetricsRequestColumnNamesEnum
	FRAMES_DISCARD  LldpMetricsRequestColumnNamesEnum
	TLVS_DISCARD    LldpMetricsRequestColumnNamesEnum
	TLVS_UNKNOWN    LldpMetricsRequestColumnNamesEnum
}{
	FRAMES_RX:       LldpMetricsRequestColumnNamesEnum("frames_rx"),
	FRAMES_TX:       LldpMetricsRequestColumnNamesEnum("frames_tx"),
	FRAMES_ERROR_RX: LldpMetricsRequestColumnNamesEnum("frames_error_rx"),
	FRAMES_DISCARD:  LldpMetricsRequestColumnNamesEnum("frames_discard"),
	TLVS_DISCARD:    LldpMetricsRequestColumnNamesEnum("tlvs_discard"),
	TLVS_UNKNOWN:    LldpMetricsRequestColumnNamesEnum("tlvs_unknown"),
}

func (obj *lldpMetricsRequest) ColumnNames() []LldpMetricsRequestColumnNamesEnum {
	items := []LldpMetricsRequestColumnNamesEnum{}
	for _, item := range obj.obj.ColumnNames {
		items = append(items, LldpMetricsRequestColumnNamesEnum(item.String()))
	}
	return items
}

// The requested list of column names for the result set. If the list is empty then metrics for all columns will be returned. The name of LLDP instance can not be excluded.
// SetColumnNames sets the []string value in the LldpMetricsRequest object
func (obj *lldpMetricsRequest) SetColumnNames(value []LldpMetricsRequestColumnNamesEnum) LldpMetricsRequest {

	items := []otg.LldpMetricsRequest_ColumnNames_Enum{}
	for _, item := range value {
		intValue := otg.LldpMetricsRequest_ColumnNames_Enum_value[string(item)]
		items = append(items, otg.LldpMetricsRequest_ColumnNames_Enum(intValue))
	}
	obj.obj.ColumnNames = items
	return obj
}

func (obj *lldpMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpMetricsRequest) setDefault() {

}
