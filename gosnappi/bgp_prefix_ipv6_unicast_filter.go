package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv6UnicastFilter *****
type bgpPrefixIpv6UnicastFilter struct {
	validation
	obj          *otg.BgpPrefixIpv6UnicastFilter
	marshaller   marshalBgpPrefixIpv6UnicastFilter
	unMarshaller unMarshalBgpPrefixIpv6UnicastFilter
}

func NewBgpPrefixIpv6UnicastFilter() BgpPrefixIpv6UnicastFilter {
	obj := bgpPrefixIpv6UnicastFilter{obj: &otg.BgpPrefixIpv6UnicastFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv6UnicastFilter) msg() *otg.BgpPrefixIpv6UnicastFilter {
	return obj.obj
}

func (obj *bgpPrefixIpv6UnicastFilter) setMsg(msg *otg.BgpPrefixIpv6UnicastFilter) BgpPrefixIpv6UnicastFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv6UnicastFilter struct {
	obj *bgpPrefixIpv6UnicastFilter
}

type marshalBgpPrefixIpv6UnicastFilter interface {
	// ToProto marshals BgpPrefixIpv6UnicastFilter to protobuf object *otg.BgpPrefixIpv6UnicastFilter
	ToProto() (*otg.BgpPrefixIpv6UnicastFilter, error)
	// ToPbText marshals BgpPrefixIpv6UnicastFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv6UnicastFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv6UnicastFilter to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpPrefixIpv6UnicastFilter to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpPrefixIpv6UnicastFilter struct {
	obj *bgpPrefixIpv6UnicastFilter
}

type unMarshalBgpPrefixIpv6UnicastFilter interface {
	// FromProto unmarshals BgpPrefixIpv6UnicastFilter from protobuf object *otg.BgpPrefixIpv6UnicastFilter
	FromProto(msg *otg.BgpPrefixIpv6UnicastFilter) (BgpPrefixIpv6UnicastFilter, error)
	// FromPbText unmarshals BgpPrefixIpv6UnicastFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv6UnicastFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv6UnicastFilter from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv6UnicastFilter) Marshal() marshalBgpPrefixIpv6UnicastFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv6UnicastFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv6UnicastFilter) Unmarshal() unMarshalBgpPrefixIpv6UnicastFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv6UnicastFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv6UnicastFilter) ToProto() (*otg.BgpPrefixIpv6UnicastFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv6UnicastFilter) FromProto(msg *otg.BgpPrefixIpv6UnicastFilter) (BgpPrefixIpv6UnicastFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv6UnicastFilter) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6UnicastFilter) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv6UnicastFilter) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6UnicastFilter) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv6UnicastFilter) ToJsonRaw() (string, error) {
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

func (m *marshalbgpPrefixIpv6UnicastFilter) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv6UnicastFilter) FromJson(value string) error {
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

func (obj *bgpPrefixIpv6UnicastFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6UnicastFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv6UnicastFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv6UnicastFilter) Clone() (BgpPrefixIpv6UnicastFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv6UnicastFilter()
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

// BgpPrefixIpv6UnicastFilter is description is TBD
type BgpPrefixIpv6UnicastFilter interface {
	Validation
	// msg marshals BgpPrefixIpv6UnicastFilter to protobuf object *otg.BgpPrefixIpv6UnicastFilter
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv6UnicastFilter
	// setMsg unmarshals BgpPrefixIpv6UnicastFilter from protobuf object *otg.BgpPrefixIpv6UnicastFilter
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv6UnicastFilter) BgpPrefixIpv6UnicastFilter
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv6UnicastFilter
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv6UnicastFilter
	// validate validates BgpPrefixIpv6UnicastFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv6UnicastFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns []string, set in BgpPrefixIpv6UnicastFilter.
	Addresses() []string
	// SetAddresses assigns []string provided by user to BgpPrefixIpv6UnicastFilter
	SetAddresses(value []string) BgpPrefixIpv6UnicastFilter
	// PrefixLength returns uint32, set in BgpPrefixIpv6UnicastFilter.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv6UnicastFilter
	SetPrefixLength(value uint32) BgpPrefixIpv6UnicastFilter
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv6UnicastFilter
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv6UnicastFilterOriginEnum, set in BgpPrefixIpv6UnicastFilter
	Origin() BgpPrefixIpv6UnicastFilterOriginEnum
	// SetOrigin assigns BgpPrefixIpv6UnicastFilterOriginEnum provided by user to BgpPrefixIpv6UnicastFilter
	SetOrigin(value BgpPrefixIpv6UnicastFilterOriginEnum) BgpPrefixIpv6UnicastFilter
	// HasOrigin checks if Origin has been set in BgpPrefixIpv6UnicastFilter
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv6UnicastFilter.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv6UnicastFilter
	SetPathId(value uint32) BgpPrefixIpv6UnicastFilter
	// HasPathId checks if PathId has been set in BgpPrefixIpv6UnicastFilter
	HasPathId() bool
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// Addresses returns a []string
func (obj *bgpPrefixIpv6UnicastFilter) Addresses() []string {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	return obj.obj.Addresses
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// SetAddresses sets the []string value in the BgpPrefixIpv6UnicastFilter object
func (obj *bgpPrefixIpv6UnicastFilter) SetAddresses(value []string) BgpPrefixIpv6UnicastFilter {

	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	obj.obj.Addresses = value

	return obj
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6UnicastFilter) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv6UnicastFilter) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv6UnicastFilter object
func (obj *bgpPrefixIpv6UnicastFilter) SetPrefixLength(value uint32) BgpPrefixIpv6UnicastFilter {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv6UnicastFilterOriginEnum string

// Enum of Origin on BgpPrefixIpv6UnicastFilter
var BgpPrefixIpv6UnicastFilterOrigin = struct {
	IGP        BgpPrefixIpv6UnicastFilterOriginEnum
	EGP        BgpPrefixIpv6UnicastFilterOriginEnum
	INCOMPLETE BgpPrefixIpv6UnicastFilterOriginEnum
}{
	IGP:        BgpPrefixIpv6UnicastFilterOriginEnum("igp"),
	EGP:        BgpPrefixIpv6UnicastFilterOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv6UnicastFilterOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv6UnicastFilter) Origin() BgpPrefixIpv6UnicastFilterOriginEnum {
	return BgpPrefixIpv6UnicastFilterOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin to match. If the origin is missing then all origins will match.
// Origin returns a string
func (obj *bgpPrefixIpv6UnicastFilter) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv6UnicastFilter) SetOrigin(value BgpPrefixIpv6UnicastFilterOriginEnum) BgpPrefixIpv6UnicastFilter {
	intValue, ok := otg.BgpPrefixIpv6UnicastFilter_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv6UnicastFilterOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv6UnicastFilter_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv6UnicastFilter) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv6UnicastFilter) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id to match. If the path id is missing then all path ids will match.
// SetPathId sets the uint32 value in the BgpPrefixIpv6UnicastFilter object
func (obj *bgpPrefixIpv6UnicastFilter) SetPathId(value uint32) BgpPrefixIpv6UnicastFilter {

	obj.obj.PathId = &value
	return obj
}

func (obj *bgpPrefixIpv6UnicastFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Addresses != nil {

		err := obj.validateIpv6Slice(obj.Addresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv6UnicastFilter.Addresses"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv6UnicastFilter.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *bgpPrefixIpv6UnicastFilter) setDefault() {

}
