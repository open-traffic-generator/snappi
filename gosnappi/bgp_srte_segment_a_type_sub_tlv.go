package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentATypeSubTlv *****
type bgpSrteSegmentATypeSubTlv struct {
	validation
	obj          *otg.BgpSrteSegmentATypeSubTlv
	marshaller   marshalBgpSrteSegmentATypeSubTlv
	unMarshaller unMarshalBgpSrteSegmentATypeSubTlv
}

func NewBgpSrteSegmentATypeSubTlv() BgpSrteSegmentATypeSubTlv {
	obj := bgpSrteSegmentATypeSubTlv{obj: &otg.BgpSrteSegmentATypeSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentATypeSubTlv) msg() *otg.BgpSrteSegmentATypeSubTlv {
	return obj.obj
}

func (obj *bgpSrteSegmentATypeSubTlv) setMsg(msg *otg.BgpSrteSegmentATypeSubTlv) BgpSrteSegmentATypeSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentATypeSubTlv struct {
	obj *bgpSrteSegmentATypeSubTlv
}

type marshalBgpSrteSegmentATypeSubTlv interface {
	// ToProto marshals BgpSrteSegmentATypeSubTlv to protobuf object *otg.BgpSrteSegmentATypeSubTlv
	ToProto() (*otg.BgpSrteSegmentATypeSubTlv, error)
	// ToPbText marshals BgpSrteSegmentATypeSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentATypeSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentATypeSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentATypeSubTlv struct {
	obj *bgpSrteSegmentATypeSubTlv
}

type unMarshalBgpSrteSegmentATypeSubTlv interface {
	// FromProto unmarshals BgpSrteSegmentATypeSubTlv from protobuf object *otg.BgpSrteSegmentATypeSubTlv
	FromProto(msg *otg.BgpSrteSegmentATypeSubTlv) (BgpSrteSegmentATypeSubTlv, error)
	// FromPbText unmarshals BgpSrteSegmentATypeSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentATypeSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentATypeSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentATypeSubTlv) Marshal() marshalBgpSrteSegmentATypeSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentATypeSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentATypeSubTlv) Unmarshal() unMarshalBgpSrteSegmentATypeSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentATypeSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentATypeSubTlv) ToProto() (*otg.BgpSrteSegmentATypeSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentATypeSubTlv) FromProto(msg *otg.BgpSrteSegmentATypeSubTlv) (BgpSrteSegmentATypeSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentATypeSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentATypeSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentATypeSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentATypeSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentATypeSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentATypeSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteSegmentATypeSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentATypeSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentATypeSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentATypeSubTlv) Clone() (BgpSrteSegmentATypeSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentATypeSubTlv()
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

// BgpSrteSegmentATypeSubTlv is type  A: SID only, in the form of MPLS Label.
type BgpSrteSegmentATypeSubTlv interface {
	Validation
	// msg marshals BgpSrteSegmentATypeSubTlv to protobuf object *otg.BgpSrteSegmentATypeSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentATypeSubTlv
	// setMsg unmarshals BgpSrteSegmentATypeSubTlv from protobuf object *otg.BgpSrteSegmentATypeSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentATypeSubTlv) BgpSrteSegmentATypeSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentATypeSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentATypeSubTlv
	// validate validates BgpSrteSegmentATypeSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentATypeSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns string, set in BgpSrteSegmentATypeSubTlv.
	Flags() string
	// SetFlags assigns string provided by user to BgpSrteSegmentATypeSubTlv
	SetFlags(value string) BgpSrteSegmentATypeSubTlv
	// HasFlags checks if Flags has been set in BgpSrteSegmentATypeSubTlv
	HasFlags() bool
	// Label returns uint32, set in BgpSrteSegmentATypeSubTlv.
	Label() uint32
	// SetLabel assigns uint32 provided by user to BgpSrteSegmentATypeSubTlv
	SetLabel(value uint32) BgpSrteSegmentATypeSubTlv
	// HasLabel checks if Label has been set in BgpSrteSegmentATypeSubTlv
	HasLabel() bool
	// Tc returns uint32, set in BgpSrteSegmentATypeSubTlv.
	Tc() uint32
	// SetTc assigns uint32 provided by user to BgpSrteSegmentATypeSubTlv
	SetTc(value uint32) BgpSrteSegmentATypeSubTlv
	// HasTc checks if Tc has been set in BgpSrteSegmentATypeSubTlv
	HasTc() bool
	// SBit returns bool, set in BgpSrteSegmentATypeSubTlv.
	SBit() bool
	// SetSBit assigns bool provided by user to BgpSrteSegmentATypeSubTlv
	SetSBit(value bool) BgpSrteSegmentATypeSubTlv
	// HasSBit checks if SBit has been set in BgpSrteSegmentATypeSubTlv
	HasSBit() bool
	// Ttl returns uint32, set in BgpSrteSegmentATypeSubTlv.
	Ttl() uint32
	// SetTtl assigns uint32 provided by user to BgpSrteSegmentATypeSubTlv
	SetTtl(value uint32) BgpSrteSegmentATypeSubTlv
	// HasTtl checks if Ttl has been set in BgpSrteSegmentATypeSubTlv
	HasTtl() bool
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentATypeSubTlv) Flags() string {

	return *obj.obj.Flags

}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// Flags returns a string
func (obj *bgpSrteSegmentATypeSubTlv) HasFlags() bool {
	return obj.obj.Flags != nil
}

// One octet bitmap for flags including V-Flag, A-Flag, S-Flag, B-Flag etc. as defined in https://datatracker.ietf.org/doc/html/draft-ietf-idr-segment-routing-te-policy-13#section-2.4.4.2.12
// SetFlags sets the string value in the BgpSrteSegmentATypeSubTlv object
func (obj *bgpSrteSegmentATypeSubTlv) SetFlags(value string) BgpSrteSegmentATypeSubTlv {

	obj.obj.Flags = &value
	return obj
}

// Label value in [0, 2^20 -1].
// Label returns a uint32
func (obj *bgpSrteSegmentATypeSubTlv) Label() uint32 {

	return *obj.obj.Label

}

// Label value in [0, 2^20 -1].
// Label returns a uint32
func (obj *bgpSrteSegmentATypeSubTlv) HasLabel() bool {
	return obj.obj.Label != nil
}

// Label value in [0, 2^20 -1].
// SetLabel sets the uint32 value in the BgpSrteSegmentATypeSubTlv object
func (obj *bgpSrteSegmentATypeSubTlv) SetLabel(value uint32) BgpSrteSegmentATypeSubTlv {

	obj.obj.Label = &value
	return obj
}

// Traffic class in bits.
// Tc returns a uint32
func (obj *bgpSrteSegmentATypeSubTlv) Tc() uint32 {

	return *obj.obj.Tc

}

// Traffic class in bits.
// Tc returns a uint32
func (obj *bgpSrteSegmentATypeSubTlv) HasTc() bool {
	return obj.obj.Tc != nil
}

// Traffic class in bits.
// SetTc sets the uint32 value in the BgpSrteSegmentATypeSubTlv object
func (obj *bgpSrteSegmentATypeSubTlv) SetTc(value uint32) BgpSrteSegmentATypeSubTlv {

	obj.obj.Tc = &value
	return obj
}

// Bottom-of-Stack bit.
// SBit returns a bool
func (obj *bgpSrteSegmentATypeSubTlv) SBit() bool {

	return *obj.obj.SBit

}

// Bottom-of-Stack bit.
// SBit returns a bool
func (obj *bgpSrteSegmentATypeSubTlv) HasSBit() bool {
	return obj.obj.SBit != nil
}

// Bottom-of-Stack bit.
// SetSBit sets the bool value in the BgpSrteSegmentATypeSubTlv object
func (obj *bgpSrteSegmentATypeSubTlv) SetSBit(value bool) BgpSrteSegmentATypeSubTlv {

	obj.obj.SBit = &value
	return obj
}

// Time To Live.
// Ttl returns a uint32
func (obj *bgpSrteSegmentATypeSubTlv) Ttl() uint32 {

	return *obj.obj.Ttl

}

// Time To Live.
// Ttl returns a uint32
func (obj *bgpSrteSegmentATypeSubTlv) HasTtl() bool {
	return obj.obj.Ttl != nil
}

// Time To Live.
// SetTtl sets the uint32 value in the BgpSrteSegmentATypeSubTlv object
func (obj *bgpSrteSegmentATypeSubTlv) SetTtl(value uint32) BgpSrteSegmentATypeSubTlv {

	obj.obj.Ttl = &value
	return obj
}

func (obj *bgpSrteSegmentATypeSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		err := obj.validateHex(obj.Flags())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteSegmentATypeSubTlv.Flags"))
		}

	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSegmentATypeSubTlv.Label <= 1048575 but Got %d", *obj.obj.Label))
		}

	}

	if obj.obj.Tc != nil {

		if *obj.obj.Tc > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSegmentATypeSubTlv.Tc <= 7 but Got %d", *obj.obj.Tc))
		}

	}

	if obj.obj.Ttl != nil {

		if *obj.obj.Ttl > 225 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrteSegmentATypeSubTlv.Ttl <= 225 but Got %d", *obj.obj.Ttl))
		}

	}

}

func (obj *bgpSrteSegmentATypeSubTlv) setDefault() {

}
