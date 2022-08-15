package log

import "go.uber.org/zap"

type Field = zap.Field

// function variables for all field types in github.com/uber-go/zap/field.go
var (
	ZapSkip        = zap.Skip
	ZapBinary      = zap.Binary
	ZapBool        = zap.Bool
	ZapBoolp       = zap.Boolp
	ZapByteString  = zap.ByteString
	ZapComplex128  = zap.Complex128
	ZapComplex128p = zap.Complex128p
	ZapComplex64   = zap.Complex64
	ZapComplex64p  = zap.Complex64p
	ZapFloat64     = zap.Float64
	ZapFloat64p    = zap.Float64p
	ZapFloat32     = zap.Float32
	ZapFloat32p    = zap.Float32p
	ZapInt         = zap.Int
	ZapIntp        = zap.Intp
	ZapInt64       = zap.Int64
	ZapInt64p      = zap.Int64p
	ZapInt32       = zap.Int32
	ZapInt32p      = zap.Int32p
	ZapInt16       = zap.Int16
	ZapInt16p      = zap.Int16p
	ZapInt8        = zap.Int8
	ZapInt8p       = zap.Int8p
	ZapString      = zap.String
	ZapStringp     = zap.Stringp
	ZapUint        = zap.Uint
	ZapUintp       = zap.Uintp
	ZapUint64      = zap.Uint64
	ZapUint64p     = zap.Uint64p
	ZapUint32      = zap.Uint32
	ZapUint32p     = zap.Uint32p
	ZapUint16      = zap.Uint16
	ZapUint16p     = zap.Uint16p
	ZapUint8       = zap.Uint8
	ZapUint8p      = zap.Uint8p
	ZapUintptr     = zap.Uintptr
	ZapUintptrp    = zap.Uintptrp
	ZapReflect     = zap.Reflect
	ZapNamespace   = zap.Namespace
	ZapStringer    = zap.Stringer
	ZapTime        = zap.Time
	ZapTimep       = zap.Timep
	ZapStack       = zap.Stack
	ZapStackSkip   = zap.StackSkip
	ZapDuration    = zap.Duration
	ZapDurationp   = zap.Durationp
	ZapAny         = zap.Any
)
