package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterAreaPrefixLsa *****
type ospfv3InterAreaPrefixLsa struct {
	validation
	obj          *otg.Ospfv3InterAreaPrefixLsa
	marshaller   marshalOspfv3InterAreaPrefixLsa
	unMarshaller unMarshalOspfv3InterAreaPrefixLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3InterAreaPrefixLsa() Ospfv3InterAreaPrefixLsa {
	obj := ospfv3InterAreaPrefixLsa{obj: &otg.Ospfv3InterAreaPrefixLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterAreaPrefixLsa) msg() *otg.Ospfv3InterAreaPrefixLsa {
	return obj.obj
}

func (obj *ospfv3InterAreaPrefixLsa) setMsg(msg *otg.Ospfv3InterAreaPrefixLsa) Ospfv3InterAreaPrefixLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterAreaPrefixLsa struct {
	obj *ospfv3InterAreaPrefixLsa
}

type marshalOspfv3InterAreaPrefixLsa interface {
	// ToProto marshals Ospfv3InterAreaPrefixLsa to protobuf object *otg.Ospfv3InterAreaPrefixLsa
	ToProto() (*otg.Ospfv3InterAreaPrefixLsa, error)
	// ToPbText marshals Ospfv3InterAreaPrefixLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterAreaPrefixLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterAreaPrefixLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterAreaPrefixLsa struct {
	obj *ospfv3InterAreaPrefixLsa
}

type unMarshalOspfv3InterAreaPrefixLsa interface {
	// FromProto unmarshals Ospfv3InterAreaPrefixLsa from protobuf object *otg.Ospfv3InterAreaPrefixLsa
	FromProto(msg *otg.Ospfv3InterAreaPrefixLsa) (Ospfv3InterAreaPrefixLsa, error)
	// FromPbText unmarshals Ospfv3InterAreaPrefixLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterAreaPrefixLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterAreaPrefixLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterAreaPrefixLsa) Marshal() marshalOspfv3InterAreaPrefixLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterAreaPrefixLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterAreaPrefixLsa) Unmarshal() unMarshalOspfv3InterAreaPrefixLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterAreaPrefixLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterAreaPrefixLsa) ToProto() (*otg.Ospfv3InterAreaPrefixLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterAreaPrefixLsa) FromProto(msg *otg.Ospfv3InterAreaPrefixLsa) (Ospfv3InterAreaPrefixLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterAreaPrefixLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterAreaPrefixLsa) FromPbText(value string) error {
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

func (m *marshalospfv3InterAreaPrefixLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterAreaPrefixLsa) FromYaml(value string) error {
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

func (m *marshalospfv3InterAreaPrefixLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterAreaPrefixLsa) FromJson(value string) error {
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

func (obj *ospfv3InterAreaPrefixLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterAreaPrefixLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterAreaPrefixLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterAreaPrefixLsa) Clone() (Ospfv3InterAreaPrefixLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterAreaPrefixLsa()
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

func (obj *ospfv3InterAreaPrefixLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3InterAreaPrefixLsa is contents of the Inter-Area-Prefix LSA - Type 3.
type Ospfv3InterAreaPrefixLsa interface {
	Validation
	// msg marshals Ospfv3InterAreaPrefixLsa to protobuf object *otg.Ospfv3InterAreaPrefixLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3InterAreaPrefixLsa
	// setMsg unmarshals Ospfv3InterAreaPrefixLsa from protobuf object *otg.Ospfv3InterAreaPrefixLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterAreaPrefixLsa) Ospfv3InterAreaPrefixLsa
	// provides marshal interface
	Marshal() marshalOspfv3InterAreaPrefixLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterAreaPrefixLsa
	// validate validates Ospfv3InterAreaPrefixLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterAreaPrefixLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3InterAreaPrefixLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3InterAreaPrefixLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3InterAreaPrefixLsa
	// HasHeader checks if Header has been set in Ospfv3InterAreaPrefixLsa
	HasHeader() bool
	// AddressPrefix returns string, set in Ospfv3InterAreaPrefixLsa.
	AddressPrefix() string
	// SetAddressPrefix assigns string provided by user to Ospfv3InterAreaPrefixLsa
	SetAddressPrefix(value string) Ospfv3InterAreaPrefixLsa
	// HasAddressPrefix checks if AddressPrefix has been set in Ospfv3InterAreaPrefixLsa
	HasAddressPrefix() bool
	// PrefixLength returns uint32, set in Ospfv3InterAreaPrefixLsa.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to Ospfv3InterAreaPrefixLsa
	SetPrefixLength(value uint32) Ospfv3InterAreaPrefixLsa
	// HasPrefixLength checks if PrefixLength has been set in Ospfv3InterAreaPrefixLsa
	HasPrefixLength() bool
	// Metric returns uint32, set in Ospfv3InterAreaPrefixLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3InterAreaPrefixLsa
	SetMetric(value uint32) Ospfv3InterAreaPrefixLsa
	// HasMetric checks if Metric has been set in Ospfv3InterAreaPrefixLsa
	HasMetric() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3InterAreaPrefixLsa) Header() Ospfv3LsaHeader {
	if obj.obj.Header == nil {
		obj.obj.Header = NewOspfv3LsaHeader().msg()
	}
	if obj.headerHolder == nil {
		obj.headerHolder = &ospfv3LsaHeader{obj: obj.obj.Header}
	}
	return obj.headerHolder
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3InterAreaPrefixLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3InterAreaPrefixLsa object
func (obj *ospfv3InterAreaPrefixLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3InterAreaPrefixLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The prefix for Inter Area Prefix LSA Address.
// AddressPrefix returns a string
func (obj *ospfv3InterAreaPrefixLsa) AddressPrefix() string {

	return *obj.obj.AddressPrefix

}

// The prefix for Inter Area Prefix LSA Address.
// AddressPrefix returns a string
func (obj *ospfv3InterAreaPrefixLsa) HasAddressPrefix() bool {
	return obj.obj.AddressPrefix != nil
}

// The prefix for Inter Area Prefix LSA Address.
// SetAddressPrefix sets the string value in the Ospfv3InterAreaPrefixLsa object
func (obj *ospfv3InterAreaPrefixLsa) SetAddressPrefix(value string) Ospfv3InterAreaPrefixLsa {

	obj.obj.AddressPrefix = &value
	return obj
}

// The prefix length for the IP address.
// PrefixLength returns a uint32
func (obj *ospfv3InterAreaPrefixLsa) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length for the IP address.
// PrefixLength returns a uint32
func (obj *ospfv3InterAreaPrefixLsa) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length for the IP address.
// SetPrefixLength sets the uint32 value in the Ospfv3InterAreaPrefixLsa object
func (obj *ospfv3InterAreaPrefixLsa) SetPrefixLength(value uint32) Ospfv3InterAreaPrefixLsa {

	obj.obj.PrefixLength = &value
	return obj
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv3InterAreaPrefixLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv3InterAreaPrefixLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// SetMetric sets the uint32 value in the Ospfv3InterAreaPrefixLsa object
func (obj *ospfv3InterAreaPrefixLsa) SetMetric(value uint32) Ospfv3InterAreaPrefixLsa {

	obj.obj.Metric = &value
	return obj
}

func (obj *ospfv3InterAreaPrefixLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.AddressPrefix != nil {

		err := obj.validateIpv6(obj.AddressPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3InterAreaPrefixLsa.AddressPrefix"))
		}

	}

}

func (obj *ospfv3InterAreaPrefixLsa) setDefault() {

}
