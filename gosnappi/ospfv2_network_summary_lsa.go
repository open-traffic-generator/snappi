package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2NetworkSummaryLsa *****
type ospfv2NetworkSummaryLsa struct {
	validation
	obj               *otg.Ospfv2NetworkSummaryLsa
	marshaller        marshalOspfv2NetworkSummaryLsa
	unMarshaller      unMarshalOspfv2NetworkSummaryLsa
	commonAttrsHolder Ospfv2CommonAttrs
}

func NewOspfv2NetworkSummaryLsa() Ospfv2NetworkSummaryLsa {
	obj := ospfv2NetworkSummaryLsa{obj: &otg.Ospfv2NetworkSummaryLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2NetworkSummaryLsa) msg() *otg.Ospfv2NetworkSummaryLsa {
	return obj.obj
}

func (obj *ospfv2NetworkSummaryLsa) setMsg(msg *otg.Ospfv2NetworkSummaryLsa) Ospfv2NetworkSummaryLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2NetworkSummaryLsa struct {
	obj *ospfv2NetworkSummaryLsa
}

type marshalOspfv2NetworkSummaryLsa interface {
	// ToProto marshals Ospfv2NetworkSummaryLsa to protobuf object *otg.Ospfv2NetworkSummaryLsa
	ToProto() (*otg.Ospfv2NetworkSummaryLsa, error)
	// ToPbText marshals Ospfv2NetworkSummaryLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2NetworkSummaryLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2NetworkSummaryLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2NetworkSummaryLsa struct {
	obj *ospfv2NetworkSummaryLsa
}

type unMarshalOspfv2NetworkSummaryLsa interface {
	// FromProto unmarshals Ospfv2NetworkSummaryLsa from protobuf object *otg.Ospfv2NetworkSummaryLsa
	FromProto(msg *otg.Ospfv2NetworkSummaryLsa) (Ospfv2NetworkSummaryLsa, error)
	// FromPbText unmarshals Ospfv2NetworkSummaryLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2NetworkSummaryLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2NetworkSummaryLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv2NetworkSummaryLsa) Marshal() marshalOspfv2NetworkSummaryLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2NetworkSummaryLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2NetworkSummaryLsa) Unmarshal() unMarshalOspfv2NetworkSummaryLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2NetworkSummaryLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2NetworkSummaryLsa) ToProto() (*otg.Ospfv2NetworkSummaryLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2NetworkSummaryLsa) FromProto(msg *otg.Ospfv2NetworkSummaryLsa) (Ospfv2NetworkSummaryLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2NetworkSummaryLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv2NetworkSummaryLsa) FromPbText(value string) error {
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

func (m *marshalospfv2NetworkSummaryLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv2NetworkSummaryLsa) FromYaml(value string) error {
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

func (m *marshalospfv2NetworkSummaryLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv2NetworkSummaryLsa) FromJson(value string) error {
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

func (obj *ospfv2NetworkSummaryLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2NetworkSummaryLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2NetworkSummaryLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2NetworkSummaryLsa) Clone() (Ospfv2NetworkSummaryLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2NetworkSummaryLsa()
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

func (obj *ospfv2NetworkSummaryLsa) setNil() {
	obj.commonAttrsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2NetworkSummaryLsa is contents of the Network Summary LSA - Type 3.
type Ospfv2NetworkSummaryLsa interface {
	Validation
	// msg marshals Ospfv2NetworkSummaryLsa to protobuf object *otg.Ospfv2NetworkSummaryLsa
	// and doesn't set defaults
	msg() *otg.Ospfv2NetworkSummaryLsa
	// setMsg unmarshals Ospfv2NetworkSummaryLsa from protobuf object *otg.Ospfv2NetworkSummaryLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv2NetworkSummaryLsa) Ospfv2NetworkSummaryLsa
	// provides marshal interface
	Marshal() marshalOspfv2NetworkSummaryLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2NetworkSummaryLsa
	// validate validates Ospfv2NetworkSummaryLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2NetworkSummaryLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonAttrs returns Ospfv2CommonAttrs, set in Ospfv2NetworkSummaryLsa.
	// Ospfv2CommonAttrs is attributes in LSA Header.
	CommonAttrs() Ospfv2CommonAttrs
	// SetCommonAttrs assigns Ospfv2CommonAttrs provided by user to Ospfv2NetworkSummaryLsa.
	// Ospfv2CommonAttrs is attributes in LSA Header.
	SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2NetworkSummaryLsa
	// HasCommonAttrs checks if CommonAttrs has been set in Ospfv2NetworkSummaryLsa
	HasCommonAttrs() bool
	// NetworkMask returns string, set in Ospfv2NetworkSummaryLsa.
	NetworkMask() string
	// SetNetworkMask assigns string provided by user to Ospfv2NetworkSummaryLsa
	SetNetworkMask(value string) Ospfv2NetworkSummaryLsa
	// HasNetworkMask checks if NetworkMask has been set in Ospfv2NetworkSummaryLsa
	HasNetworkMask() bool
	// Metric returns uint32, set in Ospfv2NetworkSummaryLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv2NetworkSummaryLsa
	SetMetric(value uint32) Ospfv2NetworkSummaryLsa
	// HasMetric checks if Metric has been set in Ospfv2NetworkSummaryLsa
	HasMetric() bool
	setNil()
}

// Common LSA attributes.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2NetworkSummaryLsa) CommonAttrs() Ospfv2CommonAttrs {
	if obj.obj.CommonAttrs == nil {
		obj.obj.CommonAttrs = NewOspfv2CommonAttrs().msg()
	}
	if obj.commonAttrsHolder == nil {
		obj.commonAttrsHolder = &ospfv2CommonAttrs{obj: obj.obj.CommonAttrs}
	}
	return obj.commonAttrsHolder
}

// Common LSA attributes.
// CommonAttrs returns a Ospfv2CommonAttrs
func (obj *ospfv2NetworkSummaryLsa) HasCommonAttrs() bool {
	return obj.obj.CommonAttrs != nil
}

// Common LSA attributes.
// SetCommonAttrs sets the Ospfv2CommonAttrs value in the Ospfv2NetworkSummaryLsa object
func (obj *ospfv2NetworkSummaryLsa) SetCommonAttrs(value Ospfv2CommonAttrs) Ospfv2NetworkSummaryLsa {

	obj.commonAttrsHolder = nil
	obj.obj.CommonAttrs = value.msg()

	return obj
}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2NetworkSummaryLsa) NetworkMask() string {

	return *obj.obj.NetworkMask

}

// The IPv4 address mask for the network.
// NetworkMask returns a string
func (obj *ospfv2NetworkSummaryLsa) HasNetworkMask() bool {
	return obj.obj.NetworkMask != nil
}

// The IPv4 address mask for the network.
// SetNetworkMask sets the string value in the Ospfv2NetworkSummaryLsa object
func (obj *ospfv2NetworkSummaryLsa) SetNetworkMask(value string) Ospfv2NetworkSummaryLsa {

	obj.obj.NetworkMask = &value
	return obj
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2NetworkSummaryLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv2NetworkSummaryLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// SetMetric sets the uint32 value in the Ospfv2NetworkSummaryLsa object
func (obj *ospfv2NetworkSummaryLsa) SetMetric(value uint32) Ospfv2NetworkSummaryLsa {

	obj.obj.Metric = &value
	return obj
}

func (obj *ospfv2NetworkSummaryLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommonAttrs != nil {

		obj.CommonAttrs().validateObj(vObj, set_default)
	}

	if obj.obj.NetworkMask != nil {

		err := obj.validateIpv4(obj.NetworkMask())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2NetworkSummaryLsa.NetworkMask"))
		}

	}

}

func (obj *ospfv2NetworkSummaryLsa) setDefault() {

}
