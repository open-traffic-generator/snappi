package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2ExternalAsLsa *****
type ospfv2ExternalAsLsa struct {
	validation
	obj          *otg.Ospfv2ExternalAsLsa
	marshaller   marshalOspfv2ExternalAsLsa
	unMarshaller unMarshalOspfv2ExternalAsLsa
	headerHolder Ospfv2LsaHeader
}

func NewOspfv2ExternalAsLsa() Ospfv2ExternalAsLsa {
	obj := ospfv2ExternalAsLsa{obj: &otg.Ospfv2ExternalAsLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2ExternalAsLsa) msg() *otg.Ospfv2ExternalAsLsa {
	return obj.obj
}

func (obj *ospfv2ExternalAsLsa) setMsg(msg *otg.Ospfv2ExternalAsLsa) Ospfv2ExternalAsLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2ExternalAsLsa struct {
	obj *ospfv2ExternalAsLsa
}

type marshalOspfv2ExternalAsLsa interface {
	// ToProto marshals Ospfv2ExternalAsLsa to protobuf object *otg.Ospfv2ExternalAsLsa
	ToProto() (*otg.Ospfv2ExternalAsLsa, error)
	// ToPbText marshals Ospfv2ExternalAsLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2ExternalAsLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2ExternalAsLsa to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Ospfv2ExternalAsLsa to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalospfv2ExternalAsLsa struct {
	obj *ospfv2ExternalAsLsa
}

type unMarshalOspfv2ExternalAsLsa interface {
	// FromProto unmarshals Ospfv2ExternalAsLsa from protobuf object *otg.Ospfv2ExternalAsLsa
	FromProto(msg *otg.Ospfv2ExternalAsLsa) (Ospfv2ExternalAsLsa, error)
	// FromPbText unmarshals Ospfv2ExternalAsLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2ExternalAsLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2ExternalAsLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2ExternalAsLsa) Marshal() marshalOspfv2ExternalAsLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2ExternalAsLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2ExternalAsLsa) Unmarshal() unMarshalOspfv2ExternalAsLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2ExternalAsLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2ExternalAsLsa) ToProto() (*otg.Ospfv2ExternalAsLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2ExternalAsLsa) FromProto(msg *otg.Ospfv2ExternalAsLsa) (Ospfv2ExternalAsLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2ExternalAsLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2ExternalAsLsa) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalospfv2ExternalAsLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2ExternalAsLsa) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalospfv2ExternalAsLsa) ToJsonRaw() (string, error) {
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

func (m *marshalospfv2ExternalAsLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2ExternalAsLsa) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ospfv2ExternalAsLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2ExternalAsLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2ExternalAsLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2ExternalAsLsa) Clone() (Ospfv2ExternalAsLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2ExternalAsLsa()
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

func (obj *ospfv2ExternalAsLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2ExternalAsLsa is contents of OSPFv2 AS-External-LSA - Type 5.
// The value of the IPv4 prefix that was received is present in header.lsa_id.
type Ospfv2ExternalAsLsa interface {
	Validation
	// msg marshals Ospfv2ExternalAsLsa to protobuf object *otg.Ospfv2ExternalAsLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2ExternalAsLsa
	// setMsg unmarshals Ospfv2ExternalAsLsa from protobuf object *otg.Ospfv2ExternalAsLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2ExternalAsLsa) Ospfv2ExternalAsLsa
	// provides marshal interface
	Marshal() marshalOspfv2ExternalAsLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2ExternalAsLsa
	// validate validates Ospfv2ExternalAsLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2ExternalAsLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv2LsaHeader, set in Ospfv2ExternalAsLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	Header() Ospfv2LsaHeader
	// SetHeader assigns Ospfv2LsaHeader provided by user to Ospfv2ExternalAsLsa.
	// Ospfv2LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv2LsaHeader) Ospfv2ExternalAsLsa
	// HasHeader checks if Header has been set in Ospfv2ExternalAsLsa
	HasHeader() bool
	// NetworkMask returns string, set in Ospfv2ExternalAsLsa.
	NetworkMask() string
	// SetNetworkMask assigns string provided by user to Ospfv2ExternalAsLsa
	SetNetworkMask(value string) Ospfv2ExternalAsLsa
	// HasNetworkMask checks if NetworkMask has been set in Ospfv2ExternalAsLsa
	HasNetworkMask() bool
	// Metric returns uint32, set in Ospfv2ExternalAsLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv2ExternalAsLsa
	SetMetric(value uint32) Ospfv2ExternalAsLsa
	// HasMetric checks if Metric has been set in Ospfv2ExternalAsLsa
	HasMetric() bool
	// MetricType returns uint32, set in Ospfv2ExternalAsLsa.
	MetricType() uint32
	// SetMetricType assigns uint32 provided by user to Ospfv2ExternalAsLsa
	SetMetricType(value uint32) Ospfv2ExternalAsLsa
	// HasMetricType checks if MetricType has been set in Ospfv2ExternalAsLsa
	HasMetricType() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv2LsaHeader
func (obj *ospfv2ExternalAsLsa) Header() Ospfv2LsaHeader {
	if obj.obj.Header == nil {
		obj.obj.Header = NewOspfv2LsaHeader().msg()
	}
	if obj.headerHolder == nil {
		obj.headerHolder = &ospfv2LsaHeader{obj: obj.obj.Header}
	}
	return obj.headerHolder
}

// Contents of the LSA header.
// Header returns a Ospfv2LsaHeader
func (obj *ospfv2ExternalAsLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv2LsaHeader value in the Ospfv2ExternalAsLsa object
func (obj *ospfv2ExternalAsLsa) SetHeader(value Ospfv2LsaHeader) Ospfv2ExternalAsLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2ExternalAsLsa) NetworkMask() string {

	return *obj.obj.NetworkMask

}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2ExternalAsLsa) HasNetworkMask() bool {
	return obj.obj.NetworkMask != nil
}

// The IPv4 address mask for the network.
// SetNetworkMask sets the string value in the Ospfv2ExternalAsLsa object
func (obj *ospfv2ExternalAsLsa) SetNetworkMask(value string) Ospfv2ExternalAsLsa {

	obj.obj.NetworkMask = &value
	return obj
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2ExternalAsLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2ExternalAsLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// SetMetric sets the uint32 value in the Ospfv2ExternalAsLsa object
func (obj *ospfv2ExternalAsLsa) SetMetric(value uint32) Ospfv2ExternalAsLsa {

	obj.obj.Metric = &value
	return obj
}

// The type of metric associated with the route range.
// MetricType returns a uint32
func (obj *ospfv2ExternalAsLsa) MetricType() uint32 {

	return *obj.obj.MetricType

}

// The type of metric associated with the route range.
// MetricType returns a uint32
func (obj *ospfv2ExternalAsLsa) HasMetricType() bool {
	return obj.obj.MetricType != nil
}

// The type of metric associated with the route range.
// SetMetricType sets the uint32 value in the Ospfv2ExternalAsLsa object
func (obj *ospfv2ExternalAsLsa) SetMetricType(value uint32) Ospfv2ExternalAsLsa {

	obj.obj.MetricType = &value
	return obj
}

func (obj *ospfv2ExternalAsLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkMask != nil {

		err := obj.validateIpv4(obj.NetworkMask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2ExternalAsLsa.NetworkMask"))
		}

	}

}

func (obj *ospfv2ExternalAsLsa) setDefault() {

}
