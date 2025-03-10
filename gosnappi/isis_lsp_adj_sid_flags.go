package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspAdjSidFlags *****
type isisLspAdjSidFlags struct {
	validation
	obj          *otg.IsisLspAdjSidFlags
	marshaller   marshalIsisLspAdjSidFlags
	unMarshaller unMarshalIsisLspAdjSidFlags
}

func NewIsisLspAdjSidFlags() IsisLspAdjSidFlags {
	obj := isisLspAdjSidFlags{obj: &otg.IsisLspAdjSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspAdjSidFlags) msg() *otg.IsisLspAdjSidFlags {
	return obj.obj
}

func (obj *isisLspAdjSidFlags) setMsg(msg *otg.IsisLspAdjSidFlags) IsisLspAdjSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspAdjSidFlags struct {
	obj *isisLspAdjSidFlags
}

type marshalIsisLspAdjSidFlags interface {
	// ToProto marshals IsisLspAdjSidFlags to protobuf object *otg.IsisLspAdjSidFlags
	ToProto() (*otg.IsisLspAdjSidFlags, error)
	// ToPbText marshals IsisLspAdjSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspAdjSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspAdjSidFlags to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspAdjSidFlags to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspAdjSidFlags struct {
	obj *isisLspAdjSidFlags
}

type unMarshalIsisLspAdjSidFlags interface {
	// FromProto unmarshals IsisLspAdjSidFlags from protobuf object *otg.IsisLspAdjSidFlags
	FromProto(msg *otg.IsisLspAdjSidFlags) (IsisLspAdjSidFlags, error)
	// FromPbText unmarshals IsisLspAdjSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspAdjSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspAdjSidFlags from JSON text
	FromJson(value string) error
}

func (obj *isisLspAdjSidFlags) Marshal() marshalIsisLspAdjSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspAdjSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspAdjSidFlags) Unmarshal() unMarshalIsisLspAdjSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspAdjSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspAdjSidFlags) ToProto() (*otg.IsisLspAdjSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspAdjSidFlags) FromProto(msg *otg.IsisLspAdjSidFlags) (IsisLspAdjSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspAdjSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisLspAdjSidFlags) FromPbText(value string) error {
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

func (m *marshalisisLspAdjSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisLspAdjSidFlags) FromYaml(value string) error {
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

func (m *marshalisisLspAdjSidFlags) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspAdjSidFlags) ToJson() (string, error) {
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

func (m *unMarshalisisLspAdjSidFlags) FromJson(value string) error {
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

func (obj *isisLspAdjSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspAdjSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspAdjSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspAdjSidFlags) Clone() (IsisLspAdjSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspAdjSidFlags()
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

// IsisLspAdjSidFlags is conatiner of 1-octet Flags associated with Adjacency Segment-ID.
type IsisLspAdjSidFlags interface {
	Validation
	// msg marshals IsisLspAdjSidFlags to protobuf object *otg.IsisLspAdjSidFlags
	// and doesn't set defaults
	msg() *otg.IsisLspAdjSidFlags
	// setMsg unmarshals IsisLspAdjSidFlags from protobuf object *otg.IsisLspAdjSidFlags
	// and doesn't set defaults
	setMsg(*otg.IsisLspAdjSidFlags) IsisLspAdjSidFlags
	// provides marshal interface
	Marshal() marshalIsisLspAdjSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspAdjSidFlags
	// validate validates IsisLspAdjSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspAdjSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FFlag returns bool, set in IsisLspAdjSidFlags.
	FFlag() bool
	// SetFFlag assigns bool provided by user to IsisLspAdjSidFlags
	SetFFlag(value bool) IsisLspAdjSidFlags
	// HasFFlag checks if FFlag has been set in IsisLspAdjSidFlags
	HasFFlag() bool
	// BFlag returns bool, set in IsisLspAdjSidFlags.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisLspAdjSidFlags
	SetBFlag(value bool) IsisLspAdjSidFlags
	// HasBFlag checks if BFlag has been set in IsisLspAdjSidFlags
	HasBFlag() bool
	// VFlag returns bool, set in IsisLspAdjSidFlags.
	VFlag() bool
	// SetVFlag assigns bool provided by user to IsisLspAdjSidFlags
	SetVFlag(value bool) IsisLspAdjSidFlags
	// HasVFlag checks if VFlag has been set in IsisLspAdjSidFlags
	HasVFlag() bool
	// LFlag returns bool, set in IsisLspAdjSidFlags.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisLspAdjSidFlags
	SetLFlag(value bool) IsisLspAdjSidFlags
	// HasLFlag checks if LFlag has been set in IsisLspAdjSidFlags
	HasLFlag() bool
	// SFlag returns bool, set in IsisLspAdjSidFlags.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisLspAdjSidFlags
	SetSFlag(value bool) IsisLspAdjSidFlags
	// HasSFlag checks if SFlag has been set in IsisLspAdjSidFlags
	HasSFlag() bool
	// PFlag returns bool, set in IsisLspAdjSidFlags.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisLspAdjSidFlags
	SetPFlag(value bool) IsisLspAdjSidFlags
	// HasPFlag checks if PFlag has been set in IsisLspAdjSidFlags
	HasPFlag() bool
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisLspAdjSidFlags) FFlag() bool {

	return *obj.obj.FFlag

}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisLspAdjSidFlags) HasFFlag() bool {
	return obj.obj.FFlag != nil
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// SetFFlag sets the bool value in the IsisLspAdjSidFlags object
func (obj *isisLspAdjSidFlags) SetFFlag(value bool) IsisLspAdjSidFlags {

	obj.obj.FFlag = &value
	return obj
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisLspAdjSidFlags) BFlag() bool {

	return *obj.obj.BFlag

}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisLspAdjSidFlags) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// SetBFlag sets the bool value in the IsisLspAdjSidFlags object
func (obj *isisLspAdjSidFlags) SetBFlag(value bool) IsisLspAdjSidFlags {

	obj.obj.BFlag = &value
	return obj
}

// The value flag. If set, then the Adj-SID carries a value.
// VFlag returns a bool
func (obj *isisLspAdjSidFlags) VFlag() bool {

	return *obj.obj.VFlag

}

// The value flag. If set, then the Adj-SID carries a value.
// VFlag returns a bool
func (obj *isisLspAdjSidFlags) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// The value flag. If set, then the Adj-SID carries a value.
// SetVFlag sets the bool value in the IsisLspAdjSidFlags object
func (obj *isisLspAdjSidFlags) SetVFlag(value bool) IsisLspAdjSidFlags {

	obj.obj.VFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance.
// LFlag returns a bool
func (obj *isisLspAdjSidFlags) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance.
// LFlag returns a bool
func (obj *isisLspAdjSidFlags) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance.
// SetLFlag sets the bool value in the IsisLspAdjSidFlags object
func (obj *isisLspAdjSidFlags) SetLFlag(value bool) IsisLspAdjSidFlags {

	obj.obj.LFlag = &value
	return obj
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisLspAdjSidFlags) SFlag() bool {

	return *obj.obj.SFlag

}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisLspAdjSidFlags) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SetSFlag sets the bool value in the IsisLspAdjSidFlags object
func (obj *isisLspAdjSidFlags) SetSFlag(value bool) IsisLspAdjSidFlags {

	obj.obj.SFlag = &value
	return obj
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisLspAdjSidFlags) PFlag() bool {

	return *obj.obj.PFlag

}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisLspAdjSidFlags) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// SetPFlag sets the bool value in the IsisLspAdjSidFlags object
func (obj *isisLspAdjSidFlags) SetPFlag(value bool) IsisLspAdjSidFlags {

	obj.obj.PFlag = &value
	return obj
}

func (obj *isisLspAdjSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspAdjSidFlags) setDefault() {

}
