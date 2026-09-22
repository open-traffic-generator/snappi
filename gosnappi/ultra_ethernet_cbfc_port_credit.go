package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcPortCredit *****
type ultraEthernetCbfcPortCredit struct {
	validation
	obj           *otg.UltraEthernetCbfcPortCredit
	marshaller    marshalUltraEthernetCbfcPortCredit
	unMarshaller  unMarshalUltraEthernetCbfcPortCredit
	perPortHolder UltraEthernetCbfcPortCreditPerPort
	perVcHolder   UltraEthernetCbfcPortCreditPerVc
}

func NewUltraEthernetCbfcPortCredit() UltraEthernetCbfcPortCredit {
	obj := ultraEthernetCbfcPortCredit{obj: &otg.UltraEthernetCbfcPortCredit{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcPortCredit) msg() *otg.UltraEthernetCbfcPortCredit {
	return obj.obj
}

func (obj *ultraEthernetCbfcPortCredit) setMsg(msg *otg.UltraEthernetCbfcPortCredit) UltraEthernetCbfcPortCredit {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcPortCredit struct {
	obj *ultraEthernetCbfcPortCredit
}

type marshalUltraEthernetCbfcPortCredit interface {
	// ToProto marshals UltraEthernetCbfcPortCredit to protobuf object *otg.UltraEthernetCbfcPortCredit
	ToProto() (*otg.UltraEthernetCbfcPortCredit, error)
	// ToPbText marshals UltraEthernetCbfcPortCredit to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcPortCredit to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcPortCredit to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcPortCredit struct {
	obj *ultraEthernetCbfcPortCredit
}

type unMarshalUltraEthernetCbfcPortCredit interface {
	// FromProto unmarshals UltraEthernetCbfcPortCredit from protobuf object *otg.UltraEthernetCbfcPortCredit
	FromProto(msg *otg.UltraEthernetCbfcPortCredit) (UltraEthernetCbfcPortCredit, error)
	// FromPbText unmarshals UltraEthernetCbfcPortCredit from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcPortCredit from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcPortCredit from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcPortCredit) Marshal() marshalUltraEthernetCbfcPortCredit {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcPortCredit{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcPortCredit) Unmarshal() unMarshalUltraEthernetCbfcPortCredit {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcPortCredit{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcPortCredit) ToProto() (*otg.UltraEthernetCbfcPortCredit, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcPortCredit) FromProto(msg *otg.UltraEthernetCbfcPortCredit) (UltraEthernetCbfcPortCredit, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcPortCredit) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcPortCredit) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcPortCredit) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcPortCredit) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcPortCredit) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcPortCredit) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcPortCredit) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcPortCredit) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcPortCredit) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcPortCredit) Clone() (UltraEthernetCbfcPortCredit, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcPortCredit()
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

func (obj *ultraEthernetCbfcPortCredit) setNil() {
	obj.perPortHolder = nil
	obj.perVcHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetCbfcPortCredit is cBFC port level credit configuration. Credits are the unit used to track
// available buffer space at the receiver.
//
// Reference: UE-Specification-1.0.3 Table 5-14.
type UltraEthernetCbfcPortCredit interface {
	Validation
	// msg marshals UltraEthernetCbfcPortCredit to protobuf object *otg.UltraEthernetCbfcPortCredit
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcPortCredit
	// setMsg unmarshals UltraEthernetCbfcPortCredit from protobuf object *otg.UltraEthernetCbfcPortCredit
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcPortCredit) UltraEthernetCbfcPortCredit
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcPortCredit
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcPortCredit
	// validate validates UltraEthernetCbfcPortCredit
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcPortCredit, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns UltraEthernetCbfcPortCreditChoiceEnum, set in UltraEthernetCbfcPortCredit
	Choice() UltraEthernetCbfcPortCreditChoiceEnum
	// setChoice assigns UltraEthernetCbfcPortCreditChoiceEnum provided by user to UltraEthernetCbfcPortCredit
	setChoice(value UltraEthernetCbfcPortCreditChoiceEnum) UltraEthernetCbfcPortCredit
	// HasChoice checks if Choice has been set in UltraEthernetCbfcPortCredit
	HasChoice() bool
	// PerPort returns UltraEthernetCbfcPortCreditPerPort, set in UltraEthernetCbfcPortCredit.
	// UltraEthernetCbfcPortCreditPerPort is port based credit limit configuration. A single total credit limit is set and
	// the sender distributes those credits among the lossless VCs.
	PerPort() UltraEthernetCbfcPortCreditPerPort
	// SetPerPort assigns UltraEthernetCbfcPortCreditPerPort provided by user to UltraEthernetCbfcPortCredit.
	// UltraEthernetCbfcPortCreditPerPort is port based credit limit configuration. A single total credit limit is set and
	// the sender distributes those credits among the lossless VCs.
	SetPerPort(value UltraEthernetCbfcPortCreditPerPort) UltraEthernetCbfcPortCredit
	// HasPerPort checks if PerPort has been set in UltraEthernetCbfcPortCredit
	HasPerPort() bool
	// PerVc returns UltraEthernetCbfcPortCreditPerVc, set in UltraEthernetCbfcPortCredit.
	// UltraEthernetCbfcPortCreditPerVc is per virtual channel credit limit configuration. The credit limit for each
	// lossless VC is set individually using the credit_limit field of the VC.
	PerVc() UltraEthernetCbfcPortCreditPerVc
	// SetPerVc assigns UltraEthernetCbfcPortCreditPerVc provided by user to UltraEthernetCbfcPortCredit.
	// UltraEthernetCbfcPortCreditPerVc is per virtual channel credit limit configuration. The credit limit for each
	// lossless VC is set individually using the credit_limit field of the VC.
	SetPerVc(value UltraEthernetCbfcPortCreditPerVc) UltraEthernetCbfcPortCredit
	// HasPerVc checks if PerVc has been set in UltraEthernetCbfcPortCredit
	HasPerVc() bool
	// CellSize returns uint32, set in UltraEthernetCbfcPortCredit.
	CellSize() uint32
	// SetCellSize assigns uint32 provided by user to UltraEthernetCbfcPortCredit
	SetCellSize(value uint32) UltraEthernetCbfcPortCredit
	// HasCellSize checks if CellSize has been set in UltraEthernetCbfcPortCredit
	HasCellSize() bool
	// PacketOverhead returns int32, set in UltraEthernetCbfcPortCredit.
	PacketOverhead() int32
	// SetPacketOverhead assigns int32 provided by user to UltraEthernetCbfcPortCredit
	SetPacketOverhead(value int32) UltraEthernetCbfcPortCredit
	// HasPacketOverhead checks if PacketOverhead has been set in UltraEthernetCbfcPortCredit
	HasPacketOverhead() bool
	setNil()
}

type UltraEthernetCbfcPortCreditChoiceEnum string

// Enum of Choice on UltraEthernetCbfcPortCredit
var UltraEthernetCbfcPortCreditChoice = struct {
	PER_PORT UltraEthernetCbfcPortCreditChoiceEnum
	PER_VC   UltraEthernetCbfcPortCreditChoiceEnum
}{
	PER_PORT: UltraEthernetCbfcPortCreditChoiceEnum("per_port"),
	PER_VC:   UltraEthernetCbfcPortCreditChoiceEnum("per_vc"),
}

func (obj *ultraEthernetCbfcPortCredit) Choice() UltraEthernetCbfcPortCreditChoiceEnum {
	return UltraEthernetCbfcPortCreditChoiceEnum(obj.obj.Choice.Enum().String())
}

// The method used to set credit limits.
//
// - per_port: the receiver sets a single total credit limit and the sender
// distributes those credits among the lossless VCs.
// - per_vc: the receiver sets an individual credit limit for each lossless VC.
// Choice returns a string
func (obj *ultraEthernetCbfcPortCredit) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *ultraEthernetCbfcPortCredit) setChoice(value UltraEthernetCbfcPortCreditChoiceEnum) UltraEthernetCbfcPortCredit {
	intValue, ok := otg.UltraEthernetCbfcPortCredit_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetCbfcPortCreditChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetCbfcPortCredit_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PerVc = nil
	obj.perVcHolder = nil
	obj.obj.PerPort = nil
	obj.perPortHolder = nil

	if value == UltraEthernetCbfcPortCreditChoice.PER_PORT {
		obj.obj.PerPort = NewUltraEthernetCbfcPortCreditPerPort().msg()
	}

	if value == UltraEthernetCbfcPortCreditChoice.PER_VC {
		obj.obj.PerVc = NewUltraEthernetCbfcPortCreditPerVc().msg()
	}

	return obj
}

// description is TBD
// PerPort returns a UltraEthernetCbfcPortCreditPerPort
func (obj *ultraEthernetCbfcPortCredit) PerPort() UltraEthernetCbfcPortCreditPerPort {
	if obj.obj.PerPort == nil {
		obj.setChoice(UltraEthernetCbfcPortCreditChoice.PER_PORT)
	}
	if obj.perPortHolder == nil {
		obj.perPortHolder = &ultraEthernetCbfcPortCreditPerPort{obj: obj.obj.PerPort}
	}
	return obj.perPortHolder
}

// description is TBD
// PerPort returns a UltraEthernetCbfcPortCreditPerPort
func (obj *ultraEthernetCbfcPortCredit) HasPerPort() bool {
	return obj.obj.PerPort != nil
}

// description is TBD
// SetPerPort sets the UltraEthernetCbfcPortCreditPerPort value in the UltraEthernetCbfcPortCredit object
func (obj *ultraEthernetCbfcPortCredit) SetPerPort(value UltraEthernetCbfcPortCreditPerPort) UltraEthernetCbfcPortCredit {
	obj.setChoice(UltraEthernetCbfcPortCreditChoice.PER_PORT)
	obj.perPortHolder = nil
	obj.obj.PerPort = value.msg()

	return obj
}

// description is TBD
// PerVc returns a UltraEthernetCbfcPortCreditPerVc
func (obj *ultraEthernetCbfcPortCredit) PerVc() UltraEthernetCbfcPortCreditPerVc {
	if obj.obj.PerVc == nil {
		obj.setChoice(UltraEthernetCbfcPortCreditChoice.PER_VC)
	}
	if obj.perVcHolder == nil {
		obj.perVcHolder = &ultraEthernetCbfcPortCreditPerVc{obj: obj.obj.PerVc}
	}
	return obj.perVcHolder
}

// description is TBD
// PerVc returns a UltraEthernetCbfcPortCreditPerVc
func (obj *ultraEthernetCbfcPortCredit) HasPerVc() bool {
	return obj.obj.PerVc != nil
}

// description is TBD
// SetPerVc sets the UltraEthernetCbfcPortCreditPerVc value in the UltraEthernetCbfcPortCredit object
func (obj *ultraEthernetCbfcPortCredit) SetPerVc(value UltraEthernetCbfcPortCreditPerVc) UltraEthernetCbfcPortCredit {
	obj.setChoice(UltraEthernetCbfcPortCreditChoice.PER_VC)
	obj.perVcHolder = nil
	obj.obj.PerVc = value.msg()

	return obj
}

// The number of bytes represented by each credit. This is a function of the
// receiver's buffer allocation unit ("cell size") and is implementation
// dependent; a given implementation may support only certain discrete cell
// sizes.
// CellSize returns a uint32
func (obj *ultraEthernetCbfcPortCredit) CellSize() uint32 {

	return *obj.obj.CellSize

}

// The number of bytes represented by each credit. This is a function of the
// receiver's buffer allocation unit ("cell size") and is implementation
// dependent; a given implementation may support only certain discrete cell
// sizes.
// CellSize returns a uint32
func (obj *ultraEthernetCbfcPortCredit) HasCellSize() bool {
	return obj.obj.CellSize != nil
}

// The number of bytes represented by each credit. This is a function of the
// receiver's buffer allocation unit ("cell size") and is implementation
// dependent; a given implementation may support only certain discrete cell
// sizes.
// SetCellSize sets the uint32 value in the UltraEthernetCbfcPortCredit object
func (obj *ultraEthernetCbfcPortCredit) SetCellSize(value uint32) UltraEthernetCbfcPortCredit {

	obj.obj.CellSize = &value
	return obj
}

// The number of bytes of per-packet overhead in the receiver's input buffer.
// May be negative if part of each packet is not stored in the receiver's
// input buffer (for example if the CRC is stripped before storing).
// PacketOverhead returns a int32
func (obj *ultraEthernetCbfcPortCredit) PacketOverhead() int32 {

	return *obj.obj.PacketOverhead

}

// The number of bytes of per-packet overhead in the receiver's input buffer.
// May be negative if part of each packet is not stored in the receiver's
// input buffer (for example if the CRC is stripped before storing).
// PacketOverhead returns a int32
func (obj *ultraEthernetCbfcPortCredit) HasPacketOverhead() bool {
	return obj.obj.PacketOverhead != nil
}

// The number of bytes of per-packet overhead in the receiver's input buffer.
// May be negative if part of each packet is not stored in the receiver's
// input buffer (for example if the CRC is stripped before storing).
// SetPacketOverhead sets the int32 value in the UltraEthernetCbfcPortCredit object
func (obj *ultraEthernetCbfcPortCredit) SetPacketOverhead(value int32) UltraEthernetCbfcPortCredit {

	obj.obj.PacketOverhead = &value
	return obj
}

func (obj *ultraEthernetCbfcPortCredit) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PerPort != nil {

		obj.PerPort().validateObj(vObj, set_default)
	}

	if obj.obj.PerVc != nil {

		obj.PerVc().validateObj(vObj, set_default)
	}

	if obj.obj.CellSize != nil {

		if *obj.obj.CellSize < 32 || *obj.obj.CellSize > 2048 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("32 <= UltraEthernetCbfcPortCredit.CellSize <= 2048 but Got %d", *obj.obj.CellSize))
		}

	}

	if obj.obj.PacketOverhead != nil {

		if *obj.obj.PacketOverhead < -16 || *obj.obj.PacketOverhead > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-16 <= UltraEthernetCbfcPortCredit.PacketOverhead <= 127 but Got %d", *obj.obj.PacketOverhead))
		}

	}

}

func (obj *ultraEthernetCbfcPortCredit) setDefault() {
	var choices_set int = 0
	var choice UltraEthernetCbfcPortCreditChoiceEnum

	if obj.obj.PerPort != nil {
		choices_set += 1
		choice = UltraEthernetCbfcPortCreditChoice.PER_PORT
	}

	if obj.obj.PerVc != nil {
		choices_set += 1
		choice = UltraEthernetCbfcPortCreditChoice.PER_VC
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(UltraEthernetCbfcPortCreditChoice.PER_PORT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UltraEthernetCbfcPortCredit")
			}
		} else {
			intVal := otg.UltraEthernetCbfcPortCredit_Choice_Enum_value[string(choice)]
			enumValue := otg.UltraEthernetCbfcPortCredit_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.CellSize == nil {
		obj.SetCellSize(64)
	}
	if obj.obj.PacketOverhead == nil {
		obj.SetPacketOverhead(-4)
	}

}
