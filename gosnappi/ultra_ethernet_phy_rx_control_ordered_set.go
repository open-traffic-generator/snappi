package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetPhyRxControlOrderedSet *****
type ultraEthernetPhyRxControlOrderedSet struct {
	validation
	obj          *otg.UltraEthernetPhyRxControlOrderedSet
	marshaller   marshalUltraEthernetPhyRxControlOrderedSet
	unMarshaller unMarshalUltraEthernetPhyRxControlOrderedSet
}

func NewUltraEthernetPhyRxControlOrderedSet() UltraEthernetPhyRxControlOrderedSet {
	obj := ultraEthernetPhyRxControlOrderedSet{obj: &otg.UltraEthernetPhyRxControlOrderedSet{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetPhyRxControlOrderedSet) msg() *otg.UltraEthernetPhyRxControlOrderedSet {
	return obj.obj
}

func (obj *ultraEthernetPhyRxControlOrderedSet) setMsg(msg *otg.UltraEthernetPhyRxControlOrderedSet) UltraEthernetPhyRxControlOrderedSet {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetPhyRxControlOrderedSet struct {
	obj *ultraEthernetPhyRxControlOrderedSet
}

type marshalUltraEthernetPhyRxControlOrderedSet interface {
	// ToProto marshals UltraEthernetPhyRxControlOrderedSet to protobuf object *otg.UltraEthernetPhyRxControlOrderedSet
	ToProto() (*otg.UltraEthernetPhyRxControlOrderedSet, error)
	// ToPbText marshals UltraEthernetPhyRxControlOrderedSet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetPhyRxControlOrderedSet to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetPhyRxControlOrderedSet to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetPhyRxControlOrderedSet struct {
	obj *ultraEthernetPhyRxControlOrderedSet
}

type unMarshalUltraEthernetPhyRxControlOrderedSet interface {
	// FromProto unmarshals UltraEthernetPhyRxControlOrderedSet from protobuf object *otg.UltraEthernetPhyRxControlOrderedSet
	FromProto(msg *otg.UltraEthernetPhyRxControlOrderedSet) (UltraEthernetPhyRxControlOrderedSet, error)
	// FromPbText unmarshals UltraEthernetPhyRxControlOrderedSet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetPhyRxControlOrderedSet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetPhyRxControlOrderedSet from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetPhyRxControlOrderedSet) Marshal() marshalUltraEthernetPhyRxControlOrderedSet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetPhyRxControlOrderedSet{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetPhyRxControlOrderedSet) Unmarshal() unMarshalUltraEthernetPhyRxControlOrderedSet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetPhyRxControlOrderedSet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetPhyRxControlOrderedSet) ToProto() (*otg.UltraEthernetPhyRxControlOrderedSet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetPhyRxControlOrderedSet) FromProto(msg *otg.UltraEthernetPhyRxControlOrderedSet) (UltraEthernetPhyRxControlOrderedSet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetPhyRxControlOrderedSet) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetPhyRxControlOrderedSet) FromPbText(value string) error {
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

func (m *marshalultraEthernetPhyRxControlOrderedSet) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetPhyRxControlOrderedSet) FromYaml(value string) error {
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

func (m *marshalultraEthernetPhyRxControlOrderedSet) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetPhyRxControlOrderedSet) FromJson(value string) error {
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

func (obj *ultraEthernetPhyRxControlOrderedSet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetPhyRxControlOrderedSet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetPhyRxControlOrderedSet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetPhyRxControlOrderedSet) Clone() (UltraEthernetPhyRxControlOrderedSet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetPhyRxControlOrderedSet()
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

// UltraEthernetPhyRxControlOrderedSet is receive side Control Ordered Set (CtlOS) spacing validation configuration.
// Each validation, when enabled, checks that the received CtlOS spacing honors
// the configured minimum spacing.
type UltraEthernetPhyRxControlOrderedSet interface {
	Validation
	// msg marshals UltraEthernetPhyRxControlOrderedSet to protobuf object *otg.UltraEthernetPhyRxControlOrderedSet
	// and doesn't set defaults
	msg() *otg.UltraEthernetPhyRxControlOrderedSet
	// setMsg unmarshals UltraEthernetPhyRxControlOrderedSet from protobuf object *otg.UltraEthernetPhyRxControlOrderedSet
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetPhyRxControlOrderedSet) UltraEthernetPhyRxControlOrderedSet
	// provides marshal interface
	Marshal() marshalUltraEthernetPhyRxControlOrderedSet
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetPhyRxControlOrderedSet
	// validate validates UltraEthernetPhyRxControlOrderedSet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetPhyRxControlOrderedSet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ValidateMinSpacing returns bool, set in UltraEthernetPhyRxControlOrderedSet.
	ValidateMinSpacing() bool
	// SetValidateMinSpacing assigns bool provided by user to UltraEthernetPhyRxControlOrderedSet
	SetValidateMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet
	// HasValidateMinSpacing checks if ValidateMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasValidateMinSpacing() bool
	// MinSpacing returns uint32, set in UltraEthernetPhyRxControlOrderedSet.
	MinSpacing() uint32
	// SetMinSpacing assigns uint32 provided by user to UltraEthernetPhyRxControlOrderedSet
	SetMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet
	// HasMinSpacing checks if MinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasMinSpacing() bool
	// ValidateLlrInitMinSpacing returns bool, set in UltraEthernetPhyRxControlOrderedSet.
	ValidateLlrInitMinSpacing() bool
	// SetValidateLlrInitMinSpacing assigns bool provided by user to UltraEthernetPhyRxControlOrderedSet
	SetValidateLlrInitMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet
	// HasValidateLlrInitMinSpacing checks if ValidateLlrInitMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasValidateLlrInitMinSpacing() bool
	// LlrInitMinSpacing returns uint32, set in UltraEthernetPhyRxControlOrderedSet.
	LlrInitMinSpacing() uint32
	// SetLlrInitMinSpacing assigns uint32 provided by user to UltraEthernetPhyRxControlOrderedSet
	SetLlrInitMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet
	// HasLlrInitMinSpacing checks if LlrInitMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasLlrInitMinSpacing() bool
	// ValidateLlrAckNackMinSpacing returns bool, set in UltraEthernetPhyRxControlOrderedSet.
	ValidateLlrAckNackMinSpacing() bool
	// SetValidateLlrAckNackMinSpacing assigns bool provided by user to UltraEthernetPhyRxControlOrderedSet
	SetValidateLlrAckNackMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet
	// HasValidateLlrAckNackMinSpacing checks if ValidateLlrAckNackMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasValidateLlrAckNackMinSpacing() bool
	// LlrAckNackMinSpacing returns uint32, set in UltraEthernetPhyRxControlOrderedSet.
	LlrAckNackMinSpacing() uint32
	// SetLlrAckNackMinSpacing assigns uint32 provided by user to UltraEthernetPhyRxControlOrderedSet
	SetLlrAckNackMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet
	// HasLlrAckNackMinSpacing checks if LlrAckNackMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasLlrAckNackMinSpacing() bool
	// ValidateCfMinSpacing returns bool, set in UltraEthernetPhyRxControlOrderedSet.
	ValidateCfMinSpacing() bool
	// SetValidateCfMinSpacing assigns bool provided by user to UltraEthernetPhyRxControlOrderedSet
	SetValidateCfMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet
	// HasValidateCfMinSpacing checks if ValidateCfMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasValidateCfMinSpacing() bool
	// CfMinSpacing returns uint32, set in UltraEthernetPhyRxControlOrderedSet.
	CfMinSpacing() uint32
	// SetCfMinSpacing assigns uint32 provided by user to UltraEthernetPhyRxControlOrderedSet
	SetCfMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet
	// HasCfMinSpacing checks if CfMinSpacing has been set in UltraEthernetPhyRxControlOrderedSet
	HasCfMinSpacing() bool
}

// Enable validation of the minimum spacing between received CtlOS.
// ValidateMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) ValidateMinSpacing() bool {

	return *obj.obj.ValidateMinSpacing

}

// Enable validation of the minimum spacing between received CtlOS.
// ValidateMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) HasValidateMinSpacing() bool {
	return obj.obj.ValidateMinSpacing != nil
}

// Enable validation of the minimum spacing between received CtlOS.
// SetValidateMinSpacing sets the bool value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetValidateMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.ValidateMinSpacing = &value
	return obj
}

// The minimum spacing, in bytes, expected between received CtlOS.
// MinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) MinSpacing() uint32 {

	return *obj.obj.MinSpacing

}

// The minimum spacing, in bytes, expected between received CtlOS.
// MinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) HasMinSpacing() bool {
	return obj.obj.MinSpacing != nil
}

// The minimum spacing, in bytes, expected between received CtlOS.
// SetMinSpacing sets the uint32 value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.MinSpacing = &value
	return obj
}

// Enable validation of the minimum spacing between received LLR_INIT CtlOS.
// ValidateLlrInitMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) ValidateLlrInitMinSpacing() bool {

	return *obj.obj.ValidateLlrInitMinSpacing

}

// Enable validation of the minimum spacing between received LLR_INIT CtlOS.
// ValidateLlrInitMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) HasValidateLlrInitMinSpacing() bool {
	return obj.obj.ValidateLlrInitMinSpacing != nil
}

// Enable validation of the minimum spacing between received LLR_INIT CtlOS.
// SetValidateLlrInitMinSpacing sets the bool value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetValidateLlrInitMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.ValidateLlrInitMinSpacing = &value
	return obj
}

// The minimum spacing, in bytes, expected between received LLR_INIT CtlOS.
// LlrInitMinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) LlrInitMinSpacing() uint32 {

	return *obj.obj.LlrInitMinSpacing

}

// The minimum spacing, in bytes, expected between received LLR_INIT CtlOS.
// LlrInitMinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) HasLlrInitMinSpacing() bool {
	return obj.obj.LlrInitMinSpacing != nil
}

// The minimum spacing, in bytes, expected between received LLR_INIT CtlOS.
// SetLlrInitMinSpacing sets the uint32 value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetLlrInitMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.LlrInitMinSpacing = &value
	return obj
}

// Enable validation of the minimum spacing between received
// LLR_ACK / LLR_NACK CtlOS.
// ValidateLlrAckNackMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) ValidateLlrAckNackMinSpacing() bool {

	return *obj.obj.ValidateLlrAckNackMinSpacing

}

// Enable validation of the minimum spacing between received
// LLR_ACK / LLR_NACK CtlOS.
// ValidateLlrAckNackMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) HasValidateLlrAckNackMinSpacing() bool {
	return obj.obj.ValidateLlrAckNackMinSpacing != nil
}

// Enable validation of the minimum spacing between received
// LLR_ACK / LLR_NACK CtlOS.
// SetValidateLlrAckNackMinSpacing sets the bool value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetValidateLlrAckNackMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.ValidateLlrAckNackMinSpacing = &value
	return obj
}

// The minimum spacing, in bytes, expected between received
// LLR_ACK / LLR_NACK CtlOS.
// LlrAckNackMinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) LlrAckNackMinSpacing() uint32 {

	return *obj.obj.LlrAckNackMinSpacing

}

// The minimum spacing, in bytes, expected between received
// LLR_ACK / LLR_NACK CtlOS.
// LlrAckNackMinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) HasLlrAckNackMinSpacing() bool {
	return obj.obj.LlrAckNackMinSpacing != nil
}

// The minimum spacing, in bytes, expected between received
// LLR_ACK / LLR_NACK CtlOS.
// SetLlrAckNackMinSpacing sets the uint32 value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetLlrAckNackMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.LlrAckNackMinSpacing = &value
	return obj
}

// Enable validation of the minimum spacing between received CBFC
// CF_Update CtlOS.
// ValidateCfMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) ValidateCfMinSpacing() bool {

	return *obj.obj.ValidateCfMinSpacing

}

// Enable validation of the minimum spacing between received CBFC
// CF_Update CtlOS.
// ValidateCfMinSpacing returns a bool
func (obj *ultraEthernetPhyRxControlOrderedSet) HasValidateCfMinSpacing() bool {
	return obj.obj.ValidateCfMinSpacing != nil
}

// Enable validation of the minimum spacing between received CBFC
// CF_Update CtlOS.
// SetValidateCfMinSpacing sets the bool value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetValidateCfMinSpacing(value bool) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.ValidateCfMinSpacing = &value
	return obj
}

// The minimum spacing, in bytes, expected between received CBFC CF_Update
// CtlOS.
// CfMinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) CfMinSpacing() uint32 {

	return *obj.obj.CfMinSpacing

}

// The minimum spacing, in bytes, expected between received CBFC CF_Update
// CtlOS.
// CfMinSpacing returns a uint32
func (obj *ultraEthernetPhyRxControlOrderedSet) HasCfMinSpacing() bool {
	return obj.obj.CfMinSpacing != nil
}

// The minimum spacing, in bytes, expected between received CBFC CF_Update
// CtlOS.
// SetCfMinSpacing sets the uint32 value in the UltraEthernetPhyRxControlOrderedSet object
func (obj *ultraEthernetPhyRxControlOrderedSet) SetCfMinSpacing(value uint32) UltraEthernetPhyRxControlOrderedSet {

	obj.obj.CfMinSpacing = &value
	return obj
}

func (obj *ultraEthernetPhyRxControlOrderedSet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MinSpacing != nil {

		if *obj.obj.MinSpacing > 131064 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetPhyRxControlOrderedSet.MinSpacing <= 131064 but Got %d", *obj.obj.MinSpacing))
		}

	}

	if obj.obj.LlrInitMinSpacing != nil {

		if *obj.obj.LlrInitMinSpacing > 131064 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetPhyRxControlOrderedSet.LlrInitMinSpacing <= 131064 but Got %d", *obj.obj.LlrInitMinSpacing))
		}

	}

	if obj.obj.LlrAckNackMinSpacing != nil {

		if *obj.obj.LlrAckNackMinSpacing > 131064 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetPhyRxControlOrderedSet.LlrAckNackMinSpacing <= 131064 but Got %d", *obj.obj.LlrAckNackMinSpacing))
		}

	}

	if obj.obj.CfMinSpacing != nil {

		if *obj.obj.CfMinSpacing > 131064 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetPhyRxControlOrderedSet.CfMinSpacing <= 131064 but Got %d", *obj.obj.CfMinSpacing))
		}

	}

}

func (obj *ultraEthernetPhyRxControlOrderedSet) setDefault() {
	if obj.obj.ValidateMinSpacing == nil {
		obj.SetValidateMinSpacing(false)
	}
	if obj.obj.MinSpacing == nil {
		obj.SetMinSpacing(400)
	}
	if obj.obj.ValidateLlrInitMinSpacing == nil {
		obj.SetValidateLlrInitMinSpacing(false)
	}
	if obj.obj.LlrInitMinSpacing == nil {
		obj.SetLlrInitMinSpacing(1600)
	}
	if obj.obj.ValidateLlrAckNackMinSpacing == nil {
		obj.SetValidateLlrAckNackMinSpacing(false)
	}
	if obj.obj.LlrAckNackMinSpacing == nil {
		obj.SetLlrAckNackMinSpacing(400)
	}
	if obj.obj.ValidateCfMinSpacing == nil {
		obj.SetValidateCfMinSpacing(false)
	}
	if obj.obj.CfMinSpacing == nil {
		obj.SetCfMinSpacing(400)
	}

}
