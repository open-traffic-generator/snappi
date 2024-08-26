package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2SummaryAsLsa *****
type ospfv2SummaryAsLsa struct {
	validation
	obj               *otg.Ospfv2SummaryAsLsa
	marshaller        marshalOspfv2SummaryAsLsa
	unMarshaller      unMarshalOspfv2SummaryAsLsa
	commonAttrsHolder Ospfv2CommonAttrs
}

func NewOspfv2SummaryAsLsa() Ospfv2SummaryAsLsa {
	obj := ospfv2SummaryAsLsa{obj: &otg.Ospfv2SummaryAsLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2SummaryAsLsa) msg() *otg.Ospfv2SummaryAsLsa {
	return obj.obj
}

func (obj *ospfv2SummaryAsLsa) setMsg(msg *otg.Ospfv2SummaryAsLsa) Ospfv2SummaryAsLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2SummaryAsLsa struct {
	obj *ospfv2SummaryAsLsa
}

type marshalOspfv2SummaryAsLsa interface {
	// ToProto marshals Ospfv2SummaryAsLsa to protobuf object *otg.Ospfv2SummaryAsLsa
	ToProto() (*otg.Ospfv2SummaryAsLsa, error)
	// ToPbText marshals Ospfv2SummaryAsLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2SummaryAsLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2SummaryAsLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2SummaryAsLsa struct {
	obj *ospfv2SummaryAsLsa
}

type unMarshalOspfv2SummaryAsLsa interface {
	// FromProto unmarshals Ospfv2SummaryAsLsa from protobuf object *otg.Ospfv2SummaryAsLsa
	FromProto(msg *otg.Ospfv2SummaryAsLsa) (Ospfv2SummaryAsLsa, error)
	// FromPbText unmarshals Ospfv2SummaryAsLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2SummaryAsLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2SummaryAsLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2SummaryAsLsa) Marshal() marshalOspfv2SummaryAsLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2SummaryAsLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2SummaryAsLsa) Unmarshal() unMarshalOspfv2SummaryAsLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2SummaryAsLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2SummaryAsLsa) ToProto() (*otg.Ospfv2SummaryAsLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2SummaryAsLsa) FromProto(msg *otg.Ospfv2SummaryAsLsa) (Ospfv2SummaryAsLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2SummaryAsLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2SummaryAsLsa) FromPbText(value string) error {
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

func (m *marshalospfv2SummaryAsLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2SummaryAsLsa) FromYaml(value string) error {
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

func (m *marshalospfv2SummaryAsLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2SummaryAsLsa) FromJson(value string) error {
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

func (obj *ospfv2SummaryAsLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2SummaryAsLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2SummaryAsLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2SummaryAsLsa) Clone() (Ospfv2SummaryAsLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2SummaryAsLsa()
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

func (obj *ospfv2SummaryAsLsa) setNil() {
	obj.commonAttrsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2SummaryAsLsa is contents of OSPFv2 Autonomous System Boundary Router (ASBR) summary LSA - Type 4.
type Ospfv2SummaryAsLsa interface {
	Validation
	// msg marshals Ospfv2SummaryAsLsa to protobuf object *otg.Ospfv2SummaryAsLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2SummaryAsLsa
	// setMsg unmarshals Ospfv2SummaryAsLsa from protobuf object *otg.Ospfv2SummaryAsLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2SummaryAsLsa) Ospfv2SummaryAsLsa
	// provides marshal interface
	Marshal() marshalOspfv2SummaryAsLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2SummaryAsLsa
	// validate validates Ospfv2SummaryAsLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2SummaryAsLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonAttrs returns Ospfv2CommonAttrs, set in Ospfv2SummaryAsLsa.
	// Ospfv2CommonAttrs is tBD
	CommonAttrs() Ospfv2CommonAttrs
	// SetCommonAttrs assigns Ospfv2CommonAttrs provided by user to Ospfv2SummaryAsLsa.
	// Ospfv2CommonAttrs is tBD
	SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2SummaryAsLsa
	// NetworkMask returns string, set in Ospfv2SummaryAsLsa.
	NetworkMask() string
	// SetNetworkMask assigns string provided by user to Ospfv2SummaryAsLsa
	SetNetworkMask(value string) Ospfv2SummaryAsLsa
	// HasNetworkMask checks if NetworkMask has been set in Ospfv2SummaryAsLsa
	HasNetworkMask() bool
	// Metric returns uint32, set in Ospfv2SummaryAsLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv2SummaryAsLsa
	SetMetric(value uint32) Ospfv2SummaryAsLsa
	// HasMetric checks if Metric has been set in Ospfv2SummaryAsLsa
	HasMetric() bool
	setNil()
}

// Common LSA attributes.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2SummaryAsLsa) CommonAttrs() Ospfv2CommonAttrs {
	if obj.obj.CommonAttrs == nil {
		obj.obj.CommonAttrs = NewOspfv2CommonAttrs().msg()
	}
	if obj.commonAttrsHolder == nil {
		obj.commonAttrsHolder = &ospfv2CommonAttrs{obj: obj.obj.CommonAttrs}
	}
	return obj.commonAttrsHolder
}

// Common LSA attributes.
// SetCommonAttrs sets the Ospfv2CommonAttrs value in the Ospfv2SummaryAsLsa object
func (obj *ospfv2SummaryAsLsa) SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2SummaryAsLsa {

	obj.commonAttrsHolder = nil
	obj.obj.CommonAttrs = value.msg()

	return obj
}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2SummaryAsLsa) NetworkMask() string {

	return *obj.obj.NetworkMask

}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2SummaryAsLsa) HasNetworkMask() bool {
	return obj.obj.NetworkMask != nil
}

// The IPv4 address mask for the network.
// SetNetworkMask sets the string value in the Ospfv2SummaryAsLsa object
func (obj *ospfv2SummaryAsLsa) SetNetworkMask(value string) Ospfv2SummaryAsLsa {

	obj.obj.NetworkMask = &value
	return obj
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2SummaryAsLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2SummaryAsLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// SetMetric sets the uint32 value in the Ospfv2SummaryAsLsa object
func (obj *ospfv2SummaryAsLsa) SetMetric(value uint32) Ospfv2SummaryAsLsa {

	obj.obj.Metric = &value
	return obj
}

func (obj *ospfv2SummaryAsLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// CommonAttrs is required
	if obj.obj.CommonAttrs == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CommonAttrs is required field on interface Ospfv2SummaryAsLsa")
	}

	if obj.obj.CommonAttrs != nil {

		obj.CommonAttrs().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkMask != nil {

		err := obj.validateIpv4(obj.NetworkMask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2SummaryAsLsa.NetworkMask"))
		}

	}

}

func (obj *ospfv2SummaryAsLsa) setDefault() {

}
