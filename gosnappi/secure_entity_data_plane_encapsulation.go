package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SecureEntityDataPlaneEncapsulation *****
type secureEntityDataPlaneEncapsulation struct {
	validation
	obj                *otg.SecureEntityDataPlaneEncapsulation
	marshaller         marshalSecureEntityDataPlaneEncapsulation
	unMarshaller       unMarshalSecureEntityDataPlaneEncapsulation
	txHolder           SecureEntityDataPlaneTx
	rxHolder           SecureEntityDataPlaneRx
	cryptoEngineHolder SecureEntityCryptoEngine
}

func NewSecureEntityDataPlaneEncapsulation() SecureEntityDataPlaneEncapsulation {
	obj := secureEntityDataPlaneEncapsulation{obj: &otg.SecureEntityDataPlaneEncapsulation{}}
	obj.setDefault()
	return &obj
}

func (obj *secureEntityDataPlaneEncapsulation) msg() *otg.SecureEntityDataPlaneEncapsulation {
	return obj.obj
}

func (obj *secureEntityDataPlaneEncapsulation) setMsg(msg *otg.SecureEntityDataPlaneEncapsulation) SecureEntityDataPlaneEncapsulation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsecureEntityDataPlaneEncapsulation struct {
	obj *secureEntityDataPlaneEncapsulation
}

type marshalSecureEntityDataPlaneEncapsulation interface {
	// ToProto marshals SecureEntityDataPlaneEncapsulation to protobuf object *otg.SecureEntityDataPlaneEncapsulation
	ToProto() (*otg.SecureEntityDataPlaneEncapsulation, error)
	// ToPbText marshals SecureEntityDataPlaneEncapsulation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SecureEntityDataPlaneEncapsulation to YAML text
	ToYaml() (string, error)
	// ToJson marshals SecureEntityDataPlaneEncapsulation to JSON text
	ToJson() (string, error)
}

type unMarshalsecureEntityDataPlaneEncapsulation struct {
	obj *secureEntityDataPlaneEncapsulation
}

type unMarshalSecureEntityDataPlaneEncapsulation interface {
	// FromProto unmarshals SecureEntityDataPlaneEncapsulation from protobuf object *otg.SecureEntityDataPlaneEncapsulation
	FromProto(msg *otg.SecureEntityDataPlaneEncapsulation) (SecureEntityDataPlaneEncapsulation, error)
	// FromPbText unmarshals SecureEntityDataPlaneEncapsulation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SecureEntityDataPlaneEncapsulation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SecureEntityDataPlaneEncapsulation from JSON text
	FromJson(value string) error
}

func (obj *secureEntityDataPlaneEncapsulation) Marshal() marshalSecureEntityDataPlaneEncapsulation {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsecureEntityDataPlaneEncapsulation{obj: obj}
	}
	return obj.marshaller
}

func (obj *secureEntityDataPlaneEncapsulation) Unmarshal() unMarshalSecureEntityDataPlaneEncapsulation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsecureEntityDataPlaneEncapsulation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsecureEntityDataPlaneEncapsulation) ToProto() (*otg.SecureEntityDataPlaneEncapsulation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsecureEntityDataPlaneEncapsulation) FromProto(msg *otg.SecureEntityDataPlaneEncapsulation) (SecureEntityDataPlaneEncapsulation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsecureEntityDataPlaneEncapsulation) ToPbText() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneEncapsulation) FromPbText(value string) error {
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

func (m *marshalsecureEntityDataPlaneEncapsulation) ToYaml() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneEncapsulation) FromYaml(value string) error {
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

func (m *marshalsecureEntityDataPlaneEncapsulation) ToJson() (string, error) {
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

func (m *unMarshalsecureEntityDataPlaneEncapsulation) FromJson(value string) error {
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

func (obj *secureEntityDataPlaneEncapsulation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *secureEntityDataPlaneEncapsulation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *secureEntityDataPlaneEncapsulation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *secureEntityDataPlaneEncapsulation) Clone() (SecureEntityDataPlaneEncapsulation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSecureEntityDataPlaneEncapsulation()
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

func (obj *secureEntityDataPlaneEncapsulation) setNil() {
	obj.txHolder = nil
	obj.rxHolder = nil
	obj.cryptoEngineHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SecureEntityDataPlaneEncapsulation is a container of encapsulation properties for a secure entity(SecY).
type SecureEntityDataPlaneEncapsulation interface {
	Validation
	// msg marshals SecureEntityDataPlaneEncapsulation to protobuf object *otg.SecureEntityDataPlaneEncapsulation
	// and doesn't set defaults
	msg() *otg.SecureEntityDataPlaneEncapsulation
	// setMsg unmarshals SecureEntityDataPlaneEncapsulation from protobuf object *otg.SecureEntityDataPlaneEncapsulation
	// and doesn't set defaults
	setMsg(*otg.SecureEntityDataPlaneEncapsulation) SecureEntityDataPlaneEncapsulation
	// provides marshal interface
	Marshal() marshalSecureEntityDataPlaneEncapsulation
	// provides unmarshal interface
	Unmarshal() unMarshalSecureEntityDataPlaneEncapsulation
	// validate validates SecureEntityDataPlaneEncapsulation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SecureEntityDataPlaneEncapsulation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Tx returns SecureEntityDataPlaneTx, set in SecureEntityDataPlaneEncapsulation.
	// SecureEntityDataPlaneTx is a container of Tx properties of SecY.
	Tx() SecureEntityDataPlaneTx
	// SetTx assigns SecureEntityDataPlaneTx provided by user to SecureEntityDataPlaneEncapsulation.
	// SecureEntityDataPlaneTx is a container of Tx properties of SecY.
	SetTx(value SecureEntityDataPlaneTx) SecureEntityDataPlaneEncapsulation
	// HasTx checks if Tx has been set in SecureEntityDataPlaneEncapsulation
	HasTx() bool
	// Rx returns SecureEntityDataPlaneRx, set in SecureEntityDataPlaneEncapsulation.
	// SecureEntityDataPlaneRx is a container for Rx settings of SecY.
	Rx() SecureEntityDataPlaneRx
	// SetRx assigns SecureEntityDataPlaneRx provided by user to SecureEntityDataPlaneEncapsulation.
	// SecureEntityDataPlaneRx is a container for Rx settings of SecY.
	SetRx(value SecureEntityDataPlaneRx) SecureEntityDataPlaneEncapsulation
	// HasRx checks if Rx has been set in SecureEntityDataPlaneEncapsulation
	HasRx() bool
	// CryptoEngine returns SecureEntityCryptoEngine, set in SecureEntityDataPlaneEncapsulation.
	// SecureEntityCryptoEngine is a container of crypto engine properties of a SecY.
	CryptoEngine() SecureEntityCryptoEngine
	// SetCryptoEngine assigns SecureEntityCryptoEngine provided by user to SecureEntityDataPlaneEncapsulation.
	// SecureEntityCryptoEngine is a container of crypto engine properties of a SecY.
	SetCryptoEngine(value SecureEntityCryptoEngine) SecureEntityDataPlaneEncapsulation
	setNil()
}

// Tx properties of SecY.
// Tx returns a SecureEntityDataPlaneTx
func (obj *secureEntityDataPlaneEncapsulation) Tx() SecureEntityDataPlaneTx {
	if obj.obj.Tx == nil {
		obj.obj.Tx = NewSecureEntityDataPlaneTx().msg()
	}
	if obj.txHolder == nil {
		obj.txHolder = &secureEntityDataPlaneTx{obj: obj.obj.Tx}
	}
	return obj.txHolder
}

// Tx properties of SecY.
// Tx returns a SecureEntityDataPlaneTx
func (obj *secureEntityDataPlaneEncapsulation) HasTx() bool {
	return obj.obj.Tx != nil
}

// Tx properties of SecY.
// SetTx sets the SecureEntityDataPlaneTx value in the SecureEntityDataPlaneEncapsulation object
func (obj *secureEntityDataPlaneEncapsulation) SetTx(value SecureEntityDataPlaneTx) SecureEntityDataPlaneEncapsulation {

	obj.txHolder = nil
	obj.obj.Tx = value.msg()

	return obj
}

// Rx properties of SecY.
// Rx returns a SecureEntityDataPlaneRx
func (obj *secureEntityDataPlaneEncapsulation) Rx() SecureEntityDataPlaneRx {
	if obj.obj.Rx == nil {
		obj.obj.Rx = NewSecureEntityDataPlaneRx().msg()
	}
	if obj.rxHolder == nil {
		obj.rxHolder = &secureEntityDataPlaneRx{obj: obj.obj.Rx}
	}
	return obj.rxHolder
}

// Rx properties of SecY.
// Rx returns a SecureEntityDataPlaneRx
func (obj *secureEntityDataPlaneEncapsulation) HasRx() bool {
	return obj.obj.Rx != nil
}

// Rx properties of SecY.
// SetRx sets the SecureEntityDataPlaneRx value in the SecureEntityDataPlaneEncapsulation object
func (obj *secureEntityDataPlaneEncapsulation) SetRx(value SecureEntityDataPlaneRx) SecureEntityDataPlaneEncapsulation {

	obj.rxHolder = nil
	obj.obj.Rx = value.msg()

	return obj
}

// Crypto engine properties of SecY.
// CryptoEngine returns a SecureEntityCryptoEngine
func (obj *secureEntityDataPlaneEncapsulation) CryptoEngine() SecureEntityCryptoEngine {
	if obj.obj.CryptoEngine == nil {
		obj.obj.CryptoEngine = NewSecureEntityCryptoEngine().msg()
	}
	if obj.cryptoEngineHolder == nil {
		obj.cryptoEngineHolder = &secureEntityCryptoEngine{obj: obj.obj.CryptoEngine}
	}
	return obj.cryptoEngineHolder
}

// Crypto engine properties of SecY.
// SetCryptoEngine sets the SecureEntityCryptoEngine value in the SecureEntityDataPlaneEncapsulation object
func (obj *secureEntityDataPlaneEncapsulation) SetCryptoEngine(value SecureEntityCryptoEngine) SecureEntityDataPlaneEncapsulation {

	obj.cryptoEngineHolder = nil
	obj.obj.CryptoEngine = value.msg()

	return obj
}

func (obj *secureEntityDataPlaneEncapsulation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Tx != nil {

		obj.Tx().validateObj(vObj, set_default)
	}

	if obj.obj.Rx != nil {

		obj.Rx().validateObj(vObj, set_default)
	}

	// CryptoEngine is required
	if obj.obj.CryptoEngine == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CryptoEngine is required field on interface SecureEntityDataPlaneEncapsulation")
	}

	if obj.obj.CryptoEngine != nil {

		obj.CryptoEngine().validateObj(vObj, set_default)
	}

}

func (obj *secureEntityDataPlaneEncapsulation) setDefault() {

}
