package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHNeighborRestartSucceded *****
type isisIIHNeighborRestartSucceded struct {
	validation
	obj          *otg.IsisIIHNeighborRestartSucceded
	marshaller   marshalIsisIIHNeighborRestartSucceded
	unMarshaller unMarshalIsisIIHNeighborRestartSucceded
}

func NewIsisIIHNeighborRestartSucceded() IsisIIHNeighborRestartSucceded {
	obj := isisIIHNeighborRestartSucceded{obj: &otg.IsisIIHNeighborRestartSucceded{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHNeighborRestartSucceded) msg() *otg.IsisIIHNeighborRestartSucceded {
	return obj.obj
}

func (obj *isisIIHNeighborRestartSucceded) setMsg(msg *otg.IsisIIHNeighborRestartSucceded) IsisIIHNeighborRestartSucceded {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHNeighborRestartSucceded struct {
	obj *isisIIHNeighborRestartSucceded
}

type marshalIsisIIHNeighborRestartSucceded interface {
	// ToProto marshals IsisIIHNeighborRestartSucceded to protobuf object *otg.IsisIIHNeighborRestartSucceded
	ToProto() (*otg.IsisIIHNeighborRestartSucceded, error)
	// ToPbText marshals IsisIIHNeighborRestartSucceded to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHNeighborRestartSucceded to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHNeighborRestartSucceded to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHNeighborRestartSucceded struct {
	obj *isisIIHNeighborRestartSucceded
}

type unMarshalIsisIIHNeighborRestartSucceded interface {
	// FromProto unmarshals IsisIIHNeighborRestartSucceded from protobuf object *otg.IsisIIHNeighborRestartSucceded
	FromProto(msg *otg.IsisIIHNeighborRestartSucceded) (IsisIIHNeighborRestartSucceded, error)
	// FromPbText unmarshals IsisIIHNeighborRestartSucceded from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHNeighborRestartSucceded from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHNeighborRestartSucceded from JSON text
	FromJson(value string) error
}

func (obj *isisIIHNeighborRestartSucceded) Marshal() marshalIsisIIHNeighborRestartSucceded {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHNeighborRestartSucceded{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHNeighborRestartSucceded) Unmarshal() unMarshalIsisIIHNeighborRestartSucceded {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHNeighborRestartSucceded{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHNeighborRestartSucceded) ToProto() (*otg.IsisIIHNeighborRestartSucceded, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHNeighborRestartSucceded) FromProto(msg *otg.IsisIIHNeighborRestartSucceded) (IsisIIHNeighborRestartSucceded, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHNeighborRestartSucceded) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHNeighborRestartSucceded) FromPbText(value string) error {
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

func (m *marshalisisIIHNeighborRestartSucceded) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHNeighborRestartSucceded) FromYaml(value string) error {
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

func (m *marshalisisIIHNeighborRestartSucceded) ToJson() (string, error) {
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

func (m *unMarshalisisIIHNeighborRestartSucceded) FromJson(value string) error {
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

func (obj *isisIIHNeighborRestartSucceded) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHNeighborRestartSucceded) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHNeighborRestartSucceded) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHNeighborRestartSucceded) Clone() (IsisIIHNeighborRestartSucceded, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHNeighborRestartSucceded()
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

// IsisIIHNeighborRestartSucceded is when neighbor router is Restarting/Starting the Graceful Restart status is successful.
type IsisIIHNeighborRestartSucceded interface {
	Validation
	// msg marshals IsisIIHNeighborRestartSucceded to protobuf object *otg.IsisIIHNeighborRestartSucceded
	// and doesn't set defaults
	msg() *otg.IsisIIHNeighborRestartSucceded
	// setMsg unmarshals IsisIIHNeighborRestartSucceded from protobuf object *otg.IsisIIHNeighborRestartSucceded
	// and doesn't set defaults
	setMsg(*otg.IsisIIHNeighborRestartSucceded) IsisIIHNeighborRestartSucceded
	// provides marshal interface
	Marshal() marshalIsisIIHNeighborRestartSucceded
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHNeighborRestartSucceded
	// validate validates IsisIIHNeighborRestartSucceded
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHNeighborRestartSucceded, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsdbSyncupTime returns uint32, set in IsisIIHNeighborRestartSucceded.
	LsdbSyncupTime() uint32
	// SetLsdbSyncupTime assigns uint32 provided by user to IsisIIHNeighborRestartSucceded
	SetLsdbSyncupTime(value uint32) IsisIIHNeighborRestartSucceded
	// HasLsdbSyncupTime checks if LsdbSyncupTime has been set in IsisIIHNeighborRestartSucceded
	HasLsdbSyncupTime() bool
	// LastAdjacencyBringUpTime returns uint32, set in IsisIIHNeighborRestartSucceded.
	LastAdjacencyBringUpTime() uint32
	// SetLastAdjacencyBringUpTime assigns uint32 provided by user to IsisIIHNeighborRestartSucceded
	SetLastAdjacencyBringUpTime(value uint32) IsisIIHNeighborRestartSucceded
	// HasLastAdjacencyBringUpTime checks if LastAdjacencyBringUpTime has been set in IsisIIHNeighborRestartSucceded
	HasLastAdjacencyBringUpTime() bool
}

// The time in second is taken to sync up the LSDB by the Restarting Neighbor. The time is counted by a Helper router when the first Restart TLV with RR bit is unset after  the recent RR bit set in the previous IIH PDU and till it receives the first LSP(s) from the Neighbor Router.
// LsdbSyncupTime returns a uint32
func (obj *isisIIHNeighborRestartSucceded) LsdbSyncupTime() uint32 {

	return *obj.obj.LsdbSyncupTime

}

// The time in second is taken to sync up the LSDB by the Restarting Neighbor. The time is counted by a Helper router when the first Restart TLV with RR bit is unset after  the recent RR bit set in the previous IIH PDU and till it receives the first LSP(s) from the Neighbor Router.
// LsdbSyncupTime returns a uint32
func (obj *isisIIHNeighborRestartSucceded) HasLsdbSyncupTime() bool {
	return obj.obj.LsdbSyncupTime != nil
}

// The time in second is taken to sync up the LSDB by the Restarting Neighbor. The time is counted by a Helper router when the first Restart TLV with RR bit is unset after  the recent RR bit set in the previous IIH PDU and till it receives the first LSP(s) from the Neighbor Router.
// SetLsdbSyncupTime sets the uint32 value in the IsisIIHNeighborRestartSucceded object
func (obj *isisIIHNeighborRestartSucceded) SetLsdbSyncupTime(value uint32) IsisIIHNeighborRestartSucceded {

	obj.obj.LsdbSyncupTime = &value
	return obj
}

// The time (in second) is taken to bring up adjacency by the Restarting Neighbor. The time is counted by a Helper router when the Restart TLV with RR bit set in a IIH PDU is received and till it receives the Restart TLV with RR bit unset in a IIH PDU from the Neighbor Router.
// LastAdjacencyBringUpTime returns a uint32
func (obj *isisIIHNeighborRestartSucceded) LastAdjacencyBringUpTime() uint32 {

	return *obj.obj.LastAdjacencyBringUpTime

}

// The time (in second) is taken to bring up adjacency by the Restarting Neighbor. The time is counted by a Helper router when the Restart TLV with RR bit set in a IIH PDU is received and till it receives the Restart TLV with RR bit unset in a IIH PDU from the Neighbor Router.
// LastAdjacencyBringUpTime returns a uint32
func (obj *isisIIHNeighborRestartSucceded) HasLastAdjacencyBringUpTime() bool {
	return obj.obj.LastAdjacencyBringUpTime != nil
}

// The time (in second) is taken to bring up adjacency by the Restarting Neighbor. The time is counted by a Helper router when the Restart TLV with RR bit set in a IIH PDU is received and till it receives the Restart TLV with RR bit unset in a IIH PDU from the Neighbor Router.
// SetLastAdjacencyBringUpTime sets the uint32 value in the IsisIIHNeighborRestartSucceded object
func (obj *isisIIHNeighborRestartSucceded) SetLastAdjacencyBringUpTime(value uint32) IsisIIHNeighborRestartSucceded {

	obj.obj.LastAdjacencyBringUpTime = &value
	return obj
}

func (obj *isisIIHNeighborRestartSucceded) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHNeighborRestartSucceded) setDefault() {

}
