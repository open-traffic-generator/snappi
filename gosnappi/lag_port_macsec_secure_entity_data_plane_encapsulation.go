package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityDataPlaneEncapsulation *****
type lagPortMacsecSecureEntityDataPlaneEncapsulation struct {
	validation
	obj                *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation
	marshaller         marshalLagPortMacsecSecureEntityDataPlaneEncapsulation
	unMarshaller       unMarshalLagPortMacsecSecureEntityDataPlaneEncapsulation
	txHolder           LagPortMacsecSecureEntityDataPlaneTx
	rxHolder           LagPortMacsecSecureEntityDataPlaneRx
	cryptoEngineHolder LagPortMacsecSecureEntityCryptoEngine
}

func NewLagPortMacsecSecureEntityDataPlaneEncapsulation() LagPortMacsecSecureEntityDataPlaneEncapsulation {
	obj := lagPortMacsecSecureEntityDataPlaneEncapsulation{obj: &otg.LagPortMacsecSecureEntityDataPlaneEncapsulation{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) msg() *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) setMsg(msg *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation) LagPortMacsecSecureEntityDataPlaneEncapsulation {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityDataPlaneEncapsulation struct {
	obj *lagPortMacsecSecureEntityDataPlaneEncapsulation
}

type marshalLagPortMacsecSecureEntityDataPlaneEncapsulation interface {
	// ToProto marshals LagPortMacsecSecureEntityDataPlaneEncapsulation to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation
	ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneEncapsulation, error)
	// ToPbText marshals LagPortMacsecSecureEntityDataPlaneEncapsulation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityDataPlaneEncapsulation to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityDataPlaneEncapsulation to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityDataPlaneEncapsulation struct {
	obj *lagPortMacsecSecureEntityDataPlaneEncapsulation
}

type unMarshalLagPortMacsecSecureEntityDataPlaneEncapsulation interface {
	// FromProto unmarshals LagPortMacsecSecureEntityDataPlaneEncapsulation from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation
	FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation) (LagPortMacsecSecureEntityDataPlaneEncapsulation, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityDataPlaneEncapsulation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityDataPlaneEncapsulation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityDataPlaneEncapsulation from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) Marshal() marshalLagPortMacsecSecureEntityDataPlaneEncapsulation {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityDataPlaneEncapsulation{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneEncapsulation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityDataPlaneEncapsulation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityDataPlaneEncapsulation) ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneEncapsulation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityDataPlaneEncapsulation) FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation) (LagPortMacsecSecureEntityDataPlaneEncapsulation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityDataPlaneEncapsulation) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneEncapsulation) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneEncapsulation) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneEncapsulation) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneEncapsulation) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneEncapsulation) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) Clone() (LagPortMacsecSecureEntityDataPlaneEncapsulation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityDataPlaneEncapsulation()
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

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) setNil() {
	obj.txHolder = nil
	obj.rxHolder = nil
	obj.cryptoEngineHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LagPortMacsecSecureEntityDataPlaneEncapsulation is a container of encapsulation properties for a secure entity(SecY).
type LagPortMacsecSecureEntityDataPlaneEncapsulation interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityDataPlaneEncapsulation to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation
	// setMsg unmarshals LagPortMacsecSecureEntityDataPlaneEncapsulation from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneEncapsulation
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityDataPlaneEncapsulation) LagPortMacsecSecureEntityDataPlaneEncapsulation
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityDataPlaneEncapsulation
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneEncapsulation
	// validate validates LagPortMacsecSecureEntityDataPlaneEncapsulation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityDataPlaneEncapsulation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Tx returns LagPortMacsecSecureEntityDataPlaneTx, set in LagPortMacsecSecureEntityDataPlaneEncapsulation.
	// LagPortMacsecSecureEntityDataPlaneTx is a container of Tx properties of SecY.
	Tx() LagPortMacsecSecureEntityDataPlaneTx
	// SetTx assigns LagPortMacsecSecureEntityDataPlaneTx provided by user to LagPortMacsecSecureEntityDataPlaneEncapsulation.
	// LagPortMacsecSecureEntityDataPlaneTx is a container of Tx properties of SecY.
	SetTx(value LagPortMacsecSecureEntityDataPlaneTx) LagPortMacsecSecureEntityDataPlaneEncapsulation
	// HasTx checks if Tx has been set in LagPortMacsecSecureEntityDataPlaneEncapsulation
	HasTx() bool
	// Rx returns LagPortMacsecSecureEntityDataPlaneRx, set in LagPortMacsecSecureEntityDataPlaneEncapsulation.
	// LagPortMacsecSecureEntityDataPlaneRx is a container for Rx settings of SecY.
	Rx() LagPortMacsecSecureEntityDataPlaneRx
	// SetRx assigns LagPortMacsecSecureEntityDataPlaneRx provided by user to LagPortMacsecSecureEntityDataPlaneEncapsulation.
	// LagPortMacsecSecureEntityDataPlaneRx is a container for Rx settings of SecY.
	SetRx(value LagPortMacsecSecureEntityDataPlaneRx) LagPortMacsecSecureEntityDataPlaneEncapsulation
	// HasRx checks if Rx has been set in LagPortMacsecSecureEntityDataPlaneEncapsulation
	HasRx() bool
	// CryptoEngine returns LagPortMacsecSecureEntityCryptoEngine, set in LagPortMacsecSecureEntityDataPlaneEncapsulation.
	// LagPortMacsecSecureEntityCryptoEngine is a container of crypto engine properties of a SecY.
	CryptoEngine() LagPortMacsecSecureEntityCryptoEngine
	// SetCryptoEngine assigns LagPortMacsecSecureEntityCryptoEngine provided by user to LagPortMacsecSecureEntityDataPlaneEncapsulation.
	// LagPortMacsecSecureEntityCryptoEngine is a container of crypto engine properties of a SecY.
	SetCryptoEngine(value LagPortMacsecSecureEntityCryptoEngine) LagPortMacsecSecureEntityDataPlaneEncapsulation
	setNil()
}

// Tx properties of SecY.
// Tx returns a LagPortMacsecSecureEntityDataPlaneTx
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) Tx() LagPortMacsecSecureEntityDataPlaneTx {
	if obj.obj.Tx == nil {
		obj.obj.Tx = NewLagPortMacsecSecureEntityDataPlaneTx().msg()
	}
	if obj.txHolder == nil {
		obj.txHolder = &lagPortMacsecSecureEntityDataPlaneTx{obj: obj.obj.Tx}
	}
	return obj.txHolder
}

// Tx properties of SecY.
// Tx returns a LagPortMacsecSecureEntityDataPlaneTx
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) HasTx() bool {
	return obj.obj.Tx != nil
}

// Tx properties of SecY.
// SetTx sets the LagPortMacsecSecureEntityDataPlaneTx value in the LagPortMacsecSecureEntityDataPlaneEncapsulation object
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) SetTx(value LagPortMacsecSecureEntityDataPlaneTx) LagPortMacsecSecureEntityDataPlaneEncapsulation {

	obj.txHolder = nil
	obj.obj.Tx = value.msg()

	return obj
}

// Rx properties of SecY.
// Rx returns a LagPortMacsecSecureEntityDataPlaneRx
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) Rx() LagPortMacsecSecureEntityDataPlaneRx {
	if obj.obj.Rx == nil {
		obj.obj.Rx = NewLagPortMacsecSecureEntityDataPlaneRx().msg()
	}
	if obj.rxHolder == nil {
		obj.rxHolder = &lagPortMacsecSecureEntityDataPlaneRx{obj: obj.obj.Rx}
	}
	return obj.rxHolder
}

// Rx properties of SecY.
// Rx returns a LagPortMacsecSecureEntityDataPlaneRx
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) HasRx() bool {
	return obj.obj.Rx != nil
}

// Rx properties of SecY.
// SetRx sets the LagPortMacsecSecureEntityDataPlaneRx value in the LagPortMacsecSecureEntityDataPlaneEncapsulation object
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) SetRx(value LagPortMacsecSecureEntityDataPlaneRx) LagPortMacsecSecureEntityDataPlaneEncapsulation {

	obj.rxHolder = nil
	obj.obj.Rx = value.msg()

	return obj
}

// Crypto engine properties of SecY.
// CryptoEngine returns a LagPortMacsecSecureEntityCryptoEngine
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) CryptoEngine() LagPortMacsecSecureEntityCryptoEngine {
	if obj.obj.CryptoEngine == nil {
		obj.obj.CryptoEngine = NewLagPortMacsecSecureEntityCryptoEngine().msg()
	}
	if obj.cryptoEngineHolder == nil {
		obj.cryptoEngineHolder = &lagPortMacsecSecureEntityCryptoEngine{obj: obj.obj.CryptoEngine}
	}
	return obj.cryptoEngineHolder
}

// Crypto engine properties of SecY.
// SetCryptoEngine sets the LagPortMacsecSecureEntityCryptoEngine value in the LagPortMacsecSecureEntityDataPlaneEncapsulation object
func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) SetCryptoEngine(value LagPortMacsecSecureEntityCryptoEngine) LagPortMacsecSecureEntityDataPlaneEncapsulation {

	obj.cryptoEngineHolder = nil
	obj.obj.CryptoEngine = value.msg()

	return obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) validateObj(vObj *validation, set_default bool) {
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
		vObj.validationErrors = append(vObj.validationErrors, "CryptoEngine is required field on interface LagPortMacsecSecureEntityDataPlaneEncapsulation")
	}

	if obj.obj.CryptoEngine != nil {

		obj.CryptoEngine().validateObj(vObj, set_default)
	}

}

func (obj *lagPortMacsecSecureEntityDataPlaneEncapsulation) setDefault() {

}
