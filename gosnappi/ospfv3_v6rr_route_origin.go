package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv3V6RRRouteOrigin *****
type ospfv3V6RRRouteOrigin struct {
	validation
	obj                *otg.Ospfv3V6RRRouteOrigin
	marshaller         marshalOspfv3V6RRRouteOrigin
	unMarshaller       unMarshalOspfv3V6RRRouteOrigin
	nssaExternalHolder Ospfv3V6RRNssaExternal
}

func NewOspfv3V6RRRouteOrigin() Ospfv3V6RRRouteOrigin {
	obj := ospfv3V6RRRouteOrigin{obj: &otg.Ospfv3V6RRRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv3V6RRRouteOrigin) msg() *otg.Ospfv3V6RRRouteOrigin {
	return obj.obj
}

func (obj *ospfv3V6RRRouteOrigin) setMsg(msg *otg.Ospfv3V6RRRouteOrigin) Ospfv3V6RRRouteOrigin {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv3V6RRRouteOrigin struct {
	obj *ospfv3V6RRRouteOrigin
}

type marshalOspfv3V6RRRouteOrigin interface {
	// ToProto marshals Ospfv3V6RRRouteOrigin to protobuf object *otg.Ospfv3V6RRRouteOrigin
	ToProto() (*otg.Ospfv3V6RRRouteOrigin, error)
	// ToPbText marshals Ospfv3V6RRRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv3V6RRRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv3V6RRRouteOrigin to JSON text
	ToJson() (string, error)
}

type unMarshalospfv3V6RRRouteOrigin struct {
	obj *ospfv3V6RRRouteOrigin
}

type unMarshalOspfv3V6RRRouteOrigin interface {
	// FromProto unmarshals Ospfv3V6RRRouteOrigin from protobuf object *otg.Ospfv3V6RRRouteOrigin
	FromProto(msg *otg.Ospfv3V6RRRouteOrigin) (Ospfv3V6RRRouteOrigin, error)
	// FromPbText unmarshals Ospfv3V6RRRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv3V6RRRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv3V6RRRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *ospfv3V6RRRouteOrigin) Marshal() marshalOspfv3V6RRRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv3V6RRRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv3V6RRRouteOrigin) Unmarshal() unMarshalOspfv3V6RRRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv3V6RRRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv3V6RRRouteOrigin) ToProto() (*otg.Ospfv3V6RRRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv3V6RRRouteOrigin) FromProto(msg *otg.Ospfv3V6RRRouteOrigin) (Ospfv3V6RRRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv3V6RRRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalospfv3V6RRRouteOrigin) FromPbText(value string) error {
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

func (m *marshalospfv3V6RRRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalospfv3V6RRRouteOrigin) FromYaml(value string) error {
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

func (m *marshalospfv3V6RRRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalospfv3V6RRRouteOrigin) FromJson(value string) error {
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

func (obj *ospfv3V6RRRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv3V6RRRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv3V6RRRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv3V6RRRouteOrigin) Clone() (Ospfv3V6RRRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv3V6RRRouteOrigin()
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

func (obj *ospfv3V6RRRouteOrigin) setNil() {
	obj.nssaExternalHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv3V6RRRouteOrigin is container of type of the OSPFv3 types correspond directly to the OSPFv3 LSAs types.
type Ospfv3V6RRRouteOrigin interface {
	Validation
	// msg marshals Ospfv3V6RRRouteOrigin to protobuf object *otg.Ospfv3V6RRRouteOrigin
	// and doesn't set defaults
	msg() *otg.Ospfv3V6RRRouteOrigin
	// setMsg unmarshals Ospfv3V6RRRouteOrigin from protobuf object *otg.Ospfv3V6RRRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.Ospfv3V6RRRouteOrigin) Ospfv3V6RRRouteOrigin
	// provides marshal interface
	Marshal() marshalOspfv3V6RRRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv3V6RRRouteOrigin
	// validate validates Ospfv3V6RRRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv3V6RRRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv3V6RRRouteOriginChoiceEnum, set in Ospfv3V6RRRouteOrigin
	Choice() Ospfv3V6RRRouteOriginChoiceEnum
	// setChoice assigns Ospfv3V6RRRouteOriginChoiceEnum provided by user to Ospfv3V6RRRouteOrigin
	setChoice(value Ospfv3V6RRRouteOriginChoiceEnum) Ospfv3V6RRRouteOrigin
	// HasChoice checks if Choice has been set in Ospfv3V6RRRouteOrigin
	HasChoice() bool
	// getter for ExternalType_2 to set choice.
	ExternalType_2()
	// getter for ExternalType_1 to set choice.
	ExternalType_1()
	// getter for IntraArea to set choice.
	IntraArea()
	// getter for InterArea to set choice.
	InterArea()
	// NssaExternal returns Ospfv3V6RRNssaExternal, set in Ospfv3V6RRRouteOrigin.
	// Ospfv3V6RRNssaExternal is container for the forwarding address of NSSA External route origin configuration.
	NssaExternal() Ospfv3V6RRNssaExternal
	// SetNssaExternal assigns Ospfv3V6RRNssaExternal provided by user to Ospfv3V6RRRouteOrigin.
	// Ospfv3V6RRNssaExternal is container for the forwarding address of NSSA External route origin configuration.
	SetNssaExternal(value Ospfv3V6RRNssaExternal) Ospfv3V6RRRouteOrigin
	// HasNssaExternal checks if NssaExternal has been set in Ospfv3V6RRRouteOrigin
	HasNssaExternal() bool
	setNil()
}

type Ospfv3V6RRRouteOriginChoiceEnum string

// Enum of Choice on Ospfv3V6RRRouteOrigin
var Ospfv3V6RRRouteOriginChoice = struct {
	INTRA_AREA      Ospfv3V6RRRouteOriginChoiceEnum
	INTER_AREA      Ospfv3V6RRRouteOriginChoiceEnum
	EXTERNAL_TYPE_1 Ospfv3V6RRRouteOriginChoiceEnum
	EXTERNAL_TYPE_2 Ospfv3V6RRRouteOriginChoiceEnum
	NSSA_EXTERNAL   Ospfv3V6RRRouteOriginChoiceEnum
}{
	INTRA_AREA:      Ospfv3V6RRRouteOriginChoiceEnum("intra_area"),
	INTER_AREA:      Ospfv3V6RRRouteOriginChoiceEnum("inter_area"),
	EXTERNAL_TYPE_1: Ospfv3V6RRRouteOriginChoiceEnum("external_type_1"),
	EXTERNAL_TYPE_2: Ospfv3V6RRRouteOriginChoiceEnum("external_type_2"),
	NSSA_EXTERNAL:   Ospfv3V6RRRouteOriginChoiceEnum("nssa_external"),
}

func (obj *ospfv3V6RRRouteOrigin) Choice() Ospfv3V6RRRouteOriginChoiceEnum {
	return Ospfv3V6RRRouteOriginChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for ExternalType_2 to set choice
func (obj *ospfv3V6RRRouteOrigin) ExternalType_2() {
	obj.setChoice(Ospfv3V6RRRouteOriginChoice.EXTERNAL_TYPE_2)
}

// getter for ExternalType_1 to set choice
func (obj *ospfv3V6RRRouteOrigin) ExternalType_1() {
	obj.setChoice(Ospfv3V6RRRouteOriginChoice.EXTERNAL_TYPE_1)
}

// getter for IntraArea to set choice
func (obj *ospfv3V6RRRouteOrigin) IntraArea() {
	obj.setChoice(Ospfv3V6RRRouteOriginChoice.INTRA_AREA)
}

// getter for InterArea to set choice
func (obj *ospfv3V6RRRouteOrigin) InterArea() {
	obj.setChoice(Ospfv3V6RRRouteOriginChoice.INTER_AREA)
}

// Supported types are: - intra_area: for Intra-Area. - inter_area: for Inter Area. - external_type_1: for Autonomous System (AS) External with internal AS metric. - external_type_2: for Autonomous System (AS) External with internal and external AS metric. - nssa_external: for type 7 Not-So-Stubby Area (NSSA) External.
// Choice returns a string
func (obj *ospfv3V6RRRouteOrigin) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv3V6RRRouteOrigin) setChoice(value Ospfv3V6RRRouteOriginChoiceEnum) Ospfv3V6RRRouteOrigin {
	intValue, ok := otg.Ospfv3V6RRRouteOrigin_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv3V6RRRouteOriginChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv3V6RRRouteOrigin_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.NssaExternal = nil
	obj.nssaExternalHolder = nil

	if value == Ospfv3V6RRRouteOriginChoice.NSSA_EXTERNAL {
		obj.obj.NssaExternal = NewOspfv3V6RRNssaExternal().msg()
	}

	return obj
}

// Configuration for the type 7 Not-So-Stubby Area (NSSA) External.
// NssaExternal returns a Ospfv3V6RRNssaExternal
func (obj *ospfv3V6RRRouteOrigin) NssaExternal() Ospfv3V6RRNssaExternal {
	if obj.obj.NssaExternal == nil {
		obj.setChoice(Ospfv3V6RRRouteOriginChoice.NSSA_EXTERNAL)
	}
	if obj.nssaExternalHolder == nil {
		obj.nssaExternalHolder = &ospfv3V6RRNssaExternal{obj: obj.obj.NssaExternal}
	}
	return obj.nssaExternalHolder
}

// Configuration for the type 7 Not-So-Stubby Area (NSSA) External.
// NssaExternal returns a Ospfv3V6RRNssaExternal
func (obj *ospfv3V6RRRouteOrigin) HasNssaExternal() bool {
	return obj.obj.NssaExternal != nil
}

// Configuration for the type 7 Not-So-Stubby Area (NSSA) External.
// SetNssaExternal sets the Ospfv3V6RRNssaExternal value in the Ospfv3V6RRRouteOrigin object
func (obj *ospfv3V6RRRouteOrigin) SetNssaExternal(value Ospfv3V6RRNssaExternal) Ospfv3V6RRRouteOrigin {
	obj.setChoice(Ospfv3V6RRRouteOriginChoice.NSSA_EXTERNAL)
	obj.nssaExternalHolder = nil
	obj.obj.NssaExternal = value.msg()

	return obj
}

func (obj *ospfv3V6RRRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NssaExternal != nil {

		obj.NssaExternal().validateObj(vObj, set_default)
	}

}

func (obj *ospfv3V6RRRouteOrigin) setDefault() {
	var choices_set int = 0
	var choice Ospfv3V6RRRouteOriginChoiceEnum

	if obj.obj.NssaExternal != nil {
		choices_set += 1
		choice = Ospfv3V6RRRouteOriginChoice.NSSA_EXTERNAL
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv3V6RRRouteOriginChoice.INTER_AREA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv3V6RRRouteOrigin")
			}
		} else {
			intVal := otg.Ospfv3V6RRRouteOrigin_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv3V6RRRouteOrigin_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
