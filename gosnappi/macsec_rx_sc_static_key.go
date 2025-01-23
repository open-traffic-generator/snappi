package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecRxScStaticKey *****
type macsecRxScStaticKey struct {
	validation
	obj           *otg.MacsecRxScStaticKey
	marshaller    marshalMacsecRxScStaticKey
	unMarshaller  unMarshalMacsecRxScStaticKey
	sakPoolHolder MacsecBasicKeyGenerationStaticSakPool
}

func NewMacsecRxScStaticKey() MacsecRxScStaticKey {
	obj := macsecRxScStaticKey{obj: &otg.MacsecRxScStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecRxScStaticKey) msg() *otg.MacsecRxScStaticKey {
	return obj.obj
}

func (obj *macsecRxScStaticKey) setMsg(msg *otg.MacsecRxScStaticKey) MacsecRxScStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecRxScStaticKey struct {
	obj *macsecRxScStaticKey
}

type marshalMacsecRxScStaticKey interface {
	// ToProto marshals MacsecRxScStaticKey to protobuf object *otg.MacsecRxScStaticKey
	ToProto() (*otg.MacsecRxScStaticKey, error)
	// ToPbText marshals MacsecRxScStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecRxScStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecRxScStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecRxScStaticKey struct {
	obj *macsecRxScStaticKey
}

type unMarshalMacsecRxScStaticKey interface {
	// FromProto unmarshals MacsecRxScStaticKey from protobuf object *otg.MacsecRxScStaticKey
	FromProto(msg *otg.MacsecRxScStaticKey) (MacsecRxScStaticKey, error)
	// FromPbText unmarshals MacsecRxScStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecRxScStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecRxScStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecRxScStaticKey) Marshal() marshalMacsecRxScStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecRxScStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecRxScStaticKey) Unmarshal() unMarshalMacsecRxScStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecRxScStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecRxScStaticKey) ToProto() (*otg.MacsecRxScStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecRxScStaticKey) FromProto(msg *otg.MacsecRxScStaticKey) (MacsecRxScStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecRxScStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecRxScStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecRxScStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecRxScStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecRxScStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecRxScStaticKey) FromJson(value string) error {
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

func (obj *macsecRxScStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecRxScStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecRxScStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecRxScStaticKey) Clone() (MacsecRxScStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecRxScStaticKey()
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

func (obj *macsecRxScStaticKey) setNil() {
	obj.sakPoolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecRxScStaticKey is rx SC settings for static key.
type MacsecRxScStaticKey interface {
	Validation
	// msg marshals MacsecRxScStaticKey to protobuf object *otg.MacsecRxScStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecRxScStaticKey
	// setMsg unmarshals MacsecRxScStaticKey from protobuf object *otg.MacsecRxScStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecRxScStaticKey) MacsecRxScStaticKey
	// provides marshal interface
	Marshal() marshalMacsecRxScStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecRxScStaticKey
	// validate validates MacsecRxScStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecRxScStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DutSystemId returns string, set in MacsecRxScStaticKey.
	DutSystemId() string
	// SetDutSystemId assigns string provided by user to MacsecRxScStaticKey
	SetDutSystemId(value string) MacsecRxScStaticKey
	// HasDutSystemId checks if DutSystemId has been set in MacsecRxScStaticKey
	HasDutSystemId() bool
	// DutPortId returns uint32, set in MacsecRxScStaticKey.
	DutPortId() uint32
	// SetDutPortId assigns uint32 provided by user to MacsecRxScStaticKey
	SetDutPortId(value uint32) MacsecRxScStaticKey
	// HasDutPortId checks if DutPortId has been set in MacsecRxScStaticKey
	HasDutPortId() bool
	// ReplayProtection returns bool, set in MacsecRxScStaticKey.
	ReplayProtection() bool
	// SetReplayProtection assigns bool provided by user to MacsecRxScStaticKey
	SetReplayProtection(value bool) MacsecRxScStaticKey
	// HasReplayProtection checks if ReplayProtection has been set in MacsecRxScStaticKey
	HasReplayProtection() bool
	// ReplayWindow returns uint32, set in MacsecRxScStaticKey.
	ReplayWindow() uint32
	// SetReplayWindow assigns uint32 provided by user to MacsecRxScStaticKey
	SetReplayWindow(value uint32) MacsecRxScStaticKey
	// HasReplayWindow checks if ReplayWindow has been set in MacsecRxScStaticKey
	HasReplayWindow() bool
	// SakPool returns MacsecBasicKeyGenerationStaticSakPool, set in MacsecRxScStaticKey.
	// MacsecBasicKeyGenerationStaticSakPool is the container for Secure Association Key(SAK) Pool.
	SakPool() MacsecBasicKeyGenerationStaticSakPool
	// SetSakPool assigns MacsecBasicKeyGenerationStaticSakPool provided by user to MacsecRxScStaticKey.
	// MacsecBasicKeyGenerationStaticSakPool is the container for Secure Association Key(SAK) Pool.
	SetSakPool(value MacsecBasicKeyGenerationStaticSakPool) MacsecRxScStaticKey
	// HasSakPool checks if SakPool has been set in MacsecRxScStaticKey
	HasSakPool() bool
	setNil()
}

// System ID in DUT SCI.
// DutSystemId returns a string
func (obj *macsecRxScStaticKey) DutSystemId() string {

	return *obj.obj.DutSystemId

}

// System ID in DUT SCI.
// DutSystemId returns a string
func (obj *macsecRxScStaticKey) HasDutSystemId() bool {
	return obj.obj.DutSystemId != nil
}

// System ID in DUT SCI.
// SetDutSystemId sets the string value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetDutSystemId(value string) MacsecRxScStaticKey {

	obj.obj.DutSystemId = &value
	return obj
}

// Port ID in DUT SCI.
// DutPortId returns a uint32
func (obj *macsecRxScStaticKey) DutPortId() uint32 {

	return *obj.obj.DutPortId

}

// Port ID in DUT SCI.
// DutPortId returns a uint32
func (obj *macsecRxScStaticKey) HasDutPortId() bool {
	return obj.obj.DutPortId != nil
}

// Port ID in DUT SCI.
// SetDutPortId sets the uint32 value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetDutPortId(value uint32) MacsecRxScStaticKey {

	obj.obj.DutPortId = &value
	return obj
}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *macsecRxScStaticKey) ReplayProtection() bool {

	return *obj.obj.ReplayProtection

}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *macsecRxScStaticKey) HasReplayProtection() bool {
	return obj.obj.ReplayProtection != nil
}

// Enable replay protection on not.
// SetReplayProtection sets the bool value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetReplayProtection(value bool) MacsecRxScStaticKey {

	obj.obj.ReplayProtection = &value
	return obj
}

// Replay window size.
// ReplayWindow returns a uint32
func (obj *macsecRxScStaticKey) ReplayWindow() uint32 {

	return *obj.obj.ReplayWindow

}

// Replay window size.
// ReplayWindow returns a uint32
func (obj *macsecRxScStaticKey) HasReplayWindow() bool {
	return obj.obj.ReplayWindow != nil
}

// Replay window size.
// SetReplayWindow sets the uint32 value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetReplayWindow(value uint32) MacsecRxScStaticKey {

	obj.obj.ReplayWindow = &value
	return obj
}

// Rx SAK pool.
// SakPool returns a MacsecBasicKeyGenerationStaticSakPool
func (obj *macsecRxScStaticKey) SakPool() MacsecBasicKeyGenerationStaticSakPool {
	if obj.obj.SakPool == nil {
		obj.obj.SakPool = NewMacsecBasicKeyGenerationStaticSakPool().msg()
	}
	if obj.sakPoolHolder == nil {
		obj.sakPoolHolder = &macsecBasicKeyGenerationStaticSakPool{obj: obj.obj.SakPool}
	}
	return obj.sakPoolHolder
}

// Rx SAK pool.
// SakPool returns a MacsecBasicKeyGenerationStaticSakPool
func (obj *macsecRxScStaticKey) HasSakPool() bool {
	return obj.obj.SakPool != nil
}

// Rx SAK pool.
// SetSakPool sets the MacsecBasicKeyGenerationStaticSakPool value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetSakPool(value MacsecBasicKeyGenerationStaticSakPool) MacsecRxScStaticKey {

	obj.sakPoolHolder = nil
	obj.obj.SakPool = value.msg()

	return obj
}

func (obj *macsecRxScStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DutSystemId != nil {

		err := obj.validateMac(obj.DutSystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecRxScStaticKey.DutSystemId"))
		}

	}

	if obj.obj.DutPortId != nil {

		if *obj.obj.DutPortId < 1 || *obj.obj.DutPortId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecRxScStaticKey.DutPortId <= 65535 but Got %d", *obj.obj.DutPortId))
		}

	}

	if obj.obj.ReplayWindow != nil {

		if *obj.obj.ReplayWindow < 1 || *obj.obj.ReplayWindow > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecRxScStaticKey.ReplayWindow <= 4294967295 but Got %d", *obj.obj.ReplayWindow))
		}

	}

	if obj.obj.SakPool != nil {

		obj.SakPool().validateObj(vObj, set_default)
	}

}

func (obj *macsecRxScStaticKey) setDefault() {
	if obj.obj.DutPortId == nil {
		obj.SetDutPortId(1)
	}
	if obj.obj.ReplayProtection == nil {
		obj.SetReplayProtection(false)
	}
	if obj.obj.ReplayWindow == nil {
		obj.SetReplayWindow(1)
	}

}
