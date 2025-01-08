package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecRxStaticKey *****
type macsecRxStaticKey struct {
	validation
	obj          *otg.MacsecRxStaticKey
	marshaller   marshalMacsecRxStaticKey
	unMarshaller unMarshalMacsecRxStaticKey
}

func NewMacsecRxStaticKey() MacsecRxStaticKey {
	obj := macsecRxStaticKey{obj: &otg.MacsecRxStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecRxStaticKey) msg() *otg.MacsecRxStaticKey {
	return obj.obj
}

func (obj *macsecRxStaticKey) setMsg(msg *otg.MacsecRxStaticKey) MacsecRxStaticKey {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecRxStaticKey struct {
	obj *macsecRxStaticKey
}

type marshalMacsecRxStaticKey interface {
	// ToProto marshals MacsecRxStaticKey to protobuf object *otg.MacsecRxStaticKey
	ToProto() (*otg.MacsecRxStaticKey, error)
	// ToPbText marshals MacsecRxStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecRxStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecRxStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecRxStaticKey struct {
	obj *macsecRxStaticKey
}

type unMarshalMacsecRxStaticKey interface {
	// FromProto unmarshals MacsecRxStaticKey from protobuf object *otg.MacsecRxStaticKey
	FromProto(msg *otg.MacsecRxStaticKey) (MacsecRxStaticKey, error)
	// FromPbText unmarshals MacsecRxStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecRxStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecRxStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecRxStaticKey) Marshal() marshalMacsecRxStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecRxStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecRxStaticKey) Unmarshal() unMarshalMacsecRxStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecRxStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecRxStaticKey) ToProto() (*otg.MacsecRxStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecRxStaticKey) FromProto(msg *otg.MacsecRxStaticKey) (MacsecRxStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecRxStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecRxStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecRxStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecRxStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecRxStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecRxStaticKey) FromJson(value string) error {
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

func (obj *macsecRxStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecRxStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecRxStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecRxStaticKey) Clone() (MacsecRxStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecRxStaticKey()
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

// MacsecRxStaticKey is static key Rx settings.
type MacsecRxStaticKey interface {
	Validation
	// msg marshals MacsecRxStaticKey to protobuf object *otg.MacsecRxStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecRxStaticKey
	// setMsg unmarshals MacsecRxStaticKey from protobuf object *otg.MacsecRxStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecRxStaticKey) MacsecRxStaticKey
	// provides marshal interface
	Marshal() marshalMacsecRxStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecRxStaticKey
	// validate validates MacsecRxStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecRxStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DutSystemId returns string, set in MacsecRxStaticKey.
	DutSystemId() string
	// SetDutSystemId assigns string provided by user to MacsecRxStaticKey
	SetDutSystemId(value string) MacsecRxStaticKey
	// HasDutSystemId checks if DutSystemId has been set in MacsecRxStaticKey
	HasDutSystemId() bool
	// DutSciPortId returns int32, set in MacsecRxStaticKey.
	DutSciPortId() int32
	// SetDutSciPortId assigns int32 provided by user to MacsecRxStaticKey
	SetDutSciPortId(value int32) MacsecRxStaticKey
	// HasDutSciPortId checks if DutSciPortId has been set in MacsecRxStaticKey
	HasDutSciPortId() bool
	// ReplayProtection returns bool, set in MacsecRxStaticKey.
	ReplayProtection() bool
	// SetReplayProtection assigns bool provided by user to MacsecRxStaticKey
	SetReplayProtection(value bool) MacsecRxStaticKey
	// HasReplayProtection checks if ReplayProtection has been set in MacsecRxStaticKey
	HasReplayProtection() bool
	// ReplayWindow returns int32, set in MacsecRxStaticKey.
	ReplayWindow() int32
	// SetReplayWindow assigns int32 provided by user to MacsecRxStaticKey
	SetReplayWindow(value int32) MacsecRxStaticKey
	// HasReplayWindow checks if ReplayWindow has been set in MacsecRxStaticKey
	HasReplayWindow() bool
}

// DUT system ID.
// DutSystemId returns a string
func (obj *macsecRxStaticKey) DutSystemId() string {

	return *obj.obj.DutSystemId

}

// DUT system ID.
// DutSystemId returns a string
func (obj *macsecRxStaticKey) HasDutSystemId() bool {
	return obj.obj.DutSystemId != nil
}

// DUT system ID.
// SetDutSystemId sets the string value in the MacsecRxStaticKey object
func (obj *macsecRxStaticKey) SetDutSystemId(value string) MacsecRxStaticKey {

	obj.obj.DutSystemId = &value
	return obj
}

// DUT SCI Port ID.
// DutSciPortId returns a int32
func (obj *macsecRxStaticKey) DutSciPortId() int32 {

	return *obj.obj.DutSciPortId

}

// DUT SCI Port ID.
// DutSciPortId returns a int32
func (obj *macsecRxStaticKey) HasDutSciPortId() bool {
	return obj.obj.DutSciPortId != nil
}

// DUT SCI Port ID.
// SetDutSciPortId sets the int32 value in the MacsecRxStaticKey object
func (obj *macsecRxStaticKey) SetDutSciPortId(value int32) MacsecRxStaticKey {

	obj.obj.DutSciPortId = &value
	return obj
}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *macsecRxStaticKey) ReplayProtection() bool {

	return *obj.obj.ReplayProtection

}

// Enable replay protection on not.
// ReplayProtection returns a bool
func (obj *macsecRxStaticKey) HasReplayProtection() bool {
	return obj.obj.ReplayProtection != nil
}

// Enable replay protection on not.
// SetReplayProtection sets the bool value in the MacsecRxStaticKey object
func (obj *macsecRxStaticKey) SetReplayProtection(value bool) MacsecRxStaticKey {

	obj.obj.ReplayProtection = &value
	return obj
}

// Replay window size.
// ReplayWindow returns a int32
func (obj *macsecRxStaticKey) ReplayWindow() int32 {

	return *obj.obj.ReplayWindow

}

// Replay window size.
// ReplayWindow returns a int32
func (obj *macsecRxStaticKey) HasReplayWindow() bool {
	return obj.obj.ReplayWindow != nil
}

// Replay window size.
// SetReplayWindow sets the int32 value in the MacsecRxStaticKey object
func (obj *macsecRxStaticKey) SetReplayWindow(value int32) MacsecRxStaticKey {

	obj.obj.ReplayWindow = &value
	return obj
}

func (obj *macsecRxStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DutSystemId != nil {

		err := obj.validateMac(obj.DutSystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecRxStaticKey.DutSystemId"))
		}

	}

}

func (obj *macsecRxStaticKey) setDefault() {
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
