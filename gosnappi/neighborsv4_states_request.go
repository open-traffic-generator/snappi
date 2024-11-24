package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Neighborsv4StatesRequest *****
type neighborsv4StatesRequest struct {
	validation
	obj          *otg.Neighborsv4StatesRequest
	marshaller   marshalNeighborsv4StatesRequest
	unMarshaller unMarshalNeighborsv4StatesRequest
}

func NewNeighborsv4StatesRequest() Neighborsv4StatesRequest {
	obj := neighborsv4StatesRequest{obj: &otg.Neighborsv4StatesRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *neighborsv4StatesRequest) msg() *otg.Neighborsv4StatesRequest {
	return obj.obj
}

func (obj *neighborsv4StatesRequest) setMsg(msg *otg.Neighborsv4StatesRequest) Neighborsv4StatesRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalneighborsv4StatesRequest struct {
	obj *neighborsv4StatesRequest
}

type marshalNeighborsv4StatesRequest interface {
	// ToProto marshals Neighborsv4StatesRequest to protobuf object *otg.Neighborsv4StatesRequest
	ToProto() (*otg.Neighborsv4StatesRequest, error)
	// ToPbText marshals Neighborsv4StatesRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Neighborsv4StatesRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Neighborsv4StatesRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Neighborsv4StatesRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalneighborsv4StatesRequest struct {
	obj *neighborsv4StatesRequest
}

type unMarshalNeighborsv4StatesRequest interface {
	// FromProto unmarshals Neighborsv4StatesRequest from protobuf object *otg.Neighborsv4StatesRequest
	FromProto(msg *otg.Neighborsv4StatesRequest) (Neighborsv4StatesRequest, error)
	// FromPbText unmarshals Neighborsv4StatesRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Neighborsv4StatesRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Neighborsv4StatesRequest from JSON text
	FromJson(value string) error
}

func (obj *neighborsv4StatesRequest) Marshal() marshalNeighborsv4StatesRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalneighborsv4StatesRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *neighborsv4StatesRequest) Unmarshal() unMarshalNeighborsv4StatesRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalneighborsv4StatesRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalneighborsv4StatesRequest) ToProto() (*otg.Neighborsv4StatesRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalneighborsv4StatesRequest) FromProto(msg *otg.Neighborsv4StatesRequest) (Neighborsv4StatesRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalneighborsv4StatesRequest) ToPbText() (string, error) {
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

func (m *unMarshalneighborsv4StatesRequest) FromPbText(value string) error {
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

func (m *marshalneighborsv4StatesRequest) ToYaml() (string, error) {
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

func (m *unMarshalneighborsv4StatesRequest) FromYaml(value string) error {
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

func (m *marshalneighborsv4StatesRequest) ToJsonRaw() (string, error) {
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

func (m *marshalneighborsv4StatesRequest) ToJson() (string, error) {
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

func (m *unMarshalneighborsv4StatesRequest) FromJson(value string) error {
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

func (obj *neighborsv4StatesRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *neighborsv4StatesRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *neighborsv4StatesRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *neighborsv4StatesRequest) Clone() (Neighborsv4StatesRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNeighborsv4StatesRequest()
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

// Neighborsv4StatesRequest is the request to retrieve IPv4 Neighbor state (ARP cache entries) of a network interface(s).
type Neighborsv4StatesRequest interface {
	Validation
	// msg marshals Neighborsv4StatesRequest to protobuf object *otg.Neighborsv4StatesRequest
	// and doesn't set defaults
	msg() *otg.Neighborsv4StatesRequest
	// setMsg unmarshals Neighborsv4StatesRequest from protobuf object *otg.Neighborsv4StatesRequest
	// and doesn't set defaults
	setMsg(*otg.Neighborsv4StatesRequest) Neighborsv4StatesRequest
	// provides marshal interface
	Marshal() marshalNeighborsv4StatesRequest
	// provides unmarshal interface
	Unmarshal() unMarshalNeighborsv4StatesRequest
	// validate validates Neighborsv4StatesRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Neighborsv4StatesRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthernetNames returns []string, set in Neighborsv4StatesRequest.
	EthernetNames() []string
	// SetEthernetNames assigns []string provided by user to Neighborsv4StatesRequest
	SetEthernetNames(value []string) Neighborsv4StatesRequest
}

// The names of Ethernet interfaces for which Neighbor state (ARP cache entries) will be retrieved. If no names are specified then the results will contain Neighbor state (ARP cache entries) for all available Ethernet interfaces.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthernetNames returns a []string
func (obj *neighborsv4StatesRequest) EthernetNames() []string {
	if obj.obj.EthernetNames == nil {
		obj.obj.EthernetNames = make([]string, 0)
	}
	return obj.obj.EthernetNames
}

// The names of Ethernet interfaces for which Neighbor state (ARP cache entries) will be retrieved. If no names are specified then the results will contain Neighbor state (ARP cache entries) for all available Ethernet interfaces.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthernetNames sets the []string value in the Neighborsv4StatesRequest object
func (obj *neighborsv4StatesRequest) SetEthernetNames(value []string) Neighborsv4StatesRequest {

	if obj.obj.EthernetNames == nil {
		obj.obj.EthernetNames = make([]string, 0)
	}
	obj.obj.EthernetNames = value

	return obj
}

func (obj *neighborsv4StatesRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *neighborsv4StatesRequest) setDefault() {

}
