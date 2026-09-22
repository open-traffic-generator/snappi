package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetPhy *****
type ultraEthernetPhy struct {
	validation
	obj                       *otg.UltraEthernetPhy
	marshaller                marshalUltraEthernetPhy
	unMarshaller              unMarshalUltraEthernetPhy
	txControlOrderedSetHolder UltraEthernetPhyTxControlOrderedSet
	rxControlOrderedSetHolder UltraEthernetPhyRxControlOrderedSet
}

func NewUltraEthernetPhy() UltraEthernetPhy {
	obj := ultraEthernetPhy{obj: &otg.UltraEthernetPhy{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetPhy) msg() *otg.UltraEthernetPhy {
	return obj.obj
}

func (obj *ultraEthernetPhy) setMsg(msg *otg.UltraEthernetPhy) UltraEthernetPhy {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetPhy struct {
	obj *ultraEthernetPhy
}

type marshalUltraEthernetPhy interface {
	// ToProto marshals UltraEthernetPhy to protobuf object *otg.UltraEthernetPhy
	ToProto() (*otg.UltraEthernetPhy, error)
	// ToPbText marshals UltraEthernetPhy to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetPhy to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetPhy to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetPhy struct {
	obj *ultraEthernetPhy
}

type unMarshalUltraEthernetPhy interface {
	// FromProto unmarshals UltraEthernetPhy from protobuf object *otg.UltraEthernetPhy
	FromProto(msg *otg.UltraEthernetPhy) (UltraEthernetPhy, error)
	// FromPbText unmarshals UltraEthernetPhy from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetPhy from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetPhy from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetPhy) Marshal() marshalUltraEthernetPhy {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetPhy{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetPhy) Unmarshal() unMarshalUltraEthernetPhy {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetPhy{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetPhy) ToProto() (*otg.UltraEthernetPhy, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetPhy) FromProto(msg *otg.UltraEthernetPhy) (UltraEthernetPhy, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetPhy) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetPhy) FromPbText(value string) error {
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

func (m *marshalultraEthernetPhy) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetPhy) FromYaml(value string) error {
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

func (m *marshalultraEthernetPhy) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetPhy) FromJson(value string) error {
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

func (obj *ultraEthernetPhy) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetPhy) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetPhy) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetPhy) Clone() (UltraEthernetPhy, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetPhy()
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

func (obj *ultraEthernetPhy) setNil() {
	obj.txControlOrderedSetHolder = nil
	obj.rxControlOrderedSetHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetPhy is ultra Ethernet PHY layer settings that control the insertion and validation
// of Control Ordered Sets (CtlOS). CtlOS carry the LLR (LLR_ACK, LLR_NACK,
// LLR_INIT, LLR_INIT_ECHO) and CBFC (CF_Update) messages in the 64B/66B block
// stream.
//
// Reference: UE-Specification-1.0.3 Section 6.2 (Control Ordered Sets); CtlOS
// spacing rules Section 5.1.3.1.1.
type UltraEthernetPhy interface {
	Validation
	// msg marshals UltraEthernetPhy to protobuf object *otg.UltraEthernetPhy
	// and doesn't set defaults
	msg() *otg.UltraEthernetPhy
	// setMsg unmarshals UltraEthernetPhy from protobuf object *otg.UltraEthernetPhy
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetPhy) UltraEthernetPhy
	// provides marshal interface
	Marshal() marshalUltraEthernetPhy
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetPhy
	// validate validates UltraEthernetPhy
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetPhy, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxControlOrderedSet returns UltraEthernetPhyTxControlOrderedSet, set in UltraEthernetPhy.
	// UltraEthernetPhyTxControlOrderedSet is transmit side Control Ordered Set (CtlOS) spacing and timing configuration.
	TxControlOrderedSet() UltraEthernetPhyTxControlOrderedSet
	// SetTxControlOrderedSet assigns UltraEthernetPhyTxControlOrderedSet provided by user to UltraEthernetPhy.
	// UltraEthernetPhyTxControlOrderedSet is transmit side Control Ordered Set (CtlOS) spacing and timing configuration.
	SetTxControlOrderedSet(value UltraEthernetPhyTxControlOrderedSet) UltraEthernetPhy
	// HasTxControlOrderedSet checks if TxControlOrderedSet has been set in UltraEthernetPhy
	HasTxControlOrderedSet() bool
	// RxControlOrderedSet returns UltraEthernetPhyRxControlOrderedSet, set in UltraEthernetPhy.
	// UltraEthernetPhyRxControlOrderedSet is receive side Control Ordered Set (CtlOS) spacing validation configuration.
	// Each validation, when enabled, checks that the received CtlOS spacing honors
	// the configured minimum spacing.
	RxControlOrderedSet() UltraEthernetPhyRxControlOrderedSet
	// SetRxControlOrderedSet assigns UltraEthernetPhyRxControlOrderedSet provided by user to UltraEthernetPhy.
	// UltraEthernetPhyRxControlOrderedSet is receive side Control Ordered Set (CtlOS) spacing validation configuration.
	// Each validation, when enabled, checks that the received CtlOS spacing honors
	// the configured minimum spacing.
	SetRxControlOrderedSet(value UltraEthernetPhyRxControlOrderedSet) UltraEthernetPhy
	// HasRxControlOrderedSet checks if RxControlOrderedSet has been set in UltraEthernetPhy
	HasRxControlOrderedSet() bool
	// OCode returns uint32, set in UltraEthernetPhy.
	OCode() uint32
	// SetOCode assigns uint32 provided by user to UltraEthernetPhy
	SetOCode(value uint32) UltraEthernetPhy
	// HasOCode checks if OCode has been set in UltraEthernetPhy
	HasOCode() bool
	setNil()
}

// Transmit side Control Ordered Set spacing and timing settings.
// TxControlOrderedSet returns a UltraEthernetPhyTxControlOrderedSet
func (obj *ultraEthernetPhy) TxControlOrderedSet() UltraEthernetPhyTxControlOrderedSet {
	if obj.obj.TxControlOrderedSet == nil {
		obj.obj.TxControlOrderedSet = NewUltraEthernetPhyTxControlOrderedSet().msg()
	}
	if obj.txControlOrderedSetHolder == nil {
		obj.txControlOrderedSetHolder = &ultraEthernetPhyTxControlOrderedSet{obj: obj.obj.TxControlOrderedSet}
	}
	return obj.txControlOrderedSetHolder
}

// Transmit side Control Ordered Set spacing and timing settings.
// TxControlOrderedSet returns a UltraEthernetPhyTxControlOrderedSet
func (obj *ultraEthernetPhy) HasTxControlOrderedSet() bool {
	return obj.obj.TxControlOrderedSet != nil
}

// Transmit side Control Ordered Set spacing and timing settings.
// SetTxControlOrderedSet sets the UltraEthernetPhyTxControlOrderedSet value in the UltraEthernetPhy object
func (obj *ultraEthernetPhy) SetTxControlOrderedSet(value UltraEthernetPhyTxControlOrderedSet) UltraEthernetPhy {

	obj.txControlOrderedSetHolder = nil
	obj.obj.TxControlOrderedSet = value.msg()

	return obj
}

// Receive side Control Ordered Set spacing validation settings.
// RxControlOrderedSet returns a UltraEthernetPhyRxControlOrderedSet
func (obj *ultraEthernetPhy) RxControlOrderedSet() UltraEthernetPhyRxControlOrderedSet {
	if obj.obj.RxControlOrderedSet == nil {
		obj.obj.RxControlOrderedSet = NewUltraEthernetPhyRxControlOrderedSet().msg()
	}
	if obj.rxControlOrderedSetHolder == nil {
		obj.rxControlOrderedSetHolder = &ultraEthernetPhyRxControlOrderedSet{obj: obj.obj.RxControlOrderedSet}
	}
	return obj.rxControlOrderedSetHolder
}

// Receive side Control Ordered Set spacing validation settings.
// RxControlOrderedSet returns a UltraEthernetPhyRxControlOrderedSet
func (obj *ultraEthernetPhy) HasRxControlOrderedSet() bool {
	return obj.obj.RxControlOrderedSet != nil
}

// Receive side Control Ordered Set spacing validation settings.
// SetRxControlOrderedSet sets the UltraEthernetPhyRxControlOrderedSet value in the UltraEthernetPhy object
func (obj *ultraEthernetPhy) SetRxControlOrderedSet(value UltraEthernetPhyRxControlOrderedSet) UltraEthernetPhy {

	obj.rxControlOrderedSetHolder = nil
	obj.obj.RxControlOrderedSet = value.msg()

	return obj
}

// The 4-bit O-code used in the CtlOS 64B/66B block format. The UE
// specification uses the value 0x6 (6).
// OCode returns a uint32
func (obj *ultraEthernetPhy) OCode() uint32 {

	return *obj.obj.OCode

}

// The 4-bit O-code used in the CtlOS 64B/66B block format. The UE
// specification uses the value 0x6 (6).
// OCode returns a uint32
func (obj *ultraEthernetPhy) HasOCode() bool {
	return obj.obj.OCode != nil
}

// The 4-bit O-code used in the CtlOS 64B/66B block format. The UE
// specification uses the value 0x6 (6).
// SetOCode sets the uint32 value in the UltraEthernetPhy object
func (obj *ultraEthernetPhy) SetOCode(value uint32) UltraEthernetPhy {

	obj.obj.OCode = &value
	return obj
}

func (obj *ultraEthernetPhy) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxControlOrderedSet != nil {

		obj.TxControlOrderedSet().validateObj(vObj, set_default)
	}

	if obj.obj.RxControlOrderedSet != nil {

		obj.RxControlOrderedSet().validateObj(vObj, set_default)
	}

	if obj.obj.OCode != nil {

		if *obj.obj.OCode > 15 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetPhy.OCode <= 15 but Got %d", *obj.obj.OCode))
		}

	}

}

func (obj *ultraEthernetPhy) setDefault() {
	if obj.obj.OCode == nil {
		obj.SetOCode(6)
	}

}
