package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisMT *****
type isisMT struct {
	validation
	obj          *otg.IsisMT
	marshaller   marshalIsisMT
	unMarshaller unMarshalIsisMT
}

func NewIsisMT() IsisMT {
	obj := isisMT{obj: &otg.IsisMT{}}
	obj.setDefault()
	return &obj
}

func (obj *isisMT) msg() *otg.IsisMT {
	return obj.obj
}

func (obj *isisMT) setMsg(msg *otg.IsisMT) IsisMT {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisMT struct {
	obj *isisMT
}

type marshalIsisMT interface {
	// ToProto marshals IsisMT to protobuf object *otg.IsisMT
	ToProto() (*otg.IsisMT, error)
	// ToPbText marshals IsisMT to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisMT to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisMT to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisMT to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisMT struct {
	obj *isisMT
}

type unMarshalIsisMT interface {
	// FromProto unmarshals IsisMT from protobuf object *otg.IsisMT
	FromProto(msg *otg.IsisMT) (IsisMT, error)
	// FromPbText unmarshals IsisMT from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisMT from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisMT from JSON text
	FromJson(value string) error
}

func (obj *isisMT) Marshal() marshalIsisMT {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisMT{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisMT) Unmarshal() unMarshalIsisMT {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisMT{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisMT) ToProto() (*otg.IsisMT, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisMT) FromProto(msg *otg.IsisMT) (IsisMT, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisMT) ToPbText() (string, error) {
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

func (m *unMarshalisisMT) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalisisMT) ToYaml() (string, error) {
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

func (m *unMarshalisisMT) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalisisMT) ToJsonRaw() (string, error) {
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

func (m *marshalisisMT) ToJson() (string, error) {
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

func (m *unMarshalisisMT) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *isisMT) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisMT) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisMT) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisMT) Clone() (IsisMT, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisMT()
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

// IsisMT is configuration of properties per interface per topology when multiple topologies are configured in an ISIS router.
// in a ISIS router.
type IsisMT interface {
	Validation
	// msg marshals IsisMT to protobuf object *otg.IsisMT
	// and doesn't set defaults
	msg() *otg.IsisMT
	// setMsg unmarshals IsisMT from protobuf object *otg.IsisMT
	// and doesn't set defaults
	setMsg(*otg.IsisMT) IsisMT
	// provides marshal interface
	Marshal() marshalIsisMT
	// provides unmarshal interface
	Unmarshal() unMarshalIsisMT
	// validate validates IsisMT
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisMT, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MtId returns uint32, set in IsisMT.
	MtId() uint32
	// SetMtId assigns uint32 provided by user to IsisMT
	SetMtId(value uint32) IsisMT
	// HasMtId checks if MtId has been set in IsisMT
	HasMtId() bool
	// LinkMetric returns uint32, set in IsisMT.
	LinkMetric() uint32
	// SetLinkMetric assigns uint32 provided by user to IsisMT
	SetLinkMetric(value uint32) IsisMT
	// HasLinkMetric checks if LinkMetric has been set in IsisMT
	HasLinkMetric() bool
}

// The Multi Topology ID for one of the topologies supported on the ISIS interface.
// MtId returns a uint32
func (obj *isisMT) MtId() uint32 {

	return *obj.obj.MtId

}

// The Multi Topology ID for one of the topologies supported on the ISIS interface.
// MtId returns a uint32
func (obj *isisMT) HasMtId() bool {
	return obj.obj.MtId != nil
}

// The Multi Topology ID for one of the topologies supported on the ISIS interface.
// SetMtId sets the uint32 value in the IsisMT object
func (obj *isisMT) SetMtId(value uint32) IsisMT {

	obj.obj.MtId = &value
	return obj
}

// Specifies the link metric for this topology on the ISIS interface.
// LinkMetric returns a uint32
func (obj *isisMT) LinkMetric() uint32 {

	return *obj.obj.LinkMetric

}

// Specifies the link metric for this topology on the ISIS interface.
// LinkMetric returns a uint32
func (obj *isisMT) HasLinkMetric() bool {
	return obj.obj.LinkMetric != nil
}

// Specifies the link metric for this topology on the ISIS interface.
// SetLinkMetric sets the uint32 value in the IsisMT object
func (obj *isisMT) SetLinkMetric(value uint32) IsisMT {

	obj.obj.LinkMetric = &value
	return obj
}

func (obj *isisMT) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MtId != nil {

		if *obj.obj.MtId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisMT.MtId <= 65535 but Got %d", *obj.obj.MtId))
		}

	}

	if obj.obj.LinkMetric != nil {

		if *obj.obj.LinkMetric > 16777215 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisMT.LinkMetric <= 16777215 but Got %d", *obj.obj.LinkMetric))
		}

	}

}

func (obj *isisMT) setDefault() {
	if obj.obj.MtId == nil {
		obj.SetMtId(0)
	}
	if obj.obj.LinkMetric == nil {
		obj.SetLinkMetric(10)
	}

}
