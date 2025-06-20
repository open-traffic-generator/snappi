package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3LsaState *****
type ospfv3LsaState struct {
	validation
	obj                       *otg.Ospfv3LsaState
	marshaller                marshalOspfv3LsaState
	unMarshaller              unMarshalOspfv3LsaState
	routerLsasHolder          Ospfv3LsaStateOspfv3RouterLsaIter
	networkLsasHolder         Ospfv3LsaStateOspfv3NetworkLsaIter
	interAreaPrefixLsasHolder Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	interAreaRouterLsasHolder Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	externalAsLsasHolder      Ospfv3LsaStateOspfv3ExternalAsLsaIter
	nssaLsasHolder            Ospfv3LsaStateOspfv3NssaLsaIter
	linkLsasHolder            Ospfv3LsaStateOspfv3LinkLsaIter
	intraAreaPrefixLsasHolder Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
}

func NewOspfv3LsaState() Ospfv3LsaState {
	obj := ospfv3LsaState{obj: &otg.Ospfv3LsaState{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3LsaState) msg() *otg.Ospfv3LsaState {
	return obj.obj
}

func (obj *ospfv3LsaState) setMsg(msg *otg.Ospfv3LsaState) Ospfv3LsaState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3LsaState struct {
	obj *ospfv3LsaState
}

type marshalOspfv3LsaState interface {
	// ToProto marshals Ospfv3LsaState to protobuf object *otg.Ospfv3LsaState
	ToProto() (*otg.Ospfv3LsaState, error)
	// ToPbText marshals Ospfv3LsaState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3LsaState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3LsaState to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3LsaState struct {
	obj *ospfv3LsaState
}

type unMarshalOspfv3LsaState interface {
	// FromProto unmarshals Ospfv3LsaState from protobuf object *otg.Ospfv3LsaState
	FromProto(msg *otg.Ospfv3LsaState) (Ospfv3LsaState, error)
	// FromPbText unmarshals Ospfv3LsaState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3LsaState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3LsaState from JSON text
	FromJson(value string) error
}

func (obj *ospfv3LsaState) Marshal() marshalOspfv3LsaState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3LsaState{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3LsaState) Unmarshal() unMarshalOspfv3LsaState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3LsaState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3LsaState) ToProto() (*otg.Ospfv3LsaState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3LsaState) FromProto(msg *otg.Ospfv3LsaState) (Ospfv3LsaState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3LsaState) ToPbText() (string, error) {
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

func (m *unMarshalospfv3LsaState) FromPbText(value string) error {
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

func (m *marshalospfv3LsaState) ToYaml() (string, error) {
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

func (m *unMarshalospfv3LsaState) FromYaml(value string) error {
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

func (m *marshalospfv3LsaState) ToJson() (string, error) {
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

func (m *unMarshalospfv3LsaState) FromJson(value string) error {
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

func (obj *ospfv3LsaState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3LsaState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3LsaState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3LsaState) Clone() (Ospfv3LsaState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3LsaState()
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

func (obj *ospfv3LsaState) setNil() {
	obj.routerLsasHolder = nil
	obj.networkLsasHolder = nil
	obj.interAreaPrefixLsasHolder = nil
	obj.interAreaRouterLsasHolder = nil
	obj.externalAsLsasHolder = nil
	obj.nssaLsasHolder = nil
	obj.linkLsasHolder = nil
	obj.intraAreaPrefixLsasHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3LsaState is the result of OSPFv3 LSA information that are retrieved.
type Ospfv3LsaState interface {
	Validation
	// msg marshals Ospfv3LsaState to protobuf object *otg.Ospfv3LsaState
	// and doesn't set defaults
	msg() *otg.Ospfv3LsaState
	// setMsg unmarshals Ospfv3LsaState from protobuf object *otg.Ospfv3LsaState
	// and doesn't set defaults
	setMsg(*otg.Ospfv3LsaState) Ospfv3LsaState
	// provides marshal interface
	Marshal() marshalOspfv3LsaState
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3LsaState
	// validate validates Ospfv3LsaState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3LsaState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterName returns string, set in Ospfv3LsaState.
	RouterName() string
	// SetRouterName assigns string provided by user to Ospfv3LsaState
	SetRouterName(value string) Ospfv3LsaState
	// HasRouterName checks if RouterName has been set in Ospfv3LsaState
	HasRouterName() bool
	// RouterLsas returns Ospfv3LsaStateOspfv3RouterLsaIterIter, set in Ospfv3LsaState
	RouterLsas() Ospfv3LsaStateOspfv3RouterLsaIter
	// NetworkLsas returns Ospfv3LsaStateOspfv3NetworkLsaIterIter, set in Ospfv3LsaState
	NetworkLsas() Ospfv3LsaStateOspfv3NetworkLsaIter
	// InterAreaPrefixLsas returns Ospfv3LsaStateOspfv3InterAreaPrefixLsaIterIter, set in Ospfv3LsaState
	InterAreaPrefixLsas() Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	// InterAreaRouterLsas returns Ospfv3LsaStateOspfv3InterAreaRouterLsaIterIter, set in Ospfv3LsaState
	InterAreaRouterLsas() Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	// ExternalAsLsas returns Ospfv3LsaStateOspfv3ExternalAsLsaIterIter, set in Ospfv3LsaState
	ExternalAsLsas() Ospfv3LsaStateOspfv3ExternalAsLsaIter
	// NssaLsas returns Ospfv3LsaStateOspfv3NssaLsaIterIter, set in Ospfv3LsaState
	NssaLsas() Ospfv3LsaStateOspfv3NssaLsaIter
	// LinkLsas returns Ospfv3LsaStateOspfv3LinkLsaIterIter, set in Ospfv3LsaState
	LinkLsas() Ospfv3LsaStateOspfv3LinkLsaIter
	// IntraAreaPrefixLsas returns Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIterIter, set in Ospfv3LsaState
	IntraAreaPrefixLsas() Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
	setNil()
}

// The name of the OSPFv3 Router that learned the LSA information.
// RouterName returns a string
func (obj *ospfv3LsaState) RouterName() string {

	return *obj.obj.RouterName

}

// The name of the OSPFv3 Router that learned the LSA information.
// RouterName returns a string
func (obj *ospfv3LsaState) HasRouterName() bool {
	return obj.obj.RouterName != nil
}

// The name of the OSPFv3 Router that learned the LSA information.
// SetRouterName sets the string value in the Ospfv3LsaState object
func (obj *ospfv3LsaState) SetRouterName(value string) Ospfv3LsaState {

	obj.obj.RouterName = &value
	return obj
}

// One or more OSPFv3 Router LSA - Type 1.
// RouterLsas returns a []Ospfv3RouterLsa
func (obj *ospfv3LsaState) RouterLsas() Ospfv3LsaStateOspfv3RouterLsaIter {
	if len(obj.obj.RouterLsas) == 0 {
		obj.obj.RouterLsas = []*otg.Ospfv3RouterLsa{}
	}
	if obj.routerLsasHolder == nil {
		obj.routerLsasHolder = newOspfv3LsaStateOspfv3RouterLsaIter(&obj.obj.RouterLsas).setMsg(obj)
	}
	return obj.routerLsasHolder
}

type ospfv3LsaStateOspfv3RouterLsaIter struct {
	obj                  *ospfv3LsaState
	ospfv3RouterLsaSlice []Ospfv3RouterLsa
	fieldPtr             *[]*otg.Ospfv3RouterLsa
}

func newOspfv3LsaStateOspfv3RouterLsaIter(ptr *[]*otg.Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter {
	return &ospfv3LsaStateOspfv3RouterLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3RouterLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3RouterLsaIter
	Items() []Ospfv3RouterLsa
	Add() Ospfv3RouterLsa
	Append(items ...Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter
	Set(index int, newObj Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter
	Clear() Ospfv3LsaStateOspfv3RouterLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3RouterLsaIter
	appendHolderSlice(item Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter
}

func (obj *ospfv3LsaStateOspfv3RouterLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3RouterLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3RouterLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3RouterLsaIter) Items() []Ospfv3RouterLsa {
	return obj.ospfv3RouterLsaSlice
}

func (obj *ospfv3LsaStateOspfv3RouterLsaIter) Add() Ospfv3RouterLsa {
	newObj := &otg.Ospfv3RouterLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3RouterLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3RouterLsaSlice = append(obj.ospfv3RouterLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3RouterLsaIter) Append(items ...Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3RouterLsaSlice = append(obj.ospfv3RouterLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3RouterLsaIter) Set(index int, newObj Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3RouterLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3RouterLsaIter) Clear() Ospfv3LsaStateOspfv3RouterLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3RouterLsa{}
		obj.ospfv3RouterLsaSlice = []Ospfv3RouterLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3RouterLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3RouterLsaIter {
	if len(obj.ospfv3RouterLsaSlice) > 0 {
		obj.ospfv3RouterLsaSlice = []Ospfv3RouterLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3RouterLsaIter) appendHolderSlice(item Ospfv3RouterLsa) Ospfv3LsaStateOspfv3RouterLsaIter {
	obj.ospfv3RouterLsaSlice = append(obj.ospfv3RouterLsaSlice, item)
	return obj
}

// One or more OSPFv3 Network LSA - Type 2.
// NetworkLsas returns a []Ospfv3NetworkLsa
func (obj *ospfv3LsaState) NetworkLsas() Ospfv3LsaStateOspfv3NetworkLsaIter {
	if len(obj.obj.NetworkLsas) == 0 {
		obj.obj.NetworkLsas = []*otg.Ospfv3NetworkLsa{}
	}
	if obj.networkLsasHolder == nil {
		obj.networkLsasHolder = newOspfv3LsaStateOspfv3NetworkLsaIter(&obj.obj.NetworkLsas).setMsg(obj)
	}
	return obj.networkLsasHolder
}

type ospfv3LsaStateOspfv3NetworkLsaIter struct {
	obj                   *ospfv3LsaState
	ospfv3NetworkLsaSlice []Ospfv3NetworkLsa
	fieldPtr              *[]*otg.Ospfv3NetworkLsa
}

func newOspfv3LsaStateOspfv3NetworkLsaIter(ptr *[]*otg.Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter {
	return &ospfv3LsaStateOspfv3NetworkLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3NetworkLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3NetworkLsaIter
	Items() []Ospfv3NetworkLsa
	Add() Ospfv3NetworkLsa
	Append(items ...Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter
	Set(index int, newObj Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter
	Clear() Ospfv3LsaStateOspfv3NetworkLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3NetworkLsaIter
	appendHolderSlice(item Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter
}

func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3NetworkLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3NetworkLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) Items() []Ospfv3NetworkLsa {
	return obj.ospfv3NetworkLsaSlice
}

func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) Add() Ospfv3NetworkLsa {
	newObj := &otg.Ospfv3NetworkLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3NetworkLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3NetworkLsaSlice = append(obj.ospfv3NetworkLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) Append(items ...Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3NetworkLsaSlice = append(obj.ospfv3NetworkLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) Set(index int, newObj Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3NetworkLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) Clear() Ospfv3LsaStateOspfv3NetworkLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3NetworkLsa{}
		obj.ospfv3NetworkLsaSlice = []Ospfv3NetworkLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3NetworkLsaIter {
	if len(obj.ospfv3NetworkLsaSlice) > 0 {
		obj.ospfv3NetworkLsaSlice = []Ospfv3NetworkLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3NetworkLsaIter) appendHolderSlice(item Ospfv3NetworkLsa) Ospfv3LsaStateOspfv3NetworkLsaIter {
	obj.ospfv3NetworkLsaSlice = append(obj.ospfv3NetworkLsaSlice, item)
	return obj
}

// One or more OSPFv3 Inter-Area-Prefix LSA - Type 3.
// InterAreaPrefixLsas returns a []Ospfv3InterAreaPrefixLsa
func (obj *ospfv3LsaState) InterAreaPrefixLsas() Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	if len(obj.obj.InterAreaPrefixLsas) == 0 {
		obj.obj.InterAreaPrefixLsas = []*otg.Ospfv3InterAreaPrefixLsa{}
	}
	if obj.interAreaPrefixLsasHolder == nil {
		obj.interAreaPrefixLsasHolder = newOspfv3LsaStateOspfv3InterAreaPrefixLsaIter(&obj.obj.InterAreaPrefixLsas).setMsg(obj)
	}
	return obj.interAreaPrefixLsasHolder
}

type ospfv3LsaStateOspfv3InterAreaPrefixLsaIter struct {
	obj                           *ospfv3LsaState
	ospfv3InterAreaPrefixLsaSlice []Ospfv3InterAreaPrefixLsa
	fieldPtr                      *[]*otg.Ospfv3InterAreaPrefixLsa
}

func newOspfv3LsaStateOspfv3InterAreaPrefixLsaIter(ptr *[]*otg.Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	return &ospfv3LsaStateOspfv3InterAreaPrefixLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	Items() []Ospfv3InterAreaPrefixLsa
	Add() Ospfv3InterAreaPrefixLsa
	Append(items ...Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	Set(index int, newObj Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	Clear() Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
	appendHolderSlice(item Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter
}

func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3InterAreaPrefixLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) Items() []Ospfv3InterAreaPrefixLsa {
	return obj.ospfv3InterAreaPrefixLsaSlice
}

func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) Add() Ospfv3InterAreaPrefixLsa {
	newObj := &otg.Ospfv3InterAreaPrefixLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3InterAreaPrefixLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3InterAreaPrefixLsaSlice = append(obj.ospfv3InterAreaPrefixLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) Append(items ...Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3InterAreaPrefixLsaSlice = append(obj.ospfv3InterAreaPrefixLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) Set(index int, newObj Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3InterAreaPrefixLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) Clear() Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3InterAreaPrefixLsa{}
		obj.ospfv3InterAreaPrefixLsaSlice = []Ospfv3InterAreaPrefixLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	if len(obj.ospfv3InterAreaPrefixLsaSlice) > 0 {
		obj.ospfv3InterAreaPrefixLsaSlice = []Ospfv3InterAreaPrefixLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3InterAreaPrefixLsaIter) appendHolderSlice(item Ospfv3InterAreaPrefixLsa) Ospfv3LsaStateOspfv3InterAreaPrefixLsaIter {
	obj.ospfv3InterAreaPrefixLsaSlice = append(obj.ospfv3InterAreaPrefixLsaSlice, item)
	return obj
}

// One or more OSPFv3 Inter-Area-Router LSA - Type 4.
// InterAreaRouterLsas returns a []Ospfv3InterAreaRouterLsa
func (obj *ospfv3LsaState) InterAreaRouterLsas() Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	if len(obj.obj.InterAreaRouterLsas) == 0 {
		obj.obj.InterAreaRouterLsas = []*otg.Ospfv3InterAreaRouterLsa{}
	}
	if obj.interAreaRouterLsasHolder == nil {
		obj.interAreaRouterLsasHolder = newOspfv3LsaStateOspfv3InterAreaRouterLsaIter(&obj.obj.InterAreaRouterLsas).setMsg(obj)
	}
	return obj.interAreaRouterLsasHolder
}

type ospfv3LsaStateOspfv3InterAreaRouterLsaIter struct {
	obj                           *ospfv3LsaState
	ospfv3InterAreaRouterLsaSlice []Ospfv3InterAreaRouterLsa
	fieldPtr                      *[]*otg.Ospfv3InterAreaRouterLsa
}

func newOspfv3LsaStateOspfv3InterAreaRouterLsaIter(ptr *[]*otg.Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	return &ospfv3LsaStateOspfv3InterAreaRouterLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3InterAreaRouterLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	Items() []Ospfv3InterAreaRouterLsa
	Add() Ospfv3InterAreaRouterLsa
	Append(items ...Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	Set(index int, newObj Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	Clear() Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
	appendHolderSlice(item Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter
}

func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3InterAreaRouterLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) Items() []Ospfv3InterAreaRouterLsa {
	return obj.ospfv3InterAreaRouterLsaSlice
}

func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) Add() Ospfv3InterAreaRouterLsa {
	newObj := &otg.Ospfv3InterAreaRouterLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3InterAreaRouterLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3InterAreaRouterLsaSlice = append(obj.ospfv3InterAreaRouterLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) Append(items ...Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3InterAreaRouterLsaSlice = append(obj.ospfv3InterAreaRouterLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) Set(index int, newObj Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3InterAreaRouterLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) Clear() Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3InterAreaRouterLsa{}
		obj.ospfv3InterAreaRouterLsaSlice = []Ospfv3InterAreaRouterLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	if len(obj.ospfv3InterAreaRouterLsaSlice) > 0 {
		obj.ospfv3InterAreaRouterLsaSlice = []Ospfv3InterAreaRouterLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3InterAreaRouterLsaIter) appendHolderSlice(item Ospfv3InterAreaRouterLsa) Ospfv3LsaStateOspfv3InterAreaRouterLsaIter {
	obj.ospfv3InterAreaRouterLsaSlice = append(obj.ospfv3InterAreaRouterLsaSlice, item)
	return obj
}

// OSPFv3 AS-External LSA - Type 5.
// ExternalAsLsas returns a []Ospfv3ExternalAsLsa
func (obj *ospfv3LsaState) ExternalAsLsas() Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	if len(obj.obj.ExternalAsLsas) == 0 {
		obj.obj.ExternalAsLsas = []*otg.Ospfv3ExternalAsLsa{}
	}
	if obj.externalAsLsasHolder == nil {
		obj.externalAsLsasHolder = newOspfv3LsaStateOspfv3ExternalAsLsaIter(&obj.obj.ExternalAsLsas).setMsg(obj)
	}
	return obj.externalAsLsasHolder
}

type ospfv3LsaStateOspfv3ExternalAsLsaIter struct {
	obj                      *ospfv3LsaState
	ospfv3ExternalAsLsaSlice []Ospfv3ExternalAsLsa
	fieldPtr                 *[]*otg.Ospfv3ExternalAsLsa
}

func newOspfv3LsaStateOspfv3ExternalAsLsaIter(ptr *[]*otg.Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	return &ospfv3LsaStateOspfv3ExternalAsLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3ExternalAsLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3ExternalAsLsaIter
	Items() []Ospfv3ExternalAsLsa
	Add() Ospfv3ExternalAsLsa
	Append(items ...Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter
	Set(index int, newObj Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter
	Clear() Ospfv3LsaStateOspfv3ExternalAsLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3ExternalAsLsaIter
	appendHolderSlice(item Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter
}

func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3ExternalAsLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) Items() []Ospfv3ExternalAsLsa {
	return obj.ospfv3ExternalAsLsaSlice
}

func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) Add() Ospfv3ExternalAsLsa {
	newObj := &otg.Ospfv3ExternalAsLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3ExternalAsLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3ExternalAsLsaSlice = append(obj.ospfv3ExternalAsLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) Append(items ...Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3ExternalAsLsaSlice = append(obj.ospfv3ExternalAsLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) Set(index int, newObj Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3ExternalAsLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) Clear() Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3ExternalAsLsa{}
		obj.ospfv3ExternalAsLsaSlice = []Ospfv3ExternalAsLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	if len(obj.ospfv3ExternalAsLsaSlice) > 0 {
		obj.ospfv3ExternalAsLsaSlice = []Ospfv3ExternalAsLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3ExternalAsLsaIter) appendHolderSlice(item Ospfv3ExternalAsLsa) Ospfv3LsaStateOspfv3ExternalAsLsaIter {
	obj.ospfv3ExternalAsLsaSlice = append(obj.ospfv3ExternalAsLsaSlice, item)
	return obj
}

// One or more OSPFv3 NSSA LSA - Type 7.
// NssaLsas returns a []Ospfv3NssaLsa
func (obj *ospfv3LsaState) NssaLsas() Ospfv3LsaStateOspfv3NssaLsaIter {
	if len(obj.obj.NssaLsas) == 0 {
		obj.obj.NssaLsas = []*otg.Ospfv3NssaLsa{}
	}
	if obj.nssaLsasHolder == nil {
		obj.nssaLsasHolder = newOspfv3LsaStateOspfv3NssaLsaIter(&obj.obj.NssaLsas).setMsg(obj)
	}
	return obj.nssaLsasHolder
}

type ospfv3LsaStateOspfv3NssaLsaIter struct {
	obj                *ospfv3LsaState
	ospfv3NssaLsaSlice []Ospfv3NssaLsa
	fieldPtr           *[]*otg.Ospfv3NssaLsa
}

func newOspfv3LsaStateOspfv3NssaLsaIter(ptr *[]*otg.Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter {
	return &ospfv3LsaStateOspfv3NssaLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3NssaLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3NssaLsaIter
	Items() []Ospfv3NssaLsa
	Add() Ospfv3NssaLsa
	Append(items ...Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter
	Set(index int, newObj Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter
	Clear() Ospfv3LsaStateOspfv3NssaLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3NssaLsaIter
	appendHolderSlice(item Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter
}

func (obj *ospfv3LsaStateOspfv3NssaLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3NssaLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3NssaLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3NssaLsaIter) Items() []Ospfv3NssaLsa {
	return obj.ospfv3NssaLsaSlice
}

func (obj *ospfv3LsaStateOspfv3NssaLsaIter) Add() Ospfv3NssaLsa {
	newObj := &otg.Ospfv3NssaLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3NssaLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3NssaLsaSlice = append(obj.ospfv3NssaLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3NssaLsaIter) Append(items ...Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3NssaLsaSlice = append(obj.ospfv3NssaLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3NssaLsaIter) Set(index int, newObj Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3NssaLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3NssaLsaIter) Clear() Ospfv3LsaStateOspfv3NssaLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3NssaLsa{}
		obj.ospfv3NssaLsaSlice = []Ospfv3NssaLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3NssaLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3NssaLsaIter {
	if len(obj.ospfv3NssaLsaSlice) > 0 {
		obj.ospfv3NssaLsaSlice = []Ospfv3NssaLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3NssaLsaIter) appendHolderSlice(item Ospfv3NssaLsa) Ospfv3LsaStateOspfv3NssaLsaIter {
	obj.ospfv3NssaLsaSlice = append(obj.ospfv3NssaLsaSlice, item)
	return obj
}

// One or more OSPFv3 Link LSA - Type 8.
// LinkLsas returns a []Ospfv3LinkLsa
func (obj *ospfv3LsaState) LinkLsas() Ospfv3LsaStateOspfv3LinkLsaIter {
	if len(obj.obj.LinkLsas) == 0 {
		obj.obj.LinkLsas = []*otg.Ospfv3LinkLsa{}
	}
	if obj.linkLsasHolder == nil {
		obj.linkLsasHolder = newOspfv3LsaStateOspfv3LinkLsaIter(&obj.obj.LinkLsas).setMsg(obj)
	}
	return obj.linkLsasHolder
}

type ospfv3LsaStateOspfv3LinkLsaIter struct {
	obj                *ospfv3LsaState
	ospfv3LinkLsaSlice []Ospfv3LinkLsa
	fieldPtr           *[]*otg.Ospfv3LinkLsa
}

func newOspfv3LsaStateOspfv3LinkLsaIter(ptr *[]*otg.Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter {
	return &ospfv3LsaStateOspfv3LinkLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3LinkLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3LinkLsaIter
	Items() []Ospfv3LinkLsa
	Add() Ospfv3LinkLsa
	Append(items ...Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter
	Set(index int, newObj Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter
	Clear() Ospfv3LsaStateOspfv3LinkLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3LinkLsaIter
	appendHolderSlice(item Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter
}

func (obj *ospfv3LsaStateOspfv3LinkLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3LinkLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3LinkLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3LinkLsaIter) Items() []Ospfv3LinkLsa {
	return obj.ospfv3LinkLsaSlice
}

func (obj *ospfv3LsaStateOspfv3LinkLsaIter) Add() Ospfv3LinkLsa {
	newObj := &otg.Ospfv3LinkLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3LinkLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3LinkLsaSlice = append(obj.ospfv3LinkLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3LinkLsaIter) Append(items ...Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3LinkLsaSlice = append(obj.ospfv3LinkLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3LinkLsaIter) Set(index int, newObj Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3LinkLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3LinkLsaIter) Clear() Ospfv3LsaStateOspfv3LinkLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3LinkLsa{}
		obj.ospfv3LinkLsaSlice = []Ospfv3LinkLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3LinkLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3LinkLsaIter {
	if len(obj.ospfv3LinkLsaSlice) > 0 {
		obj.ospfv3LinkLsaSlice = []Ospfv3LinkLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3LinkLsaIter) appendHolderSlice(item Ospfv3LinkLsa) Ospfv3LsaStateOspfv3LinkLsaIter {
	obj.ospfv3LinkLsaSlice = append(obj.ospfv3LinkLsaSlice, item)
	return obj
}

// One or more OSPFv3 Intra-Area-Prefix LSA - Type 9.
// IntraAreaPrefixLsas returns a []Ospfv3IntraAreaPrefixLsa
func (obj *ospfv3LsaState) IntraAreaPrefixLsas() Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	if len(obj.obj.IntraAreaPrefixLsas) == 0 {
		obj.obj.IntraAreaPrefixLsas = []*otg.Ospfv3IntraAreaPrefixLsa{}
	}
	if obj.intraAreaPrefixLsasHolder == nil {
		obj.intraAreaPrefixLsasHolder = newOspfv3LsaStateOspfv3IntraAreaPrefixLsaIter(&obj.obj.IntraAreaPrefixLsas).setMsg(obj)
	}
	return obj.intraAreaPrefixLsasHolder
}

type ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter struct {
	obj                           *ospfv3LsaState
	ospfv3IntraAreaPrefixLsaSlice []Ospfv3IntraAreaPrefixLsa
	fieldPtr                      *[]*otg.Ospfv3IntraAreaPrefixLsa
}

func newOspfv3LsaStateOspfv3IntraAreaPrefixLsaIter(ptr *[]*otg.Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	return &ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter{fieldPtr: ptr}
}

type Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter interface {
	setMsg(*ospfv3LsaState) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
	Items() []Ospfv3IntraAreaPrefixLsa
	Add() Ospfv3IntraAreaPrefixLsa
	Append(items ...Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
	Set(index int, newObj Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
	Clear() Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
	clearHolderSlice() Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
	appendHolderSlice(item Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter
}

func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) setMsg(msg *ospfv3LsaState) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3IntraAreaPrefixLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) Items() []Ospfv3IntraAreaPrefixLsa {
	return obj.ospfv3IntraAreaPrefixLsaSlice
}

func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) Add() Ospfv3IntraAreaPrefixLsa {
	newObj := &otg.Ospfv3IntraAreaPrefixLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3IntraAreaPrefixLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3IntraAreaPrefixLsaSlice = append(obj.ospfv3IntraAreaPrefixLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) Append(items ...Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3IntraAreaPrefixLsaSlice = append(obj.ospfv3IntraAreaPrefixLsaSlice, item)
	}
	return obj
}

func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) Set(index int, newObj Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3IntraAreaPrefixLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) Clear() Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3IntraAreaPrefixLsa{}
		obj.ospfv3IntraAreaPrefixLsaSlice = []Ospfv3IntraAreaPrefixLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) clearHolderSlice() Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	if len(obj.ospfv3IntraAreaPrefixLsaSlice) > 0 {
		obj.ospfv3IntraAreaPrefixLsaSlice = []Ospfv3IntraAreaPrefixLsa{}
	}
	return obj
}
func (obj *ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter) appendHolderSlice(item Ospfv3IntraAreaPrefixLsa) Ospfv3LsaStateOspfv3IntraAreaPrefixLsaIter {
	obj.ospfv3IntraAreaPrefixLsaSlice = append(obj.ospfv3IntraAreaPrefixLsaSlice, item)
	return obj
}

func (obj *ospfv3LsaState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.RouterLsas) != 0 {

		if set_default {
			obj.RouterLsas().clearHolderSlice()
			for _, item := range obj.obj.RouterLsas {
				obj.RouterLsas().appendHolderSlice(&ospfv3RouterLsa{obj: item})
			}
		}
		for _, item := range obj.RouterLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.NetworkLsas) != 0 {

		if set_default {
			obj.NetworkLsas().clearHolderSlice()
			for _, item := range obj.obj.NetworkLsas {
				obj.NetworkLsas().appendHolderSlice(&ospfv3NetworkLsa{obj: item})
			}
		}
		for _, item := range obj.NetworkLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.InterAreaPrefixLsas) != 0 {

		if set_default {
			obj.InterAreaPrefixLsas().clearHolderSlice()
			for _, item := range obj.obj.InterAreaPrefixLsas {
				obj.InterAreaPrefixLsas().appendHolderSlice(&ospfv3InterAreaPrefixLsa{obj: item})
			}
		}
		for _, item := range obj.InterAreaPrefixLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.InterAreaRouterLsas) != 0 {

		if set_default {
			obj.InterAreaRouterLsas().clearHolderSlice()
			for _, item := range obj.obj.InterAreaRouterLsas {
				obj.InterAreaRouterLsas().appendHolderSlice(&ospfv3InterAreaRouterLsa{obj: item})
			}
		}
		for _, item := range obj.InterAreaRouterLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExternalAsLsas) != 0 {

		if set_default {
			obj.ExternalAsLsas().clearHolderSlice()
			for _, item := range obj.obj.ExternalAsLsas {
				obj.ExternalAsLsas().appendHolderSlice(&ospfv3ExternalAsLsa{obj: item})
			}
		}
		for _, item := range obj.ExternalAsLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.NssaLsas) != 0 {

		if set_default {
			obj.NssaLsas().clearHolderSlice()
			for _, item := range obj.obj.NssaLsas {
				obj.NssaLsas().appendHolderSlice(&ospfv3NssaLsa{obj: item})
			}
		}
		for _, item := range obj.NssaLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.LinkLsas) != 0 {

		if set_default {
			obj.LinkLsas().clearHolderSlice()
			for _, item := range obj.obj.LinkLsas {
				obj.LinkLsas().appendHolderSlice(&ospfv3LinkLsa{obj: item})
			}
		}
		for _, item := range obj.LinkLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.IntraAreaPrefixLsas) != 0 {

		if set_default {
			obj.IntraAreaPrefixLsas().clearHolderSlice()
			for _, item := range obj.obj.IntraAreaPrefixLsas {
				obj.IntraAreaPrefixLsas().appendHolderSlice(&ospfv3IntraAreaPrefixLsa{obj: item})
			}
		}
		for _, item := range obj.IntraAreaPrefixLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv3LsaState) setDefault() {

}
