package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspIpv4Interface *****
type rsvpLspIpv4Interface struct {
	validation
	obj                      *otg.RsvpLspIpv4Interface
	marshaller               marshalRsvpLspIpv4Interface
	unMarshaller             unMarshalRsvpLspIpv4Interface
	p2PEgressIpv4LspsHolder  RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	p2PIngressIpv4LspsHolder RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
}

func NewRsvpLspIpv4Interface() RsvpLspIpv4Interface {
	obj := rsvpLspIpv4Interface{obj: &otg.RsvpLspIpv4Interface{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspIpv4Interface) msg() *otg.RsvpLspIpv4Interface {
	return obj.obj
}

func (obj *rsvpLspIpv4Interface) setMsg(msg *otg.RsvpLspIpv4Interface) RsvpLspIpv4Interface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspIpv4Interface struct {
	obj *rsvpLspIpv4Interface
}

type marshalRsvpLspIpv4Interface interface {
	// ToProto marshals RsvpLspIpv4Interface to protobuf object *otg.RsvpLspIpv4Interface
	ToProto() (*otg.RsvpLspIpv4Interface, error)
	// ToPbText marshals RsvpLspIpv4Interface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspIpv4Interface to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspIpv4Interface to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpLspIpv4Interface struct {
	obj *rsvpLspIpv4Interface
}

type unMarshalRsvpLspIpv4Interface interface {
	// FromProto unmarshals RsvpLspIpv4Interface from protobuf object *otg.RsvpLspIpv4Interface
	FromProto(msg *otg.RsvpLspIpv4Interface) (RsvpLspIpv4Interface, error)
	// FromPbText unmarshals RsvpLspIpv4Interface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspIpv4Interface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspIpv4Interface from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspIpv4Interface) Marshal() marshalRsvpLspIpv4Interface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspIpv4Interface{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspIpv4Interface) Unmarshal() unMarshalRsvpLspIpv4Interface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspIpv4Interface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspIpv4Interface) ToProto() (*otg.RsvpLspIpv4Interface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspIpv4Interface) FromProto(msg *otg.RsvpLspIpv4Interface) (RsvpLspIpv4Interface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspIpv4Interface) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Interface) FromPbText(value string) error {
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

func (m *marshalrsvpLspIpv4Interface) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Interface) FromYaml(value string) error {
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

func (m *marshalrsvpLspIpv4Interface) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspIpv4Interface) FromJson(value string) error {
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

func (obj *rsvpLspIpv4Interface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4Interface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspIpv4Interface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspIpv4Interface) Clone() (RsvpLspIpv4Interface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspIpv4Interface()
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

func (obj *rsvpLspIpv4Interface) setNil() {
	obj.p2PEgressIpv4LspsHolder = nil
	obj.p2PIngressIpv4LspsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RsvpLspIpv4Interface is configuration for RSVP LSP IPv4 Interface.
type RsvpLspIpv4Interface interface {
	Validation
	// msg marshals RsvpLspIpv4Interface to protobuf object *otg.RsvpLspIpv4Interface
	// and doesn't set defaults
	msg() *otg.RsvpLspIpv4Interface
	// setMsg unmarshals RsvpLspIpv4Interface from protobuf object *otg.RsvpLspIpv4Interface
	// and doesn't set defaults
	setMsg(*otg.RsvpLspIpv4Interface) RsvpLspIpv4Interface
	// provides marshal interface
	Marshal() marshalRsvpLspIpv4Interface
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspIpv4Interface
	// validate validates RsvpLspIpv4Interface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspIpv4Interface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Name returns string, set in RsvpLspIpv4Interface.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to RsvpLspIpv4Interface
	SetIpv4Name(value string) RsvpLspIpv4Interface
	// P2PEgressIpv4Lsps returns RsvpLspIpv4InterfaceP2PEgressIpv4Lsp, set in RsvpLspIpv4Interface.
	// RsvpLspIpv4InterfaceP2PEgressIpv4Lsp is configuration for RSVP Egress Point-to-Point(P2P) IPv4 LSPs.
	P2PEgressIpv4Lsps() RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
	// SetP2PEgressIpv4Lsps assigns RsvpLspIpv4InterfaceP2PEgressIpv4Lsp provided by user to RsvpLspIpv4Interface.
	// RsvpLspIpv4InterfaceP2PEgressIpv4Lsp is configuration for RSVP Egress Point-to-Point(P2P) IPv4 LSPs.
	SetP2PEgressIpv4Lsps(value RsvpLspIpv4InterfaceP2PEgressIpv4Lsp) RsvpLspIpv4Interface
	// HasP2PEgressIpv4Lsps checks if P2PEgressIpv4Lsps has been set in RsvpLspIpv4Interface
	HasP2PEgressIpv4Lsps() bool
	// P2PIngressIpv4Lsps returns RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIterIter, set in RsvpLspIpv4Interface
	P2PIngressIpv4Lsps() RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
	setNil()
}

// The globally unique name of the IPv4 or Loopback IPv4 interface acting as the RSVP ingress and egress endpoint for  the LSPs configured on this interface. This must match the "name" field of either "ipv4_addresses" or "ipv4_loopbacks"  on which this LSP interface is configured.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// Ipv4Name returns a string
func (obj *rsvpLspIpv4Interface) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The globally unique name of the IPv4 or Loopback IPv4 interface acting as the RSVP ingress and egress endpoint for  the LSPs configured on this interface. This must match the "name" field of either "ipv4_addresses" or "ipv4_loopbacks"  on which this LSP interface is configured.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
// - /components/schemas/Device.Ipv4Loopback/properties/name
//
// SetIpv4Name sets the string value in the RsvpLspIpv4Interface object
func (obj *rsvpLspIpv4Interface) SetIpv4Name(value string) RsvpLspIpv4Interface {

	obj.obj.Ipv4Name = &value
	return obj
}

// Contains properties of Tail(Egress) LSPs.
// P2PEgressIpv4Lsps returns a RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
func (obj *rsvpLspIpv4Interface) P2PEgressIpv4Lsps() RsvpLspIpv4InterfaceP2PEgressIpv4Lsp {
	if obj.obj.P2PEgressIpv4Lsps == nil {
		obj.obj.P2PEgressIpv4Lsps = NewRsvpLspIpv4InterfaceP2PEgressIpv4Lsp().msg()
	}
	if obj.p2PEgressIpv4LspsHolder == nil {
		obj.p2PEgressIpv4LspsHolder = &rsvpLspIpv4InterfaceP2PEgressIpv4Lsp{obj: obj.obj.P2PEgressIpv4Lsps}
	}
	return obj.p2PEgressIpv4LspsHolder
}

// Contains properties of Tail(Egress) LSPs.
// P2PEgressIpv4Lsps returns a RsvpLspIpv4InterfaceP2PEgressIpv4Lsp
func (obj *rsvpLspIpv4Interface) HasP2PEgressIpv4Lsps() bool {
	return obj.obj.P2PEgressIpv4Lsps != nil
}

// Contains properties of Tail(Egress) LSPs.
// SetP2PEgressIpv4Lsps sets the RsvpLspIpv4InterfaceP2PEgressIpv4Lsp value in the RsvpLspIpv4Interface object
func (obj *rsvpLspIpv4Interface) SetP2PEgressIpv4Lsps(value RsvpLspIpv4InterfaceP2PEgressIpv4Lsp) RsvpLspIpv4Interface {

	obj.p2PEgressIpv4LspsHolder = nil
	obj.obj.P2PEgressIpv4Lsps = value.msg()

	return obj
}

// Array of point-to-point RSVP-TE P2P LSPs originating from this interface.
// P2PIngressIpv4Lsps returns a []RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
func (obj *rsvpLspIpv4Interface) P2PIngressIpv4Lsps() RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	if len(obj.obj.P2PIngressIpv4Lsps) == 0 {
		obj.obj.P2PIngressIpv4Lsps = []*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp{}
	}
	if obj.p2PIngressIpv4LspsHolder == nil {
		obj.p2PIngressIpv4LspsHolder = newRsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter(&obj.obj.P2PIngressIpv4Lsps).setMsg(obj)
	}
	return obj.p2PIngressIpv4LspsHolder
}

type rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter struct {
	obj                                        *rsvpLspIpv4Interface
	rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice []RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	fieldPtr                                   *[]*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
}

func newRsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter(ptr *[]*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	return &rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter{fieldPtr: ptr}
}

type RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter interface {
	setMsg(*rsvpLspIpv4Interface) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
	Items() []RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	Add() RsvpLspIpv4InterfaceP2PIngressIpv4Lsp
	Append(items ...RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
	Set(index int, newObj RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
	Clear() RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
	clearHolderSlice() RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
	appendHolderSlice(item RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter
}

func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) setMsg(msg *rsvpLspIpv4Interface) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpLspIpv4InterfaceP2PIngressIpv4Lsp{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) Items() []RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	return obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice
}

func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) Add() RsvpLspIpv4InterfaceP2PIngressIpv4Lsp {
	newObj := &otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpLspIpv4InterfaceP2PIngressIpv4Lsp{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice = append(obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice, newLibObj)
	return newLibObj
}

func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) Append(items ...RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice = append(obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice, item)
	}
	return obj
}

func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) Set(index int, newObj RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice[index] = newObj
	return obj
}
func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) Clear() RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpLspIpv4InterfaceP2PIngressIpv4Lsp{}
		obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice = []RsvpLspIpv4InterfaceP2PIngressIpv4Lsp{}
	}
	return obj
}
func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) clearHolderSlice() RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	if len(obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice) > 0 {
		obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice = []RsvpLspIpv4InterfaceP2PIngressIpv4Lsp{}
	}
	return obj
}
func (obj *rsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter) appendHolderSlice(item RsvpLspIpv4InterfaceP2PIngressIpv4Lsp) RsvpLspIpv4InterfaceRsvpLspIpv4InterfaceP2PIngressIpv4LspIter {
	obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice = append(obj.rsvpLspIpv4InterfaceP2PIngressIpv4LspSlice, item)
	return obj
}

func (obj *rsvpLspIpv4Interface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface RsvpLspIpv4Interface")
	}

	if obj.obj.P2PEgressIpv4Lsps != nil {

		obj.P2PEgressIpv4Lsps().validateObj(vObj, set_default)
	}

	if len(obj.obj.P2PIngressIpv4Lsps) != 0 {

		if set_default {
			obj.P2PIngressIpv4Lsps().clearHolderSlice()
			for _, item := range obj.obj.P2PIngressIpv4Lsps {
				obj.P2PIngressIpv4Lsps().appendHolderSlice(&rsvpLspIpv4InterfaceP2PIngressIpv4Lsp{obj: item})
			}
		}
		for _, item := range obj.P2PIngressIpv4Lsps().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rsvpLspIpv4Interface) setDefault() {

}
