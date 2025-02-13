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
	obj           *otg.MacsecRxSc
	marshaller    marshalMacsecRxSc
	unMarshaller  unMarshalMacsecRxSc
	sakPoolHolder MacsecStaticKeySakPool
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
	obj.sakPoolHolder = nil
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
	// SakPool returns MacsecStaticKeySakPool, set in MacsecRxSc.
	// MacsecStaticKeySakPool is the container for Secure Association Key(SAK) Pool.
	SakPool() MacsecStaticKeySakPool
	// SetSakPool assigns MacsecStaticKeySakPool provided by user to MacsecRxSc.
	// MacsecStaticKeySakPool is the container for Secure Association Key(SAK) Pool.
	SetSakPool(value MacsecStaticKeySakPool) MacsecRxSc
	// HasSakPool checks if SakPool has been set in MacsecRxSc
	HasSakPool() bool
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
// SakPool returns a MacsecStaticKeySakPool
func (obj *macsecRxSc) SakPool() MacsecStaticKeySakPool {
	if obj.obj.SakPool == nil {
		obj.obj.SakPool = NewMacsecStaticKeySakPool().msg()
	}
	if obj.sakPoolHolder == nil {
		obj.sakPoolHolder = &macsecStaticKeySakPool{obj: obj.obj.SakPool}
	}
	return obj.sakPoolHolder
}

// Rx SAK pool.
// SakPool returns a MacsecStaticKeySakPool
func (obj *macsecRxSc) HasSakPool() bool {
	return obj.obj.SakPool != nil
}

// Rx SAK pool.
// SetSakPool sets the MacsecStaticKeySakPool value in the MacsecRxSc object
func (obj *macsecRxSc) SetSakPool(value MacsecStaticKeySakPool) MacsecRxSc {

	obj.sakPoolHolder = nil
	obj.obj.SakPool = value.msg()

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

	if obj.obj.SakPool != nil {

		obj.SakPool().validateObj(vObj, set_default)
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
