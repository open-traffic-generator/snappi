package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpV6SegmentRouting *****
type bgpV6SegmentRouting struct {
	validation
	obj          *otg.BgpV6SegmentRouting
	marshaller   marshalBgpV6SegmentRouting
	unMarshaller unMarshalBgpV6SegmentRouting
}

func NewBgpV6SegmentRouting() BgpV6SegmentRouting {
	obj := bgpV6SegmentRouting{obj: &otg.BgpV6SegmentRouting{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpV6SegmentRouting) msg() *otg.BgpV6SegmentRouting {
	return obj.obj
}

func (obj *bgpV6SegmentRouting) setMsg(msg *otg.BgpV6SegmentRouting) BgpV6SegmentRouting {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpV6SegmentRouting struct {
	obj *bgpV6SegmentRouting
}

type marshalBgpV6SegmentRouting interface {
	// ToProto marshals BgpV6SegmentRouting to protobuf object *otg.BgpV6SegmentRouting
	ToProto() (*otg.BgpV6SegmentRouting, error)
	// ToPbText marshals BgpV6SegmentRouting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpV6SegmentRouting to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpV6SegmentRouting to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpV6SegmentRouting to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpV6SegmentRouting struct {
	obj *bgpV6SegmentRouting
}

type unMarshalBgpV6SegmentRouting interface {
	// FromProto unmarshals BgpV6SegmentRouting from protobuf object *otg.BgpV6SegmentRouting
	FromProto(msg *otg.BgpV6SegmentRouting) (BgpV6SegmentRouting, error)
	// FromPbText unmarshals BgpV6SegmentRouting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpV6SegmentRouting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpV6SegmentRouting from JSON text
	FromJson(value string) error
}

func (obj *bgpV6SegmentRouting) Marshal() marshalBgpV6SegmentRouting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpV6SegmentRouting{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpV6SegmentRouting) Unmarshal() unMarshalBgpV6SegmentRouting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpV6SegmentRouting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpV6SegmentRouting) ToProto() (*otg.BgpV6SegmentRouting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpV6SegmentRouting) FromProto(msg *otg.BgpV6SegmentRouting) (BgpV6SegmentRouting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpV6SegmentRouting) ToPbText() (string, error) {
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

func (m *unMarshalbgpV6SegmentRouting) FromPbText(value string) error {
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

func (m *marshalbgpV6SegmentRouting) ToYaml() (string, error) {
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

func (m *unMarshalbgpV6SegmentRouting) FromYaml(value string) error {
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

func (m *marshalbgpV6SegmentRouting) ToJsonRaw() (string, error) {
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

func (m *marshalbgpV6SegmentRouting) ToJson() (string, error) {
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

func (m *unMarshalbgpV6SegmentRouting) FromJson(value string) error {
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

func (obj *bgpV6SegmentRouting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpV6SegmentRouting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpV6SegmentRouting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpV6SegmentRouting) Clone() (BgpV6SegmentRouting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpV6SegmentRouting()
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

// BgpV6SegmentRouting is configuration for BGPv6 segment routing settings.
type BgpV6SegmentRouting interface {
	Validation
	// msg marshals BgpV6SegmentRouting to protobuf object *otg.BgpV6SegmentRouting
	// and doesn't set defaults
	msg() *otg.BgpV6SegmentRouting
	// setMsg unmarshals BgpV6SegmentRouting from protobuf object *otg.BgpV6SegmentRouting
	// and doesn't set defaults
	setMsg(*otg.BgpV6SegmentRouting) BgpV6SegmentRouting
	// provides marshal interface
	Marshal() marshalBgpV6SegmentRouting
	// provides unmarshal interface
	Unmarshal() unMarshalBgpV6SegmentRouting
	// validate validates BgpV6SegmentRouting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpV6SegmentRouting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IngressSupportsVpn returns bool, set in BgpV6SegmentRouting.
	IngressSupportsVpn() bool
	// SetIngressSupportsVpn assigns bool provided by user to BgpV6SegmentRouting
	SetIngressSupportsVpn(value bool) BgpV6SegmentRouting
	// HasIngressSupportsVpn checks if IngressSupportsVpn has been set in BgpV6SegmentRouting
	HasIngressSupportsVpn() bool
	// ReducedEncapsulation returns bool, set in BgpV6SegmentRouting.
	ReducedEncapsulation() bool
	// SetReducedEncapsulation assigns bool provided by user to BgpV6SegmentRouting
	SetReducedEncapsulation(value bool) BgpV6SegmentRouting
	// HasReducedEncapsulation checks if ReducedEncapsulation has been set in BgpV6SegmentRouting
	HasReducedEncapsulation() bool
	// CopyTimeToLive returns bool, set in BgpV6SegmentRouting.
	CopyTimeToLive() bool
	// SetCopyTimeToLive assigns bool provided by user to BgpV6SegmentRouting
	SetCopyTimeToLive(value bool) BgpV6SegmentRouting
	// HasCopyTimeToLive checks if CopyTimeToLive has been set in BgpV6SegmentRouting
	HasCopyTimeToLive() bool
	// TimeToLive returns uint32, set in BgpV6SegmentRouting.
	TimeToLive() uint32
	// SetTimeToLive assigns uint32 provided by user to BgpV6SegmentRouting
	SetTimeToLive(value uint32) BgpV6SegmentRouting
	// HasTimeToLive checks if TimeToLive has been set in BgpV6SegmentRouting
	HasTimeToLive() bool
	// MaxSidsPerSrh returns uint32, set in BgpV6SegmentRouting.
	MaxSidsPerSrh() uint32
	// SetMaxSidsPerSrh assigns uint32 provided by user to BgpV6SegmentRouting
	SetMaxSidsPerSrh(value uint32) BgpV6SegmentRouting
	// HasMaxSidsPerSrh checks if MaxSidsPerSrh has been set in BgpV6SegmentRouting
	HasMaxSidsPerSrh() bool
	// AutoGenerateSegmentLeftValue returns bool, set in BgpV6SegmentRouting.
	AutoGenerateSegmentLeftValue() bool
	// SetAutoGenerateSegmentLeftValue assigns bool provided by user to BgpV6SegmentRouting
	SetAutoGenerateSegmentLeftValue(value bool) BgpV6SegmentRouting
	// HasAutoGenerateSegmentLeftValue checks if AutoGenerateSegmentLeftValue has been set in BgpV6SegmentRouting
	HasAutoGenerateSegmentLeftValue() bool
	// SegmentLeftValue returns uint32, set in BgpV6SegmentRouting.
	SegmentLeftValue() uint32
	// SetSegmentLeftValue assigns uint32 provided by user to BgpV6SegmentRouting
	SetSegmentLeftValue(value uint32) BgpV6SegmentRouting
	// HasSegmentLeftValue checks if SegmentLeftValue has been set in BgpV6SegmentRouting
	HasSegmentLeftValue() bool
	// AdvertiseSrTePolicy returns bool, set in BgpV6SegmentRouting.
	AdvertiseSrTePolicy() bool
	// SetAdvertiseSrTePolicy assigns bool provided by user to BgpV6SegmentRouting
	SetAdvertiseSrTePolicy(value bool) BgpV6SegmentRouting
	// HasAdvertiseSrTePolicy checks if AdvertiseSrTePolicy has been set in BgpV6SegmentRouting
	HasAdvertiseSrTePolicy() bool
}

// TBD
// IngressSupportsVpn returns a bool
func (obj *bgpV6SegmentRouting) IngressSupportsVpn() bool {

	return *obj.obj.IngressSupportsVpn

}

// TBD
// IngressSupportsVpn returns a bool
func (obj *bgpV6SegmentRouting) HasIngressSupportsVpn() bool {
	return obj.obj.IngressSupportsVpn != nil
}

// TBD
// SetIngressSupportsVpn sets the bool value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetIngressSupportsVpn(value bool) BgpV6SegmentRouting {

	obj.obj.IngressSupportsVpn = &value
	return obj
}

// TBD
// ReducedEncapsulation returns a bool
func (obj *bgpV6SegmentRouting) ReducedEncapsulation() bool {

	return *obj.obj.ReducedEncapsulation

}

// TBD
// ReducedEncapsulation returns a bool
func (obj *bgpV6SegmentRouting) HasReducedEncapsulation() bool {
	return obj.obj.ReducedEncapsulation != nil
}

// TBD
// SetReducedEncapsulation sets the bool value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetReducedEncapsulation(value bool) BgpV6SegmentRouting {

	obj.obj.ReducedEncapsulation = &value
	return obj
}

// TBD
// CopyTimeToLive returns a bool
func (obj *bgpV6SegmentRouting) CopyTimeToLive() bool {

	return *obj.obj.CopyTimeToLive

}

// TBD
// CopyTimeToLive returns a bool
func (obj *bgpV6SegmentRouting) HasCopyTimeToLive() bool {
	return obj.obj.CopyTimeToLive != nil
}

// TBD
// SetCopyTimeToLive sets the bool value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetCopyTimeToLive(value bool) BgpV6SegmentRouting {

	obj.obj.CopyTimeToLive = &value
	return obj
}

// TBD
// TimeToLive returns a uint32
func (obj *bgpV6SegmentRouting) TimeToLive() uint32 {

	return *obj.obj.TimeToLive

}

// TBD
// TimeToLive returns a uint32
func (obj *bgpV6SegmentRouting) HasTimeToLive() bool {
	return obj.obj.TimeToLive != nil
}

// TBD
// SetTimeToLive sets the uint32 value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetTimeToLive(value uint32) BgpV6SegmentRouting {

	obj.obj.TimeToLive = &value
	return obj
}

// TBD
// MaxSidsPerSrh returns a uint32
func (obj *bgpV6SegmentRouting) MaxSidsPerSrh() uint32 {

	return *obj.obj.MaxSidsPerSrh

}

// TBD
// MaxSidsPerSrh returns a uint32
func (obj *bgpV6SegmentRouting) HasMaxSidsPerSrh() bool {
	return obj.obj.MaxSidsPerSrh != nil
}

// TBD
// SetMaxSidsPerSrh sets the uint32 value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetMaxSidsPerSrh(value uint32) BgpV6SegmentRouting {

	obj.obj.MaxSidsPerSrh = &value
	return obj
}

// TBD
// AutoGenerateSegmentLeftValue returns a bool
func (obj *bgpV6SegmentRouting) AutoGenerateSegmentLeftValue() bool {

	return *obj.obj.AutoGenerateSegmentLeftValue

}

// TBD
// AutoGenerateSegmentLeftValue returns a bool
func (obj *bgpV6SegmentRouting) HasAutoGenerateSegmentLeftValue() bool {
	return obj.obj.AutoGenerateSegmentLeftValue != nil
}

// TBD
// SetAutoGenerateSegmentLeftValue sets the bool value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetAutoGenerateSegmentLeftValue(value bool) BgpV6SegmentRouting {

	obj.obj.AutoGenerateSegmentLeftValue = &value
	return obj
}

// TBD
// SegmentLeftValue returns a uint32
func (obj *bgpV6SegmentRouting) SegmentLeftValue() uint32 {

	return *obj.obj.SegmentLeftValue

}

// TBD
// SegmentLeftValue returns a uint32
func (obj *bgpV6SegmentRouting) HasSegmentLeftValue() bool {
	return obj.obj.SegmentLeftValue != nil
}

// TBD
// SetSegmentLeftValue sets the uint32 value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetSegmentLeftValue(value uint32) BgpV6SegmentRouting {

	obj.obj.SegmentLeftValue = &value
	return obj
}

// TBD
// AdvertiseSrTePolicy returns a bool
func (obj *bgpV6SegmentRouting) AdvertiseSrTePolicy() bool {

	return *obj.obj.AdvertiseSrTePolicy

}

// TBD
// AdvertiseSrTePolicy returns a bool
func (obj *bgpV6SegmentRouting) HasAdvertiseSrTePolicy() bool {
	return obj.obj.AdvertiseSrTePolicy != nil
}

// TBD
// SetAdvertiseSrTePolicy sets the bool value in the BgpV6SegmentRouting object
func (obj *bgpV6SegmentRouting) SetAdvertiseSrTePolicy(value bool) BgpV6SegmentRouting {

	obj.obj.AdvertiseSrTePolicy = &value
	return obj
}

func (obj *bgpV6SegmentRouting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MaxSidsPerSrh != nil {

		if *obj.obj.MaxSidsPerSrh > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpV6SegmentRouting.MaxSidsPerSrh <= 255 but Got %d", *obj.obj.MaxSidsPerSrh))
		}

	}

}

func (obj *bgpV6SegmentRouting) setDefault() {
	if obj.obj.IngressSupportsVpn == nil {
		obj.SetIngressSupportsVpn(false)
	}
	if obj.obj.ReducedEncapsulation == nil {
		obj.SetReducedEncapsulation(false)
	}
	if obj.obj.CopyTimeToLive == nil {
		obj.SetCopyTimeToLive(false)
	}
	if obj.obj.TimeToLive == nil {
		obj.SetTimeToLive(0)
	}
	if obj.obj.MaxSidsPerSrh == nil {
		obj.SetMaxSidsPerSrh(0)
	}
	if obj.obj.AutoGenerateSegmentLeftValue == nil {
		obj.SetAutoGenerateSegmentLeftValue(false)
	}
	if obj.obj.SegmentLeftValue == nil {
		obj.SetSegmentLeftValue(0)
	}
	if obj.obj.AdvertiseSrTePolicy == nil {
		obj.SetAdvertiseSrTePolicy(false)
	}

}
