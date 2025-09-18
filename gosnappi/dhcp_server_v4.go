package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DhcpServerV4 *****
type dhcpServerV4 struct {
	validation
	obj                *otg.DhcpServerV4
	marshaller         marshalDhcpServerV4
	unMarshaller       unMarshalDhcpServerV4
	addressPoolsHolder DhcpServerV4DhcpServerV4PoolIter
}

func NewDhcpServerV4() DhcpServerV4 {
	obj := dhcpServerV4{obj: &otg.DhcpServerV4{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpServerV4) msg() *otg.DhcpServerV4 {
	return obj.obj
}

func (obj *dhcpServerV4) setMsg(msg *otg.DhcpServerV4) DhcpServerV4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpServerV4 struct {
	obj *dhcpServerV4
}

type marshalDhcpServerV4 interface {
	// ToProto marshals DhcpServerV4 to protobuf object *otg.DhcpServerV4
	ToProto() (*otg.DhcpServerV4, error)
	// ToPbText marshals DhcpServerV4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DhcpServerV4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals DhcpServerV4 to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpServerV4 struct {
	obj *dhcpServerV4
}

type unMarshalDhcpServerV4 interface {
	// FromProto unmarshals DhcpServerV4 from protobuf object *otg.DhcpServerV4
	FromProto(msg *otg.DhcpServerV4) (DhcpServerV4, error)
	// FromPbText unmarshals DhcpServerV4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DhcpServerV4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DhcpServerV4 from JSON text
	FromJson(value string) error
}

func (obj *dhcpServerV4) Marshal() marshalDhcpServerV4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpServerV4{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpServerV4) Unmarshal() unMarshalDhcpServerV4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpServerV4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpServerV4) ToProto() (*otg.DhcpServerV4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpServerV4) FromProto(msg *otg.DhcpServerV4) (DhcpServerV4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpServerV4) ToPbText() (string, error) {
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

func (m *unMarshaldhcpServerV4) FromPbText(value string) error {
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

func (m *marshaldhcpServerV4) ToYaml() (string, error) {
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

func (m *unMarshaldhcpServerV4) FromYaml(value string) error {
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

func (m *marshaldhcpServerV4) ToJson() (string, error) {
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

func (m *unMarshaldhcpServerV4) FromJson(value string) error {
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

func (obj *dhcpServerV4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpServerV4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpServerV4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpServerV4) Clone() (DhcpServerV4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpServerV4()
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

func (obj *dhcpServerV4) setNil() {
	obj.addressPoolsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DhcpServerV4 is configuration for emulated DHCPv4 Server.
type DhcpServerV4 interface {
	Validation
	// msg marshals DhcpServerV4 to protobuf object *otg.DhcpServerV4
	// and doesn't set defaults
	msg() *otg.DhcpServerV4
	// setMsg unmarshals DhcpServerV4 from protobuf object *otg.DhcpServerV4
	// and doesn't set defaults
	setMsg(*otg.DhcpServerV4) DhcpServerV4
	// provides marshal interface
	Marshal() marshalDhcpServerV4
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpServerV4
	// validate validates DhcpServerV4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DhcpServerV4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DhcpServerV4.
	Name() string
	// SetName assigns string provided by user to DhcpServerV4
	SetName(value string) DhcpServerV4
	// Ipv4Name returns string, set in DhcpServerV4.
	Ipv4Name() string
	// SetIpv4Name assigns string provided by user to DhcpServerV4
	SetIpv4Name(value string) DhcpServerV4
	// AddressPools returns DhcpServerV4DhcpServerV4PoolIterIter, set in DhcpServerV4
	AddressPools() DhcpServerV4DhcpServerV4PoolIter
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *dhcpServerV4) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DhcpServerV4 object
func (obj *dhcpServerV4) SetName(value string) DhcpServerV4 {

	obj.obj.Name = &value
	return obj
}

// The unique name of the IPv4 on which DHCPv4 server will run.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// Ipv4Name returns a string
func (obj *dhcpServerV4) Ipv4Name() string {

	return *obj.obj.Ipv4Name

}

// The unique name of the IPv4 on which DHCPv4 server will run.
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// x-constraint:
// - /components/schemas/Device.Ipv4/properties/name
//
// SetIpv4Name sets the string value in the DhcpServerV4 object
func (obj *dhcpServerV4) SetIpv4Name(value string) DhcpServerV4 {

	obj.obj.Ipv4Name = &value
	return obj
}

// List of DHCPv4 Server Lease parameters
// AddressPools returns a []DhcpServerV4Pool
func (obj *dhcpServerV4) AddressPools() DhcpServerV4DhcpServerV4PoolIter {
	if len(obj.obj.AddressPools) == 0 {
		obj.obj.AddressPools = []*otg.DhcpServerV4Pool{}
	}
	if obj.addressPoolsHolder == nil {
		obj.addressPoolsHolder = newDhcpServerV4DhcpServerV4PoolIter(&obj.obj.AddressPools).setMsg(obj)
	}
	return obj.addressPoolsHolder
}

type dhcpServerV4DhcpServerV4PoolIter struct {
	obj                   *dhcpServerV4
	dhcpServerV4PoolSlice []DhcpServerV4Pool
	fieldPtr              *[]*otg.DhcpServerV4Pool
}

func newDhcpServerV4DhcpServerV4PoolIter(ptr *[]*otg.DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter {
	return &dhcpServerV4DhcpServerV4PoolIter{fieldPtr: ptr}
}

type DhcpServerV4DhcpServerV4PoolIter interface {
	setMsg(*dhcpServerV4) DhcpServerV4DhcpServerV4PoolIter
	Items() []DhcpServerV4Pool
	Add() DhcpServerV4Pool
	Append(items ...DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter
	Set(index int, newObj DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter
	Clear() DhcpServerV4DhcpServerV4PoolIter
	clearHolderSlice() DhcpServerV4DhcpServerV4PoolIter
	appendHolderSlice(item DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter
}

func (obj *dhcpServerV4DhcpServerV4PoolIter) setMsg(msg *dhcpServerV4) DhcpServerV4DhcpServerV4PoolIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpServerV4Pool{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpServerV4DhcpServerV4PoolIter) Items() []DhcpServerV4Pool {
	return obj.dhcpServerV4PoolSlice
}

func (obj *dhcpServerV4DhcpServerV4PoolIter) Add() DhcpServerV4Pool {
	newObj := &otg.DhcpServerV4Pool{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpServerV4Pool{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpServerV4PoolSlice = append(obj.dhcpServerV4PoolSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpServerV4DhcpServerV4PoolIter) Append(items ...DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpServerV4PoolSlice = append(obj.dhcpServerV4PoolSlice, item)
	}
	return obj
}

func (obj *dhcpServerV4DhcpServerV4PoolIter) Set(index int, newObj DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpServerV4PoolSlice[index] = newObj
	return obj
}
func (obj *dhcpServerV4DhcpServerV4PoolIter) Clear() DhcpServerV4DhcpServerV4PoolIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DhcpServerV4Pool{}
		obj.dhcpServerV4PoolSlice = []DhcpServerV4Pool{}
	}
	return obj
}
func (obj *dhcpServerV4DhcpServerV4PoolIter) clearHolderSlice() DhcpServerV4DhcpServerV4PoolIter {
	if len(obj.dhcpServerV4PoolSlice) > 0 {
		obj.dhcpServerV4PoolSlice = []DhcpServerV4Pool{}
	}
	return obj
}
func (obj *dhcpServerV4DhcpServerV4PoolIter) appendHolderSlice(item DhcpServerV4Pool) DhcpServerV4DhcpServerV4PoolIter {
	obj.dhcpServerV4PoolSlice = append(obj.dhcpServerV4PoolSlice, item)
	return obj
}

func (obj *dhcpServerV4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DhcpServerV4")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DhcpServerV4.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	// Ipv4Name is required
	if obj.obj.Ipv4Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4Name is required field on interface DhcpServerV4")
	}

	if len(obj.obj.AddressPools) != 0 {

		if set_default {
			obj.AddressPools().clearHolderSlice()
			for _, item := range obj.obj.AddressPools {
				obj.AddressPools().appendHolderSlice(&dhcpServerV4Pool{obj: item})
			}
		}
		for _, item := range obj.AddressPools().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *dhcpServerV4) setDefault() {

}
