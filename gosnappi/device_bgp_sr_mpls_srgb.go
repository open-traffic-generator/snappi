package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBgpSrMplsSrgb *****
type deviceBgpSrMplsSrgb struct {
	validation
	obj          *otg.DeviceBgpSrMplsSrgb
	marshaller   marshalDeviceBgpSrMplsSrgb
	unMarshaller unMarshalDeviceBgpSrMplsSrgb
}

func NewDeviceBgpSrMplsSrgb() DeviceBgpSrMplsSrgb {
	obj := deviceBgpSrMplsSrgb{obj: &otg.DeviceBgpSrMplsSrgb{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBgpSrMplsSrgb) msg() *otg.DeviceBgpSrMplsSrgb {
	return obj.obj
}

func (obj *deviceBgpSrMplsSrgb) setMsg(msg *otg.DeviceBgpSrMplsSrgb) DeviceBgpSrMplsSrgb {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBgpSrMplsSrgb struct {
	obj *deviceBgpSrMplsSrgb
}

type marshalDeviceBgpSrMplsSrgb interface {
	// ToProto marshals DeviceBgpSrMplsSrgb to protobuf object *otg.DeviceBgpSrMplsSrgb
	ToProto() (*otg.DeviceBgpSrMplsSrgb, error)
	// ToPbText marshals DeviceBgpSrMplsSrgb to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBgpSrMplsSrgb to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBgpSrMplsSrgb to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBgpSrMplsSrgb struct {
	obj *deviceBgpSrMplsSrgb
}

type unMarshalDeviceBgpSrMplsSrgb interface {
	// FromProto unmarshals DeviceBgpSrMplsSrgb from protobuf object *otg.DeviceBgpSrMplsSrgb
	FromProto(msg *otg.DeviceBgpSrMplsSrgb) (DeviceBgpSrMplsSrgb, error)
	// FromPbText unmarshals DeviceBgpSrMplsSrgb from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBgpSrMplsSrgb from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBgpSrMplsSrgb from JSON text
	FromJson(value string) error
}

func (obj *deviceBgpSrMplsSrgb) Marshal() marshalDeviceBgpSrMplsSrgb {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBgpSrMplsSrgb{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBgpSrMplsSrgb) Unmarshal() unMarshalDeviceBgpSrMplsSrgb {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBgpSrMplsSrgb{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBgpSrMplsSrgb) ToProto() (*otg.DeviceBgpSrMplsSrgb, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBgpSrMplsSrgb) FromProto(msg *otg.DeviceBgpSrMplsSrgb) (DeviceBgpSrMplsSrgb, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBgpSrMplsSrgb) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBgpSrMplsSrgb) FromPbText(value string) error {
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

func (m *marshaldeviceBgpSrMplsSrgb) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBgpSrMplsSrgb) FromYaml(value string) error {
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

func (m *marshaldeviceBgpSrMplsSrgb) ToJson() (string, error) {
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

func (m *unMarshaldeviceBgpSrMplsSrgb) FromJson(value string) error {
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

func (obj *deviceBgpSrMplsSrgb) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBgpSrMplsSrgb) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBgpSrMplsSrgb) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBgpSrMplsSrgb) Clone() (DeviceBgpSrMplsSrgb, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBgpSrMplsSrgb()
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

// DeviceBgpSrMplsSrgb is this contains the propeties of Segment Routing Global Block (SRGB) Sub-Tlv (3) that is advertised under PREFIX_SID Tlv (40). The set of global segments in the SR domain. If a node participates in multiple SR domains, there is one SRGB for each SR domain.  In SR-MPLS, SRGB is a local property of a node and identifies the set of local labels reserved for global segments. Reference: https://datatracker.ietf.org/doc/html/rfc8402
type DeviceBgpSrMplsSrgb interface {
	Validation
	// msg marshals DeviceBgpSrMplsSrgb to protobuf object *otg.DeviceBgpSrMplsSrgb
	// and doesn't set defaults
	msg() *otg.DeviceBgpSrMplsSrgb
	// setMsg unmarshals DeviceBgpSrMplsSrgb from protobuf object *otg.DeviceBgpSrMplsSrgb
	// and doesn't set defaults
	setMsg(*otg.DeviceBgpSrMplsSrgb) DeviceBgpSrMplsSrgb
	// provides marshal interface
	Marshal() marshalDeviceBgpSrMplsSrgb
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBgpSrMplsSrgb
	// validate validates DeviceBgpSrMplsSrgb
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBgpSrMplsSrgb, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StartingLabel returns uint32, set in DeviceBgpSrMplsSrgb.
	StartingLabel() uint32
	// SetStartingLabel assigns uint32 provided by user to DeviceBgpSrMplsSrgb
	SetStartingLabel(value uint32) DeviceBgpSrMplsSrgb
	// HasStartingLabel checks if StartingLabel has been set in DeviceBgpSrMplsSrgb
	HasStartingLabel() bool
	// Range returns uint32, set in DeviceBgpSrMplsSrgb.
	Range() uint32
	// SetRange assigns uint32 provided by user to DeviceBgpSrMplsSrgb
	SetRange(value uint32) DeviceBgpSrMplsSrgb
	// HasRange checks if Range has been set in DeviceBgpSrMplsSrgb
	HasRange() bool
	// AdvertiseOnce returns bool, set in DeviceBgpSrMplsSrgb.
	AdvertiseOnce() bool
	// SetAdvertiseOnce assigns bool provided by user to DeviceBgpSrMplsSrgb
	SetAdvertiseOnce(value bool) DeviceBgpSrMplsSrgb
	// HasAdvertiseOnce checks if AdvertiseOnce has been set in DeviceBgpSrMplsSrgb
	HasAdvertiseOnce() bool
}

// The first Prefix SID Label value of the SRGB while the range contains the number of SRGB elements.
// StartingLabel returns a uint32
func (obj *deviceBgpSrMplsSrgb) StartingLabel() uint32 {

	return *obj.obj.StartingLabel

}

// The first Prefix SID Label value of the SRGB while the range contains the number of SRGB elements.
// StartingLabel returns a uint32
func (obj *deviceBgpSrMplsSrgb) HasStartingLabel() bool {
	return obj.obj.StartingLabel != nil
}

// The first Prefix SID Label value of the SRGB while the range contains the number of SRGB elements.
// SetStartingLabel sets the uint32 value in the DeviceBgpSrMplsSrgb object
func (obj *deviceBgpSrMplsSrgb) SetStartingLabel(value uint32) DeviceBgpSrMplsSrgb {

	obj.obj.StartingLabel = &value
	return obj
}

// This represents the number of Prefix SID Label in a SRGB range.
// Range returns a uint32
func (obj *deviceBgpSrMplsSrgb) Range() uint32 {

	return *obj.obj.Range

}

// This represents the number of Prefix SID Label in a SRGB range.
// Range returns a uint32
func (obj *deviceBgpSrMplsSrgb) HasRange() bool {
	return obj.obj.Range != nil
}

// This represents the number of Prefix SID Label in a SRGB range.
// SetRange sets the uint32 value in the DeviceBgpSrMplsSrgb object
func (obj *deviceBgpSrMplsSrgb) SetRange(value uint32) DeviceBgpSrMplsSrgb {

	obj.obj.Range = &value
	return obj
}

// If it is true the SRGB will be advertise only in the first route under that is advertised under PREFIX_SID Tlv 40. Otherwise SRGB Sub Tlv will be advertised with the advertisement of each prefix.
// AdvertiseOnce returns a bool
func (obj *deviceBgpSrMplsSrgb) AdvertiseOnce() bool {

	return *obj.obj.AdvertiseOnce

}

// If it is true the SRGB will be advertise only in the first route under that is advertised under PREFIX_SID Tlv 40. Otherwise SRGB Sub Tlv will be advertised with the advertisement of each prefix.
// AdvertiseOnce returns a bool
func (obj *deviceBgpSrMplsSrgb) HasAdvertiseOnce() bool {
	return obj.obj.AdvertiseOnce != nil
}

// If it is true the SRGB will be advertise only in the first route under that is advertised under PREFIX_SID Tlv 40. Otherwise SRGB Sub Tlv will be advertised with the advertisement of each prefix.
// SetAdvertiseOnce sets the bool value in the DeviceBgpSrMplsSrgb object
func (obj *deviceBgpSrMplsSrgb) SetAdvertiseOnce(value bool) DeviceBgpSrMplsSrgb {

	obj.obj.AdvertiseOnce = &value
	return obj
}

func (obj *deviceBgpSrMplsSrgb) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StartingLabel != nil {

		if *obj.obj.StartingLabel < 1 || *obj.obj.StartingLabel > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceBgpSrMplsSrgb.StartingLabel <= 1048575 but Got %d", *obj.obj.StartingLabel))
		}

	}

	if obj.obj.Range != nil {

		if *obj.obj.Range < 1 || *obj.obj.Range > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= DeviceBgpSrMplsSrgb.Range <= 1048575 but Got %d", *obj.obj.Range))
		}

	}

}

func (obj *deviceBgpSrMplsSrgb) setDefault() {
	if obj.obj.StartingLabel == nil {
		obj.SetStartingLabel(16000)
	}
	if obj.obj.Range == nil {
		obj.SetRange(8000)
	}
	if obj.obj.AdvertiseOnce == nil {
		obj.SetAdvertiseOnce(true)
	}

}
