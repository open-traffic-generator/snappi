package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSegmentRouting *****
type isisSegmentRouting struct {
	validation
	obj                    *otg.IsisSegmentRouting
	marshaller             marshalIsisSegmentRouting
	unMarshaller           unMarshalIsisSegmentRouting
	routerCapabilityHolder IsisRouterCapability
	srv6LocatorsHolder     IsisSegmentRoutingIsisSRv6LocatorIter
}

func NewIsisSegmentRouting() IsisSegmentRouting {
	obj := isisSegmentRouting{obj: &otg.IsisSegmentRouting{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSegmentRouting) msg() *otg.IsisSegmentRouting {
	return obj.obj
}

func (obj *isisSegmentRouting) setMsg(msg *otg.IsisSegmentRouting) IsisSegmentRouting {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSegmentRouting struct {
	obj *isisSegmentRouting
}

type marshalIsisSegmentRouting interface {
	// ToProto marshals IsisSegmentRouting to protobuf object *otg.IsisSegmentRouting
	ToProto() (*otg.IsisSegmentRouting, error)
	// ToPbText marshals IsisSegmentRouting to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSegmentRouting to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSegmentRouting to JSON text
	ToJson() (string, error)
}

type unMarshalisisSegmentRouting struct {
	obj *isisSegmentRouting
}

type unMarshalIsisSegmentRouting interface {
	// FromProto unmarshals IsisSegmentRouting from protobuf object *otg.IsisSegmentRouting
	FromProto(msg *otg.IsisSegmentRouting) (IsisSegmentRouting, error)
	// FromPbText unmarshals IsisSegmentRouting from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSegmentRouting from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSegmentRouting from JSON text
	FromJson(value string) error
}

func (obj *isisSegmentRouting) Marshal() marshalIsisSegmentRouting {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSegmentRouting{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSegmentRouting) Unmarshal() unMarshalIsisSegmentRouting {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSegmentRouting{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSegmentRouting) ToProto() (*otg.IsisSegmentRouting, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSegmentRouting) FromProto(msg *otg.IsisSegmentRouting) (IsisSegmentRouting, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSegmentRouting) ToPbText() (string, error) {
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

func (m *unMarshalisisSegmentRouting) FromPbText(value string) error {
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

func (m *marshalisisSegmentRouting) ToYaml() (string, error) {
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

func (m *unMarshalisisSegmentRouting) FromYaml(value string) error {
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

func (m *marshalisisSegmentRouting) ToJson() (string, error) {
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

func (m *unMarshalisisSegmentRouting) FromJson(value string) error {
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

func (obj *isisSegmentRouting) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSegmentRouting) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSegmentRouting) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSegmentRouting) Clone() (IsisSegmentRouting, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSegmentRouting()
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

func (obj *isisSegmentRouting) setNil() {
	obj.routerCapabilityHolder = nil
	obj.srv6LocatorsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSegmentRouting is segment Routing (SR) allows for a flexible definition of end-to-end paths within IGP topologies by encoding paths as sequences of topological sub-paths,
// called "segments". These segments are advertised by the link-state routing protocols (IS-IS and OSPF).
// Prefix segments represent an ECMP-aware shortest path to a prefix (or a node), as per the state of the IGP topology.
// Adjacency segments represent a hop over a specific adjacency between two nodes in the IGP.
// A prefix segment is typically a multi-hop path while an adjacency segment, in most of the cases, is a one-hop path.
// These segments act as topological sub-paths that can be combined together to form the required path.
// Reference: https://datatracker.ietf.org/doc/html/rfc8667.
// An implementation may advertise Router Capability with default values if a user does not even set the properties
// of Router Capability and Segment Routing Capability.
type IsisSegmentRouting interface {
	Validation
	// msg marshals IsisSegmentRouting to protobuf object *otg.IsisSegmentRouting
	// and doesn't set defaults
	msg() *otg.IsisSegmentRouting
	// setMsg unmarshals IsisSegmentRouting from protobuf object *otg.IsisSegmentRouting
	// and doesn't set defaults
	setMsg(*otg.IsisSegmentRouting) IsisSegmentRouting
	// provides marshal interface
	Marshal() marshalIsisSegmentRouting
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSegmentRouting
	// validate validates IsisSegmentRouting
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSegmentRouting, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterCapability returns IsisRouterCapability, set in IsisSegmentRouting.
	// IsisRouterCapability is container for the configuration of IS-IS Router CAPABILITY TLV.
	// https://datatracker.ietf.org/doc/html/rfc7981#section-2.
	// An implementation should set default values appropriately if any mandatory item is not configured by a user.
	RouterCapability() IsisRouterCapability
	// SetRouterCapability assigns IsisRouterCapability provided by user to IsisSegmentRouting.
	// IsisRouterCapability is container for the configuration of IS-IS Router CAPABILITY TLV.
	// https://datatracker.ietf.org/doc/html/rfc7981#section-2.
	// An implementation should set default values appropriately if any mandatory item is not configured by a user.
	SetRouterCapability(value IsisRouterCapability) IsisSegmentRouting
	// HasRouterCapability checks if RouterCapability has been set in IsisSegmentRouting
	HasRouterCapability() bool
	// Srv6Locators returns IsisSegmentRoutingIsisSRv6LocatorIterIter, set in IsisSegmentRouting
	Srv6Locators() IsisSegmentRoutingIsisSRv6LocatorIter
	setNil()
}

// Optional IS-IS TLV named CAPABILITY, formed of multiple sub-TLVs, which allows a router to announce its
// capabilities within an IS-IS level or the entire routing domain.
// RouterCapability returns a IsisRouterCapability
func (obj *isisSegmentRouting) RouterCapability() IsisRouterCapability {
	if obj.obj.RouterCapability == nil {
		obj.obj.RouterCapability = NewIsisRouterCapability().msg()
	}
	if obj.routerCapabilityHolder == nil {
		obj.routerCapabilityHolder = &isisRouterCapability{obj: obj.obj.RouterCapability}
	}
	return obj.routerCapabilityHolder
}

// Optional IS-IS TLV named CAPABILITY, formed of multiple sub-TLVs, which allows a router to announce its
// capabilities within an IS-IS level or the entire routing domain.
// RouterCapability returns a IsisRouterCapability
func (obj *isisSegmentRouting) HasRouterCapability() bool {
	return obj.obj.RouterCapability != nil
}

// Optional IS-IS TLV named CAPABILITY, formed of multiple sub-TLVs, which allows a router to announce its
// capabilities within an IS-IS level or the entire routing domain.
// SetRouterCapability sets the IsisRouterCapability value in the IsisSegmentRouting object
func (obj *isisSegmentRouting) SetRouterCapability(value IsisRouterCapability) IsisSegmentRouting {

	obj.routerCapabilityHolder = nil
	obj.obj.RouterCapability = value.msg()

	return obj
}

// List of SRv6 Locator TLVs (TLV type 27) to be advertised by this IS-IS router. Each locator binds an IPv6 prefix to an IGP algorithm (standard SPF or Flex-Algo per RFC 9350) and carries End SID Sub-TLVs for locally instantiated node-level SRv6 SIDs. One Locator TLV is required per topology/algorithm combination. Reference: RFC 9352 Section 7.1.
// Srv6Locators returns a []IsisSRv6Locator
func (obj *isisSegmentRouting) Srv6Locators() IsisSegmentRoutingIsisSRv6LocatorIter {
	if len(obj.obj.Srv6Locators) == 0 {
		obj.obj.Srv6Locators = []*otg.IsisSRv6Locator{}
	}
	if obj.srv6LocatorsHolder == nil {
		obj.srv6LocatorsHolder = newIsisSegmentRoutingIsisSRv6LocatorIter(&obj.obj.Srv6Locators).setMsg(obj)
	}
	return obj.srv6LocatorsHolder
}

type isisSegmentRoutingIsisSRv6LocatorIter struct {
	obj                  *isisSegmentRouting
	isisSRv6LocatorSlice []IsisSRv6Locator
	fieldPtr             *[]*otg.IsisSRv6Locator
}

func newIsisSegmentRoutingIsisSRv6LocatorIter(ptr *[]*otg.IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter {
	return &isisSegmentRoutingIsisSRv6LocatorIter{fieldPtr: ptr}
}

type IsisSegmentRoutingIsisSRv6LocatorIter interface {
	setMsg(*isisSegmentRouting) IsisSegmentRoutingIsisSRv6LocatorIter
	Items() []IsisSRv6Locator
	Add() IsisSRv6Locator
	Append(items ...IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter
	Set(index int, newObj IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter
	Clear() IsisSegmentRoutingIsisSRv6LocatorIter
	clearHolderSlice() IsisSegmentRoutingIsisSRv6LocatorIter
	appendHolderSlice(item IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter
}

func (obj *isisSegmentRoutingIsisSRv6LocatorIter) setMsg(msg *isisSegmentRouting) IsisSegmentRoutingIsisSRv6LocatorIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRv6Locator{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisSegmentRoutingIsisSRv6LocatorIter) Items() []IsisSRv6Locator {
	return obj.isisSRv6LocatorSlice
}

func (obj *isisSegmentRoutingIsisSRv6LocatorIter) Add() IsisSRv6Locator {
	newObj := &otg.IsisSRv6Locator{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRv6Locator{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRv6LocatorSlice = append(obj.isisSRv6LocatorSlice, newLibObj)
	return newLibObj
}

func (obj *isisSegmentRoutingIsisSRv6LocatorIter) Append(items ...IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRv6LocatorSlice = append(obj.isisSRv6LocatorSlice, item)
	}
	return obj
}

func (obj *isisSegmentRoutingIsisSRv6LocatorIter) Set(index int, newObj IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRv6LocatorSlice[index] = newObj
	return obj
}
func (obj *isisSegmentRoutingIsisSRv6LocatorIter) Clear() IsisSegmentRoutingIsisSRv6LocatorIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRv6Locator{}
		obj.isisSRv6LocatorSlice = []IsisSRv6Locator{}
	}
	return obj
}
func (obj *isisSegmentRoutingIsisSRv6LocatorIter) clearHolderSlice() IsisSegmentRoutingIsisSRv6LocatorIter {
	if len(obj.isisSRv6LocatorSlice) > 0 {
		obj.isisSRv6LocatorSlice = []IsisSRv6Locator{}
	}
	return obj
}
func (obj *isisSegmentRoutingIsisSRv6LocatorIter) appendHolderSlice(item IsisSRv6Locator) IsisSegmentRoutingIsisSRv6LocatorIter {
	obj.isisSRv6LocatorSlice = append(obj.isisSRv6LocatorSlice, item)
	return obj
}

func (obj *isisSegmentRouting) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterCapability != nil {

		obj.RouterCapability().validateObj(vObj, set_default)
	}

	if len(obj.obj.Srv6Locators) != 0 {

		if set_default {
			obj.Srv6Locators().clearHolderSlice()
			for _, item := range obj.obj.Srv6Locators {
				obj.Srv6Locators().appendHolderSlice(&isisSRv6Locator{obj: item})
			}
		}
		for _, item := range obj.Srv6Locators().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisSegmentRouting) setDefault() {

}
