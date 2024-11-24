package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Lag *****
type lag struct {
	validation
	obj            *otg.Lag
	marshaller     marshalLag
	unMarshaller   unMarshalLag
	portsHolder    LagLagPortIter
	protocolHolder LagProtocol
}

func NewLag() Lag {
	obj := lag{obj: &otg.Lag{}}
	obj.setDefault()
	return &obj
}

func (obj *lag) msg() *otg.Lag {
	return obj.obj
}

func (obj *lag) setMsg(msg *otg.Lag) Lag {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallag struct {
	obj *lag
}

type marshalLag interface {
	// ToProto marshals Lag to protobuf object *otg.Lag
	ToProto() (*otg.Lag, error)
	// ToPbText marshals Lag to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Lag to YAML text
	ToYaml() (string, error)
	// ToJson marshals Lag to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Lag to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallag struct {
	obj *lag
}

type unMarshalLag interface {
	// FromProto unmarshals Lag from protobuf object *otg.Lag
	FromProto(msg *otg.Lag) (Lag, error)
	// FromPbText unmarshals Lag from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Lag from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Lag from JSON text
	FromJson(value string) error
}

func (obj *lag) Marshal() marshalLag {
	if obj.marshaller == nil {
		obj.marshaller = &marshallag{obj: obj}
	}
	return obj.marshaller
}

func (obj *lag) Unmarshal() unMarshalLag {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallag{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallag) ToProto() (*otg.Lag, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallag) FromProto(msg *otg.Lag) (Lag, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallag) ToPbText() (string, error) {
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

func (m *unMarshallag) FromPbText(value string) error {
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

func (m *marshallag) ToYaml() (string, error) {
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

func (m *unMarshallag) FromYaml(value string) error {
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

func (m *marshallag) ToJsonRaw() (string, error) {
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

func (m *marshallag) ToJson() (string, error) {
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

func (m *unMarshallag) FromJson(value string) error {
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

func (obj *lag) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lag) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lag) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lag) Clone() (Lag, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLag()
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

func (obj *lag) setNil() {
	obj.portsHolder = nil
	obj.protocolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Lag is the container for LAG (ports group) - aggregation of multiple LAG members (ports)
type Lag interface {
	Validation
	// msg marshals Lag to protobuf object *otg.Lag
	// and doesn't set defaults
	msg() *otg.Lag
	// setMsg unmarshals Lag from protobuf object *otg.Lag
	// and doesn't set defaults
	setMsg(*otg.Lag) Lag
	// provides marshal interface
	Marshal() marshalLag
	// provides unmarshal interface
	Unmarshal() unMarshalLag
	// validate validates Lag
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Lag, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ports returns LagLagPortIterIter, set in Lag
	Ports() LagLagPortIter
	// Protocol returns LagProtocol, set in Lag.
	// LagProtocol is description is TBD
	Protocol() LagProtocol
	// SetProtocol assigns LagProtocol provided by user to Lag.
	// LagProtocol is description is TBD
	SetProtocol(value LagProtocol) Lag
	// HasProtocol checks if Protocol has been set in Lag
	HasProtocol() bool
	// MinLinks returns uint32, set in Lag.
	MinLinks() uint32
	// SetMinLinks assigns uint32 provided by user to Lag
	SetMinLinks(value uint32) Lag
	// HasMinLinks checks if MinLinks has been set in Lag
	HasMinLinks() bool
	// Name returns string, set in Lag.
	Name() string
	// SetName assigns string provided by user to Lag
	SetName(value string) Lag
	setNil()
}

// description is TBD
// Ports returns a []LagPort
func (obj *lag) Ports() LagLagPortIter {
	if len(obj.obj.Ports) == 0 {
		obj.obj.Ports = []*otg.LagPort{}
	}
	if obj.portsHolder == nil {
		obj.portsHolder = newLagLagPortIter(&obj.obj.Ports).setMsg(obj)
	}
	return obj.portsHolder
}

type lagLagPortIter struct {
	obj          *lag
	lagPortSlice []LagPort
	fieldPtr     *[]*otg.LagPort
}

func newLagLagPortIter(ptr *[]*otg.LagPort) LagLagPortIter {
	return &lagLagPortIter{fieldPtr: ptr}
}

type LagLagPortIter interface {
	setMsg(*lag) LagLagPortIter
	Items() []LagPort
	Add() LagPort
	Append(items ...LagPort) LagLagPortIter
	Set(index int, newObj LagPort) LagLagPortIter
	Clear() LagLagPortIter
	clearHolderSlice() LagLagPortIter
	appendHolderSlice(item LagPort) LagLagPortIter
}

func (obj *lagLagPortIter) setMsg(msg *lag) LagLagPortIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&lagPort{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *lagLagPortIter) Items() []LagPort {
	return obj.lagPortSlice
}

func (obj *lagLagPortIter) Add() LagPort {
	newObj := &otg.LagPort{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &lagPort{obj: newObj}
	newLibObj.setDefault()
	obj.lagPortSlice = append(obj.lagPortSlice, newLibObj)
	return newLibObj
}

func (obj *lagLagPortIter) Append(items ...LagPort) LagLagPortIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.lagPortSlice = append(obj.lagPortSlice, item)
	}
	return obj
}

func (obj *lagLagPortIter) Set(index int, newObj LagPort) LagLagPortIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.lagPortSlice[index] = newObj
	return obj
}
func (obj *lagLagPortIter) Clear() LagLagPortIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.LagPort{}
		obj.lagPortSlice = []LagPort{}
	}
	return obj
}
func (obj *lagLagPortIter) clearHolderSlice() LagLagPortIter {
	if len(obj.lagPortSlice) > 0 {
		obj.lagPortSlice = []LagPort{}
	}
	return obj
}
func (obj *lagLagPortIter) appendHolderSlice(item LagPort) LagLagPortIter {
	obj.lagPortSlice = append(obj.lagPortSlice, item)
	return obj
}

// description is TBD
// Protocol returns a LagProtocol
func (obj *lag) Protocol() LagProtocol {
	if obj.obj.Protocol == nil {
		obj.obj.Protocol = NewLagProtocol().msg()
	}
	if obj.protocolHolder == nil {
		obj.protocolHolder = &lagProtocol{obj: obj.obj.Protocol}
	}
	return obj.protocolHolder
}

// description is TBD
// Protocol returns a LagProtocol
func (obj *lag) HasProtocol() bool {
	return obj.obj.Protocol != nil
}

// description is TBD
// SetProtocol sets the LagProtocol value in the Lag object
func (obj *lag) SetProtocol(value LagProtocol) Lag {

	obj.protocolHolder = nil
	obj.obj.Protocol = value.msg()

	return obj
}

// Specifies the mininum number of member interfaces that must be active for the aggregate interface to be available.
// If the aggregate interface is not available due to min-links criterion not being met, LACPDUs continue to be transmitted and received by the member interfaces if LACP is enabled, but other PDUs are not transmitted or received.
// MinLinks returns a uint32
func (obj *lag) MinLinks() uint32 {

	return *obj.obj.MinLinks

}

// Specifies the mininum number of member interfaces that must be active for the aggregate interface to be available.
// If the aggregate interface is not available due to min-links criterion not being met, LACPDUs continue to be transmitted and received by the member interfaces if LACP is enabled, but other PDUs are not transmitted or received.
// MinLinks returns a uint32
func (obj *lag) HasMinLinks() bool {
	return obj.obj.MinLinks != nil
}

// Specifies the mininum number of member interfaces that must be active for the aggregate interface to be available.
// If the aggregate interface is not available due to min-links criterion not being met, LACPDUs continue to be transmitted and received by the member interfaces if LACP is enabled, but other PDUs are not transmitted or received.
// SetMinLinks sets the uint32 value in the Lag object
func (obj *lag) SetMinLinks(value uint32) Lag {

	obj.obj.MinLinks = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *lag) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Lag object
func (obj *lag) SetName(value string) Lag {

	obj.obj.Name = &value
	return obj
}

func (obj *lag) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ports) != 0 {

		if set_default {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				obj.Ports().appendHolderSlice(&lagPort{obj: item})
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Protocol != nil {

		obj.Protocol().validateObj(vObj, set_default)
	}

	if obj.obj.MinLinks != nil {

		if *obj.obj.MinLinks > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Lag.MinLinks <= 32 but Got %d", *obj.obj.MinLinks))
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Lag")
	}
}

func (obj *lag) setDefault() {
	if obj.obj.MinLinks == nil {
		obj.SetMinLinks(1)
	}

}
