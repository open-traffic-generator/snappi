package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Layer1AutoNegotiation *****
type layer1AutoNegotiation struct {
	validation
	obj          *otg.Layer1AutoNegotiation
	marshaller   marshalLayer1AutoNegotiation
	unMarshaller unMarshalLayer1AutoNegotiation
}

func NewLayer1AutoNegotiation() Layer1AutoNegotiation {
	obj := layer1AutoNegotiation{obj: &otg.Layer1AutoNegotiation{}}
	obj.setDefault()
	return &obj
}

func (obj *layer1AutoNegotiation) msg() *otg.Layer1AutoNegotiation {
	return obj.obj
}

func (obj *layer1AutoNegotiation) setMsg(msg *otg.Layer1AutoNegotiation) Layer1AutoNegotiation {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallayer1AutoNegotiation struct {
	obj *layer1AutoNegotiation
}

type marshalLayer1AutoNegotiation interface {
	// ToProto marshals Layer1AutoNegotiation to protobuf object *otg.Layer1AutoNegotiation
	ToProto() (*otg.Layer1AutoNegotiation, error)
	// ToPbText marshals Layer1AutoNegotiation to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Layer1AutoNegotiation to YAML text
	ToYaml() (string, error)
	// ToJson marshals Layer1AutoNegotiation to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Layer1AutoNegotiation to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallayer1AutoNegotiation struct {
	obj *layer1AutoNegotiation
}

type unMarshalLayer1AutoNegotiation interface {
	// FromProto unmarshals Layer1AutoNegotiation from protobuf object *otg.Layer1AutoNegotiation
	FromProto(msg *otg.Layer1AutoNegotiation) (Layer1AutoNegotiation, error)
	// FromPbText unmarshals Layer1AutoNegotiation from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Layer1AutoNegotiation from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Layer1AutoNegotiation from JSON text
	FromJson(value string) error
}

func (obj *layer1AutoNegotiation) Marshal() marshalLayer1AutoNegotiation {
	if obj.marshaller == nil {
		obj.marshaller = &marshallayer1AutoNegotiation{obj: obj}
	}
	return obj.marshaller
}

func (obj *layer1AutoNegotiation) Unmarshal() unMarshalLayer1AutoNegotiation {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallayer1AutoNegotiation{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallayer1AutoNegotiation) ToProto() (*otg.Layer1AutoNegotiation, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallayer1AutoNegotiation) FromProto(msg *otg.Layer1AutoNegotiation) (Layer1AutoNegotiation, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallayer1AutoNegotiation) ToPbText() (string, error) {
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

func (m *unMarshallayer1AutoNegotiation) FromPbText(value string) error {
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

func (m *marshallayer1AutoNegotiation) ToYaml() (string, error) {
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

func (m *unMarshallayer1AutoNegotiation) FromYaml(value string) error {
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

func (m *marshallayer1AutoNegotiation) ToJsonRaw() (string, error) {
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

func (m *marshallayer1AutoNegotiation) ToJson() (string, error) {
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

func (m *unMarshallayer1AutoNegotiation) FromJson(value string) error {
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

func (obj *layer1AutoNegotiation) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *layer1AutoNegotiation) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *layer1AutoNegotiation) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *layer1AutoNegotiation) Clone() (Layer1AutoNegotiation, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLayer1AutoNegotiation()
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

// Layer1AutoNegotiation is configuration for auto negotiation settings
type Layer1AutoNegotiation interface {
	Validation
	// msg marshals Layer1AutoNegotiation to protobuf object *otg.Layer1AutoNegotiation
	// and doesn't set defaults
	msg() *otg.Layer1AutoNegotiation
	// setMsg unmarshals Layer1AutoNegotiation from protobuf object *otg.Layer1AutoNegotiation
	// and doesn't set defaults
	setMsg(*otg.Layer1AutoNegotiation) Layer1AutoNegotiation
	// provides marshal interface
	Marshal() marshalLayer1AutoNegotiation
	// provides unmarshal interface
	Unmarshal() unMarshalLayer1AutoNegotiation
	// validate validates Layer1AutoNegotiation
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Layer1AutoNegotiation, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Advertise1000Mbps returns bool, set in Layer1AutoNegotiation.
	Advertise1000Mbps() bool
	// SetAdvertise1000Mbps assigns bool provided by user to Layer1AutoNegotiation
	SetAdvertise1000Mbps(value bool) Layer1AutoNegotiation
	// HasAdvertise1000Mbps checks if Advertise1000Mbps has been set in Layer1AutoNegotiation
	HasAdvertise1000Mbps() bool
	// Advertise100FdMbps returns bool, set in Layer1AutoNegotiation.
	Advertise100FdMbps() bool
	// SetAdvertise100FdMbps assigns bool provided by user to Layer1AutoNegotiation
	SetAdvertise100FdMbps(value bool) Layer1AutoNegotiation
	// HasAdvertise100FdMbps checks if Advertise100FdMbps has been set in Layer1AutoNegotiation
	HasAdvertise100FdMbps() bool
	// Advertise100HdMbps returns bool, set in Layer1AutoNegotiation.
	Advertise100HdMbps() bool
	// SetAdvertise100HdMbps assigns bool provided by user to Layer1AutoNegotiation
	SetAdvertise100HdMbps(value bool) Layer1AutoNegotiation
	// HasAdvertise100HdMbps checks if Advertise100HdMbps has been set in Layer1AutoNegotiation
	HasAdvertise100HdMbps() bool
	// Advertise10FdMbps returns bool, set in Layer1AutoNegotiation.
	Advertise10FdMbps() bool
	// SetAdvertise10FdMbps assigns bool provided by user to Layer1AutoNegotiation
	SetAdvertise10FdMbps(value bool) Layer1AutoNegotiation
	// HasAdvertise10FdMbps checks if Advertise10FdMbps has been set in Layer1AutoNegotiation
	HasAdvertise10FdMbps() bool
	// Advertise10HdMbps returns bool, set in Layer1AutoNegotiation.
	Advertise10HdMbps() bool
	// SetAdvertise10HdMbps assigns bool provided by user to Layer1AutoNegotiation
	SetAdvertise10HdMbps(value bool) Layer1AutoNegotiation
	// HasAdvertise10HdMbps checks if Advertise10HdMbps has been set in Layer1AutoNegotiation
	HasAdvertise10HdMbps() bool
	// LinkTraining returns bool, set in Layer1AutoNegotiation.
	LinkTraining() bool
	// SetLinkTraining assigns bool provided by user to Layer1AutoNegotiation
	SetLinkTraining(value bool) Layer1AutoNegotiation
	// HasLinkTraining checks if LinkTraining has been set in Layer1AutoNegotiation
	HasLinkTraining() bool
	// RsFec returns bool, set in Layer1AutoNegotiation.
	RsFec() bool
	// SetRsFec assigns bool provided by user to Layer1AutoNegotiation
	SetRsFec(value bool) Layer1AutoNegotiation
	// HasRsFec checks if RsFec has been set in Layer1AutoNegotiation
	HasRsFec() bool
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise1000Mbps returns a bool
func (obj *layer1AutoNegotiation) Advertise1000Mbps() bool {

	return *obj.obj.Advertise_1000Mbps

}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise1000Mbps returns a bool
func (obj *layer1AutoNegotiation) HasAdvertise1000Mbps() bool {
	return obj.obj.Advertise_1000Mbps != nil
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// SetAdvertise1000Mbps sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetAdvertise1000Mbps(value bool) Layer1AutoNegotiation {

	obj.obj.Advertise_1000Mbps = &value
	return obj
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise100FdMbps returns a bool
func (obj *layer1AutoNegotiation) Advertise100FdMbps() bool {

	return *obj.obj.Advertise_100FdMbps

}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise100FdMbps returns a bool
func (obj *layer1AutoNegotiation) HasAdvertise100FdMbps() bool {
	return obj.obj.Advertise_100FdMbps != nil
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// SetAdvertise100FdMbps sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetAdvertise100FdMbps(value bool) Layer1AutoNegotiation {

	obj.obj.Advertise_100FdMbps = &value
	return obj
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise100HdMbps returns a bool
func (obj *layer1AutoNegotiation) Advertise100HdMbps() bool {

	return *obj.obj.Advertise_100HdMbps

}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise100HdMbps returns a bool
func (obj *layer1AutoNegotiation) HasAdvertise100HdMbps() bool {
	return obj.obj.Advertise_100HdMbps != nil
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// SetAdvertise100HdMbps sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetAdvertise100HdMbps(value bool) Layer1AutoNegotiation {

	obj.obj.Advertise_100HdMbps = &value
	return obj
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise10FdMbps returns a bool
func (obj *layer1AutoNegotiation) Advertise10FdMbps() bool {

	return *obj.obj.Advertise_10FdMbps

}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise10FdMbps returns a bool
func (obj *layer1AutoNegotiation) HasAdvertise10FdMbps() bool {
	return obj.obj.Advertise_10FdMbps != nil
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// SetAdvertise10FdMbps sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetAdvertise10FdMbps(value bool) Layer1AutoNegotiation {

	obj.obj.Advertise_10FdMbps = &value
	return obj
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise10HdMbps returns a bool
func (obj *layer1AutoNegotiation) Advertise10HdMbps() bool {

	return *obj.obj.Advertise_10HdMbps

}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// Advertise10HdMbps returns a bool
func (obj *layer1AutoNegotiation) HasAdvertise10HdMbps() bool {
	return obj.obj.Advertise_10HdMbps != nil
}

// If auto_negotiate is true and the interface supports this option
// then this speed will be advertised.
// SetAdvertise10HdMbps sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetAdvertise10HdMbps(value bool) Layer1AutoNegotiation {

	obj.obj.Advertise_10HdMbps = &value
	return obj
}

// Enable/disable gigabit ethernet link training.
// LinkTraining returns a bool
func (obj *layer1AutoNegotiation) LinkTraining() bool {

	return *obj.obj.LinkTraining

}

// Enable/disable gigabit ethernet link training.
// LinkTraining returns a bool
func (obj *layer1AutoNegotiation) HasLinkTraining() bool {
	return obj.obj.LinkTraining != nil
}

// Enable/disable gigabit ethernet link training.
// SetLinkTraining sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetLinkTraining(value bool) Layer1AutoNegotiation {

	obj.obj.LinkTraining = &value
	return obj
}

// Enable/disable gigabit ethernet reed solomon forward error correction (RS FEC).
// RsFec returns a bool
func (obj *layer1AutoNegotiation) RsFec() bool {

	return *obj.obj.RsFec

}

// Enable/disable gigabit ethernet reed solomon forward error correction (RS FEC).
// RsFec returns a bool
func (obj *layer1AutoNegotiation) HasRsFec() bool {
	return obj.obj.RsFec != nil
}

// Enable/disable gigabit ethernet reed solomon forward error correction (RS FEC).
// SetRsFec sets the bool value in the Layer1AutoNegotiation object
func (obj *layer1AutoNegotiation) SetRsFec(value bool) Layer1AutoNegotiation {

	obj.obj.RsFec = &value
	return obj
}

func (obj *layer1AutoNegotiation) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *layer1AutoNegotiation) setDefault() {
	if obj.obj.Advertise_1000Mbps == nil {
		obj.SetAdvertise1000Mbps(true)
	}
	if obj.obj.Advertise_100FdMbps == nil {
		obj.SetAdvertise100FdMbps(true)
	}
	if obj.obj.Advertise_100HdMbps == nil {
		obj.SetAdvertise100HdMbps(true)
	}
	if obj.obj.Advertise_10FdMbps == nil {
		obj.SetAdvertise10FdMbps(true)
	}
	if obj.obj.Advertise_10HdMbps == nil {
		obj.SetAdvertise10HdMbps(true)
	}
	if obj.obj.LinkTraining == nil {
		obj.SetLinkTraining(false)
	}
	if obj.obj.RsFec == nil {
		obj.SetRsFec(false)
	}

}
