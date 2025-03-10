package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspTlvs *****
type isisLspTlvs struct {
	validation
	obj                                *otg.IsisLspTlvs
	marshaller                         marshalIsisLspTlvs
	unMarshaller                       unMarshalIsisLspTlvs
	hostnameTlvsHolder                 IsisLspTlvsIsisLspHostnameIter
	isReachabilityTlvsHolder           IsisLspTlvsIsisLspIsReachabilityTlvIter
	extendedIsReachabilityTlvsHolder   IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	ipv4InternalReachabilityTlvsHolder IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	ipv4ExternalReachabilityTlvsHolder IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	extendedIpv4ReachabilityTlvsHolder IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	ipv6ReachabilityTlvsHolder         IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	routerCapabilitiesHolder           IsisLspTlvsIsisLspCapabilityIter
}

func NewIsisLspTlvs() IsisLspTlvs {
	obj := isisLspTlvs{obj: &otg.IsisLspTlvs{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspTlvs) msg() *otg.IsisLspTlvs {
	return obj.obj
}

func (obj *isisLspTlvs) setMsg(msg *otg.IsisLspTlvs) IsisLspTlvs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspTlvs struct {
	obj *isisLspTlvs
}

type marshalIsisLspTlvs interface {
	// ToProto marshals IsisLspTlvs to protobuf object *otg.IsisLspTlvs
	ToProto() (*otg.IsisLspTlvs, error)
	// ToPbText marshals IsisLspTlvs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspTlvs to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspTlvs to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspTlvs to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspTlvs struct {
	obj *isisLspTlvs
}

type unMarshalIsisLspTlvs interface {
	// FromProto unmarshals IsisLspTlvs from protobuf object *otg.IsisLspTlvs
	FromProto(msg *otg.IsisLspTlvs) (IsisLspTlvs, error)
	// FromPbText unmarshals IsisLspTlvs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspTlvs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspTlvs from JSON text
	FromJson(value string) error
}

func (obj *isisLspTlvs) Marshal() marshalIsisLspTlvs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspTlvs{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspTlvs) Unmarshal() unMarshalIsisLspTlvs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspTlvs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspTlvs) ToProto() (*otg.IsisLspTlvs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspTlvs) FromProto(msg *otg.IsisLspTlvs) (IsisLspTlvs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspTlvs) ToPbText() (string, error) {
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

func (m *unMarshalisisLspTlvs) FromPbText(value string) error {
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

func (m *marshalisisLspTlvs) ToYaml() (string, error) {
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

func (m *unMarshalisisLspTlvs) FromYaml(value string) error {
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

func (m *marshalisisLspTlvs) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspTlvs) ToJson() (string, error) {
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

func (m *unMarshalisisLspTlvs) FromJson(value string) error {
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

func (obj *isisLspTlvs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspTlvs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspTlvs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspTlvs) Clone() (IsisLspTlvs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspTlvs()
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

func (obj *isisLspTlvs) setNil() {
	obj.hostnameTlvsHolder = nil
	obj.isReachabilityTlvsHolder = nil
	obj.extendedIsReachabilityTlvsHolder = nil
	obj.ipv4InternalReachabilityTlvsHolder = nil
	obj.ipv4ExternalReachabilityTlvsHolder = nil
	obj.extendedIpv4ReachabilityTlvsHolder = nil
	obj.ipv6ReachabilityTlvsHolder = nil
	obj.routerCapabilitiesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspTlvs is this contains the list of TLVs present in one LSP.
type IsisLspTlvs interface {
	Validation
	// msg marshals IsisLspTlvs to protobuf object *otg.IsisLspTlvs
	// and doesn't set defaults
	msg() *otg.IsisLspTlvs
	// setMsg unmarshals IsisLspTlvs from protobuf object *otg.IsisLspTlvs
	// and doesn't set defaults
	setMsg(*otg.IsisLspTlvs) IsisLspTlvs
	// provides marshal interface
	Marshal() marshalIsisLspTlvs
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspTlvs
	// validate validates IsisLspTlvs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspTlvs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HostnameTlvs returns IsisLspTlvsIsisLspHostnameIterIter, set in IsisLspTlvs
	HostnameTlvs() IsisLspTlvsIsisLspHostnameIter
	// IsReachabilityTlvs returns IsisLspTlvsIsisLspIsReachabilityTlvIterIter, set in IsisLspTlvs
	IsReachabilityTlvs() IsisLspTlvsIsisLspIsReachabilityTlvIter
	// ExtendedIsReachabilityTlvs returns IsisLspTlvsIsisLspExtendedIsReachabilityTlvIterIter, set in IsisLspTlvs
	ExtendedIsReachabilityTlvs() IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	// Ipv4InternalReachabilityTlvs returns IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIterIter, set in IsisLspTlvs
	Ipv4InternalReachabilityTlvs() IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	// Ipv4ExternalReachabilityTlvs returns IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIterIter, set in IsisLspTlvs
	Ipv4ExternalReachabilityTlvs() IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	// ExtendedIpv4ReachabilityTlvs returns IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIterIter, set in IsisLspTlvs
	ExtendedIpv4ReachabilityTlvs() IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	// Ipv6ReachabilityTlvs returns IsisLspTlvsIsisLspIpv6ReachabilityTlvIterIter, set in IsisLspTlvs
	Ipv6ReachabilityTlvs() IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	// RouterCapabilities returns IsisLspTlvsIsisLspCapabilityIterIter, set in IsisLspTlvs
	RouterCapabilities() IsisLspTlvsIsisLspCapabilityIter
	setNil()
}

// Array of Hostname TLVs ( type 137) present in this LSP.
// HostnameTlvs returns a []IsisLspHostname
func (obj *isisLspTlvs) HostnameTlvs() IsisLspTlvsIsisLspHostnameIter {
	if len(obj.obj.HostnameTlvs) == 0 {
		obj.obj.HostnameTlvs = []*otg.IsisLspHostname{}
	}
	if obj.hostnameTlvsHolder == nil {
		obj.hostnameTlvsHolder = newIsisLspTlvsIsisLspHostnameIter(&obj.obj.HostnameTlvs).setMsg(obj)
	}
	return obj.hostnameTlvsHolder
}

type isisLspTlvsIsisLspHostnameIter struct {
	obj                  *isisLspTlvs
	isisLspHostnameSlice []IsisLspHostname
	fieldPtr             *[]*otg.IsisLspHostname
}

func newIsisLspTlvsIsisLspHostnameIter(ptr *[]*otg.IsisLspHostname) IsisLspTlvsIsisLspHostnameIter {
	return &isisLspTlvsIsisLspHostnameIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspHostnameIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspHostnameIter
	Items() []IsisLspHostname
	Add() IsisLspHostname
	Append(items ...IsisLspHostname) IsisLspTlvsIsisLspHostnameIter
	Set(index int, newObj IsisLspHostname) IsisLspTlvsIsisLspHostnameIter
	Clear() IsisLspTlvsIsisLspHostnameIter
	clearHolderSlice() IsisLspTlvsIsisLspHostnameIter
	appendHolderSlice(item IsisLspHostname) IsisLspTlvsIsisLspHostnameIter
}

func (obj *isisLspTlvsIsisLspHostnameIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspHostnameIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspHostname{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspHostnameIter) Items() []IsisLspHostname {
	return obj.isisLspHostnameSlice
}

func (obj *isisLspTlvsIsisLspHostnameIter) Add() IsisLspHostname {
	newObj := &otg.IsisLspHostname{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspHostname{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspHostnameSlice = append(obj.isisLspHostnameSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspHostnameIter) Append(items ...IsisLspHostname) IsisLspTlvsIsisLspHostnameIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspHostnameSlice = append(obj.isisLspHostnameSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspHostnameIter) Set(index int, newObj IsisLspHostname) IsisLspTlvsIsisLspHostnameIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspHostnameSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspHostnameIter) Clear() IsisLspTlvsIsisLspHostnameIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspHostname{}
		obj.isisLspHostnameSlice = []IsisLspHostname{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspHostnameIter) clearHolderSlice() IsisLspTlvsIsisLspHostnameIter {
	if len(obj.isisLspHostnameSlice) > 0 {
		obj.isisLspHostnameSlice = []IsisLspHostname{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspHostnameIter) appendHolderSlice(item IsisLspHostname) IsisLspTlvsIsisLspHostnameIter {
	obj.isisLspHostnameSlice = append(obj.isisLspHostnameSlice, item)
	return obj
}

// Array of IS-Reachability TLVs (type 2) present in this LSP.
// IsReachabilityTlvs returns a []IsisLspIsReachabilityTlv
func (obj *isisLspTlvs) IsReachabilityTlvs() IsisLspTlvsIsisLspIsReachabilityTlvIter {
	if len(obj.obj.IsReachabilityTlvs) == 0 {
		obj.obj.IsReachabilityTlvs = []*otg.IsisLspIsReachabilityTlv{}
	}
	if obj.isReachabilityTlvsHolder == nil {
		obj.isReachabilityTlvsHolder = newIsisLspTlvsIsisLspIsReachabilityTlvIter(&obj.obj.IsReachabilityTlvs).setMsg(obj)
	}
	return obj.isReachabilityTlvsHolder
}

type isisLspTlvsIsisLspIsReachabilityTlvIter struct {
	obj                           *isisLspTlvs
	isisLspIsReachabilityTlvSlice []IsisLspIsReachabilityTlv
	fieldPtr                      *[]*otg.IsisLspIsReachabilityTlv
}

func newIsisLspTlvsIsisLspIsReachabilityTlvIter(ptr *[]*otg.IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter {
	return &isisLspTlvsIsisLspIsReachabilityTlvIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspIsReachabilityTlvIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspIsReachabilityTlvIter
	Items() []IsisLspIsReachabilityTlv
	Add() IsisLspIsReachabilityTlv
	Append(items ...IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter
	Set(index int, newObj IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter
	Clear() IsisLspTlvsIsisLspIsReachabilityTlvIter
	clearHolderSlice() IsisLspTlvsIsisLspIsReachabilityTlvIter
	appendHolderSlice(item IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter
}

func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspIsReachabilityTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspIsReachabilityTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) Items() []IsisLspIsReachabilityTlv {
	return obj.isisLspIsReachabilityTlvSlice
}

func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) Add() IsisLspIsReachabilityTlv {
	newObj := &otg.IsisLspIsReachabilityTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspIsReachabilityTlv{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspIsReachabilityTlvSlice = append(obj.isisLspIsReachabilityTlvSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) Append(items ...IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspIsReachabilityTlvSlice = append(obj.isisLspIsReachabilityTlvSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) Set(index int, newObj IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspIsReachabilityTlvSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) Clear() IsisLspTlvsIsisLspIsReachabilityTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspIsReachabilityTlv{}
		obj.isisLspIsReachabilityTlvSlice = []IsisLspIsReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) clearHolderSlice() IsisLspTlvsIsisLspIsReachabilityTlvIter {
	if len(obj.isisLspIsReachabilityTlvSlice) > 0 {
		obj.isisLspIsReachabilityTlvSlice = []IsisLspIsReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIsReachabilityTlvIter) appendHolderSlice(item IsisLspIsReachabilityTlv) IsisLspTlvsIsisLspIsReachabilityTlvIter {
	obj.isisLspIsReachabilityTlvSlice = append(obj.isisLspIsReachabilityTlvSlice, item)
	return obj
}

// Array of Extended IS-Reachability TLVs (type 22) present in this LSP.
// ExtendedIsReachabilityTlvs returns a []IsisLspExtendedIsReachabilityTlv
func (obj *isisLspTlvs) ExtendedIsReachabilityTlvs() IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	if len(obj.obj.ExtendedIsReachabilityTlvs) == 0 {
		obj.obj.ExtendedIsReachabilityTlvs = []*otg.IsisLspExtendedIsReachabilityTlv{}
	}
	if obj.extendedIsReachabilityTlvsHolder == nil {
		obj.extendedIsReachabilityTlvsHolder = newIsisLspTlvsIsisLspExtendedIsReachabilityTlvIter(&obj.obj.ExtendedIsReachabilityTlvs).setMsg(obj)
	}
	return obj.extendedIsReachabilityTlvsHolder
}

type isisLspTlvsIsisLspExtendedIsReachabilityTlvIter struct {
	obj                                   *isisLspTlvs
	isisLspExtendedIsReachabilityTlvSlice []IsisLspExtendedIsReachabilityTlv
	fieldPtr                              *[]*otg.IsisLspExtendedIsReachabilityTlv
}

func newIsisLspTlvsIsisLspExtendedIsReachabilityTlvIter(ptr *[]*otg.IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	return &isisLspTlvsIsisLspExtendedIsReachabilityTlvIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	Items() []IsisLspExtendedIsReachabilityTlv
	Add() IsisLspExtendedIsReachabilityTlv
	Append(items ...IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	Set(index int, newObj IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	Clear() IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	clearHolderSlice() IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
	appendHolderSlice(item IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter
}

func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspExtendedIsReachabilityTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) Items() []IsisLspExtendedIsReachabilityTlv {
	return obj.isisLspExtendedIsReachabilityTlvSlice
}

func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) Add() IsisLspExtendedIsReachabilityTlv {
	newObj := &otg.IsisLspExtendedIsReachabilityTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspExtendedIsReachabilityTlv{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspExtendedIsReachabilityTlvSlice = append(obj.isisLspExtendedIsReachabilityTlvSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) Append(items ...IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspExtendedIsReachabilityTlvSlice = append(obj.isisLspExtendedIsReachabilityTlvSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) Set(index int, newObj IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspExtendedIsReachabilityTlvSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) Clear() IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspExtendedIsReachabilityTlv{}
		obj.isisLspExtendedIsReachabilityTlvSlice = []IsisLspExtendedIsReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) clearHolderSlice() IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	if len(obj.isisLspExtendedIsReachabilityTlvSlice) > 0 {
		obj.isisLspExtendedIsReachabilityTlvSlice = []IsisLspExtendedIsReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspExtendedIsReachabilityTlvIter) appendHolderSlice(item IsisLspExtendedIsReachabilityTlv) IsisLspTlvsIsisLspExtendedIsReachabilityTlvIter {
	obj.isisLspExtendedIsReachabilityTlvSlice = append(obj.isisLspExtendedIsReachabilityTlvSlice, item)
	return obj
}

// Array of IPv4 Internal Reachability TLVs (type 128) present in this LSP.
// Ipv4InternalReachabilityTlvs returns a []IsisLspIpv4InternalReachabilityTlv
func (obj *isisLspTlvs) Ipv4InternalReachabilityTlvs() IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	if len(obj.obj.Ipv4InternalReachabilityTlvs) == 0 {
		obj.obj.Ipv4InternalReachabilityTlvs = []*otg.IsisLspIpv4InternalReachabilityTlv{}
	}
	if obj.ipv4InternalReachabilityTlvsHolder == nil {
		obj.ipv4InternalReachabilityTlvsHolder = newIsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter(&obj.obj.Ipv4InternalReachabilityTlvs).setMsg(obj)
	}
	return obj.ipv4InternalReachabilityTlvsHolder
}

type isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter struct {
	obj                                     *isisLspTlvs
	isisLspIpv4InternalReachabilityTlvSlice []IsisLspIpv4InternalReachabilityTlv
	fieldPtr                                *[]*otg.IsisLspIpv4InternalReachabilityTlv
}

func newIsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter(ptr *[]*otg.IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	return &isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	Items() []IsisLspIpv4InternalReachabilityTlv
	Add() IsisLspIpv4InternalReachabilityTlv
	Append(items ...IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	Set(index int, newObj IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	Clear() IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	clearHolderSlice() IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
	appendHolderSlice(item IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter
}

func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspIpv4InternalReachabilityTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) Items() []IsisLspIpv4InternalReachabilityTlv {
	return obj.isisLspIpv4InternalReachabilityTlvSlice
}

func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) Add() IsisLspIpv4InternalReachabilityTlv {
	newObj := &otg.IsisLspIpv4InternalReachabilityTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspIpv4InternalReachabilityTlv{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspIpv4InternalReachabilityTlvSlice = append(obj.isisLspIpv4InternalReachabilityTlvSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) Append(items ...IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspIpv4InternalReachabilityTlvSlice = append(obj.isisLspIpv4InternalReachabilityTlvSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) Set(index int, newObj IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspIpv4InternalReachabilityTlvSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) Clear() IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspIpv4InternalReachabilityTlv{}
		obj.isisLspIpv4InternalReachabilityTlvSlice = []IsisLspIpv4InternalReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) clearHolderSlice() IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	if len(obj.isisLspIpv4InternalReachabilityTlvSlice) > 0 {
		obj.isisLspIpv4InternalReachabilityTlvSlice = []IsisLspIpv4InternalReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIpv4InternalReachabilityTlvIter) appendHolderSlice(item IsisLspIpv4InternalReachabilityTlv) IsisLspTlvsIsisLspIpv4InternalReachabilityTlvIter {
	obj.isisLspIpv4InternalReachabilityTlvSlice = append(obj.isisLspIpv4InternalReachabilityTlvSlice, item)
	return obj
}

// Array of IPv4 External Reachability TLVs (type 130) present in this LSP.
// Ipv4ExternalReachabilityTlvs returns a []IsisLspIpv4ExternalReachabilityTlv
func (obj *isisLspTlvs) Ipv4ExternalReachabilityTlvs() IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	if len(obj.obj.Ipv4ExternalReachabilityTlvs) == 0 {
		obj.obj.Ipv4ExternalReachabilityTlvs = []*otg.IsisLspIpv4ExternalReachabilityTlv{}
	}
	if obj.ipv4ExternalReachabilityTlvsHolder == nil {
		obj.ipv4ExternalReachabilityTlvsHolder = newIsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter(&obj.obj.Ipv4ExternalReachabilityTlvs).setMsg(obj)
	}
	return obj.ipv4ExternalReachabilityTlvsHolder
}

type isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter struct {
	obj                                     *isisLspTlvs
	isisLspIpv4ExternalReachabilityTlvSlice []IsisLspIpv4ExternalReachabilityTlv
	fieldPtr                                *[]*otg.IsisLspIpv4ExternalReachabilityTlv
}

func newIsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter(ptr *[]*otg.IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	return &isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	Items() []IsisLspIpv4ExternalReachabilityTlv
	Add() IsisLspIpv4ExternalReachabilityTlv
	Append(items ...IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	Set(index int, newObj IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	Clear() IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	clearHolderSlice() IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
	appendHolderSlice(item IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter
}

func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspIpv4ExternalReachabilityTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) Items() []IsisLspIpv4ExternalReachabilityTlv {
	return obj.isisLspIpv4ExternalReachabilityTlvSlice
}

func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) Add() IsisLspIpv4ExternalReachabilityTlv {
	newObj := &otg.IsisLspIpv4ExternalReachabilityTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspIpv4ExternalReachabilityTlv{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspIpv4ExternalReachabilityTlvSlice = append(obj.isisLspIpv4ExternalReachabilityTlvSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) Append(items ...IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspIpv4ExternalReachabilityTlvSlice = append(obj.isisLspIpv4ExternalReachabilityTlvSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) Set(index int, newObj IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspIpv4ExternalReachabilityTlvSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) Clear() IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspIpv4ExternalReachabilityTlv{}
		obj.isisLspIpv4ExternalReachabilityTlvSlice = []IsisLspIpv4ExternalReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) clearHolderSlice() IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	if len(obj.isisLspIpv4ExternalReachabilityTlvSlice) > 0 {
		obj.isisLspIpv4ExternalReachabilityTlvSlice = []IsisLspIpv4ExternalReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter) appendHolderSlice(item IsisLspIpv4ExternalReachabilityTlv) IsisLspTlvsIsisLspIpv4ExternalReachabilityTlvIter {
	obj.isisLspIpv4ExternalReachabilityTlvSlice = append(obj.isisLspIpv4ExternalReachabilityTlvSlice, item)
	return obj
}

// Array of IPv4 Extended Reachability TLVs (type 135) present in this LSP.
// ExtendedIpv4ReachabilityTlvs returns a []IsisLspExtendedIpv4ReachabilityTlv
func (obj *isisLspTlvs) ExtendedIpv4ReachabilityTlvs() IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	if len(obj.obj.ExtendedIpv4ReachabilityTlvs) == 0 {
		obj.obj.ExtendedIpv4ReachabilityTlvs = []*otg.IsisLspExtendedIpv4ReachabilityTlv{}
	}
	if obj.extendedIpv4ReachabilityTlvsHolder == nil {
		obj.extendedIpv4ReachabilityTlvsHolder = newIsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter(&obj.obj.ExtendedIpv4ReachabilityTlvs).setMsg(obj)
	}
	return obj.extendedIpv4ReachabilityTlvsHolder
}

type isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter struct {
	obj                                     *isisLspTlvs
	isisLspExtendedIpv4ReachabilityTlvSlice []IsisLspExtendedIpv4ReachabilityTlv
	fieldPtr                                *[]*otg.IsisLspExtendedIpv4ReachabilityTlv
}

func newIsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter(ptr *[]*otg.IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	return &isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	Items() []IsisLspExtendedIpv4ReachabilityTlv
	Add() IsisLspExtendedIpv4ReachabilityTlv
	Append(items ...IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	Set(index int, newObj IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	Clear() IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	clearHolderSlice() IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
	appendHolderSlice(item IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter
}

func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspExtendedIpv4ReachabilityTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) Items() []IsisLspExtendedIpv4ReachabilityTlv {
	return obj.isisLspExtendedIpv4ReachabilityTlvSlice
}

func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) Add() IsisLspExtendedIpv4ReachabilityTlv {
	newObj := &otg.IsisLspExtendedIpv4ReachabilityTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspExtendedIpv4ReachabilityTlv{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspExtendedIpv4ReachabilityTlvSlice = append(obj.isisLspExtendedIpv4ReachabilityTlvSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) Append(items ...IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspExtendedIpv4ReachabilityTlvSlice = append(obj.isisLspExtendedIpv4ReachabilityTlvSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) Set(index int, newObj IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspExtendedIpv4ReachabilityTlvSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) Clear() IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspExtendedIpv4ReachabilityTlv{}
		obj.isisLspExtendedIpv4ReachabilityTlvSlice = []IsisLspExtendedIpv4ReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) clearHolderSlice() IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	if len(obj.isisLspExtendedIpv4ReachabilityTlvSlice) > 0 {
		obj.isisLspExtendedIpv4ReachabilityTlvSlice = []IsisLspExtendedIpv4ReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter) appendHolderSlice(item IsisLspExtendedIpv4ReachabilityTlv) IsisLspTlvsIsisLspExtendedIpv4ReachabilityTlvIter {
	obj.isisLspExtendedIpv4ReachabilityTlvSlice = append(obj.isisLspExtendedIpv4ReachabilityTlvSlice, item)
	return obj
}

// Array of IPv6 Reachability TLVs (type 236) present in this LSP.
// Ipv6ReachabilityTlvs returns a []IsisLspIpv6ReachabilityTlv
func (obj *isisLspTlvs) Ipv6ReachabilityTlvs() IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	if len(obj.obj.Ipv6ReachabilityTlvs) == 0 {
		obj.obj.Ipv6ReachabilityTlvs = []*otg.IsisLspIpv6ReachabilityTlv{}
	}
	if obj.ipv6ReachabilityTlvsHolder == nil {
		obj.ipv6ReachabilityTlvsHolder = newIsisLspTlvsIsisLspIpv6ReachabilityTlvIter(&obj.obj.Ipv6ReachabilityTlvs).setMsg(obj)
	}
	return obj.ipv6ReachabilityTlvsHolder
}

type isisLspTlvsIsisLspIpv6ReachabilityTlvIter struct {
	obj                             *isisLspTlvs
	isisLspIpv6ReachabilityTlvSlice []IsisLspIpv6ReachabilityTlv
	fieldPtr                        *[]*otg.IsisLspIpv6ReachabilityTlv
}

func newIsisLspTlvsIsisLspIpv6ReachabilityTlvIter(ptr *[]*otg.IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	return &isisLspTlvsIsisLspIpv6ReachabilityTlvIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspIpv6ReachabilityTlvIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	Items() []IsisLspIpv6ReachabilityTlv
	Add() IsisLspIpv6ReachabilityTlv
	Append(items ...IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	Set(index int, newObj IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	Clear() IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	clearHolderSlice() IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
	appendHolderSlice(item IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter
}

func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspIpv6ReachabilityTlv{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) Items() []IsisLspIpv6ReachabilityTlv {
	return obj.isisLspIpv6ReachabilityTlvSlice
}

func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) Add() IsisLspIpv6ReachabilityTlv {
	newObj := &otg.IsisLspIpv6ReachabilityTlv{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspIpv6ReachabilityTlv{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspIpv6ReachabilityTlvSlice = append(obj.isisLspIpv6ReachabilityTlvSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) Append(items ...IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspIpv6ReachabilityTlvSlice = append(obj.isisLspIpv6ReachabilityTlvSlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) Set(index int, newObj IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspIpv6ReachabilityTlvSlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) Clear() IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspIpv6ReachabilityTlv{}
		obj.isisLspIpv6ReachabilityTlvSlice = []IsisLspIpv6ReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) clearHolderSlice() IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	if len(obj.isisLspIpv6ReachabilityTlvSlice) > 0 {
		obj.isisLspIpv6ReachabilityTlvSlice = []IsisLspIpv6ReachabilityTlv{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspIpv6ReachabilityTlvIter) appendHolderSlice(item IsisLspIpv6ReachabilityTlv) IsisLspTlvsIsisLspIpv6ReachabilityTlvIter {
	obj.isisLspIpv6ReachabilityTlvSlice = append(obj.isisLspIpv6ReachabilityTlvSlice, item)
	return obj
}

// IS-IS Router Capabilities: TLV 242.
// This container defines Router Capabilities.
// RouterCapabilities returns a []IsisLspCapability
func (obj *isisLspTlvs) RouterCapabilities() IsisLspTlvsIsisLspCapabilityIter {
	if len(obj.obj.RouterCapabilities) == 0 {
		obj.obj.RouterCapabilities = []*otg.IsisLspCapability{}
	}
	if obj.routerCapabilitiesHolder == nil {
		obj.routerCapabilitiesHolder = newIsisLspTlvsIsisLspCapabilityIter(&obj.obj.RouterCapabilities).setMsg(obj)
	}
	return obj.routerCapabilitiesHolder
}

type isisLspTlvsIsisLspCapabilityIter struct {
	obj                    *isisLspTlvs
	isisLspCapabilitySlice []IsisLspCapability
	fieldPtr               *[]*otg.IsisLspCapability
}

func newIsisLspTlvsIsisLspCapabilityIter(ptr *[]*otg.IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter {
	return &isisLspTlvsIsisLspCapabilityIter{fieldPtr: ptr}
}

type IsisLspTlvsIsisLspCapabilityIter interface {
	setMsg(*isisLspTlvs) IsisLspTlvsIsisLspCapabilityIter
	Items() []IsisLspCapability
	Add() IsisLspCapability
	Append(items ...IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter
	Set(index int, newObj IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter
	Clear() IsisLspTlvsIsisLspCapabilityIter
	clearHolderSlice() IsisLspTlvsIsisLspCapabilityIter
	appendHolderSlice(item IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter
}

func (obj *isisLspTlvsIsisLspCapabilityIter) setMsg(msg *isisLspTlvs) IsisLspTlvsIsisLspCapabilityIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspCapability{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspTlvsIsisLspCapabilityIter) Items() []IsisLspCapability {
	return obj.isisLspCapabilitySlice
}

func (obj *isisLspTlvsIsisLspCapabilityIter) Add() IsisLspCapability {
	newObj := &otg.IsisLspCapability{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspCapability{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspCapabilitySlice = append(obj.isisLspCapabilitySlice, newLibObj)
	return newLibObj
}

func (obj *isisLspTlvsIsisLspCapabilityIter) Append(items ...IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspCapabilitySlice = append(obj.isisLspCapabilitySlice, item)
	}
	return obj
}

func (obj *isisLspTlvsIsisLspCapabilityIter) Set(index int, newObj IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspCapabilitySlice[index] = newObj
	return obj
}
func (obj *isisLspTlvsIsisLspCapabilityIter) Clear() IsisLspTlvsIsisLspCapabilityIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspCapability{}
		obj.isisLspCapabilitySlice = []IsisLspCapability{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspCapabilityIter) clearHolderSlice() IsisLspTlvsIsisLspCapabilityIter {
	if len(obj.isisLspCapabilitySlice) > 0 {
		obj.isisLspCapabilitySlice = []IsisLspCapability{}
	}
	return obj
}
func (obj *isisLspTlvsIsisLspCapabilityIter) appendHolderSlice(item IsisLspCapability) IsisLspTlvsIsisLspCapabilityIter {
	obj.isisLspCapabilitySlice = append(obj.isisLspCapabilitySlice, item)
	return obj
}

func (obj *isisLspTlvs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.HostnameTlvs) != 0 {

		if set_default {
			obj.HostnameTlvs().clearHolderSlice()
			for _, item := range obj.obj.HostnameTlvs {
				obj.HostnameTlvs().appendHolderSlice(&isisLspHostname{obj: item})
			}
		}
		for _, item := range obj.HostnameTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.IsReachabilityTlvs) != 0 {

		if set_default {
			obj.IsReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.IsReachabilityTlvs {
				obj.IsReachabilityTlvs().appendHolderSlice(&isisLspIsReachabilityTlv{obj: item})
			}
		}
		for _, item := range obj.IsReachabilityTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExtendedIsReachabilityTlvs) != 0 {

		if set_default {
			obj.ExtendedIsReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.ExtendedIsReachabilityTlvs {
				obj.ExtendedIsReachabilityTlvs().appendHolderSlice(&isisLspExtendedIsReachabilityTlv{obj: item})
			}
		}
		for _, item := range obj.ExtendedIsReachabilityTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv4InternalReachabilityTlvs) != 0 {

		if set_default {
			obj.Ipv4InternalReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.Ipv4InternalReachabilityTlvs {
				obj.Ipv4InternalReachabilityTlvs().appendHolderSlice(&isisLspIpv4InternalReachabilityTlv{obj: item})
			}
		}
		for _, item := range obj.Ipv4InternalReachabilityTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv4ExternalReachabilityTlvs) != 0 {

		if set_default {
			obj.Ipv4ExternalReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.Ipv4ExternalReachabilityTlvs {
				obj.Ipv4ExternalReachabilityTlvs().appendHolderSlice(&isisLspIpv4ExternalReachabilityTlv{obj: item})
			}
		}
		for _, item := range obj.Ipv4ExternalReachabilityTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExtendedIpv4ReachabilityTlvs) != 0 {

		if set_default {
			obj.ExtendedIpv4ReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.ExtendedIpv4ReachabilityTlvs {
				obj.ExtendedIpv4ReachabilityTlvs().appendHolderSlice(&isisLspExtendedIpv4ReachabilityTlv{obj: item})
			}
		}
		for _, item := range obj.ExtendedIpv4ReachabilityTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Ipv6ReachabilityTlvs) != 0 {

		if set_default {
			obj.Ipv6ReachabilityTlvs().clearHolderSlice()
			for _, item := range obj.obj.Ipv6ReachabilityTlvs {
				obj.Ipv6ReachabilityTlvs().appendHolderSlice(&isisLspIpv6ReachabilityTlv{obj: item})
			}
		}
		for _, item := range obj.Ipv6ReachabilityTlvs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.RouterCapabilities) != 0 {

		if set_default {
			obj.RouterCapabilities().clearHolderSlice()
			for _, item := range obj.obj.RouterCapabilities {
				obj.RouterCapabilities().appendHolderSlice(&isisLspCapability{obj: item})
			}
		}
		for _, item := range obj.RouterCapabilities().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspTlvs) setDefault() {

}
