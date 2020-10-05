package godottemplate

import (
	"github.com/godot-go/godot-go/pkg/gdnative"
	"github.com/godot-go/godot-go/pkg/log"
)

var (
	signalDozen gdnative.String
	defaultName gdnative.Variant
)

type MyCounter struct {
	gdnative.NodeImpl
	gdnative.UserDataIdentifiableImpl

	name  gdnative.String
	count int64
}

func (s MyCounter) ClassName() string {
	return "MyCounter"
}

func (s MyCounter) BaseClass() string {
	return "Node"
}

func (s *MyCounter) Init() {

}

func (c *MyCounter) OnClassRegistered(e gdnative.ClassRegisteredEvent) {
	// methods
	e.RegisterMethod("_ready", "Ready")
	e.RegisterMethod("send", "Send")

	// signals
	e.RegisterSignal("dozens", gdnative.RegisterSignalArg{"count", gdnative.GODOT_VARIANT_TYPE_INT})

	// properties
	e.RegisterProperty("name", "SetName", "GetName", defaultName)
	e.RegisterProperty("count", "SetCount", "GetCount", gdnative.NewVariantInt(0))
}

func (c *MyCounter) Free() {
	log.WithFields(gdnative.WithObject(c.GetOwnerObject())).Trace("free MyCounter")
}

func MycounterCreateFunc(owner *gdnative.GodotObject, typeTag gdnative.TypeTag) gdnative.NativeScriptClass {
	log.WithFields(gdnative.WithObject(owner)).Info("create_func new MyCounter")

	m := &MyCounter{}
	m.Owner = owner
	m.TypeTag = typeTag

	return m
}

func NewSlack() MyCounter {
	log.Trace("NewMyCounter")
	return *(gdnative.CreateCustomClassInstance("MyCounter", "Node").(*MyCounter))
}

func (c MyCounter) MyCounter() {
	log.WithFields(gdnative.WithObject(c.GetOwnerObject())).Trace("Start MyCounter::Ready")
	log.WithFields(gdnative.WithObject(c.GetOwnerObject())).Trace("End MyCounter::Ready")
}

func MyCounterNativescriptInit() {
	signalDozen = gdnative.NewStringFromGoString("dozen")
	defaultName = gdnative.NewVariantString(gdnative.NewStringFromGoString("godot"))
}

func MyCounterNativescriptTerminate() {
	signalDozen.Destroy()
	defaultName.Destroy()
}

// Properties

func (c *MyCounter) SetName(v gdnative.Variant) {
	c.name = v.AsString()
}

func (c MyCounter) GetName() gdnative.Variant {
	return gdnative.NewVariantString(c.name)
}

func (c *MyCounter) SetCount(v gdnative.Variant) {
	c.count = v.AsInt()
}

func (c MyCounter) GetCount() gdnative.Variant {
	return gdnative.NewVariantInt(c.count)
}

// Method
func (c *MyCounter) Increment() {
	c.count++
	if c.count%12 == 0 {
		variant := gdnative.NewVariantInt(c.count)
		defer variant.Destroy()
		c.EmitSignal(signalDozen, &variant)
	}
}
