package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSrMplsSid *****
type bgpSrteSrMplsSid struct {
	validation
	obj          *otg.BgpSrteSrMplsSid
	marshaller   marshalBgpSrteSrMplsSid
	unMarshaller unMarshalBgpSrteSrMplsSid
}

func NewBgpSrteSrMplsSid() BgpSrteSrMplsSid {
	obj := bgpSrteSrMplsSid{obj: &otg.BgpSrteSrMplsSid{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSrMplsSid) msg() *otg.BgpSrteSrMplsSid {
	return obj.obj
}

func (obj *bgpSrteSrMplsSid) setMsg(msg *otg.BgpSrteSrMplsSid) BgpSrteSrMplsSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSrMplsSid struct {
	obj *bgpSrteSrMplsSid
}

type marshalBgpSrteSrMplsSid interface {
	// ToProto marshals BgpSrteSrMplsSid to protobuf object *otg.BgpSrteSrMplsSid
	ToProto() (*otg.BgpSrteSrMplsSid, error)
	// ToPbText marshals BgpSrteSrMplsSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSrMplsSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSrMplsSid to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSrMplsSid struct {
	obj *bgpSrteSrMplsSid
}

type unMarshalBgpSrteSrMplsSid interface {
	// FromProto unmarshals BgpSrteSrMplsSid from protobuf object *otg.BgpSrteSrMplsSid
	FromProto(msg *otg.BgpSrteSrMplsSid) (BgpSrteSrMplsSid, error)
	// FromPbText unmarshals BgpSrteSrMplsSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSrMplsSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSrMplsSid from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSrMplsSid) Marshal() marshalBgpSrteSrMplsSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSrMplsSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSrMplsSid) Unmarshal() unMarshalBgpSrteSrMplsSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSrMplsSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSrMplsSid) ToProto() (*otg.BgpSrteSrMplsSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSrMplsSid) FromProto(msg *otg.BgpSrteSrMplsSid) (BgpSrteSrMplsSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSrMplsSid) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSrMplsSid) FromPbText(value string) error {
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

func (m *marshalbgpSrteSrMplsSid) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSrMplsSid) FromYaml(value string) error {
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

func (m *marshalbgpSrteSrMplsSid) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSrMplsSid) FromJson(value string) error {
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

func (obj *bgpSrteSrMplsSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSrMplsSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSrMplsSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSrMplsSid) Clone() (BgpSrteSrMplsSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSrMplsSid()
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

// BgpSrteSrMplsSid is configuration for SR-MPLS with Label, TC, Bottom-of-Stack and TTL.
type BgpSrteSrMplsSid interface {
	Validation
	// msg marshals BgpSrteSrMplsSid to protobuf object *otg.BgpSrteSrMplsSid
	// and doesn't set defaults
	msg() *otg.BgpSrteSrMplsSid
	// setMsg unmarshals BgpSrteSrMplsSid from protobuf object *otg.BgpSrteSrMplsSid
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSrMplsSid) BgpSrteSrMplsSid
	// provides marshal interface
	Marshal() marshalBgpSrteSrMplsSid
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSrMplsSid
	// validate validates BgpSrteSrMplsSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSrMplsSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Label returns uint32, set in BgpSrteSrMplsSid.
	Label() uint32
	// SetLabel assigns uint32 provided by user to BgpSrteSrMplsSid
	SetLabel(value uint32) BgpSrteSrMplsSid
	// HasLabel checks if Label has been set in BgpSrteSrMplsSid
	HasLabel() bool
	// Tc returns uint32, set in BgpSrteSrMplsSid.
	Tc() uint32
	// SetTc assigns uint32 provided by user to BgpSrteSrMplsSid
	SetTc(value uint32) BgpSrteSrMplsSid
	// HasTc checks if Tc has been set in BgpSrteSrMplsSid
	HasTc() bool
	// SBit returns bool, set in BgpSrteSrMplsSid.
	SBit() bool
	// SetSBit assigns bool provided by user to BgpSrteSrMplsSid
	SetSBit(value bool) BgpSrteSrMplsSid
	// HasSBit checks if SBit has been set in BgpSrteSrMplsSid
	HasSBit() bool
	// Ttl returns uint32, set in BgpSrteSrMplsSid.
	Ttl() uint32
	// SetTtl assigns uint32 provided by user to BgpSrteSrMplsSid
	SetTtl(value uint32) BgpSrteSrMplsSid
	// HasTtl checks if Ttl has been set in BgpSrteSrMplsSid
	HasTtl() bool
}

// Label value in [0, 2^20 -1].
// Label returns a uint32
func (obj *bgpSrteSrMplsSid) Label() uint32 {

	return *obj.obj.Label

}

// Label value in [0, 2^20 -1].
// Label returns a uint32
func (obj *bgpSrteSrMplsSid) HasLabel() bool {
	return obj.obj.Label != nil
}

// Label value in [0, 2^20 -1].
// SetLabel sets the uint32 value in the BgpSrteSrMplsSid object
func (obj *bgpSrteSrMplsSid) SetLabel(value uint32) BgpSrteSrMplsSid {

	obj.obj.Label = &value
	return obj
}

// Traffic class in bits.
// Tc returns a uint32
func (obj *bgpSrteSrMplsSid) Tc() uint32 {

	return *obj.obj.Tc

}

// Traffic class in bits.
// Tc returns a uint32
func (obj *bgpSrteSrMplsSid) HasTc() bool {
	return obj.obj.Tc != nil
}

// Traffic class in bits.
// SetTc sets the uint32 value in the BgpSrteSrMplsSid object
func (obj *bgpSrteSrMplsSid) SetTc(value uint32) BgpSrteSrMplsSid {

	obj.obj.Tc = &value
	return obj
}

// Bottom-of-Stack bit.
// SBit returns a bool
func (obj *bgpSrteSrMplsSid) SBit() bool {

	return *obj.obj.SBit

}

// Bottom-of-Stack bit.
// SBit returns a bool
func (obj *bgpSrteSrMplsSid) HasSBit() bool {
	return obj.obj.SBit != nil
}

// Bottom-of-Stack bit.
// SetSBit sets the bool value in the BgpSrteSrMplsSid object
func (obj *bgpSrteSrMplsSid) SetSBit(value bool) BgpSrteSrMplsSid {

	obj.obj.SBit = &value
	return obj
}

// Time To Live.
// Ttl returns a uint32
func (obj *bgpSrteSrMplsSid) Ttl() uint32 {

	return *obj.obj.Ttl

}

// Time To Live.
// Ttl returns a uint32
func (obj *bgpSrteSrMplsSid) HasTtl() bool {
	return obj.obj.Ttl != nil
}

// Time To Live.
// SetTtl sets the uint32 value in the BgpSrteSrMplsSid object
func (obj *bgpSrteSrMplsSid) SetTtl(value uint32) BgpSrteSrMplsSid {

	obj.obj.Ttl = &value
	return obj
}

func (obj *bgpSrteSrMplsSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSrMplsSid.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

	}

	if obj.obj.Tc != nil {

		if *obj.obj.Tc > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSrMplsSid.Tc <= 7 but Got %d", *obj.obj.Tc))
		}

	}

	if obj.obj.Ttl != nil {

		if *obj.obj.Ttl > 225 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSrMplsSid.Ttl <= 225 but Got %d", *obj.obj.Ttl))
		}

	}

}

func (obj *bgpSrteSrMplsSid) setDefault() {

}
