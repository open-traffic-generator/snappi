package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaPrefixSid *****
type ospfv2LsaPrefixSid struct {
	validation
	obj          *otg.Ospfv2LsaPrefixSid
	marshaller   marshalOspfv2LsaPrefixSid
	unMarshaller unMarshalOspfv2LsaPrefixSid
	flagsHolder  Ospfv2LsaPrefixSidFlags
}

func NewOspfv2LsaPrefixSid() Ospfv2LsaPrefixSid {
	obj := ospfv2LsaPrefixSid{obj: &otg.Ospfv2LsaPrefixSid{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaPrefixSid) msg() *otg.Ospfv2LsaPrefixSid {
	return obj.obj
}

func (obj *ospfv2LsaPrefixSid) setMsg(msg *otg.Ospfv2LsaPrefixSid) Ospfv2LsaPrefixSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaPrefixSid struct {
	obj *ospfv2LsaPrefixSid
}

type marshalOspfv2LsaPrefixSid interface {
	// ToProto marshals Ospfv2LsaPrefixSid to protobuf object *otg.Ospfv2LsaPrefixSid
	ToProto() (*otg.Ospfv2LsaPrefixSid, error)
	// ToPbText marshals Ospfv2LsaPrefixSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaPrefixSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaPrefixSid to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaPrefixSid struct {
	obj *ospfv2LsaPrefixSid
}

type unMarshalOspfv2LsaPrefixSid interface {
	// FromProto unmarshals Ospfv2LsaPrefixSid from protobuf object *otg.Ospfv2LsaPrefixSid
	FromProto(msg *otg.Ospfv2LsaPrefixSid) (Ospfv2LsaPrefixSid, error)
	// FromPbText unmarshals Ospfv2LsaPrefixSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaPrefixSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaPrefixSid from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaPrefixSid) Marshal() marshalOspfv2LsaPrefixSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaPrefixSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaPrefixSid) Unmarshal() unMarshalOspfv2LsaPrefixSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaPrefixSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaPrefixSid) ToProto() (*otg.Ospfv2LsaPrefixSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaPrefixSid) FromProto(msg *otg.Ospfv2LsaPrefixSid) (Ospfv2LsaPrefixSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaPrefixSid) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaPrefixSid) FromPbText(value string) error {
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

func (m *marshalospfv2LsaPrefixSid) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaPrefixSid) FromYaml(value string) error {
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

func (m *marshalospfv2LsaPrefixSid) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaPrefixSid) FromJson(value string) error {
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

func (obj *ospfv2LsaPrefixSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaPrefixSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaPrefixSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaPrefixSid) Clone() (Ospfv2LsaPrefixSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaPrefixSid()
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

func (obj *ospfv2LsaPrefixSid) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2LsaPrefixSid is the learned OSPFv2 Prefix-SID and its attributes, decoded from the Prefix-SID sub-TLV of
// the Extended Prefix Opaque LSA (RFC 8665).
type Ospfv2LsaPrefixSid interface {
	Validation
	// msg marshals Ospfv2LsaPrefixSid to protobuf object *otg.Ospfv2LsaPrefixSid
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaPrefixSid
	// setMsg unmarshals Ospfv2LsaPrefixSid from protobuf object *otg.Ospfv2LsaPrefixSid
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaPrefixSid) Ospfv2LsaPrefixSid
	// provides marshal interface
	Marshal() marshalOspfv2LsaPrefixSid
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaPrefixSid
	// validate validates Ospfv2LsaPrefixSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaPrefixSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sids returns []uint32, set in Ospfv2LsaPrefixSid.
	Sids() []uint32
	// SetSids assigns []uint32 provided by user to Ospfv2LsaPrefixSid
	SetSids(value []uint32) Ospfv2LsaPrefixSid
	// Flags returns Ospfv2LsaPrefixSidFlags, set in Ospfv2LsaPrefixSid.
	// Ospfv2LsaPrefixSidFlags is one-octet flags of the OSPFv2 Prefix-SID sub-TLV (RFC 8665).
	Flags() Ospfv2LsaPrefixSidFlags
	// SetFlags assigns Ospfv2LsaPrefixSidFlags provided by user to Ospfv2LsaPrefixSid.
	// Ospfv2LsaPrefixSidFlags is one-octet flags of the OSPFv2 Prefix-SID sub-TLV (RFC 8665).
	SetFlags(value Ospfv2LsaPrefixSidFlags) Ospfv2LsaPrefixSid
	// HasFlags checks if Flags has been set in Ospfv2LsaPrefixSid
	HasFlags() bool
	// Algorithm returns uint32, set in Ospfv2LsaPrefixSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to Ospfv2LsaPrefixSid
	SetAlgorithm(value uint32) Ospfv2LsaPrefixSid
	// HasAlgorithm checks if Algorithm has been set in Ospfv2LsaPrefixSid
	HasAlgorithm() bool
	setNil()
}

// One or more SID/Label values or indices associated with the IGP Prefix segment attached to the prefix.
// Sids returns a []uint32
func (obj *ospfv2LsaPrefixSid) Sids() []uint32 {
	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	return obj.obj.Sids
}

// One or more SID/Label values or indices associated with the IGP Prefix segment attached to the prefix.
// SetSids sets the []uint32 value in the Ospfv2LsaPrefixSid object
func (obj *ospfv2LsaPrefixSid) SetSids(value []uint32) Ospfv2LsaPrefixSid {

	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	obj.obj.Sids = value

	return obj
}

// Flags associated with the Prefix-SID.
// Flags returns a Ospfv2LsaPrefixSidFlags
func (obj *ospfv2LsaPrefixSid) Flags() Ospfv2LsaPrefixSidFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewOspfv2LsaPrefixSidFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &ospfv2LsaPrefixSidFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// Flags associated with the Prefix-SID.
// Flags returns a Ospfv2LsaPrefixSidFlags
func (obj *ospfv2LsaPrefixSid) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Flags associated with the Prefix-SID.
// SetFlags sets the Ospfv2LsaPrefixSidFlags value in the Ospfv2LsaPrefixSid object
func (obj *ospfv2LsaPrefixSid) SetFlags(value Ospfv2LsaPrefixSidFlags) Ospfv2LsaPrefixSid {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The Segment Routing algorithm the Prefix-SID is associated with.
// Algorithm returns a uint32
func (obj *ospfv2LsaPrefixSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Segment Routing algorithm the Prefix-SID is associated with.
// Algorithm returns a uint32
func (obj *ospfv2LsaPrefixSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Segment Routing algorithm the Prefix-SID is associated with.
// SetAlgorithm sets the uint32 value in the Ospfv2LsaPrefixSid object
func (obj *ospfv2LsaPrefixSid) SetAlgorithm(value uint32) Ospfv2LsaPrefixSid {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *ospfv2LsaPrefixSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Ospfv2LsaPrefixSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *ospfv2LsaPrefixSid) setDefault() {

}
