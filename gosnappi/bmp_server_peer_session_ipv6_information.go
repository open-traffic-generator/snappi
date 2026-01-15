package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerPeerSessionIpv6Information *****
type bmpServerPeerSessionIpv6Information struct {
	validation
	obj          *otg.BmpServerPeerSessionIpv6Information
	marshaller   marshalBmpServerPeerSessionIpv6Information
	unMarshaller unMarshalBmpServerPeerSessionIpv6Information
}

func NewBmpServerPeerSessionIpv6Information() BmpServerPeerSessionIpv6Information {
	obj := bmpServerPeerSessionIpv6Information{obj: &otg.BmpServerPeerSessionIpv6Information{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerPeerSessionIpv6Information) msg() *otg.BmpServerPeerSessionIpv6Information {
	return obj.obj
}

func (obj *bmpServerPeerSessionIpv6Information) setMsg(msg *otg.BmpServerPeerSessionIpv6Information) BmpServerPeerSessionIpv6Information {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerPeerSessionIpv6Information struct {
	obj *bmpServerPeerSessionIpv6Information
}

type marshalBmpServerPeerSessionIpv6Information interface {
	// ToProto marshals BmpServerPeerSessionIpv6Information to protobuf object *otg.BmpServerPeerSessionIpv6Information
	ToProto() (*otg.BmpServerPeerSessionIpv6Information, error)
	// ToPbText marshals BmpServerPeerSessionIpv6Information to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerPeerSessionIpv6Information to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerPeerSessionIpv6Information to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerPeerSessionIpv6Information struct {
	obj *bmpServerPeerSessionIpv6Information
}

type unMarshalBmpServerPeerSessionIpv6Information interface {
	// FromProto unmarshals BmpServerPeerSessionIpv6Information from protobuf object *otg.BmpServerPeerSessionIpv6Information
	FromProto(msg *otg.BmpServerPeerSessionIpv6Information) (BmpServerPeerSessionIpv6Information, error)
	// FromPbText unmarshals BmpServerPeerSessionIpv6Information from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerPeerSessionIpv6Information from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerPeerSessionIpv6Information from JSON text
	FromJson(value string) error
}

func (obj *bmpServerPeerSessionIpv6Information) Marshal() marshalBmpServerPeerSessionIpv6Information {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerPeerSessionIpv6Information{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerPeerSessionIpv6Information) Unmarshal() unMarshalBmpServerPeerSessionIpv6Information {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerPeerSessionIpv6Information{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerPeerSessionIpv6Information) ToProto() (*otg.BmpServerPeerSessionIpv6Information, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerPeerSessionIpv6Information) FromProto(msg *otg.BmpServerPeerSessionIpv6Information) (BmpServerPeerSessionIpv6Information, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerPeerSessionIpv6Information) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerPeerSessionIpv6Information) FromPbText(value string) error {
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

func (m *marshalbmpServerPeerSessionIpv6Information) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerPeerSessionIpv6Information) FromYaml(value string) error {
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

func (m *marshalbmpServerPeerSessionIpv6Information) ToJson() (string, error) {
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

func (m *unMarshalbmpServerPeerSessionIpv6Information) FromJson(value string) error {
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

func (obj *bmpServerPeerSessionIpv6Information) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerPeerSessionIpv6Information) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerPeerSessionIpv6Information) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerPeerSessionIpv6Information) Clone() (BmpServerPeerSessionIpv6Information, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerPeerSessionIpv6Information()
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

// BmpServerPeerSessionIpv6Information is this object is included if session_ip_type is of type ipv6.
type BmpServerPeerSessionIpv6Information interface {
	Validation
	// msg marshals BmpServerPeerSessionIpv6Information to protobuf object *otg.BmpServerPeerSessionIpv6Information
	// and doesn't set defaults
	msg() *otg.BmpServerPeerSessionIpv6Information
	// setMsg unmarshals BmpServerPeerSessionIpv6Information from protobuf object *otg.BmpServerPeerSessionIpv6Information
	// and doesn't set defaults
	setMsg(*otg.BmpServerPeerSessionIpv6Information) BmpServerPeerSessionIpv6Information
	// provides marshal interface
	Marshal() marshalBmpServerPeerSessionIpv6Information
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerPeerSessionIpv6Information
	// validate validates BmpServerPeerSessionIpv6Information
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerPeerSessionIpv6Information, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocalAddress returns string, set in BmpServerPeerSessionIpv6Information.
	LocalAddress() string
	// SetLocalAddress assigns string provided by user to BmpServerPeerSessionIpv6Information
	SetLocalAddress(value string) BmpServerPeerSessionIpv6Information
	// HasLocalAddress checks if LocalAddress has been set in BmpServerPeerSessionIpv6Information
	HasLocalAddress() bool
	// RemoteAddress returns string, set in BmpServerPeerSessionIpv6Information.
	RemoteAddress() string
	// SetRemoteAddress assigns string provided by user to BmpServerPeerSessionIpv6Information
	SetRemoteAddress(value string) BmpServerPeerSessionIpv6Information
	// HasRemoteAddress checks if RemoteAddress has been set in BmpServerPeerSessionIpv6Information
	HasRemoteAddress() bool
}

// The local IPv6 address associated with the BGPv6 peer.
// LocalAddress returns a string
func (obj *bmpServerPeerSessionIpv6Information) LocalAddress() string {

	return *obj.obj.LocalAddress

}

// The local IPv6 address associated with the BGPv6 peer.
// LocalAddress returns a string
func (obj *bmpServerPeerSessionIpv6Information) HasLocalAddress() bool {
	return obj.obj.LocalAddress != nil
}

// The local IPv6 address associated with the BGPv6 peer.
// SetLocalAddress sets the string value in the BmpServerPeerSessionIpv6Information object
func (obj *bmpServerPeerSessionIpv6Information) SetLocalAddress(value string) BmpServerPeerSessionIpv6Information {

	obj.obj.LocalAddress = &value
	return obj
}

// The remote IPv6 address associated with the BGPv6 peer.
// RemoteAddress returns a string
func (obj *bmpServerPeerSessionIpv6Information) RemoteAddress() string {

	return *obj.obj.RemoteAddress

}

// The remote IPv6 address associated with the BGPv6 peer.
// RemoteAddress returns a string
func (obj *bmpServerPeerSessionIpv6Information) HasRemoteAddress() bool {
	return obj.obj.RemoteAddress != nil
}

// The remote IPv6 address associated with the BGPv6 peer.
// SetRemoteAddress sets the string value in the BmpServerPeerSessionIpv6Information object
func (obj *bmpServerPeerSessionIpv6Information) SetRemoteAddress(value string) BmpServerPeerSessionIpv6Information {

	obj.obj.RemoteAddress = &value
	return obj
}

func (obj *bmpServerPeerSessionIpv6Information) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocalAddress != nil {

		err := obj.validateIpv6(obj.LocalAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpServerPeerSessionIpv6Information.LocalAddress"))
		}

	}

	if obj.obj.RemoteAddress != nil {

		err := obj.validateIpv6(obj.RemoteAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpServerPeerSessionIpv6Information.RemoteAddress"))
		}

	}

}

func (obj *bmpServerPeerSessionIpv6Information) setDefault() {

}
