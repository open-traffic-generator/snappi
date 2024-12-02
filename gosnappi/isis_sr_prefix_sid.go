package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRPrefixSID *****
type isisSRPrefixSID struct {
	validation
	obj          *otg.IsisSRPrefixSID
	marshaller   marshalIsisSRPrefixSID
	unMarshaller unMarshalIsisSRPrefixSID
}

func NewIsisSRPrefixSID() IsisSRPrefixSID {
	obj := isisSRPrefixSID{obj: &otg.IsisSRPrefixSID{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRPrefixSID) msg() *otg.IsisSRPrefixSID {
	return obj.obj
}

func (obj *isisSRPrefixSID) setMsg(msg *otg.IsisSRPrefixSID) IsisSRPrefixSID {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRPrefixSID struct {
	obj *isisSRPrefixSID
}

type marshalIsisSRPrefixSID interface {
	// ToProto marshals IsisSRPrefixSID to protobuf object *otg.IsisSRPrefixSID
	ToProto() (*otg.IsisSRPrefixSID, error)
	// ToPbText marshals IsisSRPrefixSID to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRPrefixSID to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRPrefixSID to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRPrefixSID struct {
	obj *isisSRPrefixSID
}

type unMarshalIsisSRPrefixSID interface {
	// FromProto unmarshals IsisSRPrefixSID from protobuf object *otg.IsisSRPrefixSID
	FromProto(msg *otg.IsisSRPrefixSID) (IsisSRPrefixSID, error)
	// FromPbText unmarshals IsisSRPrefixSID from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRPrefixSID from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRPrefixSID from JSON text
	FromJson(value string) error
}

func (obj *isisSRPrefixSID) Marshal() marshalIsisSRPrefixSID {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRPrefixSID{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRPrefixSID) Unmarshal() unMarshalIsisSRPrefixSID {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRPrefixSID{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRPrefixSID) ToProto() (*otg.IsisSRPrefixSID, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRPrefixSID) FromProto(msg *otg.IsisSRPrefixSID) (IsisSRPrefixSID, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRPrefixSID) ToPbText() (string, error) {
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

func (m *unMarshalisisSRPrefixSID) FromPbText(value string) error {
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

func (m *marshalisisSRPrefixSID) ToYaml() (string, error) {
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

func (m *unMarshalisisSRPrefixSID) FromYaml(value string) error {
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

func (m *marshalisisSRPrefixSID) ToJson() (string, error) {
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

func (m *unMarshalisisSRPrefixSID) FromJson(value string) error {
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

func (obj *isisSRPrefixSID) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRPrefixSID) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRPrefixSID) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRPrefixSID) Clone() (IsisSRPrefixSID, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRPrefixSID()
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

// IsisSRPrefixSID is this contains the properties of IS-IS Prefix-SID and its attributes for  the extended Ipv4 and Ipv6 reachability. Refernce: https://datatracker.ietf.org/doc/html/rfc8667#name-prefix-segment-identifier-p.
type IsisSRPrefixSID interface {
	Validation
	// msg marshals IsisSRPrefixSID to protobuf object *otg.IsisSRPrefixSID
	// and doesn't set defaults
	msg() *otg.IsisSRPrefixSID
	// setMsg unmarshals IsisSRPrefixSID from protobuf object *otg.IsisSRPrefixSID
	// and doesn't set defaults
	setMsg(*otg.IsisSRPrefixSID) IsisSRPrefixSID
	// provides marshal interface
	Marshal() marshalIsisSRPrefixSID
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRPrefixSID
	// validate validates IsisSRPrefixSID
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRPrefixSID, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sid returns uint32, set in IsisSRPrefixSID.
	Sid() uint32
	// SetSid assigns uint32 provided by user to IsisSRPrefixSID
	SetSid(value uint32) IsisSRPrefixSID
	// HasSid checks if Sid has been set in IsisSRPrefixSID
	HasSid() bool
	// RFlag returns bool, set in IsisSRPrefixSID.
	RFlag() bool
	// SetRFlag assigns bool provided by user to IsisSRPrefixSID
	SetRFlag(value bool) IsisSRPrefixSID
	// HasRFlag checks if RFlag has been set in IsisSRPrefixSID
	HasRFlag() bool
	// NFlag returns bool, set in IsisSRPrefixSID.
	NFlag() bool
	// SetNFlag assigns bool provided by user to IsisSRPrefixSID
	SetNFlag(value bool) IsisSRPrefixSID
	// HasNFlag checks if NFlag has been set in IsisSRPrefixSID
	HasNFlag() bool
	// PFlag returns bool, set in IsisSRPrefixSID.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisSRPrefixSID
	SetPFlag(value bool) IsisSRPrefixSID
	// HasPFlag checks if PFlag has been set in IsisSRPrefixSID
	HasPFlag() bool
	// EFlag returns bool, set in IsisSRPrefixSID.
	EFlag() bool
	// SetEFlag assigns bool provided by user to IsisSRPrefixSID
	SetEFlag(value bool) IsisSRPrefixSID
	// HasEFlag checks if EFlag has been set in IsisSRPrefixSID
	HasEFlag() bool
	// VFlag returns bool, set in IsisSRPrefixSID.
	VFlag() bool
	// SetVFlag assigns bool provided by user to IsisSRPrefixSID
	SetVFlag(value bool) IsisSRPrefixSID
	// HasVFlag checks if VFlag has been set in IsisSRPrefixSID
	HasVFlag() bool
	// LFlag returns bool, set in IsisSRPrefixSID.
	LFlag() bool
	// SetLFlag assigns bool provided by user to IsisSRPrefixSID
	SetLFlag(value bool) IsisSRPrefixSID
	// HasLFlag checks if LFlag has been set in IsisSRPrefixSID
	HasLFlag() bool
	// Algorithm returns uint32, set in IsisSRPrefixSID.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisSRPrefixSID
	SetAlgorithm(value uint32) IsisSRPrefixSID
	// HasAlgorithm checks if Algorithm has been set in IsisSRPrefixSID
	HasAlgorithm() bool
}

// SID/Index is the SID/Label value associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// Sid returns a uint32
func (obj *isisSRPrefixSID) Sid() uint32 {

	return *obj.obj.Sid

}

// SID/Index is the SID/Label value associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// Sid returns a uint32
func (obj *isisSRPrefixSID) HasSid() bool {
	return obj.obj.Sid != nil
}

// SID/Index is the SID/Label value associated  with the IGP Prefix segment attached to the specific IPv4 or IPv6 prefix.
// SetSid sets the uint32 value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetSid(value uint32) IsisSRPrefixSID {

	obj.obj.Sid = &value
	return obj
}

// R-Flag:Re-advertisement flag
// RFlag returns a bool
func (obj *isisSRPrefixSID) RFlag() bool {

	return *obj.obj.RFlag

}

// R-Flag:Re-advertisement flag
// RFlag returns a bool
func (obj *isisSRPrefixSID) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// R-Flag:Re-advertisement flag
// SetRFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetRFlag(value bool) IsisSRPrefixSID {

	obj.obj.RFlag = &value
	return obj
}

// N-Flag: Node-SID flag.  If set, then the Prefix-SID refers to
// the router identified by the prefix
// NFlag returns a bool
func (obj *isisSRPrefixSID) NFlag() bool {

	return *obj.obj.NFlag

}

// N-Flag: Node-SID flag.  If set, then the Prefix-SID refers to
// the router identified by the prefix
// NFlag returns a bool
func (obj *isisSRPrefixSID) HasNFlag() bool {
	return obj.obj.NFlag != nil
}

// N-Flag: Node-SID flag.  If set, then the Prefix-SID refers to
// the router identified by the prefix
// SetNFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetNFlag(value bool) IsisSRPrefixSID {

	obj.obj.NFlag = &value
	return obj
}

// P-Flag: no-PHP flag.
// PFlag returns a bool
func (obj *isisSRPrefixSID) PFlag() bool {

	return *obj.obj.PFlag

}

// P-Flag: no-PHP flag.
// PFlag returns a bool
func (obj *isisSRPrefixSID) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// P-Flag: no-PHP flag.
// SetPFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetPFlag(value bool) IsisSRPrefixSID {

	obj.obj.PFlag = &value
	return obj
}

// E-Flag: Explicit-Null Flag.
// EFlag returns a bool
func (obj *isisSRPrefixSID) EFlag() bool {

	return *obj.obj.EFlag

}

// E-Flag: Explicit-Null Flag.
// EFlag returns a bool
func (obj *isisSRPrefixSID) HasEFlag() bool {
	return obj.obj.EFlag != nil
}

// E-Flag: Explicit-Null Flag.
// SetEFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetEFlag(value bool) IsisSRPrefixSID {

	obj.obj.EFlag = &value
	return obj
}

// The value flag. If set, then the Prefix-SID carries a value instead of an index.
// VFlag returns a bool
func (obj *isisSRPrefixSID) VFlag() bool {

	return *obj.obj.VFlag

}

// The value flag. If set, then the Prefix-SID carries a value instead of an index.
// VFlag returns a bool
func (obj *isisSRPrefixSID) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// The value flag. If set, then the Prefix-SID carries a value instead of an index.
// SetVFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetVFlag(value bool) IsisSRPrefixSID {

	obj.obj.VFlag = &value
	return obj
}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *isisSRPrefixSID) LFlag() bool {

	return *obj.obj.LFlag

}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance.
// LFlag returns a bool
func (obj *isisSRPrefixSID) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// The local flag.  If set, then the value/index carried by
// the Prefix-SID has local significance.
// SetLFlag sets the bool value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetLFlag(value bool) IsisSRPrefixSID {

	obj.obj.LFlag = &value
	return obj
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisSRPrefixSID) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// Algorithm returns a uint32
func (obj *isisSRPrefixSID) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// The Isis may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these
// nodes.
// SetAlgorithm sets the uint32 value in the IsisSRPrefixSID object
func (obj *isisSRPrefixSID) SetAlgorithm(value uint32) IsisSRPrefixSID {

	obj.obj.Algorithm = &value
	return obj
}

func (obj *isisSRPrefixSID) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRPrefixSID.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

}

func (obj *isisSRPrefixSID) setDefault() {
	if obj.obj.Sid == nil {
		obj.SetSid(0)
	}
	if obj.obj.RFlag == nil {
		obj.SetRFlag(false)
	}
	if obj.obj.NFlag == nil {
		obj.SetNFlag(false)
	}
	if obj.obj.PFlag == nil {
		obj.SetPFlag(false)
	}
	if obj.obj.EFlag == nil {
		obj.SetEFlag(false)
	}
	if obj.obj.VFlag == nil {
		obj.SetVFlag(false)
	}
	if obj.obj.LFlag == nil {
		obj.SetLFlag(false)
	}
	if obj.obj.Algorithm == nil {
		obj.SetAlgorithm(0)
	}

}
