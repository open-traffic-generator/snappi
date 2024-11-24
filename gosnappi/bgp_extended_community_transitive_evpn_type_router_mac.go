package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpExtendedCommunityTransitiveEvpnTypeRouterMac *****
type bgpExtendedCommunityTransitiveEvpnTypeRouterMac struct {
	validation
	obj          *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	marshaller   marshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac
	unMarshaller unMarshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac
}

func NewBgpExtendedCommunityTransitiveEvpnTypeRouterMac() BgpExtendedCommunityTransitiveEvpnTypeRouterMac {
	obj := bgpExtendedCommunityTransitiveEvpnTypeRouterMac{obj: &otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) msg() *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac {
	return obj.obj
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) setMsg(msg *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac) BgpExtendedCommunityTransitiveEvpnTypeRouterMac {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac struct {
	obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac
}

type marshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac interface {
	// ToProto marshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac to protobuf object *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	ToProto() (*otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac, error)
	// ToPbText marshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac struct {
	obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac
}

type unMarshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac interface {
	// FromProto unmarshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac from protobuf object *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	FromProto(msg *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac) (BgpExtendedCommunityTransitiveEvpnTypeRouterMac, error)
	// FromPbText unmarshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac from JSON text
	FromJson(value string) error
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) Marshal() marshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) Unmarshal() unMarshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) ToProto() (*otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) FromProto(msg *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac) (BgpExtendedCommunityTransitiveEvpnTypeRouterMac, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) ToPbText() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) FromPbText(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) ToYaml() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) FromYaml(value string) error {
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

func (m *marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) ToJsonRaw() (string, error) {
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

func (m *marshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) ToJson() (string, error) {
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

func (m *unMarshalbgpExtendedCommunityTransitiveEvpnTypeRouterMac) FromJson(value string) error {
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

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) Clone() (BgpExtendedCommunityTransitiveEvpnTypeRouterMac, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpExtendedCommunityTransitiveEvpnTypeRouterMac()
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

// BgpExtendedCommunityTransitiveEvpnTypeRouterMac is the Router MAC EVPN Community is defined in RFC9135 and normally sent only for EVPN Type-2 Routes . It is sent with sub-type 0x03.
type BgpExtendedCommunityTransitiveEvpnTypeRouterMac interface {
	Validation
	// msg marshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac to protobuf object *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// and doesn't set defaults
	msg() *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// setMsg unmarshals BgpExtendedCommunityTransitiveEvpnTypeRouterMac from protobuf object *otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// and doesn't set defaults
	setMsg(*otg.BgpExtendedCommunityTransitiveEvpnTypeRouterMac) BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// provides marshal interface
	Marshal() marshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// provides unmarshal interface
	Unmarshal() unMarshalBgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// validate validates BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpExtendedCommunityTransitiveEvpnTypeRouterMac, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterMac returns string, set in BgpExtendedCommunityTransitiveEvpnTypeRouterMac.
	RouterMac() string
	// SetRouterMac assigns string provided by user to BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	SetRouterMac(value string) BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	// HasRouterMac checks if RouterMac has been set in BgpExtendedCommunityTransitiveEvpnTypeRouterMac
	HasRouterMac() bool
}

// MAC Address of the PE Router.
// RouterMac returns a string
func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) RouterMac() string {

	return *obj.obj.RouterMac

}

// MAC Address of the PE Router.
// RouterMac returns a string
func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) HasRouterMac() bool {
	return obj.obj.RouterMac != nil
}

// MAC Address of the PE Router.
// SetRouterMac sets the string value in the BgpExtendedCommunityTransitiveEvpnTypeRouterMac object
func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) SetRouterMac(value string) BgpExtendedCommunityTransitiveEvpnTypeRouterMac {

	obj.obj.RouterMac = &value
	return obj
}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterMac != nil {

		err := obj.validateMac(obj.RouterMac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpExtendedCommunityTransitiveEvpnTypeRouterMac.RouterMac"))
		}

	}

}

func (obj *bgpExtendedCommunityTransitiveEvpnTypeRouterMac) setDefault() {
	if obj.obj.RouterMac == nil {
		obj.SetRouterMac("0:0:0:0:0:0")
	}

}
