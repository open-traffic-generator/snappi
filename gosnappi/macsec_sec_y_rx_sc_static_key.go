package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYRxScStaticKey *****
type macsecSecYRxScStaticKey struct {
	validation
	obj           *otg.MacsecSecYRxScStaticKey
	marshaller    marshalMacsecSecYRxScStaticKey
	unMarshaller  unMarshalMacsecSecYRxScStaticKey
	sakPoolHolder MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
}

func NewMacsecSecYRxScStaticKey() MacsecSecYRxScStaticKey {
	obj := macsecSecYRxScStaticKey{obj: &otg.MacsecSecYRxScStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYRxScStaticKey) msg() *otg.MacsecSecYRxScStaticKey {
	return obj.obj
}

func (obj *macsecSecYRxScStaticKey) setMsg(msg *otg.MacsecSecYRxScStaticKey) MacsecSecYRxScStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYRxScStaticKey struct {
	obj *macsecSecYRxScStaticKey
}

type marshalMacsecSecYRxScStaticKey interface {
	// ToProto marshals MacsecSecYRxScStaticKey to protobuf object *otg.MacsecSecYRxScStaticKey
	ToProto() (*otg.MacsecSecYRxScStaticKey, error)
	// ToPbText marshals MacsecSecYRxScStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYRxScStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYRxScStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYRxScStaticKey struct {
	obj *macsecSecYRxScStaticKey
}

type unMarshalMacsecSecYRxScStaticKey interface {
	// FromProto unmarshals MacsecSecYRxScStaticKey from protobuf object *otg.MacsecSecYRxScStaticKey
	FromProto(msg *otg.MacsecSecYRxScStaticKey) (MacsecSecYRxScStaticKey, error)
	// FromPbText unmarshals MacsecSecYRxScStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYRxScStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYRxScStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYRxScStaticKey) Marshal() marshalMacsecSecYRxScStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYRxScStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYRxScStaticKey) Unmarshal() unMarshalMacsecSecYRxScStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYRxScStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYRxScStaticKey) ToProto() (*otg.MacsecSecYRxScStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYRxScStaticKey) FromProto(msg *otg.MacsecSecYRxScStaticKey) (MacsecSecYRxScStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYRxScStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYRxScStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecSecYRxScStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYRxScStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecSecYRxScStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYRxScStaticKey) FromJson(value string) error {
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

func (obj *macsecSecYRxScStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYRxScStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYRxScStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYRxScStaticKey) Clone() (MacsecSecYRxScStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYRxScStaticKey()
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

func (obj *macsecSecYRxScStaticKey) setNil() {
	obj.sakPoolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYRxScStaticKey is rx SC settings for static key.
type MacsecSecYRxScStaticKey interface {
	Validation
	// msg marshals MacsecSecYRxScStaticKey to protobuf object *otg.MacsecSecYRxScStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecSecYRxScStaticKey
	// setMsg unmarshals MacsecSecYRxScStaticKey from protobuf object *otg.MacsecSecYRxScStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYRxScStaticKey) MacsecSecYRxScStaticKey
	// provides marshal interface
	Marshal() marshalMacsecSecYRxScStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYRxScStaticKey
	// validate validates MacsecSecYRxScStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYRxScStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DutSystemId returns string, set in MacsecSecYRxScStaticKey.
	DutSystemId() string
	// SetDutSystemId assigns string provided by user to MacsecSecYRxScStaticKey
	SetDutSystemId(value string) MacsecSecYRxScStaticKey
	// HasDutSystemId checks if DutSystemId has been set in MacsecSecYRxScStaticKey
	HasDutSystemId() bool
	// DutSciPortId returns int32, set in MacsecSecYRxScStaticKey.
	DutSciPortId() int32
	// SetDutSciPortId assigns int32 provided by user to MacsecSecYRxScStaticKey
	SetDutSciPortId(value int32) MacsecSecYRxScStaticKey
	// HasDutSciPortId checks if DutSciPortId has been set in MacsecSecYRxScStaticKey
	HasDutSciPortId() bool
	// ReplayProtection returns bool, set in MacsecSecYRxScStaticKey.
	ReplayProtection() bool
	// SetReplayProtection assigns bool provided by user to MacsecSecYRxScStaticKey
	SetReplayProtection(value bool) MacsecSecYRxScStaticKey
	// HasReplayProtection checks if ReplayProtection has been set in MacsecSecYRxScStaticKey
	HasReplayProtection() bool
	// ReplayWindow returns int32, set in MacsecSecYRxScStaticKey.
	ReplayWindow() int32
	// SetReplayWindow assigns int32 provided by user to MacsecSecYRxScStaticKey
	SetReplayWindow(value int32) MacsecSecYRxScStaticKey
	// HasReplayWindow checks if ReplayWindow has been set in MacsecSecYRxScStaticKey
	HasReplayWindow() bool
	// SakPool returns MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIterIter, set in MacsecSecYRxScStaticKey
	SakPool() MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	setNil()
}

// DUT system ID.
// DutSystemId returns a string
func (obj *macsecSecYRxScStaticKey) DutSystemId() string {

	return *obj.obj.DutSystemId

}

// DUT system ID.
// DutSystemId returns a string
func (obj *macsecSecYRxScStaticKey) HasDutSystemId() bool {
	return obj.obj.DutSystemId != nil
}

// DUT system ID.
// SetDutSystemId sets the string value in the MacsecSecYRxScStaticKey object
func (obj *macsecSecYRxScStaticKey) SetDutSystemId(value string) MacsecSecYRxScStaticKey {

	obj.obj.DutSystemId = &value
	return obj
}

// DUT SCI Port ID.
// DutSciPortId returns a int32
func (obj *macsecSecYRxScStaticKey) DutSciPortId() int32 {

	return *obj.obj.DutSciPortId

}

// DUT SCI Port ID.
// DutSciPortId returns a int32
func (obj *macsecSecYRxScStaticKey) HasDutSciPortId() bool {
	return obj.obj.DutSciPortId != nil
}

// DUT SCI Port ID.
// SetDutSciPortId sets the int32 value in the MacsecSecYRxScStaticKey object
func (obj *macsecSecYRxScStaticKey) SetDutSciPortId(value int32) MacsecSecYRxScStaticKey {

	obj.obj.DutSciPortId = &value
	return obj
}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *macsecSecYRxScStaticKey) ReplayProtection() bool {

	return *obj.obj.ReplayProtection

}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *macsecSecYRxScStaticKey) HasReplayProtection() bool {
	return obj.obj.ReplayProtection != nil
}

// Enable replay protection on not.
// SetReplayProtection sets the bool value in the MacsecSecYRxScStaticKey object
func (obj *macsecSecYRxScStaticKey) SetReplayProtection(value bool) MacsecSecYRxScStaticKey {

	obj.obj.ReplayProtection = &value
	return obj
}

// Replay window size.
// ReplayWindow returns a int32
func (obj *macsecSecYRxScStaticKey) ReplayWindow() int32 {

	return *obj.obj.ReplayWindow

}

// Replay window size.
// ReplayWindow returns a int32
func (obj *macsecSecYRxScStaticKey) HasReplayWindow() bool {
	return obj.obj.ReplayWindow != nil
}

// Replay window size.
// SetReplayWindow sets the int32 value in the MacsecSecYRxScStaticKey object
func (obj *macsecSecYRxScStaticKey) SetReplayWindow(value int32) MacsecSecYRxScStaticKey {

	obj.obj.ReplayWindow = &value
	return obj
}

// Rx SAK pool.
// SakPool returns a []MacsecSecYBasicKeyGenerationStaticSak
func (obj *macsecSecYRxScStaticKey) SakPool() MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	if len(obj.obj.SakPool) == 0 {
		obj.obj.SakPool = []*otg.MacsecSecYBasicKeyGenerationStaticSak{}
	}
	if obj.sakPoolHolder == nil {
		obj.sakPoolHolder = newMacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter(&obj.obj.SakPool).setMsg(obj)
	}
	return obj.sakPoolHolder
}

type macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter struct {
	obj                                        *macsecSecYRxScStaticKey
	macsecSecYBasicKeyGenerationStaticSakSlice []MacsecSecYBasicKeyGenerationStaticSak
	fieldPtr                                   *[]*otg.MacsecSecYBasicKeyGenerationStaticSak
}

func newMacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter(ptr *[]*otg.MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	return &macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter{fieldPtr: ptr}
}

type MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter interface {
	setMsg(*macsecSecYRxScStaticKey) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	Items() []MacsecSecYBasicKeyGenerationStaticSak
	Add() MacsecSecYBasicKeyGenerationStaticSak
	Append(items ...MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	Set(index int, newObj MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	Clear() MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	clearHolderSlice() MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	appendHolderSlice(item MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
}

func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) setMsg(msg *macsecSecYRxScStaticKey) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecSecYBasicKeyGenerationStaticSak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Items() []MacsecSecYBasicKeyGenerationStaticSak {
	return obj.macsecSecYBasicKeyGenerationStaticSakSlice
}

func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Add() MacsecSecYBasicKeyGenerationStaticSak {
	newObj := &otg.MacsecSecYBasicKeyGenerationStaticSak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecSecYBasicKeyGenerationStaticSak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecSecYBasicKeyGenerationStaticSakSlice = append(obj.macsecSecYBasicKeyGenerationStaticSakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Append(items ...MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecSecYBasicKeyGenerationStaticSakSlice = append(obj.macsecSecYBasicKeyGenerationStaticSakSlice, item)
	}
	return obj
}

func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Set(index int, newObj MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecSecYBasicKeyGenerationStaticSakSlice[index] = newObj
	return obj
}
func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Clear() MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecSecYBasicKeyGenerationStaticSak{}
		obj.macsecSecYBasicKeyGenerationStaticSakSlice = []MacsecSecYBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) clearHolderSlice() MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	if len(obj.macsecSecYBasicKeyGenerationStaticSakSlice) > 0 {
		obj.macsecSecYBasicKeyGenerationStaticSakSlice = []MacsecSecYBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) appendHolderSlice(item MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYRxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	obj.macsecSecYBasicKeyGenerationStaticSakSlice = append(obj.macsecSecYBasicKeyGenerationStaticSakSlice, item)
	return obj
}

func (obj *macsecSecYRxScStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DutSystemId != nil {

		err := obj.validateMac(obj.DutSystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecSecYRxScStaticKey.DutSystemId"))
		}

	}

	if len(obj.obj.SakPool) != 0 {

		if set_default {
			obj.SakPool().clearHolderSlice()
			for _, item := range obj.obj.SakPool {
				obj.SakPool().appendHolderSlice(&macsecSecYBasicKeyGenerationStaticSak{obj: item})
			}
		}
		for _, item := range obj.SakPool().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *macsecSecYRxScStaticKey) setDefault() {
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
