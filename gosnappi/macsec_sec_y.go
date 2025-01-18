package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecY *****
type macsecSecY struct {
	validation
	obj                *otg.MacsecSecY
	marshaller         marshalMacsecSecY
	unMarshaller       unMarshalMacsecSecY
	basicHolder        MacsecSecYBasic
	txscsHolder        MacsecSecYMacsecSecYTxScIter
	rxscsHolder        MacsecSecYMacsecSecYRxScIter
	cryptoEngineHolder MacsecSecYCryptoEngine
	advanceHolder      MacsecSecYAdvance
}

func NewMacsecSecY() MacsecSecY {
	obj := macsecSecY{obj: &otg.MacsecSecY{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecY) msg() *otg.MacsecSecY {
	return obj.obj
}

func (obj *macsecSecY) setMsg(msg *otg.MacsecSecY) MacsecSecY {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecY struct {
	obj *macsecSecY
}

type marshalMacsecSecY interface {
	// ToProto marshals MacsecSecY to protobuf object *otg.MacsecSecY
	ToProto() (*otg.MacsecSecY, error)
	// ToPbText marshals MacsecSecY to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecY to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecY to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecY struct {
	obj *macsecSecY
}

type unMarshalMacsecSecY interface {
	// FromProto unmarshals MacsecSecY from protobuf object *otg.MacsecSecY
	FromProto(msg *otg.MacsecSecY) (MacsecSecY, error)
	// FromPbText unmarshals MacsecSecY from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecY from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecY from JSON text
	FromJson(value string) error
}

func (obj *macsecSecY) Marshal() marshalMacsecSecY {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecY{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecY) Unmarshal() unMarshalMacsecSecY {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecY{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecY) ToProto() (*otg.MacsecSecY, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecY) FromProto(msg *otg.MacsecSecY) (MacsecSecY, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecY) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecY) FromPbText(value string) error {
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

func (m *marshalmacsecSecY) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecY) FromYaml(value string) error {
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

func (m *marshalmacsecSecY) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecY) FromJson(value string) error {
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

func (obj *macsecSecY) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecY) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecY) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecY) Clone() (MacsecSecY, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecY()
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

func (obj *macsecSecY) setNil() {
	obj.basicHolder = nil
	obj.txscsHolder = nil
	obj.rxscsHolder = nil
	obj.cryptoEngineHolder = nil
	obj.advanceHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecY is configuration for a Secure Entity (SecY).
type MacsecSecY interface {
	Validation
	// msg marshals MacsecSecY to protobuf object *otg.MacsecSecY
	// and doesn't set defaults
	msg() *otg.MacsecSecY
	// setMsg unmarshals MacsecSecY from protobuf object *otg.MacsecSecY
	// and doesn't set defaults
	setMsg(*otg.MacsecSecY) MacsecSecY
	// provides marshal interface
	Marshal() marshalMacsecSecY
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecY
	// validate validates MacsecSecY
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecY, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in MacsecSecY.
	Name() string
	// SetName assigns string provided by user to MacsecSecY
	SetName(value string) MacsecSecY
	// Basic returns MacsecSecYBasic, set in MacsecSecY.
	// MacsecSecYBasic is a container of basic properties for a SecY.
	Basic() MacsecSecYBasic
	// SetBasic assigns MacsecSecYBasic provided by user to MacsecSecY.
	// MacsecSecYBasic is a container of basic properties for a SecY.
	SetBasic(value MacsecSecYBasic) MacsecSecY
	// Txscs returns MacsecSecYMacsecSecYTxScIterIter, set in MacsecSecY
	Txscs() MacsecSecYMacsecSecYTxScIter
	// Rxscs returns MacsecSecYMacsecSecYRxScIterIter, set in MacsecSecY
	Rxscs() MacsecSecYMacsecSecYRxScIter
	// CryptoEngine returns MacsecSecYCryptoEngine, set in MacsecSecY.
	// MacsecSecYCryptoEngine is a container of crypto engine properties of a SecY.
	CryptoEngine() MacsecSecYCryptoEngine
	// SetCryptoEngine assigns MacsecSecYCryptoEngine provided by user to MacsecSecY.
	// MacsecSecYCryptoEngine is a container of crypto engine properties of a SecY.
	SetCryptoEngine(value MacsecSecYCryptoEngine) MacsecSecY
	// Advance returns MacsecSecYAdvance, set in MacsecSecY.
	// MacsecSecYAdvance is a container of advance properties SecY.
	Advance() MacsecSecYAdvance
	// SetAdvance assigns MacsecSecYAdvance provided by user to MacsecSecY.
	// MacsecSecYAdvance is a container of advance properties SecY.
	SetAdvance(value MacsecSecYAdvance) MacsecSecY
	// HasAdvance checks if Advance has been set in MacsecSecY
	HasAdvance() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *macsecSecY) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the MacsecSecY object
func (obj *macsecSecY) SetName(value string) MacsecSecY {

	obj.obj.Name = &value
	return obj
}

// This contains the basic properties of SecY.
// Basic returns a MacsecSecYBasic
func (obj *macsecSecY) Basic() MacsecSecYBasic {
	if obj.obj.Basic == nil {
		obj.obj.Basic = NewMacsecSecYBasic().msg()
	}
	if obj.basicHolder == nil {
		obj.basicHolder = &macsecSecYBasic{obj: obj.obj.Basic}
	}
	return obj.basicHolder
}

// This contains the basic properties of SecY.
// SetBasic sets the MacsecSecYBasic value in the MacsecSecY object
func (obj *macsecSecY) SetBasic(value MacsecSecYBasic) MacsecSecY {

	obj.basicHolder = nil
	obj.obj.Basic = value.msg()

	return obj
}

// Tx secure channels.
// Txscs returns a []MacsecSecYTxSc
func (obj *macsecSecY) Txscs() MacsecSecYMacsecSecYTxScIter {
	if len(obj.obj.Txscs) == 0 {
		obj.obj.Txscs = []*otg.MacsecSecYTxSc{}
	}
	if obj.txscsHolder == nil {
		obj.txscsHolder = newMacsecSecYMacsecSecYTxScIter(&obj.obj.Txscs).setMsg(obj)
	}
	return obj.txscsHolder
}

type macsecSecYMacsecSecYTxScIter struct {
	obj                 *macsecSecY
	macsecSecYTxScSlice []MacsecSecYTxSc
	fieldPtr            *[]*otg.MacsecSecYTxSc
}

func newMacsecSecYMacsecSecYTxScIter(ptr *[]*otg.MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter {
	return &macsecSecYMacsecSecYTxScIter{fieldPtr: ptr}
}

type MacsecSecYMacsecSecYTxScIter interface {
	setMsg(*macsecSecY) MacsecSecYMacsecSecYTxScIter
	Items() []MacsecSecYTxSc
	Add() MacsecSecYTxSc
	Append(items ...MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter
	Set(index int, newObj MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter
	Clear() MacsecSecYMacsecSecYTxScIter
	clearHolderSlice() MacsecSecYMacsecSecYTxScIter
	appendHolderSlice(item MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter
}

func (obj *macsecSecYMacsecSecYTxScIter) setMsg(msg *macsecSecY) MacsecSecYMacsecSecYTxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecSecYTxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecSecYMacsecSecYTxScIter) Items() []MacsecSecYTxSc {
	return obj.macsecSecYTxScSlice
}

func (obj *macsecSecYMacsecSecYTxScIter) Add() MacsecSecYTxSc {
	newObj := &otg.MacsecSecYTxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecSecYTxSc{obj: newObj}
	newLibObj.setDefault()
	obj.macsecSecYTxScSlice = append(obj.macsecSecYTxScSlice, newLibObj)
	return newLibObj
}

func (obj *macsecSecYMacsecSecYTxScIter) Append(items ...MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecSecYTxScSlice = append(obj.macsecSecYTxScSlice, item)
	}
	return obj
}

func (obj *macsecSecYMacsecSecYTxScIter) Set(index int, newObj MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecSecYTxScSlice[index] = newObj
	return obj
}
func (obj *macsecSecYMacsecSecYTxScIter) Clear() MacsecSecYMacsecSecYTxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecSecYTxSc{}
		obj.macsecSecYTxScSlice = []MacsecSecYTxSc{}
	}
	return obj
}
func (obj *macsecSecYMacsecSecYTxScIter) clearHolderSlice() MacsecSecYMacsecSecYTxScIter {
	if len(obj.macsecSecYTxScSlice) > 0 {
		obj.macsecSecYTxScSlice = []MacsecSecYTxSc{}
	}
	return obj
}
func (obj *macsecSecYMacsecSecYTxScIter) appendHolderSlice(item MacsecSecYTxSc) MacsecSecYMacsecSecYTxScIter {
	obj.macsecSecYTxScSlice = append(obj.macsecSecYTxScSlice, item)
	return obj
}

// Rx secure channels.
// Rxscs returns a []MacsecSecYRxSc
func (obj *macsecSecY) Rxscs() MacsecSecYMacsecSecYRxScIter {
	if len(obj.obj.Rxscs) == 0 {
		obj.obj.Rxscs = []*otg.MacsecSecYRxSc{}
	}
	if obj.rxscsHolder == nil {
		obj.rxscsHolder = newMacsecSecYMacsecSecYRxScIter(&obj.obj.Rxscs).setMsg(obj)
	}
	return obj.rxscsHolder
}

type macsecSecYMacsecSecYRxScIter struct {
	obj                 *macsecSecY
	macsecSecYRxScSlice []MacsecSecYRxSc
	fieldPtr            *[]*otg.MacsecSecYRxSc
}

func newMacsecSecYMacsecSecYRxScIter(ptr *[]*otg.MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter {
	return &macsecSecYMacsecSecYRxScIter{fieldPtr: ptr}
}

type MacsecSecYMacsecSecYRxScIter interface {
	setMsg(*macsecSecY) MacsecSecYMacsecSecYRxScIter
	Items() []MacsecSecYRxSc
	Add() MacsecSecYRxSc
	Append(items ...MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter
	Set(index int, newObj MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter
	Clear() MacsecSecYMacsecSecYRxScIter
	clearHolderSlice() MacsecSecYMacsecSecYRxScIter
	appendHolderSlice(item MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter
}

func (obj *macsecSecYMacsecSecYRxScIter) setMsg(msg *macsecSecY) MacsecSecYMacsecSecYRxScIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecSecYRxSc{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecSecYMacsecSecYRxScIter) Items() []MacsecSecYRxSc {
	return obj.macsecSecYRxScSlice
}

func (obj *macsecSecYMacsecSecYRxScIter) Add() MacsecSecYRxSc {
	newObj := &otg.MacsecSecYRxSc{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecSecYRxSc{obj: newObj}
	newLibObj.setDefault()
	obj.macsecSecYRxScSlice = append(obj.macsecSecYRxScSlice, newLibObj)
	return newLibObj
}

func (obj *macsecSecYMacsecSecYRxScIter) Append(items ...MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecSecYRxScSlice = append(obj.macsecSecYRxScSlice, item)
	}
	return obj
}

func (obj *macsecSecYMacsecSecYRxScIter) Set(index int, newObj MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecSecYRxScSlice[index] = newObj
	return obj
}
func (obj *macsecSecYMacsecSecYRxScIter) Clear() MacsecSecYMacsecSecYRxScIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecSecYRxSc{}
		obj.macsecSecYRxScSlice = []MacsecSecYRxSc{}
	}
	return obj
}
func (obj *macsecSecYMacsecSecYRxScIter) clearHolderSlice() MacsecSecYMacsecSecYRxScIter {
	if len(obj.macsecSecYRxScSlice) > 0 {
		obj.macsecSecYRxScSlice = []MacsecSecYRxSc{}
	}
	return obj
}
func (obj *macsecSecYMacsecSecYRxScIter) appendHolderSlice(item MacsecSecYRxSc) MacsecSecYMacsecSecYRxScIter {
	obj.macsecSecYRxScSlice = append(obj.macsecSecYRxScSlice, item)
	return obj
}

// This contains crypto engine properties of SecY.
// CryptoEngine returns a MacsecSecYCryptoEngine
func (obj *macsecSecY) CryptoEngine() MacsecSecYCryptoEngine {
	if obj.obj.CryptoEngine == nil {
		obj.obj.CryptoEngine = NewMacsecSecYCryptoEngine().msg()
	}
	if obj.cryptoEngineHolder == nil {
		obj.cryptoEngineHolder = &macsecSecYCryptoEngine{obj: obj.obj.CryptoEngine}
	}
	return obj.cryptoEngineHolder
}

// This contains crypto engine properties of SecY.
// SetCryptoEngine sets the MacsecSecYCryptoEngine value in the MacsecSecY object
func (obj *macsecSecY) SetCryptoEngine(value MacsecSecYCryptoEngine) MacsecSecY {

	obj.cryptoEngineHolder = nil
	obj.obj.CryptoEngine = value.msg()

	return obj
}

// This contains advance properties of SecY.
// Advance returns a MacsecSecYAdvance
func (obj *macsecSecY) Advance() MacsecSecYAdvance {
	if obj.obj.Advance == nil {
		obj.obj.Advance = NewMacsecSecYAdvance().msg()
	}
	if obj.advanceHolder == nil {
		obj.advanceHolder = &macsecSecYAdvance{obj: obj.obj.Advance}
	}
	return obj.advanceHolder
}

// This contains advance properties of SecY.
// Advance returns a MacsecSecYAdvance
func (obj *macsecSecY) HasAdvance() bool {
	return obj.obj.Advance != nil
}

// This contains advance properties of SecY.
// SetAdvance sets the MacsecSecYAdvance value in the MacsecSecY object
func (obj *macsecSecY) SetAdvance(value MacsecSecYAdvance) MacsecSecY {

	obj.advanceHolder = nil
	obj.obj.Advance = value.msg()

	return obj
}

func (obj *macsecSecY) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface MacsecSecY")
	}

	// Basic is required
	if obj.obj.Basic == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Basic is required field on interface MacsecSecY")
	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(vObj, set_default)
	}

	if len(obj.obj.Txscs) != 0 {

		if set_default {
			obj.Txscs().clearHolderSlice()
			for _, item := range obj.obj.Txscs {
				obj.Txscs().appendHolderSlice(&macsecSecYTxSc{obj: item})
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
				obj.Rxscs().appendHolderSlice(&macsecSecYRxSc{obj: item})
			}
		}
		for _, item := range obj.Rxscs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// CryptoEngine is required
	if obj.obj.CryptoEngine == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CryptoEngine is required field on interface MacsecSecY")
	}

	if obj.obj.CryptoEngine != nil {

		obj.CryptoEngine().validateObj(vObj, set_default)
	}

	if obj.obj.Advance != nil {

		obj.Advance().validateObj(vObj, set_default)
	}

}

func (obj *macsecSecY) setDefault() {

}
