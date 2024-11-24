package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteColorSubTlv *****
type bgpSrteColorSubTlv struct {
	validation
	obj          *otg.BgpSrteColorSubTlv
	marshaller   marshalBgpSrteColorSubTlv
	unMarshaller unMarshalBgpSrteColorSubTlv
}

func NewBgpSrteColorSubTlv() BgpSrteColorSubTlv {
	obj := bgpSrteColorSubTlv{obj: &otg.BgpSrteColorSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteColorSubTlv) msg() *otg.BgpSrteColorSubTlv {
	return obj.obj
}

func (obj *bgpSrteColorSubTlv) setMsg(msg *otg.BgpSrteColorSubTlv) BgpSrteColorSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteColorSubTlv struct {
	obj *bgpSrteColorSubTlv
}

type marshalBgpSrteColorSubTlv interface {
	// ToProto marshals BgpSrteColorSubTlv to protobuf object *otg.BgpSrteColorSubTlv
	ToProto() (*otg.BgpSrteColorSubTlv, error)
	// ToPbText marshals BgpSrteColorSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteColorSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteColorSubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrteColorSubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrteColorSubTlv struct {
	obj *bgpSrteColorSubTlv
}

type unMarshalBgpSrteColorSubTlv interface {
	// FromProto unmarshals BgpSrteColorSubTlv from protobuf object *otg.BgpSrteColorSubTlv
	FromProto(msg *otg.BgpSrteColorSubTlv) (BgpSrteColorSubTlv, error)
	// FromPbText unmarshals BgpSrteColorSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteColorSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteColorSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteColorSubTlv) Marshal() marshalBgpSrteColorSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteColorSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteColorSubTlv) Unmarshal() unMarshalBgpSrteColorSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteColorSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteColorSubTlv) ToProto() (*otg.BgpSrteColorSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteColorSubTlv) FromProto(msg *otg.BgpSrteColorSubTlv) (BgpSrteColorSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteColorSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteColorSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteColorSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteColorSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteColorSubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrteColorSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteColorSubTlv) FromJson(value string) error {
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

func (obj *bgpSrteColorSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteColorSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteColorSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteColorSubTlv) Clone() (BgpSrteColorSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteColorSubTlv()
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

// BgpSrteColorSubTlv is configuration for the Policy Color attribute sub-TLV. The Color sub-TLV MAY be used as a way to "color" the corresponding Tunnel TLV. The Value field of the sub-TLV is eight octets long and consists of a Color Extended Community. First two octets of its Value field are 0x030b as type and subtype of extended  community. Remaining six octets are are exposed to configure.
type BgpSrteColorSubTlv interface {
	Validation
	// msg marshals BgpSrteColorSubTlv to protobuf object *otg.BgpSrteColorSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteColorSubTlv
	// setMsg unmarshals BgpSrteColorSubTlv from protobuf object *otg.BgpSrteColorSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteColorSubTlv) BgpSrteColorSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteColorSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteColorSubTlv
	// validate validates BgpSrteColorSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteColorSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Color returns string, set in BgpSrteColorSubTlv.
	Color() string
	// SetColor assigns string provided by user to BgpSrteColorSubTlv
	SetColor(value string) BgpSrteColorSubTlv
	// HasColor checks if Color has been set in BgpSrteColorSubTlv
	HasColor() bool
}

// Six octet values. Example: 000000000064 for color value 100.
// Color returns a string
func (obj *bgpSrteColorSubTlv) Color() string {

	return *obj.obj.Color

}

// Six octet values. Example: 000000000064 for color value 100.
// Color returns a string
func (obj *bgpSrteColorSubTlv) HasColor() bool {
	return obj.obj.Color != nil
}

// Six octet values. Example: 000000000064 for color value 100.
// SetColor sets the string value in the BgpSrteColorSubTlv object
func (obj *bgpSrteColorSubTlv) SetColor(value string) BgpSrteColorSubTlv {

	obj.obj.Color = &value
	return obj
}

func (obj *bgpSrteColorSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Color != nil {

		err := obj.validateHex(obj.Color())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpSrteColorSubTlv.Color"))
		}

	}

}

func (obj *bgpSrteColorSubTlv) setDefault() {

}
