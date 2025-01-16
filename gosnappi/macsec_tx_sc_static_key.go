package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecTxScStaticKey *****
type macsecTxScStaticKey struct {
	validation
	obj           *otg.MacsecTxScStaticKey
	marshaller    marshalMacsecTxScStaticKey
	unMarshaller  unMarshalMacsecTxScStaticKey
	sakPoolHolder MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
}

func NewMacsecTxScStaticKey() MacsecTxScStaticKey {
	obj := macsecTxScStaticKey{obj: &otg.MacsecTxScStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecTxScStaticKey) msg() *otg.MacsecTxScStaticKey {
	return obj.obj
}

func (obj *macsecTxScStaticKey) setMsg(msg *otg.MacsecTxScStaticKey) MacsecTxScStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecTxScStaticKey struct {
	obj *macsecTxScStaticKey
}

type marshalMacsecTxScStaticKey interface {
	// ToProto marshals MacsecTxScStaticKey to protobuf object *otg.MacsecTxScStaticKey
	ToProto() (*otg.MacsecTxScStaticKey, error)
	// ToPbText marshals MacsecTxScStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecTxScStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecTxScStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecTxScStaticKey struct {
	obj *macsecTxScStaticKey
}

type unMarshalMacsecTxScStaticKey interface {
	// FromProto unmarshals MacsecTxScStaticKey from protobuf object *otg.MacsecTxScStaticKey
	FromProto(msg *otg.MacsecTxScStaticKey) (MacsecTxScStaticKey, error)
	// FromPbText unmarshals MacsecTxScStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecTxScStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecTxScStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecTxScStaticKey) Marshal() marshalMacsecTxScStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecTxScStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecTxScStaticKey) Unmarshal() unMarshalMacsecTxScStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecTxScStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecTxScStaticKey) ToProto() (*otg.MacsecTxScStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecTxScStaticKey) FromProto(msg *otg.MacsecTxScStaticKey) (MacsecTxScStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecTxScStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecTxScStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecTxScStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecTxScStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecTxScStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecTxScStaticKey) FromJson(value string) error {
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

func (obj *macsecTxScStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecTxScStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecTxScStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecTxScStaticKey) Clone() (MacsecTxScStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecTxScStaticKey()
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

func (obj *macsecTxScStaticKey) setNil() {
	obj.sakPoolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecTxScStaticKey is tx SC setting for static key.
type MacsecTxScStaticKey interface {
	Validation
	// msg marshals MacsecTxScStaticKey to protobuf object *otg.MacsecTxScStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecTxScStaticKey
	// setMsg unmarshals MacsecTxScStaticKey from protobuf object *otg.MacsecTxScStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecTxScStaticKey) MacsecTxScStaticKey
	// provides marshal interface
	Marshal() marshalMacsecTxScStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecTxScStaticKey
	// validate validates MacsecTxScStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecTxScStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in MacsecTxScStaticKey.
	SystemId() string
	// SetSystemId assigns string provided by user to MacsecTxScStaticKey
	SetSystemId(value string) MacsecTxScStaticKey
	// HasSystemId checks if SystemId has been set in MacsecTxScStaticKey
	HasSystemId() bool
	// PortId returns int32, set in MacsecTxScStaticKey.
	PortId() int32
	// SetPortId assigns int32 provided by user to MacsecTxScStaticKey
	SetPortId(value int32) MacsecTxScStaticKey
	// HasPortId checks if PortId has been set in MacsecTxScStaticKey
	HasPortId() bool
	// EndStation returns bool, set in MacsecTxScStaticKey.
	EndStation() bool
	// SetEndStation assigns bool provided by user to MacsecTxScStaticKey
	SetEndStation(value bool) MacsecTxScStaticKey
	// HasEndStation checks if EndStation has been set in MacsecTxScStaticKey
	HasEndStation() bool
	// IncludeSci returns bool, set in MacsecTxScStaticKey.
	IncludeSci() bool
	// SetIncludeSci assigns bool provided by user to MacsecTxScStaticKey
	SetIncludeSci(value bool) MacsecTxScStaticKey
	// HasIncludeSci checks if IncludeSci has been set in MacsecTxScStaticKey
	HasIncludeSci() bool
	// Confidentiality returns bool, set in MacsecTxScStaticKey.
	Confidentiality() bool
	// SetConfidentiality assigns bool provided by user to MacsecTxScStaticKey
	SetConfidentiality(value bool) MacsecTxScStaticKey
	// HasConfidentiality checks if Confidentiality has been set in MacsecTxScStaticKey
	HasConfidentiality() bool
	// SakPool returns MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIterIter, set in MacsecTxScStaticKey
	SakPool() MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	setNil()
}

// System ID.
// SystemId returns a string
func (obj *macsecTxScStaticKey) SystemId() string {

	return *obj.obj.SystemId

}

// System ID.
// SystemId returns a string
func (obj *macsecTxScStaticKey) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// System ID.
// SetSystemId sets the string value in the MacsecTxScStaticKey object
func (obj *macsecTxScStaticKey) SetSystemId(value string) MacsecTxScStaticKey {

	obj.obj.SystemId = &value
	return obj
}

// Port ID.
// PortId returns a int32
func (obj *macsecTxScStaticKey) PortId() int32 {

	return *obj.obj.PortId

}

// Port ID.
// PortId returns a int32
func (obj *macsecTxScStaticKey) HasPortId() bool {
	return obj.obj.PortId != nil
}

// Port ID.
// SetPortId sets the int32 value in the MacsecTxScStaticKey object
func (obj *macsecTxScStaticKey) SetPortId(value int32) MacsecTxScStaticKey {

	obj.obj.PortId = &value
	return obj
}

// End station on not.
// EndStation returns a bool
func (obj *macsecTxScStaticKey) EndStation() bool {

	return *obj.obj.EndStation

}

// End station on not.
// EndStation returns a bool
func (obj *macsecTxScStaticKey) HasEndStation() bool {
	return obj.obj.EndStation != nil
}

// End station on not.
// SetEndStation sets the bool value in the MacsecTxScStaticKey object
func (obj *macsecTxScStaticKey) SetEndStation(value bool) MacsecTxScStaticKey {

	obj.obj.EndStation = &value
	return obj
}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecTxScStaticKey) IncludeSci() bool {

	return *obj.obj.IncludeSci

}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecTxScStaticKey) HasIncludeSci() bool {
	return obj.obj.IncludeSci != nil
}

// Include SCI on not.
// SetIncludeSci sets the bool value in the MacsecTxScStaticKey object
func (obj *macsecTxScStaticKey) SetIncludeSci(value bool) MacsecTxScStaticKey {

	obj.obj.IncludeSci = &value
	return obj
}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *macsecTxScStaticKey) Confidentiality() bool {

	return *obj.obj.Confidentiality

}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *macsecTxScStaticKey) HasConfidentiality() bool {
	return obj.obj.Confidentiality != nil
}

// Encrypt or not.
// SetConfidentiality sets the bool value in the MacsecTxScStaticKey object
func (obj *macsecTxScStaticKey) SetConfidentiality(value bool) MacsecTxScStaticKey {

	obj.obj.Confidentiality = &value
	return obj
}

// Tx SAK pool.
// SakPool returns a []MacsecBasicKeyGenerationStaticSak
func (obj *macsecTxScStaticKey) SakPool() MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	if len(obj.obj.SakPool) == 0 {
		obj.obj.SakPool = []*otg.MacsecBasicKeyGenerationStaticSak{}
	}
	if obj.sakPoolHolder == nil {
		obj.sakPoolHolder = newMacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter(&obj.obj.SakPool).setMsg(obj)
	}
	return obj.sakPoolHolder
}

type macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter struct {
	obj                                    *macsecTxScStaticKey
	macsecBasicKeyGenerationStaticSakSlice []MacsecBasicKeyGenerationStaticSak
	fieldPtr                               *[]*otg.MacsecBasicKeyGenerationStaticSak
}

func newMacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter(ptr *[]*otg.MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	return &macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter{fieldPtr: ptr}
}

type MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter interface {
	setMsg(*macsecTxScStaticKey) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	Items() []MacsecBasicKeyGenerationStaticSak
	Add() MacsecBasicKeyGenerationStaticSak
	Append(items ...MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	Set(index int, newObj MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	Clear() MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	clearHolderSlice() MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
	appendHolderSlice(item MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter
}

func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) setMsg(msg *macsecTxScStaticKey) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecBasicKeyGenerationStaticSak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Items() []MacsecBasicKeyGenerationStaticSak {
	return obj.macsecBasicKeyGenerationStaticSakSlice
}

func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Add() MacsecBasicKeyGenerationStaticSak {
	newObj := &otg.MacsecBasicKeyGenerationStaticSak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecBasicKeyGenerationStaticSak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Append(items ...MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, item)
	}
	return obj
}

func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Set(index int, newObj MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecBasicKeyGenerationStaticSakSlice[index] = newObj
	return obj
}
func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) Clear() MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecBasicKeyGenerationStaticSak{}
		obj.macsecBasicKeyGenerationStaticSakSlice = []MacsecBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) clearHolderSlice() MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	if len(obj.macsecBasicKeyGenerationStaticSakSlice) > 0 {
		obj.macsecBasicKeyGenerationStaticSakSlice = []MacsecBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter) appendHolderSlice(item MacsecBasicKeyGenerationStaticSak) MacsecTxScStaticKeyMacsecBasicKeyGenerationStaticSakIter {
	obj.macsecBasicKeyGenerationStaticSakSlice = append(obj.macsecBasicKeyGenerationStaticSakSlice, item)
	return obj
}

func (obj *macsecTxScStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecTxScStaticKey.SystemId"))
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

func (obj *macsecTxScStaticKey) setDefault() {
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
