package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspPrefixSids *****
type isisLspPrefixSids struct {
	validation
	obj          *otg.IsisLspPrefixSids
	marshaller   marshalIsisLspPrefixSids
	unMarshaller unMarshalIsisLspPrefixSids
}

func NewIsisLspPrefixSids() IsisLspPrefixSids {
	obj := isisLspPrefixSids{obj: &otg.IsisLspPrefixSids{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspPrefixSids) msg() *otg.IsisLspPrefixSids {
	return obj.obj
}

func (obj *isisLspPrefixSids) setMsg(msg *otg.IsisLspPrefixSids) IsisLspPrefixSids {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspPrefixSids struct {
	obj *isisLspPrefixSids
}

type marshalIsisLspPrefixSids interface {
	// ToProto marshals IsisLspPrefixSids to protobuf object *otg.IsisLspPrefixSids
	ToProto() (*otg.IsisLspPrefixSids, error)
	// ToPbText marshals IsisLspPrefixSids to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspPrefixSids to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspPrefixSids to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspPrefixSids struct {
	obj *isisLspPrefixSids
}

type unMarshalIsisLspPrefixSids interface {
	// FromProto unmarshals IsisLspPrefixSids from protobuf object *otg.IsisLspPrefixSids
	FromProto(msg *otg.IsisLspPrefixSids) (IsisLspPrefixSids, error)
	// FromPbText unmarshals IsisLspPrefixSids from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspPrefixSids from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspPrefixSids from JSON text
	FromJson(value string) error
}

func (obj *isisLspPrefixSids) Marshal() marshalIsisLspPrefixSids {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspPrefixSids{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspPrefixSids) Unmarshal() unMarshalIsisLspPrefixSids {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspPrefixSids{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspPrefixSids) ToProto() (*otg.IsisLspPrefixSids, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspPrefixSids) FromProto(msg *otg.IsisLspPrefixSids) (IsisLspPrefixSids, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspPrefixSids) ToPbText() (string, error) {
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

func (m *unMarshalisisLspPrefixSids) FromPbText(value string) error {
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

func (m *marshalisisLspPrefixSids) ToYaml() (string, error) {
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

func (m *unMarshalisisLspPrefixSids) FromYaml(value string) error {
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

func (m *marshalisisLspPrefixSids) ToJson() (string, error) {
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

func (m *unMarshalisisLspPrefixSids) FromJson(value string) error {
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

func (obj *isisLspPrefixSids) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspPrefixSids) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspPrefixSids) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspPrefixSids) Clone() (IsisLspPrefixSids, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspPrefixSids()
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

// IsisLspPrefixSids is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Refernce: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
type IsisLspPrefixSids interface {
	Validation
	// msg marshals IsisLspPrefixSids to protobuf object *otg.IsisLspPrefixSids
	// and doesn't set defaults
	msg() *otg.IsisLspPrefixSids
	// setMsg unmarshals IsisLspPrefixSids from protobuf object *otg.IsisLspPrefixSids
	// and doesn't set defaults
	setMsg(*otg.IsisLspPrefixSids) IsisLspPrefixSids
	// provides marshal interface
	Marshal() marshalIsisLspPrefixSids
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspPrefixSids
	// validate validates IsisLspPrefixSids
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspPrefixSids, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sid returns uint32, set in IsisLspPrefixSids.
	Sid() uint32
	// SetSid assigns uint32 provided by user to IsisLspPrefixSids
	SetSid(value uint32) IsisLspPrefixSids
	// HasSid checks if Sid has been set in IsisLspPrefixSids
	HasSid() bool
	// RFlag returns bool, set in IsisLspPrefixSids.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisLspPrefixSids
	SetRFlag(value bool) IsisLspPrefixSids
	// HasRFlag checks if RFlag has been set in IsisLspPrefixSids
	HasRFlag() bool
	// NFlag returns bool, set in IsisLspPrefixSids.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisLspPrefixSids
	SetNFlag(value bool) IsisLspPrefixSids
	// HasNFlag checks if NFlag has been set in IsisLspPrefixSids
	HasNFlag() bool
	// PFlag returns bool, set in IsisLspPrefixSids.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisLspPrefixSids
	SetPFlag(value bool) IsisLspPrefixSids
	// HasPFlag checks if PFlag has been set in IsisLspPrefixSids
	HasPFlag() bool
	// EFlag returns bool, set in IsisLspPrefixSids.
	EFlag() bool
	// SetEFlag assigns bool provided by user to IsisLspPrefixSids
	SetEFlag(value bool) IsisLspPrefixSids
	// HasEFlag checks if EFlag has been set in IsisLspPrefixSids
	HasEFlag() bool
	// VFlag returns bool, set in IsisLspPrefixSids.
	VFlag() bool
	// SetVFlag assigns bool provided by user to IsisLspPrefixSids
	SetVFlag(value bool) IsisLspPrefixSids
	// HasVFlag checks if VFlag has been set in IsisLspPrefixSids
	HasVFlag() bool
	// LFlag returns bool, set in IsisLspPrefixSids.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisLspPrefixSids
	SetLFlag(value bool) IsisLspPrefixSids
	// HasLFlag checks if LFlag has been set in IsisLspPrefixSids
	HasLFlag() bool
	// Algorithm returns uint32, set in IsisLspPrefixSids.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisLspPrefixSids
	SetAlgorithm(value uint32) IsisLspPrefixSids
	// HasAlgorithm checks if Algorithm has been set in IsisLspPrefixSids
	HasAlgorithm() bool
}

// SID/Index is the SID/Label value associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// Sid returns a uint32
func (obj *isisLspPrefixSids) Sid() uint32 {

	return *obj.obj.Sid

}

// SID/Index is the SID/Label value associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// Sid returns a uint32
func (obj *isisLspPrefixSids) HasSid() bool {
	return obj.obj.Sid != nil
}

// SID/Index is the SID/Label value associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSid sets the uint32 value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetSid(value uint32) IsisLspPrefixSids {

	obj.obj.Sid = &value
	return obj
}

// R-Flag:Re-advertisement flag
// RFlag returns a bool
func (obj *isisLspPrefixSids) RFlag() bool {

	return *obj.obj.RFlag

}

// R-Flag:Re-advertisement flag
// RFlag returns a bool
func (obj *isisLspPrefixSids) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// R-Flag:Re-advertisement flag
// SetRFlag sets the bool value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetRFlag(value bool) IsisLspPrefixSids {

	obj.obj.RFlag = &value
	return obj
}

// N-Flag: Node-SID flag.  If set, then the Prefix-SID refers to
// the router identified by the prefix
// NFlag returns a bool
func (obj *isisLspPrefixSids) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag: Node-SID flag.  If set, then the Prefix-SID refers to
// the router identified by the prefix
// NFlag returns a bool
func (obj *isisLspPrefixSids) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag: Node-SID flag.  If set, then the Prefix-SID refers to
// the router identified by the prefix
// SetNFlag sets the bool value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetNFlag(value bool) IsisLspPrefixSids {

	obj.obj.NFlag = &value
	return obj
}

// P-Flag: no-PHP flag.
// PFlag returns a bool
func (obj *isisLspPrefixSids) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag: no-PHP flag.
// PFlag returns a bool
func (obj *isisLspPrefixSids) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag: no-PHP flag.
// SetPFlag sets the bool value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetPFlag(value bool) IsisLspPrefixSids {

	obj.obj.PFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// EFlag returns a bool
func (obj *isisLspPrefixSids) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// EFlag returns a bool
func (obj *isisLspPrefixSids) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// SetEFlag sets the bool value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetEFlag(value bool) IsisLspPrefixSids {

	obj.obj.EFlag = &value
	return obj
}

// The value flag. If set, then the Prefix-SID carries a value instead of an index.
// VFlag returns a bool
func (obj *isisLspPrefixSids) VFlag() bool {

	return *obj.obj.VFlag

}

// The value flag. If set, then the Prefix-SID carries a value instead of an index.
// VFlag returns a bool
func (obj *isisLspPrefixSids) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// The value flag. If set, then the Prefix-SID carries a value instead of an index.
// SetVFlag sets the bool value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetVFlag(value bool) IsisLspPrefixSids {

	obj.obj.VFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *isisLspPrefixSids) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *isisLspPrefixSids) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance.
// SetLFlag sets the bool value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetLFlag(value bool) IsisLspPrefixSids {

	obj.obj.LFlag = &value
	return obj
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisLspPrefixSids) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisLspPrefixSids) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// SetAlgorithm sets the uint32 value in the IsisLspPrefixSids object
func (obj *isisLspPrefixSids) SetAlgorithm(value uint32) IsisLspPrefixSids {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisLspPrefixSids) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspPrefixSids.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *isisLspPrefixSids) setDefault() {

}
