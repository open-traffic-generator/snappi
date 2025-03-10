package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpRouteAdvanced *****
type bgpRouteAdvanced struct {
	validation
	obj          *otg.BgpRouteAdvanced
	marshaller   marshalBgpRouteAdvanced
	unMarshaller unMarshalBgpRouteAdvanced
}

func NewBgpRouteAdvanced() BgpRouteAdvanced {
	obj := bgpRouteAdvanced{obj: &otg.BgpRouteAdvanced{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpRouteAdvanced) msg() *otg.BgpRouteAdvanced {
	return obj.obj
}

func (obj *bgpRouteAdvanced) setMsg(msg *otg.BgpRouteAdvanced) BgpRouteAdvanced {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpRouteAdvanced struct {
	obj *bgpRouteAdvanced
}

type marshalBgpRouteAdvanced interface {
	// ToProto marshals BgpRouteAdvanced to protobuf object *otg.BgpRouteAdvanced
	ToProto() (*otg.BgpRouteAdvanced, error)
	// ToPbText marshals BgpRouteAdvanced to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpRouteAdvanced to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpRouteAdvanced to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpRouteAdvanced to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpRouteAdvanced struct {
	obj *bgpRouteAdvanced
}

type unMarshalBgpRouteAdvanced interface {
	// FromProto unmarshals BgpRouteAdvanced from protobuf object *otg.BgpRouteAdvanced
	FromProto(msg *otg.BgpRouteAdvanced) (BgpRouteAdvanced, error)
	// FromPbText unmarshals BgpRouteAdvanced from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpRouteAdvanced from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpRouteAdvanced from JSON text
	FromJson(value string) error
}

func (obj *bgpRouteAdvanced) Marshal() marshalBgpRouteAdvanced {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpRouteAdvanced{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpRouteAdvanced) Unmarshal() unMarshalBgpRouteAdvanced {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpRouteAdvanced{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpRouteAdvanced) ToProto() (*otg.BgpRouteAdvanced, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpRouteAdvanced) FromProto(msg *otg.BgpRouteAdvanced) (BgpRouteAdvanced, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpRouteAdvanced) ToPbText() (string, error) {
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

func (m *unMarshalbgpRouteAdvanced) FromPbText(value string) error {
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

func (m *marshalbgpRouteAdvanced) ToYaml() (string, error) {
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

func (m *unMarshalbgpRouteAdvanced) FromYaml(value string) error {
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

func (m *marshalbgpRouteAdvanced) ToJsonRaw() (string, error) {
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

func (m *marshalbgpRouteAdvanced) ToJson() (string, error) {
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

func (m *unMarshalbgpRouteAdvanced) FromJson(value string) error {
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

func (obj *bgpRouteAdvanced) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpRouteAdvanced) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpRouteAdvanced) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpRouteAdvanced) Clone() (BgpRouteAdvanced, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpRouteAdvanced()
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

// BgpRouteAdvanced is configuration for advanced BGP route range settings.
type BgpRouteAdvanced interface {
	Validation
	// msg marshals BgpRouteAdvanced to protobuf object *otg.BgpRouteAdvanced
	// and doesn't set defaults
	msg() *otg.BgpRouteAdvanced
	// setMsg unmarshals BgpRouteAdvanced from protobuf object *otg.BgpRouteAdvanced
	// and doesn't set defaults
	setMsg(*otg.BgpRouteAdvanced) BgpRouteAdvanced
	// provides marshal interface
	Marshal() marshalBgpRouteAdvanced
	// provides unmarshal interface
	Unmarshal() unMarshalBgpRouteAdvanced
	// validate validates BgpRouteAdvanced
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpRouteAdvanced, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IncludeMultiExitDiscriminator returns bool, set in BgpRouteAdvanced.
	IncludeMultiExitDiscriminator() bool
	// SetIncludeMultiExitDiscriminator assigns bool provided by user to BgpRouteAdvanced
	SetIncludeMultiExitDiscriminator(value bool) BgpRouteAdvanced
	// HasIncludeMultiExitDiscriminator checks if IncludeMultiExitDiscriminator has been set in BgpRouteAdvanced
	HasIncludeMultiExitDiscriminator() bool
	// MultiExitDiscriminator returns uint32, set in BgpRouteAdvanced.
	MultiExitDiscriminator() uint32
	// SetMultiExitDiscriminator assigns uint32 provided by user to BgpRouteAdvanced
	SetMultiExitDiscriminator(value uint32) BgpRouteAdvanced
	// HasMultiExitDiscriminator checks if MultiExitDiscriminator has been set in BgpRouteAdvanced
	HasMultiExitDiscriminator() bool
	// IncludeOrigin returns bool, set in BgpRouteAdvanced.
	IncludeOrigin() bool
	// SetIncludeOrigin assigns bool provided by user to BgpRouteAdvanced
	SetIncludeOrigin(value bool) BgpRouteAdvanced
	// HasIncludeOrigin checks if IncludeOrigin has been set in BgpRouteAdvanced
	HasIncludeOrigin() bool
	// Origin returns BgpRouteAdvancedOriginEnum, set in BgpRouteAdvanced
	Origin() BgpRouteAdvancedOriginEnum
	// SetOrigin assigns BgpRouteAdvancedOriginEnum provided by user to BgpRouteAdvanced
	SetOrigin(value BgpRouteAdvancedOriginEnum) BgpRouteAdvanced
	// HasOrigin checks if Origin has been set in BgpRouteAdvanced
	HasOrigin() bool
	// IncludeLocalPreference returns bool, set in BgpRouteAdvanced.
	IncludeLocalPreference() bool
	// SetIncludeLocalPreference assigns bool provided by user to BgpRouteAdvanced
	SetIncludeLocalPreference(value bool) BgpRouteAdvanced
	// HasIncludeLocalPreference checks if IncludeLocalPreference has been set in BgpRouteAdvanced
	HasIncludeLocalPreference() bool
	// LocalPreference returns uint32, set in BgpRouteAdvanced.
	LocalPreference() uint32
	// SetLocalPreference assigns uint32 provided by user to BgpRouteAdvanced
	SetLocalPreference(value uint32) BgpRouteAdvanced
	// HasLocalPreference checks if LocalPreference has been set in BgpRouteAdvanced
	HasLocalPreference() bool
}

// BGP Multi Exit Discriminator attribute sent to the peer to help in the route selection process.  If set to true, the Multi Exit Discriminator attribute will be included in the route advertisement.
// IncludeMultiExitDiscriminator returns a bool
func (obj *bgpRouteAdvanced) IncludeMultiExitDiscriminator() bool {

	return *obj.obj.IncludeMultiExitDiscriminator

}

// BGP Multi Exit Discriminator attribute sent to the peer to help in the route selection process.  If set to true, the Multi Exit Discriminator attribute will be included in the route advertisement.
// IncludeMultiExitDiscriminator returns a bool
func (obj *bgpRouteAdvanced) HasIncludeMultiExitDiscriminator() bool {
	return obj.obj.IncludeMultiExitDiscriminator != nil
}

// BGP Multi Exit Discriminator attribute sent to the peer to help in the route selection process.  If set to true, the Multi Exit Discriminator attribute will be included in the route advertisement.
// SetIncludeMultiExitDiscriminator sets the bool value in the BgpRouteAdvanced object
func (obj *bgpRouteAdvanced) SetIncludeMultiExitDiscriminator(value bool) BgpRouteAdvanced {

	obj.obj.IncludeMultiExitDiscriminator = &value
	return obj
}

// The multi exit discriminator (MED) value used for route selection sent to the peer.
// MultiExitDiscriminator returns a uint32
func (obj *bgpRouteAdvanced) MultiExitDiscriminator() uint32 {

	return *obj.obj.MultiExitDiscriminator

}

// The multi exit discriminator (MED) value used for route selection sent to the peer.
// MultiExitDiscriminator returns a uint32
func (obj *bgpRouteAdvanced) HasMultiExitDiscriminator() bool {
	return obj.obj.MultiExitDiscriminator != nil
}

// The multi exit discriminator (MED) value used for route selection sent to the peer.
// SetMultiExitDiscriminator sets the uint32 value in the BgpRouteAdvanced object
func (obj *bgpRouteAdvanced) SetMultiExitDiscriminator(value uint32) BgpRouteAdvanced {

	obj.obj.MultiExitDiscriminator = &value
	return obj
}

// If set to true, the Origin attribute will be included in the route advertisement.
// IncludeOrigin returns a bool
func (obj *bgpRouteAdvanced) IncludeOrigin() bool {

	return *obj.obj.IncludeOrigin

}

// If set to true, the Origin attribute will be included in the route advertisement.
// IncludeOrigin returns a bool
func (obj *bgpRouteAdvanced) HasIncludeOrigin() bool {
	return obj.obj.IncludeOrigin != nil
}

// If set to true, the Origin attribute will be included in the route advertisement.
// SetIncludeOrigin sets the bool value in the BgpRouteAdvanced object
func (obj *bgpRouteAdvanced) SetIncludeOrigin(value bool) BgpRouteAdvanced {

	obj.obj.IncludeOrigin = &value
	return obj
}

type BgpRouteAdvancedOriginEnum string

// Enum of Origin on BgpRouteAdvanced
var BgpRouteAdvancedOrigin = struct {
	IGP        BgpRouteAdvancedOriginEnum
	EGP        BgpRouteAdvancedOriginEnum
	INCOMPLETE BgpRouteAdvancedOriginEnum
}{
	IGP:        BgpRouteAdvancedOriginEnum("igp"),
	EGP:        BgpRouteAdvancedOriginEnum("egp"),
	INCOMPLETE: BgpRouteAdvancedOriginEnum("incomplete"),
}

func (obj *bgpRouteAdvanced) Origin() BgpRouteAdvancedOriginEnum {
	return BgpRouteAdvancedOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin attribute of a prefix can take three values: the prefix originates from an interior routing protocol 'igp',  it originates from 'egp' or the origin is 'incomplete', if the prefix is learned through other means.
// Origin returns a string
func (obj *bgpRouteAdvanced) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpRouteAdvanced) SetOrigin(value BgpRouteAdvancedOriginEnum) BgpRouteAdvanced {
	intValue, ok := otg.BgpRouteAdvanced_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpRouteAdvancedOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpRouteAdvanced_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// BGP Local Preference attribute sent to the peer to indicate the degree of preference  for externally learned routes. If set to true, the Local Preference attribute will be included  in the route advertisement. This should be included only for internal peers.
// IncludeLocalPreference returns a bool
func (obj *bgpRouteAdvanced) IncludeLocalPreference() bool {

	return *obj.obj.IncludeLocalPreference

}

// BGP Local Preference attribute sent to the peer to indicate the degree of preference  for externally learned routes. If set to true, the Local Preference attribute will be included  in the route advertisement. This should be included only for internal peers.
// IncludeLocalPreference returns a bool
func (obj *bgpRouteAdvanced) HasIncludeLocalPreference() bool {
	return obj.obj.IncludeLocalPreference != nil
}

// BGP Local Preference attribute sent to the peer to indicate the degree of preference  for externally learned routes. If set to true, the Local Preference attribute will be included  in the route advertisement. This should be included only for internal peers.
// SetIncludeLocalPreference sets the bool value in the BgpRouteAdvanced object
func (obj *bgpRouteAdvanced) SetIncludeLocalPreference(value bool) BgpRouteAdvanced {

	obj.obj.IncludeLocalPreference = &value
	return obj
}

// Value to be set in Local Preference attribute if include_local_preference is set to true. It is  used for the selection of the path for the traffic leaving the AS. The route with the  highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpRouteAdvanced) LocalPreference() uint32 {

	return *obj.obj.LocalPreference

}

// Value to be set in Local Preference attribute if include_local_preference is set to true. It is  used for the selection of the path for the traffic leaving the AS. The route with the  highest local preference value is preferred.
// LocalPreference returns a uint32
func (obj *bgpRouteAdvanced) HasLocalPreference() bool {
	return obj.obj.LocalPreference != nil
}

// Value to be set in Local Preference attribute if include_local_preference is set to true. It is  used for the selection of the path for the traffic leaving the AS. The route with the  highest local preference value is preferred.
// SetLocalPreference sets the uint32 value in the BgpRouteAdvanced object
func (obj *bgpRouteAdvanced) SetLocalPreference(value uint32) BgpRouteAdvanced {

	obj.obj.LocalPreference = &value
	return obj
}

func (obj *bgpRouteAdvanced) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpRouteAdvanced) setDefault() {
	if obj.obj.IncludeMultiExitDiscriminator == nil {
		obj.SetIncludeMultiExitDiscriminator(true)
	}
	if obj.obj.IncludeOrigin == nil {
		obj.SetIncludeOrigin(true)
	}
	if obj.obj.Origin == nil {
		obj.SetOrigin(BgpRouteAdvancedOrigin.IGP)

	}
	if obj.obj.IncludeLocalPreference == nil {
		obj.SetIncludeLocalPreference(true)
	}
	if obj.obj.LocalPreference == nil {
		obj.SetLocalPreference(100)
	}

}
