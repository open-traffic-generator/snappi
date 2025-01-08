package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineTypeStatelessEncryptionOnly *****
type macsecCryptoEngineTypeStatelessEncryptionOnly struct {
	validation
	obj           *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly
	marshaller    marshalMacsecCryptoEngineTypeStatelessEncryptionOnly
	unMarshaller  unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnly
	txPnHolder    MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	trafficHolder MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

func NewMacsecCryptoEngineTypeStatelessEncryptionOnly() MacsecCryptoEngineTypeStatelessEncryptionOnly {
	obj := macsecCryptoEngineTypeStatelessEncryptionOnly{obj: &otg.MacsecCryptoEngineTypeStatelessEncryptionOnly{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly {
	return obj.obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) setMsg(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly) MacsecCryptoEngineTypeStatelessEncryptionOnly {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineTypeStatelessEncryptionOnly struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnly
}

type marshalMacsecCryptoEngineTypeStatelessEncryptionOnly interface {
	// ToProto marshals MacsecCryptoEngineTypeStatelessEncryptionOnly to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly
	ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnly, error)
	// ToPbText marshals MacsecCryptoEngineTypeStatelessEncryptionOnly to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineTypeStatelessEncryptionOnly to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineTypeStatelessEncryptionOnly to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnly struct {
	obj *macsecCryptoEngineTypeStatelessEncryptionOnly
}

type unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnly interface {
	// FromProto unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnly from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly
	FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly) (MacsecCryptoEngineTypeStatelessEncryptionOnly, error)
	// FromPbText unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnly from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnly from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnly from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnly {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineTypeStatelessEncryptionOnly{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnly {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnly{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnly) ToProto() (*otg.MacsecCryptoEngineTypeStatelessEncryptionOnly, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnly) FromProto(msg *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly) (MacsecCryptoEngineTypeStatelessEncryptionOnly, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnly) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnly) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnly) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnly) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineTypeStatelessEncryptionOnly) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineTypeStatelessEncryptionOnly) FromJson(value string) error {
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnly, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineTypeStatelessEncryptionOnly()
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

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) setNil() {
	obj.txPnHolder = nil
	obj.trafficHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineTypeStatelessEncryptionOnly is the container for stateless encryption only engine settings.
type MacsecCryptoEngineTypeStatelessEncryptionOnly interface {
	Validation
	// msg marshals MacsecCryptoEngineTypeStatelessEncryptionOnly to protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly
	// setMsg unmarshals MacsecCryptoEngineTypeStatelessEncryptionOnly from protobuf object *otg.MacsecCryptoEngineTypeStatelessEncryptionOnly
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineTypeStatelessEncryptionOnly) MacsecCryptoEngineTypeStatelessEncryptionOnly
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineTypeStatelessEncryptionOnly
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineTypeStatelessEncryptionOnly
	// validate validates MacsecCryptoEngineTypeStatelessEncryptionOnly
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineTypeStatelessEncryptionOnly, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn, set in MacsecCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
	// choice:
	// description: >-
	// Types of Tx PN series.
	// type: string
	// default: fixed_pn
	// x-field-uid: 1
	// x-enum:
	// fixed_pn:
	// x-field-uid: 1
	// incrementing_pn:
	// x-field-uid: 2
	// fixed:
	// $ref: '#/components/schemas/Macsec.CryptoEngine.Type.StatelessEncryptionOnly.FixedPn'
	// x-field-uid: 2
	// incrementing:
	// $ref: '#/components/schemas/Macsec.CryptoEngine.Type.StatelessEncryptionOnly.IncrementingPn'
	// x-field-uid: 3
	TxPn() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// SetTxPn assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
	// choice:
	// description: >-
	// Types of Tx PN series.
	// type: string
	// default: fixed_pn
	// x-field-uid: 1
	// x-enum:
	// fixed_pn:
	// x-field-uid: 1
	// incrementing_pn:
	// x-field-uid: 2
	// fixed:
	// $ref: '#/components/schemas/Macsec.CryptoEngine.Type.StatelessEncryptionOnly.FixedPn'
	// x-field-uid: 2
	// incrementing:
	// $ref: '#/components/schemas/Macsec.CryptoEngine.Type.StatelessEncryptionOnly.IncrementingPn'
	// x-field-uid: 3
	SetTxPn(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecCryptoEngineTypeStatelessEncryptionOnly
	// HasTxPn checks if TxPn has been set in MacsecCryptoEngineTypeStatelessEncryptionOnly
	HasTxPn() bool
	// Traffic returns MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic, set in MacsecCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic is encryption only traffic settings.
	Traffic() MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// SetTraffic assigns MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic provided by user to MacsecCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic is encryption only traffic settings.
	SetTraffic(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecCryptoEngineTypeStatelessEncryptionOnly
	// HasTraffic checks if Traffic has been set in MacsecCryptoEngineTypeStatelessEncryptionOnly
	HasTraffic() bool
	setNil()
}

// description is TBD
// TxPn returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) TxPn() MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &macsecCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn value in the MacsecCryptoEngineTypeStatelessEncryptionOnly object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) SetTxPn(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecCryptoEngineTypeStatelessEncryptionOnly {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

// description is TBD
// Traffic returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) Traffic() MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	if obj.obj.Traffic == nil {
		obj.obj.Traffic = NewMacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic().msg()
	}
	if obj.trafficHolder == nil {
		obj.trafficHolder = &macsecCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: obj.obj.Traffic}
	}
	return obj.trafficHolder
}

// description is TBD
// Traffic returns a MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) HasTraffic() bool {
	return obj.obj.Traffic != nil
}

// description is TBD
// SetTraffic sets the MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic value in the MacsecCryptoEngineTypeStatelessEncryptionOnly object
func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) SetTraffic(value MacsecCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecCryptoEngineTypeStatelessEncryptionOnly {

	obj.trafficHolder = nil
	obj.obj.Traffic = value.msg()

	return obj
}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TxPn != nil {

		obj.TxPn().validateObj(vObj, set_default)
	}

	if obj.obj.Traffic != nil {

		obj.Traffic().validateObj(vObj, set_default)
	}

}

func (obj *macsecCryptoEngineTypeStatelessEncryptionOnly) setDefault() {

}
