package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LdpFecLabelIncrement *****
type ldpFecLabelIncrement struct {
	validation
	obj          *otg.LdpFecLabelIncrement
	marshaller   marshalLdpFecLabelIncrement
	unMarshaller unMarshalLdpFecLabelIncrement
}

func NewLdpFecLabelIncrement() LdpFecLabelIncrement {
	obj := ldpFecLabelIncrement{obj: &otg.LdpFecLabelIncrement{}}
	obj.setDefault()
	return &obj
}

func (obj *ldpFecLabelIncrement) msg() *otg.LdpFecLabelIncrement {
	return obj.obj
}

func (obj *ldpFecLabelIncrement) setMsg(msg *otg.LdpFecLabelIncrement) LdpFecLabelIncrement {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalldpFecLabelIncrement struct {
	obj *ldpFecLabelIncrement
}

type marshalLdpFecLabelIncrement interface {
	// ToProto marshals LdpFecLabelIncrement to protobuf object *otg.LdpFecLabelIncrement
	ToProto() (*otg.LdpFecLabelIncrement, error)
	// ToPbText marshals LdpFecLabelIncrement to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LdpFecLabelIncrement to YAML text
	ToYaml() (string, error)
	// ToJson marshals LdpFecLabelIncrement to JSON text
	ToJson() (string, error)
}

type unMarshalldpFecLabelIncrement struct {
	obj *ldpFecLabelIncrement
}

type unMarshalLdpFecLabelIncrement interface {
	// FromProto unmarshals LdpFecLabelIncrement from protobuf object *otg.LdpFecLabelIncrement
	FromProto(msg *otg.LdpFecLabelIncrement) (LdpFecLabelIncrement, error)
	// FromPbText unmarshals LdpFecLabelIncrement from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LdpFecLabelIncrement from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LdpFecLabelIncrement from JSON text
	FromJson(value string) error
}

func (obj *ldpFecLabelIncrement) Marshal() marshalLdpFecLabelIncrement {
	if obj.marshaller == nil {
		obj.marshaller = &marshalldpFecLabelIncrement{obj: obj}
	}
	return obj.marshaller
}

func (obj *ldpFecLabelIncrement) Unmarshal() unMarshalLdpFecLabelIncrement {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalldpFecLabelIncrement{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalldpFecLabelIncrement) ToProto() (*otg.LdpFecLabelIncrement, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalldpFecLabelIncrement) FromProto(msg *otg.LdpFecLabelIncrement) (LdpFecLabelIncrement, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalldpFecLabelIncrement) ToPbText() (string, error) {
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

func (m *unMarshalldpFecLabelIncrement) FromPbText(value string) error {
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

func (m *marshalldpFecLabelIncrement) ToYaml() (string, error) {
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

func (m *unMarshalldpFecLabelIncrement) FromYaml(value string) error {
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

func (m *marshalldpFecLabelIncrement) ToJson() (string, error) {
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

func (m *unMarshalldpFecLabelIncrement) FromJson(value string) error {
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

func (obj *ldpFecLabelIncrement) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ldpFecLabelIncrement) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ldpFecLabelIncrement) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ldpFecLabelIncrement) Clone() (LdpFecLabelIncrement, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLdpFecLabelIncrement()
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

// LdpFecLabelIncrement is labels that increase by 1 from one prefix to the next.
type LdpFecLabelIncrement interface {
	Validation
	// msg marshals LdpFecLabelIncrement to protobuf object *otg.LdpFecLabelIncrement
	// and doesn't set defaults
	msg() *otg.LdpFecLabelIncrement
	// setMsg unmarshals LdpFecLabelIncrement from protobuf object *otg.LdpFecLabelIncrement
	// and doesn't set defaults
	setMsg(*otg.LdpFecLabelIncrement) LdpFecLabelIncrement
	// provides marshal interface
	Marshal() marshalLdpFecLabelIncrement
	// provides unmarshal interface
	Unmarshal() unMarshalLdpFecLabelIncrement
	// validate validates LdpFecLabelIncrement
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LdpFecLabelIncrement, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in LdpFecLabelIncrement.
	Start() uint32
	// SetStart assigns uint32 provided by user to LdpFecLabelIncrement
	SetStart(value uint32) LdpFecLabelIncrement
	// HasStart checks if Start has been set in LdpFecLabelIncrement
	HasStart() bool
}

// The label bound to the first prefix of the range. Labels 0 to 15 are reserved (RFC 3032 Section 2.1). start plus the number of prefixes in the range minus 1 must not exceed 1048575; the implementation rejects a configuration that does.
// Start returns a uint32
func (obj *ldpFecLabelIncrement) Start() uint32 {

	return *obj.obj.Start

}

// The label bound to the first prefix of the range. Labels 0 to 15 are reserved (RFC 3032 Section 2.1). start plus the number of prefixes in the range minus 1 must not exceed 1048575; the implementation rejects a configuration that does.
// Start returns a uint32
func (obj *ldpFecLabelIncrement) HasStart() bool {
	return obj.obj.Start != nil
}

// The label bound to the first prefix of the range. Labels 0 to 15 are reserved (RFC 3032 Section 2.1). start plus the number of prefixes in the range minus 1 must not exceed 1048575; the implementation rejects a configuration that does.
// SetStart sets the uint32 value in the LdpFecLabelIncrement object
func (obj *ldpFecLabelIncrement) SetStart(value uint32) LdpFecLabelIncrement {

	obj.obj.Start = &value
	return obj
}

func (obj *ldpFecLabelIncrement) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start < 16 || *obj.obj.Start > 1048575 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("16 <= LdpFecLabelIncrement.Start <= 1048575 but Got %d", *obj.obj.Start))
		}

	}

}

func (obj *ldpFecLabelIncrement) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(16)
	}

}
