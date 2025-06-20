package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagProtocolStatic *****
type lagProtocolStatic struct {
	validation
	obj          *otg.LagProtocolStatic
	marshaller   marshalLagProtocolStatic
	unMarshaller unMarshalLagProtocolStatic
}

func NewLagProtocolStatic() LagProtocolStatic {
	obj := lagProtocolStatic{obj: &otg.LagProtocolStatic{}}
	obj.setDefault()
	return &obj
}

func (obj *lagProtocolStatic) msg() *otg.LagProtocolStatic {
	return obj.obj
}

func (obj *lagProtocolStatic) setMsg(msg *otg.LagProtocolStatic) LagProtocolStatic {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagProtocolStatic struct {
	obj *lagProtocolStatic
}

type marshalLagProtocolStatic interface {
	// ToProto marshals LagProtocolStatic to protobuf object *otg.LagProtocolStatic
	ToProto() (*otg.LagProtocolStatic, error)
	// ToPbText marshals LagProtocolStatic to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagProtocolStatic to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagProtocolStatic to JSON text
	ToJson() (string, error)
}

type unMarshallagProtocolStatic struct {
	obj *lagProtocolStatic
}

type unMarshalLagProtocolStatic interface {
	// FromProto unmarshals LagProtocolStatic from protobuf object *otg.LagProtocolStatic
	FromProto(msg *otg.LagProtocolStatic) (LagProtocolStatic, error)
	// FromPbText unmarshals LagProtocolStatic from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagProtocolStatic from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagProtocolStatic from JSON text
	FromJson(value string) error
}

func (obj *lagProtocolStatic) Marshal() marshalLagProtocolStatic {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagProtocolStatic{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagProtocolStatic) Unmarshal() unMarshalLagProtocolStatic {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagProtocolStatic{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagProtocolStatic) ToProto() (*otg.LagProtocolStatic, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagProtocolStatic) FromProto(msg *otg.LagProtocolStatic) (LagProtocolStatic, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagProtocolStatic) ToPbText() (string, error) {
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

func (m *unMarshallagProtocolStatic) FromPbText(value string) error {
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

func (m *marshallagProtocolStatic) ToYaml() (string, error) {
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

func (m *unMarshallagProtocolStatic) FromYaml(value string) error {
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

func (m *marshallagProtocolStatic) ToJson() (string, error) {
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

func (m *unMarshallagProtocolStatic) FromJson(value string) error {
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

func (obj *lagProtocolStatic) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagProtocolStatic) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagProtocolStatic) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagProtocolStatic) Clone() (LagProtocolStatic, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagProtocolStatic()
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

// LagProtocolStatic is the container for static link aggregation protocol settings.
type LagProtocolStatic interface {
	Validation
	// msg marshals LagProtocolStatic to protobuf object *otg.LagProtocolStatic
	// and doesn't set defaults
	msg() *otg.LagProtocolStatic
	// setMsg unmarshals LagProtocolStatic from protobuf object *otg.LagProtocolStatic
	// and doesn't set defaults
	setMsg(*otg.LagProtocolStatic) LagProtocolStatic
	// provides marshal interface
	Marshal() marshalLagProtocolStatic
	// provides unmarshal interface
	Unmarshal() unMarshalLagProtocolStatic
	// validate validates LagProtocolStatic
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagProtocolStatic, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LagId returns uint32, set in LagProtocolStatic.
	LagId() uint32
	// SetLagId assigns uint32 provided by user to LagProtocolStatic
	SetLagId(value uint32) LagProtocolStatic
	// HasLagId checks if LagId has been set in LagProtocolStatic
	HasLagId() bool
}

// The static lag id
// LagId returns a uint32
func (obj *lagProtocolStatic) LagId() uint32 {

	return *obj.obj.LagId

}

// The static lag id
// LagId returns a uint32
func (obj *lagProtocolStatic) HasLagId() bool {
	return obj.obj.LagId != nil
}

// The static lag id
// SetLagId sets the uint32 value in the LagProtocolStatic object
func (obj *lagProtocolStatic) SetLagId(value uint32) LagProtocolStatic {

	obj.obj.LagId = &value
	return obj
}

func (obj *lagProtocolStatic) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LagId != nil {

		if *obj.obj.LagId > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= LagProtocolStatic.LagId <= 65535 but Got %d", *obj.obj.LagId))
		}

	}

}

func (obj *lagProtocolStatic) setDefault() {
	if obj.obj.LagId == nil {
		obj.SetLagId(0)
	}

}
