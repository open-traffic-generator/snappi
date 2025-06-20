package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpOrgInfo *****
type lldpOrgInfo struct {
	validation
	obj               *otg.LldpOrgInfo
	marshaller        marshalLldpOrgInfo
	unMarshaller      unMarshalLldpOrgInfo
	informationHolder LldpOrgInfoType
}

func NewLldpOrgInfo() LldpOrgInfo {
	obj := lldpOrgInfo{obj: &otg.LldpOrgInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpOrgInfo) msg() *otg.LldpOrgInfo {
	return obj.obj
}

func (obj *lldpOrgInfo) setMsg(msg *otg.LldpOrgInfo) LldpOrgInfo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpOrgInfo struct {
	obj *lldpOrgInfo
}

type marshalLldpOrgInfo interface {
	// ToProto marshals LldpOrgInfo to protobuf object *otg.LldpOrgInfo
	ToProto() (*otg.LldpOrgInfo, error)
	// ToPbText marshals LldpOrgInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpOrgInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpOrgInfo to JSON text
	ToJson() (string, error)
}

type unMarshallldpOrgInfo struct {
	obj *lldpOrgInfo
}

type unMarshalLldpOrgInfo interface {
	// FromProto unmarshals LldpOrgInfo from protobuf object *otg.LldpOrgInfo
	FromProto(msg *otg.LldpOrgInfo) (LldpOrgInfo, error)
	// FromPbText unmarshals LldpOrgInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpOrgInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpOrgInfo from JSON text
	FromJson(value string) error
}

func (obj *lldpOrgInfo) Marshal() marshalLldpOrgInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpOrgInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpOrgInfo) Unmarshal() unMarshalLldpOrgInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpOrgInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpOrgInfo) ToProto() (*otg.LldpOrgInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpOrgInfo) FromProto(msg *otg.LldpOrgInfo) (LldpOrgInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpOrgInfo) ToPbText() (string, error) {
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

func (m *unMarshallldpOrgInfo) FromPbText(value string) error {
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

func (m *marshallldpOrgInfo) ToYaml() (string, error) {
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

func (m *unMarshallldpOrgInfo) FromYaml(value string) error {
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

func (m *marshallldpOrgInfo) ToJson() (string, error) {
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

func (m *unMarshallldpOrgInfo) FromJson(value string) error {
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

func (obj *lldpOrgInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpOrgInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpOrgInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpOrgInfo) Clone() (LldpOrgInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpOrgInfo()
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

func (obj *lldpOrgInfo) setNil() {
	obj.informationHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LldpOrgInfo is the organization specific information configured in the Organization Specific TLV.
type LldpOrgInfo interface {
	Validation
	// msg marshals LldpOrgInfo to protobuf object *otg.LldpOrgInfo
	// and doesn't set defaults
	msg() *otg.LldpOrgInfo
	// setMsg unmarshals LldpOrgInfo from protobuf object *otg.LldpOrgInfo
	// and doesn't set defaults
	setMsg(*otg.LldpOrgInfo) LldpOrgInfo
	// provides marshal interface
	Marshal() marshalLldpOrgInfo
	// provides unmarshal interface
	Unmarshal() unMarshalLldpOrgInfo
	// validate validates LldpOrgInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpOrgInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Oui returns string, set in LldpOrgInfo.
	Oui() string
	// SetOui assigns string provided by user to LldpOrgInfo
	SetOui(value string) LldpOrgInfo
	// HasOui checks if Oui has been set in LldpOrgInfo
	HasOui() bool
	// Subtype returns uint32, set in LldpOrgInfo.
	Subtype() uint32
	// SetSubtype assigns uint32 provided by user to LldpOrgInfo
	SetSubtype(value uint32) LldpOrgInfo
	// HasSubtype checks if Subtype has been set in LldpOrgInfo
	HasSubtype() bool
	// Information returns LldpOrgInfoType, set in LldpOrgInfo.
	// LldpOrgInfoType is contains either the Alpha-numeric information encoded in UTF-8 (as specified in IETF RFC 3629) or include one or more information fields with  their associated field-type identifiers designators, similar to those in the Management Address TLV. Currently only one choice as info is given in future if required it can be extended to define sub tlvs.
	Information() LldpOrgInfoType
	// SetInformation assigns LldpOrgInfoType provided by user to LldpOrgInfo.
	// LldpOrgInfoType is contains either the Alpha-numeric information encoded in UTF-8 (as specified in IETF RFC 3629) or include one or more information fields with  their associated field-type identifiers designators, similar to those in the Management Address TLV. Currently only one choice as info is given in future if required it can be extended to define sub tlvs.
	SetInformation(value LldpOrgInfoType) LldpOrgInfo
	// HasInformation checks if Information has been set in LldpOrgInfo
	HasInformation() bool
	setNil()
}

// The organizationally unique identifier field shall contain the organization's OUI as defined in Clause 9 of IEEE  Std 802. It is a 24 bit number that uniquely identifies a vendor, manufacturer, or other organizations.
// Oui returns a string
func (obj *lldpOrgInfo) Oui() string {

	return *obj.obj.Oui

}

// The organizationally unique identifier field shall contain the organization's OUI as defined in Clause 9 of IEEE  Std 802. It is a 24 bit number that uniquely identifies a vendor, manufacturer, or other organizations.
// Oui returns a string
func (obj *lldpOrgInfo) HasOui() bool {
	return obj.obj.Oui != nil
}

// The organizationally unique identifier field shall contain the organization's OUI as defined in Clause 9 of IEEE  Std 802. It is a 24 bit number that uniquely identifies a vendor, manufacturer, or other organizations.
// SetOui sets the string value in the LldpOrgInfo object
func (obj *lldpOrgInfo) SetOui(value string) LldpOrgInfo {

	obj.obj.Oui = &value
	return obj
}

// The organizationally defined subtype field shall contain a unique subtype value assigned by the defining organization.
// Subtype returns a uint32
func (obj *lldpOrgInfo) Subtype() uint32 {

	return *obj.obj.Subtype

}

// The organizationally defined subtype field shall contain a unique subtype value assigned by the defining organization.
// Subtype returns a uint32
func (obj *lldpOrgInfo) HasSubtype() bool {
	return obj.obj.Subtype != nil
}

// The organizationally defined subtype field shall contain a unique subtype value assigned by the defining organization.
// SetSubtype sets the uint32 value in the LldpOrgInfo object
func (obj *lldpOrgInfo) SetSubtype(value uint32) LldpOrgInfo {

	obj.obj.Subtype = &value
	return obj
}

// Contains the organizationally defined information. The actual format of the organizationally defined information string  field is organizationally specific and can contain either binary or alpha-numeric information that is instance specific  for the particular TLV type and subtype. Alpha-numeric information are encoded in UTF-8 (as specified in IETF RFC 3629).  Or include one or more information fields with their associated field-type identifiers, designators similar to those in  the Management Address TLV.
// Information returns a LldpOrgInfoType
func (obj *lldpOrgInfo) Information() LldpOrgInfoType {
	if obj.obj.Information == nil {
		obj.obj.Information = NewLldpOrgInfoType().msg()
	}
	if obj.informationHolder == nil {
		obj.informationHolder = &lldpOrgInfoType{obj: obj.obj.Information}
	}
	return obj.informationHolder
}

// Contains the organizationally defined information. The actual format of the organizationally defined information string  field is organizationally specific and can contain either binary or alpha-numeric information that is instance specific  for the particular TLV type and subtype. Alpha-numeric information are encoded in UTF-8 (as specified in IETF RFC 3629).  Or include one or more information fields with their associated field-type identifiers, designators similar to those in  the Management Address TLV.
// Information returns a LldpOrgInfoType
func (obj *lldpOrgInfo) HasInformation() bool {
	return obj.obj.Information != nil
}

// Contains the organizationally defined information. The actual format of the organizationally defined information string  field is organizationally specific and can contain either binary or alpha-numeric information that is instance specific  for the particular TLV type and subtype. Alpha-numeric information are encoded in UTF-8 (as specified in IETF RFC 3629).  Or include one or more information fields with their associated field-type identifiers, designators similar to those in  the Management Address TLV.
// SetInformation sets the LldpOrgInfoType value in the LldpOrgInfo object
func (obj *lldpOrgInfo) SetInformation(value LldpOrgInfoType) LldpOrgInfo {

	obj.informationHolder = nil
	obj.obj.Information = value.msg()

	return obj
}

func (obj *lldpOrgInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Oui != nil {

		if len(*obj.obj.Oui) < 6 || len(*obj.obj.Oui) > 6 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"6 <= length of LldpOrgInfo.Oui <= 6 but Got %d",
					len(*obj.obj.Oui)))
		}

	}

	if obj.obj.Subtype != nil {

		if *obj.obj.Subtype < 1 || *obj.obj.Subtype > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= LldpOrgInfo.Subtype <= 127 but Got %d", *obj.obj.Subtype))
		}

	}

	if obj.obj.Information != nil {

		obj.Information().validateObj(vObj, set_default)
	}

}

func (obj *lldpOrgInfo) setDefault() {
	if obj.obj.Oui == nil {
		obj.SetOui("0080C2")
	}
	if obj.obj.Subtype == nil {
		obj.SetSubtype(1)
	}

}
