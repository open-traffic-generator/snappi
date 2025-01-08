package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecTxStaticKey *****
type macsecTxStaticKey struct {
	validation
	obj          *otg.MacsecTxStaticKey
	marshaller   marshalMacsecTxStaticKey
	unMarshaller unMarshalMacsecTxStaticKey
}

func NewMacsecTxStaticKey() MacsecTxStaticKey {
	obj := macsecTxStaticKey{obj: &otg.MacsecTxStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecTxStaticKey) msg() *otg.MacsecTxStaticKey {
	return obj.obj
}

func (obj *macsecTxStaticKey) setMsg(msg *otg.MacsecTxStaticKey) MacsecTxStaticKey {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecTxStaticKey struct {
	obj *macsecTxStaticKey
}

type marshalMacsecTxStaticKey interface {
	// ToProto marshals MacsecTxStaticKey to protobuf object *otg.MacsecTxStaticKey
	ToProto() (*otg.MacsecTxStaticKey, error)
	// ToPbText marshals MacsecTxStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecTxStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecTxStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecTxStaticKey struct {
	obj *macsecTxStaticKey
}

type unMarshalMacsecTxStaticKey interface {
	// FromProto unmarshals MacsecTxStaticKey from protobuf object *otg.MacsecTxStaticKey
	FromProto(msg *otg.MacsecTxStaticKey) (MacsecTxStaticKey, error)
	// FromPbText unmarshals MacsecTxStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecTxStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecTxStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecTxStaticKey) Marshal() marshalMacsecTxStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecTxStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecTxStaticKey) Unmarshal() unMarshalMacsecTxStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecTxStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecTxStaticKey) ToProto() (*otg.MacsecTxStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecTxStaticKey) FromProto(msg *otg.MacsecTxStaticKey) (MacsecTxStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecTxStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecTxStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecTxStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecTxStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecTxStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecTxStaticKey) FromJson(value string) error {
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

func (obj *macsecTxStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecTxStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecTxStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecTxStaticKey) Clone() (MacsecTxStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecTxStaticKey()
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

// MacsecTxStaticKey is static key Tx settings.
type MacsecTxStaticKey interface {
	Validation
	// msg marshals MacsecTxStaticKey to protobuf object *otg.MacsecTxStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecTxStaticKey
	// setMsg unmarshals MacsecTxStaticKey from protobuf object *otg.MacsecTxStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecTxStaticKey) MacsecTxStaticKey
	// provides marshal interface
	Marshal() marshalMacsecTxStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecTxStaticKey
	// validate validates MacsecTxStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecTxStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in MacsecTxStaticKey.
	SystemId() string
	// SetSystemId assigns string provided by user to MacsecTxStaticKey
	SetSystemId(value string) MacsecTxStaticKey
	// HasSystemId checks if SystemId has been set in MacsecTxStaticKey
	HasSystemId() bool
	// PortId returns int32, set in MacsecTxStaticKey.
	PortId() int32
	// SetPortId assigns int32 provided by user to MacsecTxStaticKey
	SetPortId(value int32) MacsecTxStaticKey
	// HasPortId checks if PortId has been set in MacsecTxStaticKey
	HasPortId() bool
	// EndStation returns bool, set in MacsecTxStaticKey.
	EndStation() bool
	// SetEndStation assigns bool provided by user to MacsecTxStaticKey
	SetEndStation(value bool) MacsecTxStaticKey
	// HasEndStation checks if EndStation has been set in MacsecTxStaticKey
	HasEndStation() bool
	// IncludeSci returns bool, set in MacsecTxStaticKey.
	IncludeSci() bool
	// SetIncludeSci assigns bool provided by user to MacsecTxStaticKey
	SetIncludeSci(value bool) MacsecTxStaticKey
	// HasIncludeSci checks if IncludeSci has been set in MacsecTxStaticKey
	HasIncludeSci() bool
	// Confidentiality returns bool, set in MacsecTxStaticKey.
	Confidentiality() bool
	// SetConfidentiality assigns bool provided by user to MacsecTxStaticKey
	SetConfidentiality(value bool) MacsecTxStaticKey
	// HasConfidentiality checks if Confidentiality has been set in MacsecTxStaticKey
	HasConfidentiality() bool
}

// System ID.
// SystemId returns a string
func (obj *macsecTxStaticKey) SystemId() string {

	return *obj.obj.SystemId

}

// System ID.
// SystemId returns a string
func (obj *macsecTxStaticKey) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// System ID.
// SetSystemId sets the string value in the MacsecTxStaticKey object
func (obj *macsecTxStaticKey) SetSystemId(value string) MacsecTxStaticKey {

	obj.obj.SystemId = &value
	return obj
}

// Port ID.
// PortId returns a int32
func (obj *macsecTxStaticKey) PortId() int32 {

	return *obj.obj.PortId

}

// Port ID.
// PortId returns a int32
func (obj *macsecTxStaticKey) HasPortId() bool {
	return obj.obj.PortId != nil
}

// Port ID.
// SetPortId sets the int32 value in the MacsecTxStaticKey object
func (obj *macsecTxStaticKey) SetPortId(value int32) MacsecTxStaticKey {

	obj.obj.PortId = &value
	return obj
}

// End station on not.
// EndStation returns a bool
func (obj *macsecTxStaticKey) EndStation() bool {

	return *obj.obj.EndStation

}

// End station on not.
// EndStation returns a bool
func (obj *macsecTxStaticKey) HasEndStation() bool {
	return obj.obj.EndStation != nil
}

// End station on not.
// SetEndStation sets the bool value in the MacsecTxStaticKey object
func (obj *macsecTxStaticKey) SetEndStation(value bool) MacsecTxStaticKey {

	obj.obj.EndStation = &value
	return obj
}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecTxStaticKey) IncludeSci() bool {

	return *obj.obj.IncludeSci

}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecTxStaticKey) HasIncludeSci() bool {
	return obj.obj.IncludeSci != nil
}

// Include SCI on not.
// SetIncludeSci sets the bool value in the MacsecTxStaticKey object
func (obj *macsecTxStaticKey) SetIncludeSci(value bool) MacsecTxStaticKey {

	obj.obj.IncludeSci = &value
	return obj
}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *macsecTxStaticKey) Confidentiality() bool {

	return *obj.obj.Confidentiality

}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *macsecTxStaticKey) HasConfidentiality() bool {
	return obj.obj.Confidentiality != nil
}

// Encrypt or not.
// SetConfidentiality sets the bool value in the MacsecTxStaticKey object
func (obj *macsecTxStaticKey) SetConfidentiality(value bool) MacsecTxStaticKey {

	obj.obj.Confidentiality = &value
	return obj
}

func (obj *macsecTxStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecTxStaticKey.SystemId"))
		}

	}

}

func (obj *macsecTxStaticKey) setDefault() {
	if obj.obj.PortId == nil {
		obj.SetPortId(1)
	}
	if obj.obj.EndStation == nil {
		obj.SetEndStation(false)
	}
	if obj.obj.IncludeSci == nil {
		obj.SetIncludeSci(false)
	}
	if obj.obj.Confidentiality == nil {
		obj.SetConfidentiality(true)
	}

}
