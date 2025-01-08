package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacsecEthernetInterface *****
type macsecEthernetInterface struct {
	validation
	obj                *otg.MacsecEthernetInterface
	marshaller         marshalMacsecEthernetInterface
	unMarshaller       unMarshalMacsecEthernetInterface
	basicHolder        MacsecBasic
	txHolder           MacsecTx
	rxHolder           MacsecRx
	cryptoEngineHolder MacsecCryptoEngine
	advanceHolder      MacsecAdvance
}

func NewMacsecEthernetInterface() MacsecEthernetInterface {
	obj := macsecEthernetInterface{obj: &otg.MacsecEthernetInterface{}}
	obj.setDefault()
	return &obj
}

func (obj *macsecEthernetInterface) msg() *otg.MacsecEthernetInterface {
	return obj.obj
}

func (obj *macsecEthernetInterface) setMsg(msg *otg.MacsecEthernetInterface) MacsecEthernetInterface {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacsecEthernetInterface struct {
	obj *macsecEthernetInterface
}

type marshalMacsecEthernetInterface interface {
	// ToProto marshals MacsecEthernetInterface to protobuf object *otg.MacsecEthernetInterface
	ToProto() (*otg.MacsecEthernetInterface, error)
	// ToPbText marshals MacsecEthernetInterface to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacsecEthernetInterface to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacsecEthernetInterface to JSON text
	ToJson() (string, error)
}

type unMarshalmacsecEthernetInterface struct {
	obj *macsecEthernetInterface
}

type unMarshalMacsecEthernetInterface interface {
	// FromProto unmarshals MacsecEthernetInterface from protobuf object *otg.MacsecEthernetInterface
	FromProto(msg *otg.MacsecEthernetInterface) (MacsecEthernetInterface, error)
	// FromPbText unmarshals MacsecEthernetInterface from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacsecEthernetInterface from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacsecEthernetInterface from JSON text
	FromJson(value string) error
}

func (obj *macsecEthernetInterface) Marshal() marshalMacsecEthernetInterface {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacsecEthernetInterface{obj: obj}
	}
	return obj.marshaller
}

func (obj *macsecEthernetInterface) Unmarshal() unMarshalMacsecEthernetInterface {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacsecEthernetInterface{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacsecEthernetInterface) ToProto() (*otg.MacsecEthernetInterface, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacsecEthernetInterface) FromProto(msg *otg.MacsecEthernetInterface) (MacsecEthernetInterface, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacsecEthernetInterface) ToPbText() (string, error) {
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

func (m *unMarshalmacsecEthernetInterface) FromPbText(value string) error {
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

func (m *marshalmacsecEthernetInterface) ToYaml() (string, error) {
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

func (m *unMarshalmacsecEthernetInterface) FromYaml(value string) error {
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

func (m *marshalmacsecEthernetInterface) ToJson() (string, error) {
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

func (m *unMarshalmacsecEthernetInterface) FromJson(value string) error {
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

func (obj *macsecEthernetInterface) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macsecEthernetInterface) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macsecEthernetInterface) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macsecEthernetInterface) Clone() (MacsecEthernetInterface, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacsecEthernetInterface()
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

func (obj *macsecEthernetInterface) setNil() {
	obj.basicHolder = nil
	obj.txHolder = nil
	obj.rxHolder = nil
	obj.cryptoEngineHolder = nil
	obj.advanceHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacsecEthernetInterface is configuration for single MACsec interface.
type MacsecEthernetInterface interface {
	Validation
	// msg marshals MacsecEthernetInterface to protobuf object *otg.MacsecEthernetInterface
	// and doesn't set defaults
	msg() *otg.MacsecEthernetInterface
	// setMsg unmarshals MacsecEthernetInterface from protobuf object *otg.MacsecEthernetInterface
	// and doesn't set defaults
	setMsg(*otg.MacsecEthernetInterface) MacsecEthernetInterface
	// provides marshal interface
	Marshal() marshalMacsecEthernetInterface
	// provides unmarshal interface
	Unmarshal() unMarshalMacsecEthernetInterface
	// validate validates MacsecEthernetInterface
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacsecEthernetInterface, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EthName returns string, set in MacsecEthernetInterface.
	EthName() string
	// SetEthName assigns string provided by user to MacsecEthernetInterface
	SetEthName(value string) MacsecEthernetInterface
	// Name returns string, set in MacsecEthernetInterface.
	Name() string
	// SetName assigns string provided by user to MacsecEthernetInterface
	SetName(value string) MacsecEthernetInterface
	// Basic returns MacsecBasic, set in MacsecEthernetInterface.
	// MacsecBasic is a container of basic properties for a MACsec interface.
	Basic() MacsecBasic
	// SetBasic assigns MacsecBasic provided by user to MacsecEthernetInterface.
	// MacsecBasic is a container of basic properties for a MACsec interface.
	SetBasic(value MacsecBasic) MacsecEthernetInterface
	// Tx returns MacsecTx, set in MacsecEthernetInterface.
	// MacsecTx is the container for Tx settings.
	Tx() MacsecTx
	// SetTx assigns MacsecTx provided by user to MacsecEthernetInterface.
	// MacsecTx is the container for Tx settings.
	SetTx(value MacsecTx) MacsecEthernetInterface
	// Rx returns MacsecRx, set in MacsecEthernetInterface.
	// MacsecRx is the container for Rx settings.
	Rx() MacsecRx
	// SetRx assigns MacsecRx provided by user to MacsecEthernetInterface.
	// MacsecRx is the container for Rx settings.
	SetRx(value MacsecRx) MacsecEthernetInterface
	// CryptoEngine returns MacsecCryptoEngine, set in MacsecEthernetInterface.
	// MacsecCryptoEngine is a container of crypto engine properties for a MACsec interface.
	CryptoEngine() MacsecCryptoEngine
	// SetCryptoEngine assigns MacsecCryptoEngine provided by user to MacsecEthernetInterface.
	// MacsecCryptoEngine is a container of crypto engine properties for a MACsec interface.
	SetCryptoEngine(value MacsecCryptoEngine) MacsecEthernetInterface
	// Advance returns MacsecAdvance, set in MacsecEthernetInterface.
	// MacsecAdvance is a container of advance properties for a MACsec interface.
	Advance() MacsecAdvance
	// SetAdvance assigns MacsecAdvance provided by user to MacsecEthernetInterface.
	// MacsecAdvance is a container of advance properties for a MACsec interface.
	SetAdvance(value MacsecAdvance) MacsecEthernetInterface
	// HasAdvance checks if Advance has been set in MacsecEthernetInterface
	HasAdvance() bool
	setNil()
}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// EthName returns a string
func (obj *macsecEthernetInterface) EthName() string {

	return *obj.obj.EthName

}

// The unique name of the Ethernet interface on which MACsec is enabled.
//
// x-constraint:
// - /components/schemas/Device.Ethernet/properties/name
//
// SetEthName sets the string value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetEthName(value string) MacsecEthernetInterface {

	obj.obj.EthName = &value
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *macsecEthernetInterface) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetName(value string) MacsecEthernetInterface {

	obj.obj.Name = &value
	return obj
}

// This contains the basic properties of MACsec.
// Basic returns a MacsecBasic
func (obj *macsecEthernetInterface) Basic() MacsecBasic {
	if obj.obj.Basic == nil {
		obj.obj.Basic = NewMacsecBasic().msg()
	}
	if obj.basicHolder == nil {
		obj.basicHolder = &macsecBasic{obj: obj.obj.Basic}
	}
	return obj.basicHolder
}

// This contains the basic properties of MACsec.
// SetBasic sets the MacsecBasic value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetBasic(value MacsecBasic) MacsecEthernetInterface {

	obj.basicHolder = nil
	obj.obj.Basic = value.msg()

	return obj
}

// This contains the Tx properties of MACsec.
// Tx returns a MacsecTx
func (obj *macsecEthernetInterface) Tx() MacsecTx {
	if obj.obj.Tx == nil {
		obj.obj.Tx = NewMacsecTx().msg()
	}
	if obj.txHolder == nil {
		obj.txHolder = &macsecTx{obj: obj.obj.Tx}
	}
	return obj.txHolder
}

// This contains the Tx properties of MACsec.
// SetTx sets the MacsecTx value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetTx(value MacsecTx) MacsecEthernetInterface {

	obj.txHolder = nil
	obj.obj.Tx = value.msg()

	return obj
}

// This contains the Rx properties of MACsec.
// Rx returns a MacsecRx
func (obj *macsecEthernetInterface) Rx() MacsecRx {
	if obj.obj.Rx == nil {
		obj.obj.Rx = NewMacsecRx().msg()
	}
	if obj.rxHolder == nil {
		obj.rxHolder = &macsecRx{obj: obj.obj.Rx}
	}
	return obj.rxHolder
}

// This contains the Rx properties of MACsec.
// SetRx sets the MacsecRx value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetRx(value MacsecRx) MacsecEthernetInterface {

	obj.rxHolder = nil
	obj.obj.Rx = value.msg()

	return obj
}

// This contains crypto engine properties of MACsec.
// CryptoEngine returns a MacsecCryptoEngine
func (obj *macsecEthernetInterface) CryptoEngine() MacsecCryptoEngine {
	if obj.obj.CryptoEngine == nil {
		obj.obj.CryptoEngine = NewMacsecCryptoEngine().msg()
	}
	if obj.cryptoEngineHolder == nil {
		obj.cryptoEngineHolder = &macsecCryptoEngine{obj: obj.obj.CryptoEngine}
	}
	return obj.cryptoEngineHolder
}

// This contains crypto engine properties of MACsec.
// SetCryptoEngine sets the MacsecCryptoEngine value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetCryptoEngine(value MacsecCryptoEngine) MacsecEthernetInterface {

	obj.cryptoEngineHolder = nil
	obj.obj.CryptoEngine = value.msg()

	return obj
}

// This contains advance properties of MACsec.
// Advance returns a MacsecAdvance
func (obj *macsecEthernetInterface) Advance() MacsecAdvance {
	if obj.obj.Advance == nil {
		obj.obj.Advance = NewMacsecAdvance().msg()
	}
	if obj.advanceHolder == nil {
		obj.advanceHolder = &macsecAdvance{obj: obj.obj.Advance}
	}
	return obj.advanceHolder
}

// This contains advance properties of MACsec.
// Advance returns a MacsecAdvance
func (obj *macsecEthernetInterface) HasAdvance() bool {
	return obj.obj.Advance != nil
}

// This contains advance properties of MACsec.
// SetAdvance sets the MacsecAdvance value in the MacsecEthernetInterface object
func (obj *macsecEthernetInterface) SetAdvance(value MacsecAdvance) MacsecEthernetInterface {

	obj.advanceHolder = nil
	obj.obj.Advance = value.msg()

	return obj
}

func (obj *macsecEthernetInterface) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EthName is required
	if obj.obj.EthName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EthName is required field on interface MacsecEthernetInterface")
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface MacsecEthernetInterface")
	}

	// Basic is required
	if obj.obj.Basic == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Basic is required field on interface MacsecEthernetInterface")
	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(vObj, set_default)
	}

	// Tx is required
	if obj.obj.Tx == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Tx is required field on interface MacsecEthernetInterface")
	}

	if obj.obj.Tx != nil {

		obj.Tx().validateObj(vObj, set_default)
	}

	// Rx is required
	if obj.obj.Rx == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Rx is required field on interface MacsecEthernetInterface")
	}

	if obj.obj.Rx != nil {

		obj.Rx().validateObj(vObj, set_default)
	}

	// CryptoEngine is required
	if obj.obj.CryptoEngine == nil {
		vObj.validationErrors = append(vObj.validationErrors, "CryptoEngine is required field on interface MacsecEthernetInterface")
	}

	if obj.obj.CryptoEngine != nil {

		obj.CryptoEngine().validateObj(vObj, set_default)
	}

	if obj.obj.Advance != nil {

		obj.Advance().validateObj(vObj, set_default)
	}

}

func (obj *macsecEthernetInterface) setDefault() {

}
