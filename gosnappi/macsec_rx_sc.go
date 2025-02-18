package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecRxSc *****
type macsecRxSc struct {
	validation
	obj          *otg.MacsecRxSc
	marshaller   marshalMacsecRxSc
	unMarshaller unMarshalMacsecRxSc
	saksHolder   MacsecRxScMacsecStaticKeySakIter
}

func NewMacsecRxSc() MacsecRxSc {
	obj := macsecRxSc{obj: &otg.MacsecRxSc{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecRxSc) msg() *otg.MacsecRxSc {
	return obj.obj
}

func (obj *macsecRxSc) setMsg(msg *otg.MacsecRxSc) MacsecRxSc {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecRxSc struct {
	obj *macsecRxSc
}

type marshalMacsecRxSc interface {
	// ToProto marshals MacsecRxSc to protobuf object *otg.MacsecRxSc
	ToProto() (*otg.MacsecRxSc, error)
	// ToPbText marshals MacsecRxSc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecRxSc to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecRxSc to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecRxSc struct {
	obj *macsecRxSc
}

type unMarshalMacsecRxSc interface {
	// FromProto unmarshals MacsecRxSc from protobuf object *otg.MacsecRxSc
	FromProto(msg *otg.MacsecRxSc) (MacsecRxSc, error)
	// FromPbText unmarshals MacsecRxSc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecRxSc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecRxSc from JSON text
	FromJson(value string) error
}

func (obj *macsecRxSc) Marshal() marshalMacsecRxSc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecRxSc{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecRxSc) Unmarshal() unMarshalMacsecRxSc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecRxSc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecRxSc) ToProto() (*otg.MacsecRxSc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecRxSc) FromProto(msg *otg.MacsecRxSc) (MacsecRxSc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecRxSc) ToPbText() (string, error) {
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

func (m *unMarshalmacsecRxSc) FromPbText(value string) error {
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

func (m *marshalmacsecRxSc) ToYaml() (string, error) {
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

func (m *unMarshalmacsecRxSc) FromYaml(value string) error {
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

func (m *marshalmacsecRxSc) ToJson() (string, error) {
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

func (m *unMarshalmacsecRxSc) FromJson(value string) error {
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

func (obj *macsecRxSc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecRxSc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecRxSc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecRxSc) Clone() (MacsecRxSc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecRxSc()
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

func (obj *macsecRxSc) setNil() {
	obj.saksHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecRxSc is rx SC settings.
type MacsecRxSc interface {
	Validation
	// msg marshals MacsecRxSc to protobuf object *otg.MacsecRxSc
	// and doesn't set defaults
	msg() *otg.MacsecRxSc
	// setMsg unmarshals MacsecRxSc from protobuf object *otg.MacsecRxSc
	// and doesn't set defaults
	setMsg(*otg.MacsecRxSc) MacsecRxSc
	// provides marshal interface
	Marshal() marshalMacsecRxSc
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecRxSc
	// validate validates MacsecRxSc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecRxSc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DutSystemId returns string, set in MacsecRxSc.
	DutSystemId() string
	// SetDutSystemId assigns string provided by user to MacsecRxSc
	SetDutSystemId(value string) MacsecRxSc
	// HasDutSystemId checks if DutSystemId has been set in MacsecRxSc
	HasDutSystemId() bool
	// DutPortId returns uint32, set in MacsecRxSc.
	DutPortId() uint32
	// SetDutPortId assigns uint32 provided by user to MacsecRxSc
	SetDutPortId(value uint32) MacsecRxSc
	// HasDutPortId checks if DutPortId has been set in MacsecRxSc
	HasDutPortId() bool
	// DutMsbXpn returns uint32, set in MacsecRxSc.
	DutMsbXpn() uint32
	// SetDutMsbXpn assigns uint32 provided by user to MacsecRxSc
	SetDutMsbXpn(value uint32) MacsecRxSc
	// HasDutMsbXpn checks if DutMsbXpn has been set in MacsecRxSc
	HasDutMsbXpn() bool
	// Saks returns MacsecRxScMacsecStaticKeySakIterIter, set in MacsecRxSc
	Saks() MacsecRxScMacsecStaticKeySakIter
	setNil()
}

// System ID in DUT SCI.
// DutSystemId returns a string
func (obj *macsecRxSc) DutSystemId() string {

	return *obj.obj.DutSystemId

}

// System ID in DUT SCI.
// DutSystemId returns a string
func (obj *macsecRxSc) HasDutSystemId() bool {
	return obj.obj.DutSystemId != nil
}

// System ID in DUT SCI.
// SetDutSystemId sets the string value in the MacsecRxSc object
func (obj *macsecRxSc) SetDutSystemId(value string) MacsecRxSc {

	obj.obj.DutSystemId = &value
	return obj
}

// Port ID in DUT SCI.
// DutPortId returns a uint32
func (obj *macsecRxSc) DutPortId() uint32 {

	return *obj.obj.DutPortId

}

// Port ID in DUT SCI.
// DutPortId returns a uint32
func (obj *macsecRxSc) HasDutPortId() bool {
	return obj.obj.DutPortId != nil
}

// Port ID in DUT SCI.
// SetDutPortId sets the uint32 value in the MacsecRxSc object
func (obj *macsecRxSc) SetDutPortId(value uint32) MacsecRxSc {

	obj.obj.DutPortId = &value
	return obj
}

// DUT MSB of XPN. The 32 most significant bits of the XPN that DUT will be using to construct the 64 bits XPN value when test starts.
// DutMsbXpn returns a uint32
func (obj *macsecRxSc) DutMsbXpn() uint32 {

	return *obj.obj.DutMsbXpn

}

// DUT MSB of XPN. The 32 most significant bits of the XPN that DUT will be using to construct the 64 bits XPN value when test starts.
// DutMsbXpn returns a uint32
func (obj *macsecRxSc) HasDutMsbXpn() bool {
	return obj.obj.DutMsbXpn != nil
}

// DUT MSB of XPN. The 32 most significant bits of the XPN that DUT will be using to construct the 64 bits XPN value when test starts.
// SetDutMsbXpn sets the uint32 value in the MacsecRxSc object
func (obj *macsecRxSc) SetDutMsbXpn(value uint32) MacsecRxSc {

	obj.obj.DutMsbXpn = &value
	return obj
}

// Rx SAK pool.
// Saks returns a []MacsecStaticKeySak
func (obj *macsecRxSc) Saks() MacsecRxScMacsecStaticKeySakIter {
	if len(obj.obj.Saks) == 0 {
		obj.obj.Saks = []*otg.MacsecStaticKeySak{}
	}
	if obj.saksHolder == nil {
		obj.saksHolder = newMacsecRxScMacsecStaticKeySakIter(&obj.obj.Saks).setMsg(obj)
	}
	return obj.saksHolder
}

type macsecRxScMacsecStaticKeySakIter struct {
	obj                     *macsecRxSc
	macsecStaticKeySakSlice []MacsecStaticKeySak
	fieldPtr                *[]*otg.MacsecStaticKeySak
}

func newMacsecRxScMacsecStaticKeySakIter(ptr *[]*otg.MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter {
	return &macsecRxScMacsecStaticKeySakIter{fieldPtr: ptr}
}

type MacsecRxScMacsecStaticKeySakIter interface {
	setMsg(*macsecRxSc) MacsecRxScMacsecStaticKeySakIter
	Items() []MacsecStaticKeySak
	Add() MacsecStaticKeySak
	Append(items ...MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter
	Set(index int, newObj MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter
	Clear() MacsecRxScMacsecStaticKeySakIter
	clearHolderSlice() MacsecRxScMacsecStaticKeySakIter
	appendHolderSlice(item MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter
}

func (obj *macsecRxScMacsecStaticKeySakIter) setMsg(msg *macsecRxSc) MacsecRxScMacsecStaticKeySakIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&macsecStaticKeySak{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *macsecRxScMacsecStaticKeySakIter) Items() []MacsecStaticKeySak {
	return obj.macsecStaticKeySakSlice
}

func (obj *macsecRxScMacsecStaticKeySakIter) Add() MacsecStaticKeySak {
	newObj := &otg.MacsecStaticKeySak{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &macsecStaticKeySak{obj: newObj}
	newLibObj.setDefault()
	obj.macsecStaticKeySakSlice = append(obj.macsecStaticKeySakSlice, newLibObj)
	return newLibObj
}

func (obj *macsecRxScMacsecStaticKeySakIter) Append(items ...MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.macsecStaticKeySakSlice = append(obj.macsecStaticKeySakSlice, item)
	}
	return obj
}

func (obj *macsecRxScMacsecStaticKeySakIter) Set(index int, newObj MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.macsecStaticKeySakSlice[index] = newObj
	return obj
}
func (obj *macsecRxScMacsecStaticKeySakIter) Clear() MacsecRxScMacsecStaticKeySakIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.MacsecStaticKeySak{}
		obj.macsecStaticKeySakSlice = []MacsecStaticKeySak{}
	}
	return obj
}
func (obj *macsecRxScMacsecStaticKeySakIter) clearHolderSlice() MacsecRxScMacsecStaticKeySakIter {
	if len(obj.macsecStaticKeySakSlice) > 0 {
		obj.macsecStaticKeySakSlice = []MacsecStaticKeySak{}
	}
	return obj
}
func (obj *macsecRxScMacsecStaticKeySakIter) appendHolderSlice(item MacsecStaticKeySak) MacsecRxScMacsecStaticKeySakIter {
	obj.macsecStaticKeySakSlice = append(obj.macsecStaticKeySakSlice, item)
	return obj
}

func (obj *macsecRxSc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DutSystemId != nil {

		err := obj.validateMac(obj.DutSystemId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MacsecRxSc.DutSystemId"))
		}

	}

	if obj.obj.DutPortId != nil {

		if *obj.obj.DutPortId < 1 || *obj.obj.DutPortId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= MacsecRxSc.DutPortId <= 65535 but Got %d", *obj.obj.DutPortId))
		}

	}

	if obj.obj.DutMsbXpn != nil {

		if *obj.obj.DutMsbXpn > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MacsecRxSc.DutMsbXpn <= 4294967295 but Got %d", *obj.obj.DutMsbXpn))
		}

	}

	if len(obj.obj.Saks) != 0 {

		if set_default {
			obj.Saks().clearHolderSlice()
			for _, item := range obj.obj.Saks {
				obj.Saks().appendHolderSlice(&macsecStaticKeySak{obj: item})
			}
		}
		for _, item := range obj.Saks().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *macsecRxSc) setDefault() {
	if obj.obj.DutPortId == nil {
		obj.SetDutPortId(1)
	}
	if obj.obj.DutMsbXpn == nil {
		obj.SetDutMsbXpn(0)
	}

}
