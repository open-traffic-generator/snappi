package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServerPeerSessionIpv4Information *****
type bmpServerPeerSessionIpv4Information struct {
	validation
	obj          *otg.BmpServerPeerSessionIpv4Information
	marshaller   marshalBmpServerPeerSessionIpv4Information
	unMarshaller unMarshalBmpServerPeerSessionIpv4Information
}

func NewBmpServerPeerSessionIpv4Information() BmpServerPeerSessionIpv4Information {
	obj := bmpServerPeerSessionIpv4Information{obj: &otg.BmpServerPeerSessionIpv4Information{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServerPeerSessionIpv4Information) msg() *otg.BmpServerPeerSessionIpv4Information {
	return obj.obj
}

func (obj *bmpServerPeerSessionIpv4Information) setMsg(msg *otg.BmpServerPeerSessionIpv4Information) BmpServerPeerSessionIpv4Information {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServerPeerSessionIpv4Information struct {
	obj *bmpServerPeerSessionIpv4Information
}

type marshalBmpServerPeerSessionIpv4Information interface {
	// ToProto marshals BmpServerPeerSessionIpv4Information to protobuf object *otg.BmpServerPeerSessionIpv4Information
	ToProto() (*otg.BmpServerPeerSessionIpv4Information, error)
	// ToPbText marshals BmpServerPeerSessionIpv4Information to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServerPeerSessionIpv4Information to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServerPeerSessionIpv4Information to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServerPeerSessionIpv4Information struct {
	obj *bmpServerPeerSessionIpv4Information
}

type unMarshalBmpServerPeerSessionIpv4Information interface {
	// FromProto unmarshals BmpServerPeerSessionIpv4Information from protobuf object *otg.BmpServerPeerSessionIpv4Information
	FromProto(msg *otg.BmpServerPeerSessionIpv4Information) (BmpServerPeerSessionIpv4Information, error)
	// FromPbText unmarshals BmpServerPeerSessionIpv4Information from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServerPeerSessionIpv4Information from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServerPeerSessionIpv4Information from JSON text
	FromJson(value string) error
}

func (obj *bmpServerPeerSessionIpv4Information) Marshal() marshalBmpServerPeerSessionIpv4Information {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServerPeerSessionIpv4Information{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServerPeerSessionIpv4Information) Unmarshal() unMarshalBmpServerPeerSessionIpv4Information {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServerPeerSessionIpv4Information{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServerPeerSessionIpv4Information) ToProto() (*otg.BmpServerPeerSessionIpv4Information, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServerPeerSessionIpv4Information) FromProto(msg *otg.BmpServerPeerSessionIpv4Information) (BmpServerPeerSessionIpv4Information, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServerPeerSessionIpv4Information) ToPbText() (string, error) {
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

func (m *unMarshalbmpServerPeerSessionIpv4Information) FromPbText(value string) error {
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

func (m *marshalbmpServerPeerSessionIpv4Information) ToYaml() (string, error) {
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

func (m *unMarshalbmpServerPeerSessionIpv4Information) FromYaml(value string) error {
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

func (m *marshalbmpServerPeerSessionIpv4Information) ToJson() (string, error) {
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

func (m *unMarshalbmpServerPeerSessionIpv4Information) FromJson(value string) error {
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

func (obj *bmpServerPeerSessionIpv4Information) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServerPeerSessionIpv4Information) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServerPeerSessionIpv4Information) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServerPeerSessionIpv4Information) Clone() (BmpServerPeerSessionIpv4Information, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServerPeerSessionIpv4Information()
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

// BmpServerPeerSessionIpv4Information is this object is included if session_ip_type is of type ipv4.
type BmpServerPeerSessionIpv4Information interface {
	Validation
	// msg marshals BmpServerPeerSessionIpv4Information to protobuf object *otg.BmpServerPeerSessionIpv4Information
	// and doesn't set defaults
	msg() *otg.BmpServerPeerSessionIpv4Information
	// setMsg unmarshals BmpServerPeerSessionIpv4Information from protobuf object *otg.BmpServerPeerSessionIpv4Information
	// and doesn't set defaults
	setMsg(*otg.BmpServerPeerSessionIpv4Information) BmpServerPeerSessionIpv4Information
	// provides marshal interface
	Marshal() marshalBmpServerPeerSessionIpv4Information
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServerPeerSessionIpv4Information
	// validate validates BmpServerPeerSessionIpv4Information
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServerPeerSessionIpv4Information, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocalAddress returns string, set in BmpServerPeerSessionIpv4Information.
	LocalAddress() string
	// SetLocalAddress assigns string provided by user to BmpServerPeerSessionIpv4Information
	SetLocalAddress(value string) BmpServerPeerSessionIpv4Information
	// HasLocalAddress checks if LocalAddress has been set in BmpServerPeerSessionIpv4Information
	HasLocalAddress() bool
	// RemoteAddress returns string, set in BmpServerPeerSessionIpv4Information.
	RemoteAddress() string
	// SetRemoteAddress assigns string provided by user to BmpServerPeerSessionIpv4Information
	SetRemoteAddress(value string) BmpServerPeerSessionIpv4Information
	// HasRemoteAddress checks if RemoteAddress has been set in BmpServerPeerSessionIpv4Information
	HasRemoteAddress() bool
}

// The local IPv4 address associated with the BGPv4 peer.
// LocalAddress returns a string
func (obj *bmpServerPeerSessionIpv4Information) LocalAddress() string {

	return *obj.obj.LocalAddress

}

// The local IPv4 address associated with the BGPv4 peer.
// LocalAddress returns a string
func (obj *bmpServerPeerSessionIpv4Information) HasLocalAddress() bool {
	return obj.obj.LocalAddress != nil
}

// The local IPv4 address associated with the BGPv4 peer.
// SetLocalAddress sets the string value in the BmpServerPeerSessionIpv4Information object
func (obj *bmpServerPeerSessionIpv4Information) SetLocalAddress(value string) BmpServerPeerSessionIpv4Information {

	obj.obj.LocalAddress = &value
	return obj
}

// The remote IPv4 address associated with the BGPv4 peer.
// RemoteAddress returns a string
func (obj *bmpServerPeerSessionIpv4Information) RemoteAddress() string {

	return *obj.obj.RemoteAddress

}

// The remote IPv4 address associated with the BGPv4 peer.
// RemoteAddress returns a string
func (obj *bmpServerPeerSessionIpv4Information) HasRemoteAddress() bool {
	return obj.obj.RemoteAddress != nil
}

// The remote IPv4 address associated with the BGPv4 peer.
// SetRemoteAddress sets the string value in the BmpServerPeerSessionIpv4Information object
func (obj *bmpServerPeerSessionIpv4Information) SetRemoteAddress(value string) BmpServerPeerSessionIpv4Information {

	obj.obj.RemoteAddress = &value
	return obj
}

func (obj *bmpServerPeerSessionIpv4Information) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocalAddress != nil {

		err := obj.validateIpv4(obj.LocalAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpServerPeerSessionIpv4Information.LocalAddress"))
		}

	}

	if obj.obj.RemoteAddress != nil {

		err := obj.validateIpv4(obj.RemoteAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BmpServerPeerSessionIpv4Information.RemoteAddress"))
		}

	}

}

func (obj *bmpServerPeerSessionIpv4Information) setDefault() {

}
