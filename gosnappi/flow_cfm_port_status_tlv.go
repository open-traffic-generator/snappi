package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmPortStatusTlv *****
type flowCfmPortStatusTlv struct {
	validation
	obj          *otg.FlowCfmPortStatusTlv
	marshaller   marshalFlowCfmPortStatusTlv
	unMarshaller unMarshalFlowCfmPortStatusTlv
}

func NewFlowCfmPortStatusTlv() FlowCfmPortStatusTlv {
	obj := flowCfmPortStatusTlv{obj: &otg.FlowCfmPortStatusTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmPortStatusTlv) msg() *otg.FlowCfmPortStatusTlv {
	return obj.obj
}

func (obj *flowCfmPortStatusTlv) setMsg(msg *otg.FlowCfmPortStatusTlv) FlowCfmPortStatusTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmPortStatusTlv struct {
	obj *flowCfmPortStatusTlv
}

type marshalFlowCfmPortStatusTlv interface {
	// ToProto marshals FlowCfmPortStatusTlv to protobuf object *otg.FlowCfmPortStatusTlv
	ToProto() (*otg.FlowCfmPortStatusTlv, error)
	// ToPbText marshals FlowCfmPortStatusTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmPortStatusTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmPortStatusTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmPortStatusTlv struct {
	obj *flowCfmPortStatusTlv
}

type unMarshalFlowCfmPortStatusTlv interface {
	// FromProto unmarshals FlowCfmPortStatusTlv from protobuf object *otg.FlowCfmPortStatusTlv
	FromProto(msg *otg.FlowCfmPortStatusTlv) (FlowCfmPortStatusTlv, error)
	// FromPbText unmarshals FlowCfmPortStatusTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmPortStatusTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmPortStatusTlv from JSON text
	FromJson(value string) error
}

func (obj *flowCfmPortStatusTlv) Marshal() marshalFlowCfmPortStatusTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmPortStatusTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmPortStatusTlv) Unmarshal() unMarshalFlowCfmPortStatusTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmPortStatusTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmPortStatusTlv) ToProto() (*otg.FlowCfmPortStatusTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmPortStatusTlv) FromProto(msg *otg.FlowCfmPortStatusTlv) (FlowCfmPortStatusTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmPortStatusTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmPortStatusTlv) FromPbText(value string) error {
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

func (m *marshalflowCfmPortStatusTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmPortStatusTlv) FromYaml(value string) error {
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

func (m *marshalflowCfmPortStatusTlv) ToJson() (string, error) {
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

func (m *unMarshalflowCfmPortStatusTlv) FromJson(value string) error {
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

func (obj *flowCfmPortStatusTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmPortStatusTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmPortStatusTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmPortStatusTlv) Clone() (FlowCfmPortStatusTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmPortStatusTlv()
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

// FlowCfmPortStatusTlv is port status
type FlowCfmPortStatusTlv interface {
	Validation
	// msg marshals FlowCfmPortStatusTlv to protobuf object *otg.FlowCfmPortStatusTlv
	// and doesn't set defaults
	msg() *otg.FlowCfmPortStatusTlv
	// setMsg unmarshals FlowCfmPortStatusTlv from protobuf object *otg.FlowCfmPortStatusTlv
	// and doesn't set defaults
	setMsg(*otg.FlowCfmPortStatusTlv) FlowCfmPortStatusTlv
	// provides marshal interface
	Marshal() marshalFlowCfmPortStatusTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmPortStatusTlv
	// validate validates FlowCfmPortStatusTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmPortStatusTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Status returns FlowCfmPortStatusTlvStatusEnum, set in FlowCfmPortStatusTlv
	Status() FlowCfmPortStatusTlvStatusEnum
	// SetStatus assigns FlowCfmPortStatusTlvStatusEnum provided by user to FlowCfmPortStatusTlv
	SetStatus(value FlowCfmPortStatusTlvStatusEnum) FlowCfmPortStatusTlv
	// HasStatus checks if Status has been set in FlowCfmPortStatusTlv
	HasStatus() bool
}

type FlowCfmPortStatusTlvStatusEnum string

// Enum of Status on FlowCfmPortStatusTlv
var FlowCfmPortStatusTlvStatus = struct {
	PORT_BLOCKED FlowCfmPortStatusTlvStatusEnum
	PORT_UP      FlowCfmPortStatusTlvStatusEnum
}{
	PORT_BLOCKED: FlowCfmPortStatusTlvStatusEnum("port_blocked"),
	PORT_UP:      FlowCfmPortStatusTlvStatusEnum("port_up"),
}

func (obj *flowCfmPortStatusTlv) Status() FlowCfmPortStatusTlvStatusEnum {
	return FlowCfmPortStatusTlvStatusEnum(obj.obj.Status.Enum().String())
}

// Port status can be either port blocked or port up
// Status returns a string
func (obj *flowCfmPortStatusTlv) HasStatus() bool {
	return obj.obj.Status != nil
}

func (obj *flowCfmPortStatusTlv) SetStatus(value FlowCfmPortStatusTlvStatusEnum) FlowCfmPortStatusTlv {
	intValue, ok := otg.FlowCfmPortStatusTlv_Status_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmPortStatusTlvStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmPortStatusTlv_Status_Enum(intValue)
	obj.obj.Status = &enumValue

	return obj
}

func (obj *flowCfmPortStatusTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowCfmPortStatusTlv) setDefault() {
	if obj.obj.Status == nil {
		obj.SetStatus(FlowCfmPortStatusTlvStatus.PORT_UP)

	}

}
