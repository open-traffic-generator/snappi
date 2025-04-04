package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2IPv6MetricsRequest *****
type rocev2IPv6MetricsRequest struct {
	validation
	obj           *otg.Rocev2IPv6MetricsRequest
	marshaller    marshalRocev2IPv6MetricsRequest
	unMarshaller  unMarshalRocev2IPv6MetricsRequest
	perPeerHolder Rocev2IPv6ColumnNames
}

func NewRocev2IPv6MetricsRequest() Rocev2IPv6MetricsRequest {
	obj := rocev2IPv6MetricsRequest{obj: &otg.Rocev2IPv6MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2IPv6MetricsRequest) msg() *otg.Rocev2IPv6MetricsRequest {
	return obj.obj
}

func (obj *rocev2IPv6MetricsRequest) setMsg(msg *otg.Rocev2IPv6MetricsRequest) Rocev2IPv6MetricsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2IPv6MetricsRequest struct {
	obj *rocev2IPv6MetricsRequest
}

type marshalRocev2IPv6MetricsRequest interface {
	// ToProto marshals Rocev2IPv6MetricsRequest to protobuf object *otg.Rocev2IPv6MetricsRequest
	ToProto() (*otg.Rocev2IPv6MetricsRequest, error)
	// ToPbText marshals Rocev2IPv6MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2IPv6MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2IPv6MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2IPv6MetricsRequest struct {
	obj *rocev2IPv6MetricsRequest
}

type unMarshalRocev2IPv6MetricsRequest interface {
	// FromProto unmarshals Rocev2IPv6MetricsRequest from protobuf object *otg.Rocev2IPv6MetricsRequest
	FromProto(msg *otg.Rocev2IPv6MetricsRequest) (Rocev2IPv6MetricsRequest, error)
	// FromPbText unmarshals Rocev2IPv6MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2IPv6MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2IPv6MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *rocev2IPv6MetricsRequest) Marshal() marshalRocev2IPv6MetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2IPv6MetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2IPv6MetricsRequest) Unmarshal() unMarshalRocev2IPv6MetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2IPv6MetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2IPv6MetricsRequest) ToProto() (*otg.Rocev2IPv6MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2IPv6MetricsRequest) FromProto(msg *otg.Rocev2IPv6MetricsRequest) (Rocev2IPv6MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2IPv6MetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalrocev2IPv6MetricsRequest) FromPbText(value string) error {
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

func (m *marshalrocev2IPv6MetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalrocev2IPv6MetricsRequest) FromYaml(value string) error {
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

func (m *marshalrocev2IPv6MetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalrocev2IPv6MetricsRequest) FromJson(value string) error {
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

func (obj *rocev2IPv6MetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2IPv6MetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2IPv6MetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2IPv6MetricsRequest) Clone() (Rocev2IPv6MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2IPv6MetricsRequest()
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

func (obj *rocev2IPv6MetricsRequest) setNil() {
	obj.perPeerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2IPv6MetricsRequest is request to retrieve RoCEv2 over IPv6 per peer metrics/statistics.
type Rocev2IPv6MetricsRequest interface {
	Validation
	// msg marshals Rocev2IPv6MetricsRequest to protobuf object *otg.Rocev2IPv6MetricsRequest
	// and doesn't set defaults
	msg() *otg.Rocev2IPv6MetricsRequest
	// setMsg unmarshals Rocev2IPv6MetricsRequest from protobuf object *otg.Rocev2IPv6MetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Rocev2IPv6MetricsRequest) Rocev2IPv6MetricsRequest
	// provides marshal interface
	Marshal() marshalRocev2IPv6MetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2IPv6MetricsRequest
	// validate validates Rocev2IPv6MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2IPv6MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2IPv6MetricsRequestChoiceEnum, set in Rocev2IPv6MetricsRequest
	Choice() Rocev2IPv6MetricsRequestChoiceEnum
	// setChoice assigns Rocev2IPv6MetricsRequestChoiceEnum provided by user to Rocev2IPv6MetricsRequest
	setChoice(value Rocev2IPv6MetricsRequestChoiceEnum) Rocev2IPv6MetricsRequest
	// HasChoice checks if Choice has been set in Rocev2IPv6MetricsRequest
	HasChoice() bool
	// PerPeer returns Rocev2IPv6ColumnNames, set in Rocev2IPv6MetricsRequest.
	// Rocev2IPv6ColumnNames is the names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
	PerPeer() Rocev2IPv6ColumnNames
	// SetPerPeer assigns Rocev2IPv6ColumnNames provided by user to Rocev2IPv6MetricsRequest.
	// Rocev2IPv6ColumnNames is the names of RoCEv2 over IPv6 peers to return results for. An empty list will return results for all RoCEv2 peers.
	SetPerPeer(value Rocev2IPv6ColumnNames) Rocev2IPv6MetricsRequest
	// HasPerPeer checks if PerPeer has been set in Rocev2IPv6MetricsRequest
	HasPerPeer() bool
	setNil()
}

type Rocev2IPv6MetricsRequestChoiceEnum string

// Enum of Choice on Rocev2IPv6MetricsRequest
var Rocev2IPv6MetricsRequestChoice = struct {
	PER_PEER Rocev2IPv6MetricsRequestChoiceEnum
}{
	PER_PEER: Rocev2IPv6MetricsRequestChoiceEnum("per_peer"),
}

func (obj *rocev2IPv6MetricsRequest) Choice() Rocev2IPv6MetricsRequestChoiceEnum {
	return Rocev2IPv6MetricsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// Fetch stats per_peer
// Choice returns a string
func (obj *rocev2IPv6MetricsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2IPv6MetricsRequest) setChoice(value Rocev2IPv6MetricsRequestChoiceEnum) Rocev2IPv6MetricsRequest {
	intValue, ok := otg.Rocev2IPv6MetricsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2IPv6MetricsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2IPv6MetricsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PerPeer = nil
	obj.perPeerHolder = nil

	if value == Rocev2IPv6MetricsRequestChoice.PER_PEER {
		obj.obj.PerPeer = NewRocev2IPv6ColumnNames().msg()
	}

	return obj
}

// description is TBD
// PerPeer returns a Rocev2IPv6ColumnNames
func (obj *rocev2IPv6MetricsRequest) PerPeer() Rocev2IPv6ColumnNames {
	if obj.obj.PerPeer == nil {
		obj.setChoice(Rocev2IPv6MetricsRequestChoice.PER_PEER)
	}
	if obj.perPeerHolder == nil {
		obj.perPeerHolder = &rocev2IPv6ColumnNames{obj: obj.obj.PerPeer}
	}
	return obj.perPeerHolder
}

// description is TBD
// PerPeer returns a Rocev2IPv6ColumnNames
func (obj *rocev2IPv6MetricsRequest) HasPerPeer() bool {
	return obj.obj.PerPeer != nil
}

// description is TBD
// SetPerPeer sets the Rocev2IPv6ColumnNames value in the Rocev2IPv6MetricsRequest object
func (obj *rocev2IPv6MetricsRequest) SetPerPeer(value Rocev2IPv6ColumnNames) Rocev2IPv6MetricsRequest {
	obj.setChoice(Rocev2IPv6MetricsRequestChoice.PER_PEER)
	obj.perPeerHolder = nil
	obj.obj.PerPeer = value.msg()

	return obj
}

func (obj *rocev2IPv6MetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PerPeer != nil {

		obj.PerPeer().validateObj(vObj, set_default)
	}

}

func (obj *rocev2IPv6MetricsRequest) setDefault() {
	var choices_set int = 0
	var choice Rocev2IPv6MetricsRequestChoiceEnum

	if obj.obj.PerPeer != nil {
		choices_set += 1
		choice = Rocev2IPv6MetricsRequestChoice.PER_PEER
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2IPv6MetricsRequestChoice.PER_PEER)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2IPv6MetricsRequest")
			}
		} else {
			intVal := otg.Rocev2IPv6MetricsRequest_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2IPv6MetricsRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
