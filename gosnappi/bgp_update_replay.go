package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpUpdateReplay *****
type bgpUpdateReplay struct {
	validation
	obj                  *otg.BgpUpdateReplay
	marshaller           marshalBgpUpdateReplay
	unMarshaller         unMarshalBgpUpdateReplay
	structuredPdusHolder BgpStructuredPdus
	rawBytesHolder       BgpRawBytes
}

func NewBgpUpdateReplay() BgpUpdateReplay {
	obj := bgpUpdateReplay{obj: &otg.BgpUpdateReplay{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpUpdateReplay) msg() *otg.BgpUpdateReplay {
	return obj.obj
}

func (obj *bgpUpdateReplay) setMsg(msg *otg.BgpUpdateReplay) BgpUpdateReplay {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpUpdateReplay struct {
	obj *bgpUpdateReplay
}

type marshalBgpUpdateReplay interface {
	// ToProto marshals BgpUpdateReplay to protobuf object *otg.BgpUpdateReplay
	ToProto() (*otg.BgpUpdateReplay, error)
	// ToPbText marshals BgpUpdateReplay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpUpdateReplay to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpUpdateReplay to JSON text
	ToJson() (string, error)
}

type unMarshalbgpUpdateReplay struct {
	obj *bgpUpdateReplay
}

type unMarshalBgpUpdateReplay interface {
	// FromProto unmarshals BgpUpdateReplay from protobuf object *otg.BgpUpdateReplay
	FromProto(msg *otg.BgpUpdateReplay) (BgpUpdateReplay, error)
	// FromPbText unmarshals BgpUpdateReplay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpUpdateReplay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpUpdateReplay from JSON text
	FromJson(value string) error
}

func (obj *bgpUpdateReplay) Marshal() marshalBgpUpdateReplay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpUpdateReplay{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpUpdateReplay) Unmarshal() unMarshalBgpUpdateReplay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpUpdateReplay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpUpdateReplay) ToProto() (*otg.BgpUpdateReplay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpUpdateReplay) FromProto(msg *otg.BgpUpdateReplay) (BgpUpdateReplay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpUpdateReplay) ToPbText() (string, error) {
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

func (m *unMarshalbgpUpdateReplay) FromPbText(value string) error {
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

func (m *marshalbgpUpdateReplay) ToYaml() (string, error) {
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

func (m *unMarshalbgpUpdateReplay) FromYaml(value string) error {
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

func (m *marshalbgpUpdateReplay) ToJson() (string, error) {
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

func (m *unMarshalbgpUpdateReplay) FromJson(value string) error {
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

func (obj *bgpUpdateReplay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpUpdateReplay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpUpdateReplay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpUpdateReplay) Clone() (BgpUpdateReplay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpUpdateReplay()
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

func (obj *bgpUpdateReplay) setNil() {
	obj.structuredPdusHolder = nil
	obj.rawBytesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpUpdateReplay is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
type BgpUpdateReplay interface {
	Validation
	// msg marshals BgpUpdateReplay to protobuf object *otg.BgpUpdateReplay
	// and doesn't set defaults
	msg() *otg.BgpUpdateReplay
	// setMsg unmarshals BgpUpdateReplay from protobuf object *otg.BgpUpdateReplay
	// and doesn't set defaults
	setMsg(*otg.BgpUpdateReplay) BgpUpdateReplay
	// provides marshal interface
	Marshal() marshalBgpUpdateReplay
	// provides unmarshal interface
	Unmarshal() unMarshalBgpUpdateReplay
	// validate validates BgpUpdateReplay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpUpdateReplay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns BgpUpdateReplayChoiceEnum, set in BgpUpdateReplay
	Choice() BgpUpdateReplayChoiceEnum
	// setChoice assigns BgpUpdateReplayChoiceEnum provided by user to BgpUpdateReplay
	setChoice(value BgpUpdateReplayChoiceEnum) BgpUpdateReplay
	// HasChoice checks if Choice has been set in BgpUpdateReplay
	HasChoice() bool
	// StructuredPdus returns BgpStructuredPdus, set in BgpUpdateReplay.
	// BgpStructuredPdus is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	StructuredPdus() BgpStructuredPdus
	// SetStructuredPdus assigns BgpStructuredPdus provided by user to BgpUpdateReplay.
	// BgpStructuredPdus is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	SetStructuredPdus(value BgpStructuredPdus) BgpUpdateReplay
	// HasStructuredPdus checks if StructuredPdus has been set in BgpUpdateReplay
	HasStructuredPdus() bool
	// RawBytes returns BgpRawBytes, set in BgpUpdateReplay.
	// BgpRawBytes is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	RawBytes() BgpRawBytes
	// SetRawBytes assigns BgpRawBytes provided by user to BgpUpdateReplay.
	// BgpRawBytes is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
	SetRawBytes(value BgpRawBytes) BgpUpdateReplay
	// HasRawBytes checks if RawBytes has been set in BgpUpdateReplay
	HasRawBytes() bool
	setNil()
}

type BgpUpdateReplayChoiceEnum string

// Enum of Choice on BgpUpdateReplay
var BgpUpdateReplayChoice = struct {
	STRUCTURED_PDUS BgpUpdateReplayChoiceEnum
	RAW_BYTES       BgpUpdateReplayChoiceEnum
}{
	STRUCTURED_PDUS: BgpUpdateReplayChoiceEnum("structured_pdus"),
	RAW_BYTES:       BgpUpdateReplayChoiceEnum("raw_bytes"),
}

func (obj *bgpUpdateReplay) Choice() BgpUpdateReplayChoiceEnum {
	return BgpUpdateReplayChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *bgpUpdateReplay) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *bgpUpdateReplay) setChoice(value BgpUpdateReplayChoiceEnum) BgpUpdateReplay {
	intValue, ok := otg.BgpUpdateReplay_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpUpdateReplayChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpUpdateReplay_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.RawBytes = nil
	obj.rawBytesHolder = nil
	obj.obj.StructuredPdus = nil
	obj.structuredPdusHolder = nil

	if value == BgpUpdateReplayChoice.STRUCTURED_PDUS {
		obj.obj.StructuredPdus = NewBgpStructuredPdus().msg()
	}

	if value == BgpUpdateReplayChoice.RAW_BYTES {
		obj.obj.RawBytes = NewBgpRawBytes().msg()
	}

	return obj
}

// description is TBD
// StructuredPdus returns a BgpStructuredPdus
func (obj *bgpUpdateReplay) StructuredPdus() BgpStructuredPdus {
	if obj.obj.StructuredPdus == nil {
		obj.setChoice(BgpUpdateReplayChoice.STRUCTURED_PDUS)
	}
	if obj.structuredPdusHolder == nil {
		obj.structuredPdusHolder = &bgpStructuredPdus{obj: obj.obj.StructuredPdus}
	}
	return obj.structuredPdusHolder
}

// description is TBD
// StructuredPdus returns a BgpStructuredPdus
func (obj *bgpUpdateReplay) HasStructuredPdus() bool {
	return obj.obj.StructuredPdus != nil
}

// description is TBD
// SetStructuredPdus sets the BgpStructuredPdus value in the BgpUpdateReplay object
func (obj *bgpUpdateReplay) SetStructuredPdus(value BgpStructuredPdus) BgpUpdateReplay {
	obj.setChoice(BgpUpdateReplayChoice.STRUCTURED_PDUS)
	obj.structuredPdusHolder = nil
	obj.obj.StructuredPdus = value.msg()

	return obj
}

// description is TBD
// RawBytes returns a BgpRawBytes
func (obj *bgpUpdateReplay) RawBytes() BgpRawBytes {
	if obj.obj.RawBytes == nil {
		obj.setChoice(BgpUpdateReplayChoice.RAW_BYTES)
	}
	if obj.rawBytesHolder == nil {
		obj.rawBytesHolder = &bgpRawBytes{obj: obj.obj.RawBytes}
	}
	return obj.rawBytesHolder
}

// description is TBD
// RawBytes returns a BgpRawBytes
func (obj *bgpUpdateReplay) HasRawBytes() bool {
	return obj.obj.RawBytes != nil
}

// description is TBD
// SetRawBytes sets the BgpRawBytes value in the BgpUpdateReplay object
func (obj *bgpUpdateReplay) SetRawBytes(value BgpRawBytes) BgpUpdateReplay {
	obj.setChoice(BgpUpdateReplayChoice.RAW_BYTES)
	obj.rawBytesHolder = nil
	obj.obj.RawBytes = value.msg()

	return obj
}

func (obj *bgpUpdateReplay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.StructuredPdus != nil {

		obj.StructuredPdus().validateObj(vObj, set_default)
	}

	if obj.obj.RawBytes != nil {

		obj.RawBytes().validateObj(vObj, set_default)
	}

}

func (obj *bgpUpdateReplay) setDefault() {
	var choices_set int = 0
	var choice BgpUpdateReplayChoiceEnum

	if obj.obj.StructuredPdus != nil {
		choices_set += 1
		choice = BgpUpdateReplayChoice.STRUCTURED_PDUS
	}

	if obj.obj.RawBytes != nil {
		choices_set += 1
		choice = BgpUpdateReplayChoice.RAW_BYTES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(BgpUpdateReplayChoice.STRUCTURED_PDUS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in BgpUpdateReplay")
			}
		} else {
			intVal := otg.BgpUpdateReplay_Choice_Enum_value[string(choice)]
			enumValue := otg.BgpUpdateReplay_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
