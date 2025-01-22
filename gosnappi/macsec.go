package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Macsec *****
type macsec struct {
	validation
	obj                *otg.Macsec
	marshaller         marshalMacsec
	unMarshaller       unMarshalMacsec
	basicHolder        MacsecBasic
	txscsHolder        MacsecMacsecTxScIter
	rxscsHolder        MacsecMacsecRxScIter
	cryptoEngineHolder MacsecCryptoEngine
	advanceHolder      MacsecAdvance
}

func NewMacsec() Macsec {
	obj := macsec{obj: &otg.Macsec{}}
	obj.setDefault()
	return &obj
}

func (obj *macsec) msg() *otg.Macsec {
	return obj.obj
}

func (obj *macsec) setMsg(msg *otg.Macsec) Macsec {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsec struct {
	obj *macsec
}

type marshalMacsec interface {
	// ToProto marshals Macsec to protobuf object *otg.Macsec
	ToProto() (*otg.Macsec, error)
	// ToPbText marshals Macsec to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Macsec to YAML text
	ToYaml() (string, error)
	// ToJson marshals Macsec to JSON text
	ToJson() (string, error)
}

type unMarshalmacsec struct {
	obj *macsec
}

type unMarshalMacsec interface {
	// FromProto unmarshals Macsec from protobuf object *otg.Macsec
	FromProto(msg *otg.Macsec) (Macsec, error)
	// FromPbText unmarshals Macsec from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Macsec from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Macsec from JSON text
	FromJson(value string) error
}

func (obj *macsec) Marshal() marshalMacsec {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsec{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsec) Unmarshal() unMarshalMacsec {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsec{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsec) ToProto() (*otg.Macsec, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsec) FromProto(msg *otg.Macsec) (Macsec, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsec) ToPbText() (string, error) {
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

func (m *unMarshalmacsec) FromPbText(value string) error {
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

func (m *marshalmacsec) ToYaml() (string, error) {
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

func (m *unMarshalmacsec) FromYaml(value string) error {
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

func (m *marshalmacsec) ToJson() (string, error) {
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

func (m *unMarshalmacsec) FromJson(value string) error {
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

func (obj *macsec) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsec) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsec) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsec) Clone() (Macsec, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsec()
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

func (obj *macsec) setNil() {
	obj.basicHolder = nil
	obj.txscsHolder = nil
	obj.rxscsHolder = nil
	obj.cryptoEngineHolder = nil
	obj.advanceHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Macsec is configuration of a Secure Entity (SecY).
type Macsec interface {
	Validation
	// msg marshals Macsec to protobuf object *otg.Macsec
	// and doesn't set defaults
	msg() *otg.Macsec
	// setMsg unmarshals Macsec from protobuf object *otg.Macsec
	// and doesn't set defaults
	setMsg(*otg.Macsec) Macsec
	// provides marshal interface
	Marshal() marshalMacsec
	// provides unmarshal interface
	Unmarshal() unMarshalMacsec
	// validate validates Macsec
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Macsec, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in Macsec.
	Name() string
	// SetName assigns string provided by user to Macsec
	SetName(value string) Macsec
	// Basic returns MacsecBasic, set in Macsec.
	// MacsecBasic is a container of basic properties for a SecY.
	Basic() MacsecBasic
	// SetBasic assigns MacsecBasic provided by user to Macsec.
	// MacsecBasic is a container of basic properties for a SecY.
	SetBasic(value MacsecBasic) Macsec
	// Txscs returns MacsecMacsecTxScIterIter, set in Macsec
	Txscs() MacsecMacsecTxScIter
	// Rxscs returns MacsecMacsecRxScIterIter, set in Macsec
	Rxscs() MacsecMacsecRxScIter
	// CryptoEngine returns MacsecCryptoEngine, set in Macsec.
	// MacsecCryptoEngine is a container of crypto engine properties of a SecY.
	CryptoEngine() MacsecCryptoEngine
	// SetCryptoEngine assigns MacsecCryptoEngine provided by user to Macsec.
	// MacsecCryptoEngine is a container of crypto engine properties of a SecY.
	SetCryptoEngine(value MacsecCryptoEngine) Macsec
	// HasCryptoEngine checks if CryptoEngine has been set in Macsec
	HasCryptoEngine() bool
	// Advance returns MacsecAdvance, set in Macsec.
	// MacsecAdvance is a container of advance properties SecY.
	Advance() MacsecAdvance
	// SetAdvance assigns MacsecAdvance provided by user to Macsec.
	// MacsecAdvance is a container of advance properties SecY.
	SetAdvance(value MacsecAdvance) Macsec
	// HasAdvance checks if Advance has been set in Macsec
	HasAdvance() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *macsec) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Macsec object
func (obj *macsec) SetName(value string) Macsec {

	obj.obj.Name = &value
	return obj
}

// This contains the basic properties of SecY.
// Basic returns a MacsecBasic
func (obj *macsec) Basic() MacsecBasic {
	if obj.obj.Basic == nil {
		obj.obj.Basic = NewMacsecBasic().msg()
	}
	if obj.basicHolder == nil {
		obj.basicHolder = &macsecBasic{obj: obj.obj.Basic}
	}
	return obj.basicHolder
}

// This contains the basic properties of SecY.
// SetBasic sets the MacsecBasic value in the Macsec object
func (obj *macsec) SetBasic(value MacsecBasic) Macsec {

	obj.basicHolder = nil
	obj.obj.Basic = value.msg()

	return obj
}

// Tx secure channels.
// Txscs returns a []MacsecTxSc
func (obj *macsec) Txscs() MacsecMacsecTxScIter {
	if len(obj.obj.Txscs) == 0 {
		obj.obj.Txscs = []*otg.MacsecTxSc{}
	}
	if obj.txscsHolder == nil {
		obj.txscsHolder = newMacsecMacsecTxScIter(&obj.obj.Txscs).setMsg(obj)
	}
	return obj.txscsHolder
}

type macsecMacsecTxScIter struct {
	obj             *macsec
	macsecTxScSlice []MacsecTxSc
	fieldPtr        *[]*otg.MacsecTxSc
}

func newMacsecMacsecTxScIter(ptr *[]*otg.MacsecTxSc) MacsecMacsecTxScIter {
	return &macsecMacsecTxScIter{fieldPtr: ptr}
}

type MacsecMacsecTxScIter interface {
	setMsg(*macsec) MacsecMacsecTxScIter
	Items() []MacsecTxSc
	Add() MacsecTxSc
	Append(items ...MacsecTxSc) MacsecMacsecTxScIter
	Set(index int, newObj MacsecTxSc) MacsecMacsecTxScIter
	Clear() MacsecMacsecTxScIter
	clearHolderSlice() MacsecMacsecTxScIter
	appendHolderSlice(item MacsecTxSc) MacsecMacsecTxScIter
}

func (obj *macsecMacsecTxScIter) setMsg(msg *macsec) MacsecMacsecTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecMacsecTxScIter) Items() []MacsecTxSc {
	return obj.macsecTxScSlice
}

func (obj *macsecMacsecTxScIter) Add() MacsecTxSc {
	newObj := &otg.MacsecTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.macsecTxScSlice = append(obj.macsecTxScSlice, newLibObj)
	return newLibObj
}

func (obj *macsecMacsecTxScIter) Append(items ...MacsecTxSc) MacsecMacsecTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecTxScSlice = append(obj.macsecTxScSlice, item)
	}
	return obj
}

func (obj *macsecMacsecTxScIter) Set(index int, newObj MacsecTxSc) MacsecMacsecTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecTxScSlice[index] = newObj
	return obj
}
func (obj *macsecMacsecTxScIter) Clear() MacsecMacsecTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecTxSc{}
		obj.macsecTxScSlice = []MacsecTxSc{}
	}
	return obj
}
func (obj *macsecMacsecTxScIter) clearHolderSlice() MacsecMacsecTxScIter {
	if len(obj.macsecTxScSlice) > 0 {
		obj.macsecTxScSlice = []MacsecTxSc{}
	}
	return obj
}
func (obj *macsecMacsecTxScIter) appendHolderSlice(item MacsecTxSc) MacsecMacsecTxScIter {
	obj.macsecTxScSlice = append(obj.macsecTxScSlice, item)
	return obj
}

// Rx secure channels.
// Rxscs returns a []MacsecRxSc
func (obj *macsec) Rxscs() MacsecMacsecRxScIter {
	if len(obj.obj.Rxscs) == 0 {
		obj.obj.Rxscs = []*otg.MacsecRxSc{}
	}
	if obj.rxscsHolder == nil {
		obj.rxscsHolder = newMacsecMacsecRxScIter(&obj.obj.Rxscs).setMsg(obj)
	}
	return obj.rxscsHolder
}

type macsecMacsecRxScIter struct {
	obj             *macsec
	macsecRxScSlice []MacsecRxSc
	fieldPtr        *[]*otg.MacsecRxSc
}

func newMacsecMacsecRxScIter(ptr *[]*otg.MacsecRxSc) MacsecMacsecRxScIter {
	return &macsecMacsecRxScIter{fieldPtr: ptr}
}

type MacsecMacsecRxScIter interface {
	setMsg(*macsec) MacsecMacsecRxScIter
	Items() []MacsecRxSc
	Add() MacsecRxSc
	Append(items ...MacsecRxSc) MacsecMacsecRxScIter
	Set(index int, newObj MacsecRxSc) MacsecMacsecRxScIter
	Clear() MacsecMacsecRxScIter
	clearHolderSlice() MacsecMacsecRxScIter
	appendHolderSlice(item MacsecRxSc) MacsecMacsecRxScIter
}

func (obj *macsecMacsecRxScIter) setMsg(msg *macsec) MacsecMacsecRxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecRxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecMacsecRxScIter) Items() []MacsecRxSc {
	return obj.macsecRxScSlice
}

func (obj *macsecMacsecRxScIter) Add() MacsecRxSc {
	newObj := &otg.MacsecRxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecRxSc{obj: newObj}
	newLibObj.setDefault()
	obj.macsecRxScSlice = append(obj.macsecRxScSlice, newLibObj)
	return newLibObj
}

func (obj *macsecMacsecRxScIter) Append(items ...MacsecRxSc) MacsecMacsecRxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecRxScSlice = append(obj.macsecRxScSlice, item)
	}
	return obj
}

func (obj *macsecMacsecRxScIter) Set(index int, newObj MacsecRxSc) MacsecMacsecRxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecRxScSlice[index] = newObj
	return obj
}
func (obj *macsecMacsecRxScIter) Clear() MacsecMacsecRxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecRxSc{}
		obj.macsecRxScSlice = []MacsecRxSc{}
	}
	return obj
}
func (obj *macsecMacsecRxScIter) clearHolderSlice() MacsecMacsecRxScIter {
	if len(obj.macsecRxScSlice) > 0 {
		obj.macsecRxScSlice = []MacsecRxSc{}
	}
	return obj
}
func (obj *macsecMacsecRxScIter) appendHolderSlice(item MacsecRxSc) MacsecMacsecRxScIter {
	obj.macsecRxScSlice = append(obj.macsecRxScSlice, item)
	return obj
}

// This contains crypto engine properties of SecY.
// CryptoEngine returns a MacsecCryptoEngine
func (obj *macsec) CryptoEngine() MacsecCryptoEngine {
	if obj.obj.CryptoEngine == nil {
		obj.obj.CryptoEngine = NewMacsecCryptoEngine().msg()
	}
	if obj.cryptoEngineHolder == nil {
		obj.cryptoEngineHolder = &macsecCryptoEngine{obj: obj.obj.CryptoEngine}
	}
	return obj.cryptoEngineHolder
}

// This contains crypto engine properties of SecY.
// CryptoEngine returns a MacsecCryptoEngine
func (obj *macsec) HasCryptoEngine() bool {
	return obj.obj.CryptoEngine != nil
}

// This contains crypto engine properties of SecY.
// SetCryptoEngine sets the MacsecCryptoEngine value in the Macsec object
func (obj *macsec) SetCryptoEngine(value MacsecCryptoEngine) Macsec {

	obj.cryptoEngineHolder = nil
	obj.obj.CryptoEngine = value.msg()

	return obj
}

// This contains advance properties of SecY.
// Advance returns a MacsecAdvance
func (obj *macsec) Advance() MacsecAdvance {
	if obj.obj.Advance == nil {
		obj.obj.Advance = NewMacsecAdvance().msg()
	}
	if obj.advanceHolder == nil {
		obj.advanceHolder = &macsecAdvance{obj: obj.obj.Advance}
	}
	return obj.advanceHolder
}

// This contains advance properties of SecY.
// Advance returns a MacsecAdvance
func (obj *macsec) HasAdvance() bool {
	return obj.obj.Advance != nil
}

// This contains advance properties of SecY.
// SetAdvance sets the MacsecAdvance value in the Macsec object
func (obj *macsec) SetAdvance(value MacsecAdvance) Macsec {

	obj.advanceHolder = nil
	obj.obj.Advance = value.msg()

	return obj
}

func (obj *macsec) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Macsec")
	}

	// Basic is required
	if obj.obj.Basic == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Basic is required field on interface Macsec")
	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(vObj, set_default)
	}

	if len(obj.obj.Txscs) != 0 {

		if set_default {
			obj.Txscs().clearHolderSlice()
			for _, item := range obj.obj.Txscs {
				obj.Txscs().appendHolderSlice(&macsecTxSc{obj: item})
			}
		}
		for _, item := range obj.Txscs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Rxscs) != 0 {

		if set_default {
			obj.Rxscs().clearHolderSlice()
			for _, item := range obj.obj.Rxscs {
				obj.Rxscs().appendHolderSlice(&macsecRxSc{obj: item})
			}
		}
		for _, item := range obj.Rxscs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.CryptoEngine != nil {

		obj.CryptoEngine().validateObj(vObj, set_default)
	}

	if obj.obj.Advance != nil {

		obj.Advance().validateObj(vObj, set_default)
	}

}

func (obj *macsec) setDefault() {

}
