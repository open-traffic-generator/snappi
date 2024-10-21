package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaHeader *****
type ospfv2LsaHeader struct {
	validation
	obj          *otg.Ospfv2LsaHeader
	marshaller   marshalOspfv2LsaHeader
	unMarshaller unMarshalOspfv2LsaHeader
}

func NewOspfv2LsaHeader() Ospfv2LsaHeader {
	obj := ospfv2LsaHeader{obj: &otg.Ospfv2LsaHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaHeader) msg() *otg.Ospfv2LsaHeader {
	return obj.obj
}

func (obj *ospfv2LsaHeader) setMsg(msg *otg.Ospfv2LsaHeader) Ospfv2LsaHeader {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaHeader struct {
	obj *ospfv2LsaHeader
}

type marshalOspfv2LsaHeader interface {
	// ToProto marshals Ospfv2LsaHeader to protobuf object *otg.Ospfv2LsaHeader
	ToProto() (*otg.Ospfv2LsaHeader, error)
	// ToPbText marshals Ospfv2LsaHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaHeader to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaHeader struct {
	obj *ospfv2LsaHeader
}

type unMarshalOspfv2LsaHeader interface {
	// FromProto unmarshals Ospfv2LsaHeader from protobuf object *otg.Ospfv2LsaHeader
	FromProto(msg *otg.Ospfv2LsaHeader) (Ospfv2LsaHeader, error)
	// FromPbText unmarshals Ospfv2LsaHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaHeader from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaHeader) Marshal() marshalOspfv2LsaHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaHeader) Unmarshal() unMarshalOspfv2LsaHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaHeader) ToProto() (*otg.Ospfv2LsaHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaHeader) FromProto(msg *otg.Ospfv2LsaHeader) (Ospfv2LsaHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaHeader) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaHeader) FromPbText(value string) error {
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

func (m *marshalospfv2LsaHeader) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaHeader) FromYaml(value string) error {
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

func (m *marshalospfv2LsaHeader) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaHeader) FromJson(value string) error {
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

func (obj *ospfv2LsaHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaHeader) Clone() (Ospfv2LsaHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaHeader()
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

// Ospfv2LsaHeader is attributes in LSA Header.
type Ospfv2LsaHeader interface {
	Validation
	// msg marshals Ospfv2LsaHeader to protobuf object *otg.Ospfv2LsaHeader
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaHeader
	// setMsg unmarshals Ospfv2LsaHeader from protobuf object *otg.Ospfv2LsaHeader
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaHeader) Ospfv2LsaHeader
	// provides marshal interface
	Marshal() marshalOspfv2LsaHeader
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaHeader
	// validate validates Ospfv2LsaHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsaId returns string, set in Ospfv2LsaHeader.
	LsaId() string
	// SetLsaId assigns string provided by user to Ospfv2LsaHeader
	SetLsaId(value string) Ospfv2LsaHeader
	// HasLsaId checks if LsaId has been set in Ospfv2LsaHeader
	HasLsaId() bool
	// AdvertisingRouterId returns string, set in Ospfv2LsaHeader.
	AdvertisingRouterId() string
	// SetAdvertisingRouterId assigns string provided by user to Ospfv2LsaHeader
	SetAdvertisingRouterId(value string) Ospfv2LsaHeader
	// HasAdvertisingRouterId checks if AdvertisingRouterId has been set in Ospfv2LsaHeader
	HasAdvertisingRouterId() bool
	// SequenceNumber returns uint32, set in Ospfv2LsaHeader.
	SequenceNumber() uint32
	// SetSequenceNumber assigns uint32 provided by user to Ospfv2LsaHeader
	SetSequenceNumber(value uint32) Ospfv2LsaHeader
	// HasSequenceNumber checks if SequenceNumber has been set in Ospfv2LsaHeader
	HasSequenceNumber() bool
	// Age returns uint32, set in Ospfv2LsaHeader.
	Age() uint32
	// SetAge assigns uint32 provided by user to Ospfv2LsaHeader
	SetAge(value uint32) Ospfv2LsaHeader
	// HasAge checks if Age has been set in Ospfv2LsaHeader
	HasAge() bool
	// OptionBits returns uint32, set in Ospfv2LsaHeader.
	OptionBits() uint32
	// SetOptionBits assigns uint32 provided by user to Ospfv2LsaHeader
	SetOptionBits(value uint32) Ospfv2LsaHeader
	// HasOptionBits checks if OptionBits has been set in Ospfv2LsaHeader
	HasOptionBits() bool
}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// LsaId returns a string
func (obj *ospfv2LsaHeader) LsaId() string {

	return *obj.obj.LsaId

}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// LsaId returns a string
func (obj *ospfv2LsaHeader) HasLsaId() bool {
	return obj.obj.LsaId != nil
}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// SetLsaId sets the string value in the Ospfv2LsaHeader object
func (obj *ospfv2LsaHeader) SetLsaId(value string) Ospfv2LsaHeader {

	obj.obj.LsaId = &value
	return obj
}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// AdvertisingRouterId returns a string
func (obj *ospfv2LsaHeader) AdvertisingRouterId() string {

	return *obj.obj.AdvertisingRouterId

}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// AdvertisingRouterId returns a string
func (obj *ospfv2LsaHeader) HasAdvertisingRouterId() bool {
	return obj.obj.AdvertisingRouterId != nil
}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// SetAdvertisingRouterId sets the string value in the Ospfv2LsaHeader object
func (obj *ospfv2LsaHeader) SetAdvertisingRouterId(value string) Ospfv2LsaHeader {

	obj.obj.AdvertisingRouterId = &value
	return obj
}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SequenceNumber returns a uint32
func (obj *ospfv2LsaHeader) SequenceNumber() uint32 {

	return *obj.obj.SequenceNumber

}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SequenceNumber returns a uint32
func (obj *ospfv2LsaHeader) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SetSequenceNumber sets the uint32 value in the Ospfv2LsaHeader object
func (obj *ospfv2LsaHeader) SetSequenceNumber(value uint32) Ospfv2LsaHeader {

	obj.obj.SequenceNumber = &value
	return obj
}

// The time since the LSA's generation in seconds.
// Age returns a uint32
func (obj *ospfv2LsaHeader) Age() uint32 {

	return *obj.obj.Age

}

// The time since the LSA's generation in seconds.
// Age returns a uint32
func (obj *ospfv2LsaHeader) HasAge() bool {
	return obj.obj.Age != nil
}

// The time since the LSA's generation in seconds.
// SetAge sets the uint32 value in the Ospfv2LsaHeader object
func (obj *ospfv2LsaHeader) SetAge(value uint32) Ospfv2LsaHeader {

	obj.obj.Age = &value
	return obj
}

// The optional bits.
// OptionBits returns a uint32
func (obj *ospfv2LsaHeader) OptionBits() uint32 {

	return *obj.obj.OptionBits

}

// The optional bits.
// OptionBits returns a uint32
func (obj *ospfv2LsaHeader) HasOptionBits() bool {
	return obj.obj.OptionBits != nil
}

// The optional bits.
// SetOptionBits sets the uint32 value in the Ospfv2LsaHeader object
func (obj *ospfv2LsaHeader) SetOptionBits(value uint32) Ospfv2LsaHeader {

	obj.obj.OptionBits = &value
	return obj
}

func (obj *ospfv2LsaHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LsaId != nil {

		err := obj.validateIpv4(obj.LsaId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2LsaHeader.LsaId"))
		}

	}

	if obj.obj.AdvertisingRouterId != nil {

		err := obj.validateIpv4(obj.AdvertisingRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2LsaHeader.AdvertisingRouterId"))
		}

	}

}

func (obj *ospfv2LsaHeader) setDefault() {

}
