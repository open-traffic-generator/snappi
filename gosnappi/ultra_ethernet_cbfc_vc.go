package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVc *****
type ultraEthernetCbfcVc struct {
	validation
	obj              *otg.UltraEthernetCbfcVc
	marshaller       marshalUltraEthernetCbfcVc
	unMarshaller     unMarshalUltraEthernetCbfcVc
	bestEffortHolder UltraEthernetCbfcVcBestEffort
	losslessHolder   UltraEthernetCbfcVcLossless
	mappingHolder    UltraEthernetCbfcVcMapping
}

func NewUltraEthernetCbfcVc() UltraEthernetCbfcVc {
	obj := ultraEthernetCbfcVc{obj: &otg.UltraEthernetCbfcVc{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVc) msg() *otg.UltraEthernetCbfcVc {
	return obj.obj
}

func (obj *ultraEthernetCbfcVc) setMsg(msg *otg.UltraEthernetCbfcVc) UltraEthernetCbfcVc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVc struct {
	obj *ultraEthernetCbfcVc
}

type marshalUltraEthernetCbfcVc interface {
	// ToProto marshals UltraEthernetCbfcVc to protobuf object *otg.UltraEthernetCbfcVc
	ToProto() (*otg.UltraEthernetCbfcVc, error)
	// ToPbText marshals UltraEthernetCbfcVc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVc to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVc to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVc struct {
	obj *ultraEthernetCbfcVc
}

type unMarshalUltraEthernetCbfcVc interface {
	// FromProto unmarshals UltraEthernetCbfcVc from protobuf object *otg.UltraEthernetCbfcVc
	FromProto(msg *otg.UltraEthernetCbfcVc) (UltraEthernetCbfcVc, error)
	// FromPbText unmarshals UltraEthernetCbfcVc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVc from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVc) Marshal() marshalUltraEthernetCbfcVc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVc{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVc) Unmarshal() unMarshalUltraEthernetCbfcVc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVc) ToProto() (*otg.UltraEthernetCbfcVc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVc) FromProto(msg *otg.UltraEthernetCbfcVc) (UltraEthernetCbfcVc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVc) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVc) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVc) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVc) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVc) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVc) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVc) Clone() (UltraEthernetCbfcVc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVc()
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

func (obj *ultraEthernetCbfcVc) setNil() {
	obj.bestEffortHolder = nil
	obj.losslessHolder = nil
	obj.mappingHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetCbfcVc is a CBFC virtual channel (VC). The position of the VC in the virtual_channels
// list is its VC number (zero based). A maximum of 4 VCs may be configured as
// lossless.
//
// Reference: UE-Specification-1.0.3 Section 5.2.
type UltraEthernetCbfcVc interface {
	Validation
	// msg marshals UltraEthernetCbfcVc to protobuf object *otg.UltraEthernetCbfcVc
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVc
	// setMsg unmarshals UltraEthernetCbfcVc from protobuf object *otg.UltraEthernetCbfcVc
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVc) UltraEthernetCbfcVc
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVc
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVc
	// validate validates UltraEthernetCbfcVc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UltraEthernetCbfcVcChoiceEnum, set in UltraEthernetCbfcVc
	Choice() UltraEthernetCbfcVcChoiceEnum
	// setChoice assigns UltraEthernetCbfcVcChoiceEnum provided by user to UltraEthernetCbfcVc
	setChoice(value UltraEthernetCbfcVcChoiceEnum) UltraEthernetCbfcVc
	// HasChoice checks if Choice has been set in UltraEthernetCbfcVc
	HasChoice() bool
	// BestEffort returns UltraEthernetCbfcVcBestEffort, set in UltraEthernetCbfcVc.
	// UltraEthernetCbfcVcBestEffort is best effort virtual channel settings. A best effort VC does not use CBFC
	// credits.
	BestEffort() UltraEthernetCbfcVcBestEffort
	// SetBestEffort assigns UltraEthernetCbfcVcBestEffort provided by user to UltraEthernetCbfcVc.
	// UltraEthernetCbfcVcBestEffort is best effort virtual channel settings. A best effort VC does not use CBFC
	// credits.
	SetBestEffort(value UltraEthernetCbfcVcBestEffort) UltraEthernetCbfcVc
	// HasBestEffort checks if BestEffort has been set in UltraEthernetCbfcVc
	HasBestEffort() bool
	// Lossless returns UltraEthernetCbfcVcLossless, set in UltraEthernetCbfcVc.
	// UltraEthernetCbfcVcLossless is lossless virtual channel settings.
	Lossless() UltraEthernetCbfcVcLossless
	// SetLossless assigns UltraEthernetCbfcVcLossless provided by user to UltraEthernetCbfcVc.
	// UltraEthernetCbfcVcLossless is lossless virtual channel settings.
	SetLossless(value UltraEthernetCbfcVcLossless) UltraEthernetCbfcVc
	// HasLossless checks if Lossless has been set in UltraEthernetCbfcVc
	HasLossless() bool
	// Mapping returns UltraEthernetCbfcVcMapping, set in UltraEthernetCbfcVc.
	// UltraEthernetCbfcVcMapping is the packet header field used to classify packets into a virtual channel, and
	// the associated values. At a minimum, mapping from the mac.vlan.pcp_dei and
	// ip.dscp header fields is supported.
	Mapping() UltraEthernetCbfcVcMapping
	// SetMapping assigns UltraEthernetCbfcVcMapping provided by user to UltraEthernetCbfcVc.
	// UltraEthernetCbfcVcMapping is the packet header field used to classify packets into a virtual channel, and
	// the associated values. At a minimum, mapping from the mac.vlan.pcp_dei and
	// ip.dscp header fields is supported.
	SetMapping(value UltraEthernetCbfcVcMapping) UltraEthernetCbfcVc
	// HasMapping checks if Mapping has been set in UltraEthernetCbfcVc
	HasMapping() bool
	setNil()
}

type UltraEthernetCbfcVcChoiceEnum string

// Enum of Choice on UltraEthernetCbfcVc
var UltraEthernetCbfcVcChoice = struct {
	BEST_EFFORT UltraEthernetCbfcVcChoiceEnum
	LOSSLESS    UltraEthernetCbfcVcChoiceEnum
}{
	BEST_EFFORT: UltraEthernetCbfcVcChoiceEnum("best_effort"),
	LOSSLESS:    UltraEthernetCbfcVcChoiceEnum("lossless"),
}

func (obj *ultraEthernetCbfcVc) Choice() UltraEthernetCbfcVcChoiceEnum {
	return UltraEthernetCbfcVcChoiceEnum(obj.obj.Choice.Enum().String())
}

// The disposition of the virtual channel.
//
// - best_effort: the VC does not use credits (best effort delivery).
// - lossless: the VC uses CBFC credits for lossless delivery.
// Choice returns a string
func (obj *ultraEthernetCbfcVc) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ultraEthernetCbfcVc) setChoice(value UltraEthernetCbfcVcChoiceEnum) UltraEthernetCbfcVc {
	intValue, ok := otg.UltraEthernetCbfcVc_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetCbfcVcChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetCbfcVc_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Lossless = nil
	obj.losslessHolder = nil
	obj.obj.BestEffort = nil
	obj.bestEffortHolder = nil

	if value == UltraEthernetCbfcVcChoice.BEST_EFFORT {
		obj.obj.BestEffort = NewUltraEthernetCbfcVcBestEffort().msg()
	}

	if value == UltraEthernetCbfcVcChoice.LOSSLESS {
		obj.obj.Lossless = NewUltraEthernetCbfcVcLossless().msg()
	}

	return obj
}

// description is TBD
// BestEffort returns a UltraEthernetCbfcVcBestEffort
func (obj *ultraEthernetCbfcVc) BestEffort() UltraEthernetCbfcVcBestEffort {
	if obj.obj.BestEffort == nil {
		obj.setChoice(UltraEthernetCbfcVcChoice.BEST_EFFORT)
	}
	if obj.bestEffortHolder == nil {
		obj.bestEffortHolder = &ultraEthernetCbfcVcBestEffort{obj: obj.obj.BestEffort}
	}
	return obj.bestEffortHolder
}

// description is TBD
// BestEffort returns a UltraEthernetCbfcVcBestEffort
func (obj *ultraEthernetCbfcVc) HasBestEffort() bool {
	return obj.obj.BestEffort != nil
}

// description is TBD
// SetBestEffort sets the UltraEthernetCbfcVcBestEffort value in the UltraEthernetCbfcVc object
func (obj *ultraEthernetCbfcVc) SetBestEffort(value UltraEthernetCbfcVcBestEffort) UltraEthernetCbfcVc {
	obj.setChoice(UltraEthernetCbfcVcChoice.BEST_EFFORT)
	obj.bestEffortHolder = nil
	obj.obj.BestEffort = value.msg()

	return obj
}

// description is TBD
// Lossless returns a UltraEthernetCbfcVcLossless
func (obj *ultraEthernetCbfcVc) Lossless() UltraEthernetCbfcVcLossless {
	if obj.obj.Lossless == nil {
		obj.setChoice(UltraEthernetCbfcVcChoice.LOSSLESS)
	}
	if obj.losslessHolder == nil {
		obj.losslessHolder = &ultraEthernetCbfcVcLossless{obj: obj.obj.Lossless}
	}
	return obj.losslessHolder
}

// description is TBD
// Lossless returns a UltraEthernetCbfcVcLossless
func (obj *ultraEthernetCbfcVc) HasLossless() bool {
	return obj.obj.Lossless != nil
}

// description is TBD
// SetLossless sets the UltraEthernetCbfcVcLossless value in the UltraEthernetCbfcVc object
func (obj *ultraEthernetCbfcVc) SetLossless(value UltraEthernetCbfcVcLossless) UltraEthernetCbfcVc {
	obj.setChoice(UltraEthernetCbfcVcChoice.LOSSLESS)
	obj.losslessHolder = nil
	obj.obj.Lossless = value.msg()

	return obj
}

// The packet header field values that classify packets into this virtual
// channel.
// Mapping returns a UltraEthernetCbfcVcMapping
func (obj *ultraEthernetCbfcVc) Mapping() UltraEthernetCbfcVcMapping {
	if obj.obj.Mapping == nil {
		obj.obj.Mapping = NewUltraEthernetCbfcVcMapping().msg()
	}
	if obj.mappingHolder == nil {
		obj.mappingHolder = &ultraEthernetCbfcVcMapping{obj: obj.obj.Mapping}
	}
	return obj.mappingHolder
}

// The packet header field values that classify packets into this virtual
// channel.
// Mapping returns a UltraEthernetCbfcVcMapping
func (obj *ultraEthernetCbfcVc) HasMapping() bool {
	return obj.obj.Mapping != nil
}

// The packet header field values that classify packets into this virtual
// channel.
// SetMapping sets the UltraEthernetCbfcVcMapping value in the UltraEthernetCbfcVc object
func (obj *ultraEthernetCbfcVc) SetMapping(value UltraEthernetCbfcVcMapping) UltraEthernetCbfcVc {

	obj.mappingHolder = nil
	obj.obj.Mapping = value.msg()

	return obj
}

func (obj *ultraEthernetCbfcVc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.BestEffort != nil {

		obj.BestEffort().validateObj(vObj, set_default)
	}

	if obj.obj.Lossless != nil {

		obj.Lossless().validateObj(vObj, set_default)
	}

	if obj.obj.Mapping != nil {

		obj.Mapping().validateObj(vObj, set_default)
	}

}

func (obj *ultraEthernetCbfcVc) setDefault() {
	var choices_set int = 0
	var choice UltraEthernetCbfcVcChoiceEnum

	if obj.obj.BestEffort != nil {
		choices_set += 1
		choice = UltraEthernetCbfcVcChoice.BEST_EFFORT
	}

	if obj.obj.Lossless != nil {
		choices_set += 1
		choice = UltraEthernetCbfcVcChoice.LOSSLESS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(UltraEthernetCbfcVcChoice.BEST_EFFORT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UltraEthernetCbfcVc")
			}
		} else {
			intVal := otg.UltraEthernetCbfcVc_Choice_Enum_value[string(choice)]
			enumValue := otg.UltraEthernetCbfcVc_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
