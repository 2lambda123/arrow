// Code generated by transpose_ints_s390x.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !noasm

package utils

// if building with the 'noasm' tag, then point to the pure go implementations
var (
	TransposeInt8Int8   = transposeInt8Int8
	TransposeInt8Uint8  = transposeInt8Uint8
	TransposeInt8Int16  = transposeInt8Int16
	TransposeInt8Uint16 = transposeInt8Uint16
	TransposeInt8Int32  = transposeInt8Int32
	TransposeInt8Uint32 = transposeInt8Uint32
	TransposeInt8Int64  = transposeInt8Int64
	TransposeInt8Uint64 = transposeInt8Uint64

	TransposeUint8Int8   = transposeUint8Int8
	TransposeUint8Uint8  = transposeUint8Uint8
	TransposeUint8Int16  = transposeUint8Int16
	TransposeUint8Uint16 = transposeUint8Uint16
	TransposeUint8Int32  = transposeUint8Int32
	TransposeUint8Uint32 = transposeUint8Uint32
	TransposeUint8Int64  = transposeUint8Int64
	TransposeUint8Uint64 = transposeUint8Uint64

	TransposeInt16Int8   = transposeInt16Int8
	TransposeInt16Uint8  = transposeInt16Uint8
	TransposeInt16Int16  = transposeInt16Int16
	TransposeInt16Uint16 = transposeInt16Uint16
	TransposeInt16Int32  = transposeInt16Int32
	TransposeInt16Uint32 = transposeInt16Uint32
	TransposeInt16Int64  = transposeInt16Int64
	TransposeInt16Uint64 = transposeInt16Uint64

	TransposeUint16Int8   = transposeUint16Int8
	TransposeUint16Uint8  = transposeUint16Uint8
	TransposeUint16Int16  = transposeUint16Int16
	TransposeUint16Uint16 = transposeUint16Uint16
	TransposeUint16Int32  = transposeUint16Int32
	TransposeUint16Uint32 = transposeUint16Uint32
	TransposeUint16Int64  = transposeUint16Int64
	TransposeUint16Uint64 = transposeUint16Uint64

	TransposeInt32Int8   = transposeInt32Int8
	TransposeInt32Uint8  = transposeInt32Uint8
	TransposeInt32Int16  = transposeInt32Int16
	TransposeInt32Uint16 = transposeInt32Uint16
	TransposeInt32Int32  = transposeInt32Int32
	TransposeInt32Uint32 = transposeInt32Uint32
	TransposeInt32Int64  = transposeInt32Int64
	TransposeInt32Uint64 = transposeInt32Uint64

	TransposeUint32Int8   = transposeUint32Int8
	TransposeUint32Uint8  = transposeUint32Uint8
	TransposeUint32Int16  = transposeUint32Int16
	TransposeUint32Uint16 = transposeUint32Uint16
	TransposeUint32Int32  = transposeUint32Int32
	TransposeUint32Uint32 = transposeUint32Uint32
	TransposeUint32Int64  = transposeUint32Int64
	TransposeUint32Uint64 = transposeUint32Uint64

	TransposeInt64Int8   = transposeInt64Int8
	TransposeInt64Uint8  = transposeInt64Uint8
	TransposeInt64Int16  = transposeInt64Int16
	TransposeInt64Uint16 = transposeInt64Uint16
	TransposeInt64Int32  = transposeInt64Int32
	TransposeInt64Uint32 = transposeInt64Uint32
	TransposeInt64Int64  = transposeInt64Int64
	TransposeInt64Uint64 = transposeInt64Uint64

	TransposeUint64Int8   = transposeUint64Int8
	TransposeUint64Uint8  = transposeUint64Uint8
	TransposeUint64Int16  = transposeUint64Int16
	TransposeUint64Uint16 = transposeUint64Uint16
	TransposeUint64Int32  = transposeUint64Int32
	TransposeUint64Uint32 = transposeUint64Uint32
	TransposeUint64Int64  = transposeUint64Int64
	TransposeUint64Uint64 = transposeUint64Uint64
)
