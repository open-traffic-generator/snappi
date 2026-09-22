package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetCbfcVcMetric *****
type ultraEthernetCbfcVcMetric struct {
	validation
	obj          *otg.UltraEthernetCbfcVcMetric
	marshaller   marshalUltraEthernetCbfcVcMetric
	unMarshaller unMarshalUltraEthernetCbfcVcMetric
}

func NewUltraEthernetCbfcVcMetric() UltraEthernetCbfcVcMetric {
	obj := ultraEthernetCbfcVcMetric{obj: &otg.UltraEthernetCbfcVcMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetCbfcVcMetric) msg() *otg.UltraEthernetCbfcVcMetric {
	return obj.obj
}

func (obj *ultraEthernetCbfcVcMetric) setMsg(msg *otg.UltraEthernetCbfcVcMetric) UltraEthernetCbfcVcMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetCbfcVcMetric struct {
	obj *ultraEthernetCbfcVcMetric
}

type marshalUltraEthernetCbfcVcMetric interface {
	// ToProto marshals UltraEthernetCbfcVcMetric to protobuf object *otg.UltraEthernetCbfcVcMetric
	ToProto() (*otg.UltraEthernetCbfcVcMetric, error)
	// ToPbText marshals UltraEthernetCbfcVcMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetCbfcVcMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetCbfcVcMetric to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetCbfcVcMetric struct {
	obj *ultraEthernetCbfcVcMetric
}

type unMarshalUltraEthernetCbfcVcMetric interface {
	// FromProto unmarshals UltraEthernetCbfcVcMetric from protobuf object *otg.UltraEthernetCbfcVcMetric
	FromProto(msg *otg.UltraEthernetCbfcVcMetric) (UltraEthernetCbfcVcMetric, error)
	// FromPbText unmarshals UltraEthernetCbfcVcMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetCbfcVcMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetCbfcVcMetric from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetCbfcVcMetric) Marshal() marshalUltraEthernetCbfcVcMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetCbfcVcMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetCbfcVcMetric) Unmarshal() unMarshalUltraEthernetCbfcVcMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetCbfcVcMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetCbfcVcMetric) ToProto() (*otg.UltraEthernetCbfcVcMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetCbfcVcMetric) FromProto(msg *otg.UltraEthernetCbfcVcMetric) (UltraEthernetCbfcVcMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetCbfcVcMetric) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMetric) FromPbText(value string) error {
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

func (m *marshalultraEthernetCbfcVcMetric) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMetric) FromYaml(value string) error {
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

func (m *marshalultraEthernetCbfcVcMetric) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetCbfcVcMetric) FromJson(value string) error {
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

func (obj *ultraEthernetCbfcVcMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetCbfcVcMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetCbfcVcMetric) Clone() (UltraEthernetCbfcVcMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetCbfcVcMetric()
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

// UltraEthernetCbfcVcMetric is cBFC per virtual channel (VC) credit state. Credit counters per UE-Specification-1.0.3 Tables 5-15/5-16; credit reconciliation Section 5.2.9.
type UltraEthernetCbfcVcMetric interface {
	Validation
	// msg marshals UltraEthernetCbfcVcMetric to protobuf object *otg.UltraEthernetCbfcVcMetric
	// and doesn't set defaults
	msg() *otg.UltraEthernetCbfcVcMetric
	// setMsg unmarshals UltraEthernetCbfcVcMetric from protobuf object *otg.UltraEthernetCbfcVcMetric
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetCbfcVcMetric) UltraEthernetCbfcVcMetric
	// provides marshal interface
	Marshal() marshalUltraEthernetCbfcVcMetric
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetCbfcVcMetric
	// validate validates UltraEthernetCbfcVcMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetCbfcVcMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// VcId returns uint32, set in UltraEthernetCbfcVcMetric.
	VcId() uint32
	// SetVcId assigns uint32 provided by user to UltraEthernetCbfcVcMetric
	SetVcId(value uint32) UltraEthernetCbfcVcMetric
	// HasVcId checks if VcId has been set in UltraEthernetCbfcVcMetric
	HasVcId() bool
	// VcState returns UltraEthernetCbfcVcMetricVcStateEnum, set in UltraEthernetCbfcVcMetric
	VcState() UltraEthernetCbfcVcMetricVcStateEnum
	// SetVcState assigns UltraEthernetCbfcVcMetricVcStateEnum provided by user to UltraEthernetCbfcVcMetric
	SetVcState(value UltraEthernetCbfcVcMetricVcStateEnum) UltraEthernetCbfcVcMetric
	// HasVcState checks if VcState has been set in UltraEthernetCbfcVcMetric
	HasVcState() bool
	// CreditsConsumed returns uint64, set in UltraEthernetCbfcVcMetric.
	CreditsConsumed() uint64
	// SetCreditsConsumed assigns uint64 provided by user to UltraEthernetCbfcVcMetric
	SetCreditsConsumed(value uint64) UltraEthernetCbfcVcMetric
	// HasCreditsConsumed checks if CreditsConsumed has been set in UltraEthernetCbfcVcMetric
	HasCreditsConsumed() bool
	// CreditsFreed returns uint64, set in UltraEthernetCbfcVcMetric.
	CreditsFreed() uint64
	// SetCreditsFreed assigns uint64 provided by user to UltraEthernetCbfcVcMetric
	SetCreditsFreed(value uint64) UltraEthernetCbfcVcMetric
	// HasCreditsFreed checks if CreditsFreed has been set in UltraEthernetCbfcVcMetric
	HasCreditsFreed() bool
	// CreditLimitExceededError returns bool, set in UltraEthernetCbfcVcMetric.
	CreditLimitExceededError() bool
	// SetCreditLimitExceededError assigns bool provided by user to UltraEthernetCbfcVcMetric
	SetCreditLimitExceededError(value bool) UltraEthernetCbfcVcMetric
	// HasCreditLimitExceededError checks if CreditLimitExceededError has been set in UltraEthernetCbfcVcMetric
	HasCreditLimitExceededError() bool
	// ExcessCreditFreedError returns bool, set in UltraEthernetCbfcVcMetric.
	ExcessCreditFreedError() bool
	// SetExcessCreditFreedError assigns bool provided by user to UltraEthernetCbfcVcMetric
	SetExcessCreditFreedError(value bool) UltraEthernetCbfcVcMetric
	// HasExcessCreditFreedError checks if ExcessCreditFreedError has been set in UltraEthernetCbfcVcMetric
	HasExcessCreditFreedError() bool
	// LostCreditError returns bool, set in UltraEthernetCbfcVcMetric.
	LostCreditError() bool
	// SetLostCreditError assigns bool provided by user to UltraEthernetCbfcVcMetric
	SetLostCreditError(value bool) UltraEthernetCbfcVcMetric
	// HasLostCreditError checks if LostCreditError has been set in UltraEthernetCbfcVcMetric
	HasLostCreditError() bool
}

// The zero based VC number.
// VcId returns a uint32
func (obj *ultraEthernetCbfcVcMetric) VcId() uint32 {

	return *obj.obj.VcId

}

// The zero based VC number.
// VcId returns a uint32
func (obj *ultraEthernetCbfcVcMetric) HasVcId() bool {
	return obj.obj.VcId != nil
}

// The zero based VC number.
// SetVcId sets the uint32 value in the UltraEthernetCbfcVcMetric object
func (obj *ultraEthernetCbfcVcMetric) SetVcId(value uint32) UltraEthernetCbfcVcMetric {

	obj.obj.VcId = &value
	return obj
}

type UltraEthernetCbfcVcMetricVcStateEnum string

// Enum of VcState on UltraEthernetCbfcVcMetric
var UltraEthernetCbfcVcMetricVcState = struct {
	BEST_EFFORT UltraEthernetCbfcVcMetricVcStateEnum
	LOSSLESS    UltraEthernetCbfcVcMetricVcStateEnum
}{
	BEST_EFFORT: UltraEthernetCbfcVcMetricVcStateEnum("best_effort"),
	LOSSLESS:    UltraEthernetCbfcVcMetricVcStateEnum("lossless"),
}

func (obj *ultraEthernetCbfcVcMetric) VcState() UltraEthernetCbfcVcMetricVcStateEnum {
	return UltraEthernetCbfcVcMetricVcStateEnum(obj.obj.VcState.Enum().String())
}

// The current operational state of the VC.
// VcState returns a string
func (obj *ultraEthernetCbfcVcMetric) HasVcState() bool {
	return obj.obj.VcState != nil
}

func (obj *ultraEthernetCbfcVcMetric) SetVcState(value UltraEthernetCbfcVcMetricVcStateEnum) UltraEthernetCbfcVcMetric {
	intValue, ok := otg.UltraEthernetCbfcVcMetric_VcState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetCbfcVcMetricVcStateEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetCbfcVcMetric_VcState_Enum(intValue)
	obj.obj.VcState = &enumValue

	return obj
}

// The credits consumed counter for this VC (S_VC_CC for the sender or R_VC_CC for the receiver). A 20-bit cyclic counter (modulo 2^20) per UE-Specification-1.0.3 Section 5.2.
// CreditsConsumed returns a uint64
func (obj *ultraEthernetCbfcVcMetric) CreditsConsumed() uint64 {

	return *obj.obj.CreditsConsumed

}

// The credits consumed counter for this VC (S_VC_CC for the sender or R_VC_CC for the receiver). A 20-bit cyclic counter (modulo 2^20) per UE-Specification-1.0.3 Section 5.2.
// CreditsConsumed returns a uint64
func (obj *ultraEthernetCbfcVcMetric) HasCreditsConsumed() bool {
	return obj.obj.CreditsConsumed != nil
}

// The credits consumed counter for this VC (S_VC_CC for the sender or R_VC_CC for the receiver). A 20-bit cyclic counter (modulo 2^20) per UE-Specification-1.0.3 Section 5.2.
// SetCreditsConsumed sets the uint64 value in the UltraEthernetCbfcVcMetric object
func (obj *ultraEthernetCbfcVcMetric) SetCreditsConsumed(value uint64) UltraEthernetCbfcVcMetric {

	obj.obj.CreditsConsumed = &value
	return obj
}

// The credits freed counter for this VC (S_VC_CF for the sender or R_VC_CF for the receiver). A 15-bit cyclic counter (modulo 2^15) per UE-Specification-1.0.3 Section 5.2.
// CreditsFreed returns a uint64
func (obj *ultraEthernetCbfcVcMetric) CreditsFreed() uint64 {

	return *obj.obj.CreditsFreed

}

// The credits freed counter for this VC (S_VC_CF for the sender or R_VC_CF for the receiver). A 15-bit cyclic counter (modulo 2^15) per UE-Specification-1.0.3 Section 5.2.
// CreditsFreed returns a uint64
func (obj *ultraEthernetCbfcVcMetric) HasCreditsFreed() bool {
	return obj.obj.CreditsFreed != nil
}

// The credits freed counter for this VC (S_VC_CF for the sender or R_VC_CF for the receiver). A 15-bit cyclic counter (modulo 2^15) per UE-Specification-1.0.3 Section 5.2.
// SetCreditsFreed sets the uint64 value in the UltraEthernetCbfcVcMetric object
func (obj *ultraEthernetCbfcVcMetric) SetCreditsFreed(value uint64) UltraEthernetCbfcVcMetric {

	obj.obj.CreditsFreed = &value
	return obj
}

// Indicates whether a credit limit exceeded condition was detected on this VC when the sender credit mode is per VC.
// CreditLimitExceededError returns a bool
func (obj *ultraEthernetCbfcVcMetric) CreditLimitExceededError() bool {

	return *obj.obj.CreditLimitExceededError

}

// Indicates whether a credit limit exceeded condition was detected on this VC when the sender credit mode is per VC.
// CreditLimitExceededError returns a bool
func (obj *ultraEthernetCbfcVcMetric) HasCreditLimitExceededError() bool {
	return obj.obj.CreditLimitExceededError != nil
}

// Indicates whether a credit limit exceeded condition was detected on this VC when the sender credit mode is per VC.
// SetCreditLimitExceededError sets the bool value in the UltraEthernetCbfcVcMetric object
func (obj *ultraEthernetCbfcVcMetric) SetCreditLimitExceededError(value bool) UltraEthernetCbfcVcMetric {

	obj.obj.CreditLimitExceededError = &value
	return obj
}

// Indicates whether an excess credit freed condition was detected on this VC when the sender credit mode is per VC.
// ExcessCreditFreedError returns a bool
func (obj *ultraEthernetCbfcVcMetric) ExcessCreditFreedError() bool {

	return *obj.obj.ExcessCreditFreedError

}

// Indicates whether an excess credit freed condition was detected on this VC when the sender credit mode is per VC.
// ExcessCreditFreedError returns a bool
func (obj *ultraEthernetCbfcVcMetric) HasExcessCreditFreedError() bool {
	return obj.obj.ExcessCreditFreedError != nil
}

// Indicates whether an excess credit freed condition was detected on this VC when the sender credit mode is per VC.
// SetExcessCreditFreedError sets the bool value in the UltraEthernetCbfcVcMetric object
func (obj *ultraEthernetCbfcVcMetric) SetExcessCreditFreedError(value bool) UltraEthernetCbfcVcMetric {

	obj.obj.ExcessCreditFreedError = &value
	return obj
}

// Indicates whether a lost credit condition was detected on this VC when the credit mode is per VC.
// LostCreditError returns a bool
func (obj *ultraEthernetCbfcVcMetric) LostCreditError() bool {

	return *obj.obj.LostCreditError

}

// Indicates whether a lost credit condition was detected on this VC when the credit mode is per VC.
// LostCreditError returns a bool
func (obj *ultraEthernetCbfcVcMetric) HasLostCreditError() bool {
	return obj.obj.LostCreditError != nil
}

// Indicates whether a lost credit condition was detected on this VC when the credit mode is per VC.
// SetLostCreditError sets the bool value in the UltraEthernetCbfcVcMetric object
func (obj *ultraEthernetCbfcVcMetric) SetLostCreditError(value bool) UltraEthernetCbfcVcMetric {

	obj.obj.LostCreditError = &value
	return obj
}

func (obj *ultraEthernetCbfcVcMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *ultraEthernetCbfcVcMetric) setDefault() {

}
