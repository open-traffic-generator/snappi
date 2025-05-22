package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2FlowMetricsRequest *****
type rocev2FlowMetricsRequest struct {
	validation
	obj          *otg.Rocev2FlowMetricsRequest
	marshaller   marshalRocev2FlowMetricsRequest
	unMarshaller unMarshalRocev2FlowMetricsRequest
	perQpHolder  Rocev2FlowColumnNames
}

func NewRocev2FlowMetricsRequest() Rocev2FlowMetricsRequest {
	obj := rocev2FlowMetricsRequest{obj: &otg.Rocev2FlowMetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2FlowMetricsRequest) msg() *otg.Rocev2FlowMetricsRequest {
	return obj.obj
}

func (obj *rocev2FlowMetricsRequest) setMsg(msg *otg.Rocev2FlowMetricsRequest) Rocev2FlowMetricsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2FlowMetricsRequest struct {
	obj *rocev2FlowMetricsRequest
}

type marshalRocev2FlowMetricsRequest interface {
	// ToProto marshals Rocev2FlowMetricsRequest to protobuf object *otg.Rocev2FlowMetricsRequest
	ToProto() (*otg.Rocev2FlowMetricsRequest, error)
	// ToPbText marshals Rocev2FlowMetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2FlowMetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2FlowMetricsRequest to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2FlowMetricsRequest to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2FlowMetricsRequest struct {
	obj *rocev2FlowMetricsRequest
}

type unMarshalRocev2FlowMetricsRequest interface {
	// FromProto unmarshals Rocev2FlowMetricsRequest from protobuf object *otg.Rocev2FlowMetricsRequest
	FromProto(msg *otg.Rocev2FlowMetricsRequest) (Rocev2FlowMetricsRequest, error)
	// FromPbText unmarshals Rocev2FlowMetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2FlowMetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2FlowMetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *rocev2FlowMetricsRequest) Marshal() marshalRocev2FlowMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2FlowMetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2FlowMetricsRequest) Unmarshal() unMarshalRocev2FlowMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2FlowMetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2FlowMetricsRequest) ToProto() (*otg.Rocev2FlowMetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2FlowMetricsRequest) FromProto(msg *otg.Rocev2FlowMetricsRequest) (Rocev2FlowMetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2FlowMetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalrocev2FlowMetricsRequest) FromPbText(value string) error {
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

func (m *marshalrocev2FlowMetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalrocev2FlowMetricsRequest) FromYaml(value string) error {
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

func (m *marshalrocev2FlowMetricsRequest) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2FlowMetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalrocev2FlowMetricsRequest) FromJson(value string) error {
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

func (obj *rocev2FlowMetricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2FlowMetricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2FlowMetricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2FlowMetricsRequest) Clone() (Rocev2FlowMetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2FlowMetricsRequest()
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

func (obj *rocev2FlowMetricsRequest) setNil() {
	obj.perQpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2FlowMetricsRequest is request to retrieve RoCEv2 FLow statistics.
type Rocev2FlowMetricsRequest interface {
	Validation
	// msg marshals Rocev2FlowMetricsRequest to protobuf object *otg.Rocev2FlowMetricsRequest
	// and doesn't set defaults
	msg() *otg.Rocev2FlowMetricsRequest
	// setMsg unmarshals Rocev2FlowMetricsRequest from protobuf object *otg.Rocev2FlowMetricsRequest
	// and doesn't set defaults
	setMsg(*otg.Rocev2FlowMetricsRequest) Rocev2FlowMetricsRequest
	// provides marshal interface
	Marshal() marshalRocev2FlowMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2FlowMetricsRequest
	// validate validates Rocev2FlowMetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2FlowMetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2FlowMetricsRequestChoiceEnum, set in Rocev2FlowMetricsRequest
	Choice() Rocev2FlowMetricsRequestChoiceEnum
	// setChoice assigns Rocev2FlowMetricsRequestChoiceEnum provided by user to Rocev2FlowMetricsRequest
	setChoice(value Rocev2FlowMetricsRequestChoiceEnum) Rocev2FlowMetricsRequest
	// HasChoice checks if Choice has been set in Rocev2FlowMetricsRequest
	HasChoice() bool
	// PerQp returns Rocev2FlowColumnNames, set in Rocev2FlowMetricsRequest.
	// Rocev2FlowColumnNames is the names of RoCEv2 flows. An empty list will return results for all RoCEv2 flows.
	PerQp() Rocev2FlowColumnNames
	// SetPerQp assigns Rocev2FlowColumnNames provided by user to Rocev2FlowMetricsRequest.
	// Rocev2FlowColumnNames is the names of RoCEv2 flows. An empty list will return results for all RoCEv2 flows.
	SetPerQp(value Rocev2FlowColumnNames) Rocev2FlowMetricsRequest
	// HasPerQp checks if PerQp has been set in Rocev2FlowMetricsRequest
	HasPerQp() bool
	setNil()
}

type Rocev2FlowMetricsRequestChoiceEnum string

// Enum of Choice on Rocev2FlowMetricsRequest
var Rocev2FlowMetricsRequestChoice = struct {
	PER_QP Rocev2FlowMetricsRequestChoiceEnum
}{
	PER_QP: Rocev2FlowMetricsRequestChoiceEnum("per_qp"),
}

func (obj *rocev2FlowMetricsRequest) Choice() Rocev2FlowMetricsRequestChoiceEnum {
	return Rocev2FlowMetricsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// Fetch stats per QP
// Choice returns a string
func (obj *rocev2FlowMetricsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2FlowMetricsRequest) setChoice(value Rocev2FlowMetricsRequestChoiceEnum) Rocev2FlowMetricsRequest {
	intValue, ok := otg.Rocev2FlowMetricsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2FlowMetricsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2FlowMetricsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PerQp = nil
	obj.perQpHolder = nil

	if value == Rocev2FlowMetricsRequestChoice.PER_QP {
		obj.obj.PerQp = NewRocev2FlowColumnNames().msg()
	}

	return obj
}

// description is TBD
// PerQp returns a Rocev2FlowColumnNames
func (obj *rocev2FlowMetricsRequest) PerQp() Rocev2FlowColumnNames {
	if obj.obj.PerQp == nil {
		obj.setChoice(Rocev2FlowMetricsRequestChoice.PER_QP)
	}
	if obj.perQpHolder == nil {
		obj.perQpHolder = &rocev2FlowColumnNames{obj: obj.obj.PerQp}
	}
	return obj.perQpHolder
}

// description is TBD
// PerQp returns a Rocev2FlowColumnNames
func (obj *rocev2FlowMetricsRequest) HasPerQp() bool {
	return obj.obj.PerQp != nil
}

// description is TBD
// SetPerQp sets the Rocev2FlowColumnNames value in the Rocev2FlowMetricsRequest object
func (obj *rocev2FlowMetricsRequest) SetPerQp(value Rocev2FlowColumnNames) Rocev2FlowMetricsRequest {
	obj.setChoice(Rocev2FlowMetricsRequestChoice.PER_QP)
	obj.perQpHolder = nil
	obj.obj.PerQp = value.msg()

	return obj
}

func (obj *rocev2FlowMetricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PerQp != nil {

		obj.PerQp().validateObj(vObj, set_default)
	}

}

func (obj *rocev2FlowMetricsRequest) setDefault() {
	var choices_set int = 0
	var choice Rocev2FlowMetricsRequestChoiceEnum

	if obj.obj.PerQp != nil {
		choices_set += 1
		choice = Rocev2FlowMetricsRequestChoice.PER_QP
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2FlowMetricsRequestChoice.PER_QP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2FlowMetricsRequest")
			}
		} else {
			intVal := otg.Rocev2FlowMetricsRequest_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2FlowMetricsRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
