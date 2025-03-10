package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Neighborsv6StatesRequest *****
type neighborsv6StatesRequest struct {
	validation
	obj          *otg.Neighborsv6StatesRequest
	marshaller   marshalNeighborsv6StatesRequest
	unMarshaller unMarshalNeighborsv6StatesRequest
}

func NewNeighborsv6StatesRequest() Neighborsv6StatesRequest {
	obj := neighborsv6StatesRequest{obj: &otg.Neighborsv6StatesRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *neighborsv6StatesRequest) msg() *otg.Neighborsv6StatesRequest {
	return obj.obj
}

func (obj *neighborsv6StatesRequest) setMsg(msg *otg.Neighborsv6StatesRequest) Neighborsv6StatesRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalneighborsv6StatesRequest struct {
	obj *neighborsv6StatesRequest
}

type marshalNeighborsv6StatesRequest interface {
	// ToProto marshals Neighborsv6StatesRequest to protobuf object *otg.Neighborsv6StatesRequest
	ToProto() (*otg.Neighborsv6StatesRequest, error)
	// ToPbText marshals Neighborsv6StatesRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Neighborsv6StatesRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Neighborsv6StatesRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Neighborsv6StatesRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalneighborsv6StatesRequest struct {
	obj *neighborsv6StatesRequest
}

type unMarshalNeighborsv6StatesRequest interface {
	// FromProto unmarshals Neighborsv6StatesRequest from protobuf object *otg.Neighborsv6StatesRequest
	FromProto(msg *otg.Neighborsv6StatesRequest) (Neighborsv6StatesRequest, error)
	// FromPbText unmarshals Neighborsv6StatesRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Neighborsv6StatesRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Neighborsv6StatesRequest from JSON text
	FromJson(value string) error
}

func (obj *neighborsv6StatesRequest) Marshal() marshalNeighborsv6StatesRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalneighborsv6StatesRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *neighborsv6StatesRequest) Unmarshal() unMarshalNeighborsv6StatesRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalneighborsv6StatesRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalneighborsv6StatesRequest) ToProto() (*otg.Neighborsv6StatesRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalneighborsv6StatesRequest) FromProto(msg *otg.Neighborsv6StatesRequest) (Neighborsv6StatesRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalneighborsv6StatesRequest) ToPbText() (string, error) {
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

func (m *unMarshalneighborsv6StatesRequest) FromPbText(value string) error {
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

func (m *marshalneighborsv6StatesRequest) ToYaml() (string, error) {
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

func (m *unMarshalneighborsv6StatesRequest) FromYaml(value string) error {
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

func (m *marshalneighborsv6StatesRequest) ToJsonRaw() (string, error) {
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

func (m *marshalneighborsv6StatesRequest) ToJson() (string, error) {
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

func (m *unMarshalneighborsv6StatesRequest) FromJson(value string) error {
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

func (obj *neighborsv6StatesRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *neighborsv6StatesRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *neighborsv6StatesRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *neighborsv6StatesRequest) Clone() (Neighborsv6StatesRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNeighborsv6StatesRequest()
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

// Neighborsv6StatesRequest is the request to retrieve IPv6 Neighbor state (NDISC cache entries) of a network interface(s).
type Neighborsv6StatesRequest interface {
	Validation
	// msg marshals Neighborsv6StatesRequest to protobuf object *otg.Neighborsv6StatesRequest
	// and doesn't set defaults
	msg() *otg.Neighborsv6StatesRequest
	// setMsg unmarshals Neighborsv6StatesRequest from protobuf object *otg.Neighborsv6StatesRequest
	// and doesn't set defaults
	setMsg(*otg.Neighborsv6StatesRequest) Neighborsv6StatesRequest
	// provides marshal interface
	Marshal() marshalNeighborsv6StatesRequest
	// provides unmarshal interface
	Unmarshal() unMarshalNeighborsv6StatesRequest
	// validate validates Neighborsv6StatesRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Neighborsv6StatesRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetNames returns []string, set in Neighborsv6StatesRequest.
	EthernetNames() []string
	// SetEthernetNames assigns []string provided by user to Neighborsv6StatesRequest
	SetEthernetNames(value []string) Neighborsv6StatesRequest
}

// The names of Ethernet interfaces for which Neighbor state (NDISC cache entries) will be retrieved. If no names are specified then the results will contain Neighbor state (NDISC cache entries) for all available Ethernet interfaces.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthernetNames returns a []string
func (obj *neighborsv6StatesRequest) EthernetNames() []string {
	if obj.obj.EthernetNames == nil {
		obj.obj.EthernetNames = make([]string, 0)
	}
	return obj.obj.EthernetNames
}

// The names of Ethernet interfaces for which Neighbor state (NDISC cache entries) will be retrieved. If no names are specified then the results will contain Neighbor state (NDISC cache entries) for all available Ethernet interfaces.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthernetNames sets the []string value in the Neighborsv6StatesRequest object
func (obj *neighborsv6StatesRequest) SetEthernetNames(value []string) Neighborsv6StatesRequest {

	if obj.obj.EthernetNames == nil {
		obj.obj.EthernetNames = make([]string, 0)
	}
	obj.obj.EthernetNames = value

	return obj
}

func (obj *neighborsv6StatesRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *neighborsv6StatesRequest) setDefault() {

}
