package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureRequestCaptureSlice *****
type captureRequestCaptureSlice struct {
	validation
	obj          *otg.CaptureRequestCaptureSlice
	marshaller   marshalCaptureRequestCaptureSlice
	unMarshaller unMarshalCaptureRequestCaptureSlice
}

func NewCaptureRequestCaptureSlice() CaptureRequestCaptureSlice {
	obj := captureRequestCaptureSlice{obj: &otg.CaptureRequestCaptureSlice{}}
	obj.setDefault()
	return &obj
}

func (obj *captureRequestCaptureSlice) msg() *otg.CaptureRequestCaptureSlice {
	return obj.obj
}

func (obj *captureRequestCaptureSlice) setMsg(msg *otg.CaptureRequestCaptureSlice) CaptureRequestCaptureSlice {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureRequestCaptureSlice struct {
	obj *captureRequestCaptureSlice
}

type marshalCaptureRequestCaptureSlice interface {
	// ToProto marshals CaptureRequestCaptureSlice to protobuf object *otg.CaptureRequestCaptureSlice
	ToProto() (*otg.CaptureRequestCaptureSlice, error)
	// ToPbText marshals CaptureRequestCaptureSlice to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureRequestCaptureSlice to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureRequestCaptureSlice to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureRequestCaptureSlice struct {
	obj *captureRequestCaptureSlice
}

type unMarshalCaptureRequestCaptureSlice interface {
	// FromProto unmarshals CaptureRequestCaptureSlice from protobuf object *otg.CaptureRequestCaptureSlice
	FromProto(msg *otg.CaptureRequestCaptureSlice) (CaptureRequestCaptureSlice, error)
	// FromPbText unmarshals CaptureRequestCaptureSlice from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureRequestCaptureSlice from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureRequestCaptureSlice from JSON text
	FromJson(value string) error
}

func (obj *captureRequestCaptureSlice) Marshal() marshalCaptureRequestCaptureSlice {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureRequestCaptureSlice{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureRequestCaptureSlice) Unmarshal() unMarshalCaptureRequestCaptureSlice {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureRequestCaptureSlice{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureRequestCaptureSlice) ToProto() (*otg.CaptureRequestCaptureSlice, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureRequestCaptureSlice) FromProto(msg *otg.CaptureRequestCaptureSlice) (CaptureRequestCaptureSlice, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureRequestCaptureSlice) ToPbText() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSlice) FromPbText(value string) error {
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

func (m *marshalcaptureRequestCaptureSlice) ToYaml() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSlice) FromYaml(value string) error {
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

func (m *marshalcaptureRequestCaptureSlice) ToJson() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSlice) FromJson(value string) error {
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

func (obj *captureRequestCaptureSlice) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureRequestCaptureSlice) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureRequestCaptureSlice) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureRequestCaptureSlice) Clone() (CaptureRequestCaptureSlice, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureRequestCaptureSlice()
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

// CaptureRequestCaptureSlice is packets to be captured based on specification of capture slice through start index and packet count.
type CaptureRequestCaptureSlice interface {
	Validation
	// msg marshals CaptureRequestCaptureSlice to protobuf object *otg.CaptureRequestCaptureSlice
	// and doesn't set defaults
	msg() *otg.CaptureRequestCaptureSlice
	// setMsg unmarshals CaptureRequestCaptureSlice from protobuf object *otg.CaptureRequestCaptureSlice
	// and doesn't set defaults
	setMsg(*otg.CaptureRequestCaptureSlice) CaptureRequestCaptureSlice
	// provides marshal interface
	Marshal() marshalCaptureRequestCaptureSlice
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureRequestCaptureSlice
	// validate validates CaptureRequestCaptureSlice
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureRequestCaptureSlice, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in CaptureRequestCaptureSlice.
	Start() uint64
	// SetStart assigns uint64 provided by user to CaptureRequestCaptureSlice
	SetStart(value uint64) CaptureRequestCaptureSlice
	// HasStart checks if Start has been set in CaptureRequestCaptureSlice
	HasStart() bool
	// Count returns uint64, set in CaptureRequestCaptureSlice.
	Count() uint64
	// SetCount assigns uint64 provided by user to CaptureRequestCaptureSlice
	SetCount(value uint64) CaptureRequestCaptureSlice
	// HasCount checks if Count has been set in CaptureRequestCaptureSlice
	HasCount() bool
}

// Index of the packet in the generated packet sequence from where capture would start.
// Start returns a uint64
func (obj *captureRequestCaptureSlice) Start() uint64 {

	return *obj.obj.Start

}

// Index of the packet in the generated packet sequence from where capture would start.
// Start returns a uint64
func (obj *captureRequestCaptureSlice) HasStart() bool {
	return obj.obj.Start != nil
}

// Index of the packet in the generated packet sequence from where capture would start.
// SetStart sets the uint64 value in the CaptureRequestCaptureSlice object
func (obj *captureRequestCaptureSlice) SetStart(value uint64) CaptureRequestCaptureSlice {

	obj.obj.Start = &value
	return obj
}

// Number of packets to be captured from the start index.
// Count returns a uint64
func (obj *captureRequestCaptureSlice) Count() uint64 {

	return *obj.obj.Count

}

// Number of packets to be captured from the start index.
// Count returns a uint64
func (obj *captureRequestCaptureSlice) HasCount() bool {
	return obj.obj.Count != nil
}

// Number of packets to be captured from the start index.
// SetCount sets the uint64 value in the CaptureRequestCaptureSlice object
func (obj *captureRequestCaptureSlice) SetCount(value uint64) CaptureRequestCaptureSlice {

	obj.obj.Count = &value
	return obj
}

func (obj *captureRequestCaptureSlice) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *captureRequestCaptureSlice) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}

}
