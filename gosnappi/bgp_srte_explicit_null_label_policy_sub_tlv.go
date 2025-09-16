package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteExplicitNullLabelPolicySubTlv *****
type bgpSrteExplicitNullLabelPolicySubTlv struct {
	validation
	obj          *otg.BgpSrteExplicitNullLabelPolicySubTlv
	marshaller   marshalBgpSrteExplicitNullLabelPolicySubTlv
	unMarshaller unMarshalBgpSrteExplicitNullLabelPolicySubTlv
}

func NewBgpSrteExplicitNullLabelPolicySubTlv() BgpSrteExplicitNullLabelPolicySubTlv {
	obj := bgpSrteExplicitNullLabelPolicySubTlv{obj: &otg.BgpSrteExplicitNullLabelPolicySubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) msg() *otg.BgpSrteExplicitNullLabelPolicySubTlv {
	return obj.obj
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) setMsg(msg *otg.BgpSrteExplicitNullLabelPolicySubTlv) BgpSrteExplicitNullLabelPolicySubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteExplicitNullLabelPolicySubTlv struct {
	obj *bgpSrteExplicitNullLabelPolicySubTlv
}

type marshalBgpSrteExplicitNullLabelPolicySubTlv interface {
	// ToProto marshals BgpSrteExplicitNullLabelPolicySubTlv to protobuf object *otg.BgpSrteExplicitNullLabelPolicySubTlv
	ToProto() (*otg.BgpSrteExplicitNullLabelPolicySubTlv, error)
	// ToPbText marshals BgpSrteExplicitNullLabelPolicySubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteExplicitNullLabelPolicySubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteExplicitNullLabelPolicySubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteExplicitNullLabelPolicySubTlv struct {
	obj *bgpSrteExplicitNullLabelPolicySubTlv
}

type unMarshalBgpSrteExplicitNullLabelPolicySubTlv interface {
	// FromProto unmarshals BgpSrteExplicitNullLabelPolicySubTlv from protobuf object *otg.BgpSrteExplicitNullLabelPolicySubTlv
	FromProto(msg *otg.BgpSrteExplicitNullLabelPolicySubTlv) (BgpSrteExplicitNullLabelPolicySubTlv, error)
	// FromPbText unmarshals BgpSrteExplicitNullLabelPolicySubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteExplicitNullLabelPolicySubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteExplicitNullLabelPolicySubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) Marshal() marshalBgpSrteExplicitNullLabelPolicySubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteExplicitNullLabelPolicySubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) Unmarshal() unMarshalBgpSrteExplicitNullLabelPolicySubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteExplicitNullLabelPolicySubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteExplicitNullLabelPolicySubTlv) ToProto() (*otg.BgpSrteExplicitNullLabelPolicySubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteExplicitNullLabelPolicySubTlv) FromProto(msg *otg.BgpSrteExplicitNullLabelPolicySubTlv) (BgpSrteExplicitNullLabelPolicySubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteExplicitNullLabelPolicySubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteExplicitNullLabelPolicySubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrteExplicitNullLabelPolicySubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteExplicitNullLabelPolicySubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrteExplicitNullLabelPolicySubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteExplicitNullLabelPolicySubTlv) FromJson(value string) error {
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

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) Clone() (BgpSrteExplicitNullLabelPolicySubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteExplicitNullLabelPolicySubTlv()
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

// BgpSrteExplicitNullLabelPolicySubTlv is configuration for BGP explicit null label policy sub TLV settings.
type BgpSrteExplicitNullLabelPolicySubTlv interface {
	Validation
	// msg marshals BgpSrteExplicitNullLabelPolicySubTlv to protobuf object *otg.BgpSrteExplicitNullLabelPolicySubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrteExplicitNullLabelPolicySubTlv
	// setMsg unmarshals BgpSrteExplicitNullLabelPolicySubTlv from protobuf object *otg.BgpSrteExplicitNullLabelPolicySubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrteExplicitNullLabelPolicySubTlv) BgpSrteExplicitNullLabelPolicySubTlv
	// provides marshal interface
	Marshal() marshalBgpSrteExplicitNullLabelPolicySubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteExplicitNullLabelPolicySubTlv
	// validate validates BgpSrteExplicitNullLabelPolicySubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteExplicitNullLabelPolicySubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ExplicitNullLabelPolicy returns BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum, set in BgpSrteExplicitNullLabelPolicySubTlv
	ExplicitNullLabelPolicy() BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum
	// SetExplicitNullLabelPolicy assigns BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum provided by user to BgpSrteExplicitNullLabelPolicySubTlv
	SetExplicitNullLabelPolicy(value BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum) BgpSrteExplicitNullLabelPolicySubTlv
	// HasExplicitNullLabelPolicy checks if ExplicitNullLabelPolicy has been set in BgpSrteExplicitNullLabelPolicySubTlv
	HasExplicitNullLabelPolicy() bool
}

type BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum string

// Enum of ExplicitNullLabelPolicy on BgpSrteExplicitNullLabelPolicySubTlv
var BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicy = struct {
	RESERVED_ENLP       BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum
	PUSH_IPV4_ENLP      BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum
	PUSH_IPV6_ENLP      BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum
	PUSH_IPV4_IPV6_ENLP BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum
	DO_NOT_PUSH_ENLP    BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum
}{
	RESERVED_ENLP:       BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum("reserved_enlp"),
	PUSH_IPV4_ENLP:      BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum("push_ipv4_enlp"),
	PUSH_IPV6_ENLP:      BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum("push_ipv6_enlp"),
	PUSH_IPV4_IPV6_ENLP: BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum("push_ipv4_ipv6_enlp"),
	DO_NOT_PUSH_ENLP:    BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum("do_not_push_enlp"),
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) ExplicitNullLabelPolicy() BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum {
	return BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum(obj.obj.ExplicitNullLabelPolicy.Enum().String())
}

// The value of the explicit null label policy
// ExplicitNullLabelPolicy returns a string
func (obj *bgpSrteExplicitNullLabelPolicySubTlv) HasExplicitNullLabelPolicy() bool {
	return obj.obj.ExplicitNullLabelPolicy != nil
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) SetExplicitNullLabelPolicy(value BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum) BgpSrteExplicitNullLabelPolicySubTlv {
	intValue, ok := otg.BgpSrteExplicitNullLabelPolicySubTlv_ExplicitNullLabelPolicy_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicyEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpSrteExplicitNullLabelPolicySubTlv_ExplicitNullLabelPolicy_Enum(intValue)
	obj.obj.ExplicitNullLabelPolicy = &enumValue

	return obj
}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bgpSrteExplicitNullLabelPolicySubTlv) setDefault() {
	if obj.obj.ExplicitNullLabelPolicy == nil {
		obj.SetExplicitNullLabelPolicy(BgpSrteExplicitNullLabelPolicySubTlvExplicitNullLabelPolicy.DO_NOT_PUSH_ENLP)

	}

}
