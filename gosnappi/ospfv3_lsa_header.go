package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3LsaHeader *****
type ospfv3LsaHeader struct {
	validation
	obj          *otg.Ospfv3LsaHeader
	marshaller   marshalOspfv3LsaHeader
	unMarshaller unMarshalOspfv3LsaHeader
}

func NewOspfv3LsaHeader() Ospfv3LsaHeader {
	obj := ospfv3LsaHeader{obj: &otg.Ospfv3LsaHeader{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3LsaHeader) msg() *otg.Ospfv3LsaHeader {
	return obj.obj
}

func (obj *ospfv3LsaHeader) setMsg(msg *otg.Ospfv3LsaHeader) Ospfv3LsaHeader {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3LsaHeader struct {
	obj *ospfv3LsaHeader
}

type marshalOspfv3LsaHeader interface {
	// ToProto marshals Ospfv3LsaHeader to protobuf object *otg.Ospfv3LsaHeader
	ToProto() (*otg.Ospfv3LsaHeader, error)
	// ToPbText marshals Ospfv3LsaHeader to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3LsaHeader to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3LsaHeader to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3LsaHeader struct {
	obj *ospfv3LsaHeader
}

type unMarshalOspfv3LsaHeader interface {
	// FromProto unmarshals Ospfv3LsaHeader from protobuf object *otg.Ospfv3LsaHeader
	FromProto(msg *otg.Ospfv3LsaHeader) (Ospfv3LsaHeader, error)
	// FromPbText unmarshals Ospfv3LsaHeader from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3LsaHeader from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3LsaHeader from JSON text
	FromJson(value string) error
}

func (obj *ospfv3LsaHeader) Marshal() marshalOspfv3LsaHeader {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3LsaHeader{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3LsaHeader) Unmarshal() unMarshalOspfv3LsaHeader {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3LsaHeader{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3LsaHeader) ToProto() (*otg.Ospfv3LsaHeader, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3LsaHeader) FromProto(msg *otg.Ospfv3LsaHeader) (Ospfv3LsaHeader, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3LsaHeader) ToPbText() (string, error) {
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

func (m *unMarshalospfv3LsaHeader) FromPbText(value string) error {
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

func (m *marshalospfv3LsaHeader) ToYaml() (string, error) {
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

func (m *unMarshalospfv3LsaHeader) FromYaml(value string) error {
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

func (m *marshalospfv3LsaHeader) ToJson() (string, error) {
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

func (m *unMarshalospfv3LsaHeader) FromJson(value string) error {
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

func (obj *ospfv3LsaHeader) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3LsaHeader) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3LsaHeader) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3LsaHeader) Clone() (Ospfv3LsaHeader, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3LsaHeader()
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

// Ospfv3LsaHeader is attributes in LSA Header.
type Ospfv3LsaHeader interface {
	Validation
	// msg marshals Ospfv3LsaHeader to protobuf object *otg.Ospfv3LsaHeader
	// and doesn't set defaults
	msg() *otg.Ospfv3LsaHeader
	// setMsg unmarshals Ospfv3LsaHeader from protobuf object *otg.Ospfv3LsaHeader
	// and doesn't set defaults
	setMsg(*otg.Ospfv3LsaHeader) Ospfv3LsaHeader
	// provides marshal interface
	Marshal() marshalOspfv3LsaHeader
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3LsaHeader
	// validate validates Ospfv3LsaHeader
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3LsaHeader, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsaId returns string, set in Ospfv3LsaHeader.
	LsaId() string
	// SetLsaId assigns string provided by user to Ospfv3LsaHeader
	SetLsaId(value string) Ospfv3LsaHeader
	// HasLsaId checks if LsaId has been set in Ospfv3LsaHeader
	HasLsaId() bool
	// AdvertisingRouterId returns string, set in Ospfv3LsaHeader.
	AdvertisingRouterId() string
	// SetAdvertisingRouterId assigns string provided by user to Ospfv3LsaHeader
	SetAdvertisingRouterId(value string) Ospfv3LsaHeader
	// HasAdvertisingRouterId checks if AdvertisingRouterId has been set in Ospfv3LsaHeader
	HasAdvertisingRouterId() bool
	// SequenceNumber returns uint32, set in Ospfv3LsaHeader.
	SequenceNumber() uint32
	// SetSequenceNumber assigns uint32 provided by user to Ospfv3LsaHeader
	SetSequenceNumber(value uint32) Ospfv3LsaHeader
	// HasSequenceNumber checks if SequenceNumber has been set in Ospfv3LsaHeader
	HasSequenceNumber() bool
	// Age returns uint32, set in Ospfv3LsaHeader.
	Age() uint32
	// SetAge assigns uint32 provided by user to Ospfv3LsaHeader
	SetAge(value uint32) Ospfv3LsaHeader
	// HasAge checks if Age has been set in Ospfv3LsaHeader
	HasAge() bool
}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// LsaId returns a string
func (obj *ospfv3LsaHeader) LsaId() string {

	return *obj.obj.LsaId

}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// LsaId returns a string
func (obj *ospfv3LsaHeader) HasLsaId() bool {
	return obj.obj.LsaId != nil
}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// SetLsaId sets the string value in the Ospfv3LsaHeader object
func (obj *ospfv3LsaHeader) SetLsaId(value string) Ospfv3LsaHeader {

	obj.obj.LsaId = &value
	return obj
}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// AdvertisingRouterId returns a string
func (obj *ospfv3LsaHeader) AdvertisingRouterId() string {

	return *obj.obj.AdvertisingRouterId

}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// AdvertisingRouterId returns a string
func (obj *ospfv3LsaHeader) HasAdvertisingRouterId() bool {
	return obj.obj.AdvertisingRouterId != nil
}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// SetAdvertisingRouterId sets the string value in the Ospfv3LsaHeader object
func (obj *ospfv3LsaHeader) SetAdvertisingRouterId(value string) Ospfv3LsaHeader {

	obj.obj.AdvertisingRouterId = &value
	return obj
}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SequenceNumber returns a uint32
func (obj *ospfv3LsaHeader) SequenceNumber() uint32 {

	return *obj.obj.SequenceNumber

}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SequenceNumber returns a uint32
func (obj *ospfv3LsaHeader) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SetSequenceNumber sets the uint32 value in the Ospfv3LsaHeader object
func (obj *ospfv3LsaHeader) SetSequenceNumber(value uint32) Ospfv3LsaHeader {

	obj.obj.SequenceNumber = &value
	return obj
}

// The time since the LSA's generation in seconds.
// Age returns a uint32
func (obj *ospfv3LsaHeader) Age() uint32 {

	return *obj.obj.Age

}

// The time since the LSA's generation in seconds.
// Age returns a uint32
func (obj *ospfv3LsaHeader) HasAge() bool {
	return obj.obj.Age != nil
}

// The time since the LSA's generation in seconds.
// SetAge sets the uint32 value in the Ospfv3LsaHeader object
func (obj *ospfv3LsaHeader) SetAge(value uint32) Ospfv3LsaHeader {

	obj.obj.Age = &value
	return obj
}

func (obj *ospfv3LsaHeader) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LsaId != nil {

		err := obj.validateIpv4(obj.LsaId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3LsaHeader.LsaId"))
		}

	}

	if obj.obj.AdvertisingRouterId != nil {

		err := obj.validateIpv4(obj.AdvertisingRouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv3LsaHeader.AdvertisingRouterId"))
		}

	}

}

func (obj *ospfv3LsaHeader) setDefault() {

}
