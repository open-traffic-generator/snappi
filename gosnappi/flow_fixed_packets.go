package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowFixedPackets *****
type flowFixedPackets struct {
	validation
	obj          *otg.FlowFixedPackets
	marshaller   marshalFlowFixedPackets
	unMarshaller unMarshalFlowFixedPackets
	delayHolder  FlowDelay
}

func NewFlowFixedPackets() FlowFixedPackets {
	obj := flowFixedPackets{obj: &otg.FlowFixedPackets{}}
	obj.setDefault()
	return &obj
}

func (obj *flowFixedPackets) msg() *otg.FlowFixedPackets {
	return obj.obj
}

func (obj *flowFixedPackets) setMsg(msg *otg.FlowFixedPackets) FlowFixedPackets {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowFixedPackets struct {
	obj *flowFixedPackets
}

type marshalFlowFixedPackets interface {
	// ToProto marshals FlowFixedPackets to protobuf object *otg.FlowFixedPackets
	ToProto() (*otg.FlowFixedPackets, error)
	// ToPbText marshals FlowFixedPackets to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowFixedPackets to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowFixedPackets to JSON text
	ToJson() (string, error)
}

type unMarshalflowFixedPackets struct {
	obj *flowFixedPackets
}

type unMarshalFlowFixedPackets interface {
	// FromProto unmarshals FlowFixedPackets from protobuf object *otg.FlowFixedPackets
	FromProto(msg *otg.FlowFixedPackets) (FlowFixedPackets, error)
	// FromPbText unmarshals FlowFixedPackets from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowFixedPackets from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowFixedPackets from JSON text
	FromJson(value string) error
}

func (obj *flowFixedPackets) Marshal() marshalFlowFixedPackets {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowFixedPackets{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowFixedPackets) Unmarshal() unMarshalFlowFixedPackets {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowFixedPackets{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowFixedPackets) ToProto() (*otg.FlowFixedPackets, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowFixedPackets) FromProto(msg *otg.FlowFixedPackets) (FlowFixedPackets, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowFixedPackets) ToPbText() (string, error) {
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

func (m *unMarshalflowFixedPackets) FromPbText(value string) error {
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

func (m *marshalflowFixedPackets) ToYaml() (string, error) {
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

func (m *unMarshalflowFixedPackets) FromYaml(value string) error {
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

func (m *marshalflowFixedPackets) ToJson() (string, error) {
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

func (m *unMarshalflowFixedPackets) FromJson(value string) error {
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

func (obj *flowFixedPackets) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowFixedPackets) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowFixedPackets) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowFixedPackets) Clone() (FlowFixedPackets, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowFixedPackets()
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

func (obj *flowFixedPackets) setNil() {
	obj.delayHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowFixedPackets is transmit a fixed number of packets after which the flow will stop.
type FlowFixedPackets interface {
	Validation
	// msg marshals FlowFixedPackets to protobuf object *otg.FlowFixedPackets
	// and doesn't set defaults
	msg() *otg.FlowFixedPackets
	// setMsg unmarshals FlowFixedPackets from protobuf object *otg.FlowFixedPackets
	// and doesn't set defaults
	setMsg(*otg.FlowFixedPackets) FlowFixedPackets
	// provides marshal interface
	Marshal() marshalFlowFixedPackets
	// provides unmarshal interface
	Unmarshal() unMarshalFlowFixedPackets
	// validate validates FlowFixedPackets
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowFixedPackets, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Packets returns uint32, set in FlowFixedPackets.
	Packets() uint32
	// SetPackets assigns uint32 provided by user to FlowFixedPackets
	SetPackets(value uint32) FlowFixedPackets
	// HasPackets checks if Packets has been set in FlowFixedPackets
	HasPackets() bool
	// Gap returns uint32, set in FlowFixedPackets.
	Gap() uint32
	// SetGap assigns uint32 provided by user to FlowFixedPackets
	SetGap(value uint32) FlowFixedPackets
	// HasGap checks if Gap has been set in FlowFixedPackets
	HasGap() bool
	// Delay returns FlowDelay, set in FlowFixedPackets.
	// FlowDelay is the optional container to specify the delay before starting
	// transmission of packets.
	Delay() FlowDelay
	// SetDelay assigns FlowDelay provided by user to FlowFixedPackets.
	// FlowDelay is the optional container to specify the delay before starting
	// transmission of packets.
	SetDelay(value FlowDelay) FlowFixedPackets
	// HasDelay checks if Delay has been set in FlowFixedPackets
	HasDelay() bool
	setNil()
}

// Stop transmit of the flow after this number of packets.
// Packets returns a uint32
func (obj *flowFixedPackets) Packets() uint32 {

	return *obj.obj.Packets

}

// Stop transmit of the flow after this number of packets.
// Packets returns a uint32
func (obj *flowFixedPackets) HasPackets() bool {
	return obj.obj.Packets != nil
}

// Stop transmit of the flow after this number of packets.
// SetPackets sets the uint32 value in the FlowFixedPackets object
func (obj *flowFixedPackets) SetPackets(value uint32) FlowFixedPackets {

	obj.obj.Packets = &value
	return obj
}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowFixedPackets) Gap() uint32 {

	return *obj.obj.Gap

}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowFixedPackets) HasGap() bool {
	return obj.obj.Gap != nil
}

// The minimum gap between packets expressed as bytes.
// SetGap sets the uint32 value in the FlowFixedPackets object
func (obj *flowFixedPackets) SetGap(value uint32) FlowFixedPackets {

	obj.obj.Gap = &value
	return obj
}

// description is TBD
// Delay returns a FlowDelay
func (obj *flowFixedPackets) Delay() FlowDelay {
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
func (obj *flowFixedPackets) HasDelay() bool {
	return obj.obj.Delay != nil
}

// description is TBD
// SetDelay sets the FlowDelay value in the FlowFixedPackets object
func (obj *flowFixedPackets) SetDelay(value FlowDelay) FlowFixedPackets {

	obj.delayHolder = nil
	obj.obj.Delay = value.msg()

	return obj
}

func (obj *flowFixedPackets) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Packets != nil {

		if *obj.obj.Packets < 1 || *obj.obj.Packets > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowFixedPackets.Packets <= 4294967295 but Got %d", *obj.obj.Packets))
		}

	}

	if obj.obj.Delay != nil {

		obj.Delay().validateObj(vObj, set_default)
	}

}

func (obj *flowFixedPackets) setDefault() {
	if obj.obj.Packets == nil {
		obj.SetPackets(1)
	}
	if obj.obj.Gap == nil {
		obj.SetGap(12)
	}

}
