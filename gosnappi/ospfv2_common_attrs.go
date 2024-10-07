package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2CommonAttrs *****
type ospfv2CommonAttrs struct {
	validation
	obj          *otg.Ospfv2CommonAttrs
	marshaller   marshalOspfv2CommonAttrs
	unMarshaller unMarshalOspfv2CommonAttrs
}

func NewOspfv2CommonAttrs() Ospfv2CommonAttrs {
	obj := ospfv2CommonAttrs{obj: &otg.Ospfv2CommonAttrs{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2CommonAttrs) msg() *otg.Ospfv2CommonAttrs {
	return obj.obj
}

func (obj *ospfv2CommonAttrs) setMsg(msg *otg.Ospfv2CommonAttrs) Ospfv2CommonAttrs {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2CommonAttrs struct {
	obj *ospfv2CommonAttrs
}

type marshalOspfv2CommonAttrs interface {
	// ToProto marshals Ospfv2CommonAttrs to protobuf object *otg.Ospfv2CommonAttrs
	ToProto() (*otg.Ospfv2CommonAttrs, error)
	// ToPbText marshals Ospfv2CommonAttrs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2CommonAttrs to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2CommonAttrs to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2CommonAttrs struct {
	obj *ospfv2CommonAttrs
}

type unMarshalOspfv2CommonAttrs interface {
	// FromProto unmarshals Ospfv2CommonAttrs from protobuf object *otg.Ospfv2CommonAttrs
	FromProto(msg *otg.Ospfv2CommonAttrs) (Ospfv2CommonAttrs, error)
	// FromPbText unmarshals Ospfv2CommonAttrs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2CommonAttrs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2CommonAttrs from JSON text
	FromJson(value string) error
}

func (obj *ospfv2CommonAttrs) Marshal() marshalOspfv2CommonAttrs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2CommonAttrs{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2CommonAttrs) Unmarshal() unMarshalOspfv2CommonAttrs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2CommonAttrs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2CommonAttrs) ToProto() (*otg.Ospfv2CommonAttrs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2CommonAttrs) FromProto(msg *otg.Ospfv2CommonAttrs) (Ospfv2CommonAttrs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2CommonAttrs) ToPbText() (string, error) {
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

func (m *unMarshalospfv2CommonAttrs) FromPbText(value string) error {
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

func (m *marshalospfv2CommonAttrs) ToYaml() (string, error) {
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

func (m *unMarshalospfv2CommonAttrs) FromYaml(value string) error {
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

func (m *marshalospfv2CommonAttrs) ToJson() (string, error) {
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

func (m *unMarshalospfv2CommonAttrs) FromJson(value string) error {
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

func (obj *ospfv2CommonAttrs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2CommonAttrs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2CommonAttrs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2CommonAttrs) Clone() (Ospfv2CommonAttrs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2CommonAttrs()
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

// Ospfv2CommonAttrs is attributes in LSA Header.
type Ospfv2CommonAttrs interface {
	Validation
	// msg marshals Ospfv2CommonAttrs to protobuf object *otg.Ospfv2CommonAttrs
	// and doesn't set defaults
	msg() *otg.Ospfv2CommonAttrs
	// setMsg unmarshals Ospfv2CommonAttrs from protobuf object *otg.Ospfv2CommonAttrs
	// and doesn't set defaults
	setMsg(*otg.Ospfv2CommonAttrs) Ospfv2CommonAttrs
	// provides marshal interface
	Marshal() marshalOspfv2CommonAttrs
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2CommonAttrs
	// validate validates Ospfv2CommonAttrs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2CommonAttrs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LsaId returns string, set in Ospfv2CommonAttrs.
	LsaId() string
	// SetLsaId assigns string provided by user to Ospfv2CommonAttrs
	SetLsaId(value string) Ospfv2CommonAttrs
	// HasLsaId checks if LsaId has been set in Ospfv2CommonAttrs
	HasLsaId() bool
	// AdvertisingRouter returns string, set in Ospfv2CommonAttrs.
	AdvertisingRouter() string
	// SetAdvertisingRouter assigns string provided by user to Ospfv2CommonAttrs
	SetAdvertisingRouter(value string) Ospfv2CommonAttrs
	// HasAdvertisingRouter checks if AdvertisingRouter has been set in Ospfv2CommonAttrs
	HasAdvertisingRouter() bool
	// SequenceNumber returns uint32, set in Ospfv2CommonAttrs.
	SequenceNumber() uint32
	// SetSequenceNumber assigns uint32 provided by user to Ospfv2CommonAttrs
	SetSequenceNumber(value uint32) Ospfv2CommonAttrs
	// HasSequenceNumber checks if SequenceNumber has been set in Ospfv2CommonAttrs
	HasSequenceNumber() bool
	// Age returns uint32, set in Ospfv2CommonAttrs.
	Age() uint32
	// SetAge assigns uint32 provided by user to Ospfv2CommonAttrs
	SetAge(value uint32) Ospfv2CommonAttrs
	// HasAge checks if Age has been set in Ospfv2CommonAttrs
	HasAge() bool
	// OptionBits returns uint32, set in Ospfv2CommonAttrs.
	OptionBits() uint32
	// SetOptionBits assigns uint32 provided by user to Ospfv2CommonAttrs
	SetOptionBits(value uint32) Ospfv2CommonAttrs
	// HasOptionBits checks if OptionBits has been set in Ospfv2CommonAttrs
	HasOptionBits() bool
}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// LsaId returns a string
func (obj *ospfv2CommonAttrs) LsaId() string {

	return *obj.obj.LsaId

}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// LsaId returns a string
func (obj *ospfv2CommonAttrs) HasLsaId() bool {
	return obj.obj.LsaId != nil
}

// LSA ID in the IPv4 format. The Link State ID for the specified LSA type.
// SetLsaId sets the string value in the Ospfv2CommonAttrs object
func (obj *ospfv2CommonAttrs) SetLsaId(value string) Ospfv2CommonAttrs {

	obj.obj.LsaId = &value
	return obj
}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// AdvertisingRouter returns a string
func (obj *ospfv2CommonAttrs) AdvertisingRouter() string {

	return *obj.obj.AdvertisingRouter

}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// AdvertisingRouter returns a string
func (obj *ospfv2CommonAttrs) HasAdvertisingRouter() bool {
	return obj.obj.AdvertisingRouter != nil
}

// The router ID (in the IPv4 format) of the router that originated the LSA.
// SetAdvertisingRouter sets the string value in the Ospfv2CommonAttrs object
func (obj *ospfv2CommonAttrs) SetAdvertisingRouter(value string) Ospfv2CommonAttrs {

	obj.obj.AdvertisingRouter = &value
	return obj
}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SequenceNumber returns a uint32
func (obj *ospfv2CommonAttrs) SequenceNumber() uint32 {

	return *obj.obj.SequenceNumber

}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SequenceNumber returns a uint32
func (obj *ospfv2CommonAttrs) HasSequenceNumber() bool {
	return obj.obj.SequenceNumber != nil
}

// Sequence number to detect old and duplicate LSAs. The greater the sequence number the more recent the LSA.
// SetSequenceNumber sets the uint32 value in the Ospfv2CommonAttrs object
func (obj *ospfv2CommonAttrs) SetSequenceNumber(value uint32) Ospfv2CommonAttrs {

	obj.obj.SequenceNumber = &value
	return obj
}

// The time since the LSA's generation in seconds.
// Age returns a uint32
func (obj *ospfv2CommonAttrs) Age() uint32 {

	return *obj.obj.Age

}

// The time since the LSA's generation in seconds.
// Age returns a uint32
func (obj *ospfv2CommonAttrs) HasAge() bool {
	return obj.obj.Age != nil
}

// The time since the LSA's generation in seconds.
// SetAge sets the uint32 value in the Ospfv2CommonAttrs object
func (obj *ospfv2CommonAttrs) SetAge(value uint32) Ospfv2CommonAttrs {

	obj.obj.Age = &value
	return obj
}

// The optional bits.
// OptionBits returns a uint32
func (obj *ospfv2CommonAttrs) OptionBits() uint32 {

	return *obj.obj.OptionBits

}

// The optional bits.
// OptionBits returns a uint32
func (obj *ospfv2CommonAttrs) HasOptionBits() bool {
	return obj.obj.OptionBits != nil
}

// The optional bits.
// SetOptionBits sets the uint32 value in the Ospfv2CommonAttrs object
func (obj *ospfv2CommonAttrs) SetOptionBits(value uint32) Ospfv2CommonAttrs {

	obj.obj.OptionBits = &value
	return obj
}

func (obj *ospfv2CommonAttrs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LsaId != nil {

		err := obj.validateIpv4(obj.LsaId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2CommonAttrs.LsaId"))
		}

	}

	if obj.obj.AdvertisingRouter != nil {

		err := obj.validateIpv4(obj.AdvertisingRouter())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Ospfv2CommonAttrs.AdvertisingRouter"))
		}

	}

}

func (obj *ospfv2CommonAttrs) setDefault() {

}
