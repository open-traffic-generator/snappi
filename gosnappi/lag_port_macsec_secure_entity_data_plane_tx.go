package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LagPortMacsecSecureEntityDataPlaneTx *****
type lagPortMacsecSecureEntityDataPlaneTx struct {
	validation
	obj          *otg.LagPortMacsecSecureEntityDataPlaneTx
	marshaller   marshalLagPortMacsecSecureEntityDataPlaneTx
	unMarshaller unMarshalLagPortMacsecSecureEntityDataPlaneTx
}

func NewLagPortMacsecSecureEntityDataPlaneTx() LagPortMacsecSecureEntityDataPlaneTx {
	obj := lagPortMacsecSecureEntityDataPlaneTx{obj: &otg.LagPortMacsecSecureEntityDataPlaneTx{}}
	obj.setDefault()
	return &obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) msg() *otg.LagPortMacsecSecureEntityDataPlaneTx {
	return obj.obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) setMsg(msg *otg.LagPortMacsecSecureEntityDataPlaneTx) LagPortMacsecSecureEntityDataPlaneTx {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallagPortMacsecSecureEntityDataPlaneTx struct {
	obj *lagPortMacsecSecureEntityDataPlaneTx
}

type marshalLagPortMacsecSecureEntityDataPlaneTx interface {
	// ToProto marshals LagPortMacsecSecureEntityDataPlaneTx to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneTx
	ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneTx, error)
	// ToPbText marshals LagPortMacsecSecureEntityDataPlaneTx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LagPortMacsecSecureEntityDataPlaneTx to YAML text
	ToYaml() (string, error)
	// ToJson marshals LagPortMacsecSecureEntityDataPlaneTx to JSON text
	ToJson() (string, error)
}

type unMarshallagPortMacsecSecureEntityDataPlaneTx struct {
	obj *lagPortMacsecSecureEntityDataPlaneTx
}

type unMarshalLagPortMacsecSecureEntityDataPlaneTx interface {
	// FromProto unmarshals LagPortMacsecSecureEntityDataPlaneTx from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneTx
	FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneTx) (LagPortMacsecSecureEntityDataPlaneTx, error)
	// FromPbText unmarshals LagPortMacsecSecureEntityDataPlaneTx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LagPortMacsecSecureEntityDataPlaneTx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LagPortMacsecSecureEntityDataPlaneTx from JSON text
	FromJson(value string) error
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) Marshal() marshalLagPortMacsecSecureEntityDataPlaneTx {
	if obj.marshaller == nil {
		obj.marshaller = &marshallagPortMacsecSecureEntityDataPlaneTx{obj: obj}
	}
	return obj.marshaller
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneTx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallagPortMacsecSecureEntityDataPlaneTx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallagPortMacsecSecureEntityDataPlaneTx) ToProto() (*otg.LagPortMacsecSecureEntityDataPlaneTx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallagPortMacsecSecureEntityDataPlaneTx) FromProto(msg *otg.LagPortMacsecSecureEntityDataPlaneTx) (LagPortMacsecSecureEntityDataPlaneTx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallagPortMacsecSecureEntityDataPlaneTx) ToPbText() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneTx) FromPbText(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneTx) ToYaml() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneTx) FromYaml(value string) error {
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

func (m *marshallagPortMacsecSecureEntityDataPlaneTx) ToJson() (string, error) {
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

func (m *unMarshallagPortMacsecSecureEntityDataPlaneTx) FromJson(value string) error {
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

func (obj *lagPortMacsecSecureEntityDataPlaneTx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) Clone() (LagPortMacsecSecureEntityDataPlaneTx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLagPortMacsecSecureEntityDataPlaneTx()
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

// LagPortMacsecSecureEntityDataPlaneTx is a container of Tx properties of SecY.
type LagPortMacsecSecureEntityDataPlaneTx interface {
	Validation
	// msg marshals LagPortMacsecSecureEntityDataPlaneTx to protobuf object *otg.LagPortMacsecSecureEntityDataPlaneTx
	// and doesn't set defaults
	msg() *otg.LagPortMacsecSecureEntityDataPlaneTx
	// setMsg unmarshals LagPortMacsecSecureEntityDataPlaneTx from protobuf object *otg.LagPortMacsecSecureEntityDataPlaneTx
	// and doesn't set defaults
	setMsg(*otg.LagPortMacsecSecureEntityDataPlaneTx) LagPortMacsecSecureEntityDataPlaneTx
	// provides marshal interface
	Marshal() marshalLagPortMacsecSecureEntityDataPlaneTx
	// provides unmarshal interface
	Unmarshal() unMarshalLagPortMacsecSecureEntityDataPlaneTx
	// validate validates LagPortMacsecSecureEntityDataPlaneTx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LagPortMacsecSecureEntityDataPlaneTx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EndStation returns bool, set in LagPortMacsecSecureEntityDataPlaneTx.
	EndStation() bool
	// SetEndStation assigns bool provided by user to LagPortMacsecSecureEntityDataPlaneTx
	SetEndStation(value bool) LagPortMacsecSecureEntityDataPlaneTx
	// HasEndStation checks if EndStation has been set in LagPortMacsecSecureEntityDataPlaneTx
	HasEndStation() bool
	// IncludeSci returns bool, set in LagPortMacsecSecureEntityDataPlaneTx.
	IncludeSci() bool
	// SetIncludeSci assigns bool provided by user to LagPortMacsecSecureEntityDataPlaneTx
	SetIncludeSci(value bool) LagPortMacsecSecureEntityDataPlaneTx
	// HasIncludeSci checks if IncludeSci has been set in LagPortMacsecSecureEntityDataPlaneTx
	HasIncludeSci() bool
}

// End station on not.
// EndStation returns a bool
func (obj *lagPortMacsecSecureEntityDataPlaneTx) EndStation() bool {

	return *obj.obj.EndStation

}

// End station on not.
// EndStation returns a bool
func (obj *lagPortMacsecSecureEntityDataPlaneTx) HasEndStation() bool {
	return obj.obj.EndStation != nil
}

// End station on not.
// SetEndStation sets the bool value in the LagPortMacsecSecureEntityDataPlaneTx object
func (obj *lagPortMacsecSecureEntityDataPlaneTx) SetEndStation(value bool) LagPortMacsecSecureEntityDataPlaneTx {

	obj.obj.EndStation = &value
	return obj
}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *lagPortMacsecSecureEntityDataPlaneTx) IncludeSci() bool {

	return *obj.obj.IncludeSci

}

// Include SCI on not.
// IncludeSci returns a bool
func (obj *lagPortMacsecSecureEntityDataPlaneTx) HasIncludeSci() bool {
	return obj.obj.IncludeSci != nil
}

// Include SCI on not.
// SetIncludeSci sets the bool value in the LagPortMacsecSecureEntityDataPlaneTx object
func (obj *lagPortMacsecSecureEntityDataPlaneTx) SetIncludeSci(value bool) LagPortMacsecSecureEntityDataPlaneTx {

	obj.obj.IncludeSci = &value
	return obj
}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lagPortMacsecSecureEntityDataPlaneTx) setDefault() {
	if obj.obj.EndStation == nil {
		obj.SetEndStation(false)
	}
	if obj.obj.IncludeSci == nil {
		obj.SetIncludeSci(false)
	}

}
