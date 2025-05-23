package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2IPv4MetricsRequest *****
type rocev2IPv4MetricsRequest struct {
	validation
	obj           *otg.Rocev2IPv4MetricsRequest
	marshaller    marshalRocev2IPv4MetricsRequest
	unMarshaller  unMarshalRocev2IPv4MetricsRequest
	perPeerHolder Rocev2IPv4ColumnNames
}

func NewRocev2IPv4MetricsRequest() Rocev2IPv4MetricsRequest {
	obj := rocev2IPv4MetricsRequest{obj: &otg.Rocev2IPv4MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2IPv4MetricsRequest) msg() *otg.Rocev2IPv4MetricsRequest {
	return obj.obj
}

func (obj *rocev2IPv4MetricsRequest) setMsg(msg *otg.Rocev2IPv4MetricsRequest) Rocev2IPv4MetricsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2IPv4MetricsRequest struct {
	obj *rocev2IPv4MetricsRequest
}

type marshalRocev2IPv4MetricsRequest interface {
	// ToProto marshals Rocev2IPv4MetricsRequest to protobuf object *otg.Rocev2IPv4MetricsRequest
	ToProto() (*otg.Rocev2IPv4MetricsRequest, error)
	// ToPbText marshals Rocev2IPv4MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2IPv4MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2IPv4MetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2IPv4MetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2IPv4MetricsRequest struct {
	obj *rocev2IPv4MetricsRequest
}

type unMarshalRocev2IPv4MetricsRequest interface {
	// FromProto unmarshals Rocev2IPv4MetricsRequest from protobuf object *otg.Rocev2IPv4MetricsRequest
	FromProto(msg *otg.Rocev2IPv4MetricsRequest) (Rocev2IPv4MetricsRequest, error)
	// FromPbText unmarshals Rocev2IPv4MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2IPv4MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2IPv4MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *rocev2IPv4MetricsRequest) Marshal() marshalRocev2IPv4MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2IPv4MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2IPv4MetricsRequest) Unmarshal() unMarshalRocev2IPv4MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2IPv4MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2IPv4MetricsRequest) ToProto() (*otg.Rocev2IPv4MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2IPv4MetricsRequest) FromProto(msg *otg.Rocev2IPv4MetricsRequest) (Rocev2IPv4MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2IPv4MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalrocev2IPv4MetricsRequest) FromPbText(value string) error {
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

func (m *marshalrocev2IPv4MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalrocev2IPv4MetricsRequest) FromYaml(value string) error {
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

func (m *marshalrocev2IPv4MetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2IPv4MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalrocev2IPv4MetricsRequest) FromJson(value string) error {
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

func (obj *rocev2IPv4MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2IPv4MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2IPv4MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2IPv4MetricsRequest) Clone() (Rocev2IPv4MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2IPv4MetricsRequest()
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

func (obj *rocev2IPv4MetricsRequest) setNil() {
	obj.perPeerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2IPv4MetricsRequest is request to retrieve RoCEv2 over IPv4 per peer metrics/statistics.
type Rocev2IPv4MetricsRequest interface {
	Validation
	// msg marshals Rocev2IPv4MetricsRequest to protobuf object *otg.Rocev2IPv4MetricsRequest
	// and doesn't set defaults
	msg() *otg.Rocev2IPv4MetricsRequest
	// setMsg unmarshals Rocev2IPv4MetricsRequest from protobuf object *otg.Rocev2IPv4MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Rocev2IPv4MetricsRequest) Rocev2IPv4MetricsRequest
	// provides marshal interface
	Marshal() marshalRocev2IPv4MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2IPv4MetricsRequest
	// validate validates Rocev2IPv4MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2IPv4MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2IPv4MetricsRequestChoiceEnum, set in Rocev2IPv4MetricsRequest
	Choice() Rocev2IPv4MetricsRequestChoiceEnum
	// setChoice assigns Rocev2IPv4MetricsRequestChoiceEnum provided by user to Rocev2IPv4MetricsRequest
	setChoice(value Rocev2IPv4MetricsRequestChoiceEnum) Rocev2IPv4MetricsRequest
	// HasChoice checks if Choice has been set in Rocev2IPv4MetricsRequest
	HasChoice() bool
	// PerPeer returns Rocev2IPv4ColumnNames, set in Rocev2IPv4MetricsRequest.
	// Rocev2IPv4ColumnNames is the names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
	PerPeer() Rocev2IPv4ColumnNames
	// SetPerPeer assigns Rocev2IPv4ColumnNames provided by user to Rocev2IPv4MetricsRequest.
	// Rocev2IPv4ColumnNames is the names of RoCEv2 over IPv4 peers to return results for. An empty list will return results for all RoCEv2 peers.
	SetPerPeer(value Rocev2IPv4ColumnNames) Rocev2IPv4MetricsRequest
	// HasPerPeer checks if PerPeer has been set in Rocev2IPv4MetricsRequest
	HasPerPeer() bool
	setNil()
}

type Rocev2IPv4MetricsRequestChoiceEnum string

// Enum of Choice on Rocev2IPv4MetricsRequest
var Rocev2IPv4MetricsRequestChoice = struct {
	PER_PEER Rocev2IPv4MetricsRequestChoiceEnum
}{
	PER_PEER: Rocev2IPv4MetricsRequestChoiceEnum("per_peer"),
}

func (obj *rocev2IPv4MetricsRequest) Choice() Rocev2IPv4MetricsRequestChoiceEnum {
	return Rocev2IPv4MetricsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// Fetch stats per_peer
// Choice returns a string
func (obj *rocev2IPv4MetricsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2IPv4MetricsRequest) setChoice(value Rocev2IPv4MetricsRequestChoiceEnum) Rocev2IPv4MetricsRequest {
	intValue, ok := otg.Rocev2IPv4MetricsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2IPv4MetricsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2IPv4MetricsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PerPeer = nil
	obj.perPeerHolder = nil

	if value == Rocev2IPv4MetricsRequestChoice.PER_PEER {
		obj.obj.PerPeer = NewRocev2IPv4ColumnNames().msg()
	}

	return obj
}

// description is TBD
// PerPeer returns a Rocev2IPv4ColumnNames
func (obj *rocev2IPv4MetricsRequest) PerPeer() Rocev2IPv4ColumnNames {
	if obj.obj.PerPeer == nil {
		obj.setChoice(Rocev2IPv4MetricsRequestChoice.PER_PEER)
	}
	if obj.perPeerHolder == nil {
		obj.perPeerHolder = &rocev2IPv4ColumnNames{obj: obj.obj.PerPeer}
	}
	return obj.perPeerHolder
}

// description is TBD
// PerPeer returns a Rocev2IPv4ColumnNames
func (obj *rocev2IPv4MetricsRequest) HasPerPeer() bool {
	return obj.obj.PerPeer != nil
}

// description is TBD
// SetPerPeer sets the Rocev2IPv4ColumnNames value in the Rocev2IPv4MetricsRequest object
func (obj *rocev2IPv4MetricsRequest) SetPerPeer(value Rocev2IPv4ColumnNames) Rocev2IPv4MetricsRequest {
	obj.setChoice(Rocev2IPv4MetricsRequestChoice.PER_PEER)
	obj.perPeerHolder = nil
	obj.obj.PerPeer = value.msg()

	return obj
}

func (obj *rocev2IPv4MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PerPeer != nil {

		obj.PerPeer().validateObj(vObj, set_default)
	}

}

func (obj *rocev2IPv4MetricsRequest) setDefault() {
	var choices_set int = 0
	var choice Rocev2IPv4MetricsRequestChoiceEnum

	if obj.obj.PerPeer != nil {
		choices_set += 1
		choice = Rocev2IPv4MetricsRequestChoice.PER_PEER
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2IPv4MetricsRequestChoice.PER_PEER)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2IPv4MetricsRequest")
			}
		} else {
			intVal := otg.Rocev2IPv4MetricsRequest_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2IPv4MetricsRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
