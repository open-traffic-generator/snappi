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
	// PrefixLength returns uint32, set in BgpPrefixIpv4MplsUnicastFilter.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4MplsUnicastFilter
	SetPrefixLength(value uint32) BgpPrefixIpv4MplsUnicastFilter
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4MplsUnicastFilter
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4MplsUnicastFilterOriginEnum, set in BgpPrefixIpv4MplsUnicastFilter
	Origin() BgpPrefixIpv4MplsUnicastFilterOriginEnum
	// SetOrigin assigns BgpPrefixIpv4MplsUnicastFilterOriginEnum provided by user to BgpPrefixIpv4MplsUnicastFilter
	SetOrigin(value BgpPrefixIpv4MplsUnicastFilterOriginEnum) BgpPrefixIpv4MplsUnicastFilter
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4MplsUnicastFilter
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4MplsUnicastFilter.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4MplsUnicastFilter
	SetPathId(value uint32) BgpPrefixIpv4MplsUnicastFilter
	// HasPathId checks if PathId has been set in BgpPrefixIpv4MplsUnicastFilter
	HasPathId() bool
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

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastFilter) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastFilter) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4MplsUnicastFilter object
func (obj *bgpPrefixIpv4MplsUnicastFilter) SetPrefixLength(value uint32) BgpPrefixIpv4MplsUnicastFilter {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4MplsUnicastFilterOriginEnum string

// Enum of Origin on BgpPrefixIpv4MplsUnicastFilter
var BgpPrefixIpv4MplsUnicastFilterOrigin = struct {
	IGP        BgpPrefixIpv4MplsUnicastFilterOriginEnum
	EGP        BgpPrefixIpv4MplsUnicastFilterOriginEnum
	INCOMPLETE BgpPrefixIpv4MplsUnicastFilterOriginEnum
}{
	IGP:        BgpPrefixIpv4MplsUnicastFilterOriginEnum("igp"),
	EGP:        BgpPrefixIpv4MplsUnicastFilterOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4MplsUnicastFilterOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) Origin() BgpPrefixIpv4MplsUnicastFilterOriginEnum {
	return BgpPrefixIpv4MplsUnicastFilterOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin to match. If the origin is missing then all origins will match.
// Origin returns a string
func (obj *bgpPrefixIpv4MplsUnicastFilter) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4MplsUnicastFilter) SetOrigin(value BgpPrefixIpv4MplsUnicastFilterOriginEnum) BgpPrefixIpv4MplsUnicastFilter {
	intValue, ok := otg.BgpPrefixIpv4MplsUnicastFilter_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4MplsUnicastFilterOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4MplsUnicastFilter_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastFilter) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv4MplsUnicastFilter) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id to match. If the path id is missing then all path ids will match.
// SetPathId sets the uint32 value in the BgpPrefixIpv4MplsUnicastFilter object
func (obj *bgpPrefixIpv4MplsUnicastFilter) SetPathId(value uint32) BgpPrefixIpv4MplsUnicastFilter {

	obj.obj.PathId = &value
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

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4MplsUnicastFilter.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *bgpPrefixIpv4MplsUnicastFilter) setDefault() {

}
