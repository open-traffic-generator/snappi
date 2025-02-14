package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisInterfaceAdjSidFlags *****
type isisInterfaceAdjSidFlags struct {
	validation
	obj          *otg.IsisInterfaceAdjSidFlags
	marshaller   marshalIsisInterfaceAdjSidFlags
	unMarshaller unMarshalIsisInterfaceAdjSidFlags
}

func NewIsisInterfaceAdjSidFlags() IsisInterfaceAdjSidFlags {
	obj := isisInterfaceAdjSidFlags{obj: &otg.IsisInterfaceAdjSidFlags{}}
	obj.setDefault()
	return &obj
}

func (obj *isisInterfaceAdjSidFlags) msg() *otg.IsisInterfaceAdjSidFlags {
	return obj.obj
}

func (obj *isisInterfaceAdjSidFlags) setMsg(msg *otg.IsisInterfaceAdjSidFlags) IsisInterfaceAdjSidFlags {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisInterfaceAdjSidFlags struct {
	obj *isisInterfaceAdjSidFlags
}

type marshalIsisInterfaceAdjSidFlags interface {
	// ToProto marshals IsisInterfaceAdjSidFlags to protobuf object *otg.IsisInterfaceAdjSidFlags
	ToProto() (*otg.IsisInterfaceAdjSidFlags, error)
	// ToPbText marshals IsisInterfaceAdjSidFlags to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisInterfaceAdjSidFlags to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisInterfaceAdjSidFlags to JSON text
	ToJson() (string, error)
}

type unMarshalisisInterfaceAdjSidFlags struct {
	obj *isisInterfaceAdjSidFlags
}

type unMarshalIsisInterfaceAdjSidFlags interface {
	// FromProto unmarshals IsisInterfaceAdjSidFlags from protobuf object *otg.IsisInterfaceAdjSidFlags
	FromProto(msg *otg.IsisInterfaceAdjSidFlags) (IsisInterfaceAdjSidFlags, error)
	// FromPbText unmarshals IsisInterfaceAdjSidFlags from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisInterfaceAdjSidFlags from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisInterfaceAdjSidFlags from JSON text
	FromJson(value string) error
}

func (obj *isisInterfaceAdjSidFlags) Marshal() marshalIsisInterfaceAdjSidFlags {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisInterfaceAdjSidFlags{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisInterfaceAdjSidFlags) Unmarshal() unMarshalIsisInterfaceAdjSidFlags {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisInterfaceAdjSidFlags{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisInterfaceAdjSidFlags) ToProto() (*otg.IsisInterfaceAdjSidFlags, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisInterfaceAdjSidFlags) FromProto(msg *otg.IsisInterfaceAdjSidFlags) (IsisInterfaceAdjSidFlags, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisInterfaceAdjSidFlags) ToPbText() (string, error) {
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

func (m *unMarshalisisInterfaceAdjSidFlags) FromPbText(value string) error {
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

func (m *marshalisisInterfaceAdjSidFlags) ToYaml() (string, error) {
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

func (m *unMarshalisisInterfaceAdjSidFlags) FromYaml(value string) error {
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

func (m *marshalisisInterfaceAdjSidFlags) ToJson() (string, error) {
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

func (m *unMarshalisisInterfaceAdjSidFlags) FromJson(value string) error {
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

func (obj *isisInterfaceAdjSidFlags) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisInterfaceAdjSidFlags) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisInterfaceAdjSidFlags) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisInterfaceAdjSidFlags) Clone() (IsisInterfaceAdjSidFlags, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisInterfaceAdjSidFlags()
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

// IsisInterfaceAdjSidFlags is conatiner of 1-octet Flags associated with Adjacency Segment-ID.
type IsisInterfaceAdjSidFlags interface {
	Validation
	// msg marshals IsisInterfaceAdjSidFlags to protobuf object *otg.IsisInterfaceAdjSidFlags
	// and doesn't set defaults
	msg() *otg.IsisInterfaceAdjSidFlags
	// setMsg unmarshals IsisInterfaceAdjSidFlags from protobuf object *otg.IsisInterfaceAdjSidFlags
	// and doesn't set defaults
	setMsg(*otg.IsisInterfaceAdjSidFlags) IsisInterfaceAdjSidFlags
	// provides marshal interface
	Marshal() marshalIsisInterfaceAdjSidFlags
	// provides unmarshal interface
	Unmarshal() unMarshalIsisInterfaceAdjSidFlags
	// validate validates IsisInterfaceAdjSidFlags
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisInterfaceAdjSidFlags, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FFlag returns bool, set in IsisInterfaceAdjSidFlags.
	FFlag() bool
	// SetFFlag assigns bool provided by user to IsisInterfaceAdjSidFlags
	SetFFlag(value bool) IsisInterfaceAdjSidFlags
	// HasFFlag checks if FFlag has been set in IsisInterfaceAdjSidFlags
	HasFFlag() bool
	// BFlag returns bool, set in IsisInterfaceAdjSidFlags.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisInterfaceAdjSidFlags
	SetBFlag(value bool) IsisInterfaceAdjSidFlags
	// HasBFlag checks if BFlag has been set in IsisInterfaceAdjSidFlags
	HasBFlag() bool
	// LFlag returns bool, set in IsisInterfaceAdjSidFlags.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisInterfaceAdjSidFlags
	SetLFlag(value bool) IsisInterfaceAdjSidFlags
	// HasLFlag checks if LFlag has been set in IsisInterfaceAdjSidFlags
	HasLFlag() bool
	// SFlag returns bool, set in IsisInterfaceAdjSidFlags.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisInterfaceAdjSidFlags
	SetSFlag(value bool) IsisInterfaceAdjSidFlags
	// HasSFlag checks if SFlag has been set in IsisInterfaceAdjSidFlags
	HasSFlag() bool
	// PFlag returns bool, set in IsisInterfaceAdjSidFlags.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisInterfaceAdjSidFlags
	SetPFlag(value bool) IsisInterfaceAdjSidFlags
	// HasPFlag checks if PFlag has been set in IsisInterfaceAdjSidFlags
	HasPFlag() bool
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) FFlag() bool {

	return *obj.obj.FFlag

}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) HasFFlag() bool {
	return obj.obj.FFlag != nil
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// SetFFlag sets the bool value in the IsisInterfaceAdjSidFlags object
func (obj *isisInterfaceAdjSidFlags) SetFFlag(value bool) IsisInterfaceAdjSidFlags {

	obj.obj.FFlag = &value
	return obj
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) BFlag() bool {

	return *obj.obj.BFlag

}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// SetBFlag sets the bool value in the IsisInterfaceAdjSidFlags object
func (obj *isisInterfaceAdjSidFlags) SetBFlag(value bool) IsisInterfaceAdjSidFlags {

	obj.obj.BFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// LFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance. In this case, Adjacency_sid is from device.isis.segment_routing.router_capability.srlb_ranges.
// SetLFlag sets the bool value in the IsisInterfaceAdjSidFlags object
func (obj *isisInterfaceAdjSidFlags) SetLFlag(value bool) IsisInterfaceAdjSidFlags {

	obj.obj.LFlag = &value
	return obj
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) SFlag() bool {

	return *obj.obj.SFlag

}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SetSFlag sets the bool value in the IsisInterfaceAdjSidFlags object
func (obj *isisInterfaceAdjSidFlags) SetSFlag(value bool) IsisInterfaceAdjSidFlags {

	obj.obj.SFlag = &value
	return obj
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) PFlag() bool {

	return *obj.obj.PFlag

}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisInterfaceAdjSidFlags) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// SetPFlag sets the bool value in the IsisInterfaceAdjSidFlags object
func (obj *isisInterfaceAdjSidFlags) SetPFlag(value bool) IsisInterfaceAdjSidFlags {

	obj.obj.PFlag = &value
	return obj
}

func (obj *isisInterfaceAdjSidFlags) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisInterfaceAdjSidFlags) setDefault() {
	if obj.obj.FFlag == nil {
		obj.SetFFlag(true)
	}
	if obj.obj.BFlag == nil {
		obj.SetBFlag(false)
	}
	if obj.obj.LFlag == nil {
		obj.SetLFlag(true)
	}
	if obj.obj.SFlag == nil {
		obj.SetSFlag(false)
	}
	if obj.obj.PFlag == nil {
		obj.SetPFlag(false)
	}

}
