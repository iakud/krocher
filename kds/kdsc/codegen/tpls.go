package codegen

// Protobuf
const TemplateProtobuf = `
{{- /* BEGIN DEFINE */ -}}
{{- /* ENUM */ -}}
{{- define "Enum"}}
enum {{.Name}} {
{{- range .EnumFields}}
	{{.Name}} = {{.Value}};
{{- end}}
}
{{- end}}
{{- /* MESSAGE */ -}}
{{- define "Message"}}
message {{.Name}} {
{{- range .Fields}}
{{- if .Repeated}}
	repeated {{protoType .Type}} {{.Name}} = {{.Number}};
{{- else if .Map}}
	map<{{protoType .KeyType}}, {{protoType .Type}}> {{.Name}} = {{.Number}};
{{- else}}
	{{protoType .Type}} {{.Name}} = {{.Number}};
{{- end}}
{{- end}}
}
{{- end}}
{{- /* ENTITY */ -}}
{{- define "Entity"}}
{{- template "Message" .}}
{{- end}}
{{- /* COMPONENT */ -}}
{{- define "Component"}}
{{- template "Message" .}}
{{- end}}
{{- /* END DEFINE */ -}}
{{- /* BEGIN PROTOBUF */ -}}
// Code generated by kds. DO NOT EDIT.
// source: {{.SourceFile}}

syntax = "proto3";

package {{.Package}};

{{- if len .Imports}}
{{/* EMPTY LINE */}}
{{- end}}
{{- range .Imports}}
import "{{.}}.proto";
{{- end}}

{{- if len .ProtoImports}}
{{/* EMPTY LINE */}}
{{- range .ProtoImports}}
import "{{.}}";
{{- end}}
{{- end}}

option go_package="{{.ProtoGoPackage}}";

{{- range .Defs}}
{{- if eq .Kind "enum"}}
{{/* EMPTY LINE */}}
{{- template "Enum" .}}
{{- else if eq .Kind "entity"}}
{{/* EMPTY LINE */}}
{{- template "Entity" .}}
{{- else if eq .Kind "component"}}
{{/* EMPTY LINE */}}
{{- template "Component" .}}
{{- end}}
{{- end}}
{{- /* END PROTOBUF */ -}}
`

// Kds Go File
const TemplateKdsGo = `
{{- /* BEGIN DEFINE */ -}}
{{- /* FIELD TYPE */ -}}
{{- define "FieldType"}}
{{- if eq .TypeKind "enum"}}{{.Type}}
{{- else if eq .TypeKind "component"}}*{{.Type}}
{{- else}}{{goType .Type}}
{{- end}}
{{- end}}
{{- /* LIST TYPE */ -}}
{{- define "ListValueType"}}
{{- if eq .TypeKind "enum"}}{{.Type}}
{{- else if eq .TypeKind "component"}}*{{.Type}}
{{- else}}{{goType .Type}}
{{- end}}
{{- end}}
{{- /* MAP TYPE */ -}}
{{- define "MapValueType"}}
{{- if eq .TypeKind "enum"}}{{.Type}}
{{- else if eq .TypeKind "component"}}*{{.Type}}
{{- else}}{{goType .Type}}
{{- end}}
{{- end}}
{{- /* LIST PROTO TYPE */ -}}
{{- define "ListValueProtoType"}}
{{- if eq .TypeKind "enum"}}{{goProtoPackage .Type}}.{{.Type}}
{{- else if eq .TypeKind "component"}}*{{goProtoPackage .Type}}.{{.Type}}
{{- else if eq .Type "timestamp"}}*{{goProtoType .Type}}
{{- else if eq .Type "duration"}}*{{goProtoType .Type}}
{{- else if eq .Type "empty"}}*{{goProtoType .Type}}
{{- else}}{{goProtoType .Type}}
{{- end}}
{{- end}}
{{- /* MAP PROTO TYPE */ -}}
{{- define "MapValueProtoType"}}
{{- if eq .TypeKind "enum"}}{{goProtoPackage .Type}}.{{.Type}}
{{- else if eq .TypeKind "component"}}*{{goProtoPackage .Type}}.{{.Type}}
{{- else if eq .Type "timestamp"}}*{{goProtoType .Type}}
{{- else if eq .Type "duration"}}*{{goProtoType .Type}}
{{- else if eq .Type "empty"}}*{{goProtoType .Type}}
{{- else}}{{goProtoType .Type}}
{{- end}}
{{- end}}

{{- /* LIST */ -}}
{{- define "List"}}
type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{.Name}} struct {
	syncable []{{template "ListValueType" .}}

	dirty bool
	dirtyParent dirtyParentFunc_{{.Name}}
}

func (x *{{.Name}}) Len() int {
	return len(x.syncable)
}

func (x *{{.Name}}) Get(i int) {{template "ListValueType" .}} {
	return x.syncable[i]
}

func (x *{{.Name}}) Set(i int, v {{template "ListValueType" .}}) {
{{- if eq .TypeKind "component"}}
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed from its original place first")
	}
{{- end}}
	if v == x.syncable[i] {
		return
	}
{{- if eq .TypeKind "component"}}
	if x.syncable[i] != nil {
		x.syncable[i].dirtyParent = nil
	}
{{- end}}
	x.syncable[i] = v
{{- if eq .TypeKind "component"}}
	if v != nil {
		v.dirtyParent = func() {
			x.markDirty()
		}
		v.dirty |= uint64(0x01)
	}
{{- end}}
	x.markDirty()
}

func (x *{{.Name}}) Append(v ...{{template "ListValueType" .}}) {
{{- if eq .TypeKind "component"}}
	for i := range v {
		if v[i] != nil && v[i].dirtyParent != nil {
			panic("the component should be removed from its original place first")
		}
	}
{{- end}}
	if len(v) == 0 {
		return
	}
	x.syncable = append(x.syncable, v...)
{{- if eq .TypeKind "component"}}
	for i := range v {
		if v[i] != nil {
			v[i].dirtyParent = func() {
				x.markDirty()
			}
			v[i].dirty |= uint64(0x01)
		}
	}
{{- end}}
	x.markDirty()
}

func (x *{{.Name}}) Insert(i int, v ...{{template "ListValueType" .}}) {
{{- if eq .TypeKind "component"}}
	for j := range v {
		if v[j] != nil && v[j].dirtyParent != nil {
			panic("the component should be removed from its original place first")
		}
	}
{{- end}}
	if len(v) == 0 {
		return
	}
	x.syncable = slices.Insert(x.syncable, i, v...)
{{- if eq .TypeKind "component"}}
	for j := range v {
		if v[j] != nil {
			v[j].dirtyParent = func() {
				x.markDirty()
			}
			v[j].dirty |= uint64(0x01)
		}
	}
{{- end}}
	x.markDirty()
}

func (x *{{.Name}}) Delete(i, j int) {
{{- if eq .TypeKind "component"}}
	r := x.syncable[i:j:len(x.syncable)]
	for k := range r {
		if r[k] != nil {
			r[k].dirtyParent = nil
		}
	}
{{- end}}
	if i == j {
		return
	}
	x.syncable = slices.Delete(x.syncable, i, j)
	x.markDirty()
}

func (x *{{.Name}}) Replace(i, j int, v ...{{template "ListValueType" .}}) {
{{- if eq .TypeKind "component"}}
	for k := range v {
		if v[k] != nil && v[k].dirtyParent != nil {
			panic("the component should be removed from its original place first")
		}
	}
	r := x.syncable[i:j:len(x.syncable)]
	for k := range r {
		if r[k] != nil {
			r[k].dirtyParent = nil
		}
	}
{{- end}}
	if i == j && len(v) == 0 {
		return
	}
	x.syncable = slices.Replace(x.syncable, i, j, v...)
{{- if eq .TypeKind "component"}}
	for k := range v {
		if v[k] != nil {
			v[k].dirtyParent = func() {
				x.markDirty()
			}
			v[k].dirty |= uint64(0x01)
		}
	}
{{- end}}
	x.markDirty()
}

func (x *{{.Name}}) Reverse() {
	if len(x.syncable) < 2 {
		return
	}
	slices.Reverse(x.syncable)
	x.markDirty()
}

func (x *{{.Name}}) All() iter.Seq2[int, {{template "ListValueType" .}}] {
	return slices.All(x.syncable)
}

func (x *{{.Name}}) Backward() iter.Seq2[int, {{template "ListValueType" .}}] {
	return slices.Backward(x.syncable)
}

func (x *{{.Name}}) Values() iter.Seq[{{template "ListValueType" .}}] {
	return slices.Values(x.syncable)
}

func (x *{{.Name}}) DumpChange() []{{template "ListValueProtoType" .}} {
	return x.DumpFull()
}

func (x *{{.Name}}) DumpFull() []{{template "ListValueProtoType" .}} {
	var m []{{template "ListValueProtoType" .}}
{{- if eq .TypeKind "component"}}
	for _, v := range x.syncable {
		m = append(m, v.DumpChange())
	}
{{- else if eq .TypeKind "enum"}}
	for _, v := range x.syncable {
		m = append(m, v)
	}
{{- else if eq .Type "timestamp"}}
	for _, v := range x.syncable {
		m = append(m, timestamppb.New(v))
	}
{{- else if eq .Type "duration"}}
	for _, v := range x.syncable {
		m = append(m, durationpb.New(v))
	}
{{- else if eq .Type "empty"}}
	for range x.syncable {
		m = append(m, new(emptypb.Empty))
	}
{{- else}}
	for _, v := range x.syncable {
		m = append(m, v)
	}
{{- end}}
	return m
}

func (x *{{.Name}}) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.invoke()
}

func (x *{{.Name}}) clearDirty() {
{{- if eq .TypeKind "component"}}
	for k := range x.syncable {
		if x.syncable[k] != nil {
			x.syncable[k].clearDirty()
		}
	}
{{- end}}
	x.dirty = false
}
{{- end}}

{{- /* MAP */ -}}
{{- define "Map"}}
type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{.Name}} struct {
	syncable map[{{goType .KeyType}}]{{template "MapValueType" .}}

	update map[{{goType .KeyType}}]{{template "MapValueType" .}}
	deleteKey map[{{goType .KeyType}}]struct{}
	clear bool
	dirty bool
	dirtyParent dirtyParentFunc_{{.Name}}
}

func (x *{{.Name}}) Len() int {
	return len(x.syncable)
}

func (x *{{.Name}}) Clear() {
	if len(x.syncable) == 0 && len(x.deleteKey) == 0 {
		return
	}
{{- if eq .TypeKind "component"}}
	for _, v := range x.syncable {
		if v != nil {
			v.dirtyParent = nil
		}
	}
{{- end}}
	clear(x.syncable)
	clear(x.update)
	clear(x.deleteKey)
	x.clear = true
	x.markDirty()
}

func (x *{{.Name}}) Get(k {{goType .KeyType}}) ({{template "MapValueType" .}}, bool) {
	v, ok := x.syncable[k]
	return v, ok
}

func (x *{{.Name}}) Set(k {{goType .KeyType}}, v {{template "MapValueType" .}}) {
{{- if eq .TypeKind "component"}}
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed from its original place first")
	}
{{- end}}
	if e, ok := x.syncable[k]; ok {
		if e == v {
			return
		}
{{- if eq .TypeKind "component"}}
		if e != nil {
			e.dirtyParent = nil
		}
{{- end}}
	}
	x.syncable[k] = v
{{- if eq .TypeKind "component"}}
	if v != nil {
		v.dirtyParent = func() {
			if _, ok := x.update[k]; ok {
				return
			}
			x.update[k] = v
			x.markDirty()
		}
		v.dirty |= uint64(0x01)
	}
{{- end}}
	x.update[k] = v
	delete(x.deleteKey, k)
	x.markDirty()
}

func (x *{{.Name}}) Delete(k {{goType .KeyType}}) {
{{- if eq .TypeKind "component"}}
	if v, ok := x.syncable[k]; !ok {
		return
	} else if v != nil {
		v.dirtyParent = nil
	}
{{- else}}
	if _, ok := x.syncable[k]; !ok {
		return
	}
{{- end}}
	delete(x.syncable, k)
	x.deleteKey[k] = struct{}{}
	delete(x.update, k)
	x.markDirty()
}

func (x *{{.Name}}) All() iter.Seq2[{{goType .KeyType}}, {{template "MapValueType" .}}] {
	return maps.All(x.syncable)
}

func (x *{{.Name}}) Keys() iter.Seq[{{goType .KeyType}}] {
	return maps.Keys(x.syncable)
}

func (x *{{.Name}}) Values() iter.Seq[{{template "MapValueType" .}}] {
	return maps.Values(x.syncable)
}

func (x *{{.Name}}) DumpChange() map[{{goType .KeyType}}]{{template "MapValueProtoType" .}} {
	if x.clear {
		return x.DumpFull()
	}
	m := make(map[{{goType .KeyType}}]{{template "MapValueProtoType" .}})
{{- if eq .TypeKind "component"}}
	for k, v := range x.update {
		m[k] = v.DumpFull()
	}
{{- else if eq .TypeKind "enum"}}
	for k, v := range x.update {
		m[k] = v
	}
{{- else if eq .Type "timestamp"}}
	for k, v := range x.update {
		m[k] = timestamppb.New(v)
	}
{{- else if eq .Type "duration"}}
	for k, v := range x.update {
		m[k] = durationpb.New(v)
	}
{{- else if eq .Type "empty"}}
	for k := range x.update {
		m[k] = new(emptypb.Empty)
	}
{{- else}}
	for k, v := range x.update {
		m[k] = v
	}
{{- end}}
	for k, _ := range x.deleteKey {
		_ = k // deleteKeys
	}
	return m
}

func (x *{{.Name}}) DumpFull() map[{{goType .KeyType}}]{{template "MapValueProtoType" .}} {
	m := make(map[{{goType .KeyType}}]{{template "MapValueProtoType" .}})
{{- if eq .TypeKind "component"}}
	for k, v := range x.syncable {
		m[k] = v.DumpFull()
	}
{{- else if eq .TypeKind "enum"}}
	for k, v := range x.syncable {
		m[k] = v
	}
{{- else if eq .Type "timestamp"}}
	for k, v := range x.syncable {
		m[k] = timestamppb.New(v)
	}
{{- else if eq .Type "duration"}}
	for k, v := range x.syncable {
		m[k] = durationpb.New(v)
	}
{{- else if eq .Type "empty"}}
	for k := range x.syncable {
		m[k] = new(emptypb.Empty)
	}
{{- else}}
	for k, v := range x.syncable {
		m[k] = v
	}
{{- end}}
	return m
}

func (x *{{.Name}}) markDirty() {
	if x.dirty {
		return
	}
	x.dirty = true
	x.dirtyParent.invoke()
}

func (x *{{.Name}}) clearDirty() {
	if !x.dirty {
		return
	}
{{- if eq .TypeKind "component"}}
	for _, v := range x.update {
		if v != nil {
			v.clearDirty()
		}
	}
{{- end}}
	clear(x.update)
	clear(x.deleteKey)
	x.clear = false
	x.dirty = false
}
{{- end}}

{{- /* ENUM */ -}}
{{- define "Enum"}}
{{- $EnumType := .Name}}
type {{.Name}} = {{.GoProtoPackage}}.{{.Name}}

const (
{{- range .EnumFields}}
	{{$EnumType}}_{{.Name}} {{$EnumType}} = {{.Value}}
{{- end}}
)
{{- end}}

{{- /* SYNCABLE */ -}}
{{- define "Syncable"}}
type syncable{{.Name}} struct {
{{- range .Fields}}
{{- if not (or .Repeated .Map (eq .TypeKind "component"))}}
	{{.Name}} {{template "FieldType" .}}
{{- end}}
{{- end}}
}
{{- end}}

{{- /* MESSAGE */ -}}
{{- define "Message"}}
{{- $MessageName := .Name}}

{{- range .Fields}}
{{/* EMPTY LINE */}}
{{- if .Repeated}}
func (x *{{$MessageName}}) Get{{.Name}}() *{{.ListType}} {
	return &x.syncable{{.Name}}
}

func (x *{{$MessageName}}) init{{.Name}}() {
	x.syncable{{.Name}}.dirtyParent = func() {
		x.markDirty(uint64(0x01) << {{.Number}})
	}
}
{{- else if .Map}}
func (x *{{$MessageName}}) Get{{.Name}}() *{{.MapType}} {
	return &x.syncable{{.Name}}
}

func (x *{{$MessageName}}) init{{.Name}}() {
	x.syncable{{.Name}}.syncable = make(map[{{goType .KeyType}}]{{template "FieldType" .}})
	x.syncable{{.Name}}.update = make(map[{{goType .KeyType}}]{{template "FieldType" .}})
	x.syncable{{.Name}}.deleteKey = make(map[{{goType .KeyType}}]struct{})
	x.syncable{{.Name}}.dirtyParent = func() {
		x.markDirty(uint64(0x01) << {{.Number}})
	}
}
{{- else if eq .TypeKind "component"}}
func (x *{{$MessageName}}) Get{{.Name}}() {{template "FieldType" .}} {
	return x.syncable{{.Name}}
}

func (x *{{$MessageName}}) set{{.Name}}(v *{{.Type}}) {
	if v != nil && v.dirtyParent != nil {
		panic("the component should be removed from its original place first")
	}
	if v == x.syncable{{.Name}} {
		return
	}
	if x.syncable{{.Name}} != nil {
		x.syncable{{.Name}}.dirtyParent = nil
	}
	x.syncable{{.Name}} = v
	v.dirtyParent = func() {
		x.markDirty(uint64(0x01) << {{.Number}})
	}
	x.markDirty(uint64(0x01) << {{.Number}})
	if v != nil {
		v.markDirty(uint64(0x01))
	}
}
{{- else}}
func (x *{{$MessageName}}) Get{{.Name}}() {{template "FieldType" .}} {
	return x.syncable.{{.Name}}
}

func (x *{{$MessageName}}) Set{{.Name}}(v {{template "FieldType" .}}) {
{{- if eq .Type "bytes"}}
	if v != nil || x.syncable.{{.Name}} != nil {
		return
	}
{{- else}}
	if v == x.syncable.{{.Name}} {
		return
	}
{{- end}}
	x.syncable.{{.Name}} = v
	x.markDirty(uint64(0x01) << {{.Number}})
}
{{- end}}
{{- end}}

func (x *{{$MessageName}}) DumpChange() *{{.GoProtoPackage}}.{{.Name}} {
	if x.checkDirty(uint64(0x01)) {
		return x.DumpFull()
	}
	m := new({{.GoProtoPackage}}.{{.Name}})
{{- range .Fields}}
	if x.checkDirty(uint64(0x01) << {{.Number}}) {
{{- if .Repeated}}
		m.{{.Name}} = x.syncable{{.Name}}.DumpChange()
{{- else if .Map}}
		m.{{.Name}} = x.syncable{{.Name}}.DumpChange()
{{- else if eq .TypeKind "component"}}
		m.{{.Name}} = x.syncable{{.Name}}.DumpChange()
{{- else if eq .Type "timestamp"}}
		m.{{.Name}} = timestamppb.New(x.syncable.{{.Name}})
{{- else if eq .Type "duration"}}
		m.{{.Name}} = durationpb.New(x.syncable.{{.Name}})
{{- else if eq .Type "empty"}}
		m.{{.Name}} = new(emptypb.Empty)
{{- else}}
		m.{{.Name}} = x.syncable.{{.Name}}
{{- end}}
	}
{{- end}}
	return m
}

func (x *{{$MessageName}}) DumpFull() *{{.GoProtoPackage}}.{{.Name}} {
	m := new({{.GoProtoPackage}}.{{.Name}})
{{- range .Fields}}
{{- if .Repeated}}
	m.{{.Name}} = x.syncable{{.Name}}.DumpFull()
{{- else if .Map}}
	m.{{.Name}} = x.syncable{{.Name}}.DumpFull()
{{- else if eq .TypeKind "component"}}
	m.{{.Name}} = x.syncable{{.Name}}.DumpFull()
{{- else if eq .Type "timestamp"}}
	m.{{.Name}} = timestamppb.New(x.syncable.{{.Name}})
{{- else if eq .Type "duration"}}
	m.{{.Name}} = durationpb.New(x.syncable.{{.Name}})
{{- else if eq .Type "empty"}}
	m.{{.Name}} = new(emptypb.Empty)
{{- else}}
	m.{{.Name}} = x.syncable.{{.Name}}
{{- end}}
{{- end}}
	return m
}

{{- end}}

{{- /* ENTITY */ -}}
{{- define "Entity"}}
{{- template "Syncable" .}}

type {{.Name}} struct {
	id int64
	syncable syncable{{.Name}}
{{- range .Fields}}
{{- if .Repeated}}
	syncable{{.Name}} {{.ListType}}
{{- else if .Map}}
	syncable{{.Name}} {{.MapType}}
{{- else if eq .TypeKind "component"}}
	syncable{{.Name}} *{{.Type}}
{{- end}}
{{- end}}

	dirty uint64
}

func New{{.Name}}() *{{.Name}} {
	x := new({{.Name}})
	x.dirty = 1
	x.id = 0 // FIXME: gen nextId()
{{- range .Fields}}
{{- if .Repeated}}
	x.init{{.Name}}()
{{- else if .Map}}
	x.init{{.Name}}()
{{- else if eq .TypeKind "component"}}
	x.set{{.Name}}(New{{.Type}}())
{{- end}}
{{- end}}
	return x
}

func (x *{{.Name}}) Id() int64 {
	return x.id
}
{{- template "Message" .}}

func (x *{{.Name}}) markAll() {
	x.dirty = uint64(0x01)
}

func (x *{{.Name}}) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
}

func (x *{{.Name}}) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}

func (x *{{.Name}}) clearDirty() {
	if x.dirty == 0 {
		return
	}
{{- range .Fields}}
{{- if .Repeated}}
	if x.dirty & uint64(0x01) != 0 || x.dirty & uint64(0x01) << {{.Number}} != 0 {
		x.syncable{{.Name}}.clearDirty()
	}
{{- else if .Map}}
	if x.dirty & uint64(0x01) != 0 || x.dirty & uint64(0x01) << {{.Number}} != 0 {
		x.syncable{{.Name}}.clearDirty()
	}
{{- else if eq .TypeKind "component"}}
	if x.dirty & uint64(0x01) != 0 || x.dirty & uint64(0x01) << {{.Number}} != 0 {
		x.syncable{{.Name}}.clearDirty()
	}
{{- end}}
{{- end}}
	x.dirty = 0
}
{{- end}}

{{- /* COMPONENT */ -}}
{{- define "Component"}}
{{- template "Syncable" .}}

type dirtyParentFunc_{{.Name}} func()

func (f dirtyParentFunc_{{.Name}}) invoke() {
	if f == nil {
		return
	}
	f()
}

type {{.Name}} struct {
	syncable syncable{{.Name}}
{{- range .Fields}}
{{- if .Repeated}}
	syncable{{.Name}} {{.ListType}}
{{- else if .Map}}
	syncable{{.Name}} {{.MapType}}
{{- else if eq .TypeKind "component"}}
	syncable{{.Name}} *{{.Type}}
{{- end}}
{{- end}}

	dirty uint64
	dirtyParent dirtyParentFunc_{{.Name}}
}

func New{{.Name}}() *{{.Name}} {
	x := new({{.Name}})
	x.dirty = 1
{{- range .Fields}}
{{- if .Repeated}}
	x.init{{.Name}}()
{{- else if .Map}}
	x.init{{.Name}}()
{{- else if eq .TypeKind "component"}}
	x.set{{.Name}}(New{{.Type}}())
{{- end}}
{{- end}}
	return x
}
{{- template "Message" .}}

func (x *{{.Name}}) markAll() {
	x.dirty = uint64(0x01)
}

func (x *{{.Name}}) markDirty(n uint64) {
	if x.dirty & n == n {
		return
	}
	x.dirty |= n
	x.dirtyParent.invoke()
}

func (x *{{.Name}}) checkDirty(n uint64) bool {
	return x.dirty & n != 0
}

func (x *{{.Name}}) clearDirty() {
	if x.dirty == 0 {
		return
	}
{{- range .Fields}}
{{- if .Repeated}}
	if x.dirty & uint64(0x01) != 0 || x.dirty & uint64(0x01) << {{.Number}} != 0 {
		x.syncable{{.Name}}.clearDirty()
	}
{{- else if .Map}}
	if x.dirty & uint64(0x01) != 0 || x.dirty & uint64(0x01) << {{.Number}} != 0 {
		x.syncable{{.Name}}.clearDirty()
	}
{{- else if eq .TypeKind "component"}}
	if x.dirty & uint64(0x01) != 0 || x.dirty & uint64(0x01) << {{.Number}} != 0 {
		x.syncable{{.Name}}.clearDirty()
	}
{{- end}}
{{- end}}
	x.dirty = 0
}
{{- end}}
{{- /* END DEFINE */ -}}
{{- /* BEGIN GO */ -}}
// Code generated by kds. DO NOT EDIT.
// source: {{.SourceFile}}

package {{.Package}};
{{- /* IMPORTS */ -}}
{{- if len .GoImportSpecs}}
{{/* EMPTY LINE */}}
import (
{{- range .GoImportSpecs}}
{{- if .SpacesBefore}}
{{/* EMPTY LINE */}}
{{- end}}
	"{{.Path}}"
{{- end}}
)
{{- end}}
{{- /* COMMON TYPES */ -}}
{{- range .Types}}

{{- with findList .}}
{{- if .}}
{{/* EMPTY LINE */}}
{{- template "List" .}}
{{- end}}
{{- end}}
{{- range findMap .}}
{{/* EMPTY LINE */}}
{{- template "Map" .}}
{{- end}}

{{- end}}
{{- /* TOPLEVEL DEFINE */ -}}
{{- range .Defs}}
{{/* EMPTY LINE */}}
{{- if eq .Kind "enum"}}
{{- template "Enum" .}}
{{- with findList .Name}}
{{- if .}}
{{/* EMPTY LINE */}}
{{- template "List" .}}
{{- end}}
{{- end}}
{{- range findMap .Name}}
{{/* EMPTY LINE */}}
{{- template "Map" .}}
{{- end}}

{{- else if eq .Kind "entity"}}
{{- template "Entity" .}}

{{- else if eq .Kind "component"}}
{{- template "Component" .}}
{{- with findList .Name}}
{{- if .}}
{{/* EMPTY LINE */}}
{{- template "List" .}}
{{- end}}
{{- end}}
{{- range findMap .Name}}
{{/* EMPTY LINE */}}
{{- template "Map" .}}
{{- end}}
{{- end}}

{{- end}}
{{- /* END GO */ -}}
`
