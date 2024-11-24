package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureEthernet *****
type captureEthernet struct {
	validation
	obj             *otg.CaptureEthernet
	marshaller      marshalCaptureEthernet
	unMarshaller    unMarshalCaptureEthernet
	srcHolder       CaptureField
	dstHolder       CaptureField
	etherTypeHolder CaptureField
	pfcQueueHolder  CaptureField
}

func NewCaptureEthernet() CaptureEthernet {
	obj := captureEthernet{obj: &otg.CaptureEthernet{}}
	obj.setDefault()
	return &obj
}

func (obj *captureEthernet) msg() *otg.CaptureEthernet {
	return obj.obj
}

func (obj *captureEthernet) setMsg(msg *otg.CaptureEthernet) CaptureEthernet {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureEthernet struct {
	obj *captureEthernet
}

type marshalCaptureEthernet interface {
	// ToProto marshals CaptureEthernet to protobuf object *otg.CaptureEthernet
	ToProto() (*otg.CaptureEthernet, error)
	// ToPbText marshals CaptureEthernet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureEthernet to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureEthernet to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals CaptureEthernet to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalcaptureEthernet struct {
	obj *captureEthernet
}

type unMarshalCaptureEthernet interface {
	// FromProto unmarshals CaptureEthernet from protobuf object *otg.CaptureEthernet
	FromProto(msg *otg.CaptureEthernet) (CaptureEthernet, error)
	// FromPbText unmarshals CaptureEthernet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureEthernet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureEthernet from JSON text
	FromJson(value string) error
}

func (obj *captureEthernet) Marshal() marshalCaptureEthernet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureEthernet{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureEthernet) Unmarshal() unMarshalCaptureEthernet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureEthernet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureEthernet) ToProto() (*otg.CaptureEthernet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureEthernet) FromProto(msg *otg.CaptureEthernet) (CaptureEthernet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureEthernet) ToPbText() (string, error) {
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

func (m *unMarshalcaptureEthernet) FromPbText(value string) error {
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

func (m *marshalcaptureEthernet) ToYaml() (string, error) {
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

func (m *unMarshalcaptureEthernet) FromYaml(value string) error {
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

func (m *marshalcaptureEthernet) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalcaptureEthernet) ToJson() (string, error) {
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

func (m *unMarshalcaptureEthernet) FromJson(value string) error {
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

func (obj *captureEthernet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureEthernet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureEthernet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureEthernet) Clone() (CaptureEthernet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureEthernet()
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

func (obj *captureEthernet) setNil() {
	obj.srcHolder = nil
	obj.dstHolder = nil
	obj.etherTypeHolder = nil
	obj.pfcQueueHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// CaptureEthernet is description is TBD
type CaptureEthernet interface {
	Validation
	// msg marshals CaptureEthernet to protobuf object *otg.CaptureEthernet
	// and doesn't set defaults
	msg() *otg.CaptureEthernet
	// setMsg unmarshals CaptureEthernet from protobuf object *otg.CaptureEthernet
	// and doesn't set defaults
	setMsg(*otg.CaptureEthernet) CaptureEthernet
	// provides marshal interface
	Marshal() marshalCaptureEthernet
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureEthernet
	// validate validates CaptureEthernet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureEthernet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Src returns CaptureField, set in CaptureEthernet.
	// CaptureField is description is TBD
	Src() CaptureField
	// SetSrc assigns CaptureField provided by user to CaptureEthernet.
	// CaptureField is description is TBD
	SetSrc(value CaptureField) CaptureEthernet
	// HasSrc checks if Src has been set in CaptureEthernet
	HasSrc() bool
	// Dst returns CaptureField, set in CaptureEthernet.
	// CaptureField is description is TBD
	Dst() CaptureField
	// SetDst assigns CaptureField provided by user to CaptureEthernet.
	// CaptureField is description is TBD
	SetDst(value CaptureField) CaptureEthernet
	// HasDst checks if Dst has been set in CaptureEthernet
	HasDst() bool
	// EtherType returns CaptureField, set in CaptureEthernet.
	// CaptureField is description is TBD
	EtherType() CaptureField
	// SetEtherType assigns CaptureField provided by user to CaptureEthernet.
	// CaptureField is description is TBD
	SetEtherType(value CaptureField) CaptureEthernet
	// HasEtherType checks if EtherType has been set in CaptureEthernet
	HasEtherType() bool
	// PfcQueue returns CaptureField, set in CaptureEthernet.
	// CaptureField is description is TBD
	PfcQueue() CaptureField
	// SetPfcQueue assigns CaptureField provided by user to CaptureEthernet.
	// CaptureField is description is TBD
	SetPfcQueue(value CaptureField) CaptureEthernet
	// HasPfcQueue checks if PfcQueue has been set in CaptureEthernet
	HasPfcQueue() bool
	setNil()
}

// description is TBD
// Src returns a CaptureField
func (obj *captureEthernet) Src() CaptureField {
	if obj.obj.Src == nil {
		obj.obj.Src = NewCaptureField().msg()
	}
	if obj.srcHolder == nil {
		obj.srcHolder = &captureField{obj: obj.obj.Src}
	}
	return obj.srcHolder
}

// description is TBD
// Src returns a CaptureField
func (obj *captureEthernet) HasSrc() bool {
	return obj.obj.Src != nil
}

// description is TBD
// SetSrc sets the CaptureField value in the CaptureEthernet object
func (obj *captureEthernet) SetSrc(value CaptureField) CaptureEthernet {

	obj.srcHolder = nil
	obj.obj.Src = value.msg()

	return obj
}

// description is TBD
// Dst returns a CaptureField
func (obj *captureEthernet) Dst() CaptureField {
	if obj.obj.Dst == nil {
		obj.obj.Dst = NewCaptureField().msg()
	}
	if obj.dstHolder == nil {
		obj.dstHolder = &captureField{obj: obj.obj.Dst}
	}
	return obj.dstHolder
}

// description is TBD
// Dst returns a CaptureField
func (obj *captureEthernet) HasDst() bool {
	return obj.obj.Dst != nil
}

// description is TBD
// SetDst sets the CaptureField value in the CaptureEthernet object
func (obj *captureEthernet) SetDst(value CaptureField) CaptureEthernet {

	obj.dstHolder = nil
	obj.obj.Dst = value.msg()

	return obj
}

// description is TBD
// EtherType returns a CaptureField
func (obj *captureEthernet) EtherType() CaptureField {
	if obj.obj.EtherType == nil {
		obj.obj.EtherType = NewCaptureField().msg()
	}
	if obj.etherTypeHolder == nil {
		obj.etherTypeHolder = &captureField{obj: obj.obj.EtherType}
	}
	return obj.etherTypeHolder
}

// description is TBD
// EtherType returns a CaptureField
func (obj *captureEthernet) HasEtherType() bool {
	return obj.obj.EtherType != nil
}

// description is TBD
// SetEtherType sets the CaptureField value in the CaptureEthernet object
func (obj *captureEthernet) SetEtherType(value CaptureField) CaptureEthernet {

	obj.etherTypeHolder = nil
	obj.obj.EtherType = value.msg()

	return obj
}

// description is TBD
// PfcQueue returns a CaptureField
func (obj *captureEthernet) PfcQueue() CaptureField {
	if obj.obj.PfcQueue == nil {
		obj.obj.PfcQueue = NewCaptureField().msg()
	}
	if obj.pfcQueueHolder == nil {
		obj.pfcQueueHolder = &captureField{obj: obj.obj.PfcQueue}
	}
	return obj.pfcQueueHolder
}

// description is TBD
// PfcQueue returns a CaptureField
func (obj *captureEthernet) HasPfcQueue() bool {
	return obj.obj.PfcQueue != nil
}

// description is TBD
// SetPfcQueue sets the CaptureField value in the CaptureEthernet object
func (obj *captureEthernet) SetPfcQueue(value CaptureField) CaptureEthernet {

	obj.pfcQueueHolder = nil
	obj.obj.PfcQueue = value.msg()

	return obj
}

func (obj *captureEthernet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Src != nil {

		obj.Src().validateObj(vObj, set_default)
	}

	if obj.obj.Dst != nil {

		obj.Dst().validateObj(vObj, set_default)
	}

	if obj.obj.EtherType != nil {

		obj.EtherType().validateObj(vObj, set_default)
	}

	if obj.obj.PfcQueue != nil {

		obj.PfcQueue().validateObj(vObj, set_default)
	}

}

func (obj *captureEthernet) setDefault() {

}
