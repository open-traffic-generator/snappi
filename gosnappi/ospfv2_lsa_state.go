package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaState *****
type ospfv2LsaState struct {
	validation
	obj                      *otg.Ospfv2LsaState
	marshaller               marshalOspfv2LsaState
	unMarshaller             unMarshalOspfv2LsaState
	routerLsasHolder         Ospfv2LsaStateOspfv2RouterLsaIter
	networkLsasHolder        Ospfv2LsaStateOspfv2NetworkLsaIter
	networkSummaryLsasHolder Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	summaryAsLsasHolder      Ospfv2LsaStateOspfv2SummaryAsLsaIter
	externalAsLsasHolder     Ospfv2LsaStateOspfv2ExternalAsLsaIter
	nssaLsasHolder           Ospfv2LsaStateOspfv2NssaLsaIter
	opaqueLsasHolder         Ospfv2LsaStateOspfv2OpaqueLsaIter
}

func NewOspfv2LsaState() Ospfv2LsaState {
	obj := ospfv2LsaState{obj: &otg.Ospfv2LsaState{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaState) msg() *otg.Ospfv2LsaState {
	return obj.obj
}

func (obj *ospfv2LsaState) setMsg(msg *otg.Ospfv2LsaState) Ospfv2LsaState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaState struct {
	obj *ospfv2LsaState
}

type marshalOspfv2LsaState interface {
	// ToProto marshals Ospfv2LsaState to protobuf object *otg.Ospfv2LsaState
	ToProto() (*otg.Ospfv2LsaState, error)
	// ToPbText marshals Ospfv2LsaState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaState to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaState to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaState struct {
	obj *ospfv2LsaState
}

type unMarshalOspfv2LsaState interface {
	// FromProto unmarshals Ospfv2LsaState from protobuf object *otg.Ospfv2LsaState
	FromProto(msg *otg.Ospfv2LsaState) (Ospfv2LsaState, error)
	// FromPbText unmarshals Ospfv2LsaState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaState from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaState) Marshal() marshalOspfv2LsaState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaState{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaState) Unmarshal() unMarshalOspfv2LsaState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaState) ToProto() (*otg.Ospfv2LsaState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaState) FromProto(msg *otg.Ospfv2LsaState) (Ospfv2LsaState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaState) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaState) FromPbText(value string) error {
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

func (m *marshalospfv2LsaState) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaState) FromYaml(value string) error {
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

func (m *marshalospfv2LsaState) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaState) FromJson(value string) error {
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

func (obj *ospfv2LsaState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaState) Clone() (Ospfv2LsaState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaState()
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

func (obj *ospfv2LsaState) setNil() {
	obj.routerLsasHolder = nil
	obj.networkLsasHolder = nil
	obj.networkSummaryLsasHolder = nil
	obj.summaryAsLsasHolder = nil
	obj.externalAsLsasHolder = nil
	obj.nssaLsasHolder = nil
	obj.opaqueLsasHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2LsaState is the result of OSPFv2 LSA information that are retrieved.
type Ospfv2LsaState interface {
	Validation
	// msg marshals Ospfv2LsaState to protobuf object *otg.Ospfv2LsaState
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaState
	// setMsg unmarshals Ospfv2LsaState from protobuf object *otg.Ospfv2LsaState
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaState) Ospfv2LsaState
	// provides marshal interface
	Marshal() marshalOspfv2LsaState
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaState
	// validate validates Ospfv2LsaState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterName returns string, set in Ospfv2LsaState.
	RouterName() string
	// SetRouterName assigns string provided by user to Ospfv2LsaState
	SetRouterName(value string) Ospfv2LsaState
	// HasRouterName checks if RouterName has been set in Ospfv2LsaState
	HasRouterName() bool
	// RouterLsas returns Ospfv2LsaStateOspfv2RouterLsaIterIter, set in Ospfv2LsaState
	RouterLsas() Ospfv2LsaStateOspfv2RouterLsaIter
	// NetworkLsas returns Ospfv2LsaStateOspfv2NetworkLsaIterIter, set in Ospfv2LsaState
	NetworkLsas() Ospfv2LsaStateOspfv2NetworkLsaIter
	// NetworkSummaryLsas returns Ospfv2LsaStateOspfv2NetworkSummaryLsaIterIter, set in Ospfv2LsaState
	NetworkSummaryLsas() Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	// SummaryAsLsas returns Ospfv2LsaStateOspfv2SummaryAsLsaIterIter, set in Ospfv2LsaState
	SummaryAsLsas() Ospfv2LsaStateOspfv2SummaryAsLsaIter
	// ExternalAsLsas returns Ospfv2LsaStateOspfv2ExternalAsLsaIterIter, set in Ospfv2LsaState
	ExternalAsLsas() Ospfv2LsaStateOspfv2ExternalAsLsaIter
	// NssaLsas returns Ospfv2LsaStateOspfv2NssaLsaIterIter, set in Ospfv2LsaState
	NssaLsas() Ospfv2LsaStateOspfv2NssaLsaIter
	// OpaqueLsas returns Ospfv2LsaStateOspfv2OpaqueLsaIterIter, set in Ospfv2LsaState
	OpaqueLsas() Ospfv2LsaStateOspfv2OpaqueLsaIter
	setNil()
}

// The name of the OSPFv2 Router that learned the LSA information.
// RouterName returns a string
func (obj *ospfv2LsaState) RouterName() string {

	return *obj.obj.RouterName

}

// The name of the OSPFv2 Router that learned the LSA information.
// RouterName returns a string
func (obj *ospfv2LsaState) HasRouterName() bool {
	return obj.obj.RouterName != nil
}

// The name of the OSPFv2 Router that learned the LSA information.
// SetRouterName sets the string value in the Ospfv2LsaState object
func (obj *ospfv2LsaState) SetRouterName(value string) Ospfv2LsaState {

	obj.obj.RouterName = &value
	return obj
}

// One or more OSPFv2 Router-LSA - Type 1.
// RouterLsas returns a []Ospfv2RouterLsa
func (obj *ospfv2LsaState) RouterLsas() Ospfv2LsaStateOspfv2RouterLsaIter {
	if len(obj.obj.RouterLsas) == 0 {
		obj.obj.RouterLsas = []*otg.Ospfv2RouterLsa{}
	}
	if obj.routerLsasHolder == nil {
		obj.routerLsasHolder = newOspfv2LsaStateOspfv2RouterLsaIter(&obj.obj.RouterLsas).setMsg(obj)
	}
	return obj.routerLsasHolder
}

type ospfv2LsaStateOspfv2RouterLsaIter struct {
	obj                  *ospfv2LsaState
	ospfv2RouterLsaSlice []Ospfv2RouterLsa
	fieldPtr             *[]*otg.Ospfv2RouterLsa
}

func newOspfv2LsaStateOspfv2RouterLsaIter(ptr *[]*otg.Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter {
	return &ospfv2LsaStateOspfv2RouterLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2RouterLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2RouterLsaIter
	Items() []Ospfv2RouterLsa
	Add() Ospfv2RouterLsa
	Append(items ...Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter
	Set(index int, newObj Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter
	Clear() Ospfv2LsaStateOspfv2RouterLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2RouterLsaIter
	appendHolderSlice(item Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter
}

func (obj *ospfv2LsaStateOspfv2RouterLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2RouterLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2RouterLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2RouterLsaIter) Items() []Ospfv2RouterLsa {
	return obj.ospfv2RouterLsaSlice
}

func (obj *ospfv2LsaStateOspfv2RouterLsaIter) Add() Ospfv2RouterLsa {
	newObj := &otg.Ospfv2RouterLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2RouterLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2RouterLsaSlice = append(obj.ospfv2RouterLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2RouterLsaIter) Append(items ...Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2RouterLsaSlice = append(obj.ospfv2RouterLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2RouterLsaIter) Set(index int, newObj Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2RouterLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2RouterLsaIter) Clear() Ospfv2LsaStateOspfv2RouterLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2RouterLsa{}
		obj.ospfv2RouterLsaSlice = []Ospfv2RouterLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2RouterLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2RouterLsaIter {
	if len(obj.ospfv2RouterLsaSlice) > 0 {
		obj.ospfv2RouterLsaSlice = []Ospfv2RouterLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2RouterLsaIter) appendHolderSlice(item Ospfv2RouterLsa) Ospfv2LsaStateOspfv2RouterLsaIter {
	obj.ospfv2RouterLsaSlice = append(obj.ospfv2RouterLsaSlice, item)
	return obj
}

// One or more OSPFv2 Network-LSA - Type 2.
// NetworkLsas returns a []Ospfv2NetworkLsa
func (obj *ospfv2LsaState) NetworkLsas() Ospfv2LsaStateOspfv2NetworkLsaIter {
	if len(obj.obj.NetworkLsas) == 0 {
		obj.obj.NetworkLsas = []*otg.Ospfv2NetworkLsa{}
	}
	if obj.networkLsasHolder == nil {
		obj.networkLsasHolder = newOspfv2LsaStateOspfv2NetworkLsaIter(&obj.obj.NetworkLsas).setMsg(obj)
	}
	return obj.networkLsasHolder
}

type ospfv2LsaStateOspfv2NetworkLsaIter struct {
	obj                   *ospfv2LsaState
	ospfv2NetworkLsaSlice []Ospfv2NetworkLsa
	fieldPtr              *[]*otg.Ospfv2NetworkLsa
}

func newOspfv2LsaStateOspfv2NetworkLsaIter(ptr *[]*otg.Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter {
	return &ospfv2LsaStateOspfv2NetworkLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2NetworkLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2NetworkLsaIter
	Items() []Ospfv2NetworkLsa
	Add() Ospfv2NetworkLsa
	Append(items ...Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter
	Set(index int, newObj Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter
	Clear() Ospfv2LsaStateOspfv2NetworkLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2NetworkLsaIter
	appendHolderSlice(item Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter
}

func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2NetworkLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2NetworkLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) Items() []Ospfv2NetworkLsa {
	return obj.ospfv2NetworkLsaSlice
}

func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) Add() Ospfv2NetworkLsa {
	newObj := &otg.Ospfv2NetworkLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2NetworkLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2NetworkLsaSlice = append(obj.ospfv2NetworkLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) Append(items ...Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2NetworkLsaSlice = append(obj.ospfv2NetworkLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) Set(index int, newObj Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2NetworkLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) Clear() Ospfv2LsaStateOspfv2NetworkLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2NetworkLsa{}
		obj.ospfv2NetworkLsaSlice = []Ospfv2NetworkLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2NetworkLsaIter {
	if len(obj.ospfv2NetworkLsaSlice) > 0 {
		obj.ospfv2NetworkLsaSlice = []Ospfv2NetworkLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2NetworkLsaIter) appendHolderSlice(item Ospfv2NetworkLsa) Ospfv2LsaStateOspfv2NetworkLsaIter {
	obj.ospfv2NetworkLsaSlice = append(obj.ospfv2NetworkLsaSlice, item)
	return obj
}

// One or more OSPFv2 Network summary LSA - Type 3.
// NetworkSummaryLsas returns a []Ospfv2NetworkSummaryLsa
func (obj *ospfv2LsaState) NetworkSummaryLsas() Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	if len(obj.obj.NetworkSummaryLsas) == 0 {
		obj.obj.NetworkSummaryLsas = []*otg.Ospfv2NetworkSummaryLsa{}
	}
	if obj.networkSummaryLsasHolder == nil {
		obj.networkSummaryLsasHolder = newOspfv2LsaStateOspfv2NetworkSummaryLsaIter(&obj.obj.NetworkSummaryLsas).setMsg(obj)
	}
	return obj.networkSummaryLsasHolder
}

type ospfv2LsaStateOspfv2NetworkSummaryLsaIter struct {
	obj                          *ospfv2LsaState
	ospfv2NetworkSummaryLsaSlice []Ospfv2NetworkSummaryLsa
	fieldPtr                     *[]*otg.Ospfv2NetworkSummaryLsa
}

func newOspfv2LsaStateOspfv2NetworkSummaryLsaIter(ptr *[]*otg.Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	return &ospfv2LsaStateOspfv2NetworkSummaryLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2NetworkSummaryLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	Items() []Ospfv2NetworkSummaryLsa
	Add() Ospfv2NetworkSummaryLsa
	Append(items ...Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	Set(index int, newObj Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	Clear() Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
	appendHolderSlice(item Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter
}

func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2NetworkSummaryLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) Items() []Ospfv2NetworkSummaryLsa {
	return obj.ospfv2NetworkSummaryLsaSlice
}

func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) Add() Ospfv2NetworkSummaryLsa {
	newObj := &otg.Ospfv2NetworkSummaryLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2NetworkSummaryLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2NetworkSummaryLsaSlice = append(obj.ospfv2NetworkSummaryLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) Append(items ...Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2NetworkSummaryLsaSlice = append(obj.ospfv2NetworkSummaryLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) Set(index int, newObj Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2NetworkSummaryLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) Clear() Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2NetworkSummaryLsa{}
		obj.ospfv2NetworkSummaryLsaSlice = []Ospfv2NetworkSummaryLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	if len(obj.ospfv2NetworkSummaryLsaSlice) > 0 {
		obj.ospfv2NetworkSummaryLsaSlice = []Ospfv2NetworkSummaryLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2NetworkSummaryLsaIter) appendHolderSlice(item Ospfv2NetworkSummaryLsa) Ospfv2LsaStateOspfv2NetworkSummaryLsaIter {
	obj.ospfv2NetworkSummaryLsaSlice = append(obj.ospfv2NetworkSummaryLsaSlice, item)
	return obj
}

// One or more OSPFv2 Autonomous System Boundary Router (ASBR) summary LSA - Type 4.
// SummaryAsLsas returns a []Ospfv2SummaryAsLsa
func (obj *ospfv2LsaState) SummaryAsLsas() Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	if len(obj.obj.SummaryAsLsas) == 0 {
		obj.obj.SummaryAsLsas = []*otg.Ospfv2SummaryAsLsa{}
	}
	if obj.summaryAsLsasHolder == nil {
		obj.summaryAsLsasHolder = newOspfv2LsaStateOspfv2SummaryAsLsaIter(&obj.obj.SummaryAsLsas).setMsg(obj)
	}
	return obj.summaryAsLsasHolder
}

type ospfv2LsaStateOspfv2SummaryAsLsaIter struct {
	obj                     *ospfv2LsaState
	ospfv2SummaryAsLsaSlice []Ospfv2SummaryAsLsa
	fieldPtr                *[]*otg.Ospfv2SummaryAsLsa
}

func newOspfv2LsaStateOspfv2SummaryAsLsaIter(ptr *[]*otg.Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	return &ospfv2LsaStateOspfv2SummaryAsLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2SummaryAsLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2SummaryAsLsaIter
	Items() []Ospfv2SummaryAsLsa
	Add() Ospfv2SummaryAsLsa
	Append(items ...Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter
	Set(index int, newObj Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter
	Clear() Ospfv2LsaStateOspfv2SummaryAsLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2SummaryAsLsaIter
	appendHolderSlice(item Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter
}

func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2SummaryAsLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) Items() []Ospfv2SummaryAsLsa {
	return obj.ospfv2SummaryAsLsaSlice
}

func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) Add() Ospfv2SummaryAsLsa {
	newObj := &otg.Ospfv2SummaryAsLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2SummaryAsLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2SummaryAsLsaSlice = append(obj.ospfv2SummaryAsLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) Append(items ...Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2SummaryAsLsaSlice = append(obj.ospfv2SummaryAsLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) Set(index int, newObj Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2SummaryAsLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) Clear() Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2SummaryAsLsa{}
		obj.ospfv2SummaryAsLsaSlice = []Ospfv2SummaryAsLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	if len(obj.ospfv2SummaryAsLsaSlice) > 0 {
		obj.ospfv2SummaryAsLsaSlice = []Ospfv2SummaryAsLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2SummaryAsLsaIter) appendHolderSlice(item Ospfv2SummaryAsLsa) Ospfv2LsaStateOspfv2SummaryAsLsaIter {
	obj.ospfv2SummaryAsLsaSlice = append(obj.ospfv2SummaryAsLsaSlice, item)
	return obj
}

// OSPFv2 AS-External-LSA - Type 5.
// ExternalAsLsas returns a []Ospfv2ExternalAsLsa
func (obj *ospfv2LsaState) ExternalAsLsas() Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	if len(obj.obj.ExternalAsLsas) == 0 {
		obj.obj.ExternalAsLsas = []*otg.Ospfv2ExternalAsLsa{}
	}
	if obj.externalAsLsasHolder == nil {
		obj.externalAsLsasHolder = newOspfv2LsaStateOspfv2ExternalAsLsaIter(&obj.obj.ExternalAsLsas).setMsg(obj)
	}
	return obj.externalAsLsasHolder
}

type ospfv2LsaStateOspfv2ExternalAsLsaIter struct {
	obj                      *ospfv2LsaState
	ospfv2ExternalAsLsaSlice []Ospfv2ExternalAsLsa
	fieldPtr                 *[]*otg.Ospfv2ExternalAsLsa
}

func newOspfv2LsaStateOspfv2ExternalAsLsaIter(ptr *[]*otg.Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	return &ospfv2LsaStateOspfv2ExternalAsLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2ExternalAsLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2ExternalAsLsaIter
	Items() []Ospfv2ExternalAsLsa
	Add() Ospfv2ExternalAsLsa
	Append(items ...Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter
	Set(index int, newObj Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter
	Clear() Ospfv2LsaStateOspfv2ExternalAsLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2ExternalAsLsaIter
	appendHolderSlice(item Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter
}

func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2ExternalAsLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) Items() []Ospfv2ExternalAsLsa {
	return obj.ospfv2ExternalAsLsaSlice
}

func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) Add() Ospfv2ExternalAsLsa {
	newObj := &otg.Ospfv2ExternalAsLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2ExternalAsLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2ExternalAsLsaSlice = append(obj.ospfv2ExternalAsLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) Append(items ...Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2ExternalAsLsaSlice = append(obj.ospfv2ExternalAsLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) Set(index int, newObj Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2ExternalAsLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) Clear() Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2ExternalAsLsa{}
		obj.ospfv2ExternalAsLsaSlice = []Ospfv2ExternalAsLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	if len(obj.ospfv2ExternalAsLsaSlice) > 0 {
		obj.ospfv2ExternalAsLsaSlice = []Ospfv2ExternalAsLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2ExternalAsLsaIter) appendHolderSlice(item Ospfv2ExternalAsLsa) Ospfv2LsaStateOspfv2ExternalAsLsaIter {
	obj.ospfv2ExternalAsLsaSlice = append(obj.ospfv2ExternalAsLsaSlice, item)
	return obj
}

// One or more OSPFv2 NSSA-LSA - Type 7.
// NssaLsas returns a []Ospfv2NssaLsa
func (obj *ospfv2LsaState) NssaLsas() Ospfv2LsaStateOspfv2NssaLsaIter {
	if len(obj.obj.NssaLsas) == 0 {
		obj.obj.NssaLsas = []*otg.Ospfv2NssaLsa{}
	}
	if obj.nssaLsasHolder == nil {
		obj.nssaLsasHolder = newOspfv2LsaStateOspfv2NssaLsaIter(&obj.obj.NssaLsas).setMsg(obj)
	}
	return obj.nssaLsasHolder
}

type ospfv2LsaStateOspfv2NssaLsaIter struct {
	obj                *ospfv2LsaState
	ospfv2NssaLsaSlice []Ospfv2NssaLsa
	fieldPtr           *[]*otg.Ospfv2NssaLsa
}

func newOspfv2LsaStateOspfv2NssaLsaIter(ptr *[]*otg.Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter {
	return &ospfv2LsaStateOspfv2NssaLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2NssaLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2NssaLsaIter
	Items() []Ospfv2NssaLsa
	Add() Ospfv2NssaLsa
	Append(items ...Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter
	Set(index int, newObj Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter
	Clear() Ospfv2LsaStateOspfv2NssaLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2NssaLsaIter
	appendHolderSlice(item Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter
}

func (obj *ospfv2LsaStateOspfv2NssaLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2NssaLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2NssaLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2NssaLsaIter) Items() []Ospfv2NssaLsa {
	return obj.ospfv2NssaLsaSlice
}

func (obj *ospfv2LsaStateOspfv2NssaLsaIter) Add() Ospfv2NssaLsa {
	newObj := &otg.Ospfv2NssaLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2NssaLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2NssaLsaSlice = append(obj.ospfv2NssaLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2NssaLsaIter) Append(items ...Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2NssaLsaSlice = append(obj.ospfv2NssaLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2NssaLsaIter) Set(index int, newObj Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2NssaLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2NssaLsaIter) Clear() Ospfv2LsaStateOspfv2NssaLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2NssaLsa{}
		obj.ospfv2NssaLsaSlice = []Ospfv2NssaLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2NssaLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2NssaLsaIter {
	if len(obj.ospfv2NssaLsaSlice) > 0 {
		obj.ospfv2NssaLsaSlice = []Ospfv2NssaLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2NssaLsaIter) appendHolderSlice(item Ospfv2NssaLsa) Ospfv2LsaStateOspfv2NssaLsaIter {
	obj.ospfv2NssaLsaSlice = append(obj.ospfv2NssaLsaSlice, item)
	return obj
}

// One or more OSPFv2 Link-Scope Opaque-LSA - Type 9.
// OpaqueLsas returns a []Ospfv2OpaqueLsa
func (obj *ospfv2LsaState) OpaqueLsas() Ospfv2LsaStateOspfv2OpaqueLsaIter {
	if len(obj.obj.OpaqueLsas) == 0 {
		obj.obj.OpaqueLsas = []*otg.Ospfv2OpaqueLsa{}
	}
	if obj.opaqueLsasHolder == nil {
		obj.opaqueLsasHolder = newOspfv2LsaStateOspfv2OpaqueLsaIter(&obj.obj.OpaqueLsas).setMsg(obj)
	}
	return obj.opaqueLsasHolder
}

type ospfv2LsaStateOspfv2OpaqueLsaIter struct {
	obj                  *ospfv2LsaState
	ospfv2OpaqueLsaSlice []Ospfv2OpaqueLsa
	fieldPtr             *[]*otg.Ospfv2OpaqueLsa
}

func newOspfv2LsaStateOspfv2OpaqueLsaIter(ptr *[]*otg.Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter {
	return &ospfv2LsaStateOspfv2OpaqueLsaIter{fieldPtr: ptr}
}

type Ospfv2LsaStateOspfv2OpaqueLsaIter interface {
	setMsg(*ospfv2LsaState) Ospfv2LsaStateOspfv2OpaqueLsaIter
	Items() []Ospfv2OpaqueLsa
	Add() Ospfv2OpaqueLsa
	Append(items ...Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter
	Set(index int, newObj Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter
	Clear() Ospfv2LsaStateOspfv2OpaqueLsaIter
	clearHolderSlice() Ospfv2LsaStateOspfv2OpaqueLsaIter
	appendHolderSlice(item Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter
}

func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) setMsg(msg *ospfv2LsaState) Ospfv2LsaStateOspfv2OpaqueLsaIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2OpaqueLsa{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) Items() []Ospfv2OpaqueLsa {
	return obj.ospfv2OpaqueLsaSlice
}

func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) Add() Ospfv2OpaqueLsa {
	newObj := &otg.Ospfv2OpaqueLsa{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2OpaqueLsa{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2OpaqueLsaSlice = append(obj.ospfv2OpaqueLsaSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) Append(items ...Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2OpaqueLsaSlice = append(obj.ospfv2OpaqueLsaSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) Set(index int, newObj Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2OpaqueLsaSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) Clear() Ospfv2LsaStateOspfv2OpaqueLsaIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2OpaqueLsa{}
		obj.ospfv2OpaqueLsaSlice = []Ospfv2OpaqueLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) clearHolderSlice() Ospfv2LsaStateOspfv2OpaqueLsaIter {
	if len(obj.ospfv2OpaqueLsaSlice) > 0 {
		obj.ospfv2OpaqueLsaSlice = []Ospfv2OpaqueLsa{}
	}
	return obj
}
func (obj *ospfv2LsaStateOspfv2OpaqueLsaIter) appendHolderSlice(item Ospfv2OpaqueLsa) Ospfv2LsaStateOspfv2OpaqueLsaIter {
	obj.ospfv2OpaqueLsaSlice = append(obj.ospfv2OpaqueLsaSlice, item)
	return obj
}

func (obj *ospfv2LsaState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.RouterLsas) != 0 {

		if set_default {
			obj.RouterLsas().clearHolderSlice()
			for _, item := range obj.obj.RouterLsas {
				obj.RouterLsas().appendHolderSlice(&ospfv2RouterLsa{obj: item})
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
				obj.NetworkLsas().appendHolderSlice(&ospfv2NetworkLsa{obj: item})
			}
		}
		for _, item := range obj.NetworkLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.NetworkSummaryLsas) != 0 {

		if set_default {
			obj.NetworkSummaryLsas().clearHolderSlice()
			for _, item := range obj.obj.NetworkSummaryLsas {
				obj.NetworkSummaryLsas().appendHolderSlice(&ospfv2NetworkSummaryLsa{obj: item})
			}
		}
		for _, item := range obj.NetworkSummaryLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.SummaryAsLsas) != 0 {

		if set_default {
			obj.SummaryAsLsas().clearHolderSlice()
			for _, item := range obj.obj.SummaryAsLsas {
				obj.SummaryAsLsas().appendHolderSlice(&ospfv2SummaryAsLsa{obj: item})
			}
		}
		for _, item := range obj.SummaryAsLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.ExternalAsLsas) != 0 {

		if set_default {
			obj.ExternalAsLsas().clearHolderSlice()
			for _, item := range obj.obj.ExternalAsLsas {
				obj.ExternalAsLsas().appendHolderSlice(&ospfv2ExternalAsLsa{obj: item})
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
				obj.NssaLsas().appendHolderSlice(&ospfv2NssaLsa{obj: item})
			}
		}
		for _, item := range obj.NssaLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.OpaqueLsas) != 0 {

		if set_default {
			obj.OpaqueLsas().clearHolderSlice()
			for _, item := range obj.obj.OpaqueLsas {
				obj.OpaqueLsas().appendHolderSlice(&ospfv2OpaqueLsa{obj: item})
			}
		}
		for _, item := range obj.OpaqueLsas().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2LsaState) setDefault() {

}
