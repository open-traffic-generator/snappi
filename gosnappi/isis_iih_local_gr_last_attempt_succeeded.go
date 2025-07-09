package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisIIHLocalGRLastAttemptSucceeded *****
type isisIIHLocalGRLastAttemptSucceeded struct {
	validation
	obj          *otg.IsisIIHLocalGRLastAttemptSucceeded
	marshaller   marshalIsisIIHLocalGRLastAttemptSucceeded
	unMarshaller unMarshalIsisIIHLocalGRLastAttemptSucceeded
}

func NewIsisIIHLocalGRLastAttemptSucceeded() IsisIIHLocalGRLastAttemptSucceeded {
	obj := isisIIHLocalGRLastAttemptSucceeded{obj: &otg.IsisIIHLocalGRLastAttemptSucceeded{}}
	obj.setDefault()
	return &obj
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) msg() *otg.IsisIIHLocalGRLastAttemptSucceeded {
	return obj.obj
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) setMsg(msg *otg.IsisIIHLocalGRLastAttemptSucceeded) IsisIIHLocalGRLastAttemptSucceeded {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisIIHLocalGRLastAttemptSucceeded struct {
	obj *isisIIHLocalGRLastAttemptSucceeded
}

type marshalIsisIIHLocalGRLastAttemptSucceeded interface {
	// ToProto marshals IsisIIHLocalGRLastAttemptSucceeded to protobuf object *otg.IsisIIHLocalGRLastAttemptSucceeded
	ToProto() (*otg.IsisIIHLocalGRLastAttemptSucceeded, error)
	// ToPbText marshals IsisIIHLocalGRLastAttemptSucceeded to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisIIHLocalGRLastAttemptSucceeded to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisIIHLocalGRLastAttemptSucceeded to JSON text
	ToJson() (string, error)
}

type unMarshalisisIIHLocalGRLastAttemptSucceeded struct {
	obj *isisIIHLocalGRLastAttemptSucceeded
}

type unMarshalIsisIIHLocalGRLastAttemptSucceeded interface {
	// FromProto unmarshals IsisIIHLocalGRLastAttemptSucceeded from protobuf object *otg.IsisIIHLocalGRLastAttemptSucceeded
	FromProto(msg *otg.IsisIIHLocalGRLastAttemptSucceeded) (IsisIIHLocalGRLastAttemptSucceeded, error)
	// FromPbText unmarshals IsisIIHLocalGRLastAttemptSucceeded from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisIIHLocalGRLastAttemptSucceeded from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisIIHLocalGRLastAttemptSucceeded from JSON text
	FromJson(value string) error
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) Marshal() marshalIsisIIHLocalGRLastAttemptSucceeded {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisIIHLocalGRLastAttemptSucceeded{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) Unmarshal() unMarshalIsisIIHLocalGRLastAttemptSucceeded {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisIIHLocalGRLastAttemptSucceeded{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisIIHLocalGRLastAttemptSucceeded) ToProto() (*otg.IsisIIHLocalGRLastAttemptSucceeded, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisIIHLocalGRLastAttemptSucceeded) FromProto(msg *otg.IsisIIHLocalGRLastAttemptSucceeded) (IsisIIHLocalGRLastAttemptSucceeded, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisIIHLocalGRLastAttemptSucceeded) ToPbText() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptSucceeded) FromPbText(value string) error {
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

func (m *marshalisisIIHLocalGRLastAttemptSucceeded) ToYaml() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptSucceeded) FromYaml(value string) error {
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

func (m *marshalisisIIHLocalGRLastAttemptSucceeded) ToJson() (string, error) {
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

func (m *unMarshalisisIIHLocalGRLastAttemptSucceeded) FromJson(value string) error {
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

func (obj *isisIIHLocalGRLastAttemptSucceeded) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) Clone() (IsisIIHLocalGRLastAttemptSucceeded, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisIIHLocalGRLastAttemptSucceeded()
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

// IsisIIHLocalGRLastAttemptSucceeded is this container contains details about the successful status of the last Graceful Restart initiated by this router.
type IsisIIHLocalGRLastAttemptSucceeded interface {
	Validation
	// msg marshals IsisIIHLocalGRLastAttemptSucceeded to protobuf object *otg.IsisIIHLocalGRLastAttemptSucceeded
	// and doesn't set defaults
	msg() *otg.IsisIIHLocalGRLastAttemptSucceeded
	// setMsg unmarshals IsisIIHLocalGRLastAttemptSucceeded from protobuf object *otg.IsisIIHLocalGRLastAttemptSucceeded
	// and doesn't set defaults
	setMsg(*otg.IsisIIHLocalGRLastAttemptSucceeded) IsisIIHLocalGRLastAttemptSucceeded
	// provides marshal interface
	Marshal() marshalIsisIIHLocalGRLastAttemptSucceeded
	// provides unmarshal interface
	Unmarshal() unMarshalIsisIIHLocalGRLastAttemptSucceeded
	// validate validates IsisIIHLocalGRLastAttemptSucceeded
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisIIHLocalGRLastAttemptSucceeded, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsdbSyncupTime returns uint32, set in IsisIIHLocalGRLastAttemptSucceeded.
	LsdbSyncupTime() uint32
	// SetLsdbSyncupTime assigns uint32 provided by user to IsisIIHLocalGRLastAttemptSucceeded
	SetLsdbSyncupTime(value uint32) IsisIIHLocalGRLastAttemptSucceeded
	// HasLsdbSyncupTime checks if LsdbSyncupTime has been set in IsisIIHLocalGRLastAttemptSucceeded
	HasLsdbSyncupTime() bool
	// AdjacencyBringUpTime returns uint32, set in IsisIIHLocalGRLastAttemptSucceeded.
	AdjacencyBringUpTime() uint32
	// SetAdjacencyBringUpTime assigns uint32 provided by user to IsisIIHLocalGRLastAttemptSucceeded
	SetAdjacencyBringUpTime(value uint32) IsisIIHLocalGRLastAttemptSucceeded
	// HasAdjacencyBringUpTime checks if AdjacencyBringUpTime has been set in IsisIIHLocalGRLastAttemptSucceeded
	HasAdjacencyBringUpTime() bool
}

// The time (in seconds) is taken to synchronize the L1 and L2 LSDB by this Restarting router.
// By this time, the CSNP list is cleared and all LSPs are collected by the neighbor(s).
// LsdbSyncupTime returns a uint32
func (obj *isisIIHLocalGRLastAttemptSucceeded) LsdbSyncupTime() uint32 {

	return *obj.obj.LsdbSyncupTime

}

// The time (in seconds) is taken to synchronize the L1 and L2 LSDB by this Restarting router.
// By this time, the CSNP list is cleared and all LSPs are collected by the neighbor(s).
// LsdbSyncupTime returns a uint32
func (obj *isisIIHLocalGRLastAttemptSucceeded) HasLsdbSyncupTime() bool {
	return obj.obj.LsdbSyncupTime != nil
}

// The time (in seconds) is taken to synchronize the L1 and L2 LSDB by this Restarting router.
// By this time, the CSNP list is cleared and all LSPs are collected by the neighbor(s).
// SetLsdbSyncupTime sets the uint32 value in the IsisIIHLocalGRLastAttemptSucceeded object
func (obj *isisIIHLocalGRLastAttemptSucceeded) SetLsdbSyncupTime(value uint32) IsisIIHLocalGRLastAttemptSucceeded {

	obj.obj.LsdbSyncupTime = &value
	return obj
}

// The time (in seconds) is measured from when the Restart TLV with RR bit set is sent
// in an IIH PDU upto the time when Restart TLV is sent with RR bit unset.
// AdjacencyBringUpTime returns a uint32
func (obj *isisIIHLocalGRLastAttemptSucceeded) AdjacencyBringUpTime() uint32 {

	return *obj.obj.AdjacencyBringUpTime

}

// The time (in seconds) is measured from when the Restart TLV with RR bit set is sent
// in an IIH PDU upto the time when Restart TLV is sent with RR bit unset.
// AdjacencyBringUpTime returns a uint32
func (obj *isisIIHLocalGRLastAttemptSucceeded) HasAdjacencyBringUpTime() bool {
	return obj.obj.AdjacencyBringUpTime != nil
}

// The time (in seconds) is measured from when the Restart TLV with RR bit set is sent
// in an IIH PDU upto the time when Restart TLV is sent with RR bit unset.
// SetAdjacencyBringUpTime sets the uint32 value in the IsisIIHLocalGRLastAttemptSucceeded object
func (obj *isisIIHLocalGRLastAttemptSucceeded) SetAdjacencyBringUpTime(value uint32) IsisIIHLocalGRLastAttemptSucceeded {

	obj.obj.AdjacencyBringUpTime = &value
	return obj
}

func (obj *isisIIHLocalGRLastAttemptSucceeded) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisIIHLocalGRLastAttemptSucceeded) setDefault() {

}
