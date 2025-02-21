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
	staticKeyHolder    MacsecStaticKey
	txHolder           MacsecTx
	rxHolder           MacsecRx
	cryptoEngineHolder MacsecCryptoEngine
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
	obj.staticKeyHolder = nil
	obj.txHolder = nil
	obj.rxHolder = nil
	obj.cryptoEngineHolder = nil
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
	// StaticKey returns MacsecStaticKey, set in Macsec.
	// MacsecStaticKey is a container of static key properties for a secure entity(SecY). This configuration is applicable when no dynamic key management protocol i.e. MACsec key agreement(MKA) is configured. If MKA is configured, any static key configuration is not applicable.
	StaticKey() MacsecStaticKey
	// SetStaticKey assigns MacsecStaticKey provided by user to Macsec.
	// MacsecStaticKey is a container of static key properties for a secure entity(SecY). This configuration is applicable when no dynamic key management protocol i.e. MACsec key agreement(MKA) is configured. If MKA is configured, any static key configuration is not applicable.
	SetStaticKey(value MacsecStaticKey) Macsec
	// HasStaticKey checks if StaticKey has been set in Macsec
	HasStaticKey() bool
	// Tx returns MacsecTx, set in Macsec.
	// MacsecTx is a container of Tx properties of SecY.
	Tx() MacsecTx
	// SetTx assigns MacsecTx provided by user to Macsec.
	// MacsecTx is a container of Tx properties of SecY.
	SetTx(value MacsecTx) Macsec
	// HasTx checks if Tx has been set in Macsec
	HasTx() bool
	// Rx returns MacsecRx, set in Macsec.
	// MacsecRx is a container for Rx settings of SecY.
	Rx() MacsecRx
	// SetRx assigns MacsecRx provided by user to Macsec.
	// MacsecRx is a container for Rx settings of SecY.
	SetRx(value MacsecRx) Macsec
	// HasRx checks if Rx has been set in Macsec
	HasRx() bool
	// CryptoEngine returns MacsecCryptoEngine, set in Macsec.
	// MacsecCryptoEngine is a container of crypto engine properties of a SecY.
	CryptoEngine() MacsecCryptoEngine
	// SetCryptoEngine assigns MacsecCryptoEngine provided by user to Macsec.
	// MacsecCryptoEngine is a container of crypto engine properties of a SecY.
	SetCryptoEngine(value MacsecCryptoEngine) Macsec
	// HasCryptoEngine checks if CryptoEngine has been set in Macsec
	HasCryptoEngine() bool
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

// Static key properties properties of SecY. Static key is used in absence MKA.
// StaticKey returns a MacsecStaticKey
func (obj *macsec) StaticKey() MacsecStaticKey {
	if obj.obj.StaticKey == nil {
		obj.obj.StaticKey = NewMacsecStaticKey().msg()
	}
	if obj.staticKeyHolder == nil {
		obj.staticKeyHolder = &macsecStaticKey{obj: obj.obj.StaticKey}
	}
	return obj.staticKeyHolder
}

// Static key properties properties of SecY. Static key is used in absence MKA.
// StaticKey returns a MacsecStaticKey
func (obj *macsec) HasStaticKey() bool {
	return obj.obj.StaticKey != nil
}

// Static key properties properties of SecY. Static key is used in absence MKA.
// SetStaticKey sets the MacsecStaticKey value in the Macsec object
func (obj *macsec) SetStaticKey(value MacsecStaticKey) Macsec {

	obj.staticKeyHolder = nil
	obj.obj.StaticKey = value.msg()

	return obj
}

// Tx properties of SecY.
// Tx returns a MacsecTx
func (obj *macsec) Tx() MacsecTx {
	if obj.obj.Tx == nil {
		obj.obj.Tx = NewMacsecTx().msg()
	}
	if obj.txHolder == nil {
		obj.txHolder = &macsecTx{obj: obj.obj.Tx}
	}
	return obj.txHolder
}

// Tx properties of SecY.
// Tx returns a MacsecTx
func (obj *macsec) HasTx() bool {
	return obj.obj.Tx != nil
}

// Tx properties of SecY.
// SetTx sets the MacsecTx value in the Macsec object
func (obj *macsec) SetTx(value MacsecTx) Macsec {

	obj.txHolder = nil
	obj.obj.Tx = value.msg()

	return obj
}

// Rx properties of SecY.
// Rx returns a MacsecRx
func (obj *macsec) Rx() MacsecRx {
	if obj.obj.Rx == nil {
		obj.obj.Rx = NewMacsecRx().msg()
	}
	if obj.rxHolder == nil {
		obj.rxHolder = &macsecRx{obj: obj.obj.Rx}
	}
	return obj.rxHolder
}

// Rx properties of SecY.
// Rx returns a MacsecRx
func (obj *macsec) HasRx() bool {
	return obj.obj.Rx != nil
}

// Rx properties of SecY.
// SetRx sets the MacsecRx value in the Macsec object
func (obj *macsec) SetRx(value MacsecRx) Macsec {

	obj.rxHolder = nil
	obj.obj.Rx = value.msg()

	return obj
}

// Crypto engine properties of SecY.
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

// Crypto engine properties of SecY.
// CryptoEngine returns a MacsecCryptoEngine
func (obj *macsec) HasCryptoEngine() bool {
	return obj.obj.CryptoEngine != nil
}

// Crypto engine properties of SecY.
// SetCryptoEngine sets the MacsecCryptoEngine value in the Macsec object
func (obj *macsec) SetCryptoEngine(value MacsecCryptoEngine) Macsec {

	obj.cryptoEngineHolder = nil
	obj.obj.CryptoEngine = value.msg()

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

	if obj.obj.StaticKey != nil {

		obj.StaticKey().validateObj(vObj, set_default)
	}

	if obj.obj.Tx != nil {

		obj.Tx().validateObj(vObj, set_default)
	}

	if obj.obj.Rx != nil {

		obj.Rx().validateObj(vObj, set_default)
	}

	if obj.obj.CryptoEngine != nil {

		obj.CryptoEngine().validateObj(vObj, set_default)
	}

}

func (obj *macsec) setDefault() {

}
