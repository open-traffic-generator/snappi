package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrtePolicyNameSubTlv *****
type bgpSrtePolicyNameSubTlv struct {
	validation
	obj          *otg.BgpSrtePolicyNameSubTlv
	marshaller   marshalBgpSrtePolicyNameSubTlv
	unMarshaller unMarshalBgpSrtePolicyNameSubTlv
}

func NewBgpSrtePolicyNameSubTlv() BgpSrtePolicyNameSubTlv {
	obj := bgpSrtePolicyNameSubTlv{obj: &otg.BgpSrtePolicyNameSubTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrtePolicyNameSubTlv) msg() *otg.BgpSrtePolicyNameSubTlv {
	return obj.obj
}

func (obj *bgpSrtePolicyNameSubTlv) setMsg(msg *otg.BgpSrtePolicyNameSubTlv) BgpSrtePolicyNameSubTlv {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrtePolicyNameSubTlv struct {
	obj *bgpSrtePolicyNameSubTlv
}

type marshalBgpSrtePolicyNameSubTlv interface {
	// ToProto marshals BgpSrtePolicyNameSubTlv to protobuf object *otg.BgpSrtePolicyNameSubTlv
	ToProto() (*otg.BgpSrtePolicyNameSubTlv, error)
	// ToPbText marshals BgpSrtePolicyNameSubTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrtePolicyNameSubTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrtePolicyNameSubTlv to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrtePolicyNameSubTlv struct {
	obj *bgpSrtePolicyNameSubTlv
}

type unMarshalBgpSrtePolicyNameSubTlv interface {
	// FromProto unmarshals BgpSrtePolicyNameSubTlv from protobuf object *otg.BgpSrtePolicyNameSubTlv
	FromProto(msg *otg.BgpSrtePolicyNameSubTlv) (BgpSrtePolicyNameSubTlv, error)
	// FromPbText unmarshals BgpSrtePolicyNameSubTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrtePolicyNameSubTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrtePolicyNameSubTlv from JSON text
	FromJson(value string) error
}

func (obj *bgpSrtePolicyNameSubTlv) Marshal() marshalBgpSrtePolicyNameSubTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrtePolicyNameSubTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrtePolicyNameSubTlv) Unmarshal() unMarshalBgpSrtePolicyNameSubTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrtePolicyNameSubTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrtePolicyNameSubTlv) ToProto() (*otg.BgpSrtePolicyNameSubTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrtePolicyNameSubTlv) FromProto(msg *otg.BgpSrtePolicyNameSubTlv) (BgpSrtePolicyNameSubTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrtePolicyNameSubTlv) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrtePolicyNameSubTlv) FromPbText(value string) error {
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

func (m *marshalbgpSrtePolicyNameSubTlv) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrtePolicyNameSubTlv) FromYaml(value string) error {
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

func (m *marshalbgpSrtePolicyNameSubTlv) ToJson() (string, error) {
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

func (m *unMarshalbgpSrtePolicyNameSubTlv) FromJson(value string) error {
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

func (obj *bgpSrtePolicyNameSubTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrtePolicyNameSubTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrtePolicyNameSubTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrtePolicyNameSubTlv) Clone() (BgpSrtePolicyNameSubTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrtePolicyNameSubTlv()
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

// BgpSrtePolicyNameSubTlv is configuration for the Policy Name sub-TLV. The Policy Name sub-TLV is used to attach a symbolic name to the SR Policy candidate path.
type BgpSrtePolicyNameSubTlv interface {
	Validation
	// msg marshals BgpSrtePolicyNameSubTlv to protobuf object *otg.BgpSrtePolicyNameSubTlv
	// and doesn't set defaults
	msg() *otg.BgpSrtePolicyNameSubTlv
	// setMsg unmarshals BgpSrtePolicyNameSubTlv from protobuf object *otg.BgpSrtePolicyNameSubTlv
	// and doesn't set defaults
	setMsg(*otg.BgpSrtePolicyNameSubTlv) BgpSrtePolicyNameSubTlv
	// provides marshal interface
	Marshal() marshalBgpSrtePolicyNameSubTlv
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrtePolicyNameSubTlv
	// validate validates BgpSrtePolicyNameSubTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrtePolicyNameSubTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PolicyName returns string, set in BgpSrtePolicyNameSubTlv.
	PolicyName() string
	// SetPolicyName assigns string provided by user to BgpSrtePolicyNameSubTlv
	SetPolicyName(value string) BgpSrtePolicyNameSubTlv
	// HasPolicyName checks if PolicyName has been set in BgpSrtePolicyNameSubTlv
	HasPolicyName() bool
}

// Symbolic name for the policy that should be a string of printable ASCII characters, without a NULL terminator.
// PolicyName returns a string
func (obj *bgpSrtePolicyNameSubTlv) PolicyName() string {

	return *obj.obj.PolicyName

}

// Symbolic name for the policy that should be a string of printable ASCII characters, without a NULL terminator.
// PolicyName returns a string
func (obj *bgpSrtePolicyNameSubTlv) HasPolicyName() bool {
	return obj.obj.PolicyName != nil
}

// Symbolic name for the policy that should be a string of printable ASCII characters, without a NULL terminator.
// SetPolicyName sets the string value in the BgpSrtePolicyNameSubTlv object
func (obj *bgpSrtePolicyNameSubTlv) SetPolicyName(value string) BgpSrtePolicyNameSubTlv {

	obj.obj.PolicyName = &value
	return obj
}

func (obj *bgpSrtePolicyNameSubTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PolicyName != nil {

		if len(*obj.obj.PolicyName) < 1 || len(*obj.obj.PolicyName) > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of BgpSrtePolicyNameSubTlv.PolicyName <= 32 but Got %d",
					len(*obj.obj.PolicyName)))
		}

	}

}

func (obj *bgpSrtePolicyNameSubTlv) setDefault() {

}
