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
	sakPoolHolder MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
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
	// DutSciPortId returns uint32, set in MacsecRxScStaticKey.
	DutSciPortId() uint32
	// SetDutSciPortId assigns uint32 provided by user to MacsecRxScStaticKey
	SetDutSciPortId(value uint32) MacsecRxScStaticKey
	// HasDutSciPortId checks if DutSciPortId has been set in MacsecRxScStaticKey
	HasDutSciPortId() bool
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
	// SakPool returns MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIterIter, set in MacsecRxScStaticKey
	SakPool() MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	setNil()
}

// DUT system ID.
// DutSystemId returns a string
func (obj *macsecRxScStaticKey) DutSystemId() string {

	return *obj.obj.DutSystemId

}

// DUT system ID.
// DutSystemId returns a string
func (obj *macsecRxScStaticKey) HasDutSystemId() bool {
	return obj.obj.DutSystemId != nil
}

// DUT system ID.
// SetDutSystemId sets the string value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetDutSystemId(value string) MacsecRxScStaticKey {

	obj.obj.DutSystemId = &value
	return obj
}

// DUT SCI Port ID.
// DutSciPortId returns a uint32
func (obj *macsecRxScStaticKey) DutSciPortId() uint32 {

	return *obj.obj.DutSciPortId

}

// DUT SCI Port ID.
// DutSciPortId returns a uint32
func (obj *macsecRxScStaticKey) HasDutSciPortId() bool {
	return obj.obj.DutSciPortId != nil
}

// DUT SCI Port ID.
// SetDutSciPortId sets the uint32 value in the MacsecRxScStaticKey object
func (obj *macsecRxScStaticKey) SetDutSciPortId(value uint32) MacsecRxScStaticKey {

	obj.obj.DutSciPortId = &value
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
// SakPool returns a []MacsecBasicKeyGenerationStaticSak
func (obj *macsecRxScStaticKey) SakPool() MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	if len(obj.obj.SakPool) == 0 {
		obj.obj.SakPool = []*otg.MacsecBasicKeyGenerationStaticSak{}
	}
	if obj.sakPoolHolder == nil {
		obj.sakPoolHolder = newMacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter(&obj.obj.SakPool).setMsg(obj)
	}
	return obj.sakPoolHolder
}

type macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter struct {
	obj                                    *macsecRxScStaticKey
	macsecBasicKeyGenerationStaticSakSlice []MacsecBasicKeyGenerationStaticSak
	fieldPtr                               *[]*otg.MacsecBasicKeyGenerationStaticSak
}

func newMacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter(ptr *[]*otg.MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	return &macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter{fieldPtr: ptr}
}

type MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter interface {
	setMsg(*macsecRxScStaticKey) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	Items() []MacsecBasicKeyGenerationStaticSak
	Add() MacsecBasicKeyGenerationStaticSak
	Append(items ...MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	Set(index int, newObj MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	Clear() MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	clearHolderSlice() MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	appendHolderSlice(item MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
}

func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) setMsg(msg *macsecRxScStaticKey) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecBasicKeyGenerationStaticSak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Items() []MacsecBasicKeyGenerationStaticSak {
	return obj.macsecBasicKeyGenerationStaticSakSlice
}

func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Add() MacsecBasicKeyGenerationStaticSak {
	newObj := &otg.MacsecBasicKeyGenerationStaticSak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecBasicKeyGenerationStaticSak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Append(items ...MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, item)
	}
	return obj
}

func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Set(index int, newObj MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecBasicKeyGenerationStaticSakSlice[index] = newObj
	return obj
}
func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Clear() MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecBasicKeyGenerationStaticSak{}
		obj.macsecBasicKeyGenerationStaticSakSlice = []MacsecBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) clearHolderSlice() MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	if len(obj.macsecBasicKeyGenerationStaticSakSlice) > 0 {
		obj.macsecBasicKeyGenerationStaticSakSlice = []MacsecBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) appendHolderSlice(item MacsecBasicKeyGenerationStaticSak) MacsecRxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, item)
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

	if obj.obj.DutSciPortId != nil {

		if *obj.obj.DutSciPortId < 1 || *obj.obj.DutSciPortId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecRxScStaticKey.DutSciPortId <= 65535 but Got %d", *obj.obj.DutSciPortId))
		}

	}

	if obj.obj.ReplayWindow != nil {

		if *obj.obj.ReplayWindow < 1 || *obj.obj.ReplayWindow > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecRxScStaticKey.ReplayWindow <= 4294967295 but Got %d", *obj.obj.ReplayWindow))
		}

	}

	if len(obj.obj.SakPool) != 0 {

		if set_default {
			obj.SakPool().clearHolderSlice()
			for _, item := range obj.obj.SakPool {
				obj.SakPool().appendHolderSlice(&macsecBasicKeyGenerationStaticSak{obj: item})
			}
		}
		for _, item := range obj.SakPool().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *macsecRxScStaticKey) setDefault() {
	if obj.obj.DutSciPortId == nil {
		obj.SetDutSciPortId(1)
	}
	if obj.obj.ReplayProtection == nil {
		obj.SetReplayProtection(false)
	}
	if obj.obj.ReplayWindow == nil {
		obj.SetReplayWindow(1)
	}

}
