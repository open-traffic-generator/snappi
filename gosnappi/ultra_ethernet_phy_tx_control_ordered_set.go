package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetPhyTxControlOrderedSet *****
type ultraEthernetPhyTxControlOrderedSet struct {
	validation
	obj          *otg.UltraEthernetPhyTxControlOrderedSet
	marshaller   marshalUltraEthernetPhyTxControlOrderedSet
	unMarshaller unMarshalUltraEthernetPhyTxControlOrderedSet
}

func NewUltraEthernetPhyTxControlOrderedSet() UltraEthernetPhyTxControlOrderedSet {
	obj := ultraEthernetPhyTxControlOrderedSet{obj: &otg.UltraEthernetPhyTxControlOrderedSet{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetPhyTxControlOrderedSet) msg() *otg.UltraEthernetPhyTxControlOrderedSet {
	return obj.obj
}

func (obj *ultraEthernetPhyTxControlOrderedSet) setMsg(msg *otg.UltraEthernetPhyTxControlOrderedSet) UltraEthernetPhyTxControlOrderedSet {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetPhyTxControlOrderedSet struct {
	obj *ultraEthernetPhyTxControlOrderedSet
}

type marshalUltraEthernetPhyTxControlOrderedSet interface {
	// ToProto marshals UltraEthernetPhyTxControlOrderedSet to protobuf object *otg.UltraEthernetPhyTxControlOrderedSet
	ToProto() (*otg.UltraEthernetPhyTxControlOrderedSet, error)
	// ToPbText marshals UltraEthernetPhyTxControlOrderedSet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetPhyTxControlOrderedSet to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetPhyTxControlOrderedSet to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetPhyTxControlOrderedSet struct {
	obj *ultraEthernetPhyTxControlOrderedSet
}

type unMarshalUltraEthernetPhyTxControlOrderedSet interface {
	// FromProto unmarshals UltraEthernetPhyTxControlOrderedSet from protobuf object *otg.UltraEthernetPhyTxControlOrderedSet
	FromProto(msg *otg.UltraEthernetPhyTxControlOrderedSet) (UltraEthernetPhyTxControlOrderedSet, error)
	// FromPbText unmarshals UltraEthernetPhyTxControlOrderedSet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetPhyTxControlOrderedSet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetPhyTxControlOrderedSet from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetPhyTxControlOrderedSet) Marshal() marshalUltraEthernetPhyTxControlOrderedSet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetPhyTxControlOrderedSet{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetPhyTxControlOrderedSet) Unmarshal() unMarshalUltraEthernetPhyTxControlOrderedSet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetPhyTxControlOrderedSet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetPhyTxControlOrderedSet) ToProto() (*otg.UltraEthernetPhyTxControlOrderedSet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetPhyTxControlOrderedSet) FromProto(msg *otg.UltraEthernetPhyTxControlOrderedSet) (UltraEthernetPhyTxControlOrderedSet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetPhyTxControlOrderedSet) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetPhyTxControlOrderedSet) FromPbText(value string) error {
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

func (m *marshalultraEthernetPhyTxControlOrderedSet) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetPhyTxControlOrderedSet) FromYaml(value string) error {
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

func (m *marshalultraEthernetPhyTxControlOrderedSet) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetPhyTxControlOrderedSet) FromJson(value string) error {
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

func (obj *ultraEthernetPhyTxControlOrderedSet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetPhyTxControlOrderedSet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetPhyTxControlOrderedSet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetPhyTxControlOrderedSet) Clone() (UltraEthernetPhyTxControlOrderedSet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetPhyTxControlOrderedSet()
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

// UltraEthernetPhyTxControlOrderedSet is transmit side Control Ordered Set (CtlOS) spacing and timing configuration.
type UltraEthernetPhyTxControlOrderedSet interface {
	Validation
	// msg marshals UltraEthernetPhyTxControlOrderedSet to protobuf object *otg.UltraEthernetPhyTxControlOrderedSet
	// and doesn't set defaults
	msg() *otg.UltraEthernetPhyTxControlOrderedSet
	// setMsg unmarshals UltraEthernetPhyTxControlOrderedSet from protobuf object *otg.UltraEthernetPhyTxControlOrderedSet
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetPhyTxControlOrderedSet) UltraEthernetPhyTxControlOrderedSet
	// provides marshal interface
	Marshal() marshalUltraEthernetPhyTxControlOrderedSet
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetPhyTxControlOrderedSet
	// validate validates UltraEthernetPhyTxControlOrderedSet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetPhyTxControlOrderedSet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MinSpacing returns uint32, set in UltraEthernetPhyTxControlOrderedSet.
	MinSpacing() uint32
	// SetMinSpacing assigns uint32 provided by user to UltraEthernetPhyTxControlOrderedSet
	SetMinSpacing(value uint32) UltraEthernetPhyTxControlOrderedSet
	// HasMinSpacing checks if MinSpacing has been set in UltraEthernetPhyTxControlOrderedSet
	HasMinSpacing() bool
	// MinSpacingAtFrameStart returns uint32, set in UltraEthernetPhyTxControlOrderedSet.
	MinSpacingAtFrameStart() uint32
	// SetMinSpacingAtFrameStart assigns uint32 provided by user to UltraEthernetPhyTxControlOrderedSet
	SetMinSpacingAtFrameStart(value uint32) UltraEthernetPhyTxControlOrderedSet
	// HasMinSpacingAtFrameStart checks if MinSpacingAtFrameStart has been set in UltraEthernetPhyTxControlOrderedSet
	HasMinSpacingAtFrameStart() bool
	// MinSpacingWithinFrame returns uint32, set in UltraEthernetPhyTxControlOrderedSet.
	MinSpacingWithinFrame() uint32
	// SetMinSpacingWithinFrame assigns uint32 provided by user to UltraEthernetPhyTxControlOrderedSet
	SetMinSpacingWithinFrame(value uint32) UltraEthernetPhyTxControlOrderedSet
	// HasMinSpacingWithinFrame checks if MinSpacingWithinFrame has been set in UltraEthernetPhyTxControlOrderedSet
	HasMinSpacingWithinFrame() bool
	// LlrTargetSpacing returns uint32, set in UltraEthernetPhyTxControlOrderedSet.
	LlrTargetSpacing() uint32
	// SetLlrTargetSpacing assigns uint32 provided by user to UltraEthernetPhyTxControlOrderedSet
	SetLlrTargetSpacing(value uint32) UltraEthernetPhyTxControlOrderedSet
	// HasLlrTargetSpacing checks if LlrTargetSpacing has been set in UltraEthernetPhyTxControlOrderedSet
	HasLlrTargetSpacing() bool
	// CbfcCfMinTimer returns uint32, set in UltraEthernetPhyTxControlOrderedSet.
	CbfcCfMinTimer() uint32
	// SetCbfcCfMinTimer assigns uint32 provided by user to UltraEthernetPhyTxControlOrderedSet
	SetCbfcCfMinTimer(value uint32) UltraEthernetPhyTxControlOrderedSet
	// HasCbfcCfMinTimer checks if CbfcCfMinTimer has been set in UltraEthernetPhyTxControlOrderedSet
	HasCbfcCfMinTimer() bool
	// CbfcCfMaxTimer returns uint32, set in UltraEthernetPhyTxControlOrderedSet.
	CbfcCfMaxTimer() uint32
	// SetCbfcCfMaxTimer assigns uint32 provided by user to UltraEthernetPhyTxControlOrderedSet
	SetCbfcCfMaxTimer(value uint32) UltraEthernetPhyTxControlOrderedSet
	// HasCbfcCfMaxTimer checks if CbfcCfMaxTimer has been set in UltraEthernetPhyTxControlOrderedSet
	HasCbfcCfMaxTimer() bool
}

// The minimum spacing, in bytes, between successive transmitted CtlOS.
// MinSpacing returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) MinSpacing() uint32 {

	return *obj.obj.MinSpacing

}

// The minimum spacing, in bytes, between successive transmitted CtlOS.
// MinSpacing returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) HasMinSpacing() bool {
	return obj.obj.MinSpacing != nil
}

// The minimum spacing, in bytes, between successive transmitted CtlOS.
// SetMinSpacing sets the uint32 value in the UltraEthernetPhyTxControlOrderedSet object
func (obj *ultraEthernetPhyTxControlOrderedSet) SetMinSpacing(value uint32) UltraEthernetPhyTxControlOrderedSet {

	obj.obj.MinSpacing = &value
	return obj
}

// The minimum number of bytes from the start of a frame before a CtlOS
// may be inserted. Used to mitigate the impact on per-frame switch latency.
// MinSpacingAtFrameStart returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) MinSpacingAtFrameStart() uint32 {

	return *obj.obj.MinSpacingAtFrameStart

}

// The minimum number of bytes from the start of a frame before a CtlOS
// may be inserted. Used to mitigate the impact on per-frame switch latency.
// MinSpacingAtFrameStart returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) HasMinSpacingAtFrameStart() bool {
	return obj.obj.MinSpacingAtFrameStart != nil
}

// The minimum number of bytes from the start of a frame before a CtlOS
// may be inserted. Used to mitigate the impact on per-frame switch latency.
// SetMinSpacingAtFrameStart sets the uint32 value in the UltraEthernetPhyTxControlOrderedSet object
func (obj *ultraEthernetPhyTxControlOrderedSet) SetMinSpacingAtFrameStart(value uint32) UltraEthernetPhyTxControlOrderedSet {

	obj.obj.MinSpacingAtFrameStart = &value
	return obj
}

// The minimum spacing, in bytes, between two CtlOS inserted within the same
// frame so that underrun can be managed on cut-through switches.
// MinSpacingWithinFrame returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) MinSpacingWithinFrame() uint32 {

	return *obj.obj.MinSpacingWithinFrame

}

// The minimum spacing, in bytes, between two CtlOS inserted within the same
// frame so that underrun can be managed on cut-through switches.
// MinSpacingWithinFrame returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) HasMinSpacingWithinFrame() bool {
	return obj.obj.MinSpacingWithinFrame != nil
}

// The minimum spacing, in bytes, between two CtlOS inserted within the same
// frame so that underrun can be managed on cut-through switches.
// SetMinSpacingWithinFrame sets the uint32 value in the UltraEthernetPhyTxControlOrderedSet object
func (obj *ultraEthernetPhyTxControlOrderedSet) SetMinSpacingWithinFrame(value uint32) UltraEthernetPhyTxControlOrderedSet {

	obj.obj.MinSpacingWithinFrame = &value
	return obj
}

// The target spacing, in bytes, between successive transmitted
// LLR_ACK / LLR_NACK CtlOS.
// LlrTargetSpacing returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) LlrTargetSpacing() uint32 {

	return *obj.obj.LlrTargetSpacing

}

// The target spacing, in bytes, between successive transmitted
// LLR_ACK / LLR_NACK CtlOS.
// LlrTargetSpacing returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) HasLlrTargetSpacing() bool {
	return obj.obj.LlrTargetSpacing != nil
}

// The target spacing, in bytes, between successive transmitted
// LLR_ACK / LLR_NACK CtlOS.
// SetLlrTargetSpacing sets the uint32 value in the UltraEthernetPhyTxControlOrderedSet object
func (obj *ultraEthernetPhyTxControlOrderedSet) SetLlrTargetSpacing(value uint32) UltraEthernetPhyTxControlOrderedSet {

	obj.obj.LlrTargetSpacing = &value
	return obj
}

// The minimum spacing, in bytes, between successive transmitted CBFC
// CF_Update messages. Guarantees a minimum spacing so that minimal bandwidth
// overhead is used for these messages.
// CbfcCfMinTimer returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) CbfcCfMinTimer() uint32 {

	return *obj.obj.CbfcCfMinTimer

}

// The minimum spacing, in bytes, between successive transmitted CBFC
// CF_Update messages. Guarantees a minimum spacing so that minimal bandwidth
// overhead is used for these messages.
// CbfcCfMinTimer returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) HasCbfcCfMinTimer() bool {
	return obj.obj.CbfcCfMinTimer != nil
}

// The minimum spacing, in bytes, between successive transmitted CBFC
// CF_Update messages. Guarantees a minimum spacing so that minimal bandwidth
// overhead is used for these messages.
// SetCbfcCfMinTimer sets the uint32 value in the UltraEthernetPhyTxControlOrderedSet object
func (obj *ultraEthernetPhyTxControlOrderedSet) SetCbfcCfMinTimer(value uint32) UltraEthernetPhyTxControlOrderedSet {

	obj.obj.CbfcCfMinTimer = &value
	return obj
}

// The maximum spacing, in bytes, between successive transmitted CBFC
// CF_Update messages. Used to periodically refresh credit freed values in
// case a previous CF_Update message was lost.
// CbfcCfMaxTimer returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) CbfcCfMaxTimer() uint32 {

	return *obj.obj.CbfcCfMaxTimer

}

// The maximum spacing, in bytes, between successive transmitted CBFC
// CF_Update messages. Used to periodically refresh credit freed values in
// case a previous CF_Update message was lost.
// CbfcCfMaxTimer returns a uint32
func (obj *ultraEthernetPhyTxControlOrderedSet) HasCbfcCfMaxTimer() bool {
	return obj.obj.CbfcCfMaxTimer != nil
}

// The maximum spacing, in bytes, between successive transmitted CBFC
// CF_Update messages. Used to periodically refresh credit freed values in
// case a previous CF_Update message was lost.
// SetCbfcCfMaxTimer sets the uint32 value in the UltraEthernetPhyTxControlOrderedSet object
func (obj *ultraEthernetPhyTxControlOrderedSet) SetCbfcCfMaxTimer(value uint32) UltraEthernetPhyTxControlOrderedSet {

	obj.obj.CbfcCfMaxTimer = &value
	return obj
}

func (obj *ultraEthernetPhyTxControlOrderedSet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MinSpacing != nil {

		if *obj.obj.MinSpacing < 320 || *obj.obj.MinSpacing > 32760 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("320 <= UltraEthernetPhyTxControlOrderedSet.MinSpacing <= 32760 but Got %d", *obj.obj.MinSpacing))
		}

	}

	if obj.obj.MinSpacingAtFrameStart != nil {

		if *obj.obj.MinSpacingAtFrameStart < 64 || *obj.obj.MinSpacingAtFrameStart > 1008 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("64 <= UltraEthernetPhyTxControlOrderedSet.MinSpacingAtFrameStart <= 1008 but Got %d", *obj.obj.MinSpacingAtFrameStart))
		}

	}

	if obj.obj.MinSpacingWithinFrame != nil {

		if *obj.obj.MinSpacingWithinFrame < 1024 || *obj.obj.MinSpacingWithinFrame > 8184 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1024 <= UltraEthernetPhyTxControlOrderedSet.MinSpacingWithinFrame <= 8184 but Got %d", *obj.obj.MinSpacingWithinFrame))
		}

	}

	if obj.obj.LlrTargetSpacing != nil {

		if *obj.obj.LlrTargetSpacing < 320 || *obj.obj.LlrTargetSpacing > 32760 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("320 <= UltraEthernetPhyTxControlOrderedSet.LlrTargetSpacing <= 32760 but Got %d", *obj.obj.LlrTargetSpacing))
		}

	}

	if obj.obj.CbfcCfMinTimer != nil {

		if *obj.obj.CbfcCfMinTimer < 768 || *obj.obj.CbfcCfMinTimer > 131072 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("768 <= UltraEthernetPhyTxControlOrderedSet.CbfcCfMinTimer <= 131072 but Got %d", *obj.obj.CbfcCfMinTimer))
		}

	}

	if obj.obj.CbfcCfMaxTimer != nil {

		if *obj.obj.CbfcCfMaxTimer < 16384 || *obj.obj.CbfcCfMaxTimer > 1048576 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("16384 <= UltraEthernetPhyTxControlOrderedSet.CbfcCfMaxTimer <= 1048576 but Got %d", *obj.obj.CbfcCfMaxTimer))
		}

	}

}

func (obj *ultraEthernetPhyTxControlOrderedSet) setDefault() {
	if obj.obj.MinSpacing == nil {
		obj.SetMinSpacing(400)
	}
	if obj.obj.MinSpacingAtFrameStart == nil {
		obj.SetMinSpacingAtFrameStart(256)
	}
	if obj.obj.MinSpacingWithinFrame == nil {
		obj.SetMinSpacingWithinFrame(2048)
	}
	if obj.obj.LlrTargetSpacing == nil {
		obj.SetLlrTargetSpacing(2048)
	}
	if obj.obj.CbfcCfMinTimer == nil {
		obj.SetCbfcCfMinTimer(800)
	}
	if obj.obj.CbfcCfMaxTimer == nil {
		obj.SetCbfcCfMaxTimer(65536)
	}

}
