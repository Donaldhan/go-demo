package reflect

import "testing"

// 单元（功能）测试
// cd reflect
// go test -v
func TestReflect(t *testing.T) {
	Reflect()
}

func TestReflectTypeKind(t *testing.T) {
	ReflectTypeKind()
}

func TestReflectElem(t *testing.T) {
	ReflectElem()
}

func TestReflectField(t *testing.T) {
	ReflectField()
}

func TestReflectStructTagError(t *testing.T) {
	ReflectStructTagError()
}
func TestReflectValue(t *testing.T) {
	ReflectValue()
}
func TestReflectKindX(t *testing.T) {
	ReflectKindX()
}
func TestCanSet(t *testing.T) {
	CanSet()
}
func TestStrutTypeOfT(t *testing.T) {
	StrutTypeOfT()
}
func TestStructElem(t *testing.T) {
	StructElem()
}
func TestInterfaceType(t *testing.T) {
	InterfaceType()
}

func TestNinValid(t *testing.T) {
	NinValid()
}

func TestCanAddrX(t *testing.T) {
	CanAddrX()
}

func TestReflectValueSet(t *testing.T) {
	// ReflectValueSet()
}

func TestReflectValueStructSet(t *testing.T) {
	// ReflectValueStructSet()
}
func TestReflectValueStructSetX(t *testing.T) {
	ReflectValueStructSetX()
}
func TestReflectNew(t *testing.T) {
	ReflectNew()
}

func TestReflectFunctionCall(t *testing.T) {
	ReflectFunctionCall()
}
