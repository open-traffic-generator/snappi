package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrtePolicyPrioritySubTlv *****
type bgpSrtePolicyPrioritySubTlv struct {
	validation
	obj          *otg.BgpSrtePolicyPrioritySubTlv
	marshaller   marshalBgpSrtePolicyPrioritySubTlv
	unMarshaller unMarshalBgpSrtePolicyPrioritySubTlv
}

func NewBgpSrtePolicyPrioritySubTlv() BgpSrtePolicyPrioritySubTlv {
	obj := bgpSrtePolicyPrioritySubTlv{obj: &otg.BgpSrtePolicyPrioritySubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrtePolicyPrioritySubTlv) msg() *otg.BgpSrtePolicyPrioritySubTlv {
	return obj.obj
}

func (obj *bgpSrtePolicyPrioritySubTlv) setMsg(msg *otg.BgpSrtePolicyPrioritySubTlv) BgpSrtePolicyPrioritySubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrtePolicyPrioritySubTlv struct {
	obj *bgpSrtePolicyPrioritySubTlv
}

type marshalBgpSrtePolicyPrioritySubTlv interface {
	// ToProto marshals BgpSrtePolicyPrioritySubTlv to protobuf object *otg.BgpSrtePolicyPrioritySubTlv
	ToProto() (*otg.BgpSrtePolicyPrioritySubTlv, error)
	// ToPbText marshals BgpSrtePolicyPrioritySubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrtePolicyPrioritySubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrtePolicyPrioritySubTlv to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpSrtePolicyPrioritySubTlv to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpSrtePolicyPrioritySubTlv struct {
	obj *bgpSrtePolicyPrioritySubTlv
}

type unMarshalBgpSrtePolicyPrioritySubTlv interface {
	// FromProto unmarshals BgpSrtePolicyPrioritySubTlv from protobuf object *otg.BgpSrtePolicyPrioritySubTlv
	FromProto(msg *otg.BgpSrtePolicyPrioritySubTlv) (BgpSrtePolicyPrioritySubTlv, error)
	// FromPbText unmarshals BgpSrtePolicyPrioritySubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrtePolicyPrioritySubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrtePolicyPrioritySubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrtePolicyPrioritySubTlv) Marshal() marshalBgpSrtePolicyPrioritySubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrtePolicyPrioritySubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrtePolicyPrioritySubTlv) Unmarshal() unMarshalBgpSrtePolicyPrioritySubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrtePolicyPrioritySubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrtePolicyPrioritySubTlv) ToProto() (*otg.BgpSrtePolicyPrioritySubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrtePolicyPrioritySubTlv) FromProto(msg *otg.BgpSrtePolicyPrioritySubTlv) (BgpSrtePolicyPrioritySubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrtePolicyPrioritySubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrtePolicyPrioritySubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrtePolicyPrioritySubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrtePolicyPrioritySubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrtePolicyPrioritySubTlv) ToJsonRaw() (string, error) {
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

func (m *marshalbgpSrtePolicyPrioritySubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrtePolicyPrioritySubTlv) FromJson(value string) error {
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

func (obj *bgpSrtePolicyPrioritySubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrtePolicyPrioritySubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrtePolicyPrioritySubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrtePolicyPrioritySubTlv) Clone() (BgpSrtePolicyPrioritySubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrtePolicyPrioritySubTlv()
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

// BgpSrtePolicyPrioritySubTlv is configuration for the Policy Priority sub-TLV. The Policy Priority to indicate the order in which the SR policies  are re-computed upon topological change.
type BgpSrtePolicyPrioritySubTlv interface {
	Validation
	// msg marshals BgpSrtePolicyPrioritySubTlv to protobuf object *otg.BgpSrtePolicyPrioritySubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrtePolicyPrioritySubTlv
	// setMsg unmarshals BgpSrtePolicyPrioritySubTlv from protobuf object *otg.BgpSrtePolicyPrioritySubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrtePolicyPrioritySubTlv) BgpSrtePolicyPrioritySubTlv
	// provides marshal interface
	Marshal() marshalBgpSrtePolicyPrioritySubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrtePolicyPrioritySubTlv
	// validate validates BgpSrtePolicyPrioritySubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrtePolicyPrioritySubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PolicyPriority returns uint32, set in BgpSrtePolicyPrioritySubTlv.
	PolicyPriority() uint32
	// SetPolicyPriority assigns uint32 provided by user to BgpSrtePolicyPrioritySubTlv
	SetPolicyPriority(value uint32) BgpSrtePolicyPrioritySubTlv
	// HasPolicyPriority checks if PolicyPriority has been set in BgpSrtePolicyPrioritySubTlv
	HasPolicyPriority() bool
}

// One-octet Priority value.
// PolicyPriority returns a uint32
func (obj *bgpSrtePolicyPrioritySubTlv) PolicyPriority() uint32 {

	return *obj.obj.PolicyPriority

}

// One-octet Priority value.
// PolicyPriority returns a uint32
func (obj *bgpSrtePolicyPrioritySubTlv) HasPolicyPriority() bool {
	return obj.obj.PolicyPriority != nil
}

// One-octet Priority value.
// SetPolicyPriority sets the uint32 value in the BgpSrtePolicyPrioritySubTlv object
func (obj *bgpSrtePolicyPrioritySubTlv) SetPolicyPriority(value uint32) BgpSrtePolicyPrioritySubTlv {

	obj.obj.PolicyPriority = &value
	return obj
}

func (obj *bgpSrtePolicyPrioritySubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PolicyPriority != nil {

		if *obj.obj.PolicyPriority > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpSrtePolicyPrioritySubTlv.PolicyPriority <= 255 but Got %d", *obj.obj.PolicyPriority))
		}

	}

}

func (obj *bgpSrtePolicyPrioritySubTlv) setDefault() {

}
