package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterAreaRouterLsa *****
type ospfv3InterAreaRouterLsa struct {
	validation
	obj          *otg.Ospfv3InterAreaRouterLsa
	marshaller   marshalOspfv3InterAreaRouterLsa
	unMarshaller unMarshalOspfv3InterAreaRouterLsa
	headerHolder Ospfv3LsaHeader
}

func NewOspfv3InterAreaRouterLsa() Ospfv3InterAreaRouterLsa {
	obj := ospfv3InterAreaRouterLsa{obj: &otg.Ospfv3InterAreaRouterLsa{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterAreaRouterLsa) msg() *otg.Ospfv3InterAreaRouterLsa {
	return obj.obj
}

func (obj *ospfv3InterAreaRouterLsa) setMsg(msg *otg.Ospfv3InterAreaRouterLsa) Ospfv3InterAreaRouterLsa {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterAreaRouterLsa struct {
	obj *ospfv3InterAreaRouterLsa
}

type marshalOspfv3InterAreaRouterLsa interface {
	// ToProto marshals Ospfv3InterAreaRouterLsa to protobuf object *otg.Ospfv3InterAreaRouterLsa
	ToProto() (*otg.Ospfv3InterAreaRouterLsa, error)
	// ToPbText marshals Ospfv3InterAreaRouterLsa to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterAreaRouterLsa to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterAreaRouterLsa to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterAreaRouterLsa struct {
	obj *ospfv3InterAreaRouterLsa
}

type unMarshalOspfv3InterAreaRouterLsa interface {
	// FromProto unmarshals Ospfv3InterAreaRouterLsa from protobuf object *otg.Ospfv3InterAreaRouterLsa
	FromProto(msg *otg.Ospfv3InterAreaRouterLsa) (Ospfv3InterAreaRouterLsa, error)
	// FromPbText unmarshals Ospfv3InterAreaRouterLsa from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterAreaRouterLsa from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterAreaRouterLsa from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterAreaRouterLsa) Marshal() marshalOspfv3InterAreaRouterLsa {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterAreaRouterLsa{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterAreaRouterLsa) Unmarshal() unMarshalOspfv3InterAreaRouterLsa {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterAreaRouterLsa{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterAreaRouterLsa) ToProto() (*otg.Ospfv3InterAreaRouterLsa, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterAreaRouterLsa) FromProto(msg *otg.Ospfv3InterAreaRouterLsa) (Ospfv3InterAreaRouterLsa, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterAreaRouterLsa) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterAreaRouterLsa) FromPbText(value string) error {
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

func (m *marshalospfv3InterAreaRouterLsa) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterAreaRouterLsa) FromYaml(value string) error {
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

func (m *marshalospfv3InterAreaRouterLsa) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterAreaRouterLsa) FromJson(value string) error {
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

func (obj *ospfv3InterAreaRouterLsa) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterAreaRouterLsa) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterAreaRouterLsa) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterAreaRouterLsa) Clone() (Ospfv3InterAreaRouterLsa, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterAreaRouterLsa()
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

func (obj *ospfv3InterAreaRouterLsa) setNil() {
	obj.headerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3InterAreaRouterLsa is contents of OSPFv3 Inter-Area-Router LSA - Type 4.
type Ospfv3InterAreaRouterLsa interface {
	Validation
	// msg marshals Ospfv3InterAreaRouterLsa to protobuf object *otg.Ospfv3InterAreaRouterLsa
	// and doesn't set defaults
	msg() *otg.Ospfv3InterAreaRouterLsa
	// setMsg unmarshals Ospfv3InterAreaRouterLsa from protobuf object *otg.Ospfv3InterAreaRouterLsa
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterAreaRouterLsa) Ospfv3InterAreaRouterLsa
	// provides marshal interface
	Marshal() marshalOspfv3InterAreaRouterLsa
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterAreaRouterLsa
	// validate validates Ospfv3InterAreaRouterLsa
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterAreaRouterLsa, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Header returns Ospfv3LsaHeader, set in Ospfv3InterAreaRouterLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	Header() Ospfv3LsaHeader
	// SetHeader assigns Ospfv3LsaHeader provided by user to Ospfv3InterAreaRouterLsa.
	// Ospfv3LsaHeader is attributes in LSA Header.
	SetHeader(value Ospfv3LsaHeader) Ospfv3InterAreaRouterLsa
	// HasHeader checks if Header has been set in Ospfv3InterAreaRouterLsa
	HasHeader() bool
	// Metric returns uint32, set in Ospfv3InterAreaRouterLsa.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to Ospfv3InterAreaRouterLsa
	SetMetric(value uint32) Ospfv3InterAreaRouterLsa
	// HasMetric checks if Metric has been set in Ospfv3InterAreaRouterLsa
	HasMetric() bool
	// DestinationRouterId returns string, set in Ospfv3InterAreaRouterLsa.
	DestinationRouterId() string
	// SetDestinationRouterId assigns string provided by user to Ospfv3InterAreaRouterLsa
	SetDestinationRouterId(value string) Ospfv3InterAreaRouterLsa
	// HasDestinationRouterId checks if DestinationRouterId has been set in Ospfv3InterAreaRouterLsa
	HasDestinationRouterId() bool
	setNil()
}

// Contents of the LSA header.
// Header returns a Ospfv3LsaHeader
func (obj *ospfv3InterAreaRouterLsa) Header() Ospfv3LsaHeader {
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
func (obj *ospfv3InterAreaRouterLsa) HasHeader() bool {
	return obj.obj.Header != nil
}

// Contents of the LSA header.
// SetHeader sets the Ospfv3LsaHeader value in the Ospfv3InterAreaRouterLsa object
func (obj *ospfv3InterAreaRouterLsa) SetHeader(value Ospfv3LsaHeader) Ospfv3InterAreaRouterLsa {

	obj.headerHolder = nil
	obj.obj.Header = value.msg()

	return obj
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv3InterAreaRouterLsa) Metric() uint32 {

	return *obj.obj.Metric

}

// The cost of the summary route TOS level 0 and all unspecified levels.
// Metric returns a uint32
func (obj *ospfv3InterAreaRouterLsa) HasMetric() bool {
	return obj.obj.Metric != nil
}

// The cost of the summary route TOS level 0 and all unspecified levels.
// SetMetric sets the uint32 value in the Ospfv3InterAreaRouterLsa object
func (obj *ospfv3InterAreaRouterLsa) SetMetric(value uint32) Ospfv3InterAreaRouterLsa {

	obj.obj.Metric = &value
	return obj
}

// The id of the destination router of LSA.
// DestinationRouterId returns a string
func (obj *ospfv3InterAreaRouterLsa) DestinationRouterId() string {

	return *obj.obj.DestinationRouterId

}

// The id of the destination router of LSA.
// DestinationRouterId returns a string
func (obj *ospfv3InterAreaRouterLsa) HasDestinationRouterId() bool {
	return obj.obj.DestinationRouterId != nil
}

// The id of the destination router of LSA.
// SetDestinationRouterId sets the string value in the Ospfv3InterAreaRouterLsa object
func (obj *ospfv3InterAreaRouterLsa) SetDestinationRouterId(value string) Ospfv3InterAreaRouterLsa {

	obj.obj.DestinationRouterId = &value
	return obj
}

func (obj *ospfv3InterAreaRouterLsa) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Header != nil {

		obj.Header().validateObj(vObj, set_default)
	}

	if obj.obj.DestinationRouterId != nil {

		err := obj.validateIpv4(obj.DestinationRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3InterAreaRouterLsa.DestinationRouterId"))
		}

	}

}

func (obj *ospfv3InterAreaRouterLsa) setDefault() {

}
