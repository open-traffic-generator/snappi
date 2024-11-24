package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSidSrv6 *****
type bgpAttributesSidSrv6 struct {
	validation
	obj          *otg.BgpAttributesSidSrv6
	marshaller   marshalBgpAttributesSidSrv6
	unMarshaller unMarshalBgpAttributesSidSrv6
}

func NewBgpAttributesSidSrv6() BgpAttributesSidSrv6 {
	obj := bgpAttributesSidSrv6{obj: &otg.BgpAttributesSidSrv6{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSidSrv6) msg() *otg.BgpAttributesSidSrv6 {
	return obj.obj
}

func (obj *bgpAttributesSidSrv6) setMsg(msg *otg.BgpAttributesSidSrv6) BgpAttributesSidSrv6 {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSidSrv6 struct {
	obj *bgpAttributesSidSrv6
}

type marshalBgpAttributesSidSrv6 interface {
	// ToProto marshals BgpAttributesSidSrv6 to protobuf object *otg.BgpAttributesSidSrv6
	ToProto() (*otg.BgpAttributesSidSrv6, error)
	// ToPbText marshals BgpAttributesSidSrv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSidSrv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSidSrv6 to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesSidSrv6 to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesSidSrv6 struct {
	obj *bgpAttributesSidSrv6
}

type unMarshalBgpAttributesSidSrv6 interface {
	// FromProto unmarshals BgpAttributesSidSrv6 from protobuf object *otg.BgpAttributesSidSrv6
	FromProto(msg *otg.BgpAttributesSidSrv6) (BgpAttributesSidSrv6, error)
	// FromPbText unmarshals BgpAttributesSidSrv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSidSrv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSidSrv6 from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSidSrv6) Marshal() marshalBgpAttributesSidSrv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSidSrv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSidSrv6) Unmarshal() unMarshalBgpAttributesSidSrv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSidSrv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSidSrv6) ToProto() (*otg.BgpAttributesSidSrv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSidSrv6) FromProto(msg *otg.BgpAttributesSidSrv6) (BgpAttributesSidSrv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSidSrv6) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSidSrv6) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSidSrv6) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSidSrv6) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSidSrv6) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesSidSrv6) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSidSrv6) FromJson(value string) error {
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

func (obj *bgpAttributesSidSrv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSidSrv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSidSrv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSidSrv6) Clone() (BgpAttributesSidSrv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSidSrv6()
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

// BgpAttributesSidSrv6 is an IPv6 address denoting a SRv6 SID.
type BgpAttributesSidSrv6 interface {
	Validation
	// msg marshals BgpAttributesSidSrv6 to protobuf object *otg.BgpAttributesSidSrv6
	// and doesn't set defaults
	msg() *otg.BgpAttributesSidSrv6
	// setMsg unmarshals BgpAttributesSidSrv6 from protobuf object *otg.BgpAttributesSidSrv6
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSidSrv6) BgpAttributesSidSrv6
	// provides marshal interface
	Marshal() marshalBgpAttributesSidSrv6
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSidSrv6
	// validate validates BgpAttributesSidSrv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSidSrv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ip returns string, set in BgpAttributesSidSrv6.
	Ip() string
	// SetIp assigns string provided by user to BgpAttributesSidSrv6
	SetIp(value string) BgpAttributesSidSrv6
	// HasIp checks if Ip has been set in BgpAttributesSidSrv6
	HasIp() bool
}

// description is TBD
// Ip returns a string
func (obj *bgpAttributesSidSrv6) Ip() string {

	return *obj.obj.Ip

}

// description is TBD
// Ip returns a string
func (obj *bgpAttributesSidSrv6) HasIp() bool {
	return obj.obj.Ip != nil
}

// description is TBD
// SetIp sets the string value in the BgpAttributesSidSrv6 object
func (obj *bgpAttributesSidSrv6) SetIp(value string) BgpAttributesSidSrv6 {

	obj.obj.Ip = &value
	return obj
}

func (obj *bgpAttributesSidSrv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ip != nil {

		err := obj.validateIpv6(obj.Ip())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpAttributesSidSrv6.Ip"))
		}

	}

}

func (obj *bgpAttributesSidSrv6) setDefault() {
	if obj.obj.Ip == nil {
		obj.SetIp("0::0")
	}

}
