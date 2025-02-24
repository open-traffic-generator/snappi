package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpIPv4LspState *****
type rsvpIPv4LspState struct {
	validation
	obj          *otg.RsvpIPv4LspState
	marshaller   marshalRsvpIPv4LspState
	unMarshaller unMarshalRsvpIPv4LspState
	lspHolder    RsvpLspState
	rrosHolder   RsvpIPv4LspStateRsvpLspIpv4RroIter
	erosHolder   RsvpIPv4LspStateRsvpLspIpv4EroIter
}

func NewRsvpIPv4LspState() RsvpIPv4LspState {
	obj := rsvpIPv4LspState{obj: &otg.RsvpIPv4LspState{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpIPv4LspState) msg() *otg.RsvpIPv4LspState {
	return obj.obj
}

func (obj *rsvpIPv4LspState) setMsg(msg *otg.RsvpIPv4LspState) RsvpIPv4LspState {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpIPv4LspState struct {
	obj *rsvpIPv4LspState
}

type marshalRsvpIPv4LspState interface {
	// ToProto marshals RsvpIPv4LspState to protobuf object *otg.RsvpIPv4LspState
	ToProto() (*otg.RsvpIPv4LspState, error)
	// ToPbText marshals RsvpIPv4LspState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpIPv4LspState to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpIPv4LspState to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpIPv4LspState to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpIPv4LspState struct {
	obj *rsvpIPv4LspState
}

type unMarshalRsvpIPv4LspState interface {
	// FromProto unmarshals RsvpIPv4LspState from protobuf object *otg.RsvpIPv4LspState
	FromProto(msg *otg.RsvpIPv4LspState) (RsvpIPv4LspState, error)
	// FromPbText unmarshals RsvpIPv4LspState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpIPv4LspState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpIPv4LspState from JSON text
	FromJson(value string) error
}

func (obj *rsvpIPv4LspState) Marshal() marshalRsvpIPv4LspState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpIPv4LspState{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpIPv4LspState) Unmarshal() unMarshalRsvpIPv4LspState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpIPv4LspState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpIPv4LspState) ToProto() (*otg.RsvpIPv4LspState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpIPv4LspState) FromProto(msg *otg.RsvpIPv4LspState) (RsvpIPv4LspState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpIPv4LspState) ToPbText() (string, error) {
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

func (m *unMarshalrsvpIPv4LspState) FromPbText(value string) error {
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

func (m *marshalrsvpIPv4LspState) ToYaml() (string, error) {
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

func (m *unMarshalrsvpIPv4LspState) FromYaml(value string) error {
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

func (m *marshalrsvpIPv4LspState) ToJsonRaw() (string, error) {
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

func (m *marshalrsvpIPv4LspState) ToJson() (string, error) {
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

func (m *unMarshalrsvpIPv4LspState) FromJson(value string) error {
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

func (obj *rsvpIPv4LspState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpIPv4LspState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpIPv4LspState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpIPv4LspState) Clone() (RsvpIPv4LspState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpIPv4LspState()
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

func (obj *rsvpIPv4LspState) setNil() {
	obj.lspHolder = nil
	obj.rrosHolder = nil
	obj.erosHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RsvpIPv4LspState is iPv4 RSVP-TE Discovered LSPs.
type RsvpIPv4LspState interface {
	Validation
	// msg marshals RsvpIPv4LspState to protobuf object *otg.RsvpIPv4LspState
	// and doesn't set defaults
	msg() *otg.RsvpIPv4LspState
	// setMsg unmarshals RsvpIPv4LspState from protobuf object *otg.RsvpIPv4LspState
	// and doesn't set defaults
	setMsg(*otg.RsvpIPv4LspState) RsvpIPv4LspState
	// provides marshal interface
	Marshal() marshalRsvpIPv4LspState
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpIPv4LspState
	// validate validates RsvpIPv4LspState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpIPv4LspState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SourceAddress returns string, set in RsvpIPv4LspState.
	SourceAddress() string
	// SetSourceAddress assigns string provided by user to RsvpIPv4LspState
	SetSourceAddress(value string) RsvpIPv4LspState
	// HasSourceAddress checks if SourceAddress has been set in RsvpIPv4LspState
	HasSourceAddress() bool
	// DestinationAddress returns string, set in RsvpIPv4LspState.
	DestinationAddress() string
	// SetDestinationAddress assigns string provided by user to RsvpIPv4LspState
	SetDestinationAddress(value string) RsvpIPv4LspState
	// HasDestinationAddress checks if DestinationAddress has been set in RsvpIPv4LspState
	HasDestinationAddress() bool
	// Lsp returns RsvpLspState, set in RsvpIPv4LspState.
	// RsvpLspState is iPv4 RSVP-TE Discovered LSPs.
	Lsp() RsvpLspState
	// SetLsp assigns RsvpLspState provided by user to RsvpIPv4LspState.
	// RsvpLspState is iPv4 RSVP-TE Discovered LSPs.
	SetLsp(value RsvpLspState) RsvpIPv4LspState
	// HasLsp checks if Lsp has been set in RsvpIPv4LspState
	HasLsp() bool
	// Rros returns RsvpIPv4LspStateRsvpLspIpv4RroIterIter, set in RsvpIPv4LspState
	Rros() RsvpIPv4LspStateRsvpLspIpv4RroIter
	// Eros returns RsvpIPv4LspStateRsvpLspIpv4EroIterIter, set in RsvpIPv4LspState
	Eros() RsvpIPv4LspStateRsvpLspIpv4EroIter
	setNil()
}

// The origin IPv4 address of RSVP session.
// SourceAddress returns a string
func (obj *rsvpIPv4LspState) SourceAddress() string {

	return *obj.obj.SourceAddress

}

// The origin IPv4 address of RSVP session.
// SourceAddress returns a string
func (obj *rsvpIPv4LspState) HasSourceAddress() bool {
	return obj.obj.SourceAddress != nil
}

// The origin IPv4 address of RSVP session.
// SetSourceAddress sets the string value in the RsvpIPv4LspState object
func (obj *rsvpIPv4LspState) SetSourceAddress(value string) RsvpIPv4LspState {

	obj.obj.SourceAddress = &value
	return obj
}

// The IPv4 destination address of RSVP session.
// DestinationAddress returns a string
func (obj *rsvpIPv4LspState) DestinationAddress() string {

	return *obj.obj.DestinationAddress

}

// The IPv4 destination address of RSVP session.
// DestinationAddress returns a string
func (obj *rsvpIPv4LspState) HasDestinationAddress() bool {
	return obj.obj.DestinationAddress != nil
}

// The IPv4 destination address of RSVP session.
// SetDestinationAddress sets the string value in the RsvpIPv4LspState object
func (obj *rsvpIPv4LspState) SetDestinationAddress(value string) RsvpIPv4LspState {

	obj.obj.DestinationAddress = &value
	return obj
}

// It refers to the RSVP LSP properties.
// Lsp returns a RsvpLspState
func (obj *rsvpIPv4LspState) Lsp() RsvpLspState {
	if obj.obj.Lsp == nil {
		obj.obj.Lsp = NewRsvpLspState().msg()
	}
	if obj.lspHolder == nil {
		obj.lspHolder = &rsvpLspState{obj: obj.obj.Lsp}
	}
	return obj.lspHolder
}

// It refers to the RSVP LSP properties.
// Lsp returns a RsvpLspState
func (obj *rsvpIPv4LspState) HasLsp() bool {
	return obj.obj.Lsp != nil
}

// It refers to the RSVP LSP properties.
// SetLsp sets the RsvpLspState value in the RsvpIPv4LspState object
func (obj *rsvpIPv4LspState) SetLsp(value RsvpLspState) RsvpIPv4LspState {

	obj.lspHolder = nil
	obj.obj.Lsp = value.msg()

	return obj
}

// It refers to RSVP RRO objects container.
// Rros returns a []RsvpLspIpv4Rro
func (obj *rsvpIPv4LspState) Rros() RsvpIPv4LspStateRsvpLspIpv4RroIter {
	if len(obj.obj.Rros) == 0 {
		obj.obj.Rros = []*otg.RsvpLspIpv4Rro{}
	}
	if obj.rrosHolder == nil {
		obj.rrosHolder = newRsvpIPv4LspStateRsvpLspIpv4RroIter(&obj.obj.Rros).setMsg(obj)
	}
	return obj.rrosHolder
}

type rsvpIPv4LspStateRsvpLspIpv4RroIter struct {
	obj                 *rsvpIPv4LspState
	rsvpLspIpv4RroSlice []RsvpLspIpv4Rro
	fieldPtr            *[]*otg.RsvpLspIpv4Rro
}

func newRsvpIPv4LspStateRsvpLspIpv4RroIter(ptr *[]*otg.RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter {
	return &rsvpIPv4LspStateRsvpLspIpv4RroIter{fieldPtr: ptr}
}

type RsvpIPv4LspStateRsvpLspIpv4RroIter interface {
	setMsg(*rsvpIPv4LspState) RsvpIPv4LspStateRsvpLspIpv4RroIter
	Items() []RsvpLspIpv4Rro
	Add() RsvpLspIpv4Rro
	Append(items ...RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter
	Set(index int, newObj RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter
	Clear() RsvpIPv4LspStateRsvpLspIpv4RroIter
	clearHolderSlice() RsvpIPv4LspStateRsvpLspIpv4RroIter
	appendHolderSlice(item RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) setMsg(msg *rsvpIPv4LspState) RsvpIPv4LspStateRsvpLspIpv4RroIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpLspIpv4Rro{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) Items() []RsvpLspIpv4Rro {
	return obj.rsvpLspIpv4RroSlice
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) Add() RsvpLspIpv4Rro {
	newObj := &otg.RsvpLspIpv4Rro{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpLspIpv4Rro{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpLspIpv4RroSlice = append(obj.rsvpLspIpv4RroSlice, newLibObj)
	return newLibObj
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) Append(items ...RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpLspIpv4RroSlice = append(obj.rsvpLspIpv4RroSlice, item)
	}
	return obj
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) Set(index int, newObj RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpLspIpv4RroSlice[index] = newObj
	return obj
}
func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) Clear() RsvpIPv4LspStateRsvpLspIpv4RroIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpLspIpv4Rro{}
		obj.rsvpLspIpv4RroSlice = []RsvpLspIpv4Rro{}
	}
	return obj
}
func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) clearHolderSlice() RsvpIPv4LspStateRsvpLspIpv4RroIter {
	if len(obj.rsvpLspIpv4RroSlice) > 0 {
		obj.rsvpLspIpv4RroSlice = []RsvpLspIpv4Rro{}
	}
	return obj
}
func (obj *rsvpIPv4LspStateRsvpLspIpv4RroIter) appendHolderSlice(item RsvpLspIpv4Rro) RsvpIPv4LspStateRsvpLspIpv4RroIter {
	obj.rsvpLspIpv4RroSlice = append(obj.rsvpLspIpv4RroSlice, item)
	return obj
}

// It refers to RSVP ERO objects container.
// Eros returns a []RsvpLspIpv4Ero
func (obj *rsvpIPv4LspState) Eros() RsvpIPv4LspStateRsvpLspIpv4EroIter {
	if len(obj.obj.Eros) == 0 {
		obj.obj.Eros = []*otg.RsvpLspIpv4Ero{}
	}
	if obj.erosHolder == nil {
		obj.erosHolder = newRsvpIPv4LspStateRsvpLspIpv4EroIter(&obj.obj.Eros).setMsg(obj)
	}
	return obj.erosHolder
}

type rsvpIPv4LspStateRsvpLspIpv4EroIter struct {
	obj                 *rsvpIPv4LspState
	rsvpLspIpv4EroSlice []RsvpLspIpv4Ero
	fieldPtr            *[]*otg.RsvpLspIpv4Ero
}

func newRsvpIPv4LspStateRsvpLspIpv4EroIter(ptr *[]*otg.RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter {
	return &rsvpIPv4LspStateRsvpLspIpv4EroIter{fieldPtr: ptr}
}

type RsvpIPv4LspStateRsvpLspIpv4EroIter interface {
	setMsg(*rsvpIPv4LspState) RsvpIPv4LspStateRsvpLspIpv4EroIter
	Items() []RsvpLspIpv4Ero
	Add() RsvpLspIpv4Ero
	Append(items ...RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter
	Set(index int, newObj RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter
	Clear() RsvpIPv4LspStateRsvpLspIpv4EroIter
	clearHolderSlice() RsvpIPv4LspStateRsvpLspIpv4EroIter
	appendHolderSlice(item RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) setMsg(msg *rsvpIPv4LspState) RsvpIPv4LspStateRsvpLspIpv4EroIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpLspIpv4Ero{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) Items() []RsvpLspIpv4Ero {
	return obj.rsvpLspIpv4EroSlice
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) Add() RsvpLspIpv4Ero {
	newObj := &otg.RsvpLspIpv4Ero{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpLspIpv4Ero{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpLspIpv4EroSlice = append(obj.rsvpLspIpv4EroSlice, newLibObj)
	return newLibObj
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) Append(items ...RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpLspIpv4EroSlice = append(obj.rsvpLspIpv4EroSlice, item)
	}
	return obj
}

func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) Set(index int, newObj RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpLspIpv4EroSlice[index] = newObj
	return obj
}
func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) Clear() RsvpIPv4LspStateRsvpLspIpv4EroIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpLspIpv4Ero{}
		obj.rsvpLspIpv4EroSlice = []RsvpLspIpv4Ero{}
	}
	return obj
}
func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) clearHolderSlice() RsvpIPv4LspStateRsvpLspIpv4EroIter {
	if len(obj.rsvpLspIpv4EroSlice) > 0 {
		obj.rsvpLspIpv4EroSlice = []RsvpLspIpv4Ero{}
	}
	return obj
}
func (obj *rsvpIPv4LspStateRsvpLspIpv4EroIter) appendHolderSlice(item RsvpLspIpv4Ero) RsvpIPv4LspStateRsvpLspIpv4EroIter {
	obj.rsvpLspIpv4EroSlice = append(obj.rsvpLspIpv4EroSlice, item)
	return obj
}

func (obj *rsvpIPv4LspState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SourceAddress != nil {

		err := obj.validateIpv4(obj.SourceAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpIPv4LspState.SourceAddress"))
		}

	}

	if obj.obj.DestinationAddress != nil {

		err := obj.validateIpv4(obj.DestinationAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RsvpIPv4LspState.DestinationAddress"))
		}

	}

	if obj.obj.Lsp != nil {

		obj.Lsp().validateObj(vObj, set_default)
	}

	if len(obj.obj.Rros) != 0 {

		if set_default {
			obj.Rros().clearHolderSlice()
			for _, item := range obj.obj.Rros {
				obj.Rros().appendHolderSlice(&rsvpLspIpv4Rro{obj: item})
			}
		}
		for _, item := range obj.Rros().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Eros) != 0 {

		if set_default {
			obj.Eros().clearHolderSlice()
			for _, item := range obj.obj.Eros {
				obj.Eros().appendHolderSlice(&rsvpLspIpv4Ero{obj: item})
			}
		}
		for _, item := range obj.Eros().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rsvpIPv4LspState) setDefault() {

}
