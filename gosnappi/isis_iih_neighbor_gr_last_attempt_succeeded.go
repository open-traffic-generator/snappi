package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHNeighborGRLastAttemptSucceeded *****
type isisIIHNeighborGRLastAttemptSucceeded struct {
	validation
	obj          *otg.IsisIIHNeighborGRLastAttemptSucceeded
	marshaller   marshalIsisIIHNeighborGRLastAttemptSucceeded
	unMarshaller unMarshalIsisIIHNeighborGRLastAttemptSucceeded
}

func NewIsisIIHNeighborGRLastAttemptSucceeded() IsisIIHNeighborGRLastAttemptSucceeded {
	obj := isisIIHNeighborGRLastAttemptSucceeded{obj: &otg.IsisIIHNeighborGRLastAttemptSucceeded{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) msg() *otg.IsisIIHNeighborGRLastAttemptSucceeded {
	return obj.obj
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) setMsg(msg *otg.IsisIIHNeighborGRLastAttemptSucceeded) IsisIIHNeighborGRLastAttemptSucceeded {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHNeighborGRLastAttemptSucceeded struct {
	obj *isisIIHNeighborGRLastAttemptSucceeded
}

type marshalIsisIIHNeighborGRLastAttemptSucceeded interface {
	// ToProto marshals IsisIIHNeighborGRLastAttemptSucceeded to protobuf object *otg.IsisIIHNeighborGRLastAttemptSucceeded
	ToProto() (*otg.IsisIIHNeighborGRLastAttemptSucceeded, error)
	// ToPbText marshals IsisIIHNeighborGRLastAttemptSucceeded to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHNeighborGRLastAttemptSucceeded to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHNeighborGRLastAttemptSucceeded to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHNeighborGRLastAttemptSucceeded struct {
	obj *isisIIHNeighborGRLastAttemptSucceeded
}

type unMarshalIsisIIHNeighborGRLastAttemptSucceeded interface {
	// FromProto unmarshals IsisIIHNeighborGRLastAttemptSucceeded from protobuf object *otg.IsisIIHNeighborGRLastAttemptSucceeded
	FromProto(msg *otg.IsisIIHNeighborGRLastAttemptSucceeded) (IsisIIHNeighborGRLastAttemptSucceeded, error)
	// FromPbText unmarshals IsisIIHNeighborGRLastAttemptSucceeded from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHNeighborGRLastAttemptSucceeded from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHNeighborGRLastAttemptSucceeded from JSON text
	FromJson(value string) error
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) Marshal() marshalIsisIIHNeighborGRLastAttemptSucceeded {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHNeighborGRLastAttemptSucceeded{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) Unmarshal() unMarshalIsisIIHNeighborGRLastAttemptSucceeded {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHNeighborGRLastAttemptSucceeded{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHNeighborGRLastAttemptSucceeded) ToProto() (*otg.IsisIIHNeighborGRLastAttemptSucceeded, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHNeighborGRLastAttemptSucceeded) FromProto(msg *otg.IsisIIHNeighborGRLastAttemptSucceeded) (IsisIIHNeighborGRLastAttemptSucceeded, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHNeighborGRLastAttemptSucceeded) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptSucceeded) FromPbText(value string) error {
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

func (m *marshalisisIIHNeighborGRLastAttemptSucceeded) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptSucceeded) FromYaml(value string) error {
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

func (m *marshalisisIIHNeighborGRLastAttemptSucceeded) ToJson() (string, error) {
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

func (m *unMarshalisisIIHNeighborGRLastAttemptSucceeded) FromJson(value string) error {
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

func (obj *isisIIHNeighborGRLastAttemptSucceeded) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) Clone() (IsisIIHNeighborGRLastAttemptSucceeded, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHNeighborGRLastAttemptSucceeded()
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

// IsisIIHNeighborGRLastAttemptSucceeded is this object contains the result of a successful graceful restart status in the last attempted by a Neighbor.
type IsisIIHNeighborGRLastAttemptSucceeded interface {
	Validation
	// msg marshals IsisIIHNeighborGRLastAttemptSucceeded to protobuf object *otg.IsisIIHNeighborGRLastAttemptSucceeded
	// and doesn't set defaults
	msg() *otg.IsisIIHNeighborGRLastAttemptSucceeded
	// setMsg unmarshals IsisIIHNeighborGRLastAttemptSucceeded from protobuf object *otg.IsisIIHNeighborGRLastAttemptSucceeded
	// and doesn't set defaults
	setMsg(*otg.IsisIIHNeighborGRLastAttemptSucceeded) IsisIIHNeighborGRLastAttemptSucceeded
	// provides marshal interface
	Marshal() marshalIsisIIHNeighborGRLastAttemptSucceeded
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHNeighborGRLastAttemptSucceeded
	// validate validates IsisIIHNeighborGRLastAttemptSucceeded
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHNeighborGRLastAttemptSucceeded, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AdjacencyBringUpTime returns uint32, set in IsisIIHNeighborGRLastAttemptSucceeded.
	AdjacencyBringUpTime() uint32
	// SetAdjacencyBringUpTime assigns uint32 provided by user to IsisIIHNeighborGRLastAttemptSucceeded
	SetAdjacencyBringUpTime(value uint32) IsisIIHNeighborGRLastAttemptSucceeded
	// HasAdjacencyBringUpTime checks if AdjacencyBringUpTime has been set in IsisIIHNeighborGRLastAttemptSucceeded
	HasAdjacencyBringUpTime() bool
}

// The time (in second) is measured from when the Restart TLV with RR bit set in a IIH PDU is received up to the time
// when it receives the Restart TLV with RR bit and SA bit unset in a IIH PDU from the Neighbor Router.
// AdjacencyBringUpTime returns a uint32
func (obj *isisIIHNeighborGRLastAttemptSucceeded) AdjacencyBringUpTime() uint32 {

	return *obj.obj.AdjacencyBringUpTime

}

// The time (in second) is measured from when the Restart TLV with RR bit set in a IIH PDU is received up to the time
// when it receives the Restart TLV with RR bit and SA bit unset in a IIH PDU from the Neighbor Router.
// AdjacencyBringUpTime returns a uint32
func (obj *isisIIHNeighborGRLastAttemptSucceeded) HasAdjacencyBringUpTime() bool {
	return obj.obj.AdjacencyBringUpTime != nil
}

// The time (in second) is measured from when the Restart TLV with RR bit set in a IIH PDU is received up to the time
// when it receives the Restart TLV with RR bit and SA bit unset in a IIH PDU from the Neighbor Router.
// SetAdjacencyBringUpTime sets the uint32 value in the IsisIIHNeighborGRLastAttemptSucceeded object
func (obj *isisIIHNeighborGRLastAttemptSucceeded) SetAdjacencyBringUpTime(value uint32) IsisIIHNeighborGRLastAttemptSucceeded {

	obj.obj.AdjacencyBringUpTime = &value
	return obj
}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHNeighborGRLastAttemptSucceeded) setDefault() {

}
