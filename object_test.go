package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nhat.io/cpy/v3"
)

func TestAttrString(t *testing.T) {
	cpy.Py_Initialize()

	sys := cpy.PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttrString("stdout"))

	stdout := sys.GetAttrString("stdout")

	assert.NotNil(t, stdout)
	assert.Zero(t, sys.DelAttrString("stdout"))
	assert.Nil(t, sys.GetAttrString("stdout"))

	cpy.PyErr_Clear()

	assert.Zero(t, sys.SetAttrString("stdout", stdout))
}

func TestAttr(t *testing.T) {
	cpy.Py_Initialize()

	name := cpy.PyUnicode_FromString("stdout")
	defer name.DecRef()

	sys := cpy.PyImport_ImportModule("sys")
	defer sys.DecRef()

	assert.True(t, sys.HasAttr(name))

	stdout := sys.GetAttr(name)

	assert.NotNil(t, stdout)
	assert.Zero(t, sys.DelAttr(name))
	assert.Nil(t, sys.GetAttr(name))

	cpy.PyErr_Clear()

	assert.Zero(t, sys.SetAttr(name, stdout))
}

func TestRichCompareBool(t *testing.T) {
	cpy.Py_Initialize()

	s1 := cpy.PyUnicode_FromString("test1")
	s2 := cpy.PyUnicode_FromString("test2")

	assert.Zero(t, s1.RichCompareBool(s2, cpy.Py_EQ))
	assert.NotZero(t, s1.RichCompareBool(s1, cpy.Py_EQ))
}

func TestRichCompare(t *testing.T) {
	cpy.Py_Initialize()

	s1 := cpy.PyUnicode_FromString("test1")
	s2 := cpy.PyUnicode_FromString("test2")

	b1 := s1.RichCompare(s2, cpy.Py_EQ)
	defer b1.DecRef()

	assert.Equal(t, cpy.Py_False, b1)

	b2 := s1.RichCompare(s1, cpy.Py_EQ)
	defer b2.DecRef()

	assert.Equal(t, cpy.Py_True, b2)
}

func TestRepr(t *testing.T) {
	cpy.Py_Initialize()

	list := cpy.PyList_New(0)
	defer list.DecRef()

	repr := list.Repr()

	assert.Equal(t, "[]", cpy.PyUnicode_AsUTF8(repr))
}

func TestStr(t *testing.T) {
	cpy.Py_Initialize()

	list := cpy.PyList_New(0)
	defer list.DecRef()

	str := list.Str()

	assert.Equal(t, "[]", cpy.PyUnicode_AsUTF8(str))
}

func TestASCII(t *testing.T) {
	cpy.Py_Initialize()

	list := cpy.PyList_New(0)
	defer list.DecRef()

	ascii := list.ASCII()

	assert.Equal(t, "[]", cpy.PyUnicode_AsUTF8(ascii))
}

func TestCallable(t *testing.T) {
	cpy.Py_Initialize()

	builtins := cpy.PyEval_GetBuiltins()

	assert.True(t, cpy.PyDict_Check(builtins))

	builtinsLength := cpy.PyDict_GetItemString(builtins, "len")

	assert.True(t, cpy.PyCallable_Check(builtinsLength))

	emptyList := cpy.PyList_New(0)

	assert.True(t, cpy.PyList_Check(emptyList))

	args := cpy.PyTuple_New(1)
	defer args.DecRef()

	assert.True(t, cpy.PyTuple_Check(args))

	cpy.PyTuple_SetItem(args, 0, emptyList)

	length := builtinsLength.Call(args, nil)
	defer length.DecRef()

	assert.True(t, cpy.PyLong_Check(length))
	assert.Equal(t, 0, cpy.PyLong_AsLong(length))

	length = builtinsLength.CallObject(args)
	defer length.DecRef()

	assert.True(t, cpy.PyLong_Check(length))
	assert.Equal(t, 0, cpy.PyLong_AsLong(length))

	length = builtinsLength.CallFunctionObjArgs(emptyList)
	defer length.DecRef()

	assert.True(t, cpy.PyLong_Check(length))
	assert.Equal(t, 0, cpy.PyLong_AsLong(length))
}

func TestCallMethod(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_FromString("hello world")
	defer s.DecRef()

	assert.True(t, cpy.PyUnicode_Check(s))

	sep := cpy.PyUnicode_FromString(" ")
	defer sep.DecRef()

	assert.True(t, cpy.PyUnicode_Check(sep))

	split := cpy.PyUnicode_FromString("split")
	defer split.DecRef()

	assert.True(t, cpy.PyUnicode_Check(split))

	words := s.CallMethodObjArgs(split, sep)
	defer words.DecRef()

	assert.True(t, cpy.PyList_Check(words))
	assert.Equal(t, 2, cpy.PyList_Size(words))

	hello := cpy.PyList_GetItem(words, 0)
	world := cpy.PyList_GetItem(words, 1)

	assert.True(t, cpy.PyUnicode_Check(hello))
	assert.True(t, cpy.PyUnicode_Check(world))
	assert.Equal(t, "hello", cpy.PyUnicode_AsUTF8(hello))
	assert.Equal(t, "world", cpy.PyUnicode_AsUTF8(world))

	words = s.CallMethodArgs("split", sep)
	defer words.DecRef()

	assert.True(t, cpy.PyList_Check(words))
	assert.Equal(t, 2, cpy.PyList_Size(words))

	hello = cpy.PyList_GetItem(words, 0)
	world = cpy.PyList_GetItem(words, 1)

	assert.True(t, cpy.PyUnicode_Check(hello))
	assert.True(t, cpy.PyUnicode_Check(world))
	assert.Equal(t, "hello", cpy.PyUnicode_AsUTF8(hello))
	assert.Equal(t, "world", cpy.PyUnicode_AsUTF8(world))
}

func TestIsTrue(t *testing.T) {
	cpy.Py_Initialize()

	b := cpy.Py_True.IsTrue() != 0

	assert.True(t, b)

	b = cpy.Py_False.IsTrue() != 0

	assert.False(t, b)
}

func TestNot(t *testing.T) {
	cpy.Py_Initialize()

	b := cpy.Py_True.Not() != 0

	assert.False(t, b)

	b = cpy.Py_False.Not() != 0

	assert.True(t, b)
}

func TestLength(t *testing.T) {
	cpy.Py_Initialize()

	length := 6

	list := cpy.PyList_New(length)
	defer list.DecRef()

	listLength := list.Length()

	assert.Equal(t, length, listLength)
}

func TestLengthHint(t *testing.T) {
	cpy.Py_Initialize()

	length := 6

	list := cpy.PyList_New(length)
	defer list.DecRef()

	listLength := list.LengthHint(0)

	assert.Equal(t, length, listLength)
}

func TestObjectItem(t *testing.T) {
	cpy.Py_Initialize()

	key := cpy.PyUnicode_FromString("key")
	defer key.DecRef()

	value := cpy.PyUnicode_FromString("value")
	defer value.DecRef()

	dict := cpy.PyDict_New()
	err := dict.SetItem(key, value)

	assert.Zero(t, err)

	dictValue := dict.GetItem(key)

	assert.Equal(t, value, dictValue)

	err = dict.DelItem(key)

	assert.Zero(t, err)
}

func TestDir(t *testing.T) {
	cpy.Py_Initialize()

	list := cpy.PyList_New(0)
	defer list.DecRef()

	dir := list.Dir()
	defer dir.DecRef()

	repr := dir.Repr()
	defer repr.DecRef()

	assert.Equal(t, "['__add__', '__class__', '__class_getitem__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getstate__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']", cpy.PyUnicode_AsUTF8(repr))
}

func TestReprEnterLeave(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_FromString("hello world")
	defer s.DecRef()

	assert.Zero(t, s.ReprEnter())
	assert.Greater(t, s.ReprEnter(), 0)

	s.ReprLeave()
	s.ReprLeave()
}

func TestIsSubclass(t *testing.T) {
	cpy.Py_Initialize()

	assert.Equal(t, 1, cpy.PyExc_Warning.IsSubclass(cpy.PyExc_Exception))
	assert.Equal(t, 0, cpy.Bool.IsSubclass(cpy.Float))
}

func TestHash(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_FromString("test string")
	defer s.DecRef()

	assert.NotEqual(t, -1, s.Hash())
}

func TestObjectType(t *testing.T) {
	cpy.Py_Initialize()

	i := cpy.PyLong_FromGoInt(23543)
	defer i.DecRef()

	assert.Equal(t, cpy.Long, i.Type())
}

func TestHashNotImplemented(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_FromString("test string")
	defer s.DecRef()

	assert.Equal(t, -1, s.HashNotImplemented())
	assert.True(t, cpy.PyErr_ExceptionMatches(cpy.PyExc_TypeError))

	cpy.PyErr_Clear()
}

func TestObjectIter(t *testing.T) {
	cpy.Py_Initialize()

	i := cpy.PyLong_FromGoInt(23)
	defer i.DecRef()

	assert.Nil(t, i.GetIter())
	assert.True(t, cpy.PyErr_ExceptionMatches(cpy.PyExc_TypeError))

	cpy.PyErr_Clear()

	list := cpy.PyList_New(23)
	defer list.DecRef()

	iter := list.GetIter()
	defer iter.DecRef()

	assert.NotNil(t, iter)
}
