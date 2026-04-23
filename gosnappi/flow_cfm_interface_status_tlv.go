package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmInterfaceStatusTlv *****
type flowCfmInterfaceStatusTlv struct {
	validation
	obj          *otg.FlowCfmInterfaceStatusTlv
	marshaller   marshalFlowCfmInterfaceStatusTlv
	unMarshaller unMarshalFlowCfmInterfaceStatusTlv
}

func NewFlowCfmInterfaceStatusTlv() FlowCfmInterfaceStatusTlv {
	obj := flowCfmInterfaceStatusTlv{obj: &otg.FlowCfmInterfaceStatusTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmInterfaceStatusTlv) msg() *otg.FlowCfmInterfaceStatusTlv {
	return obj.obj
}

func (obj *flowCfmInterfaceStatusTlv) setMsg(msg *otg.FlowCfmInterfaceStatusTlv) FlowCfmInterfaceStatusTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmInterfaceStatusTlv struct {
	obj *flowCfmInterfaceStatusTlv
}

type marshalFlowCfmInterfaceStatusTlv interface {
	// ToProto marshals FlowCfmInterfaceStatusTlv to protobuf object *otg.FlowCfmInterfaceStatusTlv
	ToProto() (*otg.FlowCfmInterfaceStatusTlv, error)
	// ToPbText marshals FlowCfmInterfaceStatusTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmInterfaceStatusTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmInterfaceStatusTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmInterfaceStatusTlv struct {
	obj *flowCfmInterfaceStatusTlv
}

type unMarshalFlowCfmInterfaceStatusTlv interface {
	// FromProto unmarshals FlowCfmInterfaceStatusTlv from protobuf object *otg.FlowCfmInterfaceStatusTlv
	FromProto(msg *otg.FlowCfmInterfaceStatusTlv) (FlowCfmInterfaceStatusTlv, error)
	// FromPbText unmarshals FlowCfmInterfaceStatusTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmInterfaceStatusTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmInterfaceStatusTlv from JSON text
	FromJson(value string) error
}

func (obj *flowCfmInterfaceStatusTlv) Marshal() marshalFlowCfmInterfaceStatusTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmInterfaceStatusTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmInterfaceStatusTlv) Unmarshal() unMarshalFlowCfmInterfaceStatusTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmInterfaceStatusTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmInterfaceStatusTlv) ToProto() (*otg.FlowCfmInterfaceStatusTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmInterfaceStatusTlv) FromProto(msg *otg.FlowCfmInterfaceStatusTlv) (FlowCfmInterfaceStatusTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmInterfaceStatusTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmInterfaceStatusTlv) FromPbText(value string) error {
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

func (m *marshalflowCfmInterfaceStatusTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmInterfaceStatusTlv) FromYaml(value string) error {
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

func (m *marshalflowCfmInterfaceStatusTlv) ToJson() (string, error) {
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

func (m *unMarshalflowCfmInterfaceStatusTlv) FromJson(value string) error {
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

func (obj *flowCfmInterfaceStatusTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmInterfaceStatusTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmInterfaceStatusTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmInterfaceStatusTlv) Clone() (FlowCfmInterfaceStatusTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmInterfaceStatusTlv()
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

// FlowCfmInterfaceStatusTlv is interface status
type FlowCfmInterfaceStatusTlv interface {
	Validation
	// msg marshals FlowCfmInterfaceStatusTlv to protobuf object *otg.FlowCfmInterfaceStatusTlv
	// and doesn't set defaults
	msg() *otg.FlowCfmInterfaceStatusTlv
	// setMsg unmarshals FlowCfmInterfaceStatusTlv from protobuf object *otg.FlowCfmInterfaceStatusTlv
	// and doesn't set defaults
	setMsg(*otg.FlowCfmInterfaceStatusTlv) FlowCfmInterfaceStatusTlv
	// provides marshal interface
	Marshal() marshalFlowCfmInterfaceStatusTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmInterfaceStatusTlv
	// validate validates FlowCfmInterfaceStatusTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmInterfaceStatusTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Status returns FlowCfmInterfaceStatusTlvStatusEnum, set in FlowCfmInterfaceStatusTlv
	Status() FlowCfmInterfaceStatusTlvStatusEnum
	// SetStatus assigns FlowCfmInterfaceStatusTlvStatusEnum provided by user to FlowCfmInterfaceStatusTlv
	SetStatus(value FlowCfmInterfaceStatusTlvStatusEnum) FlowCfmInterfaceStatusTlv
	// HasStatus checks if Status has been set in FlowCfmInterfaceStatusTlv
	HasStatus() bool
}

type FlowCfmInterfaceStatusTlvStatusEnum string

// Enum of Status on FlowCfmInterfaceStatusTlv
var FlowCfmInterfaceStatusTlvStatus = struct {
	UP               FlowCfmInterfaceStatusTlvStatusEnum
	DOWN             FlowCfmInterfaceStatusTlvStatusEnum
	TESTING          FlowCfmInterfaceStatusTlvStatusEnum
	UNKNOWN          FlowCfmInterfaceStatusTlvStatusEnum
	DORMANT          FlowCfmInterfaceStatusTlvStatusEnum
	NOT_PRESENT      FlowCfmInterfaceStatusTlvStatusEnum
	LOWER_LAYER_DOWN FlowCfmInterfaceStatusTlvStatusEnum
}{
	UP:               FlowCfmInterfaceStatusTlvStatusEnum("up"),
	DOWN:             FlowCfmInterfaceStatusTlvStatusEnum("down"),
	TESTING:          FlowCfmInterfaceStatusTlvStatusEnum("testing"),
	UNKNOWN:          FlowCfmInterfaceStatusTlvStatusEnum("unknown"),
	DORMANT:          FlowCfmInterfaceStatusTlvStatusEnum("dormant"),
	NOT_PRESENT:      FlowCfmInterfaceStatusTlvStatusEnum("not_present"),
	LOWER_LAYER_DOWN: FlowCfmInterfaceStatusTlvStatusEnum("lower_layer_down"),
}

func (obj *flowCfmInterfaceStatusTlv) Status() FlowCfmInterfaceStatusTlvStatusEnum {
	return FlowCfmInterfaceStatusTlvStatusEnum(obj.obj.Status.Enum().String())
}

// Defined interface status
// Status returns a string
func (obj *flowCfmInterfaceStatusTlv) HasStatus() bool {
	return obj.obj.Status != nil
}

func (obj *flowCfmInterfaceStatusTlv) SetStatus(value FlowCfmInterfaceStatusTlvStatusEnum) FlowCfmInterfaceStatusTlv {
	intValue, ok := otg.FlowCfmInterfaceStatusTlv_Status_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmInterfaceStatusTlvStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmInterfaceStatusTlv_Status_Enum(intValue)
	obj.obj.Status = &enumValue

	return obj
}

func (obj *flowCfmInterfaceStatusTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *flowCfmInterfaceStatusTlv) setDefault() {
	if obj.obj.Status == nil {
		obj.SetStatus(FlowCfmInterfaceStatusTlvStatus.UP)

	}

}
