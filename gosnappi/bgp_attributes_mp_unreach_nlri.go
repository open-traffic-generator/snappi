package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesMpUnreachNlri *****
type bgpAttributesMpUnreachNlri struct {
	validation
	obj                *otg.BgpAttributesMpUnreachNlri
	marshaller         marshalBgpAttributesMpUnreachNlri
	unMarshaller       unMarshalBgpAttributesMpUnreachNlri
	ipv4UnicastHolder  BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	ipv6UnicastHolder  BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	ipv4SrpolicyHolder BgpIpv4SrPolicyNLRIPrefix
	ipv6SrpolicyHolder BgpIpv6SrPolicyNLRIPrefix
}

func NewBgpAttributesMpUnreachNlri() BgpAttributesMpUnreachNlri {
	obj := bgpAttributesMpUnreachNlri{obj: &otg.BgpAttributesMpUnreachNlri{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesMpUnreachNlri) msg() *otg.BgpAttributesMpUnreachNlri {
	return obj.obj
}

func (obj *bgpAttributesMpUnreachNlri) setMsg(msg *otg.BgpAttributesMpUnreachNlri) BgpAttributesMpUnreachNlri {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesMpUnreachNlri struct {
	obj *bgpAttributesMpUnreachNlri
}

type marshalBgpAttributesMpUnreachNlri interface {
	// ToProto marshals BgpAttributesMpUnreachNlri to protobuf object *otg.BgpAttributesMpUnreachNlri
	ToProto() (*otg.BgpAttributesMpUnreachNlri, error)
	// ToPbText marshals BgpAttributesMpUnreachNlri to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesMpUnreachNlri to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesMpUnreachNlri to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesMpUnreachNlri to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesMpUnreachNlri struct {
	obj *bgpAttributesMpUnreachNlri
}

type unMarshalBgpAttributesMpUnreachNlri interface {
	// FromProto unmarshals BgpAttributesMpUnreachNlri from protobuf object *otg.BgpAttributesMpUnreachNlri
	FromProto(msg *otg.BgpAttributesMpUnreachNlri) (BgpAttributesMpUnreachNlri, error)
	// FromPbText unmarshals BgpAttributesMpUnreachNlri from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesMpUnreachNlri from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesMpUnreachNlri from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesMpUnreachNlri) Marshal() marshalBgpAttributesMpUnreachNlri {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesMpUnreachNlri{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesMpUnreachNlri) Unmarshal() unMarshalBgpAttributesMpUnreachNlri {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesMpUnreachNlri{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesMpUnreachNlri) ToProto() (*otg.BgpAttributesMpUnreachNlri, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesMpUnreachNlri) FromProto(msg *otg.BgpAttributesMpUnreachNlri) (BgpAttributesMpUnreachNlri, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesMpUnreachNlri) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesMpUnreachNlri) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalbgpAttributesMpUnreachNlri) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesMpUnreachNlri) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalbgpAttributesMpUnreachNlri) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesMpUnreachNlri) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesMpUnreachNlri) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *bgpAttributesMpUnreachNlri) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesMpUnreachNlri) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesMpUnreachNlri) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesMpUnreachNlri) Clone() (BgpAttributesMpUnreachNlri, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesMpUnreachNlri()
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

func (obj *bgpAttributesMpUnreachNlri) setNil() {
	obj.ipv4UnicastHolder = nil
	obj.ipv6UnicastHolder = nil
	obj.ipv4SrpolicyHolder = nil
	obj.ipv6SrpolicyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesMpUnreachNlri is the MP_UNREACH attribute is an optional attribute which can be included in the attributes of a BGP Update message as defined in https://datatracker.ietf.org/doc/html/rfc4760#section-3.
// The following AFI / SAFI combinations are supported:
// - IPv4 Unicast with AFI as 1 and SAFI as 1
// - IPv6 Unicast with AFI as 2 and SAFI as 1
// - Segment Routing Policy for IPv4 Unicast with AFI as 1 and SAFI as 73 (draft-ietf-idr-sr-policy-safi-02 Section 2.1)
// - Segment Routing Policy for IPv6 Unicast with AFI as 2 and SAFI as 73
type BgpAttributesMpUnreachNlri interface {
	Validation
	// msg marshals BgpAttributesMpUnreachNlri to protobuf object *otg.BgpAttributesMpUnreachNlri
	// and doesn't set defaults
	msg() *otg.BgpAttributesMpUnreachNlri
	// setMsg unmarshals BgpAttributesMpUnreachNlri from protobuf object *otg.BgpAttributesMpUnreachNlri
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesMpUnreachNlri) BgpAttributesMpUnreachNlri
	// provides marshal interface
	Marshal() marshalBgpAttributesMpUnreachNlri
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesMpUnreachNlri
	// validate validates BgpAttributesMpUnreachNlri
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesMpUnreachNlri, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpAttributesMpUnreachNlriChoiceEnum, set in BgpAttributesMpUnreachNlri
	Choice() BgpAttributesMpUnreachNlriChoiceEnum
	// setChoice assigns BgpAttributesMpUnreachNlriChoiceEnum provided by user to BgpAttributesMpUnreachNlri
	setChoice(value BgpAttributesMpUnreachNlriChoiceEnum) BgpAttributesMpUnreachNlri
	// HasChoice checks if Choice has been set in BgpAttributesMpUnreachNlri
	HasChoice() bool
	// Ipv4Unicast returns BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIterIter, set in BgpAttributesMpUnreachNlri
	Ipv4Unicast() BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	// Ipv6Unicast returns BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIterIter, set in BgpAttributesMpUnreachNlri
	Ipv6Unicast() BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	// Ipv4Srpolicy returns BgpIpv4SrPolicyNLRIPrefix, set in BgpAttributesMpUnreachNlri.
	// BgpIpv4SrPolicyNLRIPrefix is iPv4 Segment Routing Policy NLRI Prefix.
	Ipv4Srpolicy() BgpIpv4SrPolicyNLRIPrefix
	// SetIpv4Srpolicy assigns BgpIpv4SrPolicyNLRIPrefix provided by user to BgpAttributesMpUnreachNlri.
	// BgpIpv4SrPolicyNLRIPrefix is iPv4 Segment Routing Policy NLRI Prefix.
	SetIpv4Srpolicy(value BgpIpv4SrPolicyNLRIPrefix) BgpAttributesMpUnreachNlri
	// HasIpv4Srpolicy checks if Ipv4Srpolicy has been set in BgpAttributesMpUnreachNlri
	HasIpv4Srpolicy() bool
	// Ipv6Srpolicy returns BgpIpv6SrPolicyNLRIPrefix, set in BgpAttributesMpUnreachNlri.
	// BgpIpv6SrPolicyNLRIPrefix is one IPv6 Segment Routing Policy NLRI Prefix.
	Ipv6Srpolicy() BgpIpv6SrPolicyNLRIPrefix
	// SetIpv6Srpolicy assigns BgpIpv6SrPolicyNLRIPrefix provided by user to BgpAttributesMpUnreachNlri.
	// BgpIpv6SrPolicyNLRIPrefix is one IPv6 Segment Routing Policy NLRI Prefix.
	SetIpv6Srpolicy(value BgpIpv6SrPolicyNLRIPrefix) BgpAttributesMpUnreachNlri
	// HasIpv6Srpolicy checks if Ipv6Srpolicy has been set in BgpAttributesMpUnreachNlri
	HasIpv6Srpolicy() bool
	setNil()
}

type BgpAttributesMpUnreachNlriChoiceEnum string

// Enum of Choice on BgpAttributesMpUnreachNlri
var BgpAttributesMpUnreachNlriChoice = struct {
	IPV4_UNICAST  BgpAttributesMpUnreachNlriChoiceEnum
	IPV6_UNICAST  BgpAttributesMpUnreachNlriChoiceEnum
	IPV4_SRPOLICY BgpAttributesMpUnreachNlriChoiceEnum
	IPV6_SRPOLICY BgpAttributesMpUnreachNlriChoiceEnum
}{
	IPV4_UNICAST:  BgpAttributesMpUnreachNlriChoiceEnum("ipv4_unicast"),
	IPV6_UNICAST:  BgpAttributesMpUnreachNlriChoiceEnum("ipv6_unicast"),
	IPV4_SRPOLICY: BgpAttributesMpUnreachNlriChoiceEnum("ipv4_srpolicy"),
	IPV6_SRPOLICY: BgpAttributesMpUnreachNlriChoiceEnum("ipv6_srpolicy"),
}

func (obj *bgpAttributesMpUnreachNlri) Choice() BgpAttributesMpUnreachNlriChoiceEnum {
	return BgpAttributesMpUnreachNlriChoiceEnum(obj.obj.Choice.Enum().String())
}

// The AFI and SAFI to be sent in the MPUNREACH_NLRI in the Update.
// Choice returns a string
func (obj *bgpAttributesMpUnreachNlri) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpAttributesMpUnreachNlri) setChoice(value BgpAttributesMpUnreachNlriChoiceEnum) BgpAttributesMpUnreachNlri {
	intValue, ok := otg.BgpAttributesMpUnreachNlri_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesMpUnreachNlriChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesMpUnreachNlri_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6Srpolicy = nil
	obj.ipv6SrpolicyHolder = nil
	obj.obj.Ipv4Srpolicy = nil
	obj.ipv4SrpolicyHolder = nil
	obj.obj.Ipv6Unicast = nil
	obj.ipv6UnicastHolder = nil
	obj.obj.Ipv4Unicast = nil
	obj.ipv4UnicastHolder = nil

	if value == BgpAttributesMpUnreachNlriChoice.IPV4_UNICAST {
		obj.obj.Ipv4Unicast = []*otg.BgpOneIpv4NLRIPrefix{}
	}

	if value == BgpAttributesMpUnreachNlriChoice.IPV6_UNICAST {
		obj.obj.Ipv6Unicast = []*otg.BgpOneIpv6NLRIPrefix{}
	}

	if value == BgpAttributesMpUnreachNlriChoice.IPV4_SRPOLICY {
		obj.obj.Ipv4Srpolicy = NewBgpIpv4SrPolicyNLRIPrefix().msg()
	}

	if value == BgpAttributesMpUnreachNlriChoice.IPV6_SRPOLICY {
		obj.obj.Ipv6Srpolicy = NewBgpIpv6SrPolicyNLRIPrefix().msg()
	}

	return obj
}

// List of IPv4 prefixes being sent in the IPv4 Unicast MPUNREACH_NLRI .
// Ipv4Unicast returns a []BgpOneIpv4NLRIPrefix
func (obj *bgpAttributesMpUnreachNlri) Ipv4Unicast() BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	if len(obj.obj.Ipv4Unicast) == 0 {
		obj.setChoice(BgpAttributesMpUnreachNlriChoice.IPV4_UNICAST)
	}
	if obj.ipv4UnicastHolder == nil {
		obj.ipv4UnicastHolder = newBgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter(&obj.obj.Ipv4Unicast).setMsg(obj)
	}
	return obj.ipv4UnicastHolder
}

type bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter struct {
	obj                       *bgpAttributesMpUnreachNlri
	bgpOneIpv4NLRIPrefixSlice []BgpOneIpv4NLRIPrefix
	fieldPtr                  *[]*otg.BgpOneIpv4NLRIPrefix
}

func newBgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter(ptr *[]*otg.BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	return &bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter{fieldPtr: ptr}
}

type BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter interface {
	setMsg(*bgpAttributesMpUnreachNlri) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	Items() []BgpOneIpv4NLRIPrefix
	Add() BgpOneIpv4NLRIPrefix
	Append(items ...BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	Set(index int, newObj BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	Clear() BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	clearHolderSlice() BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
	appendHolderSlice(item BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) setMsg(msg *bgpAttributesMpUnreachNlri) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneIpv4NLRIPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) Items() []BgpOneIpv4NLRIPrefix {
	return obj.bgpOneIpv4NLRIPrefixSlice
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) Add() BgpOneIpv4NLRIPrefix {
	newObj := &otg.BgpOneIpv4NLRIPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneIpv4NLRIPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneIpv4NLRIPrefixSlice = append(obj.bgpOneIpv4NLRIPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) Append(items ...BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneIpv4NLRIPrefixSlice = append(obj.bgpOneIpv4NLRIPrefixSlice, item)
	}
	return obj
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) Set(index int, newObj BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneIpv4NLRIPrefixSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) Clear() BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneIpv4NLRIPrefix{}
		obj.bgpOneIpv4NLRIPrefixSlice = []BgpOneIpv4NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) clearHolderSlice() BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	if len(obj.bgpOneIpv4NLRIPrefixSlice) > 0 {
		obj.bgpOneIpv4NLRIPrefixSlice = []BgpOneIpv4NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter) appendHolderSlice(item BgpOneIpv4NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv4NLRIPrefixIter {
	obj.bgpOneIpv4NLRIPrefixSlice = append(obj.bgpOneIpv4NLRIPrefixSlice, item)
	return obj
}

// List of IPv6 prefixes being sent in the IPv6 Unicast MPUNREACH_NLRI .
// Ipv6Unicast returns a []BgpOneIpv6NLRIPrefix
func (obj *bgpAttributesMpUnreachNlri) Ipv6Unicast() BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	if len(obj.obj.Ipv6Unicast) == 0 {
		obj.setChoice(BgpAttributesMpUnreachNlriChoice.IPV6_UNICAST)
	}
	if obj.ipv6UnicastHolder == nil {
		obj.ipv6UnicastHolder = newBgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter(&obj.obj.Ipv6Unicast).setMsg(obj)
	}
	return obj.ipv6UnicastHolder
}

type bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter struct {
	obj                       *bgpAttributesMpUnreachNlri
	bgpOneIpv6NLRIPrefixSlice []BgpOneIpv6NLRIPrefix
	fieldPtr                  *[]*otg.BgpOneIpv6NLRIPrefix
}

func newBgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter(ptr *[]*otg.BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	return &bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter{fieldPtr: ptr}
}

type BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter interface {
	setMsg(*bgpAttributesMpUnreachNlri) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	Items() []BgpOneIpv6NLRIPrefix
	Add() BgpOneIpv6NLRIPrefix
	Append(items ...BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	Set(index int, newObj BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	Clear() BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	clearHolderSlice() BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
	appendHolderSlice(item BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) setMsg(msg *bgpAttributesMpUnreachNlri) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneIpv6NLRIPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) Items() []BgpOneIpv6NLRIPrefix {
	return obj.bgpOneIpv6NLRIPrefixSlice
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) Add() BgpOneIpv6NLRIPrefix {
	newObj := &otg.BgpOneIpv6NLRIPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneIpv6NLRIPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneIpv6NLRIPrefixSlice = append(obj.bgpOneIpv6NLRIPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) Append(items ...BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneIpv6NLRIPrefixSlice = append(obj.bgpOneIpv6NLRIPrefixSlice, item)
	}
	return obj
}

func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) Set(index int, newObj BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneIpv6NLRIPrefixSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) Clear() BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneIpv6NLRIPrefix{}
		obj.bgpOneIpv6NLRIPrefixSlice = []BgpOneIpv6NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) clearHolderSlice() BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	if len(obj.bgpOneIpv6NLRIPrefixSlice) > 0 {
		obj.bgpOneIpv6NLRIPrefixSlice = []BgpOneIpv6NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter) appendHolderSlice(item BgpOneIpv6NLRIPrefix) BgpAttributesMpUnreachNlriBgpOneIpv6NLRIPrefixIter {
	obj.bgpOneIpv6NLRIPrefixSlice = append(obj.bgpOneIpv6NLRIPrefixSlice, item)
	return obj
}

// IPv4 endpoint with Segment Routing Policy being sent in the IPv4 MPUNREACH_NLRI .
// Ipv4Srpolicy returns a BgpIpv4SrPolicyNLRIPrefix
func (obj *bgpAttributesMpUnreachNlri) Ipv4Srpolicy() BgpIpv4SrPolicyNLRIPrefix {
	if obj.obj.Ipv4Srpolicy == nil {
		obj.setChoice(BgpAttributesMpUnreachNlriChoice.IPV4_SRPOLICY)
	}
	if obj.ipv4SrpolicyHolder == nil {
		obj.ipv4SrpolicyHolder = &bgpIpv4SrPolicyNLRIPrefix{obj: obj.obj.Ipv4Srpolicy}
	}
	return obj.ipv4SrpolicyHolder
}

// IPv4 endpoint with Segment Routing Policy being sent in the IPv4 MPUNREACH_NLRI .
// Ipv4Srpolicy returns a BgpIpv4SrPolicyNLRIPrefix
func (obj *bgpAttributesMpUnreachNlri) HasIpv4Srpolicy() bool {
	return obj.obj.Ipv4Srpolicy != nil
}

// IPv4 endpoint with Segment Routing Policy being sent in the IPv4 MPUNREACH_NLRI .
// SetIpv4Srpolicy sets the BgpIpv4SrPolicyNLRIPrefix value in the BgpAttributesMpUnreachNlri object
func (obj *bgpAttributesMpUnreachNlri) SetIpv4Srpolicy(value BgpIpv4SrPolicyNLRIPrefix) BgpAttributesMpUnreachNlri {
	obj.setChoice(BgpAttributesMpUnreachNlriChoice.IPV4_SRPOLICY)
	obj.ipv4SrpolicyHolder = nil
	obj.obj.Ipv4Srpolicy = value.msg()

	return obj
}

// IPv6 endpoint with Segment Routing Policy being sent in the IPv4 MPUNREACH_NLRI .
// Ipv6Srpolicy returns a BgpIpv6SrPolicyNLRIPrefix
func (obj *bgpAttributesMpUnreachNlri) Ipv6Srpolicy() BgpIpv6SrPolicyNLRIPrefix {
	if obj.obj.Ipv6Srpolicy == nil {
		obj.setChoice(BgpAttributesMpUnreachNlriChoice.IPV6_SRPOLICY)
	}
	if obj.ipv6SrpolicyHolder == nil {
		obj.ipv6SrpolicyHolder = &bgpIpv6SrPolicyNLRIPrefix{obj: obj.obj.Ipv6Srpolicy}
	}
	return obj.ipv6SrpolicyHolder
}

// IPv6 endpoint with Segment Routing Policy being sent in the IPv4 MPUNREACH_NLRI .
// Ipv6Srpolicy returns a BgpIpv6SrPolicyNLRIPrefix
func (obj *bgpAttributesMpUnreachNlri) HasIpv6Srpolicy() bool {
	return obj.obj.Ipv6Srpolicy != nil
}

// IPv6 endpoint with Segment Routing Policy being sent in the IPv4 MPUNREACH_NLRI .
// SetIpv6Srpolicy sets the BgpIpv6SrPolicyNLRIPrefix value in the BgpAttributesMpUnreachNlri object
func (obj *bgpAttributesMpUnreachNlri) SetIpv6Srpolicy(value BgpIpv6SrPolicyNLRIPrefix) BgpAttributesMpUnreachNlri {
	obj.setChoice(BgpAttributesMpUnreachNlriChoice.IPV6_SRPOLICY)
	obj.ipv6SrpolicyHolder = nil
	obj.obj.Ipv6Srpolicy = value.msg()

	return obj
}

func (obj *bgpAttributesMpUnreachNlri) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ipv4Unicast) != 0 {

		if set_default {
			obj.Ipv4Unicast().clearHolderSlice()
			for _, item := range obj.obj.Ipv4Unicast {
				obj.Ipv4Unicast().appendHolderSlice(&bgpOneIpv4NLRIPrefix{obj: item})
			}
		}
		for _, item := range obj.Ipv4Unicast().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6Unicast) != 0 {

		if set_default {
			obj.Ipv6Unicast().clearHolderSlice()
			for _, item := range obj.obj.Ipv6Unicast {
				obj.Ipv6Unicast().appendHolderSlice(&bgpOneIpv6NLRIPrefix{obj: item})
			}
		}
		for _, item := range obj.Ipv6Unicast().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Ipv4Srpolicy != nil {

		obj.Ipv4Srpolicy().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6Srpolicy != nil {

		obj.Ipv6Srpolicy().validateObj(vObj, set_default)
	}

}

func (obj *bgpAttributesMpUnreachNlri) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesMpUnreachNlriChoiceEnum

	if len(obj.obj.Ipv4Unicast) > 0 {
		choices_set += 1
		choice = BgpAttributesMpUnreachNlriChoice.IPV4_UNICAST
	}

	if len(obj.obj.Ipv6Unicast) > 0 {
		choices_set += 1
		choice = BgpAttributesMpUnreachNlriChoice.IPV6_UNICAST
	}

	if obj.obj.Ipv4Srpolicy != nil {
		choices_set += 1
		choice = BgpAttributesMpUnreachNlriChoice.IPV4_SRPOLICY
	}

	if obj.obj.Ipv6Srpolicy != nil {
		choices_set += 1
		choice = BgpAttributesMpUnreachNlriChoice.IPV6_SRPOLICY
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesMpUnreachNlri")
			}
		} else {
			intVal := otg.BgpAttributesMpUnreachNlri_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesMpUnreachNlri_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
