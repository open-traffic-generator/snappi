package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecCryptoEngineStatelessEncryptionOnly *****
type macsecCryptoEngineStatelessEncryptionOnly struct {
	validation
	obj           *otg.MacsecCryptoEngineStatelessEncryptionOnly
	marshaller    marshalMacsecCryptoEngineStatelessEncryptionOnly
	unMarshaller  unMarshalMacsecCryptoEngineStatelessEncryptionOnly
	txPnHolder    MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	trafficHolder MacsecCryptoEngineStatelessEncryptionOnlyTraffic
}

func NewMacsecCryptoEngineStatelessEncryptionOnly() MacsecCryptoEngineStatelessEncryptionOnly {
	obj := macsecCryptoEngineStatelessEncryptionOnly{obj: &otg.MacsecCryptoEngineStatelessEncryptionOnly{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) msg() *otg.MacsecCryptoEngineStatelessEncryptionOnly {
	return obj.obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) setMsg(msg *otg.MacsecCryptoEngineStatelessEncryptionOnly) MacsecCryptoEngineStatelessEncryptionOnly {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecCryptoEngineStatelessEncryptionOnly struct {
	obj *macsecCryptoEngineStatelessEncryptionOnly
}

type marshalMacsecCryptoEngineStatelessEncryptionOnly interface {
	// ToProto marshals MacsecCryptoEngineStatelessEncryptionOnly to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnly
	ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnly, error)
	// ToPbText marshals MacsecCryptoEngineStatelessEncryptionOnly to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecCryptoEngineStatelessEncryptionOnly to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecCryptoEngineStatelessEncryptionOnly to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecCryptoEngineStatelessEncryptionOnly struct {
	obj *macsecCryptoEngineStatelessEncryptionOnly
}

type unMarshalMacsecCryptoEngineStatelessEncryptionOnly interface {
	// FromProto unmarshals MacsecCryptoEngineStatelessEncryptionOnly from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnly
	FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnly) (MacsecCryptoEngineStatelessEncryptionOnly, error)
	// FromPbText unmarshals MacsecCryptoEngineStatelessEncryptionOnly from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecCryptoEngineStatelessEncryptionOnly from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecCryptoEngineStatelessEncryptionOnly from JSON text
	FromJson(value string) error
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnly {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecCryptoEngineStatelessEncryptionOnly{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnly {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecCryptoEngineStatelessEncryptionOnly{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnly) ToProto() (*otg.MacsecCryptoEngineStatelessEncryptionOnly, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnly) FromProto(msg *otg.MacsecCryptoEngineStatelessEncryptionOnly) (MacsecCryptoEngineStatelessEncryptionOnly, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnly) ToPbText() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnly) FromPbText(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnly) ToYaml() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnly) FromYaml(value string) error {
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

func (m *marshalmacsecCryptoEngineStatelessEncryptionOnly) ToJson() (string, error) {
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

func (m *unMarshalmacsecCryptoEngineStatelessEncryptionOnly) FromJson(value string) error {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnly) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) Clone() (MacsecCryptoEngineStatelessEncryptionOnly, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecCryptoEngineStatelessEncryptionOnly()
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

func (obj *macsecCryptoEngineStatelessEncryptionOnly) setNil() {
	obj.txPnHolder = nil
	obj.trafficHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecCryptoEngineStatelessEncryptionOnly is the container for stateless encryption only engine configuration.
type MacsecCryptoEngineStatelessEncryptionOnly interface {
	Validation
	// msg marshals MacsecCryptoEngineStatelessEncryptionOnly to protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnly
	// and doesn't set defaults
	msg() *otg.MacsecCryptoEngineStatelessEncryptionOnly
	// setMsg unmarshals MacsecCryptoEngineStatelessEncryptionOnly from protobuf object *otg.MacsecCryptoEngineStatelessEncryptionOnly
	// and doesn't set defaults
	setMsg(*otg.MacsecCryptoEngineStatelessEncryptionOnly) MacsecCryptoEngineStatelessEncryptionOnly
	// provides marshal interface
	Marshal() marshalMacsecCryptoEngineStatelessEncryptionOnly
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecCryptoEngineStatelessEncryptionOnly
	// validate validates MacsecCryptoEngineStatelessEncryptionOnly
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecCryptoEngineStatelessEncryptionOnly, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns MacsecCryptoEngineStatelessEncryptionOnlyTxPn, set in MacsecCryptoEngineStatelessEncryptionOnly.
	// MacsecCryptoEngineStatelessEncryptionOnlyTxPn is tx packet number(PN) configuration.
	TxPn() MacsecCryptoEngineStatelessEncryptionOnlyTxPn
	// SetTxPn assigns MacsecCryptoEngineStatelessEncryptionOnlyTxPn provided by user to MacsecCryptoEngineStatelessEncryptionOnly.
	// MacsecCryptoEngineStatelessEncryptionOnlyTxPn is tx packet number(PN) configuration.
	SetTxPn(value MacsecCryptoEngineStatelessEncryptionOnlyTxPn) MacsecCryptoEngineStatelessEncryptionOnly
	// HasTxPn checks if TxPn has been set in MacsecCryptoEngineStatelessEncryptionOnly
	HasTxPn() bool
	// Traffic returns MacsecCryptoEngineStatelessEncryptionOnlyTraffic, set in MacsecCryptoEngineStatelessEncryptionOnly.
	// MacsecCryptoEngineStatelessEncryptionOnlyTraffic is encryption only traffic configuration.
	Traffic() MacsecCryptoEngineStatelessEncryptionOnlyTraffic
	// SetTraffic assigns MacsecCryptoEngineStatelessEncryptionOnlyTraffic provided by user to MacsecCryptoEngineStatelessEncryptionOnly.
	// MacsecCryptoEngineStatelessEncryptionOnlyTraffic is encryption only traffic configuration.
	SetTraffic(value MacsecCryptoEngineStatelessEncryptionOnlyTraffic) MacsecCryptoEngineStatelessEncryptionOnly
	// HasTraffic checks if Traffic has been set in MacsecCryptoEngineStatelessEncryptionOnly
	HasTraffic() bool
	setNil()
}

// description is TBD
// TxPn returns a MacsecCryptoEngineStatelessEncryptionOnlyTxPn
func (obj *macsecCryptoEngineStatelessEncryptionOnly) TxPn() MacsecCryptoEngineStatelessEncryptionOnlyTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewMacsecCryptoEngineStatelessEncryptionOnlyTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &macsecCryptoEngineStatelessEncryptionOnlyTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a MacsecCryptoEngineStatelessEncryptionOnlyTxPn
func (obj *macsecCryptoEngineStatelessEncryptionOnly) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the MacsecCryptoEngineStatelessEncryptionOnlyTxPn value in the MacsecCryptoEngineStatelessEncryptionOnly object
func (obj *macsecCryptoEngineStatelessEncryptionOnly) SetTxPn(value MacsecCryptoEngineStatelessEncryptionOnlyTxPn) MacsecCryptoEngineStatelessEncryptionOnly {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

// description is TBD
// Traffic returns a MacsecCryptoEngineStatelessEncryptionOnlyTraffic
func (obj *macsecCryptoEngineStatelessEncryptionOnly) Traffic() MacsecCryptoEngineStatelessEncryptionOnlyTraffic {
	if obj.obj.Traffic == nil {
		obj.obj.Traffic = NewMacsecCryptoEngineStatelessEncryptionOnlyTraffic().msg()
	}
	if obj.trafficHolder == nil {
		obj.trafficHolder = &macsecCryptoEngineStatelessEncryptionOnlyTraffic{obj: obj.obj.Traffic}
	}
	return obj.trafficHolder
}

// description is TBD
// Traffic returns a MacsecCryptoEngineStatelessEncryptionOnlyTraffic
func (obj *macsecCryptoEngineStatelessEncryptionOnly) HasTraffic() bool {
	return obj.obj.Traffic != nil
}

// description is TBD
// SetTraffic sets the MacsecCryptoEngineStatelessEncryptionOnlyTraffic value in the MacsecCryptoEngineStatelessEncryptionOnly object
func (obj *macsecCryptoEngineStatelessEncryptionOnly) SetTraffic(value MacsecCryptoEngineStatelessEncryptionOnlyTraffic) MacsecCryptoEngineStatelessEncryptionOnly {

	obj.trafficHolder = nil
	obj.obj.Traffic = value.msg()

	return obj
}

func (obj *macsecCryptoEngineStatelessEncryptionOnly) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecCryptoEngineStatelessEncryptionOnly) setDefault() {

}
