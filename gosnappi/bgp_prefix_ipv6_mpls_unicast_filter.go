package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6MplsUnicastFilter *****
type bgpPrefixIpv6MplsUnicastFilter struct {
	validation
	obj          *otg.BgpPrefixIpv6MplsUnicastFilter
	marshaller   marshalBgpPrefixIpv6MplsUnicastFilter
	unMarshaller unMarshalBgpPrefixIpv6MplsUnicastFilter
}

func NewBgpPrefixIpv6MplsUnicastFilter() BgpPrefixIpv6MplsUnicastFilter {
	obj := bgpPrefixIpv6MplsUnicastFilter{obj: &otg.BgpPrefixIpv6MplsUnicastFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) msg() *otg.BgpPrefixIpv6MplsUnicastFilter {
	return obj.obj
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) setMsg(msg *otg.BgpPrefixIpv6MplsUnicastFilter) BgpPrefixIpv6MplsUnicastFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6MplsUnicastFilter struct {
	obj *bgpPrefixIpv6MplsUnicastFilter
}

type marshalBgpPrefixIpv6MplsUnicastFilter interface {
	// ToProto marshals BgpPrefixIpv6MplsUnicastFilter to protobuf object *otg.BgpPrefixIpv6MplsUnicastFilter
	ToProto() (*otg.BgpPrefixIpv6MplsUnicastFilter, error)
	// ToPbText marshals BgpPrefixIpv6MplsUnicastFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6MplsUnicastFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6MplsUnicastFilter to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv6MplsUnicastFilter struct {
	obj *bgpPrefixIpv6MplsUnicastFilter
}

type unMarshalBgpPrefixIpv6MplsUnicastFilter interface {
	// FromProto unmarshals BgpPrefixIpv6MplsUnicastFilter from protobuf object *otg.BgpPrefixIpv6MplsUnicastFilter
	FromProto(msg *otg.BgpPrefixIpv6MplsUnicastFilter) (BgpPrefixIpv6MplsUnicastFilter, error)
	// FromPbText unmarshals BgpPrefixIpv6MplsUnicastFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6MplsUnicastFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6MplsUnicastFilter from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) Marshal() marshalBgpPrefixIpv6MplsUnicastFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6MplsUnicastFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) Unmarshal() unMarshalBgpPrefixIpv6MplsUnicastFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6MplsUnicastFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6MplsUnicastFilter) ToProto() (*otg.BgpPrefixIpv6MplsUnicastFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6MplsUnicastFilter) FromProto(msg *otg.BgpPrefixIpv6MplsUnicastFilter) (BgpPrefixIpv6MplsUnicastFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6MplsUnicastFilter) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsUnicastFilter) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsUnicastFilter) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsUnicastFilter) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6MplsUnicastFilter) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6MplsUnicastFilter) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6MplsUnicastFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) Clone() (BgpPrefixIpv6MplsUnicastFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6MplsUnicastFilter()
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

// BgpPrefixIpv6MplsUnicastFilter is description is TBD
type BgpPrefixIpv6MplsUnicastFilter interface {
	Validation
	// msg marshals BgpPrefixIpv6MplsUnicastFilter to protobuf object *otg.BgpPrefixIpv6MplsUnicastFilter
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6MplsUnicastFilter
	// setMsg unmarshals BgpPrefixIpv6MplsUnicastFilter from protobuf object *otg.BgpPrefixIpv6MplsUnicastFilter
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6MplsUnicastFilter) BgpPrefixIpv6MplsUnicastFilter
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6MplsUnicastFilter
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6MplsUnicastFilter
	// validate validates BgpPrefixIpv6MplsUnicastFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6MplsUnicastFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns []string, set in BgpPrefixIpv6MplsUnicastFilter.
	Addresses() []string
	// SetAddresses assigns []string provided by user to BgpPrefixIpv6MplsUnicastFilter
	SetAddresses(value []string) BgpPrefixIpv6MplsUnicastFilter
	// PrefixLength returns uint32, set in BgpPrefixIpv6MplsUnicastFilter.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6MplsUnicastFilter
	SetPrefixLength(value uint32) BgpPrefixIpv6MplsUnicastFilter
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6MplsUnicastFilter
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6MplsUnicastFilterOriginEnum, set in BgpPrefixIpv6MplsUnicastFilter
	Origin() BgpPrefixIpv6MplsUnicastFilterOriginEnum
	// SetOrigin assigns BgpPrefixIpv6MplsUnicastFilterOriginEnum provided by user to BgpPrefixIpv6MplsUnicastFilter
	SetOrigin(value BgpPrefixIpv6MplsUnicastFilterOriginEnum) BgpPrefixIpv6MplsUnicastFilter
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6MplsUnicastFilter
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6MplsUnicastFilter.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6MplsUnicastFilter
	SetPathId(value uint32) BgpPrefixIpv6MplsUnicastFilter
	// HasPathId checks if PathId has been set in BgpPrefixIpv6MplsUnicastFilter
	HasPathId() bool
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// Addresses returns a []string
func (obj *bgpPrefixIpv6MplsUnicastFilter) Addresses() []string {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	return obj.obj.Addresses
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// SetAddresses sets the []string value in the BgpPrefixIpv6MplsUnicastFilter object
func (obj *bgpPrefixIpv6MplsUnicastFilter) SetAddresses(value []string) BgpPrefixIpv6MplsUnicastFilter {

	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	obj.obj.Addresses = value

	return obj
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastFilter) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastFilter) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6MplsUnicastFilter object
func (obj *bgpPrefixIpv6MplsUnicastFilter) SetPrefixLength(value uint32) BgpPrefixIpv6MplsUnicastFilter {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6MplsUnicastFilterOriginEnum string

// Enum of Origin on BgpPrefixIpv6MplsUnicastFilter
var BgpPrefixIpv6MplsUnicastFilterOrigin = struct {
	IGP        BgpPrefixIpv6MplsUnicastFilterOriginEnum
	EGP        BgpPrefixIpv6MplsUnicastFilterOriginEnum
	INCOMPLETE BgpPrefixIpv6MplsUnicastFilterOriginEnum
}{
	IGP:        BgpPrefixIpv6MplsUnicastFilterOriginEnum("igp"),
	EGP:        BgpPrefixIpv6MplsUnicastFilterOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6MplsUnicastFilterOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) Origin() BgpPrefixIpv6MplsUnicastFilterOriginEnum {
	return BgpPrefixIpv6MplsUnicastFilterOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin to match. If the origin is missing then all origins will match.
// Origin returns a string
func (obj *bgpPrefixIpv6MplsUnicastFilter) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) SetOrigin(value BgpPrefixIpv6MplsUnicastFilterOriginEnum) BgpPrefixIpv6MplsUnicastFilter {
	intValue, ok := otg.BgpPrefixIpv6MplsUnicastFilter_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6MplsUnicastFilterOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6MplsUnicastFilter_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastFilter) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv6MplsUnicastFilter) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id to match. If the path id is missing then all path ids will match.
// SetPathId sets the uint32 value in the BgpPrefixIpv6MplsUnicastFilter object
func (obj *bgpPrefixIpv6MplsUnicastFilter) SetPathId(value uint32) BgpPrefixIpv6MplsUnicastFilter {

	obj.obj.PathId = &value
	return obj
}

func (obj *bgpPrefixIpv6MplsUnicastFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Addresses != nil {

		err := obj.validateIpv6Slice(obj.Addresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6MplsUnicastFilter.Addresses"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6MplsUnicastFilter.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *bgpPrefixIpv6MplsUnicastFilter) setDefault() {

}
