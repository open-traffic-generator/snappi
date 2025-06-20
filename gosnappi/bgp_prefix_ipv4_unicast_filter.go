package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpPrefixIpv4UnicastFilter *****
type bgpPrefixIpv4UnicastFilter struct {
	validation
	obj          *otg.BgpPrefixIpv4UnicastFilter
	marshaller   marshalBgpPrefixIpv4UnicastFilter
	unMarshaller unMarshalBgpPrefixIpv4UnicastFilter
}

func NewBgpPrefixIpv4UnicastFilter() BgpPrefixIpv4UnicastFilter {
	obj := bgpPrefixIpv4UnicastFilter{obj: &otg.BgpPrefixIpv4UnicastFilter{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpPrefixIpv4UnicastFilter) msg() *otg.BgpPrefixIpv4UnicastFilter {
	return obj.obj
}

func (obj *bgpPrefixIpv4UnicastFilter) setMsg(msg *otg.BgpPrefixIpv4UnicastFilter) BgpPrefixIpv4UnicastFilter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpPrefixIpv4UnicastFilter struct {
	obj *bgpPrefixIpv4UnicastFilter
}

type marshalBgpPrefixIpv4UnicastFilter interface {
	// ToProto marshals BgpPrefixIpv4UnicastFilter to protobuf object *otg.BgpPrefixIpv4UnicastFilter
	ToProto() (*otg.BgpPrefixIpv4UnicastFilter, error)
	// ToPbText marshals BgpPrefixIpv4UnicastFilter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpPrefixIpv4UnicastFilter to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpPrefixIpv4UnicastFilter to JSON text
	ToJson() (string, error)
}

type unMarshalbgpPrefixIpv4UnicastFilter struct {
	obj *bgpPrefixIpv4UnicastFilter
}

type unMarshalBgpPrefixIpv4UnicastFilter interface {
	// FromProto unmarshals BgpPrefixIpv4UnicastFilter from protobuf object *otg.BgpPrefixIpv4UnicastFilter
	FromProto(msg *otg.BgpPrefixIpv4UnicastFilter) (BgpPrefixIpv4UnicastFilter, error)
	// FromPbText unmarshals BgpPrefixIpv4UnicastFilter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpPrefixIpv4UnicastFilter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpPrefixIpv4UnicastFilter from JSON text
	FromJson(value string) error
}

func (obj *bgpPrefixIpv4UnicastFilter) Marshal() marshalBgpPrefixIpv4UnicastFilter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpPrefixIpv4UnicastFilter{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpPrefixIpv4UnicastFilter) Unmarshal() unMarshalBgpPrefixIpv4UnicastFilter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpPrefixIpv4UnicastFilter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpPrefixIpv4UnicastFilter) ToProto() (*otg.BgpPrefixIpv4UnicastFilter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpPrefixIpv4UnicastFilter) FromProto(msg *otg.BgpPrefixIpv4UnicastFilter) (BgpPrefixIpv4UnicastFilter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpPrefixIpv4UnicastFilter) ToPbText() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4UnicastFilter) FromPbText(value string) error {
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

func (m *marshalbgpPrefixIpv4UnicastFilter) ToYaml() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4UnicastFilter) FromYaml(value string) error {
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

func (m *marshalbgpPrefixIpv4UnicastFilter) ToJson() (string, error) {
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

func (m *unMarshalbgpPrefixIpv4UnicastFilter) FromJson(value string) error {
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

func (obj *bgpPrefixIpv4UnicastFilter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4UnicastFilter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpPrefixIpv4UnicastFilter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpPrefixIpv4UnicastFilter) Clone() (BgpPrefixIpv4UnicastFilter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpPrefixIpv4UnicastFilter()
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

// BgpPrefixIpv4UnicastFilter is description is TBD
type BgpPrefixIpv4UnicastFilter interface {
	Validation
	// msg marshals BgpPrefixIpv4UnicastFilter to protobuf object *otg.BgpPrefixIpv4UnicastFilter
	// and doesn't set defaults
	msg() *otg.BgpPrefixIpv4UnicastFilter
	// setMsg unmarshals BgpPrefixIpv4UnicastFilter from protobuf object *otg.BgpPrefixIpv4UnicastFilter
	// and doesn't set defaults
	setMsg(*otg.BgpPrefixIpv4UnicastFilter) BgpPrefixIpv4UnicastFilter
	// provides marshal interface
	Marshal() marshalBgpPrefixIpv4UnicastFilter
	// provides unmarshal interface
	Unmarshal() unMarshalBgpPrefixIpv4UnicastFilter
	// validate validates BgpPrefixIpv4UnicastFilter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpPrefixIpv4UnicastFilter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Addresses returns []string, set in BgpPrefixIpv4UnicastFilter.
	Addresses() []string
	// SetAddresses assigns []string provided by user to BgpPrefixIpv4UnicastFilter
	SetAddresses(value []string) BgpPrefixIpv4UnicastFilter
	// PrefixLength returns uint32, set in BgpPrefixIpv4UnicastFilter.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to BgpPrefixIpv4UnicastFilter
	SetPrefixLength(value uint32) BgpPrefixIpv4UnicastFilter
	// HasPrefixLength checks if PrefixLength has been set in BgpPrefixIpv4UnicastFilter
	HasPrefixLength() bool
	// Origin returns BgpPrefixIpv4UnicastFilterOriginEnum, set in BgpPrefixIpv4UnicastFilter
	Origin() BgpPrefixIpv4UnicastFilterOriginEnum
	// SetOrigin assigns BgpPrefixIpv4UnicastFilterOriginEnum provided by user to BgpPrefixIpv4UnicastFilter
	SetOrigin(value BgpPrefixIpv4UnicastFilterOriginEnum) BgpPrefixIpv4UnicastFilter
	// HasOrigin checks if Origin has been set in BgpPrefixIpv4UnicastFilter
	HasOrigin() bool
	// PathId returns uint32, set in BgpPrefixIpv4UnicastFilter.
	PathId() uint32
	// SetPathId assigns uint32 provided by user to BgpPrefixIpv4UnicastFilter
	SetPathId(value uint32) BgpPrefixIpv4UnicastFilter
	// HasPathId checks if PathId has been set in BgpPrefixIpv4UnicastFilter
	HasPathId() bool
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// Addresses returns a []string
func (obj *bgpPrefixIpv4UnicastFilter) Addresses() []string {
	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	return obj.obj.Addresses
}

// The addresses to match. If the addresses property is missing or empty then all addresses will match.
// SetAddresses sets the []string value in the BgpPrefixIpv4UnicastFilter object
func (obj *bgpPrefixIpv4UnicastFilter) SetAddresses(value []string) BgpPrefixIpv4UnicastFilter {

	if obj.obj.Addresses == nil {
		obj.obj.Addresses = make([]string, 0)
	}
	obj.obj.Addresses = value

	return obj
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4UnicastFilter) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// PrefixLength returns a uint32
func (obj *bgpPrefixIpv4UnicastFilter) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The prefix length to match. If the prefix length is missing then all prefix lengths will match.
// SetPrefixLength sets the uint32 value in the BgpPrefixIpv4UnicastFilter object
func (obj *bgpPrefixIpv4UnicastFilter) SetPrefixLength(value uint32) BgpPrefixIpv4UnicastFilter {

	obj.obj.PrefixLength = &value
	return obj
}

type BgpPrefixIpv4UnicastFilterOriginEnum string

// Enum of Origin on BgpPrefixIpv4UnicastFilter
var BgpPrefixIpv4UnicastFilterOrigin = struct {
	IGP        BgpPrefixIpv4UnicastFilterOriginEnum
	EGP        BgpPrefixIpv4UnicastFilterOriginEnum
	INCOMPLETE BgpPrefixIpv4UnicastFilterOriginEnum
}{
	IGP:        BgpPrefixIpv4UnicastFilterOriginEnum("igp"),
	EGP:        BgpPrefixIpv4UnicastFilterOriginEnum("egp"),
	INCOMPLETE: BgpPrefixIpv4UnicastFilterOriginEnum("incomplete"),
}

func (obj *bgpPrefixIpv4UnicastFilter) Origin() BgpPrefixIpv4UnicastFilterOriginEnum {
	return BgpPrefixIpv4UnicastFilterOriginEnum(obj.obj.Origin.Enum().String())
}

// The origin to match. If the origin is missing then all origins will match.
// Origin returns a string
func (obj *bgpPrefixIpv4UnicastFilter) HasOrigin() bool {
	return obj.obj.Origin != nil
}

func (obj *bgpPrefixIpv4UnicastFilter) SetOrigin(value BgpPrefixIpv4UnicastFilterOriginEnum) BgpPrefixIpv4UnicastFilter {
	intValue, ok := otg.BgpPrefixIpv4UnicastFilter_Origin_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpPrefixIpv4UnicastFilterOriginEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpPrefixIpv4UnicastFilter_Origin_Enum(intValue)
	obj.obj.Origin = &enumValue

	return obj
}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv4UnicastFilter) PathId() uint32 {

	return *obj.obj.PathId

}

// The path id to match. If the path id is missing then all path ids will match.
// PathId returns a uint32
func (obj *bgpPrefixIpv4UnicastFilter) HasPathId() bool {
	return obj.obj.PathId != nil
}

// The path id to match. If the path id is missing then all path ids will match.
// SetPathId sets the uint32 value in the BgpPrefixIpv4UnicastFilter object
func (obj *bgpPrefixIpv4UnicastFilter) SetPathId(value uint32) BgpPrefixIpv4UnicastFilter {

	obj.obj.PathId = &value
	return obj
}

func (obj *bgpPrefixIpv4UnicastFilter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Addresses != nil {

		err := obj.validateIpv4Slice(obj.Addresses())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on BgpPrefixIpv4UnicastFilter.Addresses"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= BgpPrefixIpv4UnicastFilter.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *bgpPrefixIpv4UnicastFilter) setDefault() {

}
