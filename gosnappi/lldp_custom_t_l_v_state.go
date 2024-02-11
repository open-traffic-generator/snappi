package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpCustomTLVState *****
type lldpCustomTLVState struct {
	validation
	obj          *otg.LldpCustomTLVState
	marshaller   marshalLldpCustomTLVState
	unMarshaller unMarshalLldpCustomTLVState
}

func NewLldpCustomTLVState() LldpCustomTLVState {
	obj := lldpCustomTLVState{obj: &otg.LldpCustomTLVState{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpCustomTLVState) msg() *otg.LldpCustomTLVState {
	return obj.obj
}

func (obj *lldpCustomTLVState) setMsg(msg *otg.LldpCustomTLVState) LldpCustomTLVState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpCustomTLVState struct {
	obj *lldpCustomTLVState
}

type marshalLldpCustomTLVState interface {
	// ToProto marshals LldpCustomTLVState to protobuf object *otg.LldpCustomTLVState
	ToProto() (*otg.LldpCustomTLVState, error)
	// ToPbText marshals LldpCustomTLVState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpCustomTLVState to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpCustomTLVState to JSON text
	ToJson() (string, error)
}

type unMarshallldpCustomTLVState struct {
	obj *lldpCustomTLVState
}

type unMarshalLldpCustomTLVState interface {
	// FromProto unmarshals LldpCustomTLVState from protobuf object *otg.LldpCustomTLVState
	FromProto(msg *otg.LldpCustomTLVState) (LldpCustomTLVState, error)
	// FromPbText unmarshals LldpCustomTLVState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpCustomTLVState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpCustomTLVState from JSON text
	FromJson(value string) error
}

func (obj *lldpCustomTLVState) Marshal() marshalLldpCustomTLVState {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpCustomTLVState{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpCustomTLVState) Unmarshal() unMarshalLldpCustomTLVState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpCustomTLVState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpCustomTLVState) ToProto() (*otg.LldpCustomTLVState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpCustomTLVState) FromProto(msg *otg.LldpCustomTLVState) (LldpCustomTLVState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpCustomTLVState) ToPbText() (string, error) {
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

func (m *unMarshallldpCustomTLVState) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshallldpCustomTLVState) ToYaml() (string, error) {
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

func (m *unMarshallldpCustomTLVState) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshallldpCustomTLVState) ToJson() (string, error) {
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

func (m *unMarshallldpCustomTLVState) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *lldpCustomTLVState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpCustomTLVState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpCustomTLVState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpCustomTLVState) Clone() (LldpCustomTLVState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpCustomTLVState()
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

// LldpCustomTLVState is custom TLV received from a neighbor.Custom TLVs are organization specific TLVs advertised with TLV type 127.
type LldpCustomTLVState interface {
	Validation
	// msg marshals LldpCustomTLVState to protobuf object *otg.LldpCustomTLVState
	// and doesn't set defaults
	msg() *otg.LldpCustomTLVState
	// setMsg unmarshals LldpCustomTLVState from protobuf object *otg.LldpCustomTLVState
	// and doesn't set defaults
	setMsg(*otg.LldpCustomTLVState) LldpCustomTLVState
	// provides marshal interface
	Marshal() marshalLldpCustomTLVState
	// provides unmarshal interface
	Unmarshal() unMarshalLldpCustomTLVState
	// validate validates LldpCustomTLVState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpCustomTLVState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CustomType returns uint32, set in LldpCustomTLVState.
	CustomType() uint32
	// SetCustomType assigns uint32 provided by user to LldpCustomTLVState
	SetCustomType(value uint32) LldpCustomTLVState
	// HasCustomType checks if CustomType has been set in LldpCustomTLVState
	HasCustomType() bool
	// Oui returns string, set in LldpCustomTLVState.
	Oui() string
	// SetOui assigns string provided by user to LldpCustomTLVState
	SetOui(value string) LldpCustomTLVState
	// HasOui checks if Oui has been set in LldpCustomTLVState
	HasOui() bool
	// OuiSubtype returns string, set in LldpCustomTLVState.
	OuiSubtype() string
	// SetOuiSubtype assigns string provided by user to LldpCustomTLVState
	SetOuiSubtype(value string) LldpCustomTLVState
	// HasOuiSubtype checks if OuiSubtype has been set in LldpCustomTLVState
	HasOuiSubtype() bool
}

// The integer value identifying the type of information contained in the value field.
// CustomType returns a uint32
func (obj *lldpCustomTLVState) CustomType() uint32 {

	return *obj.obj.CustomType

}

// The integer value identifying the type of information contained in the value field.
// CustomType returns a uint32
func (obj *lldpCustomTLVState) HasCustomType() bool {
	return obj.obj.CustomType != nil
}

// The integer value identifying the type of information contained in the value field.
// SetCustomType sets the uint32 value in the LldpCustomTLVState object
func (obj *lldpCustomTLVState) SetCustomType(value uint32) LldpCustomTLVState {

	obj.obj.CustomType = &value
	return obj
}

// The organizationally unique identifier field shall contain the organization's OUI as defined in Clause 9 of IEEE Std 802. The high-order octet is 0 and the low-order 3 octets are the SMI Network Management Private Enterprise Code of the Vendor in network byte order,  as defined in the 'Assigned Numbers' RFC [RFC3232].
// Oui returns a string
func (obj *lldpCustomTLVState) Oui() string {

	return *obj.obj.Oui

}

// The organizationally unique identifier field shall contain the organization's OUI as defined in Clause 9 of IEEE Std 802. The high-order octet is 0 and the low-order 3 octets are the SMI Network Management Private Enterprise Code of the Vendor in network byte order,  as defined in the 'Assigned Numbers' RFC [RFC3232].
// Oui returns a string
func (obj *lldpCustomTLVState) HasOui() bool {
	return obj.obj.Oui != nil
}

// The organizationally unique identifier field shall contain the organization's OUI as defined in Clause 9 of IEEE Std 802. The high-order octet is 0 and the low-order 3 octets are the SMI Network Management Private Enterprise Code of the Vendor in network byte order,  as defined in the 'Assigned Numbers' RFC [RFC3232].
// SetOui sets the string value in the LldpCustomTLVState object
func (obj *lldpCustomTLVState) SetOui(value string) LldpCustomTLVState {

	obj.obj.Oui = &value
	return obj
}

// The organizationally defined subtype field shall contain a unique subtype value assigned by the defining organization.
// OuiSubtype returns a string
func (obj *lldpCustomTLVState) OuiSubtype() string {

	return *obj.obj.OuiSubtype

}

// The organizationally defined subtype field shall contain a unique subtype value assigned by the defining organization.
// OuiSubtype returns a string
func (obj *lldpCustomTLVState) HasOuiSubtype() bool {
	return obj.obj.OuiSubtype != nil
}

// The organizationally defined subtype field shall contain a unique subtype value assigned by the defining organization.
// SetOuiSubtype sets the string value in the LldpCustomTLVState object
func (obj *lldpCustomTLVState) SetOuiSubtype(value string) LldpCustomTLVState {

	obj.obj.OuiSubtype = &value
	return obj
}

func (obj *lldpCustomTLVState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpCustomTLVState) setDefault() {

}
