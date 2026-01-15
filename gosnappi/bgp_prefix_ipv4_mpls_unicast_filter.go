package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4MplsUnicastFilter *****
type bgpPrefixIpv4MplsUnicastFilter struct {
	validation
	obj          *otg.BgpPrefixIpv4MplsUnicastFilter
	marshaller   marshalBgpPrefixIpv4MplsUnicastFilter
	unMarshaller unMarshalBgpPrefixIpv4MplsUnicastFilter
}

func NewBgpPrefixIpv4MplsUnicastFilter() BgpPrefixIpv4MplsUnicastFilter {
	obj := bgpPrefixIpv4MplsUnicastFilter{obj: &otg.BgpPrefixIpv4MplsUnicastFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) msg() *otg.BgpPrefixIpv4MplsUnicastFilter {
	return obj.obj
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) setMsg(msg *otg.BgpPrefixIpv4MplsUnicastFilter) BgpPrefixIpv4MplsUnicastFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4MplsUnicastFilter struct {
	obj *bgpPrefixIpv4MplsUnicastFilter
}

type marshalBgpPrefixIpv4MplsUnicastFilter interface {
	// ToProto marshals BgpPrefixIpv4MplsUnicastFilter to protobuf object *otg.BgpPrefixIpv4MplsUnicastFilter
	ToProto() (*otg.BgpPrefixIpv4MplsUnicastFilter, error)
	// ToPbText marshals BgpPrefixIpv4MplsUnicastFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4MplsUnicastFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4MplsUnicastFilter to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4MplsUnicastFilter struct {
	obj *bgpPrefixIpv4MplsUnicastFilter
}

type unMarshalBgpPrefixIpv4MplsUnicastFilter interface {
	// FromProto unmarshals BgpPrefixIpv4MplsUnicastFilter from protobuf object *otg.BgpPrefixIpv4MplsUnicastFilter
	FromProto(msg *otg.BgpPrefixIpv4MplsUnicastFilter) (BgpPrefixIpv4MplsUnicastFilter, error)
	// FromPbText unmarshals BgpPrefixIpv4MplsUnicastFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4MplsUnicastFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4MplsUnicastFilter from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) Marshal() marshalBgpPrefixIpv4MplsUnicastFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4MplsUnicastFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) Unmarshal() unMarshalBgpPrefixIpv4MplsUnicastFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4MplsUnicastFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4MplsUnicastFilter) ToProto() (*otg.BgpPrefixIpv4MplsUnicastFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4MplsUnicastFilter) FromProto(msg *otg.BgpPrefixIpv4MplsUnicastFilter) (BgpPrefixIpv4MplsUnicastFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4MplsUnicastFilter) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsUnicastFilter) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsUnicastFilter) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsUnicastFilter) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4MplsUnicastFilter) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4MplsUnicastFilter) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4MplsUnicastFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) Clone() (BgpPrefixIpv4MplsUnicastFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4MplsUnicastFilter()
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

// BgpPrefixIpv4MplsUnicastFilter is description is TBD
type BgpPrefixIpv4MplsUnicastFilter interface {
	Validation
	// msg marshals BgpPrefixIpv4MplsUnicastFilter to protobuf object *otg.BgpPrefixIpv4MplsUnicastFilter
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4MplsUnicastFilter
	// setMsg unmarshals BgpPrefixIpv4MplsUnicastFilter from protobuf object *otg.BgpPrefixIpv4MplsUnicastFilter
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4MplsUnicastFilter) BgpPrefixIpv4MplsUnicastFilter
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4MplsUnicastFilter
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4MplsUnicastFilter
	// validate validates BgpPrefixIpv4MplsUnicastFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4MplsUnicastFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns []string, set in BgpPrefixIpv4MplsUnicastFilter.
	Addresses() []string
	// SetAddresses assigns []string provided by user to BgpPrefixIpv4MplsUnicastFilter
	SetAddresses(value []string) BgpPrefixIpv4MplsUnicastFilter
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// Addresses returns a []string
func (obj *bgpPrefixIpv4MplsUnicastFilter) Addresses() []string {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	return obj.obj.Addresses
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// SetAddresses sets the []string value in the BgpPrefixIpv4MplsUnicastFilter object
func (obj *bgpPrefixIpv4MplsUnicastFilter) SetAddresses(value []string) BgpPrefixIpv4MplsUnicastFilter {

	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	obj.obj.Addresses = value

	return obj
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Addresses != nil {

		err := obj.validateIpv4Slice(obj.Addresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4MplsUnicastFilter.Addresses"))
		}

	}

}

func (obj *bgpPrefixIpv4MplsUnicastFilter) setDefault() {

}
