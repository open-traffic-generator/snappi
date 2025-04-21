package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2V4RRRouteOrigin *****
type ospfv2V4RRRouteOrigin struct {
	validation
	obj                  *otg.Ospfv2V4RRRouteOrigin
	marshaller           marshalOspfv2V4RRRouteOrigin
	unMarshaller         unMarshalOspfv2V4RRRouteOrigin
	intraAreaHolder      Ospfv2V4RRIntraArea
	interAreaHolder      Ospfv2V4RRInterArea
	externalType_1Holder Ospfv2V4RRExternalType1
	externalType_2Holder Ospfv2V4RRExternalType2
	nssaExternalHolder   Ospfv2V4RRNssaExternal
}

func NewOspfv2V4RRRouteOrigin() Ospfv2V4RRRouteOrigin {
	obj := ospfv2V4RRRouteOrigin{obj: &otg.Ospfv2V4RRRouteOrigin{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2V4RRRouteOrigin) msg() *otg.Ospfv2V4RRRouteOrigin {
	return obj.obj
}

func (obj *ospfv2V4RRRouteOrigin) setMsg(msg *otg.Ospfv2V4RRRouteOrigin) Ospfv2V4RRRouteOrigin {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2V4RRRouteOrigin struct {
	obj *ospfv2V4RRRouteOrigin
}

type marshalOspfv2V4RRRouteOrigin interface {
	// ToProto marshals Ospfv2V4RRRouteOrigin to protobuf object *otg.Ospfv2V4RRRouteOrigin
	ToProto() (*otg.Ospfv2V4RRRouteOrigin, error)
	// ToPbText marshals Ospfv2V4RRRouteOrigin to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2V4RRRouteOrigin to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2V4RRRouteOrigin to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2V4RRRouteOrigin struct {
	obj *ospfv2V4RRRouteOrigin
}

type unMarshalOspfv2V4RRRouteOrigin interface {
	// FromProto unmarshals Ospfv2V4RRRouteOrigin from protobuf object *otg.Ospfv2V4RRRouteOrigin
	FromProto(msg *otg.Ospfv2V4RRRouteOrigin) (Ospfv2V4RRRouteOrigin, error)
	// FromPbText unmarshals Ospfv2V4RRRouteOrigin from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2V4RRRouteOrigin from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2V4RRRouteOrigin from JSON text
	FromJson(value string) error
}

func (obj *ospfv2V4RRRouteOrigin) Marshal() marshalOspfv2V4RRRouteOrigin {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2V4RRRouteOrigin{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2V4RRRouteOrigin) Unmarshal() unMarshalOspfv2V4RRRouteOrigin {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2V4RRRouteOrigin{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2V4RRRouteOrigin) ToProto() (*otg.Ospfv2V4RRRouteOrigin, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2V4RRRouteOrigin) FromProto(msg *otg.Ospfv2V4RRRouteOrigin) (Ospfv2V4RRRouteOrigin, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2V4RRRouteOrigin) ToPbText() (string, error) {
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

func (m *unMarshalospfv2V4RRRouteOrigin) FromPbText(value string) error {
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

func (m *marshalospfv2V4RRRouteOrigin) ToYaml() (string, error) {
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

func (m *unMarshalospfv2V4RRRouteOrigin) FromYaml(value string) error {
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

func (m *marshalospfv2V4RRRouteOrigin) ToJson() (string, error) {
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

func (m *unMarshalospfv2V4RRRouteOrigin) FromJson(value string) error {
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

func (obj *ospfv2V4RRRouteOrigin) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2V4RRRouteOrigin) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2V4RRRouteOrigin) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2V4RRRouteOrigin) Clone() (Ospfv2V4RRRouteOrigin, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2V4RRRouteOrigin()
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

func (obj *ospfv2V4RRRouteOrigin) setNil() {
	obj.intraAreaHolder = nil
	obj.interAreaHolder = nil
	obj.externalType_1Holder = nil
	obj.externalType_2Holder = nil
	obj.nssaExternalHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2V4RRRouteOrigin is container of type of the OSPFv2 types correspond directly to the OSPFv2 LSAs types as
// defined in the "OSPFv2 Link State (LS) Type - http://www.iana.org/assignments/ospfv2-parameters.
type Ospfv2V4RRRouteOrigin interface {
	Validation
	// msg marshals Ospfv2V4RRRouteOrigin to protobuf object *otg.Ospfv2V4RRRouteOrigin
	// and doesn't set defaults
	msg() *otg.Ospfv2V4RRRouteOrigin
	// setMsg unmarshals Ospfv2V4RRRouteOrigin from protobuf object *otg.Ospfv2V4RRRouteOrigin
	// and doesn't set defaults
	setMsg(*otg.Ospfv2V4RRRouteOrigin) Ospfv2V4RRRouteOrigin
	// provides marshal interface
	Marshal() marshalOspfv2V4RRRouteOrigin
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2V4RRRouteOrigin
	// validate validates Ospfv2V4RRRouteOrigin
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2V4RRRouteOrigin, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Ospfv2V4RRRouteOriginChoiceEnum, set in Ospfv2V4RRRouteOrigin
	Choice() Ospfv2V4RRRouteOriginChoiceEnum
	// setChoice assigns Ospfv2V4RRRouteOriginChoiceEnum provided by user to Ospfv2V4RRRouteOrigin
	setChoice(value Ospfv2V4RRRouteOriginChoiceEnum) Ospfv2V4RRRouteOrigin
	// HasChoice checks if Choice has been set in Ospfv2V4RRRouteOrigin
	HasChoice() bool
	// IntraArea returns Ospfv2V4RRIntraArea, set in Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRIntraArea is container for Intra-Area.
	IntraArea() Ospfv2V4RRIntraArea
	// SetIntraArea assigns Ospfv2V4RRIntraArea provided by user to Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRIntraArea is container for Intra-Area.
	SetIntraArea(value Ospfv2V4RRIntraArea) Ospfv2V4RRRouteOrigin
	// HasIntraArea checks if IntraArea has been set in Ospfv2V4RRRouteOrigin
	HasIntraArea() bool
	// InterArea returns Ospfv2V4RRInterArea, set in Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRInterArea is container for Intra-Area.
	InterArea() Ospfv2V4RRInterArea
	// SetInterArea assigns Ospfv2V4RRInterArea provided by user to Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRInterArea is container for Intra-Area.
	SetInterArea(value Ospfv2V4RRInterArea) Ospfv2V4RRRouteOrigin
	// HasInterArea checks if InterArea has been set in Ospfv2V4RRRouteOrigin
	HasInterArea() bool
	// ExternalType1 returns Ospfv2V4RRExternalType1, set in Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRExternalType1 is container for Intra-Area.
	ExternalType1() Ospfv2V4RRExternalType1
	// SetExternalType1 assigns Ospfv2V4RRExternalType1 provided by user to Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRExternalType1 is container for Intra-Area.
	SetExternalType1(value Ospfv2V4RRExternalType1) Ospfv2V4RRRouteOrigin
	// HasExternalType1 checks if ExternalType1 has been set in Ospfv2V4RRRouteOrigin
	HasExternalType1() bool
	// ExternalType2 returns Ospfv2V4RRExternalType2, set in Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRExternalType2 is container for Intra-Area.
	ExternalType2() Ospfv2V4RRExternalType2
	// SetExternalType2 assigns Ospfv2V4RRExternalType2 provided by user to Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRExternalType2 is container for Intra-Area.
	SetExternalType2(value Ospfv2V4RRExternalType2) Ospfv2V4RRRouteOrigin
	// HasExternalType2 checks if ExternalType2 has been set in Ospfv2V4RRRouteOrigin
	HasExternalType2() bool
	// NssaExternal returns Ospfv2V4RRNssaExternal, set in Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRNssaExternal is container for Intra-Area.
	NssaExternal() Ospfv2V4RRNssaExternal
	// SetNssaExternal assigns Ospfv2V4RRNssaExternal provided by user to Ospfv2V4RRRouteOrigin.
	// Ospfv2V4RRNssaExternal is container for Intra-Area.
	SetNssaExternal(value Ospfv2V4RRNssaExternal) Ospfv2V4RRRouteOrigin
	// HasNssaExternal checks if NssaExternal has been set in Ospfv2V4RRRouteOrigin
	HasNssaExternal() bool
	setNil()
}

type Ospfv2V4RRRouteOriginChoiceEnum string

// Enum of Choice on Ospfv2V4RRRouteOrigin
var Ospfv2V4RRRouteOriginChoice = struct {
	INTRA_AREA      Ospfv2V4RRRouteOriginChoiceEnum
	INTER_AREA      Ospfv2V4RRRouteOriginChoiceEnum
	EXTERNAL_TYPE_1 Ospfv2V4RRRouteOriginChoiceEnum
	EXTERNAL_TYPE_2 Ospfv2V4RRRouteOriginChoiceEnum
	NSSA_EXTERNAL   Ospfv2V4RRRouteOriginChoiceEnum
}{
	INTRA_AREA:      Ospfv2V4RRRouteOriginChoiceEnum("intra_area"),
	INTER_AREA:      Ospfv2V4RRRouteOriginChoiceEnum("inter_area"),
	EXTERNAL_TYPE_1: Ospfv2V4RRRouteOriginChoiceEnum("external_type_1"),
	EXTERNAL_TYPE_2: Ospfv2V4RRRouteOriginChoiceEnum("external_type_2"),
	NSSA_EXTERNAL:   Ospfv2V4RRRouteOriginChoiceEnum("nssa_external"),
}

func (obj *ospfv2V4RRRouteOrigin) Choice() Ospfv2V4RRRouteOriginChoiceEnum {
	return Ospfv2V4RRRouteOriginChoiceEnum(obj.obj.Choice.Enum().String())
}

// Supported types are: - intra_area: for Intra-Area. - inter_area: for Inter Area. - external_type_1: for Autonomous System (AS) External with internal AS meteric. - external_type_2: for Autonomous System (AS) External with internal and external AS meteric. - nssa_external: for 7 Not-So-Stubby Area (NSSA) External.
// Choice returns a string
func (obj *ospfv2V4RRRouteOrigin) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ospfv2V4RRRouteOrigin) setChoice(value Ospfv2V4RRRouteOriginChoiceEnum) Ospfv2V4RRRouteOrigin {
	intValue, ok := otg.Ospfv2V4RRRouteOrigin_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Ospfv2V4RRRouteOriginChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Ospfv2V4RRRouteOrigin_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.NssaExternal = nil
	obj.nssaExternalHolder = nil
	obj.obj.ExternalType_2 = nil
	obj.externalType_2Holder = nil
	obj.obj.ExternalType_1 = nil
	obj.externalType_1Holder = nil
	obj.obj.InterArea = nil
	obj.interAreaHolder = nil
	obj.obj.IntraArea = nil
	obj.intraAreaHolder = nil

	if value == Ospfv2V4RRRouteOriginChoice.INTRA_AREA {
		obj.obj.IntraArea = NewOspfv2V4RRIntraArea().msg()
	}

	if value == Ospfv2V4RRRouteOriginChoice.INTER_AREA {
		obj.obj.InterArea = NewOspfv2V4RRInterArea().msg()
	}

	if value == Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_1 {
		obj.obj.ExternalType_1 = NewOspfv2V4RRExternalType1().msg()
	}

	if value == Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_2 {
		obj.obj.ExternalType_2 = NewOspfv2V4RRExternalType2().msg()
	}

	if value == Ospfv2V4RRRouteOriginChoice.NSSA_EXTERNAL {
		obj.obj.NssaExternal = NewOspfv2V4RRNssaExternal().msg()
	}

	return obj
}

// Configuration for the Intra-Area.
// IntraArea returns a Ospfv2V4RRIntraArea
func (obj *ospfv2V4RRRouteOrigin) IntraArea() Ospfv2V4RRIntraArea {
	if obj.obj.IntraArea == nil {
		obj.setChoice(Ospfv2V4RRRouteOriginChoice.INTRA_AREA)
	}
	if obj.intraAreaHolder == nil {
		obj.intraAreaHolder = &ospfv2V4RRIntraArea{obj: obj.obj.IntraArea}
	}
	return obj.intraAreaHolder
}

// Configuration for the Intra-Area.
// IntraArea returns a Ospfv2V4RRIntraArea
func (obj *ospfv2V4RRRouteOrigin) HasIntraArea() bool {
	return obj.obj.IntraArea != nil
}

// Configuration for the Intra-Area.
// SetIntraArea sets the Ospfv2V4RRIntraArea value in the Ospfv2V4RRRouteOrigin object
func (obj *ospfv2V4RRRouteOrigin) SetIntraArea(value Ospfv2V4RRIntraArea) Ospfv2V4RRRouteOrigin {
	obj.setChoice(Ospfv2V4RRRouteOriginChoice.INTRA_AREA)
	obj.intraAreaHolder = nil
	obj.obj.IntraArea = value.msg()

	return obj
}

// Configuration for the Intra-Area.
// InterArea returns a Ospfv2V4RRInterArea
func (obj *ospfv2V4RRRouteOrigin) InterArea() Ospfv2V4RRInterArea {
	if obj.obj.InterArea == nil {
		obj.setChoice(Ospfv2V4RRRouteOriginChoice.INTER_AREA)
	}
	if obj.interAreaHolder == nil {
		obj.interAreaHolder = &ospfv2V4RRInterArea{obj: obj.obj.InterArea}
	}
	return obj.interAreaHolder
}

// Configuration for the Intra-Area.
// InterArea returns a Ospfv2V4RRInterArea
func (obj *ospfv2V4RRRouteOrigin) HasInterArea() bool {
	return obj.obj.InterArea != nil
}

// Configuration for the Intra-Area.
// SetInterArea sets the Ospfv2V4RRInterArea value in the Ospfv2V4RRRouteOrigin object
func (obj *ospfv2V4RRRouteOrigin) SetInterArea(value Ospfv2V4RRInterArea) Ospfv2V4RRRouteOrigin {
	obj.setChoice(Ospfv2V4RRRouteOriginChoice.INTER_AREA)
	obj.interAreaHolder = nil
	obj.obj.InterArea = value.msg()

	return obj
}

// Configuration for the External Type 1.
// ExternalType1 returns a Ospfv2V4RRExternalType1
func (obj *ospfv2V4RRRouteOrigin) ExternalType1() Ospfv2V4RRExternalType1 {
	if obj.obj.ExternalType_1 == nil {
		obj.setChoice(Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_1)
	}
	if obj.externalType_1Holder == nil {
		obj.externalType_1Holder = &ospfv2V4RRExternalType1{obj: obj.obj.ExternalType_1}
	}
	return obj.externalType_1Holder
}

// Configuration for the External Type 1.
// ExternalType1 returns a Ospfv2V4RRExternalType1
func (obj *ospfv2V4RRRouteOrigin) HasExternalType1() bool {
	return obj.obj.ExternalType_1 != nil
}

// Configuration for the External Type 1.
// SetExternalType1 sets the Ospfv2V4RRExternalType1 value in the Ospfv2V4RRRouteOrigin object
func (obj *ospfv2V4RRRouteOrigin) SetExternalType1(value Ospfv2V4RRExternalType1) Ospfv2V4RRRouteOrigin {
	obj.setChoice(Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_1)
	obj.externalType_1Holder = nil
	obj.obj.ExternalType_1 = value.msg()

	return obj
}

// Configuration for the External Type 2.
// ExternalType2 returns a Ospfv2V4RRExternalType2
func (obj *ospfv2V4RRRouteOrigin) ExternalType2() Ospfv2V4RRExternalType2 {
	if obj.obj.ExternalType_2 == nil {
		obj.setChoice(Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_2)
	}
	if obj.externalType_2Holder == nil {
		obj.externalType_2Holder = &ospfv2V4RRExternalType2{obj: obj.obj.ExternalType_2}
	}
	return obj.externalType_2Holder
}

// Configuration for the External Type 2.
// ExternalType2 returns a Ospfv2V4RRExternalType2
func (obj *ospfv2V4RRRouteOrigin) HasExternalType2() bool {
	return obj.obj.ExternalType_2 != nil
}

// Configuration for the External Type 2.
// SetExternalType2 sets the Ospfv2V4RRExternalType2 value in the Ospfv2V4RRRouteOrigin object
func (obj *ospfv2V4RRRouteOrigin) SetExternalType2(value Ospfv2V4RRExternalType2) Ospfv2V4RRRouteOrigin {
	obj.setChoice(Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_2)
	obj.externalType_2Holder = nil
	obj.obj.ExternalType_2 = value.msg()

	return obj
}

// Configuration for the External Type 2.
// NssaExternal returns a Ospfv2V4RRNssaExternal
func (obj *ospfv2V4RRRouteOrigin) NssaExternal() Ospfv2V4RRNssaExternal {
	if obj.obj.NssaExternal == nil {
		obj.setChoice(Ospfv2V4RRRouteOriginChoice.NSSA_EXTERNAL)
	}
	if obj.nssaExternalHolder == nil {
		obj.nssaExternalHolder = &ospfv2V4RRNssaExternal{obj: obj.obj.NssaExternal}
	}
	return obj.nssaExternalHolder
}

// Configuration for the External Type 2.
// NssaExternal returns a Ospfv2V4RRNssaExternal
func (obj *ospfv2V4RRRouteOrigin) HasNssaExternal() bool {
	return obj.obj.NssaExternal != nil
}

// Configuration for the External Type 2.
// SetNssaExternal sets the Ospfv2V4RRNssaExternal value in the Ospfv2V4RRRouteOrigin object
func (obj *ospfv2V4RRRouteOrigin) SetNssaExternal(value Ospfv2V4RRNssaExternal) Ospfv2V4RRRouteOrigin {
	obj.setChoice(Ospfv2V4RRRouteOriginChoice.NSSA_EXTERNAL)
	obj.nssaExternalHolder = nil
	obj.obj.NssaExternal = value.msg()

	return obj
}

func (obj *ospfv2V4RRRouteOrigin) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IntraArea != nil {

		obj.IntraArea().validateObj(vObj, set_default)
	}

	if obj.obj.InterArea != nil {

		obj.InterArea().validateObj(vObj, set_default)
	}

	if obj.obj.ExternalType_1 != nil {

		obj.ExternalType1().validateObj(vObj, set_default)
	}

	if obj.obj.ExternalType_2 != nil {

		obj.ExternalType2().validateObj(vObj, set_default)
	}

	if obj.obj.NssaExternal != nil {

		obj.NssaExternal().validateObj(vObj, set_default)
	}

}

func (obj *ospfv2V4RRRouteOrigin) setDefault() {
	var choices_set int = 0
	var choice Ospfv2V4RRRouteOriginChoiceEnum

	if obj.obj.IntraArea != nil {
		choices_set += 1
		choice = Ospfv2V4RRRouteOriginChoice.INTRA_AREA
	}

	if obj.obj.InterArea != nil {
		choices_set += 1
		choice = Ospfv2V4RRRouteOriginChoice.INTER_AREA
	}

	if obj.obj.NssaExternal != nil {
		choices_set += 1
		choice = Ospfv2V4RRRouteOriginChoice.NSSA_EXTERNAL
	}

	if obj.obj.ExternalType_1 != nil {
		choices_set += 1
		choice = Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_1
	}

	if obj.obj.ExternalType_2 != nil {
		choices_set += 1
		choice = Ospfv2V4RRRouteOriginChoice.EXTERNAL_TYPE_2
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Ospfv2V4RRRouteOriginChoice.INTER_AREA)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Ospfv2V4RRRouteOrigin")
			}
		} else {
			intVal := otg.Ospfv2V4RRRouteOrigin_Choice_Enum_value[string(choice)]
			enumValue := otg.Ospfv2V4RRRouteOrigin_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
