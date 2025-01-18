package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYTxScStaticKey *****
type macsecSecYTxScStaticKey struct {
	validation
	obj           *otg.MacsecSecYTxScStaticKey
	marshaller    marshalMacsecSecYTxScStaticKey
	unMarshaller  unMarshalMacsecSecYTxScStaticKey
	sakPoolHolder MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
}

func NewMacsecSecYTxScStaticKey() MacsecSecYTxScStaticKey {
	obj := macsecSecYTxScStaticKey{obj: &otg.MacsecSecYTxScStaticKey{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYTxScStaticKey) msg() *otg.MacsecSecYTxScStaticKey {
	return obj.obj
}

func (obj *macsecSecYTxScStaticKey) setMsg(msg *otg.MacsecSecYTxScStaticKey) MacsecSecYTxScStaticKey {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYTxScStaticKey struct {
	obj *macsecSecYTxScStaticKey
}

type marshalMacsecSecYTxScStaticKey interface {
	// ToProto marshals MacsecSecYTxScStaticKey to protobuf object *otg.MacsecSecYTxScStaticKey
	ToProto() (*otg.MacsecSecYTxScStaticKey, error)
	// ToPbText marshals MacsecSecYTxScStaticKey to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYTxScStaticKey to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYTxScStaticKey to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYTxScStaticKey struct {
	obj *macsecSecYTxScStaticKey
}

type unMarshalMacsecSecYTxScStaticKey interface {
	// FromProto unmarshals MacsecSecYTxScStaticKey from protobuf object *otg.MacsecSecYTxScStaticKey
	FromProto(msg *otg.MacsecSecYTxScStaticKey) (MacsecSecYTxScStaticKey, error)
	// FromPbText unmarshals MacsecSecYTxScStaticKey from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYTxScStaticKey from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYTxScStaticKey from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYTxScStaticKey) Marshal() marshalMacsecSecYTxScStaticKey {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYTxScStaticKey{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYTxScStaticKey) Unmarshal() unMarshalMacsecSecYTxScStaticKey {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYTxScStaticKey{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYTxScStaticKey) ToProto() (*otg.MacsecSecYTxScStaticKey, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYTxScStaticKey) FromProto(msg *otg.MacsecSecYTxScStaticKey) (MacsecSecYTxScStaticKey, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYTxScStaticKey) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYTxScStaticKey) FromPbText(value string) error {
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

func (m *marshalmacsecSecYTxScStaticKey) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYTxScStaticKey) FromYaml(value string) error {
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

func (m *marshalmacsecSecYTxScStaticKey) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYTxScStaticKey) FromJson(value string) error {
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

func (obj *macsecSecYTxScStaticKey) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYTxScStaticKey) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYTxScStaticKey) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYTxScStaticKey) Clone() (MacsecSecYTxScStaticKey, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYTxScStaticKey()
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

func (obj *macsecSecYTxScStaticKey) setNil() {
	obj.sakPoolHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYTxScStaticKey is tx SC setting for static key.
type MacsecSecYTxScStaticKey interface {
	Validation
	// msg marshals MacsecSecYTxScStaticKey to protobuf object *otg.MacsecSecYTxScStaticKey
	// and doesn't set defaults
	msg() *otg.MacsecSecYTxScStaticKey
	// setMsg unmarshals MacsecSecYTxScStaticKey from protobuf object *otg.MacsecSecYTxScStaticKey
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYTxScStaticKey) MacsecSecYTxScStaticKey
	// provides marshal interface
	Marshal() marshalMacsecSecYTxScStaticKey
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYTxScStaticKey
	// validate validates MacsecSecYTxScStaticKey
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYTxScStaticKey, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SystemId returns string, set in MacsecSecYTxScStaticKey.
	SystemId() string
	// SetSystemId assigns string provided by user to MacsecSecYTxScStaticKey
	SetSystemId(value string) MacsecSecYTxScStaticKey
	// HasSystemId checks if SystemId has been set in MacsecSecYTxScStaticKey
	HasSystemId() bool
	// PortId returns int32, set in MacsecSecYTxScStaticKey.
	PortId() int32
	// SetPortId assigns int32 provided by user to MacsecSecYTxScStaticKey
	SetPortId(value int32) MacsecSecYTxScStaticKey
	// HasPortId checks if PortId has been set in MacsecSecYTxScStaticKey
	HasPortId() bool
	// EndStation returns bool, set in MacsecSecYTxScStaticKey.
	EndStation() bool
	// SetEndStation assigns bool provided by user to MacsecSecYTxScStaticKey
	SetEndStation(value bool) MacsecSecYTxScStaticKey
	// HasEndStation checks if EndStation has been set in MacsecSecYTxScStaticKey
	HasEndStation() bool
	// IncludeSci returns bool, set in MacsecSecYTxScStaticKey.
	IncludeSci() bool
	// SetIncludeSci assigns bool provided by user to MacsecSecYTxScStaticKey
	SetIncludeSci(value bool) MacsecSecYTxScStaticKey
	// HasIncludeSci checks if IncludeSci has been set in MacsecSecYTxScStaticKey
	HasIncludeSci() bool
	// Confidentiality returns bool, set in MacsecSecYTxScStaticKey.
	Confidentiality() bool
	// SetConfidentiality assigns bool provided by user to MacsecSecYTxScStaticKey
	SetConfidentiality(value bool) MacsecSecYTxScStaticKey
	// HasConfidentiality checks if Confidentiality has been set in MacsecSecYTxScStaticKey
	HasConfidentiality() bool
	// SakPool returns MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIterIter, set in MacsecSecYTxScStaticKey
	SakPool() MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	setNil()
}

// System ID.
// SystemId returns a string
func (obj *macsecSecYTxScStaticKey) SystemId() string {

	return *obj.obj.SystemId

}

// System ID.
// SystemId returns a string
func (obj *macsecSecYTxScStaticKey) HasSystemId() bool {
	return obj.obj.SystemId != nil
}

// System ID.
// SetSystemId sets the string value in the MacsecSecYTxScStaticKey object
func (obj *macsecSecYTxScStaticKey) SetSystemId(value string) MacsecSecYTxScStaticKey {

	obj.obj.SystemId = &value
	return obj
}

// Port ID.
// PortId returns a int32
func (obj *macsecSecYTxScStaticKey) PortId() int32 {

	return *obj.obj.PortId

}

// Port ID.
// PortId returns a int32
func (obj *macsecSecYTxScStaticKey) HasPortId() bool {
	return obj.obj.PortId != nil
}

// Port ID.
// SetPortId sets the int32 value in the MacsecSecYTxScStaticKey object
func (obj *macsecSecYTxScStaticKey) SetPortId(value int32) MacsecSecYTxScStaticKey {

	obj.obj.PortId = &value
	return obj
}

// End station on not.
// EndStation returns a bool
func (obj *macsecSecYTxScStaticKey) EndStation() bool {

	return *obj.obj.EndStation

}

// End station on not.
// EndStation returns a bool
func (obj *macsecSecYTxScStaticKey) HasEndStation() bool {
	return obj.obj.EndStation != nil
}

// End station on not.
// SetEndStation sets the bool value in the MacsecSecYTxScStaticKey object
func (obj *macsecSecYTxScStaticKey) SetEndStation(value bool) MacsecSecYTxScStaticKey {

	obj.obj.EndStation = &value
	return obj
}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecSecYTxScStaticKey) IncludeSci() bool {

	return *obj.obj.IncludeSci

}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *macsecSecYTxScStaticKey) HasIncludeSci() bool {
	return obj.obj.IncludeSci != nil
}

// Include SCI on not.
// SetIncludeSci sets the bool value in the MacsecSecYTxScStaticKey object
func (obj *macsecSecYTxScStaticKey) SetIncludeSci(value bool) MacsecSecYTxScStaticKey {

	obj.obj.IncludeSci = &value
	return obj
}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *macsecSecYTxScStaticKey) Confidentiality() bool {

	return *obj.obj.Confidentiality

}

// Encrypt or not.
// Confidentiality returns a bool
func (obj *macsecSecYTxScStaticKey) HasConfidentiality() bool {
	return obj.obj.Confidentiality != nil
}

// Encrypt or not.
// SetConfidentiality sets the bool value in the MacsecSecYTxScStaticKey object
func (obj *macsecSecYTxScStaticKey) SetConfidentiality(value bool) MacsecSecYTxScStaticKey {

	obj.obj.Confidentiality = &value
	return obj
}

// Tx SAK pool.
// SakPool returns a []MacsecSecYBasicKeyGenerationStaticSak
func (obj *macsecSecYTxScStaticKey) SakPool() MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	if len(obj.obj.SakPool) == 0 {
		obj.obj.SakPool = []*otg.MacsecSecYBasicKeyGenerationStaticSak{}
	}
	if obj.sakPoolHolder == nil {
		obj.sakPoolHolder = newMacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter(&obj.obj.SakPool).setMsg(obj)
	}
	return obj.sakPoolHolder
}

type macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter struct {
	obj                                        *macsecSecYTxScStaticKey
	macsecSecYBasicKeyGenerationStaticSakSlice []MacsecSecYBasicKeyGenerationStaticSak
	fieldPtr                                   *[]*otg.MacsecSecYBasicKeyGenerationStaticSak
}

func newMacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter(ptr *[]*otg.MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	return &macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter{fieldPtr: ptr}
}

type MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter interface {
	setMsg(*macsecSecYTxScStaticKey) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	Items() []MacsecSecYBasicKeyGenerationStaticSak
	Add() MacsecSecYBasicKeyGenerationStaticSak
	Append(items ...MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	Set(index int, newObj MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	Clear() MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	clearHolderSlice() MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
	appendHolderSlice(item MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter
}

func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) setMsg(msg *macsecSecYTxScStaticKey) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecSecYBasicKeyGenerationStaticSak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Items() []MacsecSecYBasicKeyGenerationStaticSak {
	return obj.macsecSecYBasicKeyGenerationStaticSakSlice
}

func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Add() MacsecSecYBasicKeyGenerationStaticSak {
	newObj := &otg.MacsecSecYBasicKeyGenerationStaticSak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecSecYBasicKeyGenerationStaticSak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecSecYBasicKeyGenerationStaticSakSlice = append(obj.macsecSecYBasicKeyGenerationStaticSakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Append(items ...MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecSecYBasicKeyGenerationStaticSakSlice = append(obj.macsecSecYBasicKeyGenerationStaticSakSlice, item)
	}
	return obj
}

func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Set(index int, newObj MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecSecYBasicKeyGenerationStaticSakSlice[index] = newObj
	return obj
}
func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) Clear() MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecSecYBasicKeyGenerationStaticSak{}
		obj.macsecSecYBasicKeyGenerationStaticSakSlice = []MacsecSecYBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) clearHolderSlice() MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	if len(obj.macsecSecYBasicKeyGenerationStaticSakSlice) > 0 {
		obj.macsecSecYBasicKeyGenerationStaticSakSlice = []MacsecSecYBasicKeyGenerationStaticSak{}
	}
	return obj
}
func (obj *macsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter) appendHolderSlice(item MacsecSecYBasicKeyGenerationStaticSak) MacsecSecYTxScStaticKeyMacsecSecYBasicKeyGenerationStaticSakIter {
	obj.macsecSecYBasicKeyGenerationStaticSakSlice = append(obj.macsecSecYBasicKeyGenerationStaticSakSlice, item)
	return obj
}

func (obj *macsecSecYTxScStaticKey) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SystemId != nil {

		err := obj.validateMac(obj.SystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecSecYTxScStaticKey.SystemId"))
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

func (obj *macsecSecYTxScStaticKey) setDefault() {
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
