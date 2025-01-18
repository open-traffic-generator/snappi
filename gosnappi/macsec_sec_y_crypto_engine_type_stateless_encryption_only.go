package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecSecYCryptoEngineTypeStatelessEncryptionOnly *****
type macsecSecYCryptoEngineTypeStatelessEncryptionOnly struct {
	validation
	obj           *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	marshaller    marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	unMarshaller  unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	txPnHolder    MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	trafficHolder MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
}

func NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnly() MacsecSecYCryptoEngineTypeStatelessEncryptionOnly {
	obj := macsecSecYCryptoEngineTypeStatelessEncryptionOnly{obj: &otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly {
	return obj.obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) setMsg(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly) MacsecSecYCryptoEngineTypeStatelessEncryptionOnly {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly
}

type marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly interface {
	// ToProto marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, error)
	// ToPbText marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly struct {
	obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly
}

type unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly interface {
	// FromProto unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, error)
	// FromPbText unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly from JSON text
	FromJson(value string) error
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) ToProto() (*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) FromProto(msg *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly) (MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) ToPbText() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) FromPbText(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) ToYaml() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) FromYaml(value string) error {
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

func (m *marshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) ToJson() (string, error) {
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

func (m *unMarshalmacsecSecYCryptoEngineTypeStatelessEncryptionOnly) FromJson(value string) error {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnly()
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) setNil() {
	obj.txPnHolder = nil
	obj.trafficHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecSecYCryptoEngineTypeStatelessEncryptionOnly is the container for stateless encryption only engine settings.
type MacsecSecYCryptoEngineTypeStatelessEncryptionOnly interface {
	Validation
	// msg marshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly to protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// and doesn't set defaults
	msg() *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// setMsg unmarshals MacsecSecYCryptoEngineTypeStatelessEncryptionOnly from protobuf object *otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// and doesn't set defaults
	setMsg(*otg.MacsecSecYCryptoEngineTypeStatelessEncryptionOnly) MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// provides marshal interface
	Marshal() marshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// validate validates MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecSecYCryptoEngineTypeStatelessEncryptionOnly, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TxPn returns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
	TxPn() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
	// SetTxPn assigns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn is tx PN settings.
	SetTxPn(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// HasTxPn checks if TxPn has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	HasTxPn() bool
	// Traffic returns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic, set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic is encryption only traffic settings.
	Traffic() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
	// SetTraffic assigns MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic provided by user to MacsecSecYCryptoEngineTypeStatelessEncryptionOnly.
	// MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic is encryption only traffic settings.
	SetTraffic(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	// HasTraffic checks if Traffic has been set in MacsecSecYCryptoEngineTypeStatelessEncryptionOnly
	HasTraffic() bool
	setNil()
}

// description is TBD
// TxPn returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) TxPn() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn {
	if obj.obj.TxPn == nil {
		obj.obj.TxPn = NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn().msg()
	}
	if obj.txPnHolder == nil {
		obj.txPnHolder = &macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn{obj: obj.obj.TxPn}
	}
	return obj.txPnHolder
}

// description is TBD
// TxPn returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) HasTxPn() bool {
	return obj.obj.TxPn != nil
}

// description is TBD
// SetTxPn sets the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnly object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) SetTxPn(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTxPn) MacsecSecYCryptoEngineTypeStatelessEncryptionOnly {

	obj.txPnHolder = nil
	obj.obj.TxPn = value.msg()

	return obj
}

// description is TBD
// Traffic returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) Traffic() MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic {
	if obj.obj.Traffic == nil {
		obj.obj.Traffic = NewMacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic().msg()
	}
	if obj.trafficHolder == nil {
		obj.trafficHolder = &macsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic{obj: obj.obj.Traffic}
	}
	return obj.trafficHolder
}

// description is TBD
// Traffic returns a MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) HasTraffic() bool {
	return obj.obj.Traffic != nil
}

// description is TBD
// SetTraffic sets the MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic value in the MacsecSecYCryptoEngineTypeStatelessEncryptionOnly object
func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) SetTraffic(value MacsecSecYCryptoEngineTypeStatelessEncryptionOnlyTraffic) MacsecSecYCryptoEngineTypeStatelessEncryptionOnly {

	obj.trafficHolder = nil
	obj.obj.Traffic = value.msg()

	return obj
}

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) validateObj(vObj *validation, set_default bool) {
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

func (obj *macsecSecYCryptoEngineTypeStatelessEncryptionOnly) setDefault() {

}
