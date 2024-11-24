package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSidMpls *****
type bgpAttributesSidMpls struct {
	validation
	obj          *otg.BgpAttributesSidMpls
	marshaller   marshalBgpAttributesSidMpls
	unMarshaller unMarshalBgpAttributesSidMpls
}

func NewBgpAttributesSidMpls() BgpAttributesSidMpls {
	obj := bgpAttributesSidMpls{obj: &otg.BgpAttributesSidMpls{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSidMpls) msg() *otg.BgpAttributesSidMpls {
	return obj.obj
}

func (obj *bgpAttributesSidMpls) setMsg(msg *otg.BgpAttributesSidMpls) BgpAttributesSidMpls {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSidMpls struct {
	obj *bgpAttributesSidMpls
}

type marshalBgpAttributesSidMpls interface {
	// ToProto marshals BgpAttributesSidMpls to protobuf object *otg.BgpAttributesSidMpls
	ToProto() (*otg.BgpAttributesSidMpls, error)
	// ToPbText marshals BgpAttributesSidMpls to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSidMpls to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSidMpls to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSidMpls to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSidMpls struct {
	obj *bgpAttributesSidMpls
}

type unMarshalBgpAttributesSidMpls interface {
	// FromProto unmarshals BgpAttributesSidMpls from protobuf object *otg.BgpAttributesSidMpls
	FromProto(msg *otg.BgpAttributesSidMpls) (BgpAttributesSidMpls, error)
	// FromPbText unmarshals BgpAttributesSidMpls from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSidMpls from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSidMpls from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSidMpls) Marshal() marshalBgpAttributesSidMpls {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSidMpls{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSidMpls) Unmarshal() unMarshalBgpAttributesSidMpls {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSidMpls{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSidMpls) ToProto() (*otg.BgpAttributesSidMpls, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSidMpls) FromProto(msg *otg.BgpAttributesSidMpls) (BgpAttributesSidMpls, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSidMpls) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSidMpls) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSidMpls) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSidMpls) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSidMpls) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSidMpls) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSidMpls) FromJson(value string) error {
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

func (obj *bgpAttributesSidMpls) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSidMpls) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSidMpls) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSidMpls) Clone() (BgpAttributesSidMpls, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSidMpls()
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

// BgpAttributesSidMpls is this carries a 20 bit Multi Protocol Label Switching alongwith 3 bits traffic class, 1 bit indicating presence
// or absence of Bottom-Of-Stack and 8 bits carrying the Time to Live value.
type BgpAttributesSidMpls interface {
	Validation
	// msg marshals BgpAttributesSidMpls to protobuf object *otg.BgpAttributesSidMpls
	// and doesn't set defaults
	msg() *otg.BgpAttributesSidMpls
	// setMsg unmarshals BgpAttributesSidMpls from protobuf object *otg.BgpAttributesSidMpls
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSidMpls) BgpAttributesSidMpls
	// provides marshal interface
	Marshal() marshalBgpAttributesSidMpls
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSidMpls
	// validate validates BgpAttributesSidMpls
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSidMpls, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Label returns uint32, set in BgpAttributesSidMpls.
	Label() uint32
	// SetLabel assigns uint32 provided by user to BgpAttributesSidMpls
	SetLabel(value uint32) BgpAttributesSidMpls
	// HasLabel checks if Label has been set in BgpAttributesSidMpls
	HasLabel() bool
	// TrafficClass returns uint32, set in BgpAttributesSidMpls.
	TrafficClass() uint32
	// SetTrafficClass assigns uint32 provided by user to BgpAttributesSidMpls
	SetTrafficClass(value uint32) BgpAttributesSidMpls
	// HasTrafficClass checks if TrafficClass has been set in BgpAttributesSidMpls
	HasTrafficClass() bool
	// FlagBos returns bool, set in BgpAttributesSidMpls.
	FlagBos() bool
	// SetFlagBos assigns bool provided by user to BgpAttributesSidMpls
	SetFlagBos(value bool) BgpAttributesSidMpls
	// HasFlagBos checks if FlagBos has been set in BgpAttributesSidMpls
	HasFlagBos() bool
	// Ttl returns uint32, set in BgpAttributesSidMpls.
	Ttl() uint32
	// SetTtl assigns uint32 provided by user to BgpAttributesSidMpls
	SetTtl(value uint32) BgpAttributesSidMpls
	// HasTtl checks if Ttl has been set in BgpAttributesSidMpls
	HasTtl() bool
}

// 20 bit MPLS Label value.
// Label returns a uint32
func (obj *bgpAttributesSidMpls) Label() uint32 {

	return *obj.obj.Label

}

// 20 bit MPLS Label value.
// Label returns a uint32
func (obj *bgpAttributesSidMpls) HasLabel() bool {
	return obj.obj.Label != nil
}

// 20 bit MPLS Label value.
// SetLabel sets the uint32 value in the BgpAttributesSidMpls object
func (obj *bgpAttributesSidMpls) SetLabel(value uint32) BgpAttributesSidMpls {

	obj.obj.Label = &value
	return obj
}

// 3 bits of Traffic Class.
// TrafficClass returns a uint32
func (obj *bgpAttributesSidMpls) TrafficClass() uint32 {

	return *obj.obj.TrafficClass

}

// 3 bits of Traffic Class.
// TrafficClass returns a uint32
func (obj *bgpAttributesSidMpls) HasTrafficClass() bool {
	return obj.obj.TrafficClass != nil
}

// 3 bits of Traffic Class.
// SetTrafficClass sets the uint32 value in the BgpAttributesSidMpls object
func (obj *bgpAttributesSidMpls) SetTrafficClass(value uint32) BgpAttributesSidMpls {

	obj.obj.TrafficClass = &value
	return obj
}

// Bottom of Stack
// FlagBos returns a bool
func (obj *bgpAttributesSidMpls) FlagBos() bool {

	return *obj.obj.FlagBos

}

// Bottom of Stack
// FlagBos returns a bool
func (obj *bgpAttributesSidMpls) HasFlagBos() bool {
	return obj.obj.FlagBos != nil
}

// Bottom of Stack
// SetFlagBos sets the bool value in the BgpAttributesSidMpls object
func (obj *bgpAttributesSidMpls) SetFlagBos(value bool) BgpAttributesSidMpls {

	obj.obj.FlagBos = &value
	return obj
}

// 8 bits Time to Live
// Ttl returns a uint32
func (obj *bgpAttributesSidMpls) Ttl() uint32 {

	return *obj.obj.Ttl

}

// 8 bits Time to Live
// Ttl returns a uint32
func (obj *bgpAttributesSidMpls) HasTtl() bool {
	return obj.obj.Ttl != nil
}

// 8 bits Time to Live
// SetTtl sets the uint32 value in the BgpAttributesSidMpls object
func (obj *bgpAttributesSidMpls) SetTtl(value uint32) BgpAttributesSidMpls {

	obj.obj.Ttl = &value
	return obj
}

func (obj *bgpAttributesSidMpls) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Label != nil {

		if *obj.obj.Label > 1048576 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSidMpls.Label <= 1048576 but Got %d", *obj.obj.Label))
		}

	}

	if obj.obj.TrafficClass != nil {

		if *obj.obj.TrafficClass > 7 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSidMpls.TrafficClass <= 7 but Got %d", *obj.obj.TrafficClass))
		}

	}

	if obj.obj.Ttl != nil {

		if *obj.obj.Ttl > 63 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpAttributesSidMpls.Ttl <= 63 but Got %d", *obj.obj.Ttl))
		}

	}

}

func (obj *bgpAttributesSidMpls) setDefault() {
	if obj.obj.Label == nil {
		obj.SetLabel(16)
	}
	if obj.obj.TrafficClass == nil {
		obj.SetTrafficClass(0)
	}
	if obj.obj.FlagBos == nil {
		obj.SetFlagBos(true)
	}
	if obj.obj.Ttl == nil {
		obj.SetTtl(63)
	}

}
