package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisActionProtocolIsisSubpressAdjacency *****
type isisActionProtocolIsisSubpressAdjacency struct {
	validation
	obj          *otg.IsisActionProtocolIsisSubpressAdjacency
	marshaller   marshalIsisActionProtocolIsisSubpressAdjacency
	unMarshaller unMarshalIsisActionProtocolIsisSubpressAdjacency
}

func NewIsisActionProtocolIsisSubpressAdjacency() IsisActionProtocolIsisSubpressAdjacency {
	obj := isisActionProtocolIsisSubpressAdjacency{obj: &otg.IsisActionProtocolIsisSubpressAdjacency{}}
	obj.setDefault()
	return &obj
}

func (obj *isisActionProtocolIsisSubpressAdjacency) msg() *otg.IsisActionProtocolIsisSubpressAdjacency {
	return obj.obj
}

func (obj *isisActionProtocolIsisSubpressAdjacency) setMsg(msg *otg.IsisActionProtocolIsisSubpressAdjacency) IsisActionProtocolIsisSubpressAdjacency {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisActionProtocolIsisSubpressAdjacency struct {
	obj *isisActionProtocolIsisSubpressAdjacency
}

type marshalIsisActionProtocolIsisSubpressAdjacency interface {
	// ToProto marshals IsisActionProtocolIsisSubpressAdjacency to protobuf object *otg.IsisActionProtocolIsisSubpressAdjacency
	ToProto() (*otg.IsisActionProtocolIsisSubpressAdjacency, error)
	// ToPbText marshals IsisActionProtocolIsisSubpressAdjacency to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisActionProtocolIsisSubpressAdjacency to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisActionProtocolIsisSubpressAdjacency to JSON text
	ToJson() (string, error)
}

type unMarshalisisActionProtocolIsisSubpressAdjacency struct {
	obj *isisActionProtocolIsisSubpressAdjacency
}

type unMarshalIsisActionProtocolIsisSubpressAdjacency interface {
	// FromProto unmarshals IsisActionProtocolIsisSubpressAdjacency from protobuf object *otg.IsisActionProtocolIsisSubpressAdjacency
	FromProto(msg *otg.IsisActionProtocolIsisSubpressAdjacency) (IsisActionProtocolIsisSubpressAdjacency, error)
	// FromPbText unmarshals IsisActionProtocolIsisSubpressAdjacency from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisActionProtocolIsisSubpressAdjacency from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisActionProtocolIsisSubpressAdjacency from JSON text
	FromJson(value string) error
}

func (obj *isisActionProtocolIsisSubpressAdjacency) Marshal() marshalIsisActionProtocolIsisSubpressAdjacency {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisActionProtocolIsisSubpressAdjacency{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisActionProtocolIsisSubpressAdjacency) Unmarshal() unMarshalIsisActionProtocolIsisSubpressAdjacency {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisActionProtocolIsisSubpressAdjacency{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisActionProtocolIsisSubpressAdjacency) ToProto() (*otg.IsisActionProtocolIsisSubpressAdjacency, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisActionProtocolIsisSubpressAdjacency) FromProto(msg *otg.IsisActionProtocolIsisSubpressAdjacency) (IsisActionProtocolIsisSubpressAdjacency, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisActionProtocolIsisSubpressAdjacency) ToPbText() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisSubpressAdjacency) FromPbText(value string) error {
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

func (m *marshalisisActionProtocolIsisSubpressAdjacency) ToYaml() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisSubpressAdjacency) FromYaml(value string) error {
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

func (m *marshalisisActionProtocolIsisSubpressAdjacency) ToJson() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisSubpressAdjacency) FromJson(value string) error {
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

func (obj *isisActionProtocolIsisSubpressAdjacency) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisSubpressAdjacency) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisSubpressAdjacency) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisActionProtocolIsisSubpressAdjacency) Clone() (IsisActionProtocolIsisSubpressAdjacency, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisActionProtocolIsisSubpressAdjacency()
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

// IsisActionProtocolIsisSubpressAdjacency is initiates IS-IS Suppress adjacency advertisement process for the selected IS-IS routers. If no name is specified then Start will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress adjacency advertisement) bit set.
type IsisActionProtocolIsisSubpressAdjacency interface {
	Validation
	// msg marshals IsisActionProtocolIsisSubpressAdjacency to protobuf object *otg.IsisActionProtocolIsisSubpressAdjacency
	// and doesn't set defaults
	msg() *otg.IsisActionProtocolIsisSubpressAdjacency
	// setMsg unmarshals IsisActionProtocolIsisSubpressAdjacency from protobuf object *otg.IsisActionProtocolIsisSubpressAdjacency
	// and doesn't set defaults
	setMsg(*otg.IsisActionProtocolIsisSubpressAdjacency) IsisActionProtocolIsisSubpressAdjacency
	// provides marshal interface
	Marshal() marshalIsisActionProtocolIsisSubpressAdjacency
	// provides unmarshal interface
	Unmarshal() unMarshalIsisActionProtocolIsisSubpressAdjacency
	// validate validates IsisActionProtocolIsisSubpressAdjacency
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisActionProtocolIsisSubpressAdjacency, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HoldingTime returns uint32, set in IsisActionProtocolIsisSubpressAdjacency.
	HoldingTime() uint32
	// SetHoldingTime assigns uint32 provided by user to IsisActionProtocolIsisSubpressAdjacency
	SetHoldingTime(value uint32) IsisActionProtocolIsisSubpressAdjacency
	// HasHoldingTime checks if HoldingTime has been set in IsisActionProtocolIsisSubpressAdjacency
	HasHoldingTime() bool
	// RestartAfter returns uint32, set in IsisActionProtocolIsisSubpressAdjacency.
	RestartAfter() uint32
	// SetRestartAfter assigns uint32 provided by user to IsisActionProtocolIsisSubpressAdjacency
	SetRestartAfter(value uint32) IsisActionProtocolIsisSubpressAdjacency
	// HasRestartAfter checks if RestartAfter has been set in IsisActionProtocolIsisSubpressAdjacency
	HasRestartAfter() bool
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start by initiating Restart TLV having SA bit set in IIH. The hold-timer in the IIH is updated with this time.
// HoldingTime returns a uint32
func (obj *isisActionProtocolIsisSubpressAdjacency) HoldingTime() uint32 {

	return *obj.obj.HoldingTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start by initiating Restart TLV having SA bit set in IIH. The hold-timer in the IIH is updated with this time.
// HoldingTime returns a uint32
func (obj *isisActionProtocolIsisSubpressAdjacency) HasHoldingTime() bool {
	return obj.obj.HoldingTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start by initiating Restart TLV having SA bit set in IIH. The hold-timer in the IIH is updated with this time.
// SetHoldingTime sets the uint32 value in the IsisActionProtocolIsisSubpressAdjacency object
func (obj *isisActionProtocolIsisSubpressAdjacency) SetHoldingTime(value uint32) IsisActionProtocolIsisSubpressAdjacency {

	obj.obj.HoldingTime = &value
	return obj
}

// Time (in seconds) after which the actually Restart TLV will be initiated in IIH with Restart Tlv having SA bit set.
// RestartAfter returns a uint32
func (obj *isisActionProtocolIsisSubpressAdjacency) RestartAfter() uint32 {

	return *obj.obj.RestartAfter

}

// Time (in seconds) after which the actually Restart TLV will be initiated in IIH with Restart Tlv having SA bit set.
// RestartAfter returns a uint32
func (obj *isisActionProtocolIsisSubpressAdjacency) HasRestartAfter() bool {
	return obj.obj.RestartAfter != nil
}

// Time (in seconds) after which the actually Restart TLV will be initiated in IIH with Restart Tlv having SA bit set.
// SetRestartAfter sets the uint32 value in the IsisActionProtocolIsisSubpressAdjacency object
func (obj *isisActionProtocolIsisSubpressAdjacency) SetRestartAfter(value uint32) IsisActionProtocolIsisSubpressAdjacency {

	obj.obj.RestartAfter = &value
	return obj
}

func (obj *isisActionProtocolIsisSubpressAdjacency) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HoldingTime != nil {

		if *obj.obj.HoldingTime < 10 || *obj.obj.HoldingTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= IsisActionProtocolIsisSubpressAdjacency.HoldingTime <= 4096 but Got %d", *obj.obj.HoldingTime))
		}

	}

	if obj.obj.RestartAfter != nil {

		if *obj.obj.RestartAfter < 10 || *obj.obj.RestartAfter > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= IsisActionProtocolIsisSubpressAdjacency.RestartAfter <= 4096 but Got %d", *obj.obj.RestartAfter))
		}

	}

}

func (obj *isisActionProtocolIsisSubpressAdjacency) setDefault() {
	if obj.obj.HoldingTime == nil {
		obj.SetHoldingTime(30)
	}
	if obj.obj.RestartAfter == nil {
		obj.SetRestartAfter(10)
	}

}
