package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigBgp *****
type updateProtocolConfigBgp struct {
	validation
	obj                 *otg.UpdateProtocolConfigBgp
	marshaller          marshalUpdateProtocolConfigBgp
	unMarshaller        unMarshalUpdateProtocolConfigBgp
	routersHolder       UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	v4RouteRangesHolder UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	v6RouteRangesHolder UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
}

func NewUpdateProtocolConfigBgp() UpdateProtocolConfigBgp {
	obj := updateProtocolConfigBgp{obj: &otg.UpdateProtocolConfigBgp{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigBgp) msg() *otg.UpdateProtocolConfigBgp {
	return obj.obj
}

func (obj *updateProtocolConfigBgp) setMsg(msg *otg.UpdateProtocolConfigBgp) UpdateProtocolConfigBgp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigBgp struct {
	obj *updateProtocolConfigBgp
}

type marshalUpdateProtocolConfigBgp interface {
	// ToProto marshals UpdateProtocolConfigBgp to protobuf object *otg.UpdateProtocolConfigBgp
	ToProto() (*otg.UpdateProtocolConfigBgp, error)
	// ToPbText marshals UpdateProtocolConfigBgp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigBgp to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigBgp to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigBgp struct {
	obj *updateProtocolConfigBgp
}

type unMarshalUpdateProtocolConfigBgp interface {
	// FromProto unmarshals UpdateProtocolConfigBgp from protobuf object *otg.UpdateProtocolConfigBgp
	FromProto(msg *otg.UpdateProtocolConfigBgp) (UpdateProtocolConfigBgp, error)
	// FromPbText unmarshals UpdateProtocolConfigBgp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigBgp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigBgp from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigBgp) Marshal() marshalUpdateProtocolConfigBgp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigBgp{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigBgp) Unmarshal() unMarshalUpdateProtocolConfigBgp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigBgp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigBgp) ToProto() (*otg.UpdateProtocolConfigBgp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigBgp) FromProto(msg *otg.UpdateProtocolConfigBgp) (UpdateProtocolConfigBgp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigBgp) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgp) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigBgp) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgp) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigBgp) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgp) FromJson(value string) error {
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

func (obj *updateProtocolConfigBgp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigBgp) Clone() (UpdateProtocolConfigBgp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigBgp()
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

func (obj *updateProtocolConfigBgp) setNil() {
	obj.routersHolder = nil
	obj.v4RouteRangesHolder = nil
	obj.v6RouteRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigBgp is a container for BGP properties to be updated. Presence of this object indicates that one or more BGP properties require updating. The choice field selects the category of attributes to be updated. Router-level and interface-level updates are controlled independently.
type UpdateProtocolConfigBgp interface {
	Validation
	// msg marshals UpdateProtocolConfigBgp to protobuf object *otg.UpdateProtocolConfigBgp
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigBgp
	// setMsg unmarshals UpdateProtocolConfigBgp from protobuf object *otg.UpdateProtocolConfigBgp
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigBgp) UpdateProtocolConfigBgp
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigBgp
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigBgp
	// validate validates UpdateProtocolConfigBgp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigBgp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Routers returns UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIterIter, set in UpdateProtocolConfigBgp
	Routers() UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	// V4RouteRanges returns UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIterIter, set in UpdateProtocolConfigBgp
	V4RouteRanges() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	// V6RouteRanges returns UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIterIter, set in UpdateProtocolConfigBgp
	V6RouteRanges() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
	setNil()
}

// List of BGP router-level update entries. Populate this field to update router-level attributes independently of interface-level attributes.
// Routers returns a []UpdateProtocolConfigBgpRouterUpdateGroup
func (obj *updateProtocolConfigBgp) Routers() UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	if len(obj.obj.Routers) == 0 {
		obj.obj.Routers = []*otg.UpdateProtocolConfigBgpRouterUpdateGroup{}
	}
	if obj.routersHolder == nil {
		obj.routersHolder = newUpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter(&obj.obj.Routers).setMsg(obj)
	}
	return obj.routersHolder
}

type updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter struct {
	obj                                           *updateProtocolConfigBgp
	updateProtocolConfigBgpRouterUpdateGroupSlice []UpdateProtocolConfigBgpRouterUpdateGroup
	fieldPtr                                      *[]*otg.UpdateProtocolConfigBgpRouterUpdateGroup
}

func newUpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter(ptr *[]*otg.UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	return &updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter{fieldPtr: ptr}
}

type UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter interface {
	setMsg(*updateProtocolConfigBgp) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	Items() []UpdateProtocolConfigBgpRouterUpdateGroup
	Add() UpdateProtocolConfigBgpRouterUpdateGroup
	Append(items ...UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	Set(index int, newObj UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	Clear() UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	clearHolderSlice() UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
	appendHolderSlice(item UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) setMsg(msg *updateProtocolConfigBgp) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigBgpRouterUpdateGroup{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) Items() []UpdateProtocolConfigBgpRouterUpdateGroup {
	return obj.updateProtocolConfigBgpRouterUpdateGroupSlice
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) Add() UpdateProtocolConfigBgpRouterUpdateGroup {
	newObj := &otg.UpdateProtocolConfigBgpRouterUpdateGroup{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigBgpRouterUpdateGroup{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigBgpRouterUpdateGroupSlice = append(obj.updateProtocolConfigBgpRouterUpdateGroupSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) Append(items ...UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigBgpRouterUpdateGroupSlice = append(obj.updateProtocolConfigBgpRouterUpdateGroupSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) Set(index int, newObj UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigBgpRouterUpdateGroupSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) Clear() UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigBgpRouterUpdateGroup{}
		obj.updateProtocolConfigBgpRouterUpdateGroupSlice = []UpdateProtocolConfigBgpRouterUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) clearHolderSlice() UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	if len(obj.updateProtocolConfigBgpRouterUpdateGroupSlice) > 0 {
		obj.updateProtocolConfigBgpRouterUpdateGroupSlice = []UpdateProtocolConfigBgpRouterUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter) appendHolderSlice(item UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpRouterUpdateGroupIter {
	obj.updateProtocolConfigBgpRouterUpdateGroupSlice = append(obj.updateProtocolConfigBgpRouterUpdateGroupSlice, item)
	return obj
}

// List of BGP IPv4 route range update entries. Populate this field to update route range attributes independently of router-level attributes.
// V4RouteRanges returns a []UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
func (obj *updateProtocolConfigBgp) V4RouteRanges() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	if len(obj.obj.V4RouteRanges) == 0 {
		obj.obj.V4RouteRanges = []*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup{}
	}
	if obj.v4RouteRangesHolder == nil {
		obj.v4RouteRangesHolder = newUpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter(&obj.obj.V4RouteRanges).setMsg(obj)
	}
	return obj.v4RouteRangesHolder
}

type updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter struct {
	obj                                                 *updateProtocolConfigBgp
	updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice []UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	fieldPtr                                            *[]*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
}

func newUpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter(ptr *[]*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	return &updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter{fieldPtr: ptr}
}

type UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter interface {
	setMsg(*updateProtocolConfigBgp) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	Items() []UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	Add() UpdateProtocolConfigBgpV4RouteRangeUpdateGroup
	Append(items ...UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	Set(index int, newObj UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	Clear() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	clearHolderSlice() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
	appendHolderSlice(item UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) setMsg(msg *updateProtocolConfigBgp) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigBgpV4RouteRangeUpdateGroup{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) Items() []UpdateProtocolConfigBgpV4RouteRangeUpdateGroup {
	return obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) Add() UpdateProtocolConfigBgpV4RouteRangeUpdateGroup {
	newObj := &otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigBgpV4RouteRangeUpdateGroup{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice = append(obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) Append(items ...UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice = append(obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) Set(index int, newObj UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) Clear() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigBgpV4RouteRangeUpdateGroup{}
		obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice = []UpdateProtocolConfigBgpV4RouteRangeUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) clearHolderSlice() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	if len(obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice) > 0 {
		obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice = []UpdateProtocolConfigBgpV4RouteRangeUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter) appendHolderSlice(item UpdateProtocolConfigBgpV4RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV4RouteRangeUpdateGroupIter {
	obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice = append(obj.updateProtocolConfigBgpV4RouteRangeUpdateGroupSlice, item)
	return obj
}

// List of BGP IPv6 route range update entries. Populate this field to update route range attributes independently of router-level attributes.
// V6RouteRanges returns a []UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
func (obj *updateProtocolConfigBgp) V6RouteRanges() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	if len(obj.obj.V6RouteRanges) == 0 {
		obj.obj.V6RouteRanges = []*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup{}
	}
	if obj.v6RouteRangesHolder == nil {
		obj.v6RouteRangesHolder = newUpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter(&obj.obj.V6RouteRanges).setMsg(obj)
	}
	return obj.v6RouteRangesHolder
}

type updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter struct {
	obj                                                 *updateProtocolConfigBgp
	updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice []UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	fieldPtr                                            *[]*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
}

func newUpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter(ptr *[]*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	return &updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter{fieldPtr: ptr}
}

type UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter interface {
	setMsg(*updateProtocolConfigBgp) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
	Items() []UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	Add() UpdateProtocolConfigBgpV6RouteRangeUpdateGroup
	Append(items ...UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
	Set(index int, newObj UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
	Clear() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
	clearHolderSlice() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
	appendHolderSlice(item UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) setMsg(msg *updateProtocolConfigBgp) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&updateProtocolConfigBgpV6RouteRangeUpdateGroup{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) Items() []UpdateProtocolConfigBgpV6RouteRangeUpdateGroup {
	return obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) Add() UpdateProtocolConfigBgpV6RouteRangeUpdateGroup {
	newObj := &otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &updateProtocolConfigBgpV6RouteRangeUpdateGroup{obj: newObj}
	newLibObj.setDefault()
	obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice = append(obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice, newLibObj)
	return newLibObj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) Append(items ...UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice = append(obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice, item)
	}
	return obj
}

func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) Set(index int, newObj UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice[index] = newObj
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) Clear() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UpdateProtocolConfigBgpV6RouteRangeUpdateGroup{}
		obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice = []UpdateProtocolConfigBgpV6RouteRangeUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) clearHolderSlice() UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	if len(obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice) > 0 {
		obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice = []UpdateProtocolConfigBgpV6RouteRangeUpdateGroup{}
	}
	return obj
}
func (obj *updateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter) appendHolderSlice(item UpdateProtocolConfigBgpV6RouteRangeUpdateGroup) UpdateProtocolConfigBgpUpdateProtocolConfigBgpV6RouteRangeUpdateGroupIter {
	obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice = append(obj.updateProtocolConfigBgpV6RouteRangeUpdateGroupSlice, item)
	return obj
}

func (obj *updateProtocolConfigBgp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Routers) != 0 {

		if set_default {
			obj.Routers().clearHolderSlice()
			for _, item := range obj.obj.Routers {
				obj.Routers().appendHolderSlice(&updateProtocolConfigBgpRouterUpdateGroup{obj: item})
			}
		}
		for _, item := range obj.Routers().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V4RouteRanges) != 0 {

		if set_default {
			obj.V4RouteRanges().clearHolderSlice()
			for _, item := range obj.obj.V4RouteRanges {
				obj.V4RouteRanges().appendHolderSlice(&updateProtocolConfigBgpV4RouteRangeUpdateGroup{obj: item})
			}
		}
		for _, item := range obj.V4RouteRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.V6RouteRanges) != 0 {

		if set_default {
			obj.V6RouteRanges().clearHolderSlice()
			for _, item := range obj.obj.V6RouteRanges {
				obj.V6RouteRanges().appendHolderSlice(&updateProtocolConfigBgpV6RouteRangeUpdateGroup{obj: item})
			}
		}
		for _, item := range obj.V6RouteRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateProtocolConfigBgp) setDefault() {

}
