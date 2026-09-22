package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BmpServersStateRequest *****
type bmpServersStateRequest struct {
	validation
	obj          *otg.BmpServersStateRequest
	marshaller   marshalBmpServersStateRequest
	unMarshaller unMarshalBmpServersStateRequest
}

func NewBmpServersStateRequest() BmpServersStateRequest {
	obj := bmpServersStateRequest{obj: &otg.BmpServersStateRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *bmpServersStateRequest) msg() *otg.BmpServersStateRequest {
	return obj.obj
}

func (obj *bmpServersStateRequest) setMsg(msg *otg.BmpServersStateRequest) BmpServersStateRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbmpServersStateRequest struct {
	obj *bmpServersStateRequest
}

type marshalBmpServersStateRequest interface {
	// ToProto marshals BmpServersStateRequest to protobuf object *otg.BmpServersStateRequest
	ToProto() (*otg.BmpServersStateRequest, error)
	// ToPbText marshals BmpServersStateRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BmpServersStateRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals BmpServersStateRequest to JSON text
	ToJson() (string, error)
}

type unMarshalbmpServersStateRequest struct {
	obj *bmpServersStateRequest
}

type unMarshalBmpServersStateRequest interface {
	// FromProto unmarshals BmpServersStateRequest from protobuf object *otg.BmpServersStateRequest
	FromProto(msg *otg.BmpServersStateRequest) (BmpServersStateRequest, error)
	// FromPbText unmarshals BmpServersStateRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BmpServersStateRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BmpServersStateRequest from JSON text
	FromJson(value string) error
}

func (obj *bmpServersStateRequest) Marshal() marshalBmpServersStateRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbmpServersStateRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *bmpServersStateRequest) Unmarshal() unMarshalBmpServersStateRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbmpServersStateRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbmpServersStateRequest) ToProto() (*otg.BmpServersStateRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbmpServersStateRequest) FromProto(msg *otg.BmpServersStateRequest) (BmpServersStateRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbmpServersStateRequest) ToPbText() (string, error) {
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

func (m *unMarshalbmpServersStateRequest) FromPbText(value string) error {
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

func (m *marshalbmpServersStateRequest) ToYaml() (string, error) {
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

func (m *unMarshalbmpServersStateRequest) FromYaml(value string) error {
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

func (m *marshalbmpServersStateRequest) ToJson() (string, error) {
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

func (m *unMarshalbmpServersStateRequest) FromJson(value string) error {
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

func (obj *bmpServersStateRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bmpServersStateRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bmpServersStateRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bmpServersStateRequest) Clone() (BmpServersStateRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBmpServersStateRequest()
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

// BmpServersStateRequest is the request for BMP server information.
type BmpServersStateRequest interface {
	Validation
	// msg marshals BmpServersStateRequest to protobuf object *otg.BmpServersStateRequest
	// and doesn't set defaults
	msg() *otg.BmpServersStateRequest
	// setMsg unmarshals BmpServersStateRequest from protobuf object *otg.BmpServersStateRequest
	// and doesn't set defaults
	setMsg(*otg.BmpServersStateRequest) BmpServersStateRequest
	// provides marshal interface
	Marshal() marshalBmpServersStateRequest
	// provides unmarshal interface
	Unmarshal() unMarshalBmpServersStateRequest
	// validate validates BmpServersStateRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BmpServersStateRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// BmpServerNames returns []string, set in BmpServersStateRequest.
	BmpServerNames() []string
	// SetBmpServerNames assigns []string provided by user to BmpServersStateRequest
	SetBmpServerNames(value []string) BmpServersStateRequest
}

// The names of the BMP Servers  to return results for. An empty list will return results for all BMP Servers.
//
// x-constraint:
// - /components/schemas/Device.Bmp.ServerV4/properties/name
// - /components/schemas/Device.Bmp.ServerV6/properties/name
//
// BmpServerNames returns a []string
func (obj *bmpServersStateRequest) BmpServerNames() []string {
	if obj.obj.BmpServerNames == nil {
		obj.obj.BmpServerNames = make([]string, 0)
	}
	return obj.obj.BmpServerNames
}

// The names of the BMP Servers  to return results for. An empty list will return results for all BMP Servers.
//
// x-constraint:
// - /components/schemas/Device.Bmp.ServerV4/properties/name
// - /components/schemas/Device.Bmp.ServerV6/properties/name
//
// SetBmpServerNames sets the []string value in the BmpServersStateRequest object
func (obj *bmpServersStateRequest) SetBmpServerNames(value []string) BmpServersStateRequest {

	if obj.obj.BmpServerNames == nil {
		obj.obj.BmpServerNames = make([]string, 0)
	}
	obj.obj.BmpServerNames = value

	return obj
}

func (obj *bmpServersStateRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *bmpServersStateRequest) setDefault() {

}
