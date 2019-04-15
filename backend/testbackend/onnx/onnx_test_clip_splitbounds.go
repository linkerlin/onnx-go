package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

// NewTestClipSplitbounds version: 3.
func NewTestClipSplitbounds() *testbackend.TestCase {
	return &testbackend.TestCase{
		Title:  "TestClipSplitbounds",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x65, 0xa, 0x2a, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x4, 0x43, 0x6c, 0x69, 0x70, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x61, 0x78, 0x15, 0x0, 0x0, 0xa0, 0x40, 0xa0, 0x1, 0x1, 0x2a, 0xd, 0xa, 0x3, 0x6d, 0x69, 0x6e, 0x15, 0x0, 0x0, 0xa0, 0xc0, 0xa0, 0x1, 0x1, 0x12, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x70, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Clip",
		     Attributes: ([]*pb.AttributeProto) (len=2 cap=2) {
		    (*pb.AttributeProto)(0xc00010b000)(name:"max" type:FLOAT f:5 ),
		    (*pb.AttributeProto)(0xc00010b100)(name:"min" type:FLOAT f:-5 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-1, 0, 6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-1, 0, 5}),
			),
		},
	}
}