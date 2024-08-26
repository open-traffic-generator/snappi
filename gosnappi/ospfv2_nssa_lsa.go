package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2NssaLsa *****
type ospfv2NssaLsa struct {
	validation
	obj               *otg.Ospfv2NssaLsa
	marshaller        marshalOspfv2NssaLsa
	unMarshaller      unMarshalOspfv2NssaLsa
	commonAttrsHolder Ospfv2CommonAttrs
}

func NewOspfv2NssaLsa() Ospfv2NssaLsa {
	obj := ospfv2NssaLsa{obj: &otg.Ospfv2NssaLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2NssaLsa) msg() *otg.Ospfv2NssaLsa {
	return obj.obj
}

func (obj *ospfv2NssaLsa) setMsg(msg *otg.Ospfv2NssaLsa) Ospfv2NssaLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2NssaLsa struct {
	obj *ospfv2NssaLsa
}

type marshalOspfv2NssaLsa interface {
	// ToProto marshals Ospfv2NssaLsa to protobuf object *otg.Ospfv2NssaLsa
	ToProto() (*otg.Ospfv2NssaLsa, error)
	// ToPbText marshals Ospfv2NssaLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2NssaLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2NssaLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2NssaLsa struct {
	obj *ospfv2NssaLsa
}

type unMarshalOspfv2NssaLsa interface {
	// FromProto unmarshals Ospfv2NssaLsa from protobuf object *otg.Ospfv2NssaLsa
	FromProto(msg *otg.Ospfv2NssaLsa) (Ospfv2NssaLsa, error)
	// FromPbText unmarshals Ospfv2NssaLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2NssaLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2NssaLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2NssaLsa) Marshal() marshalOspfv2NssaLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2NssaLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2NssaLsa) Unmarshal() unMarshalOspfv2NssaLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2NssaLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2NssaLsa) ToProto() (*otg.Ospfv2NssaLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2NssaLsa) FromProto(msg *otg.Ospfv2NssaLsa) (Ospfv2NssaLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2NssaLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2NssaLsa) FromPbText(value string) error {
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

func (m *marshalospfv2NssaLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2NssaLsa) FromYaml(value string) error {
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

func (m *marshalospfv2NssaLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2NssaLsa) FromJson(value string) error {
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

func (obj *ospfv2NssaLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2NssaLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2NssaLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2NssaLsa) Clone() (Ospfv2NssaLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2NssaLsa()
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

func (obj *ospfv2NssaLsa) setNil() {
	obj.commonAttrsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2NssaLsa is contents of OSPFv2 NSSA-LSA - Type 7.
type Ospfv2NssaLsa interface {
	Validation
	// msg marshals Ospfv2NssaLsa to protobuf object *otg.Ospfv2NssaLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2NssaLsa
	// setMsg unmarshals Ospfv2NssaLsa from protobuf object *otg.Ospfv2NssaLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2NssaLsa) Ospfv2NssaLsa
	// provides marshal interface
	Marshal() marshalOspfv2NssaLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2NssaLsa
	// validate validates Ospfv2NssaLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2NssaLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonAttrs returns Ospfv2CommonAttrs, set in Ospfv2NssaLsa.
	// Ospfv2CommonAttrs is tBD
	CommonAttrs() Ospfv2CommonAttrs
	// SetCommonAttrs assigns Ospfv2CommonAttrs provided by user to Ospfv2NssaLsa.
	// Ospfv2CommonAttrs is tBD
	SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2NssaLsa
	// NetworkMask returns string, set in Ospfv2NssaLsa.
	NetworkMask() string
	// SetNetworkMask assigns string provided by user to Ospfv2NssaLsa
	SetNetworkMask(value string) Ospfv2NssaLsa
	// HasNetworkMask checks if NetworkMask has been set in Ospfv2NssaLsa
	HasNetworkMask() bool
	// Metric returns uint32, set in Ospfv2NssaLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv2NssaLsa
	SetMetric(value uint32) Ospfv2NssaLsa
	// HasMetric checks if Metric has been set in Ospfv2NssaLsa
	HasMetric() bool
	// MetricType returns uint32, set in Ospfv2NssaLsa.
	MetricType() uint32
	// SetMetricType assigns uint32 provided by user to Ospfv2NssaLsa
	SetMetricType(value uint32) Ospfv2NssaLsa
	// HasMetricType checks if MetricType has been set in Ospfv2NssaLsa
	HasMetricType() bool
	// ForwardingAddress returns string, set in Ospfv2NssaLsa.
	ForwardingAddress() string
	// SetForwardingAddress assigns string provided by user to Ospfv2NssaLsa
	SetForwardingAddress(value string) Ospfv2NssaLsa
	// HasForwardingAddress checks if ForwardingAddress has been set in Ospfv2NssaLsa
	HasForwardingAddress() bool
	setNil()
}

// Common LSA attributes.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2NssaLsa) CommonAttrs() Ospfv2CommonAttrs {
	if obj.obj.CommonAttrs == nil {
		obj.obj.CommonAttrs = NewOspfv2CommonAttrs().msg()
	}
	if obj.commonAttrsHolder == nil {
		obj.commonAttrsHolder = &ospfv2CommonAttrs{obj: obj.obj.CommonAttrs}
	}
	return obj.commonAttrsHolder
}

// Common LSA attributes.
// SetCommonAttrs sets the Ospfv2CommonAttrs value in the Ospfv2NssaLsa object
func (obj *ospfv2NssaLsa) SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2NssaLsa {

	obj.commonAttrsHolder = nil
	obj.obj.CommonAttrs = value.msg()

	return obj
}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2NssaLsa) NetworkMask() string {

	return *obj.obj.NetworkMask

}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2NssaLsa) HasNetworkMask() bool {
	return obj.obj.NetworkMask != nil
}

// The IPv4 address mask for the network.
// SetNetworkMask sets the string value in the Ospfv2NssaLsa object
func (obj *ospfv2NssaLsa) SetNetworkMask(value string) Ospfv2NssaLsa {

	obj.obj.NetworkMask = &value
	return obj
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2NssaLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2NssaLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// SetMetric sets the uint32 value in the Ospfv2NssaLsa object
func (obj *ospfv2NssaLsa) SetMetric(value uint32) Ospfv2NssaLsa {

	obj.obj.Metric = &value
	return obj
}

// The type of metric associated with the route range.
// MetricType returns a uint32
func (obj *ospfv2NssaLsa) MetricType() uint32 {

	return *obj.obj.MetricType

}

// The type of metric associated with the route range.
// MetricType returns a uint32
func (obj *ospfv2NssaLsa) HasMetricType() bool {
	return obj.obj.MetricType != nil
}

// The type of metric associated with the route range.
// SetMetricType sets the uint32 value in the Ospfv2NssaLsa object
func (obj *ospfv2NssaLsa) SetMetricType(value uint32) Ospfv2NssaLsa {

	obj.obj.MetricType = &value
	return obj
}

// IPv4 Forwarding address.
// ForwardingAddress returns a string
func (obj *ospfv2NssaLsa) ForwardingAddress() string {

	return *obj.obj.ForwardingAddress

}

// IPv4 Forwarding address.
// ForwardingAddress returns a string
func (obj *ospfv2NssaLsa) HasForwardingAddress() bool {
	return obj.obj.ForwardingAddress != nil
}

// IPv4 Forwarding address.
// SetForwardingAddress sets the string value in the Ospfv2NssaLsa object
func (obj *ospfv2NssaLsa) SetForwardingAddress(value string) Ospfv2NssaLsa {

	obj.obj.ForwardingAddress = &value
	return obj
}

func (obj *ospfv2NssaLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// CommonAttrs is required
	if obj.obj.CommonAttrs == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CommonAttrs is required field on interface Ospfv2NssaLsa")
	}

	if obj.obj.CommonAttrs != nil {

		obj.CommonAttrs().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkMask != nil {

		err := obj.validateIpv4(obj.NetworkMask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2NssaLsa.NetworkMask"))
		}

	}

	if obj.obj.ForwardingAddress != nil {

		err := obj.validateIpv4(obj.ForwardingAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2NssaLsa.ForwardingAddress"))
		}

	}

}

func (obj *ospfv2NssaLsa) setDefault() {

}
