package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspAdjacencySID *****
type isisLspAdjacencySID struct {
	validation
	obj          *otg.IsisLspAdjacencySID
	marshaller   marshalIsisLspAdjacencySID
	unMarshaller unMarshalIsisLspAdjacencySID
}

func NewIsisLspAdjacencySID() IsisLspAdjacencySID {
	obj := isisLspAdjacencySID{obj: &otg.IsisLspAdjacencySID{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspAdjacencySID) msg() *otg.IsisLspAdjacencySID {
	return obj.obj
}

func (obj *isisLspAdjacencySID) setMsg(msg *otg.IsisLspAdjacencySID) IsisLspAdjacencySID {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspAdjacencySID struct {
	obj *isisLspAdjacencySID
}

type marshalIsisLspAdjacencySID interface {
	// ToProto marshals IsisLspAdjacencySID to protobuf object *otg.IsisLspAdjacencySID
	ToProto() (*otg.IsisLspAdjacencySID, error)
	// ToPbText marshals IsisLspAdjacencySID to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspAdjacencySID to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspAdjacencySID to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspAdjacencySID struct {
	obj *isisLspAdjacencySID
}

type unMarshalIsisLspAdjacencySID interface {
	// FromProto unmarshals IsisLspAdjacencySID from protobuf object *otg.IsisLspAdjacencySID
	FromProto(msg *otg.IsisLspAdjacencySID) (IsisLspAdjacencySID, error)
	// FromPbText unmarshals IsisLspAdjacencySID from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspAdjacencySID from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspAdjacencySID from JSON text
	FromJson(value string) error
}

func (obj *isisLspAdjacencySID) Marshal() marshalIsisLspAdjacencySID {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspAdjacencySID{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspAdjacencySID) Unmarshal() unMarshalIsisLspAdjacencySID {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspAdjacencySID{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspAdjacencySID) ToProto() (*otg.IsisLspAdjacencySID, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspAdjacencySID) FromProto(msg *otg.IsisLspAdjacencySID) (IsisLspAdjacencySID, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspAdjacencySID) ToPbText() (string, error) {
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

func (m *unMarshalisisLspAdjacencySID) FromPbText(value string) error {
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

func (m *marshalisisLspAdjacencySID) ToYaml() (string, error) {
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

func (m *unMarshalisisLspAdjacencySID) FromYaml(value string) error {
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

func (m *marshalisisLspAdjacencySID) ToJson() (string, error) {
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

func (m *unMarshalisisLspAdjacencySID) FromJson(value string) error {
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

func (obj *isisLspAdjacencySID) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspAdjacencySID) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspAdjacencySID) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspAdjacencySID) Clone() (IsisLspAdjacencySID, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspAdjacencySID()
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

// IsisLspAdjacencySID is this container defines segment routing adjacency SIDs.
type IsisLspAdjacencySID interface {
	Validation
	// msg marshals IsisLspAdjacencySID to protobuf object *otg.IsisLspAdjacencySID
	// and doesn't set defaults
	msg() *otg.IsisLspAdjacencySID
	// setMsg unmarshals IsisLspAdjacencySID from protobuf object *otg.IsisLspAdjacencySID
	// and doesn't set defaults
	setMsg(*otg.IsisLspAdjacencySID) IsisLspAdjacencySID
	// provides marshal interface
	Marshal() marshalIsisLspAdjacencySID
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspAdjacencySID
	// validate validates IsisLspAdjacencySID
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspAdjacencySID, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TlvType returns uint32, set in IsisLspAdjacencySID.
	TlvType() uint32
	// SetTlvType assigns uint32 provided by user to IsisLspAdjacencySID
	SetTlvType(value uint32) IsisLspAdjacencySID
	// HasTlvType checks if TlvType has been set in IsisLspAdjacencySID
	HasTlvType() bool
	// Sid returns uint32, set in IsisLspAdjacencySID.
	Sid() uint32
	// SetSid assigns uint32 provided by user to IsisLspAdjacencySID
	SetSid(value uint32) IsisLspAdjacencySID
	// HasSid checks if Sid has been set in IsisLspAdjacencySID
	HasSid() bool
	// FFlag returns bool, set in IsisLspAdjacencySID.
	FFlag() bool
	// SetFFlag assigns bool provided by user to IsisLspAdjacencySID
	SetFFlag(value bool) IsisLspAdjacencySID
	// HasFFlag checks if FFlag has been set in IsisLspAdjacencySID
	HasFFlag() bool
	// BFlag returns bool, set in IsisLspAdjacencySID.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisLspAdjacencySID
	SetBFlag(value bool) IsisLspAdjacencySID
	// HasBFlag checks if BFlag has been set in IsisLspAdjacencySID
	HasBFlag() bool
	// VFlag returns bool, set in IsisLspAdjacencySID.
	VFlag() bool
	// SetVFlag assigns bool provided by user to IsisLspAdjacencySID
	SetVFlag(value bool) IsisLspAdjacencySID
	// HasVFlag checks if VFlag has been set in IsisLspAdjacencySID
	HasVFlag() bool
	// LFlag returns bool, set in IsisLspAdjacencySID.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisLspAdjacencySID
	SetLFlag(value bool) IsisLspAdjacencySID
	// HasLFlag checks if LFlag has been set in IsisLspAdjacencySID
	HasLFlag() bool
	// SFlag returns bool, set in IsisLspAdjacencySID.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisLspAdjacencySID
	SetSFlag(value bool) IsisLspAdjacencySID
	// HasSFlag checks if SFlag has been set in IsisLspAdjacencySID
	HasSFlag() bool
	// PFlag returns bool, set in IsisLspAdjacencySID.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisLspAdjacencySID
	SetPFlag(value bool) IsisLspAdjacencySID
	// HasPFlag checks if PFlag has been set in IsisLspAdjacencySID
	HasPFlag() bool
	// Weight returns uint32, set in IsisLspAdjacencySID.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisLspAdjacencySID
	SetWeight(value uint32) IsisLspAdjacencySID
	// HasWeight checks if Weight has been set in IsisLspAdjacencySID
	HasWeight() bool
}

// The corresponding adjacency SID type: Adjacency-SID (Tlv - 31) LAN Adjacency-SID (Tlv - 31).
// TlvType returns a uint32
func (obj *isisLspAdjacencySID) TlvType() uint32 {

	return *obj.obj.TlvType

}

// The corresponding adjacency SID type: Adjacency-SID (Tlv - 31) LAN Adjacency-SID (Tlv - 31).
// TlvType returns a uint32
func (obj *isisLspAdjacencySID) HasTlvType() bool {
	return obj.obj.TlvType != nil
}

// The corresponding adjacency SID type: Adjacency-SID (Tlv - 31) LAN Adjacency-SID (Tlv - 31).
// SetTlvType sets the uint32 value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetTlvType(value uint32) IsisLspAdjacencySID {

	obj.obj.TlvType = &value
	return obj
}

// The corresponding adjacency SID for a link.
// Sid returns a uint32
func (obj *isisLspAdjacencySID) Sid() uint32 {

	return *obj.obj.Sid

}

// The corresponding adjacency SID for a link.
// Sid returns a uint32
func (obj *isisLspAdjacencySID) HasSid() bool {
	return obj.obj.Sid != nil
}

// The corresponding adjacency SID for a link.
// SetSid sets the uint32 value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetSid(value uint32) IsisLspAdjacencySID {

	obj.obj.Sid = &value
	return obj
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisLspAdjacencySID) FFlag() bool {

	return *obj.obj.FFlag

}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// FFlag returns a bool
func (obj *isisLspAdjacencySID) HasFFlag() bool {
	return obj.obj.FFlag != nil
}

// The address family flag. If unset, then the Adj-SID refers
// to an adjacency with outgoing IPv4 encapsulation.  If set then
// the Adj-SID refers to an adjacency with outgoing IPv6
// encapsulation.
// SetFFlag sets the bool value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetFFlag(value bool) IsisLspAdjacencySID {

	obj.obj.FFlag = &value
	return obj
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisLspAdjacencySID) BFlag() bool {

	return *obj.obj.BFlag

}

// The backup flag. If set, the Adj-SID is eligible for protection.
// BFlag returns a bool
func (obj *isisLspAdjacencySID) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// The backup flag. If set, the Adj-SID is eligible for protection.
// SetBFlag sets the bool value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetBFlag(value bool) IsisLspAdjacencySID {

	obj.obj.BFlag = &value
	return obj
}

// The value flag. If set, then the Adj-SID carries a value.
// VFlag returns a bool
func (obj *isisLspAdjacencySID) VFlag() bool {

	return *obj.obj.VFlag

}

// The value flag. If set, then the Adj-SID carries a value.
// VFlag returns a bool
func (obj *isisLspAdjacencySID) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// The value flag. If set, then the Adj-SID carries a value.
// SetVFlag sets the bool value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetVFlag(value bool) IsisLspAdjacencySID {

	obj.obj.VFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance.
// LFlag returns a bool
func (obj *isisLspAdjacencySID) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance.
// LFlag returns a bool
func (obj *isisLspAdjacencySID) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Adj-SID has local significance.
// SetLFlag sets the bool value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetLFlag(value bool) IsisLspAdjacencySID {

	obj.obj.LFlag = &value
	return obj
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisLspAdjacencySID) SFlag() bool {

	return *obj.obj.SFlag

}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SFlag returns a bool
func (obj *isisLspAdjacencySID) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// The set flag. When set, the S-Flag indicates that the
// Adj-SID refers to a set of adjacencies (and therefore MAY be
// assigned to other adjacencies as well).
// SetSFlag sets the bool value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetSFlag(value bool) IsisLspAdjacencySID {

	obj.obj.SFlag = &value
	return obj
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisLspAdjacencySID) PFlag() bool {

	return *obj.obj.PFlag

}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// PFlag returns a bool
func (obj *isisLspAdjacencySID) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// The persistent flag. When set, the P-Flag indicates that
// the Adj-SID is persistently allocated, i.e., the Adj-SID value
// remains consistent across router restart and/or interface flap.
// SetPFlag sets the bool value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetPFlag(value bool) IsisLspAdjacencySID {

	obj.obj.PFlag = &value
	return obj
}

// The value represents the weight of the Adj-SID for the purpose of
// load balancing.
// Weight returns a uint32
func (obj *isisLspAdjacencySID) Weight() uint32 {

	return *obj.obj.Weight

}

// The value represents the weight of the Adj-SID for the purpose of
// load balancing.
// Weight returns a uint32
func (obj *isisLspAdjacencySID) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The value represents the weight of the Adj-SID for the purpose of
// load balancing.
// SetWeight sets the uint32 value in the IsisLspAdjacencySID object
func (obj *isisLspAdjacencySID) SetWeight(value uint32) IsisLspAdjacencySID {

	obj.obj.Weight = &value
	return obj
}

func (obj *isisLspAdjacencySID) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *isisLspAdjacencySID) setDefault() {

}
