package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspAdjacencySid *****
type isisLspAdjacencySid struct {
	validation
	obj          *otg.IsisLspAdjacencySid
	marshaller   marshalIsisLspAdjacencySid
	unMarshaller unMarshalIsisLspAdjacencySid
	flagsHolder  IsisLspAdjSidFlags
}

func NewIsisLspAdjacencySid() IsisLspAdjacencySid {
	obj := isisLspAdjacencySid{obj: &otg.IsisLspAdjacencySid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspAdjacencySid) msg() *otg.IsisLspAdjacencySid {
	return obj.obj
}

func (obj *isisLspAdjacencySid) setMsg(msg *otg.IsisLspAdjacencySid) IsisLspAdjacencySid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspAdjacencySid struct {
	obj *isisLspAdjacencySid
}

type marshalIsisLspAdjacencySid interface {
	// ToProto marshals IsisLspAdjacencySid to protobuf object *otg.IsisLspAdjacencySid
	ToProto() (*otg.IsisLspAdjacencySid, error)
	// ToPbText marshals IsisLspAdjacencySid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspAdjacencySid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspAdjacencySid to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspAdjacencySid struct {
	obj *isisLspAdjacencySid
}

type unMarshalIsisLspAdjacencySid interface {
	// FromProto unmarshals IsisLspAdjacencySid from protobuf object *otg.IsisLspAdjacencySid
	FromProto(msg *otg.IsisLspAdjacencySid) (IsisLspAdjacencySid, error)
	// FromPbText unmarshals IsisLspAdjacencySid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspAdjacencySid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspAdjacencySid from JSON text
	FromJson(value string) error
}

func (obj *isisLspAdjacencySid) Marshal() marshalIsisLspAdjacencySid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspAdjacencySid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspAdjacencySid) Unmarshal() unMarshalIsisLspAdjacencySid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspAdjacencySid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspAdjacencySid) ToProto() (*otg.IsisLspAdjacencySid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspAdjacencySid) FromProto(msg *otg.IsisLspAdjacencySid) (IsisLspAdjacencySid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspAdjacencySid) ToPbText() (string, error) {
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

func (m *unMarshalisisLspAdjacencySid) FromPbText(value string) error {
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

func (m *marshalisisLspAdjacencySid) ToYaml() (string, error) {
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

func (m *unMarshalisisLspAdjacencySid) FromYaml(value string) error {
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

func (m *marshalisisLspAdjacencySid) ToJson() (string, error) {
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

func (m *unMarshalisisLspAdjacencySid) FromJson(value string) error {
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

func (obj *isisLspAdjacencySid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspAdjacencySid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspAdjacencySid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspAdjacencySid) Clone() (IsisLspAdjacencySid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspAdjacencySid()
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

func (obj *isisLspAdjacencySid) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspAdjacencySid is this container defines segment routing adjacency SIDs.
type IsisLspAdjacencySid interface {
	Validation
	// msg marshals IsisLspAdjacencySid to protobuf object *otg.IsisLspAdjacencySid
	// and doesn't set defaults
	msg() *otg.IsisLspAdjacencySid
	// setMsg unmarshals IsisLspAdjacencySid from protobuf object *otg.IsisLspAdjacencySid
	// and doesn't set defaults
	setMsg(*otg.IsisLspAdjacencySid) IsisLspAdjacencySid
	// provides marshal interface
	Marshal() marshalIsisLspAdjacencySid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspAdjacencySid
	// validate validates IsisLspAdjacencySid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspAdjacencySid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TlvType returns uint32, set in IsisLspAdjacencySid.
	TlvType() uint32
	// SetTlvType assigns uint32 provided by user to IsisLspAdjacencySid
	SetTlvType(value uint32) IsisLspAdjacencySid
	// HasTlvType checks if TlvType has been set in IsisLspAdjacencySid
	HasTlvType() bool
	// Sids returns []uint32, set in IsisLspAdjacencySid.
	Sids() []uint32
	// SetSids assigns []uint32 provided by user to IsisLspAdjacencySid
	SetSids(value []uint32) IsisLspAdjacencySid
	// Flags returns IsisLspAdjSidFlags, set in IsisLspAdjacencySid.
	// IsisLspAdjSidFlags is conatiner of 1-octet Flags associated with Adjacency Segment-ID.
	Flags() IsisLspAdjSidFlags
	// SetFlags assigns IsisLspAdjSidFlags provided by user to IsisLspAdjacencySid.
	// IsisLspAdjSidFlags is conatiner of 1-octet Flags associated with Adjacency Segment-ID.
	SetFlags(value IsisLspAdjSidFlags) IsisLspAdjacencySid
	// HasFlags checks if Flags has been set in IsisLspAdjacencySid
	HasFlags() bool
	// Weight returns uint32, set in IsisLspAdjacencySid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisLspAdjacencySid
	SetWeight(value uint32) IsisLspAdjacencySid
	// HasWeight checks if Weight has been set in IsisLspAdjacencySid
	HasWeight() bool
	setNil()
}

// The corresponding adjacency SID type.
// TlvType returns a uint32
func (obj *isisLspAdjacencySid) TlvType() uint32 {

	return *obj.obj.TlvType

}

// The corresponding adjacency SID type.
// TlvType returns a uint32
func (obj *isisLspAdjacencySid) HasTlvType() bool {
	return obj.obj.TlvType != nil
}

// The corresponding adjacency SID type.
// SetTlvType sets the uint32 value in the IsisLspAdjacencySid object
func (obj *isisLspAdjacencySid) SetTlvType(value uint32) IsisLspAdjacencySid {

	obj.obj.TlvType = &value
	return obj
}

// One or more SID/Indices are the SID/Label values associated with the IGP adjacency SID.
// Sids returns a []uint32
func (obj *isisLspAdjacencySid) Sids() []uint32 {
	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	return obj.obj.Sids
}

// One or more SID/Indices are the SID/Label values associated with the IGP adjacency SID.
// SetSids sets the []uint32 value in the IsisLspAdjacencySid object
func (obj *isisLspAdjacencySid) SetSids(value []uint32) IsisLspAdjacencySid {

	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	obj.obj.Sids = value

	return obj
}

// Flags associated with Adjacency Segment-ID.
// Flags returns a IsisLspAdjSidFlags
func (obj *isisLspAdjacencySid) Flags() IsisLspAdjSidFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewIsisLspAdjSidFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &isisLspAdjSidFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// Flags associated with Adjacency Segment-ID.
// Flags returns a IsisLspAdjSidFlags
func (obj *isisLspAdjacencySid) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Flags associated with Adjacency Segment-ID.
// SetFlags sets the IsisLspAdjSidFlags value in the IsisLspAdjacencySid object
func (obj *isisLspAdjacencySid) SetFlags(value IsisLspAdjSidFlags) IsisLspAdjacencySid {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *isisLspAdjacencySid) Weight() uint32 {

	return *obj.obj.Weight

}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// Weight returns a uint32
func (obj *isisLspAdjacencySid) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The value represents the weight of the Adj-SID for the purpose of load balancing.
// SetWeight sets the uint32 value in the IsisLspAdjacencySid object
func (obj *isisLspAdjacencySid) SetWeight(value uint32) IsisLspAdjacencySid {

	obj.obj.Weight = &value
	return obj
}

func (obj *isisLspAdjacencySid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

}

func (obj *isisLspAdjacencySid) setDefault() {

}
