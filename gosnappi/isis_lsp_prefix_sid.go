package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspPrefixSid *****
type isisLspPrefixSid struct {
	validation
	obj          *otg.IsisLspPrefixSid
	marshaller   marshalIsisLspPrefixSid
	unMarshaller unMarshalIsisLspPrefixSid
	flagsHolder  IsisLspPrefixSidFlags
}

func NewIsisLspPrefixSid() IsisLspPrefixSid {
	obj := isisLspPrefixSid{obj: &otg.IsisLspPrefixSid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspPrefixSid) msg() *otg.IsisLspPrefixSid {
	return obj.obj
}

func (obj *isisLspPrefixSid) setMsg(msg *otg.IsisLspPrefixSid) IsisLspPrefixSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspPrefixSid struct {
	obj *isisLspPrefixSid
}

type marshalIsisLspPrefixSid interface {
	// ToProto marshals IsisLspPrefixSid to protobuf object *otg.IsisLspPrefixSid
	ToProto() (*otg.IsisLspPrefixSid, error)
	// ToPbText marshals IsisLspPrefixSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspPrefixSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspPrefixSid to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspPrefixSid to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspPrefixSid struct {
	obj *isisLspPrefixSid
}

type unMarshalIsisLspPrefixSid interface {
	// FromProto unmarshals IsisLspPrefixSid from protobuf object *otg.IsisLspPrefixSid
	FromProto(msg *otg.IsisLspPrefixSid) (IsisLspPrefixSid, error)
	// FromPbText unmarshals IsisLspPrefixSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspPrefixSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspPrefixSid from JSON text
	FromJson(value string) error
}

func (obj *isisLspPrefixSid) Marshal() marshalIsisLspPrefixSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspPrefixSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspPrefixSid) Unmarshal() unMarshalIsisLspPrefixSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspPrefixSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspPrefixSid) ToProto() (*otg.IsisLspPrefixSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspPrefixSid) FromProto(msg *otg.IsisLspPrefixSid) (IsisLspPrefixSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspPrefixSid) ToPbText() (string, error) {
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

func (m *unMarshalisisLspPrefixSid) FromPbText(value string) error {
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

func (m *marshalisisLspPrefixSid) ToYaml() (string, error) {
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

func (m *unMarshalisisLspPrefixSid) FromYaml(value string) error {
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

func (m *marshalisisLspPrefixSid) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalisisLspPrefixSid) ToJson() (string, error) {
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

func (m *unMarshalisisLspPrefixSid) FromJson(value string) error {
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

func (obj *isisLspPrefixSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspPrefixSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspPrefixSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspPrefixSid) Clone() (IsisLspPrefixSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspPrefixSid()
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

func (obj *isisLspPrefixSid) setNil() {
	obj.flagsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspPrefixSid is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Refernce: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
type IsisLspPrefixSid interface {
	Validation
	// msg marshals IsisLspPrefixSid to protobuf object *otg.IsisLspPrefixSid
	// and doesn't set defaults
	msg() *otg.IsisLspPrefixSid
	// setMsg unmarshals IsisLspPrefixSid from protobuf object *otg.IsisLspPrefixSid
	// and doesn't set defaults
	setMsg(*otg.IsisLspPrefixSid) IsisLspPrefixSid
	// provides marshal interface
	Marshal() marshalIsisLspPrefixSid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspPrefixSid
	// validate validates IsisLspPrefixSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspPrefixSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sids returns []uint32, set in IsisLspPrefixSid.
	Sids() []uint32
	// SetSids assigns []uint32 provided by user to IsisLspPrefixSid
	SetSids(value []uint32) IsisLspPrefixSid
	// Flags returns IsisLspPrefixSidFlags, set in IsisLspPrefixSid.
	// IsisLspPrefixSidFlags is conatiner of 1-octet Flags associated with Prefix Segment-ID.
	Flags() IsisLspPrefixSidFlags
	// SetFlags assigns IsisLspPrefixSidFlags provided by user to IsisLspPrefixSid.
	// IsisLspPrefixSidFlags is conatiner of 1-octet Flags associated with Prefix Segment-ID.
	SetFlags(value IsisLspPrefixSidFlags) IsisLspPrefixSid
	// HasFlags checks if Flags has been set in IsisLspPrefixSid
	HasFlags() bool
	// Algorithm returns uint32, set in IsisLspPrefixSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisLspPrefixSid
	SetAlgorithm(value uint32) IsisLspPrefixSid
	// HasAlgorithm checks if Algorithm has been set in IsisLspPrefixSid
	HasAlgorithm() bool
	setNil()
}

// One or more SIDs/Indices are the SID/Label values associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// Sids returns a []uint32
func (obj *isisLspPrefixSid) Sids() []uint32 {
	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	return obj.obj.Sids
}

// One or more SIDs/Indices are the SID/Label values associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSids sets the []uint32 value in the IsisLspPrefixSid object
func (obj *isisLspPrefixSid) SetSids(value []uint32) IsisLspPrefixSid {

	if obj.obj.Sids == nil {
		obj.obj.Sids = make([]uint32, 0)
	}
	obj.obj.Sids = value

	return obj
}

// Flags associated with Prefix Segment-ID.
// Flags returns a IsisLspPrefixSidFlags
func (obj *isisLspPrefixSid) Flags() IsisLspPrefixSidFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewIsisLspPrefixSidFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &isisLspPrefixSidFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// Flags associated with Prefix Segment-ID.
// Flags returns a IsisLspPrefixSidFlags
func (obj *isisLspPrefixSid) HasFlags() bool {
	return obj.obj.Flags != nil
}

// Flags associated with Prefix Segment-ID.
// SetFlags sets the IsisLspPrefixSidFlags value in the IsisLspPrefixSid object
func (obj *isisLspPrefixSid) SetFlags(value IsisLspPrefixSidFlags) IsisLspPrefixSid {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisLspPrefixSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisLspPrefixSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// SetAlgorithm sets the uint32 value in the IsisLspPrefixSid object
func (obj *isisLspPrefixSid) SetAlgorithm(value uint32) IsisLspPrefixSid {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisLspPrefixSid) validateObj(vObj *validation, set_default bool) {
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
				fmt.Sprintf("0 <= IsisLspPrefixSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *isisLspPrefixSid) setDefault() {

}
