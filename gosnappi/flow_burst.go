package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowBurst *****
type flowBurst struct {
	validation
	obj                 *otg.FlowBurst
	marshaller          marshalFlowBurst
	unMarshaller        unMarshalFlowBurst
	interBurstGapHolder FlowDurationInterBurstGap
}

func NewFlowBurst() FlowBurst {
	obj := flowBurst{obj: &otg.FlowBurst{}}
	obj.setDefault()
	return &obj
}

func (obj *flowBurst) msg() *otg.FlowBurst {
	return obj.obj
}

func (obj *flowBurst) setMsg(msg *otg.FlowBurst) FlowBurst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowBurst struct {
	obj *flowBurst
}

type marshalFlowBurst interface {
	// ToProto marshals FlowBurst to protobuf object *otg.FlowBurst
	ToProto() (*otg.FlowBurst, error)
	// ToPbText marshals FlowBurst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowBurst to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowBurst to JSON text
	ToJson() (string, error)
}

type unMarshalflowBurst struct {
	obj *flowBurst
}

type unMarshalFlowBurst interface {
	// FromProto unmarshals FlowBurst from protobuf object *otg.FlowBurst
	FromProto(msg *otg.FlowBurst) (FlowBurst, error)
	// FromPbText unmarshals FlowBurst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowBurst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowBurst from JSON text
	FromJson(value string) error
}

func (obj *flowBurst) Marshal() marshalFlowBurst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowBurst{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowBurst) Unmarshal() unMarshalFlowBurst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowBurst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowBurst) ToProto() (*otg.FlowBurst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowBurst) FromProto(msg *otg.FlowBurst) (FlowBurst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowBurst) ToPbText() (string, error) {
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

func (m *unMarshalflowBurst) FromPbText(value string) error {
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

func (m *marshalflowBurst) ToYaml() (string, error) {
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

func (m *unMarshalflowBurst) FromYaml(value string) error {
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

func (m *marshalflowBurst) ToJson() (string, error) {
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

func (m *unMarshalflowBurst) FromJson(value string) error {
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

func (obj *flowBurst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowBurst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowBurst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowBurst) Clone() (FlowBurst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowBurst()
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

func (obj *flowBurst) setNil() {
	obj.interBurstGapHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowBurst is transmits continuous or fixed burst of packets.
// For continuous burst of packets, it will not automatically stop.
// For fixed burst of packets, it will stop after transmitting fixed number of bursts.
type FlowBurst interface {
	Validation
	// msg marshals FlowBurst to protobuf object *otg.FlowBurst
	// and doesn't set defaults
	msg() *otg.FlowBurst
	// setMsg unmarshals FlowBurst from protobuf object *otg.FlowBurst
	// and doesn't set defaults
	setMsg(*otg.FlowBurst) FlowBurst
	// provides marshal interface
	Marshal() marshalFlowBurst
	// provides unmarshal interface
	Unmarshal() unMarshalFlowBurst
	// validate validates FlowBurst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowBurst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Bursts returns uint32, set in FlowBurst.
	Bursts() uint32
	// SetBursts assigns uint32 provided by user to FlowBurst
	SetBursts(value uint32) FlowBurst
	// HasBursts checks if Bursts has been set in FlowBurst
	HasBursts() bool
	// Packets returns uint32, set in FlowBurst.
	Packets() uint32
	// SetPackets assigns uint32 provided by user to FlowBurst
	SetPackets(value uint32) FlowBurst
	// HasPackets checks if Packets has been set in FlowBurst
	HasPackets() bool
	// Gap returns uint32, set in FlowBurst.
	Gap() uint32
	// SetGap assigns uint32 provided by user to FlowBurst
	SetGap(value uint32) FlowBurst
	// HasGap checks if Gap has been set in FlowBurst
	HasGap() bool
	// InterBurstGap returns FlowDurationInterBurstGap, set in FlowBurst.
	// FlowDurationInterBurstGap is the optional container for specifying a gap between bursts.
	InterBurstGap() FlowDurationInterBurstGap
	// SetInterBurstGap assigns FlowDurationInterBurstGap provided by user to FlowBurst.
	// FlowDurationInterBurstGap is the optional container for specifying a gap between bursts.
	SetInterBurstGap(value FlowDurationInterBurstGap) FlowBurst
	// HasInterBurstGap checks if InterBurstGap has been set in FlowBurst
	HasInterBurstGap() bool
	setNil()
}

// The number of packet bursts transmitted per flow.
// A value of 0 implies continuous burst of packets.
// Bursts returns a uint32
func (obj *flowBurst) Bursts() uint32 {

	return *obj.obj.Bursts

}

// The number of packet bursts transmitted per flow.
// A value of 0 implies continuous burst of packets.
// Bursts returns a uint32
func (obj *flowBurst) HasBursts() bool {
	return obj.obj.Bursts != nil
}

// The number of packet bursts transmitted per flow.
// A value of 0 implies continuous burst of packets.
// SetBursts sets the uint32 value in the FlowBurst object
func (obj *flowBurst) SetBursts(value uint32) FlowBurst {

	obj.obj.Bursts = &value
	return obj
}

// The number of packets transmitted per burst.
// Packets returns a uint32
func (obj *flowBurst) Packets() uint32 {

	return *obj.obj.Packets

}

// The number of packets transmitted per burst.
// Packets returns a uint32
func (obj *flowBurst) HasPackets() bool {
	return obj.obj.Packets != nil
}

// The number of packets transmitted per burst.
// SetPackets sets the uint32 value in the FlowBurst object
func (obj *flowBurst) SetPackets(value uint32) FlowBurst {

	obj.obj.Packets = &value
	return obj
}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowBurst) Gap() uint32 {

	return *obj.obj.Gap

}

// The minimum gap between packets expressed as bytes.
// Gap returns a uint32
func (obj *flowBurst) HasGap() bool {
	return obj.obj.Gap != nil
}

// The minimum gap between packets expressed as bytes.
// SetGap sets the uint32 value in the FlowBurst object
func (obj *flowBurst) SetGap(value uint32) FlowBurst {

	obj.obj.Gap = &value
	return obj
}

// description is TBD
// InterBurstGap returns a FlowDurationInterBurstGap
func (obj *flowBurst) InterBurstGap() FlowDurationInterBurstGap {
	if obj.obj.InterBurstGap == nil {
		obj.obj.InterBurstGap = NewFlowDurationInterBurstGap().msg()
	}
	if obj.interBurstGapHolder == nil {
		obj.interBurstGapHolder = &flowDurationInterBurstGap{obj: obj.obj.InterBurstGap}
	}
	return obj.interBurstGapHolder
}

// description is TBD
// InterBurstGap returns a FlowDurationInterBurstGap
func (obj *flowBurst) HasInterBurstGap() bool {
	return obj.obj.InterBurstGap != nil
}

// description is TBD
// SetInterBurstGap sets the FlowDurationInterBurstGap value in the FlowBurst object
func (obj *flowBurst) SetInterBurstGap(value FlowDurationInterBurstGap) FlowBurst {

	obj.interBurstGapHolder = nil
	obj.obj.InterBurstGap = value.msg()

	return obj
}

func (obj *flowBurst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Packets != nil {

		if *obj.obj.Packets < 1 || *obj.obj.Packets > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowBurst.Packets <= 4294967295 but Got %d", *obj.obj.Packets))
		}

	}

	if obj.obj.InterBurstGap != nil {

		obj.InterBurstGap().validateObj(vObj, set_default)
	}

}

func (obj *flowBurst) setDefault() {
	if obj.obj.Bursts == nil {
		obj.SetBursts(0)
	}
	if obj.obj.Packets == nil {
		obj.SetPackets(1)
	}
	if obj.obj.Gap == nil {
		obj.SetGap(12)
	}

}
