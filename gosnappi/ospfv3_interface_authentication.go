package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3InterfaceAuthentication *****
type ospfv3InterfaceAuthentication struct {
	validation
	obj          *otg.Ospfv3InterfaceAuthentication
	marshaller   marshalOspfv3InterfaceAuthentication
	unMarshaller unMarshalOspfv3InterfaceAuthentication
	authsHolder  Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
}

func NewOspfv3InterfaceAuthentication() Ospfv3InterfaceAuthentication {
	obj := ospfv3InterfaceAuthentication{obj: &otg.Ospfv3InterfaceAuthentication{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3InterfaceAuthentication) msg() *otg.Ospfv3InterfaceAuthentication {
	return obj.obj
}

func (obj *ospfv3InterfaceAuthentication) setMsg(msg *otg.Ospfv3InterfaceAuthentication) Ospfv3InterfaceAuthentication {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3InterfaceAuthentication struct {
	obj *ospfv3InterfaceAuthentication
}

type marshalOspfv3InterfaceAuthentication interface {
	// ToProto marshals Ospfv3InterfaceAuthentication to protobuf object *otg.Ospfv3InterfaceAuthentication
	ToProto() (*otg.Ospfv3InterfaceAuthentication, error)
	// ToPbText marshals Ospfv3InterfaceAuthentication to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3InterfaceAuthentication to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3InterfaceAuthentication to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3InterfaceAuthentication struct {
	obj *ospfv3InterfaceAuthentication
}

type unMarshalOspfv3InterfaceAuthentication interface {
	// FromProto unmarshals Ospfv3InterfaceAuthentication from protobuf object *otg.Ospfv3InterfaceAuthentication
	FromProto(msg *otg.Ospfv3InterfaceAuthentication) (Ospfv3InterfaceAuthentication, error)
	// FromPbText unmarshals Ospfv3InterfaceAuthentication from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3InterfaceAuthentication from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3InterfaceAuthentication from JSON text
	FromJson(value string) error
}

func (obj *ospfv3InterfaceAuthentication) Marshal() marshalOspfv3InterfaceAuthentication {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3InterfaceAuthentication{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3InterfaceAuthentication) Unmarshal() unMarshalOspfv3InterfaceAuthentication {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3InterfaceAuthentication{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3InterfaceAuthentication) ToProto() (*otg.Ospfv3InterfaceAuthentication, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3InterfaceAuthentication) FromProto(msg *otg.Ospfv3InterfaceAuthentication) (Ospfv3InterfaceAuthentication, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3InterfaceAuthentication) ToPbText() (string, error) {
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

func (m *unMarshalospfv3InterfaceAuthentication) FromPbText(value string) error {
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

func (m *marshalospfv3InterfaceAuthentication) ToYaml() (string, error) {
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

func (m *unMarshalospfv3InterfaceAuthentication) FromYaml(value string) error {
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

func (m *marshalospfv3InterfaceAuthentication) ToJson() (string, error) {
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

func (m *unMarshalospfv3InterfaceAuthentication) FromJson(value string) error {
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

func (obj *ospfv3InterfaceAuthentication) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceAuthentication) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3InterfaceAuthentication) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3InterfaceAuthentication) Clone() (Ospfv3InterfaceAuthentication, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3InterfaceAuthentication()
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

func (obj *ospfv3InterfaceAuthentication) setNil() {
	obj.authsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3InterfaceAuthentication is this contains OSPFv3 authentication properties.
type Ospfv3InterfaceAuthentication interface {
	Validation
	// msg marshals Ospfv3InterfaceAuthentication to protobuf object *otg.Ospfv3InterfaceAuthentication
	// and doesn't set defaults
	msg() *otg.Ospfv3InterfaceAuthentication
	// setMsg unmarshals Ospfv3InterfaceAuthentication from protobuf object *otg.Ospfv3InterfaceAuthentication
	// and doesn't set defaults
	setMsg(*otg.Ospfv3InterfaceAuthentication) Ospfv3InterfaceAuthentication
	// provides marshal interface
	Marshal() marshalOspfv3InterfaceAuthentication
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3InterfaceAuthentication
	// validate validates Ospfv3InterfaceAuthentication
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3InterfaceAuthentication, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Auths returns Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIterIter, set in Ospfv3InterfaceAuthentication
	Auths() Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
	setNil()
}

// List of OSPFv3 SA Key IDs and Keys.
// Auths returns a []Ospfv3AuthenticationAuth
func (obj *ospfv3InterfaceAuthentication) Auths() Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	if len(obj.obj.Auths) == 0 {
		obj.obj.Auths = []*otg.Ospfv3AuthenticationAuth{}
	}
	if obj.authsHolder == nil {
		obj.authsHolder = newOspfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter(&obj.obj.Auths).setMsg(obj)
	}
	return obj.authsHolder
}

type ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter struct {
	obj                           *ospfv3InterfaceAuthentication
	ospfv3AuthenticationAuthSlice []Ospfv3AuthenticationAuth
	fieldPtr                      *[]*otg.Ospfv3AuthenticationAuth
}

func newOspfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter(ptr *[]*otg.Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	return &ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter{fieldPtr: ptr}
}

type Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter interface {
	setMsg(*ospfv3InterfaceAuthentication) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
	Items() []Ospfv3AuthenticationAuth
	Add() Ospfv3AuthenticationAuth
	Append(items ...Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
	Set(index int, newObj Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
	Clear() Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
	clearHolderSlice() Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
	appendHolderSlice(item Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter
}

func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) setMsg(msg *ospfv3InterfaceAuthentication) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv3AuthenticationAuth{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) Items() []Ospfv3AuthenticationAuth {
	return obj.ospfv3AuthenticationAuthSlice
}

func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) Add() Ospfv3AuthenticationAuth {
	newObj := &otg.Ospfv3AuthenticationAuth{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv3AuthenticationAuth{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv3AuthenticationAuthSlice = append(obj.ospfv3AuthenticationAuthSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) Append(items ...Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv3AuthenticationAuthSlice = append(obj.ospfv3AuthenticationAuthSlice, item)
	}
	return obj
}

func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) Set(index int, newObj Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv3AuthenticationAuthSlice[index] = newObj
	return obj
}
func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) Clear() Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv3AuthenticationAuth{}
		obj.ospfv3AuthenticationAuthSlice = []Ospfv3AuthenticationAuth{}
	}
	return obj
}
func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) clearHolderSlice() Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	if len(obj.ospfv3AuthenticationAuthSlice) > 0 {
		obj.ospfv3AuthenticationAuthSlice = []Ospfv3AuthenticationAuth{}
	}
	return obj
}
func (obj *ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter) appendHolderSlice(item Ospfv3AuthenticationAuth) Ospfv3InterfaceAuthenticationOspfv3AuthenticationAuthIter {
	obj.ospfv3AuthenticationAuthSlice = append(obj.ospfv3AuthenticationAuthSlice, item)
	return obj
}

func (obj *ospfv3InterfaceAuthentication) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Auths) != 0 {

		if set_default {
			obj.Auths().clearHolderSlice()
			for _, item := range obj.obj.Auths {
				obj.Auths().appendHolderSlice(&ospfv3AuthenticationAuth{obj: item})
			}
		}
		for _, item := range obj.Auths().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv3InterfaceAuthentication) setDefault() {

}
