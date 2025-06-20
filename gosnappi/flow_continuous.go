package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowContinuous *****
type flowContinuous struct {
	validation
	obj          *otg.FlowContinuous
	marshaller   marshalFlowContinuous
	unMarshaller unMarshalFlowContinuous
	delayHolder  FlowDelay
}

func NewFlowContinuous() FlowContinuous {
	obj := flowContinuous{obj: &otg.FlowContinuous{}}
	obj.setDefault()
	return &obj
}

func (obj *flowContinuous) msg() *otg.FlowContinuous {
	return obj.obj
}

func (obj *flowContinuous) setMsg(msg *otg.FlowContinuous) FlowContinuous {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowContinuous struct {
	obj *flowContinuous
}

type marshalFlowContinuous interface {
	// ToProto marshals FlowContinuous to protobuf object *otg.FlowContinuous
	ToProto() (*otg.FlowContinuous, error)
	// ToPbText marshals FlowContinuous to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowContinuous to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowContinuous to JSON text
	ToJson() (string, error)
}

type unMarshalflowContinuous struct {
	obj *flowContinuous
}

type unMarshalFlowContinuous interface {
	// FromProto unmarshals FlowContinuous from protobuf object *otg.FlowContinuous
	FromProto(msg *otg.FlowContinuous) (FlowContinuous, error)
	// FromPbText unmarshals FlowContinuous from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowContinuous from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowContinuous from JSON text
	FromJson(value string) error
}

func (obj *flowContinuous) Marshal() marshalFlowContinuous {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowContinuous{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowContinuous) Unmarshal() unMarshalFlowContinuous {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowContinuous{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowContinuous) ToProto() (*otg.FlowContinuous, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowContinuous) FromProto(msg *otg.FlowContinuous) (FlowContinuous, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowContinuous) ToPbText() (string, error) {
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

func (m *unMarshalflowContinuous) FromPbText(value string) error {
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

func (m *marshalflowContinuous) ToYaml() (string, error) {
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

func (m *unMarshalflowContinuous) FromYaml(value string) error {
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

func (m *marshalflowContinuous) ToJson() (string, error) {
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

func (m *unMarshalflowContinuous) FromJson(value string) error {
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

func (obj *flowContinuous) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowContinuous) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowContinuous) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowContinuous) Clone() (FlowContinuous, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowContinuous()
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

func (obj *flowContinuous) setNil() {
	obj.delayHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowContinuous is transmit will be continuous and will not stop automatically.
type FlowContinuous interface {
	Validation
	// msg marshals FlowContinuous to protobuf object *otg.FlowContinuous
	// and doesn't set defaults
	msg() *otg.FlowContinuous
	// setMsg unmarshals FlowContinuous from protobuf object *otg.FlowContinuous
	// and doesn't set defaults
	setMsg(*otg.FlowContinuous) FlowContinuous
	// provides marshal interface
	Marshal() marshalFlowContinuous
	// provides unmarshal interface
	Unmarshal() unMarshalFlowContinuous
	// validate validates FlowContinuous
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowContinuous, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Gap returns uint32, set in FlowContinuous.
	Gap() uint32
	// SetGap assigns uint32 provided by user to FlowContinuous
	SetGap(value uint32) FlowContinuous
	// HasGap checks if Gap has been set in FlowContinuous
	HasGap() bool
	// Delay returns FlowDelay, set in FlowContinuous.
	// FlowDelay is the optional container to specify the delay before starting
	// transmission of packets.
	Delay() FlowDelay
	// SetDelay assigns FlowDelay provided by user to FlowContinuous.
	// FlowDelay is the optional container to specify the delay before starting
	// transmission of packets.
	SetDelay(value FlowDelay) FlowContinuous
	// HasDelay checks if Delay has been set in FlowContinuous
	HasDelay() bool
	setNil()
}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowContinuous) Gap() uint32 {

	return *obj.obj.Gap

}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowContinuous) HasGap() bool {
	return obj.obj.Gap != nil
}

// The minimum gap between packets expressed as bytes.
// SetGap sets the uint32 value in the FlowContinuous object
func (obj *flowContinuous) SetGap(value uint32) FlowContinuous {

	obj.obj.Gap = &value
	return obj
}

// description is TBD
// Delay returns a FlowDelay
func (obj *flowContinuous) Delay() FlowDelay {
	if obj.obj.Delay == nil {
		obj.obj.Delay = NewFlowDelay().msg()
	}
	if obj.delayHolder == nil {
		obj.delayHolder = &flowDelay{obj: obj.obj.Delay}
	}
	return obj.delayHolder
}

// description is TBD
// Delay returns a FlowDelay
func (obj *flowContinuous) HasDelay() bool {
	return obj.obj.Delay != nil
}

// description is TBD
// SetDelay sets the FlowDelay value in the FlowContinuous object
func (obj *flowContinuous) SetDelay(value FlowDelay) FlowContinuous {

	obj.delayHolder = nil
	obj.obj.Delay = value.msg()

	return obj
}

func (obj *flowContinuous) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(vObj, set_default)
	}

}

func (obj *flowContinuous) setDefault() {
	if obj.obj.Gap == nil {
		obj.SetGap(12)
	}

}
