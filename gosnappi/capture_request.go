package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CaptureRequest *****
type captureRequest struct {
	validation
	obj          *otg.CaptureRequest
	marshaller   marshalCaptureRequest
	unMarshaller unMarshalCaptureRequest
}

func NewCaptureRequest() CaptureRequest {
	obj := captureRequest{obj: &otg.CaptureRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *captureRequest) msg() *otg.CaptureRequest {
	return obj.obj
}

func (obj *captureRequest) setMsg(msg *otg.CaptureRequest) CaptureRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcaptureRequest struct {
	obj *captureRequest
}

type marshalCaptureRequest interface {
	// ToProto marshals CaptureRequest to protobuf object *otg.CaptureRequest
	ToProto() (*otg.CaptureRequest, error)
	// ToPbText marshals CaptureRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CaptureRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals CaptureRequest to JSON text
	ToJson() (string, error)
}

type unMarshalcaptureRequest struct {
	obj *captureRequest
}

type unMarshalCaptureRequest interface {
	// FromProto unmarshals CaptureRequest from protobuf object *otg.CaptureRequest
	FromProto(msg *otg.CaptureRequest) (CaptureRequest, error)
	// FromPbText unmarshals CaptureRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CaptureRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CaptureRequest from JSON text
	FromJson(value string) error
}

func (obj *captureRequest) Marshal() marshalCaptureRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcaptureRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *captureRequest) Unmarshal() unMarshalCaptureRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcaptureRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcaptureRequest) ToProto() (*otg.CaptureRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcaptureRequest) FromProto(msg *otg.CaptureRequest) (CaptureRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcaptureRequest) ToPbText() (string, error) {
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

func (m *unMarshalcaptureRequest) FromPbText(value string) error {
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

func (m *marshalcaptureRequest) ToYaml() (string, error) {
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

func (m *unMarshalcaptureRequest) FromYaml(value string) error {
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

func (m *marshalcaptureRequest) ToJson() (string, error) {
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

func (m *unMarshalcaptureRequest) FromJson(value string) error {
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

func (obj *captureRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *captureRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *captureRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *captureRequest) Clone() (CaptureRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCaptureRequest()
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

// CaptureRequest is the capture result request to the traffic generator. Stops the port capture on the port_name and returns the capture.
type CaptureRequest interface {
	Validation
	// msg marshals CaptureRequest to protobuf object *otg.CaptureRequest
	// and doesn't set defaults
	msg() *otg.CaptureRequest
	// setMsg unmarshals CaptureRequest from protobuf object *otg.CaptureRequest
	// and doesn't set defaults
	setMsg(*otg.CaptureRequest) CaptureRequest
	// provides marshal interface
	Marshal() marshalCaptureRequest
	// provides unmarshal interface
	Unmarshal() unMarshalCaptureRequest
	// validate validates CaptureRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CaptureRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortName returns string, set in CaptureRequest.
	PortName() string
	// SetPortName assigns string provided by user to CaptureRequest
	SetPortName(value string) CaptureRequest
}

// The name of a port a capture is started on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *captureRequest) PortName() string {

	return *obj.obj.PortName

}

// The name of a port a capture is started on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the CaptureRequest object
func (obj *captureRequest) SetPortName(value string) CaptureRequest {

	obj.obj.PortName = &value
	return obj
}

func (obj *captureRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// PortName is required
	if obj.obj.PortName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "PortName is required field on interface CaptureRequest")
	}
}

func (obj *captureRequest) setDefault() {

}
