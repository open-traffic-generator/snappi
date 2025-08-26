package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureRequestCaptureSliceInitial *****
type captureRequestCaptureSliceInitial struct {
	validation
	obj          *otg.CaptureRequestCaptureSliceInitial
	marshaller   marshalCaptureRequestCaptureSliceInitial
	unMarshaller unMarshalCaptureRequestCaptureSliceInitial
}

func NewCaptureRequestCaptureSliceInitial() CaptureRequestCaptureSliceInitial {
	obj := captureRequestCaptureSliceInitial{obj: &otg.CaptureRequestCaptureSliceInitial{}}
	obj.setDefault()
	return &obj
}

func (obj *captureRequestCaptureSliceInitial) msg() *otg.CaptureRequestCaptureSliceInitial {
	return obj.obj
}

func (obj *captureRequestCaptureSliceInitial) setMsg(msg *otg.CaptureRequestCaptureSliceInitial) CaptureRequestCaptureSliceInitial {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureRequestCaptureSliceInitial struct {
	obj *captureRequestCaptureSliceInitial
}

type marshalCaptureRequestCaptureSliceInitial interface {
	// ToProto marshals CaptureRequestCaptureSliceInitial to protobuf object *otg.CaptureRequestCaptureSliceInitial
	ToProto() (*otg.CaptureRequestCaptureSliceInitial, error)
	// ToPbText marshals CaptureRequestCaptureSliceInitial to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureRequestCaptureSliceInitial to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureRequestCaptureSliceInitial to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureRequestCaptureSliceInitial struct {
	obj *captureRequestCaptureSliceInitial
}

type unMarshalCaptureRequestCaptureSliceInitial interface {
	// FromProto unmarshals CaptureRequestCaptureSliceInitial from protobuf object *otg.CaptureRequestCaptureSliceInitial
	FromProto(msg *otg.CaptureRequestCaptureSliceInitial) (CaptureRequestCaptureSliceInitial, error)
	// FromPbText unmarshals CaptureRequestCaptureSliceInitial from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureRequestCaptureSliceInitial from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureRequestCaptureSliceInitial from JSON text
	FromJson(value string) error
}

func (obj *captureRequestCaptureSliceInitial) Marshal() marshalCaptureRequestCaptureSliceInitial {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureRequestCaptureSliceInitial{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureRequestCaptureSliceInitial) Unmarshal() unMarshalCaptureRequestCaptureSliceInitial {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureRequestCaptureSliceInitial{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureRequestCaptureSliceInitial) ToProto() (*otg.CaptureRequestCaptureSliceInitial, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureRequestCaptureSliceInitial) FromProto(msg *otg.CaptureRequestCaptureSliceInitial) (CaptureRequestCaptureSliceInitial, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureRequestCaptureSliceInitial) ToPbText() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSliceInitial) FromPbText(value string) error {
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

func (m *marshalcaptureRequestCaptureSliceInitial) ToYaml() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSliceInitial) FromYaml(value string) error {
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

func (m *marshalcaptureRequestCaptureSliceInitial) ToJson() (string, error) {
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

func (m *unMarshalcaptureRequestCaptureSliceInitial) FromJson(value string) error {
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

func (obj *captureRequestCaptureSliceInitial) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureRequestCaptureSliceInitial) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureRequestCaptureSliceInitial) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureRequestCaptureSliceInitial) Clone() (CaptureRequestCaptureSliceInitial, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureRequestCaptureSliceInitial()
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

// CaptureRequestCaptureSliceInitial is specification of capture slice to capture packets from begining of captured packet sequence.
type CaptureRequestCaptureSliceInitial interface {
	Validation
	// msg marshals CaptureRequestCaptureSliceInitial to protobuf object *otg.CaptureRequestCaptureSliceInitial
	// and doesn't set defaults
	msg() *otg.CaptureRequestCaptureSliceInitial
	// setMsg unmarshals CaptureRequestCaptureSliceInitial from protobuf object *otg.CaptureRequestCaptureSliceInitial
	// and doesn't set defaults
	setMsg(*otg.CaptureRequestCaptureSliceInitial) CaptureRequestCaptureSliceInitial
	// provides marshal interface
	Marshal() marshalCaptureRequestCaptureSliceInitial
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureRequestCaptureSliceInitial
	// validate validates CaptureRequestCaptureSliceInitial
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureRequestCaptureSliceInitial, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint64, set in CaptureRequestCaptureSliceInitial.
	Start() uint64
	// SetStart assigns uint64 provided by user to CaptureRequestCaptureSliceInitial
	SetStart(value uint64) CaptureRequestCaptureSliceInitial
	// HasStart checks if Start has been set in CaptureRequestCaptureSliceInitial
	HasStart() bool
	// Count returns uint64, set in CaptureRequestCaptureSliceInitial.
	Count() uint64
	// SetCount assigns uint64 provided by user to CaptureRequestCaptureSliceInitial
	SetCount(value uint64) CaptureRequestCaptureSliceInitial
	// HasCount checks if Count has been set in CaptureRequestCaptureSliceInitial
	HasCount() bool
}

// Position of the packet (Nth) in the generated packet sequence from where capture would start.
// Start returns a uint64
func (obj *captureRequestCaptureSliceInitial) Start() uint64 {

	return *obj.obj.Start

}

// Position of the packet (Nth) in the generated packet sequence from where capture would start.
// Start returns a uint64
func (obj *captureRequestCaptureSliceInitial) HasStart() bool {
	return obj.obj.Start != nil
}

// Position of the packet (Nth) in the generated packet sequence from where capture would start.
// SetStart sets the uint64 value in the CaptureRequestCaptureSliceInitial object
func (obj *captureRequestCaptureSliceInitial) SetStart(value uint64) CaptureRequestCaptureSliceInitial {

	obj.obj.Start = &value
	return obj
}

// Number of packets to be captured from the position of first packet.
// Count returns a uint64
func (obj *captureRequestCaptureSliceInitial) Count() uint64 {

	return *obj.obj.Count

}

// Number of packets to be captured from the position of first packet.
// Count returns a uint64
func (obj *captureRequestCaptureSliceInitial) HasCount() bool {
	return obj.obj.Count != nil
}

// Number of packets to be captured from the position of first packet.
// SetCount sets the uint64 value in the CaptureRequestCaptureSliceInitial object
func (obj *captureRequestCaptureSliceInitial) SetCount(value uint64) CaptureRequestCaptureSliceInitial {

	obj.obj.Count = &value
	return obj
}

func (obj *captureRequestCaptureSliceInitial) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start < 1 || *obj.obj.Start > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= CaptureRequestCaptureSliceInitial.Start <= 18446744073709551615 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < 1 || *obj.obj.Count > 18446744073709551615 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= CaptureRequestCaptureSliceInitial.Count <= 18446744073709551615 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *captureRequestCaptureSliceInitial) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(100)
	}

}
