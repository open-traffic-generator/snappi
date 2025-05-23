package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesMpReachNlri *****
type bgpAttributesMpReachNlri struct {
	validation
	obj                *otg.BgpAttributesMpReachNlri
	marshaller         marshalBgpAttributesMpReachNlri
	unMarshaller       unMarshalBgpAttributesMpReachNlri
	nextHopHolder      BgpAttributesNextHop
	ipv4UnicastHolder  BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	ipv6UnicastHolder  BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	ipv4SrpolicyHolder BgpIpv4SrPolicyNLRIPrefix
	ipv6SrpolicyHolder BgpIpv6SrPolicyNLRIPrefix
}

func NewBgpAttributesMpReachNlri() BgpAttributesMpReachNlri {
	obj := bgpAttributesMpReachNlri{obj: &otg.BgpAttributesMpReachNlri{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesMpReachNlri) msg() *otg.BgpAttributesMpReachNlri {
	return obj.obj
}

func (obj *bgpAttributesMpReachNlri) setMsg(msg *otg.BgpAttributesMpReachNlri) BgpAttributesMpReachNlri {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesMpReachNlri struct {
	obj *bgpAttributesMpReachNlri
}

type marshalBgpAttributesMpReachNlri interface {
	// ToProto marshals BgpAttributesMpReachNlri to protobuf object *otg.BgpAttributesMpReachNlri
	ToProto() (*otg.BgpAttributesMpReachNlri, error)
	// ToPbText marshals BgpAttributesMpReachNlri to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesMpReachNlri to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesMpReachNlri to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesMpReachNlri to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesMpReachNlri struct {
	obj *bgpAttributesMpReachNlri
}

type unMarshalBgpAttributesMpReachNlri interface {
	// FromProto unmarshals BgpAttributesMpReachNlri from protobuf object *otg.BgpAttributesMpReachNlri
	FromProto(msg *otg.BgpAttributesMpReachNlri) (BgpAttributesMpReachNlri, error)
	// FromPbText unmarshals BgpAttributesMpReachNlri from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesMpReachNlri from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesMpReachNlri from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesMpReachNlri) Marshal() marshalBgpAttributesMpReachNlri {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesMpReachNlri{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesMpReachNlri) Unmarshal() unMarshalBgpAttributesMpReachNlri {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesMpReachNlri{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesMpReachNlri) ToProto() (*otg.BgpAttributesMpReachNlri, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesMpReachNlri) FromProto(msg *otg.BgpAttributesMpReachNlri) (BgpAttributesMpReachNlri, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesMpReachNlri) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesMpReachNlri) FromPbText(value string) error {
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

func (m *marshalbgpAttributesMpReachNlri) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesMpReachNlri) FromYaml(value string) error {
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

func (m *marshalbgpAttributesMpReachNlri) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesMpReachNlri) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesMpReachNlri) FromJson(value string) error {
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

func (obj *bgpAttributesMpReachNlri) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesMpReachNlri) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesMpReachNlri) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesMpReachNlri) Clone() (BgpAttributesMpReachNlri, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesMpReachNlri()
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

func (obj *bgpAttributesMpReachNlri) setNil() {
	obj.nextHopHolder = nil
	obj.ipv4UnicastHolder = nil
	obj.ipv6UnicastHolder = nil
	obj.ipv4SrpolicyHolder = nil
	obj.ipv6SrpolicyHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesMpReachNlri is the MP_REACH attribute is an optional attribute which can be included in the attributes of a BGP Update message as defined in https://datatracker.ietf.org/doc/html/rfc4760#section-3.
// The following AFI / SAFI combinations are supported:
// - IPv4 Unicast with AFI as 1 and SAFI as 1
// - IPv6 Unicast with AFI as 2 and SAFI as 1
// - Segment Routing Policy for IPv4 Unicast with AFI as 1 and SAFI as 73 ( draft-ietf-idr-sr-policy-safi-02 Section 2.1 )
// - Segment Routing Policy for IPv6 Unicast with AFI as 2 and SAFI as 73
type BgpAttributesMpReachNlri interface {
	Validation
	// msg marshals BgpAttributesMpReachNlri to protobuf object *otg.BgpAttributesMpReachNlri
	// and doesn't set defaults
	msg() *otg.BgpAttributesMpReachNlri
	// setMsg unmarshals BgpAttributesMpReachNlri from protobuf object *otg.BgpAttributesMpReachNlri
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesMpReachNlri) BgpAttributesMpReachNlri
	// provides marshal interface
	Marshal() marshalBgpAttributesMpReachNlri
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesMpReachNlri
	// validate validates BgpAttributesMpReachNlri
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesMpReachNlri, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NextHop returns BgpAttributesNextHop, set in BgpAttributesMpReachNlri.
	// BgpAttributesNextHop is next hop to be sent inside MP_REACH NLRI or as the NEXT_HOP attribute if advertised as traditional NLRI.
	NextHop() BgpAttributesNextHop
	// SetNextHop assigns BgpAttributesNextHop provided by user to BgpAttributesMpReachNlri.
	// BgpAttributesNextHop is next hop to be sent inside MP_REACH NLRI or as the NEXT_HOP attribute if advertised as traditional NLRI.
	SetNextHop(value BgpAttributesNextHop) BgpAttributesMpReachNlri
	// HasNextHop checks if NextHop has been set in BgpAttributesMpReachNlri
	HasNextHop() bool
	// Choice returns BgpAttributesMpReachNlriChoiceEnum, set in BgpAttributesMpReachNlri
	Choice() BgpAttributesMpReachNlriChoiceEnum
	// setChoice assigns BgpAttributesMpReachNlriChoiceEnum provided by user to BgpAttributesMpReachNlri
	setChoice(value BgpAttributesMpReachNlriChoiceEnum) BgpAttributesMpReachNlri
	// Ipv4Unicast returns BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIterIter, set in BgpAttributesMpReachNlri
	Ipv4Unicast() BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	// Ipv6Unicast returns BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIterIter, set in BgpAttributesMpReachNlri
	Ipv6Unicast() BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	// Ipv4Srpolicy returns BgpIpv4SrPolicyNLRIPrefix, set in BgpAttributesMpReachNlri.
	// BgpIpv4SrPolicyNLRIPrefix is iPv4 Segment Routing Policy NLRI Prefix.
	Ipv4Srpolicy() BgpIpv4SrPolicyNLRIPrefix
	// SetIpv4Srpolicy assigns BgpIpv4SrPolicyNLRIPrefix provided by user to BgpAttributesMpReachNlri.
	// BgpIpv4SrPolicyNLRIPrefix is iPv4 Segment Routing Policy NLRI Prefix.
	SetIpv4Srpolicy(value BgpIpv4SrPolicyNLRIPrefix) BgpAttributesMpReachNlri
	// HasIpv4Srpolicy checks if Ipv4Srpolicy has been set in BgpAttributesMpReachNlri
	HasIpv4Srpolicy() bool
	// Ipv6Srpolicy returns BgpIpv6SrPolicyNLRIPrefix, set in BgpAttributesMpReachNlri.
	// BgpIpv6SrPolicyNLRIPrefix is one IPv6 Segment Routing Policy NLRI Prefix.
	Ipv6Srpolicy() BgpIpv6SrPolicyNLRIPrefix
	// SetIpv6Srpolicy assigns BgpIpv6SrPolicyNLRIPrefix provided by user to BgpAttributesMpReachNlri.
	// BgpIpv6SrPolicyNLRIPrefix is one IPv6 Segment Routing Policy NLRI Prefix.
	SetIpv6Srpolicy(value BgpIpv6SrPolicyNLRIPrefix) BgpAttributesMpReachNlri
	// HasIpv6Srpolicy checks if Ipv6Srpolicy has been set in BgpAttributesMpReachNlri
	HasIpv6Srpolicy() bool
	setNil()
}

// description is TBD
// NextHop returns a BgpAttributesNextHop
func (obj *bgpAttributesMpReachNlri) NextHop() BgpAttributesNextHop {
	if obj.obj.NextHop == nil {
		obj.obj.NextHop = NewBgpAttributesNextHop().msg()
	}
	if obj.nextHopHolder == nil {
		obj.nextHopHolder = &bgpAttributesNextHop{obj: obj.obj.NextHop}
	}
	return obj.nextHopHolder
}

// description is TBD
// NextHop returns a BgpAttributesNextHop
func (obj *bgpAttributesMpReachNlri) HasNextHop() bool {
	return obj.obj.NextHop != nil
}

// description is TBD
// SetNextHop sets the BgpAttributesNextHop value in the BgpAttributesMpReachNlri object
func (obj *bgpAttributesMpReachNlri) SetNextHop(value BgpAttributesNextHop) BgpAttributesMpReachNlri {

	obj.nextHopHolder = nil
	obj.obj.NextHop = value.msg()

	return obj
}

type BgpAttributesMpReachNlriChoiceEnum string

// Enum of Choice on BgpAttributesMpReachNlri
var BgpAttributesMpReachNlriChoice = struct {
	IPV4_UNICAST  BgpAttributesMpReachNlriChoiceEnum
	IPV6_UNICAST  BgpAttributesMpReachNlriChoiceEnum
	IPV4_SRPOLICY BgpAttributesMpReachNlriChoiceEnum
	IPV6_SRPOLICY BgpAttributesMpReachNlriChoiceEnum
}{
	IPV4_UNICAST:  BgpAttributesMpReachNlriChoiceEnum("ipv4_unicast"),
	IPV6_UNICAST:  BgpAttributesMpReachNlriChoiceEnum("ipv6_unicast"),
	IPV4_SRPOLICY: BgpAttributesMpReachNlriChoiceEnum("ipv4_srpolicy"),
	IPV6_SRPOLICY: BgpAttributesMpReachNlriChoiceEnum("ipv6_srpolicy"),
}

func (obj *bgpAttributesMpReachNlri) Choice() BgpAttributesMpReachNlriChoiceEnum {
	return BgpAttributesMpReachNlriChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *bgpAttributesMpReachNlri) setChoice(value BgpAttributesMpReachNlriChoiceEnum) BgpAttributesMpReachNlri {
	intValue, ok := otg.BgpAttributesMpReachNlri_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAttributesMpReachNlriChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAttributesMpReachNlri_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6Srpolicy = nil
	obj.ipv6SrpolicyHolder = nil
	obj.obj.Ipv4Srpolicy = nil
	obj.ipv4SrpolicyHolder = nil
	obj.obj.Ipv6Unicast = nil
	obj.ipv6UnicastHolder = nil
	obj.obj.Ipv4Unicast = nil
	obj.ipv4UnicastHolder = nil

	if value == BgpAttributesMpReachNlriChoice.IPV4_UNICAST {
		obj.obj.Ipv4Unicast = []*otg.BgpOneIpv4NLRIPrefix{}
	}

	if value == BgpAttributesMpReachNlriChoice.IPV6_UNICAST {
		obj.obj.Ipv6Unicast = []*otg.BgpOneIpv6NLRIPrefix{}
	}

	if value == BgpAttributesMpReachNlriChoice.IPV4_SRPOLICY {
		obj.obj.Ipv4Srpolicy = NewBgpIpv4SrPolicyNLRIPrefix().msg()
	}

	if value == BgpAttributesMpReachNlriChoice.IPV6_SRPOLICY {
		obj.obj.Ipv6Srpolicy = NewBgpIpv6SrPolicyNLRIPrefix().msg()
	}

	return obj
}

// List of IPv4 prefixes being sent in the IPv4 Unicast MPREACH_NLRI .
// Ipv4Unicast returns a []BgpOneIpv4NLRIPrefix
func (obj *bgpAttributesMpReachNlri) Ipv4Unicast() BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	if len(obj.obj.Ipv4Unicast) == 0 {
		obj.setChoice(BgpAttributesMpReachNlriChoice.IPV4_UNICAST)
	}
	if obj.ipv4UnicastHolder == nil {
		obj.ipv4UnicastHolder = newBgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter(&obj.obj.Ipv4Unicast).setMsg(obj)
	}
	return obj.ipv4UnicastHolder
}

type bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter struct {
	obj                       *bgpAttributesMpReachNlri
	bgpOneIpv4NLRIPrefixSlice []BgpOneIpv4NLRIPrefix
	fieldPtr                  *[]*otg.BgpOneIpv4NLRIPrefix
}

func newBgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter(ptr *[]*otg.BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	return &bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter{fieldPtr: ptr}
}

type BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter interface {
	setMsg(*bgpAttributesMpReachNlri) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	Items() []BgpOneIpv4NLRIPrefix
	Add() BgpOneIpv4NLRIPrefix
	Append(items ...BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	Set(index int, newObj BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	Clear() BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	clearHolderSlice() BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
	appendHolderSlice(item BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) setMsg(msg *bgpAttributesMpReachNlri) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneIpv4NLRIPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) Items() []BgpOneIpv4NLRIPrefix {
	return obj.bgpOneIpv4NLRIPrefixSlice
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) Add() BgpOneIpv4NLRIPrefix {
	newObj := &otg.BgpOneIpv4NLRIPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneIpv4NLRIPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneIpv4NLRIPrefixSlice = append(obj.bgpOneIpv4NLRIPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) Append(items ...BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneIpv4NLRIPrefixSlice = append(obj.bgpOneIpv4NLRIPrefixSlice, item)
	}
	return obj
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) Set(index int, newObj BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneIpv4NLRIPrefixSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) Clear() BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneIpv4NLRIPrefix{}
		obj.bgpOneIpv4NLRIPrefixSlice = []BgpOneIpv4NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) clearHolderSlice() BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	if len(obj.bgpOneIpv4NLRIPrefixSlice) > 0 {
		obj.bgpOneIpv4NLRIPrefixSlice = []BgpOneIpv4NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter) appendHolderSlice(item BgpOneIpv4NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv4NLRIPrefixIter {
	obj.bgpOneIpv4NLRIPrefixSlice = append(obj.bgpOneIpv4NLRIPrefixSlice, item)
	return obj
}

// List of IPv6 prefixes being sent in the IPv6 Unicast MPREACH_NLRI .
// Ipv6Unicast returns a []BgpOneIpv6NLRIPrefix
func (obj *bgpAttributesMpReachNlri) Ipv6Unicast() BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	if len(obj.obj.Ipv6Unicast) == 0 {
		obj.setChoice(BgpAttributesMpReachNlriChoice.IPV6_UNICAST)
	}
	if obj.ipv6UnicastHolder == nil {
		obj.ipv6UnicastHolder = newBgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter(&obj.obj.Ipv6Unicast).setMsg(obj)
	}
	return obj.ipv6UnicastHolder
}

type bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter struct {
	obj                       *bgpAttributesMpReachNlri
	bgpOneIpv6NLRIPrefixSlice []BgpOneIpv6NLRIPrefix
	fieldPtr                  *[]*otg.BgpOneIpv6NLRIPrefix
}

func newBgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter(ptr *[]*otg.BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	return &bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter{fieldPtr: ptr}
}

type BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter interface {
	setMsg(*bgpAttributesMpReachNlri) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	Items() []BgpOneIpv6NLRIPrefix
	Add() BgpOneIpv6NLRIPrefix
	Append(items ...BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	Set(index int, newObj BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	Clear() BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	clearHolderSlice() BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
	appendHolderSlice(item BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) setMsg(msg *bgpAttributesMpReachNlri) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneIpv6NLRIPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) Items() []BgpOneIpv6NLRIPrefix {
	return obj.bgpOneIpv6NLRIPrefixSlice
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) Add() BgpOneIpv6NLRIPrefix {
	newObj := &otg.BgpOneIpv6NLRIPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneIpv6NLRIPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneIpv6NLRIPrefixSlice = append(obj.bgpOneIpv6NLRIPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) Append(items ...BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneIpv6NLRIPrefixSlice = append(obj.bgpOneIpv6NLRIPrefixSlice, item)
	}
	return obj
}

func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) Set(index int, newObj BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneIpv6NLRIPrefixSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) Clear() BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneIpv6NLRIPrefix{}
		obj.bgpOneIpv6NLRIPrefixSlice = []BgpOneIpv6NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) clearHolderSlice() BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	if len(obj.bgpOneIpv6NLRIPrefixSlice) > 0 {
		obj.bgpOneIpv6NLRIPrefixSlice = []BgpOneIpv6NLRIPrefix{}
	}
	return obj
}
func (obj *bgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter) appendHolderSlice(item BgpOneIpv6NLRIPrefix) BgpAttributesMpReachNlriBgpOneIpv6NLRIPrefixIter {
	obj.bgpOneIpv6NLRIPrefixSlice = append(obj.bgpOneIpv6NLRIPrefixSlice, item)
	return obj
}

// IPv4 endpoint with Segment Routing Policy being sent in the IPv4 MPREACH_NLRI .
// Ipv4Srpolicy returns a BgpIpv4SrPolicyNLRIPrefix
func (obj *bgpAttributesMpReachNlri) Ipv4Srpolicy() BgpIpv4SrPolicyNLRIPrefix {
	if obj.obj.Ipv4Srpolicy == nil {
		obj.setChoice(BgpAttributesMpReachNlriChoice.IPV4_SRPOLICY)
	}
	if obj.ipv4SrpolicyHolder == nil {
		obj.ipv4SrpolicyHolder = &bgpIpv4SrPolicyNLRIPrefix{obj: obj.obj.Ipv4Srpolicy}
	}
	return obj.ipv4SrpolicyHolder
}

// IPv4 endpoint with Segment Routing Policy being sent in the IPv4 MPREACH_NLRI .
// Ipv4Srpolicy returns a BgpIpv4SrPolicyNLRIPrefix
func (obj *bgpAttributesMpReachNlri) HasIpv4Srpolicy() bool {
	return obj.obj.Ipv4Srpolicy != nil
}

// IPv4 endpoint with Segment Routing Policy being sent in the IPv4 MPREACH_NLRI .
// SetIpv4Srpolicy sets the BgpIpv4SrPolicyNLRIPrefix value in the BgpAttributesMpReachNlri object
func (obj *bgpAttributesMpReachNlri) SetIpv4Srpolicy(value BgpIpv4SrPolicyNLRIPrefix) BgpAttributesMpReachNlri {
	obj.setChoice(BgpAttributesMpReachNlriChoice.IPV4_SRPOLICY)
	obj.ipv4SrpolicyHolder = nil
	obj.obj.Ipv4Srpolicy = value.msg()

	return obj
}

// IPv6 endpoint with Segment Routing Policy being sent in the IPv6 MPREACH_NLRI .
// Ipv6Srpolicy returns a BgpIpv6SrPolicyNLRIPrefix
func (obj *bgpAttributesMpReachNlri) Ipv6Srpolicy() BgpIpv6SrPolicyNLRIPrefix {
	if obj.obj.Ipv6Srpolicy == nil {
		obj.setChoice(BgpAttributesMpReachNlriChoice.IPV6_SRPOLICY)
	}
	if obj.ipv6SrpolicyHolder == nil {
		obj.ipv6SrpolicyHolder = &bgpIpv6SrPolicyNLRIPrefix{obj: obj.obj.Ipv6Srpolicy}
	}
	return obj.ipv6SrpolicyHolder
}

// IPv6 endpoint with Segment Routing Policy being sent in the IPv6 MPREACH_NLRI .
// Ipv6Srpolicy returns a BgpIpv6SrPolicyNLRIPrefix
func (obj *bgpAttributesMpReachNlri) HasIpv6Srpolicy() bool {
	return obj.obj.Ipv6Srpolicy != nil
}

// IPv6 endpoint with Segment Routing Policy being sent in the IPv6 MPREACH_NLRI .
// SetIpv6Srpolicy sets the BgpIpv6SrPolicyNLRIPrefix value in the BgpAttributesMpReachNlri object
func (obj *bgpAttributesMpReachNlri) SetIpv6Srpolicy(value BgpIpv6SrPolicyNLRIPrefix) BgpAttributesMpReachNlri {
	obj.setChoice(BgpAttributesMpReachNlriChoice.IPV6_SRPOLICY)
	obj.ipv6SrpolicyHolder = nil
	obj.obj.Ipv6Srpolicy = value.msg()

	return obj
}

func (obj *bgpAttributesMpReachNlri) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NextHop != nil {

		obj.NextHop().validateObj(vObj, set_default)
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface BgpAttributesMpReachNlri")
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

func (obj *bgpAttributesMpReachNlri) setDefault() {
	var choices_set int = 0
	var choice BgpAttributesMpReachNlriChoiceEnum

	if len(obj.obj.Ipv4Unicast) > 0 {
		choices_set += 1
		choice = BgpAttributesMpReachNlriChoice.IPV4_UNICAST
	}

	if len(obj.obj.Ipv6Unicast) > 0 {
		choices_set += 1
		choice = BgpAttributesMpReachNlriChoice.IPV6_UNICAST
	}

	if obj.obj.Ipv4Srpolicy != nil {
		choices_set += 1
		choice = BgpAttributesMpReachNlriChoice.IPV4_SRPOLICY
	}

	if obj.obj.Ipv6Srpolicy != nil {
		choices_set += 1
		choice = BgpAttributesMpReachNlriChoice.IPV6_SRPOLICY
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpAttributesMpReachNlri")
			}
		} else {
			intVal := otg.BgpAttributesMpReachNlri_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpAttributesMpReachNlri_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
