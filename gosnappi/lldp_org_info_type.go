package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpOrgInfoType *****
type lldpOrgInfoType struct {
	validation
	obj          *otg.LldpOrgInfoType
	marshaller   marshalLldpOrgInfoType
	unMarshaller unMarshalLldpOrgInfoType
}

func NewLldpOrgInfoType() LldpOrgInfoType {
	obj := lldpOrgInfoType{obj: &otg.LldpOrgInfoType{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpOrgInfoType) msg() *otg.LldpOrgInfoType {
	return obj.obj
}

func (obj *lldpOrgInfoType) setMsg(msg *otg.LldpOrgInfoType) LldpOrgInfoType {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpOrgInfoType struct {
	obj *lldpOrgInfoType
}

type marshalLldpOrgInfoType interface {
	// ToProto marshals LldpOrgInfoType to protobuf object *otg.LldpOrgInfoType
	ToProto() (*otg.LldpOrgInfoType, error)
	// ToPbText marshals LldpOrgInfoType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpOrgInfoType to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpOrgInfoType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpOrgInfoType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpOrgInfoType struct {
	obj *lldpOrgInfoType
}

type unMarshalLldpOrgInfoType interface {
	// FromProto unmarshals LldpOrgInfoType from protobuf object *otg.LldpOrgInfoType
	FromProto(msg *otg.LldpOrgInfoType) (LldpOrgInfoType, error)
	// FromPbText unmarshals LldpOrgInfoType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpOrgInfoType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpOrgInfoType from JSON text
	FromJson(value string) error
}

func (obj *lldpOrgInfoType) Marshal() marshalLldpOrgInfoType {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpOrgInfoType{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpOrgInfoType) Unmarshal() unMarshalLldpOrgInfoType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpOrgInfoType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpOrgInfoType) ToProto() (*otg.LldpOrgInfoType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpOrgInfoType) FromProto(msg *otg.LldpOrgInfoType) (LldpOrgInfoType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpOrgInfoType) ToPbText() (string, error) {
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

func (m *unMarshallldpOrgInfoType) FromPbText(value string) error {
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

func (m *marshallldpOrgInfoType) ToYaml() (string, error) {
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

func (m *unMarshallldpOrgInfoType) FromYaml(value string) error {
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

func (m *marshallldpOrgInfoType) ToJsonRaw() (string, error) {
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

func (m *marshallldpOrgInfoType) ToJson() (string, error) {
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

func (m *unMarshallldpOrgInfoType) FromJson(value string) error {
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

func (obj *lldpOrgInfoType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpOrgInfoType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpOrgInfoType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpOrgInfoType) Clone() (LldpOrgInfoType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpOrgInfoType()
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

// LldpOrgInfoType is contains either the Alpha-numeric information encoded in UTF-8 (as specified in IETF RFC 3629) or include one or more information fields with  their associated field-type identifiers designators, similar to those in the Management Address TLV. Currently only one choice as info is given in future if required it can be extended to define sub tlvs.
type LldpOrgInfoType interface {
	Validation
	// msg marshals LldpOrgInfoType to protobuf object *otg.LldpOrgInfoType
	// and doesn't set defaults
	msg() *otg.LldpOrgInfoType
	// setMsg unmarshals LldpOrgInfoType from protobuf object *otg.LldpOrgInfoType
	// and doesn't set defaults
	setMsg(*otg.LldpOrgInfoType) LldpOrgInfoType
	// provides marshal interface
	Marshal() marshalLldpOrgInfoType
	// provides unmarshal interface
	Unmarshal() unMarshalLldpOrgInfoType
	// validate validates LldpOrgInfoType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpOrgInfoType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpOrgInfoTypeChoiceEnum, set in LldpOrgInfoType
	Choice() LldpOrgInfoTypeChoiceEnum
	// setChoice assigns LldpOrgInfoTypeChoiceEnum provided by user to LldpOrgInfoType
	setChoice(value LldpOrgInfoTypeChoiceEnum) LldpOrgInfoType
	// HasChoice checks if Choice has been set in LldpOrgInfoType
	HasChoice() bool
	// Info returns string, set in LldpOrgInfoType.
	Info() string
	// SetInfo assigns string provided by user to LldpOrgInfoType
	SetInfo(value string) LldpOrgInfoType
	// HasInfo checks if Info has been set in LldpOrgInfoType
	HasInfo() bool
}

type LldpOrgInfoTypeChoiceEnum string

// Enum of Choice on LldpOrgInfoType
var LldpOrgInfoTypeChoice = struct {
	INFO LldpOrgInfoTypeChoiceEnum
}{
	INFO: LldpOrgInfoTypeChoiceEnum("info"),
}

func (obj *lldpOrgInfoType) Choice() LldpOrgInfoTypeChoiceEnum {
	return LldpOrgInfoTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// In info mode the organizationally defined information contain either binary or alpha-numeric information encoded in UTF-8  (as specified in IETF RFC 3629).
// Choice returns a string
func (obj *lldpOrgInfoType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpOrgInfoType) setChoice(value LldpOrgInfoTypeChoiceEnum) LldpOrgInfoType {
	intValue, ok := otg.LldpOrgInfoType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpOrgInfoTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpOrgInfoType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Info = nil
	return obj
}

// The organizationally defined information encoded in UTF-8 (as specified in IETF RFC 3629). This byte stream can be of any  length from 1 to 507 bytes. In the info byte stream, one byte is represented as string of 2 characters, for example 2  character string (0x)AB represents value of a single byte. So the maximum length of this attribute is 1014 (507 * 2 hex  characters per byte).
// Info returns a string
func (obj *lldpOrgInfoType) Info() string {

	if obj.obj.Info == nil {
		obj.setChoice(LldpOrgInfoTypeChoice.INFO)
	}

	return *obj.obj.Info

}

// The organizationally defined information encoded in UTF-8 (as specified in IETF RFC 3629). This byte stream can be of any  length from 1 to 507 bytes. In the info byte stream, one byte is represented as string of 2 characters, for example 2  character string (0x)AB represents value of a single byte. So the maximum length of this attribute is 1014 (507 * 2 hex  characters per byte).
// Info returns a string
func (obj *lldpOrgInfoType) HasInfo() bool {
	return obj.obj.Info != nil
}

// The organizationally defined information encoded in UTF-8 (as specified in IETF RFC 3629). This byte stream can be of any  length from 1 to 507 bytes. In the info byte stream, one byte is represented as string of 2 characters, for example 2  character string (0x)AB represents value of a single byte. So the maximum length of this attribute is 1014 (507 * 2 hex  characters per byte).
// SetInfo sets the string value in the LldpOrgInfoType object
func (obj *lldpOrgInfoType) SetInfo(value string) LldpOrgInfoType {
	obj.setChoice(LldpOrgInfoTypeChoice.INFO)
	obj.obj.Info = &value
	return obj
}

func (obj *lldpOrgInfoType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Info != nil {

		if len(*obj.obj.Info) < 1 || len(*obj.obj.Info) > 1014 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"1 <= length of LldpOrgInfoType.Info <= 1014 but Got %d",
					len(*obj.obj.Info)))
		}

	}

}

func (obj *lldpOrgInfoType) setDefault() {
	var choices_set int = 0
	var choice LldpOrgInfoTypeChoiceEnum

	if obj.obj.Info != nil {
		choices_set += 1
		choice = LldpOrgInfoTypeChoice.INFO
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(LldpOrgInfoTypeChoice.INFO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LldpOrgInfoType")
			}
		} else {
			intVal := otg.LldpOrgInfoType_Choice_Enum_value[string(choice)]
			enumValue := otg.LldpOrgInfoType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
