package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsServerIdentifier *****
type dhcpv6ClientOptionsServerIdentifier struct {
	validation
	obj            *otg.Dhcpv6ClientOptionsServerIdentifier
	marshaller     marshalDhcpv6ClientOptionsServerIdentifier
	unMarshaller   unMarshalDhcpv6ClientOptionsServerIdentifier
	duidLltHolder  Dhcpv6ClientOptionsDuidLlt
	duidEnHolder   Dhcpv6ClientOptionsDuidEn
	duidLlHolder   Dhcpv6ClientOptionsDuidLl
	duidUuidHolder Dhcpv6ClientOptionsDuidUuid
}

func NewDhcpv6ClientOptionsServerIdentifier() Dhcpv6ClientOptionsServerIdentifier {
	obj := dhcpv6ClientOptionsServerIdentifier{obj: &otg.Dhcpv6ClientOptionsServerIdentifier{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsServerIdentifier) msg() *otg.Dhcpv6ClientOptionsServerIdentifier {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsServerIdentifier) setMsg(msg *otg.Dhcpv6ClientOptionsServerIdentifier) Dhcpv6ClientOptionsServerIdentifier {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsServerIdentifier struct {
	obj *dhcpv6ClientOptionsServerIdentifier
}

type marshalDhcpv6ClientOptionsServerIdentifier interface {
	// ToProto marshals Dhcpv6ClientOptionsServerIdentifier to protobuf object *otg.Dhcpv6ClientOptionsServerIdentifier
	ToProto() (*otg.Dhcpv6ClientOptionsServerIdentifier, error)
	// ToPbText marshals Dhcpv6ClientOptionsServerIdentifier to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsServerIdentifier to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsServerIdentifier to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientOptionsServerIdentifier struct {
	obj *dhcpv6ClientOptionsServerIdentifier
}

type unMarshalDhcpv6ClientOptionsServerIdentifier interface {
	// FromProto unmarshals Dhcpv6ClientOptionsServerIdentifier from protobuf object *otg.Dhcpv6ClientOptionsServerIdentifier
	FromProto(msg *otg.Dhcpv6ClientOptionsServerIdentifier) (Dhcpv6ClientOptionsServerIdentifier, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsServerIdentifier from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsServerIdentifier from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsServerIdentifier from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsServerIdentifier) Marshal() marshalDhcpv6ClientOptionsServerIdentifier {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsServerIdentifier{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsServerIdentifier) Unmarshal() unMarshalDhcpv6ClientOptionsServerIdentifier {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsServerIdentifier{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsServerIdentifier) ToProto() (*otg.Dhcpv6ClientOptionsServerIdentifier, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsServerIdentifier) FromProto(msg *otg.Dhcpv6ClientOptionsServerIdentifier) (Dhcpv6ClientOptionsServerIdentifier, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsServerIdentifier) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsServerIdentifier) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsServerIdentifier) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsServerIdentifier) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsServerIdentifier) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsServerIdentifier) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsServerIdentifier) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsServerIdentifier) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsServerIdentifier) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsServerIdentifier) Clone() (Dhcpv6ClientOptionsServerIdentifier, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsServerIdentifier()
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

func (obj *dhcpv6ClientOptionsServerIdentifier) setNil() {
	obj.duidLltHolder = nil
	obj.duidEnHolder = nil
	obj.duidLlHolder = nil
	obj.duidUuidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsServerIdentifier is description is TBD
type Dhcpv6ClientOptionsServerIdentifier interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsServerIdentifier to protobuf object *otg.Dhcpv6ClientOptionsServerIdentifier
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsServerIdentifier
	// setMsg unmarshals Dhcpv6ClientOptionsServerIdentifier from protobuf object *otg.Dhcpv6ClientOptionsServerIdentifier
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsServerIdentifier) Dhcpv6ClientOptionsServerIdentifier
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsServerIdentifier
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsServerIdentifier
	// validate validates Dhcpv6ClientOptionsServerIdentifier
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsServerIdentifier, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsServerIdentifierChoiceEnum, set in Dhcpv6ClientOptionsServerIdentifier
	Choice() Dhcpv6ClientOptionsServerIdentifierChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsServerIdentifierChoiceEnum provided by user to Dhcpv6ClientOptionsServerIdentifier
	setChoice(value Dhcpv6ClientOptionsServerIdentifierChoiceEnum) Dhcpv6ClientOptionsServerIdentifier
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsServerIdentifier
	HasChoice() bool
	// DuidLlt returns Dhcpv6ClientOptionsDuidLlt, set in Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidLlt is dUID based on Link Layer address plus time. Hardware Type will be auto assigned to ethernet type.
	DuidLlt() Dhcpv6ClientOptionsDuidLlt
	// SetDuidLlt assigns Dhcpv6ClientOptionsDuidLlt provided by user to Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidLlt is dUID based on Link Layer address plus time. Hardware Type will be auto assigned to ethernet type.
	SetDuidLlt(value Dhcpv6ClientOptionsDuidLlt) Dhcpv6ClientOptionsServerIdentifier
	// HasDuidLlt checks if DuidLlt has been set in Dhcpv6ClientOptionsServerIdentifier
	HasDuidLlt() bool
	// DuidEn returns Dhcpv6ClientOptionsDuidEn, set in Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidEn is dUID assigned by vendor based on enterprise number.
	DuidEn() Dhcpv6ClientOptionsDuidEn
	// SetDuidEn assigns Dhcpv6ClientOptionsDuidEn provided by user to Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidEn is dUID assigned by vendor based on enterprise number.
	SetDuidEn(value Dhcpv6ClientOptionsDuidEn) Dhcpv6ClientOptionsServerIdentifier
	// HasDuidEn checks if DuidEn has been set in Dhcpv6ClientOptionsServerIdentifier
	HasDuidEn() bool
	// DuidLl returns Dhcpv6ClientOptionsDuidLl, set in Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidLl is dUID based on Link Layer address. Hardware Type will be auto assigned to ethernet type.
	DuidLl() Dhcpv6ClientOptionsDuidLl
	// SetDuidLl assigns Dhcpv6ClientOptionsDuidLl provided by user to Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidLl is dUID based on Link Layer address. Hardware Type will be auto assigned to ethernet type.
	SetDuidLl(value Dhcpv6ClientOptionsDuidLl) Dhcpv6ClientOptionsServerIdentifier
	// HasDuidLl checks if DuidLl has been set in Dhcpv6ClientOptionsServerIdentifier
	HasDuidLl() bool
	// DuidUuid returns Dhcpv6ClientOptionsDuidUuid, set in Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidUuid is dUID embedded a Universally Unique IDentifier (UUID). A UUID is an identifier that is unique across both  space and time, with respect to the space of all UUIDs.
	DuidUuid() Dhcpv6ClientOptionsDuidUuid
	// SetDuidUuid assigns Dhcpv6ClientOptionsDuidUuid provided by user to Dhcpv6ClientOptionsServerIdentifier.
	// Dhcpv6ClientOptionsDuidUuid is dUID embedded a Universally Unique IDentifier (UUID). A UUID is an identifier that is unique across both  space and time, with respect to the space of all UUIDs.
	SetDuidUuid(value Dhcpv6ClientOptionsDuidUuid) Dhcpv6ClientOptionsServerIdentifier
	// HasDuidUuid checks if DuidUuid has been set in Dhcpv6ClientOptionsServerIdentifier
	HasDuidUuid() bool
	setNil()
}

type Dhcpv6ClientOptionsServerIdentifierChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsServerIdentifier
var Dhcpv6ClientOptionsServerIdentifierChoice = struct {
	DUID_LLT  Dhcpv6ClientOptionsServerIdentifierChoiceEnum
	DUID_EN   Dhcpv6ClientOptionsServerIdentifierChoiceEnum
	DUID_LL   Dhcpv6ClientOptionsServerIdentifierChoiceEnum
	DUID_UUID Dhcpv6ClientOptionsServerIdentifierChoiceEnum
}{
	DUID_LLT:  Dhcpv6ClientOptionsServerIdentifierChoiceEnum("duid_llt"),
	DUID_EN:   Dhcpv6ClientOptionsServerIdentifierChoiceEnum("duid_en"),
	DUID_LL:   Dhcpv6ClientOptionsServerIdentifierChoiceEnum("duid_ll"),
	DUID_UUID: Dhcpv6ClientOptionsServerIdentifierChoiceEnum("duid_uuid"),
}

func (obj *dhcpv6ClientOptionsServerIdentifier) Choice() Dhcpv6ClientOptionsServerIdentifierChoiceEnum {
	return Dhcpv6ClientOptionsServerIdentifierChoiceEnum(obj.obj.Choice.Enum().String())
}

// The Identifier option is used to carry a DUID. The option code is 2. The server identifier identifies  a server. This option is used when client wants to contact a particular server.
// Choice returns a string
func (obj *dhcpv6ClientOptionsServerIdentifier) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsServerIdentifier) setChoice(value Dhcpv6ClientOptionsServerIdentifierChoiceEnum) Dhcpv6ClientOptionsServerIdentifier {
	intValue, ok := otg.Dhcpv6ClientOptionsServerIdentifier_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsServerIdentifierChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsServerIdentifier_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.DuidUuid = nil
	obj.duidUuidHolder = nil
	obj.obj.DuidLl = nil
	obj.duidLlHolder = nil
	obj.obj.DuidEn = nil
	obj.duidEnHolder = nil
	obj.obj.DuidLlt = nil
	obj.duidLltHolder = nil

	if value == Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LLT {
		obj.obj.DuidLlt = NewDhcpv6ClientOptionsDuidLlt().msg()
	}

	if value == Dhcpv6ClientOptionsServerIdentifierChoice.DUID_EN {
		obj.obj.DuidEn = NewDhcpv6ClientOptionsDuidEn().msg()
	}

	if value == Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LL {
		obj.obj.DuidLl = NewDhcpv6ClientOptionsDuidLl().msg()
	}

	if value == Dhcpv6ClientOptionsServerIdentifierChoice.DUID_UUID {
		obj.obj.DuidUuid = NewDhcpv6ClientOptionsDuidUuid().msg()
	}

	return obj
}

// description is TBD
// DuidLlt returns a Dhcpv6ClientOptionsDuidLlt
func (obj *dhcpv6ClientOptionsServerIdentifier) DuidLlt() Dhcpv6ClientOptionsDuidLlt {
	if obj.obj.DuidLlt == nil {
		obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LLT)
	}
	if obj.duidLltHolder == nil {
		obj.duidLltHolder = &dhcpv6ClientOptionsDuidLlt{obj: obj.obj.DuidLlt}
	}
	return obj.duidLltHolder
}

// description is TBD
// DuidLlt returns a Dhcpv6ClientOptionsDuidLlt
func (obj *dhcpv6ClientOptionsServerIdentifier) HasDuidLlt() bool {
	return obj.obj.DuidLlt != nil
}

// description is TBD
// SetDuidLlt sets the Dhcpv6ClientOptionsDuidLlt value in the Dhcpv6ClientOptionsServerIdentifier object
func (obj *dhcpv6ClientOptionsServerIdentifier) SetDuidLlt(value Dhcpv6ClientOptionsDuidLlt) Dhcpv6ClientOptionsServerIdentifier {
	obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LLT)
	obj.duidLltHolder = nil
	obj.obj.DuidLlt = value.msg()

	return obj
}

// description is TBD
// DuidEn returns a Dhcpv6ClientOptionsDuidEn
func (obj *dhcpv6ClientOptionsServerIdentifier) DuidEn() Dhcpv6ClientOptionsDuidEn {
	if obj.obj.DuidEn == nil {
		obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_EN)
	}
	if obj.duidEnHolder == nil {
		obj.duidEnHolder = &dhcpv6ClientOptionsDuidEn{obj: obj.obj.DuidEn}
	}
	return obj.duidEnHolder
}

// description is TBD
// DuidEn returns a Dhcpv6ClientOptionsDuidEn
func (obj *dhcpv6ClientOptionsServerIdentifier) HasDuidEn() bool {
	return obj.obj.DuidEn != nil
}

// description is TBD
// SetDuidEn sets the Dhcpv6ClientOptionsDuidEn value in the Dhcpv6ClientOptionsServerIdentifier object
func (obj *dhcpv6ClientOptionsServerIdentifier) SetDuidEn(value Dhcpv6ClientOptionsDuidEn) Dhcpv6ClientOptionsServerIdentifier {
	obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_EN)
	obj.duidEnHolder = nil
	obj.obj.DuidEn = value.msg()

	return obj
}

// description is TBD
// DuidLl returns a Dhcpv6ClientOptionsDuidLl
func (obj *dhcpv6ClientOptionsServerIdentifier) DuidLl() Dhcpv6ClientOptionsDuidLl {
	if obj.obj.DuidLl == nil {
		obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LL)
	}
	if obj.duidLlHolder == nil {
		obj.duidLlHolder = &dhcpv6ClientOptionsDuidLl{obj: obj.obj.DuidLl}
	}
	return obj.duidLlHolder
}

// description is TBD
// DuidLl returns a Dhcpv6ClientOptionsDuidLl
func (obj *dhcpv6ClientOptionsServerIdentifier) HasDuidLl() bool {
	return obj.obj.DuidLl != nil
}

// description is TBD
// SetDuidLl sets the Dhcpv6ClientOptionsDuidLl value in the Dhcpv6ClientOptionsServerIdentifier object
func (obj *dhcpv6ClientOptionsServerIdentifier) SetDuidLl(value Dhcpv6ClientOptionsDuidLl) Dhcpv6ClientOptionsServerIdentifier {
	obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LL)
	obj.duidLlHolder = nil
	obj.obj.DuidLl = value.msg()

	return obj
}

// description is TBD
// DuidUuid returns a Dhcpv6ClientOptionsDuidUuid
func (obj *dhcpv6ClientOptionsServerIdentifier) DuidUuid() Dhcpv6ClientOptionsDuidUuid {
	if obj.obj.DuidUuid == nil {
		obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_UUID)
	}
	if obj.duidUuidHolder == nil {
		obj.duidUuidHolder = &dhcpv6ClientOptionsDuidUuid{obj: obj.obj.DuidUuid}
	}
	return obj.duidUuidHolder
}

// description is TBD
// DuidUuid returns a Dhcpv6ClientOptionsDuidUuid
func (obj *dhcpv6ClientOptionsServerIdentifier) HasDuidUuid() bool {
	return obj.obj.DuidUuid != nil
}

// description is TBD
// SetDuidUuid sets the Dhcpv6ClientOptionsDuidUuid value in the Dhcpv6ClientOptionsServerIdentifier object
func (obj *dhcpv6ClientOptionsServerIdentifier) SetDuidUuid(value Dhcpv6ClientOptionsDuidUuid) Dhcpv6ClientOptionsServerIdentifier {
	obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_UUID)
	obj.duidUuidHolder = nil
	obj.obj.DuidUuid = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsServerIdentifier) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DuidLlt != nil {

		obj.DuidLlt().validateObj(vObj, set_default)
	}

	if obj.obj.DuidEn != nil {

		obj.DuidEn().validateObj(vObj, set_default)
	}

	if obj.obj.DuidLl != nil {

		obj.DuidLl().validateObj(vObj, set_default)
	}

	if obj.obj.DuidUuid != nil {

		obj.DuidUuid().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsServerIdentifier) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsServerIdentifierChoiceEnum

	if obj.obj.DuidLlt != nil {
		choices_set += 1
		choice = Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LLT
	}

	if obj.obj.DuidEn != nil {
		choices_set += 1
		choice = Dhcpv6ClientOptionsServerIdentifierChoice.DUID_EN
	}

	if obj.obj.DuidLl != nil {
		choices_set += 1
		choice = Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LL
	}

	if obj.obj.DuidUuid != nil {
		choices_set += 1
		choice = Dhcpv6ClientOptionsServerIdentifierChoice.DUID_UUID
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsServerIdentifierChoice.DUID_LL)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsServerIdentifier")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsServerIdentifier_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsServerIdentifier_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
