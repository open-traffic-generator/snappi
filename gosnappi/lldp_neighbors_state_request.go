package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpNeighborsStateRequest *****
type lldpNeighborsStateRequest struct {
	validation
	obj          *otg.LldpNeighborsStateRequest
	marshaller   marshalLldpNeighborsStateRequest
	unMarshaller unMarshalLldpNeighborsStateRequest
}

func NewLldpNeighborsStateRequest() LldpNeighborsStateRequest {
	obj := lldpNeighborsStateRequest{obj: &otg.LldpNeighborsStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpNeighborsStateRequest) msg() *otg.LldpNeighborsStateRequest {
	return obj.obj
}

func (obj *lldpNeighborsStateRequest) setMsg(msg *otg.LldpNeighborsStateRequest) LldpNeighborsStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpNeighborsStateRequest struct {
	obj *lldpNeighborsStateRequest
}

type marshalLldpNeighborsStateRequest interface {
	// ToProto marshals LldpNeighborsStateRequest to protobuf object *otg.LldpNeighborsStateRequest
	ToProto() (*otg.LldpNeighborsStateRequest, error)
	// ToPbText marshals LldpNeighborsStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpNeighborsStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpNeighborsStateRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpNeighborsStateRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpNeighborsStateRequest struct {
	obj *lldpNeighborsStateRequest
}

type unMarshalLldpNeighborsStateRequest interface {
	// FromProto unmarshals LldpNeighborsStateRequest from protobuf object *otg.LldpNeighborsStateRequest
	FromProto(msg *otg.LldpNeighborsStateRequest) (LldpNeighborsStateRequest, error)
	// FromPbText unmarshals LldpNeighborsStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpNeighborsStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpNeighborsStateRequest from JSON text
	FromJson(value string) error
}

func (obj *lldpNeighborsStateRequest) Marshal() marshalLldpNeighborsStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpNeighborsStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpNeighborsStateRequest) Unmarshal() unMarshalLldpNeighborsStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpNeighborsStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpNeighborsStateRequest) ToProto() (*otg.LldpNeighborsStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpNeighborsStateRequest) FromProto(msg *otg.LldpNeighborsStateRequest) (LldpNeighborsStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpNeighborsStateRequest) ToPbText() (string, error) {
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

func (m *unMarshallldpNeighborsStateRequest) FromPbText(value string) error {
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

func (m *marshallldpNeighborsStateRequest) ToYaml() (string, error) {
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

func (m *unMarshallldpNeighborsStateRequest) FromYaml(value string) error {
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

func (m *marshallldpNeighborsStateRequest) ToJsonRaw() (string, error) {
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

func (m *marshallldpNeighborsStateRequest) ToJson() (string, error) {
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

func (m *unMarshallldpNeighborsStateRequest) FromJson(value string) error {
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

func (obj *lldpNeighborsStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpNeighborsStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpNeighborsStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpNeighborsStateRequest) Clone() (LldpNeighborsStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpNeighborsStateRequest()
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

// LldpNeighborsStateRequest is the request to retrieve LLDP neighbor information for a given instance.
type LldpNeighborsStateRequest interface {
	Validation
	// msg marshals LldpNeighborsStateRequest to protobuf object *otg.LldpNeighborsStateRequest
	// and doesn't set defaults
	msg() *otg.LldpNeighborsStateRequest
	// setMsg unmarshals LldpNeighborsStateRequest from protobuf object *otg.LldpNeighborsStateRequest
	// and doesn't set defaults
	setMsg(*otg.LldpNeighborsStateRequest) LldpNeighborsStateRequest
	// provides marshal interface
	Marshal() marshalLldpNeighborsStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalLldpNeighborsStateRequest
	// validate validates LldpNeighborsStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpNeighborsStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LldpNames returns []string, set in LldpNeighborsStateRequest.
	LldpNames() []string
	// SetLldpNames assigns []string provided by user to LldpNeighborsStateRequest
	SetLldpNames(value []string) LldpNeighborsStateRequest
	// NeighborIdFilters returns []string, set in LldpNeighborsStateRequest.
	NeighborIdFilters() []string
	// SetNeighborIdFilters assigns []string provided by user to LldpNeighborsStateRequest
	SetNeighborIdFilters(value []string) LldpNeighborsStateRequest
}

// The names of LLDP instances for which neighbor information will be retrieved. If no names are specified then the results will contain neighbor information for all configured LLDP instances.
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// LldpNames returns a []string
func (obj *lldpNeighborsStateRequest) LldpNames() []string {
	if obj.obj.LldpNames == nil {
		obj.obj.LldpNames = make([]string, 0)
	}
	return obj.obj.LldpNames
}

// The names of LLDP instances for which neighbor information will be retrieved. If no names are specified then the results will contain neighbor information for all configured LLDP instances.
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// x-constraint:
// - /components/schemas/Lldp/properties/name
//
// SetLldpNames sets the []string value in the LldpNeighborsStateRequest object
func (obj *lldpNeighborsStateRequest) SetLldpNames(value []string) LldpNeighborsStateRequest {

	if obj.obj.LldpNames == nil {
		obj.obj.LldpNames = make([]string, 0)
	}
	obj.obj.LldpNames = value

	return obj
}

// Specify the neighbors for which information will be returned. If empty  or missing then information for all neighbors will be returned.
// NeighborIdFilters returns a []string
func (obj *lldpNeighborsStateRequest) NeighborIdFilters() []string {
	if obj.obj.NeighborIdFilters == nil {
		obj.obj.NeighborIdFilters = make([]string, 0)
	}
	return obj.obj.NeighborIdFilters
}

// Specify the neighbors for which information will be returned. If empty  or missing then information for all neighbors will be returned.
// SetNeighborIdFilters sets the []string value in the LldpNeighborsStateRequest object
func (obj *lldpNeighborsStateRequest) SetNeighborIdFilters(value []string) LldpNeighborsStateRequest {

	if obj.obj.NeighborIdFilters == nil {
		obj.obj.NeighborIdFilters = make([]string, 0)
	}
	obj.obj.NeighborIdFilters = value

	return obj
}

func (obj *lldpNeighborsStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpNeighborsStateRequest) setDefault() {

}
